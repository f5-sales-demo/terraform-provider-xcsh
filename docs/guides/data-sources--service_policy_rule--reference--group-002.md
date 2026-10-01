---
page_title: "xcsh_service_policy_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule reference."
---

# xcsh_service_policy_rule reference

<a id="canonical-363225290f667bd87910a6a6a73a0b71cc3b27e21aa0cd285ef7f491cfd548a1"></a>

## mum_action.default — mum_action.default / f3e81297f8b6 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [mum_action](data-sources--service_policy_rule--reference--group-001.md#canonical-137e55d694ff101b1eceed4b0f2170fade6d8b7a600a590d2a83bb1efb5d0eb3)
- mum_action.default

<a id="canonical-cfa1b83a8d5be05b8877e6dd9e43418e3c6a533a567ba11f1b3cc5c44148bca2"></a>

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

<a id="canonical-ef5182dc66b75cfdb7a1fd57a7001797daf001f36af9364a44659c62b7a9a7f5"></a>

## Direct properties — mum_action.default / f3e81297f8b6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6098b7ca667f4923251427c2e4bb0f4e94157d43de66adeddb61a850e69c919"></a>

## Next pages — mum_action.default / f3e81297f8b6 / 4

- [mum_action](data-sources--service_policy_rule--reference--group-001.md#canonical-137e55d694ff101b1eceed4b0f2170fade6d8b7a600a590d2a83bb1efb5d0eb3)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-e6b52c574809f67bf9b2e3cc9c7f95a712b9d48e5d7f86fe7aa78e1c637e099e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b646ab0d4586f9e627fcada3b53d3b736b7341ba40263b1f636513d9bce4d24"></a>

## mum_action.skip_processing — mum_action.skip_processing / 7c60aabfca48 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [mum_action](data-sources--service_policy_rule--reference--group-001.md#canonical-137e55d694ff101b1eceed4b0f2170fade6d8b7a600a590d2a83bb1efb5d0eb3)
- mum_action.skip_processing

<a id="canonical-10d9192d32a10830c275f5b628986ef6109007f710271b8a25aac62e0bfcb4f0"></a>

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

<a id="canonical-47180187bd49722adbd83948062b059628ce6b2d89c7da0a5490be05712e04b8"></a>

## Direct properties — mum_action.skip_processing / 7c60aabfca48 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ba55a8cf0f626d6907683f2c00df081c5d1239efede01197f49c67dee97a6425"></a>

## Next pages — mum_action.skip_processing / 7c60aabfca48 / 4

- [mum_action](data-sources--service_policy_rule--reference--group-001.md#canonical-137e55d694ff101b1eceed4b0f2170fade6d8b7a600a590d2a83bb1efb5d0eb3)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-62ecded1c4e84948a11dba301846deb27ceaebac027c0ed9896d35e68b2ae3c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21b07837f29e45c0054986698cc509543502d18e44415ab21ee082e34e3bea2b"></a>

## path — path / 8a0d3f8f8485 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- path

<a id="canonical-0d9e1e172dd195191983d5fe048baba005f96c513aae9ae293164121272c89ab"></a>

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

<a id="canonical-6087edfd43e5a1f5dc2f1927183207b953ce10b06ac91499ac53797d9b2906d5"></a>

## Direct properties — path / 8a0d3f8f8485 / 3

<a id="canonical-af553a240a535cc24966fa8d4e17937ac5f930e70455a89a23a1ecbd1e985a12"></a>

<a id="canonical-931f6e2b37ef82e2956d90ce45fd63cb6f339ef223d6f5fb59fe72a1c48ce0f5"></a>

## encoded_path_matcher property — path / 8a0d3f8f8485 / 4

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

<a id="canonical-4959f2b757808631d6ed53b96654146b6bf607ba9a3780990dd670f6d2cf6f25"></a>

<a id="canonical-9c2ee23dc820cfe356a3f0f2dfc3436f9657d9557e142c3dad6cfeb1f64420b3"></a>

## exact_values property — path / 8a0d3f8f8485 / 5

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

<a id="canonical-1a0aa613b0d6d05575147f6ca65dfc4ca09a4831e07e8b1e0ea771d582abf9b7"></a>

<a id="canonical-97e5ad1c3509dbd3ea195d156be787c141e963861bd2dca0bfa4ed867c0740a1"></a>

## invert_matcher property — path / 8a0d3f8f8485 / 6

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

<a id="canonical-fba100c8bff57663f468ee1f0a15e6d61a6bdddeb1a00e9268a259adcdf31826"></a>

<a id="canonical-0f037c20f73eeaa5be45ebda8fe98803eae1a1a211ede4eafe90f1b860ddc4a5"></a>

## prefix_values property — path / 8a0d3f8f8485 / 7

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

<a id="canonical-b5ddf173fed2571c218fc8a1700acba502054c43bb92f04ee10b16a21980034c"></a>

<a id="canonical-88295e0c04d688c3fe8efcb638d5e5a15126c19064fd49aebfb98973dfdf17c1"></a>

## regex_values property — path / 8a0d3f8f8485 / 8

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

<a id="canonical-6ab019f53aa7cf0987168cfa0251ca752904544aacb40f907228da318a94ca95"></a>

<a id="canonical-314bed3cd669d967894ba7e2773101c24b6a1a29e5b192df734d469ee868449e"></a>

## suffix_values property — path / 8a0d3f8f8485 / 9

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

<a id="canonical-1b4a2c22edd27c91ec14796e33888b37742378dd1c01ee36e1a38020ed80aa62"></a>

<a id="canonical-05ade2b10faa9ea87a6acfe2b2a3597d77e0111f872bc51fa55678de8017f125"></a>

## transformers property — path / 8a0d3f8f8485 / 10

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

<a id="canonical-d75b88d8b5fc422e9fa9e71accd3869c47caa2c88298376dc51f7cbd3c498a1b"></a>

## Next pages — path / 8a0d3f8f8485 / 11

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-91908ed3f38a27795fb956d472ae83414d3046dc441e006972b9304b7899f358"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14da83a5581855c0f47cffc91c8b496bd0fee4d21263455750bb3d1e994e0396"></a>

## port_matcher — port_matcher / c0fda6a04f60 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- port_matcher

<a id="canonical-1d9c6a10ef27afc42e86601a626d30d91f085a40c8de2dc3734f3d1205357c56"></a>

Type: `"single"`. Computed.

Port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true. Server applies default when omitted.

Upstream description:

A port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7ae2f6f1e38fd136b95368265cab00f92d87d929bd97267fdbfcfcf6ff28029d"></a>

## Direct properties — port_matcher / c0fda6a04f60 / 3

<a id="canonical-fcfb0321a3f2f2c90239e0488687ce6348f6c5256a4cc69cb9da530fdd4d6d17"></a>

<a id="canonical-f720f98512006202ee7284c396eaf0ffebfdf7dafb00a37c54e33658ebae461f"></a>

## invert_matcher property — port_matcher / c0fda6a04f60 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-d1712ff694536f28c86c57ff2f5009fa6ad7c95b11ceed181a647a57c80b6594"></a>

<a id="canonical-1a82b05d4b0fc4d73a2e5fc076f6bdb18c961136d78f8441d9bcd212317dc8b4"></a>

## ports property — port_matcher / c0fda6a04f60 / 5

Type: `["list", "string"]`. Computed.

List of strings, each of which is a single port value or a tuple of start and end port values
separated by '-'. The start and end values are considered to be part of the range.

Upstream description:

A list of strings, each of which is a single port value or a tuple of start and end port values
separated by "-". The start and end values are considered to be part of the range.

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

<a id="canonical-d4af09f9f1fc158aeceed0741ad11e9ec3c19c195283129e43bb45380a4c93f0"></a>

## Next pages — port_matcher / c0fda6a04f60 / 6

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-040e25c6c4a7cd4c4f103eb60316232c349443b45cd21f601c0b5cbaf73b42e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26f75d7339eb5b2761ba2187a779caa99e3ace244c1cebe3a25d0fa2366df57a"></a>

## query_params — query_params / a0cd2958fbb8 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- query_params

<a id="canonical-142ab1ad452aa2896d9e4246882edda7180e3353549888197881edb7741c9d70"></a>

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

<a id="canonical-90fc5c1409820b193687ffd1d83b0e3d00e17ebe62129a4485bfb3748cd81ade"></a>

## Direct properties — query_params / a0cd2958fbb8 / 3

- [check_not_present](data-sources--service_policy_rule--reference--group-002.md#canonical-e83fe9d08fd982daad7be1028f3735cb9e9b7eea4199ec40643f8036911826ca): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-002.md#canonical-bc1dfc2276984a71923214d9f3b50b2200cf5b2a02e1b019852c0caa7eb882a5): complete subsection reference.

<a id="canonical-533e091ae19e5a73ac1ffd866a36368a8a471adabc725e0089b2abd4b179bf59"></a>

<a id="canonical-f7509193af99fdc04d55f38d478e219a0e4d61631b60e44e78eb0b3fd5c55ec4"></a>

## invert_matcher property — query_params / a0cd2958fbb8 / 4

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

- [item](data-sources--service_policy_rule--reference--group-002.md#canonical-df07cef78699c3eb7546c3592d0b45e5d54b131ab060d10034a5d79232e3ccc2): complete subsection reference.

<a id="canonical-366485fae3ff67f7634f5e3fd94a9bcee1714d0000d6e37fc2eee959cbfdda38"></a>

<a id="canonical-32d8574b83ddac775084899493db9ce3fcc1d7a21ccb2de97798bfb875eb07ad"></a>

## key property — query_params / a0cd2958fbb8 / 5

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

<a id="canonical-d2ef71c58c7bf9b6519bd821969eb63dbff911030c5db6f39816a5f232ce31c4"></a>

## Next pages — query_params / a0cd2958fbb8 / 6

- [query_params.check_not_present](data-sources--service_policy_rule--reference--group-002.md#canonical-e83fe9d08fd982daad7be1028f3735cb9e9b7eea4199ec40643f8036911826ca)
- [query_params.check_present](data-sources--service_policy_rule--reference--group-002.md#canonical-bc1dfc2276984a71923214d9f3b50b2200cf5b2a02e1b019852c0caa7eb882a5)
- [query_params.item](data-sources--service_policy_rule--reference--group-002.md#canonical-df07cef78699c3eb7546c3592d0b45e5d54b131ab060d10034a5d79232e3ccc2)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-e83fe9d08fd982daad7be1028f3735cb9e9b7eea4199ec40643f8036911826ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9839677f08a981c0966408d24e3e5a5d2d86a8823e7d4b5a6fe622d3c9b86c1"></a>

## query_params.check_not_present — query_params.check_not_present / 83f92581794f / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-040e25c6c4a7cd4c4f103eb60316232c349443b45cd21f601c0b5cbaf73b42e3)
- query_params.check_not_present

<a id="canonical-c54636e470c24448f9738d61f4b2c77ea7c9403b1a86931785efed68c0c16016"></a>

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

<a id="canonical-bcef20b01bad18e1582b83657f02175e2d8694134cc11be47145d3af0e8a8fd6"></a>

## Direct properties — query_params.check_not_present / 83f92581794f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b95ef9bf6741c2c956f591dda69a380317c5f24e246ed9b29e454669a8fbe18a"></a>

## Next pages — query_params.check_not_present / 83f92581794f / 4

- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-040e25c6c4a7cd4c4f103eb60316232c349443b45cd21f601c0b5cbaf73b42e3)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-bc1dfc2276984a71923214d9f3b50b2200cf5b2a02e1b019852c0caa7eb882a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99bc6d9b488e2e2dfb06c4b7149f1e152804edc497b0829d2cc2494125228dab"></a>

## query_params.check_present — query_params.check_present / b820137b049b / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-040e25c6c4a7cd4c4f103eb60316232c349443b45cd21f601c0b5cbaf73b42e3)
- query_params.check_present

<a id="canonical-42d9f7fb2bdc3b01fd83f3427e644069866e02a5a5842ebdebbbf77873753603"></a>

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

<a id="canonical-ebe76c6245c97807fb194e8da94db07fda70d20608a9c168662788c8489370c8"></a>

## Direct properties — query_params.check_present / b820137b049b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cbdf57a26825f8fd6f8514dc74d96e04753015c7138f599b08a95fa79fd8b1de"></a>

## Next pages — query_params.check_present / b820137b049b / 4

- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-040e25c6c4a7cd4c4f103eb60316232c349443b45cd21f601c0b5cbaf73b42e3)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-df07cef78699c3eb7546c3592d0b45e5d54b131ab060d10034a5d79232e3ccc2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41dc9662933f469463a58078c4cc2f62f52a9d23eeed642792383a3c0235b5ec"></a>

## query_params.item — query_params.item / 98ef5fed9dc6 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-040e25c6c4a7cd4c4f103eb60316232c349443b45cd21f601c0b5cbaf73b42e3)
- query_params.item

<a id="canonical-3c45663aff7da99fbe6894ef7a402d5affddb55e5f2c7722299f1ccf92e40d67"></a>

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

<a id="canonical-688621830c949fb03f23aa03dfced2a983efde5aa4f6165e5a571d62f738f5e2"></a>

## Direct properties — query_params.item / 98ef5fed9dc6 / 3

<a id="canonical-2f605ba6bd5ee209fafc27ae5f24126ce4586487a6478c7f58a86deeae22f958"></a>

<a id="canonical-3b14f61d088d1b5f5250001cf251fad3da217605421e4a71ce0092dee4374923"></a>

## exact_values property — query_params.item / 98ef5fed9dc6 / 4

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

<a id="canonical-5decf1709b45fb79e9705dd3deba8cf8fd75981d02dcec298fe02fc67ddc1500"></a>

<a id="canonical-14a67b86035ccb111259fc0efa033938de23ee8114c0a3df67ef87250704b44e"></a>

## regex_values property — query_params.item / 98ef5fed9dc6 / 5

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

<a id="canonical-30e292a8c07688c7539ea7a498c4ae0b9b0ff3d6403966d10056ac0b410826b7"></a>

<a id="canonical-0172f1e6ec68f53ad8d1500493bccdcae7868b73091b8a813d063695abd48496"></a>

## transformers property — query_params.item / 98ef5fed9dc6 / 6

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

<a id="canonical-37b2c0811f098e9272f16bfad9c612995f8660ce14b2744d69c847530364b352"></a>

## Next pages — query_params.item / 98ef5fed9dc6 / 7

- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-040e25c6c4a7cd4c4f103eb60316232c349443b45cd21f601c0b5cbaf73b42e3)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95cb7c2f246705b7c10fad85f76ef2cef5442673a7f3665f6fc06a7a3e68ca38"></a>

## request_constraints — request_constraints / 81a328d895fb / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- request_constraints

<a id="canonical-19cdce4084ec9e481c5aee84aaa5b7c1ce661ee7646de0cfaf90c9518d616908"></a>

Type: `"single"`. Computed.

Configuration parameter for request constraints.

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

<a id="canonical-4f10bc60a800ff129460bb90154b315587e66adaa494619ea4fe6a382534bc89"></a>

## Direct properties — request_constraints / 81a328d895fb / 3

<a id="canonical-9fb38ae8576c8b902f150d75a8acfb9fd577beedeeea4fabbe9f6221ee796d4a"></a>

<a id="canonical-d133771c6ec779e6169046376a4e30df438960b931cce8dae40be96696ec5965"></a>

## max_cookie_count_exceeds property — request_constraints / 81a328d895fb / 4

Type: `"number"`. Computed.

Match on the Count for all Cookies that exceed this value. Exclusive with
\[max\_cookie\_count\_none\]

Upstream description:

Exclusive with \[max\_cookie\_count\_none\]

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

- [max_cookie_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-fa7224aa823f1ddb3a1d5f43e17e229eb1eda6d7a8454cc87a07e5a6060e91fe): complete subsection reference.

<a id="canonical-73b5d0584160d52e531082eefd32bf2b7d78ae8b3ae9b5e20cc2e31fc3dc9607"></a>

<a id="canonical-44ea1f75430c36a656d743b2cee985b184c4cd7601d08d6cbc03f4511da99255"></a>

## max_cookie_key_size_exceeds property — request_constraints / 81a328d895fb / 5

Type: `"number"`. Computed.

Exclusive with \[max\_cookie\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_key\_size\_none\]

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

- [max_cookie_key_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-003e8feee50461cc303fbf5bcaf7d69d6a8b6b0e9f3e5bf1a48c71a57cdcb4fc): complete subsection reference.

<a id="canonical-13425e16f3cd32e029497732af6844291ec8e08e1eb914a925d7888dd2497057"></a>

<a id="canonical-24404ccf79f9181260fb5889528e9887fd6d12e7266e35873015c742f012defe"></a>

## max_cookie_value_size_exceeds property — request_constraints / 81a328d895fb / 6

Type: `"number"`. Computed.

Exclusive with \[max\_cookie\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_value\_size\_none\]

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

- [max_cookie_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-c92a9c91760ccb690294726080f00718f48417163faf429ef99049a685def3e2): complete subsection reference.

<a id="canonical-45e3c7cb468594e1282141c0535dbe8933f4f5462aa54af019c1b1d4c6be54f1"></a>

<a id="canonical-a15ed484cd11768425d9662bd09b40817fc06682072069d2171824140b671e18"></a>

## max_header_count_exceeds property — request_constraints / 81a328d895fb / 7

Type: `"number"`. Computed.

Match on the Count for all Headers that exceed this value. Exclusive with
\[max\_header\_count\_none\]

Upstream description:

Exclusive with \[max\_header\_count\_none\]

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

- [max_header_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-5b9f0a2e919afd464b36f4555f16df71cb3c872db2d03a648db4d0f69a182933): complete subsection reference.

<a id="canonical-09964672a0c2420d458cd8c052d151daa4c1acdd79ebc662516f319b44122bb4"></a>

<a id="canonical-69edba2269c12d7bff218278f1df2c100f6389d1a929abf66571ce032df162e1"></a>

## max_header_key_size_exceeds property — request_constraints / 81a328d895fb / 8

Type: `"number"`. Computed.

Exclusive with \[max\_header\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_key\_size\_none\]

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

- [max_header_key_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-f7cc955d720fd027ed396eeb3032eb32a539f7a67c926db4d5f89912491abb07): complete subsection reference.

<a id="canonical-c810382529612d11205d678ae83a45a470dae42b36a2cccfd456bde22f03b29a"></a>

<a id="canonical-7dd1c877e7fe95ea4fe6b17d3b288db95eb59c245fda34a02118b18128f0c9f8"></a>

## max_header_value_size_exceeds property — request_constraints / 81a328d895fb / 9

Type: `"number"`. Computed.

Exclusive with \[max\_header\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_value\_size\_none\]

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

- [max_header_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-08b85dc29f6937a5c060f40ee8e2e2414ee505c0f5eec4e475215b90cfaaa333): complete subsection reference.

<a id="canonical-ac905d5de84f54b0d6883f9bb78215eac55464a2be12cb6b76c2d166b2514aae"></a>

<a id="canonical-d036a2136ab6b1075039f0679d4f63b060ac13117a24769c0bc564ed452a8681"></a>

## max_parameter_count_exceeds property — request_constraints / 81a328d895fb / 10

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_count\_none\].

Upstream description:

Exclusive with \[max\_parameter\_count\_none\]

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

- [max_parameter_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-cfa2ac30367989f7c5a3e6ac57fcf77632f306e730c3ad0af40e26432f1e96a9): complete subsection reference.

<a id="canonical-8566e4a7d5756e611a8092222e8e14e565c4ffd9610e739b613ed7a8ec9db56b"></a>

<a id="canonical-4fb4739b468eb1314c6a5a77674ea82e1709a1f68c0ea00ba72dceb6150b1858"></a>

## max_parameter_name_size_exceeds property — request_constraints / 81a328d895fb / 11

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_name\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_name\_size\_none\]

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

- [max_parameter_name_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-05d53bd708f87deec908e54a231494db69ffc3e319e29da9b5a6985b40b5b55d): complete subsection reference.

<a id="canonical-62c234d54a3885120045d0087bb7515421719197c1e616d2fe1b36836d63f930"></a>

<a id="canonical-50aa366d8f798a6557e26da3b504e38f27882cf51d6be4c2b52450c6c60ab64c"></a>

## max_parameter_value_size_exceeds property — request_constraints / 81a328d895fb / 12

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_value\_size\_none\]

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

- [max_parameter_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-769e7fc956e874ec87e695028d47a6179283c130ec8f1d815048d2f0f4d2067a): complete subsection reference.

<a id="canonical-955a5b60d99906a4a339a72e7d2ae550b3a24bf27862315ffeaa7dd498ff8859"></a>

<a id="canonical-081b1e7f9b48c4cbf3a47e4ba34619d5c6c013a143e5f8296372fae5c37498bf"></a>

## max_query_size_exceeds property — request_constraints / 81a328d895fb / 13

Type: `"number"`. Computed.

Match on the URL Query Size that exceed this value. Exclusive with \[max\_query\_size\_none\]

Upstream description:

Exclusive with \[max\_query\_size\_none\]

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

- [max_query_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-6947e4eeeabc89c3abe7d945251f32bd0e0dd0c8ebf5801db75376098c1a3580): complete subsection reference.

<a id="canonical-525b735b57ae97eb091e22ad2e82db2e9902393e0c6b0f2737daddd667f86788"></a>

<a id="canonical-a0e987b869f7477d6ff52b5e7babb57562bdb2d1d83e9e8aed861a7e40ddcfbb"></a>

## max_request_line_size_exceeds property — request_constraints / 81a328d895fb / 14

Type: `"number"`. Computed.

Exclusive with \[max\_request\_line\_size\_none\].

Upstream description:

Exclusive with \[max\_request\_line\_size\_none\]

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

- [max_request_line_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-43feabb70f8943c86d4389403361181755bd2e3e7ab6a1dd52e9b1d90a36cd5c): complete subsection reference.

<a id="canonical-2b0ad8f0da212e754204168b5f91cec253d66befd7808b08af2723fde005f337"></a>

<a id="canonical-f3bb553c71be3f7a02afcde07e14b3a9500ef29b176a9cea008dcee11516cc80"></a>

## max_request_size_exceeds property — request_constraints / 81a328d895fb / 15

Type: `"number"`. Computed.

Match on the Request Size that exceed this value. Exclusive with \[max\_request\_size\_none\]

Upstream description:

Exclusive with \[max\_request\_size\_none\]

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

- [max_request_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-d3e2b3abcc790ebbee99e1a3c14166ffdd557518fcda21ca7e20c3ff1ae359bb): complete subsection reference.

<a id="canonical-97a09d1ff6c9f0e998aae239212ede67be44d8e6bc653ceda3b0a3c9aa30b352"></a>

<a id="canonical-0d6b5d47480a052b00cbce32f04fea319c96eab397201b2bb697e88bb4f65b2b"></a>

## max_url_size_exceeds property — request_constraints / 81a328d895fb / 16

Type: `"number"`. Computed.

Match on the URL Size that exceed this value. Exclusive with \[max\_url\_size\_none\]

Upstream description:

Exclusive with \[max\_url\_size\_none\]

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

- [max_url_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-a459a6174b52f992a34fa8083a906308feaf13ab8863c98544275bcd45f37eb3): complete subsection reference.

<a id="canonical-5eecec41b4b48cdf6b19e3fc1a58c32fed464ecfa2794ee4cdfde6d6c51b672e"></a>

## Next pages — request_constraints / 81a328d895fb / 17

- [request_constraints.max_cookie_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-fa7224aa823f1ddb3a1d5f43e17e229eb1eda6d7a8454cc87a07e5a6060e91fe)
- [request_constraints.max_cookie_key_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-003e8feee50461cc303fbf5bcaf7d69d6a8b6b0e9f3e5bf1a48c71a57cdcb4fc)
- [request_constraints.max_cookie_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-c92a9c91760ccb690294726080f00718f48417163faf429ef99049a685def3e2)
- [request_constraints.max_header_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-5b9f0a2e919afd464b36f4555f16df71cb3c872db2d03a648db4d0f69a182933)
- [request_constraints.max_header_key_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-f7cc955d720fd027ed396eeb3032eb32a539f7a67c926db4d5f89912491abb07)
- [request_constraints.max_header_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-08b85dc29f6937a5c060f40ee8e2e2414ee505c0f5eec4e475215b90cfaaa333)
- [request_constraints.max_parameter_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-cfa2ac30367989f7c5a3e6ac57fcf77632f306e730c3ad0af40e26432f1e96a9)
- [request_constraints.max_parameter_name_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-05d53bd708f87deec908e54a231494db69ffc3e319e29da9b5a6985b40b5b55d)
- [request_constraints.max_parameter_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-769e7fc956e874ec87e695028d47a6179283c130ec8f1d815048d2f0f4d2067a)
- [request_constraints.max_query_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-6947e4eeeabc89c3abe7d945251f32bd0e0dd0c8ebf5801db75376098c1a3580)
- [request_constraints.max_request_line_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-43feabb70f8943c86d4389403361181755bd2e3e7ab6a1dd52e9b1d90a36cd5c)
- [request_constraints.max_request_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-d3e2b3abcc790ebbee99e1a3c14166ffdd557518fcda21ca7e20c3ff1ae359bb)
- [request_constraints.max_url_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-a459a6174b52f992a34fa8083a906308feaf13ab8863c98544275bcd45f37eb3)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-fa7224aa823f1ddb3a1d5f43e17e229eb1eda6d7a8454cc87a07e5a6060e91fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8008133f89a4b886e56cd9f7fb2920793d565bc78206c08003c7906012eae8c"></a>

## request_constraints.max_cookie_count_none — request_constraints.max_cookie_count_none / f9e04250839f / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- request_constraints.max_cookie_count_none

<a id="canonical-4c5c8290d478f861c6e9274d8341bd424741e31e074ccc9c0a6d8af9a3227411"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-9abca969b73a1c69bc4eb8b2eea0604ef786e554414cb950b76d45acc460b73c"></a>

## Direct properties — request_constraints.max_cookie_count_none / f9e04250839f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b5bca2f5e7908dd47e13585dfa014c69752a0c8be04fbb24143fe024d08a74a6"></a>

## Next pages — request_constraints.max_cookie_count_none / f9e04250839f / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-003e8feee50461cc303fbf5bcaf7d69d6a8b6b0e9f3e5bf1a48c71a57cdcb4fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62f87a1c40d613c847199b7f4b9edfe0caaeb4d736ee370d546970a73a920201"></a>

## request_constraints.max_cookie_key_size_none — request_constraints.max_cookie_key_size_none / 1e81a2a4cd4d / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- request_constraints.max_cookie_key_size_none

<a id="canonical-4fffe56ac3c8ebff2511e3edf40f5c43e63d1c734e6b3996a5c6257e4d60f6b7"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-adad9fb685d99b61d16d36ba8c8d8f737f9e5f8c91297672c81cd0068c44774c"></a>

## Direct properties — request_constraints.max_cookie_key_size_none / 1e81a2a4cd4d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e8986172327f47793cdf67a69c5b31f925fdaead81333b1344ea372531b03317"></a>

## Next pages — request_constraints.max_cookie_key_size_none / 1e81a2a4cd4d / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-c92a9c91760ccb690294726080f00718f48417163faf429ef99049a685def3e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-792475fca7d3a118914de9465905dc79741cad70ad5bdd9d7b7bc718a6779e1d"></a>

## request_constraints.max_cookie_value_size_none — request_constraints.max_cookie_value_size_none / 73558b9f4fc2 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- request_constraints.max_cookie_value_size_none

<a id="canonical-24ed5c31551833027a28e37bc96ad20f29609f3bd9b544e061b4dcab89b91dbe"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-89c1dc46a9c41d9e094a1fbc4f0d20716637f753a428e81410ee099d63c37f8f"></a>

## Direct properties — request_constraints.max_cookie_value_size_none / 73558b9f4fc2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4d7e1221e18a72b1eb43ca15e980b23ad1d90f6ace7bffd12f7c212a5355ab24"></a>

## Next pages — request_constraints.max_cookie_value_size_none / 73558b9f4fc2 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-5b9f0a2e919afd464b36f4555f16df71cb3c872db2d03a648db4d0f69a182933"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d696b72e021556e7f184b1b92ad9c956e12847afa7348c9349e3e03fdcb2e83b"></a>

## request_constraints.max_header_count_none — request_constraints.max_header_count_none / bf3fcaa39999 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- request_constraints.max_header_count_none

<a id="canonical-4582e8d8582ea6dcc6f90fd72332e04da378adbbab77edc7b46a4c6ce5e4c2f0"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-50826cf155a436a794f616e86e683de4d7d4b1abb444f19e1a9067b68117fbe2"></a>

## Direct properties — request_constraints.max_header_count_none / bf3fcaa39999 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cdba976adfdee24c21bb80b77dfa54a27e746338bef2205d03f60f2d75b8ee97"></a>

## Next pages — request_constraints.max_header_count_none / bf3fcaa39999 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-f7cc955d720fd027ed396eeb3032eb32a539f7a67c926db4d5f89912491abb07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-696971e280e453814ef53437b46ad9c6750bc10619299ddb21f18f8e825f91d2"></a>

## request_constraints.max_header_key_size_none — request_constraints.max_header_key_size_none / 1d284eae51a5 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- request_constraints.max_header_key_size_none

<a id="canonical-11ad84ec4196fc7b565f85d763cfa28185735048de0c291dbdd7e6265c555f42"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-541b016e1e5412e1852d901b7fdbabddd70e5b31739f3d375a5051d4a0342c18"></a>

## Direct properties — request_constraints.max_header_key_size_none / 1d284eae51a5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-02eef4929e293786a24782fca60169f5cd9e1d245c838a39ec5d5353f7c88613"></a>

## Next pages — request_constraints.max_header_key_size_none / 1d284eae51a5 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-08b85dc29f6937a5c060f40ee8e2e2414ee505c0f5eec4e475215b90cfaaa333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e42d88c62cb57e244f32f14ce0e63c23d36f879f775190820b3f673a0d6cd94d"></a>

## request_constraints.max_header_value_size_none — request_constraints.max_header_value_size_none / b065f29fdab7 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- request_constraints.max_header_value_size_none

<a id="canonical-e8aaee10c6b78b111b952d4571899cbb07d4ac30685f2b3e3c91625871d00a50"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-e682f374f8023cfa573bf1905c6a46239bb1e5f987875c3d65e02265e37a634f"></a>

## Direct properties — request_constraints.max_header_value_size_none / b065f29fdab7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2d629153b5e2bfdbab6fc862b33ddc87c7f432e48c3ff0a4e39630c7f8113355"></a>

## Next pages — request_constraints.max_header_value_size_none / b065f29fdab7 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-cfa2ac30367989f7c5a3e6ac57fcf77632f306e730c3ad0af40e26432f1e96a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecb49da90bd6ea306211df6b2f090f7ffeacad1e32d064be125e4ac56942954a"></a>

## request_constraints.max_parameter_count_none — request_constraints.max_parameter_count_none / f83deac6b1ff / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- request_constraints.max_parameter_count_none

<a id="canonical-70e74487b6feb26f29cb705a301c2079cbba978a20934b647b09ef798a21aaa6"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-b1662f70dcecebf319145cad79f2a8fa1257160f5b8a1d1c81c32d6fdfa8800f"></a>

## Direct properties — request_constraints.max_parameter_count_none / f83deac6b1ff / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2ca691d292e7597e8cc52881406a59c73391dff5e417cd1459e216bd3d5041f"></a>

## Next pages — request_constraints.max_parameter_count_none / f83deac6b1ff / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-05d53bd708f87deec908e54a231494db69ffc3e319e29da9b5a6985b40b5b55d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f92a414c00ba2ade9d7bfbdedfa72d58968c9f8c4b4cca7f2695459fdb2229d8"></a>

## request_constraints.max_parameter_name_size_none — request_constraints.max_parameter_name_size_none / 51b3d264b664 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- request_constraints.max_parameter_name_size_none

<a id="canonical-76e894999384d425ec37951ad2709ff510f193949ec6a3bea727e113915700c3"></a>

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

<a id="canonical-889f4056ffa0d71c0743edff15158f331dbbf102ffdaf4d64c36aebf4efea552"></a>

## Direct properties — request_constraints.max_parameter_name_size_none / 51b3d264b664 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c4a80c00a84cc6d7bdf062842661218a28d18cb8442dc85e574b23bdffb3024e"></a>

## Next pages — request_constraints.max_parameter_name_size_none / 51b3d264b664 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-769e7fc956e874ec87e695028d47a6179283c130ec8f1d815048d2f0f4d2067a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e82a88542c8be7114cab186013456d2d284a203891bd37df4686631e7205127"></a>

## request_constraints.max_parameter_value_size_none — request_constraints.max_parameter_value_size_none / 459f70f3ee9a / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- request_constraints.max_parameter_value_size_none

<a id="canonical-82fb4e4c80c9738af9561477fe2c1ca0223f9497c5f12c0d3ce99042333a1ac3"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-588c4075375f896bbf73a13767075bacef5d52f3cddeff61df7f566321c47a5b"></a>

## Direct properties — request_constraints.max_parameter_value_size_none / 459f70f3ee9a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9f1c856e88d495ff458850210a3cb04bf40b12f32fb102b5bf71d0c8f32ee4a8"></a>

## Next pages — request_constraints.max_parameter_value_size_none / 459f70f3ee9a / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-6947e4eeeabc89c3abe7d945251f32bd0e0dd0c8ebf5801db75376098c1a3580"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c938d523201a5c981c9191126443cef2f48596c844bb9ece42fdeed26b361614"></a>

## request_constraints.max_query_size_none — request_constraints.max_query_size_none / 97441047d872 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- request_constraints.max_query_size_none

<a id="canonical-7604655ac5aaded637ce980adb7e4a28b286559a891b7aaae9f88711ad5e4132"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-740612577eb2856f4359392aa490c0957fce29464c480cbc48a35951d851b9a2"></a>

## Direct properties — request_constraints.max_query_size_none / 97441047d872 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-be1265df66060c7487d38471c630ca8a97f4560ab5b8b09e19f11be6256811c4"></a>

## Next pages — request_constraints.max_query_size_none / 97441047d872 / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-43feabb70f8943c86d4389403361181755bd2e3e7ab6a1dd52e9b1d90a36cd5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff395e8576d3008f8e88527517971a5ad7594da7bfaf48cbf8498506d2874566"></a>

## request_constraints.max_request_line_size_none — request_constraints.max_request_line_size_none / 7281d782462b / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- request_constraints.max_request_line_size_none

<a id="canonical-3bfe1343298b3ea7650058a8934cff5c2c3e7eadca2ba352a247b86d74a7ad4a"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-de4037da0e91393ab70d3de3004fd47c9f62d86f18f6c7295d527c1d97c60701"></a>

## Direct properties — request_constraints.max_request_line_size_none / 7281d782462b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dfa17fefda1c83aa7ae2d8b43da830fd034f9d9518a600036a26567fc9bb96e9"></a>

## Next pages — request_constraints.max_request_line_size_none / 7281d782462b / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-d3e2b3abcc790ebbee99e1a3c14166ffdd557518fcda21ca7e20c3ff1ae359bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fd7d13b5a9b22bcb653c114295827d759047d738334930d896baa4ec67f444f"></a>

## request_constraints.max_request_size_none — request_constraints.max_request_size_none / 43ad87946b5d / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- request_constraints.max_request_size_none

<a id="canonical-df35bf462fa8d275ab6e1e0b7fc1a249f3918457ebd16c82032bd6f63b5da096"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-d5542d7aff922cfd86f44e4c9d3d3e5703b6aa3f0568a18732ee8f781abfdfd8"></a>

## Direct properties — request_constraints.max_request_size_none / 43ad87946b5d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0845d4e275a462244d23875becb13ab2accb82819997f3f105a2ef55e9cb51d6"></a>

## Next pages — request_constraints.max_request_size_none / 43ad87946b5d / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-a459a6174b52f992a34fa8083a906308feaf13ab8863c98544275bcd45f37eb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96fc2f539ce7a8d9e7b2e1cd9e31519f9d7dcfd44f6d3af7614d23125cb36890"></a>

## request_constraints.max_url_size_none — request_constraints.max_url_size_none / 0913f650967b / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- request_constraints.max_url_size_none

<a id="canonical-adfa33dc9c306e9a27f26e0a2e6f870d3f2999f79c81e788ca971f32ece4fe46"></a>

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

<a id="canonical-10b5524babbb897112e3819383b19215d09f5a8f9c85e5a815033a7b330d1e0e"></a>

## Direct properties — request_constraints.max_url_size_none / 0913f650967b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-337f8352e8b88557aa8260db432db0e8db3f9ef2d01ffd0cb5b6dccdee1f379a"></a>

## Next pages — request_constraints.max_url_size_none / 0913f650967b / 4

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-2b564bee935faf2fa4d76d6eff5e69eb7f4255797cca7309cb9b317da3508dc4)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eda5b70551c7539bc5bacd369df09d0db9ae10f2fa7924885c7936ed87682849"></a>

## segment_policy — segment_policy / c32a4ae5209e / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- segment_policy

<a id="canonical-8d2067c794f0fb372809242bc002bd7197e79d94368e3cbd68cb115e886ed23b"></a>

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

<a id="canonical-2778160f45d45816df0fb588b930dacf8f6abb5fe0aa52ff74ab491915d891a5"></a>

## Direct properties — segment_policy / c32a4ae5209e / 3

- [dst_any](data-sources--service_policy_rule--reference--group-002.md#canonical-f604397ad548b90ebf68957dba5efcbd0b34712b0409b19127aeac6ea98a0f2f): complete subsection reference.

- [dst_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-236812428e5b84250d5ea750896586aa9e2b16b95c21327c864401d90bea2805): complete subsection reference.

- [intra_segment](data-sources--service_policy_rule--reference--group-002.md#canonical-d463fac9b40c77557f162ed9deec5f8502e4fa45eb35b68de3a8091a7451e07c): complete subsection reference.

- [src_any](data-sources--service_policy_rule--reference--group-002.md#canonical-a9a829f02c0466603dd32ebec8cb19202e21f37fea50f2e6583babf1b36afe25): complete subsection reference.

- [src_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-6805aa9c33ecd624235a9c4919fbea61f9dfe90a4b297e1f5b9c9f43d9c8f4db): complete subsection reference.

<a id="canonical-db11645c276363d710ddd65480bbc4592e4bbb556c4f4807f0adb811788e33f9"></a>

## Next pages — segment_policy / c32a4ae5209e / 4

- [segment_policy.dst_any](data-sources--service_policy_rule--reference--group-002.md#canonical-f604397ad548b90ebf68957dba5efcbd0b34712b0409b19127aeac6ea98a0f2f)
- [segment_policy.dst_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-236812428e5b84250d5ea750896586aa9e2b16b95c21327c864401d90bea2805)
- [segment_policy.intra_segment](data-sources--service_policy_rule--reference--group-002.md#canonical-d463fac9b40c77557f162ed9deec5f8502e4fa45eb35b68de3a8091a7451e07c)
- [segment_policy.src_any](data-sources--service_policy_rule--reference--group-002.md#canonical-a9a829f02c0466603dd32ebec8cb19202e21f37fea50f2e6583babf1b36afe25)
- [segment_policy.src_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-6805aa9c33ecd624235a9c4919fbea61f9dfe90a4b297e1f5b9c9f43d9c8f4db)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-f604397ad548b90ebf68957dba5efcbd0b34712b0409b19127aeac6ea98a0f2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e87171dac57bd2bb1d610e83b2464d73f30f100f02457bed1f24f9924d03c04"></a>

## segment_policy.dst_any — segment_policy.dst_any / e8e531092437 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610)
- segment_policy.dst_any

<a id="canonical-f738d21847f22a4f9473da78950744d65a1a24132f601e9ee7039c16ebca6eca"></a>

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

<a id="canonical-6bc39d014d7830bd6ef05aa4954923dc21f3ca688761447c2c83904a87233e91"></a>

## Direct properties — segment_policy.dst_any / e8e531092437 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-42e9f4dbf593860ff911be2e64d26b51dc2045204cf15c92a67139b3cc4e6757"></a>

## Next pages — segment_policy.dst_any / e8e531092437 / 4

- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-236812428e5b84250d5ea750896586aa9e2b16b95c21327c864401d90bea2805"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-330466995faf262dab4f6da585ab16c9900aae01ead3e86a20f5a4f6ac3ca0a6"></a>

## segment_policy.dst_segments — segment_policy.dst_segments / 9f7184a5eba5 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610)
- segment_policy.dst_segments

<a id="canonical-3266bf29bb4c0787465d1e2ebccada5d79ea39ac688650129674f8272cb4da6f"></a>

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

<a id="canonical-1e91dbc1f51d897daee3bf8548ac18dc61160155934baeb25846b0f641ceb83b"></a>

## Direct properties — segment_policy.dst_segments / 9f7184a5eba5 / 3

- [segments](data-sources--service_policy_rule--reference--group-002.md#canonical-3e562a37652458df9115b097468cb464b5baf282d8a1538ac8d3c13a6a337c7a): complete subsection reference.

<a id="canonical-ab1f9fa0afec64f7594ce2287528fcd3132e17127b171eccb6f682e735868527"></a>

## Next pages — segment_policy.dst_segments / 9f7184a5eba5 / 4

- [segment_policy.dst_segments.segments](data-sources--service_policy_rule--reference--group-002.md#canonical-3e562a37652458df9115b097468cb464b5baf282d8a1538ac8d3c13a6a337c7a)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-3e562a37652458df9115b097468cb464b5baf282d8a1538ac8d3c13a6a337c7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0ce748aba85a15e074c564cfd2aea286a474a712af2454e3579b63d57cf984a"></a>

## segment_policy.dst_segments.segments — segment_policy.dst_segments.segments / c0e0f7725b7e / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610)
- [segment_policy.dst_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-236812428e5b84250d5ea750896586aa9e2b16b95c21327c864401d90bea2805)
- segment_policy.dst_segments.segments

<a id="canonical-eeba5610d8b3350152c3a4b1ca2f3d118883a2bdb78309bfd30971d41e7ef7b6"></a>

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

<a id="canonical-dee9f06d1a41c11706b958581da37a5a2b509eb71452bb2d6f47235ccbeb6d9c"></a>

## Direct properties — segment_policy.dst_segments.segments / c0e0f7725b7e / 3

<a id="canonical-c0fb29a12b2ab403e98ec6f2257bd5fbd9f4815cb27b06b996e829f7b26053df"></a>

<a id="canonical-dc849ad07c070b3e51f72f763562c5f66bafe5a7a9a87e359a26e4540d14141c"></a>

## name property — segment_policy.dst_segments.segments / c0e0f7725b7e / 4

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

<a id="canonical-20f27ac642e5778a96d471c9a7a277f1b7b44197897ecd5a62e0c82a736eb7f2"></a>

<a id="canonical-1c6548190c96fdb42169cb49b2ddd7a258ea31bfabacf71d40d899aca353b21e"></a>

## namespace property — segment_policy.dst_segments.segments / c0e0f7725b7e / 5

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

<a id="canonical-955c63f9b19e2637316255b522a969a6cdf8d51a5fd577c596386b30156daf98"></a>

<a id="canonical-c94264abb4df8dbf3b29573f8fe800fa8d8079d3460a089f25daf1335c18ad03"></a>

## tenant property — segment_policy.dst_segments.segments / c0e0f7725b7e / 6

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

<a id="canonical-3b4d13e77fba71a2c6c6f286b23e99fb549ea5a36b8a14b9c96a9e2d2601a605"></a>

## Next pages — segment_policy.dst_segments.segments / c0e0f7725b7e / 7

- [segment_policy.dst_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-236812428e5b84250d5ea750896586aa9e2b16b95c21327c864401d90bea2805)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-d463fac9b40c77557f162ed9deec5f8502e4fa45eb35b68de3a8091a7451e07c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22e1c779abf526c9fe70cfb0d052ecab036dda01166ccd75e7f48bb13fe6423c"></a>

## segment_policy.intra_segment — segment_policy.intra_segment / 28ca03984dde / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610)
- segment_policy.intra_segment

<a id="canonical-d691d7584225f3718c19bc4c3af89da188b208da93dac6b295c5ae4e8dd1fd18"></a>

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

<a id="canonical-3ff47a0dd2de650f19e5a40b5ee26aa7d303ab32ecb065e45349b8a571879288"></a>

## Direct properties — segment_policy.intra_segment / 28ca03984dde / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e7dc4b365dca80813caa8a696b527ad9b0b279c7464957d122d9246afbe22edf"></a>

## Next pages — segment_policy.intra_segment / 28ca03984dde / 4

- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-a9a829f02c0466603dd32ebec8cb19202e21f37fea50f2e6583babf1b36afe25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01fd05880688d7145bdc9b2b47e548c92f50917ad31bdb88315972550980cabf"></a>

## segment_policy.src_any — segment_policy.src_any / 023cedbe694c / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610)
- segment_policy.src_any

<a id="canonical-1a7a905bffddd3083ae6f5aca9c7891b46b4ad21d672b07d9157898878743fe6"></a>

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

<a id="canonical-898399c0b9909e2a43dfff131c44e45bb707ab1b9dcc62b21a40c1a9cd3536b6"></a>

## Direct properties — segment_policy.src_any / 023cedbe694c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe177786a37f6b2072e84946b049c647a4661925b684fcfdcb3787def051f4d4"></a>

## Next pages — segment_policy.src_any / 023cedbe694c / 4

- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-6805aa9c33ecd624235a9c4919fbea61f9dfe90a4b297e1f5b9c9f43d9c8f4db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7d1991eb099674ab093c05eb3deb168da6a022980eb2a017270e90aab4ca9be"></a>

## segment_policy.src_segments — segment_policy.src_segments / 71279c56ef7b / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610)
- segment_policy.src_segments

<a id="canonical-7da2a0d339edabaf6811c80325bbb654d797964223e30d53ca16584b98d8e600"></a>

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

<a id="canonical-249414f3b7cab02273a215e430a82ab1500e97ba3da3786087a4a0ea7ace0db7"></a>

## Direct properties — segment_policy.src_segments / 71279c56ef7b / 3

- [segments](data-sources--service_policy_rule--reference--group-002.md#canonical-554d77d26674e62f87c951c8377efa2ec00b2992e8bf00e37dd36c401e9f932a): complete subsection reference.

<a id="canonical-f40477f53b82d9fa79fd139555e72d8b7de9d64f837b53ab8865efcdf9302444"></a>

## Next pages — segment_policy.src_segments / 71279c56ef7b / 4

- [segment_policy.src_segments.segments](data-sources--service_policy_rule--reference--group-002.md#canonical-554d77d26674e62f87c951c8377efa2ec00b2992e8bf00e37dd36c401e9f932a)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-554d77d26674e62f87c951c8377efa2ec00b2992e8bf00e37dd36c401e9f932a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33bf7da5bde3bc6135d3b35661ddf9ed0e8f54264d1fdedbd766f96c8c824251"></a>

## segment_policy.src_segments.segments — segment_policy.src_segments.segments / 1ae42eb2b82c / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-d48fcc8707b40aa83e13692ccc2861e7f7dafc659b42967be846043ef06fc610)
- [segment_policy.src_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-6805aa9c33ecd624235a9c4919fbea61f9dfe90a4b297e1f5b9c9f43d9c8f4db)
- segment_policy.src_segments.segments

<a id="canonical-9a8bef16434e5720e50a05cefed43a3ad5fcedf0660925329159121d8941544b"></a>

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

<a id="canonical-f7967caf1e4b63caa90084f532b4e480303bc18fa79a638418994a09748d146e"></a>

## Direct properties — segment_policy.src_segments.segments / 1ae42eb2b82c / 3

<a id="canonical-338fea2490dd983c3f135abadcad4096f260055f4f120e8704f27db13f51e480"></a>

<a id="canonical-745465e7cc95aaadb5c311edcdce2eaf20864e60d230c4b6806bff9964121393"></a>

## name property — segment_policy.src_segments.segments / 1ae42eb2b82c / 4

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

<a id="canonical-e86bd035467aa796d0fc7bdf5ca8fb74c90be0c80df598be07751ca518818a6d"></a>

<a id="canonical-8ec90ff681f2d3225db690779d17c2f692a44a920017428e8bea903c6c702d5e"></a>

## namespace property — segment_policy.src_segments.segments / 1ae42eb2b82c / 5

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

<a id="canonical-908e58bf02ce5928f3c98b0b9aa523b4fae0c68fcdc087a74dd24d8c17ae3a32"></a>

<a id="canonical-05922465846ca06a25d8973e8ff690233c17f3274df09990b00a471a6a5a0c68"></a>

## tenant property — segment_policy.src_segments.segments / 1ae42eb2b82c / 6

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

<a id="canonical-bcd3845dd55d2dede5b984646fae6d47b2c187785269dbcd39c3ce1443c95dcc"></a>

## Next pages — segment_policy.src_segments.segments / 1ae42eb2b82c / 7

- [segment_policy.src_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-6805aa9c33ecd624235a9c4919fbea61f9dfe90a4b297e1f5b9c9f43d9c8f4db)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-1fc6b6f8ae22871967881106ea255eba23defd48f77829156aa2fdf56202e7ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bbac14bb4ae1b9b858eff5bccaf21f83f954f177128eede87b7a073b18f5614"></a>

## tls_fingerprint_matcher — tls_fingerprint_matcher / 79fc764ce1d1 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- tls_fingerprint_matcher

<a id="canonical-e44b775e4590083f8137f4b607de5d213b8808fc5179cdcd48fa9c663d5757c4"></a>

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

<a id="canonical-4f1d59cd1e8a88ce5911881d58b415cd51e91b6c618a504aef5df518a8c363d4"></a>

## Direct properties — tls_fingerprint_matcher / 79fc764ce1d1 / 3

<a id="canonical-61d227bf85dccc3d5524ab6730e0f6e041ddc550221193aa3262e5e3257a3af8"></a>

<a id="canonical-425d6834ae5aaf50c223044813c5caf3bcf015f280b5aff010f67a2242455994"></a>

## classes property — tls_fingerprint_matcher / 79fc764ce1d1 / 4

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

<a id="canonical-15c34b3a37360974a324c488a89a53e199ec1e0a77c02195a2afb9dfbb2cee4b"></a>

<a id="canonical-90ef64f548114cf20da8162882b69f6130eaa828224931bfb99bfce963470e02"></a>

## exact_values property — tls_fingerprint_matcher / 79fc764ce1d1 / 5

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

<a id="canonical-b5d47cca86cf9a0256983ed2cb4b2c297534ace225976ee9bb7830211ebbe703"></a>

<a id="canonical-d6cb095bcb6989db0a1c1a7a303f822b6a6df91bbc52512d87bc5ab27766711e"></a>

## excluded_values property — tls_fingerprint_matcher / 79fc764ce1d1 / 6

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

<a id="canonical-e8de9ec3248782e542f2fd5a5eda1cfffe5bfa837bfcc9519a44f47a83d0a2ab"></a>

## Next pages — tls_fingerprint_matcher / 79fc764ce1d1 / 7

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-7481d80ee8d0a60eb508aeefc8a77dad842c9e5fd19c82e895ebfa2c9826dd3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5e2efdb2115259424e5d6aa72f2a6b7aca58cb2157a5088c3097445c20e516a"></a>

## waf_action — waf_action / 6b73ceedb730 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- waf_action

<a id="canonical-68de118d4d4295e6bd3e781f066834486e0cf06b2f86b5b171d4d96966f991a2"></a>

Type: `"single"`. Computed.

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Upstream description:

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

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

<a id="canonical-ef06e3eadc840efb52c2a837ef34d078390cb304b055dda981a17d41aeb2157c"></a>

## Direct properties — waf_action / 6b73ceedb730 / 3

- [app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-50ac26b5d2591141fb33451f8196bd9e3bc711b53996394eb825331bfb4ab89b): complete subsection reference.

- [none](data-sources--service_policy_rule--reference--group-002.md#canonical-d96314117db9e4cf5445c7805c9f469642637470a6c1bafe3a70e9bc51e0bdc0): complete subsection reference.

- [waf_skip_processing](data-sources--service_policy_rule--reference--group-002.md#canonical-b5fee8d8d084562562498a9e3c5bf6847ee355e53bd587f18b52598f25b4ec6c): complete subsection reference.

<a id="canonical-eb8061c5517f4cd1e14b81ac34956da8ab9f3be7beace53e060c7e445c055169"></a>

## Next pages — waf_action / 6b73ceedb730 / 4

- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-50ac26b5d2591141fb33451f8196bd9e3bc711b53996394eb825331bfb4ab89b)
- [waf_action.none](data-sources--service_policy_rule--reference--group-002.md#canonical-d96314117db9e4cf5445c7805c9f469642637470a6c1bafe3a70e9bc51e0bdc0)
- [waf_action.waf_skip_processing](data-sources--service_policy_rule--reference--group-002.md#canonical-b5fee8d8d084562562498a9e3c5bf6847ee355e53bd587f18b52598f25b4ec6c)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-50ac26b5d2591141fb33451f8196bd9e3bc711b53996394eb825331bfb4ab89b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2fce536e7c1bd0a08a1fd20a5f79bd7377092af0534408ed1601f88770fd5826"></a>

## waf_action.app_firewall_detection_control — waf_action.app_firewall_detection_control / 95359ca8a246 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-7481d80ee8d0a60eb508aeefc8a77dad842c9e5fd19c82e895ebfa2c9826dd3a)
- waf_action.app_firewall_detection_control

<a id="canonical-bdd79412cece19057c6817382815fccd8e22e429a614569561109894560c72ba"></a>

Type: `"single"`. Computed.

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

<a id="canonical-dc3cb581803177df19eebe937d9a2e43e7c36c84a032665e462c36ee8d438a47"></a>

## Direct properties — waf_action.app_firewall_detection_control / 95359ca8a246 / 3

- [exclude_attack_type_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-be024ad900d75ca3dac5eeb6985a19904d16ea4d42893642a711a74837ae8734): complete subsection reference.

- [exclude_bot_name_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-f23c0c05c7cb95489411001984ecfbdd30d6586a73e121dba19966c0ac4b0f72): complete subsection reference.

- [exclude_signature_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-6558fda91331c1c86ab84b28cb52c49ab18686d98e86f21c579effed6a32e2cf): complete subsection reference.

- [exclude_violation_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-a95232c0f0cf4de8fb69c06e759580e023a1731ef9b42d3d201eaa8fc5e6eaf6): complete subsection reference.

<a id="canonical-5d5641322d0fe4b05b236deed7a7663fa6735beb5f1ce94bdfe99df7e016df0e"></a>

## Next pages — waf_action.app_firewall_detection_control / 95359ca8a246 / 4

- [waf_action.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-be024ad900d75ca3dac5eeb6985a19904d16ea4d42893642a711a74837ae8734)
- [waf_action.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-f23c0c05c7cb95489411001984ecfbdd30d6586a73e121dba19966c0ac4b0f72)
- [waf_action.app_firewall_detection_control.exclude_signature_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-6558fda91331c1c86ab84b28cb52c49ab18686d98e86f21c579effed6a32e2cf)
- [waf_action.app_firewall_detection_control.exclude_violation_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-a95232c0f0cf4de8fb69c06e759580e023a1731ef9b42d3d201eaa8fc5e6eaf6)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-7481d80ee8d0a60eb508aeefc8a77dad842c9e5fd19c82e895ebfa2c9826dd3a)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-be024ad900d75ca3dac5eeb6985a19904d16ea4d42893642a711a74837ae8734"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b67f799b4ae206122963cc8f6539b91e32491af9df2dae6a35226f45d53995d"></a>

## waf_action.app_firewall_detection_control.exclude_attack_type_contexts — waf_action.app_firewall_detection_control.exclude_attack_type_contexts / 761f02f6a2e9 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-7481d80ee8d0a60eb508aeefc8a77dad842c9e5fd19c82e895ebfa2c9826dd3a)
- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-50ac26b5d2591141fb33451f8196bd9e3bc711b53996394eb825331bfb4ab89b)
- waf_action.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-94fc1e8fe6f1ed8bbc2188ad19a57df311e5522c4aec2368ca11dd0d7d5e3750"></a>

Type: `"list"`. Computed.

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

<a id="canonical-3544c9b0e8ae6eacc5317bd10653cbdbbeaeb1dd7b45bff56e4102659dd0e413"></a>

## Direct properties — waf_action.app_firewall_detection_control.exclude_attack_type_contexts / 761f02f6a2e9 / 3

<a id="canonical-97d29978a70fe98644d4c2dadd31fff38b265d117f87ad871b0fc4fd1b051e30"></a>

<a id="canonical-9cc24b3487ed0bcc297a7d6f6fc31f3bc487d435ce56db8d52e9be9dc7a23b15"></a>

## context property — waf_action.app_firewall_detection_control.exclude_attack_type_contexts / 761f02f6a2e9 / 4

Type: `"string"`. Computed.

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

<a id="canonical-f0085ec310937b8154f59e4dadadcbdbedd53f62125f8d61192d240967d6b9e9"></a>

<a id="canonical-2b171375e8c7bf5c7f9274a3cad3e343a1d1622a2dbd7e0b49f391c9768abbad"></a>

## context_name property — waf_action.app_firewall_detection_control.exclude_attack_type_contexts / 761f02f6a2e9 / 5

Type: `"string"`. Computed.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

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

<a id="canonical-653d9be481f1429466191edd6d0bbc117ec72f034c1dabd0810803ae66c93e26"></a>

<a id="canonical-e0399bdbd9b4feaf598ed4c710c0bfd85c07547df075df1471afee3649fe006e"></a>

## exclude_attack_type property — waf_action.app_firewall_detection_control.exclude_attack_type_contexts / 761f02f6a2e9 / 6

Type: `"string"`. Computed.

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

<a id="canonical-7afe50fe441c8504fd8afbf0beec43fd8c6dc79d1109af7a469864d49dc08afd"></a>

## Next pages — waf_action.app_firewall_detection_control.exclude_attack_type_contexts / 761f02f6a2e9 / 7

- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-50ac26b5d2591141fb33451f8196bd9e3bc711b53996394eb825331bfb4ab89b)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-f23c0c05c7cb95489411001984ecfbdd30d6586a73e121dba19966c0ac4b0f72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e2ef883b260fb80a055118e4047094df1b153c41f7dd92ca707eebcfcfbf83c"></a>

## waf_action.app_firewall_detection_control.exclude_bot_name_contexts — waf_action.app_firewall_detection_control.exclude_bot_name_contexts / 42c624e02619 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-7481d80ee8d0a60eb508aeefc8a77dad842c9e5fd19c82e895ebfa2c9826dd3a)
- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-50ac26b5d2591141fb33451f8196bd9e3bc711b53996394eb825331bfb4ab89b)
- waf_action.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-7b18961e4d7afed2e51c84f996030ba68bd99ed32de86d93e5dc1ed728f10956"></a>

Type: `"list"`. Computed.

Bot Names to be excluded for the defined match criteria.

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

<a id="canonical-aee8e86e46b2274a495b38cd7a00016ab0f765679601c145380de28f5dc4aebe"></a>

## Direct properties — waf_action.app_firewall_detection_control.exclude_bot_name_contexts / 42c624e02619 / 3

<a id="canonical-6e90e64d925b28fe35f4ab40638e5d6603a2b1a7aaad76aa008edf201f69b211"></a>

<a id="canonical-3ab33542e5fca7b24c2efcf1f2ffd37ed070f9e707355c7715fb76d639ae112e"></a>

## bot_name property — waf_action.app_firewall_detection_control.exclude_bot_name_contexts / 42c624e02619 / 4

Type: `"string"`. Computed.

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

<a id="canonical-d51b0fb1d96295460fa8fca9f7a8df27285a5073c4cd170330f7812d8df62602"></a>

## Next pages — waf_action.app_firewall_detection_control.exclude_bot_name_contexts / 42c624e02619 / 5

- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-50ac26b5d2591141fb33451f8196bd9e3bc711b53996394eb825331bfb4ab89b)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-6558fda91331c1c86ab84b28cb52c49ab18686d98e86f21c579effed6a32e2cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8acb676054c807829b793857df062dba9397eca51debfa740bccb8747672fc21"></a>

## waf_action.app_firewall_detection_control.exclude_signature_contexts — waf_action.app_firewall_detection_control.exclude_signature_contexts / 0ca1903fcae8 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-7481d80ee8d0a60eb508aeefc8a77dad842c9e5fd19c82e895ebfa2c9826dd3a)
- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-50ac26b5d2591141fb33451f8196bd9e3bc711b53996394eb825331bfb4ab89b)
- waf_action.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-3ac0ce87875fe5c1e0549c08721a6d43620a4a184497953d826fe0e0c3191cac"></a>

Type: `"list"`. Computed.

Signature IDs to be excluded for the defined match criteria.

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

<a id="canonical-ac6e775410ac50da68f2e01954f8ab419af19d0a8bba6ad33e227a7511d97c0a"></a>

## Direct properties — waf_action.app_firewall_detection_control.exclude_signature_contexts / 0ca1903fcae8 / 3

<a id="canonical-6dc9ad7a01aa8f6b52391f8f3cd84d004f8c11afb7d17a3129d8d1cb198c7da3"></a>

<a id="canonical-7313c2a7ad5ddb1a9922707b0e800e5fa53cb87a4680e7007b4a9fe74c9a47b5"></a>

## context property — waf_action.app_firewall_detection_control.exclude_signature_contexts / 0ca1903fcae8 / 4

Type: `"string"`. Computed.

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

<a id="canonical-2324e7e3f052db760a47a873ad5849848e3f8fcb1ff95c96933c7c22e40fec47"></a>

<a id="canonical-af7008a66e1af7f1d25899d0d9bd24e4d71d6117625b993a659032516464688e"></a>

## context_name property — waf_action.app_firewall_detection_control.exclude_signature_contexts / 0ca1903fcae8 / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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

<a id="canonical-2e3c8e1fc87986f39302b33fb6def2aa8b372634ed9f52332ec68ebc80184b38"></a>

<a id="canonical-e5229c1b8c24a4243cf142c7542da8e6d3153c4489d5ed73fd1366cf3dc3994c"></a>

## signature_id property — waf_action.app_firewall_detection_control.exclude_signature_contexts / 0ca1903fcae8 / 6

Type: `"number"`. Computed.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

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

<a id="canonical-0538ce410ea6500779aeb76d07729d985b5357690a881244d67b43c91739bc95"></a>

## Next pages — waf_action.app_firewall_detection_control.exclude_signature_contexts / 0ca1903fcae8 / 7

- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-50ac26b5d2591141fb33451f8196bd9e3bc711b53996394eb825331bfb4ab89b)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-a95232c0f0cf4de8fb69c06e759580e023a1731ef9b42d3d201eaa8fc5e6eaf6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a455c675e83334d01af06d3b277673833fed88c7381d90493a72cc42080096f7"></a>

## waf_action.app_firewall_detection_control.exclude_violation_contexts — waf_action.app_firewall_detection_control.exclude_violation_contexts / 6682fd077aa9 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-7481d80ee8d0a60eb508aeefc8a77dad842c9e5fd19c82e895ebfa2c9826dd3a)
- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-50ac26b5d2591141fb33451f8196bd9e3bc711b53996394eb825331bfb4ab89b)
- waf_action.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-c3150a67ff750387b960d408c13011d9c3ae274d7b8ec8356ff14c109c371bf9"></a>

Type: `"list"`. Computed.

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

<a id="canonical-7cac8bc36937817b47d1d7a77e14710499b6d869c933aea8e5c54f86367ee813"></a>

## Direct properties — waf_action.app_firewall_detection_control.exclude_violation_contexts / 6682fd077aa9 / 3

<a id="canonical-9b8f700d604dd0844c7837769f255de745bfd4fd1d0d7498388fd3a8e2764546"></a>

<a id="canonical-c4539c5afb0baa007edb0ee08128ef9e003ca8d48a9a89f7780233551b2bfc00"></a>

## context property — waf_action.app_firewall_detection_control.exclude_violation_contexts / 6682fd077aa9 / 4

Type: `"string"`. Computed.

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

<a id="canonical-5bec8ad22a301eab8017f9feab3b198ab0d527ce5065f1ee7d21683641d1116f"></a>

<a id="canonical-d867041de9419d844918f2aabc75057134c1d81974ac07fd4601a7dbec4a93a4"></a>

## context_name property — waf_action.app_firewall_detection_control.exclude_violation_contexts / 6682fd077aa9 / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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

<a id="canonical-b9dcdef566f530c384244ccba5d017b63393e0aec5ae131a2b4e30f804a265b8"></a>

<a id="canonical-9f7d1bfef697d048af9b1501bbaa92c1ec753593ae857531966f7abefb8ecff0"></a>

## exclude_violation property — waf_action.app_firewall_detection_control.exclude_violation_contexts / 6682fd077aa9 / 6

Type: `"string"`. Computed.

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

<a id="canonical-7740f795a19e6ad870563676169062102a5aa148658bce800375732be13c248a"></a>

## Next pages — waf_action.app_firewall_detection_control.exclude_violation_contexts / 6682fd077aa9 / 7

- [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-50ac26b5d2591141fb33451f8196bd9e3bc711b53996394eb825331bfb4ab89b)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-d96314117db9e4cf5445c7805c9f469642637470a6c1bafe3a70e9bc51e0bdc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e763b4dbf46599964517fccc771c08dcec9648935c8467e1e07d9760e45cb76"></a>

## waf_action.none — waf_action.none / d284f3d831b5 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-7481d80ee8d0a60eb508aeefc8a77dad842c9e5fd19c82e895ebfa2c9826dd3a)
- waf_action.none

<a id="canonical-7619e066b5db2aad474f45c3a82cf891a5c7f3b6f2a9d080b91beba3337b4ffa"></a>

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

<a id="canonical-2f86d17c9ccb215ac8d097858c8be409f7ed17a420e7b3211b08241c3330c98d"></a>

## Direct properties — waf_action.none / d284f3d831b5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3142a33de5799de7ebb53363a00c3805eeb10c712e430ac1f54bdc757a19b4f0"></a>

## Next pages — waf_action.none / d284f3d831b5 / 4

- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-7481d80ee8d0a60eb508aeefc8a77dad842c9e5fd19c82e895ebfa2c9826dd3a)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-b5fee8d8d084562562498a9e3c5bf6847ee355e53bd587f18b52598f25b4ec6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a646e40dfe20b20f2b1913de79970b1daadfd6b73d76fa443903bc448110c016"></a>

## waf_action.waf_skip_processing — waf_action.waf_skip_processing / 031e6d86c46e / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-7481d80ee8d0a60eb508aeefc8a77dad842c9e5fd19c82e895ebfa2c9826dd3a)
- waf_action.waf_skip_processing

<a id="canonical-6a9db06d2e5f8785dc172f9af1937e7cfbe347a6da7c5b1788d9056d4a367785"></a>

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

<a id="canonical-d03a3825aec3440c84b8b84e14cbd2e6d3e199a846e4a14119037bf12a424ff3"></a>

## Direct properties — waf_action.waf_skip_processing / 031e6d86c46e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fd1b6a912ba759e6afde2fcc43b86af8c6a8a56610569be5d2216d76b70b978f"></a>

## Next pages — waf_action.waf_skip_processing / 031e6d86c46e / 4

- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-7481d80ee8d0a60eb508aeefc8a77dad842c9e5fd19c82e895ebfa2c9826dd3a)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
