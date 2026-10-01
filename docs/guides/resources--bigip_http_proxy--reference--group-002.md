---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-e3a5921f20f3f802d92d29f0fefc96d329152e61adf5b9e891cdd20ec09aeb0c"></a>

## origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network — origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network / c87829555486 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network

<a id="canonical-97e39fd170317efae818de0467542269121a32df3f193c9234381d6eece89205"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

<a id="canonical-f345485c2dbc76bb3274cf514890d276dde8ab68420e084301eaec405b321490"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network / c87829555486 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2d44e80e27418da745183af0d1edc2eb964713f8bfd7da6a3c8c8f16c1e010c3"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network / c87829555486 / 4

- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-b0d40842f196312bcb1da06740b1c3a6cdf45a24120792379c079ffe7876017f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ea8c5d01cac3911b8016748c92aab677eea4e7d274d460c8d1e041da9082b99"></a>

## origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator / 83a9353b1bed / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator

<a id="canonical-ab165180bad9d4fb6a3ecb23bc399c6f7536265d36b89f2a2a1bd71efbcdd811"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

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
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-c7fc80c1185c5f7c5e9ca4b17f52633a5d8192aa7475967134662ff059fda866"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator / 83a9353b1bed / 3

- [site](resources--bigip_http_proxy--reference--group-002.md#canonical-c7589053b800fe8bd985e544eed30e106b401b9adde42eac6717356a11ff22c5): complete subsection reference.

- [virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-78fff89c49a76e6771c81396ea364d1e118993c1ed29f3cc5480c8dbf910b6e7): complete subsection reference.

<a id="canonical-c4d176b77688eea26e8639ee6e9bb63a9acaf0b29954d10c9898e79768069807"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator / 83a9353b1bed / 4

- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site](resources--bigip_http_proxy--reference--group-002.md#canonical-c7589053b800fe8bd985e544eed30e106b401b9adde42eac6717356a11ff22c5)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-78fff89c49a76e6771c81396ea364d1e118993c1ed29f3cc5480c8dbf910b6e7)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-c7589053b800fe8bd985e544eed30e106b401b9adde42eac6717356a11ff22c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57b915f2f1bb08233f620c24ff7221c1323a83a54e1d0dc36d99f77af15540fc"></a>

## origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site / dbd58703d936 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-b0d40842f196312bcb1da06740b1c3a6cdf45a24120792379c079ffe7876017f)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site

<a id="canonical-a9d6bb5892dc4beb3374928e29301286c7cb69da7209dd4066d1b33aa302224b"></a>

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

<a id="canonical-cc49f53f95b2efa453f1384d41316c23aa0fcd4d20e26dc0ef61ddde2e7b7da1"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site / dbd58703d936 / 3

<a id="canonical-0b9ab319993f855548bb990cf286da4c05ec466c9059ca5c97cd1c34f0a32e83"></a>

<a id="canonical-a03de961069da25a9ae722bb5c5081ee95b67f7db8d5cf8752569d637ede0405"></a>

## name property — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site / dbd58703d936 / 4

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

<a id="canonical-115c8fe7abab8571863e76d23e85732b039124d53a725ca8b7d667fddf5b3964"></a>

<a id="canonical-f02c0cfabecc5235cd3e0cde5ad46fcbedc1057481b9d7bbf6a9362eaa094c8c"></a>

## namespace property — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site / dbd58703d936 / 5

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

<a id="canonical-f457abe20734855ac56accccb848bc26fda708514d5de6b27c51a25ccee68439"></a>

<a id="canonical-0165fa3d5bcae1b79f63a792b7e9afc9bb2f55595bb0622f07a845d5a136f1ea"></a>

## tenant property — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site / dbd58703d936 / 6

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

<a id="canonical-398e834e8c312dcce3282fb0f5f96e47993d809b9f711c4411647a2ed86be238"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site / dbd58703d936 / 7

- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-b0d40842f196312bcb1da06740b1c3a6cdf45a24120792379c079ffe7876017f)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-78fff89c49a76e6771c81396ea364d1e118993c1ed29f3cc5480c8dbf910b6e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fde2f90ccc2369fdc57e196a6eb53edcef7d86cfab90bbd5140c52d5a3e1a53c"></a>

## origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtua / f4eb6009ef44 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-b0d40842f196312bcb1da06740b1c3a6cdf45a24120792379c079ffe7876017f)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-fddeca3f7cf95a4b8e942237a70f40ff097fdb5065641a24aa224bbd7c5bdfbb"></a>

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

<a id="canonical-cec71633745954f890e9a8e2650451e7dfe203a7f00bfa3f4bf8f3501b331611"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtua / f4eb6009ef44 / 3

<a id="canonical-2b0ae0e8a47845973d254214c28ffda0ffab5297aee18c33af6bfad368f7b210"></a>

<a id="canonical-06af3986a9cf62cdd9f9f22d3ab4811c16f2330964c5ac14d2d70bff875663b4"></a>

## name property — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtua / f4eb6009ef44 / 4

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

<a id="canonical-daa92cc01f797073c0bbe205d6146642ae0fc7ccb99f2848796f908f95f0c376"></a>

<a id="canonical-dd488988d79f925c291f32d649370c4ed7b698726d7489056c141efdb171cd76"></a>

## namespace property — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtua / f4eb6009ef44 / 5

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

<a id="canonical-09b6cb57709790ac1201489204d788c7d5a0228cb29b2c5606af95fed6c5b605"></a>

<a id="canonical-c62d9a6845965070a1b99672bc8f8cf52e7bf5359ed776814f1b93cb87b820c6"></a>

## tenant property — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtua / f4eb6009ef44 / 6

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

<a id="canonical-46a843992beda0311d23bed71ba7e5fd5d4e1875299e0499f3239114bbf1cd74"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtua / f4eb6009ef44 / 7

- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-b0d40842f196312bcb1da06740b1c3a6cdf45a24120792379c079ffe7876017f)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-75979de1f9ffe73f55b908741e0d1d74b34db692a58bcf38186434943cf1f844"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a35cde8cbc2924e22503cb16b92c9b809ea489b0bb0e20eb4ff4eb7b6b02a619"></a>

## origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool — origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool / 595f2056b88f / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool

<a id="canonical-64346720c3aadf3f1f83329d6bad6497e49eb59a7d0bc1bbce567e4bca59e49d"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-20dfdc8293d0da841982104cd4f32b7be46744065d8e6178faedf3e76372c177"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool / 595f2056b88f / 3

- [no_snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-58c23a2f7baa8da53d3704ee4e20da091abd5cf6b09b4cbf55875dd5dadc01c0): complete subsection reference.

- [snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-0e9903dab5c938bb535744d72bd9c632c6233bae98f4288e2f92efb845c1b963): complete subsection reference.

<a id="canonical-1ae1ddc49845d0d35181a3a4e0bc41bf88eef8b4dd8c7165b0b30086022e6187"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool / 595f2056b88f / 4

- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-58c23a2f7baa8da53d3704ee4e20da091abd5cf6b09b4cbf55875dd5dadc01c0)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-0e9903dab5c938bb535744d72bd9c632c6233bae98f4288e2f92efb845c1b963)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-58c23a2f7baa8da53d3704ee4e20da091abd5cf6b09b4cbf55875dd5dadc01c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2eb99559fb41d7e7284aa8b9ca46e0517ddb07803b5875656451446431afcad"></a>

## origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool — origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_p / 5399f8a4cd4b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-75979de1f9ffe73f55b908741e0d1d74b34db692a58bcf38186434943cf1f844)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-9be5f87c5129410aa601c770dfd25c9d0c7f98b18ab2775e182e97acd2608035"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

<a id="canonical-9f59e16cfb2c011e3fc6294dd908265f85bf49b29582ac4af00995e07c7913c1"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_p / 5399f8a4cd4b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b47ca3ecf88bfcf6e27b888e3f4df6d79edb64cb41aafd8211d80d5051a42727"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_p / 5399f8a4cd4b / 4

- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-75979de1f9ffe73f55b908741e0d1d74b34db692a58bcf38186434943cf1f844)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-0e9903dab5c938bb535744d72bd9c632c6233bae98f4288e2f92efb845c1b963"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24593b30f0cb54ac20679234a99e06b8f280734e2d4562ef2471ff8648003520"></a>

## origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool — origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool / 83bfccfb1f05 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-75979de1f9ffe73f55b908741e0d1d74b34db692a58bcf38186434943cf1f844)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-e5c12022e34775fd3d0f4c22f30cd8a1831facc8afe318e55a1f726991a0c3db"></a>

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-ef811ecf9bd5ffc7fb8e5d1d84da3ecd1bbddaa08d9ddabd008ecb6ac8f2a91d"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool / 83bfccfb1f05 / 3

<a id="canonical-f2b4e68076cf3d87110ff70a8c17b04d8aca23b3bb6150393bfb295c6bf5cc09"></a>

<a id="canonical-86cadb4b4ea79d0585468f2f4bd2106d85860e95623e055f93272a5dac1e45bb"></a>

## prefixes property — origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool / 83bfccfb1f05 / 4

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

<a id="canonical-901142eeb72b01a993d9ef3bf3e78251c5241e75b458a09a753f5291a88617d5"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool / 83bfccfb1f05 / 5

- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-75979de1f9ffe73f55b908741e0d1d74b34db692a58bcf38186434943cf1f844)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-05cbeb1e8c1729eef1f772fe6daa3359311b9d5cbcc129787e11b754c80cf539"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8108c05a3e49838432bbda4572404fd65af9fb7b6973fc46adb0a6470b5b1826"></a>

## origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks — origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks / 066bece5c5f8 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks

<a id="canonical-257cd1a085a52e805633de70db18e2b6c5b5a9d48666633464b5398a93a74c8d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vk8s networks.

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
vk8s_networks = {}
```

<a id="canonical-bc657f909ceef1de9388a295ec51d6b23a3c5c7742723343424d3491c3509d03"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks / 066bece5c5f8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-08b13aea415cf5afdaf08d75678ce342439959ae2f71dfb0755ac67b25e99284"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks / 066bece5c5f8 / 4

- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-258a1a59c101f1165d5601f7da73ec3e489abf0775437cf8b659434b58f3c3ef)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2cb4f08797e04282b9b8d71e3311916b98828a9ff3cfe72d738505ff95fddc1"></a>

## origin_pools.pools.origin_servers.origin_servers.private_ip — origin_pools.pools.origin_servers.origin_servers.private_ip / 1003b9e3d203 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- origin_pools.pools.origin_servers.origin_servers.private_ip

<a id="canonical-b9ac5037b66a974ad7fc7d962c0fe41998e9c2a83efb44391f15f340aa107adf"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with private or public IP address and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "segment"),
  validators.ConflictingObjectAttributes("outside_network",
    "segment")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]",
  "x-ves-oneof-field-private_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
private_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-eac41c3e9097a596b6a5453f6b34c4b6b71eb3e054942e93b3f8222ee7c6aa2b"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.private_ip / 1003b9e3d203 / 3

- [inside_network](resources--bigip_http_proxy--reference--group-002.md#canonical-e2a28b682148e9ae1393375ae3f48b6e597494df8c15c24b7328f6f2f23e8df2): complete subsection reference.

<a id="canonical-4abed645f474da3a75ab754c019735bb512d5b52160e35355c762f248a8e91a0"></a>

<a id="canonical-ed74fac939cb794079d05d7b84d837e2137fdef8d21c8310bb4c7fd145a18cb5"></a>

## ip property — origin_pools.pools.origin_servers.origin_servers.private_ip / 1003b9e3d203 / 4

Type: `"string"`. Optional.

IP. Exclusive with \[\] Private IPv4 address.

Upstream description:

Exclusive with \[\] Private IPv4 address.

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

- [outside_network](resources--bigip_http_proxy--reference--group-002.md#canonical-7218d94176a8b2a438a400f9ff475e13499b746c8905a41f90984df1c62497b7): complete subsection reference.

- [segment](resources--bigip_http_proxy--reference--group-002.md#canonical-9bf6d9df65372bb728a0ce953bcf1f39aefa97d3e758fbfe97453b94d2d7e4cc): complete subsection reference.

- [site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-6c2dad8e53af50d6a795db1761aa0d962ad276145370ad1a9d5e99158ef0e1f7): complete subsection reference.

- [snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-06ec74e997f45c1b0a1aafd75102b0cfb8895a146e82223188623de109542185): complete subsection reference.

<a id="canonical-be9acf91630ea5dd0d3d546d331cb7fd1f1e14380f80f54cbe51c9f164846389"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.private_ip / 1003b9e3d203 / 5

- [origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network](resources--bigip_http_proxy--reference--group-002.md#canonical-e2a28b682148e9ae1393375ae3f48b6e597494df8c15c24b7328f6f2f23e8df2)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network](resources--bigip_http_proxy--reference--group-002.md#canonical-7218d94176a8b2a438a400f9ff475e13499b746c8905a41f90984df1c62497b7)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.segment](resources--bigip_http_proxy--reference--group-002.md#canonical-9bf6d9df65372bb728a0ce953bcf1f39aefa97d3e758fbfe97453b94d2d7e4cc)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-6c2dad8e53af50d6a795db1761aa0d962ad276145370ad1a9d5e99158ef0e1f7)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-06ec74e997f45c1b0a1aafd75102b0cfb8895a146e82223188623de109542185)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-e2a28b682148e9ae1393375ae3f48b6e597494df8c15c24b7328f6f2f23e8df2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7923be9dfd8f85781795bc8f881fa5cb1284d175a7cd5bf2178837a299cc1e8b"></a>

## origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network — origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network / 18939c80c9c7 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network

<a id="canonical-2f46a42df25a722dfa0edfdadc6165a5aced2898e732967c9d960747022d0111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

<a id="canonical-7066122b400f56ffa8da1a7b8f4e060c640885156dc1923b74979eeeeab1a7b5"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network / 18939c80c9c7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7055d4e06d878704d82c5cf6e1db57727176b0498a75d9401460f4eeebf4cd79"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network / 18939c80c9c7 / 4

- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-7218d94176a8b2a438a400f9ff475e13499b746c8905a41f90984df1c62497b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00da5cc6931eda827e89e2e9fcb87c99ae61bc3de033e9fe4de5c1fdc793230d"></a>

## origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network — origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network / a010d54b68a2 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network

<a id="canonical-92ad27a4a3c72d98a62b42c3b61fc652b66c6e89d3d9e7c065ad74d70c47fa88"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

<a id="canonical-e7d652338dc7773400e4255ba71dc4bf2a57fc002242e232395c40e98ecdea35"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network / a010d54b68a2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9127b3c814f4b217c9f3b9dc2da9f7b2e18be5d166c3a2ea763813a06df8465d"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network / a010d54b68a2 / 4

- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-9bf6d9df65372bb728a0ce953bcf1f39aefa97d3e758fbfe97453b94d2d7e4cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-891a02f125f6be25ba86b4ab93bf8cd6bc18881ba358514febf3b4ebd387a2ec"></a>

## origin_pools.pools.origin_servers.origin_servers.private_ip.segment — origin_pools.pools.origin_servers.origin_servers.private_ip.segment / 47965c470904 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- origin_pools.pools.origin_servers.origin_servers.private_ip.segment

<a id="canonical-bdda455da5206e5503684792d231c5f2f101d314497eba0d56468c8802270a11"></a>

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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-b3e0c7af9ed37dc9700d71107e79462a8e140e80b1a71a627552b93ed00a9f57"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.private_ip.segment / 47965c470904 / 3

<a id="canonical-182801bf521eb4f42f7cf6e4b9857fa706e17327292156ae16e3b1b59b3bbed7"></a>

<a id="canonical-33d860339523ded9054876130af430624b87b21c913905e85da5424895167670"></a>

## name property — origin_pools.pools.origin_servers.origin_servers.private_ip.segment / 47965c470904 / 4

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

<a id="canonical-db0eca5ff7f3c026ca902c95d1424f2a7b36789ce1a1231a4d27b6da73cff35c"></a>

<a id="canonical-5ce11561989e78a42d02e714c460760ef0bb186ba0aeb6848682911892054fbb"></a>

## namespace property — origin_pools.pools.origin_servers.origin_servers.private_ip.segment / 47965c470904 / 5

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

<a id="canonical-c073d1e92c0fe2ffb567fc8c04c8682fd27b2dcb2acc2da11e0e39b90576ddbc"></a>

<a id="canonical-74bff18047a6d422bf5b1e70a1dfa4d24a6658fda29e538c3f01bbf9e47c697e"></a>

## tenant property — origin_pools.pools.origin_servers.origin_servers.private_ip.segment / 47965c470904 / 6

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

<a id="canonical-647e04c73d099813445cd864e8529f98f18031e4bbb9fbb1ec8ef2d8eb930a44"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.private_ip.segment / 47965c470904 / 7

- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-6c2dad8e53af50d6a795db1761aa0d962ad276145370ad1a9d5e99158ef0e1f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a38fe243d86b913eebe56a6dd085d04bcc3b98db4122296766984063e051e514"></a>

## origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator / dc1330500d8a / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator

<a id="canonical-2fa618c80590a083fab2b84cdf4b3482b652141a44afdc1c5db3bf39469ee61c"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

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
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-b7a058c5ed7eb3a1c90f2485e4107b57d14118294553765ebd337b190c6e3b55"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator / dc1330500d8a / 3

- [site](resources--bigip_http_proxy--reference--group-002.md#canonical-b24b515af6e23853f3ad24ce37979c7b2c59fc50bc9d5db9a534a6e06f5715e7): complete subsection reference.

- [virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-e5a188c659e2a5ec34216b8d0676b54d365c3b155f238802d95f9ad386297fdb): complete subsection reference.

<a id="canonical-d26d5077ed90dd7dd11414d67463a771bc2b20c4499493bfa9f685754eb4873b"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator / dc1330500d8a / 4

- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site](resources--bigip_http_proxy--reference--group-002.md#canonical-b24b515af6e23853f3ad24ce37979c7b2c59fc50bc9d5db9a534a6e06f5715e7)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-e5a188c659e2a5ec34216b8d0676b54d365c3b155f238802d95f9ad386297fdb)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-b24b515af6e23853f3ad24ce37979c7b2c59fc50bc9d5db9a534a6e06f5715e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54a1c6638c3da0fd676080f1a030a88e08fa8b706c1677c7dcabb1bac2a4f73a"></a>

## origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site / cc1157c5c5f3 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-6c2dad8e53af50d6a795db1761aa0d962ad276145370ad1a9d5e99158ef0e1f7)
- origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site

<a id="canonical-6e66afcc6da1fc458207cd9ecc4437a9b4ff31ac3abde9b727aa43e6fad60b59"></a>

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

<a id="canonical-291f9ece1c9a1094ecbb48b57a1b92e5d367e50eeb65250325bc4350610205af"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site / cc1157c5c5f3 / 3

<a id="canonical-daf31fb7f8057b496ab5d2fc21daf0c3c46f8eae97ccf04a806b20d53935ff41"></a>

<a id="canonical-e46706868652916c34d7376901cfa0c6ff8123374e29a7a68f6b25edb62b5af4"></a>

## name property — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site / cc1157c5c5f3 / 4

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

<a id="canonical-e6ccba1a7698c1ac89c41b89c89308410a97f447a85c7fe6d6e0444ddda3ccaf"></a>

<a id="canonical-d018f3ccbf173a7214b0e59f8eabfa9bdca62ea2f59b899e5f34d312e04587c2"></a>

## namespace property — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site / cc1157c5c5f3 / 5

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

<a id="canonical-e1f0014e870f11d422a647b8718fc8a9620ccb24baed50b583193cd35c39c351"></a>

<a id="canonical-f45f302f91bcf5400abd791cd0b8784a99a8a5ba23c3a1f210a87617a0151e0d"></a>

## tenant property — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site / cc1157c5c5f3 / 6

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

<a id="canonical-a2edce02774c2cb6e4410c234e66ace2bd2722616cafb3734e9e24ac9f873f91"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site / cc1157c5c5f3 / 7

- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-6c2dad8e53af50d6a795db1761aa0d962ad276145370ad1a9d5e99158ef0e1f7)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-e5a188c659e2a5ec34216b8d0676b54d365c3b155f238802d95f9ad386297fdb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c068fc4e980e5982395bc9d185476ca5f5966cb1dcc7e3285fc85a994e696d1"></a>

## origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual / c2de8fb7c3f7 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-6c2dad8e53af50d6a795db1761aa0d962ad276145370ad1a9d5e99158ef0e1f7)
- origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site

<a id="canonical-7c65e5a6571b7045c585f2eb5c46d7c6207a38cabc134af10a05f5ac89921f5a"></a>

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

<a id="canonical-89d3b0ae734543d37570fe814e472324aa1fc6ab27fa86b5e140044e67387f4d"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual / c2de8fb7c3f7 / 3

<a id="canonical-6cfa1dc8bc409e5f7dcf443cfc4ed85cfabec60e7c490ccb4eecd3c7779d62da"></a>

<a id="canonical-61efd0729817cc217228a5410f3e85e8e0c9f849adadb84f30ae8505650d786d"></a>

## name property — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual / c2de8fb7c3f7 / 4

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

<a id="canonical-560b092e51cc06fbf452526a7ba134d21a68668fa7da6dc5659f6ff6a3bf12df"></a>

<a id="canonical-6a46314021845d8888ec4c42d4f8454859e7e4dd5f015c4d417fb1a5547de216"></a>

## namespace property — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual / c2de8fb7c3f7 / 5

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

<a id="canonical-0324e4d62ad11c626d38811dc4091490ff9f3feab130579c6e0b27a55a3b9092"></a>

<a id="canonical-46cb058c41e5cf9aca5eeaea368ff0b7dff8393c6c37b38babc21f0add1326b8"></a>

## tenant property — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual / c2de8fb7c3f7 / 6

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

<a id="canonical-bfb4278443d351a72ea11d5eb9f3e10dd26578ad96ddc7466c8c97f541c5cf89"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual / c2de8fb7c3f7 / 7

- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-6c2dad8e53af50d6a795db1761aa0d962ad276145370ad1a9d5e99158ef0e1f7)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-06ec74e997f45c1b0a1aafd75102b0cfb8895a146e82223188623de109542185"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3327086391d024fdaab3a1bbf7afa3f29898c5313c7be0a7a4ededcfc5dd21b6"></a>

## origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool — origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool / 01139f0783a7 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool

<a id="canonical-8e05525a92746aead25ca2a74f1ddef662dd274b6e69946f8f0a01dfbb8f7d82"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-39a065b99d196d0f9323908482d50748135009b66841ada28baec1c4bbe09778"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool / 01139f0783a7 / 3

- [no_snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-c42a551b0edcc05f764d05266a52810def4d6cad62caa512c8915e0f0b88cf17): complete subsection reference.

- [snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-7af175222dd651934bb42ae08b2ca5e7d7b216c4d925fd017f36f949f508d3ae): complete subsection reference.

<a id="canonical-21d3de9fcaf6e6b9f6cb397dd96d89c6889019dcad8db35dde330e491cea53a4"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool / 01139f0783a7 / 4

- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-c42a551b0edcc05f764d05266a52810def4d6cad62caa512c8915e0f0b88cf17)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-7af175222dd651934bb42ae08b2ca5e7d7b216c4d925fd017f36f949f508d3ae)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-c42a551b0edcc05f764d05266a52810def4d6cad62caa512c8915e0f0b88cf17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6cd267c9a0db4dd1e0ac97ead0a496f34f02c7ab61904b04ad2565a08f3735f"></a>

## origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool — origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_po / b6613ec30185 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-06ec74e997f45c1b0a1aafd75102b0cfb8895a146e82223188623de109542185)
- origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool

<a id="canonical-934bf957d3afc37edb659a6c8cf20de0fef1608f2d55b555a812ee997644bfcf"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

<a id="canonical-4ec90ab9e532d5fe968fc9a16beec5ebd5723cc48be59867d9a0327dae366b5d"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_po / b6613ec30185 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2ed9ce3953aed4f4718cd10db47b1fb3f27b96c3f856af7df7d76090bbbdfd9a"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_po / b6613ec30185 / 4

- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-06ec74e997f45c1b0a1aafd75102b0cfb8895a146e82223188623de109542185)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-7af175222dd651934bb42ae08b2ca5e7d7b216c4d925fd017f36f949f508d3ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3af81158f02ef78eddf114cc355e7740638a5f0fdd233de829025e35074fa5f9"></a>

## origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool — origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool / b5cdabde6f6c / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-ee77f5763b191d4d1e99ff7c0a2b2a8bf144d74ca3f15564221a8ec4ee4638f8)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-06ec74e997f45c1b0a1aafd75102b0cfb8895a146e82223188623de109542185)
- origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool

<a id="canonical-4ef35c03f7e343139ba72f3f8dcd7e8d2ec047d599aef0c29e4df287c9f50a05"></a>

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-809fb802ec7e4f7d4e29730255d66da7f79a1b097a1fec235fda623e2caee01e"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool / b5cdabde6f6c / 3

<a id="canonical-5b54d399f8eefd2929d76c5d9ac9c31574677f57814c8dd8ba50c2c273f2d7a8"></a>

<a id="canonical-da27518bf0d9dbf972bd1f18885da22210832ed0b704babda510f3df63465be5"></a>

## prefixes property — origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool / b5cdabde6f6c / 4

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

<a id="canonical-4101a6f579cbf81f0c0f11fe94208006a74c3c8acfcae815b515078bf2ce5175"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool / b5cdabde6f6c / 5

- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-06ec74e997f45c1b0a1aafd75102b0cfb8895a146e82223188623de109542185)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-a0fb68a40e29f1563916c0752435bd05b633fe257df629a801bd86b305b8aedf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f35baa64ceb9c338192e3b45cf206dc06578cee1c65adbd398a094b95d5fa1be"></a>

## origin_pools.pools.origin_servers.origin_servers.public_ip — origin_pools.pools.origin_servers.origin_servers.public_ip / 41b9ab5442bb / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- origin_pools.pools.origin_servers.origin_servers.public_ip

<a id="canonical-84acc5e32f941a19084cc6863a24cfc32aaee1694e1539008fff3cec4e7a812c"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-51d1b135acbe19f3daddeec104d9f63cce20c0b6dc91535a3f43e4d6778bc52c"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.public_ip / 41b9ab5442bb / 3

<a id="canonical-7a78c119ffbc8370c6319bdb971b5a231774c04c267c5c68b325b93a2e7ee966"></a>

<a id="canonical-a709bba82f80dd7142a214da4b23ac9aee9fb48f36fe33898211e8eef9c33263"></a>

## ip property — origin_pools.pools.origin_servers.origin_servers.public_ip / 41b9ab5442bb / 4

Type: `"string"`. Optional.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Upstream description:

Exclusive with \[\] Public IPv4 address.

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

<a id="canonical-e2886718c6b18b2ac74e40f6661bc454acf86ea586b1f48e924cfef010cd6a43"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.public_ip / 41b9ab5442bb / 5

- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-82f1b165508c822fa0aa39d03b6ca6cbb3c668bc683b9852c23ec6c73c6ef894"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9957e040e89f467cfb62958343b50263248a38f5f2e973d6f2613c270d97c92e"></a>

## origin_pools.pools.origin_servers.origin_servers.public_name — origin_pools.pools.origin_servers.origin_servers.public_name / 57a81bd1eb14 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-dcc909ce5f3e7042bd791e8723f5abfdc95cff2cb76bf07e334dfb94827884e2)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-4b35d46b3f08010a85262ec2990d11b81ddc25dc322eb522a3c2f537f80e3b8a)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-76aa74d039c147ddf634f1fe44cf7229f51b32e64421201edc9f14cf1a13d4ca)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- origin_pools.pools.origin_servers.origin_servers.public_name

<a id="canonical-f2c5a0876cdfafe0e2e9580568fe6ec1d1078b40d96520fa7346f612ba09b30f"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
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
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-80099234d0c3c6c554ad184dc3e0696bf98275a4caf7cb1ab6c2b22d4dd1dbc1"></a>

## Direct properties — origin_pools.pools.origin_servers.origin_servers.public_name / 57a81bd1eb14 / 3

<a id="canonical-b5f87f4d37b2c3cd7cca75bd4a6ced9d2cb14f760b209eb39272a1c7a74889f8"></a>

<a id="canonical-2208d2c0eb2df0ce02b0ec4b0a6e6b714555ef8a3626890d8b2afa47e543728d"></a>

## dns_name property — origin_pools.pools.origin_servers.origin_servers.public_name / 57a81bd1eb14 / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

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
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1be4b21479e9947e2bf2534b9a789c78899768094e242198fc02677069f58cff"></a>

<a id="canonical-6d7a8c254ab945a934930e6887149a2a45a60fc715c2622155c78e9284a9c85a"></a>

## refresh_interval property — origin_pools.pools.origin_servers.origin_servers.public_name / 57a81bd1eb14 / 5

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-9146b2da49155887e7861f1a330591e5d82dd5b35eacdfede3e1b08c0fe95eaf"></a>

## Next pages — origin_pools.pools.origin_servers.origin_servers.public_name / 57a81bd1eb14 / 6

- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-dabfd280dc0a2512a3c39ee97956cc9c845f217e88ce62c6b4ccbcc6d8214638)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9251faea4dd06e9dd5ae705660a23649fdb50d22489344bf4321577a138996f"></a>

## proxy_advertisement — proxy_advertisement / cc387b9a4b51 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- proxy_advertisement

<a id="canonical-de0220efa04dfbaddb3b33e08a1cb7605d9751c475a3eb675c2d280db4a3f28c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for proxy advertisement.

Upstream description:

Proxy Advertisement Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_custom",
    "do_not_advertise")}
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
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"do_not_advertise\"]"
}
```

Terraform syntax:

```terraform
proxy_advertisement {
  # Configure direct properties listed below.
}
```

<a id="canonical-abc31210c3c17b3c7a15e9e8110afdc3093c89d2566d44e9de2f262750f9805f"></a>

## Direct properties — proxy_advertisement / cc387b9a4b51 / 3

- [advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc): complete subsection reference.

- [do_not_advertise](resources--bigip_http_proxy--reference--group-003.md#canonical-7e3d9d10858cdc4cd01a24952d34d04710876f9cb534696e3342242cd698fd1a): complete subsection reference.

<a id="canonical-b25c6cd657648e73f696db96956b7b54cd3ee35635f86b00561b577ba8584615"></a>

## Next pages — proxy_advertisement / cc387b9a4b51 / 4

- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.do_not_advertise](resources--bigip_http_proxy--reference--group-003.md#canonical-7e3d9d10858cdc4cd01a24952d34d04710876f9cb534696e3342242cd698fd1a)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6d268a7348dcac47ab6228568ec58c5d1a66f9b6e5f84187c9096b44060cff4"></a>

## proxy_advertisement.advertise_custom — proxy_advertisement.advertise_custom / 302e43e49384 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- proxy_advertisement.advertise_custom

<a id="canonical-755582f63681a94aa38c421dcc269820798375b29da54890002ae03df0131c05"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a VIP on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where")}
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
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-08e3d9f42bd8716f9bea8988b69a9c638b913003b8d8681f9aa0c0d1d562f350"></a>

## Direct properties — proxy_advertisement.advertise_custom / 302e43e49384 / 3

- [advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf): complete subsection reference.

<a id="canonical-fb685157bee3870754dfb03143a8dc015bcbbfba387e86a039e108b0cba158d9"></a>

## Next pages — proxy_advertisement.advertise_custom / 302e43e49384 / 4

- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d5d0637ece763c34fe4664afda7b9879de720df558b3959bced5176d6a7d943"></a>

## proxy_advertisement.advertise_custom.advertise_where — proxy_advertisement.advertise_custom.advertise_where / 3d15b9c7e61b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- proxy_advertisement.advertise_custom.advertise_where

<a id="canonical-9592b4b1df16d17e7db73b173ae486f7884f5145c480e5cb7bc4f0105800f3ed"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingListObjectAttributes("port",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("port_ranges",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site_with_vip",
    "vk8s_service")}
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-f7a073d358a4cd7c2541502d4e4890e7004432f2ba028d73faeaf0d023570fda"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where / 3d15b9c7e61b / 3

- [advertise_dualstack_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-cd68ea499895e8d9b1aeb06c1f12128603525aca03d2f3aff71e472c253d38d6): complete subsection reference.

- [advertise_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-15504373f8aa2a06e590d8b40a999b29b684cbf9f12eee1cb8806443fa308aa3): complete subsection reference.

- [advertise_v6_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-c2c4a00906c27609afe9f76d46c551aace43aa2d27bf5362280691d9e886c933): complete subsection reference.

<a id="canonical-315969f14aa8272ade8ce195fb24743cfd0d8263fbb6b58f7e94586e18af758a"></a>

<a id="canonical-f28e2ce0c628fe81311f4bffb44a37500ef426500908c9684be76f9bf0300b1a"></a>

## port property — proxy_advertisement.advertise_custom.advertise_where / 3d15b9c7e61b / 4

Type: `"number"`. Optional.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Upstream description:

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-7f65c3cb0f8f4ca7b68404819b0e1229c4bfbfe31618c6b65e026697a5065c7d"></a>

<a id="canonical-4fc1c0111021f4b0d42111a4d874f5096269eb45ae755c7423a9c2b589bdf175"></a>

## port_ranges property — proxy_advertisement.advertise_custom.advertise_where / 3d15b9c7e61b / 5

Type: `"string"`. Optional.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

- [site](resources--bigip_http_proxy--reference--group-002.md#canonical-4346e75717d7baa0182be4e789ff3916a8769328679cdbc3c9b8172bb7070cbc): complete subsection reference.

- [use_default_port](resources--bigip_http_proxy--reference--group-002.md#canonical-8523aa6b919957142880fc7cac8c90c376f5030a138c7f0b9cc9f243e1c85314): complete subsection reference.

- [virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-332e3445f4286df33482e97fb8624323c1f67e0858e5ec6338d4021b519dbe4d): complete subsection reference.

- [virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-bf980df380bb8ef45ed1c6e249fdb114b08e941962ba5718f99d6d1245595f2c): complete subsection reference.

- [virtual_site_with_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-a4a832f2fabaff385fb8593ab2c4e64dfe2a90762bff1a616622779c64693c73): complete subsection reference.

- [vk8s_service](resources--bigip_http_proxy--reference--group-002.md#canonical-e11a4d6a07eb2cc00ffc458854833379b43949e4abab740e6a8c423949a6674e): complete subsection reference.

<a id="canonical-f9c6378e15fc8afb9999e597b56dafb2401450b8754503e8ed18be04d76d3e79"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where / 3d15b9c7e61b / 6

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-cd68ea499895e8d9b1aeb06c1f12128603525aca03d2f3aff71e472c253d38d6)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-15504373f8aa2a06e590d8b40a999b29b684cbf9f12eee1cb8806443fa308aa3)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-c2c4a00906c27609afe9f76d46c551aace43aa2d27bf5362280691d9e886c933)
- [proxy_advertisement.advertise_custom.advertise_where.site](resources--bigip_http_proxy--reference--group-002.md#canonical-4346e75717d7baa0182be4e789ff3916a8769328679cdbc3c9b8172bb7070cbc)
- [proxy_advertisement.advertise_custom.advertise_where.use_default_port](resources--bigip_http_proxy--reference--group-002.md#canonical-8523aa6b919957142880fc7cac8c90c376f5030a138c7f0b9cc9f243e1c85314)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-332e3445f4286df33482e97fb8624323c1f67e0858e5ec6338d4021b519dbe4d)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-bf980df380bb8ef45ed1c6e249fdb114b08e941962ba5718f99d6d1245595f2c)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-a4a832f2fabaff385fb8593ab2c4e64dfe2a90762bff1a616622779c64693c73)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--bigip_http_proxy--reference--group-002.md#canonical-e11a4d6a07eb2cc00ffc458854833379b43949e4abab740e6a8c423949a6674e)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-cd68ea499895e8d9b1aeb06c1f12128603525aca03d2f3aff71e472c253d38d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4faaf515df94e6c191ee0990e6c485f4442207b871fef339fb75289ed15a7fbb"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / a9df1e666e9b / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-182716908c40545b1d73f0eaee45aaf610773bc5c6012d9313051b445e407bc4"></a>

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

<a id="canonical-3eafffcd58d94f6cfddfb048c58712a8d12f87e18388a2f242a1648c7aea0794"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / a9df1e666e9b / 3

- [public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-d5708a61e0be02e3d97fdae8118a289218988d960409d1b96805170785b3a9b7): complete subsection reference.

<a id="canonical-afca4aab84254c16f8d962833a2f004aba086f0b5c1250de91534a2eac2a88c9"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / a9df1e666e9b / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-d5708a61e0be02e3d97fdae8118a289218988d960409d1b96805170785b3a9b7)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-d5708a61e0be02e3d97fdae8118a289218988d960409d1b96805170785b3a9b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be93adc3cda8a5a76075d4e529bf414ae8e90ece9b895c8e09da2875cdcdd6ff"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / bb5f3940bc1a / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-cd68ea499895e8d9b1aeb06c1f12128603525aca03d2f3aff71e472c253d38d6)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-5d6eca053f9283a4e7e45a30d0cedfb69092c8484976e083df7581389e9b6909"></a>

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

<a id="canonical-7a91111d4a35c888c218af8aecf9a239485bad9edd6fc715c13d3fbfe83eebe9"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / bb5f3940bc1a / 3

<a id="canonical-97d0882e0e816e2cebf956f53f5a2331b06364435cd0367a1b4a4493731334b9"></a>

<a id="canonical-506b5756238d0aa46c86b1f646a2bfa37a18e9973c63d454ef1808b434d75887"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / bb5f3940bc1a / 4

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

<a id="canonical-5b6595d513bddce700b6822cf86ee991e00e5f3ed7f5264ef471e27fd043fbde"></a>

<a id="canonical-2c86d7181c6a00a9f91df796cdfa29e7434727cc5f2c5ffe416b653a44e6a12f"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / bb5f3940bc1a / 5

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

<a id="canonical-06a4f2f060c8353300daf94c5b3fe7aefd80ee9881dc81af18a4178bfc5840e0"></a>

<a id="canonical-d5e75eee20397439561c49e73496a135162bf20765c06b8422b2294d3dfc1fff"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / bb5f3940bc1a / 6

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

<a id="canonical-1b72f2517c98067ae9de270269179d831bab614fbc2bf8f92906d0d21150b111"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / bb5f3940bc1a / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-cd68ea499895e8d9b1aeb06c1f12128603525aca03d2f3aff71e472c253d38d6)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-15504373f8aa2a06e590d8b40a999b29b684cbf9f12eee1cb8806443fa308aa3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-740da1d8eb595691b961d185c768a38210d4692af82bc2601ed371e7ea56338e"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_on_public — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public / dacf3cfe170f / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public

<a id="canonical-11c923eeea590c7be6948024b83ac6d61182cf95ac2b1338f54425a532381dff"></a>

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

<a id="canonical-6cd4d271fd72bb48b1029cbfaafa2dbc6b1c9b0a469337fb6d587a9d99435403"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public / dacf3cfe170f / 3

- [public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-8a7b6a633bb7fd1119a76eeac3f30f0882de7ae5a54d8513484cb49e421fba25): complete subsection reference.

<a id="canonical-c938be72e92c75ecb548f1727a1455b853ee29bc0f4642a7bc7b573de97d76bb"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public / dacf3cfe170f / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-8a7b6a633bb7fd1119a76eeac3f30f0882de7ae5a54d8513484cb49e421fba25)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-8a7b6a633bb7fd1119a76eeac3f30f0882de7ae5a54d8513484cb49e421fba25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc5e3a9308a5a4d2825f7f3e0e0ef282bd4c1c68eb5f1f5e7fe94815f0cb9fac"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 2b34e4e4d003 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-15504373f8aa2a06e590d8b40a999b29b684cbf9f12eee1cb8806443fa308aa3)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-fce3fe0e31050b884b6b8ab2b6b87e4a0e6a57522acfc4437f23290872951611"></a>

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

<a id="canonical-180598ce0bd76250cc59c96ddc85f011bfa2a1fa5916b05d7c293828be496f1c"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 2b34e4e4d003 / 3

<a id="canonical-e5c66680a5cd48a01aeaa4ce3ffa0e8d38ad3bc3b719daa57899dcf77c53a28e"></a>

<a id="canonical-57cf1adcbef60581b240c9640ee05b5473467e812e586e94aabdbe327c740502"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 2b34e4e4d003 / 4

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

<a id="canonical-d54c1a11f451eed0ffef1c6de8046cdedf115a2cffc840d03f849615d0548f35"></a>

<a id="canonical-85763d8435034777b2cf1f82aa2fd61dd2070b63fb9a06c34d7a3c31f6d747d8"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 2b34e4e4d003 / 5

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

<a id="canonical-b0d8eb69bf116aef9d7a9e3a466db64efe1cd1fd4cde36499ff86be5a0cf6db7"></a>

<a id="canonical-55ae8f2e01b219b5594b1ea517ebfb5193f82516dc75a17a2abb2d99d56ddff8"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 2b34e4e4d003 / 6

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

<a id="canonical-2a17a2b25fd2de1347097f758e0fbeb2843777da3fa00fd5b3c0ea07974e294a"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 2b34e4e4d003 / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-15504373f8aa2a06e590d8b40a999b29b684cbf9f12eee1cb8806443fa308aa3)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-c2c4a00906c27609afe9f76d46c551aace43aa2d27bf5362280691d9e886c933"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7afa19437a71e00da1a36fa616df2e69df05887c155d980bee69bea96a2ee22"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public / 33ebbde617da / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-39f0be748941fc91cdb2279340e238968721da63c3267e6b3d5f26cfc198f541"></a>

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

<a id="canonical-bbb213019bebca462e785e0f4c1e4ea0c16511893b18ffba3e0e64cf6b3b3c0e"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public / 33ebbde617da / 3

- [public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-b696f971b4c113bbf52756f2018f6b80438bf552307322890b4ca6163d827a5e): complete subsection reference.

<a id="canonical-834ee813f4b3e381ce747f150181b69c83a2cc1ab9d03d29560f425adbbeb239"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public / 33ebbde617da / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-b696f971b4c113bbf52756f2018f6b80438bf552307322890b4ca6163d827a5e)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-b696f971b4c113bbf52756f2018f6b80438bf552307322890b4ca6163d827a5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04c3d74b08457d4c61a3ab4853c276e834cb771c60f973730fb95066a5d8d9bb"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / 53e605e0b64a / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-c2c4a00906c27609afe9f76d46c551aace43aa2d27bf5362280691d9e886c933)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-aedc0702d3ab0c5ccd461b0ae45cda6836640d6111e16f81ab869b04fc858e42"></a>

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

<a id="canonical-0f2bc8c6834965963fd13be39eaf759cf8cf8258ce9dae0b22c3a0863de58aa1"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / 53e605e0b64a / 3

<a id="canonical-4c69ca6a5bb663e799aaf73145095d3eb6326c731ba7cfdb9543c423db998d06"></a>

<a id="canonical-09bada34f6bba8841653e174f4d8771f0f6c802ef0adc4aa95609c99b098a47a"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / 53e605e0b64a / 4

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

<a id="canonical-9ab86565a10ac424d7e020463991265fd6612670f4dc4564baec29a0f187bc4b"></a>

<a id="canonical-5fc21bbf7674a3ced9e7052c1e7ca453730c151ac16031cef3310c5a3d8958b5"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / 53e605e0b64a / 5

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

<a id="canonical-4de201bc27912c9b172e53d701f213be2a7365e3180f08afdd8677b913f76ada"></a>

<a id="canonical-7608c3ff354f251882de408bd9df90da1c22965aabd3dfe7f93cab3927286e51"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / 53e605e0b64a / 6

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

<a id="canonical-3c29a986b5e367b4868ada75f4f724c0cd5042b83670ca5b43f8b862a8b43644"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / 53e605e0b64a / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-c2c4a00906c27609afe9f76d46c551aace43aa2d27bf5362280691d9e886c933)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-4346e75717d7baa0182be4e789ff3916a8769328679cdbc3c9b8172bb7070cbc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f8f088eb5f46cc3b4647dc653c8cf7a7184a41b33e6dba4ea32d88a853be4d2"></a>

## proxy_advertisement.advertise_custom.advertise_where.site — proxy_advertisement.advertise_custom.advertise_where.site / 20404ee8ff54 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- proxy_advertisement.advertise_custom.advertise_where.site

<a id="canonical-1c69909f5aa462fa6fdcd880c074cc2d3d08596a63dcd29bbc4e83aa94692c4c"></a>

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

<a id="canonical-aafa18f27c357c9d41c5a2f7d2d24fc0ddfe5cbc58a21674a39c4d9d12f33bd7"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.site / 20404ee8ff54 / 3

<a id="canonical-7db859a85787b73335e8433014cedbfb0e7dbc7fb8a8bf738dd48d99dccc026d"></a>

<a id="canonical-cd303596f9acfc9c72f9c1183c09a34eceb3743464a2304419af877a5683c8c5"></a>

## ip property — proxy_advertisement.advertise_custom.advertise_where.site / 20404ee8ff54 / 4

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

<a id="canonical-dac11af03fd7221b8ccbf4bc5cc93e8cabfa1d8ddbf1a498a431d9c9997925ca"></a>

<a id="canonical-9826699f79a8707225b5d8bb64a1257919ded7a15d7e56a82d63f23c7057fb60"></a>

## network property — proxy_advertisement.advertise_custom.advertise_where.site / 20404ee8ff54 / 5

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

- [site](resources--bigip_http_proxy--reference--group-002.md#canonical-d35093d06ec753cbf1695cd64644f6f92526b830d0383f427c22ab6b864edcea): complete subsection reference.

<a id="canonical-eb69f139c05fe977ce2b2c982e09b2748c00077861b490ad4422081d415ac185"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.site / 20404ee8ff54 / 6

- [proxy_advertisement.advertise_custom.advertise_where.site.site](resources--bigip_http_proxy--reference--group-002.md#canonical-d35093d06ec753cbf1695cd64644f6f92526b830d0383f427c22ab6b864edcea)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-d35093d06ec753cbf1695cd64644f6f92526b830d0383f427c22ab6b864edcea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e7644ef40acea072a8e7ce1657b341d3b61165484dc21350ea4100283417b20"></a>

## proxy_advertisement.advertise_custom.advertise_where.site.site — proxy_advertisement.advertise_custom.advertise_where.site.site / 2a850b1352c3 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [proxy_advertisement.advertise_custom.advertise_where.site](resources--bigip_http_proxy--reference--group-002.md#canonical-4346e75717d7baa0182be4e789ff3916a8769328679cdbc3c9b8172bb7070cbc)
- proxy_advertisement.advertise_custom.advertise_where.site.site

<a id="canonical-8738bd360de97dacfe818b46010911db1020378900683bfaa059ecde291adc63"></a>

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

<a id="canonical-e5f39edda69fc1f22b9432cce98399ce8eac120329d1b9197e6324bf1e8aa13b"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.site.site / 2a850b1352c3 / 3

<a id="canonical-2115b2293c3dd747abb667ca9e3b62e4f67bdd8f7f7cbe39947c3ed854015e57"></a>

<a id="canonical-7fe4c044abaed7aaf47f36fa59caf6a197b93983304eaf2c6bdb96d62e98d621"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.site.site / 2a850b1352c3 / 4

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

<a id="canonical-09fe766955784b930329d3d882f176b1a9ca7204662ce2fb0df1bf7d32bb52b4"></a>

<a id="canonical-818a2ca1497ffa860164a92c37af50103e34248e9f1bff3f2330e9ca296bf07f"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.site.site / 2a850b1352c3 / 5

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

<a id="canonical-9c538802947c2587b8f2f31ddb98510ab2e31b5a87430dec3df9063d31605f98"></a>

<a id="canonical-ba950a3818d5045b7da8d98d1e9d11bf51addb28998ffe3e9f2e756e22330db3"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.site.site / 2a850b1352c3 / 6

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

<a id="canonical-8f19132a2251c941586e2dbd7097c2a0abe5e2127d906d9d57914eb3ccb076a7"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.site.site / 2a850b1352c3 / 7

- [proxy_advertisement.advertise_custom.advertise_where.site](resources--bigip_http_proxy--reference--group-002.md#canonical-4346e75717d7baa0182be4e789ff3916a8769328679cdbc3c9b8172bb7070cbc)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-8523aa6b919957142880fc7cac8c90c376f5030a138c7f0b9cc9f243e1c85314"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a3c7ef752962c08687fa5ddb5fbf7e3290a5503e4ca54a3f1292c39f05ccf48"></a>

## proxy_advertisement.advertise_custom.advertise_where.use_default_port — proxy_advertisement.advertise_custom.advertise_where.use_default_port / c54b9b053bf7 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- proxy_advertisement.advertise_custom.advertise_where.use_default_port

<a id="canonical-87a3cb118377001c4680bf117fc15155871afe6823cd03a9b9c59513c0e93eb4"></a>

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

<a id="canonical-adf7bead3921b32f8fea1ca9f271d26b1e2154f6e58760bac882ba88d2793517"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.use_default_port / c54b9b053bf7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-86e628ee8e05cd58ea8fcc2b8bf727078b07eeaabd8360c640924617f5523c24"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.use_default_port / c54b9b053bf7 / 4

- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-332e3445f4286df33482e97fb8624323c1f67e0858e5ec6338d4021b519dbe4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab9b5195d89ca77cad3d1aca76bfe7eeaf8305db9eb4ded6fbc3b1dff783fecd"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network — proxy_advertisement.advertise_custom.advertise_where.virtual_network / b8bde516513e / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network

<a id="canonical-80246c1f2b956d0fbd276631b7798159d7ffef977a200e0e20774867acfb03cf"></a>

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

<a id="canonical-7905e6efaee6b3770a3a2686f2d84547a89c69a52a7fe3b9feb6019eb54eb143"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_network / b8bde516513e / 3

- [default_v6_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-ae7858ac0520e6adba454937a793c766e74950af0bd178ba160150b298d29495): complete subsection reference.

- [default_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-c812c249dc8328ac37e040d6e5689897cfa14e2dcb8a2d50f3b168449b697fa5): complete subsection reference.

<a id="canonical-ce241ed070877134d2377eb857388e3b4a5a690c3ab3f970af5bffdcc46cd54c"></a>

<a id="canonical-002744b3a18704cedfe432443190b68183e052c31293873dc8b0cb9625af4f09"></a>

## specific_v6_vip property — proxy_advertisement.advertise_custom.advertise_where.virtual_network / b8bde516513e / 4

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

<a id="canonical-348c7c6e0a48e5dcf06d8b5c3824491b2d680a8a89bc9f3e5895bb8408549ac6"></a>

<a id="canonical-9688a5bc41f1de76e8bfb7dad7ae8422bcd403d435c497e06d73071bcaf9e439"></a>

## specific_vip property — proxy_advertisement.advertise_custom.advertise_where.virtual_network / b8bde516513e / 5

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

- [virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-ea27f783bf7a7f41df8fde31159fb353bd72736bd8b39c1f6857248dc207f6bc): complete subsection reference.

<a id="canonical-2efb59d750fe4ebf898772e838c33e235780d21741ff3086263b10809161ecd0"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_network / b8bde516513e / 6

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-ae7858ac0520e6adba454937a793c766e74950af0bd178ba160150b298d29495)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-c812c249dc8328ac37e040d6e5689897cfa14e2dcb8a2d50f3b168449b697fa5)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-ea27f783bf7a7f41df8fde31159fb353bd72736bd8b39c1f6857248dc207f6bc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-ae7858ac0520e6adba454937a793c766e74950af0bd178ba160150b298d29495"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80f1bab14ec0fb912ad7518d28f45a35d640f054810d92d175c308c11fde3d0f"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_ / 05cdf17179ad / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-332e3445f4286df33482e97fb8624323c1f67e0858e5ec6338d4021b519dbe4d)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-cd4b92f103be0b6690a75828490935b59978f28f516fd5c4e2199506e7ec5f53"></a>

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

<a id="canonical-fccfc986c86e3258aae8d9ee3b27b4abb21bfa60a1a8cc3a334c14f7ae0a99c0"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_ / 05cdf17179ad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4d93e37487ce9868995fbdd027bbda553beeb96e3947c5212fd53ef659386dbf"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_ / 05cdf17179ad / 4

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-332e3445f4286df33482e97fb8624323c1f67e0858e5ec6338d4021b519dbe4d)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-c812c249dc8328ac37e040d6e5689897cfa14e2dcb8a2d50f3b168449b697fa5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f067ebc61f14dd1602f936e5962f0f980a11e6bd09f45bcde0841eaf714adc6"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip / 4bba43025ae9 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-332e3445f4286df33482e97fb8624323c1f67e0858e5ec6338d4021b519dbe4d)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-bf6f54fa3f0a357a84dcbe813d6ab4c06ac120660f839be3df74564e7cd7f7bf"></a>

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

<a id="canonical-5ced9ac3a7a48ccfdbad6158f916de61ae8c400b8f12065975e4eabe604c4a81"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip / 4bba43025ae9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7903a452deecb5d8995e4e7a80a0cbb759eeda8df9362ac836428dddd2533195"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip / 4bba43025ae9 / 4

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-332e3445f4286df33482e97fb8624323c1f67e0858e5ec6338d4021b519dbe4d)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-ea27f783bf7a7f41df8fde31159fb353bd72736bd8b39c1f6857248dc207f6bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d2a075384cc13eedffa95fcf9f5c64916accecb65182b5ad815cbf4401e32ab"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / 78c3d515e769 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-332e3445f4286df33482e97fb8624323c1f67e0858e5ec6338d4021b519dbe4d)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-4a353091edc4ee2d679bde8f5a8d2eeb4168e33fd43c8416796a0fff222aa652"></a>

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

<a id="canonical-511870f4a62c0e89f4cf2cf6ddb672046134857aa0d49067e10d131754df2a51"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / 78c3d515e769 / 3

<a id="canonical-380824253800a008eb5499764dbaea5d6dbae4489cdd441e5d33c14607230072"></a>

<a id="canonical-7f2e9dfe997da8e2eeaeebacc898253c75ea82f72020a3920c6cd95bff645034"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / 78c3d515e769 / 4

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

<a id="canonical-dbfb66b43d236c1324fac70758d1b05222c4283a43bdde5b1ba908555b330589"></a>

<a id="canonical-de15668f45baf2d6a0bf831f079ce8f78a4386fa5a7223f00fc3df030eb7dc5a"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / 78c3d515e769 / 5

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

<a id="canonical-b962ea151483ea8f1b4ea48b1dd8defa0a9bdc4177fc2c22c29b69359ccfd413"></a>

<a id="canonical-5c98b95998aa1736cd2ee3166a6f34db1850f9605e9691120ccd11fdf30dd1a9"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / 78c3d515e769 / 6

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

<a id="canonical-c48985c88db4499982591256c8c894b31fd4c8bc7cd637ee98c99c238b41f719"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / 78c3d515e769 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-332e3445f4286df33482e97fb8624323c1f67e0858e5ec6338d4021b519dbe4d)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-bf980df380bb8ef45ed1c6e249fdb114b08e941962ba5718f99d6d1245595f2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b885517ecd8c93b89ef8654aa0718f7bb955390646b85b5b20ebedf4fa8f50c"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site — proxy_advertisement.advertise_custom.advertise_where.virtual_site / f263afd7bbcd / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site

<a id="canonical-dbf4770aef71706cad2881037199031396b77850bbb3c0d3abde75faa06e21fc"></a>

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

<a id="canonical-29fbfb75268bca53e0ef4e0c8b9bb4bb718209dcefd41fd811f150d653f67764"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_site / f263afd7bbcd / 3

<a id="canonical-de7f32319bdd3d7669f32d512496c84ae416e1676dc9671628f02d40891a32db"></a>

<a id="canonical-7950e642871e1978e492e892af51305592bc6a6507daefcc41b9f0f961443558"></a>

## network property — proxy_advertisement.advertise_custom.advertise_where.virtual_site / f263afd7bbcd / 4

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

- [virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-6aae49674a0ed2a36c4e449485ec3989eecfadf106b338be32a7ba279ec2a3cb): complete subsection reference.

<a id="canonical-60b82fc01950ed9a14f1d3341305bdc0c7c4d99420e1a2af266c635b64cd700b"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_site / f263afd7bbcd / 5

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-6aae49674a0ed2a36c4e449485ec3989eecfadf106b338be32a7ba279ec2a3cb)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-6aae49674a0ed2a36c4e449485ec3989eecfadf106b338be32a7ba279ec2a3cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38d96c2b99a9f8cec304d1f815f25ebd0d455c9e34414b441e1b858360969168"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / c404e4abe295 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-bf980df380bb8ef45ed1c6e249fdb114b08e941962ba5718f99d6d1245595f2c)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-e63b8f6f148c1543675f59fd0c9c8907e1943c821ff18292e374ddb54d0da10a"></a>

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

<a id="canonical-ea611b780eaa4647dff39242bf0bc6e5a2fd86850b2e7aa48bcc88ded284771f"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / c404e4abe295 / 3

<a id="canonical-bc23ac8f1f16d1389125dcad15c58a66f91af38dde7a98e80b7229633d7c3eac"></a>

<a id="canonical-2ab90e0e79a1bb2e6c3ff916fd6da8ec43fc82466f57b4a1740f7ba2990b0e56"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / c404e4abe295 / 4

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

<a id="canonical-88680d258b39fff6410521f73e464ce14481dbefc3eeedb88ed6cfa95216dac6"></a>

<a id="canonical-a9cbe8067bad2a2a09508658fd522f3fd8a55cf0dcbeee46113d0bd7fb782cd7"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / c404e4abe295 / 5

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

<a id="canonical-3bfaaf47181961ac99edbb9a6940af02b39632c8f97d822c24941d66a5d732fa"></a>

<a id="canonical-f9fdc857bfe0df026b97e5956a32d9788e15519051bbe558f8a560e36002e5ba"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / c404e4abe295 / 6

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

<a id="canonical-1c4cd397837e999e407acac043e3cb231ba1e8d172765b938d40f227ea536869"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / c404e4abe295 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-bf980df380bb8ef45ed1c6e249fdb114b08e941962ba5718f99d6d1245595f2c)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-a4a832f2fabaff385fb8593ab2c4e64dfe2a90762bff1a616622779c64693c73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4645e5d93aa0ef66673db299d3cd89cb744d323d3450ab42a4846d5f9baca40"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / 244e01699619 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-28c2dde4fed71e5ccce17e37f0d124cdc3d81f509b5f2a8b805fcc3d769cf459"></a>

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

<a id="canonical-e3a317d38bb2d9b1f7c104b19223390db789176927368fdda7cc82518cb3cf99"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / 244e01699619 / 3

<a id="canonical-cbbd452b1618d3ebdadfa020da4bc2debc59beb706240c802df23cc33467052b"></a>

<a id="canonical-8a1cdb3510bc9471e9b9e74594880a218140ebaf0999d83b3a7ae37d463d6695"></a>

## ip property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / 244e01699619 / 4

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

<a id="canonical-85449426bfa455c0af996a0532d936ce6809d27583133c2f2d3cf82b254cded1"></a>

<a id="canonical-1907bf47505cc63990b9e7c94f1c0cbb01c26708088d01a9cf9df7a2381e2418"></a>

## network property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / 244e01699619 / 5

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

- [virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-e9089aee5f6ee40234a28e9b3ab1c946a773bd158ab2edbfc949177658dceb5f): complete subsection reference.

<a id="canonical-c2713388193a1113c6809848cc9cd73c1fe3c0ae764f462699e433e9c035b876"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / 244e01699619 / 6

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-e9089aee5f6ee40234a28e9b3ab1c946a773bd158ab2edbfc949177658dceb5f)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-e9089aee5f6ee40234a28e9b3ab1c946a773bd158ab2edbfc949177658dceb5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d2cdf2d2af5d1f668ecc652fe715a53ddccf99e4863f00c1efbeb6d4d47195b"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / 05f120862b75 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-a4a832f2fabaff385fb8593ab2c4e64dfe2a90762bff1a616622779c64693c73)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-4206592ea7b27b2996e52d3463b2a535dfcf79fc80a9c82abcbcf7ae820b555b"></a>

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

<a id="canonical-1f9aa7535182f68029eb84b573194648ac4f3211147fae794243ec4b56c67d86"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / 05f120862b75 / 3

<a id="canonical-fa8c99780e2f0d51918945b36bcaf8d74df7281cf5f33a3e725790003341b5b5"></a>

<a id="canonical-c857bb7d0a89a33737b163f130243cef5dc95ccfacb59e7e842dd212396e29a9"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / 05f120862b75 / 4

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

<a id="canonical-857b421813c57394af9a9ec74c2a0ea55061971590ca925def86cc7adbf98395"></a>

<a id="canonical-5d00b65d825eddc22e6785dc1a40221411805c524781a6b25a03824b9b0fb752"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / 05f120862b75 / 5

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

<a id="canonical-93c6d422bd06b7fad4df483fa6253dd112fe4acda4df4bc966a59a78d143fb4c"></a>

<a id="canonical-540f60640e7be56e9a31a06fdd67f26efdaf944c7aca68836f84ce703bff90c6"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / 05f120862b75 / 6

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

<a id="canonical-93f2384720f81d29ba2e157bac9188c9ffb7ca8639921dc73970c7d53e4ca7b7"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / 05f120862b75 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-a4a832f2fabaff385fb8593ab2c4e64dfe2a90762bff1a616622779c64693c73)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-e11a4d6a07eb2cc00ffc458854833379b43949e4abab740e6a8c423949a6674e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b58417ef297ab5ffd3e9edc9d953d62aaf843c923afda65bc6be11ea6c72d66"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service — proxy_advertisement.advertise_custom.advertise_where.vk8s_service / 46f4bd7653c2 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service

<a id="canonical-f10cd8e87c0cc3b92f2f267f51611f15c31703bacebec7c68af99e68087bc24c"></a>

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

<a id="canonical-954656e551041e3bad992dda99c8274614ce460445ea43762ed55e785aacc636"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.vk8s_service / 46f4bd7653c2 / 3

- [site](resources--bigip_http_proxy--reference--group-002.md#canonical-ce0b3c13ee8d382cfd05bec5dd870f98a8e1b8e171bac0019727fe6a8400dd4e): complete subsection reference.

- [virtual_site](resources--bigip_http_proxy--reference--group-003.md#canonical-a1324d8ef4347390b9f1e100b6ece47ce241b1c260e0eb27d90921134eca1262): complete subsection reference.

<a id="canonical-56889c1e3bb70f1aad74c325f79ccac37e63addf2e25455a75e07ecdbd714cc2"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.vk8s_service / 46f4bd7653c2 / 4

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](resources--bigip_http_proxy--reference--group-002.md#canonical-ce0b3c13ee8d382cfd05bec5dd870f98a8e1b8e171bac0019727fe6a8400dd4e)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--bigip_http_proxy--reference--group-003.md#canonical-a1324d8ef4347390b9f1e100b6ece47ce241b1c260e0eb27d90921134eca1262)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)

<a id="canonical-ce0b3c13ee8d382cfd05bec5dd870f98a8e1b8e171bac0019727fe6a8400dd4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec06dc57f18e912bca2e69fcfde3f21fc050ca2285f68e18d9098384e74591d5"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / 71de83fa7095 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-4e40823c68de58167d94875d3845c4b48a0bcf37c511b874ff25e8e1d45d2815)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-1187eaa1f23c907a72464d7da4fe4e03023ed0e55f203523207d775d11055680)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-2a89bae749fbef43ff1ed28b601d6404c949ab646437965ef3776fba8218ad5a)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-03c906c318cb23d88fa6dc4c79d25c56ba8d805b6ce1151d6c3faeb58afd76fc)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-14ff5e5a12ebe72a0729ad520c70dc8b1e16ed514b4a72e99b09379e44f920cf)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--bigip_http_proxy--reference--group-002.md#canonical-e11a4d6a07eb2cc00ffc458854833379b43949e4abab740e6a8c423949a6674e)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-3418d5a3ab7b68e301977b1daab521119ea3863ffcf6e790e868b29240bfe174"></a>

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

<a id="canonical-aca0a38a47e3d632ead9fb971f3b89a666c5b5520e7b866b9b6cb71a7f21b4e9"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / 71de83fa7095 / 3

<a id="canonical-fe67a041527f98c67e9cefcd9284b4bddb47925a250c424e4914b8387159c6b9"></a>

<a id="canonical-0b1e6fd2bf5054275ea15f621baaa187169291bfca1f0a2868b575f3cd6f5d69"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / 71de83fa7095 / 4

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

<a id="canonical-a2ab1b5194e4582e61ebb41ced16ce8f39eb84e6a884fc4771584d52fe59902e"></a>
