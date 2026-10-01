---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-6861943caeed0abc98a0ef28d4ca31feecb29d9b0d6e86861da9c7ebb3b6afcf"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / f35f62449e7f / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-0331c255f6ff70be96a95e3052c85fc779fcdb41020f0b314cf7dae2199050cc)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-8ac78519df25142076f609cb18bfb050cbed63e6c1253c69984ac7d7e932ab31)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-d4b199f18d955fd86b070279e4b5d070aa5d70fea944727fde8e175881b5e61a"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-f75ffe31f3ae598b2b9b9dd8be18d14a3f7fafae049999da8b702159e6d5e216"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / f35f62449e7f / 3

<a id="canonical-daedee21f2b82700145605c74a12ee3883e52dcb9d0cada80f2ea68f694092e1"></a>

<a id="canonical-34a0f13bfbec06783908d0aaeefe99b02490d397288bd369cb8df4b796d23a91"></a>

## addr property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / f35f62449e7f / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-efec92095e32ac03bafbf9e072a20acbe2817431e4cb0e9ceba571e975d496fb"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / f35f62449e7f / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-8ac78519df25142076f609cb18bfb050cbed63e6c1253c69984ac7d7e932ab31)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-83a0fe8b56ca6eb7d2af98d9cad14496727a329df97551812930bd5ceebee4b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08bf0cbc48d9b6ad984e7760c332abf5172795ff7d6ccd070dda70c614ff68d5"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.subnets — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets / c94497e93f3f / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-cffde52c11e4ec77a0ffea28fcebd409e94ba14e97284057a8f3134dee329ef1"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-a2f504b8641e92bb247f774aa584666cafdc4547dc48a1c921931807965900e9"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets / c94497e93f3f / 3

- [ipv4](resources--aws_tgw_site--reference--group-004.md#canonical-507a293d8c378584adfa6e98eb4bc47c3146194c36ac391c0fe7397f2b4c8367): complete subsection reference.

- [ipv6](resources--aws_tgw_site--reference--group-004.md#canonical-4a3d1d485f799a4811267cd7be401a45758892eb334e4decd2c5927ce2efe510): complete subsection reference.

<a id="canonical-e5572c4428cf81d0a1e8c5c22dec983e91c679ffceedc78d65a317383d4e0bd5"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets / c94497e93f3f / 4

- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_tgw_site--reference--group-004.md#canonical-507a293d8c378584adfa6e98eb4bc47c3146194c36ac391c0fe7397f2b4c8367)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_tgw_site--reference--group-004.md#canonical-4a3d1d485f799a4811267cd7be401a45758892eb334e4decd2c5927ce2efe510)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-507a293d8c378584adfa6e98eb4bc47c3146194c36ac391c0fe7397f2b4c8367"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e42c6a84d3444c7566365f26df09a010b79d2c5f8bb0225806300090c50a01fe"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 8b2834f68915 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-83a0fe8b56ca6eb7d2af98d9cad14496727a329df97551812930bd5ceebee4b9)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-3bb4c4b68dbb1b8428aa0f2e2dada765e72f4eab37dc7d0cb1796cc258cd67ca"></a>

Type: `"object"`. single nested block, Optional.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-0f9cff2f39ae440d384e5500c46e7933c71e0c5a9d8bf0c79c221fe520ab9709"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 8b2834f68915 / 3

<a id="canonical-6e12564e2e26cfce4ac2889ffbb118c8c5551e2121de7cc07fd3f3eb3443c9a9"></a>

<a id="canonical-e294dda8901bae5eae96850b228030ec82d358cd9db16f489366e8590a466e9b"></a>

## plen property — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 8b2834f68915 / 4

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-d32b7ca202631164c4c8c108461db5c41bc3e675b00d2498f89849a007d3645b"></a>

<a id="canonical-41f13e523db8399545dd09bb1dd6a10df589ab9e8d8d2f6975bb80603bfd6c56"></a>

## prefix property — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 8b2834f68915 / 5

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-892cd067d61b636b9917d4d8f1c6472db6a9a577166465b695ef20cdee1077e4"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 8b2834f68915 / 6

- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-83a0fe8b56ca6eb7d2af98d9cad14496727a329df97551812930bd5ceebee4b9)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-4a3d1d485f799a4811267cd7be401a45758892eb334e4decd2c5927ce2efe510"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f470cebf695bea29ac6a4436d78cec2ece075e5faadf5714d558ff6a606cc885"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / aca741d34260 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-b15c55c556e2c92237426dce5dfc798be5dff980940797925d11c1b591d9279f)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-06f20fe043f267b00f3e91a1fc52e5a407b0407d03f878e10796f9cd21bfdebb)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-d68ad92169f51204b53ba53fe3d3482eace4e869f717097f1d3680f6afd24d34)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-83a0fe8b56ca6eb7d2af98d9cad14496727a329df97551812930bd5ceebee4b9)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-4bbba6f398597b9013af3d3e77324a2b0eca83fcbc8cc97dff409822d66446ce"></a>

Type: `"object"`. single nested block, Optional.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-88c696e9395059bd2b159b8094c403badf80f75baa18d0452a1e2565c50240f7"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / aca741d34260 / 3

<a id="canonical-1e6a24e125335c225e163e83a846e4cc14219ddfa4090af2634d82e4306c9ebb"></a>

<a id="canonical-339f903f4752e2e39046c9f3d8a7807458c53d649efa8e8300491e19fe9249ce"></a>

## plen property — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / aca741d34260 / 4

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-d5d7b44bf5c7f514bc820ecf98b92987ffc639042b047e6795380ec138dfb1e9"></a>

<a id="canonical-5b1b1f17cab4084ccf6dc325e143305a920385c552f5bf96c60516a08da065ba"></a>

## prefix property — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / aca741d34260 / 5

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-c5363013aaea0d41350e9394223bbeb48e8f7771533f88c48041977672025ca4"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / aca741d34260 / 6

- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-83a0fe8b56ca6eb7d2af98d9cad14496727a329df97551812930bd5ceebee4b9)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-f320fa4d1906e5ee3d44cc33f3445eb2733ff209b91d622b368f1fdcc13dd93f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64061fe6b738897980e588f8f0ebed69a32d91c830a3b6cd5d9cf2e10281a42e"></a>

## vn_config.sm_connection_public_ip — vn_config.sm_connection_public_ip / d2111b65cd63 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- vn_config.sm_connection_public_ip

<a id="canonical-62934311464c78654206088afcad0e66574bc52752dfbd07e6237934a5b3b17a"></a>

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
sm_connection_public_ip = {}
```

<a id="canonical-eb32e3dabde19108d6f2dc992fec2ece3ca23a01bdca01290598a7f390579c16"></a>

## Direct properties — vn_config.sm_connection_public_ip / d2111b65cd63 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c0d777eeed88a06edf334ce99a87a237b93e00dcbd0f2ca8a829ea67cffba8b4"></a>

## Next pages — vn_config.sm_connection_public_ip / d2111b65cd63 / 4

- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-0919a404c2ee92ed821488904172bb0250b874ac4d2a9ff45cbc802eb327c0c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f115676999e7cdd500c9446be6fe3cb27c31fd41ddb37c42640b1c941daf08c0"></a>

## vn_config.sm_connection_pvt_ip — vn_config.sm_connection_pvt_ip / d097848cb8e6 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- vn_config.sm_connection_pvt_ip

<a id="canonical-c1a93302970bffa7d1c78b6891b41c0cfc2bc5fe9c1d75f09afb4763c25ade09"></a>

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
sm_connection_pvt_ip = {}
```

<a id="canonical-4f7b3f55135d525e92e4e938a2dca5493f4021adbbf14d527291c68b93e8f45c"></a>

## Direct properties — vn_config.sm_connection_pvt_ip / d097848cb8e6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cfa2678f3fb3b8e1a30866e9bec45d1023a22437b5bb0b9a9b55c315a09f21ac"></a>

## Next pages — vn_config.sm_connection_pvt_ip / d097848cb8e6 / 4

- [vn_config](resources--aws_tgw_site--reference--group-002.md#canonical-110845d859c3a4bab87ea9730548ddd8e2e2bde8e808ef623db5e3fb6a8f808d)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-6a136b15777bb8a98743650367f2764c51a7b782bf04d63d34a2defafee0ddf3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa353674f015fd8f492613bb3be2ea37dbc5c06b6524cdd6585094550a18b712"></a>

## vpc_attachments — vpc_attachments / b1e3280f5229 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- vpc_attachments

<a id="canonical-bce0d35c2b76f1d20d33b4f67b533b620e7758adb0eae939a5d0a9d989834d5e"></a>

Type: `"object"`. single nested block, Optional.

Spoke VPCs to be attached to the AWS TGW Site.

Receipt-pinned upstream constraints:

```json
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
vpc_attachments {
  # Configure direct properties listed below.
}
```

<a id="canonical-921233a555d4c8048379b6a1f62e66d44d72703e2f4c21fe8e4a2773905f1a35"></a>

## Direct properties — vpc_attachments / b1e3280f5229 / 3

- [vpc_list](resources--aws_tgw_site--reference--group-004.md#canonical-aa5c9688063c576b6645fa665fdc41678e4f6f7a42dae214b068cecd1701a572): complete subsection reference.

<a id="canonical-c3709079780d1c7254b61c1c4803bba96ba52c4532b1b3e1816e7de301bd8fcb"></a>

## Next pages — vpc_attachments / b1e3280f5229 / 4

- [vpc_attachments.vpc_list](resources--aws_tgw_site--reference--group-004.md#canonical-aa5c9688063c576b6645fa665fdc41678e4f6f7a42dae214b068cecd1701a572)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-aa5c9688063c576b6645fa665fdc41678e4f6f7a42dae214b068cecd1701a572"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c99c03c6711a1284f23b6be67865d3afcbdccab41c65cb1b67f99fe2075065b1"></a>

## vpc_attachments.vpc_list — vpc_attachments.vpc_list / 54e550a986cd / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vpc_attachments](resources--aws_tgw_site--reference--group-004.md#canonical-6a136b15777bb8a98743650367f2764c51a7b782bf04d63d34a2defafee0ddf3)
- vpc_attachments.vpc_list

<a id="canonical-64e51b24a01c68a625a6ce81995a1274e5be7c48822cf59ddfe5e2282aa55e68"></a>

Type: `"object"`. list nested block, Optional.

List of VPC attachments to transit gateway.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

Terraform syntax:

```terraform
vpc_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-c6d9bf666db55c83c45d4da14a3e232aed6eed7b6702a92d9898521a12b2dc74"></a>

## Direct properties — vpc_attachments.vpc_list / 54e550a986cd / 3

- [labels](resources--aws_tgw_site--reference--group-004.md#canonical-cc62237053094934b8be5c45d2cadf7be109bdb78ed9abb1c698ac45f7313c02): complete subsection reference.

<a id="canonical-e0043c45eb68468b70a72db18dd59949567ad5a7309ad2bbf28db29944ac2fab"></a>

<a id="canonical-ff56c7100429a192a31ef5907fca753cbfbf4ac423de3f1d6b2718904999e874"></a>

## vpc_id property — vpc_attachments.vpc_list / 54e550a986cd / 4

Type: `"string"`. Optional.

VPC ID. Information about existing VPC.

Upstream description:

Information about existing VPC.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-dd0ee3372bdc609cecca2e8d8acb1b47ab14090b18a45e617220b6596e5831c1"></a>

## Next pages — vpc_attachments.vpc_list / 54e550a986cd / 5

- [vpc_attachments.vpc_list.labels](resources--aws_tgw_site--reference--group-004.md#canonical-cc62237053094934b8be5c45d2cadf7be109bdb78ed9abb1c698ac45f7313c02)
- [vpc_attachments](resources--aws_tgw_site--reference--group-004.md#canonical-6a136b15777bb8a98743650367f2764c51a7b782bf04d63d34a2defafee0ddf3)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-cc62237053094934b8be5c45d2cadf7be109bdb78ed9abb1c698ac45f7313c02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f37687553672df0759d293a1ccdcf131df30f1589a7533083fda2bac0c12fb29"></a>

## vpc_attachments.vpc_list.labels — vpc_attachments.vpc_list.labels / e6a1a9d88b7a / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [vpc_attachments](resources--aws_tgw_site--reference--group-004.md#canonical-6a136b15777bb8a98743650367f2764c51a7b782bf04d63d34a2defafee0ddf3)
- [vpc_attachments.vpc_list](resources--aws_tgw_site--reference--group-004.md#canonical-aa5c9688063c576b6645fa665fdc41678e4f6f7a42dae214b068cecd1701a572)
- vpc_attachments.vpc_list.labels

<a id="canonical-e41605155bcea2cdd94df0e6f184d47db2e624175241c3f3941c25ee6f9f2822"></a>

Type: `"object"`. single nested block, Optional.

Add labels for the VPC attachment. These labels can then be used in policies such as enhanced
firewall.

Receipt-pinned upstream constraints:

```json
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
labels {}
```

<a id="canonical-f8b11bed6e4f865c83f2d60272220aa3f40344dd88614fef15317fa75d07cbe3"></a>

## Direct properties — vpc_attachments.vpc_list.labels / e6a1a9d88b7a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6b6e383aad7145ac354dd7db554f5cff4a57d50140e7f88879b6ff4fe95f4b24"></a>

## Next pages — vpc_attachments.vpc_list.labels / e6a1a9d88b7a / 4

- [vpc_attachments.vpc_list](resources--aws_tgw_site--reference--group-004.md#canonical-aa5c9688063c576b6645fa665fdc41678e4f6f7a42dae214b068cecd1701a572)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-6d8e9aff961991a95f631650b225f1a27401e6966c47ee24453e16093a40748e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98ba81e39922f3c0a1de8a9d851ec50a073d346a6139cd6b527c7e6fd410459b"></a>

## waf_signatures — waf_signatures / 09fb3a604ab6 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- waf_signatures

<a id="canonical-ea9e313c1fd8aba779020efd5a0f862c453f8cbffa600e7e2d723416b917abed"></a>

Type: `"object"`. single nested block, Optional.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("automatic",
    "manual")}
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
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

Terraform syntax:

```terraform
waf_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-3921834b6113fee48f0e7df613232c8912c36e183ae95257bc766707bdf314e1"></a>

## Direct properties — waf_signatures / 09fb3a604ab6 / 3

- [automatic](resources--aws_tgw_site--reference--group-004.md#canonical-ee757b984ccb8ca3265a3ffa60d1438c6e49d48e7f0d2182a28d53b5082508c5): complete subsection reference.

- [manual](resources--aws_tgw_site--reference--group-004.md#canonical-ddd668e33b873cacd10cbaef4190776594721043f2748a3fa1822e47489b761c): complete subsection reference.

<a id="canonical-a297277bb76b6c9fbff3eec271290ad11ccc3b113bdf15349f60b7c2ce45eb81"></a>

## Next pages — waf_signatures / 09fb3a604ab6 / 4

- [waf_signatures.automatic](resources--aws_tgw_site--reference--group-004.md#canonical-ee757b984ccb8ca3265a3ffa60d1438c6e49d48e7f0d2182a28d53b5082508c5)
- [waf_signatures.manual](resources--aws_tgw_site--reference--group-004.md#canonical-ddd668e33b873cacd10cbaef4190776594721043f2748a3fa1822e47489b761c)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-ee757b984ccb8ca3265a3ffa60d1438c6e49d48e7f0d2182a28d53b5082508c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05ef5ec5eac865437102283c14d0c2c7ac5dc387ed6cb7ab3fa32f3c9f0b7f29"></a>

## waf_signatures.automatic — waf_signatures.automatic / 49bd4c509cc8 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-6d8e9aff961991a95f631650b225f1a27401e6966c47ee24453e16093a40748e)
- waf_signatures.automatic

<a id="canonical-37f040218e8a7bd1a144d6c14bba7ab887b37a8791ed9e1ecd6fb8198cb87ae7"></a>

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
automatic = {}
```

<a id="canonical-03e807395c029ef289471b49433455f509669f6f466b6367f762fa4dbd8a5804"></a>

## Direct properties — waf_signatures.automatic / 49bd4c509cc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a2bc8a8d3e1f870ef4f2a3d66542ee7f928c729577d491f00fdc39909b30a4c"></a>

## Next pages — waf_signatures.automatic / 49bd4c509cc8 / 4

- [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-6d8e9aff961991a95f631650b225f1a27401e6966c47ee24453e16093a40748e)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)

<a id="canonical-ddd668e33b873cacd10cbaef4190776594721043f2748a3fa1822e47489b761c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1894ec6ff22e82e938f35230b9938bb623041479b2dd21e5fc944a91bcf28f4"></a>

## waf_signatures.manual — waf_signatures.manual / e64638fc9e93 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0430b1fa6acd6331f1ceff8713c35af0a2b119dd32e981384cafe9bed387a194)
- [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-6d8e9aff961991a95f631650b225f1a27401e6966c47ee24453e16093a40748e)
- waf_signatures.manual

<a id="canonical-06472057b4314a3c5613b1ee1254a6d14e3c15386a27d2f6f2ec773a252689a3"></a>

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
manual = {}
```

<a id="canonical-93406f9d253c3b9e61ca972ddec431c78f558390fe6cf44fd697bae17925394f"></a>

## Direct properties — waf_signatures.manual / e64638fc9e93 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-34e8058b241a594d2ab4b2f20e1a7d58261b2ffdbcefa4cc3e95bbd1678689ac"></a>

## Next pages — waf_signatures.manual / e64638fc9e93 / 4

- [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-6d8e9aff961991a95f631650b225f1a27401e6966c47ee24453e16093a40748e)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-ac580ecb6e0dcd05701c325c2680eec16da28fd82385686735e5bd9759502116)
