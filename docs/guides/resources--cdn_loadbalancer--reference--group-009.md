---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-b3a08e41d8833ca68a5654ab8765ab57d22fcd2da27247be6df6d8d0a099424e"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers / d5fe95f51372 / 4

- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--cdn_loadbalancer--reference--group-008.md#canonical-c081d088e9fa71bea49ec98adfd57c9bf4954faeac73172908314923ef2dab94)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-77dfc9d77746c25e588f51f9e549713a504c6a0a3b7a7fa39ea343a5578da32f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82a7674fff877fd864c440201f419a496f55c4d71a2d4ee021e1be1130f19cc4"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.redirect — bot_defense.policy.protected_app_endpoints.mitigation.redirect / 8af4a2dd98aa / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-008.md#canonical-f16163a87e781bdfadca2e4dbb0ad838c5a685bd6f14b94283036c08c9cca611)
- bot_defense.policy.protected_app_endpoints.mitigation.redirect

<a id="canonical-a264899c0df1704f925414ee23096216bf31be0e30dde31520f98a2560290b8f"></a>

Type: `"object"`. single nested block, Optional.

Redirect bot mitigation. Redirect request to a custom URI.

Upstream description:

Redirect request to a custom URI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("uri")}
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
redirect {
  # Configure direct properties listed below.
}
```

<a id="canonical-43bfac2b418e7acfe0e6ea24c1a872887a526e6b27ec329baa6202c4acc71b2d"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.mitigation.redirect / 8af4a2dd98aa / 3

<a id="canonical-a1c99b40290b8e23e7e7622de4460717ea1ca4afd11b4a34b6bd6d43897e88aa"></a>

<a id="canonical-f92699d4057f21a966af345fb17ee9949566b22576b419f90953e95548eea8d7"></a>

## uri property — bot_defense.policy.protected_app_endpoints.mitigation.redirect / 8af4a2dd98aa / 4

Type: `"string"`. Optional.

URI location for redirect may be relative or absolute.

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
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  }
}
```

<a id="canonical-4ebca4425f61c0b78784a9bba0bc4534ca0cfcccf7a18f74993cd20897d2511f"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.mitigation.redirect / 8af4a2dd98aa / 5

- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-008.md#canonical-f16163a87e781bdfadca2e4dbb0ad838c5a685bd6f14b94283036c08c9cca611)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b8b5785988c2538a52493ad8157015d8f1ba98bac382a915fd589b9380999a6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73022fdde3554d544c81f286d8c618114ab06036a5ce1e141050e90f4bdf69bd"></a>

## bot_defense.policy.protected_app_endpoints.mobile — bot_defense.policy.protected_app_endpoints.mobile / a81bad573725 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- bot_defense.policy.protected_app_endpoints.mobile

<a id="canonical-6d7523808b7dd3eb019141d4cbd65d8f76455568b0d7981ac771df6965a7c551"></a>

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
mobile = {}
```

<a id="canonical-56c36554afca026a4d284a2f63bf2d46c272666eecb5ad67be2eab17718b9b7d"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.mobile / a81bad573725 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-76e69da57844856fc325b64e6f24b4666a8a59142e2d6325106ae49f96f4dcf3"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.mobile / a81bad573725 / 4

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-abdeb6e2af99c135f6189ecbc2cbb3bddda5ff11a216c11dc75c9a3a6bc39bac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cd828c3ec0d46b377b8971a1ccd34c8c383c8acdc38df870aebd9ef37469dd2"></a>

## bot_defense.policy.protected_app_endpoints.path — bot_defense.policy.protected_app_endpoints.path / d30cbc3825f2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- bot_defense.policy.protected_app_endpoints.path

<a id="canonical-9a40d8f7eecca03e99ae53a8fdd29d6edab7697ee619a51a440ead2e8144f107"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-623cd2c7142c5632cb873ad2691d66f19cbf07527977056df45c0ae4fe4e7887"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.path / d30cbc3825f2 / 3

<a id="canonical-176d6813c3f268110bd3acf368f770ec36e95ba24e8310a28865d892a02bbcc1"></a>

<a id="canonical-9fe4178ff63355ca29a395d27c08b722655cd46d419c88a1aa36154b3a398b06"></a>

## path property — bot_defense.policy.protected_app_endpoints.path / d30cbc3825f2 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-6f9849b68a0deced7ae9ed5286da26b5346ce80952be10b604d41355a6238c90"></a>

<a id="canonical-027d635194aa60dd50d9dda0683b01ec77d6b128475eb1a40684d4a7917407b7"></a>

## prefix property — bot_defense.policy.protected_app_endpoints.path / d30cbc3825f2 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-0557446bd4bac6207127ae4db4c5954b6851f38e116f9e2d76201a392732fa9e"></a>

<a id="canonical-2ee720b5693563e834bb7e50a1906cb6f0ef14bcf24da584e7ed48ce4a74cd2e"></a>

## regex property — bot_defense.policy.protected_app_endpoints.path / d30cbc3825f2 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-a6de4713476f8f378b608a5ce583adc50fe3991548f1a375154463dc2411fdbc"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.path / d30cbc3825f2 / 7

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-aa258c2188acb4b9e93bd90cf23364aef6cbcd679c3c7bd2d0dd58a22587455b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73565ba323ad8f408309be43a1e33d044e9b9c6c6a0f533078bc75c347410222"></a>

## bot_defense.policy.protected_app_endpoints.query_params — bot_defense.policy.protected_app_endpoints.query_params / 0d5b7cec4671 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- bot_defense.policy.protected_app_endpoints.query_params

<a id="canonical-8df68d8a86542275e1782513dd55ada1ea4529404512e3197434b3f5ee5444ad"></a>

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

<a id="canonical-35e78664c1459164e36411d91f015e20988dda553d5baafe0f9c4cc506533d0a"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.query_params / 0d5b7cec4671 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-009.md#canonical-cd1c474e8237c0b4989d1cb9584faf6afcb805b4684c7de59f2a633e6db92d2a): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-009.md#canonical-ca0a4077c2c2b0e608f34aaba229b260b6374a942d1884895041700276061cd3): complete subsection reference.

<a id="canonical-fef6e143742e4d56d9c0f4f1f862ff1f2cf7ff43fd009e751a54a8f82ae44580"></a>

<a id="canonical-ed4d8a9fd7b502473090869d46b778981609563fa6e78caf221e332c0668a89a"></a>

## invert_matcher property — bot_defense.policy.protected_app_endpoints.query_params / 0d5b7cec4671 / 4

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

- [item](resources--cdn_loadbalancer--reference--group-009.md#canonical-e5aeb62bb892ca5631d647187301a78609637bd4c80596d88952ee811b5d3e07): complete subsection reference.

<a id="canonical-edfec2dd974e88f1a9bb503149d0b44580358307403b0cffdbf81d52c817cd43"></a>

<a id="canonical-f241062c136941fcf692e7cab16ad75b53346b1b1b44d472875a91886003aad0"></a>

## key property — bot_defense.policy.protected_app_endpoints.query_params / 0d5b7cec4671 / 5

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

<a id="canonical-effbde20f47e4c3fdcd0499106aa95a5582543de9ebffb34437b285b82269edd"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.query_params / 0d5b7cec4671 / 6

- [bot_defense.policy.protected_app_endpoints.query_params.check_not_present](resources--cdn_loadbalancer--reference--group-009.md#canonical-cd1c474e8237c0b4989d1cb9584faf6afcb805b4684c7de59f2a633e6db92d2a)
- [bot_defense.policy.protected_app_endpoints.query_params.check_present](resources--cdn_loadbalancer--reference--group-009.md#canonical-ca0a4077c2c2b0e608f34aaba229b260b6374a942d1884895041700276061cd3)
- [bot_defense.policy.protected_app_endpoints.query_params.item](resources--cdn_loadbalancer--reference--group-009.md#canonical-e5aeb62bb892ca5631d647187301a78609637bd4c80596d88952ee811b5d3e07)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-cd1c474e8237c0b4989d1cb9584faf6afcb805b4684c7de59f2a633e6db92d2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d5b7ec57851fd02fa854b29c47fa8138761e5e4190474c3bc14b261db625de4"></a>

## bot_defense.policy.protected_app_endpoints.query_params.check_not_present — bot_defense.policy.protected_app_endpoints.query_params.check_not_present / 5de809b39355 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-009.md#canonical-aa258c2188acb4b9e93bd90cf23364aef6cbcd679c3c7bd2d0dd58a22587455b)
- bot_defense.policy.protected_app_endpoints.query_params.check_not_present

<a id="canonical-01306574bef9736cadad163e403abe9e305906a474ea47890980d8865031fd09"></a>

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

<a id="canonical-868787da2425649fcaa1f70ce911e6a541daeb8bc41980b1ecee8f116843c0d9"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.query_params.check_not_present / 5de809b39355 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-15c843efb4c4152856b6ee3d2f3cdf60ce3c61a0dbff19114350107d1fcbd43c"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.query_params.check_not_present / 5de809b39355 / 4

- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-009.md#canonical-aa258c2188acb4b9e93bd90cf23364aef6cbcd679c3c7bd2d0dd58a22587455b)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ca0a4077c2c2b0e608f34aaba229b260b6374a942d1884895041700276061cd3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed1b8fbbca889f84affbb138bbb6598b240c97ab3411fbd1d8c44465e823e77b"></a>

## bot_defense.policy.protected_app_endpoints.query_params.check_present — bot_defense.policy.protected_app_endpoints.query_params.check_present / 9a09edb32b55 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-009.md#canonical-aa258c2188acb4b9e93bd90cf23364aef6cbcd679c3c7bd2d0dd58a22587455b)
- bot_defense.policy.protected_app_endpoints.query_params.check_present

<a id="canonical-0eb3c0386f9909ddd1343fff4201a75cf77fb89f79ea16554213a0ceed435c08"></a>

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

<a id="canonical-81615d3e779c6d627b38e8bc163a21b9f2a1269ed0fe7701c40487bcbd8c326f"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.query_params.check_present / 9a09edb32b55 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-641e7dff611d833cb613a1485fb68e18a4f3bda76cd81e78b655125a3e541b54"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.query_params.check_present / 9a09edb32b55 / 4

- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-009.md#canonical-aa258c2188acb4b9e93bd90cf23364aef6cbcd679c3c7bd2d0dd58a22587455b)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e5aeb62bb892ca5631d647187301a78609637bd4c80596d88952ee811b5d3e07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df5c1176e50929b47135187040fc8bee7f77646e23f71aaed285f657f056cb93"></a>

## bot_defense.policy.protected_app_endpoints.query_params.item — bot_defense.policy.protected_app_endpoints.query_params.item / ea4c6c138fce / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-009.md#canonical-aa258c2188acb4b9e93bd90cf23364aef6cbcd679c3c7bd2d0dd58a22587455b)
- bot_defense.policy.protected_app_endpoints.query_params.item

<a id="canonical-1cf733cf1654c0a19fd4bb37f6c2dacee1fb23060281615f2c976d40eaf6a0db"></a>

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

<a id="canonical-72971c2973ffb17ffc221aaa455e5d603ab34b8474c4bc80932a033d19d2245b"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.query_params.item / ea4c6c138fce / 3

<a id="canonical-d1d0b7ed9150a2bd0ca67b13c914e28ac6765bacc007c8312661b9014b0d54d7"></a>

<a id="canonical-809854c11e03151caa58b745d3fd746f06ea18a0058fd78308da164933066ead"></a>

## exact_values property — bot_defense.policy.protected_app_endpoints.query_params.item / ea4c6c138fce / 4

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

<a id="canonical-8ae8b1411b3d18f2e816a6fe2ddf5c885bc00a81e7474e0417a936db1349d8ad"></a>

<a id="canonical-fb06609875c34a550f02c384c3390930f412b33f0213312ef78afd2e64e3fea1"></a>

## regex_values property — bot_defense.policy.protected_app_endpoints.query_params.item / ea4c6c138fce / 5

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

<a id="canonical-b49abcac12e847c3c0df6309e9f2cab4fe934dd9db6bcd00fa9eb60650f40f2c"></a>

<a id="canonical-03aa27dc426b3951720fa68645343d079f858ebb2aea6b5899b882c030733abb"></a>

## transformers property — bot_defense.policy.protected_app_endpoints.query_params.item / ea4c6c138fce / 6

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

<a id="canonical-197c656585c93ab44f54eeb9d84f7aeee19b8f76a984f0b6e8bd7ebce1fb7713"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.query_params.item / ea4c6c138fce / 7

- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-009.md#canonical-aa258c2188acb4b9e93bd90cf23364aef6cbcd679c3c7bd2d0dd58a22587455b)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-536994a1b4b82c44b0a6cdce0bb07153922fd8b35f3f716b9800ddc3784246cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9400d0e2338a8e579d086f4ba19ebdd6592de932a1471b793574dd443717db0"></a>

## bot_defense.policy.protected_app_endpoints.undefined_flow_label — bot_defense.policy.protected_app_endpoints.undefined_flow_label / ad8231ee9594 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- bot_defense.policy.protected_app_endpoints.undefined_flow_label

<a id="canonical-bec22a7136e8a84e9c0bc12108a9c9ed3745fac42cd0ae4e12bc4d24ba9984d2"></a>

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
undefined_flow_label = {}
```

<a id="canonical-ffa38c8990fb4080fd1c126108da0a5c517605790b93c63fb50fac203cbfb4e8"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.undefined_flow_label / ad8231ee9594 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7c726b3cd650ec0a20ec49d9c4e7cbab3d0d25fa46ec853fb48a0e76116314e2"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.undefined_flow_label / ad8231ee9594 / 4

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1783c8921d2e090218141c9bc3dc8b752ed8a54e2470ed1838c571db1c056084"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01272999c17bf01cba9d5aaf01d5c0332ea24d09eda4dc04915c4d22d5b2a557"></a>

## bot_defense.policy.protected_app_endpoints.web — bot_defense.policy.protected_app_endpoints.web / bacc2187fcc0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- bot_defense.policy.protected_app_endpoints.web

<a id="canonical-7fe9afc95e95e09623b19e36b5313a5d1eef6b2b7bf3b5cc255c3e9cbd9b273b"></a>

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
web = {}
```

<a id="canonical-895a88f9717895408e3b243e904995429320fe016443be894b333aac745b813f"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.web / bacc2187fcc0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c8251b554de1c137a268ffb76a588ba9405fe934e014a44eccf873f2db09e3b6"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.web / bacc2187fcc0 / 4

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-4fcf4fe1ba8a4f2d73ee2ae019775f61bddbe047d0f65fd6e7f1652cece45590"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8de9475e2f6406397e29ee277006e25f78290145297e50fe7ad4f6ab02bce931"></a>

## bot_defense.policy.protected_app_endpoints.web_mobile — bot_defense.policy.protected_app_endpoints.web_mobile / c9a2403b6fd5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-6dba61de9c04eaa8383d41e91ddcb45852b267031ad3837ac7beb232a510a110)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-efaa8e28866e10568155a83fc0ba38e9c443cdac4922e0cd2de7fe954615c1d5)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- bot_defense.policy.protected_app_endpoints.web_mobile

<a id="canonical-eb7fba0f039763650f42a0041ad96866fd20921b9d2f40ad2ed9c7fa5b7487b1"></a>

Type: `"object"`. single nested block, Optional.

Web and Mobile traffic type. Web and Mobile traffic type.

Upstream description:

Web and Mobile traffic type.

Receipt-pinned upstream constraints:

```json
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
web_mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-6e3a75e676630b88efc380e86927263ab961aa12c550f32d11a53ba9ab71fe78"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.web_mobile / c9a2403b6fd5 / 3

<a id="canonical-b424bd175ef7bd4dcca1c00e519a826eabedee27ef8937bf03fb054840f3aa8d"></a>

<a id="canonical-a0c48dc5ab279f5a68ddd6a65da4a4de51cb0407fd6b297876f3ca90df66ef24"></a>

## mobile_identifier property — bot_defense.policy.protected_app_endpoints.web_mobile / c9a2403b6fd5 / 4

Type: `"string"`. Optional.

\[Enum: HEADERS\] Mobile identifier type - HEADERS: Headers Headers. The only possible value is
\`HEADERS\`. Defaults to \`HEADERS\`.

Upstream description:

Mobile identifier type

&#8203;- HEADERS: Headers

Headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("HEADERS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "HEADERS",
  "enum": [
    "HEADERS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d9a381076a78a7d13f8bfb5ea11b7025c97f9a4a8033e397d261f04e8c91ed57"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.web_mobile / c9a2403b6fd5 / 5

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-f00539ea8e728c1a2846d3d9fb3907268733276eb08d8e3ce27c90b3c0503447)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1c2555638ed3b0821ac370fa14c5cddaae5615d0e10f75df287474926bd68764"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-765676246d1ea71454eeba20bbec15e28803163ec4c43f6172865edf7f3fe246"></a>

## captcha_challenge — captcha_challenge / b45b0a2beed2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- captcha_challenge

<a id="canonical-b0269a59f80b535b569b011f2b5ec74c9800e821a409a0b851f442b43284e6a2"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: captcha\_challenge, enable\_challenge, js\_challenge, no\_challenge,
policy\_based\_challenge; Default: no\_challenge\] Enables loadbalancer to perform captcha challenge
Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that
pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is
configured to do Captcha Challenge, it will redirect..

Upstream description:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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

OneOf alternatives in this subsection:

- [captcha_challenge](resources--cdn_loadbalancer--reference--group-009.md#canonical-b0269a59f80b535b569b011f2b5ec74c9800e821a409a0b851f442b43284e6a2)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-61db621bec5658e33b4d89b4822b6deea0e32dd5026aeca0af822d6d26be4067)
- [js_challenge](resources--cdn_loadbalancer--reference--group-011.md#canonical-52c4ddba5ad61b31899f4590c585ce96cd8ae7ff6c559dc9b64266a1842ad358)
- [no_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-daead701143ebd9525dd2c685a8cb14eb0f1a03fb72c0f48599432246cfec8f5)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-d4be8fe0e687fee1885e9b29faa62a129ea348a6ad5c2bd7c73688b1e3a5fa88)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
captcha_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-73fa7145a6da6bbb524f9485e88f414c87728892805e1c3f599321701261916d"></a>

## Direct properties — captcha_challenge / b45b0a2beed2 / 3

<a id="canonical-8cba7c14aa8446c52528cb6067628e0dddc5b687f6006172d08ec63b261ef504"></a>

<a id="canonical-5cc8cc7a48e875b948a79e19bb6fc152b780a97ed40049327472a10a8e694656"></a>

## cookie_expiry property — captcha_challenge / b45b0a2beed2 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-9b38cb4956051c1fbfd5ad96f811c2141645edc84c6aac2754f90e82e2346ffc"></a>

<a id="canonical-6ad858d7f60146b0ad31fc65ea5a9632a0ab7e062e462636321ce5f513e69d42"></a>

## custom_page property — captcha_challenge / b45b0a2beed2 / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-49f18036c70f855d9e4ceed0025162a432eb4f0ef9ed6b597c3f8c65a6ae3e9e"></a>

## Next pages — captcha_challenge / b45b0a2beed2 / 6

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b22c8437c714a8a2d8a874d245d2109a308ea7f9f5b1610fc8ab589612e8de33"></a>

## client_side_defense — client_side_defense / 4c364086855b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- client_side_defense

<a id="canonical-91931fcf38038d50e1b980032add4957bb0494b6f4b524271da05db5d2be9859"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: client\_side\_defense, disable\_client\_side\_defense; Default:
disable\_client\_side\_defense\] Defines various configuration OPTIONS for Client-Side Defense
Policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense Policy.

Receipt-pinned upstream constraints:

```json
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

- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-91931fcf38038d50e1b980032add4957bb0494b6f4b524271da05db5d2be9859)
- [disable_client_side_defense](resources--cdn_loadbalancer--reference--group-010.md#canonical-bb2a3dbac6975406cdc1e357f4ac5be5184d342f5138c0ba692f130782f07fbf)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
client_side_defense {
  # Configure direct properties listed below.
}
```

<a id="canonical-f1d75811c4d092ffa6eb9858a3e136aa12d79120b172e1124a0b6810e587b488"></a>

## Direct properties — client_side_defense / 4c364086855b / 3

- [policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745): complete subsection reference.

<a id="canonical-4c08aff5f748fe84ffca5493a904840cf8acb705f5c330804330497bbb52fb89"></a>

## Next pages — client_side_defense / 4c364086855b / 4

- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e674ac914d6f81b9571590d680737d76605e1011f7981d5fedf504ec3d950d7a"></a>

## client_side_defense.policy — client_side_defense.policy / ac7f7711dfaf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- client_side_defense.policy

<a id="canonical-5f768096dc5c6637729fe456037654909560c0732430bdd388ca4271dab1a242"></a>

Type: `"object"`. single nested block, Optional.

Defines various configuration OPTIONS for Client-Side Defense policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

Terraform syntax:

```terraform
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-146868d87c9fb63cc4c61237b3acd9579be87b1260da25f4d8fdd83bc1a58931"></a>

## Direct properties — client_side_defense.policy / ac7f7711dfaf / 3

- [disable_js_insert](resources--cdn_loadbalancer--reference--group-009.md#canonical-9cc38909e2a1d3846c438d02bbe5782754a2e3883560787350168dce42f31297): complete subsection reference.

- [js_insert_all_pages](resources--cdn_loadbalancer--reference--group-009.md#canonical-a1d7cdf86e5d4568abb7b30ff8322b88a2c5220105d83a06952805801559e9a9): complete subsection reference.

- [js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-7f7202e7d9e9da30d15caed6e2f70b70d85b0c758b62d371ad427b4cb68e9254): complete subsection reference.

- [js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233): complete subsection reference.

<a id="canonical-700eeda1813911d69e3380b052febbcf6201f0a1c8f18fd9176d6b4156429fe9"></a>

## Next pages — client_side_defense.policy / ac7f7711dfaf / 4

- [client_side_defense.policy.disable_js_insert](resources--cdn_loadbalancer--reference--group-009.md#canonical-9cc38909e2a1d3846c438d02bbe5782754a2e3883560787350168dce42f31297)
- [client_side_defense.policy.js_insert_all_pages](resources--cdn_loadbalancer--reference--group-009.md#canonical-a1d7cdf86e5d4568abb7b30ff8322b88a2c5220105d83a06952805801559e9a9)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-7f7202e7d9e9da30d15caed6e2f70b70d85b0c758b62d371ad427b4cb68e9254)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-9cc38909e2a1d3846c438d02bbe5782754a2e3883560787350168dce42f31297"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac1e65ee8f7a46c79b76cb0e3540f39a8aa3eeea188ae9c53dd5e061432f82b1"></a>

## client_side_defense.policy.disable_js_insert — client_side_defense.policy.disable_js_insert / d3ea555968a0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- client_side_defense.policy.disable_js_insert

<a id="canonical-275baa2334be2ba9acdc3bb224a19fda89c8f0e9789871829191e818368ab88e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

<a id="canonical-006e0538ab965cf501c79fb7a3d5e1c7aeb448d2551bb1187ade585a8ea80812"></a>

## Direct properties — client_side_defense.policy.disable_js_insert / d3ea555968a0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ede8b236993cd35d76fd1b597716bb38bbde597954d1adeb03e42d0f95a5c443"></a>

## Next pages — client_side_defense.policy.disable_js_insert / d3ea555968a0 / 4

- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a1d7cdf86e5d4568abb7b30ff8322b88a2c5220105d83a06952805801559e9a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-251b911f1eb447e8333b7b07de29978262d2aea8c932f9e131be2e61e681c281"></a>

## client_side_defense.policy.js_insert_all_pages — client_side_defense.policy.js_insert_all_pages / 97261c09d10c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- client_side_defense.policy.js_insert_all_pages

<a id="canonical-99a0f9422fffe1d0ba66020f31b800482aa41f5d3525a41149882662eceb1948"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for js insert all pages.

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
js_insert_all_pages = {}
```

<a id="canonical-821334ecc292c24efe819da9484df1c5ae005a623c386984cbcbef65e4bbddd5"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages / 97261c09d10c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-85d11916ac6c843644507a60f4580af1f1d4d99bd3d63959ff888bdebdc8ca77"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages / 97261c09d10c / 4

- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7f7202e7d9e9da30d15caed6e2f70b70d85b0c758b62d371ad427b4cb68e9254"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f83048cdb6ffa1cea466461dc52ee643d7b57c806e2637d24b15930aa2e97ac"></a>

## client_side_defense.policy.js_insert_all_pages_except — client_side_defense.policy.js_insert_all_pages_except / 377cc3ae9b0b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- client_side_defense.policy.js_insert_all_pages_except

<a id="canonical-cadb41b6e89c380fcb254afbd3c3f4f6a63a7b8f27dfd23e940e0c0e4d5a8323"></a>

Type: `"object"`. single nested block, Optional.

Insert Client-Side Defense JavaScript in all pages with the exceptions.

Receipt-pinned upstream constraints:

```json
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
js_insert_all_pages_except {
  # Configure direct properties listed below.
}
```

<a id="canonical-bbcf15f5dbe272e2b6e4608605c9d5e1d61f63c678f2798fb6490008bd405527"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except / 377cc3ae9b0b / 3

- [exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-99ba8ab2d40498e8cfc622e05d5a6c5545794388ce5a37c886afd599690f2908): complete subsection reference.

<a id="canonical-48a11d7002bbd3dc1e0118806a5a7924758e823092a132ab727d1f4c30f64c2d"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except / 377cc3ae9b0b / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-99ba8ab2d40498e8cfc622e05d5a6c5545794388ce5a37c886afd599690f2908)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-99ba8ab2d40498e8cfc622e05d5a6c5545794388ce5a37c886afd599690f2908"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1b2175b2bec242fb4f5bedb5ada1b6efd5f7b201ef7406af5ef7bf8cf279b98"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list — client_side_defense.policy.js_insert_all_pages_except.exclude_list / 004f3b364559 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-7f7202e7d9e9da30d15caed6e2f70b70d85b0c758b62d371ad427b4cb68e9254)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-83d90c180de3247b409916d0b3d91d4895f96c47467d7f63cdff91751ed2fbe1"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-b8e15e95cb4dcc95195279a6a2d152fae6d8a846dafc8ee9f2824ee20009fe9b"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list / 004f3b364559 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-2faac59f49aaad9f10354d1b5d086a2a559fc1675933715cb7b6fe333d7749b4): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-75e0a7a1a8445837302693b26c67d291caf03d6037a6e87299116b5f00868e45): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-009.md#canonical-a93cba8e196e667e71cce4c3222657944129643fb180a206df1ffd88e70f1195): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-009.md#canonical-a43c45f48db479cacbc31b7d7a889db5fdd4f67e9c54a8cae349962f495f59c6): complete subsection reference.

<a id="canonical-b597ac60984774d8ee0b9437933c340340c4bbb606afa2e443cf82dd4599d1ca"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list / 004f3b364559 / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-2faac59f49aaad9f10354d1b5d086a2a559fc1675933715cb7b6fe333d7749b4)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-75e0a7a1a8445837302693b26c67d291caf03d6037a6e87299116b5f00868e45)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata](resources--cdn_loadbalancer--reference--group-009.md#canonical-a93cba8e196e667e71cce4c3222657944129643fb180a206df1ffd88e70f1195)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.path](resources--cdn_loadbalancer--reference--group-009.md#canonical-a43c45f48db479cacbc31b7d7a889db5fdd4f67e9c54a8cae349962f495f59c6)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-7f7202e7d9e9da30d15caed6e2f70b70d85b0c758b62d371ad427b4cb68e9254)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2faac59f49aaad9f10354d1b5d086a2a559fc1675933715cb7b6fe333d7749b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bdab1facc488cc2e324da04a6a2e7605f87e331b750e72a21ca1241824b355b"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain — client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 9d6b2717888e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-7f7202e7d9e9da30d15caed6e2f70b70d85b0c758b62d371ad427b4cb68e9254)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-99ba8ab2d40498e8cfc622e05d5a6c5545794388ce5a37c886afd599690f2908)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-5cc0377713697f42a7ff58e9f3195168f2b5c99831ff3a99844e52140e4d9562"></a>

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

<a id="canonical-67ed9ca00387a86d3f12bf76886ae8674f17c979b3b7a3106abe164eb051996e"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 9d6b2717888e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a1e2a964d598f727db1d955daf06e6d3b23e5b8e86855c51d6e58f3359200adb"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 9d6b2717888e / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-99ba8ab2d40498e8cfc622e05d5a6c5545794388ce5a37c886afd599690f2908)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-75e0a7a1a8445837302693b26c67d291caf03d6037a6e87299116b5f00868e45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3308fef2edfc323449a2c4f953e1bffe687829249efccd951069fe396bee6a4"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / 9c76bfca335e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-7f7202e7d9e9da30d15caed6e2f70b70d85b0c758b62d371ad427b4cb68e9254)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-99ba8ab2d40498e8cfc622e05d5a6c5545794388ce5a37c886afd599690f2908)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-f72bc2159fe7f8651e88d0cb4050c4a703aa7b8f478562c943194b61786764d4"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-c7aec1d3518689a95d46572cc055317295e20e1a6fa901e417344eb730af5200"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / 9c76bfca335e / 3

<a id="canonical-cdba1c23a7a835803cef06a548e6b7d5c25930b9776494b2fa87aa663900969f"></a>

<a id="canonical-a6f2bd09736e11a9f1dd9392b0da98ec54f751f7fcff084578689c3543a9d1c7"></a>

## exact_value property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / 9c76bfca335e / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-aa1b683f5b4527835321c58e067ce70e7be21b2ccaa0bdfe85fb30aee3e1ec27"></a>

<a id="canonical-feb0c9c4508d04aa5c0b175584d262548d955d546d567684a4f3e95ac1d5e1de"></a>

## regex_value property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / 9c76bfca335e / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-fb9ba80c648b9701f462b1a3751d9b28d531b1ba475a73741040439cbc76408a"></a>

<a id="canonical-f15bdbe2ceb0302e39eec0ec6c2363743eafc7d859978cd1ab9fb49d3acf9085"></a>

## suffix_value property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / 9c76bfca335e / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-0534fe174df81ead725d32334502fc8202f22f89866f10da5cb95743d56861b0"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / 9c76bfca335e / 7

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-99ba8ab2d40498e8cfc622e05d5a6c5545794388ce5a37c886afd599690f2908)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a93cba8e196e667e71cce4c3222657944129643fb180a206df1ffd88e70f1195"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-170d06e948e32f47c66ef366691fc1e088de76ba969a29d30468e5b0684487d1"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 2509dfee2330 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-7f7202e7d9e9da30d15caed6e2f70b70d85b0c758b62d371ad427b4cb68e9254)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-99ba8ab2d40498e8cfc622e05d5a6c5545794388ce5a37c886afd599690f2908)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-7fd966e66c1638b270316d48132580ae7a9e208519e6c6d6c8ce8da1775bbb54"></a>

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

<a id="canonical-9fcabea5eb337d8aea0da8261f0ed56b80c329ca76ceae0c6d95ed5e83c203e4"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 2509dfee2330 / 3

<a id="canonical-2d0c57639823c24c7034584f7531481de70dc7b2d8c43d52d04c89c876e29776"></a>

<a id="canonical-dd2046a36be19c11b3373062b0283b46e4cb77be431046ee49d4f1950b5edd57"></a>

## description_spec property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 2509dfee2330 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-e558f10178377972dbcbc8e14980f8b0b8cede35f9331853ec57effe985701bd"></a>

<a id="canonical-ffd813415284e6edbab05e91baa210b4ce8af59ec0109c54015014f52209155e"></a>

## name property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 2509dfee2330 / 5

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

<a id="canonical-37ba1ff2ab103c65760e5289bf5ddf169f995ab894bef72d23094e8e40bf748f"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 2509dfee2330 / 6

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-99ba8ab2d40498e8cfc622e05d5a6c5545794388ce5a37c886afd599690f2908)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a43c45f48db479cacbc31b7d7a889db5fdd4f67e9c54a8cae349962f495f59c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a7960f9c7816cf79163d87dac6f72d8ed7d32a1b4a6ebfed6b461334ded0e3b"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.path — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / 4f438c9989b1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-009.md#canonical-7f7202e7d9e9da30d15caed6e2f70b70d85b0c758b62d371ad427b4cb68e9254)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-99ba8ab2d40498e8cfc622e05d5a6c5545794388ce5a37c886afd599690f2908)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-67774fd1ba61404c057aa1f86409a405803ce86a4a7bb6a4f353a46df97db585"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-3befb6ed05ff1977b33d6dd8a7ce8fdf45d7541068b9469549f6ec37e3d54096"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / 4f438c9989b1 / 3

<a id="canonical-35c30ef581bf6476eabaae0cce6d8668efaf3ee3266bec451102b4d084ade279"></a>

<a id="canonical-64668ac21db56990e297b7b2145b3bdca2d48696d0c413ff5b590a419ebe8661"></a>

## path property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / 4f438c9989b1 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-e39e80760a76678e9409e422d36ef823a48bbbf7a70a70477c68591f45639061"></a>

<a id="canonical-e0acbd8f12c35d0d437c76b4d2ee8618e56b080bc4416ad7e955cd078b748bfa"></a>

## prefix property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / 4f438c9989b1 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-1669a819e7417efb68fa5e2b34eb35f961077b1d5a282a76adf76ec1cbbc71ff"></a>

<a id="canonical-5172e4e3f81e544f60af99dc36b0fa1dc40cbf8a8c643b0dc8c3b2ffd9eb53f2"></a>

## regex property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / 4f438c9989b1 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-80e74d97596e58f947fca0d3bfc15e3cc9f17b1c7f3ba3c62ca1222c9d54897e"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / 4f438c9989b1 / 7

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-99ba8ab2d40498e8cfc622e05d5a6c5545794388ce5a37c886afd599690f2908)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4043a61c76f60bd9f8282bb063449589ba6ab11a14eef0273b778298034e647a"></a>

## client_side_defense.policy.js_insertion_rules — client_side_defense.policy.js_insertion_rules / 58edb5304ef6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- client_side_defense.policy.js_insertion_rules

<a id="canonical-0e6e587a63f4d3e23c59c0efdd0d1da2da11a97e451ad988ddbc0031d92f3dcb"></a>

Type: `"object"`. single nested block, Optional.

Defines custom JavaScript insertion rules for Client-Side Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Client-Side Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-cd455cf9787f06ec28f408a168a01a5d3f82989a82fa340be077e8034e30f1ea"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules / 58edb5304ef6 / 3

- [exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-38b6a07d0368bb1576c6340f308ade2c951be44eab40e8b86914c4817e0cfb2d): complete subsection reference.

- [rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-527813129cd3992956c0121fbee75a3816f8aea0c407cee7f7f67886eedf243d): complete subsection reference.

<a id="canonical-5efa382fceaf1aca84a7391efd3e52f28f012fbe5ed6edd8276a03d012b76610"></a>

## Next pages — client_side_defense.policy.js_insertion_rules / 58edb5304ef6 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-38b6a07d0368bb1576c6340f308ade2c951be44eab40e8b86914c4817e0cfb2d)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-527813129cd3992956c0121fbee75a3816f8aea0c407cee7f7f67886eedf243d)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-38b6a07d0368bb1576c6340f308ade2c951be44eab40e8b86914c4817e0cfb2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48b7e0a4e7a3a1be7e2131f725f859f23fbf3c7d76ecb242817dff9ee2b0200d"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list — client_side_defense.policy.js_insertion_rules.exclude_list / c380f2a233e5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233)
- client_side_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-0f2494ccd4e18c045a17322909f2c2cb3a3aae920a35afa097971a5cf5bc006f"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-82260dac0494c9735393915a54ca8f79a3070bf956151ad120df4f28a067d891"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list / c380f2a233e5 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-96eb00fa6d36257b9bc21dbb698ffa7b3b54a8ae724101e43db55240039e9ebe): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-ba6ced596ac43cab61039db789eed92b321a7f47f2b893ab23c11e906218d519): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-009.md#canonical-3d7b17930953064da94fe93ae07f3532fb8595850bea43de1e9c05f307f9a078): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-009.md#canonical-3d305cf0750c24484c0b39ce7d67161eaf98f706b0808b013b10a52202a98132): complete subsection reference.

<a id="canonical-2ad6cd4bc51124654d2b58cc0b603818ffa62dae4fcb5166a72ea490c4f19c8d"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list / c380f2a233e5 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list.any_domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-96eb00fa6d36257b9bc21dbb698ffa7b3b54a8ae724101e43db55240039e9ebe)
- [client_side_defense.policy.js_insertion_rules.exclude_list.domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-ba6ced596ac43cab61039db789eed92b321a7f47f2b893ab23c11e906218d519)
- [client_side_defense.policy.js_insertion_rules.exclude_list.metadata](resources--cdn_loadbalancer--reference--group-009.md#canonical-3d7b17930953064da94fe93ae07f3532fb8595850bea43de1e9c05f307f9a078)
- [client_side_defense.policy.js_insertion_rules.exclude_list.path](resources--cdn_loadbalancer--reference--group-009.md#canonical-3d305cf0750c24484c0b39ce7d67161eaf98f706b0808b013b10a52202a98132)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-96eb00fa6d36257b9bc21dbb698ffa7b3b54a8ae724101e43db55240039e9ebe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6832b8b3cf049014bdf2f8955153c1de0ab9e091d0e5b39bf451e85aecf66afc"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.any_domain — client_side_defense.policy.js_insertion_rules.exclude_list.any_domain / 475038910468 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-38b6a07d0368bb1576c6340f308ade2c951be44eab40e8b86914c4817e0cfb2d)
- client_side_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-1afffe624154bca8280b4d06b1995bfc2c8163ad44f5f1c74d9d7f948188be99"></a>

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

<a id="canonical-3bde4165571b9786a9f936b7829fc5faa1f75a212da9aa7c0aa37b92d3d6f631"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.any_domain / 475038910468 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-300ccf0e32be4470c75feaaf827243371633cd54e704c0478214ba4ab9857ead"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.any_domain / 475038910468 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-38b6a07d0368bb1576c6340f308ade2c951be44eab40e8b86914c4817e0cfb2d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ba6ced596ac43cab61039db789eed92b321a7f47f2b893ab23c11e906218d519"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-720b3d24533a76ef5cc49379a8fd15be516e88d25836075835864c881fec7886"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.domain — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 033d005c95cf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-38b6a07d0368bb1576c6340f308ade2c951be44eab40e8b86914c4817e0cfb2d)
- client_side_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-b2aaf81bb46e1a18fa0dba7d0052b106feb601a952f4f34b8ac3c23d7fe035d9"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-8c4f71bd78be4c70fb71dbd2abbbf75d04d085073bef7c060bc7ea4256d97bf8"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 033d005c95cf / 3

<a id="canonical-a37bd40d2e8ce1031a8dfb48e6f8656114eb6d43cf74669099b32f1522e18386"></a>

<a id="canonical-b1070c1c660bdd53bbde1dc81fa0f5ecbbd158dbdcf794d3aadb5aa6592ce84b"></a>

## exact_value property — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 033d005c95cf / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-c3a95ace27225668a7321fce2764cd1449fdca089579333bc13e382a7306dbce"></a>

<a id="canonical-2b190796d411d5ccfc5984d22c1b3811bbd447d37e8decb49559b5d488282693"></a>

## regex_value property — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 033d005c95cf / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-f0c82d02b9b5461e5ffbfb753a2db14f58c8af8e4f60b8c7c0accb45e9270513"></a>

<a id="canonical-a16fc6bc6072fe16aae951186491cde0ecc91bb94cc61031795a4f8ee8a41daf"></a>

## suffix_value property — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 033d005c95cf / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-19cce7ee8ecff3444e247e7966dccdd8e6427146f86d4788d8fcbd52bb94ac13"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 033d005c95cf / 7

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-38b6a07d0368bb1576c6340f308ade2c951be44eab40e8b86914c4817e0cfb2d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3d7b17930953064da94fe93ae07f3532fb8595850bea43de1e9c05f307f9a078"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86db3a622f1fc709efa825da21c7792a0e71353ca2ee19e283e129939675e27d"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.metadata — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 289b444b8a7b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-38b6a07d0368bb1576c6340f308ade2c951be44eab40e8b86914c4817e0cfb2d)
- client_side_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-be66a3b5fb72fe993eab8a8dcefb211c0c52b25baca2a94cc45f4c70bfba5866"></a>

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

<a id="canonical-3dce923914d0fcd7e3e7586fc4a5aaef1267636bd84687d625499efbc0c17569"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 289b444b8a7b / 3

<a id="canonical-8fa81d139424d9d1ca0385e867e961f9eef26c5d6f2dfb1f5096cd6f6b2ff42c"></a>

<a id="canonical-e53ea860871eaff67158927ffa5f1c8b984a51aee6ef7b999da5a3f2a9f11e4c"></a>

## description_spec property — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 289b444b8a7b / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-5cb3ed13c5cfd3334100234930bb1aa5ebf92ae6ff3786ced28ac08cac7c6dd1"></a>

<a id="canonical-61cdbe06a2ca1a21f727fce4efb41150f643a53bcc21db002d0b8fb01b6e7522"></a>

## name property — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 289b444b8a7b / 5

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

<a id="canonical-e36debcfd4b46a85afb17db32172e7422a300bdf19e6a3e915729a32a2bf94fe"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 289b444b8a7b / 6

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-38b6a07d0368bb1576c6340f308ade2c951be44eab40e8b86914c4817e0cfb2d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3d305cf0750c24484c0b39ce7d67161eaf98f706b0808b013b10a52202a98132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6174fa45de493003b9f71734941e73d170768b6c47fc7f94ce487a27148f95e5"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.path — client_side_defense.policy.js_insertion_rules.exclude_list.path / e707a75892d5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-38b6a07d0368bb1576c6340f308ade2c951be44eab40e8b86914c4817e0cfb2d)
- client_side_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-44ce3034f196acf1ec184cded62472c976eb1497f04c1435a0335565adde9fce"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-45c8c763d87dbce560423c9cbce99bd1daa23a5308fc3ff0a7e44659b57f6527"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.path / e707a75892d5 / 3

<a id="canonical-232c32d89c0cb25af43a82b6ea9fc9710211cb94db7568c17d1320697d96888d"></a>

<a id="canonical-826fd591adccabec596988561e36ac732671e422d10a05e4aa9380b047d57cbf"></a>

## path property — client_side_defense.policy.js_insertion_rules.exclude_list.path / e707a75892d5 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-e7d02078f6cbd6daa1584ffcb3bd7e023c279fc16c0abbdc9fc02a1dcca0015c"></a>

<a id="canonical-ff0ff85b16db8c6b7f02ea574356919076207ff0e6a6c5c4ddc771a4fc9225fc"></a>

## prefix property — client_side_defense.policy.js_insertion_rules.exclude_list.path / e707a75892d5 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-a113b14401082c68b5dfb84d2920ec5e18c6dee9297736b13936d2a5961016da"></a>

<a id="canonical-968a2f3da2e5d58b7fe22165efba03cc73954f48d021330d8e014968df713aed"></a>

## regex property — client_side_defense.policy.js_insertion_rules.exclude_list.path / e707a75892d5 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-cc3254cfcc468a4e52aae06230429dd6e0a7c57463bfaf72b9981fb06066092a"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.path / e707a75892d5 / 7

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-38b6a07d0368bb1576c6340f308ade2c951be44eab40e8b86914c4817e0cfb2d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-527813129cd3992956c0121fbee75a3816f8aea0c407cee7f7f67886eedf243d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-415baeaec320d3f2dc6055c3e94009508be031fdc259ac5a10bbfb21ff313d06"></a>

## client_side_defense.policy.js_insertion_rules.rules — client_side_defense.policy.js_insertion_rules.rules / 185cc19a5335 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233)
- client_side_defense.policy.js_insertion_rules.rules

<a id="canonical-97b50718a33b8b23ee0abb43a0e6ed78073b9c8c21a394c1cd0f0b0125e7e747"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Client-Side Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1ad794cadcdcdb0e610a2fca1f9d5d1f82671451493dfcf5c7b51c51b52bd432"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules / 185cc19a5335 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-94498085de4eb3ab6bb25fb13f0265f0d516b6f0dde01934f57bb64e0bd04e83): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-0a0eddea6c3d25c37138894a5cf823af8a43fe33608629ef18760165392fe988): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-009.md#canonical-ea1eaed037ca921d49897d19996c85b6dc656b0d1b03a655a9587c960672aaa7): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-009.md#canonical-83fca8cd595b32aebde8dae65b4c89c4644022d2a58b2cb0607712cf7542f800): complete subsection reference.

<a id="canonical-8ee324511997747e48cf0e0785b9e93d209ae297c8d9f279250393019b0a3f40"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules / 185cc19a5335 / 4

- [client_side_defense.policy.js_insertion_rules.rules.any_domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-94498085de4eb3ab6bb25fb13f0265f0d516b6f0dde01934f57bb64e0bd04e83)
- [client_side_defense.policy.js_insertion_rules.rules.domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-0a0eddea6c3d25c37138894a5cf823af8a43fe33608629ef18760165392fe988)
- [client_side_defense.policy.js_insertion_rules.rules.metadata](resources--cdn_loadbalancer--reference--group-009.md#canonical-ea1eaed037ca921d49897d19996c85b6dc656b0d1b03a655a9587c960672aaa7)
- [client_side_defense.policy.js_insertion_rules.rules.path](resources--cdn_loadbalancer--reference--group-009.md#canonical-83fca8cd595b32aebde8dae65b4c89c4644022d2a58b2cb0607712cf7542f800)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-94498085de4eb3ab6bb25fb13f0265f0d516b6f0dde01934f57bb64e0bd04e83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09f1b5df2039d3e84ee1f2c2106e8c64c8e131b9104d6e86216b65432ba64c05"></a>

## client_side_defense.policy.js_insertion_rules.rules.any_domain — client_side_defense.policy.js_insertion_rules.rules.any_domain / dcff5e5fa8a4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-527813129cd3992956c0121fbee75a3816f8aea0c407cee7f7f67886eedf243d)
- client_side_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-c383db6d8fb637f01386ae0fe99bbcaacb8e205da95a44cbf714ec3e0ddcbd8d"></a>

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

<a id="canonical-8b074e67d214919f3d396e08154f9656f0ed71a9f02635e4d85a677d17ad4d47"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.any_domain / dcff5e5fa8a4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a530dfe2b2b071e1409bf5515d240218e2f752ba9bbc067389e89a03db663c6e"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.any_domain / dcff5e5fa8a4 / 4

- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-527813129cd3992956c0121fbee75a3816f8aea0c407cee7f7f67886eedf243d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0a0eddea6c3d25c37138894a5cf823af8a43fe33608629ef18760165392fe988"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da2b3065a760f0333d979b82acf6a0afce5ffe98ae1490ed2bf995652f78e0b2"></a>

## client_side_defense.policy.js_insertion_rules.rules.domain — client_side_defense.policy.js_insertion_rules.rules.domain / abe310ec7c93 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-527813129cd3992956c0121fbee75a3816f8aea0c407cee7f7f67886eedf243d)
- client_side_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-d6383a281e2f9d87beab23e004cd4fadadc12ba3498ad040adeb438996a8681f"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-8f358e4ecd21fd19e2caee135236cdf3f59d75682bb5d3580846b3ca6563e648"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.domain / abe310ec7c93 / 3

<a id="canonical-2587f35b198be76f99c9cc415cec79ac1221fb298bf705100df414644f8245ec"></a>

<a id="canonical-f9cf7a208d12e7860e601da7f99d58b771a91f86826f6a336e7eba409e43d47c"></a>

## exact_value property — client_side_defense.policy.js_insertion_rules.rules.domain / abe310ec7c93 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-ae6458b6cdb87b88e509f3ad93aba0b97887cfffa7691116a58a2c5cba8d1816"></a>

<a id="canonical-7f683543001755a1c3227fb53d9706901ab371c0779f170992f48d8b5909683f"></a>

## regex_value property — client_side_defense.policy.js_insertion_rules.rules.domain / abe310ec7c93 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-a29f2c60945f784fd59f3eff1e8345ca41ed2a5cd6a813226975e3523872a5ce"></a>

<a id="canonical-39eddb5229b96df28521c4829329e72a8790073007017fa97b1ac618d7fab23d"></a>

## suffix_value property — client_side_defense.policy.js_insertion_rules.rules.domain / abe310ec7c93 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-2df4007c7ee4f8c4d86b4848e56f3f981cfe6d2b160153ccff592247b69dedd2"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.domain / abe310ec7c93 / 7

- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-527813129cd3992956c0121fbee75a3816f8aea0c407cee7f7f67886eedf243d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ea1eaed037ca921d49897d19996c85b6dc656b0d1b03a655a9587c960672aaa7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5052f284b2d3e29c2bf8c3e1a266b0eb6b24132207c1b1893df177eac415c7d5"></a>

## client_side_defense.policy.js_insertion_rules.rules.metadata — client_side_defense.policy.js_insertion_rules.rules.metadata / 9fc747b1ec73 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-527813129cd3992956c0121fbee75a3816f8aea0c407cee7f7f67886eedf243d)
- client_side_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-d695706fdabb720c58c2fa52c9cc08ffd81d84c1cc25bf207243d7a1bc60ac44"></a>

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

<a id="canonical-aca436dc136568b0a0cfa12c2a06d9bc28c135d1fb5bdf6508bbf4a3dc3d53f2"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.metadata / 9fc747b1ec73 / 3

<a id="canonical-4f809c53c3e10f6e70caea58f560e98c583fe71efa41588aa0268611a02eaaed"></a>

<a id="canonical-f4ef15e669eaa786142dd93ddf693f4c503fa4943de33871319ccbac30e372d1"></a>

## description_spec property — client_side_defense.policy.js_insertion_rules.rules.metadata / 9fc747b1ec73 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1314bcf95f10e26252d2511bdb396b02422b481da29d8fe0376007a9c32a1db3"></a>

<a id="canonical-0f68b9658a1008035dcd854f42d6b911d1975ad49b5ffb1abbdb20ce02082247"></a>

## name property — client_side_defense.policy.js_insertion_rules.rules.metadata / 9fc747b1ec73 / 5

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

<a id="canonical-7f398d5e2a73dfc16db658c810b1a5164916ed422f4cab2ad5c16ef48fac86a2"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.metadata / 9fc747b1ec73 / 6

- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-527813129cd3992956c0121fbee75a3816f8aea0c407cee7f7f67886eedf243d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-83fca8cd595b32aebde8dae65b4c89c4644022d2a58b2cb0607712cf7542f800"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0405d5be71bcd7b4ee6030595ed45c1c5c8dc73735fd6503322b4489ccbe394a"></a>

## client_side_defense.policy.js_insertion_rules.rules.path — client_side_defense.policy.js_insertion_rules.rules.path / 50b380b891e5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-92d9bacb40a57f2a71d3c5e4eb4d78f1a1e03ed61a6b91ef2acc7469ee69bed6)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-fedbe63d1beb7c43954a2cbbdea90174f3e62b81acc2e2b67b61dfd634c59745)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0955b819554ba01a9965cf8e4369105009bbb24d3e2ceba2643b89f3c54ad233)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-527813129cd3992956c0121fbee75a3816f8aea0c407cee7f7f67886eedf243d)
- client_side_defense.policy.js_insertion_rules.rules.path

<a id="canonical-4ae260427d79a8a3e5198d14a44705d17574e91331023c7cc2d251048f0772bc"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-5ec180842ae3d4a9c694a4faf7d38593a9245f5619d4a4dd34ebc2bbcf95f1dc"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.path / 50b380b891e5 / 3

<a id="canonical-5e74df0d5184d45ea645f3572e384c37dd4c58154846c30f73dfb8c02d8b54f4"></a>

<a id="canonical-0afd02949408e30a74390cacfca88d930d57b3148fe47412c26e3ee589185b70"></a>

## path property — client_side_defense.policy.js_insertion_rules.rules.path / 50b380b891e5 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-e67b55549ac74fbb4601b5c3d9661d5c6992d92aa8344599c13acb2139ba934f"></a>

<a id="canonical-b28a34481da4cfd0534ed5c00ac7aa87f71a3cd8fc591cf66e56d276327780cc"></a>

## prefix property — client_side_defense.policy.js_insertion_rules.rules.path / 50b380b891e5 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-2b608b03dc3ba3f1a01bf3e0639c10076f68a23473139dbc31f9cf660db8a809"></a>

<a id="canonical-099b040f206e43584fae8aa604a1333da9384f3e737a8c5e91f8dd182f844f63"></a>

## regex property — client_side_defense.policy.js_insertion_rules.rules.path / 50b380b891e5 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-b1fea8b3496deb41c3c26096387fdd175fbf5badd28a251765630c8ac92d0ea0"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.path / 50b380b891e5 / 7

- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-527813129cd3992956c0121fbee75a3816f8aea0c407cee7f7f67886eedf243d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-61274ef5700bf47a6640898b8f6b648c3daecd806a5fa82ea3a5c5d02fec2f66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43fdd135545314b3681126599bb8d060153325b2d66d0bbaf07d25c10f32e11a"></a>

## cors_policy — cors_policy / 9d1c1a5adcaf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- cors_policy

<a id="canonical-1e6edd28ac48121286f4014129ac57c8fa59c1075733af2a5333710b31856121"></a>

Type: `"object"`. single nested block, Optional.

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence. An example of an Cross origin HTTP request GET
/resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS
X 10.5..

Upstream description:

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence.

An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130
Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8
Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset:
ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive Referrer:
http&#58;//foo.example/examples/access-control/simplexsinvocation.html Origin:
http&#58;//foo.example

HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61
Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive
Transfer-Encoding: chunked Content-Type: application/XML

An example for cross origin HTTP OPTIONS request with Access-Control-Request-\* header

OPTIONS /resources/POST-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel
MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept:
text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8 Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive
Origin: http&#58;//foo.example Access-Control-Request-Method: POST Access-Control-Request-Headers:
X-PINGOTHER, Content-Type

HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http&#58;//foo.example Access-Control-Allow-Methods: POST, GET, OPTIONS
Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary:
Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive.

Receipt-pinned upstream constraints:

```json
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
cors_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-8a3ecb8f10efc89e8da539301cb9037c3cbc8855d470b0d9e84fde54adc7c805"></a>

## Direct properties — cors_policy / 9d1c1a5adcaf / 3

<a id="canonical-0a2e700d813310d69cbf23081771e529f968bb28743fc5ba760b1a0402a61e3f"></a>

<a id="canonical-9fe463dce70279729ff3028505879ce95da3b949e62fe758a4332f5e8903eb54"></a>

## allow_credentials property — cors_policy / 9d1c1a5adcaf / 4

Type: `"bool"`. Optional.

Specifies whether the resource allows credentials.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-be3ce3eb7b4529d3d894aed428c1a94355cb94aeba1b1cd40f2d640662d00e25"></a>

<a id="canonical-f5d6446eb48973a2e0e71de1342b590652442bf1ce23c2ddf8c872f0cc1045ee"></a>

## allow_headers property — cors_policy / 9d1c1a5adcaf / 5

Type: `"string"`. Optional.

Specifies the content for the access-control-allow-headers header.

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

<a id="canonical-108fbbf76e19ff573d6b2f928cfd086751d48025de4f0113687b9c1d38d597fd"></a>

<a id="canonical-932763b2e10d6fc049a918c8c5e753f2ebf66c3926a3d87dd9eced707526dcdd"></a>

## allow_methods property — cors_policy / 9d1c1a5adcaf / 6

Type: `"string"`. Optional.

Specifies the content for the access-control-allow-methods header.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-900193deda8883dd1d4a0537a1b2a62d70e57fef8d8aa69c3f9ddcc3c364553c"></a>

<a id="canonical-86413784ee7278a0f2bb335117ca0d2d23b15448de848dad86ba7f3569337b6a"></a>

## allow_origin property — cors_policy / 9d1c1a5adcaf / 7

Type: `["list", "string"]`. Optional.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Upstream description:

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-bd88d78e862ad378a574a6c6f183fb4d79a23167a3445c991bf16f258d5ca2d6"></a>

<a id="canonical-d336c7a6429a04a8491668e3400f5f25c62d53f0297e9428806bc5c8750c3169"></a>

## allow_origin_regex property — cors_policy / 9d1c1a5adcaf / 8

Type: `["list", "string"]`. Optional.

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Upstream description:

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-49b234aba461a5253df8a6ddd288cbfc971946a38583f1458083eb7a64224676"></a>

<a id="canonical-a0797d61137918694c762a46b827b5ac4df4ae7267e7d07c6d3d068af9a2ef76"></a>

## disabled property — cors_policy / 9d1c1a5adcaf / 9

Type: `"bool"`. Optional.

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-449d2e2fcce412fa60f5ce9381fde0a149afeb18d4b32f1b3d02c59c4d3e1eaa"></a>

<a id="canonical-22320421fb0b0ee6533eb8eecc3103e56bed8560fa06bf679e876f0dc7dea90d"></a>

## expose_headers property — cors_policy / 9d1c1a5adcaf / 10

Type: `"string"`. Optional.

Specifies the content for the access-control-expose-headers header.

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

<a id="canonical-71018cad4567063976f6c91f15c78a6a96e73497bf51f890d582bfeb7eee4205"></a>

<a id="canonical-32cde5a5b3301dc508c776f133c6478c938ffa339ac338d724b800bc59a8c41d"></a>

## maximum_age property — cors_policy / 9d1c1a5adcaf / 11

Type: `"number"`. Optional.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Upstream description:

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(-1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": -1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  }
}
```

<a id="canonical-ccc4415c6db7b9b1d5728df04435fceee5690d69b9446741ea557233c297ceb4"></a>

## Next pages — cors_policy / 9d1c1a5adcaf / 12

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-407303fdc296fc1cb771072072f9f38707eff9e1eb19854190f5f811b1803371"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-006381a56dc5548467868dccfa8f9fbd92ff0c6d2a791be847d2eb55916554c7"></a>

## csrf_policy — csrf_policy / 99d9bd7a353d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- csrf_policy

<a id="canonical-6da8f3a8ba23fbf2e6b3b9bc819057f67208ea56559901a2cc35367711f53958"></a>

Type: `"object"`. single nested block, Optional.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host. 1. The origin that caused the user
agent..

Upstream description:

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host.

&#8203;1. The origin that caused the user agent to issue the request (source origin). &#8203;2. The
origin that the request is going to (target origin). When the policy evaluating a request, it
ensures both pieces of information are present and compare their values. If the source origin is
missing or origins do not match the request is rejected. The exception to this being if the
source-origin has been added to they policy as valid. Because CSRF attacks specifically target
state-changing requests, the policy only acts on the HTTP requests that have state-changing method
(PUT,POST, etc.).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "custom_domain_list"),
  validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "disabled"),
  validators.ConflictingObjectAttributes("custom_domain_list",
    "disabled")}
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
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

Terraform syntax:

```terraform
csrf_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0ebf35e8014ffce602c7f7b432a9269b02fbb70a94a2e43c77b81f7354608bd1"></a>

## Direct properties — csrf_policy / 99d9bd7a353d / 3

- [all_load_balancer_domains](resources--cdn_loadbalancer--reference--group-009.md#canonical-f5785bb9c5e7343719b2a97bf23afc1ea62f379b0c863476533c6655c7b85d65): complete subsection reference.

- [custom_domain_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-278a1bc91352081a8cb2ef23e6e9b261d9333981bd1bf46ca4aceff152c4c496): complete subsection reference.

- [disabled](resources--cdn_loadbalancer--reference--group-009.md#canonical-dfd27067227ce6f1d86d6fb1a28893d8913178133acd41315de076e4945de847): complete subsection reference.

<a id="canonical-1ca569625ef174a5229f32240a9a4381a8aeff82ce3e23e1a327bb9961ef3cf3"></a>

## Next pages — csrf_policy / 99d9bd7a353d / 4

- [csrf_policy.all_load_balancer_domains](resources--cdn_loadbalancer--reference--group-009.md#canonical-f5785bb9c5e7343719b2a97bf23afc1ea62f379b0c863476533c6655c7b85d65)
- [csrf_policy.custom_domain_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-278a1bc91352081a8cb2ef23e6e9b261d9333981bd1bf46ca4aceff152c4c496)
- [csrf_policy.disabled](resources--cdn_loadbalancer--reference--group-009.md#canonical-dfd27067227ce6f1d86d6fb1a28893d8913178133acd41315de076e4945de847)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f5785bb9c5e7343719b2a97bf23afc1ea62f379b0c863476533c6655c7b85d65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8563c88c985dcee13f6b026b2c6674f2c01e589afacb0f4df59a05b8e1336ce6"></a>

## csrf_policy.all_load_balancer_domains — csrf_policy.all_load_balancer_domains / 7ba1d965ab21 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [csrf_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-407303fdc296fc1cb771072072f9f38707eff9e1eb19854190f5f811b1803371)
- csrf_policy.all_load_balancer_domains

<a id="canonical-3b9f814209fe9677cd1573c8dec9c2d0fd35a250352debfcb5fe03cef3c1771d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all load balancer domains.

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
all_load_balancer_domains = {}
```

<a id="canonical-c1b898be96beefb5d97d589244ba2a7a7e8b09d52b5563284d11486c0b02fe42"></a>

## Direct properties — csrf_policy.all_load_balancer_domains / 7ba1d965ab21 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e83f17a7b26225a3b44b3c072ebad557018b8a772c2cab0a92ac8370a4b07ef7"></a>

## Next pages — csrf_policy.all_load_balancer_domains / 7ba1d965ab21 / 4

- [csrf_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-407303fdc296fc1cb771072072f9f38707eff9e1eb19854190f5f811b1803371)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-278a1bc91352081a8cb2ef23e6e9b261d9333981bd1bf46ca4aceff152c4c496"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63bfba29b8aa91d087e09a801372ce9c78fa01b5c9e16bcad10a9f6f677e6bdb"></a>

## csrf_policy.custom_domain_list — csrf_policy.custom_domain_list / a9768fcbd8d0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [csrf_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-407303fdc296fc1cb771072072f9f38707eff9e1eb19854190f5f811b1803371)
- csrf_policy.custom_domain_list

<a id="canonical-9c0859437d73bbadf858135ba7c87bb7cd51c4966f2d759c0053089fd852b6b0"></a>

Type: `"object"`. single nested block, Optional.

List of domain names used for Host header matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
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
custom_domain_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-c4dcbe373bed850d193341c6e3e8c38240f8e94f5eed4ad4c4a8f4f6ace84eff"></a>

## Direct properties — csrf_policy.custom_domain_list / a9768fcbd8d0 / 3

<a id="canonical-d4d9b17d166bf3a425145009cbd8e985e31691ceb23d0c31218f49c3a29cb227"></a>

<a id="canonical-d529ea4c419dedf660bb9cef94d7a58ab6b7b8f78b8514a6bc634ba8aa43f3ef"></a>

## domains property — csrf_policy.custom_domain_list / a9768fcbd8d0 / 4

Type: `["list", "string"]`. Optional.

List of domain names that will be matched to loadbalancer. These domains are not used for SNI match.
Wildcard names are supported in the suffix or prefix form.

Upstream description:

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9bd4a6d48d1e64eb06412e475d19d1e39fa3d8b2865fbe252bee4d7f08fe1ba5"></a>

## Next pages — csrf_policy.custom_domain_list / a9768fcbd8d0 / 5

- [csrf_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-407303fdc296fc1cb771072072f9f38707eff9e1eb19854190f5f811b1803371)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-dfd27067227ce6f1d86d6fb1a28893d8913178133acd41315de076e4945de847"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94f96e61b6ae2e5b7fd4ab3dc6473e06c7bf5db72232b0536b3ba5b5afa3e95b"></a>

## csrf_policy.disabled — csrf_policy.disabled / f93d394e105d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [csrf_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-407303fdc296fc1cb771072072f9f38707eff9e1eb19854190f5f811b1803371)
- csrf_policy.disabled

<a id="canonical-ebfc75c7c6d925c18596bfcc91639360cfb7ec56857c10da443ab006a03a3e18"></a>

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
disabled = {}
```

<a id="canonical-12eba8c816007e752f30c615de640c479988cac3c3675205b05e80fc8bf7b6ba"></a>

## Direct properties — csrf_policy.disabled / f93d394e105d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-af49f2941295f0e52351ae2833ca2ad5c59ef7c77c853a521530f05ced1b4450"></a>

## Next pages — csrf_policy.disabled / f93d394e105d / 4

- [csrf_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-407303fdc296fc1cb771072072f9f38707eff9e1eb19854190f5f811b1803371)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a4f8e5202487add45d34bd5e279421532fa89e3abc164a114fdf93433b0f2134"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed8463ca893b5ac73cc89cfdb446b4b954fcb5545b5ac076b8a8179201395a53"></a>

## custom_cache_rule — custom_cache_rule / 468a8631d108 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- custom_cache_rule

<a id="canonical-42bf8b7ee32650674911a29bd06497468f1827d280534fc0081f3b8337650fae"></a>

Type: `"object"`. single nested block, Optional.

Custom Cache Rules. Caching policies for CDN.

Upstream description:

Caching policies for CDN.

Receipt-pinned upstream constraints:

```json
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
custom_cache_rule {
  # Configure direct properties listed below.
}
```

<a id="canonical-78f1564f3408e069dad9b0be70010245d56264bb9520e029cec6711e08bb8975"></a>

## Direct properties — custom_cache_rule / 468a8631d108 / 3

- [cdn_cache_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-ef76be86207e14e5a075afb2be6915c7c5b241a9ad10451c9a99092c7563c630): complete subsection reference.

<a id="canonical-4436040016d6ad4c3b7054b7ae795a79e746eb7d8905d4333bdfdb1d9e0dc064"></a>

## Next pages — custom_cache_rule / 468a8631d108 / 4

- [custom_cache_rule.cdn_cache_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-ef76be86207e14e5a075afb2be6915c7c5b241a9ad10451c9a99092c7563c630)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ef76be86207e14e5a075afb2be6915c7c5b241a9ad10451c9a99092c7563c630"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-077bb11e591ab710c256e65ff5438cbabd50e32f7d78cc7f45452ee06fc11e04"></a>

## custom_cache_rule.cdn_cache_rules — custom_cache_rule.cdn_cache_rules / 93df4128c5a6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [custom_cache_rule](resources--cdn_loadbalancer--reference--group-009.md#canonical-a4f8e5202487add45d34bd5e279421532fa89e3abc164a114fdf93433b0f2134)
- custom_cache_rule.cdn_cache_rules

<a id="canonical-dbd67867dbc6ffee5a3d647308e5c06c3aad07ae62b4ea9da2433fe37460aa18"></a>

Type: `"object"`. list nested block, Optional.

Reference to CDN Cache Rule configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
cdn_cache_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-816743802beee31203c6e8d92adde7d0672baeea5034a0d5ff73c37963bfd5f5"></a>

## Direct properties — custom_cache_rule.cdn_cache_rules / 93df4128c5a6 / 3

<a id="canonical-c3c56971d88f1cd3128eb546c227c5f52cb8d7e6d59912134576ea4c283d000d"></a>

<a id="canonical-fec16c2af66cdae6b4cabdf1e88b5661b09dadd966311c7db6c7b2d7f2db2e85"></a>

## name property — custom_cache_rule.cdn_cache_rules / 93df4128c5a6 / 4

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

<a id="canonical-1e933676738fb5a1558c5515f34cf2a30abe258a033ce485d994f9c028907c03"></a>

<a id="canonical-c09c812e6589d8d1a64c01937cf037cde99c8cf89d317eb5ccc36022be657349"></a>

## namespace property — custom_cache_rule.cdn_cache_rules / 93df4128c5a6 / 5

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

<a id="canonical-7e7fb649266e8af2cf5bcb888c151e9dd84528a2ac8aeff2e80a403110ff5b8f"></a>

<a id="canonical-8e771d17334713844a2f1027411e677052a4626c0362f52707ec3d4ef9360e35"></a>

## tenant property — custom_cache_rule.cdn_cache_rules / 93df4128c5a6 / 6

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

<a id="canonical-02f89aa43c7da30a4783644671f79955f584b70e5ca5c1c84c7f56912ee0a129"></a>

## Next pages — custom_cache_rule.cdn_cache_rules / 93df4128c5a6 / 7

- [custom_cache_rule](resources--cdn_loadbalancer--reference--group-009.md#canonical-a4f8e5202487add45d34bd5e279421532fa89e3abc164a114fdf93433b0f2134)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1dccf7057060794e294ec86ddd031213e249d14c153ed77e26fdb890afbba54f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c50fca06ea8f00b444c9daf31e4a466b543967f659b99ef39e0e6b4bce5841d"></a>

## data_guard_rules — data_guard_rules / 13be0b0bc838 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- data_guard_rules

<a id="canonical-359469d762bb3a6392c40e97a1940ccd4eb5e57e65e7a238ae21ed82e3d820f0"></a>

Type: `"object"`. list nested block, Optional.

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*).

Upstream description:

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*). Note: App Firewall should be enabled, to use Data
Guard feature.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("apply_data_guard",
    "skip_data_guard"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value")}
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
data_guard_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-30268e8c0b8b4afc31710af8375d627a9bad3e874c873d87d6b8f183bc5077de"></a>

## Direct properties — data_guard_rules / 13be0b0bc838 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-192e5da5702b469a4fc8f99776bebe1bea9e913b4e948cf642bb36ba24edbc5a): complete subsection reference.

- [apply_data_guard](resources--cdn_loadbalancer--reference--group-009.md#canonical-61a9f19244349b263d731223d891edbd72824e2f538f6dc99c827e3f0658a389): complete subsection reference.

<a id="canonical-5285d656968ec6190e870311702b72b63d4d79db42a82311e9e08e6d755388e8"></a>

<a id="canonical-c5a287106449097da572f37c3e5ccaa49c5b22c79b005c42eb4106088bb65bad"></a>

## exact_value property — data_guard_rules / 13be0b0bc838 / 4

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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

- [metadata](resources--cdn_loadbalancer--reference--group-009.md#canonical-527cf676226d2098ba97d9c75588e12a0ae936c141380cc3f1cc03fd3f3c294c): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-009.md#canonical-a9e61be850a080f6c090d80692c7ce8dac312f32f78db3848f9c967bd2335be6): complete subsection reference.

- [skip_data_guard](resources--cdn_loadbalancer--reference--group-009.md#canonical-6de85f2e03329a37173466f5a49f9e4a73c8a32622fd7f1ecb54d7fe908a04fa): complete subsection reference.

<a id="canonical-eb44eccafb91a01772c1671061b6046a5a6aa1bcdc5b520b7c2db118e5c828a3"></a>

<a id="canonical-01a71661afe4e4ecfb504383cb11dcd5706f2dff596a4037f0de22af41e42138"></a>

## suffix_value property — data_guard_rules / 13be0b0bc838 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-c09f75f792917ca2caf1ca83c7a3810f9bab63be0b2c7daa4b9eddd05f587f6c"></a>

## Next pages — data_guard_rules / 13be0b0bc838 / 6

- [data_guard_rules.any_domain](resources--cdn_loadbalancer--reference--group-009.md#canonical-192e5da5702b469a4fc8f99776bebe1bea9e913b4e948cf642bb36ba24edbc5a)
- [data_guard_rules.apply_data_guard](resources--cdn_loadbalancer--reference--group-009.md#canonical-61a9f19244349b263d731223d891edbd72824e2f538f6dc99c827e3f0658a389)
- [data_guard_rules.metadata](resources--cdn_loadbalancer--reference--group-009.md#canonical-527cf676226d2098ba97d9c75588e12a0ae936c141380cc3f1cc03fd3f3c294c)
- [data_guard_rules.path](resources--cdn_loadbalancer--reference--group-009.md#canonical-a9e61be850a080f6c090d80692c7ce8dac312f32f78db3848f9c967bd2335be6)
- [data_guard_rules.skip_data_guard](resources--cdn_loadbalancer--reference--group-009.md#canonical-6de85f2e03329a37173466f5a49f9e4a73c8a32622fd7f1ecb54d7fe908a04fa)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-192e5da5702b469a4fc8f99776bebe1bea9e913b4e948cf642bb36ba24edbc5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c45ce3bcf7ea91a47143edcce1a59598e571de2f1337bcea1ac316c14fe13a76"></a>

## data_guard_rules.any_domain — data_guard_rules.any_domain / d5b4695cf321 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [data_guard_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1dccf7057060794e294ec86ddd031213e249d14c153ed77e26fdb890afbba54f)
- data_guard_rules.any_domain

<a id="canonical-e01342d4616f7675add2719564951398bb6501dbc9b26708663966add8847dae"></a>

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

<a id="canonical-e809e69429ff61ead6b6d8cef2810c77faa11d37962abd5757c25d7a712b6599"></a>

## Direct properties — data_guard_rules.any_domain / d5b4695cf321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-55b3eeefe447b1183601900c2c68ddf07736f7d7f7e9d5e24f4bbf9af3dd545e"></a>

## Next pages — data_guard_rules.any_domain / d5b4695cf321 / 4

- [data_guard_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1dccf7057060794e294ec86ddd031213e249d14c153ed77e26fdb890afbba54f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-61a9f19244349b263d731223d891edbd72824e2f538f6dc99c827e3f0658a389"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45024f910ed61c4207d9c3620cb91702b2e42994ea54ad9a12bbdcebcb70db90"></a>

## data_guard_rules.apply_data_guard — data_guard_rules.apply_data_guard / fde14fadcf38 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [data_guard_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1dccf7057060794e294ec86ddd031213e249d14c153ed77e26fdb890afbba54f)
- data_guard_rules.apply_data_guard

<a id="canonical-b5cd16f27673382a82c2bc068eec3897c47ddb203d1d7f241b28ede35b5a3f94"></a>

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
apply_data_guard = {}
```

<a id="canonical-01eb67cee3d3a8a3cd9b7746d4e6fad07da69fb97f9574132c4c8babdabad23d"></a>

## Direct properties — data_guard_rules.apply_data_guard / fde14fadcf38 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4b1cd2a2576a475fe61f811e09de35a59bb5e95d8fb3a2482d229ce794683eda"></a>

## Next pages — data_guard_rules.apply_data_guard / fde14fadcf38 / 4

- [data_guard_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1dccf7057060794e294ec86ddd031213e249d14c153ed77e26fdb890afbba54f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-527cf676226d2098ba97d9c75588e12a0ae936c141380cc3f1cc03fd3f3c294c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0fe4645c4310ab2d649e68b27c11e2d1078b821b6f36cec4f9a03d8b4c0f4cc"></a>

## data_guard_rules.metadata — data_guard_rules.metadata / 34abc4c4e600 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [data_guard_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1dccf7057060794e294ec86ddd031213e249d14c153ed77e26fdb890afbba54f)
- data_guard_rules.metadata

<a id="canonical-82060ec05cdbd0a3950d908d5a86d0063040fbb23b29f3a387633705e9d04785"></a>

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

<a id="canonical-030e51bbb9b8fa9653baeaa852322f9946da3edbb91ca879af59abf09637b6b1"></a>

## Direct properties — data_guard_rules.metadata / 34abc4c4e600 / 3

<a id="canonical-2014aab8b63ed3fb557ea1607d38f8f885c081e2a7729553714dc337f124b89c"></a>

<a id="canonical-b3d414239074836f43667d44e050ec722020e7e1e429606c58b8e81fa58861d6"></a>

## description_spec property — data_guard_rules.metadata / 34abc4c4e600 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-7a0dc3cfde62ae979a0564b8a2d55d52ec5bbaac8ceee328ad87cb232ec50e1d"></a>

<a id="canonical-db5c91866977b39b0f1f5e2df7676d263b569850af24ff0661c06cfc37f32427"></a>

## name property — data_guard_rules.metadata / 34abc4c4e600 / 5

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

<a id="canonical-27dc36f6841f9d18477b1f620812ee2c6b4410f62d2056b12b03eb66e27f74b7"></a>

## Next pages — data_guard_rules.metadata / 34abc4c4e600 / 6

- [data_guard_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1dccf7057060794e294ec86ddd031213e249d14c153ed77e26fdb890afbba54f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a9e61be850a080f6c090d80692c7ce8dac312f32f78db3848f9c967bd2335be6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff1847e8723b7e39ae94c7f836cb2adaf0747487a04b7c0e07b1edec2ae8130e"></a>

## data_guard_rules.path — data_guard_rules.path / dd762c58108f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [data_guard_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1dccf7057060794e294ec86ddd031213e249d14c153ed77e26fdb890afbba54f)
- data_guard_rules.path

<a id="canonical-d6713192807642081d2a1dd11b3d5bd0b6031ab0cf02c1867fb41dd6d78a5cf5"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-d41d6c62c7a35188343dad15d507ff8621b4a9640dfbd93372ad3b69b6da05c8"></a>

## Direct properties — data_guard_rules.path / dd762c58108f / 3

<a id="canonical-bb4d30fc1fc01666f6934a372ef13ed0b09746a31bbe458ed3cd3be4602cd810"></a>

<a id="canonical-585ede41314e786c92bab3c82e199f9fda33d559069f6d7311c19d43ed9eef82"></a>

## path property — data_guard_rules.path / dd762c58108f / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-4721c65efdc5c34a0ba523fda1d3f002e9f17ea8e5c31c2d0266e04059a6bd0d"></a>

<a id="canonical-13f9cd553449771dcd66fb92668a32f59513fcc2b9cd38198d1a472c57615f16"></a>

## prefix property — data_guard_rules.path / dd762c58108f / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-78a898359b278fd030d6c61a11a44b63e28c71a77616b7b1c90b58cf0a2afeb5"></a>

<a id="canonical-46e1872921dc14454b129bd90da510e22cde0635f8a4db2de9a2f5fb79962c71"></a>

## regex property — data_guard_rules.path / dd762c58108f / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-bd36908e7b7591f252f407f321fa7dee50625e67a43d5b8f5f944c380a909d81"></a>

## Next pages — data_guard_rules.path / dd762c58108f / 7

- [data_guard_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1dccf7057060794e294ec86ddd031213e249d14c153ed77e26fdb890afbba54f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-6de85f2e03329a37173466f5a49f9e4a73c8a32622fd7f1ecb54d7fe908a04fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d610e3238e4a2624411d3537e51ad6665048e00790f3d619dbde6cc7e04d12f5"></a>

## data_guard_rules.skip_data_guard — data_guard_rules.skip_data_guard / e8ee82822241 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [data_guard_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1dccf7057060794e294ec86ddd031213e249d14c153ed77e26fdb890afbba54f)
- data_guard_rules.skip_data_guard

<a id="canonical-cfca4e760a25ca2d7da25816d9d52378ca6797ef6b2fa32b1f78851b06dff272"></a>

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
skip_data_guard = {}
```

<a id="canonical-9afc55625de7400c97012699d74bd8e1b72bf5510524972fa4d7d1f19d259824"></a>

## Direct properties — data_guard_rules.skip_data_guard / e8ee82822241 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-731fd3af10659dc99e9ab1e547ac5d42ec40a9777d8b303d764cf14dd257567e"></a>

## Next pages — data_guard_rules.skip_data_guard / e8ee82822241 / 4

- [data_guard_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1dccf7057060794e294ec86ddd031213e249d14c153ed77e26fdb890afbba54f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5746a5502dfc8e0895cb2a3ef2a31f60a71459706a1427f5bec63040fb069f07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d2513ac873c9c2c6123242992b473cf3ba42a67e1590ef4da854780b1c0107a"></a>

## ddos_mitigation_rules — ddos_mitigation_rules / 394ee84e395e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- ddos_mitigation_rules

<a id="canonical-4ac6d76661f90d0290c96dbf0419b349ad526b58a10d49d08a469d23e5ead022"></a>

Type: `"object"`. list nested block, Optional.

Define manual mitigation rules to block L7 DDoS attacks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ddos_client_source",
    "ip_prefix_list")}
```

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
    }
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
ddos_mitigation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-ab8ae0af55e6a6c0094e7dc773e383ea6d5ad540d64f808c6dd0f21bdfcbeffb"></a>

## Direct properties — ddos_mitigation_rules / 394ee84e395e / 3

- [block](resources--cdn_loadbalancer--reference--group-009.md#canonical-51e51744bb5a77078255ac5bf72cc5c816fd02abcb1d426aecaf10ea98944b5f): complete subsection reference.

- [ddos_client_source](resources--cdn_loadbalancer--reference--group-009.md#canonical-84a504b5b14092795e6fa5369922fa6b16e4fd9a7fa887303f90b8fb04294ab8): complete subsection reference.

<a id="canonical-e3b5be3466f209ef1f84a77bf9a18d7c2bceef62396b5bb671bbfe0729a99f14"></a>

<a id="canonical-1414cd8c85a800b96511dc5e04f8968f6b3f1896cf6a3bf8430e4affeb1ced00"></a>

## expiration_timestamp property — ddos_mitigation_rules / 394ee84e395e / 4

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [ip_prefix_list](resources--cdn_loadbalancer--reference--group-010.md#canonical-4e88f110db2623c28d213b72c36de4caca775a7f8faff5e4415068d1977e8db2): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-010.md#canonical-09837e2faeb7e137d6648a5128ae52925de69eb93c24b0f3acdeb1a88df11b86): complete subsection reference.

<a id="canonical-c455c4692b7d2c60f716e69617c75283e8c98b640b6ed35146cc34e844ad90f5"></a>

## Next pages — ddos_mitigation_rules / 394ee84e395e / 5

- [ddos_mitigation_rules.block](resources--cdn_loadbalancer--reference--group-009.md#canonical-51e51744bb5a77078255ac5bf72cc5c816fd02abcb1d426aecaf10ea98944b5f)
- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-009.md#canonical-84a504b5b14092795e6fa5369922fa6b16e4fd9a7fa887303f90b8fb04294ab8)
- [ddos_mitigation_rules.ip_prefix_list](resources--cdn_loadbalancer--reference--group-010.md#canonical-4e88f110db2623c28d213b72c36de4caca775a7f8faff5e4415068d1977e8db2)
- [ddos_mitigation_rules.metadata](resources--cdn_loadbalancer--reference--group-010.md#canonical-09837e2faeb7e137d6648a5128ae52925de69eb93c24b0f3acdeb1a88df11b86)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-51e51744bb5a77078255ac5bf72cc5c816fd02abcb1d426aecaf10ea98944b5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52e3c1ceed5e67bbb5ecc602dd7c97658ebbff3d6942aeddbbab295dc0bcfe01"></a>

## ddos_mitigation_rules.block — ddos_mitigation_rules.block / f5e7c12ed556 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-5746a5502dfc8e0895cb2a3ef2a31f60a71459706a1427f5bec63040fb069f07)
- ddos_mitigation_rules.block

<a id="canonical-b5c5810e83e2687a76068703a0ca00a55722ca9beb3f826eb7f4a1e5ccfde0f0"></a>

Type: `"object"`. single nested block, Optional.

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
block {}
```

<a id="canonical-3f07a25d8a1503f53580bd93306fa3575bdd8730c2c0bc476a17e181dfe44600"></a>

## Direct properties — ddos_mitigation_rules.block / f5e7c12ed556 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ec7d891587493ca66d49d93941225c7fbef6c8fb3d5810d36e44c7b14821f9b9"></a>

## Next pages — ddos_mitigation_rules.block / f5e7c12ed556 / 4

- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-5746a5502dfc8e0895cb2a3ef2a31f60a71459706a1427f5bec63040fb069f07)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-84a504b5b14092795e6fa5369922fa6b16e4fd9a7fa887303f90b8fb04294ab8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd435f01b5908a3bd0d6a19731d9523906646236c4ba7a59c65a78dc249a76fe"></a>

## ddos_mitigation_rules.ddos_client_source — ddos_mitigation_rules.ddos_client_source / c675fd2d9405 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-5746a5502dfc8e0895cb2a3ef2a31f60a71459706a1427f5bec63040fb069f07)
- ddos_mitigation_rules.ddos_client_source

<a id="canonical-85c4dd19035183af62c802d83876c5a503b6ad91bacb4334fa55c80b45b3a83c"></a>

Type: `"object"`. single nested block, Optional.

DDoS Client Source Choice. DDoS Mitigation sources to be blocked.

Upstream description:

DDoS Mitigation sources to be blocked.

Receipt-pinned upstream constraints:

```json
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
ddos_client_source {
  # Configure direct properties listed below.
}
```

<a id="canonical-c9da5f65b56fe8e33344241cc7a622031bff306f6a7a774f10865e984431f355"></a>

## Direct properties — ddos_mitigation_rules.ddos_client_source / c675fd2d9405 / 3

- [asn_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-dc383e23334e13441b25ce49785f82a5937bfef8e36eb239bb22641b565b99e3): complete subsection reference.

<a id="canonical-11b32b1522d51392b75363b1f9d74bb04c83d144af8144220cd8943610fb3c0c"></a>

<a id="canonical-e5f13d5becd04cef5260a520d490779bbbe644d092628fffafac8f014329197c"></a>

## country_list property — ddos_mitigation_rules.ddos_client_source / c675fd2d9405 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Sources that are located in one of the countries in the given list. Possible values are
\`COUNTRY\_NONE\`, \`COUNTRY\_AD\`, \`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`,
\`COUNTRY\_AI\`, \`COUNTRY\_AL\`, \`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`,
\`COUNTRY\_AQ\`, \`COUNTRY\_AR\`, \`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`,
\`COUNTRY\_AW\`, \`COUNTRY\_AX\`, \`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`,
\`COUNTRY\_BD\`, \`COUNTRY\_BE\`, \`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`,
\`COUNTRY\_BI\`, \`COUNTRY\_BJ\`, \`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`,
\`COUNTRY\_BO\`, \`COUNTRY\_BQ\`, \`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`,
\`COUNTRY\_BV\`, \`COUNTRY\_BW\`, \`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`,
\`COUNTRY\_CC\`, \`COUNTRY\_CD\`, \`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`,
\`COUNTRY\_CI\`, \`COUNTRY\_CK\`, \`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`,
\`COUNTRY\_CO\`, \`COUNTRY\_CR\`, \`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`,
\`COUNTRY\_CW\`, \`COUNTRY\_CX\`, \`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`,
\`COUNTRY\_DJ\`, \`COUNTRY\_DK\`, \`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`,
\`COUNTRY\_EC\`, \`COUNTRY\_EE\`, \`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`,
\`COUNTRY\_ES\`, \`COUNTRY\_ET\`, \`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`,
\`COUNTRY\_FM\`, \`COUNTRY\_FO\`, \`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`,
\`COUNTRY\_GD\`, \`COUNTRY\_GE\`, \`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`,
\`COUNTRY\_GI\`, \`COUNTRY\_GL\`, \`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`,
\`COUNTRY\_GQ\`, \`COUNTRY\_GR\`, \`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`,
\`COUNTRY\_GW\`, \`COUNTRY\_GY\`, \`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`,
\`COUNTRY\_HR\`, \`COUNTRY\_HT\`, \`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`,
\`COUNTRY\_IL\`, \`COUNTRY\_IM\`, \`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`,
\`COUNTRY\_IR\`, \`COUNTRY\_IS\`, \`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`,
\`COUNTRY\_JO\`, \`COUNTRY\_JP\`, \`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`,
\`COUNTRY\_KI\`, \`COUNTRY\_KM\`, \`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`,
\`COUNTRY\_KW\`, \`COUNTRY\_KY\`, \`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`,
\`COUNTRY\_LC\`, \`COUNTRY\_LI\`, \`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`,
\`COUNTRY\_LT\`, \`COUNTRY\_LU\`, \`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`,
\`COUNTRY\_MC\`, \`COUNTRY\_MD\`, \`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`,
\`COUNTRY\_MH\`, \`COUNTRY\_MK\`, \`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`,
\`COUNTRY\_MO\`, \`COUNTRY\_MP\`, \`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`,
\`COUNTRY\_MT\`, \`COUNTRY\_MU\`, \`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`,
\`COUNTRY\_MY\`, \`COUNTRY\_MZ\`, \`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`,
\`COUNTRY\_NF\`, \`COUNTRY\_NG\`, \`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`,
\`COUNTRY\_NP\`, \`COUNTRY\_NR\`, \`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`,
\`COUNTRY\_PA\`, \`COUNTRY\_PE\`, \`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`,
\`COUNTRY\_PK\`, \`COUNTRY\_PL\`, \`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`,
\`COUNTRY\_PS\`, \`COUNTRY\_PT\`, \`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`,
\`COUNTRY\_RE\`, \`COUNTRY\_RO\`, \`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`,
\`COUNTRY\_SA\`, \`COUNTRY\_SB\`, \`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`,
\`COUNTRY\_SG\`, \`COUNTRY\_SH\`, \`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`,
\`COUNTRY\_SL\`, \`COUNTRY\_SM\`, \`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`,
\`COUNTRY\_SS\`, \`COUNTRY\_ST\`, \`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`,
\`COUNTRY\_SZ\`, \`COUNTRY\_TC\`, \`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`,
\`COUNTRY\_TH\`, \`COUNTRY\_TJ\`, \`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`,
\`COUNTRY\_TN\`, \`COUNTRY\_TO\`, \`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`,
\`COUNTRY\_TW\`, \`COUNTRY\_TZ\`, \`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`,
\`COUNTRY\_US\`, \`COUNTRY\_UY\`, \`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`,
\`COUNTRY\_VE\`, \`COUNTRY\_VG\`, \`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`,
\`COUNTRY\_WF\`, \`COUNTRY\_WS\`, \`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`,
\`COUNTRY\_YT\`, \`COUNTRY\_ZA\`, \`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Upstream description:

Sources that are located in one of the countries in the given list.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [ja4_tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-009.md#canonical-69803ad85b7bcaca1b370bc0ce6f10cbae2d4feb08cdd49f1e120725790117cb): complete subsection reference.

- [tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-009.md#canonical-2e76b32a822c178d178b8cb9acff4844a34b654978c3b52f55a9fb04078b75a0): complete subsection reference.

<a id="canonical-5d7f3c3115271198088c5cde8fbc898a5e678fc304bddc3ecff2657b71203747"></a>

## Next pages — ddos_mitigation_rules.ddos_client_source / c675fd2d9405 / 5

- [ddos_mitigation_rules.ddos_client_source.asn_list](resources--cdn_loadbalancer--reference--group-009.md#canonical-dc383e23334e13441b25ce49785f82a5937bfef8e36eb239bb22641b565b99e3)
- [ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-009.md#canonical-69803ad85b7bcaca1b370bc0ce6f10cbae2d4feb08cdd49f1e120725790117cb)
- [ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-009.md#canonical-2e76b32a822c178d178b8cb9acff4844a34b654978c3b52f55a9fb04078b75a0)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-5746a5502dfc8e0895cb2a3ef2a31f60a71459706a1427f5bec63040fb069f07)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-dc383e23334e13441b25ce49785f82a5937bfef8e36eb239bb22641b565b99e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-882fdbe3bd3443288c45e8154f914669a5f2938a97d9ffadeb27ab92d7d02996"></a>

## ddos_mitigation_rules.ddos_client_source.asn_list — ddos_mitigation_rules.ddos_client_source.asn_list / dc14d516a713 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-5746a5502dfc8e0895cb2a3ef2a31f60a71459706a1427f5bec63040fb069f07)
- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-009.md#canonical-84a504b5b14092795e6fa5369922fa6b16e4fd9a7fa887303f90b8fb04294ab8)
- ddos_mitigation_rules.ddos_client_source.asn_list

<a id="canonical-9d04ccaccf158f5bb2160cb708ea92c807e77798ef441de04ac0a4308bcef591"></a>

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

<a id="canonical-8d14e58e45a052581962d99b0cea7a6359c79f0f41fe0f069614f1e93c0ba12f"></a>

## Direct properties — ddos_mitigation_rules.ddos_client_source.asn_list / dc14d516a713 / 3

<a id="canonical-41d962dbb94981cf87eba1b865a994cf9a6986535774672d890c0c828bb5cb25"></a>

<a id="canonical-b2e4f47ed78bc0c4bfe54d81ff430f1f438ec9e06e5b34fa99c6d351b21af988"></a>

## as_numbers property — ddos_mitigation_rules.ddos_client_source.asn_list / dc14d516a713 / 4

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

<a id="canonical-382117eba7a339d305b36f25e3bd6877fb295a12d30b29178b6a0ad44f30664d"></a>

## Next pages — ddos_mitigation_rules.ddos_client_source.asn_list / dc14d516a713 / 5

- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-009.md#canonical-84a504b5b14092795e6fa5369922fa6b16e4fd9a7fa887303f90b8fb04294ab8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-69803ad85b7bcaca1b370bc0ce6f10cbae2d4feb08cdd49f1e120725790117cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de7b077b1258dee1d645c84db0217a3330efbc8206817e49ea1e78dbc33ce2bb"></a>

## ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher — ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher / 7e60af850ac7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-5746a5502dfc8e0895cb2a3ef2a31f60a71459706a1427f5bec63040fb069f07)
- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-009.md#canonical-84a504b5b14092795e6fa5369922fa6b16e4fd9a7fa887303f90b8fb04294ab8)
- ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher

<a id="canonical-bc5a19d2bed45a9da8ab32a1c3fb6bf58888a78e2bf18baee106319f6005c71d"></a>

Type: `"object"`. single nested block, Optional.

Extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

Upstream description:

An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

Receipt-pinned upstream constraints:

```json
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
ja4_tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-35eb7a49e33b7e69666d2eeb9cffd556f192e7337a4ed5b753076453f9e42e2e"></a>

## Direct properties — ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher / 7e60af850ac7 / 3

<a id="canonical-8c9fce7f6819248722dbeef249eac40c901251da2669bf3ffddd9f520be11d7c"></a>

<a id="canonical-db4956be1be7478bf51f1cb3b4f9db7b4274ea39fb7c4a37dbc74f47354a7c1a"></a>

## exact_values property — ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher / 7e60af850ac7 / 4

Type: `["list", "string"]`. Optional.

List of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Upstream description:

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a8d8074d526cede82295b1ab661f670241db72f719a16074b98378f03690a678"></a>

## Next pages — ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher / 7e60af850ac7 / 5

- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-009.md#canonical-84a504b5b14092795e6fa5369922fa6b16e4fd9a7fa887303f90b8fb04294ab8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2e76b32a822c178d178b8cb9acff4844a34b654978c3b52f55a9fb04078b75a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
