---
page_title: "xcsh_securemesh_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site reference."
---

# xcsh_securemesh_site reference

<a id="canonical-5246751e2bbccaca80b2abb97e654050172f7e6402dc51ac8cf9e97542671263"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / b351f0a01eed / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-4e5ad1a125eae31e32a88e6ebab4b027d2ca3599af65c962f4bf11181c237cf6)
- [custom_network_config.slo_config.static_routes](resources--securemesh_site--reference--group-003.md#canonical-c45e3b69dba14b22975bf0b48d8863dd20ea55965c10cde8076397181a182e0a)
- [custom_network_config.slo_config.static_routes.static_routes](resources--securemesh_site--reference--group-003.md#canonical-a3f8137ccff9cfcb7187d723b350936a97089754378321e7e49622b5a8f39033)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](resources--securemesh_site--reference--group-003.md#canonical-25a5de252ee700cd3b10432a7d3c4bec47a4f973e5a954fae1c4668ae9ff3c42)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-003.md#canonical-eb6293e8559380a62630af1fce6f8ddd5e05203478bbeb48609d9530e62730ed)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-ef5e7f72db09d7d0d9ea6c1ec43f9c46a2a80a52a150dff2fb1f313d67b2b7fa"></a>

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

<a id="canonical-9e59bd2dd9d55ce1bbf45ef4982fcc1ebd9736b4b5cfe53ca3dfefd941909694"></a>

## Direct properties — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / b351f0a01eed / 3

<a id="canonical-b5b935065e7fcb0bc938168dc97bcfbf221a7da885b9a2f5b2a7696411076d9e"></a>

<a id="canonical-0d99926bdff57a9a0e07c43abe299dd5f8b4c090678fb55cfabdfab2732d78a8"></a>

## kind property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / b351f0a01eed / 4

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

<a id="canonical-1e56305fc1ce0bfe476fef1d9ea4311dddc860787a74d33983db13d69721d3c4"></a>

<a id="canonical-87ecec1d4053da6d301cc4c30dcc54243f2c7b0da8845dbb0f846a08d4ff30e9"></a>

## name property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / b351f0a01eed / 5

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

<a id="canonical-ebcd2551414e597ac198baca04da5bde0cacfb025e204e7a0e466815c440c828"></a>

<a id="canonical-3b856f574110b9a5576eb1135c0a25bb4817f5d9c5507e79b5cae9cdc0ef8406"></a>

## namespace property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / b351f0a01eed / 6

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

<a id="canonical-02f57965552f0b4675eb16a66bb41f58a290bc8230a83ffceee93d4447a3f086"></a>

<a id="canonical-0478ca1ef0c2bb46bc8c2dc63f8529a9379f1658a3711cd3611de7969ba9e4a3"></a>

## tenant property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / b351f0a01eed / 7

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

<a id="canonical-f3bf5dc3ba5d8b9d572519cd8501fe647393f91925b6a61a8f6b63aa4145b760"></a>

<a id="canonical-e8ff78e4d92ff248bd74ca7b22960fec9c0b2ee2b8cc6dd9dce337b67018399b"></a>

## uid property — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / b351f0a01eed / 8

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

<a id="canonical-aa0d424b509b10f241b200eac102933874b15180a0972633687b82eded133147"></a>

## Next pages — custom_network_config.slo_config.static_routes.static_routes.node_interface.list / b351f0a01eed / 9

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-003.md#canonical-eb6293e8559380a62630af1fce6f8ddd5e05203478bbeb48609d9530e62730ed)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-2ef3ef525fcf5c2e84040f0a59757519078457887cc7fa4cebd33a88e6b127ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12fa35bf995279474fe5f2ca0e681293a752387a68159af5cde982d75ced5666"></a>

## custom_network_config.slo_config.static_v6_routes — custom_network_config.slo_config.static_v6_routes / bf63d18d45fc / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-4e5ad1a125eae31e32a88e6ebab4b027d2ca3599af65c962f4bf11181c237cf6)
- custom_network_config.slo_config.static_v6_routes

<a id="canonical-a904b1fdd92e2dd12cdf49133e6e80ca5a8a671eee65dfcee29e7fdb725bce9d"></a>

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

<a id="canonical-773b5748e5e85424794ea79cf3979000c87bbb43dce69796f6c76b159ed88fdb"></a>

## Direct properties — custom_network_config.slo_config.static_v6_routes / bf63d18d45fc / 3

- [static_routes](resources--securemesh_site--reference--group-004.md#canonical-c78defbebc9a33045380b30673711c8de31a5bede00cf1da0c301ee08485f939): complete subsection reference.

<a id="canonical-89fe8bb713905aa34e3185bbb488cd09fb367e937480ce54342aa7e8561077cc"></a>

## Next pages — custom_network_config.slo_config.static_v6_routes / bf63d18d45fc / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-c78defbebc9a33045380b30673711c8de31a5bede00cf1da0c301ee08485f939)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-4e5ad1a125eae31e32a88e6ebab4b027d2ca3599af65c962f4bf11181c237cf6)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-c78defbebc9a33045380b30673711c8de31a5bede00cf1da0c301ee08485f939"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a93372eea459b498a7b0236f49cbc68f23cf48cd921e4cd65ed12bd4f5dd1610"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes — custom_network_config.slo_config.static_v6_routes.static_routes / 6592d1582375 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-4e5ad1a125eae31e32a88e6ebab4b027d2ca3599af65c962f4bf11181c237cf6)
- [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-2ef3ef525fcf5c2e84040f0a59757519078457887cc7fa4cebd33a88e6b127ff)
- custom_network_config.slo_config.static_v6_routes.static_routes

<a id="canonical-3e0246b764b4fb20420a912221315c3c856adb7de9c57605484c6b075dbebd0d"></a>

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

<a id="canonical-e2450bd6c200d9f5d4a9918e8735d7f878650114d814765aa990058af4087569"></a>

## Direct properties — custom_network_config.slo_config.static_v6_routes.static_routes / 6592d1582375 / 3

<a id="canonical-e4abed99cfab9e315924ad1f8f7a5499c0e4380e7132a39cefc9090a004290dd"></a>

<a id="canonical-eb1b8c575bf217c7ad8c23d1a9492549a6cc62d0af3bee6cf0b3aa1dc6cd55e9"></a>

## attrs property — custom_network_config.slo_config.static_v6_routes.static_routes / 6592d1582375 / 4

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

- [default_gateway](resources--securemesh_site--reference--group-004.md#canonical-61f8d15ee720def8cd4b28aff9339b94d90b750298993217f404c485c9812cb6): complete subsection reference.

<a id="canonical-acb0fc34bee658e5046e0d5e2650f1018bfeeb54685f4cfad87f277d55571584"></a>

<a id="canonical-996890a6605bf9fbf5e1651066c1a13d1aa4bb85775cde2016204b97a8e856e9"></a>

## ip_address property — custom_network_config.slo_config.static_v6_routes.static_routes / 6592d1582375 / 5

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

<a id="canonical-b4ef4c954332b6bd9881597337ac4b74d15e4ada8dc3ba3dccc6accfb9c49242"></a>

<a id="canonical-3b2604aa6f34dafaee19fe8eb398a697c85d2b1d20b0f38c4092166ea45249bb"></a>

## ip_prefixes property — custom_network_config.slo_config.static_v6_routes.static_routes / 6592d1582375 / 6

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

- [node_interface](resources--securemesh_site--reference--group-004.md#canonical-4c64b9c6056c33fc1a9d7f3ec2a749c7087c7c1a16befba223e3d210748db384): complete subsection reference.

<a id="canonical-60c1bdacf5fdaa8050f8e8d1d2a4e526dcb88844a176527f1ddeb1bee167ca49"></a>

## Next pages — custom_network_config.slo_config.static_v6_routes.static_routes / 6592d1582375 / 7

- [custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site--reference--group-004.md#canonical-61f8d15ee720def8cd4b28aff9339b94d90b750298993217f404c485c9812cb6)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-4c64b9c6056c33fc1a9d7f3ec2a749c7087c7c1a16befba223e3d210748db384)
- [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-2ef3ef525fcf5c2e84040f0a59757519078457887cc7fa4cebd33a88e6b127ff)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-61f8d15ee720def8cd4b28aff9339b94d90b750298993217f404c485c9812cb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-516808fb305ee328fba805fef1832e18ff8ee6106b37960c6f9d34d73d221f3e"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway — custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway / 0bbdf06e885d / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-4e5ad1a125eae31e32a88e6ebab4b027d2ca3599af65c962f4bf11181c237cf6)
- [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-2ef3ef525fcf5c2e84040f0a59757519078457887cc7fa4cebd33a88e6b127ff)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-c78defbebc9a33045380b30673711c8de31a5bede00cf1da0c301ee08485f939)
- custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-70bfff4ba4cd4edba9622c865525b4e6c2b506d7ba8ff70e7f86898a9dbbc29f"></a>

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

<a id="canonical-05c7f02dd9f886da26907e19eba682cfad9a92dbd22c521f0f1c8d6210a43718"></a>

## Direct properties — custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway / 0bbdf06e885d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2da8768419207e0afba1b5ff0b64b774b25756ad23927ad5cd73937a77c47fe1"></a>

## Next pages — custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway / 0bbdf06e885d / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-c78defbebc9a33045380b30673711c8de31a5bede00cf1da0c301ee08485f939)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-4c64b9c6056c33fc1a9d7f3ec2a749c7087c7c1a16befba223e3d210748db384"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52bd934ad605f7cddf67eb0be4eb751484d5bfe8bf6cadca2cf2f493b1c16238"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface / e79b09afca16 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-4e5ad1a125eae31e32a88e6ebab4b027d2ca3599af65c962f4bf11181c237cf6)
- [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-2ef3ef525fcf5c2e84040f0a59757519078457887cc7fa4cebd33a88e6b127ff)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-c78defbebc9a33045380b30673711c8de31a5bede00cf1da0c301ee08485f939)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface

<a id="canonical-311cdb0729c05aeedd98201e612d0df0357aac2725d5aace8ffe118e7f39b088"></a>

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

<a id="canonical-2e247ea4107308cb5d4a054959f6f0419a0d4bbe8fd41cb611048bea46c99e21"></a>

## Direct properties — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface / e79b09afca16 / 3

- [list](resources--securemesh_site--reference--group-004.md#canonical-b0560a606c1fce2fd395d55c5b26293e9bbc7dafda0263b2c350e4c2ba8ae3f7): complete subsection reference.

<a id="canonical-6f770611fe84e2d3fb1af1c705567eac421704a38f90a6e63723a133041f2648"></a>

## Next pages — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface / e79b09afca16 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-004.md#canonical-b0560a606c1fce2fd395d55c5b26293e9bbc7dafda0263b2c350e4c2ba8ae3f7)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-c78defbebc9a33045380b30673711c8de31a5bede00cf1da0c301ee08485f939)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-b0560a606c1fce2fd395d55c5b26293e9bbc7dafda0263b2c350e4c2ba8ae3f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44c6b38d6987664bc678d112ec41c13aff7793f5b8dcfc83ed5ead1d9ce89770"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 51be24c6a289 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-4e5ad1a125eae31e32a88e6ebab4b027d2ca3599af65c962f4bf11181c237cf6)
- [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-2ef3ef525fcf5c2e84040f0a59757519078457887cc7fa4cebd33a88e6b127ff)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-c78defbebc9a33045380b30673711c8de31a5bede00cf1da0c301ee08485f939)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-4c64b9c6056c33fc1a9d7f3ec2a749c7087c7c1a16befba223e3d210748db384)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-a70845ec6315fdd7e64db083d8ea4b145a68cbd8fa6d1f0b7072d74dc6851f1a"></a>

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

<a id="canonical-337e4e9521103e642ade2a554d586489ce74a16c84aaae2cbf6709d7d5467c83"></a>

## Direct properties — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 51be24c6a289 / 3

- [interface](resources--securemesh_site--reference--group-004.md#canonical-d42505ead3e354e7099f2b4073ef8b3b569ab6ddff818010d0ffc72cd3c33aec): complete subsection reference.

<a id="canonical-4bfa73f630657482c707d01ea9476f4ae1d20acfc78c96225d64a88faa45e917"></a>

<a id="canonical-79609517f74f9f5f99497ce9f2414cf88072f8b925f4f2ac63b16374c3ad114f"></a>

## node property — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 51be24c6a289 / 4

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

<a id="canonical-c6ff7e78f5beb373e46bfae3fbe75c15e804c2aa4b4eb7965ae889d5056901ef"></a>

## Next pages — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / 51be24c6a289 / 5

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site--reference--group-004.md#canonical-d42505ead3e354e7099f2b4073ef8b3b569ab6ddff818010d0ffc72cd3c33aec)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-4c64b9c6056c33fc1a9d7f3ec2a749c7087c7c1a16befba223e3d210748db384)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-d42505ead3e354e7099f2b4073ef8b3b569ab6ddff818010d0ffc72cd3c33aec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-234823ef085fcb7d167fd4ae5c3c8c56fc1104d3d4ca97f13ea024072d112e9b"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / af8d486c6073 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [custom_network_config.slo_config](resources--securemesh_site--reference--group-003.md#canonical-4e5ad1a125eae31e32a88e6ebab4b027d2ca3599af65c962f4bf11181c237cf6)
- [custom_network_config.slo_config.static_v6_routes](resources--securemesh_site--reference--group-004.md#canonical-2ef3ef525fcf5c2e84040f0a59757519078457887cc7fa4cebd33a88e6b127ff)
- [custom_network_config.slo_config.static_v6_routes.static_routes](resources--securemesh_site--reference--group-004.md#canonical-c78defbebc9a33045380b30673711c8de31a5bede00cf1da0c301ee08485f939)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site--reference--group-004.md#canonical-4c64b9c6056c33fc1a9d7f3ec2a749c7087c7c1a16befba223e3d210748db384)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-004.md#canonical-b0560a606c1fce2fd395d55c5b26293e9bbc7dafda0263b2c350e4c2ba8ae3f7)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-a670428ef76047a7ab2687a261fbe5a25ee7297dd04daddd1463c61d1fd65cd0"></a>

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

<a id="canonical-43a5bfa8d0ca9039195b79a9dad4c71072e3ad80861abc859277e9d63ec40284"></a>

## Direct properties — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / af8d486c6073 / 3

<a id="canonical-fd9c8405b29800ca92f5394e175f0362b965d7a4bc7dc1931080e1a464022a5b"></a>

<a id="canonical-22dd38928e6542c9d156109d2934d75c11b0d69cd3c960d499ee449a7f1e1c4d"></a>

## kind property — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / af8d486c6073 / 4

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

<a id="canonical-d3b6e5e0ad420df54c6c4163c73ce1364574b492b0f9cff174404bd5387350c9"></a>

<a id="canonical-511cc4bd76c7392a884f5ed38a0606f1d9698851bf9dd64c272d05055c3aaaed"></a>

## name property — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / af8d486c6073 / 5

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

<a id="canonical-6b653db4e1ad0e1221e577fcae7e68a9fc325f227165e3baeb5563380450f788"></a>

<a id="canonical-b9810fcee9896296155bc6c796687df54a0a5f08bc94ebe6964b6c03a0a66a3a"></a>

## namespace property — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / af8d486c6073 / 6

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

<a id="canonical-9298f4fa3bdb9fff39ab713cb98e847a89d15c1fd4c5d0e4b0e141e7d8ae8a29"></a>

<a id="canonical-60192497df815c2c7814457909e8c1a8e67b2cca1ad79f9e6277a1e69df99f60"></a>

## tenant property — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / af8d486c6073 / 7

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

<a id="canonical-dba9bdebb224d568389f8b3e87465a35c0818da8bce1184f65cff70f14667c3f"></a>

<a id="canonical-104136b1220dcc2f9bb17df2b74d610aab0a9de06971ab5cc350894531191da7"></a>

## uid property — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / af8d486c6073 / 8

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

<a id="canonical-e0357ba4ef893f1d17b79ca6c059b221be6e5df172b45fbdc9c5ce55866c469e"></a>

## Next pages — custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.l / af8d486c6073 / 9

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site--reference--group-004.md#canonical-b0560a606c1fce2fd395d55c5b26293e9bbc7dafda0263b2c350e4c2ba8ae3f7)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-9d47714be24e01c0fbd23354a4e2f5bc22bedd857689e3612df4118fc0681053"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b49ea2eeef71f2809065cf2936706cc277a0f0cfb9c1e001fbd27bd45c5164b2"></a>

## custom_network_config.sm_connection_public_ip — custom_network_config.sm_connection_public_ip / cc3ca4c66533 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- custom_network_config.sm_connection_public_ip

<a id="canonical-f0fe59ddbc39010fd803346b06a0dfd06ae175e7be0510456de61fd4b394dffb"></a>

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

<a id="canonical-1810cbb0b441ad198f56d630f4d6d4b999fe71b9b301ed99a8a29b38b8ab517d"></a>

## Direct properties — custom_network_config.sm_connection_public_ip / cc3ca4c66533 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-75e6a6e506b1bfa2215fb5cf12873e5b60019ee9baded246d688e70e2bc66156"></a>

## Next pages — custom_network_config.sm_connection_public_ip / cc3ca4c66533 / 4

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-a1caa42f5fc7bc4f4b913c717ce9b31da256bb3e598e725d5049b3208030f6d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-288c03b56f548a2bd8ccf7655697d2b8ffd6d780a36e96726086b2847961c37d"></a>

## custom_network_config.sm_connection_pvt_ip — custom_network_config.sm_connection_pvt_ip / 9179739000bf / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- custom_network_config.sm_connection_pvt_ip

<a id="canonical-d911e9af9bc1b8a477d016be97dc2e2b07da59f1c1b166a80fbdf6e0f38c68b4"></a>

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

<a id="canonical-baf52630a6ac979fdd86a3c60b90796ed482945ec6837732c59d51d0a97948ac"></a>

## Direct properties — custom_network_config.sm_connection_pvt_ip / 9179739000bf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-32dc4b2d278e26aa5c2b9b3b0b6851a7aca258396c19402c290dd622844719d8"></a>

## Next pages — custom_network_config.sm_connection_pvt_ip / 9179739000bf / 4

- [custom_network_config](resources--securemesh_site--reference--group-001.md#canonical-3f87ddd2820f6bbd9b51a46fce669cc3fbd36f93ac1400ea551f52772e5125ec)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-9ce856978b9d1a3834c3b63a67ff7d42db5ed137144b2db86037212335e19331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8115c2dc9146a3c23a55c019a6c45f47e86b97af3163e507d9e41688157adf6e"></a>

## default_blocked_services — default_blocked_services / dbf760de6d35 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- default_blocked_services

<a id="canonical-73afebe2d256020451a57e09c99766c20af7833ca105c606be8813c42a418f9c"></a>

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
default_blocked_services = {}
```

<a id="canonical-03bfdb463ee97e0d98161ee542e0e06986da38f17b86adab4765037d1c2d5511"></a>

## Direct properties — default_blocked_services / dbf760de6d35 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a88349994e86e4faa004486e780adff8ccb1713059cdb5516398e457aa266db4"></a>

## Next pages — default_blocked_services / dbf760de6d35 / 4

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-6598b3c9c785b28b0be5365f6ae8bc3b1d78ecfffef33c5a8ea18ada89496ed4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13e4acf0ad69385f9f6537a65e3d49ce879632ee3df5661904b06d86ccbf7efc"></a>

## default_network_config — default_network_config / 35be8fd81f90 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- default_network_config

<a id="canonical-871ddd6e9fca923ed67c5ad66b8affcef0f12462f8565e9538776a3e61f8efb5"></a>

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
default_network_config = {}
```

<a id="canonical-fed7418789a0d8e479e42700c76f01cfb24b84e80fb3327994fede79f326b8b9"></a>

## Direct properties — default_network_config / 35be8fd81f90 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c0cd9940aedc1a154304ac1afa2a2d02629871bdc186e6f9ba8dc2087b871c66"></a>

## Next pages — default_network_config / 35be8fd81f90 / 4

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-a9e0287fe45dbce2f97d69ddb1db27fec38f2212e9065dbda674fc9fa7aadd69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6560b1c2eb5624d4278eb089f10e12ce2e9cd3929a678543283f0b608706a8a"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / 5aa9a77966fa / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- kubernetes_upgrade_drain

<a id="canonical-60a04a3b53b7ee4e977a3f20a3bf33612673816cf8e23ee80f0828844a5fefca"></a>

Type: `"object"`. single nested block, Optional.

Specify how worker nodes within a site will be upgraded.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_upgrade_drain",
    "enable_upgrade_drain")}
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
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

Terraform syntax:

```terraform
kubernetes_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-ea9c77869e5c939eb404c855411b5546a89d96422216b9a03f40c97df7586995"></a>

## Direct properties — kubernetes_upgrade_drain / 5aa9a77966fa / 3

- [disable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-2bba3f21873ac9b02d15803a49f38d47846cc488938a549f5b9844d33a5db1dc): complete subsection reference.

- [enable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-dac6f0e5dff414c4a1ea72812615be1f00d185e078af4fa053a27e9f725e5b84): complete subsection reference.

<a id="canonical-fcb8cebed6092151ed4d90389488d5af17207c9db63fbc98ef501f492518be24"></a>

## Next pages — kubernetes_upgrade_drain / 5aa9a77966fa / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-2bba3f21873ac9b02d15803a49f38d47846cc488938a549f5b9844d33a5db1dc)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-dac6f0e5dff414c4a1ea72812615be1f00d185e078af4fa053a27e9f725e5b84)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-2bba3f21873ac9b02d15803a49f38d47846cc488938a549f5b9844d33a5db1dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e3b566f4afadf5ceb4f2145dfb71d9e0905b6be9ea18793ba62f176ba776467"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — kubernetes_upgrade_drain.disable_upgrade_drain / 9b64b2d82ec4 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-a9e0287fe45dbce2f97d69ddb1db27fec38f2212e9065dbda674fc9fa7aadd69)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-513cfbb07b012f77e4dd6c86a4b24df90399ee81104edc80f79c05480118e5bc"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable upgrade drain.

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
disable_upgrade_drain = {}
```

<a id="canonical-4893e671e3210bbef219b375b7ad042c0274c1fcbc330ab545ea4eee0928263f"></a>

## Direct properties — kubernetes_upgrade_drain.disable_upgrade_drain / 9b64b2d82ec4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fba97f32ebeaf1f13814a0ae10d35e5c6b37d57cc4a1db0510d9a99b08dad27e"></a>

## Next pages — kubernetes_upgrade_drain.disable_upgrade_drain / 9b64b2d82ec4 / 4

- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-a9e0287fe45dbce2f97d69ddb1db27fec38f2212e9065dbda674fc9fa7aadd69)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-dac6f0e5dff414c4a1ea72812615be1f00d185e078af4fa053a27e9f725e5b84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6974623ffe75229ba286a4d6f01be5cf9d83fb6a6c17bdd00b3a19f3bf6dee97"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — kubernetes_upgrade_drain.enable_upgrade_drain / 1406e51f7bcd / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-a9e0287fe45dbce2f97d69ddb1db27fec38f2212e9065dbda674fc9fa7aadd69)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-02ae73d7973662f340b89c01cf0507b04bb67b51cec717b99023bdc7c36a5638"></a>

Type: `"object"`. single nested block, Optional.

Specify batch upgrade settings for worker nodes within a site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("drain_node_timeout"),
  validators.ConflictingObjectAttributes("disable_vega_upgrade_mode",
    "enable_vega_upgrade_mode"),
  validators.ConflictingObjectAttributes("drain_max_unavailable_node_count",
    "drain_max_unavailable_node_percentage")}
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
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

Terraform syntax:

```terraform
enable_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-29154a55120e582d679e97e923d1e238f3fca6919f74a8ff40e1933e6635104f"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain / 1406e51f7bcd / 3

- [disable_vega_upgrade_mode](resources--securemesh_site--reference--group-004.md#canonical-77075d4e1e6e91f1849bd563bd70b14b740ff24d7f0e9516894990939cf01654): complete subsection reference.

<a id="canonical-fcd4ed04036b22fa5930d40a646c68fe6e953e2f20a40883af11d9d199ccb565"></a>

<a id="canonical-183df77928446c44dc17ff85c1079c2b0777fef542dc1426064fadab4c1f2d4e"></a>

## drain_max_unavailable_node_count property — kubernetes_upgrade_drain.enable_upgrade_drain / 1406e51f7bcd / 4

Type: `"number"`. Optional.

Node Batch Size Count. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5000),
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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-a3def8c765fd65a9ba09291c64d9bdd59eb9cb7cc92f2d666e2b16a529d33fa6"></a>

<a id="canonical-a7cffc5aa1df7cec349b941d41a4d3439de651b352b6e20d0b463a84a5119cec"></a>

## drain_max_unavailable_node_percentage property — kubernetes_upgrade_drain.enable_upgrade_drain / 1406e51f7bcd / 5

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-16fe608a27b26482e9cd2a8ee977cc0f1e24f84589141b0262439d51a59f9e59"></a>

<a id="canonical-714bfc080b9fcd9b81758c6fcd5ad784c71b72744166b60ee3d022d5fa506115"></a>

## drain_node_timeout property — kubernetes_upgrade_drain.enable_upgrade_drain / 1406e51f7bcd / 6

Type: `"number"`. Optional.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

Upstream description:

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 900),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
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
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](resources--securemesh_site--reference--group-004.md#canonical-9768465db20e6c12616cf2f075da69409a254945515a259050cb1079fe5c2c0d): complete subsection reference.

<a id="canonical-00ca012c8c6ea22ed8d03cdcc8251abdef3d0055512bb28a684b2e75a6ba8803"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain / 1406e51f7bcd / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--securemesh_site--reference--group-004.md#canonical-77075d4e1e6e91f1849bd563bd70b14b740ff24d7f0e9516894990939cf01654)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--securemesh_site--reference--group-004.md#canonical-9768465db20e6c12616cf2f075da69409a254945515a259050cb1079fe5c2c0d)
- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-a9e0287fe45dbce2f97d69ddb1db27fec38f2212e9065dbda674fc9fa7aadd69)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-77075d4e1e6e91f1849bd563bd70b14b740ff24d7f0e9516894990939cf01654"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7f8fa53ba91e66f7b22af196d404b79f0dd6d819065c24d7f6d7b1261aad196"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 61a185d2bc45 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-a9e0287fe45dbce2f97d69ddb1db27fec38f2212e9065dbda674fc9fa7aadd69)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-dac6f0e5dff414c4a1ea72812615be1f00d185e078af4fa053a27e9f725e5b84)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-6a3508f965c6874365e370c5fb1e2ac8ec60a65546e04a9442aa250138faf4c7"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable vega upgrade mode.

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
disable_vega_upgrade_mode = {}
```

<a id="canonical-401dd98d77722a3baee4dd8b446490f5f064c5cf5470d69add6b8113fd727897"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 61a185d2bc45 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6ff81b1577dbb14bef3873331afd3130ffa3b6a72f5b5a18126ee0d4fee7f049"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 61a185d2bc45 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-dac6f0e5dff414c4a1ea72812615be1f00d185e078af4fa053a27e9f725e5b84)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-9768465db20e6c12616cf2f075da69409a254945515a259050cb1079fe5c2c0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc730dfdbe1f112dae6c8807d1e58a0b1432fa7c3f592ad011a3de87a5b3cbb4"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 781ac77e60d6 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [kubernetes_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-a9e0287fe45dbce2f97d69ddb1db27fec38f2212e9065dbda674fc9fa7aadd69)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-dac6f0e5dff414c4a1ea72812615be1f00d185e078af4fa053a27e9f725e5b84)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-ddc1746b969a8572e3332bf364d17bdfdce24ac48d33b87de34dd70292385418"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable vega upgrade mode.

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
enable_vega_upgrade_mode = {}
```

<a id="canonical-6f9204c1e08b9aff7054afdc28d7a98d61d1f81e017b832c90742ef22bccda97"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 781ac77e60d6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c632e7bc2a360b3cfcc755dc5fdc5308b846d17636e9ad8b5c2a55bffdc350b3"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 781ac77e60d6 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site--reference--group-004.md#canonical-dac6f0e5dff414c4a1ea72812615be1f00d185e078af4fa053a27e9f725e5b84)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-64b4390e24c8d3cd7963fefe17327c30c38775d5b9fc0c3a76696addb3c8473b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd00166efa43bffa6c16b620852fe162552699e9b5e98d76f0dac70cdf9a6457"></a>

## log_receiver — log_receiver / 208bfafdb3d1 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- log_receiver

<a id="canonical-c3429a9023d31ca9356e3d2ceb1e9ca7237c8ccdd903da6dbe8252205cbf9bce"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [log_receiver](resources--securemesh_site--reference--group-004.md#canonical-c3429a9023d31ca9356e3d2ceb1e9ca7237c8ccdd903da6dbe8252205cbf9bce)
- [logs_streaming_disabled](resources--securemesh_site--reference--group-004.md#canonical-1a185866c3136ce73bbc47c579c8cc7ef9d53505e40412b5e387d04158d63237)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-c5ecfdf0d0468c9125eeaed7d55eef72809d73ac59e2e0a126db0c2118dc5333"></a>

## Direct properties — log_receiver / 208bfafdb3d1 / 3

<a id="canonical-880168d21c2f4f1d2ea5333ecb1cfe4b3c5a8a56f75f8c0a792d01b744070d8c"></a>

<a id="canonical-ea6f7901310f1115d4a7f24e62698bcbf2296ff0901994d6979e29ff4ebc180e"></a>

## name property — log_receiver / 208bfafdb3d1 / 4

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

<a id="canonical-45100d480fabe69ccad69c66552f4264173d6cc89b9216b74d3540835f44d7c9"></a>

<a id="canonical-833f9426d980497148b719a2e6008406758630d7bcacdf80e79f61bc2be6a9f6"></a>

## namespace property — log_receiver / 208bfafdb3d1 / 5

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

<a id="canonical-4a4d52b332d1366e79ecc92abb691e2571efe7b060b54b519a89a10f192e337a"></a>

<a id="canonical-258d3c559862a7404a24dcc146b01bca155a9ab6dc64a55f8b9c7a970757fb50"></a>

## tenant property — log_receiver / 208bfafdb3d1 / 6

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

<a id="canonical-8c8c783d712a5bfba4c82be795d81eb817e13e3bfbd5b2378a1ce9e490e3a07d"></a>

## Next pages — log_receiver / 208bfafdb3d1 / 7

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-31eb5954f742df58c87702640dbcd9722b462484fa62264b963a0128a396fc02"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a545982ba4434c2077b92284311d8ad96c6718e81143d53c01e1fdce02b1c5f5"></a>

## logs_streaming_disabled — logs_streaming_disabled / 6c4d96358893 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- logs_streaming_disabled

<a id="canonical-1a185866c3136ce73bbc47c579c8cc7ef9d53505e40412b5e387d04158d63237"></a>

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

<a id="canonical-f638475795e5f992fc7b7638c8cb1eccceae09faddcb2909bc68c1b1a7c798ca"></a>

## Direct properties — logs_streaming_disabled / 6c4d96358893 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6981503204c81998ff1a667018e7e473450f22eec3881a4da98bd96962e08f47"></a>

## Next pages — logs_streaming_disabled / 6c4d96358893 / 4

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-8c2774d9115547c29c565d10884a46a3d040cea98f976355ca3e3dace10ebd9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25f5c19ce929265e9f15817bccfc568f934564bf7b48b557f9eda6753f47ac96"></a>

## master_node_configuration — master_node_configuration / 28e87c0a7e95 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- master_node_configuration

<a id="canonical-9c95772323135cd8a70331c22a79957f3c75c30d1a108658788e4bfbfebe52bd"></a>

Type: `"object"`. list nested block, Optional.

Master Nodes. Configuration of master nodes.

Upstream description:

Configuration of master nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

Terraform syntax:

```terraform
master_node_configuration {
  # Configure direct properties listed below.
}
```

<a id="canonical-8a735286401b3fa66ab9a5fc3fde404a256998fc8be466246539acaabf9fbe09"></a>

## Direct properties — master_node_configuration / 28e87c0a7e95 / 3

<a id="canonical-867de7efda32bacb12c4d4ee3a58b0008c732d8664d3dd811279e65e9019258f"></a>

<a id="canonical-b2f42e2228c7196fbf5808c3a9997e7efcc96bcebaf3ff6d123424fda6eca8e6"></a>

## name property — master_node_configuration / 28e87c0a7e95 / 4

Type: `"string"`. Optional.

Name. Names of master node.

Upstream description:

Names of master node.

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

<a id="canonical-0d4bdf755625ce5809ddc50eb4b9a3b7cb6e41dd7df006d18b3c1e3844e10169"></a>

<a id="canonical-9171505139f0068f7877ecab0eb2e14087ec4afa24d85f8d10ffbe9083c5bebb"></a>

## public_ip property — master_node_configuration / 28e87c0a7e95 / 5

Type: `"string"`. Optional.

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

Upstream description:

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

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

<a id="canonical-e21d069bed1cf615647b37eda248fed2135739be4a0047325d64685d4cf47793"></a>

## Next pages — master_node_configuration / 28e87c0a7e95 / 6

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-3f591357d9730ab49c161f878602fbcbad1f798979daab263ba21dd5e07fe37c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0088f3eaad835bcfbb3284e16fbce4c76875c80a7c3de3042a4d47f543497989"></a>

## no_bond_devices — no_bond_devices / 25f620ba5007 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- no_bond_devices

<a id="canonical-17491a1b8d51d7cc729febc7b394fbeeefd2c87cd44aee4af438dbf46de71db7"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no bond devices.

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
no_bond_devices = {}
```

<a id="canonical-d9fdf0082eec213fd355b8b2cb142a248419cc8339a9f93f9188827a546eeeaf"></a>

## Direct properties — no_bond_devices / 25f620ba5007 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4b9ca2cf37a22669cf9bdfd1fa21b3ea3396e6b3b016bea3e3128bc8719b8e53"></a>

## Next pages — no_bond_devices / 25f620ba5007 / 4

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-8876dd170b5fea86ae9ef472fdff18fdf1d5c6c740e6d812f7d8ab5751811350"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9de60ab6ee4ffad5401b92843a6b29eb48b9884f93208f7252fac06d3d05f795"></a>

## offline_survivability_mode — offline_survivability_mode / 27410943e23c / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- offline_survivability_mode

<a id="canonical-1a249b8b55ca5b0995a31662206e87d414bb2db0089bef5ad38111871ebfa1a4"></a>

Type: `"object"`. single nested block, Optional.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("enable_offline_survivability_mode",
    "no_offline_survivability_mode")}
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
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

Terraform syntax:

```terraform
offline_survivability_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-b42c76ad26957f5e4680566bbc782cc1c5c56789fa4442003601e95f6a11a4e1"></a>

## Direct properties — offline_survivability_mode / 27410943e23c / 3

- [enable_offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-a2eba802f84d90a0a77ec3bf0085dd63ee0aa3e28d7c2e7ca6a2e71b2e1a48a2): complete subsection reference.

- [no_offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-014baef5ea2d126e615fdb725c80312c71f38269569f2bfd559b00f9e7093c48): complete subsection reference.

<a id="canonical-1e08c514c93007b48cfa146b5fda912ed6c0c7ef19d60cd1dcbf820984802bee"></a>

## Next pages — offline_survivability_mode / 27410943e23c / 4

- [offline_survivability_mode.enable_offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-a2eba802f84d90a0a77ec3bf0085dd63ee0aa3e28d7c2e7ca6a2e71b2e1a48a2)
- [offline_survivability_mode.no_offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-014baef5ea2d126e615fdb725c80312c71f38269569f2bfd559b00f9e7093c48)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-a2eba802f84d90a0a77ec3bf0085dd63ee0aa3e28d7c2e7ca6a2e71b2e1a48a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edb9c84b6bfea1245c3dbc752d9af672fa14f88a544211b3842b890149691357"></a>

## offline_survivability_mode.enable_offline_survivability_mode — offline_survivability_mode.enable_offline_survivability_mode / 806a47a68335 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-8876dd170b5fea86ae9ef472fdff18fdf1d5c6c740e6d812f7d8ab5751811350)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-d04e0a58c12a7d867762a5676ce962f7dda257a005f3923d95a9c87c1a4381db"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable offline survivability mode.

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
enable_offline_survivability_mode = {}
```

<a id="canonical-ed7df2969877f57d284a46f8ba4595f78f6e80634ebb4d5ffbb319512b55d8b4"></a>

## Direct properties — offline_survivability_mode.enable_offline_survivability_mode / 806a47a68335 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe9e4316a5a111a5b80e39552009be690418b8a50220d6b951eb14347b02d2fc"></a>

## Next pages — offline_survivability_mode.enable_offline_survivability_mode / 806a47a68335 / 4

- [offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-8876dd170b5fea86ae9ef472fdff18fdf1d5c6c740e6d812f7d8ab5751811350)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-014baef5ea2d126e615fdb725c80312c71f38269569f2bfd559b00f9e7093c48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27ad3f560d18c7650f6e0843eb6d4f4c4fe2086d6b4bf39a58ee7932c92f7d93"></a>

## offline_survivability_mode.no_offline_survivability_mode — offline_survivability_mode.no_offline_survivability_mode / bc53419d614c / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-8876dd170b5fea86ae9ef472fdff18fdf1d5c6c740e6d812f7d8ab5751811350)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-d7f3e9231ade0350555feaf6a50217f60e986fa4d2efd97a4d4fdc5d1768bd60"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no offline survivability mode.

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
no_offline_survivability_mode = {}
```

<a id="canonical-3d4139b7268a8c6cc63c17090b1288cda0984e2bfca86a54009838ccee3ab255"></a>

## Direct properties — offline_survivability_mode.no_offline_survivability_mode / bc53419d614c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-86eea9a83126afbd80ef0cdfd870c923c30078bde1fd92207b6ab6f96bd355e1"></a>

## Next pages — offline_survivability_mode.no_offline_survivability_mode / bc53419d614c / 4

- [offline_survivability_mode](resources--securemesh_site--reference--group-004.md#canonical-8876dd170b5fea86ae9ef472fdff18fdf1d5c6c740e6d812f7d8ab5751811350)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-127eee1dad136acb23d6139da4b0993a34521376b95d09aa218b63fcc060cafa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18edd9544410854944d5f3442b8047b1bec2a61be67e108b5976304df7f37275"></a>

## os — os / c82951ea2f94 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- os

<a id="canonical-7ace6872dc9b8c6d170516a8433498d0a83fafab0b7b830465761914927b0946"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Upstream description:

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_os_version",
    "operating_system_version")}
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
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

Terraform syntax:

```terraform
os {
  # Configure direct properties listed below.
}
```

<a id="canonical-4001d998950b705e68a577de63cc171f08a3a0424f2ee27841d10f06c0c9309f"></a>

## Direct properties — os / c82951ea2f94 / 3

- [default_os_version](resources--securemesh_site--reference--group-004.md#canonical-dca71fe3477a1dfd3479a364987935fbdcd4e142fcc56d4bbcb2a60381aa09e0): complete subsection reference.

<a id="canonical-2ed03b6e53c24249ca879c576eb32df96eb7ba4e66f2bf60348137d6685136ae"></a>

<a id="canonical-139c16455bc914be0de0ad63bf903f47eeb19808c2b0f73858a992095a83fe62"></a>

## operating_system_version property — os / c82951ea2f94 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Upstream description:

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-4cb15ba0badff60fa15dcca300a26ba42491319a943a36f95fd5d6435ae006f2"></a>

## Next pages — os / c82951ea2f94 / 5

- [os.default_os_version](resources--securemesh_site--reference--group-004.md#canonical-dca71fe3477a1dfd3479a364987935fbdcd4e142fcc56d4bbcb2a60381aa09e0)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-dca71fe3477a1dfd3479a364987935fbdcd4e142fcc56d4bbcb2a60381aa09e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-920d56985779237a81bc30307d6a2c51dbcd6346bf47d4b3f63bccea014f534d"></a>

## os.default_os_version — os.default_os_version / bfac03d81ffd / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [os](resources--securemesh_site--reference--group-004.md#canonical-127eee1dad136acb23d6139da4b0993a34521376b95d09aa218b63fcc060cafa)
- os.default_os_version

<a id="canonical-cd92a3cd16202fd5927a4179689fdfb5264e6f01431245dbaacbd11cad8ab28f"></a>

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
default_os_version = {}
```

<a id="canonical-66cdaa6e653b793c4aae8013187948c98e4dedc05ecd15ed01c48abc07048b21"></a>

## Direct properties — os.default_os_version / bfac03d81ffd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe5812988c274467be68f650acadf94ee885047f79a7ff00ed87d2ffc8fee4de"></a>

## Next pages — os.default_os_version / bfac03d81ffd / 4

- [os](resources--securemesh_site--reference--group-004.md#canonical-127eee1dad136acb23d6139da4b0993a34521376b95d09aa218b63fcc060cafa)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-05bbccf708211046ecbf23eb7d63a5d01bf0f6cf2f154cc39615874b9c89574f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1807cf9cda9e7ee326e5fe929dba0cb3d33710ad6879d16c8b0bf32fc0af064d"></a>

## performance_enhancement_mode — performance_enhancement_mode / bb10ebe8b9c5 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- performance_enhancement_mode

<a id="canonical-2086c93376d1db431dd641b4eb603ae3d64a5c3ec094bca55a212b51437d92f2"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("perf_mode_l3_enhanced",
    "perf_mode_l7_enhanced")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

Terraform syntax:

```terraform
performance_enhancement_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-f713c38492acce740646a01dcfcd8d7b5e358746714009af63dd1cd4d88f454c"></a>

## Direct properties — performance_enhancement_mode / bb10ebe8b9c5 / 3

- [perf_mode_l3_enhanced](resources--securemesh_site--reference--group-004.md#canonical-bfde62b1aecc91000ebff7ee69c88dd75c12837304ddb00b3af9e52693e0871f): complete subsection reference.

- [perf_mode_l7_enhanced](resources--securemesh_site--reference--group-004.md#canonical-ee614cb375e6ad98ee8524afb6e44c523b79a9cd8db9b16200a260829ad2c5d7): complete subsection reference.

<a id="canonical-0a11f5cba7b0de53b57af20292371124a24a82684bba54b62c74460463c3c986"></a>

## Next pages — performance_enhancement_mode / bb10ebe8b9c5 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site--reference--group-004.md#canonical-bfde62b1aecc91000ebff7ee69c88dd75c12837304ddb00b3af9e52693e0871f)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site--reference--group-004.md#canonical-ee614cb375e6ad98ee8524afb6e44c523b79a9cd8db9b16200a260829ad2c5d7)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-bfde62b1aecc91000ebff7ee69c88dd75c12837304ddb00b3af9e52693e0871f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3824c959f9a7a967afc8a698b5e460f3d1fcdd4ca8c0caba65eab8ce2f4341c3"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced — performance_enhancement_mode.perf_mode_l3_enhanced / a04f6c465d1f / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-05bbccf708211046ecbf23eb7d63a5d01bf0f6cf2f154cc39615874b9c89574f)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-7d7fc8dfac75ca797b002996d9f713317c55b2048d1dabf7fefaf0311566603e"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo",
    "no_jumbo")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l3_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-44dca39dbe0b7d7f94669ae46124f8fcdd2bb4143790c22dc48f265b8facc681"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l3_enhanced / a04f6c465d1f / 3

- [jumbo](resources--securemesh_site--reference--group-004.md#canonical-d69e782329ed07acb035525139da8b755e4004ed19fc5dfeb283602a985ddb80): complete subsection reference.

- [no_jumbo](resources--securemesh_site--reference--group-004.md#canonical-930781c3e4dc2e26a55b41447188026ad45a4c5c27dbcea38730447894b62298): complete subsection reference.

<a id="canonical-e85e3c9593d7c2100dca29e3290f99a1736b5a9e5acdd95843c2572b8a1c137c"></a>

## Next pages — performance_enhancement_mode.perf_mode_l3_enhanced / a04f6c465d1f / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--securemesh_site--reference--group-004.md#canonical-d69e782329ed07acb035525139da8b755e4004ed19fc5dfeb283602a985ddb80)
- [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--securemesh_site--reference--group-004.md#canonical-930781c3e4dc2e26a55b41447188026ad45a4c5c27dbcea38730447894b62298)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-05bbccf708211046ecbf23eb7d63a5d01bf0f6cf2f154cc39615874b9c89574f)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-d69e782329ed07acb035525139da8b755e4004ed19fc5dfeb283602a985ddb80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78116efa69e4091d1722013dbf25c83e2f5457d27530cf8912fbb66dcebd599d"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / f8ca8b567dbe / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-05bbccf708211046ecbf23eb7d63a5d01bf0f6cf2f154cc39615874b9c89574f)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site--reference--group-004.md#canonical-bfde62b1aecc91000ebff7ee69c88dd75c12837304ddb00b3af9e52693e0871f)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-e4b0d757ebbe608213d7b5555c10172d9878f23edb5ede4d180314d6a17c4850"></a>

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
jumbo = {}
```

<a id="canonical-35ded4bb3f98a50dc475b6a7481ba36f8888945b5ae2c1f4dfa747c400278f09"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / f8ca8b567dbe / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-02c874cb461295fc071440fbc2c5e45657d5f97b9b1cf11f3ff20c0e970dacc9"></a>

## Next pages — performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / f8ca8b567dbe / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site--reference--group-004.md#canonical-bfde62b1aecc91000ebff7ee69c88dd75c12837304ddb00b3af9e52693e0871f)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-930781c3e4dc2e26a55b41447188026ad45a4c5c27dbcea38730447894b62298"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21a68ab3adaec461f23d0c9ee4c0e360293407b14df75c457cb6f4699650e1c8"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 6a66e9155226 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-05bbccf708211046ecbf23eb7d63a5d01bf0f6cf2f154cc39615874b9c89574f)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site--reference--group-004.md#canonical-bfde62b1aecc91000ebff7ee69c88dd75c12837304ddb00b3af9e52693e0871f)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-0d2d462520408d21a91197c46a29b53cd671fd5265e0790a89cd258767c230b4"></a>

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
no_jumbo = {}
```

<a id="canonical-e5b5e004c7d651c047fc6fb9b55298f07daaedb7268b08f7b396548cc6400b25"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 6a66e9155226 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5ddbf9cfe8db178c1e0fe052d9598914ef5b2bb5467898edfb6e1c9e9d1ce5a6"></a>

## Next pages — performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 6a66e9155226 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--securemesh_site--reference--group-004.md#canonical-bfde62b1aecc91000ebff7ee69c88dd75c12837304ddb00b3af9e52693e0871f)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-ee614cb375e6ad98ee8524afb6e44c523b79a9cd8db9b16200a260829ad2c5d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c20b314da65a4185c3ecea1a49ef61077f123e6f4df2ecfa5c8686c56155f301"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced — performance_enhancement_mode.perf_mode_l7_enhanced / c1a36c464a4c / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-05bbccf708211046ecbf23eb7d63a5d01bf0f6cf2f154cc39615874b9c89574f)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-47c5c38b2a80d0652acf45c4a150ff6266f29b005c3cea3980315c40aefae9ca"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo_disabled",
    "jumbo_enabled")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l7_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-59c4ac282fad7ec465cc513b451337743301dc6753c6dd75a75a6ebc086c86b9"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l7_enhanced / c1a36c464a4c / 3

- [jumbo_disabled](resources--securemesh_site--reference--group-004.md#canonical-b72fad071fbc38f791792e56e9ba98fa2d4ef401a67bced1e8d6c9bb54e8b70a): complete subsection reference.

- [jumbo_enabled](resources--securemesh_site--reference--group-004.md#canonical-16a61df606ee655f4facf7fd0a2ba19781170c0d855244864e00521a4c6cd80a): complete subsection reference.

<a id="canonical-a04a70a7fc9f274b46038863ef994213e5b2f407c2f23a0df38e2eaa54217614"></a>

## Next pages — performance_enhancement_mode.perf_mode_l7_enhanced / c1a36c464a4c / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--securemesh_site--reference--group-004.md#canonical-b72fad071fbc38f791792e56e9ba98fa2d4ef401a67bced1e8d6c9bb54e8b70a)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--securemesh_site--reference--group-004.md#canonical-16a61df606ee655f4facf7fd0a2ba19781170c0d855244864e00521a4c6cd80a)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-05bbccf708211046ecbf23eb7d63a5d01bf0f6cf2f154cc39615874b9c89574f)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-b72fad071fbc38f791792e56e9ba98fa2d4ef401a67bced1e8d6c9bb54e8b70a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4695674fa737cbaba632e63673a1dc3d69c4aa34dcf1030b8ac27647098cc7f6"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / b906c1ee2cdf / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-05bbccf708211046ecbf23eb7d63a5d01bf0f6cf2f154cc39615874b9c89574f)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site--reference--group-004.md#canonical-ee614cb375e6ad98ee8524afb6e44c523b79a9cd8db9b16200a260829ad2c5d7)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-ead2ae05e2e008a3f6f842395a8b169f244bac2a5031a445000978a3c51866e5"></a>

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
jumbo_disabled = {}
```

<a id="canonical-df5beff589780935678a5473abe977e811977ffc856472cc7875e85981cb452a"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / b906c1ee2cdf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d18996bf210e5d520534050595d8630d81b68730edf8a8e2eaaaff5d54661b54"></a>

## Next pages — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / b906c1ee2cdf / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site--reference--group-004.md#canonical-ee614cb375e6ad98ee8524afb6e44c523b79a9cd8db9b16200a260829ad2c5d7)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-16a61df606ee655f4facf7fd0a2ba19781170c0d855244864e00521a4c6cd80a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3f8d201dc6e43a50eec17662dbc03096d932919530c3d6568f8b49211fc99d4"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / 92a5ff5bdcfa / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [performance_enhancement_mode](resources--securemesh_site--reference--group-004.md#canonical-05bbccf708211046ecbf23eb7d63a5d01bf0f6cf2f154cc39615874b9c89574f)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site--reference--group-004.md#canonical-ee614cb375e6ad98ee8524afb6e44c523b79a9cd8db9b16200a260829ad2c5d7)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-1e54ae898297ac7a935bc66bcd4f7505fc6a0dcd39bea7265d1cd3c54f168e01"></a>

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
jumbo_enabled = {}
```

<a id="canonical-125da26424718714b5110e196ce27d5294f552f0bc22e0b129ce3161906f191e"></a>

## Direct properties — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / 92a5ff5bdcfa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f27d6179804c4f1ec3e5c1f9ef2448a2d90f5195b80cb6d346724585b0796f6e"></a>

## Next pages — performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / 92a5ff5bdcfa / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--securemesh_site--reference--group-004.md#canonical-ee614cb375e6ad98ee8524afb6e44c523b79a9cd8db9b16200a260829ad2c5d7)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-401806d8a416f123ff291c4a5f724bf9513d7fad7f5ce8928e9b0ac494e6f2e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb620f51abd8f2f40bc2bfaf4cf2952f7dd0f8b98820272b9b4b569f601a9b0b"></a>

## sw — sw / ccd71df862c0 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- sw

<a id="canonical-ebd1a47213748f6d21b5631cdc1f5ea280664bb65f93be44b8d4b3514c92bfd6"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Upstream description:

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_sw_version",
    "volterra_software_version")}
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
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

Terraform syntax:

```terraform
sw {
  # Configure direct properties listed below.
}
```

<a id="canonical-c720c68c8f07109a9926939346c676eefaf09f39be2d2b61861de813c4eedb5d"></a>

## Direct properties — sw / ccd71df862c0 / 3

- [default_sw_version](resources--securemesh_site--reference--group-004.md#canonical-9ca3e76c6fcb60e69cdd1cb6abad39d01836e7f653e3e96bf1266d73cae6b7a4): complete subsection reference.

<a id="canonical-7e6cce4a55528ddc5ad808132fde43c6011f1544a0c2e1644739a6f2a1ddc647"></a>

<a id="canonical-97b2e0a2b445b8946bedc40e089ecbf098eb17df31bd4e2a4a689b844867fc4d"></a>

## volterra_software_version property — sw / ccd71df862c0 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-9cd3cb74a34e3113763626a89feeeadb619ff98c8360fab6541fc474008cfec2"></a>

## Next pages — sw / ccd71df862c0 / 5

- [sw.default_sw_version](resources--securemesh_site--reference--group-004.md#canonical-9ca3e76c6fcb60e69cdd1cb6abad39d01836e7f653e3e96bf1266d73cae6b7a4)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-9ca3e76c6fcb60e69cdd1cb6abad39d01836e7f653e3e96bf1266d73cae6b7a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee7ac05c2640b8b10f21eb8a11af5e9aff0758bfee74c8ba511a90c9582d270f"></a>

## sw.default_sw_version — sw.default_sw_version / 493b036eb86a / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [sw](resources--securemesh_site--reference--group-004.md#canonical-401806d8a416f123ff291c4a5f724bf9513d7fad7f5ce8928e9b0ac494e6f2e7)
- sw.default_sw_version

<a id="canonical-523cec08581f6fc9d3b4b0997a9da3f9a9e1a54e310c5705622f10e5d12fa976"></a>

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
default_sw_version = {}
```

<a id="canonical-dceccd6fdc94b3d3b52d2281a02786dd1d68d0b2d10cecd69e297d26d94f9c6b"></a>

## Direct properties — sw.default_sw_version / 493b036eb86a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d4f62c62d200d8409dcfb431de9416a7473690277c1d036fee851f574a8165b8"></a>

## Next pages — sw.default_sw_version / 493b036eb86a / 4

- [sw](resources--securemesh_site--reference--group-004.md#canonical-401806d8a416f123ff291c4a5f724bf9513d7fad7f5ce8928e9b0ac494e6f2e7)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-edc9ffffe4ec4adfdc3b384259fa9b388799b14d92d94a5b7fb79492e7d5346e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1dd862f4fa13b499e0e35ae330c3d1fc4de7784ac2520e386113c33ee0c7abbb"></a>

## timeouts — timeouts / 2e7f5a0038bd / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- timeouts

<a id="canonical-3c6815cb7a5ba39b1d2f28d9a4762f6e08d75f8bb1ab429662cf90dcc33579b4"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-20ffba8ae3fba1b53b8b8c21a625d417af0ccf512e8dbb452945c73201a7e8d0"></a>

## Direct properties — timeouts / 2e7f5a0038bd / 3

<a id="canonical-2f49a683e88c514a953d4084dc9c94e93a9d2bc5505b8ea3137a10481eaced1d"></a>

<a id="canonical-95d9107608340edfe60b8824b4e876e4cb0fbb126c17827aa370a8edc657074c"></a>

## create property — timeouts / 2e7f5a0038bd / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-91597a72b6b0c7c2a594e097c780954e8d6603f3585fc6817c2d9590eb41fc40"></a>

<a id="canonical-f09f26611bc5a285bd13c5b9182985a86b142a59f0f06f68b43524becf8c393d"></a>

## delete property — timeouts / 2e7f5a0038bd / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-a7d1de996d2b3aa8c40fbe5304cd66a23a2798fdf4486f50bddee9877351deda"></a>

<a id="canonical-fecfdec2df42cbf8f46260d14eba32878febf8fef88ed06e46c6fff48a78610e"></a>

## read property — timeouts / 2e7f5a0038bd / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-e48a8692d84734640045b4b068039e7d89020c20d80b8810a3d6e05b7651e2b6"></a>

<a id="canonical-dbe5dd33d6694db46122a6a45ebce260e5a2bc0fccadbdd5f40caa396c04529d"></a>

## update property — timeouts / 2e7f5a0038bd / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-000152fae1d935dcac832349d4c21e30d7e9516dd042b694ad9df3787b37746a"></a>

## Next pages — timeouts / 2e7f5a0038bd / 8

- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-e3d3b3dd04a2677e5757751a4242bcaa7ace22229d58fc2ed27e711cfcaac7e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee5632e3c4f7399a8ce564746ffbed1edaba88e244139562b122838a7c7fcae3"></a>

## waf_signatures — waf_signatures / 02ba784e0ee8 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- waf_signatures

<a id="canonical-7ff695956cfab12063a7891a9595a839c9b28064deb96d7610db5b7d001a0c74"></a>

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

<a id="canonical-b2d6aa2d6526ebf637e9642f6dd3f7bad59daf97e09d8b69ccc679a0b6fe4437"></a>

## Direct properties — waf_signatures / 02ba784e0ee8 / 3

- [automatic](resources--securemesh_site--reference--group-004.md#canonical-5856774581eba496bb8cf5d8e81df92cc1d3250c77d12b458235ae003a678d87): complete subsection reference.

- [manual](resources--securemesh_site--reference--group-004.md#canonical-31bad4e88791915c267b1d34d4dfefed3bbc3f9d1da7d3d2e276cb5b5804f3df): complete subsection reference.

<a id="canonical-d2adef1e551be6c125ef13bb7440cbaaf84f9ddc1f47f62b7b76a8b7d58c54f2"></a>

## Next pages — waf_signatures / 02ba784e0ee8 / 4

- [waf_signatures.automatic](resources--securemesh_site--reference--group-004.md#canonical-5856774581eba496bb8cf5d8e81df92cc1d3250c77d12b458235ae003a678d87)
- [waf_signatures.manual](resources--securemesh_site--reference--group-004.md#canonical-31bad4e88791915c267b1d34d4dfefed3bbc3f9d1da7d3d2e276cb5b5804f3df)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-5856774581eba496bb8cf5d8e81df92cc1d3250c77d12b458235ae003a678d87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-820def00503891384caf88e1da887b08bf5a7e906a1476ae443fc71e4e8b33a3"></a>

## waf_signatures.automatic — waf_signatures.automatic / 6955595a8cb1 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [waf_signatures](resources--securemesh_site--reference--group-004.md#canonical-e3d3b3dd04a2677e5757751a4242bcaa7ace22229d58fc2ed27e711cfcaac7e0)
- waf_signatures.automatic

<a id="canonical-e0a2e0fb03242afe6f34e284dd4fa2b494afe790205f6342d14cacb30cad20b8"></a>

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

<a id="canonical-c1bcb39a0ee05f9268ba427e0da62355a34036674301ebb1d9507412d78cd56b"></a>

## Direct properties — waf_signatures.automatic / 6955595a8cb1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-15695132e00ae19168a30e49179c3d96aeb47b9350cbe570b587370aa5e13003"></a>

## Next pages — waf_signatures.automatic / 6955595a8cb1 / 4

- [waf_signatures](resources--securemesh_site--reference--group-004.md#canonical-e3d3b3dd04a2677e5757751a4242bcaa7ace22229d58fc2ed27e711cfcaac7e0)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)

<a id="canonical-31bad4e88791915c267b1d34d4dfefed3bbc3f9d1da7d3d2e276cb5b5804f3df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d57dc8702977b2b6ea9e2ffecef0e165505985879831f524cc2db021a803c64b"></a>

## waf_signatures.manual — waf_signatures.manual / 20b05069f3c6 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
- [Property reference](resources--securemesh_site--reference--group-001.md#canonical-2ce4986a55070676c0c3f96d500bff087a583e8c305caafef2f055a9f2c0ec9a)
- [waf_signatures](resources--securemesh_site--reference--group-004.md#canonical-e3d3b3dd04a2677e5757751a4242bcaa7ace22229d58fc2ed27e711cfcaac7e0)
- waf_signatures.manual

<a id="canonical-9f796897ea6a52714a7aa563a691507ef56f3439e483504798b73026c23541aa"></a>

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

<a id="canonical-73d8e299730aa7ad7ffc8f035405ecb9dd8427e4b0154e3072e47dee9d0c129f"></a>

## Direct properties — waf_signatures.manual / 20b05069f3c6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-afa38ca1f26b57cf929164a4f7a6fda0cead6679f55ee89a75f6a51f94b5f331"></a>

## Next pages — waf_signatures.manual / 20b05069f3c6 / 4

- [waf_signatures](resources--securemesh_site--reference--group-004.md#canonical-e3d3b3dd04a2677e5757751a4242bcaa7ace22229d58fc2ed27e711cfcaac7e0)
- [xcsh_securemesh_site](../resources/securemesh_site.md#canonical-29910c0e7bc405230381f804594785cf63014342be390afce858ee8c458a45bc)
