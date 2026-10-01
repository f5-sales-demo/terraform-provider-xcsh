---
page_title: "xcsh_service_policy_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule reference."
---

# xcsh_service_policy_rule reference

<a id="canonical-b7f017ed53fdb22da43c016e91fa7027665306b15864a7ba0c0ded1d1f40bbb4"></a>

## regex_values property — jwt_claims.item / 1ea83c790ef2 / 5

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

<a id="canonical-a14020c707d8861ddb93d22eef6fe3dbf9c9c23207352a9ddef77ed0a3575a0a"></a>

<a id="canonical-a7f8fa4c4c0fe85ba864b7ce244e332d5e84000fbd29379333b9c5dd1e5285eb"></a>

## transformers property — jwt_claims.item / 1ea83c790ef2 / 6

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

<a id="canonical-f0fe29c4a8ed6655971d63e20c77403184ecc7f30fe8bb9451ff3b988334350f"></a>

## Next pages — jwt_claims.item / 1ea83c790ef2 / 7

- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-d9913c042ec992e4e948ffc9f73a49c0bbaa7e16b2cbe0bee34c0b7e863df397)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-885654158fb78e99195654bfd7eb4ede70c6b02f2d395cef3f137f6410b798c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db53ba3017d5860654155bb5c5c87a87758a85354ccfeac837b7e77723785f1b"></a>

## label_matcher — label_matcher / 44d2866523ae / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- label_matcher

<a id="canonical-760e05075b5d983c33bbc79bdce61f39bf38bce4d424ccc0d72e2925b2dbebf0"></a>

Type: `"object"`. single nested block, Optional.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Receipt-pinned upstream constraints:

```json
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
label_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-b0f50f13b1966072b81796b180c37740e7e834a24fbc8c993a9c0294ca67d666"></a>

## Direct properties — label_matcher / 44d2866523ae / 3

<a id="canonical-b316ea3d9da36d4f504fbc815ee8fc222baff5f65396203b051b00fb3660af06"></a>

<a id="canonical-d202dbaa7c8ab5ad186adbd026ae7af40a6f7f86f61fdfb990249ac1d1a808ec"></a>

## keys property — label_matcher / 44d2866523ae / 4

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-dabc1dd4aec8fafc610ea44fc2f4c771c43643a5b9a6de3fbd251e78ee7f958e"></a>

## Next pages — label_matcher / 44d2866523ae / 5

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-8173a7b12b6dbcabe15c5a24896b2850104c373bdd419858980760f784308f9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28852064096f6813418633b5cb1c8d288a8d5681e9640b07e9ae878675013789"></a>

## mum_action — mum_action / 47a02e44baf0 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- mum_action

<a id="canonical-a96746381ab87215d5dc2b130643029aaf328e776d49cadf3ba42382ce035150"></a>

Type: `"object"`. single nested block, Optional.

Modify behavior for a matching request. The modification could be to entirely skip processing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default",
    "skip_processing")}
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
  "x-ves-oneof-field-action_type": "[\"default\",\"skip_processing\"]"
}
```

Terraform syntax:

```terraform
mum_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-3e18dfd9368ae67228d23e3737cf94e000efe16857add75a76dd800a4ddc42c7"></a>

## Direct properties — mum_action / 47a02e44baf0 / 3

- [default](resources--service_policy_rule--reference--group-002.md#canonical-e5481ace354f83dfdc6a2d223d5c50e7e1638cd0b90e09f26061be7bd664d852): complete subsection reference.

- [skip_processing](resources--service_policy_rule--reference--group-002.md#canonical-1aefe79674859d7e00276ddcb378cc2053c7f3b70b9f596cf04d3a2aee9c543f): complete subsection reference.

<a id="canonical-cb52527d065e1e1cf03f94e367f7515187c1c85339c3247538dd23c0839b3bb9"></a>

## Next pages — mum_action / 47a02e44baf0 / 4

- [mum_action.default](resources--service_policy_rule--reference--group-002.md#canonical-e5481ace354f83dfdc6a2d223d5c50e7e1638cd0b90e09f26061be7bd664d852)
- [mum_action.skip_processing](resources--service_policy_rule--reference--group-002.md#canonical-1aefe79674859d7e00276ddcb378cc2053c7f3b70b9f596cf04d3a2aee9c543f)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-e5481ace354f83dfdc6a2d223d5c50e7e1638cd0b90e09f26061be7bd664d852"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fa1918f31cc81c603c27f80c7a90f4eed0ef3e1f99ba175dd0b9bf88f80c54a"></a>

## mum_action.default — mum_action.default / b3b01b995a56 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [mum_action](resources--service_policy_rule--reference--group-002.md#canonical-8173a7b12b6dbcabe15c5a24896b2850104c373bdd419858980760f784308f9d)
- mum_action.default

<a id="canonical-925d59f1cc3b8c5a11b829e63a969d20519140c09d18fa424aa595c74d4248f2"></a>

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
default = {}
```

<a id="canonical-0fc989295fdd4ad49d15ae72a8f152cc54b912d04bfe92715eb592d08d82f5cc"></a>

## Direct properties — mum_action.default / b3b01b995a56 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3489b426b2bf32e698552dd812a369ac8ecdb8e125f4ef3e3f9c1bdcb72f62c3"></a>

## Next pages — mum_action.default / b3b01b995a56 / 4

- [mum_action](resources--service_policy_rule--reference--group-002.md#canonical-8173a7b12b6dbcabe15c5a24896b2850104c373bdd419858980760f784308f9d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-1aefe79674859d7e00276ddcb378cc2053c7f3b70b9f596cf04d3a2aee9c543f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0e14570f8d2a8f46382088302e1d0b3e52ade8f6c93abb3b934a39d9d9befad"></a>

## mum_action.skip_processing — mum_action.skip_processing / 20bdfad0ea23 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [mum_action](resources--service_policy_rule--reference--group-002.md#canonical-8173a7b12b6dbcabe15c5a24896b2850104c373bdd419858980760f784308f9d)
- mum_action.skip_processing

<a id="canonical-36d25b7c63255352d3bc21986471f3267a2305ddf5633971fc4b1dc9272159f5"></a>

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
skip_processing = {}
```

<a id="canonical-da01dfbb575c28d5650163c4e47bc69f18b0d9b5a1d755218d70b4e9ec08aa7d"></a>

## Direct properties — mum_action.skip_processing / 20bdfad0ea23 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fdfbc21192ba7405247305d71cb57c13486a5b829c3eb1c07a28338c56778b53"></a>

## Next pages — mum_action.skip_processing / 20bdfad0ea23 / 4

- [mum_action](resources--service_policy_rule--reference--group-002.md#canonical-8173a7b12b6dbcabe15c5a24896b2850104c373bdd419858980760f784308f9d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-c453af0b534656a15fed7afe77bb9ad4115339a74c815501323d7c06337a777c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c9c26d75016e10e7fc8ef361a8a42c6b22c8a1cd109661da1e8593a6d1b0d2f"></a>

## path — path / 3debc46a3632 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- path

<a id="canonical-6af93d47ba88d484b4fe57b19fcde6ddd13a83e9beed16060baa2c6f10a54bf0"></a>

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

<a id="canonical-ee5e9deda1be5d28e3c94a9580113ea8bd4334aa9bd0c29dd89764a992a0e9c9"></a>

## Direct properties — path / 3debc46a3632 / 3

<a id="canonical-63e41612ec6d0152f469426f443a8eb3798b555191ba3debcd5dba431bf57d97"></a>

<a id="canonical-d262c86319ec7099df916b63c393823e0d01cc339d6627ac7961e8874ae78976"></a>

## encoded_path_matcher property — path / 3debc46a3632 / 4

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

<a id="canonical-f094b2c510f0220e29ba80438b7378d7545cb2522353ecda3178cb472e96508a"></a>

<a id="canonical-df504e2c7a6a7d1d8e34b343522582e8ddf03bdef43cf097c3a2920ad5f3273b"></a>

## exact_values property — path / 3debc46a3632 / 5

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

<a id="canonical-1a893eeefedc018f0d0ce860739d00295a90e224e9e8602f2c8c907083197f14"></a>

<a id="canonical-f89a3a32c22faf1695c72690d9510607786b5a3d10811ad184c88c8e2994182a"></a>

## invert_matcher property — path / 3debc46a3632 / 6

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

<a id="canonical-0f76def79fccdab7fca45f761e2a30acefc2dd4c6e0fc2bc439de16e8f1b159c"></a>

<a id="canonical-3de79dadb1555f3097c36eb4825b78bf8c3f55c46a80dad9655f71c81ad97380"></a>

## prefix_values property — path / 3debc46a3632 / 7

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

<a id="canonical-ed8dbf45442704dfd0134b7f960a0e52a7c45d879953722ef190e129a3ec4d79"></a>

<a id="canonical-5948a989319978af48a2577d465ce867f469619c40be5854a6ce8c3a01d9682d"></a>

## regex_values property — path / 3debc46a3632 / 8

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

<a id="canonical-06f8ab0d215fd1a6293c348b8c851f70ba8b526e528b34930afaa402db9d2001"></a>

<a id="canonical-395a0fe11ac42b47e17074496dd118fc66cb5cc22d18080d3d93b99284eb6ec4"></a>

## suffix_values property — path / 3debc46a3632 / 9

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

<a id="canonical-bea7491b7392e0c8459bc85a382fa3f93ec48d536b1f8c822711eccc6695d52c"></a>

<a id="canonical-caf332d2b68bf0b0cb410c140f9220191781ae28d65831e57b1ef2ff8098a065"></a>

## transformers property — path / 3debc46a3632 / 10

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

<a id="canonical-ec2c6313db8cf88b5acbe1b6edb3dfd06f4f260bf60aa0a9134e39c564df6faf"></a>

## Next pages — path / 3debc46a3632 / 11

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-a3aa60eda2c10f207494ee2277014d7f3dde42d696ff2dcc542c02648fd6d21f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e731ce1f4fa8dad948db86ebcd16f698676f5640507cf9e14a2e29a5714531d"></a>

## port_matcher — port_matcher / 300e9df60eb4 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- port_matcher

<a id="canonical-ba598e2ad70c6507d4b690fa29610e739e52ad5f8c95cd922db927fa8b2ed825"></a>

Type: `"object"`. single nested block, Optional.

Port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true. Server applies default when omitted.

Upstream description:

A port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
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
port_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2ecdfa66fe76b81799f4350edd3e5b0d618a06f69d68abba53c145805d5fd28d"></a>

## Direct properties — port_matcher / 300e9df60eb4 / 3

<a id="canonical-18a750d040b0da907aec0d8e4d73e96d6454710ce418d272d391201ecee54b09"></a>

<a id="canonical-2c5d052219d404747869d502078911f97480234fe7503fd7815dc646b65cf419"></a>

## invert_matcher property — port_matcher / 300e9df60eb4 / 4

Type: `"bool"`. Optional.

Invert Port Matcher. Invert the match result.

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

<a id="canonical-abe394e30b6a6ec1c2ac8e09e788842c6d72a7821ecee80cdd2c11d18466bf73"></a>

<a id="canonical-4fd411b324ae9971bc5bfa1f0db7cea41f6108d9715934ed58f14f1c80970fa2"></a>

## ports property — port_matcher / 300e9df60eb4 / 5

Type: `["list", "string"]`. Optional.

List of strings, each of which is a single port value or a tuple of start and end port values
separated by '-'. The start and end values are considered to be part of the range.

Upstream description:

A list of strings, each of which is a single port value or a tuple of start and end port values
separated by "-". The start and end values are considered to be part of the range.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-cc5119517f6e4da9b9cc2567634aab53ea0493e8bf63b18650cd7017846fbc5b"></a>

## Next pages — port_matcher / 300e9df60eb4 / 6

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-4639c0528507de8e532608ba3a369a541f88419059d914ce8a0f99a6366e9111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-718300c9b3f609699b7201a2618154fb57ca96939600ed7c4f7f2aba3113f591"></a>

## query_params — query_params / 1adcc5fb638c / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- query_params

<a id="canonical-5e673c0aea73dd35c6a73e2a30cfa90986748ad923cadef6d7f5933f0558e13a"></a>

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

<a id="canonical-65525c4d13564489a13c8045d3704e18f10b71e03a11b231a246342cadcd09be"></a>

## Direct properties — query_params / 1adcc5fb638c / 3

- [check_not_present](resources--service_policy_rule--reference--group-002.md#canonical-af404d7907c084fcae6c0b9cb7ff0eee8965a4c6f962050eaa1fa1d5d5ef8152): complete subsection reference.

- [check_present](resources--service_policy_rule--reference--group-002.md#canonical-a78e879a5bd57d1820799eaa7852016de79f0433906a43ed58e546804125e640): complete subsection reference.

<a id="canonical-8f44b135345aac5290bb604550df17e35d29b00d129ec54173a3c63026c05d62"></a>

<a id="canonical-39ca1966be300c8e017c05f3586d371349103e5f745f67a45c5647b292ac1ee9"></a>

## invert_matcher property — query_params / 1adcc5fb638c / 4

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

- [item](resources--service_policy_rule--reference--group-002.md#canonical-75f64a686b468e85311ba78b3f57b19f05ddc6e0d2b7cdc2baa3a44962ef77c6): complete subsection reference.

<a id="canonical-90c647f4bfaf6c61dd033d1d7fafc991999fd9998dab93ff22731df6ed017df6"></a>

<a id="canonical-7035e97b5e6d755f9b7d099f5d3473f6c71cbf0bd3a6bf1202a07903e5a7a171"></a>

## key property — query_params / 1adcc5fb638c / 5

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

<a id="canonical-f48f10cda396aec29269cb4b087387b1bbe23fe75aa257794e0a49a8c27e3e80"></a>

## Next pages — query_params / 1adcc5fb638c / 6

- [query_params.check_not_present](resources--service_policy_rule--reference--group-002.md#canonical-af404d7907c084fcae6c0b9cb7ff0eee8965a4c6f962050eaa1fa1d5d5ef8152)
- [query_params.check_present](resources--service_policy_rule--reference--group-002.md#canonical-a78e879a5bd57d1820799eaa7852016de79f0433906a43ed58e546804125e640)
- [query_params.item](resources--service_policy_rule--reference--group-002.md#canonical-75f64a686b468e85311ba78b3f57b19f05ddc6e0d2b7cdc2baa3a44962ef77c6)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-af404d7907c084fcae6c0b9cb7ff0eee8965a4c6f962050eaa1fa1d5d5ef8152"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c258476e9688b90fb863738dbbf93e88d12af79191240216b4afb1f086343aa"></a>

## query_params.check_not_present — query_params.check_not_present / fb86d74bf061 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [query_params](resources--service_policy_rule--reference--group-002.md#canonical-4639c0528507de8e532608ba3a369a541f88419059d914ce8a0f99a6366e9111)
- query_params.check_not_present

<a id="canonical-e2f016bd5cf28071a1626f2980fcda1e3d1fff9f0c16060d2ecae66660a7f882"></a>

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

<a id="canonical-7e4b9afb39f9d10dd0798dcdce23c816e3e26c6e4a2f1e661f3936dd98c488e9"></a>

## Direct properties — query_params.check_not_present / fb86d74bf061 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-59b7d23b6e93b797e25555ebdf8acb796ef966743f7115b018027da573686f18"></a>

## Next pages — query_params.check_not_present / fb86d74bf061 / 4

- [query_params](resources--service_policy_rule--reference--group-002.md#canonical-4639c0528507de8e532608ba3a369a541f88419059d914ce8a0f99a6366e9111)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-a78e879a5bd57d1820799eaa7852016de79f0433906a43ed58e546804125e640"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e32ecd11121d591d7e22575912f7e387be71dc653766645b3bb380d7aaa459f7"></a>

## query_params.check_present — query_params.check_present / 461d7910c036 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [query_params](resources--service_policy_rule--reference--group-002.md#canonical-4639c0528507de8e532608ba3a369a541f88419059d914ce8a0f99a6366e9111)
- query_params.check_present

<a id="canonical-2336808eb9e1aef93c5949fc245bf3b5b00d5c515facb34e533806b6db7f45d9"></a>

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

<a id="canonical-3f7cbaf0ff43fa3b2c7428620ad855db41c77958271e202e871313163a5b66fd"></a>

## Direct properties — query_params.check_present / 461d7910c036 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f0f430d3b15f33ce54576fabad0fa84235226ae75f0e369550cd240793163232"></a>

## Next pages — query_params.check_present / 461d7910c036 / 4

- [query_params](resources--service_policy_rule--reference--group-002.md#canonical-4639c0528507de8e532608ba3a369a541f88419059d914ce8a0f99a6366e9111)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-75f64a686b468e85311ba78b3f57b19f05ddc6e0d2b7cdc2baa3a44962ef77c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-465163530872cc02532d119d2b1c0f6317a000063c4cd02e77f9e880b8d3ea5f"></a>

## query_params.item — query_params.item / e213190c9b9e / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [query_params](resources--service_policy_rule--reference--group-002.md#canonical-4639c0528507de8e532608ba3a369a541f88419059d914ce8a0f99a6366e9111)
- query_params.item

<a id="canonical-af4ee6679d50bdfa377f8c985903450f95cdc5fc1e780cb19124e0e7de3718f6"></a>

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

<a id="canonical-e12a6825ea8b12b4c82fd52dc66e1681d36443c6d8912c9a5a008c1ca1316684"></a>

## Direct properties — query_params.item / e213190c9b9e / 3

<a id="canonical-fb56437d5f6d7654dde992b75ac941a5059c7b71b101eb27fed90c7397d63b96"></a>

<a id="canonical-b3466c4fd3ded4c01ff2780596d82c38ffa5213fb0ec7cd86a0ae6db3820605f"></a>

## exact_values property — query_params.item / e213190c9b9e / 4

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

<a id="canonical-43d0e053feb8a6a24d97ab0dcfb6b14ce1e4fc7dc21dc729c0b5ac849b777949"></a>

<a id="canonical-e6cb4edac796ea29d24d0a14a013f3062e6d0ad88cf4c58fdaa0a25dc94c35e7"></a>

## regex_values property — query_params.item / e213190c9b9e / 5

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

<a id="canonical-212e604288705df5ac7c5a948e7adc1173b355e8f8ea985e7e321ee9acc26a31"></a>

<a id="canonical-16a17ea81da5a142151eb4ec6a67d68ed9b6363564d8f9946e1627f0054bfcd6"></a>

## transformers property — query_params.item / e213190c9b9e / 6

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

<a id="canonical-beab798df8038537730b8a4555c98728170ae3cd138c4d3dff32fc18340a5460"></a>

## Next pages — query_params.item / e213190c9b9e / 7

- [query_params](resources--service_policy_rule--reference--group-002.md#canonical-4639c0528507de8e532608ba3a369a541f88419059d914ce8a0f99a6366e9111)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84df9b84608a7f639c034516b33200fed47c51e25584ded60f3466b8c433cd7b"></a>

## request_constraints — request_constraints / c7923af9bdab / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- request_constraints

<a id="canonical-f9c2b05098e790f0f50c4e03b2a23eb2ffa42d922bc0c62423baf521f93257f0"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request constraints.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_cookie_count_exceeds",
    "max_cookie_count_none"),
  validators.ConflictingObjectAttributes("max_cookie_key_size_exceeds",
    "max_cookie_key_size_none"),
  validators.ConflictingObjectAttributes("max_cookie_value_size_exceeds",
    "max_cookie_value_size_none"),
  validators.ConflictingObjectAttributes("max_header_count_exceeds",
    "max_header_count_none"),
  validators.ConflictingObjectAttributes("max_header_key_size_exceeds",
    "max_header_key_size_none"),
  validators.ConflictingObjectAttributes("max_header_value_size_exceeds",
    "max_header_value_size_none"),
  validators.ConflictingObjectAttributes("max_parameter_count_exceeds",
    "max_parameter_count_none"),
  validators.ConflictingObjectAttributes("max_parameter_name_size_exceeds",
    "max_parameter_name_size_none"),
  validators.ConflictingObjectAttributes("max_parameter_value_size_exceeds",
    "max_parameter_value_size_none"),
  validators.ConflictingObjectAttributes("max_query_size_exceeds",
    "max_query_size_none"),
  validators.ConflictingObjectAttributes("max_request_line_size_exceeds",
    "max_request_line_size_none"),
  validators.ConflictingObjectAttributes("max_request_size_exceeds",
    "max_request_size_none"),
  validators.ConflictingObjectAttributes("max_url_size_exceeds",
    "max_url_size_none")}
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
  "x-ves-oneof-field-max_cookie_count_choice": "[\"max_cookie_count_exceeds\",\"max_cookie_count_none\"]",
  "x-ves-oneof-field-max_cookie_key_size_choice": "[\"max_cookie_key_size_exceeds\",\"max_cookie_key_size_none\"]",
  "x-ves-oneof-field-max_cookie_value_size_choice": "[\"max_cookie_value_size_exceeds\",\"max_cookie_value_size_none\"]",
  "x-ves-oneof-field-max_header_count_choice": "[\"max_header_count_exceeds\",\"max_header_count_none\"]",
  "x-ves-oneof-field-max_header_key_size_choice": "[\"max_header_key_size_exceeds\",\"max_header_key_size_none\"]",
  "x-ves-oneof-field-max_header_value_size_choice": "[\"max_header_value_size_exceeds\",\"max_header_value_size_none\"]",
  "x-ves-oneof-field-max_parameter_count_choice": "[\"max_parameter_count_exceeds\",\"max_parameter_count_none\"]",
  "x-ves-oneof-field-max_parameter_name_size_choice": "[\"max_parameter_name_size_exceeds\",\"max_parameter_name_size_none\"]",
  "x-ves-oneof-field-max_parameter_value_size_choice": "[\"max_parameter_value_size_exceeds\",\"max_parameter_value_size_none\"]",
  "x-ves-oneof-field-max_query_size_choice": "[\"max_query_size_exceeds\",\"max_query_size_none\"]",
  "x-ves-oneof-field-max_request_line_size_choice": "[\"max_request_line_size_exceeds\",\"max_request_line_size_none\"]",
  "x-ves-oneof-field-max_request_size_choice": "[\"max_request_size_exceeds\",\"max_request_size_none\"]",
  "x-ves-oneof-field-max_url_size_choice": "[\"max_url_size_exceeds\",\"max_url_size_none\"]"
}
```

Terraform syntax:

```terraform
request_constraints {
  # Configure direct properties listed below.
}
```

<a id="canonical-9edb5bb25fcc7623b6e67aad5bad50f046a4d7580f59e3d9ccceb517706dc1ab"></a>

## Direct properties — request_constraints / c7923af9bdab / 3

<a id="canonical-b7834970258e25db1418339e952bcfe485b8adba324bb364af27776cbdc239f0"></a>

<a id="canonical-d04bb8ff69fbf8fc502f6903cfb8a4623395fabab324243ed055dc809c7f0647"></a>

## max_cookie_count_exceeds property — request_constraints / c7923af9bdab / 4

Type: `"number"`. Optional.

Match on the Count for all Cookies that exceed this value. Exclusive with
\[max\_cookie\_count\_none\]

Upstream description:

Exclusive with \[max\_cookie\_count\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_cookie_count_none](resources--service_policy_rule--reference--group-002.md#canonical-5a28501eb4b984f572ea0c2cad7c819b115e3a5cd3f5111c203a4a1abf83c2b4): complete subsection reference.

<a id="canonical-88991a529d4d9b0ee5458c074c39b76f81972f77131815f597a9f074db4203d0"></a>

<a id="canonical-b5cc98ee1e043dc487fd4efd8073d60b9c10b7b4c87172f432c6b356aa20a00f"></a>

## max_cookie_key_size_exceeds property — request_constraints / c7923af9bdab / 5

Type: `"number"`. Optional.

Exclusive with \[max\_cookie\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_key\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_cookie_key_size_none](resources--service_policy_rule--reference--group-002.md#canonical-e4d7cde44c3831eeb41be5db209cd154ce3644bd9d9e845df9865567d56f06da): complete subsection reference.

<a id="canonical-4934ef3754d74a2b5caed9a85b114bd708fa066058402ab6b9fc29f41df5e0a9"></a>

<a id="canonical-fd0af916f2e4cb1afaae135205997fba6826d7be739d1e3efc981b80e7ba5dfe"></a>

## max_cookie_value_size_exceeds property — request_constraints / c7923af9bdab / 6

Type: `"number"`. Optional.

Exclusive with \[max\_cookie\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_value\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 32768),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

- [max_cookie_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-a4fbc952fbc89a183f205fdcedbe9f08351112cd67624c9cce2396f3e81cf038): complete subsection reference.

<a id="canonical-eb15bcde201e57f87f7670bbc519158d2d13cd55585c940366d20b6e123bba33"></a>

<a id="canonical-5e35dc2158ca4bcec2d75085728291ec5053f782107f5e87dd8cc00a9a50118d"></a>

## max_header_count_exceeds property — request_constraints / c7923af9bdab / 7

Type: `"number"`. Optional.

Match on the Count for all Headers that exceed this value. Exclusive with
\[max\_header\_count\_none\]

Upstream description:

Exclusive with \[max\_header\_count\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "40"
  }
}
```

- [max_header_count_none](resources--service_policy_rule--reference--group-002.md#canonical-8dfc81add07b8fe87909596659e96e174f5b1695e93bf066fe0147fd4d1dadc0): complete subsection reference.

<a id="canonical-10a116c4d617b294577631c1fa31b80c95c894779f8775b6336ade57576c0956"></a>

<a id="canonical-aad2b30ce242080349534bfd5a9b84182f79d9a8fb1d92101d5d93c9f1593fd5"></a>

## max_header_key_size_exceeds property — request_constraints / c7923af9bdab / 8

Type: `"number"`. Optional.

Exclusive with \[max\_header\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_key\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_header_key_size_none](resources--service_policy_rule--reference--group-002.md#canonical-5e449dd2fe3f57c3a47c103a04bbd8ceba7d1179623b8fb6de3c63b7fe502f54): complete subsection reference.

<a id="canonical-8acda1972bf9a58cabdb3f5a58e8d09b3ac4606f7d92edd7428a32aa3056d147"></a>

<a id="canonical-94782b1b9f92f882344c97c64f81109418bafccc7cd8617d6c193316c80547f2"></a>

## max_header_value_size_exceeds property — request_constraints / c7923af9bdab / 9

Type: `"number"`. Optional.

Exclusive with \[max\_header\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_value\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 64000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [max_header_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-9262c92d904535185b8c5841cdfe5ebc64b5517cb448f052789c51f50961efd3): complete subsection reference.

<a id="canonical-6f3238576fa36484c7fe387a68d844ecea249a8d924fddae267fee4309160780"></a>

<a id="canonical-73635f8d96a756109d8e94b2a6b893fe733cca168fa04cc00f4bd25f5fd99421"></a>

## max_parameter_count_exceeds property — request_constraints / c7923af9bdab / 10

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_count\_none\].

Upstream description:

Exclusive with \[max\_parameter\_count\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_parameter_count_none](resources--service_policy_rule--reference--group-002.md#canonical-e9c316ff2f641a4975c83005848d33c5d0767c4a05bbb6a830b57df477d15fee): complete subsection reference.

<a id="canonical-08fbe2fa9b079c2f14fecab50e37a502524564d64fe31a4f7c3a26a8728da845"></a>

<a id="canonical-5ef4db500d8e3c8345a75602eff1ad78038cda22de857320e24ee7636b37e62d"></a>

## max_parameter_name_size_exceeds property — request_constraints / c7923af9bdab / 11

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_name\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_name\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_parameter_name_size_none](resources--service_policy_rule--reference--group-002.md#canonical-ac89f75ebd0b44b9a3fac6f6d84a174d80a9750f9a4a7501f613f90e0e9c556a): complete subsection reference.

<a id="canonical-0ce1c05cffb71f89b89d255fbed9d105b830841ebf24eae9e17acab772f1156c"></a>

<a id="canonical-060511b79157ceb581218253924bdd63aaddac6e9b0a1ffd44f3bac84e2f30d6"></a>

## max_parameter_value_size_exceeds property — request_constraints / c7923af9bdab / 12

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_value\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1073741824),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1073741824,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1073741824"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1073741824"
  }
}
```

- [max_parameter_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-4b652bacd47947068316b1bb3c753e8a6fd4c430176f4b86ef43103de00bc5a0): complete subsection reference.

<a id="canonical-eb244083c53ecfd9debecc5e1f0742e9fd4395cc6f65591f82df9d22fcfbd07e"></a>

<a id="canonical-af563bc2d71bedff631977b0f9debe91fc8e16efa06ca4315fb2b948ea5687ea"></a>

## max_query_size_exceeds property — request_constraints / c7923af9bdab / 13

Type: `"number"`. Optional.

Match on the URL Query Size that exceed this value. Exclusive with \[max\_query\_size\_none\]

Upstream description:

Exclusive with \[max\_query\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 60000),
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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

- [max_query_size_none](resources--service_policy_rule--reference--group-002.md#canonical-104c3764a83230dc671d74469bdde202e568752c2222ef65d5c22e5a71f52720): complete subsection reference.

<a id="canonical-9528912dcc8fad9d1fa851cfac4689f2bb5a04379a63f7cdc8ffd4eadf1b113e"></a>

<a id="canonical-1234d0d275fb7c33b6ebe9b9cd7d2dd7fae1d134880686ebca721317006f8198"></a>

## max_request_line_size_exceeds property — request_constraints / c7923af9bdab / 14

Type: `"number"`. Optional.

Exclusive with \[max\_request\_line\_size\_none\].

Upstream description:

Exclusive with \[max\_request\_line\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_line_size_none](resources--service_policy_rule--reference--group-002.md#canonical-001fb5cd8f1dfe0e84be4f6fa20814bfeaec670f9e29b6452de0a3a71f1c201c): complete subsection reference.

<a id="canonical-c50de14a1cd202da88934803d55dc1c440aa4cc7b2abd9857b6d8b99b3444b03"></a>

<a id="canonical-992924279a5a42ce454b56cda4a8db97c4c89d35f74f8b0698dd1ef860deca1e"></a>

## max_request_size_exceeds property — request_constraints / c7923af9bdab / 15

Type: `"number"`. Optional.

Match on the Request Size that exceed this value. Exclusive with \[max\_request\_size\_none\]

Upstream description:

Exclusive with \[max\_request\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_size_none](resources--service_policy_rule--reference--group-002.md#canonical-2526596e2d713a0ba5c148327c6d808e6ff5d6023d74410e410fa211ae3623a0): complete subsection reference.

<a id="canonical-c2f9271e4b4a722ff1e77173dcd75c8965369e21f61dc0d8c122cd2a1544c64a"></a>

<a id="canonical-4e5d9bd42b4661b9c3eace8384e41447f2e2a9ee50edd1ee440d90dc9d3476d8"></a>

## max_url_size_exceeds property — request_constraints / c7923af9bdab / 16

Type: `"number"`. Optional.

Match on the URL Size that exceed this value. Exclusive with \[max\_url\_size\_none\]

Upstream description:

Exclusive with \[max\_url\_size\_none\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 128000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128000"
  }
}
```

- [max_url_size_none](resources--service_policy_rule--reference--group-002.md#canonical-bcf7aa40a79a14c558b10975e11785685b1694164272c1792754b65df37072fd): complete subsection reference.

<a id="canonical-807de24fd1b1bccb78f57be6631b1f11efa6c1259f4af42122daa7a37d4127c6"></a>

## Next pages — request_constraints / c7923af9bdab / 17

- [request_constraints.max_cookie_count_none](resources--service_policy_rule--reference--group-002.md#canonical-5a28501eb4b984f572ea0c2cad7c819b115e3a5cd3f5111c203a4a1abf83c2b4)
- [request_constraints.max_cookie_key_size_none](resources--service_policy_rule--reference--group-002.md#canonical-e4d7cde44c3831eeb41be5db209cd154ce3644bd9d9e845df9865567d56f06da)
- [request_constraints.max_cookie_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-a4fbc952fbc89a183f205fdcedbe9f08351112cd67624c9cce2396f3e81cf038)
- [request_constraints.max_header_count_none](resources--service_policy_rule--reference--group-002.md#canonical-8dfc81add07b8fe87909596659e96e174f5b1695e93bf066fe0147fd4d1dadc0)
- [request_constraints.max_header_key_size_none](resources--service_policy_rule--reference--group-002.md#canonical-5e449dd2fe3f57c3a47c103a04bbd8ceba7d1179623b8fb6de3c63b7fe502f54)
- [request_constraints.max_header_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-9262c92d904535185b8c5841cdfe5ebc64b5517cb448f052789c51f50961efd3)
- [request_constraints.max_parameter_count_none](resources--service_policy_rule--reference--group-002.md#canonical-e9c316ff2f641a4975c83005848d33c5d0767c4a05bbb6a830b57df477d15fee)
- [request_constraints.max_parameter_name_size_none](resources--service_policy_rule--reference--group-002.md#canonical-ac89f75ebd0b44b9a3fac6f6d84a174d80a9750f9a4a7501f613f90e0e9c556a)
- [request_constraints.max_parameter_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-4b652bacd47947068316b1bb3c753e8a6fd4c430176f4b86ef43103de00bc5a0)
- [request_constraints.max_query_size_none](resources--service_policy_rule--reference--group-002.md#canonical-104c3764a83230dc671d74469bdde202e568752c2222ef65d5c22e5a71f52720)
- [request_constraints.max_request_line_size_none](resources--service_policy_rule--reference--group-002.md#canonical-001fb5cd8f1dfe0e84be4f6fa20814bfeaec670f9e29b6452de0a3a71f1c201c)
- [request_constraints.max_request_size_none](resources--service_policy_rule--reference--group-002.md#canonical-2526596e2d713a0ba5c148327c6d808e6ff5d6023d74410e410fa211ae3623a0)
- [request_constraints.max_url_size_none](resources--service_policy_rule--reference--group-002.md#canonical-bcf7aa40a79a14c558b10975e11785685b1694164272c1792754b65df37072fd)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-5a28501eb4b984f572ea0c2cad7c819b115e3a5cd3f5111c203a4a1abf83c2b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26417203e1d26f130712b12b98d63c30a3dbafa1743df83d7bfb4108f482fcdb"></a>

## request_constraints.max_cookie_count_none — request_constraints.max_cookie_count_none / c989bef9baa1 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- request_constraints.max_cookie_count_none

<a id="canonical-fc0fa0b3d6153df4b47f004e31597090adc7cd0a57d316d565c2f3212df3d255"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie count none.

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
max_cookie_count_none = {}
```

<a id="canonical-d38838dfd5e9e18de99336ac4844630ca2732679e4e2bd38af0c943d059f45c9"></a>

## Direct properties — request_constraints.max_cookie_count_none / c989bef9baa1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-43c7a8d2c8cb21cbdf112cde9976575e2638acb390ad479ea75605a30bc36146"></a>

## Next pages — request_constraints.max_cookie_count_none / c989bef9baa1 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-e4d7cde44c3831eeb41be5db209cd154ce3644bd9d9e845df9865567d56f06da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b612e52e5b1eb4a876e8e6751de5ae9960bd660c39b4428f54a4b87362c941a0"></a>

## request_constraints.max_cookie_key_size_none — request_constraints.max_cookie_key_size_none / cf2e4ceb9c38 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- request_constraints.max_cookie_key_size_none

<a id="canonical-1a78462c48b944d6e919b2e171462ce231d03d17e4cf7e7137f7bb6f7b0219f8"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie key size none.

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
max_cookie_key_size_none = {}
```

<a id="canonical-8d1b2ab414c3a3187fd75a2ac86e9dbdb7304376f22fa907060c7135ebe95eac"></a>

## Direct properties — request_constraints.max_cookie_key_size_none / cf2e4ceb9c38 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0c74e35cd3c62026757a6485d09dd09287cbe5b27d1ac9b4ac058ec5db98512a"></a>

## Next pages — request_constraints.max_cookie_key_size_none / cf2e4ceb9c38 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-a4fbc952fbc89a183f205fdcedbe9f08351112cd67624c9cce2396f3e81cf038"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c65a90fa89cdcf07908353b445e403c7459907a13a7b30eddc40bedb8522bead"></a>

## request_constraints.max_cookie_value_size_none — request_constraints.max_cookie_value_size_none / 647294fd6621 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- request_constraints.max_cookie_value_size_none

<a id="canonical-1c8b538dc92a8d67a2542c27d67da59cdd43885e6c2c34e28ac5da995c3e2f92"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie value size none.

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
max_cookie_value_size_none = {}
```

<a id="canonical-f7412782a4869134845b7fd76aea8bd93f4b9379701b1c5d5535c3a4e1a16057"></a>

## Direct properties — request_constraints.max_cookie_value_size_none / 647294fd6621 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f24c52e315220f58338c54fa95d2f4d375f36593e54b14ba965921ca3413986a"></a>

## Next pages — request_constraints.max_cookie_value_size_none / 647294fd6621 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-8dfc81add07b8fe87909596659e96e174f5b1695e93bf066fe0147fd4d1dadc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51dce0ddb45a2b13a5aa5bc36878eedcf7c66589d1ade4678e9a92953f970740"></a>

## request_constraints.max_header_count_none — request_constraints.max_header_count_none / 2923dc94ea1b / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- request_constraints.max_header_count_none

<a id="canonical-2d5f318525a2ac86604417b6c41297a4effb06aa3a256c04308d94cbc6924277"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header count none.

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
max_header_count_none = {}
```

<a id="canonical-32fa16e16e48d1293df881d39d8bb5134a7ee0ee3911d772a4f5705e23705ba5"></a>

## Direct properties — request_constraints.max_header_count_none / 2923dc94ea1b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a9d2478d89e262afceb032739ab72b38a1b5c7bcbfe817fa20e41e4d8ed0cee0"></a>

## Next pages — request_constraints.max_header_count_none / 2923dc94ea1b / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-5e449dd2fe3f57c3a47c103a04bbd8ceba7d1179623b8fb6de3c63b7fe502f54"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c22ebaec30ad7f8460bae9ab5f7bf5f7868250ee8747874fe900bbe06e1e92f"></a>

## request_constraints.max_header_key_size_none — request_constraints.max_header_key_size_none / 1322afeb20ec / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- request_constraints.max_header_key_size_none

<a id="canonical-23d7f11367093bd4fe0b1a10134fe20b01b889ed1e441edb0989f288c2399ace"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header key size none.

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
max_header_key_size_none = {}
```

<a id="canonical-fc721625b1ea19a5ad3e0ee1b3ae483fc98c8a55ea98d2f5f6ed171cdf3b8156"></a>

## Direct properties — request_constraints.max_header_key_size_none / 1322afeb20ec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-330762a3f2ca297fc2bace9121974e6b9aaf02f2141ad1b213b053eb703b8a85"></a>

## Next pages — request_constraints.max_header_key_size_none / 1322afeb20ec / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-9262c92d904535185b8c5841cdfe5ebc64b5517cb448f052789c51f50961efd3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46da0deb235c90be548359df0021e080c7fbc6db699985b98f3458def6049ea4"></a>

## request_constraints.max_header_value_size_none — request_constraints.max_header_value_size_none / 3887f3ce2fe0 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- request_constraints.max_header_value_size_none

<a id="canonical-fe10e0da82d0da1620ede5c32e935ebe52237c6d1086abdc575d9e31db9387ae"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header value size none.

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
max_header_value_size_none = {}
```

<a id="canonical-ac18f8195607fbd94e4f180a1bd4b521980fc09a1bdc2d7fe3ce8a46b1a14c64"></a>

## Direct properties — request_constraints.max_header_value_size_none / 3887f3ce2fe0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4302864d484de1d94a6273412797235fdd05864989f01383bd521192cd418ed5"></a>

## Next pages — request_constraints.max_header_value_size_none / 3887f3ce2fe0 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-e9c316ff2f641a4975c83005848d33c5d0767c4a05bbb6a830b57df477d15fee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e812a086b1fc52451c5d7849935a42782b9022ad8c72dc91603d990d00cb9266"></a>

## request_constraints.max_parameter_count_none — request_constraints.max_parameter_count_none / 0555aae9f622 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- request_constraints.max_parameter_count_none

<a id="canonical-54fec6f536eb9d9a7c95d6e104e1e56c15bba57609113eff56a4689f735b8fe1"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max parameter count none.

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
max_parameter_count_none = {}
```

<a id="canonical-4c07e127b88042f3da91eabc15a891cc59f6bf1470ef297851f5edcbc62ee4b9"></a>

## Direct properties — request_constraints.max_parameter_count_none / 0555aae9f622 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1c7daa6fb3c0fb17475898387efdc05171ed282c26d457fc6fd8bdc4e11f3b3f"></a>

## Next pages — request_constraints.max_parameter_count_none / 0555aae9f622 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-ac89f75ebd0b44b9a3fac6f6d84a174d80a9750f9a4a7501f613f90e0e9c556a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f141512a4f71aeb6039ac67b18cf144cbb078e8d3c9d3be795c29e5f5d1add01"></a>

## request_constraints.max_parameter_name_size_none — request_constraints.max_parameter_name_size_none / 898aabb6ae5e / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- request_constraints.max_parameter_name_size_none

<a id="canonical-6beabc2f1c5c9f7cbb876a1e7a2b9d1422b93631466e6a23452e9e89da3e8c65"></a>

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
max_parameter_name_size_none = {}
```

<a id="canonical-58035342285dfcc20fa4408f41ff76fc2424e1eff1af07f6084f3fc717bfefef"></a>

## Direct properties — request_constraints.max_parameter_name_size_none / 898aabb6ae5e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1d753d2dd439bdc835364ae9e8e06456c6bd260df02726ae0718f0344cf0aa85"></a>

## Next pages — request_constraints.max_parameter_name_size_none / 898aabb6ae5e / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-4b652bacd47947068316b1bb3c753e8a6fd4c430176f4b86ef43103de00bc5a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b00c289eac2ede696b26afbb7d4ca1fd79943c5f177692937212dbef942edf8"></a>

## request_constraints.max_parameter_value_size_none — request_constraints.max_parameter_value_size_none / 24e47d7d4054 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- request_constraints.max_parameter_value_size_none

<a id="canonical-7bdec7e82a48311c9c0358d1a98be173b3e4dc91f52b3a0918e63ff8699da708"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max parameter value size none.

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
max_parameter_value_size_none = {}
```

<a id="canonical-3e2b27a6799b6fcd0131e6c204dfba68bd38d6fc69cbe6cba2b4913849685fee"></a>

## Direct properties — request_constraints.max_parameter_value_size_none / 24e47d7d4054 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d28dc452f8a51f10923fb9efccd8e597a07b7336e2e6841b9c951fd579dcbbc0"></a>

## Next pages — request_constraints.max_parameter_value_size_none / 24e47d7d4054 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-104c3764a83230dc671d74469bdde202e568752c2222ef65d5c22e5a71f52720"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fb5bc0e7003e6701411925d777ad11e2d6c4a8c154815f7364efb258ffd320c"></a>

## request_constraints.max_query_size_none — request_constraints.max_query_size_none / 5c7c01db09ff / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- request_constraints.max_query_size_none

<a id="canonical-63ba896825420c330a7ca3f560376955ba1eddd42cb8317f50a70ad6fbaae114"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max query size none.

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
max_query_size_none = {}
```

<a id="canonical-07c4596f275513b3198c9c9df12234138ca9a2bb060ad86d791977916947465a"></a>

## Direct properties — request_constraints.max_query_size_none / 5c7c01db09ff / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2387bb36d83fa6f7f845a92b60847707cff2cef5f2762439fe5ab0432ee9c937"></a>

## Next pages — request_constraints.max_query_size_none / 5c7c01db09ff / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-001fb5cd8f1dfe0e84be4f6fa20814bfeaec670f9e29b6452de0a3a71f1c201c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3778ad8399e8ab110964d757e51550fc66620c28df54619ed6145e233f4895f6"></a>

## request_constraints.max_request_line_size_none — request_constraints.max_request_line_size_none / 0feaa7ccad44 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- request_constraints.max_request_line_size_none

<a id="canonical-cf14153bda216fbc2c93900467bc0fd5241339dba3a1c36609462dab422778e9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max request line size none.

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
max_request_line_size_none = {}
```

<a id="canonical-78233318e2fc08e631af538a24da7481bcc119a023e26112b88992288efbbb59"></a>

## Direct properties — request_constraints.max_request_line_size_none / 0feaa7ccad44 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7453e8c893740f195ce07b70bacbac2add8aa0f8595bc44c6a512a03fdfb6f26"></a>

## Next pages — request_constraints.max_request_line_size_none / 0feaa7ccad44 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-2526596e2d713a0ba5c148327c6d808e6ff5d6023d74410e410fa211ae3623a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6999c5d2eb6f429bb1691cca1f9382886564c6768bbb2b692f384fd4a3c06154"></a>

## request_constraints.max_request_size_none — request_constraints.max_request_size_none / 45f0fe608e64 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- request_constraints.max_request_size_none

<a id="canonical-5426e810c9d4ab01fdb5640f7e530a9b2a2a124e8ec32cd2e0e15e1fcc287b37"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max request size none.

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
max_request_size_none = {}
```

<a id="canonical-39dcd7163c56e1b1fb2699daff8bc00260d21c18ffb725b79e91afdec872b3dc"></a>

## Direct properties — request_constraints.max_request_size_none / 45f0fe608e64 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2c7a2b91c1568dc87d66a6a0f564decd73ec5aabf7afed21c6e5832ddd7bd55b"></a>

## Next pages — request_constraints.max_request_size_none / 45f0fe608e64 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-bcf7aa40a79a14c558b10975e11785685b1694164272c1792754b65df37072fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-687f41f7b33da39c3c2ea8e8a4ecbdcb3fa46a673b087dbecd72e0327bba2e93"></a>

## request_constraints.max_url_size_none — request_constraints.max_url_size_none / 3bd15f3ecb60 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- request_constraints.max_url_size_none

<a id="canonical-e691a6c709605fd6bf2d448a4f89717c7ec2733b5a7b1faebf080bc0bbfd4e96"></a>

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
max_url_size_none = {}
```

<a id="canonical-2a5753164c845775427b6665f5816b344a26452b432c5b7c16d384dddee93cca"></a>

## Direct properties — request_constraints.max_url_size_none / 3bd15f3ecb60 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bed5b1c3bee82bb8b0ba2db18570d38ab25d0034808eeba71fae803992561bee"></a>

## Next pages — request_constraints.max_url_size_none / 3bd15f3ecb60 / 4

- [request_constraints](resources--service_policy_rule--reference--group-002.md#canonical-ea33ea0defd617899f8e9a768ed152310fe7d7d43854e28babdf6b358a05c886)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-529b58c9b8cc569fb2bfa7a068b2fea3f966f3818b47de925966df4ce718e56d"></a>

## segment_policy — segment_policy / 88c5e83d4253 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- segment_policy

<a id="canonical-00d30008ccdbf9b04fdb050086064d49ff6f6e9cd059d638fa7b75607ebd845d"></a>

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

<a id="canonical-9cf72b3d5a76bf6d808b8139d2045760dd06f6fd0fbfa385b7c8a89d0a6007b1"></a>

## Direct properties — segment_policy / 88c5e83d4253 / 3

- [dst_any](resources--service_policy_rule--reference--group-002.md#canonical-9785ca494e90ded8f23c1d5cc434cbe7ab9d794328a31b1c171df5320fd2ad1b): complete subsection reference.

- [dst_segments](resources--service_policy_rule--reference--group-002.md#canonical-8afab99f6a5d2a7415d39d560ea0419635d0d7fa6990fb1ade8d434576f8e087): complete subsection reference.

- [intra_segment](resources--service_policy_rule--reference--group-002.md#canonical-bb799e4ad85ad64969f0a9288bff95611bda9637bb24a68f1375b4e0a1136ffb): complete subsection reference.

- [src_any](resources--service_policy_rule--reference--group-002.md#canonical-493d037a24fae394341a40d2f4f3786e476bf421442554902928c0c1451fd555): complete subsection reference.

- [src_segments](resources--service_policy_rule--reference--group-002.md#canonical-11fb85daa08d75df5c8c2dfa22a460d4b95c37c5dc69461e7a9d412504685b34): complete subsection reference.

<a id="canonical-3b8e532ab097661d2f35b1f33bc0766f4ba1b1ef24267152065cd5de9e71aee1"></a>

## Next pages — segment_policy / 88c5e83d4253 / 4

- [segment_policy.dst_any](resources--service_policy_rule--reference--group-002.md#canonical-9785ca494e90ded8f23c1d5cc434cbe7ab9d794328a31b1c171df5320fd2ad1b)
- [segment_policy.dst_segments](resources--service_policy_rule--reference--group-002.md#canonical-8afab99f6a5d2a7415d39d560ea0419635d0d7fa6990fb1ade8d434576f8e087)
- [segment_policy.intra_segment](resources--service_policy_rule--reference--group-002.md#canonical-bb799e4ad85ad64969f0a9288bff95611bda9637bb24a68f1375b4e0a1136ffb)
- [segment_policy.src_any](resources--service_policy_rule--reference--group-002.md#canonical-493d037a24fae394341a40d2f4f3786e476bf421442554902928c0c1451fd555)
- [segment_policy.src_segments](resources--service_policy_rule--reference--group-002.md#canonical-11fb85daa08d75df5c8c2dfa22a460d4b95c37c5dc69461e7a9d412504685b34)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-9785ca494e90ded8f23c1d5cc434cbe7ab9d794328a31b1c171df5320fd2ad1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200bc4d9b2e1b31c92511e33dd446da11b756034144cdd7aa4de64778d8fc37"></a>

## segment_policy.dst_any — segment_policy.dst_any / 7dd88e6a0145 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9)
- segment_policy.dst_any

<a id="canonical-bee6e6af73eea55968109fad8d9789bc81191d5c344a54a43aaf6a0ca9b35eb7"></a>

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

<a id="canonical-018f761355e1ae66873b6d657e30133904a7df428a4cac6e65d050d94b031434"></a>

## Direct properties — segment_policy.dst_any / 7dd88e6a0145 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310a2af45c3d8c925f08f9ed899036e7ca4a8fd542441056da0af30ec4d2b12"></a>

## Next pages — segment_policy.dst_any / 7dd88e6a0145 / 4

- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-8afab99f6a5d2a7415d39d560ea0419635d0d7fa6990fb1ade8d434576f8e087"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15c7fff4bb70281c4269faad873d6be0b3e9b80a09e7a271bbebff1b52d73be7"></a>

## segment_policy.dst_segments — segment_policy.dst_segments / 28dfb3e90517 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9)
- segment_policy.dst_segments

<a id="canonical-d71707b3b9c1a09888ee79454d2819ac9a17ef0ad587841fc2e2aaca212bc5fb"></a>

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

<a id="canonical-30268621fd05a7ecfbb0def99fbe25d821ec8db261f75f36a83e549105e8e37b"></a>

## Direct properties — segment_policy.dst_segments / 28dfb3e90517 / 3

- [segments](resources--service_policy_rule--reference--group-002.md#canonical-124859838035322c89e3db7777681ff0c207589be2e34354624681b90fa17b78): complete subsection reference.

<a id="canonical-b59ff313af65d3109016bc85cb59e44269d68a4f4aab4c6cccb77b79ce4b4386"></a>

## Next pages — segment_policy.dst_segments / 28dfb3e90517 / 4

- [segment_policy.dst_segments.segments](resources--service_policy_rule--reference--group-002.md#canonical-124859838035322c89e3db7777681ff0c207589be2e34354624681b90fa17b78)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-124859838035322c89e3db7777681ff0c207589be2e34354624681b90fa17b78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c0fd3b334bb32e8c955d60a368b0abc47429bd68bb6e33dbb9e6dc348ebcf7d"></a>

## segment_policy.dst_segments.segments — segment_policy.dst_segments.segments / 611b7824a94a / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9)
- [segment_policy.dst_segments](resources--service_policy_rule--reference--group-002.md#canonical-8afab99f6a5d2a7415d39d560ea0419635d0d7fa6990fb1ade8d434576f8e087)
- segment_policy.dst_segments.segments

<a id="canonical-a21e47c705b45b3cd504b78b408415d06de0dacc51120cb5892555321cd91040"></a>

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

<a id="canonical-d3e4cc952d414e7e21ffe966abfdf8601437b78fbd493016fb888c14f81ec594"></a>

## Direct properties — segment_policy.dst_segments.segments / 611b7824a94a / 3

<a id="canonical-0d04f116ff42a697bf8959fe7e3135e1571f30d830957018ebc934a837c4fb5f"></a>

<a id="canonical-f760c4767c1bc5ed96ff1ad318301474644374da7acb1002b30b2dfcb3d2b01e"></a>

## name property — segment_policy.dst_segments.segments / 611b7824a94a / 4

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

<a id="canonical-27e9c498fb263e23682a08d045ecbd673e9ab3141db5a5f714a367a9f1e6d371"></a>

<a id="canonical-592292e3a37dcfd789d7b537db59dd04ddbff5d3225b4982dfac0dc8382aba87"></a>

## namespace property — segment_policy.dst_segments.segments / 611b7824a94a / 5

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

<a id="canonical-a9837cb0b771f0cc01ad332311761a46053bb4d35e5896c8adfdfe8d96febfc0"></a>

<a id="canonical-acdd9be3bc323a1e7681c3041c5447c336402fab2d30cd0a03f96265ad64334d"></a>

## tenant property — segment_policy.dst_segments.segments / 611b7824a94a / 6

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

<a id="canonical-d93ebab977e0521d8e7d4cbfb82df6c8e1460de8c9f43cfb5c6a8a6237a7c8c0"></a>

## Next pages — segment_policy.dst_segments.segments / 611b7824a94a / 7

- [segment_policy.dst_segments](resources--service_policy_rule--reference--group-002.md#canonical-8afab99f6a5d2a7415d39d560ea0419635d0d7fa6990fb1ade8d434576f8e087)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-bb799e4ad85ad64969f0a9288bff95611bda9637bb24a68f1375b4e0a1136ffb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2cb40e558f63d8ab94c4d7d60f7ff610da199ba038363937f95966fb28113dcf"></a>

## segment_policy.intra_segment — segment_policy.intra_segment / 66c112c98ea2 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9)
- segment_policy.intra_segment

<a id="canonical-2c7b775b1b38ffe563b94954df6dc1d33ffdeda3f7d37d521f249064e800f365"></a>

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

<a id="canonical-9c7db2770a96c9431f59c1b9034acb55e02553b2d09802104c7a35cd7f86f7ea"></a>

## Direct properties — segment_policy.intra_segment / 66c112c98ea2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ede4fd199b054b2813ffcc33ea10793f98005c0d5682c263e3b0365e5511b58f"></a>

## Next pages — segment_policy.intra_segment / 66c112c98ea2 / 4

- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-493d037a24fae394341a40d2f4f3786e476bf421442554902928c0c1451fd555"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48eb6eab4ce0943efc37e01fc7625aeffce102abca97bf800399fdfae2bfc572"></a>

## segment_policy.src_any — segment_policy.src_any / 24ec4c6c5621 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9)
- segment_policy.src_any

<a id="canonical-9fc2e32cf2dec5f7ac59fe587e9437dede8f5b0539b9e8fdb37ea319eb76bad0"></a>

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

<a id="canonical-53abfaca7a9c8be3408d33aef317607cb3be8e3d542409c53eaa01848e68ffc3"></a>

## Direct properties — segment_policy.src_any / 24ec4c6c5621 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9cad687f3c0836457284f8d5e7ce96e2351b67c36cc7d39d24698db6239e5a63"></a>

## Next pages — segment_policy.src_any / 24ec4c6c5621 / 4

- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-11fb85daa08d75df5c8c2dfa22a460d4b95c37c5dc69461e7a9d412504685b34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1893313aa629b18d4ad48311832f8d8615fec7876df63a7b1bca8a406ba2916f"></a>

## segment_policy.src_segments — segment_policy.src_segments / 416990067c03 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9)
- segment_policy.src_segments

<a id="canonical-44d5b0b5c5aeaccd04042c3346814703f68c3e261ea03a82c4f247228bd3ebfd"></a>

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

<a id="canonical-b6495121a10bd752333439d1d8706d92b898c96da507d5851dbc40dd5203bcfb"></a>

## Direct properties — segment_policy.src_segments / 416990067c03 / 3

- [segments](resources--service_policy_rule--reference--group-002.md#canonical-2a08c893c2054ad855c971fd9aefe9721d8a0124534f294bee0716ffded2dcbb): complete subsection reference.

<a id="canonical-84a2fcdbfd517c244e507c62abe49e6dd8cfcc956451a1e2b1af9fea216b56a6"></a>

## Next pages — segment_policy.src_segments / 416990067c03 / 4

- [segment_policy.src_segments.segments](resources--service_policy_rule--reference--group-002.md#canonical-2a08c893c2054ad855c971fd9aefe9721d8a0124534f294bee0716ffded2dcbb)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-2a08c893c2054ad855c971fd9aefe9721d8a0124534f294bee0716ffded2dcbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b567daf173ba41eadaa50d767161f36c0cfc14d5ff9374aac7bc9195ee51aaef"></a>

## segment_policy.src_segments.segments — segment_policy.src_segments.segments / f9af5d219246 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-39ed0a76d6abb2ae046489b4bc093c8c456233b10c3d63debbcf11521d5ffde9)
- [segment_policy.src_segments](resources--service_policy_rule--reference--group-002.md#canonical-11fb85daa08d75df5c8c2dfa22a460d4b95c37c5dc69461e7a9d412504685b34)
- segment_policy.src_segments.segments

<a id="canonical-cfa1d3f129a3e6212c669a54703c4f78ed62f0f86fd36bb9673a43fb3e1295e7"></a>

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

<a id="canonical-361e8340b568122656ea2519f0d38512edee3be727fc193e903fa717decc8b97"></a>

## Direct properties — segment_policy.src_segments.segments / f9af5d219246 / 3

<a id="canonical-49c0c26a784da1e28a6c324d2827450d7620da81078f6c723e3781de0dffb3a7"></a>

<a id="canonical-1fa1fb3da00ec9a54a3b5914cd0ecb9a12e5dce6f14c86912e7f8a60f0650966"></a>

## name property — segment_policy.src_segments.segments / f9af5d219246 / 4

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

<a id="canonical-b7fb798e441421c20357f0973c3b92d37963aeb00f222876bdb00b8d63138b7d"></a>

<a id="canonical-cd6b3ce27ca905c085b380c679e12c2b58f9735bb43aedf4989b76dcfeff2a16"></a>

## namespace property — segment_policy.src_segments.segments / f9af5d219246 / 5

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

<a id="canonical-31b4fe1d337f8c60095e42b0336c5491ca3b1e7e6927997b7c4c4493559a8cb9"></a>

<a id="canonical-649227463caa547cd54074d5d339c606e1ff2f7e1d0ebda0efb0306e16b01345"></a>

## tenant property — segment_policy.src_segments.segments / f9af5d219246 / 6

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

<a id="canonical-f1f1de36907ca77c6c559cac7b1ff7a64d36226336e7b9e04c7c6413e5ea6da9"></a>

## Next pages — segment_policy.src_segments.segments / f9af5d219246 / 7

- [segment_policy.src_segments](resources--service_policy_rule--reference--group-002.md#canonical-11fb85daa08d75df5c8c2dfa22a460d4b95c37c5dc69461e7a9d412504685b34)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-49d1dcedd566f98a405b94194bc73262f7147467a3f23b9ded4df6aaafb63e75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-242b7f7cfaefd84c73cd31b0af2a4d853cc9ee45bef5f33530fcc347510d5949"></a>

## timeouts — timeouts / a7e3a37e1424 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- timeouts

<a id="canonical-e8d4a4db8c37bbc139975eac7fab78a32fe33ee7dff1c263ef6de7c838963d40"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3bdb95044429104918b9b3043a4128bd07dd49f85014d5f71c09145d44301c73"></a>

## Direct properties — timeouts / a7e3a37e1424 / 3

<a id="canonical-2afdc9249a3aadf139d1bc74edcaa4434ca595129d225216d6c1d5e68c8dc570"></a>

<a id="canonical-52cad6bafd3dfc4d6cde0368a93e0fd3657033af9dcfd9232425d777c9d0ffd3"></a>

## create property — timeouts / a7e3a37e1424 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f425c20ce69983bf50624410f75ebdda2a1e4304250e2c38651074936a2d02af"></a>

<a id="canonical-846ad280c44d13e70ed0c7c1beed51296203119072055ae65d114e835de71c37"></a>

## delete property — timeouts / a7e3a37e1424 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-c776f046ac6043d6fdffa17ecd30232d6ab9ea2f0bd34068269628168bcc45d1"></a>

<a id="canonical-b0ec85a6dd4ac3f8f73e7fa967d00ce1d8ad1ff219649e70961d89d1e4e3b1f3"></a>

## read property — timeouts / a7e3a37e1424 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-5829707d9ce173638483a737b2d61f9bb6c546ff5be9634e5a038152bc7604f9"></a>

<a id="canonical-61fb81245ad584fb90c9f9409660aa4d66741d4d72f4c20142b4e646aac39ee0"></a>

## update property — timeouts / a7e3a37e1424 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-9e62acf1f32524dbd10e4224242e145c698af76a82388593c191c5f74594c653"></a>

## Next pages — timeouts / a7e3a37e1424 / 8

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-a00396ea59b4ee1bd6e1ea84d1322f649375d00c7f7360e57c8af940324e416b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47817fef6eca7c52b0cc8675a222a124b963f1dd9d62ba86b3753e98387e3216"></a>

## tls_fingerprint_matcher — tls_fingerprint_matcher / cfaa19eeb7be / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- tls_fingerprint_matcher

<a id="canonical-b02fed0ec6f83276f5054e82dee19e02e679bcd5a8c478e201220c87e66bfe6b"></a>

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

<a id="canonical-f4807aa203a4aad75126c0d54eff32ff90c302c7d0ea1610bdef5f210d30a834"></a>

## Direct properties — tls_fingerprint_matcher / cfaa19eeb7be / 3

<a id="canonical-a75c1e7eac11e7a1dd8191bad45ab6214a8e4d4b56db2b4b4e7a1e114ee53245"></a>

<a id="canonical-0dd528550cf4e7c3ad5598b4822369d48f4b64072fff7b0cc3ed3bef29a3760d"></a>

## classes property — tls_fingerprint_matcher / cfaa19eeb7be / 4

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

<a id="canonical-35b1c82d4f4362107a7bf3d8210cfd344b47a7293dadbc7e7e14dd227543a04d"></a>

<a id="canonical-e88b4e071afe344e4867e86c6d4bbc088e47d4767d953bd024d0956f3ce2d7ba"></a>

## exact_values property — tls_fingerprint_matcher / cfaa19eeb7be / 5

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

<a id="canonical-d8179763066ddebb518f5499e7a6990432d3f04d524284a500265b724fc20fbb"></a>

<a id="canonical-815977abbd32d54d7ede0d8d9e96e092444ad745bd0be80dbd58fe112e2de095"></a>

## excluded_values property — tls_fingerprint_matcher / cfaa19eeb7be / 6

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

<a id="canonical-e8894f69883fff9b240adf25e21da5eef5e64125ae8f720a549b9903b778d082"></a>

## Next pages — tls_fingerprint_matcher / cfaa19eeb7be / 7

- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-61e3a92cf25302dbe0ec3ece5b26351c9fa5f5673b5dddca9940806ace055e76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6316f8237401bbcb7a7d3d9c051d1836b967e587d7ae0003b54a797f2ca66db0"></a>

## waf_action — waf_action / 9204880c2ed7 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- waf_action

<a id="canonical-165aaa4d7dd235e1809f7ecc1c949e698490430debc8667276bc964700f63892"></a>

Type: `"object"`. single nested block, Optional.

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Upstream description:

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall_detection_control",
    "none"),
  validators.ConflictingObjectAttributes("app_firewall_detection_control",
    "waf_skip_processing"),
  validators.ConflictingObjectAttributes("none",
    "waf_skip_processing")}
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
  "x-ves-oneof-field-action_type": "[\"app_firewall_detection_control\",\"none\",\"waf_skip_processing\"]"
}
```

Terraform syntax:

```terraform
waf_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-6cfc377199f3722621cd7c9c12dd916b8efa5fd05d8095e41c0de9d331f85a2d"></a>

## Direct properties — waf_action / 9204880c2ed7 / 3

- [app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-8ff1d7a60dc33cf1b833e8d925e5fe038ec9a7471343645442649192192dd690): complete subsection reference.

- [none](resources--service_policy_rule--reference--group-002.md#canonical-1a64625e402b5495ddc7ba5e28cd2935eb47e8d9a1cd52a1b833dd43e5420fe6): complete subsection reference.

- [waf_skip_processing](resources--service_policy_rule--reference--group-002.md#canonical-b0d21ffcdf1c5a82929ac888d7ee8cf80d09d83767e92c93298db04fb3f32ab6): complete subsection reference.

<a id="canonical-72515ed1b88da84b0ce7ac2e609012ee97793b8d2c04568695a7c49755c74a0e"></a>

## Next pages — waf_action / 9204880c2ed7 / 4

- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-8ff1d7a60dc33cf1b833e8d925e5fe038ec9a7471343645442649192192dd690)
- [waf_action.none](resources--service_policy_rule--reference--group-002.md#canonical-1a64625e402b5495ddc7ba5e28cd2935eb47e8d9a1cd52a1b833dd43e5420fe6)
- [waf_action.waf_skip_processing](resources--service_policy_rule--reference--group-002.md#canonical-b0d21ffcdf1c5a82929ac888d7ee8cf80d09d83767e92c93298db04fb3f32ab6)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-8ff1d7a60dc33cf1b833e8d925e5fe038ec9a7471343645442649192192dd690"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebf20f733e852c31faa9bcc4ef3373a5b25e8c55559c697b300673b9dfdcfa88"></a>

## waf_action.app_firewall_detection_control — waf_action.app_firewall_detection_control / 26c56fbfe965 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-61e3a92cf25302dbe0ec3ece5b26351c9fa5f5673b5dddca9940806ace055e76)
- waf_action.app_firewall_detection_control

<a id="canonical-b222bf073801535e0ee7715168a722014bd0ad58e8e9dcb296e2f52e610e93cb"></a>

Type: `"object"`. single nested block, Optional.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

Receipt-pinned upstream constraints:

```json
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
app_firewall_detection_control {
  # Configure direct properties listed below.
}
```

<a id="canonical-8cf70881f7b30a2945c6320f6bafe60393e2cd09043c5e6b18bb198a02d283e5"></a>

## Direct properties — waf_action.app_firewall_detection_control / 26c56fbfe965 / 3

- [exclude_attack_type_contexts](resources--service_policy_rule--reference--group-002.md#canonical-293dbdf59d7f1043a952f2219400b170327f0df6b82bf2955fefb766dea19b64): complete subsection reference.

- [exclude_bot_name_contexts](resources--service_policy_rule--reference--group-002.md#canonical-3972e219042d533bcb43a178ea0548e166a66ec307884b6e1d8f77bd8ba5cef2): complete subsection reference.

- [exclude_signature_contexts](resources--service_policy_rule--reference--group-002.md#canonical-8bc61aad35df6212ee8fa9997b69fc5bf9a3cec0e23cffc89990ac67c948c621): complete subsection reference.

- [exclude_violation_contexts](resources--service_policy_rule--reference--group-002.md#canonical-bb2d133c353e219fd810eb24b5431af5290f55c0b1a83f511f8a2a782a675e0f): complete subsection reference.

<a id="canonical-05794d061725d1d9adb29cdecca566d0afc8f6faf97b10fae839de5d8774591a"></a>

## Next pages — waf_action.app_firewall_detection_control / 26c56fbfe965 / 4

- [waf_action.app_firewall_detection_control.exclude_attack_type_contexts](resources--service_policy_rule--reference--group-002.md#canonical-293dbdf59d7f1043a952f2219400b170327f0df6b82bf2955fefb766dea19b64)
- [waf_action.app_firewall_detection_control.exclude_bot_name_contexts](resources--service_policy_rule--reference--group-002.md#canonical-3972e219042d533bcb43a178ea0548e166a66ec307884b6e1d8f77bd8ba5cef2)
- [waf_action.app_firewall_detection_control.exclude_signature_contexts](resources--service_policy_rule--reference--group-002.md#canonical-8bc61aad35df6212ee8fa9997b69fc5bf9a3cec0e23cffc89990ac67c948c621)
- [waf_action.app_firewall_detection_control.exclude_violation_contexts](resources--service_policy_rule--reference--group-002.md#canonical-bb2d133c353e219fd810eb24b5431af5290f55c0b1a83f511f8a2a782a675e0f)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-61e3a92cf25302dbe0ec3ece5b26351c9fa5f5673b5dddca9940806ace055e76)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-293dbdf59d7f1043a952f2219400b170327f0df6b82bf2955fefb766dea19b64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-830acf1586a3fbdc1094ea6b8c3ef761f1e26181ccfbf6ab725142daedfaf0f3"></a>

## waf_action.app_firewall_detection_control.exclude_attack_type_contexts — waf_action.app_firewall_detection_control.exclude_attack_type_contexts / b2852b871e3c / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-61e3a92cf25302dbe0ec3ece5b26351c9fa5f5673b5dddca9940806ace055e76)
- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-8ff1d7a60dc33cf1b833e8d925e5fe038ec9a7471343645442649192192dd690)
- waf_action.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-0141ea8fef8d5cde72450c9f0162648107f851fafe45ace0a9f860147ec4ac5a"></a>

Type: `"object"`. list nested block, Optional.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_attack_type_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2e0e9673a3adeeb56682e7888a7ea081b5dda7554cf91f577636f5e5d6229cec"></a>

## Direct properties — waf_action.app_firewall_detection_control.exclude_attack_type_contexts / b2852b871e3c / 3

<a id="canonical-565b7dbaf3dfb9b0402e314d970beb67d15c02018fd6a5a070756d47fd256c41"></a>

<a id="canonical-42dca42404212089b16a0f123f45e8c2179ead731e77dc1d83f8ce09cdc7dda3"></a>

## context property — waf_action.app_firewall_detection_control.exclude_attack_type_contexts / b2852b871e3c / 4

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4249d3c942ed9b3ae919d149b40e9ee09aa613c909b53b1c3cd5c62e94c1294f"></a>

<a id="canonical-df13f818bb317d1294aa52b32c2127c84cb81965a6cb9f055585bf2a02bb74b1"></a>

## context_name property — waf_action.app_firewall_detection_control.exclude_attack_type_contexts / b2852b871e3c / 5

Type: `"string"`. Optional.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-47ffff15f1fbfdc13e11a052bf4b3e6a98872b5572cb6869ae28b80c415e261d"></a>

<a id="canonical-cde65af6db8c21130fcc40819003ef537b20314ac883477ef8f8d761a1461cee"></a>

## exclude_attack_type property — waf_action.app_firewall_detection_control.exclude_attack_type_contexts / b2852b871e3c / 6

Type: `"string"`. Optional.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Upstream description:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-026aeebd0bc627aef50a7f05a5efe1d0cad65c3955a9af237cf8b234a1d34dcb"></a>

## Next pages — waf_action.app_firewall_detection_control.exclude_attack_type_contexts / b2852b871e3c / 7

- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-8ff1d7a60dc33cf1b833e8d925e5fe038ec9a7471343645442649192192dd690)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-3972e219042d533bcb43a178ea0548e166a66ec307884b6e1d8f77bd8ba5cef2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2fe896494844cfb4624eb41e9d181ebe6e68bfc0244cd52d3db641432ac619ba"></a>

## waf_action.app_firewall_detection_control.exclude_bot_name_contexts — waf_action.app_firewall_detection_control.exclude_bot_name_contexts / 9c75d4341752 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-61e3a92cf25302dbe0ec3ece5b26351c9fa5f5673b5dddca9940806ace055e76)
- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-8ff1d7a60dc33cf1b833e8d925e5fe038ec9a7471343645442649192192dd690)
- waf_action.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-ad676add7878cdbfa7120a9be8c2910d7e2654dffab880b47c421344cb73b732"></a>

Type: `"object"`. list nested block, Optional.

Bot Names to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("bot_name")}
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_bot_name_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-d2fdda6d974be27f0403b9379e41981a60443f70a3f7c57ad0b53dc6c72dbad2"></a>

## Direct properties — waf_action.app_firewall_detection_control.exclude_bot_name_contexts / 9c75d4341752 / 3

<a id="canonical-79c52146348d21bf97af927ed738af07e5058446bf1cc7e07c364f0344fbd9d9"></a>

<a id="canonical-6a6d41afe51d03fbdd7664f80ad1e60ea9d3fb082ed2ea81721bb0f0519594b1"></a>

## bot_name property — waf_action.app_firewall_detection_control.exclude_bot_name_contexts / 9c75d4341752 / 4

Type: `"string"`. Optional.

Bot Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-515f540edce5ceb7a7ae2e96a866de8014a6cdd686f92c488eb862d0461f6f02"></a>

## Next pages — waf_action.app_firewall_detection_control.exclude_bot_name_contexts / 9c75d4341752 / 5

- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-8ff1d7a60dc33cf1b833e8d925e5fe038ec9a7471343645442649192192dd690)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-8bc61aad35df6212ee8fa9997b69fc5bf9a3cec0e23cffc89990ac67c948c621"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0bb87493348bc1675669b3dde047cbf6e93803aa175e0a93ab24f40798d9a7aa"></a>

## waf_action.app_firewall_detection_control.exclude_signature_contexts — waf_action.app_firewall_detection_control.exclude_signature_contexts / 09a30e87e3bd / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-61e3a92cf25302dbe0ec3ece5b26351c9fa5f5673b5dddca9940806ace055e76)
- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-8ff1d7a60dc33cf1b833e8d925e5fe038ec9a7471343645442649192192dd690)
- waf_action.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-82b99c25eee77b989b8d931a54110abce86f7d2d637a08b5a21c8f2f3ca065a4"></a>

Type: `"object"`. list nested block, Optional.

Signature IDs to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("signature_id")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_signature_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-6939dcf56eddf89661037fbedb927175657134b8a843e9931d08f088bc6f22a8"></a>

## Direct properties — waf_action.app_firewall_detection_control.exclude_signature_contexts / 09a30e87e3bd / 3

<a id="canonical-32f70242a2c1d1daf64b9fe67df3798f23366006d6b1860df342786e4039fccc"></a>

<a id="canonical-f89044b8331bc3fa4a68b210da64a6d0162a5f0d96e1a53de5260f77e0ada643"></a>

## context property — waf_action.app_firewall_detection_control.exclude_signature_contexts / 09a30e87e3bd / 4

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2590fee4736265af2844c871416ae80a3c70bd10bbf96197c9eef8fad10472ea"></a>

<a id="canonical-2fb02bd55421270a056b8f80ee6d6e9db887eedfdc57a6ac096a7c1a3ed26350"></a>

## context_name property — waf_action.app_firewall_detection_control.exclude_signature_contexts / 09a30e87e3bd / 5

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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

<a id="canonical-73bfcbc190821141dfa07aa31d428f86d7947e608deb4e4fe60d4175176d9acb"></a>

<a id="canonical-6e409485e655f0e38d60957b2a7614264212351efb836eb9289c531bfeee0f2d"></a>

## signature_id property — waf_action.app_firewall_detection_control.exclude_signature_contexts / 09a30e87e3bd / 6

Type: `"number"`. Optional.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 299999999),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

<a id="canonical-62b6f0e22610011288a358b2e495ecb9585f5804caf19910c73ba0e6d2e26056"></a>

## Next pages — waf_action.app_firewall_detection_control.exclude_signature_contexts / 09a30e87e3bd / 7

- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-8ff1d7a60dc33cf1b833e8d925e5fe038ec9a7471343645442649192192dd690)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-bb2d133c353e219fd810eb24b5431af5290f55c0b1a83f511f8a2a782a675e0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c7522a6eb119e4a916d1bc286001a3ffe798e289e7af3f1a8d776f4f4dc48b3"></a>

## waf_action.app_firewall_detection_control.exclude_violation_contexts — waf_action.app_firewall_detection_control.exclude_violation_contexts / 8e7395f8c180 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-61e3a92cf25302dbe0ec3ece5b26351c9fa5f5673b5dddca9940806ace055e76)
- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-8ff1d7a60dc33cf1b833e8d925e5fe038ec9a7471343645442649192192dd690)
- waf_action.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-829d489cb333d4c1b7dbb00179b4adce85e82d1e9465e95aac54dbfad01cf079"></a>

Type: `"object"`. list nested block, Optional.

Violations to be excluded for the defined match criteria.

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_violation_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-7aba685c782b038ec700daab9a836863ac15d63c9a6d875951796510318af000"></a>

## Direct properties — waf_action.app_firewall_detection_control.exclude_violation_contexts / 8e7395f8c180 / 3

<a id="canonical-a77b6e7c95478cdeb4570343f550b560877f331f00947c695779d95719845615"></a>

<a id="canonical-efca0ba02738b7bd3b739ea189c09b240fc5ee36da3f8f4a467b5509d5189ed7"></a>

## context property — waf_action.app_firewall_detection_control.exclude_violation_contexts / 8e7395f8c180 / 4

Type: `"string"`. Optional.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-365da9ec4d8481be8508de9fbae7b046fd8253f805f49c9d2fc450ede799142d"></a>

<a id="canonical-703de778922b6da07bdbded3ba50a7ac050bfed6c2eb73ff23b585ffeca14cf5"></a>

## context_name property — waf_action.app_firewall_detection_control.exclude_violation_contexts / 8e7395f8c180 / 5

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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

<a id="canonical-3f522efa589cdb4cfefa9745324ae230863c08545c41ffd6642c4c93e1012b0b"></a>

<a id="canonical-c73f96e29693fa80b932a958fb3cdd9cff05504b00c4495774ce6a6dff1840ad"></a>

## exclude_violation property — waf_action.app_firewall_detection_control.exclude_violation_contexts / 8e7395f8c180 / 6

Type: `"string"`. Optional.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
List of all supported Violation Types VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER
VIOL\_HTTP\_RESPONSE\_STATUS VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD
VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED
VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS.. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Upstream description:

List of all supported Violation Types

VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER VIOL\_HTTP\_RESPONSE\_STATUS
VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED
VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS
VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT
VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION
VIOL\_HTTP\_PROTOCOL\_CRLF\_CHARACTERS\_BEFORE\_REQUEST\_START
VIOL\_HTTP\_PROTOCOL\_NO\_HOST\_HEADER\_IN\_HTTP\_1\_1\_REQUEST
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_PARAMETERS\_PARSING
VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS
VIOL\_HTTP\_PROTOCOL\_CONTENT\_LENGTH\_SHOULD\_BE\_A\_POSITIVE\_NUMBER
VIOL\_EVASION\_DIRECTORY\_TRAVERSALS VIOL\_MALFORMED\_REQUEST VIOL\_EVASION\_MULTIPLE\_DECODING
VIOL\_DATA\_GUARD VIOL\_EVASION\_APACHE\_WHITESPACE VIOL\_COOKIE\_MODIFIED
VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS VIOL\_EVASION\_IIS\_BACKSLASHES
VIOL\_EVASION\_PERCENT\_U\_DECODING VIOL\_EVASION\_BARE\_BYTE\_DECODING VIOL\_EVASION\_BAD\_UNESCAPE
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_FORMDATA\_REQUEST\_PARSING
VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST
VIOL\_HTTP\_PROTOCOL\_HIGH\_ASCII\_CHARACTERS\_IN\_HEADERS VIOL\_ENCODING VIOL\_COOKIE\_MALFORMED
VIOL\_GRAPHQL\_FORMAT VIOL\_GRAPHQL\_MALFORMED VIOL\_GRAPHQL\_INTROSPECTION\_QUERY.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIOL_NONE",
  "enum": [
    "VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START",
    "VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-339925f514d23b438ec014db5bd96b0c81d4c5937d62fb49219ad4894e77dcdd"></a>

## Next pages — waf_action.app_firewall_detection_control.exclude_violation_contexts / 8e7395f8c180 / 7

- [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-8ff1d7a60dc33cf1b833e8d925e5fe038ec9a7471343645442649192192dd690)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-1a64625e402b5495ddc7ba5e28cd2935eb47e8d9a1cd52a1b833dd43e5420fe6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77266013399e6f45b75e017238b777198655e102c7416653a62bb924f26e050b"></a>

## waf_action.none — waf_action.none / 53aeabb3197c / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-61e3a92cf25302dbe0ec3ece5b26351c9fa5f5673b5dddca9940806ace055e76)
- waf_action.none

<a id="canonical-23dc649e6ea1d40792de0b55a22c583d39035e7d969b0c7bb3de265023dc7e12"></a>

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
none = {}
```

<a id="canonical-188a5d2d97aef2f448b5466cafefbaff7eb315ddaaaee2de7d0842737f506923"></a>

## Direct properties — waf_action.none / 53aeabb3197c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4ccfe1f2c69b7c9aac5451da83caf9f62ea1117122d94c28dc968af23129fbec"></a>

## Next pages — waf_action.none / 53aeabb3197c / 4

- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-61e3a92cf25302dbe0ec3ece5b26351c9fa5f5673b5dddca9940806ace055e76)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-b0d21ffcdf1c5a82929ac888d7ee8cf80d09d83767e92c93298db04fb3f32ab6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10edd80647697bf4987a86cadc8f236904d8fe702e06f8c3c210a3fae335e407"></a>

## waf_action.waf_skip_processing — waf_action.waf_skip_processing / 9b42588e7c03 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-db658ade9088ef2b50587930ef0c3849ed77076522cfcd5c96a0ce8d5cea3d6d)
- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-61e3a92cf25302dbe0ec3ece5b26351c9fa5f5673b5dddca9940806ace055e76)
- waf_action.waf_skip_processing

<a id="canonical-f70a698546ade0558fad984e3f0d5f4d7b92ccc4d7e29e09b0804699138e21b5"></a>

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
waf_skip_processing = {}
```

<a id="canonical-964eb5966f66616a40241eede44e88f2d6b59a8a9f1bae5f3f0fbbc46ce49b2f"></a>

## Direct properties — waf_action.waf_skip_processing / 9b42588e7c03 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-80c4e21df4cdd4bcda46dba3d6ee229c33234d1adb72c796fe81982f0e37a8f0"></a>

## Next pages — waf_action.waf_skip_processing / 9b42588e7c03 / 4

- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-61e3a92cf25302dbe0ec3ece5b26351c9fa5f5673b5dddca9940806ace055e76)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
