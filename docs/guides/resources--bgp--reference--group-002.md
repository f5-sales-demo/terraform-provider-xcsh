---
page_title: "xcsh_bgp reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp reference."
---

# xcsh_bgp reference

<a id="canonical-289f771c0e17b800dbb86a0f73ea97cf9ca9a881cd895321c2588d6f2db09846"></a>

## Direct properties — where.virtual_site.ref / cc6a2a8af6e1 / 3

<a id="canonical-0e3ffbcc9dcade2a8373f5bc6391e871857e6d774b54154e48198a018a4cc02c"></a>

<a id="canonical-a4f5566a156a42ee1a512881b492073a73b4ce86f3f745d4f6ead77ea164f866"></a>

## kind property — where.virtual_site.ref / cc6a2a8af6e1 / 4

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

<a id="canonical-947b96f34277781a2083f1597c04a27694a4657ec146b18ee43065633b34f60c"></a>

<a id="canonical-8d5bbc7949aab7facbbee97631c03996ccdf867105de2b8311c056f7a3ae8d4b"></a>

## name property — where.virtual_site.ref / cc6a2a8af6e1 / 5

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

<a id="canonical-cec6ab3c637feb75beddc9f3e4482918771e9d9009063ef6faa6c26a95491dc4"></a>

<a id="canonical-1670259af7a669c19ed3313680b4bb8584b5316e6c3ec5433046abbd2bb8678d"></a>

## namespace property — where.virtual_site.ref / cc6a2a8af6e1 / 6

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

<a id="canonical-1f60238c9fdf7e06f6bd0e6d1285729baaafaa4ca34095d1b0a9bed58d0e0820"></a>

<a id="canonical-5bac232d72e9ae11d55c7f3c575a8b139727e8dae1a76318545ffbcb2d0d61d9"></a>

## tenant property — where.virtual_site.ref / cc6a2a8af6e1 / 7

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

<a id="canonical-d1419ee4c9cb18225a3b9779ea78d35b5d49b31d8f1ec39d49513caad3eac5fd"></a>

<a id="canonical-26d9ce1615868e0444a3bda5ef9d75d297be159f084a8e30a88de95423142cf9"></a>

## uid property — where.virtual_site.ref / cc6a2a8af6e1 / 8

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

<a id="canonical-b6629f34bf58b5c951d3a2350605e32cebc66dbb254b3de5b8f654f6ac1961d2"></a>

## Next pages — where.virtual_site.ref / cc6a2a8af6e1 / 9

- [where.virtual_site](resources--bgp--reference--group-001.md#canonical-2d2a08b78ff2cade32a52bc98a6ea5e6a1b7db40299352ff19864c67ad4749e0)
- [xcsh_bgp](../resources/bgp.md#canonical-f452c8cd3b423a60b688afe1a7a3387739d94692cebe36d5a5d7dcaba10144e0)
