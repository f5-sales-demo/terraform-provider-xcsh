---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-81ae4a16966727a9b5fc6f9503aa6b536710ca0341de0796ad87eb1abc35f99a"></a>

## Direct properties — default_pool.origin_servers.private_ip.site_locator / 02a7ed928798 / 3

- [site](data-sources--http_loadbalancer--reference--group-016.md#canonical-76c611b9771b40c4f9a41148062fbbf351e8d1a19d709c6f67616afdfa95235c): complete subsection reference.

- [virtual_site](data-sources--http_loadbalancer--reference--group-016.md#canonical-7d87dd6996e594901886d1b3d01fb27a08024dd5f319bb09e6aede10678e4183): complete subsection reference.

<a id="canonical-09fe5f693c144c882c8398a62165e6f64ae93dacd7b6f51fb1cd0c79b32804fe"></a>

## Next pages — default_pool.origin_servers.private_ip.site_locator / 02a7ed928798 / 4

- [default_pool.origin_servers.private_ip.site_locator.site](data-sources--http_loadbalancer--reference--group-016.md#canonical-76c611b9771b40c4f9a41148062fbbf351e8d1a19d709c6f67616afdfa95235c)
- [default_pool.origin_servers.private_ip.site_locator.virtual_site](data-sources--http_loadbalancer--reference--group-016.md#canonical-7d87dd6996e594901886d1b3d01fb27a08024dd5f319bb09e6aede10678e4183)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-76c611b9771b40c4f9a41148062fbbf351e8d1a19d709c6f67616afdfa95235c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f5ae642b0152a6aaaeb5ddd0cd7f6733e4693d28d0e95f291d4a3c8c382e8a5"></a>

## default_pool.origin_servers.private_ip.site_locator.site — default_pool.origin_servers.private_ip.site_locator.site / 715cce8d1882 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- [default_pool.origin_servers.private_ip.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-1808bea5ce681ad3cdeb2595342ab193ebbbbfe62f172282c5236ffa8b1c02d5)
- default_pool.origin_servers.private_ip.site_locator.site

<a id="canonical-64f19c0e354e8f6518d1014130cbf7e0fd2e9558cbeda0fd2776e95d4469d56d"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-e6e88600d14aef7ae03e365e36ee1eadb30ef41ebd31c4009ae4947cae1a146d"></a>

## Direct properties — default_pool.origin_servers.private_ip.site_locator.site / 715cce8d1882 / 3

<a id="canonical-47d0671930b524c8f77806b11f59de6ecdd1a261fa132a013f1fa619307f1191"></a>

<a id="canonical-f2c61c2d46735c7b38982726e42474a0d7dfdf1f22cab79d08f812af4a5d3a34"></a>

## name property — default_pool.origin_servers.private_ip.site_locator.site / 715cce8d1882 / 4

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

<a id="canonical-dd1f30ed60bb0831acfa619c50274be155dddce6a9e3da94c5f17a6391f75dcc"></a>

<a id="canonical-bec2a62680a6f9db14a6eee54a965c25a9687469e1cace2d570ab8cda3038e1f"></a>

## namespace property — default_pool.origin_servers.private_ip.site_locator.site / 715cce8d1882 / 5

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

<a id="canonical-3bbf79748ec200350489e28dd565c439561e31b3894c3a283cfdf980125b2c9b"></a>

<a id="canonical-5418b370c9114edba063b959f5fe80d5cc8e57cac8ce2eb024451899406cd115"></a>

## tenant property — default_pool.origin_servers.private_ip.site_locator.site / 715cce8d1882 / 6

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

<a id="canonical-d6e97b0abedf4d1589fab23bc9acc59121135fc286414466900ee48c01c704aa"></a>

## Next pages — default_pool.origin_servers.private_ip.site_locator.site / 715cce8d1882 / 7

- [default_pool.origin_servers.private_ip.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-1808bea5ce681ad3cdeb2595342ab193ebbbbfe62f172282c5236ffa8b1c02d5)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7d87dd6996e594901886d1b3d01fb27a08024dd5f319bb09e6aede10678e4183"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89e3ab16132d15877265e3f230e32a3e00a26501d386f532c3563d56aa5b783e"></a>

## default_pool.origin_servers.private_ip.site_locator.virtual_site — default_pool.origin_servers.private_ip.site_locator.virtual_site / 6183abeb8f48 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- [default_pool.origin_servers.private_ip.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-1808bea5ce681ad3cdeb2595342ab193ebbbbfe62f172282c5236ffa8b1c02d5)
- default_pool.origin_servers.private_ip.site_locator.virtual_site

<a id="canonical-87caa88b83230b9410683717579a55f8b9a56dbc3b7168a4d6470f9f78b136d4"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-65babbfe75751b89d9dae6a5fbfbf953424972e80fe2c3fd0d7f668b9673082a"></a>

## Direct properties — default_pool.origin_servers.private_ip.site_locator.virtual_site / 6183abeb8f48 / 3

<a id="canonical-c31f1149c6e811feac08a8214c60f871df1afc0eb29ac014c77b07228191bdbe"></a>

<a id="canonical-a2c72b3f9f7964a7b570214495b6b72af04d9e3ccb6d288ae3e4ab78c9416d7e"></a>

## name property — default_pool.origin_servers.private_ip.site_locator.virtual_site / 6183abeb8f48 / 4

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

<a id="canonical-9a0b3e81b48e5fb6c961564542159e863f6382b1ac506b38ce43a0f55a0d2b20"></a>

<a id="canonical-dd0c5be385c626560de55d744ee708ff51ad64ee5b1171eefb9c81e447d866ab"></a>

## namespace property — default_pool.origin_servers.private_ip.site_locator.virtual_site / 6183abeb8f48 / 5

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

<a id="canonical-f9aeecdfa48def714d54fb3290a88e02e63530b7ed7fa5f7712743ab94e04246"></a>

<a id="canonical-c7e47b0f1d90fce41118aabe00822f44fd01c1fcc1abbb17fcd9b1d6bea6fe28"></a>

## tenant property — default_pool.origin_servers.private_ip.site_locator.virtual_site / 6183abeb8f48 / 6

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

<a id="canonical-bbac08948160bf8740b4dc62cc62f6ec430c43c4107ef3ed409c27d0f97f0427"></a>

## Next pages — default_pool.origin_servers.private_ip.site_locator.virtual_site / 6183abeb8f48 / 7

- [default_pool.origin_servers.private_ip.site_locator](data-sources--http_loadbalancer--reference--group-015.md#canonical-1808bea5ce681ad3cdeb2595342ab193ebbbbfe62f172282c5236ffa8b1c02d5)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c15d4b52c33c2964a14aef16cd11d3e8fe77e59cb9bb4126369969e6e68d5d42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d29c58e1cc4e7fca00df29bdd1f5149fe499affbf594f287a304f9258690f847"></a>

## default_pool.origin_servers.private_ip.snat_pool — default_pool.origin_servers.private_ip.snat_pool / c7aa24651a5a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- default_pool.origin_servers.private_ip.snat_pool

<a id="canonical-894559559ad9a7f9c1a8fb98973c1c7089070c225de3c3eec5cf7425104ac47b"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

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

<a id="canonical-e1e6c5fc00c925d2538baa9a553443283af459d6f576c819e08cf3d79128efa7"></a>

## Direct properties — default_pool.origin_servers.private_ip.snat_pool / c7aa24651a5a / 3

- [no_snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-68d0ad3041136d416e2b98284cafd285fea04b89f482030ba5a39de6362fa3d5): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-704d089c235633909fdc858087d2b5c042690c8026ebec9f94436a6aa1a70c36): complete subsection reference.

<a id="canonical-d929155c0a6b55258d57b7debc11982884c6468c198a97ebf2d4a1bf2e8a6b17"></a>

## Next pages — default_pool.origin_servers.private_ip.snat_pool / c7aa24651a5a / 4

- [default_pool.origin_servers.private_ip.snat_pool.no_snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-68d0ad3041136d416e2b98284cafd285fea04b89f482030ba5a39de6362fa3d5)
- [default_pool.origin_servers.private_ip.snat_pool.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-704d089c235633909fdc858087d2b5c042690c8026ebec9f94436a6aa1a70c36)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-68d0ad3041136d416e2b98284cafd285fea04b89f482030ba5a39de6362fa3d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38de99bb6c7f32c4725043993f87bc8c0c87555adb9f9aa7a6c7b5fe5f270748"></a>

## default_pool.origin_servers.private_ip.snat_pool.no_snat_pool — default_pool.origin_servers.private_ip.snat_pool.no_snat_pool / aaf8f017a34e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- [default_pool.origin_servers.private_ip.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-c15d4b52c33c2964a14aef16cd11d3e8fe77e59cb9bb4126369969e6e68d5d42)
- default_pool.origin_servers.private_ip.snat_pool.no_snat_pool

<a id="canonical-51119117c54d3c4334e35e25a55f42785924e15d83f1b93aff438822ef0e9ff2"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-fc6a0bd15df482217e5881ee6858b070ee5e0cab470ed563bf674a99fdbafc5f"></a>

## Direct properties — default_pool.origin_servers.private_ip.snat_pool.no_snat_pool / aaf8f017a34e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6736a4db4b2ecd8d975a8573ae7eb95e3933c3b91bbcd2c58698cb6494f298ad"></a>

## Next pages — default_pool.origin_servers.private_ip.snat_pool.no_snat_pool / aaf8f017a34e / 4

- [default_pool.origin_servers.private_ip.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-c15d4b52c33c2964a14aef16cd11d3e8fe77e59cb9bb4126369969e6e68d5d42)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-704d089c235633909fdc858087d2b5c042690c8026ebec9f94436a6aa1a70c36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4d341a27ff2b038ff7cf73aa3333ba32e227683c2e65c63035d86437a35d354"></a>

## default_pool.origin_servers.private_ip.snat_pool.snat_pool — default_pool.origin_servers.private_ip.snat_pool.snat_pool / 7d9e22a82a39 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_ip](data-sources--http_loadbalancer--reference--group-015.md#canonical-053e0285bce6022ecffd3a8fa6de9584f39c8b9c73d4ad04162294d0eb44528f)
- [default_pool.origin_servers.private_ip.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-c15d4b52c33c2964a14aef16cd11d3e8fe77e59cb9bb4126369969e6e68d5d42)
- default_pool.origin_servers.private_ip.snat_pool.snat_pool

<a id="canonical-d2da9926bb313244fc69824e20f1454ed9da3049616f80f6da30f210c8315edd"></a>

Type: `"single"`. Computed.

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

<a id="canonical-c2322833aa86893d43e997ece6a14e6dcde477e6173e5b814df072ced6c6cdbd"></a>

## Direct properties — default_pool.origin_servers.private_ip.snat_pool.snat_pool / 7d9e22a82a39 / 3

<a id="canonical-a882354b6efaa3662b5d6af3fac3e7756c8cb352a6b36aede8cdf15e67be2200"></a>

<a id="canonical-343cc1b3d344830a14a0effb0e4255547ecf053e1f2100a6cf50cfe5544e2be0"></a>

## prefixes property — default_pool.origin_servers.private_ip.snat_pool.snat_pool / 7d9e22a82a39 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-dac66cc054fcdf08449ccbc3a75d63610b85de7dcc6e83a644a3aa6da0959986"></a>

## Next pages — default_pool.origin_servers.private_ip.snat_pool.snat_pool / 7d9e22a82a39 / 5

- [default_pool.origin_servers.private_ip.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-c15d4b52c33c2964a14aef16cd11d3e8fe77e59cb9bb4126369969e6e68d5d42)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d0dd69ba27e6836eb72775e8fce9b10646ccecafd2155d8fa96058e1506c055"></a>

## default_pool.origin_servers.private_name — default_pool.origin_servers.private_name / d6eafa8cba59 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- default_pool.origin_servers.private_name

<a id="canonical-4e4c1176915698d0ddb23c02ce94ea09e14d320696aca8d3c3df2afb4cac7c35"></a>

Type: `"single"`. Computed.

Specify origin server with private or public DNS name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]"
}
```

<a id="canonical-e73c69e5ed6d1e7df813fe9de103e96d410655b9135bfaf8d296605c0ee96492"></a>

## Direct properties — default_pool.origin_servers.private_name / d6eafa8cba59 / 3

<a id="canonical-7e1b719521989b542df1149667b5896ce6d311dbefe6924deeb2b8ed3d63bae9"></a>

<a id="canonical-a9a536366a38b4759fd5012813f1b07e35b14824f7bb118cbf601f1254359832"></a>

## dns_name property — default_pool.origin_servers.private_name / d6eafa8cba59 / 4

Type: `"string"`. Computed.

DNS Name. DNS Name

Upstream description:

DNS Name

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [inside_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-80098c912ec18847f98d95b74234df479f088d22126e550dfd477db102371f36): complete subsection reference.

- [outside_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-0320ca85f347b503ef9df4ab68dc51725c0ad6c2e0dc466f7f4bcc6256000467): complete subsection reference.

<a id="canonical-e4aef199dbba234f9a13ceec64c27e41c9b39816175e4baa44c9f2ea291dba8b"></a>

<a id="canonical-ab32a4dadb94798e044501144996ed9d33982e4664d9332c395349ffdca21a69"></a>

## refresh_interval property — default_pool.origin_servers.private_name / d6eafa8cba59 / 5

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

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

- [segment](data-sources--http_loadbalancer--reference--group-016.md#canonical-f8f24ec9993f0c2bfd3055979e1c7ecb68ede31ed0d10673e428fdb86f39ea45): complete subsection reference.

- [site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-34e0539ab5c7173733837865caf472792fc988865a36a8f33543abcf3172f65b): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-f42944e6328fc00ce36f0ec6eeab626ec6d5cc372d0c0ce8733659b7b9ff53fd): complete subsection reference.

<a id="canonical-00d58294ca8c00df66ee6fcfacde9a9ede9c6cd886c23e7422fb4ed2c0d3ffbe"></a>

## Next pages — default_pool.origin_servers.private_name / d6eafa8cba59 / 6

- [default_pool.origin_servers.private_name.inside_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-80098c912ec18847f98d95b74234df479f088d22126e550dfd477db102371f36)
- [default_pool.origin_servers.private_name.outside_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-0320ca85f347b503ef9df4ab68dc51725c0ad6c2e0dc466f7f4bcc6256000467)
- [default_pool.origin_servers.private_name.segment](data-sources--http_loadbalancer--reference--group-016.md#canonical-f8f24ec9993f0c2bfd3055979e1c7ecb68ede31ed0d10673e428fdb86f39ea45)
- [default_pool.origin_servers.private_name.site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-34e0539ab5c7173733837865caf472792fc988865a36a8f33543abcf3172f65b)
- [default_pool.origin_servers.private_name.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-f42944e6328fc00ce36f0ec6eeab626ec6d5cc372d0c0ce8733659b7b9ff53fd)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-80098c912ec18847f98d95b74234df479f088d22126e550dfd477db102371f36"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5893b56cd2a46285d2ea634a672a47ed19aef50ef70995f2f0688fb7e0c537cf"></a>

## default_pool.origin_servers.private_name.inside_network — default_pool.origin_servers.private_name.inside_network / cc7f2461eb1d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- default_pool.origin_servers.private_name.inside_network

<a id="canonical-bf9d73779a3ddcccba1f8703200edda513150ae41a1cc6b12deb5eb7a9931842"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-361f3215debb74ba2887fe398392cc4f3d88f15125a9c7da71efd931097093ef"></a>

## Direct properties — default_pool.origin_servers.private_name.inside_network / cc7f2461eb1d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03725f31b56e7ba630605938f381707c7d4451384a8ddb4278db3d6cc39a0a72"></a>

## Next pages — default_pool.origin_servers.private_name.inside_network / cc7f2461eb1d / 4

- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0320ca85f347b503ef9df4ab68dc51725c0ad6c2e0dc466f7f4bcc6256000467"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a756ebc11905f579a08b9ce518ade1e5bc79d2b1b9b5cb9ea373c0dbb343195"></a>

## default_pool.origin_servers.private_name.outside_network — default_pool.origin_servers.private_name.outside_network / a4e46e0c9bfb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- default_pool.origin_servers.private_name.outside_network

<a id="canonical-6dd3ca993c94e3138b96ad881bffc8c28c8fd79883b2e7e360a6252730cf8105"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-185d7f0b1f8fb28c9847a67b5f51d5f36cf4fe7acb6fb18e4c0b9d8651bb3461"></a>

## Direct properties — default_pool.origin_servers.private_name.outside_network / a4e46e0c9bfb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ca4b167b6d634e09e3b7a949ec0755b308ae82cdf1d4155f6f5b5a8da83f6e87"></a>

## Next pages — default_pool.origin_servers.private_name.outside_network / a4e46e0c9bfb / 4

- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f8f24ec9993f0c2bfd3055979e1c7ecb68ede31ed0d10673e428fdb86f39ea45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-752886bc9ec8f916e13b5bb0e5e69493a433247a2c1c7f36153d63dd6a99fbdb"></a>

## default_pool.origin_servers.private_name.segment — default_pool.origin_servers.private_name.segment / 38fda5bfc540 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- default_pool.origin_servers.private_name.segment

<a id="canonical-50c3db4befd84a0a41a425ec9226de76ef85ef067c5e3ad2a05baf5194aa3bc5"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-e0e306ae55c514aafd0432b1f270b74eefcb37aa1129cfaa46b03774693a4e05"></a>

## Direct properties — default_pool.origin_servers.private_name.segment / 38fda5bfc540 / 3

<a id="canonical-fddb29992717b4552483b9da361291d23dd6cc55cb13d19c05cea48df99c90c8"></a>

<a id="canonical-d850c51705711a0b6b61a7bc2d9d4ed0f6c20e499e66bd5c0c5efd79517ea5df"></a>

## name property — default_pool.origin_servers.private_name.segment / 38fda5bfc540 / 4

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

<a id="canonical-73180367ef8cd95c01920e8e3f5cb55518f428669cff00cc901ee0289d40de9a"></a>

<a id="canonical-192a0428bca0d46c2ca5d7eeaf96ebb427ef920b37206517585c2801c7ebb7fc"></a>

## namespace property — default_pool.origin_servers.private_name.segment / 38fda5bfc540 / 5

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

<a id="canonical-32ddaafc59ef29549dedfd1f50ba1dfa8889dc930677b777c60e4d2382f2f7b2"></a>

<a id="canonical-66b5fbaa9ba70d43d6faca5fbe06cb96a1bd5563a3f8ec1d6bead3be9c9a02e3"></a>

## tenant property — default_pool.origin_servers.private_name.segment / 38fda5bfc540 / 6

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

<a id="canonical-9e2c1ff38a5716c7a4c8ee44a8969bd2c550aa99e20f68fb5ad20bfbd3bd1d9b"></a>

## Next pages — default_pool.origin_servers.private_name.segment / 38fda5bfc540 / 7

- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-34e0539ab5c7173733837865caf472792fc988865a36a8f33543abcf3172f65b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a03cf9ec1213caa9623ef0ce249a975977febc110c3da668e1ac7757a48bda0"></a>

## default_pool.origin_servers.private_name.site_locator — default_pool.origin_servers.private_name.site_locator / 60f23d0344de / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- default_pool.origin_servers.private_name.site_locator

<a id="canonical-0ffab1b1b0f1e604d32e777b02816f1cff7c934879e1730e76bf61fec0e59677"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

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

<a id="canonical-d3ab2fcbdc850c8fbfe1fff5e078c52dfdeb354d18d215e99f6ed7aed373506c"></a>

## Direct properties — default_pool.origin_servers.private_name.site_locator / 60f23d0344de / 3

- [site](data-sources--http_loadbalancer--reference--group-016.md#canonical-a75a68ad2b5130068d09bdfb68de4c55a02ea8002851911c6d3fb4f4d7f04988): complete subsection reference.

- [virtual_site](data-sources--http_loadbalancer--reference--group-016.md#canonical-481f9ef05a8ba87840cf0d8be05811a47935debbc1953cb1787528f0841a2388): complete subsection reference.

<a id="canonical-3f11bb5ab924cc60525055c968a85f7f3081ffe774ac33a4e729e90791378f3f"></a>

## Next pages — default_pool.origin_servers.private_name.site_locator / 60f23d0344de / 4

- [default_pool.origin_servers.private_name.site_locator.site](data-sources--http_loadbalancer--reference--group-016.md#canonical-a75a68ad2b5130068d09bdfb68de4c55a02ea8002851911c6d3fb4f4d7f04988)
- [default_pool.origin_servers.private_name.site_locator.virtual_site](data-sources--http_loadbalancer--reference--group-016.md#canonical-481f9ef05a8ba87840cf0d8be05811a47935debbc1953cb1787528f0841a2388)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a75a68ad2b5130068d09bdfb68de4c55a02ea8002851911c6d3fb4f4d7f04988"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5032b1aa6f407ec7e19512609e04cff1a658983cb98dc92b8a7051c782e75102"></a>

## default_pool.origin_servers.private_name.site_locator.site — default_pool.origin_servers.private_name.site_locator.site / 8824b9c825d6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- [default_pool.origin_servers.private_name.site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-34e0539ab5c7173733837865caf472792fc988865a36a8f33543abcf3172f65b)
- default_pool.origin_servers.private_name.site_locator.site

<a id="canonical-5cd4f0ba86a8a7be624d476888caef275d4a057b53313cd42aeed0954630d14a"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-7fc56b203de29fb47f4096709ab8f67be9b2c322871bfbf963e74b708948948d"></a>

## Direct properties — default_pool.origin_servers.private_name.site_locator.site / 8824b9c825d6 / 3

<a id="canonical-445e19cfad66ab59724881a52e556f779699593316e8aaf657c9c18c78fea18f"></a>

<a id="canonical-d911353548951d870eaad919d6971e8d8b6f08e2b20c321eade069d74b87fa88"></a>

## name property — default_pool.origin_servers.private_name.site_locator.site / 8824b9c825d6 / 4

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

<a id="canonical-360b7d6e8cd25379d9124f6c0e7457fd01b20c72ac006ed64843208556ea3706"></a>

<a id="canonical-636732c6edd45c23901c32f864a2e961a2cf06ca1262ea0d2dd3eb6e8cad6691"></a>

## namespace property — default_pool.origin_servers.private_name.site_locator.site / 8824b9c825d6 / 5

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

<a id="canonical-bbf53b16f926d43c3bbd5de378a6cffe69d7e4bbc72d152259f61a38bee43cad"></a>

<a id="canonical-4418efab9dd9c7fa69544bf2395ea8522144fe4e8fdee400a624be101554247e"></a>

## tenant property — default_pool.origin_servers.private_name.site_locator.site / 8824b9c825d6 / 6

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

<a id="canonical-2d1f0ddff55ef8692b4ce4dfd5a330df6d359c7c84436ddf06555e35c1b1c0f6"></a>

## Next pages — default_pool.origin_servers.private_name.site_locator.site / 8824b9c825d6 / 7

- [default_pool.origin_servers.private_name.site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-34e0539ab5c7173733837865caf472792fc988865a36a8f33543abcf3172f65b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-481f9ef05a8ba87840cf0d8be05811a47935debbc1953cb1787528f0841a2388"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d064ed0111cd3c31e3cc7e713efb8cff711a0e40a7151ba1ddd2497685eee88"></a>

## default_pool.origin_servers.private_name.site_locator.virtual_site — default_pool.origin_servers.private_name.site_locator.virtual_site / 0154dac2e0d0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- [default_pool.origin_servers.private_name.site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-34e0539ab5c7173733837865caf472792fc988865a36a8f33543abcf3172f65b)
- default_pool.origin_servers.private_name.site_locator.virtual_site

<a id="canonical-3d0fddd0b2278ca3f6c9b2a0b9b239310aa0801f7a1dae294fbb346a64ec3336"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-e00cfd58ffce5431c5e9ec5bf5a4384cb37884ec5312cee3fe6137910221a491"></a>

## Direct properties — default_pool.origin_servers.private_name.site_locator.virtual_site / 0154dac2e0d0 / 3

<a id="canonical-25064270deabedb3fdd49f1529fad4701114c1880b7a0ba16cf0df65ff470dfb"></a>

<a id="canonical-00af58382272f8ee191f2b93d5fdccb3cf25e44b04aa28b131c51c9140d9aa65"></a>

## name property — default_pool.origin_servers.private_name.site_locator.virtual_site / 0154dac2e0d0 / 4

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

<a id="canonical-726f62c4a8a1561ea52626667cf24674015ee495ec75f22640154e469d9ba22f"></a>

<a id="canonical-ca80443944c656810017f85c858452f41cf39e06a30d798130874010ae54887a"></a>

## namespace property — default_pool.origin_servers.private_name.site_locator.virtual_site / 0154dac2e0d0 / 5

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

<a id="canonical-56b1596419b386c3e12471dc787b23c2219cda41cf3a884ade572615b0cfa85e"></a>

<a id="canonical-c27a2214e592dae31ac97148a5a07abd92c74746fb3292ae37af646e972181c5"></a>

## tenant property — default_pool.origin_servers.private_name.site_locator.virtual_site / 0154dac2e0d0 / 6

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

<a id="canonical-980879ef60e73dea942eee9c81d1ecc8ec501dca046f346d7c6fff03cea0ed99"></a>

## Next pages — default_pool.origin_servers.private_name.site_locator.virtual_site / 0154dac2e0d0 / 7

- [default_pool.origin_servers.private_name.site_locator](data-sources--http_loadbalancer--reference--group-016.md#canonical-34e0539ab5c7173733837865caf472792fc988865a36a8f33543abcf3172f65b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f42944e6328fc00ce36f0ec6eeab626ec6d5cc372d0c0ce8733659b7b9ff53fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a4b7a4b1d13fd49c325b7a1ac7eb6dab54feb66319f8a31bf2ddc80e4eddeaa"></a>

## default_pool.origin_servers.private_name.snat_pool — default_pool.origin_servers.private_name.snat_pool / 848d8bdc208a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- default_pool.origin_servers.private_name.snat_pool

<a id="canonical-08c7d1af56ba0b25f36a1af2930e4840f9049ba2075dc7b0a5718ecaddc7f8d8"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

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

<a id="canonical-112ec41e41ca4ee2d41221e3fc753edf6d83464e4a4062b88c434c62654cddc8"></a>

## Direct properties — default_pool.origin_servers.private_name.snat_pool / 848d8bdc208a / 3

- [no_snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-3a4e8e19ce6aeb7e65fa9253a1044725efbd23a07f1590a0f178e9ca63b3a9c6): complete subsection reference.

- [snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-22c2ef04f2f05bd994387478b3f78c1893201eb6a45f987a27046605163cb23d): complete subsection reference.

<a id="canonical-d46526e5adc76b5c818e6e2cbe7208176650eb92ae74baee515cbe789d23bf41"></a>

## Next pages — default_pool.origin_servers.private_name.snat_pool / 848d8bdc208a / 4

- [default_pool.origin_servers.private_name.snat_pool.no_snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-3a4e8e19ce6aeb7e65fa9253a1044725efbd23a07f1590a0f178e9ca63b3a9c6)
- [default_pool.origin_servers.private_name.snat_pool.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-22c2ef04f2f05bd994387478b3f78c1893201eb6a45f987a27046605163cb23d)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3a4e8e19ce6aeb7e65fa9253a1044725efbd23a07f1590a0f178e9ca63b3a9c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d87a6481eae24fa5d3c06af6ed0e3d6b59bc48b546c524f5aaa391749bec76d7"></a>

## default_pool.origin_servers.private_name.snat_pool.no_snat_pool — default_pool.origin_servers.private_name.snat_pool.no_snat_pool / 425dbdc51d8f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- [default_pool.origin_servers.private_name.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-f42944e6328fc00ce36f0ec6eeab626ec6d5cc372d0c0ce8733659b7b9ff53fd)
- default_pool.origin_servers.private_name.snat_pool.no_snat_pool

<a id="canonical-dcd69be322fb41fc1124fed81370db63ab84f7ec0384b38c56b1a8cbb452305f"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3734fecf4f504635ae5b13258f27c8c7220de7c9ccf216287d4784a7af4cb2f4"></a>

## Direct properties — default_pool.origin_servers.private_name.snat_pool.no_snat_pool / 425dbdc51d8f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d34363df554ca75274d835d45b1616a397f2e56f468d20a1e14d6c9509ada9f5"></a>

## Next pages — default_pool.origin_servers.private_name.snat_pool.no_snat_pool / 425dbdc51d8f / 4

- [default_pool.origin_servers.private_name.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-f42944e6328fc00ce36f0ec6eeab626ec6d5cc372d0c0ce8733659b7b9ff53fd)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-22c2ef04f2f05bd994387478b3f78c1893201eb6a45f987a27046605163cb23d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-186bf8e85f2fac90884278a0f6bf0075c2f4ed1b2c550c14b65424a0268aa553"></a>

## default_pool.origin_servers.private_name.snat_pool.snat_pool — default_pool.origin_servers.private_name.snat_pool.snat_pool / 1db90b45e8a6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-cd7baa82705dda4e08f2d6f7861184cba645eb3e65f7fbfd9d205f8f58750aaa)
- [default_pool.origin_servers.private_name.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-f42944e6328fc00ce36f0ec6eeab626ec6d5cc372d0c0ce8733659b7b9ff53fd)
- default_pool.origin_servers.private_name.snat_pool.snat_pool

<a id="canonical-703c06e666a8b375f3173bd63276f28b70ba087de9c39bc12a678c3535836d1d"></a>

Type: `"single"`. Computed.

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

<a id="canonical-ce6703c003b9aadc76296bd1712aa6ea9f0c2937e180688bc257c235cdae8a70"></a>

## Direct properties — default_pool.origin_servers.private_name.snat_pool.snat_pool / 1db90b45e8a6 / 3

<a id="canonical-3e3466633858f46eb080f4dff8449bf10f4a0d7a91265885d03f38a44ec2699c"></a>

<a id="canonical-f7bd451984dc3d28ebeba6db24a709945f218c30b0504a08c82c141abd647309"></a>

## prefixes property — default_pool.origin_servers.private_name.snat_pool.snat_pool / 1db90b45e8a6 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-46d18ecff0f0663b714142f15590942aa4f734177dd809c3da16181bb1640bba"></a>

## Next pages — default_pool.origin_servers.private_name.snat_pool.snat_pool / 1db90b45e8a6 / 5

- [default_pool.origin_servers.private_name.snat_pool](data-sources--http_loadbalancer--reference--group-016.md#canonical-f42944e6328fc00ce36f0ec6eeab626ec6d5cc372d0c0ce8733659b7b9ff53fd)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ecb9b44d81af36c8674aac3d22ff09119e51f5fc66f0d2b9426df39ca56d60b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1baa30e37312878500f414e92ae4cab5e3beb1b27b3c96f227cb3d43de88c24e"></a>

## default_pool.origin_servers.public_ip — default_pool.origin_servers.public_ip / 4cd6a1676f3c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- default_pool.origin_servers.public_ip

<a id="canonical-31a85a70adcaec0958ad198943eccdfb6f025ae621d46ac24f21841834fade32"></a>

Type: `"single"`. Computed.

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

<a id="canonical-917f779232f5888cedfc093741a49a783622e24287b7914591e9cbe36ab06e10"></a>

## Direct properties — default_pool.origin_servers.public_ip / 4cd6a1676f3c / 3

<a id="canonical-a9dfab6e5d3eb243becfe33304e43107ba19286d12abe15e27ab182210de8e67"></a>

<a id="canonical-8b96d057c3ef8f4a53f397424d5fa62c4698e624750fe08e50ee779ad8136254"></a>

## ip property — default_pool.origin_servers.public_ip / 4cd6a1676f3c / 4

Type: `"string"`. Computed.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Upstream description:

Exclusive with \[\] Public IPv4 address.

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

<a id="canonical-b62884bcf003cbad69a5466e33b95acb743c85d0aa8dbda415c16be5f790e451"></a>

## Next pages — default_pool.origin_servers.public_ip / 4cd6a1676f3c / 5

- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e7dee21ac44d2e1f9a83eb7e70c863336725b13d8bb07f428db0db52c8b898ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a060ed163a8a19fc33233e0b1a76c85cf91ad13a04faade1d7a864bf1b26952"></a>

## default_pool.origin_servers.public_name — default_pool.origin_servers.public_name / 9789e3d1b78a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- default_pool.origin_servers.public_name

<a id="canonical-0a8648885d14400448da0cf0a1d0977f89225d4de8ec5ce26d5961931a1bc5d7"></a>

Type: `"single"`. Computed.

Specify origin server with public DNS name.

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

<a id="canonical-9046767d477660532bcc06e60ff539ce25a7b8922540c44ce2fcc003ebc3dd0b"></a>

## Direct properties — default_pool.origin_servers.public_name / 9789e3d1b78a / 3

<a id="canonical-b7561aa7c04a0a046fb3d32426766bb02fe6f3aca29f8688c108d17dbc181cd9"></a>

<a id="canonical-ff1ec10f3973c497332b7393075ca53ca9b5e70aaa59cd1e645d86577ae6644a"></a>

## dns_name property — default_pool.origin_servers.public_name / 9789e3d1b78a / 4

Type: `"string"`. Computed.

DNS Name. DNS Name

Upstream description:

DNS Name

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

<a id="canonical-fb908a246d69ddf3a5285488bd5cf655cf21d6da266ae05f7a7cd61f907e343b"></a>

<a id="canonical-02ff658a05b15efb74bc36d8c4e3893576f1c86013fd94298d6b578d93d1936c"></a>

## refresh_interval property — default_pool.origin_servers.public_name / 9789e3d1b78a / 5

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

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

<a id="canonical-d5a6eca0002f4b9994dde122d585f3fc166c45158a85073fb2dab355f944bfd4"></a>

## Next pages — default_pool.origin_servers.public_name / 9789e3d1b78a / 6

- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-25ede9f0af75f70a6978df18cd382af126dd908afc4ffbf058f5808d071e3b56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9f825b5c2322acbefb1ecfc1145bac0cd65ff3cc680fc44aba4f6f34457c9f5"></a>

## default_pool.origin_servers.vn_private_ip — default_pool.origin_servers.vn_private_ip / b765eea18b91 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- default_pool.origin_servers.vn_private_ip

<a id="canonical-e0ced45adaa24610ca01a906f5de04a4be4f68bb2ffb5f0ba5d56462a510ba99"></a>

Type: `"single"`. Computed.

Specify origin server with IP on Virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_network_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-d344f1f80907c77fbf349285df65348a2588ccfed66bd0837b3560e9e02befd4"></a>

## Direct properties — default_pool.origin_servers.vn_private_ip / b765eea18b91 / 3

<a id="canonical-e5c7a4485ea8b08a2d67d789a988f87e0db05d2f825488ae50c075441599f27a"></a>

<a id="canonical-702b83ec8417da63b707913f091f5277e259ea5e92b51cfdc477db615f5f68b1"></a>

## ip property — default_pool.origin_servers.vn_private_ip / b765eea18b91 / 4

Type: `"string"`. Computed.

IPv4. Exclusive with \[\] IPv4 address.

Upstream description:

Exclusive with \[\] IPv4 address.

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

- [virtual_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-3381d31c56b1566a2cea18506c969bc81c82190de9004085ce17745d7df414a8): complete subsection reference.

<a id="canonical-a4ea3c1ec96e05067d249989ca6e5863af9b752313c7db4f59f944cb8983bba8"></a>

## Next pages — default_pool.origin_servers.vn_private_ip / b765eea18b91 / 5

- [default_pool.origin_servers.vn_private_ip.virtual_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-3381d31c56b1566a2cea18506c969bc81c82190de9004085ce17745d7df414a8)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3381d31c56b1566a2cea18506c969bc81c82190de9004085ce17745d7df414a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11df7522975e26425621505c74e0bc8f2ee9e572a90da93e1e2ba8b1f851b829"></a>

## default_pool.origin_servers.vn_private_ip.virtual_network — default_pool.origin_servers.vn_private_ip.virtual_network / f6d5782f5d32 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.vn_private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-25ede9f0af75f70a6978df18cd382af126dd908afc4ffbf058f5808d071e3b56)
- default_pool.origin_servers.vn_private_ip.virtual_network

<a id="canonical-39811f85e28d96ca12c230eb36374d774582e37e8570d66b5dd899cf9f1d7f85"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-e888dc11454fae62a1db2ea44f78d3172f3a6c7c4190e3c59146eb682b6b95a2"></a>

## Direct properties — default_pool.origin_servers.vn_private_ip.virtual_network / f6d5782f5d32 / 3

<a id="canonical-f16103a7c24c4946c4fbc66219bbf75f4601f6b89746662e65501fed02694dfd"></a>

<a id="canonical-fcc341683d568d32422ca02961361bbf3e7302872d943025aed7b12e74469e74"></a>

## name property — default_pool.origin_servers.vn_private_ip.virtual_network / f6d5782f5d32 / 4

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

<a id="canonical-0c7ba5745791b5f5866ec692006d80b394a4b70ab0ea02bee4d86979ab26ee9c"></a>

<a id="canonical-617dd87374dbb54d99dbc607737f9095b0a85836083b1517e9949ca37405db20"></a>

## namespace property — default_pool.origin_servers.vn_private_ip.virtual_network / f6d5782f5d32 / 5

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

<a id="canonical-57b7115b72182c2daf1bf54f35791c7b444ba9fb9d2da97518255428cea22023"></a>

<a id="canonical-be734308d9d3f747acdb8c3f3b0a4daa7e57b2dd071adc404e291eace105ab30"></a>

## tenant property — default_pool.origin_servers.vn_private_ip.virtual_network / f6d5782f5d32 / 6

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

<a id="canonical-d0c0c1c5cc618f6ebff239f38a06473b823d881bb8d06688ae085d7f2ceb2a37"></a>

## Next pages — default_pool.origin_servers.vn_private_ip.virtual_network / f6d5782f5d32 / 7

- [default_pool.origin_servers.vn_private_ip](data-sources--http_loadbalancer--reference--group-016.md#canonical-25ede9f0af75f70a6978df18cd382af126dd908afc4ffbf058f5808d071e3b56)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-04ae16beb927798a2b09499d45d1b859c7e1e1e6ee5ba2be614139e7bd8049f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3b1e7d2748d6c67d503225cd5cebe6ea6c87b605751e2dfdfad5dd420a71166"></a>

## default_pool.origin_servers.vn_private_name — default_pool.origin_servers.vn_private_name / 703e0e17c07e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- default_pool.origin_servers.vn_private_name

<a id="canonical-c00935cd4eb2004a003b00a5c15dd736f51c5f47eeaa32599a90f30e3d0b813a"></a>

Type: `"single"`. Computed.

Specify origin server with DNS name on Virtual Network.

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

<a id="canonical-d148a781db816cbcc8e686a2f615024e8015920eb95fadcf7f254c64d8c6f7a4"></a>

## Direct properties — default_pool.origin_servers.vn_private_name / 703e0e17c07e / 3

<a id="canonical-7028af9d4a99e92eb99f50ec16c5bb1db5e14b2c0f08ef4b857edb7594045f6c"></a>

<a id="canonical-52457cf9b945dfb3cf0de5a12d16a87cc621f9fb667e7880433f1e6af28e0934"></a>

## dns_name property — default_pool.origin_servers.vn_private_name / 703e0e17c07e / 4

Type: `"string"`. Computed.

DNS Name. DNS Name

Upstream description:

DNS Name

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [private_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-debe1da90c190abdb9750caa0ce0b60bcbe4fd00cc3982bc12a6a7212e8becaf): complete subsection reference.

<a id="canonical-5c56733a4ebd75cb33a2ae9a6971347517bfc92673b2f896e4e4a2c9e8274d82"></a>

## Next pages — default_pool.origin_servers.vn_private_name / 703e0e17c07e / 5

- [default_pool.origin_servers.vn_private_name.private_network](data-sources--http_loadbalancer--reference--group-016.md#canonical-debe1da90c190abdb9750caa0ce0b60bcbe4fd00cc3982bc12a6a7212e8becaf)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-debe1da90c190abdb9750caa0ce0b60bcbe4fd00cc3982bc12a6a7212e8becaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e509c0bae5cab8dbd18431e911f3bc3107da77c4ac0325d30e4fe8f6f206960b"></a>

## default_pool.origin_servers.vn_private_name.private_network — default_pool.origin_servers.vn_private_name.private_network / 5197ee4952f2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.origin_servers](data-sources--http_loadbalancer--reference--group-015.md#canonical-062dea120acaba1aa750451f5903f389cdfb2e127854ccab162c33f2f9cfcbc0)
- [default_pool.origin_servers.vn_private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-04ae16beb927798a2b09499d45d1b859c7e1e1e6ee5ba2be614139e7bd8049f6)
- default_pool.origin_servers.vn_private_name.private_network

<a id="canonical-2de1794526f1788d40e190ff28afa4749b010fd47c6bf54fd1adc8d55b42953e"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-ebe6b4029d76561edba9fa9b1bdcf7b4681e18cbdc00c52531c21e00832f93e5"></a>

## Direct properties — default_pool.origin_servers.vn_private_name.private_network / 5197ee4952f2 / 3

<a id="canonical-153e6ef590d074fa4b47a703e071b0cf30f726a66591bf8634a5622002eccf78"></a>

<a id="canonical-5e6db55c8b8e521dfa4b2cd3dc3a7db6130093798f1ffe86bcdf63cf3ba568da"></a>

## name property — default_pool.origin_servers.vn_private_name.private_network / 5197ee4952f2 / 4

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

<a id="canonical-1f4768ddaa437ddc5972c7f59bba22a5c108d4a4ec38f59209391217f77fa3da"></a>

<a id="canonical-ad8958ff96255885422371806c1fe5973b59d23235459b7faef1788dc25854ec"></a>

## namespace property — default_pool.origin_servers.vn_private_name.private_network / 5197ee4952f2 / 5

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

<a id="canonical-eb461f5d0de286bfd5de4fe6879c21cc7528d091c8eb0a5c1f616a03d23a8604"></a>

<a id="canonical-7ea25e3bcb9e68af43c7c35a49f82280090ef56b3c0161ca1a6d68b775f87d8b"></a>

## tenant property — default_pool.origin_servers.vn_private_name.private_network / 5197ee4952f2 / 6

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

<a id="canonical-00cc8bc0d3eb21ff6a111d88515d191a1fee5ca829125bc65b82028e566c46ba"></a>

## Next pages — default_pool.origin_servers.vn_private_name.private_network / 5197ee4952f2 / 7

- [default_pool.origin_servers.vn_private_name](data-sources--http_loadbalancer--reference--group-016.md#canonical-04ae16beb927798a2b09499d45d1b859c7e1e1e6ee5ba2be614139e7bd8049f6)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-68b742eb8998eeee0ab6a94544e8f3a137e49bf7bf3db74dc69111084a1ab5f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3316b71c5cce706d1428afa37238a9d7c9d456a9ec28d3adfe7fe070c89745b"></a>

## default_pool.same_as_endpoint_port — default_pool.same_as_endpoint_port / b85ba50f9d4f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- default_pool.same_as_endpoint_port

<a id="canonical-ae6c19583a8459df826aa25f5f8e0a161ef09c947110fa3f1ec77ddfa84713e7"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-960da5868ecc698d6bb62eb088913d258cfd903cb8e80f7baf6327156d41c7fe"></a>

## Direct properties — default_pool.same_as_endpoint_port / b85ba50f9d4f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6beee9603cb5ed750115c647dd125ef0d9d4a134d2437785e833db046784e0ad"></a>

## Next pages — default_pool.same_as_endpoint_port / b85ba50f9d4f / 4

- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-16cee54849090462c0081b0bf98320e268d1a39c839416736c74d9d455f82f06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47203dd638cc8ab6523df77a5757d82f96a4fa79b45d36bdc9132995684ccaf7"></a>

## default_pool.upstream_conn_pool_reuse_type — default_pool.upstream_conn_pool_reuse_type / 9044add8504d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- default_pool.upstream_conn_pool_reuse_type

<a id="canonical-6336289147078516b25f4cd55c25d812096b361eb5f912ca64657b5c62fe34b1"></a>

Type: `"single"`. Computed.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

<a id="canonical-e4d8fc79fdc8b38d1a53a2fe8fb4328c3ac1c6344845e5ba9416229ef0503aab"></a>

## Direct properties — default_pool.upstream_conn_pool_reuse_type / 9044add8504d / 3

- [disable_conn_pool_reuse](data-sources--http_loadbalancer--reference--group-016.md#canonical-16778b416271d63cfc4a15d26425a5b16dbfc499cb427097a4414e833a10e718): complete subsection reference.

- [enable_conn_pool_reuse](data-sources--http_loadbalancer--reference--group-016.md#canonical-2f1b2054569157c514715772024f878eee678131f20ec7477214e88b4057d0e7): complete subsection reference.

<a id="canonical-5c6499003382b5ad8d4b600a92416ab90720db697eb2f11562c93990a82c9f17"></a>

## Next pages — default_pool.upstream_conn_pool_reuse_type / 9044add8504d / 4

- [default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse](data-sources--http_loadbalancer--reference--group-016.md#canonical-16778b416271d63cfc4a15d26425a5b16dbfc499cb427097a4414e833a10e718)
- [default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse](data-sources--http_loadbalancer--reference--group-016.md#canonical-2f1b2054569157c514715772024f878eee678131f20ec7477214e88b4057d0e7)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-16778b416271d63cfc4a15d26425a5b16dbfc499cb427097a4414e833a10e718"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8673dce6fa6921df2f5e0ddae22885a23b199e3a1a0036fd46bd34c99fb9f96a"></a>

## default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse — default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse / 12f30110c162 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.upstream_conn_pool_reuse_type](data-sources--http_loadbalancer--reference--group-016.md#canonical-16cee54849090462c0081b0bf98320e268d1a39c839416736c74d9d455f82f06)
- default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-49a9ce67ec6f4c634d942570e06a39649e08bab31d9d0003d582f449489a6595"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable conn pool reuse.

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

<a id="canonical-b77f901aa7d7fa4442a6dd9ec3e202da265d1e6131809e36cd286590d980517e"></a>

## Direct properties — default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse / 12f30110c162 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-32815e73ad954c4a58003eb722ff13d5bf1554145152270b8e779b322e55785a"></a>

## Next pages — default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse / 12f30110c162 / 4

- [default_pool.upstream_conn_pool_reuse_type](data-sources--http_loadbalancer--reference--group-016.md#canonical-16cee54849090462c0081b0bf98320e268d1a39c839416736c74d9d455f82f06)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2f1b2054569157c514715772024f878eee678131f20ec7477214e88b4057d0e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba253cae041afbebcd8f3883374d3f2385007b87d86b5c99ab575207ef18662d"></a>

## default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse — default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse / 6bb701b08b3c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.upstream_conn_pool_reuse_type](data-sources--http_loadbalancer--reference--group-016.md#canonical-16cee54849090462c0081b0bf98320e268d1a39c839416736c74d9d455f82f06)
- default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-e8a5d36b98db2b705a5774d550149220a22a4e4b7000fdc5a128b146f57f1540"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable conn pool reuse.

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

<a id="canonical-4b7ec2a4282fe1b15fc69b5fe7a9e03cf2c7966e4a083d2b595821e637663034"></a>

## Direct properties — default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse / 6bb701b08b3c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9c6361336b8f2dc8605ace195bfdf362109cb0305e32e97828ec0cb33f68a93"></a>

## Next pages — default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse / 6bb701b08b3c / 4

- [default_pool.upstream_conn_pool_reuse_type](data-sources--http_loadbalancer--reference--group-016.md#canonical-16cee54849090462c0081b0bf98320e268d1a39c839416736c74d9d455f82f06)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-753406765ec542e7619cb32ff576a3cea2ef55ba052bdac6379c5d1cc5979d49"></a>

## default_pool.use_tls — default_pool.use_tls / 47ee3116f292 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- default_pool.use_tls

<a id="canonical-e5f138f271ebca8b02b214d03e841ed5de549b8e10a6eb61bdd137441a94600f"></a>

Type: `"single"`. Computed.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Upstream description:

Upstream TLS Parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

<a id="canonical-43878557bbdf8a03e03d6409516f21485b8fdfd0117fa158dd855de02f38756f"></a>

## Direct properties — default_pool.use_tls / 47ee3116f292 / 3

- [default_session_key_caching](data-sources--http_loadbalancer--reference--group-016.md#canonical-612784aca870338664dfa5c914132d7de6d3259af22b2b547568955bb55a21fd): complete subsection reference.

- [disable_session_key_caching](data-sources--http_loadbalancer--reference--group-016.md#canonical-5f96bca3c6a9df76253ea85c7731ea8b3e99e03ccc6299d716baae1af1d34363): complete subsection reference.

- [disable_sni](data-sources--http_loadbalancer--reference--group-016.md#canonical-eb759ddffb5d1447db632dd186f378d2060e2fd4e62acb27935be352b85eeeda): complete subsection reference.

<a id="canonical-220230ec04f38004c323044c2f902cda9289bb753c940adec8d646f7488fc274"></a>

<a id="canonical-b89045d5809f5a6d24a09c7107a0fd4dc0289cdedc2b1fb0f9fb2cf96bc3a04a"></a>

## max_session_keys property — default_pool.use_tls / 47ee3116f292 / 4

Type: `"number"`. Computed.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

- [no_mtls](data-sources--http_loadbalancer--reference--group-016.md#canonical-3f75f46c8bc8c669153ac3507a8e1d5e36808321e1f13482759d20caed5bdfc3): complete subsection reference.

- [skip_server_verification](data-sources--http_loadbalancer--reference--group-016.md#canonical-8f8c459a285ae240533905ea8525e94bf26f68c96b8a72cc9b207eee13314043): complete subsection reference.

<a id="canonical-f259ed05d6d3c701a5af75537f35e3772bf8abcbc046b264e8a2da166b15b535"></a>

<a id="canonical-aedd3faeaa41c017b74a6690f60bd6a00f5b60a4e7322e7271947d397f00e4b3"></a>

## sni property — default_pool.use_tls / 47ee3116f292 / 5

Type: `"string"`. Computed.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-d48f46f2225c8a0765d680ee3d9bf87353314158f5113e90a59ed9bb7842fb78): complete subsection reference.

- [use_host_header_as_sni](data-sources--http_loadbalancer--reference--group-016.md#canonical-7b30c66cec97e1edac2aed15a04792bb0e12802729e26a8cba7b08909faed4f4): complete subsection reference.

- [use_mtls](data-sources--http_loadbalancer--reference--group-016.md#canonical-3e169ba250c300a84127aa938522e0d208dfe868961386f27c15a094b537db22): complete subsection reference.

- [use_mtls_obj](data-sources--http_loadbalancer--reference--group-016.md#canonical-ab440c344da9bc74438bdbf0f8a245f84d688e11bcb53afadf30434bfaa7cd1a): complete subsection reference.

- [use_server_verification](data-sources--http_loadbalancer--reference--group-016.md#canonical-f17bbf0d149f3fd736049faae1a17b0660fbfc5021a76d772cc366d24b529f62): complete subsection reference.

- [volterra_trusted_ca](data-sources--http_loadbalancer--reference--group-016.md#canonical-f5be10e3c4a94b3a54a720fef7a6006a2da62aebad01c8f9019a3604f17a09b3): complete subsection reference.

<a id="canonical-29fd56fd13caf02830cc10a330bb0204c43f3a99b9dad9ad5cabb8b5798cd32a"></a>

## Next pages — default_pool.use_tls / 47ee3116f292 / 6

- [default_pool.use_tls.default_session_key_caching](data-sources--http_loadbalancer--reference--group-016.md#canonical-612784aca870338664dfa5c914132d7de6d3259af22b2b547568955bb55a21fd)
- [default_pool.use_tls.disable_session_key_caching](data-sources--http_loadbalancer--reference--group-016.md#canonical-5f96bca3c6a9df76253ea85c7731ea8b3e99e03ccc6299d716baae1af1d34363)
- [default_pool.use_tls.disable_sni](data-sources--http_loadbalancer--reference--group-016.md#canonical-eb759ddffb5d1447db632dd186f378d2060e2fd4e62acb27935be352b85eeeda)
- [default_pool.use_tls.no_mtls](data-sources--http_loadbalancer--reference--group-016.md#canonical-3f75f46c8bc8c669153ac3507a8e1d5e36808321e1f13482759d20caed5bdfc3)
- [default_pool.use_tls.skip_server_verification](data-sources--http_loadbalancer--reference--group-016.md#canonical-8f8c459a285ae240533905ea8525e94bf26f68c96b8a72cc9b207eee13314043)
- [default_pool.use_tls.tls_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-d48f46f2225c8a0765d680ee3d9bf87353314158f5113e90a59ed9bb7842fb78)
- [default_pool.use_tls.use_host_header_as_sni](data-sources--http_loadbalancer--reference--group-016.md#canonical-7b30c66cec97e1edac2aed15a04792bb0e12802729e26a8cba7b08909faed4f4)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-016.md#canonical-3e169ba250c300a84127aa938522e0d208dfe868961386f27c15a094b537db22)
- [default_pool.use_tls.use_mtls_obj](data-sources--http_loadbalancer--reference--group-016.md#canonical-ab440c344da9bc74438bdbf0f8a245f84d688e11bcb53afadf30434bfaa7cd1a)
- [default_pool.use_tls.use_server_verification](data-sources--http_loadbalancer--reference--group-016.md#canonical-f17bbf0d149f3fd736049faae1a17b0660fbfc5021a76d772cc366d24b529f62)
- [default_pool.use_tls.volterra_trusted_ca](data-sources--http_loadbalancer--reference--group-016.md#canonical-f5be10e3c4a94b3a54a720fef7a6006a2da62aebad01c8f9019a3604f17a09b3)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-612784aca870338664dfa5c914132d7de6d3259af22b2b547568955bb55a21fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4be90333c86cae176ebe80423e2fa6652a7ff55a1eb5379d4b0e2cc3d9fed66b"></a>

## default_pool.use_tls.default_session_key_caching — default_pool.use_tls.default_session_key_caching / 19d26631c7fd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- default_pool.use_tls.default_session_key_caching

<a id="canonical-7e6acdd53ee55b331ede35be76b6229f3a228312e3f6b7d3661b5babf7ff524e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

<a id="canonical-6a10d54e1c71146ad7cd0cae2440d81a49f05439e038b91ea1519bd49d944082"></a>

## Direct properties — default_pool.use_tls.default_session_key_caching / 19d26631c7fd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5288d9522d3fa536363748373af82ebd53dd9ffdfb84bbfe209c44184a34d770"></a>

## Next pages — default_pool.use_tls.default_session_key_caching / 19d26631c7fd / 4

- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5f96bca3c6a9df76253ea85c7731ea8b3e99e03ccc6299d716baae1af1d34363"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7807e7ca5becd676ae57a085a51b3d5e9c4667dd140d811c73a133e63c911c65"></a>

## default_pool.use_tls.disable_session_key_caching — default_pool.use_tls.disable_session_key_caching / 0ef8fb4c7b05 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- default_pool.use_tls.disable_session_key_caching

<a id="canonical-87d8ec5c2a94216b8b0d6d4fdb55ee9d634ae5a6cf8db078ba7884d5fe02800d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable session key caching.

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

<a id="canonical-ab229817b38dc85ac1ff54db99f0414212ebb16aa6911048f837008c2965650a"></a>

## Direct properties — default_pool.use_tls.disable_session_key_caching / 0ef8fb4c7b05 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d4a564c6778d37251b4d9920459d691183518d8c1b34026f6deb01cbe696102a"></a>

## Next pages — default_pool.use_tls.disable_session_key_caching / 0ef8fb4c7b05 / 4

- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-eb759ddffb5d1447db632dd186f378d2060e2fd4e62acb27935be352b85eeeda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f869e9ab299b2c69a93ad9fddc586544346ae6bc087b0d4f68c2f20722aaf67"></a>

## default_pool.use_tls.disable_sni — default_pool.use_tls.disable_sni / 2dace8dce47c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- default_pool.use_tls.disable_sni

<a id="canonical-bcd06de4cbcb923824605f38c6e16623b60c29a742cb69823767c27a57c05ccd"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable sni.

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

<a id="canonical-1ad30b4f24d75cbc9c647d05c26b3b61e50f9dcc6ee16af1f8f87d087da74437"></a>

## Direct properties — default_pool.use_tls.disable_sni / 2dace8dce47c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dd7c18f866e51f57709110fa705c66643c9e63b54b90cb9241553ad83697e1ad"></a>

## Next pages — default_pool.use_tls.disable_sni / 2dace8dce47c / 4

- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3f75f46c8bc8c669153ac3507a8e1d5e36808321e1f13482759d20caed5bdfc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04b7c42781c0cc7f8bf5c259127e5372f3dded9499e4ed02316030f666e8ba0f"></a>

## default_pool.use_tls.no_mtls — default_pool.use_tls.no_mtls / 0da08d79d6fd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- default_pool.use_tls.no_mtls

<a id="canonical-4d46afabfdffae933be32e76b5ba16499d6e6641aae7df90e762eee293ecaedf"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-a05b96ec01e2f7f6621f2b3f19712c241e5f2cd44e39aadf5108b44e1874d0ce"></a>

## Direct properties — default_pool.use_tls.no_mtls / 0da08d79d6fd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cecb106f87ae00dcf030943132e9e1690745e9cbe55f06bc63c4ff0bd2762811"></a>

## Next pages — default_pool.use_tls.no_mtls / 0da08d79d6fd / 4

- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8f8c459a285ae240533905ea8525e94bf26f68c96b8a72cc9b207eee13314043"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-221c24cbc330be7ba540ba06871e271b6b5bb5f43217952988bf2a0b1d2450e5"></a>

## default_pool.use_tls.skip_server_verification — default_pool.use_tls.skip_server_verification / 881290a3362f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- default_pool.use_tls.skip_server_verification

<a id="canonical-fa4561b10e288e55ab147d964aa50368762503da0a91ec265bf8308a698d23fc"></a>

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

<a id="canonical-c9c843297bd87d927ac1eafb470b7f6c8d46ced07ae499471cc65353bd3d4130"></a>

## Direct properties — default_pool.use_tls.skip_server_verification / 881290a3362f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e779bc90ac86fe2ee2b60dab30f38889abe021d609ca3c95f2d51aa1639f503d"></a>

## Next pages — default_pool.use_tls.skip_server_verification / 881290a3362f / 4

- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d48f46f2225c8a0765d680ee3d9bf87353314158f5113e90a59ed9bb7842fb78"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09a2cef548097c015190963dda6ed9e9b9926eabfb53718bbeebe0cb9f2cb783"></a>

## default_pool.use_tls.tls_config — default_pool.use_tls.tls_config / e9cd73737ffb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- default_pool.use_tls.tls_config

<a id="canonical-90cf721706594ae20706f0683d0528487ba615c2740021643a1e0c72c841d184"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-58622c00560c185c7d4023921977c5288c26a2dbd1dd2e8c4392339d82e3b7b0"></a>

## Direct properties — default_pool.use_tls.tls_config / e9cd73737ffb / 3

- [custom_security](data-sources--http_loadbalancer--reference--group-016.md#canonical-9783840463b615e8118789c034e5e1870058a446bc0350190c8b0c200e29debc): complete subsection reference.

- [default_security](data-sources--http_loadbalancer--reference--group-016.md#canonical-7fbf2f8fbeaf2786af9f033d33f15e6432fadbb478282622c424cba830c0b83b): complete subsection reference.

- [low_security](data-sources--http_loadbalancer--reference--group-016.md#canonical-073bd68b1aa6d7d39cd0775d27750ce5bca99a6432dcaa7df6290a3fd595cd20): complete subsection reference.

- [medium_security](data-sources--http_loadbalancer--reference--group-016.md#canonical-4eed3ab19444485162ce0ab8d24b7d0f951837c9164e90795e18542b7cb1d5e1): complete subsection reference.

<a id="canonical-94696080f5b60138784d228e95fce2d4806dc7ce6fbf0d5b0bed50defdb86d9e"></a>

## Next pages — default_pool.use_tls.tls_config / e9cd73737ffb / 4

- [default_pool.use_tls.tls_config.custom_security](data-sources--http_loadbalancer--reference--group-016.md#canonical-9783840463b615e8118789c034e5e1870058a446bc0350190c8b0c200e29debc)
- [default_pool.use_tls.tls_config.default_security](data-sources--http_loadbalancer--reference--group-016.md#canonical-7fbf2f8fbeaf2786af9f033d33f15e6432fadbb478282622c424cba830c0b83b)
- [default_pool.use_tls.tls_config.low_security](data-sources--http_loadbalancer--reference--group-016.md#canonical-073bd68b1aa6d7d39cd0775d27750ce5bca99a6432dcaa7df6290a3fd595cd20)
- [default_pool.use_tls.tls_config.medium_security](data-sources--http_loadbalancer--reference--group-016.md#canonical-4eed3ab19444485162ce0ab8d24b7d0f951837c9164e90795e18542b7cb1d5e1)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9783840463b615e8118789c034e5e1870058a446bc0350190c8b0c200e29debc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5e30d9441b7210ac24096bd7ecd3879c6c67ec885a55ce5ab6e13ab54d195b2"></a>

## default_pool.use_tls.tls_config.custom_security — default_pool.use_tls.tls_config.custom_security / ca6924aace70 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [default_pool.use_tls.tls_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-d48f46f2225c8a0765d680ee3d9bf87353314158f5113e90a59ed9bb7842fb78)
- default_pool.use_tls.tls_config.custom_security

<a id="canonical-671f788c9f8d2e6b23612d33c836ca3f14b711b03e3c5f4e1221f0b1d9108822"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-aef8b0afaeaff671ffc21063f211e56460deefbb0d7c92ff7888523439f0d0e5"></a>

## Direct properties — default_pool.use_tls.tls_config.custom_security / ca6924aace70 / 3

<a id="canonical-ec6cd9540608bad4aff20651d3e88742d5be1ea9505bdf42103c3dc31fc1ca37"></a>

<a id="canonical-901a910e0ef94f22507ba0d3ce14e8fd6896ff4bb1d12ac55a8075a8523d8460"></a>

## cipher_suites property — default_pool.use_tls.tls_config.custom_security / ca6924aace70 / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-57a06135f458c035d1b995d67fe6c3aa61b974a4bfc719e35bb0ca84a6db81ce"></a>

<a id="canonical-0fab1f4742db2fe33cbf3e00e43da2be67ed23f8ef1709c4f53f37606987da86"></a>

## max_version property — default_pool.use_tls.tls_config.custom_security / ca6924aace70 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c25cbd25649d51b43da10469dc5568e55a1ad6e38507c6756ef71fec27b0ed98"></a>

<a id="canonical-15c8207e4c9b7d242cb9e4fca6ae9df180fd69bb2c41f68179509b70e9832dd4"></a>

## min_version property — default_pool.use_tls.tls_config.custom_security / ca6924aace70 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ccfd585a0d69691096d33256bc57d2b8e4bcd1671dd3b4fab1310c4b492d4c00"></a>

## Next pages — default_pool.use_tls.tls_config.custom_security / ca6924aace70 / 7

- [default_pool.use_tls.tls_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-d48f46f2225c8a0765d680ee3d9bf87353314158f5113e90a59ed9bb7842fb78)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7fbf2f8fbeaf2786af9f033d33f15e6432fadbb478282622c424cba830c0b83b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7496286b1a02ebff9841737fa8f114de16a5a62fec78ee44f7abff5fe3d43741"></a>

## default_pool.use_tls.tls_config.default_security — default_pool.use_tls.tls_config.default_security / d7db7d686096 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [default_pool.use_tls.tls_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-d48f46f2225c8a0765d680ee3d9bf87353314158f5113e90a59ed9bb7842fb78)
- default_pool.use_tls.tls_config.default_security

<a id="canonical-e1da2f7512a5bf0cc585ba30755e86aed4253a13f25d251d5b798af821b71f44"></a>

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

<a id="canonical-a948479dd752da2312962938dc366d3d47e4dd40355a992845cfb6f8845c7522"></a>

## Direct properties — default_pool.use_tls.tls_config.default_security / d7db7d686096 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1dc3ef896c23d1184d297fddd2b50ca4414b931b187fe3424d12f8afb8e26da0"></a>

## Next pages — default_pool.use_tls.tls_config.default_security / d7db7d686096 / 4

- [default_pool.use_tls.tls_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-d48f46f2225c8a0765d680ee3d9bf87353314158f5113e90a59ed9bb7842fb78)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-073bd68b1aa6d7d39cd0775d27750ce5bca99a6432dcaa7df6290a3fd595cd20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3370a696884d4709e5dbb46299d5710c92b812427101dddd0c2745be4e8c412"></a>

## default_pool.use_tls.tls_config.low_security — default_pool.use_tls.tls_config.low_security / bfc5c6511ed2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [default_pool.use_tls.tls_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-d48f46f2225c8a0765d680ee3d9bf87353314158f5113e90a59ed9bb7842fb78)
- default_pool.use_tls.tls_config.low_security

<a id="canonical-5aa9a05c2bf277a4762a787b7505efa248b6907dd834083412622c82bd9f6fbe"></a>

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

<a id="canonical-a13e6955693b8b195631ccd3d39d322e55fe3bda0e0c945db30d621ec1d89e23"></a>

## Direct properties — default_pool.use_tls.tls_config.low_security / bfc5c6511ed2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ce4b05e3b5868a44db921642577a0631605adf86ab8715aa842aa21c15ba6561"></a>

## Next pages — default_pool.use_tls.tls_config.low_security / bfc5c6511ed2 / 4

- [default_pool.use_tls.tls_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-d48f46f2225c8a0765d680ee3d9bf87353314158f5113e90a59ed9bb7842fb78)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4eed3ab19444485162ce0ab8d24b7d0f951837c9164e90795e18542b7cb1d5e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12c55567a3a6c05496c9455958a95d597de626c09d84ca16174659d851952acf"></a>

## default_pool.use_tls.tls_config.medium_security — default_pool.use_tls.tls_config.medium_security / 597de9417def / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [default_pool.use_tls.tls_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-d48f46f2225c8a0765d680ee3d9bf87353314158f5113e90a59ed9bb7842fb78)
- default_pool.use_tls.tls_config.medium_security

<a id="canonical-750f9e49f8751dd6f4e8920bd7d4f037ab4159a9ede50f28740559f31dedc7fd"></a>

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

<a id="canonical-1d62a5ac81373176b5eedb08a75cb4708af9ecdf4eb3a055ad6b47270ce6b216"></a>

## Direct properties — default_pool.use_tls.tls_config.medium_security / 597de9417def / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d3b42a6c0b73a8427eafa933c00c25f8c6e74794ab17ec5508e33f8e8cfc7a89"></a>

## Next pages — default_pool.use_tls.tls_config.medium_security / 597de9417def / 4

- [default_pool.use_tls.tls_config](data-sources--http_loadbalancer--reference--group-016.md#canonical-d48f46f2225c8a0765d680ee3d9bf87353314158f5113e90a59ed9bb7842fb78)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7b30c66cec97e1edac2aed15a04792bb0e12802729e26a8cba7b08909faed4f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-585b62dacdeedf29300c3e66386aeefc8cd835a61a2182bde842d3c6489d7846"></a>

## default_pool.use_tls.use_host_header_as_sni — default_pool.use_tls.use_host_header_as_sni / adbf623d9095 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- default_pool.use_tls.use_host_header_as_sni

<a id="canonical-2d8f8a1f256c7e8b4e3549fe2642b1fde81da6d38ff998cfafb102b44e557c8c"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-6e4beca0bb0be6d4acd77d98edc22b2753d2528520335fc3fe01258e1f31e391"></a>

## Direct properties — default_pool.use_tls.use_host_header_as_sni / adbf623d9095 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e5275c19e4c30398e92a8a27f075dd5b6792f962c4ad527a333ae2224b7f6d38"></a>

## Next pages — default_pool.use_tls.use_host_header_as_sni / adbf623d9095 / 4

- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3e169ba250c300a84127aa938522e0d208dfe868961386f27c15a094b537db22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26923c4b393aeb78f5b67bd5007004f7b8f83fd938799cfba7d293aaa259da78"></a>

## default_pool.use_tls.use_mtls — default_pool.use_tls.use_mtls / fe2b0e736058 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- default_pool.use_tls.use_mtls

<a id="canonical-0edd59f11de2fab74d241b6a54b5c891fc3e515593238c6d073437597b63a17a"></a>

Type: `"single"`. Computed.

MTLS Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

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

<a id="canonical-7f63460d6579789f516d934d79138a2f1e77195e51d79d6f10fbc1a9716295cb"></a>

## Direct properties — default_pool.use_tls.use_mtls / fe2b0e736058 / 3

- [tls_certificates](data-sources--http_loadbalancer--reference--group-016.md#canonical-8a4204be6dc0b835193c002ae826ecd0f95dd3cf0a728b9f0bcc324d7851bdc9): complete subsection reference.

<a id="canonical-31d182302f5f9f8e1abeca798152d469ce4b2e893e331e31cbbf36123450fb11"></a>

## Next pages — default_pool.use_tls.use_mtls / fe2b0e736058 / 4

- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-016.md#canonical-8a4204be6dc0b835193c002ae826ecd0f95dd3cf0a728b9f0bcc324d7851bdc9)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8a4204be6dc0b835193c002ae826ecd0f95dd3cf0a728b9f0bcc324d7851bdc9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0959dc382204be64259b85f060a6361ce9b273cbd42ffcd8469d0805b720739d"></a>

## default_pool.use_tls.use_mtls.tls_certificates — default_pool.use_tls.use_mtls.tls_certificates / a3ccf6874bf1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-016.md#canonical-3e169ba250c300a84127aa938522e0d208dfe868961386f27c15a094b537db22)
- default_pool.use_tls.use_mtls.tls_certificates

<a id="canonical-f2d39ed8c77a9293413df2a30f5ec50594f48f5ec4c9bb568fe05304223ef894"></a>

Type: `"list"`. Computed.

MTLS Client Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1e3f753f5f6581bcca99d4efe44abdee7d6a2f94c526d791ff777b98ccb6813d"></a>

## Direct properties — default_pool.use_tls.use_mtls.tls_certificates / a3ccf6874bf1 / 3

<a id="canonical-b60a508d1db619d0353ffaa5722d001f517d91a8abe6506e2e0686b391c40db4"></a>

<a id="canonical-0985be614fc58f46c5d10feffe2f387679d5f302a7dbbebf81dd600f8f70583a"></a>

## certificate_url property — default_pool.use_tls.use_mtls.tls_certificates / a3ccf6874bf1 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](data-sources--http_loadbalancer--reference--group-016.md#canonical-8732b393b18c94964ee6c6410968b6d9d1cf7867bf8580c8e1867d703ea36983): complete subsection reference.

<a id="canonical-bebfc1b98f03d7edffe927030780492d26c169690c4ce5b038a46227690beaad"></a>

<a id="canonical-7cef77b6e89ecc3f26b121d6babb6ab493d24ccdb192fc2370116e296c0a788c"></a>

## description_spec property — default_pool.use_tls.use_mtls.tls_certificates / a3ccf6874bf1 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--http_loadbalancer--reference--group-016.md#canonical-4ac77719c9ab73f62f96a44084e7c7176be5995088830959ad3dcf6f5edd80ac): complete subsection reference.

- [private_key](data-sources--http_loadbalancer--reference--group-016.md#canonical-39b93a497467c8544cf12531015a4008692516efccacea2034e610a95a1e1c6c): complete subsection reference.

- [use_system_defaults](data-sources--http_loadbalancer--reference--group-016.md#canonical-ed418969ee4dd1f80580ec67334c5a235c50ca18c043e39013736ce3a437a688): complete subsection reference.

<a id="canonical-23b570c22f7e17170df91eeec3d6b0ee05aea4ba8978a74fa8bdcfae6f21a864"></a>

## Next pages — default_pool.use_tls.use_mtls.tls_certificates / a3ccf6874bf1 / 6

- [default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms](data-sources--http_loadbalancer--reference--group-016.md#canonical-8732b393b18c94964ee6c6410968b6d9d1cf7867bf8580c8e1867d703ea36983)
- [default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](data-sources--http_loadbalancer--reference--group-016.md#canonical-4ac77719c9ab73f62f96a44084e7c7176be5995088830959ad3dcf6f5edd80ac)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-016.md#canonical-39b93a497467c8544cf12531015a4008692516efccacea2034e610a95a1e1c6c)
- [default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults](data-sources--http_loadbalancer--reference--group-016.md#canonical-ed418969ee4dd1f80580ec67334c5a235c50ca18c043e39013736ce3a437a688)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-016.md#canonical-3e169ba250c300a84127aa938522e0d208dfe868961386f27c15a094b537db22)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8732b393b18c94964ee6c6410968b6d9d1cf7867bf8580c8e1867d703ea36983"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6baca9c5bf97b7c9e18a2682ea269199fd57006b3ec38730c2a1432efac872b0"></a>

## default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms — default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 3559fa4d7070 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-016.md#canonical-3e169ba250c300a84127aa938522e0d208dfe868961386f27c15a094b537db22)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-016.md#canonical-8a4204be6dc0b835193c002ae826ecd0f95dd3cf0a728b9f0bcc324d7851bdc9)
- default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-a9a8cbb0e9b87914ceb1739eb883b8e619b9493240555d431881642d8f88398a"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

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

<a id="canonical-41cad95703c7721d37b012c862a0b75cc5fd33b4d31ad046e8c2e42bb7f5c445"></a>

## Direct properties — default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 3559fa4d7070 / 3

<a id="canonical-1b07b45579e7c9413dc0e4c4c13074e80c4bdb687d445d38f7d0b376fd98d681"></a>

<a id="canonical-cd570999201f908633c66ad6b53db9a769ccc9f6d613df18ee39ea530bf95c53"></a>

## hash_algorithms property — default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 3559fa4d7070 / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3e346abc32ffe165d294d30f7ed558411540ae9710d10d7a1c26fa208e27e0c7"></a>

## Next pages — default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms / 3559fa4d7070 / 5

- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-016.md#canonical-8a4204be6dc0b835193c002ae826ecd0f95dd3cf0a728b9f0bcc324d7851bdc9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4ac77719c9ab73f62f96a44084e7c7176be5995088830959ad3dcf6f5edd80ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e94f36d10e4c00c6d96fb9758b8d5a60a20a88687b48933799f4c9fdc47a123d"></a>

## default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling — default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / 43c3f6bdbb53 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-016.md#canonical-3e169ba250c300a84127aa938522e0d208dfe868961386f27c15a094b537db22)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-016.md#canonical-8a4204be6dc0b835193c002ae826ecd0f95dd3cf0a728b9f0bcc324d7851bdc9)
- default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-c66d07e48bc549cf4190ede494d71ab7d638c8ef2b04f4832ec75df40ea01585"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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

<a id="canonical-21b6308528822cf7d69e4e9d88f20c61d0594ef197365d2a375061757ac3a968"></a>

## Direct properties — default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / 43c3f6bdbb53 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-30c57a80c151025a7a43a8f73fd36e40a64db8e6c1899a1274924133e440d571"></a>

## Next pages — default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling / 43c3f6bdbb53 / 4

- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-016.md#canonical-8a4204be6dc0b835193c002ae826ecd0f95dd3cf0a728b9f0bcc324d7851bdc9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-39b93a497467c8544cf12531015a4008692516efccacea2034e610a95a1e1c6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e97b7e6d89f3754dd4a6543b1fcb06bc6db961c489382670cca0f16b5824ab1"></a>

## default_pool.use_tls.use_mtls.tls_certificates.private_key — default_pool.use_tls.use_mtls.tls_certificates.private_key / 0ee660fb602a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-016.md#canonical-3e169ba250c300a84127aa938522e0d208dfe868961386f27c15a094b537db22)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-016.md#canonical-8a4204be6dc0b835193c002ae826ecd0f95dd3cf0a728b9f0bcc324d7851bdc9)
- default_pool.use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-031b05e08b9ccee62f489889a52902c4ff536fc1f7c906a24786ec5d2cd32523"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-d655bc4a33257badfe6b75426264eea367db12f97e45d358f0accd7bc70135b6"></a>

## Direct properties — default_pool.use_tls.use_mtls.tls_certificates.private_key / 0ee660fb602a / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-016.md#canonical-77b4ebbfc4bffb81fa6b9d6b076ad9578f6fe2b345a84c6a20fe9cb4922d28bb): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-016.md#canonical-cc2da59d053f115b296f8139b2cffb5f9bfb6e26e85b32b7ef4c432ffcd1bcc3): complete subsection reference.

<a id="canonical-3ec424cfaebc815f9c8a4fa5254a08e8d57a4ea4a979cddfed09cbf60855d7b8"></a>

## Next pages — default_pool.use_tls.use_mtls.tls_certificates.private_key / 0ee660fb602a / 4

- [default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-016.md#canonical-77b4ebbfc4bffb81fa6b9d6b076ad9578f6fe2b345a84c6a20fe9cb4922d28bb)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](data-sources--http_loadbalancer--reference--group-016.md#canonical-cc2da59d053f115b296f8139b2cffb5f9bfb6e26e85b32b7ef4c432ffcd1bcc3)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-016.md#canonical-8a4204be6dc0b835193c002ae826ecd0f95dd3cf0a728b9f0bcc324d7851bdc9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-77b4ebbfc4bffb81fa6b9d6b076ad9578f6fe2b345a84c6a20fe9cb4922d28bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3adfad4dff16bdc53e5bf0539cfb9b54aeb4397769c613d951023fdbcb2ce1ed"></a>

## default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info — default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 569826316955 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-016.md#canonical-3e169ba250c300a84127aa938522e0d208dfe868961386f27c15a094b537db22)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-016.md#canonical-8a4204be6dc0b835193c002ae826ecd0f95dd3cf0a728b9f0bcc324d7851bdc9)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-016.md#canonical-39b93a497467c8544cf12531015a4008692516efccacea2034e610a95a1e1c6c)
- default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-4d4827d5f9f794b7fb643fd78387f3eec182770f2f6d0445c67bd6e520d7daad"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-7802e8242dbda7710c282e9045f4d5c79ff8a9be7862626510353d41501e3851"></a>

## Direct properties — default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 569826316955 / 3

<a id="canonical-cc89999954ddcd1b937183bd479754c5b44d1a2355cb95dea8664ff849357751"></a>

<a id="canonical-ba3a8f464deb35954ea73c530a234011d0956db6268e8f1db89fc6e049e2c071"></a>

## decryption_provider property — default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 569826316955 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-ed452dfac41f10a21490da4650567c2464ae1d14bd8ed950ccb871fa46951f57"></a>

<a id="canonical-e690329e9857c1a08231cd4cc3069591090d7ff8468dc5a5671e1773382e8ebe"></a>

## location property — default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 569826316955 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3f6d4895b3bd5f52d65c7806028e8345f85d5c0b3c43719427a3148c9c19c183"></a>

<a id="canonical-97c188eef8e3764e612acaaeafa668a8b9dfffc203cb49f5522e8e04d73e32db"></a>

## store_provider property — default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 569826316955 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-c464e881d2a0f06a1f90fa6c8cb559f8a5069cd01597d6e7efd7330936850ac7"></a>

## Next pages — default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info / 569826316955 / 7

- [default_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-016.md#canonical-39b93a497467c8544cf12531015a4008692516efccacea2034e610a95a1e1c6c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-cc2da59d053f115b296f8139b2cffb5f9bfb6e26e85b32b7ef4c432ffcd1bcc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecfd657cdec0c68666ec97d9411508c7ab03e19f6113387b0213d0baef84a9e4"></a>

## default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info — default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 58cf40627b65 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-016.md#canonical-3e169ba250c300a84127aa938522e0d208dfe868961386f27c15a094b537db22)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-016.md#canonical-8a4204be6dc0b835193c002ae826ecd0f95dd3cf0a728b9f0bcc324d7851bdc9)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-016.md#canonical-39b93a497467c8544cf12531015a4008692516efccacea2034e610a95a1e1c6c)
- default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-2930027600edf0950c5e72e4f0e4c321c3fd291f8a23ab84abef336c8214bbe0"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-ee47d60314beb0d35436a420338f7234a2fc99218100c0335e4de5ec2384fc5a"></a>

## Direct properties — default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 58cf40627b65 / 3

<a id="canonical-aefa4460228f582e2db0049cbfdf948288944b6b3f4a5d5bccaf224f11597590"></a>

<a id="canonical-26b913efee75d50285f6d094130e612d793072db7b67985c1885c57510bbf925"></a>

## provider_ref property — default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 58cf40627b65 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-4239577192cbd3f2892cc00fabc99cd9bbf138b63ff029626b4961fca9db4d61"></a>

<a id="canonical-e17c5d5380682a8943c5472596819a776901184daad96732f53f159ea08faf18"></a>

## url property — default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 58cf40627b65 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1f6be91f5ebb99fa91eaf9d903cc81f1f50b36eaacbbe42c82b6a1df940bb19a"></a>

## Next pages — default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info / 58cf40627b65 / 6

- [default_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-016.md#canonical-39b93a497467c8544cf12531015a4008692516efccacea2034e610a95a1e1c6c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ed418969ee4dd1f80580ec67334c5a235c50ca18c043e39013736ce3a437a688"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5707d9f1e018a0e2803794e134bddf7da8e922c0b7d644810781b031b437ec0e"></a>

## default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults — default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults / eb042288ed37 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [default_pool.use_tls.use_mtls](data-sources--http_loadbalancer--reference--group-016.md#canonical-3e169ba250c300a84127aa938522e0d208dfe868961386f27c15a094b537db22)
- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-016.md#canonical-8a4204be6dc0b835193c002ae826ecd0f95dd3cf0a728b9f0bcc324d7851bdc9)
- default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-c1d4cef98ebcc8d2f90512369b6fb02921339bb96a160729c5a4778cbde41a55"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

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

<a id="canonical-c88b1e6d25267b9e280b0f20a72e04db6318dd4845674562270fe65ed975fdd7"></a>

## Direct properties — default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults / eb042288ed37 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-40f77ae6d081790a59261328b788f006a09fe3dc5d37c56baee706d07c3327da"></a>

## Next pages — default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults / eb042288ed37 / 4

- [default_pool.use_tls.use_mtls.tls_certificates](data-sources--http_loadbalancer--reference--group-016.md#canonical-8a4204be6dc0b835193c002ae826ecd0f95dd3cf0a728b9f0bcc324d7851bdc9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ab440c344da9bc74438bdbf0f8a245f84d688e11bcb53afadf30434bfaa7cd1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e729a5d7ee23b843c9f854ba692b78ed134aea0bfbd2dc95e501c5a962d4545"></a>

## default_pool.use_tls.use_mtls_obj — default_pool.use_tls.use_mtls_obj / bc5360325292 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- default_pool.use_tls.use_mtls_obj

<a id="canonical-c07893a851a686911310a434cd0a5bd636bb11a919afc2206a1973e946607698"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-7725ad980c7e0a0ae551d10f07ba9c574325d40ffb5224f87b1709bca218f577"></a>

## Direct properties — default_pool.use_tls.use_mtls_obj / bc5360325292 / 3

<a id="canonical-cf85ea81b96638947bb4277bb066143714fe05330c77b276cca850af2765d3b0"></a>

<a id="canonical-0348f4d93bb34a1d6310c7844e87676d40ce74974182fe145ef316f13094dd91"></a>

## name property — default_pool.use_tls.use_mtls_obj / bc5360325292 / 4

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

<a id="canonical-fc21490e1c124fb1679fdc0c62ee40df79778dd8e0554bdca9daf331131e9028"></a>

<a id="canonical-5f6d659cd52c84974f85508b1bad1f2c1f8374e66bcbdad10e79deb2cf600328"></a>

## namespace property — default_pool.use_tls.use_mtls_obj / bc5360325292 / 5

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

<a id="canonical-397077abb6ff96ef5b652386d0f80635f25c1f052e6cc418f66b7bb6a13e6e85"></a>

<a id="canonical-b9356ada0636c51204b309d12afe0939e4e52236fda6e8a413c038288d477edd"></a>

## tenant property — default_pool.use_tls.use_mtls_obj / bc5360325292 / 6

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

<a id="canonical-444def820eac55596ce03beb066ca4c331146d8479708fae3a8bef60c5647402"></a>

## Next pages — default_pool.use_tls.use_mtls_obj / bc5360325292 / 7

- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f17bbf0d149f3fd736049faae1a17b0660fbfc5021a76d772cc366d24b529f62"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5bf4063f3d97a6ab34c43a6b564531881e5f18254dfd143c83906e4cce31e70"></a>

## default_pool.use_tls.use_server_verification — default_pool.use_tls.use_server_verification / e3568f33a99d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- default_pool.use_tls.use_server_verification

<a id="canonical-52639b051e28b0fff915edb922dfbee37995661d60e2c217e40f5543b6d056af"></a>

Type: `"single"`. Computed.

Configuration parameter for use server verification.

Upstream description:

Upstream TLS Validation Context.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

<a id="canonical-e62a0bf468baa8bc6fbf5e32fb0fb2138f2eab7e4d7f8b040416c1443e88c3e5"></a>

## Direct properties — default_pool.use_tls.use_server_verification / e3568f33a99d / 3

- [trusted_ca](data-sources--http_loadbalancer--reference--group-016.md#canonical-492a117f2790707a665184a1057d2aacc097642a454a09d3043a48f4ca4574b2): complete subsection reference.

<a id="canonical-ec4dfd89ac3df100255f00d7dcab5536011c6f24f8bc3c2a63361db82f56f353"></a>

<a id="canonical-ce190229ae1858029f818de744441d2d55738476762673e7012af8f967f3b451"></a>

## trusted_ca_url property — default_pool.use_tls.use_server_verification / e3568f33a99d / 4

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-20cef4dc9a958ee5128995c6fb0274c3151902030a948ea3fc234713681dc750"></a>

## Next pages — default_pool.use_tls.use_server_verification / e3568f33a99d / 5

- [default_pool.use_tls.use_server_verification.trusted_ca](data-sources--http_loadbalancer--reference--group-016.md#canonical-492a117f2790707a665184a1057d2aacc097642a454a09d3043a48f4ca4574b2)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-492a117f2790707a665184a1057d2aacc097642a454a09d3043a48f4ca4574b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5330aefda84cef7ee4194383dbff134876c590a2affad21b2b7f02424603b0f4"></a>

## default_pool.use_tls.use_server_verification.trusted_ca — default_pool.use_tls.use_server_verification.trusted_ca / b5e564180685 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [default_pool.use_tls.use_server_verification](data-sources--http_loadbalancer--reference--group-016.md#canonical-f17bbf0d149f3fd736049faae1a17b0660fbfc5021a76d772cc366d24b529f62)
- default_pool.use_tls.use_server_verification.trusted_ca

<a id="canonical-97467c4a4aa32e32594408cb58f821e89dadc946679b63508868c8dfdbc8decb"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-683733c593b39378e258e3787de97df4658ad5e0671f80ea88bbe04cc7932a0b"></a>

## Direct properties — default_pool.use_tls.use_server_verification.trusted_ca / b5e564180685 / 3

<a id="canonical-ee314bedf6d9cf215bbd98fccf5e2a3313823f83b6f6a4b5315e494f219d1f51"></a>

<a id="canonical-2dee19a15010caf06d1923cea05e8795ddde892bfcffb11805dade4f34cbb9e7"></a>

## name property — default_pool.use_tls.use_server_verification.trusted_ca / b5e564180685 / 4

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

<a id="canonical-a2a3b72fe0ed77580b50a68bfbfb6ea62ef112ef139918ec4e3d0a0dc50e1cf3"></a>

<a id="canonical-d2bc3f079949164415850811a78e50ed1c46527e7b685742fc1afd86a5eca10d"></a>

## namespace property — default_pool.use_tls.use_server_verification.trusted_ca / b5e564180685 / 5

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

<a id="canonical-ee08e3c35684da95c298b43146af6a127c5bf48671164cca38370e5eb58b43b8"></a>

<a id="canonical-162e9658a5a708a00623b3c65538b3bdf105cc7308b4a0e8a2a4cac9df972c04"></a>

## tenant property — default_pool.use_tls.use_server_verification.trusted_ca / b5e564180685 / 6

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

<a id="canonical-7360f262f2f3c0f639b257daa7af294086d2904515437343f9c49d20f00754be"></a>

## Next pages — default_pool.use_tls.use_server_verification.trusted_ca / b5e564180685 / 7

- [default_pool.use_tls.use_server_verification](data-sources--http_loadbalancer--reference--group-016.md#canonical-f17bbf0d149f3fd736049faae1a17b0660fbfc5021a76d772cc366d24b529f62)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f5be10e3c4a94b3a54a720fef7a6006a2da62aebad01c8f9019a3604f17a09b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cecaf775bb0dc27a34d54032704c870ed3e483e2e8859b2ecc0e39ba77f78c63"></a>

## default_pool.use_tls.volterra_trusted_ca — default_pool.use_tls.volterra_trusted_ca / 6a21b08c08d9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- default_pool.use_tls.volterra_trusted_ca

<a id="canonical-189ea1c207073d82d283da3c154efd0be81535178d52f454c855a65ad9cf40cb"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra trusted ca. Defaults to \`map\[\]\`. Server applies default
when omitted.

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

<a id="canonical-d0576fd58fe4d79bb01c4ea57198dd529d370f43942fa896f9ed4a9af757874e"></a>

## Direct properties — default_pool.use_tls.volterra_trusted_ca / 6a21b08c08d9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a22cb0bf4b0df4acd4d8d79d5371413ce8ceb22a003a0b467a4546368eab5b88"></a>

## Next pages — default_pool.use_tls.volterra_trusted_ca / 6a21b08c08d9 / 4

- [default_pool.use_tls](data-sources--http_loadbalancer--reference--group-016.md#canonical-b6b59b429180642c1b66481f1d9872fa41857a587ffff84201c996bd63e6658d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-baa3be394b34d6e133fe6427378cd11e2a3f0092dddbc18782efd9521777f1bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9212b7a3f045c6406be8ae2824f78cdafdc1fffe716f33b6914bfb1a52dd525c"></a>

## default_pool.view_internal — default_pool.view_internal / f673df61fe4b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- default_pool.view_internal

<a id="canonical-6c22735727ba514d566f0bf0caa82ff40c61d822c70f69da150f016b6312d5ef"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-112b773f16fb968ec6e539153262116c319c9530ec15579acc822ad7c92bce64"></a>

## Direct properties — default_pool.view_internal / f673df61fe4b / 3

<a id="canonical-a80fd0282110127112a30ecab5248d1856e150b939b823b4bdc1505476815186"></a>

<a id="canonical-2ff208430b15d2825cb63dcf22a101b638778e0145352622682b0ebf164bddf9"></a>

## name property — default_pool.view_internal / f673df61fe4b / 4

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

<a id="canonical-eea247551b9f5160dcc638b5d6f5abce05c23e50e7ee92ca4590b31e1c9d224b"></a>

<a id="canonical-5e1ea5cf41bcc3de3a2964d4f7a47f3e1f32518d232cfe0d16e53e7dbf2a37c7"></a>

## namespace property — default_pool.view_internal / f673df61fe4b / 5

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

<a id="canonical-91f1d008ae02e8ba90c8c1bf5379941a3066ec726d1f750e732b5bbc023c8d4f"></a>

<a id="canonical-b904406148e05f56ad2ba374203d9635e803281e7e9f14a368028b7b82e8fe2e"></a>

## tenant property — default_pool.view_internal / f673df61fe4b / 6

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

<a id="canonical-df0702cd8226e18665bb59746f4e8e646c2261196995089524b7b52fde5746cd"></a>

## Next pages — default_pool.view_internal / f673df61fe4b / 7

- [default_pool](data-sources--http_loadbalancer--reference--group-014.md#canonical-7f933d13a24bf8869459e368d1f75c1a39c4ee146d0b03222363be0cbf638c3d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-98278c5aa3948605fd01fc43b6fea08aee92fa3c57a3fbe8bb93568f2c8dce49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-659fcbce5b3c7ec73349017c968690aee843bdcc5e38b7e983da55aa749110f5"></a>

## default_pool_list — default_pool_list / 398d12ff779e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- default_pool_list

<a id="canonical-5c3dedffe9b0bcb499f2f8cf2008c3c1d46d36456cce13f7c55e6b7a5fdb63ff"></a>

Type: `"single"`. Computed.

Origin Pool List Type. List of Origin Pools.

Upstream description:

List of Origin Pools.

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

<a id="canonical-8992450b61041cb68a4e41cbf551c3197d6824abac11b825229f10ac7a583063"></a>

## Direct properties — default_pool_list / 398d12ff779e / 3

- [pools](data-sources--http_loadbalancer--reference--group-016.md#canonical-44f48be9057db76965fc6c656ad8de6945056acae0f0c83686ac05e7c2a2b85c): complete subsection reference.

<a id="canonical-ae8c81f89d114f9d3997c358632881a3d58416c820821e1408b349b903c1af79"></a>

## Next pages — default_pool_list / 398d12ff779e / 4

- [default_pool_list.pools](data-sources--http_loadbalancer--reference--group-016.md#canonical-44f48be9057db76965fc6c656ad8de6945056acae0f0c83686ac05e7c2a2b85c)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-44f48be9057db76965fc6c656ad8de6945056acae0f0c83686ac05e7c2a2b85c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b455370a444948a3d626b325a20e0d0cbfad9f72a215e4a09153428b9038020d"></a>

## default_pool_list.pools — default_pool_list.pools / c04c9dfeafc6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool_list](data-sources--http_loadbalancer--reference--group-016.md#canonical-98278c5aa3948605fd01fc43b6fea08aee92fa3c57a3fbe8bb93568f2c8dce49)
- default_pool_list.pools

<a id="canonical-7b3c17557e484d7f8f3e79815cfe45c6789ed65c40d6adf28a756e665a87a93d"></a>

Type: `"list"`. Computed.

Origin Pools. List of Origin Pools.

Upstream description:

List of Origin Pools.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-6e926daf4d98e45544c6077ce0ff7c75aa1b7dae8e1de5bf8689a3805a212231"></a>

## Direct properties — default_pool_list.pools / c04c9dfeafc6 / 3

- [cluster](data-sources--http_loadbalancer--reference--group-017.md#canonical-bb4c492680a67df0207338e70442c575d5700b687a1d51308c0349b04aa15014): complete subsection reference.

- [endpoint_subsets](data-sources--http_loadbalancer--reference--group-017.md#canonical-0341a191fc6d320937cf715ebf5561c9581519a536847abe7b8d19b07bb36338): complete subsection reference.

- [pool](data-sources--http_loadbalancer--reference--group-017.md#canonical-9f8856343f149e30258c8a4eec1a7ffa6c87fbcd8ff90b82a5d90d54d7d37696): complete subsection reference.

<a id="canonical-ac254c4fcb8fc52fbd82452cd233a444d68aefb4dd502d7cd39f404018b4adf8"></a>

<a id="canonical-fe47322f7d7b29a4fa8eb859bcbf1e3f561b2613e6fafa4e05159f8b36f5fd9a"></a>

## priority property — default_pool_list.pools / c04c9dfeafc6 / 4

Type: `"number"`. Computed.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

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
    },
    "minimum": 0,
    "multipleOf": 1
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

<a id="canonical-9962217dcedf7c19bd81539282326c0b9ac0ce18fd195eb3182164b37b30d993"></a>

<a id="canonical-b8feaec61e471e0061a5c1dcb9495406df63c522facd98092a0f903b8ff147e8"></a>

## weight property — default_pool_list.pools / c04c9dfeafc6 / 5

Type: `"number"`. Computed.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
