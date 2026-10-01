---
page_title: "xcsh_dns_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy reference."
---

# xcsh_dns_proxy reference

<a id="canonical-0b6367e63d7e9a3a0cfc8de4c8d47cbaa55db70ba4bddbcc158d0d3da9a5f7d2"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where / 6d5e952296d6 / 6

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-64cad82be0b9729c28ad2caec0ed571d5167a65c50865c5ae234f216cc2787b7)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-bc85bf57f08c03a2d77a6d61eff33c6119ccc05ccc69e954c7a9df611d2108b3)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-58ffa668c93929171fa32ce6e8cad94a6d7acc87c7c420fc64ae00b2e72e587b)
- [proxy_advertisement.advertise_custom.advertise_where.site](resources--dns_proxy--reference--group-002.md#canonical-3fa02dda6c2cf1ab054c037552de464ad85510fe3096f8358bc981e7e2306c8b)
- [proxy_advertisement.advertise_custom.advertise_where.use_default_port](resources--dns_proxy--reference--group-002.md#canonical-96c4edf8994da1e38c347e3702f02883bc62d3195ceb0a99c7865bf7821da37f)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-b448646677519bce2a1a9bfe3e0f5cb2c136e8074cae4328fccb54e8224e1efa)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-9913484b063cdc95a027af29b400730c1268977168fbc50afae97345195b30ae)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--dns_proxy--reference--group-002.md#canonical-d893ad774dbda1e1c44ddb1f685e238f0ab412a959b195d2a3c5862a5c12c3a4)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-fd22c2efb9959a92cb3da99b917df0844bd4bf08a80b9f4522b22aa495a9223f)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-64cad82be0b9729c28ad2caec0ed571d5167a65c50865c5ae234f216cc2787b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-897b84999f7fc554d2ed8d8292d1e2debc5d899eea7c4d4311054ed6f74a9a8f"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 25d6c7767da7 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-a7641a638b7e2ad7d0d958b3e8bf9f6ca75e9142b47a5d99a19bfb221cd67294"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_dualstack_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-ae6f6741d930c62635500c79c7307fb3430f1b80bd6c389c4a2396bd8f07601b"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 25d6c7767da7 / 3

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-1703c22f7a8cae07908d21011cbd9be16233699e214ddea10e37eec0d506b163): complete subsection reference.

<a id="canonical-7c491633fab67113f442c163717078a222793c98625889b0496fc7f39e8e73e5"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 25d6c7767da7 / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-1703c22f7a8cae07908d21011cbd9be16233699e214ddea10e37eec0d506b163)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-1703c22f7a8cae07908d21011cbd9be16233699e214ddea10e37eec0d506b163"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17bf717d2dca5e43b1b1a704cdaf179d3ebe0012eeb9b6589ebc4f9113de9480"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 94461d799e90 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-64cad82be0b9729c28ad2caec0ed571d5167a65c50865c5ae234f216cc2787b7)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-11084f7e0f4cdf6e7666142a08909d716fe03fc1f1156d2de318f9e4076fcc39"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-03831e9325caed6fe85504be9357b53628ad60f1537301b72b2e41a6b58cff2d"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 94461d799e90 / 3

<a id="canonical-11094d4841274368b88bc4cc91415afab2376a73c5b17625a57080e151b0c053"></a>

<a id="canonical-67cc2663bdf31bbc9f08aea799ecbddae1a47fbd7a02df1d0a9d0cf9328f67d8"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 94461d799e90 / 4

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

<a id="canonical-4cefeeb8898061a3fe573fdb24c517c281b42c0a47d7b2b3b390910b62effc9a"></a>

<a id="canonical-742c8dec87b8f64f0220aa30f7356b669d71973c3458ae55cb07501b04120403"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 94461d799e90 / 5

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

<a id="canonical-b25816f2199967924d641dc202e6b39fcd653042fcc1cad92cc0399b18f9cebe"></a>

<a id="canonical-61cad18409d28f649b356ff25558b6855bd9648390684df2f4e4b818ea56ae88"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 94461d799e90 / 6

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

<a id="canonical-57bfd3202284ef4a296fcbe1dafa1b0660f39203034fbac9fc93d4de1dfe0794"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 94461d799e90 / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-64cad82be0b9729c28ad2caec0ed571d5167a65c50865c5ae234f216cc2787b7)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-bc85bf57f08c03a2d77a6d61eff33c6119ccc05ccc69e954c7a9df611d2108b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4916b588716d5747d39a5617cc93c0d016ea9fd61d0f834616776984efe78510"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_on_public — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public / 9c48f21908d2 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public

<a id="canonical-b1c364f236d2079d77104c457b2d7c8854d76cdea96a02b48e233943dabd888d"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-7c8f3c2358a975df32c9fa99dd06ee4b5e27b66bfd622ff51bd57cc4b3895cfe"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public / 9c48f21908d2 / 3

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-d7b57da53805912a818c5bcc3f608ceff2445233061bd42daad3b252afddae6b): complete subsection reference.

<a id="canonical-a32c56c68b404e2d9337c8c1d60bdd2ec11e0e5971d841afc787b297472188a1"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public / 9c48f21908d2 / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-d7b57da53805912a818c5bcc3f608ceff2445233061bd42daad3b252afddae6b)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-d7b57da53805912a818c5bcc3f608ceff2445233061bd42daad3b252afddae6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-334be8893668433f1c52ddca3f1f55dfd80db774823785ad38f5614eec5e5452"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 3e96decd092f / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-bc85bf57f08c03a2d77a6d61eff33c6119ccc05ccc69e954c7a9df611d2108b3)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-5572bc785157d47d707cc86fa6f795731f428b660bca7640329f72bae52788d6"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-e318dbb0ed4c1051c284b4dcdfdfa5a3a760ca82fcbf5a7dff8e065972180415"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 3e96decd092f / 3

<a id="canonical-ed4719e0ad2b20a22cb2fbbfe283d2db8c56adefd89028c5c1dfc0c1f70e2057"></a>

<a id="canonical-6f3968e348bef7ab24fb69254894074bcba84814377ede8893e3a634604a6163"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 3e96decd092f / 4

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

<a id="canonical-a28ac0787e0339d51e9ded30eda8f426e5a19a05831f55d3692bcc1a37be817f"></a>

<a id="canonical-8695ab09044288cc50ab2583d036f8ab36fc3776c7e333d67ecccea769f2bc87"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 3e96decd092f / 5

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

<a id="canonical-fb884ae872ac642f18b7ccd1eaba791fd22db37bc9f9604a64fbed6e475dd56d"></a>

<a id="canonical-5d9faeb6035c5a0c99f1871d70ac08e16a40bc092b97c53eb8463d0fc30372ee"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 3e96decd092f / 6

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

<a id="canonical-01554bfc377d5074710513423419897a6facaac8feaa7748f940220481edc2f5"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 3e96decd092f / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-bc85bf57f08c03a2d77a6d61eff33c6119ccc05ccc69e954c7a9df611d2108b3)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-58ffa668c93929171fa32ce6e8cad94a6d7acc87c7c420fc64ae00b2e72e587b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2d189641c92cfc3704acf43d3fcdf31c05e4b2629245a3f73e81cecd01d0578"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public / b4504fdbacc2 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-a6c78f7a1b02d8235566e7ce4e656d383c20351cd653474860d430d92deb1f09"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_v6_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-1e5fd6113f60936f107b7cfbcc3235a828c762ec5b7d0bc8beac5146bd58a3de"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public / b4504fdbacc2 / 3

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-e188f7510df75d77342c9f1f91c10b5c29e5906a45174f94cc3d5bb884bc69cd): complete subsection reference.

<a id="canonical-51f5c7bef63cbc681a36e92def5d50ec5512515a7bc38d53c10424409b62c69c"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public / b4504fdbacc2 / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-e188f7510df75d77342c9f1f91c10b5c29e5906a45174f94cc3d5bb884bc69cd)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-e188f7510df75d77342c9f1f91c10b5c29e5906a45174f94cc3d5bb884bc69cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db01cf8e36fe16d3f80992210e6b9a18de3876a87973b9e9d5f56f7f2ecd03ff"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / 5859844ed425 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-58ffa668c93929171fa32ce6e8cad94a6d7acc87c7c420fc64ae00b2e72e587b)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-a9872ef0ff295e13a5e00d9b527465190f208baea7e11e5a91565945bbd03180"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-b7a59a99e35f447057886f114b25030a00497203641798162d5af470fe4d962e"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / 5859844ed425 / 3

<a id="canonical-e112615451bf5885220006e59ff915d919a9aa587f8490b21f688eade504fbd0"></a>

<a id="canonical-e89f4b739ae17e03313b179d28ae0ffcc4713f9f649576da22d77520dfd1c437"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / 5859844ed425 / 4

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

<a id="canonical-d26825daf0c1d61a576bfb20b210415d5da7a544b4f210e6fe89f28379c5e40b"></a>

<a id="canonical-93212a3f9f3387c9408474d4e03a750f3ff0406bfc9061005dd47f741d568295"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / 5859844ed425 / 5

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

<a id="canonical-88c077842f33179ed99ab72540ef2cd9c2bcac8eb586354e8c134d034a37d5c3"></a>

<a id="canonical-484223424b2114a5692977083b1092a3f303e8bd8dad670b2e3f7693afc00957"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / 5859844ed425 / 6

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

<a id="canonical-2b21bc8b4b95501c737998cae87a7dc3243acee3b61e01da324102d6efd43bec"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / 5859844ed425 / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-58ffa668c93929171fa32ce6e8cad94a6d7acc87c7c420fc64ae00b2e72e587b)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-3fa02dda6c2cf1ab054c037552de464ad85510fe3096f8358bc981e7e2306c8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fddfa23a731deb3872b949a0d6420ee102c83f1c3c912a6b9d0d7287a55d86c"></a>

## proxy_advertisement.advertise_custom.advertise_where.site — proxy_advertisement.advertise_custom.advertise_where.site / a88da27f7a38 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- proxy_advertisement.advertise_custom.advertise_where.site

<a id="canonical-b7c58531fe2ef386616ed925862a3713e7d6b6d126aa278f80f03e53a9c8328e"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a CE site along with network type and an optional IP address where a load
balancer could be advertised.

Upstream description:

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3bce9a28be49401d996cddb592922727b2a6bacaee5c18d317e1e14a7e489884"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.site / a88da27f7a38 / 3

<a id="canonical-e15601da128dfc99ceb371f45bc9d21d0c9514d31e9ca783e88b2afec468909a"></a>

<a id="canonical-8d544444552c3cd99924b970f82c7b6d0fe356c730f2cfcd27635246913f78fe"></a>

## ip property — proxy_advertisement.advertise_custom.advertise_where.site / a88da27f7a38 / 4

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

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

<a id="canonical-5135e3c354d4932d78b6b94782a17ce4e586c0f871e0e194c1d1ef310f636f85"></a>

<a id="canonical-fdda124f13a18df217f2a1e78eefcfa5a9ca55b54c0400a25b5bf39e9c897603"></a>

## network property — proxy_advertisement.advertise_custom.advertise_where.site / a88da27f7a38 / 5

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](resources--dns_proxy--reference--group-002.md#canonical-5930e7e923213b21b50ba0d3a3b8a9505ffff1eba8131a547be062291bce87de): complete subsection reference.

<a id="canonical-572634aa9d1b007629bfb8a2e9290f50d5be515c98d013f30eb41280cc922834"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.site / a88da27f7a38 / 6

- [proxy_advertisement.advertise_custom.advertise_where.site.site](resources--dns_proxy--reference--group-002.md#canonical-5930e7e923213b21b50ba0d3a3b8a9505ffff1eba8131a547be062291bce87de)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-5930e7e923213b21b50ba0d3a3b8a9505ffff1eba8131a547be062291bce87de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89deefbe202ce54abe9264045b10306c20055e1a9cad135b7d42e5d2a58764fc"></a>

## proxy_advertisement.advertise_custom.advertise_where.site.site — proxy_advertisement.advertise_custom.advertise_where.site.site / edf1049c5f71 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [proxy_advertisement.advertise_custom.advertise_where.site](resources--dns_proxy--reference--group-002.md#canonical-3fa02dda6c2cf1ab054c037552de464ad85510fe3096f8358bc981e7e2306c8b)
- proxy_advertisement.advertise_custom.advertise_where.site.site

<a id="canonical-d9189e4a4c191823dc0ada305436900ea3e670bf4565b7ce2e8dbc12af7bdbb4"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-4b9694cdc223ea3c91ac8ccd1a1aca73d852849288dc39912310982cf8a49d3b"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.site.site / edf1049c5f71 / 3

<a id="canonical-77aff12486af39edcb0a4a8fe18d1de8eeafbef6652213c45349efc8379a8b24"></a>

<a id="canonical-c76d42a15cbdfcb466e32ff553ad05b324cc0e76ab3d38c897e27c2e129ea28f"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.site.site / edf1049c5f71 / 4

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

<a id="canonical-6def463b6ea00e2b552ae2ac79e74a3eb48b43bcfad0404b01ee12edb65139ae"></a>

<a id="canonical-17f6443707a2cd41eb6c10fa2dce4418578d36a74d443ec567c140d7378a4e4f"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.site.site / edf1049c5f71 / 5

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

<a id="canonical-1003845a00905ac39fa1368262e207fa45a445abb3d9945a1294d3facc73bb5f"></a>

<a id="canonical-ac4495fe1bde6fbf2b396fb855a8942ddf0527907e6138a2d2ab43a5c52410e8"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.site.site / edf1049c5f71 / 6

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

<a id="canonical-eb29bd718a359166020aa0316306ea0485d64c18559f184426f84e57edff1fce"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.site.site / edf1049c5f71 / 7

- [proxy_advertisement.advertise_custom.advertise_where.site](resources--dns_proxy--reference--group-002.md#canonical-3fa02dda6c2cf1ab054c037552de464ad85510fe3096f8358bc981e7e2306c8b)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-96c4edf8994da1e38c347e3702f02883bc62d3195ceb0a99c7865bf7821da37f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e86178743bfa3cd7b68a9c0a216dcd30a558e8dc68f896584744a91fc7b9c255"></a>

## proxy_advertisement.advertise_custom.advertise_where.use_default_port — proxy_advertisement.advertise_custom.advertise_where.use_default_port / b20387638a6a / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- proxy_advertisement.advertise_custom.advertise_where.use_default_port

<a id="canonical-5f5f98d3d133309f4aa0b1a0462bd32e2599cb8fbe061121f8196b0ba451be71"></a>

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
use_default_port = {}
```

<a id="canonical-c4c05cc1ecb495e6013fba0831170b1316efebf4559af3775dbda4946a6522a9"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.use_default_port / b20387638a6a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b0c8024b9856f13919ce4f9cd5c7698906615a9a355f157b89649da52ddb934e"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.use_default_port / b20387638a6a / 4

- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-b448646677519bce2a1a9bfe3e0f5cb2c136e8074cae4328fccb54e8224e1efa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-366fe9059be55ccfa51d6e9a600e47b34d4df30056159625bc2b1e0ee1334e26"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network — proxy_advertisement.advertise_custom.advertise_where.virtual_network / fde7cdf473aa / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network

<a id="canonical-ca180aa443b2a8bbc323d0aaa827c63017c4471510953fe633e231bf59e1620b"></a>

Type: `"object"`. single nested block, Optional.

Parameters to advertise on a given virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_v6_vip",
    "specific_v6_vip"),
  validators.ConflictingObjectAttributes("default_vip",
    "specific_vip")}
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
  "x-ves-oneof-field-v6_vip_choice": "[\"default_v6_vip\",\"specific_v6_vip\"]",
  "x-ves-oneof-field-vip_choice": "[\"default_vip\",\"specific_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-941066cd110129095fe3042a4bab1e6501fa594db2bc5d00d02ad7cf193a7dd8"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_network / fde7cdf473aa / 3

- [default_v6_vip](resources--dns_proxy--reference--group-002.md#canonical-686a41e2d41155267a43a47f0fe1b7bcb1407ade37bf187781922674256e3ef0): complete subsection reference.

- [default_vip](resources--dns_proxy--reference--group-002.md#canonical-a3afdc6cfce1ca0445d6b83c2cff4a265bc4f6ba752a9a6eb2a446caf9b2faf3): complete subsection reference.

<a id="canonical-d49927a0d8bb4e63bacf2b11dcb1b43940a8cca2e170f231831259c5035957d4"></a>

<a id="canonical-0d879956a708061ace19c2af2b8d15f4030be6fe6a96703b0e142c0a605cfa54"></a>

## specific_v6_vip property — proxy_advertisement.advertise_custom.advertise_where.virtual_network / fde7cdf473aa / 4

Type: `"string"`. Optional.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

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

<a id="canonical-6c007d4d9dcdc87043df472f9954090212c2845eb86a5ebc6dd16e6781d20006"></a>

<a id="canonical-67d58847e60d29f133709de9693f2bc9c2bc73fde28b85f7435795c1f06d18ff"></a>

## specific_vip property — proxy_advertisement.advertise_custom.advertise_where.virtual_network / fde7cdf473aa / 5

Type: `"string"`. Optional.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

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

- [virtual_network](resources--dns_proxy--reference--group-002.md#canonical-5467480a7decd867142bb237804d84e083a45d76ed8be61576e35788b998d3b8): complete subsection reference.

<a id="canonical-ff89d8978eada4bd047c1ca4317c6e0b4521ffc1ac310236a19f6ce05b140c00"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_network / fde7cdf473aa / 6

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--dns_proxy--reference--group-002.md#canonical-686a41e2d41155267a43a47f0fe1b7bcb1407ade37bf187781922674256e3ef0)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](resources--dns_proxy--reference--group-002.md#canonical-a3afdc6cfce1ca0445d6b83c2cff4a265bc4f6ba752a9a6eb2a446caf9b2faf3)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-5467480a7decd867142bb237804d84e083a45d76ed8be61576e35788b998d3b8)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-686a41e2d41155267a43a47f0fe1b7bcb1407ade37bf187781922674256e3ef0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3df03796f5adb3e71b2b2f01748bd4fb71ef6cbc7c432805afde60d31741932"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_ / 53b9ca55018f / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-b448646677519bce2a1a9bfe3e0f5cb2c136e8074cae4328fccb54e8224e1efa)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-a3a080da6bf6bf1d679090a57dc285c8b8fc0d12dd821615a71c9e06b47f54c4"></a>

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
default_v6_vip = {}
```

<a id="canonical-0c3e0a0de4999d026bb0c663c901fdd31c1aabaadd705271ef622680bfa42c8c"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_ / 53b9ca55018f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-afa9060cf664d63d7898e78d4484c9386360fe670c63180ce186d37de8615767"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_ / 53b9ca55018f / 4

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-b448646677519bce2a1a9bfe3e0f5cb2c136e8074cae4328fccb54e8224e1efa)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-a3afdc6cfce1ca0445d6b83c2cff4a265bc4f6ba752a9a6eb2a446caf9b2faf3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fe9df8cb19bf093ccb8b94330b6d32a5101d92591c394c14ba00d5cca91a1c6"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip / ca3f48794129 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-b448646677519bce2a1a9bfe3e0f5cb2c136e8074cae4328fccb54e8224e1efa)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-c21dbdc282a723ea9812674157896af46b68f66583b7e6a984c90b3e884849cd"></a>

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
default_vip = {}
```

<a id="canonical-ae8dac7194251d992416fa15c0e262e0e421e1e0b72d1989e65de1487a6050af"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip / ca3f48794129 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0252d806f0c594526e25782465bf1a2ccbc3296b6578d9f59b7e600f6751444d"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip / ca3f48794129 / 4

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-b448646677519bce2a1a9bfe3e0f5cb2c136e8074cae4328fccb54e8224e1efa)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-5467480a7decd867142bb237804d84e083a45d76ed8be61576e35788b998d3b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5824d363e93fdb0a43396b6a5da78caf931f1c638698b2e0a16ce61749c0cfbe"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / ced7c163405d / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-b448646677519bce2a1a9bfe3e0f5cb2c136e8074cae4328fccb54e8224e1efa)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-2a9de88859f038ac8ded12e10682a96da16a07008d95111d990b620184ffbe6b"></a>

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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-458de8ec829af9cbd50f68a3dc3778dfe5b4c462b4d487bfbea1941f1ddb1317"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / ced7c163405d / 3

<a id="canonical-c031840ba96dcdcdafd7e0ee0234712af69cfb45c5405abc91416b6c6c643fa6"></a>

<a id="canonical-3d1e88bd2756f14209c590a6fa83c361d92da0e5ff7e6c49229fde5666c2da7d"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / ced7c163405d / 4

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

<a id="canonical-6787810ebb2262130927f15b05f04e2a7da9a9da547a0a7c73ccf677a1d24d67"></a>

<a id="canonical-b1d5f2e052c2a9443080e0ece102c42cb8958ef6b60ef39f04b68963cc6159dd"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / ced7c163405d / 5

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

<a id="canonical-a20fdeb4ab0de52ec011db347e14f3306bbb31c94c121a4eeb0511439d5d7de0"></a>

<a id="canonical-f233e9f1e6ef412afb1bb51928f2bf8713e2885a9e9795a45dafcea777ef506e"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / ced7c163405d / 6

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

<a id="canonical-7e39a66842be9105b71673df6d91f0513e8f9a76f14f39b4bd98a2249565c117"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / ced7c163405d / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-b448646677519bce2a1a9bfe3e0f5cb2c136e8074cae4328fccb54e8224e1efa)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-9913484b063cdc95a027af29b400730c1268977168fbc50afae97345195b30ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf081ef5984700f9e8a6e4a4f04de9e71aa79222b38d4d1deb38fef47f2e5221"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site — proxy_advertisement.advertise_custom.advertise_where.virtual_site / 546bbf845cba / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site

<a id="canonical-80f7220d364accf1332d905e392a4603f713feae2747d3af7e21c1399a2d856e"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a customer site virtual site along with network type where a load balancer
could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-b15a4f5dae4be98a137157cd0f350b5cb5b04a346b1690ece70d4a652c3e34d6"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_site / 546bbf845cba / 3

<a id="canonical-578d4a435d89ee492fd765f8d10c91ca82875b57fbb975f0e3edfce083fe9d58"></a>

<a id="canonical-4a447f830626db94c06c3be582a7515665b98287fe768c99e3536e78f7a97063"></a>

## network property — proxy_advertisement.advertise_custom.advertise_where.virtual_site / 546bbf845cba / 4

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](resources--dns_proxy--reference--group-002.md#canonical-8cdbc4604008aae295c1e68cfed201d03703c1798ba70f62ace966a3d5001b60): complete subsection reference.

<a id="canonical-2c738ffc401efe9aae78d37066a16065714df1f396a80c21952d976f1db4a8b4"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_site / 546bbf845cba / 5

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-8cdbc4604008aae295c1e68cfed201d03703c1798ba70f62ace966a3d5001b60)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-8cdbc4604008aae295c1e68cfed201d03703c1798ba70f62ace966a3d5001b60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90401cc744d77aa553d5b1a2cd2a20683b80154b1fa94ba3649ce7ca45c9f166"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / 55b2534b6ec9 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-9913484b063cdc95a027af29b400730c1268977168fbc50afae97345195b30ae)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-09402303b6b01b815aa84d443c46ff61a60b7de64ac603ae39938a7b20c0cc88"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-f06f90958b0b053270920f7077778e383a558aaa67b5a7539b2e7e32fbef2180"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / 55b2534b6ec9 / 3

<a id="canonical-c8d6931a620e726cdd38925b4b89a9e2166a787c60104a7e2566d9fad553a15f"></a>

<a id="canonical-f55d75b0d7301a19940fb5ad62e15614aa712080c09970518a7fd17396df1186"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / 55b2534b6ec9 / 4

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

<a id="canonical-07a0d737c704d127a93b0fe18632fad4b10f2bdaff8efcc67a918cba5e05e255"></a>

<a id="canonical-d908c5154a1424f4a9eb0f4a981de5669444775643cf3d6c74556b674be0d706"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / 55b2534b6ec9 / 5

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

<a id="canonical-138e49748931f123595157123f5ce1d62e5feb8d33ca4d94361f31bc328dfcdb"></a>

<a id="canonical-9d778a490f0c1860d9da3194916564f681092f453c04be6720644e1b06ee860b"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / 55b2534b6ec9 / 6

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

<a id="canonical-733aacc8ba6182ffd2722f18d4adbc6d9680c9baafbb912b10d9725acab32f95"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / 55b2534b6ec9 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-9913484b063cdc95a027af29b400730c1268977168fbc50afae97345195b30ae)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-d893ad774dbda1e1c44ddb1f685e238f0ab412a959b195d2a3c5862a5c12c3a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98b64b747231e8ab76ff074b5e6896db76ed10201962fbdf6f7f03b4f570829d"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / d6bad3c69863 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-ba0195b00537a658c98fdc01ca2466f9dea9e5b21bc6d130cd3a18263f9fbb21"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

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
virtual_site_with_vip {
  # Configure direct properties listed below.
}
```

<a id="canonical-cd0e09914290634d113dfeebdc37dfa2f11f9c38cbcb1cd27c574c9760aed259"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / d6bad3c69863 / 3

<a id="canonical-9e923eca821102d5f6f93f785d8fba93b67b822c1596c159acaae132417f7eed"></a>

<a id="canonical-091dd89ef09c2b0c548e81edc97119fc3e197be3d68c490fcab27fea077065f3"></a>

## ip property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / d6bad3c69863 / 4

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

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

<a id="canonical-a02d4095eff145223445c3826b66bb6ec60c16612df06fc05101f93efea6e4eb"></a>

<a id="canonical-e035d4a063be91c151fb62c325bd976ed3ab6833031ef84c6b1e774b7eccfa46"></a>

## network property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / d6bad3c69863 / 5

Type: `"string"`. Optional.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Upstream description:

This defines network types to be used on virtual-site with specified VIP

All outside networks. All inside networks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
  "enum": [
    "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](resources--dns_proxy--reference--group-002.md#canonical-360c5ac17323824df96c791a2469c2015b8af13ac47a9515d88bde5c6bd2ef3b): complete subsection reference.

<a id="canonical-07103b9da6f360e086a535f6e155b5cff792bc1a66f03b830f8ef22e90892629"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / d6bad3c69863 / 6

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-360c5ac17323824df96c791a2469c2015b8af13ac47a9515d88bde5c6bd2ef3b)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-360c5ac17323824df96c791a2469c2015b8af13ac47a9515d88bde5c6bd2ef3b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4536b115b93568e6d8ff46e68433b935fb7876605118e6e5c6a118750762aee9"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / d5eede481b70 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--dns_proxy--reference--group-002.md#canonical-d893ad774dbda1e1c44ddb1f685e238f0ab412a959b195d2a3c5862a5c12c3a4)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-df4d46bad32448ed50b9d8dc0ce0878b51c94b79f545c5a817574dee7ed90666"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-f8c9f84e2c0e70e9801ce7d3a77a7e5e55a40b447b83c9a20f62a3084a95b0dd"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / d5eede481b70 / 3

<a id="canonical-b86bd2ce52276a614a1c7ce0cc61c6e33c8d8f61524fd8d38304e37478881fee"></a>

<a id="canonical-d15e644d240a190b09837432742709bfe8d8e2a29b8680c36fcdb3ea7830d4ea"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / d5eede481b70 / 4

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

<a id="canonical-456a3a4c5b5e61cbfa61dadf9b6cd37d1a8d69ddf16d994cab7e6d0c98898d23"></a>

<a id="canonical-e9369ddba26150905cc41d1d3dbeb88a7c824918df1acb7969d40ccf9ee35469"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / d5eede481b70 / 5

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

<a id="canonical-497f3aeca106c9f756058fae79f15c55e2e340f0cbaffbb93483fd4806ce4eea"></a>

<a id="canonical-7c8f17417ff19a7c3acf1abe819603f96d0808561b0ef2cbfce003616ee00129"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / d5eede481b70 / 6

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

<a id="canonical-057ed77e0dddb29d97bcf38d8dfe62afd1b95f66f78c84f4cb6c2330641e8ae9"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / d5eede481b70 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--dns_proxy--reference--group-002.md#canonical-d893ad774dbda1e1c44ddb1f685e238f0ab412a959b195d2a3c5862a5c12c3a4)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-fd22c2efb9959a92cb3da99b917df0844bd4bf08a80b9f4522b22aa495a9223f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3da7d7702fd4927f4f0e4daf257122d357d454ef3179570d364c4f3ab57d1743"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service — proxy_advertisement.advertise_custom.advertise_where.vk8s_service / 8764d5008c2d / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service

<a id="canonical-1c58c9d224bcdceac97b955e16352b9dc4c9f6fa2294881dbca976f50eaa1df3"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
vk8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-27fdb596491684952c5791e7a3ec3a71b2720ffa89e7b5c5e5f6840f185a83fe"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.vk8s_service / 8764d5008c2d / 3

- [site](resources--dns_proxy--reference--group-002.md#canonical-f8fc8f7ea1112fb638cc7fd95e0b9621b026217cc41723dc277db04f88ce01ac): complete subsection reference.

- [virtual_site](resources--dns_proxy--reference--group-002.md#canonical-d3445ce38a4f4189c1df5967fae48e74287cc11f56fc05b5f1ca748b2da76871): complete subsection reference.

<a id="canonical-aac4ae1c7f76c151b8771cb5bf43cc4c9064ba20aa599d587b416df82a5ac1bd"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.vk8s_service / 8764d5008c2d / 4

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](resources--dns_proxy--reference--group-002.md#canonical-f8fc8f7ea1112fb638cc7fd95e0b9621b026217cc41723dc277db04f88ce01ac)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-d3445ce38a4f4189c1df5967fae48e74287cc11f56fc05b5f1ca748b2da76871)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-f8fc8f7ea1112fb638cc7fd95e0b9621b026217cc41723dc277db04f88ce01ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49ad639f32b12cf4d030ee3953145f9dc77309183bf0a7bf22dfcf2592985012"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / 20dd589e6bc4 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-fd22c2efb9959a92cb3da99b917df0844bd4bf08a80b9f4522b22aa495a9223f)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-816dc0ebd210ac2f405d272ec230ca9abf5e29b4afb0119c5dfc682bc51cd352"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-8ed577a0ac553521d88e851ba82d2cac91e70a17ed948461bfe897a537bcdba2"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / 20dd589e6bc4 / 3

<a id="canonical-01a5e2a26a8c19bbe58f87099d36a3fa66461477065911c64e39af3332ddf986"></a>

<a id="canonical-4095cb9c59fa0c535064942b10a7b825645498b1d65b7964665bc4b6c1b16226"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / 20dd589e6bc4 / 4

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

<a id="canonical-4c54e33544a462d7c433d431d47d8b87fe01fe6ae7e5acaf43d5dc6d3813b49b"></a>

<a id="canonical-aec168063d171a01609de3d2fd99ab1ed1611e15155407f49722dbaa7ce15232"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / 20dd589e6bc4 / 5

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

<a id="canonical-7c81efcd329b307a7328698d378d8eefe4948e50b0359fa297f72265b16adfc9"></a>

<a id="canonical-4dd105b25a897e1281f9d43b96824bd99377455f31b6c7d77ec0f6aa19862d65"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / 20dd589e6bc4 / 6

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

<a id="canonical-e8934244a2ee68fe504f397a47635ea912dcce71621726cd2429295080d93b6f"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / 20dd589e6bc4 / 7

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-fd22c2efb9959a92cb3da99b917df0844bd4bf08a80b9f4522b22aa495a9223f)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-d3445ce38a4f4189c1df5967fae48e74287cc11f56fc05b5f1ca748b2da76871"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa0ebfb64af63db93a5edd0026a576601bf7c57d4f1c20c083de2a78619601dc"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 02c094a42815 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-ec8ca912df59a88bf2b5ef3ab1b37246ecc9f41b954a2a8a77f5a4f1eef02a91)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-2649d97a7a81981f3f12b33ef51981bfca3e3d05ed2fd7dcf620fdae2ae71fa8)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-fd22c2efb9959a92cb3da99b917df0844bd4bf08a80b9f4522b22aa495a9223f)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-282ac9db2580842dfffd6856c0507080a3b3614e351ccd2da6f590227cbf0496"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-cf105c3d9be8a6c1690785afdb2c0315eefd724e5eba7c59466bd32f1ed6ed3c"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 02c094a42815 / 3

<a id="canonical-085baa29e23b35a322523acf8e20d810188aed1e5a51cce28d9a3150b05a8289"></a>

<a id="canonical-86cb4cf2d014dc1ced8b26027a7ffe2a6946311464d906b0b23fa3afe7e0b7dd"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 02c094a42815 / 4

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

<a id="canonical-2e1550638810dc5e6659da8ab7de9e6c3d57d4e3c38dddfc8f99e62ec9aef5cc"></a>

<a id="canonical-dff5c96f2c8d3426661a85e86348cd070f93dc410336a43ad22554094c93993b"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 02c094a42815 / 5

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

<a id="canonical-03f76c87d493070fb330c35464806b73c7513eeb4321fea25b45af6ee7454fdb"></a>

<a id="canonical-582619906f79af55ac1f800acf8112bb41daf7bdde11b35090ad21787ea6158b"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 02c094a42815 / 6

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

<a id="canonical-3fd80405ecca61c104e6c84231082eab9d5adc11b6f0b5a9e0d8b3846b97c205"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 02c094a42815 / 7

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-fd22c2efb9959a92cb3da99b917df0844bd4bf08a80b9f4522b22aa495a9223f)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-36e413c3804a479487c72a00782daad55ec340e06316387ea5579e5f34a9cdb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fdb227ca0498717e8152a120582482ea80a8a161668380570f58013b342f03ce"></a>

## proxy_advertisement.advertise_dualstack_on_public — proxy_advertisement.advertise_dualstack_on_public / 68a22f864e1a / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- proxy_advertisement.advertise_dualstack_on_public

<a id="canonical-c820766fcdd4fb770c94ca5c5b0fd54f15cafcf06975ec25ae613b56da319ba7"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_dualstack_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-a77acbc5c2c2e09949a07e5ceda86957cafa2e4ab215632c46cc5195ad67e7cb"></a>

## Direct properties — proxy_advertisement.advertise_dualstack_on_public / 68a22f864e1a / 3

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-ae0a6e5ecf36d39dce752c50aa846df247f8be5004f33b1110040c33ad8ba86b): complete subsection reference.

<a id="canonical-0a7b2f28633c162b3e05222ed627085fc1ad640093a28b692b3e7289ae1a51c0"></a>

## Next pages — proxy_advertisement.advertise_dualstack_on_public / 68a22f864e1a / 4

- [proxy_advertisement.advertise_dualstack_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-ae0a6e5ecf36d39dce752c50aa846df247f8be5004f33b1110040c33ad8ba86b)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-ae0a6e5ecf36d39dce752c50aa846df247f8be5004f33b1110040c33ad8ba86b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8611e5af2a31613c231fd9714a11b0497d6d772b5146b2876f388af3bf63da50"></a>

## proxy_advertisement.advertise_dualstack_on_public.public_ip — proxy_advertisement.advertise_dualstack_on_public.public_ip / a66f913396b4 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-36e413c3804a479487c72a00782daad55ec340e06316387ea5579e5f34a9cdb2)
- proxy_advertisement.advertise_dualstack_on_public.public_ip

<a id="canonical-a399c938bb8bc4d822898858f8dc4f3971560e4f446651ece4993faa82d7388a"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-ffeb122d05fd6ee4ba46738d31334a04b8c154726fe68bbbe67fd6145eb49076"></a>

## Direct properties — proxy_advertisement.advertise_dualstack_on_public.public_ip / a66f913396b4 / 3

<a id="canonical-a9dd5c6d15f11d94e5e2833f807625c04df13673f9e2d1092cd2aa176874bd80"></a>

<a id="canonical-c1b55e95bb2f6aec927d79538b5b2756bb70d2b430d025c9362e7420c359e615"></a>

## name property — proxy_advertisement.advertise_dualstack_on_public.public_ip / a66f913396b4 / 4

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

<a id="canonical-b047659fcd0e096390aea5439f39590cd885aa1e6c4abe0b578edc3b72c5795f"></a>

<a id="canonical-dbef9e7136eca56c0b48d5693bf6bc4314de917dc9ca8e5d624e106b5ebdfca9"></a>

## namespace property — proxy_advertisement.advertise_dualstack_on_public.public_ip / a66f913396b4 / 5

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

<a id="canonical-ca1694eaf86b6d6eaab67ef64be2eb57269c82a0f08177f558e092ff0033f687"></a>

<a id="canonical-0ccf228594ed5cef81c612284b502f6ad922edbf426b5482ab69fcb2b6e3eeed"></a>

## tenant property — proxy_advertisement.advertise_dualstack_on_public.public_ip / a66f913396b4 / 6

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

<a id="canonical-ec7820b5491d1a211d40630a6b8aff32b216d575c42e41e397beac1424917229"></a>

## Next pages — proxy_advertisement.advertise_dualstack_on_public.public_ip / a66f913396b4 / 7

- [proxy_advertisement.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-36e413c3804a479487c72a00782daad55ec340e06316387ea5579e5f34a9cdb2)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-a8cfb9d394509b726e0d82c9c369b0980923800bee81478d2a1193a893c5ab94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-925294dd0f47b9a79556eb2d9b7296533ee1e50fe95838ddec6a0b8319906380"></a>

## proxy_advertisement.advertise_on_public — proxy_advertisement.advertise_on_public / 3b217d4ff6ff / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- proxy_advertisement.advertise_on_public

<a id="canonical-54e8f45ac32b59b60a7ae822690d0b26ab76fa220d1a6bdadab1c0923a818d69"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-a89f9c283f3f187cd813a89fd2f72a075ac8d859515139e2a4b4a81b66901957"></a>

## Direct properties — proxy_advertisement.advertise_on_public / 3b217d4ff6ff / 3

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-6505eb32d6e1fdc9f49d2ae0efa31b982b1890148b4bb602e890f5f656edc4a1): complete subsection reference.

<a id="canonical-b99352d18495e602eebaecf3cac14fcfe02f4665c9a8aed3bc999c9a398a4b85"></a>

## Next pages — proxy_advertisement.advertise_on_public / 3b217d4ff6ff / 4

- [proxy_advertisement.advertise_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-6505eb32d6e1fdc9f49d2ae0efa31b982b1890148b4bb602e890f5f656edc4a1)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-6505eb32d6e1fdc9f49d2ae0efa31b982b1890148b4bb602e890f5f656edc4a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4109a198e92101b565957ad4f5b3592fc85809198f24be868b7810dadaa575c5"></a>

## proxy_advertisement.advertise_on_public.public_ip — proxy_advertisement.advertise_on_public.public_ip / 7761fe5226d3 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-a8cfb9d394509b726e0d82c9c369b0980923800bee81478d2a1193a893c5ab94)
- proxy_advertisement.advertise_on_public.public_ip

<a id="canonical-c124cd17917e7d5cfdee0b0c59b8f7c7e8c01c31f9df7d4edea0c3c37f85a77c"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-864a24613d0c81187193de045b058814079bae67ad246cf3d414a2f318779df1"></a>

## Direct properties — proxy_advertisement.advertise_on_public.public_ip / 7761fe5226d3 / 3

<a id="canonical-b78ba0d652c45de3b5e0e29ca94e4bb3cbb98c8deba1ab635b5f1640ac223e37"></a>

<a id="canonical-9df15582dcf1c5814b6a0519b9b0a3b4c0bde4df886bec47fe0a7dcce908f534"></a>

## name property — proxy_advertisement.advertise_on_public.public_ip / 7761fe5226d3 / 4

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

<a id="canonical-d0654a9912e1315803ede06839c8a8ac78561d358bf9bfabbcdcbf71e50563f2"></a>

<a id="canonical-4f938793df42959a276e22bf3686d20c2b1b84ebc90942fb63899372f8d903b5"></a>

## namespace property — proxy_advertisement.advertise_on_public.public_ip / 7761fe5226d3 / 5

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

<a id="canonical-52b1b257ecbfda48770df3db0ff3568b2c941a9ece75c1dbdf2c86f7d5e03b55"></a>

<a id="canonical-2a7284236cdc9ce7a834842e77c069635863d3c0eaa03dc8c9b3bbad6511c26d"></a>

## tenant property — proxy_advertisement.advertise_on_public.public_ip / 7761fe5226d3 / 6

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

<a id="canonical-c93b72ed4d15d43a0ff2002ab0614b874f71cf8801df3a3eb098044827c9d1a9"></a>

## Next pages — proxy_advertisement.advertise_on_public.public_ip / 7761fe5226d3 / 7

- [proxy_advertisement.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-a8cfb9d394509b726e0d82c9c369b0980923800bee81478d2a1193a893c5ab94)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-d9018890345e9c38927c60cd2abc03a1c13887da81cdf0e2d11570f102d379da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-663e42683ba4e9ffe009ff1ae3fd11948c60b5a76bcae6838931fb79b34ce3ae"></a>

## proxy_advertisement.advertise_on_public_default_dualstack_vip — proxy_advertisement.advertise_on_public_default_dualstack_vip / 7197410c4100 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- proxy_advertisement.advertise_on_public_default_dualstack_vip

<a id="canonical-84d2d13d5761b42c5c7e54727bc052b67717e3b2adbb8be52d66b2aa667f398a"></a>

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
advertise_on_public_default_dualstack_vip = {}
```

<a id="canonical-518faa1a378e6f1e2a1f8097f8180b6263cfb3dbfecfac8fb4a4e8bd4993c324"></a>

## Direct properties — proxy_advertisement.advertise_on_public_default_dualstack_vip / 7197410c4100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a83a5ecec1438ec9355a2490513e81bb9b4535b162a3122d465a8b41810331f"></a>

## Next pages — proxy_advertisement.advertise_on_public_default_dualstack_vip / 7197410c4100 / 4

- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-a8dabdd0094f62327d00befc565720566f64203ef5501c846fa47158a9d172f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9aaf59fb8c5489b5503ae407c75f3699ebbdf1f4ac0f4d451956fda08dbb1a9f"></a>

## proxy_advertisement.advertise_on_public_default_ipv6_vip — proxy_advertisement.advertise_on_public_default_ipv6_vip / 4aed38b9fae6 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- proxy_advertisement.advertise_on_public_default_ipv6_vip

<a id="canonical-7e48d8055fe6a2f1f4fd140f85356ddc41192d400af831537768ccc5e4db340f"></a>

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
advertise_on_public_default_ipv6_vip = {}
```

<a id="canonical-34cf72c094c8f60164aaaa7c9390510b2c8dce704d268d92c68a564df49ef245"></a>

## Direct properties — proxy_advertisement.advertise_on_public_default_ipv6_vip / 4aed38b9fae6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-efdbee77f69cbf3dedf1ccd5b0b49c8773f0b0f759f4cabaff7ff9aa24d68f29"></a>

## Next pages — proxy_advertisement.advertise_on_public_default_ipv6_vip / 4aed38b9fae6 / 4

- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-6a8eec7071e18f350fff1d62d082b941c2c8761c5e70f805cf71c559c9691d27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2604592fcd063abb085330918bb27f036617adedeb29cdfee46eb110faa2d07c"></a>

## proxy_advertisement.advertise_on_public_default_vip — proxy_advertisement.advertise_on_public_default_vip / 7b4f74d2075b / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- proxy_advertisement.advertise_on_public_default_vip

<a id="canonical-b229c26418bda71329600f5016b949b398b760327bbb18e88c8da04c9c067adb"></a>

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
advertise_on_public_default_vip = {}
```

<a id="canonical-a023f5a7995195188aa7617353c3aa07963573cbd564c224086d798147de48f4"></a>

## Direct properties — proxy_advertisement.advertise_on_public_default_vip / 7b4f74d2075b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-494525d4d608bbaa6bdcbb9c349ecd8bdc6167e4a5dbcdbcc8d040232d2d3426"></a>

## Next pages — proxy_advertisement.advertise_on_public_default_vip / 7b4f74d2075b / 4

- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-049a38713e827da66813f326a1263c600ce757ba01a9f6926640be9f17033e1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eddd3746234660b90d80d7dd056d5d113cc8ca6e06b900bfe21be5248e2f9459"></a>

## proxy_advertisement.advertise_v6_on_public — proxy_advertisement.advertise_v6_on_public / d9f49701df2b / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- proxy_advertisement.advertise_v6_on_public

<a id="canonical-526771c2166cf95233da0632290598a5d30732d85bd6cb7985c8fb5238dd3ff2"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_v6_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-64ce2f7dc030077db41d62adab25580a7da9e3ef29c4ef86dc037e5b856f757b"></a>

## Direct properties — proxy_advertisement.advertise_v6_on_public / d9f49701df2b / 3

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-e8ff2182ce91ef6f77c4029199659476992d2129aab22902aec5ac79ee675aeb): complete subsection reference.

<a id="canonical-73c867a4110be3b9894ed73d5264f5f9f57be5d1b6599fc4c95d07590e57f2ba"></a>

## Next pages — proxy_advertisement.advertise_v6_on_public / d9f49701df2b / 4

- [proxy_advertisement.advertise_v6_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-e8ff2182ce91ef6f77c4029199659476992d2129aab22902aec5ac79ee675aeb)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-e8ff2182ce91ef6f77c4029199659476992d2129aab22902aec5ac79ee675aeb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e2d102b12c76dfc6dc3ef5be1f983494bb1f718ba040bfcb963bd93f269e018"></a>

## proxy_advertisement.advertise_v6_on_public.public_ip — proxy_advertisement.advertise_v6_on_public.public_ip / 846480471369 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [proxy_advertisement.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-049a38713e827da66813f326a1263c600ce757ba01a9f6926640be9f17033e1b)
- proxy_advertisement.advertise_v6_on_public.public_ip

<a id="canonical-f024433cd07fd0cb19ee3eb5768761ab712f94233cb1360b5314239281394069"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-a25fcadfeeceb466a921cc2214a043019e61e8b4e86506741b6fcf51ab8e1735"></a>

## Direct properties — proxy_advertisement.advertise_v6_on_public.public_ip / 846480471369 / 3

<a id="canonical-1ffcb351004491851561252a220ee66520830cd6fe224ae85736de2055f7c0fc"></a>

<a id="canonical-a28bebfb84e34a2598c80ea5a29185956e24b921d1ba912a48c02e30934358cc"></a>

## name property — proxy_advertisement.advertise_v6_on_public.public_ip / 846480471369 / 4

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

<a id="canonical-e374814a597d550ec06cd74f075717af41293ed89621631a73d72eac3f689819"></a>

<a id="canonical-49525dc8560fb614d32a8fe225e4df34572edbb67b57518bb9e3f036f443245a"></a>

## namespace property — proxy_advertisement.advertise_v6_on_public.public_ip / 846480471369 / 5

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

<a id="canonical-d00b56add2a6f0fc5431966cda55582552dcda0f0e49a45ac143822d5b32d4ea"></a>

<a id="canonical-cacf6552c3bf3bed453ebfdd7b787bb24c507acb982f66baa6c246b037337f70"></a>

## tenant property — proxy_advertisement.advertise_v6_on_public.public_ip / 846480471369 / 6

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

<a id="canonical-31ad6353415cf6660c770ba7b3ef4276afcbbc20ec912d12bd6bc47d02ef4d7a"></a>

## Next pages — proxy_advertisement.advertise_v6_on_public.public_ip / 846480471369 / 7

- [proxy_advertisement.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-049a38713e827da66813f326a1263c600ce757ba01a9f6926640be9f17033e1b)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-3bb4e78ed32c4a8cda0a7dc8a0f096120494038738f8935d0887a0697ff759b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35095e136aec04eeb34696011d5c1184bfea35f2252bf39326b514f9ed4a6feb"></a>

## proxy_advertisement.do_not_advertise — proxy_advertisement.do_not_advertise / 29e186a1495a / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- proxy_advertisement.do_not_advertise

<a id="canonical-c5f205abf9af66657cec11a7dc597469311c4d34a70223a22586b71ab4272fe7"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise.

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
do_not_advertise = {}
```

<a id="canonical-668b0575061c6d4ad62a49def41d23ceb2cfba878f71e89cc918ebe887d7fd0b"></a>

## Direct properties — proxy_advertisement.do_not_advertise / 29e186a1495a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-75b49efc527344177520da1fa463353806cad3868e5b039df9c7fe265062717a"></a>

## Next pages — proxy_advertisement.do_not_advertise / 29e186a1495a / 4

- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-52507c9a61455520770b06df7e652c9d48cd01b1dabe747f532d8221b6b9bc6c)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-fbff3022aa3c1cd24e5dd973753c2686428fb896d6a2c568f5465e06f7e025a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94b92ee94a47ed37c84d5d0f1a035cea77ac7613636915cb5c87601d00dd3703"></a>

## timeouts — timeouts / 7501344e2b85 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- timeouts

<a id="canonical-997c29c6fbf12e74d56064161260a45cd928709c531cea946edcba3d30f6c84f"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0337220a868e6a0fa08755ebde2a6805d52065f21c3e2b613e20d836d1b5e512"></a>

## Direct properties — timeouts / 7501344e2b85 / 3

<a id="canonical-20885ebbf11b16653e443eeec8970c4b82e7092737a15d02b1ec79e2d425a301"></a>

<a id="canonical-74e7d6297c13d705bd3469badf3ec4bfc982383df89a4b86d19dfb4be0b69e1b"></a>

## create property — timeouts / 7501344e2b85 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-4a7058d79099b37eba4630afe530deacae2cc48bf3eb179d2c8553cac4d018e8"></a>

<a id="canonical-3fa11e81efceda8d6c4de463c8b2899878bab7a7ad5fa46177148be59eb002aa"></a>

## delete property — timeouts / 7501344e2b85 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-8578c36fe6e7f1710d75d1f34a97521a7629e3364861d0c4a3e5de6e4841c23e"></a>

<a id="canonical-4916affb9b3bac723e021d78de508ffcf466b385b10e93fb67be57b6878964cc"></a>

## read property — timeouts / 7501344e2b85 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-ed5cfbbef02441248562d4b562ab0be69b540b31e8754582b125f8cf2a23a3e4"></a>

<a id="canonical-2a50b8f57e6cac75bf50e199bc5335644752fc05d138e4b921e5200901fe6a4f"></a>

## update property — timeouts / 7501344e2b85 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-d151631ab403597b90637b73d12d67d3c0a6dbdaf8ca9a8bfbae4a7a817b1bc4"></a>

## Next pages — timeouts / 7501344e2b85 / 8

- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-f39ea3767ecfecd09544d5c662c3766d1f5bdf0102a3f65e7d6021feefd85da4)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
