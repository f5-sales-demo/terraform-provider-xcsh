---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-71395c03978cd2a0348c7aed76f799604c114e5740ba2694d76002465eaaeda3"></a>

## Next pages — default_pool_list.pools / c04c9dfeafc6 / 6

- [default_pool_list.pools.cluster](data-sources--http_loadbalancer--reference--group-017.md#canonical-bb4c492680a67df0207338e70442c575d5700b687a1d51308c0349b04aa15014)
- [default_pool_list.pools.endpoint_subsets](data-sources--http_loadbalancer--reference--group-017.md#canonical-0341a191fc6d320937cf715ebf5561c9581519a536847abe7b8d19b07bb36338)
- [default_pool_list.pools.pool](data-sources--http_loadbalancer--reference--group-017.md#canonical-9f8856343f149e30258c8a4eec1a7ffa6c87fbcd8ff90b82a5d90d54d7d37696)
- [default_pool_list](data-sources--http_loadbalancer--reference--group-016.md#canonical-98278c5aa3948605fd01fc43b6fea08aee92fa3c57a3fbe8bb93568f2c8dce49)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-bb4c492680a67df0207338e70442c575d5700b687a1d51308c0349b04aa15014"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21b45fe33bc99232a226c619c844dd73348493c1142161f74c23b6c24f10ffa8"></a>

## default_pool_list.pools.cluster — default_pool_list.pools.cluster / be723df70bbb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool_list](data-sources--http_loadbalancer--reference--group-016.md#canonical-98278c5aa3948605fd01fc43b6fea08aee92fa3c57a3fbe8bb93568f2c8dce49)
- [default_pool_list.pools](data-sources--http_loadbalancer--reference--group-016.md#canonical-44f48be9057db76965fc6c656ad8de6945056acae0f0c83686ac05e7c2a2b85c)
- default_pool_list.pools.cluster

<a id="canonical-2d26d07b6dc066607fa7ee7eb5adf826e63835c9384640d55c248e4a09840e5e"></a>

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

<a id="canonical-583ca119903ee9a15589b090bce4135909108c890023ced6c9628eb994995859"></a>

## Direct properties — default_pool_list.pools.cluster / be723df70bbb / 3

<a id="canonical-15e41cc883464321ca946eb8a9396c6c8f311175b5eb6ad51b1d702462c245c4"></a>

<a id="canonical-ee8d15725db7ab0edbb872af1f75a610972970883f2faf1b1797c4605f2b1df8"></a>

## name property — default_pool_list.pools.cluster / be723df70bbb / 4

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

<a id="canonical-6ad456c2a24322a9dcc7f7e93834cc4bab2c0c7989e3e80fdd6d4f74183e19df"></a>

<a id="canonical-0b4a6c2ed88b956bb55b629cbedac6e53b990a249fedf89950eb4d35c2ef3cec"></a>

## namespace property — default_pool_list.pools.cluster / be723df70bbb / 5

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

<a id="canonical-e31b5855d01fc8d47ce8f8a989da0920dea373d29cac26a6d84c3251a1d9014e"></a>

<a id="canonical-c903fdd702db6dfb911cc53c31b373c4967f5db4d9b6e54196463dcd35bb6acc"></a>

## tenant property — default_pool_list.pools.cluster / be723df70bbb / 6

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

<a id="canonical-b8d66a74893cf65b94534d914e2ad65fb823e3321e96c54930d2a5522ce23d80"></a>

## Next pages — default_pool_list.pools.cluster / be723df70bbb / 7

- [default_pool_list.pools](data-sources--http_loadbalancer--reference--group-016.md#canonical-44f48be9057db76965fc6c656ad8de6945056acae0f0c83686ac05e7c2a2b85c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0341a191fc6d320937cf715ebf5561c9581519a536847abe7b8d19b07bb36338"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45ac2650a5043e02b8111cbbcb8e7775a7e2dfaba302f1524911633d3a7a1c0e"></a>

## default_pool_list.pools.endpoint_subsets — default_pool_list.pools.endpoint_subsets / d1f91faeb318 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool_list](data-sources--http_loadbalancer--reference--group-016.md#canonical-98278c5aa3948605fd01fc43b6fea08aee92fa3c57a3fbe8bb93568f2c8dce49)
- [default_pool_list.pools](data-sources--http_loadbalancer--reference--group-016.md#canonical-44f48be9057db76965fc6c656ad8de6945056acae0f0c83686ac05e7c2a2b85c)
- default_pool_list.pools.endpoint_subsets

<a id="canonical-ff3cc92dfaeb48746569894a7c8adf8d07ff97c5ce8eb0eace1714548529ed0c"></a>

Type: `"single"`. Computed.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

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
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

<a id="canonical-215c6cc77d2a72dd82d5a68766a998d41361e6c7f22c3bd3f655dfdc3f3f7540"></a>

## Direct properties — default_pool_list.pools.endpoint_subsets / d1f91faeb318 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f40169ffc631fb7a1dbae987f9365c6a015378faf7562416806614da46b66e95"></a>

## Next pages — default_pool_list.pools.endpoint_subsets / d1f91faeb318 / 4

- [default_pool_list.pools](data-sources--http_loadbalancer--reference--group-016.md#canonical-44f48be9057db76965fc6c656ad8de6945056acae0f0c83686ac05e7c2a2b85c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9f8856343f149e30258c8a4eec1a7ffa6c87fbcd8ff90b82a5d90d54d7d37696"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51097362cafcd2147db7e9e9e179e476fb4b9c7edb3c6c57238e8a49d432544d"></a>

## default_pool_list.pools.pool — default_pool_list.pools.pool / 21ff8bed94bd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_pool_list](data-sources--http_loadbalancer--reference--group-016.md#canonical-98278c5aa3948605fd01fc43b6fea08aee92fa3c57a3fbe8bb93568f2c8dce49)
- [default_pool_list.pools](data-sources--http_loadbalancer--reference--group-016.md#canonical-44f48be9057db76965fc6c656ad8de6945056acae0f0c83686ac05e7c2a2b85c)
- default_pool_list.pools.pool

<a id="canonical-0792789dab74e8587ceb1c7a12ecd83e5ba95e92356453c78b05fb7c21ba150d"></a>

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

<a id="canonical-82f91cefb2ac59350f4b39235a9067ec67fc8219304c6c47786edc43842f7f4f"></a>

## Direct properties — default_pool_list.pools.pool / 21ff8bed94bd / 3

<a id="canonical-b6d3c0c8270f441c0b47488d52332aa2b674acbfd928dd76387ad29110f204db"></a>

<a id="canonical-739518b44038ace83e2c9b954a93696095b3b91d42eb49f55444dc335545c12c"></a>

## name property — default_pool_list.pools.pool / 21ff8bed94bd / 4

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

<a id="canonical-6452ca1dce1393218d7546a132003d4df2aa8bf6cfa1154cf011e0d6ad19b0a9"></a>

<a id="canonical-c36adb52f4df35a320140a18ee7203ed07b8768be8302714602eda18d3c6aa34"></a>

## namespace property — default_pool_list.pools.pool / 21ff8bed94bd / 5

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

<a id="canonical-e5bbaddccbfd085a2cf6959bee4397318057784e675e612947fcf54a265254c0"></a>

<a id="canonical-d895b0417ac45c18360748b8c6a4aaf0088200198c35e7cfea545a664cdd1103"></a>

## tenant property — default_pool_list.pools.pool / 21ff8bed94bd / 6

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

<a id="canonical-a0c2c187736b0fc646c9a6e0e23920530d69f3c78b207817743712491a1a0d4d"></a>

## Next pages — default_pool_list.pools.pool / 21ff8bed94bd / 7

- [default_pool_list.pools](data-sources--http_loadbalancer--reference--group-016.md#canonical-44f48be9057db76965fc6c656ad8de6945056acae0f0c83686ac05e7c2a2b85c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9be3fca6f6542504cece244b3240752dde0016a01ffb4d8c39941639d5bba0f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96c7048513d0f707f9856e7db3b7d936f294f3c18dacfd2f8b3f0e0c2ff745c7"></a>

## default_route_pools — default_route_pools / 741ae5e8b3fb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- default_route_pools

<a id="canonical-7837cff4f9d317ed698153717b362162d8023771a4e47fb19855691c5930826b"></a>

Type: `"list"`. Computed.

Origin Pools used when no route is specified (default route).

Upstream description:

Origin Pools used when no route is specified (default route)

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

<a id="canonical-f4211c6b313c7b8bfbdbcb8cf7982bd6c8661bbb74d11b9c30fc3bda30685c32"></a>

## Direct properties — default_route_pools / 741ae5e8b3fb / 3

- [cluster](data-sources--http_loadbalancer--reference--group-017.md#canonical-62930efc94bd31f3a630cde1afad040e8f003f67fccbc4f8c794b838f7f7ec64): complete subsection reference.

- [endpoint_subsets](data-sources--http_loadbalancer--reference--group-017.md#canonical-5c9e4472ccebe17152e75a09acb9700dd5985b0b4ea2980547496bd9f2c68734): complete subsection reference.

- [pool](data-sources--http_loadbalancer--reference--group-017.md#canonical-2a75f3a36861e76a35dce00e5849fe79c1a41016b215a5c3ba47353dd5cbe2a7): complete subsection reference.

<a id="canonical-51a30966cc1cfc0b7dc77794c9ca98490d9f489e1cc3f451da9c0a55bdadbce2"></a>

<a id="canonical-7f3dc8b4cebace581e6dc475341331c7b489a6930657ed4893fccc498d768481"></a>

## priority property — default_route_pools / 741ae5e8b3fb / 4

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

<a id="canonical-7e42962cbc442c0834519e19787843ba4f712cb9d6a12e7189ec14d5db115312"></a>

<a id="canonical-348cf10535c5f12bf859075ed486689b7af54a59c877a2c967179433f1ae1ba1"></a>

## weight property — default_route_pools / 741ae5e8b3fb / 5

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

<a id="canonical-d5483bf4058592b5a8964c95d9bec023485d6fbb089d0acddad62445846f6710"></a>

## Next pages — default_route_pools / 741ae5e8b3fb / 6

- [default_route_pools.cluster](data-sources--http_loadbalancer--reference--group-017.md#canonical-62930efc94bd31f3a630cde1afad040e8f003f67fccbc4f8c794b838f7f7ec64)
- [default_route_pools.endpoint_subsets](data-sources--http_loadbalancer--reference--group-017.md#canonical-5c9e4472ccebe17152e75a09acb9700dd5985b0b4ea2980547496bd9f2c68734)
- [default_route_pools.pool](data-sources--http_loadbalancer--reference--group-017.md#canonical-2a75f3a36861e76a35dce00e5849fe79c1a41016b215a5c3ba47353dd5cbe2a7)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-62930efc94bd31f3a630cde1afad040e8f003f67fccbc4f8c794b838f7f7ec64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78e28586ffca24da2da1318053db6f58046db4f079a60909da332b850d2072f7"></a>

## default_route_pools.cluster — default_route_pools.cluster / 46c9ee13e4c5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_route_pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-9be3fca6f6542504cece244b3240752dde0016a01ffb4d8c39941639d5bba0f6)
- default_route_pools.cluster

<a id="canonical-57aff22fce137674ccd37789282f8467f2b11fe8f83ddbead4b6c6ecbe894db2"></a>

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

<a id="canonical-748d44735b422ed94afdfc8c9257966d1afb9d792c561ffd0f28c2591743efa7"></a>

## Direct properties — default_route_pools.cluster / 46c9ee13e4c5 / 3

<a id="canonical-1cbfa25762049b01e1a496df812c3f8201a3759f659e7abc2a64d31cde924cee"></a>

<a id="canonical-ac3f4115ee5a6e70563c38475f87d04a9de49e1d0ba74554a2ec6f7af8e91d8f"></a>

## name property — default_route_pools.cluster / 46c9ee13e4c5 / 4

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

<a id="canonical-ac04396e12c6a5479af66b915e0a66d9ce18a117599b012a9e53e374de64a7ef"></a>

<a id="canonical-9f02c5641735aded7e6c685e2cd3e7d128c04765d693d54a5109d8ddb6fa29c5"></a>

## namespace property — default_route_pools.cluster / 46c9ee13e4c5 / 5

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

<a id="canonical-32933d68bbb97a23c9454b94f69898da771619e0120e7fadf09df147cf25ed09"></a>

<a id="canonical-f70870139e30cdcf50de95ef8c5fa72b2c917a83ecf5d40524a92aba75e34406"></a>

## tenant property — default_route_pools.cluster / 46c9ee13e4c5 / 6

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

<a id="canonical-6aeb0ff3dbe56e22f4e9344c8ff752085e6345ad7a9b1765dd018559a76b8c31"></a>

## Next pages — default_route_pools.cluster / 46c9ee13e4c5 / 7

- [default_route_pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-9be3fca6f6542504cece244b3240752dde0016a01ffb4d8c39941639d5bba0f6)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5c9e4472ccebe17152e75a09acb9700dd5985b0b4ea2980547496bd9f2c68734"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b67680499a26e5ff0ff6178f577a78a9891991e305b329f76a579481b6525b56"></a>

## default_route_pools.endpoint_subsets — default_route_pools.endpoint_subsets / d7e52c53cc04 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_route_pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-9be3fca6f6542504cece244b3240752dde0016a01ffb4d8c39941639d5bba0f6)
- default_route_pools.endpoint_subsets

<a id="canonical-56a00bdf92851b3d949e6ad9007624960949ee7771ae0e0410abb8d76f6a1bdd"></a>

Type: `"single"`. Computed.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

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
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

<a id="canonical-9de3bbf6ec18e1a1fa2f572a62961ee84a3e0812c71600f4a0805837d60ccdd7"></a>

## Direct properties — default_route_pools.endpoint_subsets / d7e52c53cc04 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9dfb83e0bbab10036844a3ddf4be98e37dae64fbdaefed7aec12533422e06566"></a>

## Next pages — default_route_pools.endpoint_subsets / d7e52c53cc04 / 4

- [default_route_pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-9be3fca6f6542504cece244b3240752dde0016a01ffb4d8c39941639d5bba0f6)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2a75f3a36861e76a35dce00e5849fe79c1a41016b215a5c3ba47353dd5cbe2a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb1326c303352f818b7af7015f2896fb2adb2860b76acb8d60cba244d3a770b8"></a>

## default_route_pools.pool — default_route_pools.pool / 37049d16a591 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [default_route_pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-9be3fca6f6542504cece244b3240752dde0016a01ffb4d8c39941639d5bba0f6)
- default_route_pools.pool

<a id="canonical-06a1ef55c1a266341870355bfbed5e52be6c299e4ce7a655c907e9cfd5ccf61a"></a>

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

<a id="canonical-3774811495e1cef41e5c337ec131a9e1fec9ead1a7b13c5ce6ce05c05a34d1ce"></a>

## Direct properties — default_route_pools.pool / 37049d16a591 / 3

<a id="canonical-a4b5a9f2faaa9d051d51513e0323604b598d695859ad0e8c9b8257e6bbf97f2b"></a>

<a id="canonical-6c5abd57c3e4d517821091a1e52b52ee1cd46ad137c865d5c12d4a968fdcca4b"></a>

## name property — default_route_pools.pool / 37049d16a591 / 4

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

<a id="canonical-dc43a910e6f927aef34ea77177add52d67c66fc25c4e02b0cc893082c1589558"></a>

<a id="canonical-7a4eb75445cec4f2bff2657d69c528c951d98f4e1051fb360a81936a0da01b19"></a>

## namespace property — default_route_pools.pool / 37049d16a591 / 5

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

<a id="canonical-2f50ef3d7e4dbfcf8dd86ed01d4fccd013f67ff181be24114984ca64023d79c9"></a>

<a id="canonical-5db0307beadc16bcbbd7bebd90d6780acd4074a837c085e5b29114d1d1a01483"></a>

## tenant property — default_route_pools.pool / 37049d16a591 / 6

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

<a id="canonical-dc40832a354386893fd62c47725a2368c6f18f1d0d1f29e00f571b7995dc2159"></a>

## Next pages — default_route_pools.pool / 37049d16a591 / 7

- [default_route_pools](data-sources--http_loadbalancer--reference--group-017.md#canonical-9be3fca6f6542504cece244b3240752dde0016a01ffb4d8c39941639d5bba0f6)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-15b893a4a3c0f97fe3a61685768548528500e3486c5d0f3aec45a9409148202b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-089c9e3fa21df344070da7da0d5017473527a4aa7ee21dd0d7bc2d6094ccad0e"></a>

## default_sensitive_data_policy — default_sensitive_data_policy / 4e9d21b901de / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- default_sensitive_data_policy

<a id="canonical-12c8d487e36ed7952d83ac56f82d20170849647d63b38326d0b3295cebaca664"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_sensitive\_data\_policy, sensitive\_data\_policy; Default:
default\_sensitive\_data\_policy\] Policy configuration for this feature. Defaults to \`map\[\]\`.
Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [default_sensitive_data_policy](data-sources--http_loadbalancer--reference--group-017.md#canonical-12c8d487e36ed7952d83ac56f82d20170849647d63b38326d0b3295cebaca664)
- [sensitive_data_policy](data-sources--http_loadbalancer--reference--group-025.md#canonical-6f1aff19cc28fb21bf8c1d5a1a8c4b73143011f8d76a167baae77b17eb1dbba2)

Select alternatives according to the provider validators above.

<a id="canonical-1613496495d6603253987d5c676569fdd182bf50079d51c4ead698fbf30de4b5"></a>

## Direct properties — default_sensitive_data_policy / 4e9d21b901de / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b091d28a43349a0a809821c4d84a24158c3eb5f3887aff9e96228fa9a274320b"></a>

## Next pages — default_sensitive_data_policy / 4e9d21b901de / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d3e1f6a451e30b0e93bc8549fd2dbeb90aafb192811d671ec70ac7e5beac5f8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76b2c47ae037c057ed44bfddcb229ad0239cfb8767b74e3d44523af49887eb7c"></a>

## disable_api_definition — disable_api_definition / b60ee9fab19b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- disable_api_definition

<a id="canonical-343427cfdfef7c824b866ce80898525f5fbf5b84e5bed9e61ebe35e1f2ae9bad"></a>

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

<a id="canonical-b3d6fbc6bce1fe60f5753accb62463aea23a8825b46c18302668a8cbc18d8697"></a>

## Direct properties — disable_api_definition / b60ee9fab19b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa64e97a63605c6015258294fd8154d333db42a5752ed0f6cfc110cb75bc63df"></a>

## Next pages — disable_api_definition / b60ee9fab19b / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-72f2d8912b0e440553040e939c1c7c570deff6b547de214884cc46b0b9e5aba1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ce2acd8d0aa9f99ca40fab2585b7d7eb1dd715f1d569dbc9f37fc52fd0d18c6"></a>

## disable_api_discovery — disable_api_discovery / ccd1b62fb5d1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- disable_api_discovery

<a id="canonical-f1473ada4094bfae92374c019cdaa0a565c30c6563e9de6db3f348df6aaf9569"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_api\_discovery, enable\_api\_discovery; Default: disable\_api\_discovery\] Enable
this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [disable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-f1473ada4094bfae92374c019cdaa0a565c30c6563e9de6db3f348df6aaf9569)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-7ed00ade30dae48c97488897ddb2db5a8e3bdab9e557057141257e247abd3b06)

Select alternatives according to the provider validators above.

<a id="canonical-67651f8ff9930f8001bc0e0d6ff72ca4cdf0d1188669b46ac102255fb2a76baa"></a>

## Direct properties — disable_api_discovery / ccd1b62fb5d1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-680fb0ce1bab3f4a08af020b7157d4a283190e0a2130720c06547dec13454631"></a>

## Next pages — disable_api_discovery / ccd1b62fb5d1 / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ff803d8b6937c7510caf2b002088badbec57751bcc7d1f09b82ea1227edec199"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c6bcb849abee18461270bb1cd42c47fceade1d0ab3c42122c1ea5c37b4f9080"></a>

## disable_api_testing — disable_api_testing / 0d870042223a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- disable_api_testing

<a id="canonical-bca8ea2132d6126063f053b84a89cc22d57c2e7fdb82651d031def1788f0961f"></a>

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

<a id="canonical-f03615bbf7814f74cb492560bc81a31153f5ae19fc4b1bcd2272d1f56a16f4de"></a>

## Direct properties — disable_api_testing / 0d870042223a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ab3f38b195b9262d6427bc36f7a6cb812acd6f0e1633207f028fb7bcdfce07bb"></a>

## Next pages — disable_api_testing / 0d870042223a / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-025c607a7db6cf103750aee084ffecd0bb52e176f09763341fbda5b42f40b997"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-997f57c77197058b14b9435678f29d25c5d55568c6d9e811bddffb25f9f0f654"></a>

## disable_bot_defense — disable_bot_defense / 2cdaf001bd28 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- disable_bot_defense

<a id="canonical-5c061bffc3b08eafa7c1378220ec071484387b114a62bbbf0083b591f4310159"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable bot defense. Defaults to \`map\[\]\`. Server applies default
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

<a id="canonical-05915a409f88df6aab413409768f55d5d5bb081194ad454681c6ef72233b1739"></a>

## Direct properties — disable_bot_defense / 2cdaf001bd28 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-62646f145ddedd4f4b60f779edc3f4f30fa5cfec6c2be19cee2bad43c9501e7a"></a>

## Next pages — disable_bot_defense / 2cdaf001bd28 / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-02c9534d5c7cf318702bcb048091f816fa8ecfdee8e128dd3512d341989d9953"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fd893c5eeaa9dfba4a8ca6f2c947d43cbe403b2c19b2987b1a8103c6b4cd153"></a>

## disable_caching — disable_caching / 57fe4657593c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- disable_caching

<a id="canonical-766f6acdde7728d639af6252950ccfa5735917daf33c46f87b55bc042df36218"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable caching.

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

<a id="canonical-65dbadc53ced2930d27651c2d5a23a6af1a2b31f261ce2a7d28519a697f07180"></a>

## Direct properties — disable_caching / 57fe4657593c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8b355ba73a7e833c5801591b1d6edd6e783fc352abb370831c046bef1638cfe1"></a>

## Next pages — disable_caching / 57fe4657593c / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9bf9f7383d97a540bcb94e62bc865711c55889bb695677d1cec4d59eeb47ac0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-606a6132af3163f3503e28a1418b8298258c9a63f03b879c322da16779869cec"></a>

## disable_client_side_defense — disable_client_side_defense / 0fa312ade7dd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- disable_client_side_defense

<a id="canonical-f7af9d8ade6622f5a01b3a38d83cbba5133e427ad9964c3c53288dc93afc760b"></a>

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

<a id="canonical-35261e10a8d80d4bd6ea114ed96d2c834d03f991ebf617cca1cf08c8c7958c54"></a>

## Direct properties — disable_client_side_defense / 0fa312ade7dd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c130ae3502070ec353d9848b17cbcf31ed167c7d196478f0f2fd4631bf022939"></a>

## Next pages — disable_client_side_defense / 0fa312ade7dd / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1674ba81c74af0eb269b77b9e757b1b7d36b13b46456f35ed9410bc40de849ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04072cc54f8680c5a5a95c4d3787e81eee2d752e9740058cbad9165d14f3ce10"></a>

## disable_ip_reputation — disable_ip_reputation / 34f7a4154b61 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- disable_ip_reputation

<a id="canonical-7f51dba987e1e6622a82856aaf17a84d1cde56da5c231a46940e85dd2482eff2"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_ip\_reputation, enable\_ip\_reputation; Default: disable\_ip\_reputation\] Enable
this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [disable_ip_reputation](data-sources--http_loadbalancer--reference--group-017.md#canonical-7f51dba987e1e6622a82856aaf17a84d1cde56da5c231a46940e85dd2482eff2)
- [enable_ip_reputation](data-sources--http_loadbalancer--reference--group-017.md#canonical-736d168bd3a13e81b8705b518c88767c85fde91faac9ca33aa4d6943e142b3c4)

Select alternatives according to the provider validators above.

<a id="canonical-afd905d1073e164f765f8f0d7552697b57ceeacae3aecb8078eb8f73538e08b0"></a>

## Direct properties — disable_ip_reputation / 34f7a4154b61 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-72045f088140637746a335e1254ad149168924dd81f3e19ff89bea40e4286d05"></a>

## Next pages — disable_ip_reputation / 34f7a4154b61 / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7cedce4a45d8ac191cec74f6233758f2984112f3a55123fdb66eb1ccc4b5ffbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fee20ee405e91fd92df35c6ef6fcf808b460d835d5ec5ec585a72278b6728d07"></a>

## disable_malicious_user_detection — disable_malicious_user_detection / 5564f9a8557d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- disable_malicious_user_detection

<a id="canonical-2562cb37ae09520a34c330d6a898658650054efd8dab057fa5a87b7aefae8a07"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_malicious\_user\_detection, enable\_malicious\_user\_detection; Default:
disable\_malicious\_user\_detection\] Configuration parameter for disable malicious user detection.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [disable_malicious_user_detection](data-sources--http_loadbalancer--reference--group-017.md#canonical-2562cb37ae09520a34c330d6a898658650054efd8dab057fa5a87b7aefae8a07)
- [enable_malicious_user_detection](data-sources--http_loadbalancer--reference--group-017.md#canonical-94fbd15e16a2324db4830d97005cdf7b0a55079744d95ae905991326a953aad6)

Select alternatives according to the provider validators above.

<a id="canonical-8a21839b0db49c01fa59962ba20fa9e05aea26f9d5025b39960d95fd75c29975"></a>

## Direct properties — disable_malicious_user_detection / 5564f9a8557d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5ef5a394875c552c1b0982d281fe3d8cd0a0dd191bad546afe6af6c84c6f2f75"></a>

## Next pages — disable_malicious_user_detection / 5564f9a8557d / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b57e296971fc4052df817977dfea27b735768e109893ad2b7100587e0126c8cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-984cfcf8d599a461e9095b2260d1fbe1eb69d709121af5c58ada8c9f045ce32f"></a>

## disable_malware_protection — disable_malware_protection / 784bc9dc25ce / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- disable_malware_protection

<a id="canonical-1456120d9e861f65b0e4d88668cd177b591c50e74b82401fe6e2b435e292ef47"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_malware\_protection, malware\_protection\_settings; Default:
disable\_malware\_protection\] Configuration parameter for disable malware protection. Defaults to
\`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [disable_malware_protection](data-sources--http_loadbalancer--reference--group-017.md#canonical-1456120d9e861f65b0e4d88668cd177b591c50e74b82401fe6e2b435e292ef47)
- [malware_protection_settings](data-sources--http_loadbalancer--reference--group-019.md#canonical-b47410ab141501851aaa558786ceb4c62b0f4c46e4ab173cbdceb13ea7b7da2c)

Select alternatives according to the provider validators above.

<a id="canonical-032ab89f2faf79306684a59190236290f93f106af2bcf5706ca21dbee79b74ac"></a>

## Direct properties — disable_malware_protection / 784bc9dc25ce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-33d642b41e97d30703ea6c6b3ee20a2dac00568671d1a76c2bad61f0209ed1cb"></a>

## Next pages — disable_malware_protection / 784bc9dc25ce / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-28e55b9e2b6b6493d430059ab19205160e6ac3bb8c4402689d3b95c786c22b0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c70ab7225bd1fb1e5d9668b2a42d71759abc82a928ac11e20b12267044e1290f"></a>

## disable_rate_limit — disable_rate_limit / 97071d6359fb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- disable_rate_limit

<a id="canonical-b2480f6fd6f9a92a4645c5b143d29c07d358fd24e7d4b7bd9fe593cb83a47bd3"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable rate limit. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-a1c1e03e5d91ce5f39d5c6e71d933332a3bb60117978385aa8cf84c72281e323"></a>

## Direct properties — disable_rate_limit / 97071d6359fb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c8b52f4167cb761c3406fd0768a4f02fb9b3ee2c9e53400e9e52f06370665f65"></a>

## Next pages — disable_rate_limit / 97071d6359fb / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-78bb759eccbac27147998b2f88d803fd565a71c69942e5d39ca6ed8e300eca05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b9a2830a17f7de0e64580e802ebff3311a16613818927c330bf0612c0eb5e23"></a>

## disable_threat_mesh — disable_threat_mesh / c7270fbb3bc6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- disable_threat_mesh

<a id="canonical-c17df30b8b7082b71c18a211302df85ad75ff281f9cdeabcd98af2882aa8f126"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_threat\_mesh, enable\_threat\_mesh; Default: disable\_threat\_mesh\] Enable this
option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [disable_threat_mesh](data-sources--http_loadbalancer--reference--group-017.md#canonical-c17df30b8b7082b71c18a211302df85ad75ff281f9cdeabcd98af2882aa8f126)
- [enable_threat_mesh](data-sources--http_loadbalancer--reference--group-017.md#canonical-e2092de6997caef2bee3fdeb76ccaecccc7c1280386bca74033aed25945d5326)

Select alternatives according to the provider validators above.

<a id="canonical-264596197cae85805cc4ec21d47d97bcb3c45c338ea3d502c193181334b0fe4d"></a>

## Direct properties — disable_threat_mesh / c7270fbb3bc6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a5b8bb8f76a2b9203346bd90569331bb526683b82aac71c603189c5aabfd37d4"></a>

## Next pages — disable_threat_mesh / c7270fbb3bc6 / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c79bdd7560149381f27ac6617a5e2c14f63acf9bb45638c01728fdcb0bd6a0c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efe565d55a6ff3c26b096851f5b5bf825b03896cf4d143c51daa28f06aac6eef"></a>

## disable_trust_client_ip_headers — disable_trust_client_ip_headers / 905e577cb2e0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- disable_trust_client_ip_headers

<a id="canonical-b0b8d40f1f3f8b74682af73d345e3dc6055041e45bf4a494cf1f5b9d46d5dd8f"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_trust\_client\_ip\_headers, enable\_trust\_client\_ip\_headers; Default:
disable\_trust\_client\_ip\_headers\] Enable this option. Defaults to \`map\[\]\`. Server applies
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

OneOf alternatives in this subsection:

- [disable_trust_client_ip_headers](data-sources--http_loadbalancer--reference--group-017.md#canonical-b0b8d40f1f3f8b74682af73d345e3dc6055041e45bf4a494cf1f5b9d46d5dd8f)
- [enable_trust_client_ip_headers](data-sources--http_loadbalancer--reference--group-017.md#canonical-a02a4674cca8bcf25c4a652cbbc98f54bdbd508a4850f8c22dffd62fc7666cd4)

Select alternatives according to the provider validators above.

<a id="canonical-a252184dc43afac017aad40ef6b675c3f8d39b721b1ab23ab53a32cba2404f99"></a>

## Direct properties — disable_trust_client_ip_headers / 905e577cb2e0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-78d47a019be08095100ed0daedf26e6382aa9e976fd4d53c74cf4f3b20205c97"></a>

## Next pages — disable_trust_client_ip_headers / 905e577cb2e0 / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-67933adc1825eb46b905495bf5963c44e2169f663f6b3c8ea5b9b08c6abe87c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e43a1b67844f8838d7bdf61b11a2441444e2a4aee7d53e997083cfe80e3626b3"></a>

## disable_waf — disable_waf / 0e9138794a1e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- disable_waf

<a id="canonical-b3ddcdbb722e309fe97016afb736bebfe8bf3480d4ad438b6a705f36bf04cfd7"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable waf. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-952ac6534ac62c428a03a0dde579146df4ca7d08162a0ef00574919bbff1e3af"></a>

## Direct properties — disable_waf / 0e9138794a1e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bf090eec043c24203b304f4ea3470eafaa0f74d161b66afee0834b56365ffa13"></a>

## Next pages — disable_waf / 0e9138794a1e / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-362e3ea7ff37ef9971de06d9656e93e613d193bd2a68f5912c4f7f01f05203aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f99d75e6b0087ee92d23db25b3ede5f5bee727c01186b46845db1486088b11f9"></a>

## do_not_advertise — do_not_advertise / d632160b6e17 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- do_not_advertise

<a id="canonical-09c848d6307ef2c26f82d2ad02a5882d77d806d1554fe8f59813a1da89ddc2d8"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-d03abaafd0507a3535598a74995b38b19c334fc6a5458df7d0eda7d4690c96de"></a>

## Direct properties — do_not_advertise / d632160b6e17 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a904b6005a4f7e6d6cfc0d3ef7f78b8d79ac41c7a38735c03e0c7fb10a1e7262"></a>

## Next pages — do_not_advertise / d632160b6e17 / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7811c78ad51ac7281faf926833a5a501343034e639e541d490160eec5615bd11"></a>

## enable_api_discovery — enable_api_discovery / 69d1ee7d26ce / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- enable_api_discovery

<a id="canonical-7ed00ade30dae48c97488897ddb2db5a8e3bdab9e557057141257e247abd3b06"></a>

Type: `"single"`. Computed.

Specifies the settings used for API discovery.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

<a id="canonical-a69ac84ca4246b943cbc6feb7c890c5a2a50addfc8df2d88016548d8e7372c8e"></a>

## Direct properties — enable_api_discovery / 69d1ee7d26ce / 3

- [api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-983cfe32f9f682ac478622c5f4d643e09d23f24c343113587cf9b6fff5bb7c4d): complete subsection reference.

- [api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-017.md#canonical-a3b54ff308f284af6637283b847b4c20bbcdbdcecbd8b1185e5c902d434ca3e7): complete subsection reference.

- [custom_api_auth_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-c181d980f942067057a6fdda87533a0d58442fadc86f5bc523334c7a6bf6fd61): complete subsection reference.

- [default_api_auth_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-be2e8d6d8a446947846e16ef6e63f19cc705472ef057ab10a6ab88fa3f696c14): complete subsection reference.

- [disable_learn_from_redirect_traffic](data-sources--http_loadbalancer--reference--group-017.md#canonical-c64c042e7f28e25ffa220b64c88fed26e5162008970f56d3aba71003f9d7d405): complete subsection reference.

- [discovered_api_settings](data-sources--http_loadbalancer--reference--group-017.md#canonical-9b576280bf7d48fc28af2814a23aa47a62f81020f5ef8285659d0eb86fe32b7b): complete subsection reference.

- [enable_learn_from_redirect_traffic](data-sources--http_loadbalancer--reference--group-017.md#canonical-deed01e1e44c0ee218ce7865e0a436b73764161e6d7f566c49842915cd78cf57): complete subsection reference.

<a id="canonical-780771e4a56a37a905aeeeac4275d40cbbb97fef682d8c913335974afd02daa5"></a>

## Next pages — enable_api_discovery / 69d1ee7d26ce / 4

- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-983cfe32f9f682ac478622c5f4d643e09d23f24c343113587cf9b6fff5bb7c4d)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-017.md#canonical-a3b54ff308f284af6637283b847b4c20bbcdbdcecbd8b1185e5c902d434ca3e7)
- [enable_api_discovery.custom_api_auth_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-c181d980f942067057a6fdda87533a0d58442fadc86f5bc523334c7a6bf6fd61)
- [enable_api_discovery.default_api_auth_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-be2e8d6d8a446947846e16ef6e63f19cc705472ef057ab10a6ab88fa3f696c14)
- [enable_api_discovery.disable_learn_from_redirect_traffic](data-sources--http_loadbalancer--reference--group-017.md#canonical-c64c042e7f28e25ffa220b64c88fed26e5162008970f56d3aba71003f9d7d405)
- [enable_api_discovery.discovered_api_settings](data-sources--http_loadbalancer--reference--group-017.md#canonical-9b576280bf7d48fc28af2814a23aa47a62f81020f5ef8285659d0eb86fe32b7b)
- [enable_api_discovery.enable_learn_from_redirect_traffic](data-sources--http_loadbalancer--reference--group-017.md#canonical-deed01e1e44c0ee218ce7865e0a436b73764161e6d7f566c49842915cd78cf57)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-983cfe32f9f682ac478622c5f4d643e09d23f24c343113587cf9b6fff5bb7c4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6654ce9bfd30060e4e14c6583f4ff04cd577f35063416ccf7b626ddb3650ecff"></a>

## enable_api_discovery.api_crawler — enable_api_discovery.api_crawler / 151b4eb36623 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- enable_api_discovery.api_crawler

<a id="canonical-7a01548d7f058b1c5cc14b16112ae9c5543456a4c7efcb86892ae88d75532b34"></a>

Type: `"single"`. Computed.

API Crawling. API Crawler message.

Upstream description:

API Crawler message.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

<a id="canonical-729efcc38bbbf0ba39bb80414a734956152cc4826a75882584fecc47739fead5"></a>

## Direct properties — enable_api_discovery.api_crawler / 151b4eb36623 / 3

- [api_crawler_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-f21a7a7aa7bb9207b52894bde7f43cb515dbf57d21cd715563447c44bf127e3c): complete subsection reference.

- [disable_api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-0832ff6f352fa87724efe9c74834444053b4f9b9fc88a11645a034c809937ae3): complete subsection reference.

<a id="canonical-51fdf8ef9873ccc62a30c20bb9b0d38268fca8b8e9b2a5db16d80e22a9ba1294"></a>

## Next pages — enable_api_discovery.api_crawler / 151b4eb36623 / 4

- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-f21a7a7aa7bb9207b52894bde7f43cb515dbf57d21cd715563447c44bf127e3c)
- [enable_api_discovery.api_crawler.disable_api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-0832ff6f352fa87724efe9c74834444053b4f9b9fc88a11645a034c809937ae3)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f21a7a7aa7bb9207b52894bde7f43cb515dbf57d21cd715563447c44bf127e3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e070a0fe401c8438ddc3799eb8a3145adfc29d8c66d85000a58a5714407ee39"></a>

## enable_api_discovery.api_crawler.api_crawler_config — enable_api_discovery.api_crawler.api_crawler_config / 189125c11976 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-983cfe32f9f682ac478622c5f4d643e09d23f24c343113587cf9b6fff5bb7c4d)
- enable_api_discovery.api_crawler.api_crawler_config

<a id="canonical-b48db56d60c66da71947c67daa9a53bb41135f81a717590cd0b8c52effcd2d23"></a>

Type: `"single"`. Computed.

Crawler Configure.

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

<a id="canonical-4948e10269b32f82e85b72dad5b780864fc93d667b05814114b0d23c5801e441"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config / 189125c11976 / 3

- [domains](data-sources--http_loadbalancer--reference--group-017.md#canonical-e85e23109db0e693b05636fcc4d33cb0cc7b164e19a081ea349ebd6bd0998fff): complete subsection reference.

<a id="canonical-a77dc1256a1471e9ed00a0d172f0a5fad32be0756ac7c9dda32707e7a738bc78"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config / 189125c11976 / 4

- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-017.md#canonical-e85e23109db0e693b05636fcc4d33cb0cc7b164e19a081ea349ebd6bd0998fff)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-983cfe32f9f682ac478622c5f4d643e09d23f24c343113587cf9b6fff5bb7c4d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e85e23109db0e693b05636fcc4d33cb0cc7b164e19a081ea349ebd6bd0998fff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2173cc664eccb93ce0577ccf59d578dbafb69f3655259c10706ba61e1b26733"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains — enable_api_discovery.api_crawler.api_crawler_config.domains / 3bcb59675967 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-983cfe32f9f682ac478622c5f4d643e09d23f24c343113587cf9b6fff5bb7c4d)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-f21a7a7aa7bb9207b52894bde7f43cb515dbf57d21cd715563447c44bf127e3c)
- enable_api_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-72402667c23a63382751b7d57eae70cfa3e219e5cacab7c0a0f6832f6b2dbd05"></a>

Type: `"list"`. Computed.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-844c82b77ba4b09a333f13b45d5361b8d947833abfe6f0989b5cf4aad2547a94"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains / 3bcb59675967 / 3

<a id="canonical-a717713752c6ce4db4c50b560dbd7d01b8a5b975a1f080f6dc05b63dcab7322d"></a>

<a id="canonical-447788f70ab4cec439fb00a66f4d582c634968e779781405847dda2c08e215eb"></a>

## domain property — enable_api_discovery.api_crawler.api_crawler_config.domains / 3bcb59675967 / 4

Type: `"string"`. Computed.

Select the domain to execute API Crawling with given credentials.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](data-sources--http_loadbalancer--reference--group-017.md#canonical-68e4c7f995e49b59dadaf5e573508d32bf8977a026eae5c5bfb3cff001b52e70): complete subsection reference.

<a id="canonical-c3c191be1a63777c42d6721457b6492c20377b3294c81bc7ad10a3cd0970bd45"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains / 3bcb59675967 / 5

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-017.md#canonical-68e4c7f995e49b59dadaf5e573508d32bf8977a026eae5c5bfb3cff001b52e70)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-f21a7a7aa7bb9207b52894bde7f43cb515dbf57d21cd715563447c44bf127e3c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-68e4c7f995e49b59dadaf5e573508d32bf8977a026eae5c5bfb3cff001b52e70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b011054a7e7725b9686e0d41b003a9291d976fcff0ebfd45fcacb46190bf1d66"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / 9b2a7c653f65 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-983cfe32f9f682ac478622c5f4d643e09d23f24c343113587cf9b6fff5bb7c4d)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-f21a7a7aa7bb9207b52894bde7f43cb515dbf57d21cd715563447c44bf127e3c)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-017.md#canonical-e85e23109db0e693b05636fcc4d33cb0cc7b164e19a081ea349ebd6bd0998fff)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-33b21d92d8f235be7213d4697e3f17f39d50244bb52e67a3d0d66bac747aa002"></a>

Type: `"single"`. Computed.

Configuration parameter for simple login.

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

<a id="canonical-306d12ef1820684c4d18b94711771d0885aff00fe52b0a7cbd6023d00ff25cde"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / 9b2a7c653f65 / 3

- [password](data-sources--http_loadbalancer--reference--group-017.md#canonical-0adb08b32cdbb1d098047c38d3ea839cc96641f48bf39d952879f2d574adaa95): complete subsection reference.

<a id="canonical-e337c3ef096c4ced7f506347a695231b7a3300a440224123b2457e4fb31375da"></a>

<a id="canonical-614976d66594184c2a51073c28dfc8635f48a76e74381f111a5fade2c0d403e1"></a>

## user property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / 9b2a7c653f65 / 4

Type: `"string"`. Computed.

Enter the username to assign credentials for the selected domain to crawl.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-fa1a50a2611ed90638cd9d71e8f0c85071dfa28438f0df202047041c21e79e9b"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / 9b2a7c653f65 / 5

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--http_loadbalancer--reference--group-017.md#canonical-0adb08b32cdbb1d098047c38d3ea839cc96641f48bf39d952879f2d574adaa95)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-017.md#canonical-e85e23109db0e693b05636fcc4d33cb0cc7b164e19a081ea349ebd6bd0998fff)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0adb08b32cdbb1d098047c38d3ea839cc96641f48bf39d952879f2d574adaa95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37970e5f7a5fce0cbf30c607424ea36b0f7d057c4c9bf5ba88e411137b48b347"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 51ae0f51cbac / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-983cfe32f9f682ac478622c5f4d643e09d23f24c343113587cf9b6fff5bb7c4d)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-f21a7a7aa7bb9207b52894bde7f43cb515dbf57d21cd715563447c44bf127e3c)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-017.md#canonical-e85e23109db0e693b05636fcc4d33cb0cc7b164e19a081ea349ebd6bd0998fff)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-017.md#canonical-68e4c7f995e49b59dadaf5e573508d32bf8977a026eae5c5bfb3cff001b52e70)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-65b12154109767b6770608d49a461f68a2148ae339c31974efa4cccaacfb991d"></a>

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

<a id="canonical-4a8011d4b0405f936cc94cfe91b45f550039b0b6bf05546509c79c5cf334adeb"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 51ae0f51cbac / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-017.md#canonical-ef9d87485bb6acc4856ad35df337d0b70076faa204d25be6f655c9b1a3823bf0): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-017.md#canonical-5f11867caddb7877cb744edcefae537b375333c5b6ce3c56d5a8c765f834b45d): complete subsection reference.

<a id="canonical-59d8acc3e095ffbbf36190c29ee6b4e8228898a65a3ea4e5bdb67687210f2c5a"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 51ae0f51cbac / 4

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-017.md#canonical-ef9d87485bb6acc4856ad35df337d0b70076faa204d25be6f655c9b1a3823bf0)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info](data-sources--http_loadbalancer--reference--group-017.md#canonical-5f11867caddb7877cb744edcefae537b375333c5b6ce3c56d5a8c765f834b45d)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-017.md#canonical-68e4c7f995e49b59dadaf5e573508d32bf8977a026eae5c5bfb3cff001b52e70)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ef9d87485bb6acc4856ad35df337d0b70076faa204d25be6f655c9b1a3823bf0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fcc0108717559baf39342593f87f5b18be38710a48b58bda7f67a47f3955b5b0"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / df453b0fe46c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-983cfe32f9f682ac478622c5f4d643e09d23f24c343113587cf9b6fff5bb7c4d)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-f21a7a7aa7bb9207b52894bde7f43cb515dbf57d21cd715563447c44bf127e3c)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-017.md#canonical-e85e23109db0e693b05636fcc4d33cb0cc7b164e19a081ea349ebd6bd0998fff)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-017.md#canonical-68e4c7f995e49b59dadaf5e573508d32bf8977a026eae5c5bfb3cff001b52e70)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--http_loadbalancer--reference--group-017.md#canonical-0adb08b32cdbb1d098047c38d3ea839cc96641f48bf39d952879f2d574adaa95)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-aa78dfdc9aaabe262d989fc3b98ed1edb8b6e7fdebb44fc533f83bb9fcea99e1"></a>

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

<a id="canonical-50b14e1a09ec2ec1ff285a0a0572328d9126eb12f1cb4b15b18704b214fd7f94"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / df453b0fe46c / 3

<a id="canonical-701c91ba1c5ffe3fd1832cde8f322a4d25d0db31bfe96e2853e3ecf2a09e85d5"></a>

<a id="canonical-61d64a6ebc1bd124351edec8f3acfefa7e4462a66b8613f4bdb9d4e1c552d0a4"></a>

## decryption_provider property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / df453b0fe46c / 4

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

<a id="canonical-1a6c5559ca0c5a8474e68994fda9a1421a460175171c7b7b02d706966ff21266"></a>

<a id="canonical-96ef8595bcdabcaf8b3dee7bda4427214243eb5fc48ae4f66c208a53b758e181"></a>

## location property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / df453b0fe46c / 5

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

<a id="canonical-b358833a01b48aaf6e312bca19dfc77acebc26f131cc05ebac80a26e9ee958d7"></a>

<a id="canonical-cc212487ef253afcfdde7daeacbc5b0c28f31e8e4c30797a7c8142ae213f3f2c"></a>

## store_provider property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / df453b0fe46c / 6

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

<a id="canonical-77be76d9c3bed5708369eb5a7df0d5445f10f182ad3b50295f39a4e12e859884"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / df453b0fe46c / 7

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--http_loadbalancer--reference--group-017.md#canonical-0adb08b32cdbb1d098047c38d3ea839cc96641f48bf39d952879f2d574adaa95)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5f11867caddb7877cb744edcefae537b375333c5b6ce3c56d5a8c765f834b45d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f3b6badb7d7291e52efe08ba7cf4f86d4804221e7c8c722169e30bdfa3871d2"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / d728dae12ba7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-983cfe32f9f682ac478622c5f4d643e09d23f24c343113587cf9b6fff5bb7c4d)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-017.md#canonical-f21a7a7aa7bb9207b52894bde7f43cb515dbf57d21cd715563447c44bf127e3c)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-017.md#canonical-e85e23109db0e693b05636fcc4d33cb0cc7b164e19a081ea349ebd6bd0998fff)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-017.md#canonical-68e4c7f995e49b59dadaf5e573508d32bf8977a026eae5c5bfb3cff001b52e70)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--http_loadbalancer--reference--group-017.md#canonical-0adb08b32cdbb1d098047c38d3ea839cc96641f48bf39d952879f2d574adaa95)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-96bb18a770ab0df82487a8801a944369d0be0339f399e014955bc5bfc50bed2e"></a>

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

<a id="canonical-52edd34201e5a054e124770e829fb736fe7c349b913843c8db3ec2efe52c7786"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / d728dae12ba7 / 3

<a id="canonical-3b9f123a98fdfa894f63cdbf02c041a3b2183d4e0b80d795760c67ec52c3ec59"></a>

<a id="canonical-782a79cb535213ceb4479d5046035fb66ca0106f29801ce1df8aeed28a74c75e"></a>

## provider_ref property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / d728dae12ba7 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-8b9b21578463cb9a5b4db333ebc5870532b02868e8c89c68c12d6997ba890518"></a>

<a id="canonical-74aa8bae30fac76c493dbe5e3cf8039c2fd93fa6f79104647da7a4e46f4ccaae"></a>

## url property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / d728dae12ba7 / 5

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

<a id="canonical-0ab983c2f40af8c4f61b4b4805fbd3ff9008c0a526f9ed5e37fc7442c76d9fab"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / d728dae12ba7 / 6

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--http_loadbalancer--reference--group-017.md#canonical-0adb08b32cdbb1d098047c38d3ea839cc96641f48bf39d952879f2d574adaa95)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0832ff6f352fa87724efe9c74834444053b4f9b9fc88a11645a034c809937ae3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03083dbf8eee6d54cbeb1b3f83fc0cb688818c0a3de595217dce973a7cf5a0ea"></a>

## enable_api_discovery.api_crawler.disable_api_crawler — enable_api_discovery.api_crawler.disable_api_crawler / 404c11c6f46d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-983cfe32f9f682ac478622c5f4d643e09d23f24c343113587cf9b6fff5bb7c4d)
- enable_api_discovery.api_crawler.disable_api_crawler

<a id="canonical-23a977202d76687e48f75b42d19a239f0bb161f1f60b72e46550f67f0bcedd1f"></a>

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

<a id="canonical-856339887ab6067ea4cfef0f5c61f1abbf66f0086837b034f42478b58db33636"></a>

## Direct properties — enable_api_discovery.api_crawler.disable_api_crawler / 404c11c6f46d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e2c93342cac7d08139434f6e226bb0f82975845b9d12e3f5fd9fcfb50e949fe9"></a>

## Next pages — enable_api_discovery.api_crawler.disable_api_crawler / 404c11c6f46d / 4

- [enable_api_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-017.md#canonical-983cfe32f9f682ac478622c5f4d643e09d23f24c343113587cf9b6fff5bb7c4d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a3b54ff308f284af6637283b847b4c20bbcdbdcecbd8b1185e5c902d434ca3e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4be4ff622bd349edae2db0bf7d6cfa73fe0822ba0a9f77542d6eb5a41eeb0a5"></a>

## enable_api_discovery.api_discovery_from_code_scan — enable_api_discovery.api_discovery_from_code_scan / e25bfbe87e35 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- enable_api_discovery.api_discovery_from_code_scan

<a id="canonical-91224009bf4b285b5af031165139728f04a1733251552a42bfedb306f0d59fab"></a>

Type: `"single"`. Computed.

Select Code Base and Repositories.

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

<a id="canonical-55daa2894200428e19db159904f973c93bf182189fec90c3aef23728c0ccf555"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan / e25bfbe87e35 / 3

- [code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-4c944c464703ac14b27abb9d6a864ee212b0ccef8d5d511942faf23672484820): complete subsection reference.

<a id="canonical-9e75167cb533603424603233c69c6c99f58b08a5bcac8aef4682191802b3ac64"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan / e25bfbe87e35 / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-4c944c464703ac14b27abb9d6a864ee212b0ccef8d5d511942faf23672484820)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4c944c464703ac14b27abb9d6a864ee212b0ccef8d5d511942faf23672484820"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f2905bcabb63bb021d55f1d2a04dd669190274044d37fe519166088dc82eaf9"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations / 89897590284b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-017.md#canonical-a3b54ff308f284af6637283b847b4c20bbcdbdcecbd8b1185e5c902d434ca3e7)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-06ae9da0b515cd5b96dd838f59e734b211eadf5777e7bfd9c6c2684ca9b866f2"></a>

Type: `"list"`. Computed.

Configuration parameter for code base integrations.

Upstream description:

Configuration parameter for code base integrations

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3532a75927538025a53e16008ef45309f24f7ab68d43683cd32727b88202bf68"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations / 89897590284b / 3

- [all_repos](data-sources--http_loadbalancer--reference--group-017.md#canonical-c75b076666fd6118697355b92fb043e55926a20319bdb2076a19276d43eebf00): complete subsection reference.

- [code_base_integration](data-sources--http_loadbalancer--reference--group-017.md#canonical-7e9d27bd61b7616b8f7900557e8d64b5c81882c7610292815e3d997a302ca67c): complete subsection reference.

- [selected_repos](data-sources--http_loadbalancer--reference--group-017.md#canonical-b9275539ad0a2b2ce0b86e3295d596e1afeb0659146f578c45da82c073fe0d99): complete subsection reference.

<a id="canonical-36a23d75ff0edafda022f4ce04763cee969b4fbf18d4eb599c0e428de482bca9"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations / 89897590284b / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos](data-sources--http_loadbalancer--reference--group-017.md#canonical-c75b076666fd6118697355b92fb043e55926a20319bdb2076a19276d43eebf00)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration](data-sources--http_loadbalancer--reference--group-017.md#canonical-7e9d27bd61b7616b8f7900557e8d64b5c81882c7610292815e3d997a302ca67c)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos](data-sources--http_loadbalancer--reference--group-017.md#canonical-b9275539ad0a2b2ce0b86e3295d596e1afeb0659146f578c45da82c073fe0d99)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-017.md#canonical-a3b54ff308f284af6637283b847b4c20bbcdbdcecbd8b1185e5c902d434ca3e7)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c75b076666fd6118697355b92fb043e55926a20319bdb2076a19276d43eebf00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f952e13804bde5729af2f4852f58fd16b84cdd4ef1379c8740948606bab02549"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_rep / 62bf149e054a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-017.md#canonical-a3b54ff308f284af6637283b847b4c20bbcdbdcecbd8b1185e5c902d434ca3e7)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-4c944c464703ac14b27abb9d6a864ee212b0ccef8d5d511942faf23672484820)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-dd708d5816fd4cbb756e0d6cad8462eae70960572c00840f71c4bffb57fa18c8"></a>

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

<a id="canonical-c1c94a62ea74f491448e173ddffe62816f5f71b965afd4d5fa17520ddb7ab853"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_rep / 62bf149e054a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-46fbe7d59e6c25b079fd805fede1b56dca11272bb5267477138853ccd143a57c"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_rep / 62bf149e054a / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-4c944c464703ac14b27abb9d6a864ee212b0ccef8d5d511942faf23672484820)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7e9d27bd61b7616b8f7900557e8d64b5c81882c7610292815e3d997a302ca67c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62f30e7d737dd4b968e7f5c86c813e905a8c1ae11fd910f926a91015c2ab9adc"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 63bc93335f6e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-017.md#canonical-a3b54ff308f284af6637283b847b4c20bbcdbdcecbd8b1185e5c902d434ca3e7)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-4c944c464703ac14b27abb9d6a864ee212b0ccef8d5d511942faf23672484820)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-53983ad577049507e618c8d73dce452d77bf4d16ee080b9dad57b07dc177e621"></a>

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

<a id="canonical-1b5dc744cf216d67b5865f5385041d9e8854956c6f02029cbb5fda3ecc70e8ea"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 63bc93335f6e / 3

<a id="canonical-0c5d4803748e50a9ec6bb7c1c7995eb091de5e99e5a95136cbfd2899b36cc9b4"></a>

<a id="canonical-463cb2eccdbb0efd9a43264fee1de26abd7cb04cfa8d913c89ecc9c0a73c58e0"></a>

## name property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 63bc93335f6e / 4

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

<a id="canonical-6c9a8cbde548f272954dead9062ccb0c6753c634ce219d01884d660276cf657a"></a>

<a id="canonical-911ca316eb1c08bc85a08106df8499736f14b8e2f2f21df8ea6d089cf44c3b42"></a>

## namespace property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 63bc93335f6e / 5

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

<a id="canonical-ba395bfee511940c46522b7d37187d77aa229f19ee2ee445f0793b297dee17d7"></a>

<a id="canonical-e7af6baf7fc0004aa4b7c7b8c41dd2d929dc5f508a467190f1ae32455d22887a"></a>

## tenant property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 63bc93335f6e / 6

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

<a id="canonical-82a641c6c9481dac54db283e3fd7c06d1515cc01e0f875dcb26c594cd56bd8fa"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / 63bc93335f6e / 7

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-4c944c464703ac14b27abb9d6a864ee212b0ccef8d5d511942faf23672484820)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b9275539ad0a2b2ce0b86e3295d596e1afeb0659146f578c45da82c073fe0d99"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5eefb3cfcf0e073610e8502122312cf5e55cbb7f8570b33eb0c3a2cb0cf092bd"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / 5cdccb54c3ab / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-017.md#canonical-a3b54ff308f284af6637283b847b4c20bbcdbdcecbd8b1185e5c902d434ca3e7)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-4c944c464703ac14b27abb9d6a864ee212b0ccef8d5d511942faf23672484820)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-b83b6284fd7a44fb2e6373ff17dac9b98e6434089b0f4cc45f88b3a4925b2065"></a>

Type: `"single"`. Computed.

Select which API repositories represent the LB applications.

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

<a id="canonical-9268bec5872c9f6a095bd083aec4e6e24e382375168cf7b32f87fd5d35ba11ee"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / 5cdccb54c3ab / 3

<a id="canonical-05ec68fd8ed40d3de2686af0bf53c245ba3fab632ffcc16365cdd7de0fecd184"></a>

<a id="canonical-d40598af8677fff158d4188f0fe4b9f031e34acc5f1b4c3d9520cd7909eb515e"></a>

## api_code_repo property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / 5cdccb54c3ab / 4

Type: `["list", "string"]`. Computed.

Code repository which contain API endpoints.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-76d55484995a7db979769c4cd6c5b7e58d1732a363d11e7bc26dd7ebf3234bbf"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / 5cdccb54c3ab / 5

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-4c944c464703ac14b27abb9d6a864ee212b0ccef8d5d511942faf23672484820)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c181d980f942067057a6fdda87533a0d58442fadc86f5bc523334c7a6bf6fd61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aefcd9dc7d4a4cd37314d88585f873236f23d54add6d3d727c23897938695de3"></a>

## enable_api_discovery.custom_api_auth_discovery — enable_api_discovery.custom_api_auth_discovery / 885d81c7e296 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- enable_api_discovery.custom_api_auth_discovery

<a id="canonical-cc1764e4fce1d14f0ca960699fc93f7b9788f6e803e76452bd824d43c07e618e"></a>

Type: `"single"`. Computed.

API Discovery Advanced Settings. API Discovery Advanced settings.

Upstream description:

API Discovery Advanced settings.

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

<a id="canonical-167e495ea6c0a0773c40ca84857361ce45250f279e1248fd0211e25e29dec389"></a>

## Direct properties — enable_api_discovery.custom_api_auth_discovery / 885d81c7e296 / 3

- [api_discovery_ref](data-sources--http_loadbalancer--reference--group-017.md#canonical-86aa79d33b34744f28be8fcfe9dc1d49693e77910d915d3b49b90775c273dbbb): complete subsection reference.

<a id="canonical-5c543b64e7f8b3c835ecd3a99fadb274f2f3ab031273d2d7252d127990f7a7c2"></a>

## Next pages — enable_api_discovery.custom_api_auth_discovery / 885d81c7e296 / 4

- [enable_api_discovery.custom_api_auth_discovery.api_discovery_ref](data-sources--http_loadbalancer--reference--group-017.md#canonical-86aa79d33b34744f28be8fcfe9dc1d49693e77910d915d3b49b90775c273dbbb)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-86aa79d33b34744f28be8fcfe9dc1d49693e77910d915d3b49b90775c273dbbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-815679c9f0100f2bafc76442ec14abfd28f56a217ac28aa056dd6256161cfd74"></a>

## enable_api_discovery.custom_api_auth_discovery.api_discovery_ref — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / 6e6d5a4cca82 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [enable_api_discovery.custom_api_auth_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-c181d980f942067057a6fdda87533a0d58442fadc86f5bc523334c7a6bf6fd61)
- enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-870564b0fdeabd9c4bcd58cfd61ef032d0a4aba399100ed3dbd24f8cf0f2eaf0"></a>

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

<a id="canonical-fbc0dd4aab29cbcb7bf44f0bebeabccd53bc1bafcbb355910818ee0044225ff1"></a>

## Direct properties — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / 6e6d5a4cca82 / 3

<a id="canonical-2da90844270e326981bed7f823a456507b31f7873e66a057bede495e8f7c0cf3"></a>

<a id="canonical-ecd693202edf2fb3d161180d53fe783b5415e80d98911e833a3d3f7565b6e7ed"></a>

## name property — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / 6e6d5a4cca82 / 4

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

<a id="canonical-05970f026c12b5fffed22f298075a9e494403b76ded2c28ac6c44e690a12e280"></a>

<a id="canonical-e1af939c09fc826b22e32fc75350ce8e212d4afb8545324737b87fb18961d5f1"></a>

## namespace property — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / 6e6d5a4cca82 / 5

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

<a id="canonical-a55e17a683a327985d1bc828f216512cc990f803ca9b79ac354071e77955a87b"></a>

<a id="canonical-a1414e3792b8274df021c1809207b0085e06997db99380af82958205c507b22a"></a>

## tenant property — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / 6e6d5a4cca82 / 6

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

<a id="canonical-46bbfe8851aaee996bd09f41c6a702d2c049626d0599ad2d6e7ff4f4f4727cac"></a>

## Next pages — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / 6e6d5a4cca82 / 7

- [enable_api_discovery.custom_api_auth_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-c181d980f942067057a6fdda87533a0d58442fadc86f5bc523334c7a6bf6fd61)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-be2e8d6d8a446947846e16ef6e63f19cc705472ef057ab10a6ab88fa3f696c14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e29063b0b86b7f36ab22b491e66fda36eba42cbf376c1e4f06736f37a6622ef8"></a>

## enable_api_discovery.default_api_auth_discovery — enable_api_discovery.default_api_auth_discovery / 1615de9c061d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- enable_api_discovery.default_api_auth_discovery

<a id="canonical-53fb336b8fd6f805fb651210e897fb53c83a98e156ff8965b7c75ab4ba627b40"></a>

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

<a id="canonical-6d53d2bbc33c1e83f08c182eb22f84edc1b6210bc04164a13a4ca46516c66a01"></a>

## Direct properties — enable_api_discovery.default_api_auth_discovery / 1615de9c061d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7ba52aab86451db8f4df299c524c0eb474005556cef80498f81c484c565a83e8"></a>

## Next pages — enable_api_discovery.default_api_auth_discovery / 1615de9c061d / 4

- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c64c042e7f28e25ffa220b64c88fed26e5162008970f56d3aba71003f9d7d405"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07bf5e63b16ade2993a667d7b75ce78cc61c71010184396b65c64c186c594c7a"></a>

## enable_api_discovery.disable_learn_from_redirect_traffic — enable_api_discovery.disable_learn_from_redirect_traffic / 53c11bcdb266 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- enable_api_discovery.disable_learn_from_redirect_traffic

<a id="canonical-0dbcdac45ebc8a6ec589a6d1000a61d97bd841d29aade4e06cc8300131d6920f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable learn from redirect traffic.

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

<a id="canonical-5ca2a6021cf1f6f59bc21ad53d81f854cfdd3f20d7eea047e692ac8383f98bff"></a>

## Direct properties — enable_api_discovery.disable_learn_from_redirect_traffic / 53c11bcdb266 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-79fe7a0bacaeef990ef25d8eeb7a931ab712993b75e3905f64ffe62b9db87630"></a>

## Next pages — enable_api_discovery.disable_learn_from_redirect_traffic / 53c11bcdb266 / 4

- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9b576280bf7d48fc28af2814a23aa47a62f81020f5ef8285659d0eb86fe32b7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4443b3669c6ffde9d1824af58c713f586773f55d72987916319c5eb0cbff055"></a>

## enable_api_discovery.discovered_api_settings — enable_api_discovery.discovered_api_settings / a29a07d898b3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- enable_api_discovery.discovered_api_settings

<a id="canonical-7584aef0cf4adcebc5ba00727b7f6805006e8e47f5afd86639d1755b3da2cb2b"></a>

Type: `"single"`. Computed.

Discovered API Settings. Configure Discovered API Settings.

Upstream description:

Configure Discovered API Settings.

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

<a id="canonical-1c0a9949a82f943ae56e78fbda2c745a5c58b414fc6e6d70b7d09b63c78d8597"></a>

## Direct properties — enable_api_discovery.discovered_api_settings / a29a07d898b3 / 3

<a id="canonical-c470797efc624dd1a2bcacd71437a4a754b56215a584083b7e9ea736db56e374"></a>

<a id="canonical-b5770d46a97dadd71c223bf8f89fcb898ac8b8fc3b243ee8a3efcd6ebe906b4f"></a>

## purge_duration_for_inactive_discovered_apis property — enable_api_discovery.discovered_api_settings / a29a07d898b3 / 4

Type: `"number"`. Computed.

Inactive discovered API will be deleted after configured duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
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
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

<a id="canonical-66358c95031689e0900b2d48a2a329996389f8c6b94ad9cc2189079897f82037"></a>

## Next pages — enable_api_discovery.discovered_api_settings / a29a07d898b3 / 5

- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-deed01e1e44c0ee218ce7865e0a436b73764161e6d7f566c49842915cd78cf57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3b0d5fdd747450c682622a9c2c1ff813c8a42d22a94b145b323599151579cad"></a>

## enable_api_discovery.enable_learn_from_redirect_traffic — enable_api_discovery.enable_learn_from_redirect_traffic / 16f9f1e9c1bf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- enable_api_discovery.enable_learn_from_redirect_traffic

<a id="canonical-8923c7312c8d44a76cfe78eb2b2aa98805efe516290bbf288f3cbeaebd42009c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable learn from redirect traffic.

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

<a id="canonical-4997e28315d95e714ddc36e6a28991ced21b53d16167ef04be1078be92d07745"></a>

## Direct properties — enable_api_discovery.enable_learn_from_redirect_traffic / 16f9f1e9c1bf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bda35b885b98a122fea2bcb82c173cdfc94348e65a023079ed55b051e7b331aa"></a>

## Next pages — enable_api_discovery.enable_learn_from_redirect_traffic / 16f9f1e9c1bf / 4

- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-d2b53cfe0548b07f41457dec340e33022c24ff0508c5cab482f0470528dc0dab)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0580b32e533934838810b561171d27d15f1c857edc4a222427c359305befa4aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c6130aefa17a0d1e865e7c6468e3387db65a2613194ee818b72f47cf1b1e0f7"></a>

## enable_challenge — enable_challenge / 68eae646265c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- enable_challenge

<a id="canonical-ac78562305fc93dc44bf704750de89f8edddab648fba24d6d00a069000afc00d"></a>

Type: `"single"`. Computed.

Configure auto mitigation i.e risk based challenges for malicious users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]"
}
```

<a id="canonical-9fe83bf38be9df058fd664d6c9ecdca50b518600342516ab552bb592444fde6c"></a>

## Direct properties — enable_challenge / 68eae646265c / 3

- [captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-017.md#canonical-193279397aca4dc7d6f9bec8796ddfd99d5e5b88eb556a81054bc248d157d7c5): complete subsection reference.

- [default_captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-017.md#canonical-f8e253f2cf60712a19631f8faf6bbbf841cbf1ef34f96243fe2968b16b67b7c9): complete subsection reference.

- [default_js_challenge_parameters](data-sources--http_loadbalancer--reference--group-017.md#canonical-d23266822f19790dc0d7159ec25368b3196a4400696884ca3242aba5d60336dd): complete subsection reference.

- [default_mitigation_settings](data-sources--http_loadbalancer--reference--group-017.md#canonical-cfd934b6e835a4e1c34c290552da16a278535ab7f75aaa22abdc507a95332b86): complete subsection reference.

- [js_challenge_parameters](data-sources--http_loadbalancer--reference--group-017.md#canonical-9277a1f9d301296ba3ac01c20ce410fbdbf2472b55ca03f88c2907db3f158026): complete subsection reference.

- [malicious_user_mitigation](data-sources--http_loadbalancer--reference--group-017.md#canonical-784a5cac7d9732a3c5ad26d4c6d739ecefd8ae3c1c8c60d3081bcb73e6307b58): complete subsection reference.

<a id="canonical-e162adbd1c70a35f4d0fe3c77a65fe7012e731d978019efd66c41efc6176fca1"></a>

## Next pages — enable_challenge / 68eae646265c / 4

- [enable_challenge.captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-017.md#canonical-193279397aca4dc7d6f9bec8796ddfd99d5e5b88eb556a81054bc248d157d7c5)
- [enable_challenge.default_captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-017.md#canonical-f8e253f2cf60712a19631f8faf6bbbf841cbf1ef34f96243fe2968b16b67b7c9)
- [enable_challenge.default_js_challenge_parameters](data-sources--http_loadbalancer--reference--group-017.md#canonical-d23266822f19790dc0d7159ec25368b3196a4400696884ca3242aba5d60336dd)
- [enable_challenge.default_mitigation_settings](data-sources--http_loadbalancer--reference--group-017.md#canonical-cfd934b6e835a4e1c34c290552da16a278535ab7f75aaa22abdc507a95332b86)
- [enable_challenge.js_challenge_parameters](data-sources--http_loadbalancer--reference--group-017.md#canonical-9277a1f9d301296ba3ac01c20ce410fbdbf2472b55ca03f88c2907db3f158026)
- [enable_challenge.malicious_user_mitigation](data-sources--http_loadbalancer--reference--group-017.md#canonical-784a5cac7d9732a3c5ad26d4c6d739ecefd8ae3c1c8c60d3081bcb73e6307b58)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-193279397aca4dc7d6f9bec8796ddfd99d5e5b88eb556a81054bc248d157d7c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99bef4f1d60c551a8090bcf8f79726312ae510c1070d7fb1fc2901d7a7cef6b3"></a>

## enable_challenge.captcha_challenge_parameters — enable_challenge.captcha_challenge_parameters / 8ecc142bd5e0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-017.md#canonical-0580b32e533934838810b561171d27d15f1c857edc4a222427c359305befa4aa)
- enable_challenge.captcha_challenge_parameters

<a id="canonical-ce42bd60298e791745f5f02e5cb3762eb375077e5859fef30485b18db1a5d4fa"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google
Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed
to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will
redirect..

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

<a id="canonical-3fc8d52f0a2383728eb445adaeea7994ab5907bb025217cca0aed4baeee01005"></a>

## Direct properties — enable_challenge.captcha_challenge_parameters / 8ecc142bd5e0 / 3

<a id="canonical-e26ec7a97c2cd8440cbbed1fe289529bd0a4a10e721f5a5066f386cb78a489d2"></a>

<a id="canonical-35327e01d30caba91b411e0f9e81130f8770d48350b19c94317bcb18307c8dab"></a>

## cookie_expiry property — enable_challenge.captcha_challenge_parameters / 8ecc142bd5e0 / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

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

<a id="canonical-de652a2f8ca7c92471b891fa3b2de64e13cd391f717761b72106a4ee542ca434"></a>

<a id="canonical-b085e1e3fe4675099daf823b4f39fad648dc129b14f96445f847f615d6f6925d"></a>

## custom_page property — enable_challenge.captcha_challenge_parameters / 8ecc142bd5e0 / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

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

<a id="canonical-e27300d5871a057d606c5ba3ccd2b37639c58976a0447c7d0b23d9285cd07c08"></a>

## Next pages — enable_challenge.captcha_challenge_parameters / 8ecc142bd5e0 / 6

- [enable_challenge](data-sources--http_loadbalancer--reference--group-017.md#canonical-0580b32e533934838810b561171d27d15f1c857edc4a222427c359305befa4aa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f8e253f2cf60712a19631f8faf6bbbf841cbf1ef34f96243fe2968b16b67b7c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae54311c0b56c046d5cb017f6c75f0a17cc577f7ce6db971673f6e1d8bbd3149"></a>

## enable_challenge.default_captcha_challenge_parameters — enable_challenge.default_captcha_challenge_parameters / 98989327decf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-017.md#canonical-0580b32e533934838810b561171d27d15f1c857edc4a222427c359305befa4aa)
- enable_challenge.default_captcha_challenge_parameters

<a id="canonical-759a487ff40bbdedcef6cd7f4fa72ea5f6ef2187ea7a7a5c1401ff105a56891c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default captcha challenge parameters.

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

<a id="canonical-e9e8984145136dab509a1c7108a5ccde366ac7dc9d0e04d62e7bfd36a71fa8ce"></a>

## Direct properties — enable_challenge.default_captcha_challenge_parameters / 98989327decf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5d0833ab0b5a5e77fe2a9e32cda0664e54ee3e4bb7a48f932700eb41a13b6ac1"></a>

## Next pages — enable_challenge.default_captcha_challenge_parameters / 98989327decf / 4

- [enable_challenge](data-sources--http_loadbalancer--reference--group-017.md#canonical-0580b32e533934838810b561171d27d15f1c857edc4a222427c359305befa4aa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d23266822f19790dc0d7159ec25368b3196a4400696884ca3242aba5d60336dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-478b0b4a7232d62d42bb22b13785bc20522c145d15c4ca661970e8fe49857a7a"></a>

## enable_challenge.default_js_challenge_parameters — enable_challenge.default_js_challenge_parameters / 3045787ab2fb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-017.md#canonical-0580b32e533934838810b561171d27d15f1c857edc4a222427c359305befa4aa)
- enable_challenge.default_js_challenge_parameters

<a id="canonical-c8863860addaa66a24fae03069a3b142ab02887a8c7856989f5d79d038fd086d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default js challenge parameters.

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

<a id="canonical-94cf2bc59c98cbcecf1ab1dd5c7381832d826322ffd16019fa0b8e05df4d4949"></a>

## Direct properties — enable_challenge.default_js_challenge_parameters / 3045787ab2fb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e853a77a6e2eeb75e96dc9c1d8f5edf31b8a817232955073e542cfeae6a66696"></a>

## Next pages — enable_challenge.default_js_challenge_parameters / 3045787ab2fb / 4

- [enable_challenge](data-sources--http_loadbalancer--reference--group-017.md#canonical-0580b32e533934838810b561171d27d15f1c857edc4a222427c359305befa4aa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-cfd934b6e835a4e1c34c290552da16a278535ab7f75aaa22abdc507a95332b86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43428e1a24ac6848392ecd7cbcc72b1b9fca40243a36077e56d44b61c0245001"></a>

## enable_challenge.default_mitigation_settings — enable_challenge.default_mitigation_settings / d090907e0560 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-017.md#canonical-0580b32e533934838810b561171d27d15f1c857edc4a222427c359305befa4aa)
- enable_challenge.default_mitigation_settings

<a id="canonical-4515825768bc87a135d94376bc0ef7dbd34a596d10dcf7fa4f79f731a7aae220"></a>

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

<a id="canonical-e019a79ea892a82068feb0b45cf0b03321abf5c3532a74f7265ddf7d6d9e77da"></a>

## Direct properties — enable_challenge.default_mitigation_settings / d090907e0560 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-40be1b3fd8ea4ed65d34b09f822ffe4f1840f6a7dd1bda1937eddfb5bf27ef86"></a>

## Next pages — enable_challenge.default_mitigation_settings / d090907e0560 / 4

- [enable_challenge](data-sources--http_loadbalancer--reference--group-017.md#canonical-0580b32e533934838810b561171d27d15f1c857edc4a222427c359305befa4aa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9277a1f9d301296ba3ac01c20ce410fbdbf2472b55ca03f88c2907db3f158026"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6eeb1adf2e08f0d6ed6aa6c718da53a82ec1c50f517826995966df69ad7a131c"></a>

## enable_challenge.js_challenge_parameters — enable_challenge.js_challenge_parameters / f43b43288424 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-017.md#canonical-0580b32e533934838810b561171d27d15f1c857edc4a222427c359305befa4aa)
- enable_challenge.js_challenge_parameters

<a id="canonical-1628138bc60ce9c0d2f17bacd419de1a0f1df9204bed044df7e5dfb2027af01c"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript.

With this feature enabled, only clients that are capable of executing Javascript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the Javascript. Javascript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid Javascript challenge for subsequent requests.

Javascript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

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

<a id="canonical-d1d16b9ad5ceb7bfb9a6216517247439e28f554b20a13e9bbfc6573a42f74bb5"></a>

## Direct properties — enable_challenge.js_challenge_parameters / f43b43288424 / 3

<a id="canonical-e0f6be13e1a951a83a61ccce2bab68a03efb3c7f54a351ceb2ddda9d9893f040"></a>

<a id="canonical-0838554da9cfc6d09b36564997d259d022c788de121161c3b1157c70297715d1"></a>

## cookie_expiry property — enable_challenge.js_challenge_parameters / f43b43288424 / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

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

<a id="canonical-323f14511f920ad936f4cec6ace3de6384e369016aba18df6124711470a0f817"></a>

<a id="canonical-c5c1bcbb02f47074311695170bcdaaebf4912a16322529682e065b396b3ff3ff"></a>

## custom_page property — enable_challenge.js_challenge_parameters / f43b43288424 / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

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

<a id="canonical-5c7a429cfb0aa38d3a934527e0980445e205da5a79a900f7b5d11651c31acb0d"></a>

<a id="canonical-00ffa5543f55d032b7235c97b96e906ad92150a7c00bfc90c6f76225c84e38a4"></a>

## js_script_delay property — enable_challenge.js_challenge_parameters / f43b43288424 / 6

Type: `"number"`. Computed.

Delay introduced by Javascript, in milliseconds.

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
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-b0b57079c382a1c063629f4a4294b05a6530bf7adb595fdba0eb17b3972ac429"></a>

## Next pages — enable_challenge.js_challenge_parameters / f43b43288424 / 7

- [enable_challenge](data-sources--http_loadbalancer--reference--group-017.md#canonical-0580b32e533934838810b561171d27d15f1c857edc4a222427c359305befa4aa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-784a5cac7d9732a3c5ad26d4c6d739ecefd8ae3c1c8c60d3081bcb73e6307b58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d650a054fe9b1c85176125454edace318e044aee78bb95f7c4a368a11110460"></a>

## enable_challenge.malicious_user_mitigation — enable_challenge.malicious_user_mitigation / 7cb20b0fbe6e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-017.md#canonical-0580b32e533934838810b561171d27d15f1c857edc4a222427c359305befa4aa)
- enable_challenge.malicious_user_mitigation

<a id="canonical-6ba5c5523c136970ae8fcaa7dd1e5e0c72a2bb0778fce9a93754f2e6dd9f6446"></a>

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

<a id="canonical-b1216dc9c599d42928ba38ccdb953a2834f7b02977fdbc603ebafbbd407eb9b2"></a>

## Direct properties — enable_challenge.malicious_user_mitigation / 7cb20b0fbe6e / 3

<a id="canonical-c6a42070c3ccb3b5b000e127cf46793ea7cf9e5671c63520934be6012ecd7987"></a>

<a id="canonical-590352f4e928be53fad23ed65271ae6dd403f2a537931d2bb5c63e7371291aaa"></a>

## name property — enable_challenge.malicious_user_mitigation / 7cb20b0fbe6e / 4

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

<a id="canonical-fd63fc0024a0ae6ea635ae6c50cf9080f97aeffce1b34424d33004358b09220b"></a>

<a id="canonical-4ed6a075143d549eaf91d11c2129755a196d20f229892ef8617561e6a804881a"></a>

## namespace property — enable_challenge.malicious_user_mitigation / 7cb20b0fbe6e / 5

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

<a id="canonical-ab266ff67e9d26cdc02de4e3bcd5b8d3236aebddc7cc3518b190e8cc41fd4f0b"></a>

<a id="canonical-9b9f776936950a2007bc183fe83c9c0009f708c3542e8fd4c7d8b93f3e1d840c"></a>

## tenant property — enable_challenge.malicious_user_mitigation / 7cb20b0fbe6e / 6

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

<a id="canonical-d1e44b8c2adab2303a226584645fcdb4014f1b27049fa991de5ddb5098b3c01d"></a>

## Next pages — enable_challenge.malicious_user_mitigation / 7cb20b0fbe6e / 7

- [enable_challenge](data-sources--http_loadbalancer--reference--group-017.md#canonical-0580b32e533934838810b561171d27d15f1c857edc4a222427c359305befa4aa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-aaa2e62f5cdd092b54ec1c2b8d102d67ce63fd12ca4f464b132fd31428945e25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bdd31f1773efd4d51e8fff644303f48eb994d39fc14686f90f22a7718f01fec"></a>

## enable_ip_reputation — enable_ip_reputation / 1e687b64b913 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- enable_ip_reputation

<a id="canonical-736d168bd3a13e81b8705b518c88767c85fde91faac9ca33aa4d6943e142b3c4"></a>

Type: `"single"`. Computed.

IP Threat Category List. List of IP threat categories.

Upstream description:

List of IP threat categories.

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

<a id="canonical-3d8d9d58bdceef4473cecc7af2b5cd09ef39a0d17258176f3ba1e88314585f35"></a>

## Direct properties — enable_ip_reputation / 1e687b64b913 / 3

<a id="canonical-465bb9c419174d6b1fbb9f244e588aca257d71aa3c3736a353cf9049d8032005"></a>

<a id="canonical-1a702ecf94ce99360a8647621f9434ab21563b1277a5a3d171631009133df500"></a>

## ip_threat_categories property — enable_ip_reputation / 1e687b64b913 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`, \`WEB\_ATTACKS\`, \`BOTNETS\`,
\`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`, \`MOBILE\_THREATS\`, \`TOR\_PROXY\`,
\`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to \`SPAM\_SOURCES\`.

Upstream description:

If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-52dece9d18c6cb60c4e7036bc9950581b5141fa6712bfe0adb11c656464a277b"></a>

## Next pages — enable_ip_reputation / 1e687b64b913 / 5

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d6e47ac52e62bc579bc00d01b31e62ac48677812d725d736decd3be81c8e489a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d60d8c2c31e60b2dbd75ff92980cd81d8c8603d3b18df5d06fc31d8b46b84811"></a>

## enable_malicious_user_detection — enable_malicious_user_detection / c19638682b2b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- enable_malicious_user_detection

<a id="canonical-94fbd15e16a2324db4830d97005cdf7b0a55079744d95ae905991326a953aad6"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable malicious user detection.

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

<a id="canonical-af79f1d29f05113beb1744b13fc30a5a294be0cd985d706c9c7ceddcd81a4b2a"></a>

## Direct properties — enable_malicious_user_detection / c19638682b2b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-25c277e868d42b79c47da2137ed25dcc3f7d6a96693372f0e670a57b44346caa"></a>

## Next pages — enable_malicious_user_detection / c19638682b2b / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-55b4a75f7dde05f4f4c851830cf84dcb2c4adca68704256b98a57e131fbdd9f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bd5b5ac411f752214656c0a2c5345f6f71259eac4385c7b3c858edba98c19b1"></a>

## enable_threat_mesh — enable_threat_mesh / 44b3e3f3ebd7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- enable_threat_mesh

<a id="canonical-e2092de6997caef2bee3fdeb76ccaecccc7c1280386bca74033aed25945d5326"></a>

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

<a id="canonical-1f631dd1f7020bd5f184d2c91d4ed36997f888b0d4012b58d5ccf3a8f2c359af"></a>

## Direct properties — enable_threat_mesh / 44b3e3f3ebd7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c3e7299ce7708ea36ca3f89dbfac5a1d44e65b7d5e457b26a5fa661e01f78176"></a>

## Next pages — enable_threat_mesh / 44b3e3f3ebd7 / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2bcad775662bd18661990a48eac1087fb2944081ffe2b61081ec4abf1649fa10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f342d111b4a984c91bc1c22178ae8e782d843cc5be0eafbcc5c88680e1f1b24b"></a>

## enable_trust_client_ip_headers — enable_trust_client_ip_headers / 2ce08fd7deac / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- enable_trust_client_ip_headers

<a id="canonical-a02a4674cca8bcf25c4a652cbbc98f54bdbd508a4850f8c22dffd62fc7666cd4"></a>

Type: `"single"`. Computed.

Trust Client IP Headers List. List of Client IP Headers.

Upstream description:

List of Client IP Headers.

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

<a id="canonical-1e4280340493e5b9cc64b02cbbbce50e06ca7d46757c8fc745c1c2b1004ac03e"></a>

## Direct properties — enable_trust_client_ip_headers / 2ce08fd7deac / 3

<a id="canonical-b53c0c926cda7cc034b988bb9a292a896d84ce3c7ff5afe0aa267de36da866da"></a>

<a id="canonical-1ead6a1fca66c453165a7f797f1260d653c571b4615438bfe93b3fdc77d1fb4f"></a>

## client_ip_headers property — enable_trust_client_ip_headers / 2ce08fd7deac / 4

Type: `["list", "string"]`. Computed.

Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom,
meaning if the first header is not present in the request, the system will proceed to check for the
second header, and so on, until one of the listed headers is found. If none of the defined..

Upstream description:

Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom,
meaning if the first header is not present in the request, the system will proceed to check for the
second header, and so on, until one of the listed headers is found. If none of the defined headers
exist, or the value is not an IP address, then the system will use the source IP of the packet. If
multiple defined headers with different names are present in the request, the value of the first
header name in the configuration will be used. If multiple defined headers with the same name are
present in the request, values of all those headers will be combined. The system will read the
right-most IP address from header, if there are multiple IP addresses in the header value. For
X-Forwarded-For header, the system will read the IP address(rightmost - 1), as the client IP.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-c1beeaaa00bf6a5c9d729c635054d57c7ee1d50ea816b96f123aad36126cc6e0"></a>

## Next pages — enable_trust_client_ip_headers / 2ce08fd7deac / 5

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1370b5a83b8a2ba11050926c0526c4373a074187e59fec6d2d3dff150b688eae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63b7c57e6b47dc93be2ea05a3e6ed44c6f4dff48aa249389ab948b106ad74e75"></a>

## graphql_rules — graphql_rules / 42d1d3cbe919 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- graphql_rules

<a id="canonical-c3c31abb16492c42178f03a7ce866255e00e6a6b0c58b211e98d2751b7bab4cd"></a>

Type: `"list"`. Computed.

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy..

Upstream description:

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks.

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

<a id="canonical-73301da198282fb3853c1f58dd3f7df6b3c12f1ed8ecf10fa84562be78896033"></a>

## Direct properties — graphql_rules / 42d1d3cbe919 / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-017.md#canonical-538a3953e5f2b2765b5d66056f3595864942d20988fe4df1096eecd630014ba2): complete subsection reference.

<a id="canonical-655ae07c059eb360db082a4a580df40028e73e5c30e7f2d4846f6e8582d2805e"></a>

<a id="canonical-627d8ce08603f9b8f4397b608e841ef36af3bea8e400342b8198a4a1382eb663"></a>

## exact_path property — graphql_rules / 42d1d3cbe919 / 4

Type: `"string"`. Computed.

Specifies the exact path to GraphQL endpoint. Defaults to \`/graphql\`.

Upstream description:

Specifies the exact path to GraphQL endpoint. Default value is /graphql.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-b5793c78dc515174d9c4fa17a0a32da23d9fd069bc7e0e1e3e12152c4696e14a"></a>

<a id="canonical-639fa3680d77f16bc4e11549e6b1a7019e8c71444afc102f5e046430415ecebc"></a>

## exact_value property — graphql_rules / 42d1d3cbe919 / 5

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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

- [graphql_settings](data-sources--http_loadbalancer--reference--group-017.md#canonical-1b0589e8af679b11596496aeaa7f6c690264daa4a7c01ea459587c5b2f916fc4): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-017.md#canonical-3d77d9d3d3d8bc7033e584577b7a1f9ac48e40fe1e5aaef730b184ac20c18312): complete subsection reference.

- [method_get](data-sources--http_loadbalancer--reference--group-017.md#canonical-a77581fbf1a5dd2b09f711a28622d697b8303a658529594c8f599380191fc4de): complete subsection reference.

- [method_post](data-sources--http_loadbalancer--reference--group-017.md#canonical-ea5d7fea54b0ef90ec792ccfc93f736b7b4dc0966c5d8576cade0e46895d5c9c): complete subsection reference.

<a id="canonical-ba19b139ebc72c943c1690fd7fa2d06f2a6a1cdcf937f05f318f540fd64e9d63"></a>

<a id="canonical-227dd84610498e32e89b168336fdee5f3234dcf4483edfb847b727127f599a97"></a>

## suffix_value property — graphql_rules / 42d1d3cbe919 / 6

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-21d844d7c8f94e6ed8c663a819005cfe0ee96eb18e00b15dbf914e64d9d63d8c"></a>

## Next pages — graphql_rules / 42d1d3cbe919 / 7

- [graphql_rules.any_domain](data-sources--http_loadbalancer--reference--group-017.md#canonical-538a3953e5f2b2765b5d66056f3595864942d20988fe4df1096eecd630014ba2)
- [graphql_rules.graphql_settings](data-sources--http_loadbalancer--reference--group-017.md#canonical-1b0589e8af679b11596496aeaa7f6c690264daa4a7c01ea459587c5b2f916fc4)
- [graphql_rules.metadata](data-sources--http_loadbalancer--reference--group-017.md#canonical-3d77d9d3d3d8bc7033e584577b7a1f9ac48e40fe1e5aaef730b184ac20c18312)
- [graphql_rules.method_get](data-sources--http_loadbalancer--reference--group-017.md#canonical-a77581fbf1a5dd2b09f711a28622d697b8303a658529594c8f599380191fc4de)
- [graphql_rules.method_post](data-sources--http_loadbalancer--reference--group-017.md#canonical-ea5d7fea54b0ef90ec792ccfc93f736b7b4dc0966c5d8576cade0e46895d5c9c)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-538a3953e5f2b2765b5d66056f3595864942d20988fe4df1096eecd630014ba2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9292b7a5bf182ee4b2e873166436cda906ab07f8bd1bcf71f253a76526c16cd"></a>

## graphql_rules.any_domain — graphql_rules.any_domain / 2de95f0c05ea / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-017.md#canonical-1370b5a83b8a2ba11050926c0526c4373a074187e59fec6d2d3dff150b688eae)
- graphql_rules.any_domain

<a id="canonical-b69ee08b48b0c8349cdb242b0c8d866cf5338049b597ad04739aa86dbb4c66f2"></a>

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

<a id="canonical-1678483c18e14b5dfb0968415d88ba0504f46b4039d6ab90c9f0751378ce6b96"></a>

## Direct properties — graphql_rules.any_domain / 2de95f0c05ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c675756fada54b541771d00f68f6177b10596e0b6944a750a0464ba0af9085e0"></a>

## Next pages — graphql_rules.any_domain / 2de95f0c05ea / 4

- [graphql_rules](data-sources--http_loadbalancer--reference--group-017.md#canonical-1370b5a83b8a2ba11050926c0526c4373a074187e59fec6d2d3dff150b688eae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1b0589e8af679b11596496aeaa7f6c690264daa4a7c01ea459587c5b2f916fc4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f105fed99ae26978c255b934ab059ef14b91243c3bd205f92f2cc1a128aba89"></a>

## graphql_rules.graphql_settings — graphql_rules.graphql_settings / fa3cfcdb93ce / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-017.md#canonical-1370b5a83b8a2ba11050926c0526c4373a074187e59fec6d2d3dff150b688eae)
- graphql_rules.graphql_settings

<a id="canonical-332f4ce3462a645576cb1777aa2b615ddc4d2f57a71eba9fa17f8d8be8bd97b1"></a>

Type: `"single"`. Computed.

Configuration parameter for graphql settings.

Upstream description:

GraphQL configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allow_introspection_queries_choice": "[\"disable_introspection\",\"enable_introspection\"]"
}
```

<a id="canonical-4ebfabd04c4a5ac4764d6d6e818ea09521eaacc99a7193ea8d4a765bb2233b3e"></a>

## Direct properties — graphql_rules.graphql_settings / fa3cfcdb93ce / 3

- [disable_introspection](data-sources--http_loadbalancer--reference--group-017.md#canonical-069390c91dd25871ad900105c6234144795e883cfb5c0f5b056642d51b7ffd2e): complete subsection reference.

- [enable_introspection](data-sources--http_loadbalancer--reference--group-017.md#canonical-9251acccbccc700393738342131aa8fe9a7c4bfa51dc1d4f9de639f493644059): complete subsection reference.

<a id="canonical-5776a058809f4c79c31fa592a8032e767bf1ddedac9be65804e7e5290a64ff6d"></a>

<a id="canonical-fa8108f23e1e4f57b058b45714d4988ca9a4b5c100b8c7adebf0b9d9da4bdf13"></a>

## max_batched_queries property — graphql_rules.graphql_settings / fa3cfcdb93ce / 4

Type: `"number"`. Computed.

Specify maximum number of queries in a single batched request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-d9e6c32b737993cfdc5bbdee040431065d066a92147eb8342d1436b340d0e180"></a>

<a id="canonical-e1eb7dbe01d88d279890303da1ca148c6a90d93e0f4e2ccb9e03b8d57786936b"></a>

## max_depth property — graphql_rules.graphql_settings / fa3cfcdb93ce / 5

Type: `"number"`. Computed.

Specify maximum depth for the GraphQL query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-2000f72afe06f8674fa9972395223ac6af8a4dc8e2a4c4435cae693b5c31b051"></a>

<a id="canonical-41966b3167bbd284279bae6ba15cf0b2060aff6e3d5804563ae02536e414fbf4"></a>

## max_total_length property — graphql_rules.graphql_settings / fa3cfcdb93ce / 6

Type: `"number"`. Computed.

Specify maximum length in bytes for the GraphQL query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16386,
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
    "ves.io.schema.rules.uint32.lte": "16386"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  }
}
```

<a id="canonical-cc8c0a0ac6f0a1838a2b5e7a33f39564edb277a211dd64e96c2d988e9cf98719"></a>

## Next pages — graphql_rules.graphql_settings / fa3cfcdb93ce / 7

- [graphql_rules.graphql_settings.disable_introspection](data-sources--http_loadbalancer--reference--group-017.md#canonical-069390c91dd25871ad900105c6234144795e883cfb5c0f5b056642d51b7ffd2e)
- [graphql_rules.graphql_settings.enable_introspection](data-sources--http_loadbalancer--reference--group-017.md#canonical-9251acccbccc700393738342131aa8fe9a7c4bfa51dc1d4f9de639f493644059)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-017.md#canonical-1370b5a83b8a2ba11050926c0526c4373a074187e59fec6d2d3dff150b688eae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-069390c91dd25871ad900105c6234144795e883cfb5c0f5b056642d51b7ffd2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e04f7774c515f54519a763a83f0ecbdf3f56385aba0147cdc5d5db1365f0db6"></a>

## graphql_rules.graphql_settings.disable_introspection — graphql_rules.graphql_settings.disable_introspection / e75572a73b25 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-017.md#canonical-1370b5a83b8a2ba11050926c0526c4373a074187e59fec6d2d3dff150b688eae)
- [graphql_rules.graphql_settings](data-sources--http_loadbalancer--reference--group-017.md#canonical-1b0589e8af679b11596496aeaa7f6c690264daa4a7c01ea459587c5b2f916fc4)
- graphql_rules.graphql_settings.disable_introspection

<a id="canonical-72d2440d5b4605d3e80ca68132173c6ddfc6c833f34f384dc858af6cefec60e9"></a>

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

<a id="canonical-e1ca51c0860550ba7a1d143547d8b781413e52e41ab0567f33d8d93e04f69c9b"></a>

## Direct properties — graphql_rules.graphql_settings.disable_introspection / e75572a73b25 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-80d2f21c159b1a2ddafeaf7b3efc476806102422c4d121947cc86153d1620efc"></a>

## Next pages — graphql_rules.graphql_settings.disable_introspection / e75572a73b25 / 4

- [graphql_rules.graphql_settings](data-sources--http_loadbalancer--reference--group-017.md#canonical-1b0589e8af679b11596496aeaa7f6c690264daa4a7c01ea459587c5b2f916fc4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9251acccbccc700393738342131aa8fe9a7c4bfa51dc1d4f9de639f493644059"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afe11af71472a38dd30831d0939747d3568f49e8ab448c4907124b738da501b1"></a>

## graphql_rules.graphql_settings.enable_introspection — graphql_rules.graphql_settings.enable_introspection / ee8730f8e81d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-017.md#canonical-1370b5a83b8a2ba11050926c0526c4373a074187e59fec6d2d3dff150b688eae)
- [graphql_rules.graphql_settings](data-sources--http_loadbalancer--reference--group-017.md#canonical-1b0589e8af679b11596496aeaa7f6c690264daa4a7c01ea459587c5b2f916fc4)
- graphql_rules.graphql_settings.enable_introspection

<a id="canonical-2617a7a516f4d24309df4fcb6ec21fe4957982f5a6928d1fd3940615edfc7e86"></a>

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

<a id="canonical-b62d914174198bd93d003cee91d8682f5bffae9c4e8532123e136c2f6627c77d"></a>

## Direct properties — graphql_rules.graphql_settings.enable_introspection / ee8730f8e81d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ccbc67ff8177238ce5ba42be24b0b7a95e0eb349df09344fcf7da1e08583c5d5"></a>

## Next pages — graphql_rules.graphql_settings.enable_introspection / ee8730f8e81d / 4

- [graphql_rules.graphql_settings](data-sources--http_loadbalancer--reference--group-017.md#canonical-1b0589e8af679b11596496aeaa7f6c690264daa4a7c01ea459587c5b2f916fc4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3d77d9d3d3d8bc7033e584577b7a1f9ac48e40fe1e5aaef730b184ac20c18312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4ba85c147b9ca97b4b021df65745871d6fe54f8950ae4a72f272bdfe3c1d4b6"></a>

## graphql_rules.metadata — graphql_rules.metadata / 46b4589ab4f7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-017.md#canonical-1370b5a83b8a2ba11050926c0526c4373a074187e59fec6d2d3dff150b688eae)
- graphql_rules.metadata

<a id="canonical-475cfc423ac0a179563e6e71d9cf02de8eaedb090c9f8314a8f20f206781e879"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-8a557228cdbad2df35d37528ca81a42f595077d7c088c7f2c65a2a3d1a341f5c"></a>

## Direct properties — graphql_rules.metadata / 46b4589ab4f7 / 3

<a id="canonical-c54e22b0b5f3d6ab037f7c6c120087ec54e36b49b9a65a2dd7ae2388db0b7d33"></a>

<a id="canonical-5ad99a0ffb5dc5b1204045c44253848c14727d0136f89d6985348c8290e41aaf"></a>

## description_spec property — graphql_rules.metadata / 46b4589ab4f7 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-ddec9591ff7ec5b06c571b71ddb7dbcc811923733310514f5b7bb76cfc0763bc"></a>

<a id="canonical-7a1d62382de54be9086f24587ce8f567c748237f21c3f8a186edfe00f5077113"></a>

## name property — graphql_rules.metadata / 46b4589ab4f7 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-babfa7b481d4207f5a593e32460763ad8da91494d5b7e80bbecc7b1063d4550a"></a>

## Next pages — graphql_rules.metadata / 46b4589ab4f7 / 6

- [graphql_rules](data-sources--http_loadbalancer--reference--group-017.md#canonical-1370b5a83b8a2ba11050926c0526c4373a074187e59fec6d2d3dff150b688eae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a77581fbf1a5dd2b09f711a28622d697b8303a658529594c8f599380191fc4de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03e5daa933291e4442d15a9ef13df82a8e747d6ffd607e5525778a528bdfc1ce"></a>

## graphql_rules.method_get — graphql_rules.method_get / 5eada1d8463b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-017.md#canonical-1370b5a83b8a2ba11050926c0526c4373a074187e59fec6d2d3dff150b688eae)
- graphql_rules.method_get

<a id="canonical-da53f0cdf1c6cecc4f0b51496efb7d8ef7d629886a66c45e169e490aa088f924"></a>

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

<a id="canonical-1cae05d4f9928a29e8930c6e9d8c0edce61f6cae8282fae88f6916e9853f5ca5"></a>

## Direct properties — graphql_rules.method_get / 5eada1d8463b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5d147d688a4d9b7910a9d33f92c037220d1c4e54c808929fc6c9495e8185a230"></a>

## Next pages — graphql_rules.method_get / 5eada1d8463b / 4

- [graphql_rules](data-sources--http_loadbalancer--reference--group-017.md#canonical-1370b5a83b8a2ba11050926c0526c4373a074187e59fec6d2d3dff150b688eae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ea5d7fea54b0ef90ec792ccfc93f736b7b4dc0966c5d8576cade0e46895d5c9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c328770f442e6196764d3ac4dbe5ba9b8d6a8c74a9b6f199ef47485774dca48"></a>

## graphql_rules.method_post — graphql_rules.method_post / e257e68398e4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-017.md#canonical-1370b5a83b8a2ba11050926c0526c4373a074187e59fec6d2d3dff150b688eae)
- graphql_rules.method_post

<a id="canonical-927f5b69ed4bec48604df748ff0c9580480e2161558b2d945c6147335ae4f8df"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for method post.

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

<a id="canonical-1fd2430209a80f0aa7d83387a999b8dadb848785de059a74060a0a688a9fff0e"></a>

## Direct properties — graphql_rules.method_post / e257e68398e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-030dee6e7a9557542f4d267193513de0cb8a560a3ba86ec2aad5ba52d283e617"></a>

## Next pages — graphql_rules.method_post / e257e68398e4 / 4

- [graphql_rules](data-sources--http_loadbalancer--reference--group-017.md#canonical-1370b5a83b8a2ba11050926c0526c4373a074187e59fec6d2d3dff150b688eae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6901a453d09cfcee9e11fc8486b63aaba8056dfe0273a80fdfafc2c0920ef5e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-711e46b4e4b449f2821ec13520202acb1c80d9a45751884afb5b0adcf6ea2224"></a>

## http — http / cc20cd09879b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- http

<a id="canonical-3cd8621c4ea57491bb9106fe3bc27cf376134149500c7ff76c1009272257af7a"></a>

Type: `"single"`. Computed.

\[OneOf: http, https, https\_auto\_cert; Default: https\_auto\_cert\] HTTP Choice. Choice for
selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

OneOf alternatives in this subsection:

- [http](data-sources--http_loadbalancer--reference--group-017.md#canonical-3cd8621c4ea57491bb9106fe3bc27cf376134149500c7ff76c1009272257af7a)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-6c21a9ef846c1a3e84fbb425cb076c4110e624f44cb5f46205344143821dc2d1)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-94bcef80ac42c4b17d5d9e65955f67b3339f56bd0be651f841e5fe2c67636962)

Select alternatives according to the provider validators above.

<a id="canonical-6154374f2be45f837ef1c894d94bcab38c03899c1488489448a754265e794945"></a>

## Direct properties — http / cc20cd09879b / 3

<a id="canonical-bfe1b6ad172dfc145cffe3acbb2746acbd6dea5410c6ef0858ee3d1b43e28c59"></a>

<a id="canonical-e81cede84837e6e7e75b7cb7f06363cc2a5f28f23749ba8f16b631deb2778ce6"></a>

## dns_volterra_managed property — http / cc20cd09879b / 4

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

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

<a id="canonical-b0a93204a7c94c8b59cbd8bc5b894874bbb05b7a20b6fbfa77766cd9c6a0b09b"></a>

<a id="canonical-9850444f7b28062d6c3819b3baa637846c31b56e5ea4f5190985fcae5c93863f"></a>

## port property — http / cc20cd09879b / 5

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-eae8a719db0716b344ec89ced1caae76dfd73517922098ea52b24444ffa2ee04"></a>

<a id="canonical-36848d665c8f5dc3e3796feaebf019328b444d02dbae93f27c6d512ef723d808"></a>

## port_ranges property — http / cc20cd09879b / 6

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

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

<a id="canonical-0c732c4dd116e0ea5df7be572bc8693c06e03c70ffa0130dc898c7b87a6cb8e8"></a>

## Next pages — http / cc20cd09879b / 7

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26e2db22d52f38eba8d49cb65d7581a088e0fc73e6dc575c531c5b8e0549c506"></a>

## https — https / 230ec52ee4e6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- https

<a id="canonical-6c21a9ef846c1a3e84fbb425cb076c4110e624f44cb5f46205344143821dc2d1"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

<a id="canonical-1f26021d4439e2b351d96dbe87642ecb19697a1f4720be01771cc1b97bfcc5c3"></a>

## Direct properties — https / 230ec52ee4e6 / 3

<a id="canonical-bc436d285c6f7ce7b80af94ba9688f2f7a2a7cb820131b9d2587c1ae3544e37d"></a>

<a id="canonical-44d0a76ba5fc385df29ba757d1c10c7d6e6ea771fd35c1de7897aa44dcdb8ebc"></a>

## add_hsts property — https / 230ec52ee4e6 / 4

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header.

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

<a id="canonical-76e67ca0b571ed97a90fc82ac7c15563be444c887e8d955c6a18882352a44c6d"></a>

<a id="canonical-7c26828d06746c7db9f0b2e13dd0be5a76f054dd6dccbe0d3d2d14f1ce9a94df"></a>

## append_server_name property — https / 230ec52ee4e6 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-41a23dbccd1c97f2dc1e26c46e237d1a27fb06d4debe54e928158e55e838b0be): complete subsection reference.

<a id="canonical-41c6b00e48c0632a63d524308caec8c3ae5608939a67c0f5588d8e66b198983d"></a>

<a id="canonical-372a3b18abfd8b6c5c88f22031a0f7bcf5097fe54eb1751dae9d78c338a6e910"></a>

## connection_idle_timeout property — https / 230ec52ee4e6 / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](data-sources--http_loadbalancer--reference--group-018.md#canonical-45a289eb0b2b5755f9439993b2190f90f49954fd8c1bc921503aa970bf4dddaf): complete subsection reference.

- [default_loadbalancer](data-sources--http_loadbalancer--reference--group-018.md#canonical-5b0efb11f9fd0b7666a36992cab5d99e5ead20fbcc2aa2434c8092674ca51ad1): complete subsection reference.

- [disable_path_normalize](data-sources--http_loadbalancer--reference--group-018.md#canonical-77b368d92971088c9c30dd54008c34933cacc279e688507abbf7369af2972041): complete subsection reference.

- [enable_path_normalize](data-sources--http_loadbalancer--reference--group-018.md#canonical-e5127c793408a0cfb866a5a6fa3cd627e1b66df8afb6371ae5bb2c998660acaf): complete subsection reference.

- [http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-032234c0bf40bef5559c9957b38561f39bb4cd94e21c9e1efcf689c954b3a809): complete subsection reference.

<a id="canonical-cd066b58d594aba43d4c131817d6f00bba8b8a52d8b4bd47c4c8f442068444d4"></a>

<a id="canonical-419b30dbfb33f4f725270401e6a8c163117d21ba6a419fb4bcb80707415f2660"></a>

## http_redirect property — https / 230ec52ee4e6 / 7

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

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

- [non_default_loadbalancer](data-sources--http_loadbalancer--reference--group-018.md#canonical-51f361932fee4d694f1924fb28f3c46f54163b6aad91f8294a4062ebd6122245): complete subsection reference.

- [pass_through](data-sources--http_loadbalancer--reference--group-018.md#canonical-e4ba0983025b8422431c4ada41aeace4ba6fe934fa5855e9b51e2c4a21edc80d): complete subsection reference.

<a id="canonical-a95e08430e2cfc728861a98b03ffd7d27d85cd0a0f2dc8d7becab980a7e86072"></a>

<a id="canonical-761ed2288424e589ae5b4cffd2d7324bd4fee5ca74b9528d430f1337bcc940bd"></a>

## port property — https / 230ec52ee4e6 / 8

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-c157e054339cf234c3259d08103daf1878d0302c881f909a07ffcb8b610cc795"></a>
