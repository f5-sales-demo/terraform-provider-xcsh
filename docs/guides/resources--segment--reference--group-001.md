---
page_title: "xcsh_segment reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment reference."
---

# xcsh_segment reference

<a id="canonical-3fd2ddf918ccce02c13c05b12c882f76272305dd050891bce28f013df6b7f704"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b71e9b84c810d13f326db36023f1668d3505b5661498600f9edd47b2579faa6c"></a>

## Property reference — Property reference / 622b4b64fd9d / 2

Breadcrumbs:

- [xcsh_segment](../resources/segment.md#canonical-115b348cc2da26907c0d3a2200fee1c45353481468d739f4370fd6bf8ca9a410)
- Property reference

<a id="canonical-afdd5be43479c6485803c2630623f92ba48905636efbdd7c0bdcd9634ea75ace"></a>

## Direct properties — Property reference / 622b4b64fd9d / 3

<a id="canonical-499e6828001f89d252df59131fe48d948b4d3ac58ed5ee5a804b430ad5d9709e"></a>

<a id="canonical-ecf0eef2fd39754d692d0fd451efb0d8b6a6263296f00f861823cdf62bcfa2ad"></a>

## annotations property — Property reference / 622b4b64fd9d / 4

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

<a id="canonical-35f54726f863f6ef17c7fcda372b3e3979e0dd62502c9bf4dfce9c017c69c779"></a>

<a id="canonical-94f7eb84cc9a028b436d3b14ae61863ae047390e8645fe220ab4f998c8e965f8"></a>

## description property — Property reference / 622b4b64fd9d / 5

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

- [disable_spec](resources--segment--reference--group-001.md#canonical-acc6baed6480f8db86a7e22a39929e914d3ee4d4d1c15c4f58d277d392d24381): complete subsection reference.

- [enable](resources--segment--reference--group-001.md#canonical-d0ae754493da4a3aaf8c215aeb2f217dbfe57f22baf2d5937f94d59b3ddd95fc): complete subsection reference.

<a id="canonical-1f756b89b493d921b3c57949904a642ede32f080cb64ae0efbab8526b00cde47"></a>

<a id="canonical-d02be23d45c8b4a9eba7e38979d91b8cb9863a5501187d57a91258c503ddd4f9"></a>

## id property — Property reference / 622b4b64fd9d / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-a590006080a32e9bb9d70de4bda1c70d81de6859583977e80b1f4c5bd3c8d918"></a>

<a id="canonical-cdc8229667bdddb38f8c5f66b963a9aa7be39aca73a3c72d14cceae879277d75"></a>

## labels property — Property reference / 622b4b64fd9d / 7

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

<a id="canonical-766f90c45f755f816c9b456d12aa01e09e443df8758b339b1e79d8d072ee77ab"></a>

<a id="canonical-091459d5f101ad56eee7afd9e1ff2b607a30413394eede3000d60966e52811fb"></a>

## name property — Property reference / 622b4b64fd9d / 8

Type: `"string"`. Required.

Name of the Segment. Must be unique within the namespace.

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

<a id="canonical-824777e49e098e8b1f9b3e1385bbba3c3e78646787d2f7beb41da24b317a32ce"></a>

<a id="canonical-6588a99728173570416d59075119f7277b9592585f4746ec18ccf28af782a0e3"></a>

## namespace property — Property reference / 622b4b64fd9d / 9

Type: `"string"`. Optional, Computed.

Namespace for the Segment. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [timeouts](resources--segment--reference--group-001.md#canonical-6c30ec15955bbfec9888fe8ea3b5a4d93a4ad9cb60790787de0df5f5164b218e): complete subsection reference.

<a id="canonical-2f348f02b46931304fd60713928e1abfe8e87ab2dfeac027d7a94984687c81b2"></a>

## All schema paths — Property reference / 622b4b64fd9d / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--segment--reference--group-001.md#canonical-499e6828001f89d252df59131fe48d948b4d3ac58ed5ee5a804b430ad5d9709e) |
| `description` | [description](resources--segment--reference--group-001.md#canonical-35f54726f863f6ef17c7fcda372b3e3979e0dd62502c9bf4dfce9c017c69c779) |
| `disable_spec` | [disable_spec](resources--segment--reference--group-001.md#canonical-368edc61c1fb4fedbcd1522f89c241e371bf13b73bbb743b43b3f8452702ed1c) |
| `enable` | [enable](resources--segment--reference--group-001.md#canonical-e8b5f278093f318ddeb3e0bffd6799cbfcb3845052c29c67d5479a4c70280a09) |
| `id` | [id](resources--segment--reference--group-001.md#canonical-1f756b89b493d921b3c57949904a642ede32f080cb64ae0efbab8526b00cde47) |
| `labels` | [labels](resources--segment--reference--group-001.md#canonical-a590006080a32e9bb9d70de4bda1c70d81de6859583977e80b1f4c5bd3c8d918) |
| `name` | [name](resources--segment--reference--group-001.md#canonical-766f90c45f755f816c9b456d12aa01e09e443df8758b339b1e79d8d072ee77ab) |
| `namespace` | [namespace](resources--segment--reference--group-001.md#canonical-824777e49e098e8b1f9b3e1385bbba3c3e78646787d2f7beb41da24b317a32ce) |
| `timeouts` | [timeouts](resources--segment--reference--group-001.md#canonical-b88534d80e4b4972986b658f2724ad9d0663d4641e86ae62cd57f44d44bf4074) |
| `timeouts.create` | [timeouts.create](resources--segment--reference--group-001.md#canonical-bde71bbc616c52b50f3f4c7ca51ebf712eacbefc84aeaa80aa52856d1b42a665) |
| `timeouts.delete` | [timeouts.delete](resources--segment--reference--group-001.md#canonical-d90ceea5a0ba995d927f5b70eed6154a2b1610744cbbf6eeaac99b731f39a34d) |
| `timeouts.read` | [timeouts.read](resources--segment--reference--group-001.md#canonical-940a4a68bf58e8eb8e5e3f916f49c9aca773caccd9425dc74f64df8db4dd13de) |
| `timeouts.update` | [timeouts.update](resources--segment--reference--group-001.md#canonical-eae0b5f6489ba8f2ba8b112650c2d50a1f3951537ca2894924ab2af4d031216a) |

<a id="canonical-aabad8d86cd035d507382cd0c0e1c1b372f55ec9f938aaffdd6c2e6ef7368165"></a>

## Next pages — Property reference / 622b4b64fd9d / 11

- [disable_spec](resources--segment--reference--group-001.md#canonical-acc6baed6480f8db86a7e22a39929e914d3ee4d4d1c15c4f58d277d392d24381)
- [enable](resources--segment--reference--group-001.md#canonical-d0ae754493da4a3aaf8c215aeb2f217dbfe57f22baf2d5937f94d59b3ddd95fc)
- [timeouts](resources--segment--reference--group-001.md#canonical-6c30ec15955bbfec9888fe8ea3b5a4d93a4ad9cb60790787de0df5f5164b218e)
- [xcsh_segment](../resources/segment.md#canonical-115b348cc2da26907c0d3a2200fee1c45353481468d739f4370fd6bf8ca9a410)

<a id="canonical-acc6baed6480f8db86a7e22a39929e914d3ee4d4d1c15c4f58d277d392d24381"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64fcd9bd85474f248c484648817578c376f426e2bc2e9bc0e1b726bff174de6b"></a>

## disable_spec — disable_spec / 37b07138e32f / 2

Breadcrumbs:

- [xcsh_segment](../resources/segment.md#canonical-115b348cc2da26907c0d3a2200fee1c45353481468d739f4370fd6bf8ca9a410)
- [Property reference](resources--segment--reference--group-001.md#canonical-3fd2ddf918ccce02c13c05b12c882f76272305dd050891bce28f013df6b7f704)
- disable_spec

<a id="canonical-368edc61c1fb4fedbcd1522f89c241e371bf13b73bbb743b43b3f8452702ed1c"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable, enable\] Enable this option

OneOf alternatives in this subsection:

- `disable`
- [enable](resources--segment--reference--group-001.md#canonical-e8b5f278093f318ddeb3e0bffd6799cbfcb3845052c29c67d5479a4c70280a09)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-1826edaa292270108d6fa0920e260aef6128f4ad5c937dcb14a6deabebfdc9f2"></a>

## Direct properties — disable_spec / 37b07138e32f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cd0ff327a5212d45c7d00f039394ad7dac8da432ca937d691efcff351bdd4fa7"></a>

## Next pages — disable_spec / 37b07138e32f / 4

- [Property reference](resources--segment--reference--group-001.md#canonical-3fd2ddf918ccce02c13c05b12c882f76272305dd050891bce28f013df6b7f704)
- [xcsh_segment](../resources/segment.md#canonical-115b348cc2da26907c0d3a2200fee1c45353481468d739f4370fd6bf8ca9a410)

<a id="canonical-d0ae754493da4a3aaf8c215aeb2f217dbfe57f22baf2d5937f94d59b3ddd95fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7960f0d24dde233bd92638ae917a5e943783bef673dff029459e9677b05d06fa"></a>

## enable — enable / b37ec1faf255 / 2

Breadcrumbs:

- [xcsh_segment](../resources/segment.md#canonical-115b348cc2da26907c0d3a2200fee1c45353481468d739f4370fd6bf8ca9a410)
- [Property reference](resources--segment--reference--group-001.md#canonical-3fd2ddf918ccce02c13c05b12c882f76272305dd050891bce28f013df6b7f704)
- enable

<a id="canonical-e8b5f278093f318ddeb3e0bffd6799cbfcb3845052c29c67d5479a4c70280a09"></a>

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
enable = {}
```

<a id="canonical-71c6406bdc82c525c9fd4da20f7bbe0be4abdfbce4646084465e750fa58e1ccc"></a>

## Direct properties — enable / b37ec1faf255 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d5191d829cdd5b0db39d0a7f12e9f69392829ee8f8ed24e9a28bea488caeb37a"></a>

## Next pages — enable / b37ec1faf255 / 4

- [Property reference](resources--segment--reference--group-001.md#canonical-3fd2ddf918ccce02c13c05b12c882f76272305dd050891bce28f013df6b7f704)
- [xcsh_segment](../resources/segment.md#canonical-115b348cc2da26907c0d3a2200fee1c45353481468d739f4370fd6bf8ca9a410)

<a id="canonical-6c30ec15955bbfec9888fe8ea3b5a4d93a4ad9cb60790787de0df5f5164b218e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a89a1d2309df173ed13bf3256158c49c322e394e568d931b2afef4dbc463e4f4"></a>

## timeouts — timeouts / f76c16dfc36d / 2

Breadcrumbs:

- [xcsh_segment](../resources/segment.md#canonical-115b348cc2da26907c0d3a2200fee1c45353481468d739f4370fd6bf8ca9a410)
- [Property reference](resources--segment--reference--group-001.md#canonical-3fd2ddf918ccce02c13c05b12c882f76272305dd050891bce28f013df6b7f704)
- timeouts

<a id="canonical-b88534d80e4b4972986b658f2724ad9d0663d4641e86ae62cd57f44d44bf4074"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-d053350392811e094236f842cbeb45a509c3a49fa9ee660ec7bace6cf728cb5e"></a>

## Direct properties — timeouts / f76c16dfc36d / 3

<a id="canonical-bde71bbc616c52b50f3f4c7ca51ebf712eacbefc84aeaa80aa52856d1b42a665"></a>

<a id="canonical-6158036e642f7c1c394cd7c88cc645f4a3dc43c39657f4851f006c9a44cdff69"></a>

## create property — timeouts / f76c16dfc36d / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-d90ceea5a0ba995d927f5b70eed6154a2b1610744cbbf6eeaac99b731f39a34d"></a>

<a id="canonical-ed3a08efd4ab4133fa6c9097a4bda6af7e98bece1aea212b1975e41478b2c03b"></a>

## delete property — timeouts / f76c16dfc36d / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-940a4a68bf58e8eb8e5e3f916f49c9aca773caccd9425dc74f64df8db4dd13de"></a>

<a id="canonical-e05a3ee1a3bd8bab23f0a674a320ac841aa32a47592d13f4d8ea9f21f92f920d"></a>

## read property — timeouts / f76c16dfc36d / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-eae0b5f6489ba8f2ba8b112650c2d50a1f3951537ca2894924ab2af4d031216a"></a>

<a id="canonical-4c93b61ca3a28db451341bc8478e9bb5730ba933ff6344b98678545fce72393f"></a>

## update property — timeouts / f76c16dfc36d / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-be6bb49ffdf61d3a6c52187c56ada7a002f95b08e9e8a126f9a44b96498f73b4"></a>

## Next pages — timeouts / f76c16dfc36d / 8

- [Property reference](resources--segment--reference--group-001.md#canonical-3fd2ddf918ccce02c13c05b12c882f76272305dd050891bce28f013df6b7f704)
- [xcsh_segment](../resources/segment.md#canonical-115b348cc2da26907c0d3a2200fee1c45353481468d739f4370fd6bf8ca9a410)
