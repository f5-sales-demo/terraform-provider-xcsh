---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-47448c65305fb1c84696b09f69244e1b21554db9da00b9b8d26f6de2e3da04ef"></a>

## name property — default_pool.origin_servers.consul_service.site_locator.site / 6b1b73976f41 / 4

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

<a id="canonical-e7f300abbb32ccd7ba48e558520e28f5e9e7c0258fbccea37681d9d24d6107a3"></a>

<a id="canonical-6ae6dd236a304ba458489ac5f539bf923923c83d7d8754c4379569b70f16d9d2"></a>

## namespace property — default_pool.origin_servers.consul_service.site_locator.site / 6b1b73976f41 / 5

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

<a id="canonical-8a647c3e6715003290e50308417a7fc71e3e26c1161b20d2a8371f291cd8ee1a"></a>

<a id="canonical-9dab8d4fb474e18fa2410a1be255bf6f9c63f3c5b2ce363d7882217b615d16d7"></a>

## tenant property — default_pool.origin_servers.consul_service.site_locator.site / 6b1b73976f41 / 6

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

<a id="canonical-3dbeb2ea3d69154e5a3497bd1eb9ca408ace19bba2cda81e47f55921f765ebc3"></a>

## Next pages — default_pool.origin_servers.consul_service.site_locator.site / 6b1b73976f41 / 7

- [default_pool.origin_servers.consul_service.site_locator](resources--http_loadbalancer--reference--group-015.md#canonical-7b550af37617ec41d35a67185bca1602a1724fa8e3f577e9d5d71e95519db17e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1fc8d8eed953d0af1fc113c275e18a7cbc94fd76d524e914e6d0447f6ef9ecad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fbcfd0fc62f221e985bdb9d9f0156130bc9dc5e11945f9dde580f50a3f8fac3"></a>

## default_pool.origin_servers.consul_service.site_locator.virtual_site — default_pool.origin_servers.consul_service.site_locator.virtual_site / f789971be50b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.consul_service](resources--http_loadbalancer--reference--group-015.md#canonical-882d5272609b857112020b2484260710807f5d78a8033fc8024b8611e0d86843)
- [default_pool.origin_servers.consul_service.site_locator](resources--http_loadbalancer--reference--group-015.md#canonical-7b550af37617ec41d35a67185bca1602a1724fa8e3f577e9d5d71e95519db17e)
- default_pool.origin_servers.consul_service.site_locator.virtual_site

<a id="canonical-44f1661c98904c8eac2703a320f4ee163851812d7bcc352ded25d188427b4004"></a>

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

<a id="canonical-694c96cb9692822213a3e812d6150810cb952e472e9ddff620b1429763c92488"></a>

## Direct properties — default_pool.origin_servers.consul_service.site_locator.virtual_site / f789971be50b / 3

<a id="canonical-69c2472c3a8b3188093f53b395941202264f8711cbae34a793ee274036dee949"></a>

<a id="canonical-ae103fb98914a21e0a3f4bbc35bc279834eed3b314553458cdb280afa986f0b5"></a>

## name property — default_pool.origin_servers.consul_service.site_locator.virtual_site / f789971be50b / 4

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

<a id="canonical-b6a2d00653bfb61e9df9c9eb83101b4ebf43a41bd39957aa98424293e92638bf"></a>

<a id="canonical-ef4388eb1755df62ae63adf8e789ac6e65b21182337b7f34d5b608c88d194378"></a>

## namespace property — default_pool.origin_servers.consul_service.site_locator.virtual_site / f789971be50b / 5

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

<a id="canonical-d0dea90cc54ad55e361d2685a5442f14878aa04b8f1e3b96aa8d37b721ae44ad"></a>

<a id="canonical-a92cda4691a5a624b355e772f30022dae1ef61c544d85b6a253702c20a4fc38a"></a>

## tenant property — default_pool.origin_servers.consul_service.site_locator.virtual_site / f789971be50b / 6

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

<a id="canonical-6547f36db5c90082541d8b3d1baecf8fa66b1c9eaaa34f2254689f3cc5f2b587"></a>

## Next pages — default_pool.origin_servers.consul_service.site_locator.virtual_site / f789971be50b / 7

- [default_pool.origin_servers.consul_service.site_locator](resources--http_loadbalancer--reference--group-015.md#canonical-7b550af37617ec41d35a67185bca1602a1724fa8e3f577e9d5d71e95519db17e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-63835deddb723b6526d8107880fd893003156b59a4af0db7a77d6b9ba62f1181"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2fd5840592ad53a0ae279105960774d8abf1bac0a13f72ef042d994d78b0f337"></a>

## default_pool.origin_servers.consul_service.snat_pool — default_pool.origin_servers.consul_service.snat_pool / a8c6dc5f9c13 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.consul_service](resources--http_loadbalancer--reference--group-015.md#canonical-882d5272609b857112020b2484260710807f5d78a8033fc8024b8611e0d86843)
- default_pool.origin_servers.consul_service.snat_pool

<a id="canonical-5c09980c64f7918c513d2726b2ef8f8897d18b182969fafa9273953d0188b986"></a>

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

<a id="canonical-b6c3abe2941b652f4fe2a762d503055212065bafbdc285b2c5b81980527ad86f"></a>

## Direct properties — default_pool.origin_servers.consul_service.snat_pool / a8c6dc5f9c13 / 3

- [no_snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-b10d5a6847547eab552880f53f0a745808104ae16dc0ab64ff1a7d3b8b7f6909): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3f4b45b7a0215265fb27bdd3234c5dabf3c2e1297f63299e949e9f97c7d88696): complete subsection reference.

<a id="canonical-5c3fef224d002c9ec6a791d501545b64e8f4664c84be33ac9d901f7e7032172d"></a>

## Next pages — default_pool.origin_servers.consul_service.snat_pool / a8c6dc5f9c13 / 4

- [default_pool.origin_servers.consul_service.snat_pool.no_snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-b10d5a6847547eab552880f53f0a745808104ae16dc0ab64ff1a7d3b8b7f6909)
- [default_pool.origin_servers.consul_service.snat_pool.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3f4b45b7a0215265fb27bdd3234c5dabf3c2e1297f63299e949e9f97c7d88696)
- [default_pool.origin_servers.consul_service](resources--http_loadbalancer--reference--group-015.md#canonical-882d5272609b857112020b2484260710807f5d78a8033fc8024b8611e0d86843)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b10d5a6847547eab552880f53f0a745808104ae16dc0ab64ff1a7d3b8b7f6909"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffff3b378472ddd089a3e24f347f25905ad50e96c04f7bdcf880369a32e12a43"></a>

## default_pool.origin_servers.consul_service.snat_pool.no_snat_pool — default_pool.origin_servers.consul_service.snat_pool.no_snat_pool / 54a93feb2cd4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.consul_service](resources--http_loadbalancer--reference--group-015.md#canonical-882d5272609b857112020b2484260710807f5d78a8033fc8024b8611e0d86843)
- [default_pool.origin_servers.consul_service.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-63835deddb723b6526d8107880fd893003156b59a4af0db7a77d6b9ba62f1181)
- default_pool.origin_servers.consul_service.snat_pool.no_snat_pool

<a id="canonical-89ad730d54a4dfbec16a8946474db1fef75699c8544d20b45901c862fc38e2c8"></a>

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

<a id="canonical-a791fa58fe00f1ce9843f089c92aa0c812550069aba6fa5b22ad616543265d68"></a>

## Direct properties — default_pool.origin_servers.consul_service.snat_pool.no_snat_pool / 54a93feb2cd4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fcd95a6d7198b1e66e08bbc24376e8ed751cdd26451c673f0cf12253f6154cd0"></a>

## Next pages — default_pool.origin_servers.consul_service.snat_pool.no_snat_pool / 54a93feb2cd4 / 4

- [default_pool.origin_servers.consul_service.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-63835deddb723b6526d8107880fd893003156b59a4af0db7a77d6b9ba62f1181)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3f4b45b7a0215265fb27bdd3234c5dabf3c2e1297f63299e949e9f97c7d88696"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63b5eb1a00c333c83690e6e1736be5a8686525e357b23d1504b8fcaa22624adf"></a>

## default_pool.origin_servers.consul_service.snat_pool.snat_pool — default_pool.origin_servers.consul_service.snat_pool.snat_pool / 68b6154a83f4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.consul_service](resources--http_loadbalancer--reference--group-015.md#canonical-882d5272609b857112020b2484260710807f5d78a8033fc8024b8611e0d86843)
- [default_pool.origin_servers.consul_service.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-63835deddb723b6526d8107880fd893003156b59a4af0db7a77d6b9ba62f1181)
- default_pool.origin_servers.consul_service.snat_pool.snat_pool

<a id="canonical-0b41f5104c9024285e7e92c006beb72f6955513a1e3bcc7a3a424ffb715314b6"></a>

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

<a id="canonical-fef792bcb60938846bbd905e4b02276a8e3f42af87fd088d23ba4e24b11bb89e"></a>

## Direct properties — default_pool.origin_servers.consul_service.snat_pool.snat_pool / 68b6154a83f4 / 3

<a id="canonical-c1acdab88e4e8dfea38c59970e7400bb51288b614affba9164f21becbdc4fcaf"></a>

<a id="canonical-7e575d81f25d693751ab21a975eb870ca6d8cfefb46fd7271baec57231be1798"></a>

## prefixes property — default_pool.origin_servers.consul_service.snat_pool.snat_pool / 68b6154a83f4 / 4

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

<a id="canonical-8c997dd16a1dea682c22516b2314f213f3c62be5e030428744968b91564876ee"></a>

## Next pages — default_pool.origin_servers.consul_service.snat_pool.snat_pool / 68b6154a83f4 / 5

- [default_pool.origin_servers.consul_service.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-63835deddb723b6526d8107880fd893003156b59a4af0db7a77d6b9ba62f1181)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f846047f11f88994cb51966325c84bbd318dc82fe9be62217adfc9505485871b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-823e16e8bc88b269cc0e1f1d6abe82b87d10852bd17e3cf928643f9e062de9f7"></a>

## default_pool.origin_servers.custom_endpoint_object — default_pool.origin_servers.custom_endpoint_object / 9aff4684ee19 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- default_pool.origin_servers.custom_endpoint_object

<a id="canonical-bd2e56a4606195fb83f73dbe0428550c063781eb17044ed1df37c4b5d4aca9a4"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with a reference to endpoint object.

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
custom_endpoint_object {
  # Configure direct properties listed below.
}
```

<a id="canonical-5b0a894b710404303090c2bd98b47ae98bfe0d3abb4d72d561b1fdd3ea0be35a"></a>

## Direct properties — default_pool.origin_servers.custom_endpoint_object / 9aff4684ee19 / 3

- [endpoint](resources--http_loadbalancer--reference--group-016.md#canonical-019484d7d58b227f0098d01018e9fc831a7601d77d5cba07282d9a1376b36962): complete subsection reference.

<a id="canonical-645662d14843ae9b94b098bb3d7b9b8a8550b8c543cceccbedcef870c5dd7c0f"></a>

## Next pages — default_pool.origin_servers.custom_endpoint_object / 9aff4684ee19 / 4

- [default_pool.origin_servers.custom_endpoint_object.endpoint](resources--http_loadbalancer--reference--group-016.md#canonical-019484d7d58b227f0098d01018e9fc831a7601d77d5cba07282d9a1376b36962)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-019484d7d58b227f0098d01018e9fc831a7601d77d5cba07282d9a1376b36962"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e6bbce1e72dbe05fb6b6758427386755460838cbcd5dd5f68ee4d955a818497"></a>

## default_pool.origin_servers.custom_endpoint_object.endpoint — default_pool.origin_servers.custom_endpoint_object.endpoint / f89fa0a75aa8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.custom_endpoint_object](resources--http_loadbalancer--reference--group-016.md#canonical-f846047f11f88994cb51966325c84bbd318dc82fe9be62217adfc9505485871b)
- default_pool.origin_servers.custom_endpoint_object.endpoint

<a id="canonical-aac616d4e0cb9739c033449ea6e35899cad9dfdbbc8783281c83282ee8299646"></a>

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
endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-74c8eb1a50a6b0d7130167c3e76a91e18f710041670099abc191ee47a0647b62"></a>

## Direct properties — default_pool.origin_servers.custom_endpoint_object.endpoint / f89fa0a75aa8 / 3

<a id="canonical-8f10d37b0c3a55f9d67e5671a894416a87724e56d89b0fa5f2b2a5a6fe4baeff"></a>

<a id="canonical-018d7f261e69f81951152e6b90d7175c8f309db2d10443986ac42335a76735d6"></a>

## name property — default_pool.origin_servers.custom_endpoint_object.endpoint / f89fa0a75aa8 / 4

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

<a id="canonical-1e9577220ba9197c3e3364e1e5b7b315f318116fe3bdb04644e9e47f20af08ff"></a>

<a id="canonical-677d3eaa770ad8ba5334c515bfd7d114f673e8a49d4fedd503e01d872e6e2bda"></a>

## namespace property — default_pool.origin_servers.custom_endpoint_object.endpoint / f89fa0a75aa8 / 5

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

<a id="canonical-a4f48a38b4b58bc991c4eb0d38af74e9e763e9f01d25ad4dd7a01710224e6183"></a>

<a id="canonical-c2d719362153189f98f284c2406eacc90c0f9f74c59b5945ce3188d2aaf5c327"></a>

## tenant property — default_pool.origin_servers.custom_endpoint_object.endpoint / f89fa0a75aa8 / 6

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

<a id="canonical-1c9c99e4cff6df3c9c91589d112289ac1e62d7dfeec6903286e46870a03b92a9"></a>

## Next pages — default_pool.origin_servers.custom_endpoint_object.endpoint / f89fa0a75aa8 / 7

- [default_pool.origin_servers.custom_endpoint_object](resources--http_loadbalancer--reference--group-016.md#canonical-f846047f11f88994cb51966325c84bbd318dc82fe9be62217adfc9505485871b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5b00c1b4651b1e68d95159487f98ce2b52ec91cde619fef3f53ef40f49a72b8"></a>

## default_pool.origin_servers.k8s_service — default_pool.origin_servers.k8s_service / 547d4734d0f2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- default_pool.origin_servers.k8s_service

<a id="canonical-d7f00e896059383e260e7837c68f5c9bf39ad67ffe8a5eeb8174850da1f4f0e1"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with K8s service name and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "vk8s_networks"),
  validators.ConflictingObjectAttributes("outside_network",
    "vk8s_networks")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

Terraform syntax:

```terraform
k8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-8674a75338bdaa02432b56db3de01898d857b216a696b1b97a093af654b19efd"></a>

## Direct properties — default_pool.origin_servers.k8s_service / 547d4734d0f2 / 3

- [inside_network](resources--http_loadbalancer--reference--group-016.md#canonical-1c06e483345c30fbc3c5b16d887bed0987d99e9c161a7faba79dc6b2b7f01f31): complete subsection reference.

- [outside_network](resources--http_loadbalancer--reference--group-016.md#canonical-cf185ef09e8c0dc13d8fa58f69f66e1804a98fa63a9c98b9d75f8363defd5b45): complete subsection reference.

<a id="canonical-6f8ace056fa109c6aff505f5e3ab9939c18b00f1604a14e33cbe947ef2d15949"></a>

<a id="canonical-27d7060382de204fc6c18c516a113010ae62debb68699fbd957a761b51b185e2"></a>

## protocol property — default_pool.origin_servers.k8s_service / 547d4734d0f2 / 4

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

&#8203;- PROTOCOL\_UDP: UDP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0fb971ac163f2f4b9632de8ffe1d32a64d4ef5604d36d2683cf4bb0b9226c7d8"></a>

<a id="canonical-75ae7fcb50dae3445047a590ff3c9868474248cf00943db231bced41d6bb03ea"></a>

## service_name property — default_pool.origin_servers.k8s_service / 547d4734d0f2 / 5

Type: `"string"`. Optional.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Upstream description:

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-b3e34f4a474f0153625579f9efd043ca3765bd9ecad139ca627212a8a03f7efe): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-b18f25eee537497ad5e3d59073d97aa92d8f358f28ff0050a204518d1c9f92b3): complete subsection reference.

- [vk8s_networks](resources--http_loadbalancer--reference--group-016.md#canonical-e9d4477a5ef3fffceb08a01e5df112ad52bf3523bddc28abe7eff388abd8ccff): complete subsection reference.

<a id="canonical-d1a6547c8814ed83be5b8d8371c578efa59b30968b75dbf0565b85bcb1f76124"></a>

## Next pages — default_pool.origin_servers.k8s_service / 547d4734d0f2 / 6

- [default_pool.origin_servers.k8s_service.inside_network](resources--http_loadbalancer--reference--group-016.md#canonical-1c06e483345c30fbc3c5b16d887bed0987d99e9c161a7faba79dc6b2b7f01f31)
- [default_pool.origin_servers.k8s_service.outside_network](resources--http_loadbalancer--reference--group-016.md#canonical-cf185ef09e8c0dc13d8fa58f69f66e1804a98fa63a9c98b9d75f8363defd5b45)
- [default_pool.origin_servers.k8s_service.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-b3e34f4a474f0153625579f9efd043ca3765bd9ecad139ca627212a8a03f7efe)
- [default_pool.origin_servers.k8s_service.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-b18f25eee537497ad5e3d59073d97aa92d8f358f28ff0050a204518d1c9f92b3)
- [default_pool.origin_servers.k8s_service.vk8s_networks](resources--http_loadbalancer--reference--group-016.md#canonical-e9d4477a5ef3fffceb08a01e5df112ad52bf3523bddc28abe7eff388abd8ccff)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1c06e483345c30fbc3c5b16d887bed0987d99e9c161a7faba79dc6b2b7f01f31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2e346c0803ae4965773e6d31ed3a4e98ff4e48478215e9b3b6cdbee49712acb"></a>

## default_pool.origin_servers.k8s_service.inside_network — default_pool.origin_servers.k8s_service.inside_network / bc9173b76d7d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7)
- default_pool.origin_servers.k8s_service.inside_network

<a id="canonical-207d8270773508c4596570f1123b7444c269ddb0b27cfdfd7b2335301ed3bea2"></a>

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

<a id="canonical-f64b4490ff44f73ce40cc81f83b617146f05404ae23d06aa4d20b7cf440fa1bc"></a>

## Direct properties — default_pool.origin_servers.k8s_service.inside_network / bc9173b76d7d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-42717b453f0a315b59990f8db33ab6abb4a65ff7f6ef2fda96e5490fc403118a"></a>

## Next pages — default_pool.origin_servers.k8s_service.inside_network / bc9173b76d7d / 4

- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-cf185ef09e8c0dc13d8fa58f69f66e1804a98fa63a9c98b9d75f8363defd5b45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb8ad6d9eaaf0c7a856a9396dc8c0437018b065cffab369bc5807e9bcd0f0d12"></a>

## default_pool.origin_servers.k8s_service.outside_network — default_pool.origin_servers.k8s_service.outside_network / a12194e0690a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7)
- default_pool.origin_servers.k8s_service.outside_network

<a id="canonical-947c70ec0d2c6e73a5696fbd774b6802c51ac1da6d6251d5ad671802b0d4482b"></a>

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

<a id="canonical-6a2a6bb23104dcb1afb27ca6a74540610a72c16a61fecac4956e96906b1aede7"></a>

## Direct properties — default_pool.origin_servers.k8s_service.outside_network / a12194e0690a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a9474553cb01ba07cc025a503bb2c5a418cff34cbe89a77cd7d06b7983736dd"></a>

## Next pages — default_pool.origin_servers.k8s_service.outside_network / a12194e0690a / 4

- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b3e34f4a474f0153625579f9efd043ca3765bd9ecad139ca627212a8a03f7efe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80e34527a5dff0ab55d4f1b8a20d66b0f2092ee98a03de544d65e2002d47a8fa"></a>

## default_pool.origin_servers.k8s_service.site_locator — default_pool.origin_servers.k8s_service.site_locator / 53596a7bc16b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7)
- default_pool.origin_servers.k8s_service.site_locator

<a id="canonical-2e76f91698ce5a49e866a2e6621a7ad06ac8a68482e09478fdaade064a0faccd"></a>

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

<a id="canonical-8215857c205215aeb2b4262dd4cd579a02e8147e8e6f5fb28dbde238188fd42c"></a>

## Direct properties — default_pool.origin_servers.k8s_service.site_locator / 53596a7bc16b / 3

- [site](resources--http_loadbalancer--reference--group-016.md#canonical-e1c2e24a2c9faddf875168de5ee19c1b77534f80bca22947f396e6b978154621): complete subsection reference.

- [virtual_site](resources--http_loadbalancer--reference--group-016.md#canonical-9acb055e75276348fcc3bbe61b66422021163690c6edd535194532f944769950): complete subsection reference.

<a id="canonical-3cc1012c54bde9acc6409c5888fce130e4c518d1d32de872b735c1386905cbcc"></a>

## Next pages — default_pool.origin_servers.k8s_service.site_locator / 53596a7bc16b / 4

- [default_pool.origin_servers.k8s_service.site_locator.site](resources--http_loadbalancer--reference--group-016.md#canonical-e1c2e24a2c9faddf875168de5ee19c1b77534f80bca22947f396e6b978154621)
- [default_pool.origin_servers.k8s_service.site_locator.virtual_site](resources--http_loadbalancer--reference--group-016.md#canonical-9acb055e75276348fcc3bbe61b66422021163690c6edd535194532f944769950)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e1c2e24a2c9faddf875168de5ee19c1b77534f80bca22947f396e6b978154621"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30a85e14eea8994a630d99b9d4d375d8132c608c45b1af17302e44d62df52ef8"></a>

## default_pool.origin_servers.k8s_service.site_locator.site — default_pool.origin_servers.k8s_service.site_locator.site / fac80465e3b1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7)
- [default_pool.origin_servers.k8s_service.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-b3e34f4a474f0153625579f9efd043ca3765bd9ecad139ca627212a8a03f7efe)
- default_pool.origin_servers.k8s_service.site_locator.site

<a id="canonical-a48f04e780d33161a83ba9855f9a7994789f4d3475ea0ee3bd62ad84444d88dd"></a>

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

<a id="canonical-c054cc2c7d445584a00d6b2880f9f435433961273cc29e73c5d4f68d2b38b4fb"></a>

## Direct properties — default_pool.origin_servers.k8s_service.site_locator.site / fac80465e3b1 / 3

<a id="canonical-4efeea5099f1db42ea05be0d768f8024543393eb2c4bdd170027b0d612d25253"></a>

<a id="canonical-64dde08237dd6a661f8630ea5a58056659c4baaab6c8fa270e78bd291a4b269a"></a>

## name property — default_pool.origin_servers.k8s_service.site_locator.site / fac80465e3b1 / 4

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

<a id="canonical-3f2412efe1076c30ba127bb93f79834711ebbe14acba392f4bf9f07884a7188e"></a>

<a id="canonical-107a9c110b82f2e68db9d79bbb1df441a98837e821f65482e7573ee7565d7553"></a>

## namespace property — default_pool.origin_servers.k8s_service.site_locator.site / fac80465e3b1 / 5

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

<a id="canonical-76335590acf48af9f625d60a886630ef3edca1af71e941aa326853d1f671cdf1"></a>

<a id="canonical-c60a3c289ae51cefa180c4156abb9f5a2c31a52dbce0272a79f7a7ef908a8479"></a>

## tenant property — default_pool.origin_servers.k8s_service.site_locator.site / fac80465e3b1 / 6

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

<a id="canonical-0ae1fd24e9078f1cdc4f7670aa72f38737e195a69510f6f518d021aff355731e"></a>

## Next pages — default_pool.origin_servers.k8s_service.site_locator.site / fac80465e3b1 / 7

- [default_pool.origin_servers.k8s_service.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-b3e34f4a474f0153625579f9efd043ca3765bd9ecad139ca627212a8a03f7efe)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9acb055e75276348fcc3bbe61b66422021163690c6edd535194532f944769950"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f06f9c52de5eb76fd4d470f4385257a685cace8d50146083dd8fe50f54ae94b"></a>

## default_pool.origin_servers.k8s_service.site_locator.virtual_site — default_pool.origin_servers.k8s_service.site_locator.virtual_site / ac09dcaf1cc4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7)
- [default_pool.origin_servers.k8s_service.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-b3e34f4a474f0153625579f9efd043ca3765bd9ecad139ca627212a8a03f7efe)
- default_pool.origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-892755fef2fac00616d3d46d73bd6aa2fd6189fb3ed1747b0cd50abcdda2f4f7"></a>

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

<a id="canonical-11f08225cabe4bf3b936a229fcec8779431bbb5d4575400300c617091c611529"></a>

## Direct properties — default_pool.origin_servers.k8s_service.site_locator.virtual_site / ac09dcaf1cc4 / 3

<a id="canonical-c2302ab1ef3854b5fb594bfab4d0486573d8b954574944fc3a433788d1aa41bf"></a>

<a id="canonical-68609f2e64b316fbb6ba968f10747d16cd9a24b7a1370e9fc6ee29301b486834"></a>

## name property — default_pool.origin_servers.k8s_service.site_locator.virtual_site / ac09dcaf1cc4 / 4

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

<a id="canonical-28c7426717c6c26d0e233413d034e80a7a8d512ce0101e4a31515b0e7b9ea0d8"></a>

<a id="canonical-bb83bbdc45fc631b08d7842c1b263e9adc12ea9902de95f4932da646f2bf5f86"></a>

## namespace property — default_pool.origin_servers.k8s_service.site_locator.virtual_site / ac09dcaf1cc4 / 5

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

<a id="canonical-266d4ef41b96eaf2545a782f591ad4f837c22f177f1e96927d81ea775ee9bcf2"></a>

<a id="canonical-862dcaf01171c10aa87e555afde2992b3a22588081d8e42f3081f0b97d6f01bb"></a>

## tenant property — default_pool.origin_servers.k8s_service.site_locator.virtual_site / ac09dcaf1cc4 / 6

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

<a id="canonical-809f05aaf4fd1c0bfe825b842987d13c5e4e84c4efa38c1cec16e1c8b26455d6"></a>

## Next pages — default_pool.origin_servers.k8s_service.site_locator.virtual_site / ac09dcaf1cc4 / 7

- [default_pool.origin_servers.k8s_service.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-b3e34f4a474f0153625579f9efd043ca3765bd9ecad139ca627212a8a03f7efe)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b18f25eee537497ad5e3d59073d97aa92d8f358f28ff0050a204518d1c9f92b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78d52883a4c772e33c9cd351ffd84033e0a673e58b4f1e4936f2dc2887b58ee8"></a>

## default_pool.origin_servers.k8s_service.snat_pool — default_pool.origin_servers.k8s_service.snat_pool / a1a01179ebf7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7)
- default_pool.origin_servers.k8s_service.snat_pool

<a id="canonical-e9850548d83298199168fdf16e277e1af771c243aa269ddbc38093c42d19d83d"></a>

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

<a id="canonical-1826aa97bacb2766e189e14fdbb6e35cf39d44fbd854d0b63f78f7b1756084df"></a>

## Direct properties — default_pool.origin_servers.k8s_service.snat_pool / a1a01179ebf7 / 3

- [no_snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-55e2bca03c9bb46268a244246b70468c20077bc08fcda8f715994834197e6a75): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-c46c8061014a8ae55f4859e3e3bfbe07fdf2316c8bcff517554c7568962daf57): complete subsection reference.

<a id="canonical-159ff34d0bd638148af8bd7fb3bf250d12a3f56e48e0d33e3ac01de9562c59c1"></a>

## Next pages — default_pool.origin_servers.k8s_service.snat_pool / a1a01179ebf7 / 4

- [default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-55e2bca03c9bb46268a244246b70468c20077bc08fcda8f715994834197e6a75)
- [default_pool.origin_servers.k8s_service.snat_pool.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-c46c8061014a8ae55f4859e3e3bfbe07fdf2316c8bcff517554c7568962daf57)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-55e2bca03c9bb46268a244246b70468c20077bc08fcda8f715994834197e6a75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-204d3cf20a54b40ebb83e38797735235e23fe63793e1e7a5a5aeff4f9943daf3"></a>

## default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool — default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool / 56d8262fbfc8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7)
- [default_pool.origin_servers.k8s_service.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-b18f25eee537497ad5e3d59073d97aa92d8f358f28ff0050a204518d1c9f92b3)
- default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-44331d41a9f44a296a22793f299a660802fb7a47c7d6bcad36610f100e4babeb"></a>

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

<a id="canonical-e2a4c82e872ec2a7de02116626f218037a14e7a1bc3278b73ae8a68e25a3d335"></a>

## Direct properties — default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool / 56d8262fbfc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c022284396255909ef6eb3c5ecf3d1cacf94847bad258c434cb966a829e30af9"></a>

## Next pages — default_pool.origin_servers.k8s_service.snat_pool.no_snat_pool / 56d8262fbfc8 / 4

- [default_pool.origin_servers.k8s_service.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-b18f25eee537497ad5e3d59073d97aa92d8f358f28ff0050a204518d1c9f92b3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c46c8061014a8ae55f4859e3e3bfbe07fdf2316c8bcff517554c7568962daf57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f85220c0c67d2e833e2da6b577cfd2693f121f8300afaf643932ddd768047dea"></a>

## default_pool.origin_servers.k8s_service.snat_pool.snat_pool — default_pool.origin_servers.k8s_service.snat_pool.snat_pool / ef28ef583739 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7)
- [default_pool.origin_servers.k8s_service.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-b18f25eee537497ad5e3d59073d97aa92d8f358f28ff0050a204518d1c9f92b3)
- default_pool.origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-f13d360e3c9355e5ccc0f71f154f264e2dccae75da048a390b644c3870475c0e"></a>

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

<a id="canonical-8c0373ecfad24ef41d98ebb56f1baea527803ba206d214fdce66a79547a86d07"></a>

## Direct properties — default_pool.origin_servers.k8s_service.snat_pool.snat_pool / ef28ef583739 / 3

<a id="canonical-7dc2d27c7e69ce9e4b806f4bc74ea61fcebbc680eb17a0e4346b310f303c8279"></a>

<a id="canonical-105148aa94da8e8341cf23e46711b6442b4b68e201bbaba7347596951e607a66"></a>

## prefixes property — default_pool.origin_servers.k8s_service.snat_pool.snat_pool / ef28ef583739 / 4

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

<a id="canonical-429d199a30912f54940940d9475e0b74c90879c1777fca971fb229db1d64c476"></a>

## Next pages — default_pool.origin_servers.k8s_service.snat_pool.snat_pool / ef28ef583739 / 5

- [default_pool.origin_servers.k8s_service.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-b18f25eee537497ad5e3d59073d97aa92d8f358f28ff0050a204518d1c9f92b3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e9d4477a5ef3fffceb08a01e5df112ad52bf3523bddc28abe7eff388abd8ccff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1eedc0af35b6ad79b399ec3fa9c7aece2cb002b0ec8c93f049b673f28a71bd60"></a>

## default_pool.origin_servers.k8s_service.vk8s_networks — default_pool.origin_servers.k8s_service.vk8s_networks / 6b4dfb27edab / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7)
- default_pool.origin_servers.k8s_service.vk8s_networks

<a id="canonical-368bc4413e43532a90b552eb84a065bb198f8819422ca2bf4cf49af31717162b"></a>

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

<a id="canonical-8061337ff5c95276da7a441f199aa3dbf63050c93822b7478c6a50ee7c00ad3b"></a>

## Direct properties — default_pool.origin_servers.k8s_service.vk8s_networks / 6b4dfb27edab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d3f70aa8bb2ebcd57c6503e13aa0f96d7434bbc3ec5ad95033cb0f43820119e3"></a>

## Next pages — default_pool.origin_servers.k8s_service.vk8s_networks / 6b4dfb27edab / 4

- [default_pool.origin_servers.k8s_service](resources--http_loadbalancer--reference--group-016.md#canonical-10b24e4b0b64c6557d4d20f24506fb25fa41b1ba7f98b8966adbed9be4f8f6e7)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edae249d437c3d8fc0aa47a7aa57123da22eb534ac20f14fc92c0a13aa9aadce"></a>

## default_pool.origin_servers.private_ip — default_pool.origin_servers.private_ip / 22fcf741d5ee / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- default_pool.origin_servers.private_ip

<a id="canonical-d056f265fe5c6d3fc6e8089b0762e34ca8c6c1de57dd3901977bdb6226a30a0e"></a>

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

<a id="canonical-304af2b1c49a958be6ff70e7d7b9ccf59ffb853750d95d27e72248d44567c44b"></a>

## Direct properties — default_pool.origin_servers.private_ip / 22fcf741d5ee / 3

- [inside_network](resources--http_loadbalancer--reference--group-016.md#canonical-2ff6c6223e6d1034c39d064ee3e594ec6613e20e6d80fa3715cdb5cad2d76952): complete subsection reference.

<a id="canonical-6aec524c2506fee07060ce31457ff7f136ac7358d5890f3815f6cb2a6729b5a6"></a>

<a id="canonical-a44a5c0872eed55f0bec599c7a6286c356cb023c338fb51595d9af52eeb99a91"></a>

## ip property — default_pool.origin_servers.private_ip / 22fcf741d5ee / 4

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

- [outside_network](resources--http_loadbalancer--reference--group-016.md#canonical-668452e6bf308b7893ae45df344ac80cab695dfbc8e9ea176539bc6c7635d42c): complete subsection reference.

- [segment](resources--http_loadbalancer--reference--group-016.md#canonical-f5337e14fc461571edab3b9bc058276c0100b6d18a7f7ba621acdca0b24678e0): complete subsection reference.

- [site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-ac4b896b5d82bbc6349631af95578a4e1397c6493a4471e9e95ba47a5cea2e53): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-42f3e9afe8fde5bb95a0069ec5725d28b987400cb9590af9c830988b2afd0c91): complete subsection reference.

<a id="canonical-b2283c90f0ed31ba639ba2a60c1dd8092e6102270f5a20077c94997d9214f32f"></a>

## Next pages — default_pool.origin_servers.private_ip / 22fcf741d5ee / 5

- [default_pool.origin_servers.private_ip.inside_network](resources--http_loadbalancer--reference--group-016.md#canonical-2ff6c6223e6d1034c39d064ee3e594ec6613e20e6d80fa3715cdb5cad2d76952)
- [default_pool.origin_servers.private_ip.outside_network](resources--http_loadbalancer--reference--group-016.md#canonical-668452e6bf308b7893ae45df344ac80cab695dfbc8e9ea176539bc6c7635d42c)
- [default_pool.origin_servers.private_ip.segment](resources--http_loadbalancer--reference--group-016.md#canonical-f5337e14fc461571edab3b9bc058276c0100b6d18a7f7ba621acdca0b24678e0)
- [default_pool.origin_servers.private_ip.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-ac4b896b5d82bbc6349631af95578a4e1397c6493a4471e9e95ba47a5cea2e53)
- [default_pool.origin_servers.private_ip.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-42f3e9afe8fde5bb95a0069ec5725d28b987400cb9590af9c830988b2afd0c91)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2ff6c6223e6d1034c39d064ee3e594ec6613e20e6d80fa3715cdb5cad2d76952"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47835aabfc6db8e9a477800399e510d5201724ed5de471cc49e8997272eca6d8"></a>

## default_pool.origin_servers.private_ip.inside_network — default_pool.origin_servers.private_ip.inside_network / 69594620e896 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09)
- default_pool.origin_servers.private_ip.inside_network

<a id="canonical-28bbad1d5a071e87bf6aab7195d0218fb10893c2ee9a4a96a0464b2667d76ed9"></a>

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

<a id="canonical-c243e2ec891828f02f9bb638f3b1df6c81aa44fb987a228a715ee3e03a7a2b55"></a>

## Direct properties — default_pool.origin_servers.private_ip.inside_network / 69594620e896 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a9fa7a99dca97fdbdb36e6aac158e49b71403e069302b54f4347e3f2db4d3413"></a>

## Next pages — default_pool.origin_servers.private_ip.inside_network / 69594620e896 / 4

- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-668452e6bf308b7893ae45df344ac80cab695dfbc8e9ea176539bc6c7635d42c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35698255be1dd428ce8e55cebbea2341789c7b95658e434f8991c1f4d7b9c6e4"></a>

## default_pool.origin_servers.private_ip.outside_network — default_pool.origin_servers.private_ip.outside_network / dc9068addbad / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09)
- default_pool.origin_servers.private_ip.outside_network

<a id="canonical-e7a0e94e84346ab5e300de13ad495c7a523adb5b2a32adee2e1881e921aad231"></a>

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

<a id="canonical-ecdc6d0702f063f153c06f36460dc7a6ea8a3b1fe6eba10ef02f3282fd6e51af"></a>

## Direct properties — default_pool.origin_servers.private_ip.outside_network / dc9068addbad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1e173e364c829c560d36d9d325134ad9d1bce75397761782259ef0054c3a0e23"></a>

## Next pages — default_pool.origin_servers.private_ip.outside_network / dc9068addbad / 4

- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f5337e14fc461571edab3b9bc058276c0100b6d18a7f7ba621acdca0b24678e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d484d3406e9ca2df00b4c9d370f3a425deda485e3cf6b3f3f41b347e32c23aa"></a>

## default_pool.origin_servers.private_ip.segment — default_pool.origin_servers.private_ip.segment / 4b389984adc8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09)
- default_pool.origin_servers.private_ip.segment

<a id="canonical-195ae2b9fff29e81dbc3adb1602ee909d8f2d5d67ba6402a3b29f46aff9010d3"></a>

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

<a id="canonical-7eae43662b0281f5b23e4d385793964069e87fa689bd68a3d7b2914839d62a1f"></a>

## Direct properties — default_pool.origin_servers.private_ip.segment / 4b389984adc8 / 3

<a id="canonical-ff7fd9861c81f4963bafb4805164e525c88fcb6440f81edd2944e97663afac49"></a>

<a id="canonical-e6d946b29085f137af17f673901ef7e91f7d7ebe89a2c8ae56800502eb48e2fa"></a>

## name property — default_pool.origin_servers.private_ip.segment / 4b389984adc8 / 4

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

<a id="canonical-a12ce570a1a606ff5c12129ac970c91e978644434c440429d2ae766b4f4d958f"></a>

<a id="canonical-fde2e14f4b8879bd115a745893fa43684102cd965651fbf3e95f5547e3db4a03"></a>

## namespace property — default_pool.origin_servers.private_ip.segment / 4b389984adc8 / 5

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

<a id="canonical-f6b1b357c313017a40c5b432fbf5f9f15b692f7c61ae582e871b17130d23b3c9"></a>

<a id="canonical-9a2f42f14a4c29261e73e586857fa25819c3ffabea1f3d58d8146eed314e3a62"></a>

## tenant property — default_pool.origin_servers.private_ip.segment / 4b389984adc8 / 6

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

<a id="canonical-1eefa6b2366d6602232c74da0e87ca4c10b7448becff23dde942ac1847c04147"></a>

## Next pages — default_pool.origin_servers.private_ip.segment / 4b389984adc8 / 7

- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ac4b896b5d82bbc6349631af95578a4e1397c6493a4471e9e95ba47a5cea2e53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c1ff81e63082ffefde162a6d5b9319ebf897a4ab0ad3f6e16c082034127d3af"></a>

## default_pool.origin_servers.private_ip.site_locator — default_pool.origin_servers.private_ip.site_locator / 11348e61026d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09)
- default_pool.origin_servers.private_ip.site_locator

<a id="canonical-2507ade0b1596da82aa65c865b7d09051d932ac28ff24ab13d66c7f2faf98a05"></a>

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

<a id="canonical-76bd34e28333ad40be393c014acc980a97fa06f0210e00e7826b504bb890ece7"></a>

## Direct properties — default_pool.origin_servers.private_ip.site_locator / 11348e61026d / 3

- [site](resources--http_loadbalancer--reference--group-016.md#canonical-58d4c04c643a014ac640eb71061c2473c528be8cc9f117791d807bf8b060b5a9): complete subsection reference.

- [virtual_site](resources--http_loadbalancer--reference--group-016.md#canonical-e1b31624bff9bf9c30cc78d2507a8f4f790e978387698dd5f65f37847a4d1175): complete subsection reference.

<a id="canonical-77f875f4796e306124ab2e62ff5f8b88e7fe68ed94215ec81bb1624d1d563af1"></a>

## Next pages — default_pool.origin_servers.private_ip.site_locator / 11348e61026d / 4

- [default_pool.origin_servers.private_ip.site_locator.site](resources--http_loadbalancer--reference--group-016.md#canonical-58d4c04c643a014ac640eb71061c2473c528be8cc9f117791d807bf8b060b5a9)
- [default_pool.origin_servers.private_ip.site_locator.virtual_site](resources--http_loadbalancer--reference--group-016.md#canonical-e1b31624bff9bf9c30cc78d2507a8f4f790e978387698dd5f65f37847a4d1175)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-58d4c04c643a014ac640eb71061c2473c528be8cc9f117791d807bf8b060b5a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20813e714b8c7f00e4bb73a370bf4eb4e1d7c4b2ef71d0d2d7a4c9e28116294d"></a>

## default_pool.origin_servers.private_ip.site_locator.site — default_pool.origin_servers.private_ip.site_locator.site / 6e6b6c7ee49f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09)
- [default_pool.origin_servers.private_ip.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-ac4b896b5d82bbc6349631af95578a4e1397c6493a4471e9e95ba47a5cea2e53)
- default_pool.origin_servers.private_ip.site_locator.site

<a id="canonical-32a83de3f26a61293eb1a623aa0f790433e2578dd35a0d0182f4fbc0fac8f105"></a>

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

<a id="canonical-0222ec84260bc599917729f82c6b5235d4a9580ed196b7718a3bd9b6b84c67b0"></a>

## Direct properties — default_pool.origin_servers.private_ip.site_locator.site / 6e6b6c7ee49f / 3

<a id="canonical-69f813bb5711588ccb5fa036962a7aa6c5443e779a70424e07a954f080716681"></a>

<a id="canonical-88db5c4bef191ab754db0b91a9e0d78051f3ed4328be76f8c2c9b1192c92752b"></a>

## name property — default_pool.origin_servers.private_ip.site_locator.site / 6e6b6c7ee49f / 4

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

<a id="canonical-fa410360741c1a49a41ff91a7178d61aecbf328e8b254f194b544f20c0446d61"></a>

<a id="canonical-7534dd46470a6ee1d6cff9b3e8ef3dd8d70b399e61dacd99d0585e83e6be22a7"></a>

## namespace property — default_pool.origin_servers.private_ip.site_locator.site / 6e6b6c7ee49f / 5

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

<a id="canonical-e64fdb0dea24add78b6c94893827867a8c819d2f617b7f30336d7f610d29e4c4"></a>

<a id="canonical-b1f1a1f4cdd033910d0580e36df45182c7a540e80d9ced8931605f0787e0a02b"></a>

## tenant property — default_pool.origin_servers.private_ip.site_locator.site / 6e6b6c7ee49f / 6

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

<a id="canonical-51025e9dd0bf4613f36cff0e1d407375dfee0edf637381fdc2d18e1e81f2c167"></a>

## Next pages — default_pool.origin_servers.private_ip.site_locator.site / 6e6b6c7ee49f / 7

- [default_pool.origin_servers.private_ip.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-ac4b896b5d82bbc6349631af95578a4e1397c6493a4471e9e95ba47a5cea2e53)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e1b31624bff9bf9c30cc78d2507a8f4f790e978387698dd5f65f37847a4d1175"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79ddeae3026baa1f1965decf5d0a5284395e5d1f39f473c1d16cc46ae98e3dd8"></a>

## default_pool.origin_servers.private_ip.site_locator.virtual_site — default_pool.origin_servers.private_ip.site_locator.virtual_site / 31f732ab6d91 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09)
- [default_pool.origin_servers.private_ip.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-ac4b896b5d82bbc6349631af95578a4e1397c6493a4471e9e95ba47a5cea2e53)
- default_pool.origin_servers.private_ip.site_locator.virtual_site

<a id="canonical-26df5750c1b1526bfe3c3a2119ec4289b2d5583658d29c681dc54711f084aef4"></a>

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

<a id="canonical-bc38d2800aee595d22267c782abc81837987c8f16b45bea09a741a2982ed832b"></a>

## Direct properties — default_pool.origin_servers.private_ip.site_locator.virtual_site / 31f732ab6d91 / 3

<a id="canonical-470c307ad2a4febc1896035a61fe461cbcbd2018134cb6dda4f04c73e3456d34"></a>

<a id="canonical-220096c0b745242165a0c2f5e8b75278b9de235578a8c5192a5677a5ac312db3"></a>

## name property — default_pool.origin_servers.private_ip.site_locator.virtual_site / 31f732ab6d91 / 4

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

<a id="canonical-4989ea8f39aa68afe861d92a3ca42ab72becf085ea3cf6fc163aa21b35dea7db"></a>

<a id="canonical-bfea778079f51d110bcf6a4c09b5c89d821cc5cef01467b349ed643f269f3fb6"></a>

## namespace property — default_pool.origin_servers.private_ip.site_locator.virtual_site / 31f732ab6d91 / 5

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

<a id="canonical-f333c08eb9f46974dff0b50cc20c82af4f65c3f3339c84b108f96d4a6b9f10cd"></a>

<a id="canonical-6b8a90aa97873bda00988306a5184ae99c4d90acf61b2a95604090f70a777b77"></a>

## tenant property — default_pool.origin_servers.private_ip.site_locator.virtual_site / 31f732ab6d91 / 6

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

<a id="canonical-290ce75954ed007931e8a46adda084a1acfa1872369e06abae32f602488ee4c5"></a>

## Next pages — default_pool.origin_servers.private_ip.site_locator.virtual_site / 31f732ab6d91 / 7

- [default_pool.origin_servers.private_ip.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-ac4b896b5d82bbc6349631af95578a4e1397c6493a4471e9e95ba47a5cea2e53)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-42f3e9afe8fde5bb95a0069ec5725d28b987400cb9590af9c830988b2afd0c91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e5ff01b4f19c1c58a707f4d2436951d058c3c9adebddcdf1888549b9a893f50"></a>

## default_pool.origin_servers.private_ip.snat_pool — default_pool.origin_servers.private_ip.snat_pool / 2caa5326827c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09)
- default_pool.origin_servers.private_ip.snat_pool

<a id="canonical-78a7bd69e355ee06de426b9cb67d9d948fd138922786ffe3cf2fee4f08452905"></a>

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

<a id="canonical-8e636be386ec42ed12156ba8830ae5723dd52e0fd998dbd16866ab5354138d79"></a>

## Direct properties — default_pool.origin_servers.private_ip.snat_pool / 2caa5326827c / 3

- [no_snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-0fd506dd7470e844be364ed8a3feaa4a06370626bb0d5a6fca951480c2d1b021): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-a1db855a3acad3b9af779248ea21024732d9d21188ec29f4647f8c71f7e5b75a): complete subsection reference.

<a id="canonical-0de32bfd32429f24f4aae93a8c290d111a2ab06c7b35db8d5ad7871ce5fb9a65"></a>

## Next pages — default_pool.origin_servers.private_ip.snat_pool / 2caa5326827c / 4

- [default_pool.origin_servers.private_ip.snat_pool.no_snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-0fd506dd7470e844be364ed8a3feaa4a06370626bb0d5a6fca951480c2d1b021)
- [default_pool.origin_servers.private_ip.snat_pool.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-a1db855a3acad3b9af779248ea21024732d9d21188ec29f4647f8c71f7e5b75a)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0fd506dd7470e844be364ed8a3feaa4a06370626bb0d5a6fca951480c2d1b021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecf5b65d4bcd450db693bafdee82b4abfcb08e35fa677bc8d150b8834437caaf"></a>

## default_pool.origin_servers.private_ip.snat_pool.no_snat_pool — default_pool.origin_servers.private_ip.snat_pool.no_snat_pool / 55035a79d393 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09)
- [default_pool.origin_servers.private_ip.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-42f3e9afe8fde5bb95a0069ec5725d28b987400cb9590af9c830988b2afd0c91)
- default_pool.origin_servers.private_ip.snat_pool.no_snat_pool

<a id="canonical-aea48009708c410a13b8eca179990cda754b4ac2737f44e79908bba65a2562d8"></a>

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

<a id="canonical-bbfd2d16f7ef787b7825177a37b234eafb543b64a7251ec6c2a3afff134c0a20"></a>

## Direct properties — default_pool.origin_servers.private_ip.snat_pool.no_snat_pool / 55035a79d393 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ce0464abe36b2ba3c9ca065e0f9e1463698d37fc7d8e7489dbf2060f3e151759"></a>

## Next pages — default_pool.origin_servers.private_ip.snat_pool.no_snat_pool / 55035a79d393 / 4

- [default_pool.origin_servers.private_ip.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-42f3e9afe8fde5bb95a0069ec5725d28b987400cb9590af9c830988b2afd0c91)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a1db855a3acad3b9af779248ea21024732d9d21188ec29f4647f8c71f7e5b75a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a28bba27f6abe30af570f6fee4eec43ed4db90b63ebfd8101be160bb87c42b46"></a>

## default_pool.origin_servers.private_ip.snat_pool.snat_pool — default_pool.origin_servers.private_ip.snat_pool.snat_pool / 921477447db1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-90b9de47259be9c2c38b0bc9518101d73940e584413783175a6ece8a0d700d09)
- [default_pool.origin_servers.private_ip.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-42f3e9afe8fde5bb95a0069ec5725d28b987400cb9590af9c830988b2afd0c91)
- default_pool.origin_servers.private_ip.snat_pool.snat_pool

<a id="canonical-aeb979bdcda868c781f6e23cb170dc3e3cf967eb2d168c176aa8751a624e7506"></a>

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

<a id="canonical-8f8fe380281130f963c859359c27d3afe72b803a8cb0fd3c4c17348c210f96e5"></a>

## Direct properties — default_pool.origin_servers.private_ip.snat_pool.snat_pool / 921477447db1 / 3

<a id="canonical-59caadbadbe5af3e1ff6f2fc362e4ba6558d895f36bd14b2fa99442539edbf09"></a>

<a id="canonical-18b4a01f3452c58133314fc0a4d92232570aa9be840a3cbc887459099cd8df4e"></a>

## prefixes property — default_pool.origin_servers.private_ip.snat_pool.snat_pool / 921477447db1 / 4

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

<a id="canonical-90fbfaa2477b628f1d752ba87a322ddd6fd34d8a5348bf21694a1ec197658398"></a>

## Next pages — default_pool.origin_servers.private_ip.snat_pool.snat_pool / 921477447db1 / 5

- [default_pool.origin_servers.private_ip.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-42f3e9afe8fde5bb95a0069ec5725d28b987400cb9590af9c830988b2afd0c91)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-582e99d89bcdddd8bc61a76730f087d8265b7bf713442b1fe8374ab27866802c"></a>

## default_pool.origin_servers.private_name — default_pool.origin_servers.private_name / 7d5c18f3302d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- default_pool.origin_servers.private_name

<a id="canonical-a94f8102683667230c3207c0781b765d569bdce45ed56619c6d724a9638c97d6"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with private or public DNS name and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name"),
  validators.ConflictingObjectAttributes("inside_network",
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]"
}
```

Terraform syntax:

```terraform
private_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-457d0098be15ccf83a6ba73d910e18cb81acbad8a8a0b37f0b4e3339acce3b40"></a>

## Direct properties — default_pool.origin_servers.private_name / 7d5c18f3302d / 3

<a id="canonical-b908aa7e57b54b3bfc3cf0bfe3db3dce36011db68b1b18db2626a80c36510eb5"></a>

<a id="canonical-cc7d612b335c18d626d3a58498330800d87368bc2f66b1b8604f00f0f6837e2c"></a>

## dns_name property — default_pool.origin_servers.private_name / 7d5c18f3302d / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

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

- [inside_network](resources--http_loadbalancer--reference--group-016.md#canonical-46f12b70f0958ddb4f0300f58fa0a0f4d41ecc3fd500b70e4ece0dbe2bd155de): complete subsection reference.

- [outside_network](resources--http_loadbalancer--reference--group-016.md#canonical-35c5ab1ce2ebca50c4af22e7c042d2ff4058e4c25e6271e6ddb8fa2f06aa6589): complete subsection reference.

<a id="canonical-6903183afc16def22da2875ea7fca44671ea3dffbf80f6ca4e37353bd1b5be24"></a>

<a id="canonical-e81dcba078823403f9553b2f6a3288102d768d06dba531d1cc6bade2cfe37260"></a>

## refresh_interval property — default_pool.origin_servers.private_name / 7d5c18f3302d / 5

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

- [segment](resources--http_loadbalancer--reference--group-016.md#canonical-fae7459bfd8e45f19c71e7da8e4dc5d692d8f32327b2e70907baf48ef923a32f): complete subsection reference.

- [site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-25b86ad541b435cb669813ed0d775359919ad58b3ea47e88e4727d2601688efc): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3670a6574f0e6f580b452826761d784823d489db2b3510fceed14161b61ac07c): complete subsection reference.

<a id="canonical-669a8f1b14532f10474b768e651c21dc5b29f8b1b130794713094c42fdc967e8"></a>

## Next pages — default_pool.origin_servers.private_name / 7d5c18f3302d / 6

- [default_pool.origin_servers.private_name.inside_network](resources--http_loadbalancer--reference--group-016.md#canonical-46f12b70f0958ddb4f0300f58fa0a0f4d41ecc3fd500b70e4ece0dbe2bd155de)
- [default_pool.origin_servers.private_name.outside_network](resources--http_loadbalancer--reference--group-016.md#canonical-35c5ab1ce2ebca50c4af22e7c042d2ff4058e4c25e6271e6ddb8fa2f06aa6589)
- [default_pool.origin_servers.private_name.segment](resources--http_loadbalancer--reference--group-016.md#canonical-fae7459bfd8e45f19c71e7da8e4dc5d692d8f32327b2e70907baf48ef923a32f)
- [default_pool.origin_servers.private_name.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-25b86ad541b435cb669813ed0d775359919ad58b3ea47e88e4727d2601688efc)
- [default_pool.origin_servers.private_name.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3670a6574f0e6f580b452826761d784823d489db2b3510fceed14161b61ac07c)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-46f12b70f0958ddb4f0300f58fa0a0f4d41ecc3fd500b70e4ece0dbe2bd155de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ada472851c1354ad90f35109bf1d783b352da95019d9109bda8699f6959d0010"></a>

## default_pool.origin_servers.private_name.inside_network — default_pool.origin_servers.private_name.inside_network / a5a538a30b00 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e)
- default_pool.origin_servers.private_name.inside_network

<a id="canonical-21c15f897224ce7bff2678af60243acc4a9b89b4f72e574a1f54ba9cf71a0ed2"></a>

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

<a id="canonical-d9969bc033a83c4084649a33d79fb12fdac9b82350bfbdcf5a6ec07bd5fe7f7e"></a>

## Direct properties — default_pool.origin_servers.private_name.inside_network / a5a538a30b00 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0672f4554e4d2b62f77457714b00301a60653697a3d576d2eceb748349d86a87"></a>

## Next pages — default_pool.origin_servers.private_name.inside_network / a5a538a30b00 / 4

- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-35c5ab1ce2ebca50c4af22e7c042d2ff4058e4c25e6271e6ddb8fa2f06aa6589"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bafcf637835ac6989ebe7bc5905397127e2e4a6e6828abb2f6c3b0047af666e"></a>

## default_pool.origin_servers.private_name.outside_network — default_pool.origin_servers.private_name.outside_network / 6fa15ca148d7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e)
- default_pool.origin_servers.private_name.outside_network

<a id="canonical-73bbc3647360d00ed9479991c5313f6130249b9d7b6ab5f3320a01879cfb332a"></a>

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

<a id="canonical-92835edcd69281a6246fd7146d7a706a6f06b6cfb97646f26fa65e226ee4fb0b"></a>

## Direct properties — default_pool.origin_servers.private_name.outside_network / 6fa15ca148d7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dbb269fa42aaf98fba5407e260b813a926b4e865c8d0bc27c4a4c23ae177ba72"></a>

## Next pages — default_pool.origin_servers.private_name.outside_network / 6fa15ca148d7 / 4

- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-fae7459bfd8e45f19c71e7da8e4dc5d692d8f32327b2e70907baf48ef923a32f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5fef60ca54269f6fe92f5b7ddebd9c203cc6db40867467ec71feefc6b62aa19"></a>

## default_pool.origin_servers.private_name.segment — default_pool.origin_servers.private_name.segment / 80ad0bcb9069 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e)
- default_pool.origin_servers.private_name.segment

<a id="canonical-ca2730cd9bbccc8285a04fddcac17a1e29c4b7c2f888aa9f8e2e741f5f9c59c0"></a>

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

<a id="canonical-51cb10d8ab762c8a79669b8977433e36c3891ce90659df235741f30e4dd4f3d9"></a>

## Direct properties — default_pool.origin_servers.private_name.segment / 80ad0bcb9069 / 3

<a id="canonical-dfa0a31a815873d831f9f85729dbf7bbc315f97ca1df986c0c805a14a14aecf6"></a>

<a id="canonical-2430fb1e389942c8ed74290f2897d1578daf6ffaf937b32cd376f0678375ccf2"></a>

## name property — default_pool.origin_servers.private_name.segment / 80ad0bcb9069 / 4

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

<a id="canonical-dd38b44051dda2eee20afc11e6a4bf453b56d09731b84225abcb3caf8a9b2492"></a>

<a id="canonical-67484ba0fac9adc9412ac124e630c69b04718abff520046f38aed7b110b3b006"></a>

## namespace property — default_pool.origin_servers.private_name.segment / 80ad0bcb9069 / 5

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

<a id="canonical-d786463618da80b9d38c2bd604ddbbe6f96b08598d8580d0b3c2a4856f946c4d"></a>

<a id="canonical-ee941b347290e492b8b34f721bd12bbb1e571a8399f4b898c74f9d33ccaa9ddf"></a>

## tenant property — default_pool.origin_servers.private_name.segment / 80ad0bcb9069 / 6

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

<a id="canonical-67063403abdd5d83fff0afcbc7b97863b756dcb0058111c8fa2ff33b6a9f881b"></a>

## Next pages — default_pool.origin_servers.private_name.segment / 80ad0bcb9069 / 7

- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-25b86ad541b435cb669813ed0d775359919ad58b3ea47e88e4727d2601688efc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5940b8abd62e7361505995f0ace1106f0eb2fb6da3a5b3036a899cc92727ae45"></a>

## default_pool.origin_servers.private_name.site_locator — default_pool.origin_servers.private_name.site_locator / 3c766d7b0cc5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e)
- default_pool.origin_servers.private_name.site_locator

<a id="canonical-328a96914804326f0e2daa85692166939e13825298560139300ac369acefe281"></a>

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

<a id="canonical-9dcb0e68f15b30e3e28d6bfee3b57737d8ea9c302bb03626e211516392547792"></a>

## Direct properties — default_pool.origin_servers.private_name.site_locator / 3c766d7b0cc5 / 3

- [site](resources--http_loadbalancer--reference--group-016.md#canonical-7edc562244088375b3e38d126d7e44dd90154f6e98f562a4d1a269c34bea0ccf): complete subsection reference.

- [virtual_site](resources--http_loadbalancer--reference--group-016.md#canonical-dc8031a6a099bd15d644ed7c86e37f698475eeb4296cd57e5408721c3922b120): complete subsection reference.

<a id="canonical-c496d085e2b28a43f67971b78b23ec4e1fad7250b9145f4c88a0735d8eafd85b"></a>

## Next pages — default_pool.origin_servers.private_name.site_locator / 3c766d7b0cc5 / 4

- [default_pool.origin_servers.private_name.site_locator.site](resources--http_loadbalancer--reference--group-016.md#canonical-7edc562244088375b3e38d126d7e44dd90154f6e98f562a4d1a269c34bea0ccf)
- [default_pool.origin_servers.private_name.site_locator.virtual_site](resources--http_loadbalancer--reference--group-016.md#canonical-dc8031a6a099bd15d644ed7c86e37f698475eeb4296cd57e5408721c3922b120)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7edc562244088375b3e38d126d7e44dd90154f6e98f562a4d1a269c34bea0ccf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6cdb6814bf9fb4fb0fc128dd5025bb360b4637d5c4a1607ac78acffeb846061"></a>

## default_pool.origin_servers.private_name.site_locator.site — default_pool.origin_servers.private_name.site_locator.site / 344996781693 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e)
- [default_pool.origin_servers.private_name.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-25b86ad541b435cb669813ed0d775359919ad58b3ea47e88e4727d2601688efc)
- default_pool.origin_servers.private_name.site_locator.site

<a id="canonical-c5ba605e06606e39e3ec1dac7e8f744c1a6b42e457495da3470684a55b336411"></a>

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

<a id="canonical-668a7327f741d7555a8bb004e467e02970f7282850430c2c3f086a9fe18ed0d3"></a>

## Direct properties — default_pool.origin_servers.private_name.site_locator.site / 344996781693 / 3

<a id="canonical-5610959be14a2cb92839dbdf34899f4b1031384f23b0b1d41ac5c55fbe2990e3"></a>

<a id="canonical-519451ba75fd637a126cea70a3fc601d42b1597e704e78ad8ccf8931714b547a"></a>

## name property — default_pool.origin_servers.private_name.site_locator.site / 344996781693 / 4

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

<a id="canonical-62c31a90862a9fc71c48cc64e550d8855672dcedf1bc0cd0c97011d8116daedb"></a>

<a id="canonical-d039be4859a62f383f45c262f07a200f4e6a3c0684396362dd4117bd81210f53"></a>

## namespace property — default_pool.origin_servers.private_name.site_locator.site / 344996781693 / 5

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

<a id="canonical-f9be53e95636a744157fe5994881f4e604f865b80463ae9d76499c8a397664ce"></a>

<a id="canonical-94b2ef3a968a20f24ea31d6cd5db54c5dd753d5d899db97ae59e46cf297472d3"></a>

## tenant property — default_pool.origin_servers.private_name.site_locator.site / 344996781693 / 6

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

<a id="canonical-88662d88ff94833544a43f5385231827350c14c1aafad51549658070cefbc443"></a>

## Next pages — default_pool.origin_servers.private_name.site_locator.site / 344996781693 / 7

- [default_pool.origin_servers.private_name.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-25b86ad541b435cb669813ed0d775359919ad58b3ea47e88e4727d2601688efc)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-dc8031a6a099bd15d644ed7c86e37f698475eeb4296cd57e5408721c3922b120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eba7a926d5e8ec80995d6ffa1e2aa80427c46c5897158ce4f2864d46f448b512"></a>

## default_pool.origin_servers.private_name.site_locator.virtual_site — default_pool.origin_servers.private_name.site_locator.virtual_site / 0f5e0eabab65 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e)
- [default_pool.origin_servers.private_name.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-25b86ad541b435cb669813ed0d775359919ad58b3ea47e88e4727d2601688efc)
- default_pool.origin_servers.private_name.site_locator.virtual_site

<a id="canonical-3292f769ee632cf45f3d3561f0166a521110b5b9b9a827872b093848a71b7741"></a>

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

<a id="canonical-d1cf9def7e793ba588b7e726ea41f7485b047cde43916ec1d8ea65871be406f9"></a>

## Direct properties — default_pool.origin_servers.private_name.site_locator.virtual_site / 0f5e0eabab65 / 3

<a id="canonical-ce90666508299ed1c25a0717f95c96ef0edceaaa756050fdbb42ecb8da6fe9cd"></a>

<a id="canonical-0ba7d4ed290a0c8dd6e74dd17ed160743f83d6d7de7cf408d7dd8fabb0360d11"></a>

## name property — default_pool.origin_servers.private_name.site_locator.virtual_site / 0f5e0eabab65 / 4

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

<a id="canonical-e0c9a034e88ae2ae092e8d21dfbb23285e741388692d782d9fbe65e3ba7d1158"></a>

<a id="canonical-51f2ea310586cf5ed87cf97b53ae3181d6b8207834a117c31481e72217097991"></a>

## namespace property — default_pool.origin_servers.private_name.site_locator.virtual_site / 0f5e0eabab65 / 5

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

<a id="canonical-9681c69ffff332dbbbac0333f966a437fdbbc0d870b26bac53a44090cd0efe81"></a>

<a id="canonical-7dcc323f83c53eea8c41256fc9bafb28b98c1c93fa9451c1606566176300b349"></a>

## tenant property — default_pool.origin_servers.private_name.site_locator.virtual_site / 0f5e0eabab65 / 6

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

<a id="canonical-fe7cb2a3b7fc1eaece586442badde761a87a891ef1177ca92d85cddad0c4f69e"></a>

## Next pages — default_pool.origin_servers.private_name.site_locator.virtual_site / 0f5e0eabab65 / 7

- [default_pool.origin_servers.private_name.site_locator](resources--http_loadbalancer--reference--group-016.md#canonical-25b86ad541b435cb669813ed0d775359919ad58b3ea47e88e4727d2601688efc)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3670a6574f0e6f580b452826761d784823d489db2b3510fceed14161b61ac07c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1e08713a83b30a6f30944b956c094f596b9c4a4f02345f8eaec9d616f2ffe34"></a>

## default_pool.origin_servers.private_name.snat_pool — default_pool.origin_servers.private_name.snat_pool / 7e97c36ad0df / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e)
- default_pool.origin_servers.private_name.snat_pool

<a id="canonical-b8c58b11794679942da70f81cb6e0513697f640efaed22b42c1dd775d7947fbd"></a>

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

<a id="canonical-9a55211a6ac111f103970af016c830fadd55398741c36c60e2af0556499ed470"></a>

## Direct properties — default_pool.origin_servers.private_name.snat_pool / 7e97c36ad0df / 3

- [no_snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-69d4f2540ab6907affdb9810a8d3b2debbeea79ad1bc73e81f5f339dd4830232): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-0b87cdcfc6d6804bc853cb595b09fc3f2ccf053ea3021248545c3424b1b3dd47): complete subsection reference.

<a id="canonical-c23c5b44b12bc2ddd785060016491a2bec24c1ecb64590687b2038aeeb00079c"></a>

## Next pages — default_pool.origin_servers.private_name.snat_pool / 7e97c36ad0df / 4

- [default_pool.origin_servers.private_name.snat_pool.no_snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-69d4f2540ab6907affdb9810a8d3b2debbeea79ad1bc73e81f5f339dd4830232)
- [default_pool.origin_servers.private_name.snat_pool.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-0b87cdcfc6d6804bc853cb595b09fc3f2ccf053ea3021248545c3424b1b3dd47)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-69d4f2540ab6907affdb9810a8d3b2debbeea79ad1bc73e81f5f339dd4830232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82e6572d0e478fe686390609757898b33a5f30430a8bbc61050d9cd0879867c1"></a>

## default_pool.origin_servers.private_name.snat_pool.no_snat_pool — default_pool.origin_servers.private_name.snat_pool.no_snat_pool / fdc98bb59423 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e)
- [default_pool.origin_servers.private_name.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3670a6574f0e6f580b452826761d784823d489db2b3510fceed14161b61ac07c)
- default_pool.origin_servers.private_name.snat_pool.no_snat_pool

<a id="canonical-c41755a03cf5abe4236cb695edba8b047fca739ec7a042d578057bfc71bd437e"></a>

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

<a id="canonical-903c4cf1d76d8ac8536b2cacef2249a6faf95f024a5ab04b4e4a1feeab8c5926"></a>

## Direct properties — default_pool.origin_servers.private_name.snat_pool.no_snat_pool / fdc98bb59423 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe9adf03f23d981a59d11c1064df6f9f137b66b2e5e61e1d11000bb3076e5393"></a>

## Next pages — default_pool.origin_servers.private_name.snat_pool.no_snat_pool / fdc98bb59423 / 4

- [default_pool.origin_servers.private_name.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3670a6574f0e6f580b452826761d784823d489db2b3510fceed14161b61ac07c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0b87cdcfc6d6804bc853cb595b09fc3f2ccf053ea3021248545c3424b1b3dd47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee4a495d0c4ffe0b94b6cd6170d222053c61d797f2748c4be8bff149a809b0d5"></a>

## default_pool.origin_servers.private_name.snat_pool.snat_pool — default_pool.origin_servers.private_name.snat_pool.snat_pool / 9a4564a6c4d8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-fd2bc05fe9b3ef436c32b78ff311202c4774b7636e1e98b9d26fd361a51e8e4e)
- [default_pool.origin_servers.private_name.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3670a6574f0e6f580b452826761d784823d489db2b3510fceed14161b61ac07c)
- default_pool.origin_servers.private_name.snat_pool.snat_pool

<a id="canonical-d36aa4a446fa66c37c10b570692d156405f86ce057ecfeff33a8fa13b25f273f"></a>

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

<a id="canonical-b5fb85b7d8ad18c286ff5da86e2fd76ffbec7ffc8b51ad733845e7f0623cfa47"></a>

## Direct properties — default_pool.origin_servers.private_name.snat_pool.snat_pool / 9a4564a6c4d8 / 3

<a id="canonical-46d60e7d7836f729b212a85008d64d3fc2170e96044c0eab3bda2efc1bd2b187"></a>

<a id="canonical-56240bba73e728126d4b8b0d9ec18846371a11644e9d607c00e9f8e1b0c81be7"></a>

## prefixes property — default_pool.origin_servers.private_name.snat_pool.snat_pool / 9a4564a6c4d8 / 4

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

<a id="canonical-8e3155a8bca742fdcfbbe11e474379a978133f4791391a2ef7468f45644a1019"></a>

## Next pages — default_pool.origin_servers.private_name.snat_pool.snat_pool / 9a4564a6c4d8 / 5

- [default_pool.origin_servers.private_name.snat_pool](resources--http_loadbalancer--reference--group-016.md#canonical-3670a6574f0e6f580b452826761d784823d489db2b3510fceed14161b61ac07c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5bee059e85c64ff82cff9be8d882b9bb291b25f95217cf99294d7b1505716cdd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38643a8ef4d6d3a5a63581d34459764622b6cbb46d2d5ce4a6dd9cb8f70b1ebb"></a>

## default_pool.origin_servers.public_ip — default_pool.origin_servers.public_ip / e7cefd4bf374 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- default_pool.origin_servers.public_ip

<a id="canonical-af2350d9d9ff925dc168379771ab0288585fd1b336eafbd8128541e788db032f"></a>

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

<a id="canonical-22b6c90a90fca5e855d76b5936123eada1a708d9ebcefdee2ed814482b1430aa"></a>

## Direct properties — default_pool.origin_servers.public_ip / e7cefd4bf374 / 3

<a id="canonical-401f0bd4066285e6847b6cf3ed35284e9d0430e54633872c45c1c9c960d88127"></a>

<a id="canonical-c5946b7cba57332bc1444573514703ce26ccde0be9b7ecbbfbb26340a8766e18"></a>

## ip property — default_pool.origin_servers.public_ip / e7cefd4bf374 / 4

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

<a id="canonical-138dd6c2152b698c3ec3412e61bf9206311c0ab6c81b8e6d5d00ec5310b68043"></a>

## Next pages — default_pool.origin_servers.public_ip / e7cefd4bf374 / 5

- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f2e32ad24b3ffb04aa89fcd991562a6c9733a3697e44402d62355c2c43dfe0b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f26c63c6ec5a03926aed3712424ac92d7109dc6a8f8f1eff53ec386f50ae2bf9"></a>

## default_pool.origin_servers.public_name — default_pool.origin_servers.public_name / 057be0a74b29 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- default_pool.origin_servers.public_name

<a id="canonical-4178514022ace0e0a1d2bea1b129c0b2b4c3212ad19d6383bb38d68d87e91bf1"></a>

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

<a id="canonical-78f433aa4c2bde9b47732e2bdc6aa9abee087203358bc3615486da22ce4cc976"></a>

## Direct properties — default_pool.origin_servers.public_name / 057be0a74b29 / 3

<a id="canonical-3fca00e5645a2c22285cc26a806d534dd4fea12cc0cb9dd2810e0fe4e2c9c096"></a>

<a id="canonical-d19741af2083853f2c5fefca6cb54bdfe7b96689f46d3267b14572bdbcc34db7"></a>

## dns_name property — default_pool.origin_servers.public_name / 057be0a74b29 / 4

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

<a id="canonical-0b6f3fb6da16a9c025b78a2cbdeaad416e41ba04eb630a50a90bf9d4266fdb71"></a>

<a id="canonical-83a0ce4c9573c4236aeaba69ff184f81756ebebcd2fa648475b00e5362c8ee6e"></a>

## refresh_interval property — default_pool.origin_servers.public_name / 057be0a74b29 / 5

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

<a id="canonical-3e9b0d7e8ff3db1da9a891a0112ce05596dfadd216135e733d6699cc659a9193"></a>

## Next pages — default_pool.origin_servers.public_name / 057be0a74b29 / 6

- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0ad0770c95ce89278886a3af584ddc06558cfec6495fcfefd9b9620c20c6389a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3b536266ed81de9abfde7dfb39df3aa30738e9929a41596a86e8c7ca9611f78"></a>

## default_pool.origin_servers.vn_private_ip — default_pool.origin_servers.vn_private_ip / 0e26bec57714 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- default_pool.origin_servers.vn_private_ip

<a id="canonical-a465bd9cd382ca090706784b3d07725d238b6832865b418a6ca4e7dedb833b9f"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
vn_private_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-03825e36783b544fbc5eb16b0f4d60b99759260f5878e1fc7e8edddb8de7f43e"></a>

## Direct properties — default_pool.origin_servers.vn_private_ip / 0e26bec57714 / 3

<a id="canonical-510c773dc1160de630ee69e7b5524e54d02f8c9933426a02d6f3a6db345f50ba"></a>

<a id="canonical-d7ec2e034d29147f65e49ad03dcb2eb16833f8df113fc1a62bffe4cab2bc4ec5"></a>

## ip property — default_pool.origin_servers.vn_private_ip / 0e26bec57714 / 4

Type: `"string"`. Optional.

IPv4. Exclusive with \[\] IPv4 address.

Upstream description:

Exclusive with \[\] IPv4 address.

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

- [virtual_network](resources--http_loadbalancer--reference--group-016.md#canonical-0a54e7ba3e0de8eb89af2468bd7338c9cb6f71dcfc4612956b22717a5f4fb4d9): complete subsection reference.

<a id="canonical-f20a5566d9cab6199e1135aa6765ce873c4343e285eb002bdfb839d5e780fe2e"></a>

## Next pages — default_pool.origin_servers.vn_private_ip / 0e26bec57714 / 5

- [default_pool.origin_servers.vn_private_ip.virtual_network](resources--http_loadbalancer--reference--group-016.md#canonical-0a54e7ba3e0de8eb89af2468bd7338c9cb6f71dcfc4612956b22717a5f4fb4d9)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0a54e7ba3e0de8eb89af2468bd7338c9cb6f71dcfc4612956b22717a5f4fb4d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd204435c9ff330425cb068808b90560ef562776120a40dedeb66c1cb79b0f7b"></a>

## default_pool.origin_servers.vn_private_ip.virtual_network — default_pool.origin_servers.vn_private_ip.virtual_network / ab8145c02386 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.vn_private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-0ad0770c95ce89278886a3af584ddc06558cfec6495fcfefd9b9620c20c6389a)
- default_pool.origin_servers.vn_private_ip.virtual_network

<a id="canonical-aa1636ef6dc1b7c472c24540930a3b9f8d6d60a0f95fbcd7391d8ff633f478d0"></a>

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

<a id="canonical-adf394fcea2793124dfede442778cde106e8403ae77db30990f7ca643f9a79a7"></a>

## Direct properties — default_pool.origin_servers.vn_private_ip.virtual_network / ab8145c02386 / 3

<a id="canonical-1579112b312ae6ba7556a39a4d2634f472bf31f8d06151727b3d34be0c76d87e"></a>

<a id="canonical-4a953e15519066048771fc94fb66d378ddc15a63a7a112c8ba2ea806f6d913ae"></a>

## name property — default_pool.origin_servers.vn_private_ip.virtual_network / ab8145c02386 / 4

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

<a id="canonical-f96da44ec55b6246d940aaf0c0a1718bab0c4bdf0695701491cde84d86e4a1c3"></a>

<a id="canonical-b86b80be0eb36960f029893f8d023fa0c915773d0772a774c48d82aca2d7dd05"></a>

## namespace property — default_pool.origin_servers.vn_private_ip.virtual_network / ab8145c02386 / 5

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

<a id="canonical-d7d074f26543d570876ff29122f80ca5ff583eb91d553d88c9164953dc0fb427"></a>

<a id="canonical-f5092ba7ca941341411e77a64e8a1c98bac24275bbb7b837a73b10aa9b38e7c9"></a>

## tenant property — default_pool.origin_servers.vn_private_ip.virtual_network / ab8145c02386 / 6

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

<a id="canonical-43523c6e041b1594fada0aa7b634c580ffe649324f0cd6102c44a31124604ca7"></a>

## Next pages — default_pool.origin_servers.vn_private_ip.virtual_network / ab8145c02386 / 7

- [default_pool.origin_servers.vn_private_ip](resources--http_loadbalancer--reference--group-016.md#canonical-0ad0770c95ce89278886a3af584ddc06558cfec6495fcfefd9b9620c20c6389a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-913f706cd0a71542850ef8c550736a5942b7a05f022e4ccff63a7758be36e6b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b1704059840fe88013d32922143de979730f5e0a627ac18be209b0b9d434dc6"></a>

## default_pool.origin_servers.vn_private_name — default_pool.origin_servers.vn_private_name / a6d24361bce2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- default_pool.origin_servers.vn_private_name

<a id="canonical-a87bd023d4f676676c7698e0e7e6d5796749a8c28806f51152220e4ca87fff97"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with DNS name on Virtual Network.

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
vn_private_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-db997a4eb1b383de86c626906d1c8f69d152fcba17b637cb502e8ee42b39018f"></a>

## Direct properties — default_pool.origin_servers.vn_private_name / a6d24361bce2 / 3

<a id="canonical-75cd169cf06f06a9f191aaa3467b2db196e15070ac9ce2df1ebe6a5b5d2ec9cc"></a>

<a id="canonical-a45de78a45d4c9835abe7410ecbd15edc5d1640ada636672359637e658e81059"></a>

## dns_name property — default_pool.origin_servers.vn_private_name / a6d24361bce2 / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

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

- [private_network](resources--http_loadbalancer--reference--group-016.md#canonical-2323801612b4fb846b4bb0129060948838e188a4878c969f85000c4c85c23c55): complete subsection reference.

<a id="canonical-ffd411deffcce929716f9c7989406cfb23c205da33b53b13b1aa29c75f1f3aab"></a>

## Next pages — default_pool.origin_servers.vn_private_name / a6d24361bce2 / 5

- [default_pool.origin_servers.vn_private_name.private_network](resources--http_loadbalancer--reference--group-016.md#canonical-2323801612b4fb846b4bb0129060948838e188a4878c969f85000c4c85c23c55)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2323801612b4fb846b4bb0129060948838e188a4878c969f85000c4c85c23c55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb831295232df0bf8527c44089748c5c563b6763c72da4c846e3a5ef6d69852f"></a>

## default_pool.origin_servers.vn_private_name.private_network — default_pool.origin_servers.vn_private_name.private_network / ead9376b1204 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-015.md#canonical-ce916683c1c2244d9b9fa4eb0bd7eae8fbfe4a91df6fa1b86427c35b76fe860d)
- [default_pool.origin_servers.vn_private_name](resources--http_loadbalancer--reference--group-016.md#canonical-913f706cd0a71542850ef8c550736a5942b7a05f022e4ccff63a7758be36e6b6)
- default_pool.origin_servers.vn_private_name.private_network

<a id="canonical-bd9f866368baafec7aa921e85613e21448dee50a1f0986161635b072973ca03a"></a>

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
private_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-5b3fc918511145664991c53c3291771c03d8da8a51270da0e9a999e09c0fc298"></a>

## Direct properties — default_pool.origin_servers.vn_private_name.private_network / ead9376b1204 / 3

<a id="canonical-052950a4503f70e71e63f4d8a3b6586def745c334a646757d1008ce7159cb8eb"></a>

<a id="canonical-f53fcf27107bec60ba6a97e1f82b147eca91f6af913efed0ca48c85bbeeb2012"></a>

## name property — default_pool.origin_servers.vn_private_name.private_network / ead9376b1204 / 4

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

<a id="canonical-41c16062897bf5d020120fc7d78e88eec7f07020a96964e7d0a0ff9ff69558e0"></a>

<a id="canonical-d07e7060ce3855f5e71bc61541ec69c94e4e8e85413973e46c5d23aae24bc808"></a>

## namespace property — default_pool.origin_servers.vn_private_name.private_network / ead9376b1204 / 5

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

<a id="canonical-ebf6c0c61cdfcb9c2cbf59a0ca910c3478ef5230183163138ce4bf0c4f6d4032"></a>

<a id="canonical-77da720b2f08384f38c83c5a9749b6d3d8f0bf104112ab9d77d33034ea8fd6c3"></a>

## tenant property — default_pool.origin_servers.vn_private_name.private_network / ead9376b1204 / 6

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

<a id="canonical-6b353da5b1d3cb3fc4cfcc21f9c67bcbf4b9b585fda8c1a07be14a0453a4c75f"></a>

## Next pages — default_pool.origin_servers.vn_private_name.private_network / ead9376b1204 / 7

- [default_pool.origin_servers.vn_private_name](resources--http_loadbalancer--reference--group-016.md#canonical-913f706cd0a71542850ef8c550736a5942b7a05f022e4ccff63a7758be36e6b6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6e43f3ff29119393843e3c446b9898f88148e1e43fc6af036d8075eb8495e529"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae67a588f6d19824c0b0aa097d993ca20b1ac522c188b4e338709d452199ceb7"></a>

## default_pool.same_as_endpoint_port — default_pool.same_as_endpoint_port / 476b4d342f46 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- default_pool.same_as_endpoint_port

<a id="canonical-8f9e9b019f2d1eb76c81842846eda220aa82d2ce468a2759d2f1938e4cb7ae17"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
same_as_endpoint_port = {}
```

<a id="canonical-e2f66708b51154554b09965f7f84771822da0786d142c3f4678d4e958b9b474e"></a>

## Direct properties — default_pool.same_as_endpoint_port / 476b4d342f46 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8710d030c41c9f6350616d46769ebcac2cd545fd3dcd10b43ef39fe1a781df6f"></a>

## Next pages — default_pool.same_as_endpoint_port / 476b4d342f46 / 4

- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-560c4c12019253450993494d981fb6f16358ab25b5d078bdbb365ea2790e52ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32d0fccefbf20175c063f803043f5e57698a12cf14b2dc0f196ef2e8da49f402"></a>

## default_pool.upstream_conn_pool_reuse_type — default_pool.upstream_conn_pool_reuse_type / f20d84a212ab / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- default_pool.upstream_conn_pool_reuse_type

<a id="canonical-1bea24407943cb81b0adf39a9003b696089caf78e79128c318632acca4cbe537"></a>

Type: `"object"`. single nested block, Optional.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_conn_pool_reuse",
    "enable_conn_pool_reuse")}
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
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

Terraform syntax:

```terraform
upstream_conn_pool_reuse_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-137c63473b7d403737ded04975e466e293ecbba6383d136c16acc223cdcacfb1"></a>

## Direct properties — default_pool.upstream_conn_pool_reuse_type / f20d84a212ab / 3

- [disable_conn_pool_reuse](resources--http_loadbalancer--reference--group-016.md#canonical-ca0fe5a3f764083a48191ee645139cf7289ed85fdb9db4c227c8141607226080): complete subsection reference.

- [enable_conn_pool_reuse](resources--http_loadbalancer--reference--group-016.md#canonical-1abf5c6c8f49c9a27304f5f6b373d5ff19cf846c6713af571b664034f4fb3a1f): complete subsection reference.

<a id="canonical-9bdca0859f5cd3505312a53fda3b7c5c9157e017b43f71ff13657cbde117756c"></a>

## Next pages — default_pool.upstream_conn_pool_reuse_type / f20d84a212ab / 4

- [default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse](resources--http_loadbalancer--reference--group-016.md#canonical-ca0fe5a3f764083a48191ee645139cf7289ed85fdb9db4c227c8141607226080)
- [default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse](resources--http_loadbalancer--reference--group-016.md#canonical-1abf5c6c8f49c9a27304f5f6b373d5ff19cf846c6713af571b664034f4fb3a1f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ca0fe5a3f764083a48191ee645139cf7289ed85fdb9db4c227c8141607226080"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9dd81803892f8a158e7a011d28941c3df68c7702959cc824bf2d908d7f20cb8a"></a>

## default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse — default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse / 4c60689a9e00 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-016.md#canonical-560c4c12019253450993494d981fb6f16358ab25b5d078bdbb365ea2790e52ec)
- default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-4f9263db3a7b5b5d8306140f6a371709c86297cf1f2698c1703e4078fa3bb949"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_conn_pool_reuse = {}
```

<a id="canonical-a8974f95675adf5501ba6f9a187b839f640bf50f801103f6bb276b546664ac4b"></a>

## Direct properties — default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse / 4c60689a9e00 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f91ccc23ceae16379dbf472543aecf7ca536818fe46cb9625fe4a27b1714ca32"></a>

## Next pages — default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse / 4c60689a9e00 / 4

- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-016.md#canonical-560c4c12019253450993494d981fb6f16358ab25b5d078bdbb365ea2790e52ec)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1abf5c6c8f49c9a27304f5f6b373d5ff19cf846c6713af571b664034f4fb3a1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd6d546d518fbb6fe937f4e31b40e9b77433c81388d61ceaa2297c644b3834a6"></a>

## default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse — default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse / 1a0133e1beee / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-016.md#canonical-560c4c12019253450993494d981fb6f16358ab25b5d078bdbb365ea2790e52ec)
- default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-d7dc829f3cc2df2dce99485b2d47e2ded9d8d5c3c8b13ee6f36a4382e7f307b8"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable_conn_pool_reuse = {}
```

<a id="canonical-704f6d283bdc68aa5bea63f228e5090980da44f61b5c6183728bc6066f4efb91"></a>

## Direct properties — default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse / 1a0133e1beee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-315eaf172aed895fce9099976c01da0dfac30721dd87ac1a609b898d30145c44"></a>

## Next pages — default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse / 1a0133e1beee / 4

- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-016.md#canonical-560c4c12019253450993494d981fb6f16358ab25b5d078bdbb365ea2790e52ec)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-68492bfc5ceddba344a680db2d846470aadb36d450381b9ceca31b745a6e954a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-936b63df2e21aa60faeaedfac66fcf8e0065835549fc2bda359606fc63354dbc"></a>

## default_pool.use_tls — default_pool.use_tls / 9968264bf97b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-c5364c4bd5cda7551a10e887854412a029f9be946b9e974764fe3c1ac76ed345)
- default_pool.use_tls

<a id="canonical-be9cb37e47efb364106d6c85bc1ec4c4127b5f38840f20cfb7633f3d3649a6e9"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Upstream description:

Upstream TLS Parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_session_key_caching",
    "disable_session_key_caching"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("disable_sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "use_server_verification"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "volterra_trusted_ca"),
  validators.ConflictingObjectAttributes("sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("use_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("use_server_verification",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-bf3a96c576e1238c2a70659030dd09d6f4c0fd16f0483ab821623ad7bb058dec"></a>

## Direct properties — default_pool.use_tls / 9968264bf97b / 3

- [default_session_key_caching](resources--http_loadbalancer--reference--group-017.md#canonical-a2acfe534bdcaa0e5053d2f480bea97b11a74d89d2f878881191a1949391d9f1): complete subsection reference.

- [disable_session_key_caching](resources--http_loadbalancer--reference--group-017.md#canonical-47f48b7f228f8c096ef34438b03482235bf3677bc15e0db31a30286aef3ed799): complete subsection reference.

- [disable_sni](resources--http_loadbalancer--reference--group-017.md#canonical-b09ff81e1b6e78eb20e610fe5f76673a18292c7aded52fc2609604bbeff8d198): complete subsection reference.

<a id="canonical-67f0724b27ec22eab45c366d04f2ec7e65668d603d2be0cb3ca62ca1721d4f75"></a>

<a id="canonical-01252ad242ce7e58b9788257c1686c61f6395ba6bb1a9ddaaf36cae96758a5c1"></a>

## max_session_keys property — default_pool.use_tls / 9968264bf97b / 4

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 64),
}
```

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

- [no_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-33f4e0a4361011987c5811f404e04aba9e70703e6d77518bc949bafdbcaac10e): complete subsection reference.

- [skip_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-33f9e5356bcfd9228c0a7d2d9124d4c2ea9ac37da7b1cf02f374d37484d15c4c): complete subsection reference.

<a id="canonical-e34ea02f813ba725aea0e8bffec4f33b875fd3d2cf1cd5d9158d3f785a2d4b5b"></a>

<a id="canonical-8f9b830a10a5ce472001e73465322e12b8e8553d7b119a44ab526be6efd69228"></a>

## sni property — default_pool.use_tls / 9968264bf97b / 5

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

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

- [tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-ca0761358fe533402c9626e5cea7f3ff8bf7312ad3597f4ef3c885cc176b2a16): complete subsection reference.

- [use_host_header_as_sni](resources--http_loadbalancer--reference--group-017.md#canonical-7a219fecc30a63d0bec0125952244ad9dd6f6eb848785eec222ba376f7eb9bc7): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-f49fe1b0b9d777d1ef353f580d43bcbd164194292c3ba075662faa9cc7f5f14d): complete subsection reference.

- [use_mtls_obj](resources--http_loadbalancer--reference--group-017.md#canonical-85d339096a8a94d45c3ee4ac5d310e910ea0d15e221b1ad1c7454382122a6ec6): complete subsection reference.

- [use_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-f26b369f3bbe3d3bc658c0b08395510e8df06b4445cd510817ab56c342ad0a0c): complete subsection reference.

- [volterra_trusted_ca](resources--http_loadbalancer--reference--group-017.md#canonical-f913ba70e9bad1bf406bf57ac712f432b527b62659c57f4aea342e0ad11610cb): complete subsection reference.
