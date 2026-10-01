---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-69673370f7ca96d46e0af2f58feb36d880749f56ccd34965f7bb9e94c34ab14f"></a>

## local_vrf.slo_config.static_routes.static_routes.node_interface.list — local_vrf.slo_config.static_routes.static_routes.node_interface.list / 907a0101815e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0b8aabb054a1f98dea1148aa06b8a93ba7c86b7255222f986fc6c9a38693f564)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-7b017141760f6e2414b83b28951b395ad2ba1c2be38d4e7ddd68e507166390a0)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-e5a2d87110f268db0b060ad49ac4400b9813118d308e9ae9c97941f9d717a41a)
- local_vrf.slo_config.static_routes.static_routes.node_interface.list

<a id="canonical-25d48e0e46c7f7cb5ab5a8e230fbd1a7e07587bd2704bbc026b6b77edc4333cb"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0cfd0303c2de38736fe6f945e6197312129630d23b5a4159d3b2fd55fd60df88"></a>

## Direct properties — local_vrf.slo_config.static_routes.static_routes.node_interface.list / 907a0101815e / 3

- [interface](resources--securemesh_site_v2--reference--group-012.md#canonical-b25778fa2e024f1ac7e39d5a1df2af6ce2df25eff79cbafe5282151ba5823ffe): complete subsection reference.

<a id="canonical-c9e00daa014fef302a8e490fff02829e5dbc9c98b97928635356cc57cbc39a20"></a>

<a id="canonical-fb83a6f85da73a7e8de3b43dcf5f580811cd6778be9d00ef2581bfdac75fd21f"></a>

## node property — local_vrf.slo_config.static_routes.static_routes.node_interface.list / 907a0101815e / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-b065340cbb9143a51e8b60f248118e0fb9a73d9edc9db094fc281a8699a0a3b7"></a>

## Next pages — local_vrf.slo_config.static_routes.static_routes.node_interface.list / 907a0101815e / 5

- [local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface](resources--securemesh_site_v2--reference--group-012.md#canonical-b25778fa2e024f1ac7e39d5a1df2af6ce2df25eff79cbafe5282151ba5823ffe)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-e5a2d87110f268db0b060ad49ac4400b9813118d308e9ae9c97941f9d717a41a)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-b25778fa2e024f1ac7e39d5a1df2af6ce2df25eff79cbafe5282151ba5823ffe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa40226dcc817ee592bc80b9f7d7f0428fea05633939a6ef3b35af50df755c4f"></a>

## local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / b934854851a5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [local_vrf.slo_config.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-0b8aabb054a1f98dea1148aa06b8a93ba7c86b7255222f986fc6c9a38693f564)
- [local_vrf.slo_config.static_routes.static_routes](resources--securemesh_site_v2--reference--group-011.md#canonical-7b017141760f6e2414b83b28951b395ad2ba1c2be38d4e7ddd68e507166390a0)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-e5a2d87110f268db0b060ad49ac4400b9813118d308e9ae9c97941f9d717a41a)
- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-011.md#canonical-08b51f16fa711b354f0067ee85777bc4870af5335cf76f1e4aebea417e3d5317)
- local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-34b6d057e14782aee7462010052284c2fb008e082dcc87c1e057b053c2b3ee7a"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-09a9e46fd79d2375141863cf244495bbc3896aeb6c0c258dca224682f2be5b7b"></a>

## Direct properties — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / b934854851a5 / 3

<a id="canonical-630223776e8d820b50be35082bb0527e97885d1efadd07d075d83a1b5f22ed58"></a>

<a id="canonical-ded11443b28c2371c2c9a31f9798e5c76a5e408cdeab31b8d536acca63e579f3"></a>

## kind property — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / b934854851a5 / 4

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

<a id="canonical-88dbe94fa3889b1d58a2eebef7ae1c57f2c1c7ab7b2b86f670ba583ed1415ddd"></a>

<a id="canonical-2e32e6f216a2ef76a81e351dbee349ab3578e5e828c7bd8618ac945fde451bdb"></a>

## name property — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / b934854851a5 / 5

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

<a id="canonical-3d1304944a91b89c493783e8b3267a0046a9f5d56c3f48b9d5cba6a225c23a86"></a>

<a id="canonical-983f8388e2ea64468b0fdd31cae1ee1ad2e570480e759fde6e02d5af2d2fa672"></a>

## namespace property — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / b934854851a5 / 6

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

<a id="canonical-b2aa29ac271f9073705715e70b08dce63ad64b1c20e23ad4efad34cc97d37027"></a>

<a id="canonical-d208bb77cda9dcfc0be44375ebac2a2711f158c94dfdde3dbab0783383d174d6"></a>

## tenant property — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / b934854851a5 / 7

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

<a id="canonical-aeba86485cb09c4e0945e5c8d42a1fbb481fdc5f3932bf955bd90b9bfd03348d"></a>

<a id="canonical-6c102b87748d4cd37a84f81ca31e0517a28a7bb863c2c8d0ae0cc9f71d888739"></a>

## uid property — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / b934854851a5 / 8

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

<a id="canonical-05b3accc630a8ed9f23a9dba0a1a79ef4253350a429c4cef862725e91fad6b23"></a>

## Next pages — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / b934854851a5 / 9

- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-011.md#canonical-08b51f16fa711b354f0067ee85777bc4870af5335cf76f1e4aebea417e3d5317)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9872008a55e1a5af865c6190b362df07766bd3871fa42d568d7c27e0ef9afcc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63a84487fb332c5c891291535f8f9bd154d7d458ccfe1700906e7d5f93bd3fd9"></a>

## local_vrf.slo_config.static_v6_routes — local_vrf.slo_config.static_v6_routes / 2a87b2a77a4a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- local_vrf.slo_config.static_v6_routes

<a id="canonical-1eee8721a06f6eaa1885f91abff5303bee40ab3f81ee8233e5458d3e752e0de4"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_v6_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-b709c693849d86bdae813c2a37149228cbe9befa097f693a2dcc1bd4c6d72132"></a>

## Direct properties — local_vrf.slo_config.static_v6_routes / 2a87b2a77a4a / 3

- [static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-d8d585be51694f4b06c2c04054fc7ee5f2879cad1e72214c9c8135bed533e622): complete subsection reference.

<a id="canonical-73c7fef799e75059bd830267d4c73a112a2f2dbe22c384c6cb16b53973f2d26c"></a>

## Next pages — local_vrf.slo_config.static_v6_routes / 2a87b2a77a4a / 4

- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-d8d585be51694f4b06c2c04054fc7ee5f2879cad1e72214c9c8135bed533e622)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d8d585be51694f4b06c2c04054fc7ee5f2879cad1e72214c9c8135bed533e622"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b923e91a0e55a5804076e8b288cad99e124daa73865489694fd9b52ee8542c90"></a>

## local_vrf.slo_config.static_v6_routes.static_routes — local_vrf.slo_config.static_v6_routes.static_routes / ffde78da448a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-9872008a55e1a5af865c6190b362df07766bd3871fa42d568d7c27e0ef9afcc8)
- local_vrf.slo_config.static_v6_routes.static_routes

<a id="canonical-5358d377c5a6538b565631e677dfb624a4e7128601f514e39257526b1caafb78"></a>

Type: `"object"`. list nested block, Optional.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_prefixes"),
  validators.RequiredOneOfListObjectAttributes("default_gateway",
    "ip_address",
    "node_interface"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "ip_address"),
  validators.ConflictingListObjectAttributes("default_gateway",
    "node_interface"),
  validators.ConflictingListObjectAttributes("ip_address",
    "node_interface")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-bc9719638146a9e85aa61b70a851ee22458d53a433a0c21200e7cd6a11c5627a"></a>

## Direct properties — local_vrf.slo_config.static_v6_routes.static_routes / ffde78da448a / 3

<a id="canonical-b84514422dad02a4457bce9664718be0ad9bda6574bf89c81afe92123436cb44"></a>

<a id="canonical-826f6387770c59680b71ce37fc5f20c4967e436832339076e7ff8775902b4951"></a>

## attrs property — local_vrf.slo_config.static_v6_routes.static_routes / ffde78da448a / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](resources--securemesh_site_v2--reference--group-012.md#canonical-e635fe756d0c7ca596362238ccb813e2b79b637fd0e4d6fc0ba7746a61868d91): complete subsection reference.

<a id="canonical-3f9a22122bcb36314258d8304db9a5b06a6324fdcdf10a3a78f5d5100ad9bf87"></a>

<a id="canonical-fe45342625f85ca43e51830a4162c34a2a26520c66b7c9cf896987dbc114c086"></a>

## ip_address property — local_vrf.slo_config.static_v6_routes.static_routes / ffde78da448a / 5

Type: `"string"`. Optional.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-ea2eda2cf440e0120cb14b57524e774d2aaa0f914cdb7252f1df153c78a7df84"></a>

<a id="canonical-c680c8a13f4444ede985d2bdda01fdc3e86267b2a5deaefc0160fb0b0d9b8e02"></a>

## ip_prefixes property — local_vrf.slo_config.static_v6_routes.static_routes / ffde78da448a / 6

Type: `["list", "string"]`. Optional.

List of IPv6 route prefixes that have common next hop and attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-00144d688e6220fa7671abf920675500b0b0252a24c1a5c40f6d0c31ec1978dc): complete subsection reference.

<a id="canonical-4ea4d58ec5d17237d9a6282e2b58aef02e022a7ebb2a475941801dad2862897a"></a>

## Next pages — local_vrf.slo_config.static_v6_routes.static_routes / ffde78da448a / 7

- [local_vrf.slo_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site_v2--reference--group-012.md#canonical-e635fe756d0c7ca596362238ccb813e2b79b637fd0e4d6fc0ba7746a61868d91)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-00144d688e6220fa7671abf920675500b0b0252a24c1a5c40f6d0c31ec1978dc)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-9872008a55e1a5af865c6190b362df07766bd3871fa42d568d7c27e0ef9afcc8)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-e635fe756d0c7ca596362238ccb813e2b79b637fd0e4d6fc0ba7746a61868d91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8443bb72d5647692d20726e276cd138c5a0d924739f20b71e8cabd8e2961592d"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.default_gateway — local_vrf.slo_config.static_v6_routes.static_routes.default_gateway / e5915719d4f3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-9872008a55e1a5af865c6190b362df07766bd3871fa42d568d7c27e0ef9afcc8)
- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-d8d585be51694f4b06c2c04054fc7ee5f2879cad1e72214c9c8135bed533e622)
- local_vrf.slo_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-7706d791f397fdbda33bac0fd855ab774bcce8b4ef7a03820173d6a974f1053e"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

<a id="canonical-09bf9a5adfd897f60401088aa49e09733300c6aa20c321bb6590835437ecf85e"></a>

## Direct properties — local_vrf.slo_config.static_v6_routes.static_routes.default_gateway / e5915719d4f3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c7c571e5367bdc98aa43514f6e180d9cb01e55ecb2da59c65d4df4c3d945e61a"></a>

## Next pages — local_vrf.slo_config.static_v6_routes.static_routes.default_gateway / e5915719d4f3 / 4

- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-d8d585be51694f4b06c2c04054fc7ee5f2879cad1e72214c9c8135bed533e622)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-00144d688e6220fa7671abf920675500b0b0252a24c1a5c40f6d0c31ec1978dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7aa7284c03be740d4601fa369b159bfe109816f615fb7f005f8306eee5688341"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.node_interface — local_vrf.slo_config.static_v6_routes.static_routes.node_interface / 3502c3fa7d24 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-9872008a55e1a5af865c6190b362df07766bd3871fa42d568d7c27e0ef9afcc8)
- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-d8d585be51694f4b06c2c04054fc7ee5f2879cad1e72214c9c8135bed533e622)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface

<a id="canonical-b7c55052ac5b0e27452d2a0cdf5c9698cbe98733c654bbbf412ec3ae1b053036"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-6027fab8c4282df868abc12cdafb53f8d0de240ccbe0705592e2eca3760ab7a3"></a>

## Direct properties — local_vrf.slo_config.static_v6_routes.static_routes.node_interface / 3502c3fa7d24 / 3

- [list](resources--securemesh_site_v2--reference--group-012.md#canonical-4b55725ff545a68ad636a70f384b981fcc3556a04a33c490ee61e1c0362e5a7c): complete subsection reference.

<a id="canonical-5011ce1ea881e6045be059a69b1bf6f64e68afa54fa10d88dc1270b785115876"></a>

## Next pages — local_vrf.slo_config.static_v6_routes.static_routes.node_interface / 3502c3fa7d24 / 4

- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-4b55725ff545a68ad636a70f384b981fcc3556a04a33c490ee61e1c0362e5a7c)
- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-d8d585be51694f4b06c2c04054fc7ee5f2879cad1e72214c9c8135bed533e622)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-4b55725ff545a68ad636a70f384b981fcc3556a04a33c490ee61e1c0362e5a7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e75da1345b1846c7bc8df3f8c8a3e0ea7c0a46c7ad886c8cdaced7f3b57f28f"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list / 29427449652c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-9872008a55e1a5af865c6190b362df07766bd3871fa42d568d7c27e0ef9afcc8)
- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-d8d585be51694f4b06c2c04054fc7ee5f2879cad1e72214c9c8135bed533e622)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-00144d688e6220fa7671abf920675500b0b0252a24c1a5c40f6d0c31ec1978dc)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-8cc91d2f93c28689d681a3850b53eed2e2e61ac3f49e02ffb14de2de1ada7e43"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-485c17dbaa8e854e61dda3d43cee569ebc0f32ab1cc8b8d2b5ad77ffddd350f9"></a>

## Direct properties — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list / 29427449652c / 3

- [interface](resources--securemesh_site_v2--reference--group-012.md#canonical-9aaaa010fd8bab22df331537324f787909fae4950be9784fe8c2a13651ee2bfc): complete subsection reference.

<a id="canonical-bafac934fa414978b54d746099cca22f4a7b59d89c34706549ccfa84c0eae808"></a>

<a id="canonical-2f55b7ebbbed74061870fe3e2ff9e01ecf83c0e7745d43dd1c2a6aea18fc6bec"></a>

## node property — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list / 29427449652c / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-4a21a09fe9975b0086847cabfc2032b7564594bea0d35347cb85d3f48f40e5a6"></a>

## Next pages — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list / 29427449652c / 5

- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site_v2--reference--group-012.md#canonical-9aaaa010fd8bab22df331537324f787909fae4950be9784fe8c2a13651ee2bfc)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-00144d688e6220fa7671abf920675500b0b0252a24c1a5c40f6d0c31ec1978dc)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9aaaa010fd8bab22df331537324f787909fae4950be9784fe8c2a13651ee2bfc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c1d849669835485081a82fec83dd02f13d2c3cde14c07fd00c0ae2ebdde9a45"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / c0dad2416b6d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [local_vrf](resources--securemesh_site_v2--reference--group-011.md#canonical-53331fb506d9637ed7bdd223f6561d958515be45ce89ef42efc31265406f7589)
- [local_vrf.slo_config](resources--securemesh_site_v2--reference--group-011.md#canonical-0ce8615e00ae72c60dba2329d7cd6aaba03093a7af767fcf474e3073dcfd45bd)
- [local_vrf.slo_config.static_v6_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-9872008a55e1a5af865c6190b362df07766bd3871fa42d568d7c27e0ef9afcc8)
- [local_vrf.slo_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-012.md#canonical-d8d585be51694f4b06c2c04054fc7ee5f2879cad1e72214c9c8135bed533e622)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-00144d688e6220fa7671abf920675500b0b0252a24c1a5c40f6d0c31ec1978dc)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-4b55725ff545a68ad636a70f384b981fcc3556a04a33c490ee61e1c0362e5a7c)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-390f6df85bd6f8bab105138d44ce6f1ad5d2835c2584a4747e16a66bf0c56cfe"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-847fe181d9f91eb41c90dc5d765671044b191263629eab3e414b7b1740a3f5ff"></a>

## Direct properties — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / c0dad2416b6d / 3

<a id="canonical-40b68d4827a015ab3ba6151aaed5213b2af2d23ebdbcf5cc4e99cff82e0bd54b"></a>

<a id="canonical-26b22ea02916d9815c4a450bd304def2ce555dbdbf9b0a07dc1be4db6106404f"></a>

## kind property — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / c0dad2416b6d / 4

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

<a id="canonical-5efaf6c3a685e4664349defa981ebf3b00f95587dd8063428236413cecfdc214"></a>

<a id="canonical-3124cb3aa5ffb823316dca3b99c16d8aefa0f5f770f8ee052516d6c4d21513f9"></a>

## name property — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / c0dad2416b6d / 5

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

<a id="canonical-3be040eaa2f63787185e5b10cf32625aa3906b89fab529223a7ea5ba4ce43fea"></a>

<a id="canonical-aba427570b733c616e1ce3bc51b541c5c99e5bf40c2c61356b974329f76958f0"></a>

## namespace property — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / c0dad2416b6d / 6

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

<a id="canonical-0abe0fb798ca0086608af34231a2fcedac4c1020d12ecd7fcae532552485d196"></a>

<a id="canonical-4765b5ea1d526f97760456220883de485a19ee8e971197aba525e266c738046c"></a>

## tenant property — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / c0dad2416b6d / 7

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

<a id="canonical-e196da562de5e180c336192850220f5c6f6285db76410d523a7301498cb666f0"></a>

<a id="canonical-0f94f7ebd6998d7a25ff8455aeab48f7b803e8d80318245789aadb4765ddc800"></a>

## uid property — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / c0dad2416b6d / 8

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

<a id="canonical-cf57614017245f0589bfdca14708023d06bf1d433e5af633ec18e2615018c0d2"></a>

## Next pages — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / c0dad2416b6d / 9

- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-012.md#canonical-4b55725ff545a68ad636a70f384b981fcc3556a04a33c490ee61e1c0362e5a7c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-993aba33f992fb4b41f8843fc0f570269f549fd2bd386d4eba1a525fcc444d4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59d0f0370a3b1d2881f9c54378aab07fb3997141802e094200fb064ffbedfb25"></a>

## log_receiver_with_net — log_receiver_with_net / a0994bd8144d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- log_receiver_with_net

<a id="canonical-6a7462d3b08b309fd69bb4b86c84eb9fbe6d5d063f78a829a317c5b8abc80b08"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: log\_receiver\_with\_net, logs\_streaming\_disabled\] Select log receiver for logs
streaming with network option.

Upstream description:

Select log receiver for logs streaming with network option.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("use_management_network",
    "use_slo_sli")}
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
  "x-ves-oneof-field-network_choice": "[\"use_management_network\",\"use_slo_sli\"]"
}
```

OneOf alternatives in this subsection:

- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-6a7462d3b08b309fd69bb4b86c84eb9fbe6d5d063f78a829a317c5b8abc80b08)
- [logs_streaming_disabled](resources--securemesh_site_v2--reference--group-012.md#canonical-15065f7fff40fee6948365b7b102e3823ddb68256f68246fbb9a9d887f2c4bbd)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver_with_net {
  # Configure direct properties listed below.
}
```

<a id="canonical-925a10d39bd1a212e09eac94747c1e2b8ec1559faf4c33b078097839e017627f"></a>

## Direct properties — log_receiver_with_net / a0994bd8144d / 3

- [log_receiver](resources--securemesh_site_v2--reference--group-012.md#canonical-581871dbca944f2fc21d18ada77043d9f91f168e5015789704cca8ad436c3a20): complete subsection reference.

- [use_management_network](resources--securemesh_site_v2--reference--group-012.md#canonical-40b3fd90b67c58e3b1dc0373e71e962de52c3e9888dd9630d99d31640b159568): complete subsection reference.

- [use_slo_sli](resources--securemesh_site_v2--reference--group-012.md#canonical-0d242cf9d67060d6457519650f5a9630f5637349e48ffd9960a67c5e4f97222c): complete subsection reference.

<a id="canonical-e58472177de8b7ecf4449b49ab115b8774179dd641948fd37473f911519291e2"></a>

## Next pages — log_receiver_with_net / a0994bd8144d / 4

- [log_receiver_with_net.log_receiver](resources--securemesh_site_v2--reference--group-012.md#canonical-581871dbca944f2fc21d18ada77043d9f91f168e5015789704cca8ad436c3a20)
- [log_receiver_with_net.use_management_network](resources--securemesh_site_v2--reference--group-012.md#canonical-40b3fd90b67c58e3b1dc0373e71e962de52c3e9888dd9630d99d31640b159568)
- [log_receiver_with_net.use_slo_sli](resources--securemesh_site_v2--reference--group-012.md#canonical-0d242cf9d67060d6457519650f5a9630f5637349e48ffd9960a67c5e4f97222c)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-581871dbca944f2fc21d18ada77043d9f91f168e5015789704cca8ad436c3a20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-310bc889cd4fd4f4a5b4bad770dca036c12885e30c8569e77aa5f951cbb0ae7e"></a>

## log_receiver_with_net.log_receiver — log_receiver_with_net.log_receiver / 44aa8538d77e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-993aba33f992fb4b41f8843fc0f570269f549fd2bd386d4eba1a525fcc444d4c)
- log_receiver_with_net.log_receiver

<a id="canonical-fc4e0c1784dd6ec8c5899273a0ae8bea4cad16a20acb54cecf3c56574d560759"></a>

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
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-bbe115c5ead1afb1e2b7f8e96dec87d5dfb011a54981c1fc7eaaa8b1bc2b4ffd"></a>

## Direct properties — log_receiver_with_net.log_receiver / 44aa8538d77e / 3

<a id="canonical-ed3b968780cdd6309ac0a94bcafdd8c5239adcaf41e642f2da5464c3bb5d41e4"></a>

<a id="canonical-0bd24554ddf0ebd80b7b646ea1401c4d6fca74700c1c17bcf34c706647ad14a6"></a>

## name property — log_receiver_with_net.log_receiver / 44aa8538d77e / 4

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

<a id="canonical-67630fb7bc74df2d1926dd41b9161119bc3dc42aedc2138d0db3cbb1a50b9e53"></a>

<a id="canonical-c6dfb9b72d4a990505b6ee20a8a18c4a30f7bafecfcffdc3a379e013969d3fce"></a>

## namespace property — log_receiver_with_net.log_receiver / 44aa8538d77e / 5

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

<a id="canonical-c6d117bb3bf5f18afbf2e9131b32cf4a7e1709f894d161e4da87258bd5d60b94"></a>

<a id="canonical-7f642d9f1ed809657631800627641c50750e9288646392be9b7b848b6aa461fe"></a>

## tenant property — log_receiver_with_net.log_receiver / 44aa8538d77e / 6

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

<a id="canonical-6f048d319dc9a15a1fb52112d4ab676e64315114b951de09974d58a80fd3edc4"></a>

## Next pages — log_receiver_with_net.log_receiver / 44aa8538d77e / 7

- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-993aba33f992fb4b41f8843fc0f570269f549fd2bd386d4eba1a525fcc444d4c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-40b3fd90b67c58e3b1dc0373e71e962de52c3e9888dd9630d99d31640b159568"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad733055b56fe7e01a364ab8a65e972ff24a32c75594a57031b067d6ac75c1a0"></a>

## log_receiver_with_net.use_management_network — log_receiver_with_net.use_management_network / ef3edbcb6265 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-993aba33f992fb4b41f8843fc0f570269f549fd2bd386d4eba1a525fcc444d4c)
- log_receiver_with_net.use_management_network

<a id="canonical-8fe169d7de9d0f3ba21843c33ad52614cba1b230b35c4b5d20baf0c19ffa1452"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use management network.

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
use_management_network = {}
```

<a id="canonical-aac352f74774c92d95418d6023f42b0305e9a20a89ae1fa7ba75b8c31f74a5aa"></a>

## Direct properties — log_receiver_with_net.use_management_network / ef3edbcb6265 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d1fc88b6506b5ea4cca7d7da189b1de8e198c9dd9a66550da08b2990433f02a3"></a>

## Next pages — log_receiver_with_net.use_management_network / ef3edbcb6265 / 4

- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-993aba33f992fb4b41f8843fc0f570269f549fd2bd386d4eba1a525fcc444d4c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0d242cf9d67060d6457519650f5a9630f5637349e48ffd9960a67c5e4f97222c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b56a0e5691f6d6f84998f821d88eb399ccefa27327daa18a3a6fd7aeacd3f32a"></a>

## log_receiver_with_net.use_slo_sli — log_receiver_with_net.use_slo_sli / 7b843d83ba26 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-993aba33f992fb4b41f8843fc0f570269f549fd2bd386d4eba1a525fcc444d4c)
- log_receiver_with_net.use_slo_sli

<a id="canonical-6b37a4d70da8f157dc60c8a73904acbf32f41639bb8725a2cf151b8a702818cf"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use slo sli.

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
use_slo_sli = {}
```

<a id="canonical-ae82fb915f8e6e0f1ab21f7b0664d2e7a11b2cb7f994bfe5eab02403124a2a6f"></a>

## Direct properties — log_receiver_with_net.use_slo_sli / 7b843d83ba26 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d3e964fcc5016872fb34db8c47ffe940c88fc9cb5a0f87b4c96cae225873c3ca"></a>

## Next pages — log_receiver_with_net.use_slo_sli / 7b843d83ba26 / 4

- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-012.md#canonical-993aba33f992fb4b41f8843fc0f570269f549fd2bd386d4eba1a525fcc444d4c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-a4bf722d3477ef61cdfac72a1f44fbf42bb3d727935964b96c0b646e651df9a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e5e325a313c6c38fbedb715486ac2913125498acc5676266ef81c4f1d93b358"></a>

## logs_streaming_disabled — logs_streaming_disabled / 568f7f15191a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- logs_streaming_disabled

<a id="canonical-15065f7fff40fee6948365b7b102e3823ddb68256f68246fbb9a9d887f2c4bbd"></a>

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
logs_streaming_disabled = {}
```

<a id="canonical-8ba9671a71097028657766eb48648bfae01809a34f577f3476bd8e84b9c837cc"></a>

## Direct properties — logs_streaming_disabled / 568f7f15191a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9d044c69d62d556587b9b36f8ffd69c3b40662ce7ceb106c061820d0b3ca9b39"></a>

## Next pages — logs_streaming_disabled / 568f7f15191a / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ed04394f5786d7ba076683bb8d8c3b295faccdf2dd9bbb6f7127a1214cd25f5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38e42e1da4e844544ae3ed4e96a7cdd4e5bc1c49529d7bbe585e8567a498ec5e"></a>

## no_forward_proxy — no_forward_proxy / 7ac2e022f907 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- no_forward_proxy

<a id="canonical-7a3380413c43f9eb28e6c7afb7dcb6fb8ccbe4e58ff268f0c626d193eca975cd"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no forward proxy.

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
no_forward_proxy = {}
```

<a id="canonical-2caf5fcc7e72cf4029557210e63021ccbd1478ef28d311e806a40e51693a685a"></a>

## Direct properties — no_forward_proxy / 7ac2e022f907 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-022cc0be4acb86e8d7406f9ea2043490dec03a1989842cce01fcc8c5194bcd98"></a>

## Next pages — no_forward_proxy / 7ac2e022f907 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-55f68f318cc4eb83391191144f1bc143f9a78ff7aa479bf7e19f3e94c5313c05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c60b4c97c9a07a425cccf933aa90110f94f4af1537d6ccfe14e528d18aa1c2f"></a>

## no_network_policy — no_network_policy / d27c954e22db / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- no_network_policy

<a id="canonical-a713a35898fc5fe73ae07fc3d82e4b872a9048db4ee12dd76f083a1a6e515e4d"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_network_policy = {}
```

<a id="canonical-972f919b69ff9a57e614c80ba460bddae40650995bded9a33d7154ae6182cc36"></a>

## Direct properties — no_network_policy / d27c954e22db / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bd021ac46d8460ae86978e314e6fbc3494b4d0c8918c8330e83a2fe6d04edb9f"></a>

## Next pages — no_network_policy / d27c954e22db / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-dbf7f694dafceef77a34b3fcca1e1e924ce856e24e5d48c4539c2cbbc5ae83fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a827f6d5188e61fac48931956d14356775eac30143043ec73a799eb07134b846"></a>

## no_proxy_bypass — no_proxy_bypass / f2ecc90455c0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- no_proxy_bypass

<a id="canonical-ca2c2b4ed1b36ea3644b35c8df651a17fe251ad10d3edd37a911ddea5d3f024a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no proxy bypass.

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
no_proxy_bypass = {}
```

<a id="canonical-502af8f045495e8743e19f29f03db725574a3f2fd046ba721f04fdebb391dbf7"></a>

## Direct properties — no_proxy_bypass / f2ecc90455c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65d0eebffa17cbae4f155e23c81e041f39c6294820e3a76c42ce3bbcdf7c81ce"></a>

## Next pages — no_proxy_bypass / f2ecc90455c0 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-fa1aa13eac43da651e8e0edc8322e51c0691628168db723973a8397b8861fb63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93f7000d141b9da54401561f4de8c11e8c4c20b2fcf3c740b70b48ad667da111"></a>

## no_s2s_connectivity_sli — no_s2s_connectivity_sli / b02d3581d79d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- no_s2s_connectivity_sli

<a id="canonical-3ccc2aaa670af29300d6a03d6ccbc04a0a4b22708dea6251171785cb26005b36"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no s2s connectivity sli.

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
no_s2s_connectivity_sli = {}
```

<a id="canonical-8b5737d28fdfc37d42bd301507a166e657f6467ca79db28421315b7003c1ac5b"></a>

## Direct properties — no_s2s_connectivity_sli / b02d3581d79d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6ed41c17607f1ba7f281cfa282b3ae7b5b41cd921334c866f0149ba2f4b6f545"></a>

## Next pages — no_s2s_connectivity_sli / b02d3581d79d / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-4e5a9b729b39300b66e4c6167141f27fda0a34d68aa8603d3fba3eb5e17d8722"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8b79dd047d24744a3ffb6b193015f10978c05d1a5c8be4ae6f91e22683fa633"></a>

## no_s2s_connectivity_slo — no_s2s_connectivity_slo / f37e8b9f4edc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- no_s2s_connectivity_slo

<a id="canonical-4ccc3da3d6e75369b4fbe0186fdd9d7795127b9fb25719afa5fe057674cb44d9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no s2s connectivity slo.

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
no_s2s_connectivity_slo = {}
```

<a id="canonical-229dbb9b23f496eca3ee9a5a80073670be85bfddbb961cc641719aa3cf5b6bd3"></a>

## Direct properties — no_s2s_connectivity_slo / f37e8b9f4edc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7c0106c900e49ad5f44ea58f0074b1fb6d0b728e909184d65c597333155dda10"></a>

## Next pages — no_s2s_connectivity_slo / f37e8b9f4edc / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0da6386dc27278430447d29293919790fdc22424187c7694816bbf234593776"></a>

## nutanix — nutanix / 6cb3950f6ed0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- nutanix

<a id="canonical-6419bacfe7e82130499848b85bc2e920be918823eaa75ad0c8981bf48ee12612"></a>

Type: `"object"`. single nested block, Optional.

Nutanix Provider Type. Nutanix Provider Type.

Upstream description:

Nutanix Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

Terraform syntax:

```terraform
nutanix {
  # Configure direct properties listed below.
}
```

<a id="canonical-029a7c340075e9de62f9fdd08115cf9d2d41971802f1c0306630cd899e0a62e6"></a>

## Direct properties — nutanix / 6cb3950f6ed0 / 3

- [not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0): complete subsection reference.

<a id="canonical-0ffcafcf26858590f41e834f078d977c591910d3252b8a8f1ae33545d1646b5d"></a>

## Next pages — nutanix / 6cb3950f6ed0 / 4

- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ff53370ae51a69e014ed201cdaf7c0c8859325fd2d41f9f2a32303b95633b23"></a>

## nutanix.not_managed — nutanix.not_managed / 94e30e878a02 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- nutanix.not_managed

<a id="canonical-fe4ebfdf7dceb6791181c2141b05ccda69c23ae41918d65aa87bcce68045516d"></a>

Type: `"object"`. single nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
not_managed {
  # Configure direct properties listed below.
}
```

<a id="canonical-2f56e83bad2d44247f4e67656b1ab241cd820d41989adcca270464eca66b4f4b"></a>

## Direct properties — nutanix.not_managed / 94e30e878a02 / 3

- [node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d): complete subsection reference.

<a id="canonical-c4b8bcab1c2783716223ecd7895bfcff626752fe29cc956ba1f9e5e19d1d329c"></a>

## Next pages — nutanix.not_managed / 94e30e878a02 / 4

- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68967fb720eb6dcc1d5213e5ff2c95ad2ba19eae40518d94025383bfcf975457"></a>

## nutanix.not_managed.node_list — nutanix.not_managed.node_list / 194e299942bc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- nutanix.not_managed.node_list

<a id="canonical-6437ae46ea50f0d02614b08b6b1210eecb41999703974f72382a87bd7bd0655a"></a>

Type: `"object"`. list nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
node_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-c811946d3dc3a9e77086d8b27f29582ca49bb27462149465730973fecc01181b"></a>

## Direct properties — nutanix.not_managed.node_list / 194e299942bc / 3

<a id="canonical-e1aa957c1ea93f0236b3f8c9d1683b96d12255a0190c3ed57e070cf1ef5595de"></a>

<a id="canonical-dd6fda8fd966a0a0ad0e4ba6bbd0157d9e753a47798e35c5033296cc23bac1ae"></a>

## hostname property — nutanix.not_managed.node_list / 194e299942bc / 4

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

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
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522): complete subsection reference.

<a id="canonical-d9ad382426cb4461ae4a41f76b36313a50c94075aefb8dbb4d3afac03aadec96"></a>

<a id="canonical-2f0bfb74fc25102eaaa66fa69487300bba82f4d6ea2c0787acc7ef63307ec14b"></a>

## public_ip property — nutanix.not_managed.node_list / 194e299942bc / 5

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-00c0c1487161b5b8353e65a4fa6dabd69f5ec5fdfc325a63b9fd9a75fd709b0f"></a>

<a id="canonical-34b722764cad64609d423aefaeb8c151b64b0aa791518286f2e53f3ac198b7af"></a>

## type property — nutanix.not_managed.node_list / 194e299942bc / 6

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("Control",
    "Worker"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-f5443786ca6b0f8fc84f1b51d8efccaae61054e5484499cabfbbc21f21571dcf"></a>

## Next pages — nutanix.not_managed.node_list / 194e299942bc / 7

- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e5c560398d1087bc346c227ec3284e8c6aeb2af990cb3f2f2cfafd9fa761fc5"></a>

## nutanix.not_managed.node_list.interface_list — nutanix.not_managed.node_list.interface_list / af0533f76d51 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- nutanix.not_managed.node_list.interface_list

<a id="canonical-b9ca75c300987a8749d96ea4c0561db614446923990b40fe9cd8182859c26556"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bond_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("bond_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingListObjectAttributes("ethernet_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingListObjectAttributes("no_ipv4_address",
    "static_ip"),
  validators.ConflictingListObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("site_to_site_connectivity_interface_disabled",
    "site_to_site_connectivity_interface_enabled")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3bce5d4649b29b57606aecd943eee6ae0daf3ca9f1a9c4c3c52f6b1b3d6e028a"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list / af0533f76d51 / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-ec0a9b34757895be550a9150b07a0a1038c641c2b9c32b675a30c5b2c643ecb6): complete subsection reference.

<a id="canonical-04d8470bbbcdcafb98c73b39bbd40ce2b60160b7931a69b0b27608cfe4d30712"></a>

<a id="canonical-a5c6a1025b53126355fe00ce724f8927970695b7c9dcb654bd74c5f85af8e71d"></a>

## description_spec property — nutanix.not_managed.node_list.interface_list / af0533f76d51 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-012.md#canonical-1ac4da81c604c38d68b1c1ed9731af0febb350dfbc51cd33536d184a42c35507): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-68cd3ae102c3e83c24a35967e85cf26e55db30814cafd730fac46b72bde5d0fb): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-aa3ab597c7a5276227d7ee995e3732ffc6f21607dc80315bb319df6b1d333ae2): complete subsection reference.

<a id="canonical-d1dc1e171c8ed87095c2d4882e5c7477ae1f365dca85023f5c5b098a48d7edfe"></a>

<a id="canonical-5d3d7cc81c0ea6205fd6ebdcd7d3fc0fe2c33a05be2e271e78ede9f4e7b92acb"></a>

## is_management property — nutanix.not_managed.node_list.interface_list / af0533f76d51 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-dc2acdac56a35e518dc2c4356b35501f49f8144c837da891e9b6cde5271aff6b"></a>

<a id="canonical-b920c736c25d8b9d19c0cad56cd64648365ceec2caea29622a104cc10fe3d992"></a>

## is_primary property — nutanix.not_managed.node_list.interface_list / af0533f76d51 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-4af8d130500660fc6bb460936865552a8ac618086384b3c87e772afea959eb6a"></a>

<a id="canonical-2e91916f18d84d46bb4de433baf80d50137cf0954c74abd1a23b50c07493f26c"></a>

## labels property — nutanix.not_managed.node_list.interface_list / af0533f76d51 / 7

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

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
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [monitor](resources--securemesh_site_v2--reference--group-013.md#canonical-dbe0e4014d7ddaed057318e15b6a1c4918a67b5d54ff25630e631cd9ee6cdcb1): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-013.md#canonical-dbe81e92d10d2d41c92664f865b8d01dc926e8166fa6cd3ef3a4a589dd286687): complete subsection reference.

<a id="canonical-1f59130a6de5ad77c4ba0edda5b7b426bb578f0379f46e9cc611b2a9cba38c28"></a>

<a id="canonical-8b55a758ef66ff3449fedd48c7e7f87eda3dce139739774a18b47c74f7d16f78"></a>

## mtu property — nutanix.not_managed.node_list.interface_list / af0533f76d51 / 8

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 8000},
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
    "maximum": 8000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-f3dfe15832505fecb0b8614a358bf2f6d1ef6a7d085f316db717ffad30084607"></a>

<a id="canonical-5c171aa8ac27e081b1e80684f9e51d60bc32a7239d402e1f1703fd75bb39e7f6"></a>

## name property — nutanix.not_managed.node_list.interface_list / af0533f76d51 / 9

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [network_option](resources--securemesh_site_v2--reference--group-013.md#canonical-8fbf11caa19bbf9ab5107aa52f6b1ba78a8cd2708283d1bbe13dc8069fb215dd): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-013.md#canonical-705d846b684d1e338f9a636332bb25e5f83e1b17fd8d54036ad2fc07831212f8): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-013.md#canonical-4319c5f112c43c7debcd908c53de6600d1c7153e64820754a924f13477c958d9): complete subsection reference.

<a id="canonical-0aab56e0f18d53b07e0d0dfac215657349537e0fdf688aab2ef1d3566116c9ed"></a>

<a id="canonical-4f0d9feb56c0498e2448196ce66b1346c5ce11ddc5129a3f5fde75649e31bc8a"></a>

## priority property — nutanix.not_managed.node_list.interface_list / af0533f76d51 / 10

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-013.md#canonical-6dc76137a7b805aac7b06ba91f657d74530a4f67d8eb9932faaf9bdc97b16c77): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-013.md#canonical-e7fb5c863a507e8e8bcb995531644d8de149994714f477f5c0532e377ff3e3ad): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-013.md#canonical-c0d65e4f48d95cf33fe518e218c66cb89f748e033225df8973de16352e9e0aa2): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-013.md#canonical-9e8888be541059bba45ce48e4e149ddeb88cffe65c0abda27086402b11d5073f): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-013.md#canonical-0128ee0584b515edd21ad3b984a508489ddf6fa2dbe9f07bf62e5d1850b21ae4): complete subsection reference.

<a id="canonical-a1abf4e6ae9724d952a40fc6ed504602ed81359c753cab9c506dd6283ba78079"></a>

## Next pages — nutanix.not_managed.node_list.interface_list / af0533f76d51 / 11

- [nutanix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-ec0a9b34757895be550a9150b07a0a1038c641c2b9c32b675a30c5b2c643ecb6)
- [nutanix.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-012.md#canonical-1ac4da81c604c38d68b1c1ed9731af0febb350dfbc51cd33536d184a42c35507)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8)
- [nutanix.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-68cd3ae102c3e83c24a35967e85cf26e55db30814cafd730fac46b72bde5d0fb)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-aa3ab597c7a5276227d7ee995e3732ffc6f21607dc80315bb319df6b1d333ae2)
- [nutanix.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-013.md#canonical-dbe0e4014d7ddaed057318e15b6a1c4918a67b5d54ff25630e631cd9ee6cdcb1)
- [nutanix.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-013.md#canonical-dbe81e92d10d2d41c92664f865b8d01dc926e8166fa6cd3ef3a4a589dd286687)
- [nutanix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-013.md#canonical-8fbf11caa19bbf9ab5107aa52f6b1ba78a8cd2708283d1bbe13dc8069fb215dd)
- [nutanix.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-013.md#canonical-705d846b684d1e338f9a636332bb25e5f83e1b17fd8d54036ad2fc07831212f8)
- [nutanix.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-013.md#canonical-4319c5f112c43c7debcd908c53de6600d1c7153e64820754a924f13477c958d9)
- [nutanix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-013.md#canonical-6dc76137a7b805aac7b06ba91f657d74530a4f67d8eb9932faaf9bdc97b16c77)
- [nutanix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-013.md#canonical-e7fb5c863a507e8e8bcb995531644d8de149994714f477f5c0532e377ff3e3ad)
- [nutanix.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-013.md#canonical-c0d65e4f48d95cf33fe518e218c66cb89f748e033225df8973de16352e9e0aa2)
- [nutanix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-013.md#canonical-9e8888be541059bba45ce48e4e149ddeb88cffe65c0abda27086402b11d5073f)
- [nutanix.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-013.md#canonical-0128ee0584b515edd21ad3b984a508489ddf6fa2dbe9f07bf62e5d1850b21ae4)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ec0a9b34757895be550a9150b07a0a1038c641c2b9c32b675a30c5b2c643ecb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71ace85902d2988029ad77ad8478dfec30f69f997c1f7fe248555c2e93173545"></a>

## nutanix.not_managed.node_list.interface_list.bond_interface — nutanix.not_managed.node_list.interface_list.bond_interface / ce8d9deb3e15 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- nutanix.not_managed.node_list.interface_list.bond_interface

<a id="canonical-d86da59a3775158591cb339ccebc746a99b60dfb77428027924302cebfcdabf9"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingObjectAttributes("active_backup",
    "lacp")}
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
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

Terraform syntax:

```terraform
bond_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-26836c95d4001c89aecfde27cae13a4c33149f309670524dc6243fd4b59a75e6"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.bond_interface / ce8d9deb3e15 / 3

- [active_backup](resources--securemesh_site_v2--reference--group-012.md#canonical-d5897743276297575bef0e9ad9b12fea147d4d472c933d346831660bc61ff5bc): complete subsection reference.

<a id="canonical-d23cd3f324c1bc4ba02cba73211c31c933adc7c7c215c44ba0268b51ebc38d70"></a>

<a id="canonical-bc973d012c7c8950eee0d2f827832a7694c4084450f1b9b8f81dd741105b93f8"></a>

## devices property — nutanix.not_managed.node_list.interface_list.bond_interface / ce8d9deb3e15 / 4

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](resources--securemesh_site_v2--reference--group-012.md#canonical-a27d10194f02c48f8743aedf874524d3268c72c1707c136c4c766c959cebdf58): complete subsection reference.

<a id="canonical-b44e1a9681994bbf49c830998c26f90919a0c85baa726acbcef41d1da57475cb"></a>

<a id="canonical-f67165980c104e7b8fdbba8da96b9fa3c2a46b1ac3f45f0c553d223c19fd66d5"></a>

## link_polling_interval property — nutanix.not_managed.node_list.interface_list.bond_interface / ce8d9deb3e15 / 5

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(500, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-aefa2304e07d6fb3d1f1e682b16fc3fd8d3aa65dd77a6fb1b91f07127417e098"></a>

<a id="canonical-da394e7b876f6f7c429e6eaf89689240e27b703eb489dc8504eb54fb1aa4a76c"></a>

## link_up_delay property — nutanix.not_managed.node_list.interface_list.bond_interface / ce8d9deb3e15 / 6

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 1000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
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
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="canonical-247e26353a8b41bdbf421a2c3d850fcfabe1b73a5bef482c829003c6b504c31a"></a>

<a id="canonical-8f1346fb566e62dfee923fe596e3509bb537f9200f20c5d692276224dd7b0c4a"></a>

## name property — nutanix.not_managed.node_list.interface_list.bond_interface / ce8d9deb3e15 / 7

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-7aa91aa74e27866fe08b2898d1e7b46d3279d77c25304e27eaad661f7ad3ab2b"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.bond_interface / ce8d9deb3e15 / 8

- [nutanix.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-012.md#canonical-d5897743276297575bef0e9ad9b12fea147d4d472c933d346831660bc61ff5bc)
- [nutanix.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-012.md#canonical-a27d10194f02c48f8743aedf874524d3268c72c1707c136c4c766c959cebdf58)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d5897743276297575bef0e9ad9b12fea147d4d472c933d346831660bc61ff5bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e9277223aaf984da01b64a45b19eae8a9e3491386ae8658a6582efb9c007ee7"></a>

## nutanix.not_managed.node_list.interface_list.bond_interface.active_backup — nutanix.not_managed.node_list.interface_list.bond_interface.active_backup / 51bf4cd39682 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-ec0a9b34757895be550a9150b07a0a1038c641c2b9c32b675a30c5b2c643ecb6)
- nutanix.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-9bd0aa24b59f9e0c82346eb7d3aede078ccce7250b53d2bf0e92eb21018be19d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for active backup.

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
active_backup = {}
```

<a id="canonical-1a60c669476f69f6bed22bff2a38e66afad910ac17a94f98facbf6bcb10432e0"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.bond_interface.active_backup / 51bf4cd39682 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-23af2f30261edaee83f56d487cb7479015efd38222be543f1d9fd7e3985576c4"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.bond_interface.active_backup / 51bf4cd39682 / 4

- [nutanix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-ec0a9b34757895be550a9150b07a0a1038c641c2b9c32b675a30c5b2c643ecb6)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-a27d10194f02c48f8743aedf874524d3268c72c1707c136c4c766c959cebdf58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe62611ff2d68183df6ed28ad304dcff6e736d4d2abcf31ef53b6749a0282c0b"></a>

## nutanix.not_managed.node_list.interface_list.bond_interface.lacp — nutanix.not_managed.node_list.interface_list.bond_interface.lacp / 55b91f697700 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-ec0a9b34757895be550a9150b07a0a1038c641c2b9c32b675a30c5b2c643ecb6)
- nutanix.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-0c93f48675f7c8e619f19cb0e57dacc3a0f72adf4ce01eb4da5c2228715bee5d"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate")}
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
lacp {
  # Configure direct properties listed below.
}
```

<a id="canonical-a6b281101ce73668f5c4d776d37781490758467184a089703049fd5688b48ab9"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.bond_interface.lacp / 55b91f697700 / 3

<a id="canonical-8727a3c73cae3fdd7dad8d5a43ed04a0b901443a31d7bad3e74c98add61d29a5"></a>

<a id="canonical-7a4740738c76b9acd75377f16307d8e6f26ea04e7e3a0929f317e87745c3e147"></a>

## rate property — nutanix.not_managed.node_list.interface_list.bond_interface.lacp / 55b91f697700 / 4

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
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
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-a4e48344de88ac5549aa0afcd3780f64a281435f8e442634b39f72b4ff24a3a4"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.bond_interface.lacp / 55b91f697700 / 5

- [nutanix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-012.md#canonical-ec0a9b34757895be550a9150b07a0a1038c641c2b9c32b675a30c5b2c643ecb6)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1ac4da81c604c38d68b1c1ed9731af0febb350dfbc51cd33536d184a42c35507"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2af0d6b93e78a456b29452e6e186749d05b12f1093a3a65b662483525168b7e"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_client — nutanix.not_managed.node_list.interface_list.dhcp_client / e95c263d388b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- nutanix.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-25682d1812c765f4a0e9792665dcc32da68ed612493e2edd33b606dbd949639c"></a>

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
dhcp_client = {}
```

<a id="canonical-c6eb4fa26e3f3c01fdeb16d7424d383f2f380efa537a24587640015b6b777dc8"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.dhcp_client / e95c263d388b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-01b0f7760fa242f4e4abf3495b481438c8b3dc70a1cb8d85af2195d2302c9374"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.dhcp_client / e95c263d388b / 4

- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-926691d7bbc15686d42a24aa9122f4f8dbd7769753baebdbf6c49a969ed58622"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server — nutanix.not_managed.node_list.interface_list.dhcp_server / 299e5d1c954b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- nutanix.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-1c0d6c482887f89213d9ec7cc14fee62d326d8c2dd65552ecd7655e1d4eb6fd7"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-03181df0f2bd8795b22e61bacf0118d16267e0e346f83b4a644488f8894512e8"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.dhcp_server / 299e5d1c954b / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-012.md#canonical-1916ddafa7d2995c6503aa1eff14587ebee434936b72f38d3ca90f8008eeab2c): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-012.md#canonical-122f6af00ff15c7b17b5749eee9922dd0f02d05936f8f5fc70846dc5ee2320bd): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-34d7ff8d51a6a3e355aa2681f9173a619572c6882c439e44e973572b51d242d1): complete subsection reference.

<a id="canonical-39f1e47442e7baab088e07e54b5ccce886a4ae9300900e7c34b1fdc9a588ea2d"></a>

<a id="canonical-452cca71f4c25d346a3f38cbc7e15f83554272be4485c54b8a81a5bfc111f69a"></a>

## dhcp_option82_tag property — nutanix.not_managed.node_list.interface_list.dhcp_server / 299e5d1c954b / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-a83cb8a29af0d1683e74fa75094c9f872f8343f95543c33cbb4099bda990487f"></a>

<a id="canonical-a45d6ed2f3864eca0221459f59e3f9c9d2b8e3ebfa452ebeb850b88a4b323c88"></a>

## fixed_ip_map property — nutanix.not_managed.node_list.interface_list.dhcp_server / 299e5d1c954b / 5

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](resources--securemesh_site_v2--reference--group-012.md#canonical-70b2805598b1f70335e50fa0288258f4b9307af85bb8f7c41fda98606f424461): complete subsection reference.

<a id="canonical-054db354c69633dff942fc87cde891acba51c02c56b12ecfaa224386c2b674e9"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.dhcp_server / 299e5d1c954b / 6

- [nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-012.md#canonical-1916ddafa7d2995c6503aa1eff14587ebee434936b72f38d3ca90f8008eeab2c)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-012.md#canonical-122f6af00ff15c7b17b5749eee9922dd0f02d05936f8f5fc70846dc5ee2320bd)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-34d7ff8d51a6a3e355aa2681f9173a619572c6882c439e44e973572b51d242d1)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-012.md#canonical-70b2805598b1f70335e50fa0288258f4b9307af85bb8f7c41fda98606f424461)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1916ddafa7d2995c6503aa1eff14587ebee434936b72f38d3ca90f8008eeab2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-470e7335ef98f20dc5520d40acaf46e2dd73e0fa99805f7cfa9b5167cde355e6"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / ce156980f2df / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8)
- nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-f64536877427691e396a17b2082d36e966e8dc2b8f0a3e4fc2a28a8b8dfaf8cc"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

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
automatic_from_end = {}
```

<a id="canonical-57c26c9970ea77ce2f584f1f74238794485b99fafdec5cb8afca70a625742a93"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / ce156980f2df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5c4abf3845321bd3d0ff6fbc643a406a5b370311586bf204362b6d523d4cbb7d"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / ce156980f2df / 4

- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-122f6af00ff15c7b17b5749eee9922dd0f02d05936f8f5fc70846dc5ee2320bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c3a3b44b68e7f3beb186b1ce6d23b55eea461042dfb7bc7ecb2b09498ebac9a"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 4b460b529863 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8)
- nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-7d44c9bc898de12519e4109c0ebe24be0d3477fbf482263e79ca12cba5f39b14"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from start.

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
automatic_from_start = {}
```

<a id="canonical-f34f221fbb521b2f74f921ae7be03589b8f25f67b42901b66b155e7db9d6e23e"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 4b460b529863 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6367a43e3b113286bfac3443788ae063bd3f301894e05636358c979148cf14f4"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 4b460b529863 / 4

- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-34d7ff8d51a6a3e355aa2681f9173a619572c6882c439e44e973572b51d242d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54a29ec4c4d57790dbae743879028d67b93fe914366e57fdbeb9bb87a12f1884"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 99748bcf3414 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-c2e68d99d5ed0bcb33f36ea763bbc41db20992863601def0259831d432cf06dc"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dgw_address",
    "first_address"),
  validators.ConflictingListObjectAttributes("dgw_address",
    "last_address"),
  validators.ConflictingListObjectAttributes("dns_address",
    "same_as_dgw"),
  validators.ConflictingListObjectAttributes("first_address",
    "last_address")}
```

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-f1408feb1412c0634950694fb6c6bc5ec889a21e1dce9e24ebc4cc1100e85923"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 99748bcf3414 / 3

<a id="canonical-c175e63ab9bf322af45f3b4130d7d23b1c3bc240f06639ad21f22fac88fa6dc6"></a>

<a id="canonical-a4bc024cc877d7316ffcbc9d7949443bddc57247045890ac1991f66a808db614"></a>

## dgw_address property — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 99748bcf3414 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

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

<a id="canonical-aea1a75734df8302c7933f1fcbb35f8d66673188bd44ebc5fa1b5b6d7c8c5228"></a>

<a id="canonical-e65131ca915094bb556223e60a8129c097413c26ca6cc81af380de1bfee449db"></a>

## dns_address property — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 99748bcf3414 / 5

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

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

- [first_address](resources--securemesh_site_v2--reference--group-012.md#canonical-f570600b840046abe88738d585584bbea8103ac7b0713af94b2e3251d83e3b3d): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-012.md#canonical-af72ca76d8a3e469053569151e8f679cd7cab7edc533e82e4630279a40e419a2): complete subsection reference.

<a id="canonical-2cef494ea014492072a01c0a75257b8758e0c46562f39c366253d5718e3be62c"></a>

<a id="canonical-bced59a3cb9c278a238be387a63a56a4e59c985023994dc561bc9fb0a084ff74"></a>

## network_prefix property — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 99748bcf3414 / 6

Type: `"string"`. Optional.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Upstream description:

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-39b873212c08e4b6e92bb5e02310c4f9c2c89e1244963fc10058e0d915768994"></a>

<a id="canonical-c0c1d8ab0a00dcb490681aa2395870bb175d95e18e5ea86c0c62556ab0ce974b"></a>

## pool_settings property — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 99748bcf3414 / 7

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](resources--securemesh_site_v2--reference--group-012.md#canonical-8f2f1f933d87f107b3cb715611d250554f31b3f5b5df70120eed9a8e8a488c0f): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-012.md#canonical-baa968d626086fb3dc602b225d915806b1130d0ea6e9caedbf110b2523dfcfe1): complete subsection reference.

<a id="canonical-a03809935c4237e15d9faa80eb648ea7cdae74668da6dc12be687701d7d21516"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 99748bcf3414 / 8

- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-012.md#canonical-f570600b840046abe88738d585584bbea8103ac7b0713af94b2e3251d83e3b3d)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-012.md#canonical-af72ca76d8a3e469053569151e8f679cd7cab7edc533e82e4630279a40e419a2)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-012.md#canonical-8f2f1f933d87f107b3cb715611d250554f31b3f5b5df70120eed9a8e8a488c0f)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-012.md#canonical-baa968d626086fb3dc602b225d915806b1130d0ea6e9caedbf110b2523dfcfe1)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f570600b840046abe88738d585584bbea8103ac7b0713af94b2e3251d83e3b3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ade33e479e01b0b5d0af325e9906dafa55ddd4220ad3167926b490cc4729a80"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_add / abd72a1a7787 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-34d7ff8d51a6a3e355aa2681f9173a619572c6882c439e44e973572b51d242d1)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-831f2e174ad9440edc6cf49b800305eeda066b325412fe7e1f985db3bbd20821"></a>

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
first_address = {}
```

<a id="canonical-d78d1629e5195d368e91d7322f717ad5603f31d815836c70672d67e82b43522e"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_add / abd72a1a7787 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-29c1e6c036f2cb23da32566a85145d7e0a651a12d37c9a665cd257ac39168ed0"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_add / abd72a1a7787 / 4

- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-34d7ff8d51a6a3e355aa2681f9173a619572c6882c439e44e973572b51d242d1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-af72ca76d8a3e469053569151e8f679cd7cab7edc533e82e4630279a40e419a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0daa5ec013b561572bcb7bda628d3714348fb1d6c3879fa7eb1dae5a7fea6c2c"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addr / ed0213ef2728 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-34d7ff8d51a6a3e355aa2681f9173a619572c6882c439e44e973572b51d242d1)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-ec7369f1671b788c2f6db193e38bc7a5772677bfcef9d13fe86e3746309daa30"></a>

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
last_address = {}
```

<a id="canonical-b2c3e5b4f67472c433a1aae6e5cd316679fd8850029a2c81813205ade8dc0868"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addr / ed0213ef2728 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-19339863c6cee9868aa97208f10e35796675b0198d1c7c6bd32faa0bcfc4b6cd"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addr / ed0213ef2728 / 4

- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-34d7ff8d51a6a3e355aa2681f9173a619572c6882c439e44e973572b51d242d1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-8f2f1f933d87f107b3cb715611d250554f31b3f5b5df70120eed9a8e8a488c0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34c8dc717b6e8529b64e3707c40ab8a7cb367bb0e76b9ab50a2f952fd07d2dba"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 476c1833f611 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-34d7ff8d51a6a3e355aa2681f9173a619572c6882c439e44e973572b51d242d1)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-e61a4ddc1d4d6fb69aa8ce0cd4e3ff7f0c540d981cbbdd6b7bb4b3b327c7c894"></a>

Type: `"object"`. list nested block, Optional.

List of non overlapping IP address ranges.

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
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-43ef97ca905928cb408a22ccce48e40fb73d961560674b5a3b2416fb05096624"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 476c1833f611 / 3

<a id="canonical-703170c75154c15c0ec649b30b81e8c43fd1c0bc0c928b460cc32eafef6f8d99"></a>

<a id="canonical-ada33c9b1b9ae731ced8e798976d61a59925964dbdb790eae880bee694c4caa2"></a>

## end_ip property — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 476c1833f611 / 4

Type: `"string"`. Optional.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

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

<a id="canonical-e02e299a8e379d53e90509caf4184d462331c911a92082cb14548a84dccfac1b"></a>

<a id="canonical-17d752d9ca41f45531a8eccec46847f1602864619544fd720582936b7c120451"></a>

## exclude property — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 476c1833f611 / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-eb7adcaea4bac08ffe5bb900fa6a0e463a63b3c0e5f417df4626fa7316b0f41b"></a>

<a id="canonical-e1c51550ecc45257fe8bb08eaab564a586625b5318091daa67c69584f4b70fd8"></a>

## start_ip property — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 476c1833f611 / 6

Type: `"string"`. Optional.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

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

<a id="canonical-a668484348efb4b31bfda066c17ccbac32098350fa535b8beec0a118ffa8dc18"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 476c1833f611 / 7

- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-34d7ff8d51a6a3e355aa2681f9173a619572c6882c439e44e973572b51d242d1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-baa968d626086fb3dc602b225d915806b1130d0ea6e9caedbf110b2523dfcfe1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8dd67863ef56e684bf1c940d3357051dc277c79d8fad2d08471b5bef4eb772c6"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_d / 43e2748ab306 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-34d7ff8d51a6a3e355aa2681f9173a619572c6882c439e44e973572b51d242d1)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-d2cb81c82c7a45a9fbe4e5e9cffcd1dc9b816606d9b263be69dce059021db8a7"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for same as dgw.

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
same_as_dgw = {}
```

<a id="canonical-ef3eb69b5e91de277db6c7e737c40fd52fde08681d9a1e74736e05ae8e32e3a6"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_d / 43e2748ab306 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-798286abcf9d0d8e73e485e606ffc1acc78d47006dac9aa8a5009bc7c1db6e7a"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_d / 43e2748ab306 / 4

- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-34d7ff8d51a6a3e355aa2681f9173a619572c6882c439e44e973572b51d242d1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-70b2805598b1f70335e50fa0288258f4b9307af85bb8f7c41fda98606f424461"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ccb58f2cecf0ae362f1b437e492107470166a6281d36f160862b985d1b632d36"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 91115a0357cf / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8)
- nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-9a3f408f00fa7684d0c02ffc9b5f1c3b9331bda9270650907c8779efc7573ec7"></a>

Type: `"object"`. single nested block, Optional.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

Upstream description:

Specify static IPv4 addresses per node.

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
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-a92cf6e5c866e0476963956700a432456fb72f0cbe214d3e99e7a7be15a8b4d6"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 91115a0357cf / 3

<a id="canonical-a04bc349593a14d1011d8c42716629968f0eaf2b8e3e6eef3dc3c8c92f40427e"></a>

<a id="canonical-e3d653b9f671849a9f421e505044904263f1d5d058b55468e53c8e6ebbe7d206"></a>

## interface_ip_map property — nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 91115a0357cf / 4

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

<a id="canonical-992e37188d5a4efc08141ae6ed22d9c6499178ffa033924380142ce58ef26abf"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 91115a0357cf / 5

- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-54f2a5fbcb1c54873aeb77368475022609f7b57e6cf35d2494fd80936605d2a8)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-68cd3ae102c3e83c24a35967e85cf26e55db30814cafd730fac46b72bde5d0fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-649192322c174c57fcab24cd7e57bae70076fa4b585aab9006968c2362c1bd6e"></a>

## nutanix.not_managed.node_list.interface_list.ethernet_interface — nutanix.not_managed.node_list.interface_list.ethernet_interface / b2ec01eade7b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- nutanix.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-2dbb3c98dc90194bfacfb01988091fb74b60be3b544d8f23c6330b918148812f"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mac")}
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
ethernet_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3f6de6c1f2d156b9c64666b5dcc2e1d9e5fa370f66d6c5e72b81637213c93731"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.ethernet_interface / b2ec01eade7b / 3

<a id="canonical-a29533cad6865cc79a4c708cb94772e614c285c4e6b46213aee6b6f72fa19dc1"></a>

<a id="canonical-dfcd8331f3fc55c443f73792dc74c53f362b74e9f98025232ea92bfdc496fb7f"></a>

## device property — nutanix.not_managed.node_list.interface_list.ethernet_interface / b2ec01eade7b / 4

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-8cd91a588a0e9089ce6e83233a3f6891ae97b7e8789fec96d6602e19771c8b99"></a>

<a id="canonical-cc2f7e874d1a5855a5c3e13a9cfe97f077c0b0f40477caf409536b9bc0bc3a7b"></a>

## mac property — nutanix.not_managed.node_list.interface_list.ethernet_interface / b2ec01eade7b / 5

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.MACValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```

<a id="canonical-85f1a266242a2f699f4aeae7fa2e0185a3f6c68b7c89c89e3b27d56f555138a5"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.ethernet_interface / b2ec01eade7b / 6

- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-aa3ab597c7a5276227d7ee995e3732ffc6f21607dc80315bb319df6b1d333ae2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca33670ee4c298cde57234098101d3ef36d756af5cf514f708a3c2e4ef21757a"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config — nutanix.not_managed.node_list.interface_list.ipv6_auto_config / 61f1f754df87 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-75c0497d434cb7bbdd0b988d2269b7e5d98f5921353470b1a9652ce9eae4e734"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("host",
    "router")}
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
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

Terraform syntax:

```terraform
ipv6_auto_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-c7cdd2caaf422a350b6482d71657b4f55c6f80d791cf0ed0fc1b3cb036380d67"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.ipv6_auto_config / 61f1f754df87 / 3

- [host](resources--securemesh_site_v2--reference--group-012.md#canonical-b563aeec9d06a69645bb0cb1854ae0217f777528c2e1f60b4b39ec9c8ef2d4fd): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-012.md#canonical-972cde8c89c13ee666e9dd3a268d0261d7aaf37cbfaff2a46cd8fa5e8bdea9a4): complete subsection reference.

<a id="canonical-465d1a9cb667dca416452e2001457b345bd74edc08771be7100c4b2b4a473114"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.ipv6_auto_config / 61f1f754df87 / 4

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-012.md#canonical-b563aeec9d06a69645bb0cb1854ae0217f777528c2e1f60b4b39ec9c8ef2d4fd)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-972cde8c89c13ee666e9dd3a268d0261d7aaf37cbfaff2a46cd8fa5e8bdea9a4)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-b563aeec9d06a69645bb0cb1854ae0217f777528c2e1f60b4b39ec9c8ef2d4fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58c1086051d70500d566c1d0dee73c3a29000e08439e4440a82545fdb2faeb91"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config.host — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.host / 064ae252eca4 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-aa3ab597c7a5276227d7ee995e3732ffc6f21607dc80315bb319df6b1d333ae2)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-dde6311a825523ffccff3d9cb7948eb2d3ccbf545cbf5823a5e3ad5aad980327"></a>

Type: `["object", {}]`. Optional.

Hostname or IP address of the target server.

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
host = {}
```

<a id="canonical-7d7e7d9556acc2c586f42c8d687cb139196b55abd23be7108fd3476af414fbe6"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.host / 064ae252eca4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-71c0a441595cddc24a53787260fb3dacaf676ce59de34c50c0fc3b887686dcaf"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.host / 064ae252eca4 / 4

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-aa3ab597c7a5276227d7ee995e3732ffc6f21607dc80315bb319df6b1d333ae2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-972cde8c89c13ee666e9dd3a268d0261d7aaf37cbfaff2a46cd8fa5e8bdea9a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6ea5d63ad22a94735e4882f5e21f8b1fdbb6d1b8e8d211d3b9ffe5521aa8f60"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router / 9ef323946732 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-aa3ab597c7a5276227d7ee995e3732ffc6f21607dc80315bb319df6b1d333ae2)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-75963b96d1452547a949e1261bb63bfba2b18aed3ea2eb4ef0e4b9c3e60eaf6c"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("network_prefix",
    "stateful")}
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
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

Terraform syntax:

```terraform
router {
  # Configure direct properties listed below.
}
```

<a id="canonical-6d3705ecd679d52a66abb060aea458b5bd7c1682267435ab47a123519ede08ab"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router / 9ef323946732 / 3

- [dns_config](resources--securemesh_site_v2--reference--group-012.md#canonical-21802b7b36a5cd72b9935053673f6c8763eb42bc56924760a3b0c09accd8cf2c): complete subsection reference.

<a id="canonical-90d2faed080a118b1afc497d8b7b8d96d97da778738e138f907ec7eb406e6aa8"></a>

<a id="canonical-c57be9b218fb67342966ef8c4462636514b23c93a1ada61ed0ec96998a0653cc"></a>

## network_prefix property — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router / 9ef323946732 / 4

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](resources--securemesh_site_v2--reference--group-012.md#canonical-868ec13f83f053d7ca18a3eb6ae292f62fbf3838ad89b5f531ea02fe5f89ae08): complete subsection reference.

<a id="canonical-80eee90267ad0a60a6181c47b2174fcb8c72b16f8383af0de59f903a8bfefecc"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router / 9ef323946732 / 5

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-012.md#canonical-21802b7b36a5cd72b9935053673f6c8763eb42bc56924760a3b0c09accd8cf2c)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-012.md#canonical-868ec13f83f053d7ca18a3eb6ae292f62fbf3838ad89b5f531ea02fe5f89ae08)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-aa3ab597c7a5276227d7ee995e3732ffc6f21607dc80315bb319df6b1d333ae2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-21802b7b36a5cd72b9935053673f6c8763eb42bc56924760a3b0c09accd8cf2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80d8077437bb234a689bb20c231d534ddcdcd2a883111dd04b86c94894a5e455"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / a994ecc94e6e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-aa3ab597c7a5276227d7ee995e3732ffc6f21607dc80315bb319df6b1d333ae2)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-972cde8c89c13ee666e9dd3a268d0261d7aaf37cbfaff2a46cd8fa5e8bdea9a4)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-f390f83833f944bb57d6249a000269305d17de91bd75331d9711a417129a5ccf"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_list",
    "local_dns")}
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
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

Terraform syntax:

```terraform
dns_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-a1f2454ef146def990b694d62d25fe8b2c297b1f63baa7a0ef6a44cc6aa013d2"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / a994ecc94e6e / 3

- [configured_list](resources--securemesh_site_v2--reference--group-012.md#canonical-ab559bc2cd6f13546683d75ba021fa9b2035a39db87e7b80a240ce9a71de967c): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-012.md#canonical-dd3c1367f84acb5df64b2c1f7b3da87ce0ef182541a0067d9a3abf27217036cc): complete subsection reference.

<a id="canonical-c78250bf0a476b3b76e69f7fa3a2974534e5933abbe49df696a1ed9b66945c50"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / a994ecc94e6e / 4

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-012.md#canonical-ab559bc2cd6f13546683d75ba021fa9b2035a39db87e7b80a240ce9a71de967c)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-012.md#canonical-dd3c1367f84acb5df64b2c1f7b3da87ce0ef182541a0067d9a3abf27217036cc)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-972cde8c89c13ee666e9dd3a268d0261d7aaf37cbfaff2a46cd8fa5e8bdea9a4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ab559bc2cd6f13546683d75ba021fa9b2035a39db87e7b80a240ce9a71de967c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-db3856f568a12d20e52b078b53c9dc0dbcf477bb05f04b31abaaed6155d29112"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 8301ed6d72f6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-aa3ab597c7a5276227d7ee995e3732ffc6f21607dc80315bb319df6b1d333ae2)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-972cde8c89c13ee666e9dd3a268d0261d7aaf37cbfaff2a46cd8fa5e8bdea9a4)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-012.md#canonical-21802b7b36a5cd72b9935053673f6c8763eb42bc56924760a3b0c09accd8cf2c)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-d74ab0d1ebfec8b2295578ac241a919087c73250c923ef6feff862868e67fd97"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_list")}
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
configured_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-10833eb0440de58340e2c8e5a03ee62b5b87457468d9a7b913141651fa206ec4"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 8301ed6d72f6 / 3

<a id="canonical-776c20dc7f212435fdfedca50f9e23e046603ec8761ad7d36b8e57e8152e0d87"></a>

<a id="canonical-51a335d4f6764c31fe20ddffbddcf2c980ed2646bd2d9a1025d6476306f6a724"></a>

## dns_list property — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 8301ed6d72f6 / 4

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8cc37a3d6825a85ace791984cb1c22af46e848f5223e1c0b0b55f585ef7615cd"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 8301ed6d72f6 / 5

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-012.md#canonical-21802b7b36a5cd72b9935053673f6c8763eb42bc56924760a3b0c09accd8cf2c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-dd3c1367f84acb5df64b2c1f7b3da87ce0ef182541a0067d9a3abf27217036cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5d4321e69fe8ff0c70923aab7cb4ab2113ecb0de2b6d04e9c73212321411fe1"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / c10e20dd6863 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-aa3ab597c7a5276227d7ee995e3732ffc6f21607dc80315bb319df6b1d333ae2)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-972cde8c89c13ee666e9dd3a268d0261d7aaf37cbfaff2a46cd8fa5e8bdea9a4)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-012.md#canonical-21802b7b36a5cd72b9935053673f6c8763eb42bc56924760a3b0c09accd8cf2c)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-af30ed682df15aef444b6a85fd5f7fd900815ab6d826bdc27a9fa94c1aa91171"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_address",
    "first_address"),
  validators.ConflictingObjectAttributes("configured_address",
    "last_address"),
  validators.ConflictingObjectAttributes("first_address",
    "last_address")}
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
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

Terraform syntax:

```terraform
local_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-624885ad6034de8c483c406e2920007b85005dfebff47dd58f174aee2df916ac"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / c10e20dd6863 / 3

<a id="canonical-1753a1c07e0277a031786b50baf8ad8082c77617fd4cd119070679fe9b4acd1b"></a>

<a id="canonical-906057b309b75fc57a6a7d175cb9ce9f5a3d30d82a33f74e821b8fe6781cdd52"></a>

## configured_address property — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / c10e20dd6863 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

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

- [first_address](resources--securemesh_site_v2--reference--group-012.md#canonical-0d6a5b3d5fb2c293055c28bfabb7776e5cd928362febb7b177bfd47a0e6c09e3): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-012.md#canonical-86b64c75a79a5f6ec527cfa12168a949a15a2843675f495e31c54241117c05a8): complete subsection reference.

<a id="canonical-e9b8e3158e158ff3d294100c48ec8c687d5722e47a919de7db4fd714a0e4f6c6"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / c10e20dd6863 / 5

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-012.md#canonical-0d6a5b3d5fb2c293055c28bfabb7776e5cd928362febb7b177bfd47a0e6c09e3)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-012.md#canonical-86b64c75a79a5f6ec527cfa12168a949a15a2843675f495e31c54241117c05a8)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-012.md#canonical-21802b7b36a5cd72b9935053673f6c8763eb42bc56924760a3b0c09accd8cf2c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0d6a5b3d5fb2c293055c28bfabb7776e5cd928362febb7b177bfd47a0e6c09e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-143c361a858d591514b701557acf484856e987f9ca1c00c63b5a368df7572d74"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 7adaf28b40c0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-aa3ab597c7a5276227d7ee995e3732ffc6f21607dc80315bb319df6b1d333ae2)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-972cde8c89c13ee666e9dd3a268d0261d7aaf37cbfaff2a46cd8fa5e8bdea9a4)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-012.md#canonical-21802b7b36a5cd72b9935053673f6c8763eb42bc56924760a3b0c09accd8cf2c)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-012.md#canonical-dd3c1367f84acb5df64b2c1f7b3da87ce0ef182541a0067d9a3abf27217036cc)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-0e152cb5397041149bfb568110445af066dd8b464c7b3f51ff1fcb58b39ed1be"></a>

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
first_address = {}
```

<a id="canonical-7345e2405a693e9f1f9f152ef2c5e9aafc2e61a10790eb2f6679ce97564e85dd"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 7adaf28b40c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-47bef413922e19adbbf38d49fe0aa183a84c5c5a54b2bd5951cdb0b46d19930d"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 7adaf28b40c0 / 4

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-012.md#canonical-dd3c1367f84acb5df64b2c1f7b3da87ce0ef182541a0067d9a3abf27217036cc)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-86b64c75a79a5f6ec527cfa12168a949a15a2843675f495e31c54241117c05a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-850545d505d282b7668b890e4c74a711012586fe38a9abc6c38416a1b4aa79d3"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 66656abc046d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-aa3ab597c7a5276227d7ee995e3732ffc6f21607dc80315bb319df6b1d333ae2)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-972cde8c89c13ee666e9dd3a268d0261d7aaf37cbfaff2a46cd8fa5e8bdea9a4)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-012.md#canonical-21802b7b36a5cd72b9935053673f6c8763eb42bc56924760a3b0c09accd8cf2c)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-012.md#canonical-dd3c1367f84acb5df64b2c1f7b3da87ce0ef182541a0067d9a3abf27217036cc)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-39978a388aa8fa3c223b7273f3f55a3fa00cb8725b17c684b18f7cced00f3573"></a>

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
last_address = {}
```

<a id="canonical-8ca90a8058275ed111d5cf775489c68110d2a9db43cd21c1c2f3d99bae2b5478"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 66656abc046d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-53d3aa2de9343d743fd280030f874b92547c1c4667d27f8946d989a329851d5e"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 66656abc046d / 4

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-012.md#canonical-dd3c1367f84acb5df64b2c1f7b3da87ce0ef182541a0067d9a3abf27217036cc)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-868ec13f83f053d7ca18a3eb6ae292f62fbf3838ad89b5f531ea02fe5f89ae08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0fb81cea8dfb344c96023d630144e972c456151ecf5d1c8ba23253b87852789"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 3ee9a67426fe / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [nutanix](resources--securemesh_site_v2--reference--group-012.md#canonical-f7f10140791a954ba385da0509ddf368b0c00d74712036f2468bfdb9f0de4d39)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-012.md#canonical-e858945ad47398b0c0ba13088c7045f0dc2e5302f14b23a59db4cc9332552fc0)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-012.md#canonical-6bb28bffc9664e54e13993d09bcc787d1232c1cffdb95ad3e1fbd9ae9411b56d)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-012.md#canonical-de1d567a5d76040cd5c80da41282e8beeb0fffab7f9f990687581277bf00d522)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-aa3ab597c7a5276227d7ee995e3732ffc6f21607dc80315bb319df6b1d333ae2)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-972cde8c89c13ee666e9dd3a268d0261d7aaf37cbfaff2a46cd8fa5e8bdea9a4)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-2ad36db5fb53d2ab1276446e44ce86199090f302e3bbf15dd9948eb44b1fc74c"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
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
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
stateful {
  # Configure direct properties listed below.
}
```

<a id="canonical-710e0ed85387398cb6d67ca554acc1084eeb18c52b1f377f6fd84b04aec82dc7"></a>

## Direct properties — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 3ee9a67426fe / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-012.md#canonical-68544c15ee2131d6e3d2ff170c2996e9f710b5007e8e29267d9be7c37a7bffc4): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-013.md#canonical-897c74b1a095c784e39634a1011f2f9ab1ce0c8b58e9f8f5ad5aa17082df423b): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-013.md#canonical-eac2d3a95b13fed71cc4972afbd8fd2c3c6fa7fa25cc03a14380807ddc99f09d): complete subsection reference.

<a id="canonical-624a2178dfa0a6a18c9ad20ccaf370cccdb445f705a49bc4d5d6ad48f122f107"></a>

<a id="canonical-c0a449f0b1c9e04791e211870627b4b2ba11462d9610c2e5a35d7aa9034fbb0d"></a>

## fixed_ip_map property — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 3ee9a67426fe / 4

Type: `["map", "string"]`. Optional.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

- [interface_ip_map](resources--securemesh_site_v2--reference--group-013.md#canonical-ab39e78bbf501c37ee28845f6848206e40ccb3f955a8edb1ad51ed9e41e4d490): complete subsection reference.

<a id="canonical-c8d5b477ba057ab7b117144d7e9e3dd1e4dcdf86b0db50c76185fcfda17e19b1"></a>

## Next pages — nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 3ee9a67426fe / 5

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-012.md#canonical-68544c15ee2131d6e3d2ff170c2996e9f710b5007e8e29267d9be7c37a7bffc4)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-013.md#canonical-897c74b1a095c784e39634a1011f2f9ab1ce0c8b58e9f8f5ad5aa17082df423b)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-013.md#canonical-eac2d3a95b13fed71cc4972afbd8fd2c3c6fa7fa25cc03a14380807ddc99f09d)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-013.md#canonical-ab39e78bbf501c37ee28845f6848206e40ccb3f955a8edb1ad51ed9e41e4d490)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-972cde8c89c13ee666e9dd3a268d0261d7aaf37cbfaff2a46cd8fa5e8bdea9a4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-68544c15ee2131d6e3d2ff170c2996e9f710b5007e8e29267d9be7c37a7bffc4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
