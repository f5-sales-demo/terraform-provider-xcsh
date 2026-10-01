---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-aa6024e93f49bd8b9c7d83cff4e646047c06420d783d96cfe7c20a5e89e9aceb"></a>

## description_spec property — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadat / edd29c073f02 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-ce0c627f72bcc7f61c1f6b76161dfdc59279c290da3bf8fa1e2783f0aba6f0c8"></a>

<a id="canonical-637501d69dce00dc7eaf6f850b249e40475066646cced6f462cd8f2ca46af9d5"></a>

## name property — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadat / edd29c073f02 / 5

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

<a id="canonical-1e74f782ef6fd459db67cfe6e72e3525727f9481c211b0466c5f15219cbe67d3"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadat / edd29c073f02 / 6

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-f6d0f9297117a3167801c698e2a43d0d0d89caaac1f4bf8114b1de0c14cf1738)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-568742d77262697064e45db08ec16b1643bca8ccbd9b25a7b77195fea55885bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-904a9eab0d4dbf07f4ea633a0434b72cfeefe4deaf9fc7cdcc1dcb7705a33f2f"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path / ebcc23435bf6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-012.md#canonical-c97868542280c86c9322eccdb6092ef736f7e9c0f70d01a8ccbbbfeabcdc263d)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-013.md#canonical-ae66eebe64a1b355feb3b51809c7a41951b2c542a106e845900ae24627bd4626)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-0a985c497432bb4b357496cf5b2cd6e86dde9ccfedf5273bbd8ce88b551dee71)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-f6d0f9297117a3167801c698e2a43d0d0d89caaac1f4bf8114b1de0c14cf1738)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path

<a id="canonical-dc5bd699a45afa810427a8c28ec0c2c3c362065c4a7937d8530ed6d1e6f39612"></a>

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

<a id="canonical-47bc7a6e3877eeb1013a27894c066677bd6b5df82c2b9b11c6044f356a3f2754"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path / ebcc23435bf6 / 3

<a id="canonical-22ebe7018506a04ac0265c6480bf0c9ba212a5b7207ccec025ec3b3c36a373a4"></a>

<a id="canonical-9383c3efc83e7c898dac88dc418c01b31730cdc980d7eb720d83969c0a782ef3"></a>

## path property — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path / ebcc23435bf6 / 4

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

<a id="canonical-61eb896547130361a4b22f87f8265aca914f16cf8e739cdaf7a4ee2aaf441d50"></a>

<a id="canonical-17e70d13c55fce6e36efbf9d2c21df33160cd5bb818c81773b1d1e6b20fc0faf"></a>

## prefix property — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path / ebcc23435bf6 / 5

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

<a id="canonical-6d3a3b33a11dacd3b29ff938ea0d927a0b53ce1fbea3d1cc12a928af5e6b2817"></a>

<a id="canonical-b33827558d258a934c66d1f3f8e3219b0dbc9667ad20a67b2d54e274c8e431d3"></a>

## regex property — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path / ebcc23435bf6 / 6

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

<a id="canonical-08bf27ffde4577097a4723c274bdf6447ee95bddd75363ce48c2c357b41083f3"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path / ebcc23435bf6 / 7

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-013.md#canonical-f6d0f9297117a3167801c698e2a43d0d0d89caaac1f4bf8114b1de0c14cf1738)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ec78ce49852d8bde21f7f6098649b3b277233fa08cd92b5ff2443cda9ada87b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a0cdc7b54d64b04beef41bd4926135c4d36de21ff384313c88401202482217a"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules — bot_defense_advanced_protection.web_only.js_insertion_rules.rules / b69aef7267a3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-012.md#canonical-c97868542280c86c9322eccdb6092ef736f7e9c0f70d01a8ccbbbfeabcdc263d)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-013.md#canonical-ae66eebe64a1b355feb3b51809c7a41951b2c542a106e845900ae24627bd4626)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-0a985c497432bb4b357496cf5b2cd6e86dde9ccfedf5273bbd8ce88b551dee71)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules

<a id="canonical-3006a41581a424eaf341f93c9dd9144eb1746147bdef03d9403c14e8e5c43a9e"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

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

<a id="canonical-f67bcd33aba525630dffe420bf0443cbe8f0bf0e841e0baa457025cbd63f78dc"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.rules / b69aef7267a3 / 3

- [any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-60f4e5b21f0b6269074de126fdf2ecca3f6b275f8b85abfa5113c341397c1507): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-014.md#canonical-71bd286f2f1b54e24b902205af52dfdb7056acc74d3cc325281a31c8d48449d0): complete subsection reference.

<a id="canonical-0e489a706fd6da3465ef99e909d109c898f135ca215757eace6df1a3febf70e6"></a>

<a id="canonical-24f5d5f5ec0737af28374fb8a75616d5f7afb1e35ed5f9510043817f9f0f24eb"></a>

## javascript_location property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules / b69aef7267a3 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

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

- [metadata](resources--http_loadbalancer--reference--group-014.md#canonical-42c55474af9d3a3cecf11e449ea9432f9adbb3405d4c13ec38e52da8d5dae941): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-014.md#canonical-f3ffd54db49c1d186a4b78d70359bd519f4488752ac361872cdb54bc49adcf7f): complete subsection reference.

<a id="canonical-a2582fa96ba2f0f7ff5fc31e6da613e685e8dcfc71cbcad5f31ec0bd17a69454"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.rules / b69aef7267a3 / 5

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-60f4e5b21f0b6269074de126fdf2ecca3f6b275f8b85abfa5113c341397c1507)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain](resources--http_loadbalancer--reference--group-014.md#canonical-71bd286f2f1b54e24b902205af52dfdb7056acc74d3cc325281a31c8d48449d0)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata](resources--http_loadbalancer--reference--group-014.md#canonical-42c55474af9d3a3cecf11e449ea9432f9adbb3405d4c13ec38e52da8d5dae941)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path](resources--http_loadbalancer--reference--group-014.md#canonical-f3ffd54db49c1d186a4b78d70359bd519f4488752ac361872cdb54bc49adcf7f)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-0a985c497432bb4b357496cf5b2cd6e86dde9ccfedf5273bbd8ce88b551dee71)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-60f4e5b21f0b6269074de126fdf2ecca3f6b275f8b85abfa5113c341397c1507"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f1df6c0f1d9c435cc50976192bda5cde93dd721e7ff1e0c63c9ac422b1902d9"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain / d59aebeb6175 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-012.md#canonical-c97868542280c86c9322eccdb6092ef736f7e9c0f70d01a8ccbbbfeabcdc263d)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-013.md#canonical-ae66eebe64a1b355feb3b51809c7a41951b2c542a106e845900ae24627bd4626)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-0a985c497432bb4b357496cf5b2cd6e86dde9ccfedf5273bbd8ce88b551dee71)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-ec78ce49852d8bde21f7f6098649b3b277233fa08cd92b5ff2443cda9ada87b3)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain

<a id="canonical-91f30b8bb23cfcb7ccc263b82e6462bc84a80e90d2528c03bed37d73944061ca"></a>

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

<a id="canonical-f2b79d1b4d9de9c816bba781d028196d41ecdaf8be295551966d2a71f1ae944b"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain / d59aebeb6175 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-444ffe23cd7de3f7cdbe4bd580b907e951fc4d71fe004ee9bf4b57bcb5d12a6b"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain / d59aebeb6175 / 4

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-ec78ce49852d8bde21f7f6098649b3b277233fa08cd92b5ff2443cda9ada87b3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-71bd286f2f1b54e24b902205af52dfdb7056acc74d3cc325281a31c8d48449d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a45ccd9a636be04dc9d7c0ec3bcb7108e01bb404f68f549ef246143cb037ba5f"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain / eeabc7e354c1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-012.md#canonical-c97868542280c86c9322eccdb6092ef736f7e9c0f70d01a8ccbbbfeabcdc263d)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-013.md#canonical-ae66eebe64a1b355feb3b51809c7a41951b2c542a106e845900ae24627bd4626)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-0a985c497432bb4b357496cf5b2cd6e86dde9ccfedf5273bbd8ce88b551dee71)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-ec78ce49852d8bde21f7f6098649b3b277233fa08cd92b5ff2443cda9ada87b3)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain

<a id="canonical-d0a9ba83dddc42a5b95aa8533f119b40e7de3265df5d7d2255ad1f452dbb4c47"></a>

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

<a id="canonical-7677a815f4da475f3ffae643e55c23aa1966db520130826e5452d92a5565e99d"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain / eeabc7e354c1 / 3

<a id="canonical-b12097c075f43829496c36773f3c05c595f8516118509b1de58fd9d12d8a7c1e"></a>

<a id="canonical-de2069ca9d96236e5657ded35653dc6f719a375be1fd13a3d3463f9c06a485f4"></a>

## exact_value property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain / eeabc7e354c1 / 4

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

<a id="canonical-d8efedb30b7d9c6048d994722d229ed53ff2a050e82fc205e6ebf45f2720b390"></a>

<a id="canonical-0231a5a33b28b6dc06675794ab3d5f274c6eb992c84fe3a8ad507246284d1d38"></a>

## regex_value property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain / eeabc7e354c1 / 5

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

<a id="canonical-160b95548fc10f7ad103e1e95d0aa5808dd5cfbd3911ff4ec1a281f8ce56e27f"></a>

<a id="canonical-17c091e10205450808421df811ad2432cc6ffc963f66e327ba1440443eeabf55"></a>

## suffix_value property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain / eeabc7e354c1 / 6

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

<a id="canonical-83cc74cbc72154d446ab435773f295a0c9fdf56e782db6d1286e4034fef7505e"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain / eeabc7e354c1 / 7

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-ec78ce49852d8bde21f7f6098649b3b277233fa08cd92b5ff2443cda9ada87b3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-42c55474af9d3a3cecf11e449ea9432f9adbb3405d4c13ec38e52da8d5dae941"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41392805a3c03fad441db850ff2f3018ca42d5f16dffe6c7873087f53ac623a3"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata / 6d76a4feac6e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-012.md#canonical-c97868542280c86c9322eccdb6092ef736f7e9c0f70d01a8ccbbbfeabcdc263d)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-013.md#canonical-ae66eebe64a1b355feb3b51809c7a41951b2c542a106e845900ae24627bd4626)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-0a985c497432bb4b357496cf5b2cd6e86dde9ccfedf5273bbd8ce88b551dee71)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-ec78ce49852d8bde21f7f6098649b3b277233fa08cd92b5ff2443cda9ada87b3)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata

<a id="canonical-c03230a612e805cea5597452d02c16966e670e084e4073be05a944c62a877afe"></a>

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

<a id="canonical-7fc886ea455e0694ff574e4e62c7aeac0af87a8e97c42ad8adfd7e0c5d280379"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata / 6d76a4feac6e / 3

<a id="canonical-9e45a451d2269b03f3a367883cc764932cc8b08a929d58ed5883b71cd130a6e1"></a>

<a id="canonical-2ebf09b4c567b448ec60fd14626e187121efb9970bda6b335acfddee9faa408d"></a>

## description_spec property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata / 6d76a4feac6e / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-ccc515b916691d0ace1256370ac2127a6c10b1cfe873e0ea9a9b8e6c0314ec2a"></a>

<a id="canonical-c056b9de8444abcaa9c8ee26a7a3b69fa441ee15fa16767adb50397aaa66c8e4"></a>

## name property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata / 6d76a4feac6e / 5

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

<a id="canonical-51df2ad578b7618bc1192457ee5deed0e8197c245c061ff636df974ac3649856"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata / 6d76a4feac6e / 6

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-ec78ce49852d8bde21f7f6098649b3b277233fa08cd92b5ff2443cda9ada87b3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f3ffd54db49c1d186a4b78d70359bd519f4488752ac361872cdb54bc49adcf7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d400e7c181640f4b4dbe02e261e52329fcf59d384f966967abf610a5c585db28"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path / fcc7f0096353 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-012.md#canonical-c97868542280c86c9322eccdb6092ef736f7e9c0f70d01a8ccbbbfeabcdc263d)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-013.md#canonical-ae66eebe64a1b355feb3b51809c7a41951b2c542a106e845900ae24627bd4626)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-013.md#canonical-0a985c497432bb4b357496cf5b2cd6e86dde9ccfedf5273bbd8ce88b551dee71)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-ec78ce49852d8bde21f7f6098649b3b277233fa08cd92b5ff2443cda9ada87b3)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path

<a id="canonical-238f794df4f127453c63118a180bc3bc41413c4bf0c89abd8eef7d48085c4321"></a>

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

<a id="canonical-12fc9597343d7fe4ffd094a74b1e0c1aadf36dc708dc2739e40b2dd9b7d88289"></a>

## Direct properties — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path / fcc7f0096353 / 3

<a id="canonical-84b7e638ed09259f460cf6a03c49ba9d581045369d661a12a77dbef197a9c011"></a>

<a id="canonical-b186aad78f845b49cb9b39bd2f4963ef416f78ba9717eb795729c6d8165e87a2"></a>

## path property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path / fcc7f0096353 / 4

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

<a id="canonical-f2ff69f6d5a8d762fe52eb8cf22b3ab6c5d5f93231ae5cd8d69a1a4145205519"></a>

<a id="canonical-7d53fb50a9cd4f4b89358c9e3f756b1c666871473e0b2918847efc45dacf5dab"></a>

## prefix property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path / fcc7f0096353 / 5

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

<a id="canonical-391cce3dadda0093164ad389bc39ee9788fb837557d5f1843bada46386a11eda"></a>

<a id="canonical-f537ccfbecd3d2bec0619a381472093415d9d4ee7df57c57320ba03db5c81eb0"></a>

## regex property — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path / fcc7f0096353 / 6

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

<a id="canonical-f2a0b1bae1a457c996b1e073b5695625c7a3ba598e8e23f7977554cd956d7b92"></a>

## Next pages — bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path / fcc7f0096353 / 7

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-ec78ce49852d8bde21f7f6098649b3b277233fa08cd92b5ff2443cda9ada87b3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0efb6811c6ea8316785224b62886e488a271ca47310daf8c3d667cdd1fff2b24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2767e720f10fd284e561d07f25a47b00e1cb510e449288c145721aee35f9e1c2"></a>

## bot_defense_advanced_protection.web_only.web — bot_defense_advanced_protection.web_only.web / 0d5f51bea89a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-012.md#canonical-c97868542280c86c9322eccdb6092ef736f7e9c0f70d01a8ccbbbfeabcdc263d)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-013.md#canonical-ae66eebe64a1b355feb3b51809c7a41951b2c542a106e845900ae24627bd4626)
- bot_defense_advanced_protection.web_only.web

<a id="canonical-e2418ed587fe85addfb907fdb8c1a778981565b8c0bc647b4ac5345029a29447"></a>

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
web {
  # Configure direct properties listed below.
}
```

<a id="canonical-a15626fa1d7cefc60d43d18bcf427153f450c7a844c3befbe1dca7f3b76bccea"></a>

## Direct properties — bot_defense_advanced_protection.web_only.web / 0d5f51bea89a / 3

<a id="canonical-989d9571439ff60ccbbd83020129830aaa3ebe009d4680cc7da3c196b9434ad7"></a>

<a id="canonical-fae8b8599816feea1b2a0dbf5e3994d13ab05d2662ef1b78e0883109a2537f11"></a>

## name property — bot_defense_advanced_protection.web_only.web / 0d5f51bea89a / 4

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

<a id="canonical-1ef11aea2555ae94bb4e34e2947f5f11b2791334e6ce4a2d427f745afb1d2a2a"></a>

<a id="canonical-a3efbc41de5fc23d8b9dccfd77b613e6ca4a938960dbf124fff51388a4cf362f"></a>

## namespace property — bot_defense_advanced_protection.web_only.web / 0d5f51bea89a / 5

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

<a id="canonical-b7a7ab15cdba519beb65c9f841bf8adf68625053eaf92765f4893e791624dec8"></a>

<a id="canonical-85b33ef744933f769c99f454c2837100a8f17eea62ad7b4162f78d47bffe77af"></a>

## tenant property — bot_defense_advanced_protection.web_only.web / 0d5f51bea89a / 6

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

<a id="canonical-c128c90789e37d1f8e0ae41bf15bc95ef73b1776e51d60ca9b04a57d012454c7"></a>

## Next pages — bot_defense_advanced_protection.web_only.web / 0d5f51bea89a / 7

- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-013.md#canonical-ae66eebe64a1b355feb3b51809c7a41951b2c542a106e845900ae24627bd4626)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d20150bcbc1a30a287cae322ba14c8cf07c2866aab4bca3a45ecca53fbc870fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd711c21a6d3c66a6cdd954be2e01466dac1773bd3fe871e079854839679f6ae"></a>

## caching_policy — caching_policy / b1cf0937598f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- caching_policy

<a id="canonical-e3876f087b9afd88e733cdb2e17e2b8146861bcb01900d94bf4fcc3ed71f4ab8"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: caching\_policy, disable\_caching; Default: disable\_caching\] Policy configuration for
this feature.

Upstream description:

Caching Policies for the CDN.

Receipt-pinned upstream constraints:

```json
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

- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-e3876f087b9afd88e733cdb2e17e2b8146861bcb01900d94bf4fcc3ed71f4ab8)
- [disable_caching](resources--http_loadbalancer--reference--group-017.md#canonical-b413cccae648cb2bf5e5fe30f01c130252244e7839d08a39058ea4f65a9b067a)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
caching_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-1b1a8d8897dc40aad3c3b16aedb12087c81ed7635cfedb9b5e067cc63a963c84"></a>

## Direct properties — caching_policy / b1cf0937598f / 3

- [custom_cache_rule](resources--http_loadbalancer--reference--group-014.md#canonical-56eecf01e47f1c9c414de61257560fe28495f51ee3de27c94c00f3c8664af991): complete subsection reference.

- [default_cache_action](resources--http_loadbalancer--reference--group-014.md#canonical-92cdb117ac78b5d4317af9401b57dfc21f6b382f09c5b064a63bc8757739de59): complete subsection reference.

<a id="canonical-5b937bdc6bef41368370a390f97ec5cca243c19859df9a342885e66eada2cdd4"></a>

## Next pages — caching_policy / b1cf0937598f / 4

- [caching_policy.custom_cache_rule](resources--http_loadbalancer--reference--group-014.md#canonical-56eecf01e47f1c9c414de61257560fe28495f51ee3de27c94c00f3c8664af991)
- [caching_policy.default_cache_action](resources--http_loadbalancer--reference--group-014.md#canonical-92cdb117ac78b5d4317af9401b57dfc21f6b382f09c5b064a63bc8757739de59)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-56eecf01e47f1c9c414de61257560fe28495f51ee3de27c94c00f3c8664af991"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74d21c3ba1622e49e66cbcd711a55c7e19b930473efe9064eaedb26f2cc512b2"></a>

## caching_policy.custom_cache_rule — caching_policy.custom_cache_rule / f77a059220bb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-d20150bcbc1a30a287cae322ba14c8cf07c2866aab4bca3a45ecca53fbc870fd)
- caching_policy.custom_cache_rule

<a id="canonical-4cb818ed58eb74ba743e9e48359e9354cdea6bd72000f9523fe8a49e888f36b5"></a>

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

<a id="canonical-c660e7309fe80edf7ca531e4231e2bc2178b35a795f7172ac087545af2d4790a"></a>

## Direct properties — caching_policy.custom_cache_rule / f77a059220bb / 3

- [cdn_cache_rules](resources--http_loadbalancer--reference--group-014.md#canonical-38a16683e020f6be71df20125a7c0dd2482b70ed7a4a48f1c4b1f9eec37ac28e): complete subsection reference.

<a id="canonical-c581f8db38c22f2eed16ce7e898ef69375396ef5a30c7a63dc3b6087746e7ca5"></a>

## Next pages — caching_policy.custom_cache_rule / f77a059220bb / 4

- [caching_policy.custom_cache_rule.cdn_cache_rules](resources--http_loadbalancer--reference--group-014.md#canonical-38a16683e020f6be71df20125a7c0dd2482b70ed7a4a48f1c4b1f9eec37ac28e)
- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-d20150bcbc1a30a287cae322ba14c8cf07c2866aab4bca3a45ecca53fbc870fd)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-38a16683e020f6be71df20125a7c0dd2482b70ed7a4a48f1c4b1f9eec37ac28e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16d13c008a0413fde532a9586418a2c41a339cd09f51a9d112ec53099e547c05"></a>

## caching_policy.custom_cache_rule.cdn_cache_rules — caching_policy.custom_cache_rule.cdn_cache_rules / 3c5c8248972a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-d20150bcbc1a30a287cae322ba14c8cf07c2866aab4bca3a45ecca53fbc870fd)
- [caching_policy.custom_cache_rule](resources--http_loadbalancer--reference--group-014.md#canonical-56eecf01e47f1c9c414de61257560fe28495f51ee3de27c94c00f3c8664af991)
- caching_policy.custom_cache_rule.cdn_cache_rules

<a id="canonical-7a4595ca0a5a63cfacd65dde3b76765339b7bde4b43fe024ea1f2066da5cd7a6"></a>

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

<a id="canonical-2e8cf0ab1ca76af217ede50c8c8f9ac85f529839b9c5a917b910b347e9156f47"></a>

## Direct properties — caching_policy.custom_cache_rule.cdn_cache_rules / 3c5c8248972a / 3

<a id="canonical-e96047caeeace89fbd31ab633cf627a77a22b8dcd6d0211ad42ae02abd76bc9b"></a>

<a id="canonical-d8c482137a5728a40fe6fef9b1fa9ad1cad4fdbb1706979503c09b274376db1d"></a>

## name property — caching_policy.custom_cache_rule.cdn_cache_rules / 3c5c8248972a / 4

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

<a id="canonical-de877960cef6815855d0d58942051e624c6becbb66f402d315d6a1b3aad48afb"></a>

<a id="canonical-5b902604b080c46fc86f398e87e3cba0d683bb8526cd7d2aa9d15c2f506f217f"></a>

## namespace property — caching_policy.custom_cache_rule.cdn_cache_rules / 3c5c8248972a / 5

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

<a id="canonical-02aec683068f0a7a97f883d80d5eeacd1b506b4f19a58df43f8c6ebed1abef05"></a>

<a id="canonical-fa33119b2e618b2128eb634364c8b0d3c19cdaad6bccc1499aa0293d59500c07"></a>

## tenant property — caching_policy.custom_cache_rule.cdn_cache_rules / 3c5c8248972a / 6

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

<a id="canonical-f250d2d5d2ef08e979dd023939722d4aea4e279e97620e7c10b99b48a2621e8a"></a>

## Next pages — caching_policy.custom_cache_rule.cdn_cache_rules / 3c5c8248972a / 7

- [caching_policy.custom_cache_rule](resources--http_loadbalancer--reference--group-014.md#canonical-56eecf01e47f1c9c414de61257560fe28495f51ee3de27c94c00f3c8664af991)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-92cdb117ac78b5d4317af9401b57dfc21f6b382f09c5b064a63bc8757739de59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bab251356fc9440c8e4760293e3ab0af11e2a25dc21d2bba298bd11e0b66e9c"></a>

## caching_policy.default_cache_action — caching_policy.default_cache_action / 5961ae7a4ca1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-d20150bcbc1a30a287cae322ba14c8cf07c2866aab4bca3a45ecca53fbc870fd)
- caching_policy.default_cache_action

<a id="canonical-2d45d6c91581700f73aa1a9e3492fd23a69f062ec22473d272d3ad2927cca8ad"></a>

Type: `"object"`. single nested block, Optional.

Default Cache Behaviour. This defines a Default Cache Action.

Upstream description:

This defines a Default Cache Action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cache_disabled",
    "cache_ttl_default"),
  validators.ConflictingObjectAttributes("cache_disabled",
    "cache_ttl_override"),
  validators.ConflictingObjectAttributes("cache_ttl_default",
    "cache_ttl_override")}
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
  "x-ves-oneof-field-cache_actions": "[\"cache_disabled\",\"cache_ttl_default\",\"cache_ttl_override\"]"
}
```

Terraform syntax:

```terraform
default_cache_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-611bbf7726a0831c0e37b2d3b38acb9cba6d2bb8248e2f5aa491509cd63f3db1"></a>

## Direct properties — caching_policy.default_cache_action / 5961ae7a4ca1 / 3

- [cache_disabled](resources--http_loadbalancer--reference--group-014.md#canonical-e45ba19aa2ae1bec2d9a36007fca0d0acc0b7b4f6cd211b1a91a2439c60e0e11): complete subsection reference.

<a id="canonical-a3343b50cafddf89a92d5cc45593535d8280d4e790d2cc440d1437ce48a8a0d0"></a>

<a id="canonical-677ff852042304601942a1e07d8b8a4d67d7e3c9f5f833bac906c56faa8785bf"></a>

## cache_ttl_default property — caching_policy.default_cache_action / 5961ae7a4ca1 / 4

Type: `"string"`. Optional.

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-53f732ac8cd48120b60d22abfd3c11615e091ea07dd3b793800475563c180acd"></a>

<a id="canonical-ad6d10dadbb52c83f15acaa2fa63d0a3383de6fe8f7116a4af8cbbbb740ff86c"></a>

## cache_ttl_override property — caching_policy.default_cache_action / 5961ae7a4ca1 / 5

Type: `"string"`. Optional.

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-c568d82fe563e39432b7f1954cd1e0a337d50e1f039040df70c293b243e386ec"></a>

## Next pages — caching_policy.default_cache_action / 5961ae7a4ca1 / 6

- [caching_policy.default_cache_action.cache_disabled](resources--http_loadbalancer--reference--group-014.md#canonical-e45ba19aa2ae1bec2d9a36007fca0d0acc0b7b4f6cd211b1a91a2439c60e0e11)
- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-d20150bcbc1a30a287cae322ba14c8cf07c2866aab4bca3a45ecca53fbc870fd)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e45ba19aa2ae1bec2d9a36007fca0d0acc0b7b4f6cd211b1a91a2439c60e0e11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d60fe30dcc436d9d1321e40b6cb45ec6aec513049d00835d3f1f58ca09c5405c"></a>

## caching_policy.default_cache_action.cache_disabled — caching_policy.default_cache_action.cache_disabled / bd4ae12b8916 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-d20150bcbc1a30a287cae322ba14c8cf07c2866aab4bca3a45ecca53fbc870fd)
- [caching_policy.default_cache_action](resources--http_loadbalancer--reference--group-014.md#canonical-92cdb117ac78b5d4317af9401b57dfc21f6b382f09c5b064a63bc8757739de59)
- caching_policy.default_cache_action.cache_disabled

<a id="canonical-b9fa8f70091c4e869a2e50843793f5796e7bb2a102712c5d66aec29ff8f97f5e"></a>

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
cache_disabled = {}
```

<a id="canonical-84bcedc61491c5203ace5225a134aedfda94a3cdad2d0c7c520041a8867345b7"></a>

## Direct properties — caching_policy.default_cache_action.cache_disabled / bd4ae12b8916 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3bab76eff83152e3ef703a631d2c85c00d34ccba6d4b52d45d0569b1eeca4947"></a>

## Next pages — caching_policy.default_cache_action.cache_disabled / bd4ae12b8916 / 4

- [caching_policy.default_cache_action](resources--http_loadbalancer--reference--group-014.md#canonical-92cdb117ac78b5d4317af9401b57dfc21f6b382f09c5b064a63bc8757739de59)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-fd05a1137cab68c6922bcd30d9888aeee26a67af28a835d3b19f4d7b6e67bbb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5fcda3b9523414ee1cee29541d79be4e90a4bd960fccb8711d23eb6d70d61092"></a>

## captcha_challenge — captcha_challenge / 6501364caea9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- captcha_challenge

<a id="canonical-6e3be2d2440406a2160ed2acef37f1aafdfe8bc0f0f70bcd673d14f518018ed7"></a>

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

- [captcha_challenge](resources--http_loadbalancer--reference--group-014.md#canonical-6e3be2d2440406a2160ed2acef37f1aafdfe8bc0f0f70bcd673d14f518018ed7)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-3dfa8587960b4f35004d16184a7dd796a97d011b48fea0d273cfa032b9b6832d)
- [js_challenge](resources--http_loadbalancer--reference--group-019.md#canonical-4557dc29e7f082374f2bec8955f6733d0e0c82acc02f9f5f823afb5f2f14e7d0)
- [no_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-de253bc4120f307cd44e3e5625cf404100e08678cf40944fa82eb5b7f4a67ccb)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-39ad349c197b736b689021789fb7cd999d794a0277cbe70b543cc70730b18cae)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
captcha_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-7548a70c3e1c5b134484720ff72d9147b88bb6e0b6c25c5961922a8052c8f79a"></a>

## Direct properties — captcha_challenge / 6501364caea9 / 3

<a id="canonical-343fd48931d9672e173e7d047c56b281f085634a91bff707e6e694956e4a8c08"></a>

<a id="canonical-8dda96990e9325db5925b7a218a683a149e8b227e0af3e65127019e287cf28b0"></a>

## cookie_expiry property — captcha_challenge / 6501364caea9 / 4

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

<a id="canonical-eb482b3ec0359793b842e87219b03e7a3a9bf358477fd9620c028793cb925b62"></a>

<a id="canonical-033b33933cb4fcd4ae91b1db415b180b657b92a12a0c9a74c41aaf6aa1c6fcba"></a>

## custom_page property — captcha_challenge / 6501364caea9 / 5

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

<a id="canonical-0e329232baa40735de9c3eeeb908e58dea76770e88930ad20e17a3772c211c07"></a>

## Next pages — captcha_challenge / 6501364caea9 / 6

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e3b12d7d0b53c796c31dac3e3b2dbc5841fc24b5310dc3d178b2d38e5b57b99"></a>

## client_side_defense — client_side_defense / 85396d5ce367 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- client_side_defense

<a id="canonical-0ff888502d4981ae758b3a4993edf37f93d25ac75363c3e0fe07200c2f861a9e"></a>

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

- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0ff888502d4981ae758b3a4993edf37f93d25ac75363c3e0fe07200c2f861a9e)
- [disable_client_side_defense](resources--http_loadbalancer--reference--group-017.md#canonical-7701f8ceb4f61af7f586280ab35cd8654495e369b75bf9f3819499706530bef2)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
client_side_defense {
  # Configure direct properties listed below.
}
```

<a id="canonical-8f1798e13f404576278e72db30461ebbd88fa948a1344a75ee256e56e735451f"></a>

## Direct properties — client_side_defense / 85396d5ce367 / 3

- [policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74): complete subsection reference.

<a id="canonical-8390fa948f0bfd88ef44f30d8a237ff7b6a323c2a641e5a3444ec3800edb1588"></a>

## Next pages — client_side_defense / 85396d5ce367 / 4

- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e23af7382b796c0f70473d048e505476b8f966038d8fff6cd777b3bbcb6d9763"></a>

## client_side_defense.policy — client_side_defense.policy / 54cd73aa049b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- client_side_defense.policy

<a id="canonical-67a4395a5e155e4f875db70ed5e82b6593ccdaec216e06ce6d41ae894d35102c"></a>

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

<a id="canonical-2cde952fa25cc07653ee5622972943f012f335b706315a6a48d4fc3fc8ae45d2"></a>

## Direct properties — client_side_defense.policy / 54cd73aa049b / 3

- [disable_js_insert](resources--http_loadbalancer--reference--group-014.md#canonical-177515c1f963af5fb3aa4f3030ae41b5fa7c681e54015dd837c34b4d19893aae): complete subsection reference.

- [js_insert_all_pages](resources--http_loadbalancer--reference--group-014.md#canonical-f9647ef721ea1e6aa04952dfb67aa871abe22a0abe7dfe94d6dfc63785ed00e0): complete subsection reference.

- [js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-e35eed5363c4bd97535b7099b0c67740891a71d420b45d9f0fe0667cfaadd92d): complete subsection reference.

- [js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b): complete subsection reference.

<a id="canonical-48d05471d7002b422691415c1a6e2d2b6bbe31b1a3dda3a9157db0551bd5f8a1"></a>

## Next pages — client_side_defense.policy / 54cd73aa049b / 4

- [client_side_defense.policy.disable_js_insert](resources--http_loadbalancer--reference--group-014.md#canonical-177515c1f963af5fb3aa4f3030ae41b5fa7c681e54015dd837c34b4d19893aae)
- [client_side_defense.policy.js_insert_all_pages](resources--http_loadbalancer--reference--group-014.md#canonical-f9647ef721ea1e6aa04952dfb67aa871abe22a0abe7dfe94d6dfc63785ed00e0)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-e35eed5363c4bd97535b7099b0c67740891a71d420b45d9f0fe0667cfaadd92d)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-177515c1f963af5fb3aa4f3030ae41b5fa7c681e54015dd837c34b4d19893aae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f82eef79e20463d98122ec497e9caf3a8575510c3059cf13cbec3449d413693"></a>

## client_side_defense.policy.disable_js_insert — client_side_defense.policy.disable_js_insert / 459b054488e0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- client_side_defense.policy.disable_js_insert

<a id="canonical-4dd7de672389dd3b0ab0008294dc3d1e5b34bc48e2e09af060123eacdfaf97c1"></a>

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

<a id="canonical-f4bad4181486f2fa00f9d7fd223e1ce36eda729eaf39bf7587543904359abf81"></a>

## Direct properties — client_side_defense.policy.disable_js_insert / 459b054488e0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f5833af34ebef8a517721319d1d86e5d2db98c55a13b1c56f1c46663ed78ae6c"></a>

## Next pages — client_side_defense.policy.disable_js_insert / 459b054488e0 / 4

- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f9647ef721ea1e6aa04952dfb67aa871abe22a0abe7dfe94d6dfc63785ed00e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fdcaa5d66ff4cc958f171e7612a8d2c8565c94a04af8894a44f40616df33f25f"></a>

## client_side_defense.policy.js_insert_all_pages — client_side_defense.policy.js_insert_all_pages / 2651bfc50631 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- client_side_defense.policy.js_insert_all_pages

<a id="canonical-8ed77003b0eb9a4e59cffd264a221aec7c2bdca0dd0748b1f191d106c92b5548"></a>

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

<a id="canonical-15be47993e21dd77adfe25af5aa0520d7a148bbc3692459d1e80310fea1f6cc2"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages / 2651bfc50631 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c90e86f5b0757bff7edfcbb9d3dd50abd2fe48759bf6edaa22c2ef32d7ce359a"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages / 2651bfc50631 / 4

- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e35eed5363c4bd97535b7099b0c67740891a71d420b45d9f0fe0667cfaadd92d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59a92eae1d923b2b819e3ffac4db2d7957c99821edeebb2d2798a45dbbe42f5a"></a>

## client_side_defense.policy.js_insert_all_pages_except — client_side_defense.policy.js_insert_all_pages_except / da7b32a68f84 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- client_side_defense.policy.js_insert_all_pages_except

<a id="canonical-dfb26605ced05341f4f006c51e0147fc2d69d9b34593a8b9337ae42ee2be6c87"></a>

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

<a id="canonical-51183fd7a4be7fdb35ee5a9146826e5b1857d0c1464b8a4e88c2f849c8cd2a0b"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except / da7b32a68f84 / 3

- [exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0fb2b7902a7a9a19077e2ed4fc8cba546bdbc031a18df7f360f551bb3d4f4177): complete subsection reference.

<a id="canonical-b09479038323091e073b058672d7c9eb5c3f1908b6cc68decfac4791fd32ce70"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except / da7b32a68f84 / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0fb2b7902a7a9a19077e2ed4fc8cba546bdbc031a18df7f360f551bb3d4f4177)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0fb2b7902a7a9a19077e2ed4fc8cba546bdbc031a18df7f360f551bb3d4f4177"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d40958a61778b5efeda337d5ed8e0bdfdf53d639dc699a7f4ff14d00d6685c5"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list — client_side_defense.policy.js_insert_all_pages_except.exclude_list / 312cbc38ec4c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-e35eed5363c4bd97535b7099b0c67740891a71d420b45d9f0fe0667cfaadd92d)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-dd03a627bb650ad7f50dccc53927b520d2eca51af57823e3a41e3dfd2ec34643"></a>

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

<a id="canonical-854fe3ec0555fa98cfb060889d275ae6335267fa8d0e89ed91774c935dc87d01"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list / 312cbc38ec4c / 3

- [any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-e1f5b86937b3f9801e8e16588ea71e84047b07c0f15c6139df7c605bd1297cb6): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-014.md#canonical-82fe2d154a8101816b69976cabd0d885b842d5e732d34f70f55ce3762c9dce9f): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-014.md#canonical-a778e9904c5a223b4e5c9fb5e3ed9c9780269d576d1d8eaf0b550c5532ea95e5): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-014.md#canonical-375e03cd2785789f260004466847cac22f87d6b9868126592e9a4c998ba8af08): complete subsection reference.

<a id="canonical-f1ba16f54e0a9d05e8571dbb578530c05e69c6751ba5d8cb2101dac067ff04cc"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list / 312cbc38ec4c / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-e1f5b86937b3f9801e8e16588ea71e84047b07c0f15c6139df7c605bd1297cb6)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain](resources--http_loadbalancer--reference--group-014.md#canonical-82fe2d154a8101816b69976cabd0d885b842d5e732d34f70f55ce3762c9dce9f)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata](resources--http_loadbalancer--reference--group-014.md#canonical-a778e9904c5a223b4e5c9fb5e3ed9c9780269d576d1d8eaf0b550c5532ea95e5)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.path](resources--http_loadbalancer--reference--group-014.md#canonical-375e03cd2785789f260004466847cac22f87d6b9868126592e9a4c998ba8af08)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-e35eed5363c4bd97535b7099b0c67740891a71d420b45d9f0fe0667cfaadd92d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e1f5b86937b3f9801e8e16588ea71e84047b07c0f15c6139df7c605bd1297cb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da67e6596cdaaa749dbe9f379c755d1ac480c550b3d98ac9d6481fc37566c644"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain — client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 7dc0852d975f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-e35eed5363c4bd97535b7099b0c67740891a71d420b45d9f0fe0667cfaadd92d)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0fb2b7902a7a9a19077e2ed4fc8cba546bdbc031a18df7f360f551bb3d4f4177)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-197b50a6fd26bdd065b39b575a6c92259f9dc8ee577a038d11d1574144ef5b1c"></a>

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

<a id="canonical-0c2dcbce3c80b31aa2cb5264d6ff7b9b2ff4703d579295560ad74e62a3898fe3"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 7dc0852d975f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4f0bb3ce4a1a614bd0606e16a0360e432716326e01b64974f12b8239e8be1094"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 7dc0852d975f / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0fb2b7902a7a9a19077e2ed4fc8cba546bdbc031a18df7f360f551bb3d4f4177)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-82fe2d154a8101816b69976cabd0d885b842d5e732d34f70f55ce3762c9dce9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43d120596d92a340f759011a7a0f020b0f7f8b4ddfa36239f81316f8df2d85d7"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / f9084fd9fbe8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-e35eed5363c4bd97535b7099b0c67740891a71d420b45d9f0fe0667cfaadd92d)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0fb2b7902a7a9a19077e2ed4fc8cba546bdbc031a18df7f360f551bb3d4f4177)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-dfa2186288c999b23859cca75e7be2564d9c350336615ec10b3434fbab717500"></a>

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

<a id="canonical-e1fbf8d79b253c0c743d9133fd1ad1e92bd3009ee0882e5965a10905bdbb2539"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / f9084fd9fbe8 / 3

<a id="canonical-fd6cd13be1334a5486b3e1df101ed75bbc81a5eb6d0a815d967fe821ea239c98"></a>

<a id="canonical-e91d5fb800baafb12808ecd1621e6ea33427c30cf251dc34400e7736cc71089b"></a>

## exact_value property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / f9084fd9fbe8 / 4

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

<a id="canonical-586cc1c793056cd769a38b5697555389a4c7a3a284bad1709d13f568f656acf8"></a>

<a id="canonical-173b5265a65ffe2ca9bf48f48a7d4fa55b9642302d993c38bdc2375f6a2c389a"></a>

## regex_value property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / f9084fd9fbe8 / 5

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

<a id="canonical-d08cda41d6721fe7499ef0b2e515d8c08eb4abed51393e4d62d654727254401a"></a>

<a id="canonical-9ad853a9ffb5c24b1b746b087cf16b45f70f04809e390acc9ecbb36038ba5d96"></a>

## suffix_value property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / f9084fd9fbe8 / 6

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

<a id="canonical-874b69d0bb1d8c19b7720f560e373d7f56f95c70b5c93d0a697a9904eaff4229"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain / f9084fd9fbe8 / 7

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0fb2b7902a7a9a19077e2ed4fc8cba546bdbc031a18df7f360f551bb3d4f4177)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a778e9904c5a223b4e5c9fb5e3ed9c9780269d576d1d8eaf0b550c5532ea95e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-878ba068c1aad59ee5359ce0703f3eb867805189d6478b5932baf1d53b93f096"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 99723f92e8d3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-e35eed5363c4bd97535b7099b0c67740891a71d420b45d9f0fe0667cfaadd92d)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0fb2b7902a7a9a19077e2ed4fc8cba546bdbc031a18df7f360f551bb3d4f4177)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-89066db944ea006219e262b002e447dff171c620312cd5e7a144266f790f7d1a"></a>

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

<a id="canonical-e2efc03a4f444932e4ec569d667c1241be586bf6e62561d95ee8442588a34d59"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 99723f92e8d3 / 3

<a id="canonical-561d689f48b32047d757371fb144c10df56c14f8e0673e936d8b28ae12fc4268"></a>

<a id="canonical-b1495ac736724ea75cba8020c7b7b895eb1148d6945d3347774fd8b73fc5e09d"></a>

## description_spec property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 99723f92e8d3 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-5538d8357acaac5e348c067432dced4d85338594b2d542ca8143bcee6278251f"></a>

<a id="canonical-e6091fc6616b1dd859df01212707f4dac3dc0b4a4c35cf8b519f4105154fc04b"></a>

## name property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 99723f92e8d3 / 5

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

<a id="canonical-cc7e34c74a2f1f0cb7a06b04431faddabdfe5c0602783e3e962521fe1917e7f7"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 99723f92e8d3 / 6

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0fb2b7902a7a9a19077e2ed4fc8cba546bdbc031a18df7f360f551bb3d4f4177)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-375e03cd2785789f260004466847cac22f87d6b9868126592e9a4c998ba8af08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-871cac39440ad36cfe5498adf893246c7c71c7b0f385a08135c5c9ebf3f51157"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.path — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / a634bcf2fb5f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-e35eed5363c4bd97535b7099b0c67740891a71d420b45d9f0fe0667cfaadd92d)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0fb2b7902a7a9a19077e2ed4fc8cba546bdbc031a18df7f360f551bb3d4f4177)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-a9010faab3af9748151f59e7f4eeffdbc0357401afc99229096c605de6189275"></a>

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

<a id="canonical-29d367eca37e7ca208acd79b2d7d079561630caba2b06fee517072b695a6b26f"></a>

## Direct properties — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / a634bcf2fb5f / 3

<a id="canonical-5598e838c5c4d2fbe9f3292400d4ca69d11d831c1d726a2483c881a20555d794"></a>

<a id="canonical-c936c58e87511c3c7d2e9ba582749d0b7814ed63775671bb87b9396cce057e4d"></a>

## path property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / a634bcf2fb5f / 4

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

<a id="canonical-3fcf1f38364d2c8b3d3ca3d13f8c1050ddfc7572a38516756948abd8d12d5849"></a>

<a id="canonical-28c4e7be21bf563594679e3df602f28c389c5f74cac6d69676022652eb2b69e6"></a>

## prefix property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / a634bcf2fb5f / 5

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

<a id="canonical-9034d31c86c56fb22b76a60a321b4880bf2c10f6b08bd11ad3eae8a72993507a"></a>

<a id="canonical-0f55e5a31792b5e58e469a5a347bc57e73ce06a11df6ceaa2d7c1694b1d44312"></a>

## regex property — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / a634bcf2fb5f / 6

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

<a id="canonical-c64a0bd87c3c809c0fe1d3c3af2db4aaa6d65e3d896089782b9c175404f874cc"></a>

## Next pages — client_side_defense.policy.js_insert_all_pages_except.exclude_list.path / a634bcf2fb5f / 7

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0fb2b7902a7a9a19077e2ed4fc8cba546bdbc031a18df7f360f551bb3d4f4177)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a73642cbc01668bba50877a4f39b54a9a5453ccb105a2a48f0c633346e6813c0"></a>

## client_side_defense.policy.js_insertion_rules — client_side_defense.policy.js_insertion_rules / 6ffd43660b36 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- client_side_defense.policy.js_insertion_rules

<a id="canonical-abf1810ffca16c4f310db075f8a39def5efbf522c36dd7be5cb63bda27ac745c"></a>

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

<a id="canonical-16914dec03fca3bdeb140a0a06be365062a345415639ef54bfb5f27e8dce88d4"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules / 6ffd43660b36 / 3

- [exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2663d10d9a694c4e907463d77925b7e38dd872d7821a4520813e4e28911c7fea): complete subsection reference.

- [rules](resources--http_loadbalancer--reference--group-014.md#canonical-8372d741bc57378682f7cdb27834f51ac15283d2b827223afb79f2192488e539): complete subsection reference.

<a id="canonical-247d1a3435ddb75afd76458dfb74e3909d0060d8b8dd1bfca209edbc61b145f8"></a>

## Next pages — client_side_defense.policy.js_insertion_rules / 6ffd43660b36 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2663d10d9a694c4e907463d77925b7e38dd872d7821a4520813e4e28911c7fea)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-8372d741bc57378682f7cdb27834f51ac15283d2b827223afb79f2192488e539)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2663d10d9a694c4e907463d77925b7e38dd872d7821a4520813e4e28911c7fea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-632b22c6ace0a0774b725c185f087b6aea4c054d07d6f4ff693179ed17fb9035"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list — client_side_defense.policy.js_insertion_rules.exclude_list / 42d7746329ba / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b)
- client_side_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-328251fcd8198d944b9b4e7a518ef5e1614aa65e68439ea4daf0fcf6aeeb0738"></a>

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

<a id="canonical-5c275830ed4d8d26a01bae69d5c1eb1b621682effc9bd9359922abce3089da51"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list / 42d7746329ba / 3

- [any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-653979cd763e8ccc8074e8cc2444bc3d7203122bc9e2d8c2cee9718acbade3a2): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-014.md#canonical-727f2771328c4f84ce8b2dbd7ad8b2f81993d097fde8d353e6798457e404002d): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-014.md#canonical-3be71e9dcd60d456a3bfc8266abdd403340aa24c65d9c481b8da34e7469074b6): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-014.md#canonical-e511b1a8546ed2d74bfcfacfda8f6184f9b028190f9212fd9ec5e620c9ddceb9): complete subsection reference.

<a id="canonical-18fb7a6f80e5235f5b1d2cfd94f5a0f4a25698349f2c817ceb0568af49fa64b7"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list / 42d7746329ba / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list.any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-653979cd763e8ccc8074e8cc2444bc3d7203122bc9e2d8c2cee9718acbade3a2)
- [client_side_defense.policy.js_insertion_rules.exclude_list.domain](resources--http_loadbalancer--reference--group-014.md#canonical-727f2771328c4f84ce8b2dbd7ad8b2f81993d097fde8d353e6798457e404002d)
- [client_side_defense.policy.js_insertion_rules.exclude_list.metadata](resources--http_loadbalancer--reference--group-014.md#canonical-3be71e9dcd60d456a3bfc8266abdd403340aa24c65d9c481b8da34e7469074b6)
- [client_side_defense.policy.js_insertion_rules.exclude_list.path](resources--http_loadbalancer--reference--group-014.md#canonical-e511b1a8546ed2d74bfcfacfda8f6184f9b028190f9212fd9ec5e620c9ddceb9)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-653979cd763e8ccc8074e8cc2444bc3d7203122bc9e2d8c2cee9718acbade3a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bab38debc0d7adcd2c1f9e422b8ed990e1df1eb88c397e90c8d2d22d8efc9c6d"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.any_domain — client_side_defense.policy.js_insertion_rules.exclude_list.any_domain / 39a2a640da78 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2663d10d9a694c4e907463d77925b7e38dd872d7821a4520813e4e28911c7fea)
- client_side_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-b30465f10b83811de81d7bb35fa28d3850b7ded7ba3d9b633f5055b8431be82b"></a>

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

<a id="canonical-af244dd00564a9f0965a6ef77e7e6e41a7c81372576d1c5a914165f5b4ceaf33"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.any_domain / 39a2a640da78 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-47a10f28950970fcb5a01946406671702ccac0f3e4001e9ec24ad6334162fbb6"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.any_domain / 39a2a640da78 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2663d10d9a694c4e907463d77925b7e38dd872d7821a4520813e4e28911c7fea)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-727f2771328c4f84ce8b2dbd7ad8b2f81993d097fde8d353e6798457e404002d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd9274642fde8cd77642c4b27ff555ba7bf45bbf6ab28dbc515e89e8a4bddff3"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.domain — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 56aad4451f94 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2663d10d9a694c4e907463d77925b7e38dd872d7821a4520813e4e28911c7fea)
- client_side_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-6964ffb58603169b973401b2c6c5fde96a47497042b52e533e2c6c3770798d97"></a>

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

<a id="canonical-691dc7bde41bb3ec1c24084230d35c2cbce6c2e1c7ea5562a7d025d313c4eafd"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 56aad4451f94 / 3

<a id="canonical-5a4280ee4ade317a58371d2c20aead2e3752e8dcfc2bbf2f0a908b6be80498a5"></a>

<a id="canonical-b5774d479a3120a546dd0443f66874d2a1e9e7780a720b50565beeb2e9390847"></a>

## exact_value property — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 56aad4451f94 / 4

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

<a id="canonical-0d27c29af823132764281891616e2f5eb06c0ba3b4b831e1374799d54a905e47"></a>

<a id="canonical-14f38173b048269bd91b10968dcdfdd2f2d5a94dc9c460817409b99062d29ce5"></a>

## regex_value property — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 56aad4451f94 / 5

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

<a id="canonical-e66458fd5845cdcdffd33493662fdb753c245f479ba0c99cfe408b25886362a6"></a>

<a id="canonical-0472648abd389c3d4311f943f221d1c07cedbe24a8ed172d3f7c3ac5f696c29d"></a>

## suffix_value property — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 56aad4451f94 / 6

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

<a id="canonical-5094da3f10c4a0ccf63c41cd2a377abd50b1b242b980c3c8c8443ef6b7d1f5ce"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.domain / 56aad4451f94 / 7

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2663d10d9a694c4e907463d77925b7e38dd872d7821a4520813e4e28911c7fea)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3be71e9dcd60d456a3bfc8266abdd403340aa24c65d9c481b8da34e7469074b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d1f30fa2cef4756815a8d26b83e70d231b937b63164fb6cee30480befd96ab9"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.metadata — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 55d275bc52c8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2663d10d9a694c4e907463d77925b7e38dd872d7821a4520813e4e28911c7fea)
- client_side_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-2d661741744ebedfa49134194cfaf4fc803986b510e9edfe04eebc4da0ecfac4"></a>

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

<a id="canonical-37bc8c5781520a80915cd76eba0b993904128d5c7cca897e975356f3f34245bb"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 55d275bc52c8 / 3

<a id="canonical-7d8eb483a625fff22c2c9c0142b5faf45f44b5f8d910d637738b23f71081c083"></a>

<a id="canonical-a1a2c97b25e3c6f970363dd283b0ab3c3a521b9217c817cff7e363e74cb29a23"></a>

## description_spec property — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 55d275bc52c8 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-db6c33e6038fa64e21cefe79362b1af1571430bf43cbc6d177fe1b6df5036b59"></a>

<a id="canonical-f637e8f46d18d16c5bd1e5cc3a9f3d7e972c846f245dc2a6315153a8f17a3488"></a>

## name property — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 55d275bc52c8 / 5

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

<a id="canonical-5e19817ffa0651061fb55e03a1ac4d16d98167c2e044ec2d5e7912b6303ac665"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.metadata / 55d275bc52c8 / 6

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2663d10d9a694c4e907463d77925b7e38dd872d7821a4520813e4e28911c7fea)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e511b1a8546ed2d74bfcfacfda8f6184f9b028190f9212fd9ec5e620c9ddceb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-393997b12ffcdafc09bc87ea0cbfc845df66714fb594736fedf1188f8928afb0"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.path — client_side_defense.policy.js_insertion_rules.exclude_list.path / 0ebf0e07f84e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2663d10d9a694c4e907463d77925b7e38dd872d7821a4520813e4e28911c7fea)
- client_side_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-270681911e3d1e91fd0b35958f6903df5a2e55312534477827bdc1cb39b6d22a"></a>

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

<a id="canonical-618c77441fbe46480510edb0ec87db003cfa38199322d7a46a12695a21b5caca"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.exclude_list.path / 0ebf0e07f84e / 3

<a id="canonical-c19da9ef3258eb170463fa2e38f395b8d62f97985634c8a7bfc3e2c007f656ba"></a>

<a id="canonical-be9cac45912ba6466a8205fe59fc3a714c911dd4e6b344421f32af64ce5a26e6"></a>

## path property — client_side_defense.policy.js_insertion_rules.exclude_list.path / 0ebf0e07f84e / 4

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

<a id="canonical-28ee20cfd0c8b57843209e190f647e8301a8208458c0b93f45e4fd1015c0de2f"></a>

<a id="canonical-01c89795b99c796593d5ffee55d751c3ced253f281c9b3c0d8678ced0b1496cc"></a>

## prefix property — client_side_defense.policy.js_insertion_rules.exclude_list.path / 0ebf0e07f84e / 5

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

<a id="canonical-25fbc960b0d0611ff51a0f286791de277d0469cf4c5c1548bd113b17b8f3badc"></a>

<a id="canonical-9907613bf1a77fc2bb35ba734a8d96a2628f3efa403994bb6c10721a3be70067"></a>

## regex property — client_side_defense.policy.js_insertion_rules.exclude_list.path / 0ebf0e07f84e / 6

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

<a id="canonical-b579e185a91a93316925bd10c52c46bc8642fea20de3e4045e47cdb59c58fa03"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.exclude_list.path / 0ebf0e07f84e / 7

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2663d10d9a694c4e907463d77925b7e38dd872d7821a4520813e4e28911c7fea)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8372d741bc57378682f7cdb27834f51ac15283d2b827223afb79f2192488e539"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a45217581f645679c6fd02c73bbe0ef0a00e112bac7c8ec2f0d823c62686a3f0"></a>

## client_side_defense.policy.js_insertion_rules.rules — client_side_defense.policy.js_insertion_rules.rules / 7cfb7e3600a0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b)
- client_side_defense.policy.js_insertion_rules.rules

<a id="canonical-2c073f0a7295ae6a560edbe95a03da09e0a126097fe25aa0c2dd6fabf91e49d2"></a>

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

<a id="canonical-3d71ebbe69fa4681222ac70ee7c12568d974eda7dce0b5a5bc05e8182ee8989a"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules / 7cfb7e3600a0 / 3

- [any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-87d8f3b57b85b98386b7129763c049a13d6f30568f33cbfab131c084848a9a48): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-014.md#canonical-d9a3cd6d1d5cc9fe2c2de824bdc0eb3da1de4db501414a730ea79ea961061f37): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-014.md#canonical-10f18b6ed8aac51169e42bc424121c6ebec7f695977e5c9467ac0a3035579df0): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-014.md#canonical-50fdc007289f5f029ad72537229802f8978240656d0cd2a9d3fe7e6eaf71e7a4): complete subsection reference.

<a id="canonical-41e55921a872ec587d771c231a4f10bcf24bed1e42bcef325114a276c6e9fb12"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules / 7cfb7e3600a0 / 4

- [client_side_defense.policy.js_insertion_rules.rules.any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-87d8f3b57b85b98386b7129763c049a13d6f30568f33cbfab131c084848a9a48)
- [client_side_defense.policy.js_insertion_rules.rules.domain](resources--http_loadbalancer--reference--group-014.md#canonical-d9a3cd6d1d5cc9fe2c2de824bdc0eb3da1de4db501414a730ea79ea961061f37)
- [client_side_defense.policy.js_insertion_rules.rules.metadata](resources--http_loadbalancer--reference--group-014.md#canonical-10f18b6ed8aac51169e42bc424121c6ebec7f695977e5c9467ac0a3035579df0)
- [client_side_defense.policy.js_insertion_rules.rules.path](resources--http_loadbalancer--reference--group-014.md#canonical-50fdc007289f5f029ad72537229802f8978240656d0cd2a9d3fe7e6eaf71e7a4)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-87d8f3b57b85b98386b7129763c049a13d6f30568f33cbfab131c084848a9a48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f23fb8dd12b352e5d7687f1124669f7b65bfc40a8e78e5bd62fb3addfeb8072"></a>

## client_side_defense.policy.js_insertion_rules.rules.any_domain — client_side_defense.policy.js_insertion_rules.rules.any_domain / f13c968ed9e1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-8372d741bc57378682f7cdb27834f51ac15283d2b827223afb79f2192488e539)
- client_side_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-ac2c35cae994f652389d223c69bd373d9e593f4feeed7b27b7ade42dc0bbde41"></a>

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

<a id="canonical-b74041d8b829c4350b172a63fe001a87590eae63a956ee907f9061f9ffd42b24"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.any_domain / f13c968ed9e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c767cab3a9b689c6e34d8a365261ae00ed856c06cf8c8f1236280f288514e5ce"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.any_domain / f13c968ed9e1 / 4

- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-8372d741bc57378682f7cdb27834f51ac15283d2b827223afb79f2192488e539)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d9a3cd6d1d5cc9fe2c2de824bdc0eb3da1de4db501414a730ea79ea961061f37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6ee241d465d9aa07885c9bed64f65b1010c036ca29e54f059dbab1c899340c7"></a>

## client_side_defense.policy.js_insertion_rules.rules.domain — client_side_defense.policy.js_insertion_rules.rules.domain / 6cc4d64ae339 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-8372d741bc57378682f7cdb27834f51ac15283d2b827223afb79f2192488e539)
- client_side_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-72982f3758229dd759db73cc3a59701004b21145a2c20be418b0dffcaabfb40e"></a>

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

<a id="canonical-507eb7f151ce2cfa780881394613b22679ca4ca8fd2bdcc911a4047c7a642b19"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.domain / 6cc4d64ae339 / 3

<a id="canonical-ea4d2b84788827342a032660efe4835aa504cb1f35d6a1c30c4383210ad5a5fe"></a>

<a id="canonical-cb40e4ebc61b5330151bb50dbbf5725a5b2cb176f986de622594e9ef104362f6"></a>

## exact_value property — client_side_defense.policy.js_insertion_rules.rules.domain / 6cc4d64ae339 / 4

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

<a id="canonical-7d48eeb9e65d3453e84235256f8c2afb2b380b6a558a8b2e0de1d7198aab0a06"></a>

<a id="canonical-f2f1eba1b681b011ed93f47430cbe43473f63301e2dc3f75dc8b1eaf4ca03847"></a>

## regex_value property — client_side_defense.policy.js_insertion_rules.rules.domain / 6cc4d64ae339 / 5

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

<a id="canonical-5dc9dca723d6eebd50dc4dc07fd7a9ee8daf31e7c0ddc012cdab531eb93e24a8"></a>

<a id="canonical-606e80cd47ed3360b84a4493f87d9ab43d21e47a9be8a192d6dfaa40169ffd08"></a>

## suffix_value property — client_side_defense.policy.js_insertion_rules.rules.domain / 6cc4d64ae339 / 6

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

<a id="canonical-e09cc9568fccd220404350db2164afb5d9a8d8e726c38bec9c3763eee28a30d9"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.domain / 6cc4d64ae339 / 7

- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-8372d741bc57378682f7cdb27834f51ac15283d2b827223afb79f2192488e539)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-10f18b6ed8aac51169e42bc424121c6ebec7f695977e5c9467ac0a3035579df0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60556640cb5abe0435ebf0111178ba8d31b413f5cdd3192555ee81ad02c6b8e9"></a>

## client_side_defense.policy.js_insertion_rules.rules.metadata — client_side_defense.policy.js_insertion_rules.rules.metadata / 8271cd7e215a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-8372d741bc57378682f7cdb27834f51ac15283d2b827223afb79f2192488e539)
- client_side_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-62e7665e5034b59d8d2db9adfb8d088702a020173efd1372c1b675d561a4aee3"></a>

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

<a id="canonical-c4a0df9f132d8a4b805eee0f66bd1ec2bae62a93d81764b8d050795ab03d0f45"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.metadata / 8271cd7e215a / 3

<a id="canonical-87e701d5221c66a428fc2f9b88285a5f0678b9df1a7e754708e8426a197d8f6c"></a>

<a id="canonical-5fce70092551909abf404e06a640cb62f2703df9a17868851a323c1961b397e4"></a>

## description_spec property — client_side_defense.policy.js_insertion_rules.rules.metadata / 8271cd7e215a / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-4c0cae52ea88d9119f4afd0784b06473a82b14abe16e8b51b55834a4d42f55cc"></a>

<a id="canonical-1b19ee210bca5fd8ebb56dfa5f138d4a86a66e0a72251a9b0fe59bcd0dd0d1b0"></a>

## name property — client_side_defense.policy.js_insertion_rules.rules.metadata / 8271cd7e215a / 5

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

<a id="canonical-d817c9316634cb5791e9077737054f554144ef316062147e45a8af4cdeb08b5d"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.metadata / 8271cd7e215a / 6

- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-8372d741bc57378682f7cdb27834f51ac15283d2b827223afb79f2192488e539)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-50fdc007289f5f029ad72537229802f8978240656d0cd2a9d3fe7e6eaf71e7a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84c3f67c7a237b800f1c4d9acbc9c099c95245ac0dce863fbc83d52e3cf10761"></a>

## client_side_defense.policy.js_insertion_rules.rules.path — client_side_defense.policy.js_insertion_rules.rules.path / 0b67d1e6df71 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-354a75b4365663abe51f2bf7166a7658702f51f87b9dd01d146f09ceac63c1aa)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-68c5693b0b6005d4526b8ce8a8197548da5cc005bcae4f7483da83d55b903e74)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-ecf5029b40b6bb1ffbe978f17210bd6bb4882910a8c33ddbbbfe2d0c586f120b)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-8372d741bc57378682f7cdb27834f51ac15283d2b827223afb79f2192488e539)
- client_side_defense.policy.js_insertion_rules.rules.path

<a id="canonical-c5ef32cc3514fadf95f547deb460d62408c21cf5aecd4818a5f9eb20be461c6d"></a>

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

<a id="canonical-4d2863830c70e6141de1666a05b0ac6f6a906eda2f84af1eb9625cfd5369b15a"></a>

## Direct properties — client_side_defense.policy.js_insertion_rules.rules.path / 0b67d1e6df71 / 3

<a id="canonical-c8eb6bb47f3b3bdcf71190b5b6d7c52dcfa0e6b66eadd21b39a1996e3206a76f"></a>

<a id="canonical-b69357f0132db5d87fa39f059ae49b2a7c9aa6aeda40aff18656c82b77b01bf4"></a>

## path property — client_side_defense.policy.js_insertion_rules.rules.path / 0b67d1e6df71 / 4

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

<a id="canonical-39d45d66967c0405bddff27a7f36740bbd4a4725aadf8249fb2d023050d6cd06"></a>

<a id="canonical-4cb08da2f335a3f4631847840b1094b510f156673c7063a3befa7e441699be8e"></a>

## prefix property — client_side_defense.policy.js_insertion_rules.rules.path / 0b67d1e6df71 / 5

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

<a id="canonical-1c14ff135ce5460ed77d53fc509bdb036f90a2b3ccb1c4e8084cd109df0142d7"></a>

<a id="canonical-51bfbfd2b950253eac84cbbdced60bd6f88bab7d4c876b325cfb870f5d2f8a8c"></a>

## regex property — client_side_defense.policy.js_insertion_rules.rules.path / 0b67d1e6df71 / 6

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

<a id="canonical-d09ae600fe25b96cdb6f6cc3a4c5ecf653196d032bcc405b886a74c1de6c7fe8"></a>

## Next pages — client_side_defense.policy.js_insertion_rules.rules.path / 0b67d1e6df71 / 7

- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-8372d741bc57378682f7cdb27834f51ac15283d2b827223afb79f2192488e539)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d4c290efdfa8fdf49863e9bc1584cc85c7a2cdb431534a160fba71ab00032f5"></a>

## cookie_stickiness — cookie_stickiness / 82ac088af244 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- cookie_stickiness

<a id="canonical-ca41c73ba5200268de509e46e1cee06e17e0c2fae25b07bdbf85150b5e4719e3"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: cookie\_stickiness, least\_active, random, ring\_hash, round\_robin,
source\_ip\_stickiness; Default: round\_robin\] Two types of cookie affinity: 1. Passive. Takes a
cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and
sets a cookie with an expiration (TTL) on the first request from the client in its response to the
client, based on the endpoint the request gets..

Upstream description:

Two types of cookie affinity:

&#8203;1. Passive. Takes a cookie that's present in the cookies header and hashes on its value.

&#8203;2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from
the client in its response to the client, based on the endpoint the request gets sent to. The client
then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
these requests GET sent to the same endpoint. The cookie is generated by hashing the source and
destination ports and addresses so that multiple independent HTTP2 streams on the same connection
will independently receive the same cookie, even if they arrive simultaneously.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name"),
  validators.ConflictingObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingObjectAttributes("samesite_none",
    "samesite_strict")}
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
  "x-ves-oneof-field-httponly": "[\"add_httponly\",\"ignore_httponly\"]",
  "x-ves-oneof-field-samesite": "[\"ignore_samesite\",\"samesite_lax\",\"samesite_none\",\"samesite_strict\"]",
  "x-ves-oneof-field-secure": "[\"add_secure\",\"ignore_secure\"]"
}
```

OneOf alternatives in this subsection:

- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-ca41c73ba5200268de509e46e1cee06e17e0c2fae25b07bdbf85150b5e4719e3)
- [least_active](resources--http_loadbalancer--reference--group-020.md#canonical-063684637d11748256fee74bca48a125d8741c4075b2785483e6fc0b832132a5)
- [random](resources--http_loadbalancer--reference--group-023.md#canonical-49dc362bd194bac07ff10ed2cf8972b7955e80db56085bfb7631604655215439)
- [ring_hash](resources--http_loadbalancer--reference--group-023.md#canonical-c791a5991aa7af7807b966fd867b8f6a5eb72d0e91be60fd43fc1923f680defe)
- [round_robin](resources--http_loadbalancer--reference--group-023.md#canonical-58b46b94c19530fe94fa7794af84b6c4fd342bbba1ddaf22f9be84dffc3a3aa3)
- [source_ip_stickiness](resources--http_loadbalancer--reference--group-026.md#canonical-22b74d22d28f94180a0f3f4f765f6c85bd492b35036db22fb390f2d3ba3cfd15)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
cookie_stickiness {
  # Configure direct properties listed below.
}
```

<a id="canonical-eb656187446f6fde034caecd5d82cda9ffa6656e8028c10976f5b3c0727ad275"></a>

## Direct properties — cookie_stickiness / 82ac088af244 / 3

- [add_httponly](resources--http_loadbalancer--reference--group-014.md#canonical-8c75888b94fa11242bd01e00b71b43d6de67b240fb3931afe35a4996bcf0893f): complete subsection reference.

- [add_secure](resources--http_loadbalancer--reference--group-014.md#canonical-84169db5a1dd3f3834dcf7a7ec1a9ab81abccbdfa77d4293a10c52ce749f889f): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-014.md#canonical-3e3a168a945ed453cc61992d93ba848883553cba8c224d9b0bcc19a0a2f21864): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-014.md#canonical-2008ec024ed98c0b9819b7d67e79aa075063dc458627fdce4561e26d164c747c): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-014.md#canonical-0fb332f161e0a89604c7d1a803c0d02d2b5f6a287f6a9ab48a128f2cf80a01b1): complete subsection reference.

<a id="canonical-f046ae0b1b16b37102bdc0da05a8089803ea0b9d563c6ddb64a0e7593ee600a5"></a>

<a id="canonical-c57d89a6cbdbc3b18de0c805b9664517bbe16084231138360944d7c029b641eb"></a>

## name property — cookie_stickiness / 82ac088af244 / 4

Type: `"string"`. Optional.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Upstream description:

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-8b0e17c70ddae5dafc2920dfa4764c0342eb4f0c29247f29508c4dc6b02a3f2f"></a>

<a id="canonical-551943a675b1e8c20b037e2902ebac15dcc653cfe55b47f21f0afd3bc829528b"></a>

## path property — cookie_stickiness / 82ac088af244 / 5

Type: `"string"`. Optional.

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Upstream description:

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
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
  }
}
```

- [samesite_lax](resources--http_loadbalancer--reference--group-014.md#canonical-79063befb9b8324a2315d1b8e329bc2bc246d23cf26f9c33ab9e2fe2fcad769e): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-014.md#canonical-5e756c805ac7bcdf840de22bc699a78c6226a6c82bdf76de983b0f36b01d57e0): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-014.md#canonical-4584885e0247c4f003760adfad05f1b46c8b45694ffb22ab6e58ad5a65e0a6c6): complete subsection reference.

<a id="canonical-2eb5723aca1963faff989da8078d079735be64c13e126c08139a36649d5df291"></a>

<a id="canonical-1f7e1febeae5d72155eaeef4a82d83802010eb54eb6176626af2935b4efbde5d"></a>

## ttl property — cookie_stickiness / 82ac088af244 / 6

Type: `"number"`. Optional.

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

Upstream description:

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-389da0f64db4269afe52b02cc1eba547ca088aa14d0a84aff69d647c69119ffc"></a>

## Next pages — cookie_stickiness / 82ac088af244 / 7

- [cookie_stickiness.add_httponly](resources--http_loadbalancer--reference--group-014.md#canonical-8c75888b94fa11242bd01e00b71b43d6de67b240fb3931afe35a4996bcf0893f)
- [cookie_stickiness.add_secure](resources--http_loadbalancer--reference--group-014.md#canonical-84169db5a1dd3f3834dcf7a7ec1a9ab81abccbdfa77d4293a10c52ce749f889f)
- [cookie_stickiness.ignore_httponly](resources--http_loadbalancer--reference--group-014.md#canonical-3e3a168a945ed453cc61992d93ba848883553cba8c224d9b0bcc19a0a2f21864)
- [cookie_stickiness.ignore_samesite](resources--http_loadbalancer--reference--group-014.md#canonical-2008ec024ed98c0b9819b7d67e79aa075063dc458627fdce4561e26d164c747c)
- [cookie_stickiness.ignore_secure](resources--http_loadbalancer--reference--group-014.md#canonical-0fb332f161e0a89604c7d1a803c0d02d2b5f6a287f6a9ab48a128f2cf80a01b1)
- [cookie_stickiness.samesite_lax](resources--http_loadbalancer--reference--group-014.md#canonical-79063befb9b8324a2315d1b8e329bc2bc246d23cf26f9c33ab9e2fe2fcad769e)
- [cookie_stickiness.samesite_none](resources--http_loadbalancer--reference--group-014.md#canonical-5e756c805ac7bcdf840de22bc699a78c6226a6c82bdf76de983b0f36b01d57e0)
- [cookie_stickiness.samesite_strict](resources--http_loadbalancer--reference--group-014.md#canonical-4584885e0247c4f003760adfad05f1b46c8b45694ffb22ab6e58ad5a65e0a6c6)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8c75888b94fa11242bd01e00b71b43d6de67b240fb3931afe35a4996bcf0893f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-062c3c858b4e8953642159c797348e1afa934b28b252a3a4521fc6e9145e3af6"></a>

## cookie_stickiness.add_httponly — cookie_stickiness.add_httponly / 319aabe145a3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- cookie_stickiness.add_httponly

<a id="canonical-fe22accbd16a75d4d1c8a460e5dcb99d11eab04d724f54965512d1401c492367"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

<a id="canonical-8c2c085059a2c77ab9d27d14ce4cc0eee6f7b6b23a6f895788d535cdb4c79321"></a>

## Direct properties — cookie_stickiness.add_httponly / 319aabe145a3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d3e20f720dfec1c2871b93a856294739157eb2a1700b3b4da88b7995736776a9"></a>

## Next pages — cookie_stickiness.add_httponly / 319aabe145a3 / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-84169db5a1dd3f3834dcf7a7ec1a9ab81abccbdfa77d4293a10c52ce749f889f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7698a9796a629f0b1e14f6fedf783f268930c4611e43cbd6952db192fddbfb41"></a>

## cookie_stickiness.add_secure — cookie_stickiness.add_secure / 0801186af33e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- cookie_stickiness.add_secure

<a id="canonical-492fa91aa640bf48f80e7876f7ff1336b9c270ac38a50e025d93e5d59c7893ed"></a>

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
add_secure = {}
```

<a id="canonical-c9925010b656170f9a243e31188c15a90d95c3b4b3f63cb949b7b6d8619181b6"></a>

## Direct properties — cookie_stickiness.add_secure / 0801186af33e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8146278499011330ff3b17fd49c136d6706e63b5c051f21b735d7a89acd9f2a0"></a>

## Next pages — cookie_stickiness.add_secure / 0801186af33e / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3e3a168a945ed453cc61992d93ba848883553cba8c224d9b0bcc19a0a2f21864"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa269e5b799cb951baa40cd4e32cb87fc6acf56e7d1bb3451cac017201e6d396"></a>

## cookie_stickiness.ignore_httponly — cookie_stickiness.ignore_httponly / 2b7c82c99c6c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- cookie_stickiness.ignore_httponly

<a id="canonical-b4423422f6bd1f8fa8c9d01dc132dd3f8c7fc13c1f04694982fe9c2cd35fd571"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

<a id="canonical-c3aebf95d72230244fc83a11195d9f76740ea665df38bfa140f06135f624c7f2"></a>

## Direct properties — cookie_stickiness.ignore_httponly / 2b7c82c99c6c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-75cc0def3505f0dd07c87401989466d5776954f454e0f6524d5d2037ae0bc99e"></a>

## Next pages — cookie_stickiness.ignore_httponly / 2b7c82c99c6c / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2008ec024ed98c0b9819b7d67e79aa075063dc458627fdce4561e26d164c747c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c119b09ed5e3dc7f0455253dd9999c9af4faf720025f8bc5f62425cd2930020"></a>

## cookie_stickiness.ignore_samesite — cookie_stickiness.ignore_samesite / 7c6ec772cc8d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- cookie_stickiness.ignore_samesite

<a id="canonical-23ee44adc76714601514e01195f56c9f187ca6f63a1cc0e6514478c1bb39ac84"></a>

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
ignore_samesite = {}
```

<a id="canonical-ef13887781fb495c3d408a41e5dab476dcfcd38cddf8506dcecabe456ad7116c"></a>

## Direct properties — cookie_stickiness.ignore_samesite / 7c6ec772cc8d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ff67ece78b76b9788da17d6c1e3429c58ae420ceb94debc0a533f3423ea96781"></a>

## Next pages — cookie_stickiness.ignore_samesite / 7c6ec772cc8d / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0fb332f161e0a89604c7d1a803c0d02d2b5f6a287f6a9ab48a128f2cf80a01b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df0e2e2c7613a590961043a353b81b4bbc63baf9ed1af38c26b5c6d6b7c5d3e5"></a>

## cookie_stickiness.ignore_secure — cookie_stickiness.ignore_secure / 83966a8c3038 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- cookie_stickiness.ignore_secure

<a id="canonical-cd5e832701314936b3a3ea4234a36c98798841a226a75a89c97ee4252e3d16ba"></a>

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
ignore_secure = {}
```

<a id="canonical-6b809939acfd584882cab5ffb1d532b972693e9968aabeb218a7a25ee82d42f0"></a>

## Direct properties — cookie_stickiness.ignore_secure / 83966a8c3038 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-36ee6ce32ca9df93400fd5311f501bf4fd3e75f10e45183fc7b3cb1d597fe661"></a>

## Next pages — cookie_stickiness.ignore_secure / 83966a8c3038 / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-79063befb9b8324a2315d1b8e329bc2bc246d23cf26f9c33ab9e2fe2fcad769e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17651a9d3cc15f1dc92b329357d6612879ce2dd49c729b771f983c5a7993fb64"></a>

## cookie_stickiness.samesite_lax — cookie_stickiness.samesite_lax / 6a7d3544d025 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- cookie_stickiness.samesite_lax

<a id="canonical-a2b26674adaa59fe87b7525405b6d6984de3c5898e62ec04ba1dfb7836da7910"></a>

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
samesite_lax = {}
```

<a id="canonical-d8ab50baffacc3d9de0fc72864b7fcd57930ab90715da2dec1a96ea1e875e278"></a>

## Direct properties — cookie_stickiness.samesite_lax / 6a7d3544d025 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9ebb7faf9dbfa42ed49c84e25c7ceef3a53ff454a5a2592510b02c5d8d1c99d8"></a>

## Next pages — cookie_stickiness.samesite_lax / 6a7d3544d025 / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5e756c805ac7bcdf840de22bc699a78c6226a6c82bdf76de983b0f36b01d57e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e39a5e17b7aa20190c21351e0ee39f3aa0675198df67d5a3912c87d09a19d262"></a>

## cookie_stickiness.samesite_none — cookie_stickiness.samesite_none / ab04c41e7037 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- cookie_stickiness.samesite_none

<a id="canonical-2566ea8726ce8fc89faec4c10392ca4ee384bcacee77b319f2e1e81534a8d516"></a>

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
samesite_none = {}
```

<a id="canonical-e53047f7fd748ec2a94ccb2067ac25276e49cc7dc7c7295e891fa4dc0b5bd09a"></a>

## Direct properties — cookie_stickiness.samesite_none / ab04c41e7037 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e155ce076527f21763d3f4832556ff9f0273e2f2c36a3a4b259367b7b5095e18"></a>

## Next pages — cookie_stickiness.samesite_none / ab04c41e7037 / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4584885e0247c4f003760adfad05f1b46c8b45694ffb22ab6e58ad5a65e0a6c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a9866b7dd21df6a0b0f34744ec0065098fbbe0d52019ed4d18bd728f480b407"></a>

## cookie_stickiness.samesite_strict — cookie_stickiness.samesite_strict / 80d1ec7d6404 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- cookie_stickiness.samesite_strict

<a id="canonical-50bee3b128cb38977169f28576d44c4672987955a5a4e8a31214e5375a768a0f"></a>

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
samesite_strict = {}
```

<a id="canonical-e36f8dc17fde1e23a8284ebe913c0bf3f1d295ad9671cf0e7520f4b8b7123300"></a>

## Direct properties — cookie_stickiness.samesite_strict / 80d1ec7d6404 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-22ac11b2a22c9974e39b03c8d3510d17c46ac25b074dda44a87409458bd20cc3"></a>

## Next pages — cookie_stickiness.samesite_strict / 80d1ec7d6404 / 4

- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-faaa86207cf7661dbbb4ee2619564b3e45dd1f0a02c98c62148376af44cd3e95)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-98b1bc48ab39bfd42acdddd7722ced6b2d8018703544313e54793d9a49e8306d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe8dc3429c950da0fe7da24209bc1ed633cb2b7c9418ec25fc5aeb7af61236c4"></a>

## cors_policy — cors_policy / 41ea9f0f8843 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- cors_policy

<a id="canonical-9aa24f232819f19913728680f04f3203f179243a2de06f810be04510073c171c"></a>

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

<a id="canonical-7f664efba9a516918180143ff06e50966e2f5ac86b819c694b75d3b1ae11fe4b"></a>

## Direct properties — cors_policy / 41ea9f0f8843 / 3

<a id="canonical-93d88223ffa06a7845a6747a4de3b04da9065dfd9359c09be89bcb7d3b53d013"></a>

<a id="canonical-916e22c58ef171b0a014dbfecd3f52076192457fdcb9d1737056efca590191ee"></a>

## allow_credentials property — cors_policy / 41ea9f0f8843 / 4

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

<a id="canonical-3679b9d83f30656d2c72e51d545a80209ff337358ef90a8164199edee64b6868"></a>

<a id="canonical-a6f05ed1a9767414b841b5000def2903b3049b9c4bc0e474fa9551463de03e3c"></a>

## allow_headers property — cors_policy / 41ea9f0f8843 / 5

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

<a id="canonical-67b5d475d576b6a84d4f1e562c35e04b320a8fb4d8fb72deb47ba5ac9527c52e"></a>

<a id="canonical-9d9a6be6f29c8e4488edef6f0e72ef42952d776d699d7d5fc63c577fca3af52c"></a>

## allow_methods property — cors_policy / 41ea9f0f8843 / 6

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

<a id="canonical-109bb5893dd40d35c3a4350daae3a9cc55fba4c2ca597ad1e358d40cc38ddc4b"></a>

<a id="canonical-ed514b33cf28eb0b36a856095f2e653219ce22595833e9236a33fb1bb3922cbb"></a>

## allow_origin property — cors_policy / 41ea9f0f8843 / 7

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

<a id="canonical-1f236cb8b91fab23db64c66151ca07f9e21d91be427ad8bd6a171886f777f31f"></a>

<a id="canonical-f38457c8660fca4b69e4cd36c3ddcedb73514d217c1285c53aa8502e69703f6f"></a>

## allow_origin_regex property — cors_policy / 41ea9f0f8843 / 8

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

<a id="canonical-8a44b12418a84cea49acd64495ae65b676f11d8631b0eb64f20df234efd90dcb"></a>

<a id="canonical-dd1ca717110d1d166a4f417fca96958d63c6ffea23046ea32ea30e59dfe87103"></a>

## disabled property — cors_policy / 41ea9f0f8843 / 9

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

<a id="canonical-3d8f16f7ab6d233befd8f6c7b6935f05e8addc56506fb9995c73b7fd1d0950b9"></a>

<a id="canonical-0800da17f3f86e6cc2e9b7bd409c3b81aa0670c1a89ab23597ee4bc589e3317a"></a>

## expose_headers property — cors_policy / 41ea9f0f8843 / 10

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

<a id="canonical-6fff16ad923a8fc9bdb3b4ddab35c4618f5b8bfc3116f36d1784a4a3275543a0"></a>

<a id="canonical-c21a044df0e45b68e67809dcdac4af28ed5e5d610a81b4d01667d1f21f8346c8"></a>

## maximum_age property — cors_policy / 41ea9f0f8843 / 11

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

<a id="canonical-73c83f3ff7bab77b241fb47d1d56b2a8fc0b35f0d921067ef57fb12d15a73d13"></a>

## Next pages — cors_policy / 41ea9f0f8843 / 12

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c9a870f8ec486129a3bfb92e157af90b7ed2f407741b40a6ff596b41269bd975"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-626c2389f1a5bdcec976e589e185933ee9b987879c289c099e7525bf08cc8335"></a>

## csrf_policy — csrf_policy / b203d8f11816 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- csrf_policy

<a id="canonical-d456ba9da9a96c880afa48699fa29b6f2b35697b479dc97167068450bb0cec32"></a>

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

<a id="canonical-a30b712a7ba6a0d5651c5e4c8027ad1fb4187f2fab196fe5c5fec186a487c137"></a>

## Direct properties — csrf_policy / b203d8f11816 / 3

- [all_load_balancer_domains](resources--http_loadbalancer--reference--group-014.md#canonical-5015920141e7cc362388a6c39e05bebeeefb05458f792bf58fa4854b2341b656): complete subsection reference.

- [custom_domain_list](resources--http_loadbalancer--reference--group-014.md#canonical-2ae675010a79e2194ce73f6da5693b1905a39bcb817bdbc2fdfb63c7ff154be2): complete subsection reference.

- [disabled](resources--http_loadbalancer--reference--group-014.md#canonical-b25619312b34585b5447618dd5f6f2e7e0d8e3e35c72fb33335afc10b0f511dd): complete subsection reference.

<a id="canonical-78f08621114e865f1d6dd04ccaa44d01618c0fdde9340867e72412f262098234"></a>

## Next pages — csrf_policy / b203d8f11816 / 4

- [csrf_policy.all_load_balancer_domains](resources--http_loadbalancer--reference--group-014.md#canonical-5015920141e7cc362388a6c39e05bebeeefb05458f792bf58fa4854b2341b656)
- [csrf_policy.custom_domain_list](resources--http_loadbalancer--reference--group-014.md#canonical-2ae675010a79e2194ce73f6da5693b1905a39bcb817bdbc2fdfb63c7ff154be2)
- [csrf_policy.disabled](resources--http_loadbalancer--reference--group-014.md#canonical-b25619312b34585b5447618dd5f6f2e7e0d8e3e35c72fb33335afc10b0f511dd)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5015920141e7cc362388a6c39e05bebeeefb05458f792bf58fa4854b2341b656"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9b5c6c7dab9b409c730c22859849b426498bcee0e9092e19de1f6106d70efcf"></a>

## csrf_policy.all_load_balancer_domains — csrf_policy.all_load_balancer_domains / 0682b22414bb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [csrf_policy](resources--http_loadbalancer--reference--group-014.md#canonical-c9a870f8ec486129a3bfb92e157af90b7ed2f407741b40a6ff596b41269bd975)
- csrf_policy.all_load_balancer_domains

<a id="canonical-1175ea836426803c252b3009b5d3fb14286b1536546b873a410823058348ca0b"></a>

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

<a id="canonical-47d64a6a290dedc64e0cfea20a3153aa5d5fb7a519f6eb11a3a88206e221b698"></a>

## Direct properties — csrf_policy.all_load_balancer_domains / 0682b22414bb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-87a7e2721df0bfaf609a3aae38c26f321e99c522c039d91ff9bea32547446c1d"></a>

## Next pages — csrf_policy.all_load_balancer_domains / 0682b22414bb / 4

- [csrf_policy](resources--http_loadbalancer--reference--group-014.md#canonical-c9a870f8ec486129a3bfb92e157af90b7ed2f407741b40a6ff596b41269bd975)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2ae675010a79e2194ce73f6da5693b1905a39bcb817bdbc2fdfb63c7ff154be2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf207ee5d3435ef43e049389f9ef62bfa1889f478bb42dce2c38c5dc6d679036"></a>

## csrf_policy.custom_domain_list — csrf_policy.custom_domain_list / 465c75eaf4d6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [csrf_policy](resources--http_loadbalancer--reference--group-014.md#canonical-c9a870f8ec486129a3bfb92e157af90b7ed2f407741b40a6ff596b41269bd975)
- csrf_policy.custom_domain_list

<a id="canonical-1ba7a1b49fa71c588c7c233eb9158bd51296ad2699492ce97b1bef30977de99c"></a>

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

<a id="canonical-9d6b298bc81a299fff82463ca62051277c7073e2172fb29ab8581a6fb1de13c4"></a>

## Direct properties — csrf_policy.custom_domain_list / 465c75eaf4d6 / 3

<a id="canonical-3f94e13083b356a132b296e55ae8ffffa6fdc639a366eb3861bfacbc94852351"></a>

<a id="canonical-0ee31ad0592f5d7a5e9a8d703303ae6de25d6148dbe7769d9d369718000fa4c2"></a>

## domains property — csrf_policy.custom_domain_list / 465c75eaf4d6 / 4

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

<a id="canonical-82db493ae2ad1779486922ac1358883ab3989fe97bd52c1445b00f04b0b6d174"></a>

## Next pages — csrf_policy.custom_domain_list / 465c75eaf4d6 / 5

- [csrf_policy](resources--http_loadbalancer--reference--group-014.md#canonical-c9a870f8ec486129a3bfb92e157af90b7ed2f407741b40a6ff596b41269bd975)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b25619312b34585b5447618dd5f6f2e7e0d8e3e35c72fb33335afc10b0f511dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ade5e0763fa05f02e845cc63f20f408b9850ea1617481d0545ec8c8009b25701"></a>

## csrf_policy.disabled — csrf_policy.disabled / 82d45f2df4fb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [csrf_policy](resources--http_loadbalancer--reference--group-014.md#canonical-c9a870f8ec486129a3bfb92e157af90b7ed2f407741b40a6ff596b41269bd975)
- csrf_policy.disabled

<a id="canonical-8bd2eced3e675b0f28d7ad6f8972cbe543a425f3edf0f65fd063d6dd87b70992"></a>

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

<a id="canonical-df65e71e08171a40c4e53c459f0027ecdc251071b753113bb0450121d87db4eb"></a>

## Direct properties — csrf_policy.disabled / 82d45f2df4fb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-802a3ff42644498912aff73a98571667b173883866dae1e1595c87ecfabf307c"></a>

## Next pages — csrf_policy.disabled / 82d45f2df4fb / 4

- [csrf_policy](resources--http_loadbalancer--reference--group-014.md#canonical-c9a870f8ec486129a3bfb92e157af90b7ed2f407741b40a6ff596b41269bd975)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3442b1c895a9f57fa1bf6b5e124c769e15a335cd5a6ebc494c98b2dc3d1f9ca9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-effc7ccec39b8f0d291f648471fc559336b845df40a79245920f890ab8d32656"></a>

## data_guard_rules — data_guard_rules / e3c52206ca62 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- data_guard_rules

<a id="canonical-8a1384ffceb744fa70e0f54ce367420b813626d34c3802f4a6ccaac8a588434e"></a>

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

<a id="canonical-d85dcffcdc20f03ce01efac41ba70c02ae92dd0de4a22fbab0a7e547ce517ae5"></a>

## Direct properties — data_guard_rules / e3c52206ca62 / 3

- [any_domain](resources--http_loadbalancer--reference--group-015.md#canonical-50ae7860c79718f1aea712e5280c8bdcbab4c8e8a539db6c1bc48fab2dad8cee): complete subsection reference.

- [apply_data_guard](resources--http_loadbalancer--reference--group-015.md#canonical-c149a320f10b642542c67ae2eed447271172a1db8390670cce07f76d3c098579): complete subsection reference.

<a id="canonical-4c5c2c4dfc73172b542cbc3cae9bdbebd036a5378a4fa0cd5f3895df206b5656"></a>

<a id="canonical-225525a1b35013d9aa22abbb2e863b3ab72ac683f30054f86127e7c848570c33"></a>

## exact_value property — data_guard_rules / e3c52206ca62 / 4

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

- [metadata](resources--http_loadbalancer--reference--group-015.md#canonical-3ba4ee498cc35a542aa9e35e5a911b167bcb66cd38c52c270253d78dd214db70): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-015.md#canonical-1bc9705c15a39fc67161a44b0a872b95e8354a97561f255ed9e8c5ebe77bc278): complete subsection reference.

- [skip_data_guard](resources--http_loadbalancer--reference--group-015.md#canonical-abac2b960dfad5a837e46c29dce72957d9278644732de05c7e08324c8fa67dfa): complete subsection reference.

<a id="canonical-662842cd480da3c66886557e46f93f821a6ff6c2e521f3c796a780e1bb1b70c3"></a>

<a id="canonical-eda9e99e0e1b5b5099d8a24bf9baebf31171041b8b0c0a5f30f37dae2a9f16fc"></a>

## suffix_value property — data_guard_rules / e3c52206ca62 / 5

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
