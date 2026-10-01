---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-2d084f49c4a6a292a710676612534aa7926ed42e126acf0af1e4a0e2e9087b46"></a>

## advertise_custom.advertise_where.vk8s_service.virtual_site — advertise_custom.advertise_where.vk8s_service.virtual_site / a7555620a25a / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [advertise_custom.advertise_where.vk8s_service](resources--tcp_loadbalancer--reference--group-001.md#canonical-b4392d7e78138a30c0757783cdc3c073c5da6748af1bf6264da06f58d293c38a)
- advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-966d0364d83a9ed98fcbfe5cfb83e0eb0f634809b46f91110b61efa4768e9d46"></a>

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

<a id="canonical-f67685b39b477647d293adf71f438c23cfa587aec4635445fc381cca17eaa4aa"></a>

## Direct properties — advertise_custom.advertise_where.vk8s_service.virtual_site / a7555620a25a / 3

<a id="canonical-15c3669a45b81e82bf8d58c62a26d9cd769d96656d19273e08cc933da1220350"></a>

<a id="canonical-c09ea1f68e4fcfc4856be3453f1ae5ade4a2db73ca9d1914c54e98aef9261684"></a>

## name property — advertise_custom.advertise_where.vk8s_service.virtual_site / a7555620a25a / 4

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

<a id="canonical-56cc8c0463b677f58f80b081e530848dbfe3e338fc8477c70368b179ed87da2a"></a>

<a id="canonical-317ca6b8bf4ece655ab3dab07e352b3580b5b2ea5de97ae2934ceab61050206e"></a>

## namespace property — advertise_custom.advertise_where.vk8s_service.virtual_site / a7555620a25a / 5

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

<a id="canonical-a06692555aa0a9222ae73ee852b7d52571008ef5eec10162f1bee11889732db2"></a>

<a id="canonical-c79feee2c62edb6c10a27b13c117c174547bb01a6ca092dff35a02a0c5568624"></a>

## tenant property — advertise_custom.advertise_where.vk8s_service.virtual_site / a7555620a25a / 6

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

<a id="canonical-2fbaedc39ac2396dedae88dbf114a97ca6a91e6e6b652c7eb9209fb3813e14c7"></a>

## Next pages — advertise_custom.advertise_where.vk8s_service.virtual_site / a7555620a25a / 7

- [advertise_custom.advertise_where.vk8s_service](resources--tcp_loadbalancer--reference--group-001.md#canonical-b4392d7e78138a30c0757783cdc3c073c5da6748af1bf6264da06f58d293c38a)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-5145a903bc57dcc49c9709e298e99c068fe602cad3de3f2ab20093ee69049729"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4810ff0f42bb8d4a510e693ccdb25b39833daee117495bd7a505dc20102fd4d"></a>

## advertise_on_public — advertise_on_public / a3741ebaf7b0 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- advertise_on_public

<a id="canonical-32ddddd3b88cd727dfc738e4a881d56c2f176fb82c812cd6b5af1a6bceba73a5"></a>

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

<a id="canonical-038b8fe193a3a376b47dd82a0567742e40ae2fff97cbcb38992361f23734311a"></a>

## Direct properties — advertise_on_public / a3741ebaf7b0 / 3

- [public_ip](resources--tcp_loadbalancer--reference--group-002.md#canonical-62695909c68df812937cd31e41e40100395ad2eee66fe6ae2c8c89f5b6631e12): complete subsection reference.

<a id="canonical-35f9993b34d88902eee136b4ecbec6fa8428a93d119d9dd65085c730ad8e478a"></a>

## Next pages — advertise_on_public / a3741ebaf7b0 / 4

- [advertise_on_public.public_ip](resources--tcp_loadbalancer--reference--group-002.md#canonical-62695909c68df812937cd31e41e40100395ad2eee66fe6ae2c8c89f5b6631e12)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-62695909c68df812937cd31e41e40100395ad2eee66fe6ae2c8c89f5b6631e12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e35377248c58f76201b2f3c3a07db10ed8d93950f164c6c75081b5460d98e592"></a>

## advertise_on_public.public_ip — advertise_on_public.public_ip / f9fdf8a14950 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_on_public](resources--tcp_loadbalancer--reference--group-002.md#canonical-5145a903bc57dcc49c9709e298e99c068fe602cad3de3f2ab20093ee69049729)
- advertise_on_public.public_ip

<a id="canonical-baa551955ff9311107bba0f7822fdcdde5a5a0f1f1992bbee147f7b3a5931e6f"></a>

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

<a id="canonical-c9b97d4b347318b76ef03332e90766a9f6c491372266002d5c0985f337440a91"></a>

## Direct properties — advertise_on_public.public_ip / f9fdf8a14950 / 3

<a id="canonical-0283f86159db777fea17a1ab85b4fc403294497eaea5705e4b96b62c30830f9d"></a>

<a id="canonical-b7505786baf467d62e22e2dbd2755d7ee9d15f8ce55ea0f4bc88866d46f7dfbb"></a>

## name property — advertise_on_public.public_ip / f9fdf8a14950 / 4

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

<a id="canonical-7687b3dadb92306fb8af8a2faa861c662cc63fcef498a00fc08d9e9bbffe7b40"></a>

<a id="canonical-20382d00d208cee17c4924192e5f63588ebb34eed83b751560938351763c776c"></a>

## namespace property — advertise_on_public.public_ip / f9fdf8a14950 / 5

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

<a id="canonical-7356de531c6b5f5d247a1d1a340e199641612607aa53c74080d1f87c1d29117f"></a>

<a id="canonical-7995b2c2658889fbaa6803b66640dfa5073a62b6716836e3fda02cca3a2a63d2"></a>

## tenant property — advertise_on_public.public_ip / f9fdf8a14950 / 6

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

<a id="canonical-7c2426603b77ab15b8d9dda7a037d901a54e141d647f10a40b10cf98e0b66cca"></a>

## Next pages — advertise_on_public.public_ip / f9fdf8a14950 / 7

- [advertise_on_public](resources--tcp_loadbalancer--reference--group-002.md#canonical-5145a903bc57dcc49c9709e298e99c068fe602cad3de3f2ab20093ee69049729)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-ff42e920406b0375c6cbaac8ce29eedff4194fc6bf1fc0585caa1787584c2174"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cfafd40fe093082b71e86c0b639a28436b5ce33635b69c6de547058c9785a436"></a>

## advertise_on_public_default_vip — advertise_on_public_default_vip / 8e0cdb53033d / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- advertise_on_public_default_vip

<a id="canonical-398ab0b68068225a40ff55c122883245ccb581e7f66f1534fbb67519359d5b79"></a>

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

<a id="canonical-3190f70cc35947520e75f18119cd94d89ddef6c4c1da3db5f702919bc59b507d"></a>

## Direct properties — advertise_on_public_default_vip / 8e0cdb53033d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-54ed0ec8aee01b4a327ab5fbed19449cf57b32a3d1cb57471ace71c6aa93795b"></a>

## Next pages — advertise_on_public_default_vip / 8e0cdb53033d / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-6fc17a77d48f9cf85e426e65ef541a24a263cb7378eb08dbdebeaddb4c6c5404"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ac99d83016b809e62580865cd9e0ef6521205f6e2f7f40022048899a933af6c"></a>

## default_lb_with_sni — default_lb_with_sni / eaf6127e29dc / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- default_lb_with_sni

<a id="canonical-3788fb217091bb364117acb918953f23d8b90c1c58ced8a152d22577b6555028"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_lb\_with\_sni, no\_sni, sni; Default: default\_lb\_with\_sni\] Configuration
parameter for default lb with sni.

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

- [default_lb_with_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-3788fb217091bb364117acb918953f23d8b90c1c58ced8a152d22577b6555028)
- [no_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-c7ca021715eb3eac0052ca372e3b473fbb49cd1d446d5f30a8b5d6c0a5e2d789)
- [sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-99f5e75dcf3506323c7b2293c0a12580c5d75d61a0d59b6bc5a40868e1d1663a)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_lb_with_sni = {}
```

<a id="canonical-57331a1a6256d86a762f66b638eba8850d4fb7e1aed2e72e5ba9feb384551541"></a>

## Direct properties — default_lb_with_sni / eaf6127e29dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-daa7e4357c6716aed8421829d7bd674cb933f8626b1624276e6507de701106cb"></a>

## Next pages — default_lb_with_sni / eaf6127e29dc / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-ba9e95942d760250bb35617997afdb3d548b484a3b5ac545e364fd505c61a698"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c36011f0e25cc0f9ebb7f92c682fa937f361109e49016a4938784bdbde95e4fc"></a>

## do_not_advertise — do_not_advertise / 751b32a1c871 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- do_not_advertise

<a id="canonical-de7e64a189c21e9f3f46ce0121947d119ee5bee16f16d6a10e89d0504f2a673d"></a>

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

<a id="canonical-e0ed8d38bdbafabdaf4c15782639d28b3fc81b7b7344352554ed726f01dbd395"></a>

## Direct properties — do_not_advertise / 751b32a1c871 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-458f506f35e4adefe6533585d1a226085ef417c16084fe401256f3167f48ea54"></a>

## Next pages — do_not_advertise / 751b32a1c871 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-ba5e7e3c123f000e5c21dcf85c3f1ab41075becc703e6fcc8f87971d9c0fce07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-633db3d12291fee46a99550136e6762501a6ad342366a03b79ed9ff97f60224d"></a>

## do_not_retract_cluster — do_not_retract_cluster / f6ffd36612df / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- do_not_retract_cluster

<a id="canonical-9e6628b6f34d48dd1b66dc924ac1d5ec91d3c3774d7fbcafb402b3f8dc162bef"></a>

Type: `["object", {}]`. Optional.

\[OneOf: do\_not\_retract\_cluster, retract\_cluster\] Enable this option

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

- [do_not_retract_cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-9e6628b6f34d48dd1b66dc924ac1d5ec91d3c3774d7fbcafb402b3f8dc162bef)
- [retract_cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-2b9a9a7b43fc5b3bb54f0086ca58dfff2563b0e53b5fac4e073e87d03972460f)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
do_not_retract_cluster = {}
```

<a id="canonical-00db4e89bca6f7195b15e30ea7e10f5ac6ef8fb124090104ad080ae572ca5293"></a>

## Direct properties — do_not_retract_cluster / f6ffd36612df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7659335588994a73d62ce1a7505e0f99643b5d47b9b0dd8ba8777215a99520b1"></a>

## Next pages — do_not_retract_cluster / f6ffd36612df / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-edd915446ceae49d7455932f6f248f30860d874aec8e428a7dd297607454e0bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bebd63f9d243896bba8c591ddca9952647612734b7f84edb73d9571868283374"></a>

## hash_policy_choice_least_active — hash_policy_choice_least_active / c5d62216c53f / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- hash_policy_choice_least_active

<a id="canonical-982847a6e067eab3f677405e68402501c7e29da80d0971dc6b854489d213b593"></a>

Type: `["object", {}]`. Optional.

\[OneOf: hash\_policy\_choice\_least\_active, hash\_policy\_choice\_random,
hash\_policy\_choice\_round\_robin, hash\_policy\_choice\_source\_ip\_stickiness\] Enable this
option

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

- [hash_policy_choice_least_active](resources--tcp_loadbalancer--reference--group-002.md#canonical-982847a6e067eab3f677405e68402501c7e29da80d0971dc6b854489d213b593)
- [hash_policy_choice_random](resources--tcp_loadbalancer--reference--group-002.md#canonical-38521a8224ad394cfe8cd652d3035a77824b9a41bc3a4efbf14ab6614d1dac6d)
- [hash_policy_choice_round_robin](resources--tcp_loadbalancer--reference--group-002.md#canonical-fae47bac70c13a3fa8e46c01d2322e558fc2fde428219f768ed6920169f060a6)
- [hash_policy_choice_source_ip_stickiness](resources--tcp_loadbalancer--reference--group-002.md#canonical-d12c6f4812985b415deec0717b71a02ce26d8da7e057db994ed37e035a4594ea)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
hash_policy_choice_least_active = {}
```

<a id="canonical-c139bc4a7cdc8799ad63e7566bef5cc70e5825c93bd1a14c3ac38452981243cb"></a>

## Direct properties — hash_policy_choice_least_active / c5d62216c53f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e3a33d1ffcfdd41ff75dbebf535019a4578a65158211601c46bcfbef14baa457"></a>

## Next pages — hash_policy_choice_least_active / c5d62216c53f / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-9f8669f3e8de1af30d714b546776d46ef9bc8683f210f0b322894e386a825f0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6ef382aa7ae4d5eb4ea662536b5ced32159a82cada331909146efdc1b478174"></a>

## hash_policy_choice_random — hash_policy_choice_random / 0bbb31985733 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- hash_policy_choice_random

<a id="canonical-38521a8224ad394cfe8cd652d3035a77824b9a41bc3a4efbf14ab6614d1dac6d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for hash policy choice random.

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
hash_policy_choice_random = {}
```

<a id="canonical-fa42cd4cef58f0ad23a3e5893fb876255e9914be64145fd427b1c530785b39db"></a>

## Direct properties — hash_policy_choice_random / 0bbb31985733 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-768cd5b57c2b54ab7ae41998c473cc739686b22efc7683608b47ad7f529ffe22"></a>

## Next pages — hash_policy_choice_random / 0bbb31985733 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-85cdf1528965a6285ff2c75a0267cb449348c1ca88921d060042419de75f1ace"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd5f5e9ff0831c28affec2d5e4a535aa196990d7cddc643f0cdf17dde4bd7470"></a>

## hash_policy_choice_round_robin — hash_policy_choice_round_robin / a6ce632bd4ed / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- hash_policy_choice_round_robin

<a id="canonical-fae47bac70c13a3fa8e46c01d2322e558fc2fde428219f768ed6920169f060a6"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for hash policy choice round robin. Defaults to \`map\[\]\`. Server applies
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

Terraform syntax:

```terraform
hash_policy_choice_round_robin = {}
```

<a id="canonical-43d57f4607890e097a4a8d89477b615652d95bf33ed768221cc7ef443a8cf9bf"></a>

## Direct properties — hash_policy_choice_round_robin / a6ce632bd4ed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a22ea0fa8189ebff6bbe4fda3a5260938767c172e615e675927d169bb9963ca"></a>

## Next pages — hash_policy_choice_round_robin / a6ce632bd4ed / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-dbedf978ab25250f8b6bdd8d76481f4d8b6a7e6f331b982a6fc699ddc46b8ac6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ceefddb219e2e89e3c68b43e76c963742d55a6b265e596e9767cd8c6245da01"></a>

## hash_policy_choice_source_ip_stickiness — hash_policy_choice_source_ip_stickiness / 2f14a6f41b51 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- hash_policy_choice_source_ip_stickiness

<a id="canonical-d12c6f4812985b415deec0717b71a02ce26d8da7e057db994ed37e035a4594ea"></a>

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
hash_policy_choice_source_ip_stickiness = {}
```

<a id="canonical-863cdfbccd9563a8f5353a807bb9404fcad1266fe798c5214e46204b6aa43069"></a>

## Direct properties — hash_policy_choice_source_ip_stickiness / 2f14a6f41b51 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-761b53b985400d06bcba9a2245acae28d4a47a50f724f2c896a92e14e3a62602"></a>

## Next pages — hash_policy_choice_source_ip_stickiness / 2f14a6f41b51 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-25ccd69fc2452f16f34ad58d3b9af414cbf6a2ca1af6864921afc4c6286aca01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bebccba383e32cb4638a46fd35bbe12b4c28e29f4f728fe36882d52ff2a40a27"></a>

## no_service_policies — no_service_policies / e1fddfc835ab / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- no_service_policies

<a id="canonical-12c893c9b154a54bbd2643e61d7c9e8fc00d6788f4f4fe74c14f0612bacefd95"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no service policies.

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
no_service_policies = {}
```

<a id="canonical-3429a7bc7bf063aaa9c2c7891ce3b1666d473b05c0b6bce3af9ae26a6dacc992"></a>

## Direct properties — no_service_policies / e1fddfc835ab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-104c3e010a70ff321eca0e5e26920d8ce38fa984ac611de3eb472c78ef452c70"></a>

## Next pages — no_service_policies / e1fddfc835ab / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-3e6782861d51d0a9e5a0f8e8ffa45db8737335ce850aafcfc19a13424444586b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3862b5bc0e155837080d08ad169272448fbf1ead7c2a8746d300998d72b7bf6c"></a>

## no_sni — no_sni / ffbada7e738c / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- no_sni

<a id="canonical-c7ca021715eb3eac0052ca372e3b473fbb49cd1d446d5f30a8b5d6c0a5e2d789"></a>

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
no_sni = {}
```

<a id="canonical-5ddd48d4a133cfe97c48ec6a9643855a66c3bb493d2e6ea0f56300409f7e547d"></a>

## Direct properties — no_sni / ffbada7e738c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-608b272c71da7f585624e64b151d80089bfb57f6b26887afea38669b02c62118"></a>

## Next pages — no_sni / ffbada7e738c / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-9010378f7d918df4260eb3f27869e01272844da0ffa7dc25fe2fa5c5121b4bda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4201503cccad78d29ffc0e44b592bf9db6904f9b2ffcbda3bf8dea32efbfc840"></a>

## origin_pools_weights — origin_pools_weights / 83a8a0b5a84d / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- origin_pools_weights

<a id="canonical-f04ad50b5ad516b9df6f8c3cd2a82dfb4162bccfeb44f6910e644751ef3e536e"></a>

Type: `"object"`. list nested block, Optional.

Origin pools and weights used for this load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cluster",
    "pool")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
origin_pools_weights {
  # Configure direct properties listed below.
}
```

<a id="canonical-ab6ca160230e53e18be75cfe97f3694c85fd58cf121998a34a6d0558b62cb38a"></a>

## Direct properties — origin_pools_weights / 83a8a0b5a84d / 3

- [cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-4563c7d7e9d12526c65254ce66dfc2beb3f2e111e65fcc21bd414081fc5571ff): complete subsection reference.

- [endpoint_subsets](resources--tcp_loadbalancer--reference--group-002.md#canonical-2f7b20574768b92a6a2b3ee0d844f3592e88e8950d4327d973bfcb21716f43db): complete subsection reference.

- [pool](resources--tcp_loadbalancer--reference--group-002.md#canonical-1118e536064c7a49806d688fa5e16d21275004c4792ea5e5cdf66ed9b81a8a4d): complete subsection reference.

<a id="canonical-8b2f0ac94c1faca03b213d64d99636f61e6b4c55b4fc302cb8d0e77aed694cb0"></a>

<a id="canonical-c17db9acfba85e71416225f3c794186108a639ee7ff64ef3712c68c53551bc9c"></a>

## priority property — origin_pools_weights / 83a8a0b5a84d / 4

Type: `"number"`. Optional.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
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

<a id="canonical-702ed77400d2d458554f69b652a904a7576f5cc7599fa01cbc01ed2e50e70cdc"></a>

<a id="canonical-b34cff32c37fb1cee302006ec5fdd3207c6a1ab275a1f218edd72b1e54ca929e"></a>

## weight property — origin_pools_weights / 83a8a0b5a84d / 5

Type: `"number"`. Optional.

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

<a id="canonical-9b37d6dd87d7321d0fa13476e3a9a06a46ba9308c25812a861c4c432bf5b911a"></a>

## Next pages — origin_pools_weights / 83a8a0b5a84d / 6

- [origin_pools_weights.cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-4563c7d7e9d12526c65254ce66dfc2beb3f2e111e65fcc21bd414081fc5571ff)
- [origin_pools_weights.endpoint_subsets](resources--tcp_loadbalancer--reference--group-002.md#canonical-2f7b20574768b92a6a2b3ee0d844f3592e88e8950d4327d973bfcb21716f43db)
- [origin_pools_weights.pool](resources--tcp_loadbalancer--reference--group-002.md#canonical-1118e536064c7a49806d688fa5e16d21275004c4792ea5e5cdf66ed9b81a8a4d)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-4563c7d7e9d12526c65254ce66dfc2beb3f2e111e65fcc21bd414081fc5571ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d603785cea78536f29d88a4c5ce5a3ac5a101d7da3a6ada51b4c3622b851855"></a>

## origin_pools_weights.cluster — origin_pools_weights.cluster / ea1b455d9894 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-9010378f7d918df4260eb3f27869e01272844da0ffa7dc25fe2fa5c5121b4bda)
- origin_pools_weights.cluster

<a id="canonical-942bd378bb92030aa1cd9915b20be97862c55f3bdad9b1ddea6c1e596280437c"></a>

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
cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-fe1cd68dceb6b939c00c8703b369fd2d49aa6365bc0793c0bd05c520a17ac321"></a>

## Direct properties — origin_pools_weights.cluster / ea1b455d9894 / 3

<a id="canonical-a984aaa851ba57bd65eb5d462deac346d8bbdc71a1690bf84a086c466256a46f"></a>

<a id="canonical-2d89063b6ecd4fca29080a19a1b6fff99e1ffaeeb7fc2eedc04e198658438d9a"></a>

## name property — origin_pools_weights.cluster / ea1b455d9894 / 4

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

<a id="canonical-76e5f22b8b233e3257e5fa65f43eaed9a817d9486b26f64a8273e8b351506c82"></a>

<a id="canonical-70dceaf01bf481c0185371101e33daee83410ac0e433b6f4359b3d463c687171"></a>

## namespace property — origin_pools_weights.cluster / ea1b455d9894 / 5

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

<a id="canonical-3370414cdcc3b8c966737e9eaec3a22cfcd4707dd555b616ff248e2c3586b869"></a>

<a id="canonical-451ee1149946c78368be9a0d1220f2cd26a7b60e487656606959b706665d9bf2"></a>

## tenant property — origin_pools_weights.cluster / ea1b455d9894 / 6

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

<a id="canonical-ac36d248ceb486db3a1240f6dac2f12867a438523cbd07a74700be0013986edb"></a>

## Next pages — origin_pools_weights.cluster / ea1b455d9894 / 7

- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-9010378f7d918df4260eb3f27869e01272844da0ffa7dc25fe2fa5c5121b4bda)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-2f7b20574768b92a6a2b3ee0d844f3592e88e8950d4327d973bfcb21716f43db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0ee37f008f03b9520c1861cef2cb5e1bb8c9330b0256246797a3957497413fe"></a>

## origin_pools_weights.endpoint_subsets — origin_pools_weights.endpoint_subsets / 4d1a7f1f8cbc / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-9010378f7d918df4260eb3f27869e01272844da0ffa7dc25fe2fa5c5121b4bda)
- origin_pools_weights.endpoint_subsets

<a id="canonical-d8bd47f63991028ac0e577f4dc4512696bc7ef1f53ad4b2cac2031492f71f2ac"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
endpoint_subsets {}
```

<a id="canonical-10750ecf49c820d1c7aad123e7179b2078a296d407b342ac435d3aad9b04cd66"></a>

## Direct properties — origin_pools_weights.endpoint_subsets / 4d1a7f1f8cbc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c3c54c219df14d788b32debac285901c098fdbeff563862924dbfe25180a2061"></a>

## Next pages — origin_pools_weights.endpoint_subsets / 4d1a7f1f8cbc / 4

- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-9010378f7d918df4260eb3f27869e01272844da0ffa7dc25fe2fa5c5121b4bda)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-1118e536064c7a49806d688fa5e16d21275004c4792ea5e5cdf66ed9b81a8a4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4977ee9da9ab104bc0c7abeb9f7874e52eb81bbb7c257db80adcfb3950a72222"></a>

## origin_pools_weights.pool — origin_pools_weights.pool / 300aabaa76ef / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-9010378f7d918df4260eb3f27869e01272844da0ffa7dc25fe2fa5c5121b4bda)
- origin_pools_weights.pool

<a id="canonical-a5504463dea2b40665148d07dfbf6b838cc82d0a218f03b6a2e90550953f2900"></a>

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
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-942788662608516caeeb73bd4473c14ddaa888aa71941090a43ed04a3cadee4d"></a>

## Direct properties — origin_pools_weights.pool / 300aabaa76ef / 3

<a id="canonical-ed7e56172412822cce07fc1efca6093483967a488ae7c8c1e1e83ffdac939ff1"></a>

<a id="canonical-45fea1fe6954e4db361ad6554771786b785bf01a4c0c1e7e5182cc1cc5094803"></a>

## name property — origin_pools_weights.pool / 300aabaa76ef / 4

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

<a id="canonical-75c72e7953856558b310d1dfcddf8e71a7a170d1b867517ca4f3ee13fbe3b83b"></a>

<a id="canonical-77e7a561324c0b5c4fd85a081b80f6715cce97a2fccf847a3c66b1d513862a9b"></a>

## namespace property — origin_pools_weights.pool / 300aabaa76ef / 5

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

<a id="canonical-2437b6cacfa240a19aeb58c09d060fefc8e7f0e29014b1ff769019879e4f5422"></a>

<a id="canonical-e4309e94e6ad8f22d5855002db5c2c501f3ef5024b726c47df4e2df0e1afa32d"></a>

## tenant property — origin_pools_weights.pool / 300aabaa76ef / 6

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

<a id="canonical-ff8754c212ec720faab942c7a6cd68a9510d5d3edee83db775013bbd5c884b17"></a>

## Next pages — origin_pools_weights.pool / 300aabaa76ef / 7

- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-9010378f7d918df4260eb3f27869e01272844da0ffa7dc25fe2fa5c5121b4bda)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-a3f2eb692d42a4518e81b4d2ca346c903a6e9aedbc28daf1e229ac85d2217fbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5209eaf5b7e2b05c94ca5c2735a92dabd9397ecf9d949d2437088bbbbf984b08"></a>

## retract_cluster — retract_cluster / b9d645a1fdb0 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- retract_cluster

<a id="canonical-2b9a9a7b43fc5b3bb54f0086ca58dfff2563b0e53b5fac4e073e87d03972460f"></a>

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
retract_cluster = {}
```

<a id="canonical-c9bf9afd0b15021f59a07c440ce4327d476f9157e96386c63b4aa36edf5781fd"></a>

## Direct properties — retract_cluster / b9d645a1fdb0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4067ed7753925feb483942c04074147a5f49be840c9d2a7da427cae5d21e05f2"></a>

## Next pages — retract_cluster / b9d645a1fdb0 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-7202f6813ce28cd3211969f35f4a551cbf7c6921c78f64f4a67273125a5d0ceb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f9800ea552c7ae7270998700b65bacd569433559961567202ffaceaf6715fb8"></a>

## service_policies_from_namespace — service_policies_from_namespace / 56e36f31dca0 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- service_policies_from_namespace

<a id="canonical-c1ad3fff48340b5a358cdb53dfdbe34e72962062e51d663e4320d1daef574a2c"></a>

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
service_policies_from_namespace = {}
```

<a id="canonical-e77335961853bd29d0c5e289cca9f07253aab9a801889387d2a170cf19d14269"></a>

## Direct properties — service_policies_from_namespace / 56e36f31dca0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-489dea6f0a2590c4664ac3099968d14eaa58e72b01ef47b3361d04dc12a0059e"></a>

## Next pages — service_policies_from_namespace / 56e36f31dca0 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-65d837fc6b970ea897c23eaec474f1cabe159cddf3ae011ef00b692480d12e3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6488fa3cc4b8016979c7dee85cad7b00b3387713ac859b4f6ca48dbe9c7aa29"></a>

## sni — sni / fa90350f7a69 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- sni

<a id="canonical-99f5e75dcf3506323c7b2293c0a12580c5d75d61a0d59b6bc5a40868e1d1663a"></a>

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
sni = {}
```

<a id="canonical-d0199d0ea46695b8c4f0b01890851417daf850714fb1dd108f9e59199dc01aa2"></a>

## Direct properties — sni / fa90350f7a69 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-56450471cfddbcb85151fa19f2d932cb23678932bf01e4809dc4529c0fe34b13"></a>

## Next pages — sni / fa90350f7a69 / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-a8b8d21b3fa7b0073cd75420b779305e70da290853df37c205c279e25e1912e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc1b72a8065285d9e9da7bcfdfa594a9a5422dd79bc1cb05bc80df77db0b6bb0"></a>

## tcp — tcp / 3234ddcfccfa / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- tcp

<a id="canonical-f8a3300666d29fc38bb80075835da84e389d0f20e01b4bfd63d8b3b38f6c0fa1"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: tcp, tls\_tcp, tls\_tcp\_auto\_cert\] Enable this option. Defaults to \`map\[\]\`. Server
applies default when omitted.

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

- [tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-f8a3300666d29fc38bb80075835da84e389d0f20e01b4bfd63d8b3b38f6c0fa1)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-f5d11b5136db56e396813bcb2d1d10476e78227b0523aa6575c4c8c4a64cc8fb)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-490a6432751fa2d70f483cf952ae0e4fd5499dee0876f285d95600aa388da7e3)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
tcp = {}
```

<a id="canonical-0bb97cdd27eb3984481a67e6575f59721a9d692bcfbc73dcbf21c63e0f9e62e9"></a>

## Direct properties — tcp / 3234ddcfccfa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-969173216c54e46bfc54a6424a67a60bf79ae8197f44e1e193629c2acb62a600"></a>

## Next pages — tcp / 3234ddcfccfa / 4

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-6b17fb1ab324a8664c2799ba750302145ca7f5afebaf304cad2896a517809a45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62f72ebead8d9ae89cc52c6cc649674e048486b69fa53f3ab2dd04912e7a5b46"></a>

## timeouts — timeouts / b07600b5c2b6 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- timeouts

<a id="canonical-2ddf1f8d291bf93ba2a0d30af0129e2812f85e6039aafd9c5fb7bd067fff75aa"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-7a445dfc3a4fef51dfce33b5f7a4de52f7cc315051a037375fa937ed9f758eb6"></a>

## Direct properties — timeouts / b07600b5c2b6 / 3

<a id="canonical-6bbc654538118f3b9442b4420925d63ae60e555308b8d1e8f2999236feedd49e"></a>

<a id="canonical-c0a49a27c2b50e77fef4134644e3a59f7f173e7df2c5a141b21abf023fd33e19"></a>

## create property — timeouts / b07600b5c2b6 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-b1e05b0dd7f397839eda1921cd879fcc361bcba32112bcacaff2ad9b9e5ac993"></a>

<a id="canonical-1610723c8b56027f44b42be174c6ef1116e15856e558a842f10756748cc98abf"></a>

## delete property — timeouts / b07600b5c2b6 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-b63277a70989155692a5117d3eccf29a336d8a8141a54381253d73b84469202d"></a>

<a id="canonical-f607037a9dbe86887c14989edf8f8157ac9eda23c1f9e6af30d506a1ba60a18f"></a>

## read property — timeouts / b07600b5c2b6 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-fbf72386fc53c991b3b9fe9c2b0fa1deb2422471a3db9f77855ac737743bdcc4"></a>

<a id="canonical-a3d8edff6a985596ae5f45a1250dc4b3172405a64f0a4717fe05ad9fd59e80b1"></a>

## update property — timeouts / b07600b5c2b6 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-295665cf9d41b092d9da32c7cc02eec63c8af18ab4d7eaf33c6e3a9ef36e95d7"></a>

## Next pages — timeouts / b07600b5c2b6 / 8

- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01ccae5df0828823f0e22d4565245f1b3c700c4fd182b650abb557a4014a93cb"></a>

## tls_tcp — tls_tcp / ee43121eb5b8 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- tls_tcp

<a id="canonical-f5d11b5136db56e396813bcb2d1d10476e78227b0523aa6575c4c8c4a64cc8fb"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting TLS over TCP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_parameters")}
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
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
tls_tcp {
  # Configure direct properties listed below.
}
```

<a id="canonical-2fd7fdcf9fbcd823c4ae858802d5c576d3f037f29e02b50cb71666c908a601db"></a>

## Direct properties — tls_tcp / ee43121eb5b8 / 3

- [tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71): complete subsection reference.

- [tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e): complete subsection reference.

<a id="canonical-219e7c66a54985e2040ca42de2cb1e39ac790fddc4587e7e3a782f4ea80ffe49"></a>

## Next pages — tls_tcp / ee43121eb5b8 / 4

- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-937e038bb1a5f952bab297751a522b1c0049f18976f019590012c7627db417db"></a>

## tls_tcp.tls_cert_params — tls_tcp.tls_cert_params / 0530253505a0 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- tls_tcp.tls_cert_params

<a id="canonical-839adb6c9acef80e1b4a4ae9b1fd31890a98e94682061014c982ad34abdca7ab"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-b5bfea644e30c6d0ab9e87dddf717caf719c86fd65fff97636785b9f83164562"></a>

## Direct properties — tls_tcp.tls_cert_params / 0530253505a0 / 3

- [certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-e750348f735958838bf78538df58a316a3e30f862d1605e41b551315218fb803): complete subsection reference.

- [no_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-05efd86baf3eaecd1b27e74b9170860f7a1131ca0028b822bf999eee4d78e563): complete subsection reference.

- [tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3a55666bced2da14311845cecc0411ceea22567a34cacae0b435db588a70ac9c): complete subsection reference.

- [use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-83c630c92f25a665af9d5956afb83400342df1dd00376ece7ebb575449364af7): complete subsection reference.

<a id="canonical-12dbc5233001f4f93e143cb8acd2061e7ec01d2ef5ae65d1b4585bb8917c4cb0"></a>

## Next pages — tls_tcp.tls_cert_params / 0530253505a0 / 4

- [tls_tcp.tls_cert_params.certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-e750348f735958838bf78538df58a316a3e30f862d1605e41b551315218fb803)
- [tls_tcp.tls_cert_params.no_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-05efd86baf3eaecd1b27e74b9170860f7a1131ca0028b822bf999eee4d78e563)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3a55666bced2da14311845cecc0411ceea22567a34cacae0b435db588a70ac9c)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-83c630c92f25a665af9d5956afb83400342df1dd00376ece7ebb575449364af7)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-e750348f735958838bf78538df58a316a3e30f862d1605e41b551315218fb803"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bdf272139056bc3d7f2e9fb7a724c699790d45ba2905da6b3b33124da12e6ba7"></a>

## tls_tcp.tls_cert_params.certificates — tls_tcp.tls_cert_params.certificates / 19a0e03b7d01 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- tls_tcp.tls_cert_params.certificates

<a id="canonical-4bf7c44f4821368c9bf9db15ac4591e64bc6b15685ccf63d895a097417ad77bf"></a>

Type: `"object"`. list nested block, Optional.

Select one or more certificates with any domain names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-9138c9ed9a1a89417579d97a286327cf2c6829933fe5bf91ce0ee0bc77c7e4fc"></a>

## Direct properties — tls_tcp.tls_cert_params.certificates / 19a0e03b7d01 / 3

<a id="canonical-065d444cda11e11777d8b0c74b0bfce39060f6056cc12a6a091ce8048de4c4af"></a>

<a id="canonical-1269d7ec7e0cffa6569dada7ab9360effee0343bc6aee4f5fa7b9c1abb911573"></a>

## name property — tls_tcp.tls_cert_params.certificates / 19a0e03b7d01 / 4

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

<a id="canonical-e119b8645677106a5f8706c4278fe3f3f331cf1b38c667246ff70793776a8877"></a>

<a id="canonical-2a7f7d282ade1f8e907eeaae9d1c6faa2ce8c5c883978f3b4ee5dc4836083b57"></a>

## namespace property — tls_tcp.tls_cert_params.certificates / 19a0e03b7d01 / 5

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

<a id="canonical-e954d9bc1a092d2cb105bce4ffb42ad8c46f08fcb84a188ca0f782623d2225c4"></a>

<a id="canonical-ad5de7d113f79dc528507890e73a62b40fe53fd044ac8ad2f1d397911df489a5"></a>

## tenant property — tls_tcp.tls_cert_params.certificates / 19a0e03b7d01 / 6

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

<a id="canonical-1b14e30bf3651771f17faaf168d36d53afa4ddab9b8e2a348f7c46948faafe82"></a>

## Next pages — tls_tcp.tls_cert_params.certificates / 19a0e03b7d01 / 7

- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-05efd86baf3eaecd1b27e74b9170860f7a1131ca0028b822bf999eee4d78e563"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2c1123168d423eb2972df570c3dae32887dddadbb3577c5b14bc3c62ac35bb4"></a>

## tls_tcp.tls_cert_params.no_mtls — tls_tcp.tls_cert_params.no_mtls / 1aa722a647e6 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- tls_tcp.tls_cert_params.no_mtls

<a id="canonical-2d3bdbf3e98bb9d69ac82de23ada90ed2c82656feab7063d25cda4b6817fe5be"></a>

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
no_mtls = {}
```

<a id="canonical-c6b9bde3f7fb3347b429596d7de3661c81e46b642359db68b2304b0a181ffad6"></a>

## Direct properties — tls_tcp.tls_cert_params.no_mtls / 1aa722a647e6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-daa3d42c1ec13df359e0c713ca7c4a08a5c3cc3c1c99822c47347f66f2315bcf"></a>

## Next pages — tls_tcp.tls_cert_params.no_mtls / 1aa722a647e6 / 4

- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-3a55666bced2da14311845cecc0411ceea22567a34cacae0b435db588a70ac9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-17a1791e32d01598709404d90e20cde1ddb661d68fb9a54af6d09e7c135d5747"></a>

## tls_tcp.tls_cert_params.tls_config — tls_tcp.tls_cert_params.tls_config / d0c1b6944f98 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- tls_tcp.tls_cert_params.tls_config

<a id="canonical-30ba8f8e6931cac9d549246cf322a7be2950e02978374eb7816a600afeb22f09"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-506ac35520be8e7f29f4834e0c5842c488f98523471667f0b782eee312548be8"></a>

## Direct properties — tls_tcp.tls_cert_params.tls_config / d0c1b6944f98 / 3

- [custom_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-f8a23923c4de78b20763d5c6fd928dc033a80ea6c066cdc1bee195b4d6a4cd3d): complete subsection reference.

- [default_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-a56fdbf69d08173c4f3c088257632c85b3e6eba70c9ed9ca096f4223268e2431): complete subsection reference.

- [low_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-e58fcdc5714ac2ec699c92079341522d6d42bbe8d2b978f32a51858e9f356217): complete subsection reference.

- [medium_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-fd228c80022dc65cf2b74c513c06e4e1a82b04e8a93aa0f4ea24ae82b6b9c615): complete subsection reference.

<a id="canonical-59fb2389d8c9c8aeddd7e256d102941d64f39d97a1c521cc0989930c14f6cb6f"></a>

## Next pages — tls_tcp.tls_cert_params.tls_config / d0c1b6944f98 / 4

- [tls_tcp.tls_cert_params.tls_config.custom_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-f8a23923c4de78b20763d5c6fd928dc033a80ea6c066cdc1bee195b4d6a4cd3d)
- [tls_tcp.tls_cert_params.tls_config.default_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-a56fdbf69d08173c4f3c088257632c85b3e6eba70c9ed9ca096f4223268e2431)
- [tls_tcp.tls_cert_params.tls_config.low_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-e58fcdc5714ac2ec699c92079341522d6d42bbe8d2b978f32a51858e9f356217)
- [tls_tcp.tls_cert_params.tls_config.medium_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-fd228c80022dc65cf2b74c513c06e4e1a82b04e8a93aa0f4ea24ae82b6b9c615)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-f8a23923c4de78b20763d5c6fd928dc033a80ea6c066cdc1bee195b4d6a4cd3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d7ca1d89489d302b9ab4adb061e2440052af058f6d09440f79c71af23076984"></a>

## tls_tcp.tls_cert_params.tls_config.custom_security — tls_tcp.tls_cert_params.tls_config.custom_security / d7648fc3f612 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3a55666bced2da14311845cecc0411ceea22567a34cacae0b435db588a70ac9c)
- tls_tcp.tls_cert_params.tls_config.custom_security

<a id="canonical-97efd0877528b98ff62a6dce1a9c9d997394c4587c0b3e2c0efd85c2bbda6469"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-9ea5ebd5768ebf215f7815f80e5b3a9bd8e5f7677a999486504f26f6c06c70a5"></a>

## Direct properties — tls_tcp.tls_cert_params.tls_config.custom_security / d7648fc3f612 / 3

<a id="canonical-d958ecdd97ffc497b20aa15b791b99dab14168ca7203dc9109cb4ab14b9b8e1e"></a>

<a id="canonical-9a3322996f820aa926d93aaf20de48616e44a72be37674335667e9f26d9153d1"></a>

## cipher_suites property — tls_tcp.tls_cert_params.tls_config.custom_security / d7648fc3f612 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-a73e9cce9fbaf9a80461468d5dfac06d4e0c63006188843392d5deae76112830"></a>

<a id="canonical-369617d9838a4b827920690ae1450f432d3d1bc844855fee810197f54271846a"></a>

## max_version property — tls_tcp.tls_cert_params.tls_config.custom_security / d7648fc3f612 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-59af86cf863b8a767395402ffeb09e4ca231596c37f04d12297f363b1256a33e"></a>

<a id="canonical-b03357e1ff00f35bdfe6b70f3a4a15ad21fa424c492aecb84b51faa82a03a32a"></a>

## min_version property — tls_tcp.tls_cert_params.tls_config.custom_security / d7648fc3f612 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-2f472593d76bdc4cb05db88c7b15548864b6286f86b865b548f849078b411b93"></a>

## Next pages — tls_tcp.tls_cert_params.tls_config.custom_security / d7648fc3f612 / 7

- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3a55666bced2da14311845cecc0411ceea22567a34cacae0b435db588a70ac9c)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-a56fdbf69d08173c4f3c088257632c85b3e6eba70c9ed9ca096f4223268e2431"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc9ab823ffb9484dbebf18d539edbe904250d202011b1195223c23af58fc59fa"></a>

## tls_tcp.tls_cert_params.tls_config.default_security — tls_tcp.tls_cert_params.tls_config.default_security / 6d99173dab4b / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3a55666bced2da14311845cecc0411ceea22567a34cacae0b435db588a70ac9c)
- tls_tcp.tls_cert_params.tls_config.default_security

<a id="canonical-76445d70996e5f723a036b8c79095576a550eb4e0c9c85068ae5809cf7302c06"></a>

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
default_security = {}
```

<a id="canonical-5a6658a477b6596d8319e8df71cb37666546b9e745bec3121f4416f10972c714"></a>

## Direct properties — tls_tcp.tls_cert_params.tls_config.default_security / 6d99173dab4b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-346c56dcd4002dfec6d754d17b27cd45ff47897874eae42c25f473ff8733d984"></a>

## Next pages — tls_tcp.tls_cert_params.tls_config.default_security / 6d99173dab4b / 4

- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3a55666bced2da14311845cecc0411ceea22567a34cacae0b435db588a70ac9c)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-e58fcdc5714ac2ec699c92079341522d6d42bbe8d2b978f32a51858e9f356217"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4efcb21f4aee72366834417988c0ccccd8eaed9403e5f5594e0e19f4541f82b8"></a>

## tls_tcp.tls_cert_params.tls_config.low_security — tls_tcp.tls_cert_params.tls_config.low_security / 9a1dba6dd045 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3a55666bced2da14311845cecc0411ceea22567a34cacae0b435db588a70ac9c)
- tls_tcp.tls_cert_params.tls_config.low_security

<a id="canonical-7c537e85348f88cdd9d2368711b4dfd2e1a44e59b592414380af4d0add268ef8"></a>

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
low_security = {}
```

<a id="canonical-2496a5836a3fa3ffbbaa0536ed296aff83eea0f25195a1e6848ea761ca8bda1a"></a>

## Direct properties — tls_tcp.tls_cert_params.tls_config.low_security / 9a1dba6dd045 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dc20f8af20c4bc23b6d7940f10e648b6a5acc5617a807ab1e4dd804ec7a49166"></a>

## Next pages — tls_tcp.tls_cert_params.tls_config.low_security / 9a1dba6dd045 / 4

- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3a55666bced2da14311845cecc0411ceea22567a34cacae0b435db588a70ac9c)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-fd228c80022dc65cf2b74c513c06e4e1a82b04e8a93aa0f4ea24ae82b6b9c615"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58f02b38f460af19ad9bef95f71c3123672fa8221d8f6cfb6ab03ba70846d2cd"></a>

## tls_tcp.tls_cert_params.tls_config.medium_security — tls_tcp.tls_cert_params.tls_config.medium_security / fcb67b8da3d2 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3a55666bced2da14311845cecc0411ceea22567a34cacae0b435db588a70ac9c)
- tls_tcp.tls_cert_params.tls_config.medium_security

<a id="canonical-a62bf1ba751fa8a13eb16dcd085d47c9e04fc8ca4dd711b1843267fed018f980"></a>

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
medium_security = {}
```

<a id="canonical-498f10ff12e4078b0c9703de22ae518adf7beb4b165036ff282d32c44d797a5b"></a>

## Direct properties — tls_tcp.tls_cert_params.tls_config.medium_security / fcb67b8da3d2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f804186116d9a3c35c90cdef857fb70cb7efddd7a0da374b017d3963b7919e91"></a>

## Next pages — tls_tcp.tls_cert_params.tls_config.medium_security / fcb67b8da3d2 / 4

- [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-3a55666bced2da14311845cecc0411ceea22567a34cacae0b435db588a70ac9c)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-83c630c92f25a665af9d5956afb83400342df1dd00376ece7ebb575449364af7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a8da8cc0189bf6849d0d2f1b61c17b2372573652f11c810a2b5681d5d7f3ed6"></a>

## tls_tcp.tls_cert_params.use_mtls — tls_tcp.tls_cert_params.use_mtls / 75afe51217d5 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- tls_tcp.tls_cert_params.use_mtls

<a id="canonical-d564850dc8103f3a5a2e16b1fa8495a26b5e284453498560ca6c4c4167472ace"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-06c0f519fbc92f95bbd1f29a6be7e4ed7c7bd77b16229949ed91c04f7343f241"></a>

## Direct properties — tls_tcp.tls_cert_params.use_mtls / 75afe51217d5 / 3

<a id="canonical-c7e25aae57d84d4287f7acceda456ba799fdcf08b7b9df696d84d89480f20df4"></a>

<a id="canonical-db2e8e480bc51f7f156e0d02f062ffbceb3d1b094203253bbd456979922e317a"></a>

## client_certificate_optional property — tls_tcp.tls_cert_params.use_mtls / 75afe51217d5 / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-078eea5add57e80c63e8e3a1c8c695fa8da7ff8a6be7938f6cc6b660c85a2fca): complete subsection reference.

- [no_crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-d6b152f5d149bef6060ae04cb43c1b54ecda9c52811db23e5760f5742096846e): complete subsection reference.

- [trusted_ca](resources--tcp_loadbalancer--reference--group-002.md#canonical-200d0320946fecacc9bce889970f1dea4413d8da7d3b11abef0c1e8fe126db2e): complete subsection reference.

<a id="canonical-7d50d6041f589aaff46e34535832fc3c926e94278e4e9d7bc9861839c20d1427"></a>

<a id="canonical-506391bca2cd8590e298d6fe288638bf00d92155535a0485d2f0ed4d50737841"></a>

## trusted_ca_url property — tls_tcp.tls_cert_params.use_mtls / 75afe51217d5 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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

- [xfcc_disabled](resources--tcp_loadbalancer--reference--group-002.md#canonical-2e0e80e281c83d06d1c16b5b807731c72cddc156a1b16bedb551b7a4299d6375): complete subsection reference.

- [xfcc_options](resources--tcp_loadbalancer--reference--group-002.md#canonical-a5be3541169fd9b88c4189515b8eba0d27d524047a7b98ac64131144ed23d6f8): complete subsection reference.

<a id="canonical-5292d8b566adcd70e806aff67c0e13236f4f6f5a85c338f3eb2dced716085459"></a>

## Next pages — tls_tcp.tls_cert_params.use_mtls / 75afe51217d5 / 6

- [tls_tcp.tls_cert_params.use_mtls.crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-078eea5add57e80c63e8e3a1c8c695fa8da7ff8a6be7938f6cc6b660c85a2fca)
- [tls_tcp.tls_cert_params.use_mtls.no_crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-d6b152f5d149bef6060ae04cb43c1b54ecda9c52811db23e5760f5742096846e)
- [tls_tcp.tls_cert_params.use_mtls.trusted_ca](resources--tcp_loadbalancer--reference--group-002.md#canonical-200d0320946fecacc9bce889970f1dea4413d8da7d3b11abef0c1e8fe126db2e)
- [tls_tcp.tls_cert_params.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--reference--group-002.md#canonical-2e0e80e281c83d06d1c16b5b807731c72cddc156a1b16bedb551b7a4299d6375)
- [tls_tcp.tls_cert_params.use_mtls.xfcc_options](resources--tcp_loadbalancer--reference--group-002.md#canonical-a5be3541169fd9b88c4189515b8eba0d27d524047a7b98ac64131144ed23d6f8)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-078eea5add57e80c63e8e3a1c8c695fa8da7ff8a6be7938f6cc6b660c85a2fca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4632c9aabebf69fd12f64384c39d69e0b8450954578f1800bc0efb300214931"></a>

## tls_tcp.tls_cert_params.use_mtls.crl — tls_tcp.tls_cert_params.use_mtls.crl / e6f420b04b96 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-83c630c92f25a665af9d5956afb83400342df1dd00376ece7ebb575449364af7)
- tls_tcp.tls_cert_params.use_mtls.crl

<a id="canonical-2102e226a2d5162573054dda7b5c5ff20bc00f5f91ade2ac81e4278e5578a47b"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-557a7b2735b7477b761691dff10b6c8462636270f865600dd2389d7d78fd1138"></a>

## Direct properties — tls_tcp.tls_cert_params.use_mtls.crl / e6f420b04b96 / 3

<a id="canonical-1aa399ffce5889d600757cb51cde34dd0dd8a5b1b78ba73b95fc0fce9faefe15"></a>

<a id="canonical-75a1c4867266324cf8bac53a59c3593b01240cace568061333f391a10fa851ca"></a>

## name property — tls_tcp.tls_cert_params.use_mtls.crl / e6f420b04b96 / 4

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

<a id="canonical-ea960781cda343ce74d0e3f8f0a23a2153e5a12f4b0623428e0f9362abab779a"></a>

<a id="canonical-75d5ceb8a3db84d0884ee64681c75f6119936d39187a0384fef9169a7d4cef83"></a>

## namespace property — tls_tcp.tls_cert_params.use_mtls.crl / e6f420b04b96 / 5

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

<a id="canonical-e241de93108b158dafd74cb1fbad1fa5c8d30c0fa55fddd3b5db058ff2ccfc02"></a>

<a id="canonical-c68900fcc80ab0e845511792fc039cb0aad46b9d3bfefc8543946556e632250e"></a>

## tenant property — tls_tcp.tls_cert_params.use_mtls.crl / e6f420b04b96 / 6

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

<a id="canonical-ccc9370d9620d8e385014f65a33ef9e7d645681ca0049627f531588145aa70cb"></a>

## Next pages — tls_tcp.tls_cert_params.use_mtls.crl / e6f420b04b96 / 7

- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-83c630c92f25a665af9d5956afb83400342df1dd00376ece7ebb575449364af7)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-d6b152f5d149bef6060ae04cb43c1b54ecda9c52811db23e5760f5742096846e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7dcded8f8fc3a66f755c52b947558069209ac0d82b5cf1f97a06d6cba85b8b61"></a>

## tls_tcp.tls_cert_params.use_mtls.no_crl — tls_tcp.tls_cert_params.use_mtls.no_crl / 3d0e43b3fb0a / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-83c630c92f25a665af9d5956afb83400342df1dd00376ece7ebb575449364af7)
- tls_tcp.tls_cert_params.use_mtls.no_crl

<a id="canonical-f91a2cbab02af12b020b84d0a19b0f669dbb2bccb5e755d650ad7de0a35b487f"></a>

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
no_crl = {}
```

<a id="canonical-a12f0805382d75b6da76b17f8d78a2c7de06019dce8bdbf69c5d43c86f24e84c"></a>

## Direct properties — tls_tcp.tls_cert_params.use_mtls.no_crl / 3d0e43b3fb0a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a6b585c4dcaa5c7fd90d3fbf5e4931c7743f96f0305883113c7d7d184b9a0ff8"></a>

## Next pages — tls_tcp.tls_cert_params.use_mtls.no_crl / 3d0e43b3fb0a / 4

- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-83c630c92f25a665af9d5956afb83400342df1dd00376ece7ebb575449364af7)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-200d0320946fecacc9bce889970f1dea4413d8da7d3b11abef0c1e8fe126db2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bda8ab98c79110d1a51e7b6acfd5d5d9cd3d9d57c52cf60f23b93f90a8a713ea"></a>

## tls_tcp.tls_cert_params.use_mtls.trusted_ca — tls_tcp.tls_cert_params.use_mtls.trusted_ca / 8875f96ee766 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-83c630c92f25a665af9d5956afb83400342df1dd00376ece7ebb575449364af7)
- tls_tcp.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-81662e01a2b658548318786c28df285fbc043aa0467f119ebebfbf3272370026"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-097e6962c47eecf28d01ca0d5b9b1b9df12f4b50595474c5d53a72d308be3ab7"></a>

## Direct properties — tls_tcp.tls_cert_params.use_mtls.trusted_ca / 8875f96ee766 / 3

<a id="canonical-0c0fe0b1cf9d367ea2439cc8ce05f41378df1e4353bd7757989481353da56c4f"></a>

<a id="canonical-efbb6c66e7372bc8baba9f1e5e8d6e8b8f5d675c4cfff59babdc0153fd5ea243"></a>

## name property — tls_tcp.tls_cert_params.use_mtls.trusted_ca / 8875f96ee766 / 4

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

<a id="canonical-cfceeff7da90d114521d0bffe90ca9ee5fbc422a5169f49b38800149776be32a"></a>

<a id="canonical-4f90f8289f3751b1686d6f916b7ca822984c8568e7f040e514827cf96e6937a2"></a>

## namespace property — tls_tcp.tls_cert_params.use_mtls.trusted_ca / 8875f96ee766 / 5

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

<a id="canonical-1f771d732064a50e77c8c6cf80ff8c7b8b302e2e43c913fe96cb0546007c9aa7"></a>

<a id="canonical-639496a4689fb1af29a72be46bb62ee9ac5a5d270760f35137ea249ddf19e7d3"></a>

## tenant property — tls_tcp.tls_cert_params.use_mtls.trusted_ca / 8875f96ee766 / 6

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

<a id="canonical-1eb8d21fa660fd14d5e4850e893707b163cee559d4322c21db0d423e8ed1c027"></a>

## Next pages — tls_tcp.tls_cert_params.use_mtls.trusted_ca / 8875f96ee766 / 7

- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-83c630c92f25a665af9d5956afb83400342df1dd00376ece7ebb575449364af7)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-2e0e80e281c83d06d1c16b5b807731c72cddc156a1b16bedb551b7a4299d6375"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6119092062c68fbb89f47dc7bc1958dba7c46e51df1a1b17e49f7dff7bcc9a92"></a>

## tls_tcp.tls_cert_params.use_mtls.xfcc_disabled — tls_tcp.tls_cert_params.use_mtls.xfcc_disabled / e9ab13b627a0 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-83c630c92f25a665af9d5956afb83400342df1dd00376ece7ebb575449364af7)
- tls_tcp.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-520e0626013ad608c41aedf492771ce3949f9e4213407b8e461bc798c6546b2c"></a>

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
xfcc_disabled = {}
```

<a id="canonical-76449cda5f3c26c2503efa3c6ce0993f910d77f039d4957214c62c6f93146658"></a>

## Direct properties — tls_tcp.tls_cert_params.use_mtls.xfcc_disabled / e9ab13b627a0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-43aed4bab2db427109505f1bb83eda56964d1428b08aaa436703c6c50297b575"></a>

## Next pages — tls_tcp.tls_cert_params.use_mtls.xfcc_disabled / e9ab13b627a0 / 4

- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-83c630c92f25a665af9d5956afb83400342df1dd00376ece7ebb575449364af7)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-a5be3541169fd9b88c4189515b8eba0d27d524047a7b98ac64131144ed23d6f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91cf0f89567f4f150a67fbed68e14f0ab876eeb98a640141ebeb5506e9878051"></a>

## tls_tcp.tls_cert_params.use_mtls.xfcc_options — tls_tcp.tls_cert_params.use_mtls.xfcc_options / a7c31bbfcfe8 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-92c860fcd66dc198bb9b51e5ae9dfb0891f265341b0814a0cb90e68b8c7a8e71)
- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-83c630c92f25a665af9d5956afb83400342df1dd00376ece7ebb575449364af7)
- tls_tcp.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-817f51e33de2a08d48b000736ce01f7a0f77bfbe94c4b1f33f2d858ffff02d0a"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-d93986e6f60e1f0da841191a5c9dc04355d3fb25554257770b545fff987a70b8"></a>

## Direct properties — tls_tcp.tls_cert_params.use_mtls.xfcc_options / a7c31bbfcfe8 / 3

<a id="canonical-d427cd4bbc4532c4b4824041fc4b26196d877455265ba7eed62e559f2b8f6b35"></a>

<a id="canonical-3eb05a4d1c69f2642b47df32b380212f1942ab4d26e47f67b1bf3d3d5b076c9a"></a>

## xfcc_header_elements property — tls_tcp.tls_cert_params.use_mtls.xfcc_options / a7c31bbfcfe8 / 4

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-9c025853cadd1d550be4e6aa251ae054186efea4612bedcdf84f2b08841c4ac7"></a>

## Next pages — tls_tcp.tls_cert_params.use_mtls.xfcc_options / a7c31bbfcfe8 / 5

- [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-83c630c92f25a665af9d5956afb83400342df1dd00376ece7ebb575449364af7)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c9c151ebb106d2bcc218e96ba8836d1f1eb305f9a0ddeaa833e9d95f3d750ce"></a>

## tls_tcp.tls_parameters — tls_tcp.tls_parameters / c588dd436cc7 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- tls_tcp.tls_parameters

<a id="canonical-d4769c860887e9b2ae5d731f715a34ff5b5839a189b45decccd16f21cb5ef71d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-a99e2e3116f208853fb18540d44396446c80e63a58a42f72374cd82e7c983407"></a>

## Direct properties — tls_tcp.tls_parameters / c588dd436cc7 / 3

- [no_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-f6392d1ef7fcfb4f8f1e198d360fbaba7f49c1bef7c1f680479d5b2d7f76a4e9): complete subsection reference.

- [tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-02e0a85c343dfad6647457492f01e1820c87872bc877abb87c17c5902dc67d58): complete subsection reference.

- [tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-f694c7eee4c07b43d127a66c1d602afb6d309130b757dccd7620fc214cefee83): complete subsection reference.

- [use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-e6c9f64ac3d283ff93179b5d7b14091b6266db6afffc6ee45c4e1f53d3445ec0): complete subsection reference.

<a id="canonical-ae891592785f3e93c43747cc7e6dfdc098edc8a49defcef572167563ae460d13"></a>

## Next pages — tls_tcp.tls_parameters / c588dd436cc7 / 4

- [tls_tcp.tls_parameters.no_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-f6392d1ef7fcfb4f8f1e198d360fbaba7f49c1bef7c1f680479d5b2d7f76a4e9)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-02e0a85c343dfad6647457492f01e1820c87872bc877abb87c17c5902dc67d58)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-f694c7eee4c07b43d127a66c1d602afb6d309130b757dccd7620fc214cefee83)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-e6c9f64ac3d283ff93179b5d7b14091b6266db6afffc6ee45c4e1f53d3445ec0)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-f6392d1ef7fcfb4f8f1e198d360fbaba7f49c1bef7c1f680479d5b2d7f76a4e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ba3e978bf443f83afce415af5d230e5882693fd5bcfc6178bd6b72b55579364"></a>

## tls_tcp.tls_parameters.no_mtls — tls_tcp.tls_parameters.no_mtls / 91a25ab2d7c5 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- tls_tcp.tls_parameters.no_mtls

<a id="canonical-51e63514831425ccde074d2f430f29f8315ee7a2512e9ab1ea172009876d9efb"></a>

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
no_mtls = {}
```

<a id="canonical-2f8474ab55b46e414c7786bc9cbf65690294edc683183495948ada5046435b63"></a>

## Direct properties — tls_tcp.tls_parameters.no_mtls / 91a25ab2d7c5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e902ca77f0cf88781324274a70a980ccb918b32340ad0f19eb3167842e6c5656"></a>

## Next pages — tls_tcp.tls_parameters.no_mtls / 91a25ab2d7c5 / 4

- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-02e0a85c343dfad6647457492f01e1820c87872bc877abb87c17c5902dc67d58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d7a5b4ff9bd06604613af6bf825ed3b59e00415ed8a126349c16d70b22df55c"></a>

## tls_tcp.tls_parameters.tls_certificates — tls_tcp.tls_parameters.tls_certificates / 77d580e6812f / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- tls_tcp.tls_parameters.tls_certificates

<a id="canonical-62b666102cdee1a411a3fb10df10554fd099d0dc8e559389e7607e206bae9ec6"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-102dfffa970cc4f458be42d0c6b5b11b90da7df98ccb8863be4541734f76dc95"></a>

## Direct properties — tls_tcp.tls_parameters.tls_certificates / 77d580e6812f / 3

<a id="canonical-fbdb0eb311cd17f81736c112164a63cb509fb2272fb32b2d8d90eb6cc181c5a6"></a>

<a id="canonical-a9c9e8137301a37eab77fc9d536163fc4cb5d905c5943f60d3431eac3ff25932"></a>

## certificate_url property — tls_tcp.tls_parameters.tls_certificates / 77d580e6812f / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](resources--tcp_loadbalancer--reference--group-002.md#canonical-f1cda88e10b001eac9166d442c71964679348b0aa9dbb936bf11346d37599d97): complete subsection reference.

<a id="canonical-a7c113dcf28f29d73afbd90621e603a4c8ff43401f1722776f0a4e16d7ecf6bc"></a>

<a id="canonical-b03695f8d1e6789a20d4e3e104aecd577d35f25a20be30141048e5110a361b37"></a>

## description_spec property — tls_tcp.tls_parameters.tls_certificates / 77d580e6812f / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--tcp_loadbalancer--reference--group-002.md#canonical-5c7ee54c77ba3cd4dd80cfeb3f3eb0002a91bec8227306d25a519fef43293826): complete subsection reference.

- [private_key](resources--tcp_loadbalancer--reference--group-002.md#canonical-10fd2470e881e06b0c62cc1afd86ebb26ec57e583ca4532394fbd904895bb0c7): complete subsection reference.

- [use_system_defaults](resources--tcp_loadbalancer--reference--group-002.md#canonical-121189fdebacb5211ab94ca63b53d4f4cd1a115347142805e211bcf96dab275e): complete subsection reference.

<a id="canonical-adb1921021a476b22b6d5475464542b1124defeab7c13dbba9c0766e1c83a828"></a>

## Next pages — tls_tcp.tls_parameters.tls_certificates / 77d580e6812f / 6

- [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms](resources--tcp_loadbalancer--reference--group-002.md#canonical-f1cda88e10b001eac9166d442c71964679348b0aa9dbb936bf11346d37599d97)
- [tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--tcp_loadbalancer--reference--group-002.md#canonical-5c7ee54c77ba3cd4dd80cfeb3f3eb0002a91bec8227306d25a519fef43293826)
- [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--reference--group-002.md#canonical-10fd2470e881e06b0c62cc1afd86ebb26ec57e583ca4532394fbd904895bb0c7)
- [tls_tcp.tls_parameters.tls_certificates.use_system_defaults](resources--tcp_loadbalancer--reference--group-002.md#canonical-121189fdebacb5211ab94ca63b53d4f4cd1a115347142805e211bcf96dab275e)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-f1cda88e10b001eac9166d442c71964679348b0aa9dbb936bf11346d37599d97"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6cb5d35bee9724bedb70a131e4252996ec7dc7ebb4e7d1a1a23ca2d7a6eab261"></a>

## tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms — tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms / 9a93ca378a8e / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-02e0a85c343dfad6647457492f01e1820c87872bc877abb87c17c5902dc67d58)
- tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-cb3848221be67cb4201786c8b13fa885e91479c4c0b43ee0d1aadd44100cfbaf"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-acf7464da0e4fc179ff6063724e73418501056747cb251512501ba2c9b68abbe"></a>

## Direct properties — tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms / 9a93ca378a8e / 3

<a id="canonical-a6a186f99f370f3d15afeafeed3dded1928e5689408af507fdcd786a104f4025"></a>

<a id="canonical-398b1a1396ca3dfdd237f909a1c20b8f544c8e4b8007341fdba2104f8afd6bd3"></a>

## hash_algorithms property — tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms / 9a93ca378a8e / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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

<a id="canonical-f0dbe31725fbcb1b6c928e1bbcccf2def2702ae291e7f3163a04645631fd8035"></a>

## Next pages — tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms / 9a93ca378a8e / 5

- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-02e0a85c343dfad6647457492f01e1820c87872bc877abb87c17c5902dc67d58)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-5c7ee54c77ba3cd4dd80cfeb3f3eb0002a91bec8227306d25a519fef43293826"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c03601a080659bb2b86f7a26a4156f4ec16c6e1670bcd1a66d954c9a9f6a3ae2"></a>

## tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling — tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling / 161adefecdab / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-02e0a85c343dfad6647457492f01e1820c87872bc877abb87c17c5902dc67d58)
- tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-cbd45497c9ae17d67b95902e3393b833bc7073c4c238b40126ef7b6e99b9c6bf"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

<a id="canonical-df9479ae40b35a830f0b6230f221faecc13f979c2e8a3db354ec80cf0aa85a5e"></a>

## Direct properties — tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling / 161adefecdab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d03fa563789b5aca84836ed6e018194354571ed8a719c6840d9a1a7ef7f964bc"></a>

## Next pages — tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling / 161adefecdab / 4

- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-02e0a85c343dfad6647457492f01e1820c87872bc877abb87c17c5902dc67d58)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-10fd2470e881e06b0c62cc1afd86ebb26ec57e583ca4532394fbd904895bb0c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a5ec4982bb509299cb0d14ad906e276c2b1e1cf90fc9e72e3d81a718c52c554"></a>

## tls_tcp.tls_parameters.tls_certificates.private_key — tls_tcp.tls_parameters.tls_certificates.private_key / 16732b9f62a8 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-02e0a85c343dfad6647457492f01e1820c87872bc877abb87c17c5902dc67d58)
- tls_tcp.tls_parameters.tls_certificates.private_key

<a id="canonical-5b1748d7c91697be4f4f1d574040d2e87f587b86005a97a7fa3f34edabb5450e"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-00944030c00e431b7b64771e93328578cf3dffbd35fad65f5731b3cb48d57e65"></a>

## Direct properties — tls_tcp.tls_parameters.tls_certificates.private_key / 16732b9f62a8 / 3

- [blindfold_secret_info](resources--tcp_loadbalancer--reference--group-002.md#canonical-38df892c426fa42d73958d05412c08441427112858934af1f80bbf813beee724): complete subsection reference.

- [clear_secret_info](resources--tcp_loadbalancer--reference--group-002.md#canonical-2363805873d3a47d7fb9d2bf513acfe9a143d69d35adafbc230f074784823bfa): complete subsection reference.

<a id="canonical-0659b2b5295ff206072308e178f12065e09293c0f1b0485fca205ece2be1a140"></a>

## Next pages — tls_tcp.tls_parameters.tls_certificates.private_key / 16732b9f62a8 / 4

- [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--tcp_loadbalancer--reference--group-002.md#canonical-38df892c426fa42d73958d05412c08441427112858934af1f80bbf813beee724)
- [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--tcp_loadbalancer--reference--group-002.md#canonical-2363805873d3a47d7fb9d2bf513acfe9a143d69d35adafbc230f074784823bfa)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-02e0a85c343dfad6647457492f01e1820c87872bc877abb87c17c5902dc67d58)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-38df892c426fa42d73958d05412c08441427112858934af1f80bbf813beee724"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3bca759e7015a7d9c10b3bf3e4fde6d0e0029262af7c83ccc65383f4fdecf70"></a>

## tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info — tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 565db3000d8a / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-02e0a85c343dfad6647457492f01e1820c87872bc877abb87c17c5902dc67d58)
- [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--reference--group-002.md#canonical-10fd2470e881e06b0c62cc1afd86ebb26ec57e583ca4532394fbd904895bb0c7)
- tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-fec46dcc08b92908c0d55fea9f584a98ea1b138234b91b82fbb1d687bdbb139e"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-f01ac9df7271be62c76ad53ffcbb7215de983a9d02fd10c0d0ca6bddd82c3180"></a>

## Direct properties — tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 565db3000d8a / 3

<a id="canonical-8692c4721e8eb2ae682cad958e59eb3cb1a8fe153c237539200c3837059c05fd"></a>

<a id="canonical-7c1982d350c9782ea5e3edf8dfe78f928ce65cebe4ae9e4fc343e61a2554efb9"></a>

## decryption_provider property — tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 565db3000d8a / 4

Type: `"string"`. Optional.

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

<a id="canonical-f3c6f480b6ecaa2e2bb1fb47918457f5d7ab318a07ed3b59d9b3765dc5c33a1e"></a>

<a id="canonical-2a009bcd6d6e11c57a4f63ef65cf6bb6ff948fa146444d7af9a17ac8f2c32407"></a>

## location property — tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 565db3000d8a / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-4bf6be6b098efb4fd44eb8c833a644bb6e8dda1c9702230b481c75d08d2b92b9"></a>

<a id="canonical-f1467f3e4ac1e4ac7772634dc1d0690bd5e7211e10959cc7f747c4ed948ed820"></a>

## store_provider property — tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 565db3000d8a / 6

Type: `"string"`. Optional.

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

<a id="canonical-f41219709d00a91a563846467cfdb303ace43b4504fea9ac61776b228b838770"></a>

## Next pages — tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 565db3000d8a / 7

- [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--reference--group-002.md#canonical-10fd2470e881e06b0c62cc1afd86ebb26ec57e583ca4532394fbd904895bb0c7)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-2363805873d3a47d7fb9d2bf513acfe9a143d69d35adafbc230f074784823bfa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7aa0e7b065aa5ca91573184ada10fe5973392487884d8ca96d4bfbeb3990eaf"></a>

## tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info — tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info / 6be2feb22d40 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-02e0a85c343dfad6647457492f01e1820c87872bc877abb87c17c5902dc67d58)
- [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--reference--group-002.md#canonical-10fd2470e881e06b0c62cc1afd86ebb26ec57e583ca4532394fbd904895bb0c7)
- tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-59f990737c34c0febe2e09685755198517f82d62b68c2112c02c044b5856f203"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-f4f85f1857c5216208e9e67cfe4b79ed0d74599981e3ef83ffd134bc7d8efbb9"></a>

## Direct properties — tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info / 6be2feb22d40 / 3

<a id="canonical-12d71bbe2f6e433ca300a70f8e8a48bc3d488b93f77d536123334f99e41f9aee"></a>

<a id="canonical-eabed6d18ace49fadb82bd79e4c1e64ce90f47f2e5b245223ab5dbb72be710d2"></a>

## provider_ref property — tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info / 6be2feb22d40 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1c891faae28b14d8261107515821358dcf87043983c3147852720dc5ddaef4b4"></a>

<a id="canonical-3727fd9bc8fb4f053f49ef926262f71fbb3985c217870a7a62f33b4a12717f8e"></a>

## url property — tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info / 6be2feb22d40 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-0d2a4b6e2948cdb5ee5e55d0c862aef070bfe8707cd990f22af8c4080f9a2834"></a>

## Next pages — tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info / 6be2feb22d40 / 6

- [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--reference--group-002.md#canonical-10fd2470e881e06b0c62cc1afd86ebb26ec57e583ca4532394fbd904895bb0c7)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-121189fdebacb5211ab94ca63b53d4f4cd1a115347142805e211bcf96dab275e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1619560eb8619a03c31c595a1fa01eda4d8bc361648b7d9f482ad9687f60b57e"></a>

## tls_tcp.tls_parameters.tls_certificates.use_system_defaults — tls_tcp.tls_parameters.tls_certificates.use_system_defaults / 324652b36217 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-02e0a85c343dfad6647457492f01e1820c87872bc877abb87c17c5902dc67d58)
- tls_tcp.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-f9f968bd3ea088f4333ea94dd5381109caf22f242cca206d63129879f8a3add9"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_system_defaults = {}
```

<a id="canonical-113e9ad7e2c3eb2dde26f262152c723342e3961c978dbb1b359260b1e09dc410"></a>

## Direct properties — tls_tcp.tls_parameters.tls_certificates.use_system_defaults / 324652b36217 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bac4efc4b46c401efa4f24fa9d6dd7a01013f4bcd1011878a1f8454fbd8d4646"></a>

## Next pages — tls_tcp.tls_parameters.tls_certificates.use_system_defaults / 324652b36217 / 4

- [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-02e0a85c343dfad6647457492f01e1820c87872bc877abb87c17c5902dc67d58)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-f694c7eee4c07b43d127a66c1d602afb6d309130b757dccd7620fc214cefee83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1ecf3aabdccd362440a43f2113053c15267d8e61e2047a2644d6c02daa5a3f2"></a>

## tls_tcp.tls_parameters.tls_config — tls_tcp.tls_parameters.tls_config / 5da5aa2a2006 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- tls_tcp.tls_parameters.tls_config

<a id="canonical-26c0ca004c48c60eec18c10c0026fb93281961932e9c86faae2490720608f07a"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-471cf9e4b0b26fd3113a771d6e3613e4426575c8a71e8af4ea6a03bdb97d586f"></a>

## Direct properties — tls_tcp.tls_parameters.tls_config / 5da5aa2a2006 / 3

- [custom_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-1c6338238c5b83231f6b98876f71faa5ba81a4eb182e650f2d18d81d6f1b1b67): complete subsection reference.

- [default_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-c5dd24d9edbe48690a9c2d374eeb0bf920c87b9ff0dd190fbc9122b3b31c7724): complete subsection reference.

- [low_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-b00c2e128aabbfe4b14555292261de74450c979946ac740132fc8bc615f9a3b3): complete subsection reference.

- [medium_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-1a46641a1babeccacddd1d56deed7f62492d6db20354c88d78a9fc1caa6699e6): complete subsection reference.

<a id="canonical-c83988df7bb74e721a9c3e3c441d8d554f9a7c8e7d22e446c4e05db78519d316"></a>

## Next pages — tls_tcp.tls_parameters.tls_config / 5da5aa2a2006 / 4

- [tls_tcp.tls_parameters.tls_config.custom_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-1c6338238c5b83231f6b98876f71faa5ba81a4eb182e650f2d18d81d6f1b1b67)
- [tls_tcp.tls_parameters.tls_config.default_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-c5dd24d9edbe48690a9c2d374eeb0bf920c87b9ff0dd190fbc9122b3b31c7724)
- [tls_tcp.tls_parameters.tls_config.low_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-b00c2e128aabbfe4b14555292261de74450c979946ac740132fc8bc615f9a3b3)
- [tls_tcp.tls_parameters.tls_config.medium_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-1a46641a1babeccacddd1d56deed7f62492d6db20354c88d78a9fc1caa6699e6)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-1c6338238c5b83231f6b98876f71faa5ba81a4eb182e650f2d18d81d6f1b1b67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2465410ee66d0228749fcceac59c81db69f32b8d2404e29d8ad1d2ca2dfb8a76"></a>

## tls_tcp.tls_parameters.tls_config.custom_security — tls_tcp.tls_parameters.tls_config.custom_security / fb61bff187e3 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-f694c7eee4c07b43d127a66c1d602afb6d309130b757dccd7620fc214cefee83)
- tls_tcp.tls_parameters.tls_config.custom_security

<a id="canonical-918ec4c30acce8ad1daeb546ef1b64e7412c93c19028d16963481ef971401a19"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-a5427b57f08d6617f9c06c9807f9bc906492a9389388f32f77dbf4c8a6c036d3"></a>

## Direct properties — tls_tcp.tls_parameters.tls_config.custom_security / fb61bff187e3 / 3

<a id="canonical-9856c3a6f10498084b7ef3a654ce1505892aca615d22031fabb99a3793d68697"></a>

<a id="canonical-b85b7549692a86f2489df60e75ff90d121bbe7ed38857a6f810aae09424e5bec"></a>

## cipher_suites property — tls_tcp.tls_parameters.tls_config.custom_security / fb61bff187e3 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-4dfdec101f6e26c20047907f8d69d0f38aa41a202fc2f039beefa79477264ff8"></a>

<a id="canonical-429a050017644b5d3d1a3128cc3b3745da1050d6f69d48fc5f40a0c5c582f249"></a>

## max_version property — tls_tcp.tls_parameters.tls_config.custom_security / fb61bff187e3 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-3e1c44b62f035948bd02a5e118dc280e676b036e43a8090bdabf7b8934ced109"></a>

<a id="canonical-44ae82b26d8f0d6feac19058bb37d93e8e40ddb84bb557598cb7d8d0a6dc48cb"></a>

## min_version property — tls_tcp.tls_parameters.tls_config.custom_security / fb61bff187e3 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-c56a63cf5cd8395a8db78c2fc174891ca0c00e57c81fbee2527663bb0e44aa1f"></a>

## Next pages — tls_tcp.tls_parameters.tls_config.custom_security / fb61bff187e3 / 7

- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-f694c7eee4c07b43d127a66c1d602afb6d309130b757dccd7620fc214cefee83)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-c5dd24d9edbe48690a9c2d374eeb0bf920c87b9ff0dd190fbc9122b3b31c7724"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca58a4093306c51dec9a0129f976707b6425bf4e0a78e8df6ef0a4a19c506c82"></a>

## tls_tcp.tls_parameters.tls_config.default_security — tls_tcp.tls_parameters.tls_config.default_security / 38e0ff0d14b7 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-f694c7eee4c07b43d127a66c1d602afb6d309130b757dccd7620fc214cefee83)
- tls_tcp.tls_parameters.tls_config.default_security

<a id="canonical-90405f4e788751e268de7fc91e700f3dc84b6ea8a1acde55d737e6e3f9c6d93e"></a>

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
default_security = {}
```

<a id="canonical-3ab07250f234810b45d07d4393d703a7a9a5e80dbaa343a3c6b9e7d660c664a0"></a>

## Direct properties — tls_tcp.tls_parameters.tls_config.default_security / 38e0ff0d14b7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e05ad77f75b59a34cf7fa1d7df8f02e9f66099a662d283e6379076e4e5542670"></a>

## Next pages — tls_tcp.tls_parameters.tls_config.default_security / 38e0ff0d14b7 / 4

- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-f694c7eee4c07b43d127a66c1d602afb6d309130b757dccd7620fc214cefee83)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-b00c2e128aabbfe4b14555292261de74450c979946ac740132fc8bc615f9a3b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dab51bf3fcbcdf9b872e204a01119af3677eb656dc01985623f3da417d08563b"></a>

## tls_tcp.tls_parameters.tls_config.low_security — tls_tcp.tls_parameters.tls_config.low_security / 6c6d8950d09e / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-f694c7eee4c07b43d127a66c1d602afb6d309130b757dccd7620fc214cefee83)
- tls_tcp.tls_parameters.tls_config.low_security

<a id="canonical-69fa96b5dc15cc29622e1bacc77a628f8af0ccf2acd65f48c7eec38299b8fd71"></a>

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
low_security = {}
```

<a id="canonical-9ba9b83c1dc18b3f4ea04b0152f8fc56bb84430f4cfafbb58ee6931e15f12890"></a>

## Direct properties — tls_tcp.tls_parameters.tls_config.low_security / 6c6d8950d09e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-49f99ceb80930087f2b28cba0d71cf421b3a4457f4f19be4cb7f151b2576fe90"></a>

## Next pages — tls_tcp.tls_parameters.tls_config.low_security / 6c6d8950d09e / 4

- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-f694c7eee4c07b43d127a66c1d602afb6d309130b757dccd7620fc214cefee83)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-1a46641a1babeccacddd1d56deed7f62492d6db20354c88d78a9fc1caa6699e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03ab381dceac731d68904084b5a4b5a746f7b74a80f952cde16c4fa9ebdeff64"></a>

## tls_tcp.tls_parameters.tls_config.medium_security — tls_tcp.tls_parameters.tls_config.medium_security / e6c927dec9f7 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-f694c7eee4c07b43d127a66c1d602afb6d309130b757dccd7620fc214cefee83)
- tls_tcp.tls_parameters.tls_config.medium_security

<a id="canonical-4834a793215722a3dfb4701b5a78a353ac1e18034db90ce141998c9839f80b82"></a>

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
medium_security = {}
```

<a id="canonical-41ca1f420402bdff99f13db207ae0f56f9d7b1df5d62ac700cef396d5c9c5e40"></a>

## Direct properties — tls_tcp.tls_parameters.tls_config.medium_security / e6c927dec9f7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-17ea82906ad090a8299a4aa41192c9d5442279ee8b291a04e37f9927aa42e1a3"></a>

## Next pages — tls_tcp.tls_parameters.tls_config.medium_security / e6c927dec9f7 / 4

- [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-f694c7eee4c07b43d127a66c1d602afb6d309130b757dccd7620fc214cefee83)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-e6c9f64ac3d283ff93179b5d7b14091b6266db6afffc6ee45c4e1f53d3445ec0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38053b935e684d1775133075c70a87b86988d61d7d16cea0c08042c21c2b3db0"></a>

## tls_tcp.tls_parameters.use_mtls — tls_tcp.tls_parameters.use_mtls / e60723db02f7 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- tls_tcp.tls_parameters.use_mtls

<a id="canonical-6ea72b670d9342d72553586b580b69d885e70506e030706a904fba9ef26b618c"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-308c22711c3a19189067084ad0a4e4e648839d2dc758e392060b049d5a019673"></a>

## Direct properties — tls_tcp.tls_parameters.use_mtls / e60723db02f7 / 3

<a id="canonical-e9547855e72e4233dd339b81e348b0d312044ca4864ed6bd63868b50e995e61b"></a>

<a id="canonical-fa9970a91fb9627e53b6dea744f6a726bcc3300bce9d9f8f0b8ab371b89199e3"></a>

## client_certificate_optional property — tls_tcp.tls_parameters.use_mtls / e60723db02f7 / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-2916e792c7ad3394668fbe5b60ad4048b93d36b298e3c37bb32d0a5ba03da9cf): complete subsection reference.

- [no_crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-1bf08a451cf987661bddfb96d5b043c5a90912502006d8a16693ff0505754715): complete subsection reference.

- [trusted_ca](resources--tcp_loadbalancer--reference--group-002.md#canonical-85bf3e93c547d076e4493d80a0c12cc3112acf1347abedc0828576896327007c): complete subsection reference.

<a id="canonical-8396a709121ca368355627f16b026c3c183d967492e49dd2704b6ab9b992273b"></a>

<a id="canonical-e0ccee76c182cb6afdc424a363542e496690a595b18db358abb6f086fdab0e94"></a>

## trusted_ca_url property — tls_tcp.tls_parameters.use_mtls / e60723db02f7 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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

- [xfcc_disabled](resources--tcp_loadbalancer--reference--group-002.md#canonical-e96327e0875fdfa33bfbafbd97319ddee147bbe8426edf50ccbf4f98cfb614c6): complete subsection reference.

- [xfcc_options](resources--tcp_loadbalancer--reference--group-002.md#canonical-00d830725aba0426b5b05f41a11e43702ca38d39c78f5fcffaf21f55a7b7a2d5): complete subsection reference.

<a id="canonical-ddfb6aac73169ab655cb49fd3d55950e202e6fbac0a9a45869d792e8154ab952"></a>

## Next pages — tls_tcp.tls_parameters.use_mtls / e60723db02f7 / 6

- [tls_tcp.tls_parameters.use_mtls.crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-2916e792c7ad3394668fbe5b60ad4048b93d36b298e3c37bb32d0a5ba03da9cf)
- [tls_tcp.tls_parameters.use_mtls.no_crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-1bf08a451cf987661bddfb96d5b043c5a90912502006d8a16693ff0505754715)
- [tls_tcp.tls_parameters.use_mtls.trusted_ca](resources--tcp_loadbalancer--reference--group-002.md#canonical-85bf3e93c547d076e4493d80a0c12cc3112acf1347abedc0828576896327007c)
- [tls_tcp.tls_parameters.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--reference--group-002.md#canonical-e96327e0875fdfa33bfbafbd97319ddee147bbe8426edf50ccbf4f98cfb614c6)
- [tls_tcp.tls_parameters.use_mtls.xfcc_options](resources--tcp_loadbalancer--reference--group-002.md#canonical-00d830725aba0426b5b05f41a11e43702ca38d39c78f5fcffaf21f55a7b7a2d5)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-2916e792c7ad3394668fbe5b60ad4048b93d36b298e3c37bb32d0a5ba03da9cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d9a9ae5ca66bf53268f38a8ee1d38b68a01bb1a092c4f8b57e210ba283ca649"></a>

## tls_tcp.tls_parameters.use_mtls.crl — tls_tcp.tls_parameters.use_mtls.crl / 8f521ab60406 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-e6c9f64ac3d283ff93179b5d7b14091b6266db6afffc6ee45c4e1f53d3445ec0)
- tls_tcp.tls_parameters.use_mtls.crl

<a id="canonical-ad0b02b934e081f9c7b41cbc2a2a12037fd89a881b0a378de92d2b13bb936b30"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-114eb17b85b518e60317837b0f85cdfe9d5128dfdb3081dc1f41270ee373b7c7"></a>

## Direct properties — tls_tcp.tls_parameters.use_mtls.crl / 8f521ab60406 / 3

<a id="canonical-2c1b79330f12b7c730816f5e396a9acfdec46325f28c7477ede3abecc03839eb"></a>

<a id="canonical-c39c955e41e7a470d121b1ada38bb757f942a66d91482017baf5edf96e18d717"></a>

## name property — tls_tcp.tls_parameters.use_mtls.crl / 8f521ab60406 / 4

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

<a id="canonical-17a7128c69ab97b1046d94e23026dce626e298196b543634f9d61580b789a6ab"></a>

<a id="canonical-10207dc454752aa6b0e44ddee93fe9c78e5362c56cf125b25af2319fc4d8b5cc"></a>

## namespace property — tls_tcp.tls_parameters.use_mtls.crl / 8f521ab60406 / 5

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

<a id="canonical-f94cbcba5cc93cb0eec999ade7e2134fee9d053391ccadbf935b36ae68e26680"></a>

<a id="canonical-50d48af69935abf68f932457a595e1d1f0fedeac86e166019040a4f69e0554fc"></a>

## tenant property — tls_tcp.tls_parameters.use_mtls.crl / 8f521ab60406 / 6

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

<a id="canonical-4491ab95f204efb36a8f7b4f0d4f1afb5de77c4ab6a128c873e9ad2f0fcde788"></a>

## Next pages — tls_tcp.tls_parameters.use_mtls.crl / 8f521ab60406 / 7

- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-e6c9f64ac3d283ff93179b5d7b14091b6266db6afffc6ee45c4e1f53d3445ec0)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-1bf08a451cf987661bddfb96d5b043c5a90912502006d8a16693ff0505754715"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-375d8147b2b7f1cde3e971ba39aab5ef531d0b61febcbdd671d26090b8144dbe"></a>

## tls_tcp.tls_parameters.use_mtls.no_crl — tls_tcp.tls_parameters.use_mtls.no_crl / 163f82cbcf28 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-e6c9f64ac3d283ff93179b5d7b14091b6266db6afffc6ee45c4e1f53d3445ec0)
- tls_tcp.tls_parameters.use_mtls.no_crl

<a id="canonical-dc67fa54b332672f119c4aeee658ee0ef43fff943cefbe0b5a02a0a3ac5aba81"></a>

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
no_crl = {}
```

<a id="canonical-85e9eca9fc772d540d1ee2870450c3b4e93e4508184725e3ed0ed0e07e33e312"></a>

## Direct properties — tls_tcp.tls_parameters.use_mtls.no_crl / 163f82cbcf28 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c93824046dbfafd32bb70841a767a301d2a1bae03f9dee9389a0d474057bb8b"></a>

## Next pages — tls_tcp.tls_parameters.use_mtls.no_crl / 163f82cbcf28 / 4

- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-e6c9f64ac3d283ff93179b5d7b14091b6266db6afffc6ee45c4e1f53d3445ec0)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-85bf3e93c547d076e4493d80a0c12cc3112acf1347abedc0828576896327007c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2692d07520b627211cecd1202e3ed0f82c8f2574c3e9be6770203cc834b69749"></a>

## tls_tcp.tls_parameters.use_mtls.trusted_ca — tls_tcp.tls_parameters.use_mtls.trusted_ca / 5df9a5fba590 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-e6c9f64ac3d283ff93179b5d7b14091b6266db6afffc6ee45c4e1f53d3445ec0)
- tls_tcp.tls_parameters.use_mtls.trusted_ca

<a id="canonical-2b89dc2634daa717de787e5f57118572562eb480d78e5036703ee740b3472935"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-a9c4074034104df7530a19544ce91a533d4d33d9ac697326494b4cc636148214"></a>

## Direct properties — tls_tcp.tls_parameters.use_mtls.trusted_ca / 5df9a5fba590 / 3

<a id="canonical-a698b662b72fce792de59525dbc4b07adca20b49e6dfad29f78a44516e366aed"></a>

<a id="canonical-3a2aac520d4255e693dd8ca84698275f577162703c721748458ca327ddad0bed"></a>

## name property — tls_tcp.tls_parameters.use_mtls.trusted_ca / 5df9a5fba590 / 4

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

<a id="canonical-ac731bacc11592d2b21532cb83666dcb5448c48c8337baaf180aa1b2683fc7d7"></a>

<a id="canonical-e6e8ddf2f87caf6d01f42f47713d0260976428966f37814a48b70512fe49c34d"></a>

## namespace property — tls_tcp.tls_parameters.use_mtls.trusted_ca / 5df9a5fba590 / 5

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

<a id="canonical-1c6cd973263adb0ded51346501d07a09ceaa1c0181ddec1ba00cde00e9af5e3f"></a>

<a id="canonical-f9d7999eb09bba45f66ea184df46e1aea85586b425042bd64694aaf97e586e4f"></a>

## tenant property — tls_tcp.tls_parameters.use_mtls.trusted_ca / 5df9a5fba590 / 6

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

<a id="canonical-5f109db037f6980c3d6ffde54235b4e1bdf70c26d8adb159fb299f527db5464f"></a>

## Next pages — tls_tcp.tls_parameters.use_mtls.trusted_ca / 5df9a5fba590 / 7

- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-e6c9f64ac3d283ff93179b5d7b14091b6266db6afffc6ee45c4e1f53d3445ec0)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-e96327e0875fdfa33bfbafbd97319ddee147bbe8426edf50ccbf4f98cfb614c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd2cb0f6e70191e5c0811410b036f5c6c325f9c72a4d3e29a2fbedbcf7d47ade"></a>

## tls_tcp.tls_parameters.use_mtls.xfcc_disabled — tls_tcp.tls_parameters.use_mtls.xfcc_disabled / 7f4254bcda5d / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-e6c9f64ac3d283ff93179b5d7b14091b6266db6afffc6ee45c4e1f53d3445ec0)
- tls_tcp.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-71c57554afe54c60e15deff1db2e063f9c33ce6cdc467771f83e4050b22ddf2b"></a>

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
xfcc_disabled = {}
```

<a id="canonical-4e2293a406bf8302e32faadcef9bfd5f8866259a105b1d51d017d77b9641b857"></a>

## Direct properties — tls_tcp.tls_parameters.use_mtls.xfcc_disabled / 7f4254bcda5d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ba049c2f0efe37352fd085721151d79da1d2be62bd9762dc6e6384209a313dda"></a>

## Next pages — tls_tcp.tls_parameters.use_mtls.xfcc_disabled / 7f4254bcda5d / 4

- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-e6c9f64ac3d283ff93179b5d7b14091b6266db6afffc6ee45c4e1f53d3445ec0)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-00d830725aba0426b5b05f41a11e43702ca38d39c78f5fcffaf21f55a7b7a2d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73f5c78b3a4f2978130c57cae536fa536e723f3adab9b9ccf61ef4a21898d416"></a>

## tls_tcp.tls_parameters.use_mtls.xfcc_options — tls_tcp.tls_parameters.use_mtls.xfcc_options / 05d99bec33a2 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-19570632ac35f7dc9b767a5c421915699bf6c6831d61bacbf894ef04542d745e)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-e6c9f64ac3d283ff93179b5d7b14091b6266db6afffc6ee45c4e1f53d3445ec0)
- tls_tcp.tls_parameters.use_mtls.xfcc_options

<a id="canonical-6e143add6dd848c238c2a03f88fe0226071145f139a1bde7fdf2a07519836a12"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-05d277ee2e4b9ffe579a12a81e6b54603d0bacd9abb623c7efbd284f73663516"></a>

## Direct properties — tls_tcp.tls_parameters.use_mtls.xfcc_options / 05d99bec33a2 / 3

<a id="canonical-6d113ed51f2aa65dcde5738e97f78ef10445efca76b5c0907d830ef4c1f8f7c9"></a>

<a id="canonical-bcc4c903a595f5f74394e4fc5df1ad26bb75a81b30ad28a42ef015a10d72ab45"></a>

## xfcc_header_elements property — tls_tcp.tls_parameters.use_mtls.xfcc_options / 05d99bec33a2 / 4

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-50d6c8ef92fe1f13a440d79dce1711c400606715cf381b3d94998a40f48c4a6a"></a>

## Next pages — tls_tcp.tls_parameters.use_mtls.xfcc_options / 05d99bec33a2 / 5

- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-e6c9f64ac3d283ff93179b5d7b14091b6266db6afffc6ee45c4e1f53d3445ec0)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31fab9cbbb14414a9f407e8576b1ed45e6318f9914285aad701c3c94cfa92e68"></a>

## tls_tcp_auto_cert — tls_tcp_auto_cert / 4e5bd02f0b1c / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- tls_tcp_auto_cert

<a id="canonical-490a6432751fa2d70f483cf952ae0e4fd5499dee0876f285d95600aa388da7e3"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting TLS over TCP proxy with automatic certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_tcp_auto_cert {
  # Configure direct properties listed below.
}
```

<a id="canonical-11ae6d96dbc860fb5c930bd992ffc81756c17ecae5f75ba0e75d2ec64cca47ee"></a>

## Direct properties — tls_tcp_auto_cert / 4e5bd02f0b1c / 3

- [no_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-becec3643210393334e20d4f4ff8f5040235ed81c5c5509aec9c9de254904d8f): complete subsection reference.

- [tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-8ae3c42a422bdd4db0ff92d6a72576a4c1e8287fc9a5b49323e5fa2423033731): complete subsection reference.

- [use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-9a7e24b8b186e72dc3e0ca17d39d4a51c54a20cc31cddb17aef00e32d5b4687c): complete subsection reference.

<a id="canonical-796e59b2513b4d25d15ba04f5b87f8bb9ff302d0e85404a9c096138904173d34"></a>

## Next pages — tls_tcp_auto_cert / 4e5bd02f0b1c / 4

- [tls_tcp_auto_cert.no_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-becec3643210393334e20d4f4ff8f5040235ed81c5c5509aec9c9de254904d8f)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-8ae3c42a422bdd4db0ff92d6a72576a4c1e8287fc9a5b49323e5fa2423033731)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-9a7e24b8b186e72dc3e0ca17d39d4a51c54a20cc31cddb17aef00e32d5b4687c)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-becec3643210393334e20d4f4ff8f5040235ed81c5c5509aec9c9de254904d8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
