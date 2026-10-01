---
page_title: "xcsh_certificate_chain reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate_chain reference."
---

# xcsh_certificate_chain reference

<a id="canonical-c976c618810c3d923d3d6cdad9e66893de920464905dd953c308630b99c54079"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bef64c34a42a81c11cc5ed5b6454187197eca9362570e2d134fbf9d6fd4deb35"></a>

## Property reference — Property reference / cdb159f999bd / 2

Breadcrumbs:

- [xcsh_certificate_chain](../resources/certificate_chain.md#canonical-f0d6989e5267b6b55a7b8070bbe48d5745c060de8472924023939e563ef7d93a)
- Property reference

<a id="canonical-0db8bdface069ec19ab23c436814fecd369b04b2629e8d7145181a0eb93d3718"></a>

## Direct properties — Property reference / cdb159f999bd / 3

<a id="canonical-0c5af76f332776f30025e82b6c799cc8569e17a515ea7c1d5dfb644dc89848a7"></a>

<a id="canonical-3955a826fa3374438df65d2cda67346c2a7d1d2d0bfcd3be8d46cc45e6003dae"></a>

## annotations property — Property reference / cdb159f999bd / 4

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

<a id="canonical-c776d0542e8f4502e883d10301c056c7f0fb24860596a29449957ba1b82dd244"></a>

<a id="canonical-0be80c5aab24b47a8efcc0d6b81fa2df7704ec7f737747940d99dd0a1b557546"></a>

## certificate_url property — Property reference / cdb159f999bd / 5

Type: `"string"`. Required.

Certificate chain is the list of intermediate certificates in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.intermediate_certificate_chain_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.intermediate_certificate_chain_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-f9bc80976261cc073cdaf093b7bc728ed35b87ee42577600f3bc19e3ed438e65"></a>

<a id="canonical-78d447f6f3e8af1407d811582a5f62c57c2238748afca89eb9965a4259e6ebef"></a>

## description property — Property reference / cdb159f999bd / 6

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

<a id="canonical-4b9d0b2b9e54c57b070b4504a4276d482c34d0b3d7b6a2efb94e68d807b45649"></a>

<a id="canonical-eaa82a5f201fb3eaeee7360260cc35f19c7b97cf34e534e877305173d217f3e0"></a>

## disable property — Property reference / cdb159f999bd / 7

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

<a id="canonical-e0ea0a88d1ed8150c05415c05ac194a5a52e965cc1dacfc3f18d78dbfe8f63d2"></a>

<a id="canonical-a249e5d3ef7292d0259f6ae05ed80e070fc3a2d4e286d0d224fc207a642f12b9"></a>

## id property — Property reference / cdb159f999bd / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-72a828a6090a22eefaac3e6744e9362c35e1467c1337bbd92483aa9c06b0526f"></a>

<a id="canonical-6a6e0c58c693c27e3907ff92f858333a74a237adac739e5f1dc0e9b059276cbd"></a>

## labels property — Property reference / cdb159f999bd / 9

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

<a id="canonical-ff6eda4d84be20a1278496e2e3fc895a9a4868dad66a95db43290c01517500d6"></a>

<a id="canonical-b8423b0f6ed3e23d3ba89bd79b980f5fd87012d1986c56bb209f26c1c0f83c62"></a>

## name property — Property reference / cdb159f999bd / 10

Type: `"string"`. Required.

Name of the Certificate Chain. Must be unique within the namespace.

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

<a id="canonical-59c485ccae5907ccb9688d84ddd7b9971e07819d8525f5a9099859f74f99c1da"></a>

<a id="canonical-d35adbd69545b63ef9af7e4ba5eb04e54c0c070f349268c9592a3cd7aab9f695"></a>

## namespace property — Property reference / cdb159f999bd / 11

Type: `"string"`. Required.

Namespace where the Certificate Chain is created.

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

- [timeouts](resources--certificate_chain--reference--group-001.md#canonical-71660efa42bf25d132859a51b5cdd7aec87c3d4f14e82fde8d4e3e38ca85c529): complete subsection reference.

<a id="canonical-cc8d8b3f9900458e1fb55d16f192b4f48e663ec8f3af621e719034a62ba2e92a"></a>

## All schema paths — Property reference / cdb159f999bd / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--certificate_chain--reference--group-001.md#canonical-0c5af76f332776f30025e82b6c799cc8569e17a515ea7c1d5dfb644dc89848a7) |
| `certificate_url` | [certificate_url](resources--certificate_chain--reference--group-001.md#canonical-c776d0542e8f4502e883d10301c056c7f0fb24860596a29449957ba1b82dd244) |
| `description` | [description](resources--certificate_chain--reference--group-001.md#canonical-f9bc80976261cc073cdaf093b7bc728ed35b87ee42577600f3bc19e3ed438e65) |
| `disable` | [disable](resources--certificate_chain--reference--group-001.md#canonical-4b9d0b2b9e54c57b070b4504a4276d482c34d0b3d7b6a2efb94e68d807b45649) |
| `id` | [id](resources--certificate_chain--reference--group-001.md#canonical-e0ea0a88d1ed8150c05415c05ac194a5a52e965cc1dacfc3f18d78dbfe8f63d2) |
| `labels` | [labels](resources--certificate_chain--reference--group-001.md#canonical-72a828a6090a22eefaac3e6744e9362c35e1467c1337bbd92483aa9c06b0526f) |
| `name` | [name](resources--certificate_chain--reference--group-001.md#canonical-ff6eda4d84be20a1278496e2e3fc895a9a4868dad66a95db43290c01517500d6) |
| `namespace` | [namespace](resources--certificate_chain--reference--group-001.md#canonical-59c485ccae5907ccb9688d84ddd7b9971e07819d8525f5a9099859f74f99c1da) |
| `timeouts` | [timeouts](resources--certificate_chain--reference--group-001.md#canonical-c3d4c15d86ca67999a2a79996fe20812e4654c8a7c30238f8e37c1f4222d1c7b) |
| `timeouts.create` | [timeouts.create](resources--certificate_chain--reference--group-001.md#canonical-eb6768b4b1504d2ae908fd6c05656a4f64f3c92a5732f61b71c8b1ff78f9e362) |
| `timeouts.delete` | [timeouts.delete](resources--certificate_chain--reference--group-001.md#canonical-af9fa0cdb2a94060006af1501a957dc6cf045b6059516241e26adfb911f49574) |
| `timeouts.read` | [timeouts.read](resources--certificate_chain--reference--group-001.md#canonical-be8be526a7dc168c70c49a5a2c99cd4410cc9bc3eaa83a70e174e1fbe252e248) |
| `timeouts.update` | [timeouts.update](resources--certificate_chain--reference--group-001.md#canonical-3074f7494c32e3e4449b3c6bba6292e540d859ebe9e9a83384830fd993090666) |

<a id="canonical-feb6ab1f96add3893b5b1c3fb89688dd1a4cc7a5d0ad24644a200c09987f7f8b"></a>

## Next pages — Property reference / cdb159f999bd / 13

- [timeouts](resources--certificate_chain--reference--group-001.md#canonical-71660efa42bf25d132859a51b5cdd7aec87c3d4f14e82fde8d4e3e38ca85c529)
- [xcsh_certificate_chain](../resources/certificate_chain.md#canonical-f0d6989e5267b6b55a7b8070bbe48d5745c060de8472924023939e563ef7d93a)

<a id="canonical-71660efa42bf25d132859a51b5cdd7aec87c3d4f14e82fde8d4e3e38ca85c529"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e295d221e644cddb5becdfc948bfc59e371fa530670b2cf6d991ab821d9df40"></a>

## timeouts — timeouts / 047b48ea1f68 / 2

Breadcrumbs:

- [xcsh_certificate_chain](../resources/certificate_chain.md#canonical-f0d6989e5267b6b55a7b8070bbe48d5745c060de8472924023939e563ef7d93a)
- [Property reference](resources--certificate_chain--reference--group-001.md#canonical-c976c618810c3d923d3d6cdad9e66893de920464905dd953c308630b99c54079)
- timeouts

<a id="canonical-c3d4c15d86ca67999a2a79996fe20812e4654c8a7c30238f8e37c1f4222d1c7b"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-502770968958311438c801810d6ecbf961ddfbcba0f8d77514ab889936c03535"></a>

## Direct properties — timeouts / 047b48ea1f68 / 3

<a id="canonical-eb6768b4b1504d2ae908fd6c05656a4f64f3c92a5732f61b71c8b1ff78f9e362"></a>

<a id="canonical-c0aa8907e8f04523bdd4fc1ceb2d2bbc8fe8f5d41b89cc0c44bd9efdb57690fc"></a>

## create property — timeouts / 047b48ea1f68 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-af9fa0cdb2a94060006af1501a957dc6cf045b6059516241e26adfb911f49574"></a>

<a id="canonical-083eab2a8a5dd9ce1d6f0698e1718227ae733f5d880fd2b2b6d4fd56f236ec5b"></a>

## delete property — timeouts / 047b48ea1f68 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-be8be526a7dc168c70c49a5a2c99cd4410cc9bc3eaa83a70e174e1fbe252e248"></a>

<a id="canonical-0a82f292c4d5ef5120aeb7c80fcd9e2907fce12800ac25e66aa198fa336329bd"></a>

## read property — timeouts / 047b48ea1f68 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3074f7494c32e3e4449b3c6bba6292e540d859ebe9e9a83384830fd993090666"></a>

<a id="canonical-3f16653fca78601142f21aec672a29a4eeaa0743fbfa05b17e5bc98baa6fc6a8"></a>

## update property — timeouts / 047b48ea1f68 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-99baed36fd4aee17ed2fe3bf8207a791035b7ec39d0248fb86b48795d4732944"></a>

## Next pages — timeouts / 047b48ea1f68 / 8

- [Property reference](resources--certificate_chain--reference--group-001.md#canonical-c976c618810c3d923d3d6cdad9e66893de920464905dd953c308630b99c54079)
- [xcsh_certificate_chain](../resources/certificate_chain.md#canonical-f0d6989e5267b6b55a7b8070bbe48d5745c060de8472924023939e563ef7d93a)
