---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-459a153586e0328997a0f476b21589c96e852c71fe40cac7bc843e03273b2dff"></a>

## segment_vrf.segment_config.static_v6_routes — segment_vrf.segment_config.static_v6_routes / 7b90c0473233 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- segment_vrf.segment_config.static_v6_routes

<a id="canonical-1ea077a4e39324aeeb534f27843957c643e3954bce39b75a06b4e4e9b21d5965"></a>

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

<a id="canonical-00979f6c8c48b3d946d2711d58ad2ff9ff919b0967ac875ad88348606e18e6c4"></a>

## Direct properties — segment_vrf.segment_config.static_v6_routes / 7b90c0473233 / 3

- [static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-845ce022853e13601ebc2ee4af16ef7ebb4f01aade2d7d8ea9ebe39fcad50976): complete subsection reference.

<a id="canonical-859c1304e281d2ef6027cafd2bf9d2da4336331f95f1d42a03f1eaa31960887b"></a>

## Next pages — segment_vrf.segment_config.static_v6_routes / 7b90c0473233 / 4

- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-845ce022853e13601ebc2ee4af16ef7ebb4f01aade2d7d8ea9ebe39fcad50976)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-845ce022853e13601ebc2ee4af16ef7ebb4f01aade2d7d8ea9ebe39fcad50976"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8d2f00947e7ba8fd3b56ca0a3675cc154dce132bd82f6b1c9e608783b032933"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes — segment_vrf.segment_config.static_v6_routes.static_routes / 8334ab3939be / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-cf1073d409f03914b2f9656b033c34eed3eb86ced851ca580d443da4e6df778a)
- segment_vrf.segment_config.static_v6_routes.static_routes

<a id="canonical-f5c0317fd5592d689b5239dadd0b301e9d677d27497ad1bfaa0b46b92167263a"></a>

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

<a id="canonical-16350865a6fb004a951981392f1cf1fccd2305a13b293104796d0be05caa42ae"></a>

## Direct properties — segment_vrf.segment_config.static_v6_routes.static_routes / 8334ab3939be / 3

<a id="canonical-cf964a2b52082c780f9d9be11be4e0a68d526b545a5e503acce831c57cf14066"></a>

<a id="canonical-c72da668dd00499bcc49f349ee7b075bee3693cb2cb3896d883d4bb0cf80a14a"></a>

## attrs property — segment_vrf.segment_config.static_v6_routes.static_routes / 8334ab3939be / 4

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

- [default_gateway](resources--securemesh_site_v2--reference--group-017.md#canonical-d9ddc2e9361e0cb36a3837e0fa3f931e54a8e84a593024d6c2e113a6c31c9da8): complete subsection reference.

<a id="canonical-ff95c9bae79b0e0cffbb74d0acdec806061031681d5c940948b13cbff2be6c03"></a>

<a id="canonical-ef42f9704675125d58adc5f5153617baeb1bf6ba99e6e7748aad934ac61bf7ea"></a>

## ip_address property — segment_vrf.segment_config.static_v6_routes.static_routes / 8334ab3939be / 5

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

<a id="canonical-19c6de42a2120400b80f908d7c3fa6da4e896ec781bd57b81757c34a7bc60a07"></a>

<a id="canonical-fc8c790c7e758240e9e6f5aa3999fb73132b2c26701b7581df82a47dd317da99"></a>

## ip_prefixes property — segment_vrf.segment_config.static_v6_routes.static_routes / 8334ab3939be / 6

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

- [node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-a5bb1d4af7bd53bc15053fecadbdd8e17781d679635027afbc108c4a0e2527af): complete subsection reference.

<a id="canonical-4a2540074a18131f9e388e0cdcd976b1d6ea44cc0f4a56367a909d33bea2ccfd"></a>

## Next pages — segment_vrf.segment_config.static_v6_routes.static_routes / 8334ab3939be / 7

- [segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway](resources--securemesh_site_v2--reference--group-017.md#canonical-d9ddc2e9361e0cb36a3837e0fa3f931e54a8e84a593024d6c2e113a6c31c9da8)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-a5bb1d4af7bd53bc15053fecadbdd8e17781d679635027afbc108c4a0e2527af)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-cf1073d409f03914b2f9656b033c34eed3eb86ced851ca580d443da4e6df778a)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d9ddc2e9361e0cb36a3837e0fa3f931e54a8e84a593024d6c2e113a6c31c9da8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95edfe02b1d293e1eb09410ceb8334227be28cd35a799d6055dcf369694160ab"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway — segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway / 13e401cb2c7e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-cf1073d409f03914b2f9656b033c34eed3eb86ced851ca580d443da4e6df778a)
- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-845ce022853e13601ebc2ee4af16ef7ebb4f01aade2d7d8ea9ebe39fcad50976)
- segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-7ffcc7c39a0c03730aa7f000a2afae70b3106f55317fff302e85b174a3112e0f"></a>

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

<a id="canonical-df4171e478310b7b00e4751c9bf73c5d9984a371fdee21a0c46becd2ba886c46"></a>

## Direct properties — segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway / 13e401cb2c7e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bd243f8ac4f250a765a3717bd6759ee750dfc5aa9a0cbea2bbc165ccd1d47ed5"></a>

## Next pages — segment_vrf.segment_config.static_v6_routes.static_routes.default_gateway / 13e401cb2c7e / 4

- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-845ce022853e13601ebc2ee4af16ef7ebb4f01aade2d7d8ea9ebe39fcad50976)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-a5bb1d4af7bd53bc15053fecadbdd8e17781d679635027afbc108c4a0e2527af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a26d9790cfc8409c0b969ef75f94a8bf976841adf98d0a5643217106a68a0f1"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.node_interface — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface / 219f75ebd235 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-cf1073d409f03914b2f9656b033c34eed3eb86ced851ca580d443da4e6df778a)
- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-845ce022853e13601ebc2ee4af16ef7ebb4f01aade2d7d8ea9ebe39fcad50976)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface

<a id="canonical-e75d7ff3b6f1bea53d3f62fc368ff2b0a64e40a2469af81bf2277b49f26b9121"></a>

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

<a id="canonical-54080ed6994222a5a92ee3b1e0a0a4fa5e7a4ce1d0490bbd3274c265b1acb34b"></a>

## Direct properties — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface / 219f75ebd235 / 3

- [list](resources--securemesh_site_v2--reference--group-017.md#canonical-d296d2fa7df1bac959a2c0fa80392edbdc4accb98c8448020fe50fb99bb3ea8d): complete subsection reference.

<a id="canonical-23ff5b791f17afaa7595d72b7d8492bba0ff5cfe39e69f003b7bd8265ecf853d"></a>

## Next pages — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface / 219f75ebd235 / 4

- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-017.md#canonical-d296d2fa7df1bac959a2c0fa80392edbdc4accb98c8448020fe50fb99bb3ea8d)
- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-845ce022853e13601ebc2ee4af16ef7ebb4f01aade2d7d8ea9ebe39fcad50976)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d296d2fa7df1bac959a2c0fa80392edbdc4accb98c8448020fe50fb99bb3ea8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d84e4afa935f7338cba3fecaf934fbf3901c7724bb10264faf4bbee4c7ae1766"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list / 73d6685dff5a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-cf1073d409f03914b2f9656b033c34eed3eb86ced851ca580d443da4e6df778a)
- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-845ce022853e13601ebc2ee4af16ef7ebb4f01aade2d7d8ea9ebe39fcad50976)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-a5bb1d4af7bd53bc15053fecadbdd8e17781d679635027afbc108c4a0e2527af)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-8b73992a46365a057a3f4fdd5b813899172bddc24afcae5f86c67444fd886083"></a>

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

<a id="canonical-6ff57637dd1eee34f0f370d49d8afd70093e808c8e69908f3a5bbe10358379db"></a>

## Direct properties — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list / 73d6685dff5a / 3

- [interface](resources--securemesh_site_v2--reference--group-017.md#canonical-231e4a23c5923201b3c80bf878eecc88fc857e7a8f9f777345e0ef18b2e8f5a6): complete subsection reference.

<a id="canonical-d942cbfdb5687507da4e0f9fc6159d2e27e37b5cd8ddd8a04783e01d529693d1"></a>

<a id="canonical-1bb5d9fd0d88b399428e5be327a920ce727c970ddf0e4b7ba45cb43cd5fa667b"></a>

## node property — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list / 73d6685dff5a / 4

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

<a id="canonical-3d07f61f31f01585675105af5765f98357a6fa6cc477a3c36973ffbea614af50"></a>

## Next pages — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list / 73d6685dff5a / 5

- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface](resources--securemesh_site_v2--reference--group-017.md#canonical-231e4a23c5923201b3c80bf878eecc88fc857e7a8f9f777345e0ef18b2e8f5a6)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-a5bb1d4af7bd53bc15053fecadbdd8e17781d679635027afbc108c4a0e2527af)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-231e4a23c5923201b3c80bf878eecc88fc857e7a8f9f777345e0ef18b2e8f5a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f91b00e206f3c48e96298123d43b917de125c2bbaf85a50d422bf0a5cd4a7f0f"></a>

## segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ffd79031e64f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-9a2c70076000966c30333cfebd9d38f45c2c2a57741d1c073d374ed697323109)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-cf1073d409f03914b2f9656b033c34eed3eb86ced851ca580d443da4e6df778a)
- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-017.md#canonical-845ce022853e13601ebc2ee4af16ef7ebb4f01aade2d7d8ea9ebe39fcad50976)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-a5bb1d4af7bd53bc15053fecadbdd8e17781d679635027afbc108c4a0e2527af)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-017.md#canonical-d296d2fa7df1bac959a2c0fa80392edbdc4accb98c8448020fe50fb99bb3ea8d)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-fa2d6021b75d58743707ad75180fa5aff3739ea000c258c5169e426574e2fc8e"></a>

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

<a id="canonical-7429d64e23d599880922376c3559e98e2af2e91894b3b1d64f2f2938e44181c9"></a>

## Direct properties — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ffd79031e64f / 3

<a id="canonical-90bfdd7114ff5fe6a2ed91a9a7ff558abe8790a46df0ff44da204bc305897144"></a>

<a id="canonical-148a07dd8cd7a5a7cc44a5e97439a02307b208ee443dfaf468706f156f795663"></a>

## kind property — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ffd79031e64f / 4

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

<a id="canonical-444789e1fd1901cde4ccd218ba1ed78ecb2d2c6a2f8ea65042b2d58d22ca7dd8"></a>

<a id="canonical-730530bd9876db858c3edff1bde3a289fa6f74d17c26214dc1fefd5c791db467"></a>

## name property — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ffd79031e64f / 5

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

<a id="canonical-4e7292a49609becf244b9338309de299e3096473685aa66c7d8a62bc5db01e8e"></a>

<a id="canonical-e28aabaed33f8364fcc2f546a77c7af0359e4fb48fd1f49ad06baf537241ba19"></a>

## namespace property — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ffd79031e64f / 6

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

<a id="canonical-53e6f9f01abbf2f351d3854e301244c3e727d118571ad3909ea46ae1b4ac4239"></a>

<a id="canonical-085a5b4faa74ff77b52f367d0c53f6dd58f1815bb69c01bc623c4f6980b1c75f"></a>

## tenant property — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ffd79031e64f / 7

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

<a id="canonical-fc35428f33149f4f9a4cf9ae97018ed43479c22fab8e39596840779bc901defc"></a>

<a id="canonical-b413b923be5a6d12e80002e468ed40cc9557f2dfed55e9965c1d72f7bec75982"></a>

## uid property — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ffd79031e64f / 8

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

<a id="canonical-f8192ff56ef186e729f0414b65b4d8709055c65c0d6f6961a782565234f180c1"></a>

## Next pages — segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.in / ffd79031e64f / 9

- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-017.md#canonical-d296d2fa7df1bac959a2c0fa80392edbdc4accb98c8448020fe50fb99bb3ea8d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d49d241ba738275d8158a3a45f825bae3e3f5fd7e47bcac432b0148e376762a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e36177d95ae5826dec45ea64b8f812d319bbc5aec0c2683a250d7541eda6c4c6"></a>

## segment_vrf.segment_network — segment_vrf.segment_network / 9cbe7332ab71 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- segment_vrf.segment_network

<a id="canonical-f4103cad288441b0f8a0660dc067a6a886a1d67dd33230d0b0a8eab314f1814d"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a 'direct reference' from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name for public API and Uid for private API This type of
reference is called direct because the relation is explicit and concrete (as opposed to selector..

Upstream description:

This type establishes a 'direct reference' from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name for public API and Uid for private API This
type of reference is called direct because the relation is explicit and concrete (as opposed to
selector reference which builds a group based on labels of selectee objects)

Terraform syntax:

```terraform
segment_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-6e40e6c3a09ff6967555883c9f2100aed14326745264a8b2c7a57967e2689765"></a>

## Direct properties — segment_vrf.segment_network / 9cbe7332ab71 / 3

<a id="canonical-8f89e1f3478be36d79aba542f03e76a5bab0bbe4f920e0f781dd1c8893476269"></a>

<a id="canonical-05f833edfe6e9f54bac5b4657bd742c9bdf20adcb86335afe8545a247322a085"></a>

## kind property — segment_vrf.segment_network / 9cbe7332ab71 / 4

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

<a id="canonical-39dd730cc48d6e61edaccf70ab3366a4675d5483700f5ca0838730bf62b72d2c"></a>

<a id="canonical-209c5c411f8904eabcd1acc48c262d82803902b872775342fb5a6430ef9c0b3a"></a>

## name property — segment_vrf.segment_network / 9cbe7332ab71 / 5

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

<a id="canonical-d1d3740c638325bbe5062e91ce5d8fecd8aa57c95e2e25a1c19a073b6e343a6c"></a>

<a id="canonical-39846f3947dc843e38643252afe4d7e5400eb62dc49513994f3874ff9348b1b2"></a>

## namespace property — segment_vrf.segment_network / 9cbe7332ab71 / 6

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

<a id="canonical-a4a4e3f65114d82ff352005a89b5b54bfa206468be50d9d64115ae03e4ec56fe"></a>

<a id="canonical-a33e2136ca2275cf8bcdc54bf82e4505f2e237e8d245ec82fcb3a4c0e7685dec"></a>

## tenant property — segment_vrf.segment_network / 9cbe7332ab71 / 7

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

<a id="canonical-98224935ae4f3f496f43726e2f1dab2d8d11756ed9b9343df468429f1d148d2e"></a>

<a id="canonical-ee136260a77ca1887125071176a1881b0c9cdaff21c973e4ec47ec610efc7b94"></a>

## uid property — segment_vrf.segment_network / 9cbe7332ab71 / 8

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

<a id="canonical-79b8bc3d1de2f17db2c17c894c4c1e7d75842862c28df672950d5d75fe01973e"></a>

## Next pages — segment_vrf.segment_network / 9cbe7332ab71 / 9

- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-873cd77da7deb220967cdf22b3e2e21067959e739c1197b600c388aa7d2c11f4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5ee0c85a2a68677e24ea62314ec318ffe24b5ba2c088b8ed4c462dd230662c35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a8e70144c809b94a600ee8208a4b6b38b29ccf7007c8f7439d18092f1323291"></a>

## site_mesh_group_on_slo — site_mesh_group_on_slo / dd301b66f8bd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- site_mesh_group_on_slo

<a id="canonical-e6abf615ac8d994b377dc1eeb4c7b4806a5812d0f2dae6d0e52b5656fbd005ad"></a>

Type: `"object"`. single nested block, Optional.

Select how the site mesh group will be connected. By default, public IPs of the control nodes of the
site will be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_site_mesh_group",
    "site_mesh_group"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-site_mesh_group_choice": "[\"no_site_mesh_group\",\"site_mesh_group\"]",
  "x-ves-oneof-field-site_mesh_group_ip_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

Terraform syntax:

```terraform
site_mesh_group_on_slo {
  # Configure direct properties listed below.
}
```

<a id="canonical-5dc871be60eb717bbacde68cab253c0bf80636a28acf5925d50358cbd5ecd0cd"></a>

## Direct properties — site_mesh_group_on_slo / dd301b66f8bd / 3

- [no_site_mesh_group](resources--securemesh_site_v2--reference--group-017.md#canonical-e71a9428d0279919a14005994413e87b9dab19e42c172574a398a51dec243254): complete subsection reference.

- [site_mesh_group](resources--securemesh_site_v2--reference--group-017.md#canonical-29aef7019fc5cd9de268cd125cc0a0b168ca7789c7761e4f37a5a5068bd76182): complete subsection reference.

- [sm_connection_public_ip](resources--securemesh_site_v2--reference--group-017.md#canonical-adf72efc279b1fd6cfd090f91e7da58ccd7a327ca7ccfbd3ca6b3528506ca14f): complete subsection reference.

- [sm_connection_pvt_ip](resources--securemesh_site_v2--reference--group-017.md#canonical-1e6e4746375e5a1dbfceb9dfd13aaa50c301f8fe52207ac8c05885a138a9ce1c): complete subsection reference.

<a id="canonical-47a2789cdb981e1176a15f8da2f8f9df1dff577bfdb7e64d1a0a3f36e1a19a96"></a>

## Next pages — site_mesh_group_on_slo / dd301b66f8bd / 4

- [site_mesh_group_on_slo.no_site_mesh_group](resources--securemesh_site_v2--reference--group-017.md#canonical-e71a9428d0279919a14005994413e87b9dab19e42c172574a398a51dec243254)
- [site_mesh_group_on_slo.site_mesh_group](resources--securemesh_site_v2--reference--group-017.md#canonical-29aef7019fc5cd9de268cd125cc0a0b168ca7789c7761e4f37a5a5068bd76182)
- [site_mesh_group_on_slo.sm_connection_public_ip](resources--securemesh_site_v2--reference--group-017.md#canonical-adf72efc279b1fd6cfd090f91e7da58ccd7a327ca7ccfbd3ca6b3528506ca14f)
- [site_mesh_group_on_slo.sm_connection_pvt_ip](resources--securemesh_site_v2--reference--group-017.md#canonical-1e6e4746375e5a1dbfceb9dfd13aaa50c301f8fe52207ac8c05885a138a9ce1c)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-e71a9428d0279919a14005994413e87b9dab19e42c172574a398a51dec243254"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-701489214b8ee57b35673d1d58ec72f801886673f735738a103ba74b5ba70557"></a>

## site_mesh_group_on_slo.no_site_mesh_group — site_mesh_group_on_slo.no_site_mesh_group / 62d16e6cde32 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-5ee0c85a2a68677e24ea62314ec318ffe24b5ba2c088b8ed4c462dd230662c35)
- site_mesh_group_on_slo.no_site_mesh_group

<a id="canonical-c6e009e5a6e1f64aa97f1da5fe7b443b15c35132335719241f096bc5b4d47a8d"></a>

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
no_site_mesh_group = {}
```

<a id="canonical-d4ea9a6b77130dfe8732f8a56672033d3cd5679153369d97268ab323ee6b0543"></a>

## Direct properties — site_mesh_group_on_slo.no_site_mesh_group / 62d16e6cde32 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f0cbb9495c62b6e7c17a1268116ab373351cfa5b21bc5339ab072870e777f171"></a>

## Next pages — site_mesh_group_on_slo.no_site_mesh_group / 62d16e6cde32 / 4

- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-5ee0c85a2a68677e24ea62314ec318ffe24b5ba2c088b8ed4c462dd230662c35)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-29aef7019fc5cd9de268cd125cc0a0b168ca7789c7761e4f37a5a5068bd76182"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-997297d633bea80c0d1842a5b8672c7f58b2cb0eb76988e3e796846a7b744d77"></a>

## site_mesh_group_on_slo.site_mesh_group — site_mesh_group_on_slo.site_mesh_group / f4a84d9f2ff2 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-5ee0c85a2a68677e24ea62314ec318ffe24b5ba2c088b8ed4c462dd230662c35)
- site_mesh_group_on_slo.site_mesh_group

<a id="canonical-e3351943d33a9e709abb15d9df1caef0679f2a006337e00156d1fd8fad92f299"></a>

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
site_mesh_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-590c432be2b8586e2a8c172bcd3c04d89008d9a0494fd85c48de8b518772fc79"></a>

## Direct properties — site_mesh_group_on_slo.site_mesh_group / f4a84d9f2ff2 / 3

<a id="canonical-ab4983001d4e1c1a4fc00377229b1466516dce09ca18ea45de640f6a4fa04842"></a>

<a id="canonical-6581517179315bea686823939da042da28a6c07cf3aeca6c2be688d7e40d0fa2"></a>

## name property — site_mesh_group_on_slo.site_mesh_group / f4a84d9f2ff2 / 4

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

<a id="canonical-99a4b6731a55b3169a22440c451d0451409e7d073c6bc9f81bf028a5aaa658f8"></a>

<a id="canonical-b5e8ce928e35c25f739504b3c0ef5079e9d0914445256b56f044ec33ebb636f8"></a>

## namespace property — site_mesh_group_on_slo.site_mesh_group / f4a84d9f2ff2 / 5

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

<a id="canonical-254d631355b272ba31e714769551a619ce3e5a09fa3788393f55dcf461f02942"></a>

<a id="canonical-82e8bffe3b08d797b660cc0d33130d454b96b62d5ab7492c1d22b26581486195"></a>

## tenant property — site_mesh_group_on_slo.site_mesh_group / f4a84d9f2ff2 / 6

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

<a id="canonical-d34d99e7c623b73587ca42c8f33e9e0adc4eb7afd772e98883743743e2bb336f"></a>

## Next pages — site_mesh_group_on_slo.site_mesh_group / f4a84d9f2ff2 / 7

- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-5ee0c85a2a68677e24ea62314ec318ffe24b5ba2c088b8ed4c462dd230662c35)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-adf72efc279b1fd6cfd090f91e7da58ccd7a327ca7ccfbd3ca6b3528506ca14f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c34b769cb82940c1f70268dccf39826002676f48eb4abbdb536c3b7ba469e287"></a>

## site_mesh_group_on_slo.sm_connection_public_ip — site_mesh_group_on_slo.sm_connection_public_ip / 8fb96501d59a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-5ee0c85a2a68677e24ea62314ec318ffe24b5ba2c088b8ed4c462dd230662c35)
- site_mesh_group_on_slo.sm_connection_public_ip

<a id="canonical-787f08685bd7dc4e129a3f41b1efb011e571580f1dcc0cd5abd3671d524953bf"></a>

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

<a id="canonical-bd4a7b8ae236da7219198779eb4bcecae9ec865b5d0eb546d33645924883732f"></a>

## Direct properties — site_mesh_group_on_slo.sm_connection_public_ip / 8fb96501d59a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4f0913f35cbd30752df3b14352585f2d2a0cedf97f11036ec24c2a7cac127c01"></a>

## Next pages — site_mesh_group_on_slo.sm_connection_public_ip / 8fb96501d59a / 4

- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-5ee0c85a2a68677e24ea62314ec318ffe24b5ba2c088b8ed4c462dd230662c35)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1e6e4746375e5a1dbfceb9dfd13aaa50c301f8fe52207ac8c05885a138a9ce1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-164abb8609b02c13facc2b1b916bab056c1753990d87c02db24c2eeb4b9172ff"></a>

## site_mesh_group_on_slo.sm_connection_pvt_ip — site_mesh_group_on_slo.sm_connection_pvt_ip / 580a87a784cd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-5ee0c85a2a68677e24ea62314ec318ffe24b5ba2c088b8ed4c462dd230662c35)
- site_mesh_group_on_slo.sm_connection_pvt_ip

<a id="canonical-e21d65dfa4e1744f28107392bc92eebb6502e3b10d709682ebde4ed289abe8fb"></a>

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

<a id="canonical-bb81e8347deffcdda697c987d67a26b80b23e35fbd9391f82124a44cbb7ea10c"></a>

## Direct properties — site_mesh_group_on_slo.sm_connection_pvt_ip / 580a87a784cd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2379030b2e6cad2e638c63ea11511a8d6380cad13719c2020172487701013614"></a>

## Next pages — site_mesh_group_on_slo.sm_connection_pvt_ip / 580a87a784cd / 4

- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-5ee0c85a2a68677e24ea62314ec318ffe24b5ba2c088b8ed4c462dd230662c35)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0ab73d55383bba9a5301602766550fe795ca7248f7723978f17cfa74ba089f1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6e50868c334dad396f4121b7ca45c4614985f0f5d32767678319de09cf7ac4d"></a>

## software_settings — software_settings / acdba1463b8f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- software_settings

<a id="canonical-89cec147da5ccf4f5d02145c6b0d9e2c21c89d51bddd25d24e2949882fed7e5d"></a>

Type: `"object"`. single nested block, Optional.

Select OS and Software version for the site. All nodes in the site will run the same OS and Software
version. These settings cannot be changed after the site is created. This block is a create-only,
write-only input; changing it replaces the resource, and refresh preserves the configured value
without claiming XC observed it.

Upstream description:

Select OS and Software version for the site. All nodes in the site will run the same OS and Software
version. These settings cannot be changed after the site is created.

Receipt-pinned upstream constraints:

```json
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
software_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-9c55901381a4b19737d9c2ba6f98b2394ee180076f783a7e60786e24611e7b85"></a>

## Direct properties — software_settings / acdba1463b8f / 3

- [os](resources--securemesh_site_v2--reference--group-017.md#canonical-3affa0bace6c77c9719dcdae4a013518bb250df4fc4aa084bcdce8f0e3fde9f2): complete subsection reference.

- [sw](resources--securemesh_site_v2--reference--group-017.md#canonical-4a1403e1dcc8d0f377f130c5e8ddb4e410074f2e8b98c7929224e371f5bd3370): complete subsection reference.

- [waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-041ed37c5980030a99afd02e01479bdfc6bb40ae5b38f07a8172d03aafc57e8d): complete subsection reference.

<a id="canonical-962968672d25737fba0b160a638542bebd589b7f45b60f1926069a5b25e556a7"></a>

## Next pages — software_settings / acdba1463b8f / 4

- [software_settings.os](resources--securemesh_site_v2--reference--group-017.md#canonical-3affa0bace6c77c9719dcdae4a013518bb250df4fc4aa084bcdce8f0e3fde9f2)
- [software_settings.sw](resources--securemesh_site_v2--reference--group-017.md#canonical-4a1403e1dcc8d0f377f130c5e8ddb4e410074f2e8b98c7929224e371f5bd3370)
- [software_settings.waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-041ed37c5980030a99afd02e01479bdfc6bb40ae5b38f07a8172d03aafc57e8d)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-3affa0bace6c77c9719dcdae4a013518bb250df4fc4aa084bcdce8f0e3fde9f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af27d48fb105e1a345591c87ad0bb1f511ee142275231256803792af3a2e8aa3"></a>

## software_settings.os — software_settings.os / f1a245c1a4ab / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0ab73d55383bba9a5301602766550fe795ca7248f7723978f17cfa74ba089f1f)
- software_settings.os

<a id="canonical-74a1155d92515879d53acbe4be953a928b7f92ed6c269abb3e61269397cbcc78"></a>

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

<a id="canonical-11354d5ea0912d6ba0d5324f0b721474f83c3ace084110762298e023152664f3"></a>

## Direct properties — software_settings.os / f1a245c1a4ab / 3

- [default_os_version](resources--securemesh_site_v2--reference--group-017.md#canonical-4ca1a2b26e2d95082d044efd7f60633ef1d135186c46660c87cdcfe69dbdc16e): complete subsection reference.

<a id="canonical-d415bc7acbf53ab4a1e9b458460ad84a69a3f19bc41a77fb1b20942c419e2547"></a>

<a id="canonical-c17b7f1c93e3192cc2d5ddef3076767c231bf1c01883373ac1eaeec082789b22"></a>

## operating_system_version property — software_settings.os / f1a245c1a4ab / 4

Type: `"string"`. Optional, Sensitive.

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

<a id="canonical-06a9516cf442650a7e09e5db03f1681db7237671737cfb4dc1c600758c071b50"></a>

## Next pages — software_settings.os / f1a245c1a4ab / 5

- [software_settings.os.default_os_version](resources--securemesh_site_v2--reference--group-017.md#canonical-4ca1a2b26e2d95082d044efd7f60633ef1d135186c46660c87cdcfe69dbdc16e)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0ab73d55383bba9a5301602766550fe795ca7248f7723978f17cfa74ba089f1f)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-4ca1a2b26e2d95082d044efd7f60633ef1d135186c46660c87cdcfe69dbdc16e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9edde72a8a118c8b132f93b9a944a39598fcd527c385fa69133cf320db7edff1"></a>

## software_settings.os.default_os_version — software_settings.os.default_os_version / 08dc5d612ccc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0ab73d55383bba9a5301602766550fe795ca7248f7723978f17cfa74ba089f1f)
- [software_settings.os](resources--securemesh_site_v2--reference--group-017.md#canonical-3affa0bace6c77c9719dcdae4a013518bb250df4fc4aa084bcdce8f0e3fde9f2)
- software_settings.os.default_os_version

<a id="canonical-d55029a412ebd93c41848fac8d3728c253a5897b374d6d31bfc7defa1f72aa76"></a>

Type: `["object", {}]`. Optional, Sensitive.

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

<a id="canonical-2283f7523bf59126f208d7b5f28f6f5dfbfb83b03c2f19923f3dfc88189cb87a"></a>

## Direct properties — software_settings.os.default_os_version / 08dc5d612ccc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7072b96ed68c574a9a3670404c7531925add75835361b86f40e176b5098280f3"></a>

## Next pages — software_settings.os.default_os_version / 08dc5d612ccc / 4

- [software_settings.os](resources--securemesh_site_v2--reference--group-017.md#canonical-3affa0bace6c77c9719dcdae4a013518bb250df4fc4aa084bcdce8f0e3fde9f2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-4a1403e1dcc8d0f377f130c5e8ddb4e410074f2e8b98c7929224e371f5bd3370"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c6d5a517ce8b9985b5fca4f0060503f941cc04db08c182dd5a5dbd2b9e3e5de"></a>

## software_settings.sw — software_settings.sw / b72f1cc0903c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0ab73d55383bba9a5301602766550fe795ca7248f7723978f17cfa74ba089f1f)
- software_settings.sw

<a id="canonical-545e72f752d5e531fb25cd8590fe56bd983ebb1f4dbe5d4f58b95b2897f211ad"></a>

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

<a id="canonical-ab03d3296506736cd86910a2ac3ee11f0d304ee6a198bb2b716ecf120a6dac9a"></a>

## Direct properties — software_settings.sw / b72f1cc0903c / 3

- [default_sw_version](resources--securemesh_site_v2--reference--group-017.md#canonical-b9c630ba3d5aefffff94f63681da51ab481c75cc099b2cb81ae365890dc94843): complete subsection reference.

<a id="canonical-a4ec068aede2f9324bca520e9dbdb32f6802279a8d4ccbd50faf144853abfba2"></a>

<a id="canonical-7f766a24b208564200b28155731904eb00dedd4941eb15063f5295884729ecbe"></a>

## volterra_software_version property — software_settings.sw / b72f1cc0903c / 4

Type: `"string"`. Optional, Sensitive.

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

<a id="canonical-efea593ccd539779e391e20a9ffa5fc3c919916cac6f3c4ac78f8147524a7dc9"></a>

## Next pages — software_settings.sw / b72f1cc0903c / 5

- [software_settings.sw.default_sw_version](resources--securemesh_site_v2--reference--group-017.md#canonical-b9c630ba3d5aefffff94f63681da51ab481c75cc099b2cb81ae365890dc94843)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0ab73d55383bba9a5301602766550fe795ca7248f7723978f17cfa74ba089f1f)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-b9c630ba3d5aefffff94f63681da51ab481c75cc099b2cb81ae365890dc94843"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43f8d4497153ed2656a8cc7f18358bf1f32a12ec9b502ed82f1643cf2e2e37a6"></a>

## software_settings.sw.default_sw_version — software_settings.sw.default_sw_version / f84394ee0f96 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0ab73d55383bba9a5301602766550fe795ca7248f7723978f17cfa74ba089f1f)
- [software_settings.sw](resources--securemesh_site_v2--reference--group-017.md#canonical-4a1403e1dcc8d0f377f130c5e8ddb4e410074f2e8b98c7929224e371f5bd3370)
- software_settings.sw.default_sw_version

<a id="canonical-e3a027f268ffa619961bdcb8a041b52c193c56f498dc5824947b2f23cea2377f"></a>

Type: `["object", {}]`. Optional, Sensitive.

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

<a id="canonical-b530d9d1239c0e1ce5e20537d7760c4fbc29322bf054c2d73291514d5d14db33"></a>

## Direct properties — software_settings.sw.default_sw_version / f84394ee0f96 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f874ee2d9b48368fd17d8f4578b1ac7f2f56a928c052b970050569fef14b45cd"></a>

## Next pages — software_settings.sw.default_sw_version / f84394ee0f96 / 4

- [software_settings.sw](resources--securemesh_site_v2--reference--group-017.md#canonical-4a1403e1dcc8d0f377f130c5e8ddb4e410074f2e8b98c7929224e371f5bd3370)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-041ed37c5980030a99afd02e01479bdfc6bb40ae5b38f07a8172d03aafc57e8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64bcff187a99b014a1f18950ea4a7163e4088f94aafaa6227c5314ed7b3fe19b"></a>

## software_settings.waf_signatures — software_settings.waf_signatures / 1d768399a923 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0ab73d55383bba9a5301602766550fe795ca7248f7723978f17cfa74ba089f1f)
- software_settings.waf_signatures

<a id="canonical-9bfe5868498c642e29e6c275d4bfce94e0f809bc47091838bf7a6b2987f6fe22"></a>

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

<a id="canonical-2cb6704f4e8ffb35f5de054b7ffa13d0d0851b9aa6d93e788b4de2e5a7eb5e22"></a>

## Direct properties — software_settings.waf_signatures / 1d768399a923 / 3

- [automatic](resources--securemesh_site_v2--reference--group-017.md#canonical-79593c9e35f03606e2a8b692aaabfe7bb7b6078e0147d08f562a6d56c4d158ef): complete subsection reference.

- [manual](resources--securemesh_site_v2--reference--group-017.md#canonical-7cea338b9839a051a598d5985a362c3196c7b7e311dd05d379610c05cecd154b): complete subsection reference.

<a id="canonical-1dafabdc7a36f98ae065ef78efc2cd805de7e4311d15ee4b58b62b2a59d3d8b3"></a>

## Next pages — software_settings.waf_signatures / 1d768399a923 / 4

- [software_settings.waf_signatures.automatic](resources--securemesh_site_v2--reference--group-017.md#canonical-79593c9e35f03606e2a8b692aaabfe7bb7b6078e0147d08f562a6d56c4d158ef)
- [software_settings.waf_signatures.manual](resources--securemesh_site_v2--reference--group-017.md#canonical-7cea338b9839a051a598d5985a362c3196c7b7e311dd05d379610c05cecd154b)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0ab73d55383bba9a5301602766550fe795ca7248f7723978f17cfa74ba089f1f)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-79593c9e35f03606e2a8b692aaabfe7bb7b6078e0147d08f562a6d56c4d158ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86db2475d3d67e5a79e5db07916db4cbb0a6ac3f2b71f93c2a4250a7c55189e4"></a>

## software_settings.waf_signatures.automatic — software_settings.waf_signatures.automatic / 40aac5b165c7 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0ab73d55383bba9a5301602766550fe795ca7248f7723978f17cfa74ba089f1f)
- [software_settings.waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-041ed37c5980030a99afd02e01479bdfc6bb40ae5b38f07a8172d03aafc57e8d)
- software_settings.waf_signatures.automatic

<a id="canonical-263d4659a98cd582a8ef58d40c305740d547b5992ebcef1ff44645de2babc8c0"></a>

Type: `["object", {}]`. Optional, Sensitive.

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

<a id="canonical-844aea3b06debd48580cfe923d86c576946363d45cd4c67eeb76309693614116"></a>

## Direct properties — software_settings.waf_signatures.automatic / 40aac5b165c7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-805c2b30634961d6ea1b94d700e6b044618288300620aec16c9917d33a1cbe89"></a>

## Next pages — software_settings.waf_signatures.automatic / 40aac5b165c7 / 4

- [software_settings.waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-041ed37c5980030a99afd02e01479bdfc6bb40ae5b38f07a8172d03aafc57e8d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-7cea338b9839a051a598d5985a362c3196c7b7e311dd05d379610c05cecd154b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3616da3266bef98e638c51228697a1edac80366f163fb38b14025095162458f4"></a>

## software_settings.waf_signatures.manual — software_settings.waf_signatures.manual / 5ae6774d5741 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0ab73d55383bba9a5301602766550fe795ca7248f7723978f17cfa74ba089f1f)
- [software_settings.waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-041ed37c5980030a99afd02e01479bdfc6bb40ae5b38f07a8172d03aafc57e8d)
- software_settings.waf_signatures.manual

<a id="canonical-e9fa9a51126794f8e55630b108a934c726e9d58b92d50b64eb829aacf99c16ce"></a>

Type: `["object", {}]`. Optional, Sensitive.

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

<a id="canonical-f054a54c2ba60b2325df18ae8b3a5a2241317d67d292eec9afd590a63bfd68a7"></a>

## Direct properties — software_settings.waf_signatures.manual / 5ae6774d5741 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f7f51c9f06ad8e3381180ff6949c131a2fbdd9783a0f8be9644555a702b206cf"></a>

## Next pages — software_settings.waf_signatures.manual / 5ae6774d5741 / 4

- [software_settings.waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-041ed37c5980030a99afd02e01479bdfc6bb40ae5b38f07a8172d03aafc57e8d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9d7e81a82ada0db8be76c91fbb07dbb2361aa24927ad606494ddbb6ef6e4e50c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89e511cb1787d1396bde7c6352688f1771855fda104ac660522b4a599b55037d"></a>

## timeouts — timeouts / f44037eb8e66 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- timeouts

<a id="canonical-1bb9d07005efec6a7dc0d5c577c4cd03513e9809f8d7624f02268b67c445b92b"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-a484a1cac4451127398c928533257f0cab61885b28689d5678362c75908c4099"></a>

## Direct properties — timeouts / f44037eb8e66 / 3

<a id="canonical-542de2d6504ee603af65abe1d2e4c6d9958428f10429962c79767d4c127f6c24"></a>

<a id="canonical-99b85e62a32f6b16f3b16aa96c9fa510f412913890ed0711ebd877ae57f8c1eb"></a>

## create property — timeouts / f44037eb8e66 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-dc7c52beac6db01089f3d99e5a4f45315294f85d4a9fe0bbce04c5c09108944e"></a>

<a id="canonical-a35cd026166f37d26b6b7196d7de031099b7962872b1fe50183bb2b81af2b87e"></a>

## delete property — timeouts / f44037eb8e66 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-a71dfc1188de0e20c62a55e5af95c03da7704deaac6d89b8ed94a03b42adbb7a"></a>

<a id="canonical-1eb537aedabb6e86b21c330664e29f2c8038f2c3c37b7bd38d110e52a5530593"></a>

## read property — timeouts / f44037eb8e66 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-03dc99efa80fc552048290b6feb250aa41471db18c7334dba30da300b30c2ec9"></a>

<a id="canonical-cc7238c28e2393ad9668996863dff7a22ff7dfc5e3907f1204f428f7e2a027ec"></a>

## update property — timeouts / f44037eb8e66 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-58b02a0e478db5febf69848d623b3cd49e2ce50ad6f8e6c4d850b8db0852a11b"></a>

## Next pages — timeouts / f44037eb8e66 / 8

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9a35c15998ab830d8e4f6a16d90aab3abe7cb293b94d66d487897fffa3e40f6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a4cc15c5c36b8d9b46d8525d38cf75af30d0d1381a1c4de68f87186687c721e"></a>

## upgrade_settings — upgrade_settings / 7024dc227938 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- upgrade_settings

<a id="canonical-46c1c2c8d94c5685e0b5514d5ec8e1914ab70d357a01f8a5fe5dd4ebc88ac802"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for upgrade settings.

Upstream description:

Specify how a site will be upgraded.

Receipt-pinned upstream constraints:

```json
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
upgrade_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-bb7b4abcc8a14dc47cb65be074763172d8a51b17d50f49467fdc6bc590d14c7b"></a>

## Direct properties — upgrade_settings / 7024dc227938 / 3

- [kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-c387dccd020176a132a44c675133002a9b9d0d0c0f87270854c82931da4ffb30): complete subsection reference.

<a id="canonical-a760e337a79aa29970cd7b0701b725e7d3b5f5dc2ff35f7de659ccaf3b146ab5"></a>

## Next pages — upgrade_settings / 7024dc227938 / 4

- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-c387dccd020176a132a44c675133002a9b9d0d0c0f87270854c82931da4ffb30)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c387dccd020176a132a44c675133002a9b9d0d0c0f87270854c82931da4ffb30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11e6015056dcf91e82d7c32aad3611432051c2e1ed5554bfe649ab79e8b785c2"></a>

## upgrade_settings.kubernetes_upgrade_drain — upgrade_settings.kubernetes_upgrade_drain / 72425ee0620c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-9a35c15998ab830d8e4f6a16d90aab3abe7cb293b94d66d487897fffa3e40f6b)
- upgrade_settings.kubernetes_upgrade_drain

<a id="canonical-e877f607a2e8a7404012a4e787ab0e049a905048c14d989713aa5eae5908d4ba"></a>

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

<a id="canonical-a49337a86fc0a0714c7f907bd6ded17366fceb574bb9e7da0e6dd29278e323ef"></a>

## Direct properties — upgrade_settings.kubernetes_upgrade_drain / 72425ee0620c / 3

- [disable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-ea737a849780023cb6b89bd77c3e8ee4a0e51c2dcebd990b58b165ac006e2933): complete subsection reference.

- [enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-2724871f77a611994c2a4bcb163cd697ce7938f83446905b4145e8f5892f09fd): complete subsection reference.

<a id="canonical-aa080bfeeb4139b6fa765a1d9e22079e123e7b1317bdcea252b62b38df53a80c"></a>

## Next pages — upgrade_settings.kubernetes_upgrade_drain / 72425ee0620c / 4

- [upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-ea737a849780023cb6b89bd77c3e8ee4a0e51c2dcebd990b58b165ac006e2933)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-2724871f77a611994c2a4bcb163cd697ce7938f83446905b4145e8f5892f09fd)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-9a35c15998ab830d8e4f6a16d90aab3abe7cb293b94d66d487897fffa3e40f6b)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ea737a849780023cb6b89bd77c3e8ee4a0e51c2dcebd990b58b165ac006e2933"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa5089e32eec46f997bdb40cf55f943a500c0d3ea01cdcf70fae9a31187f898b"></a>

## upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain — upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain / 1d6fcf301d7e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-9a35c15998ab830d8e4f6a16d90aab3abe7cb293b94d66d487897fffa3e40f6b)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-c387dccd020176a132a44c675133002a9b9d0d0c0f87270854c82931da4ffb30)
- upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-b91972ab47bfaa4e972d732c818e6687852e33fe2ab372c36375bc34989bddb3"></a>

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

<a id="canonical-f098f971f6f56681604f1f3c049843170b594d325b103f85d0fb6aacf74bff85"></a>

## Direct properties — upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain / 1d6fcf301d7e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a1af84a8695e051ffc633ec5a9972d2bd7f5ff39cc7711eec042bcc8d785e395"></a>

## Next pages — upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain / 1d6fcf301d7e / 4

- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-c387dccd020176a132a44c675133002a9b9d0d0c0f87270854c82931da4ffb30)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-2724871f77a611994c2a4bcb163cd697ce7938f83446905b4145e8f5892f09fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f34329cec828817bca29fccbfa6abb1e35fe5ab72cdb4d1f4f607aa23f8fb3c"></a>

## upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain / 5b579014a9d9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-9a35c15998ab830d8e4f6a16d90aab3abe7cb293b94d66d487897fffa3e40f6b)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-c387dccd020176a132a44c675133002a9b9d0d0c0f87270854c82931da4ffb30)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-6c43c0bb14681f9a498e7b49e98345234fc1526006f5287a69a8c9773d04bbc5"></a>

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

<a id="canonical-b030a6c7329eec5ca0bb6fbb056eb1603326eb45bc6987a1fe6144192cc33c60"></a>

## Direct properties — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain / 5b579014a9d9 / 3

- [disable_vega_upgrade_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-d621b272e30ab4e8b1a42ea3db95e4524c401f2c8fef1442ce7cd93528ec4322): complete subsection reference.

<a id="canonical-3173fbd5e89143fb558b957f0d998e79c9664dc225c0140bc8c516c65c6f805e"></a>

<a id="canonical-cced08edbdf24f822572a6a4c83197974ca5866d018d2c0aba01d6c771a9a03c"></a>

## drain_max_unavailable_node_count property — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain / 5b579014a9d9 / 4

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

<a id="canonical-8e902f027c0b978486a5da4773113ac7fd9ea96cff26e4e6884796db22339f38"></a>

<a id="canonical-fe983992c33ae89237fa5c0d1c58a01fec1b0302caa6f55a30746b1bef1d62d1"></a>

## drain_max_unavailable_node_percentage property — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain / 5b579014a9d9 / 5

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-f8390bb6cdd6429e2adca2a53d9d2b0c304e48d8d06953040dc3b0bad1bf607d"></a>

<a id="canonical-e6301fb1a6d0e707689a3e4638fa2ce377e75aff2c6a45d4eeca5a0deb9e30e7"></a>

## drain_node_timeout property — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain / 5b579014a9d9 / 6

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

- [enable_vega_upgrade_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-9c449138e2d5b869eda6b0e64f30b1918967e5d5af76cfbcdd0a42dc3901dbd0): complete subsection reference.

<a id="canonical-fed79858caaf0ce816827e5d0b6057e7f69a793e5deefea7be94a7d4f9d60861"></a>

## Next pages — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain / 5b579014a9d9 / 7

- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-d621b272e30ab4e8b1a42ea3db95e4524c401f2c8fef1442ce7cd93528ec4322)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-9c449138e2d5b869eda6b0e64f30b1918967e5d5af76cfbcdd0a42dc3901dbd0)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-c387dccd020176a132a44c675133002a9b9d0d0c0f87270854c82931da4ffb30)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d621b272e30ab4e8b1a42ea3db95e4524c401f2c8fef1442ce7cd93528ec4322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-839495766b05d07f3b176a6d218ef8450f80ecd06a9f01b494eff855f3831726"></a>

## upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgr / 8f639e256d89 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-9a35c15998ab830d8e4f6a16d90aab3abe7cb293b94d66d487897fffa3e40f6b)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-c387dccd020176a132a44c675133002a9b9d0d0c0f87270854c82931da4ffb30)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-2724871f77a611994c2a4bcb163cd697ce7938f83446905b4145e8f5892f09fd)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-d228d6437b1ecf4faa9d2314834decfd36cafc6757029564e1138f410951c22c"></a>

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

<a id="canonical-50a1c7085a2c7d70b10e55b97c3b70a76163c18f301081f1d38f89340218935f"></a>

## Direct properties — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgr / 8f639e256d89 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa8f1ccd52cf10c2bf48d7aabf7168453d046316634adc27c11e5718c789cc38"></a>

## Next pages — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgr / 8f639e256d89 / 4

- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-2724871f77a611994c2a4bcb163cd697ce7938f83446905b4145e8f5892f09fd)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9c449138e2d5b869eda6b0e64f30b1918967e5d5af76cfbcdd0a42dc3901dbd0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbf6c8af54348750db0ce55568f693623aecafc8a223e7e1a498edbf0d08a31d"></a>

## upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgra / ff5cbc5f6619 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-9a35c15998ab830d8e4f6a16d90aab3abe7cb293b94d66d487897fffa3e40f6b)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-c387dccd020176a132a44c675133002a9b9d0d0c0f87270854c82931da4ffb30)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-2724871f77a611994c2a4bcb163cd697ce7938f83446905b4145e8f5892f09fd)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-473f16627633d7476acfba3002302b9f1d05f9633a502d25dee8c7166fe368f2"></a>

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

<a id="canonical-ae5f5bffa9ee12f64ed159638ea2649b32067edc68b153d1ffc1d5a62fe601e0"></a>

## Direct properties — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgra / ff5cbc5f6619 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-88f5688f073c6bc596972b208c2d5626b53f5662515089a1a2884b30665dc0a5"></a>

## Next pages — upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgra / ff5cbc5f6619 / 4

- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-2724871f77a611994c2a4bcb163cd697ce7938f83446905b4145e8f5892f09fd)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-213f5ef616cf5aba47be4049f3b7e5389a567630ef8ae2a212d391aecdd1adb6"></a>

## vmware — vmware / e51c6bca8393 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- vmware

<a id="canonical-0767dfd11436e8d87e6f047b27f708f3d2eaf8501176857aa7b008cd5f5ae7ac"></a>

Type: `"object"`. single nested block, Optional.

VMware Provider Type. VMware Provider Type.

Upstream description:

VMware Provider Type.

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
vmware {
  # Configure direct properties listed below.
}
```

<a id="canonical-664fdfe611b48454742f64f554ab0ebfb172cb134b82f8f68b841098f7c41b8c"></a>

## Direct properties — vmware / e51c6bca8393 / 3

- [not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f): complete subsection reference.

<a id="canonical-fb353fc71bdf36575063a07f2c4343d10399fc5860dfb17b01461d0bfd156244"></a>

## Next pages — vmware / e51c6bca8393 / 4

- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e5087d64ce90b423bd1a91bda101d14c0f22027fa5d56e9790d05e549e6da15"></a>

## vmware.not_managed — vmware.not_managed / d99a48572d56 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- vmware.not_managed

<a id="canonical-e6e0eea7a5a664510c5afb9ad9b6bf6455be98b8701065df39bde427e58dd27d"></a>

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

<a id="canonical-bf80b014ec724ddcebb812b612f4a60c9a7e3c645db36fdd678313a638159cd0"></a>

## Direct properties — vmware.not_managed / d99a48572d56 / 3

- [node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3): complete subsection reference.

<a id="canonical-7c48ab34569ea924a56dd838d930cf360bc57bbc34ffa75340a59a94d644cad7"></a>

## Next pages — vmware.not_managed / d99a48572d56 / 4

- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-115ef8b096f81e2625d2fd2d36dbd77730e83e94aee4d7879f738904b7abc26a"></a>

## vmware.not_managed.node_list — vmware.not_managed.node_list / 955740fbe824 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- vmware.not_managed.node_list

<a id="canonical-d5d122ab592e0dccd70cf586c90b49d179656c6367ea711a9f2daf55bb5bddf7"></a>

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

<a id="canonical-667991d8ec0b6ad5fa0b642f7e8071fe7e9b3f00712f0a22b6f6f95ab24338c2"></a>

## Direct properties — vmware.not_managed.node_list / 955740fbe824 / 3

<a id="canonical-9e7116c10bf1a300ba8edfb68f4b61696efa16c7d4ae985de4c324a3e05770d1"></a>

<a id="canonical-3dcc382123125bf1bc5ca4cef8faa89509b5f860625437c04f3730a6806c9b5d"></a>

## hostname property — vmware.not_managed.node_list / 955740fbe824 / 4

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

- [interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5): complete subsection reference.

<a id="canonical-4dc53b7e7027e098a55aa5a0e5c73b848d57d7464db95ba06aad070b48582b9f"></a>

<a id="canonical-e32b7a200bae4c8ad86252a7e82726f317b03de1835344dec5424de4692f2d1b"></a>

## public_ip property — vmware.not_managed.node_list / 955740fbe824 / 5

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

<a id="canonical-f82b8ccd73386e7f4d12d29294cea959b5f42292126c96eb1a4de2a670a89008"></a>

<a id="canonical-a85a19dda72564c98454adf417ed2d626598b459f00cfdcd162441e77349d695"></a>

## type property — vmware.not_managed.node_list / 955740fbe824 / 6

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

<a id="canonical-c107946ac988d857f88a403bf180290e06ab9379da3dce7b314efe1482339f60"></a>

## Next pages — vmware.not_managed.node_list / 955740fbe824 / 7

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-336c5e368648eedcf305499dc4038cc63f04ad2af151b9d7a79440a693015bcd"></a>

## vmware.not_managed.node_list.interface_list — vmware.not_managed.node_list.interface_list / cc898528172c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- vmware.not_managed.node_list.interface_list

<a id="canonical-34265f8337b75be88d3b01a9223bd0356304400a13edb081db7d0d3423c7a95f"></a>

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

<a id="canonical-593ea77cec36f54a5c800e68933fc579607fb8c7112e33585ff4cda51e25c2ed"></a>

## Direct properties — vmware.not_managed.node_list.interface_list / cc898528172c / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-fa0212e02a884ac9787883b18ed631499dd1a3f80bb0f1ba692e2edeecdd9d30): complete subsection reference.

<a id="canonical-a3c1f786a70e794cd1f15aa85260bb5cc788c0a5df1223c60293f936968ad580"></a>

<a id="canonical-772ff475a3398c3a2add1b8171cb4e7ad8b0e3693e18a692ba1eacdbefbdb423"></a>

## description_spec property — vmware.not_managed.node_list.interface_list / cc898528172c / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-017.md#canonical-c3990a634f865d7b66b6b9a0d44b5a364a2b6be61107aa3dacf618944f87462b): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-017.md#canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-728b11da29730bc0b5e34c0970782488096a8d6f8760d7c9142ea82fcb480756): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f): complete subsection reference.

<a id="canonical-96df16bcee328cce3c6cbd57b56b261e6a288e338457031546204c596a5bcd21"></a>

<a id="canonical-97d9a0e0785819f30b10d4746d1328dcdd3792c41c3d8666e7ffb69dc05409e7"></a>

## is_management property — vmware.not_managed.node_list.interface_list / cc898528172c / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-cffe0d4476b19b4fde6fba18ea7ef82360f41fe849274e5298a34627df6999f7"></a>

<a id="canonical-98cd0aeba144339cf3266a94e20e4108b4112f69594e8c32b04c7208d1ebc9c9"></a>

## is_primary property — vmware.not_managed.node_list.interface_list / cc898528172c / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-e999ef0d4808321bcd0b705e4dcb58d69bf4a04248f0ca18f8366978c5f8a19b"></a>

<a id="canonical-a16a6260cbfec4f1204a23254bf1df8841cf91c103eb4b40d028255c3c05ef22"></a>

## labels property — vmware.not_managed.node_list.interface_list / cc898528172c / 7

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

- [monitor](resources--securemesh_site_v2--reference--group-018.md#canonical-9b0b544b6e519fd2f2fd04ceb1d404d8c1fa9f1e1e5f095d06e30599522e0ebc): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-018.md#canonical-9f32a06af7c948d6bb7e5d21fa67f7084d73a977512950f8ffa1bbb8cb23579c): complete subsection reference.

<a id="canonical-4c8a27cee3ab9c4a26f8d4c966b641f22b8b732921b1b142570c546699eb699b"></a>

<a id="canonical-f7f389ff8dafeb5bd7847d668d1d34e8b878c0b4bb17cbc572d653f26a80dea9"></a>

## mtu property — vmware.not_managed.node_list.interface_list / cc898528172c / 8

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

<a id="canonical-90ea16fedf68ee638c80dd02c89ccbd2e19d7139a17f16842825a55113f1db2a"></a>

<a id="canonical-78071b6289bb5ef0ac646d4736d65983ca450545a1344fd0f28b02133c40e023"></a>

## name property — vmware.not_managed.node_list.interface_list / cc898528172c / 9

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

- [network_option](resources--securemesh_site_v2--reference--group-018.md#canonical-5ebfba2a52831d712df58de767013f1f5a79056918e8170689ab47792a430ebf): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-018.md#canonical-40a0319a571432be435379ef7835965d96a611c1ab4e5ed201e4e13825a3a7d6): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-28b9acf8bfb7a6aa850449e054ed55aa389f96f6e827218c26e7a0f14894fd36): complete subsection reference.

<a id="canonical-6dc120e8eb8352ef476d5411f7ff51cb4eba2ddb3ec4cd5bc059b7d09cd692f0"></a>

<a id="canonical-a51af755b69b8ded50ed33368c9611623ec16b619bed5298f61554ef6085d617"></a>

## priority property — vmware.not_managed.node_list.interface_list / cc898528172c / 10

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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-018.md#canonical-ba07765ef645996ba29d90fea43545c54e0b60bee7de81bfbfc90b950ce32925): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-018.md#canonical-85a9dab358e3ed5eac5cc0be0abf30d385589cd8cd221b8cd36a13a57b4da4a2): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-018.md#canonical-f2d28f83ef14c14c7dfd32288a76cbab391c55e2ec73d753f0208ae3dd28ab80): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-4de96dd703c5f5036299256506315b1ffb4e0bd083b71d7fa30914ebd2b15d00): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-1a243e1484d9cde82529035943a86f47d148392a171c0b3e9bb0e156e7d96c89): complete subsection reference.

<a id="canonical-7b1191c2df6ad1060cb3a41d3cf53cf3ad6d1189dbe26787270a57e6bb673606"></a>

## Next pages — vmware.not_managed.node_list.interface_list / cc898528172c / 11

- [vmware.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-fa0212e02a884ac9787883b18ed631499dd1a3f80bb0f1ba692e2edeecdd9d30)
- [vmware.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-017.md#canonical-c3990a634f865d7b66b6b9a0d44b5a364a2b6be61107aa3dacf618944f87462b)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-017.md#canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b)
- [vmware.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-728b11da29730bc0b5e34c0970782488096a8d6f8760d7c9142ea82fcb480756)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- [vmware.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-018.md#canonical-9b0b544b6e519fd2f2fd04ceb1d404d8c1fa9f1e1e5f095d06e30599522e0ebc)
- [vmware.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-018.md#canonical-9f32a06af7c948d6bb7e5d21fa67f7084d73a977512950f8ffa1bbb8cb23579c)
- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-018.md#canonical-5ebfba2a52831d712df58de767013f1f5a79056918e8170689ab47792a430ebf)
- [vmware.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-018.md#canonical-40a0319a571432be435379ef7835965d96a611c1ab4e5ed201e4e13825a3a7d6)
- [vmware.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-28b9acf8bfb7a6aa850449e054ed55aa389f96f6e827218c26e7a0f14894fd36)
- [vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-018.md#canonical-ba07765ef645996ba29d90fea43545c54e0b60bee7de81bfbfc90b950ce32925)
- [vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-018.md#canonical-85a9dab358e3ed5eac5cc0be0abf30d385589cd8cd221b8cd36a13a57b4da4a2)
- [vmware.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-018.md#canonical-f2d28f83ef14c14c7dfd32288a76cbab391c55e2ec73d753f0208ae3dd28ab80)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-4de96dd703c5f5036299256506315b1ffb4e0bd083b71d7fa30914ebd2b15d00)
- [vmware.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-1a243e1484d9cde82529035943a86f47d148392a171c0b3e9bb0e156e7d96c89)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-fa0212e02a884ac9787883b18ed631499dd1a3f80bb0f1ba692e2edeecdd9d30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5e1d2f6039ae054a2d0e46902d6acd582180e05a6aa0724ac5a44937e0b200a"></a>

## vmware.not_managed.node_list.interface_list.bond_interface — vmware.not_managed.node_list.interface_list.bond_interface / f695b96e87b9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.bond_interface

<a id="canonical-d580fea7a519ec7814175cb958c6b3b913fa5133c00ee79b893dd37860ac37b8"></a>

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

<a id="canonical-15e7ef810439458b9c3efd9049aad8890f0f545fdb087a7ad842d03a7f5026dd"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.bond_interface / f695b96e87b9 / 3

- [active_backup](resources--securemesh_site_v2--reference--group-017.md#canonical-be1f5c94336017489e4282278360ab3acb9db0f5c50058d79919d29fc2db8898): complete subsection reference.

<a id="canonical-a9499b688c2759b29cdbc358561bf20888c2a0698ccd13386dead6cacf136c36"></a>

<a id="canonical-3cde837fe5827acbfdeb830bf295bb57b18a89c42f9c2756093f961ea4dc1e07"></a>

## devices property — vmware.not_managed.node_list.interface_list.bond_interface / f695b96e87b9 / 4

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

- [lacp](resources--securemesh_site_v2--reference--group-017.md#canonical-5c7c981167b150f5998e233a568f263e0aec90277082296cebc39f8a7b5734e4): complete subsection reference.

<a id="canonical-ab2f9c39989823e553cadf3fedf561a3ffaf5cd4282c854ad278f2e77ae14a5b"></a>

<a id="canonical-ddba3e9ae10d9d08bfc3f9bc32a4aab377e3d4b98dff53087b71e540edd2c2f1"></a>

## link_polling_interval property — vmware.not_managed.node_list.interface_list.bond_interface / f695b96e87b9 / 5

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

<a id="canonical-e9bf6949ad9da5a3af47e1247b788d3856c69378bebed80ec13751add6d34308"></a>

<a id="canonical-478a4eb16ec5799f0079b949ee226e92e0732bbf12cff695265a172dc034ae81"></a>

## link_up_delay property — vmware.not_managed.node_list.interface_list.bond_interface / f695b96e87b9 / 6

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

<a id="canonical-029a26ac0a22aabdb2e3a53f7e60f50470d35bcc79e8112369a4d3f275b0970e"></a>

<a id="canonical-01a4e29643cb3545ee346e791117281910f93b616eb8900430344991d9c7c3d0"></a>

## name property — vmware.not_managed.node_list.interface_list.bond_interface / f695b96e87b9 / 7

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

<a id="canonical-810b4f25c3fb10086d4014845b43e05e011e60d74455949f325c12ebae058b52"></a>

## Next pages — vmware.not_managed.node_list.interface_list.bond_interface / f695b96e87b9 / 8

- [vmware.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-017.md#canonical-be1f5c94336017489e4282278360ab3acb9db0f5c50058d79919d29fc2db8898)
- [vmware.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-017.md#canonical-5c7c981167b150f5998e233a568f263e0aec90277082296cebc39f8a7b5734e4)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-be1f5c94336017489e4282278360ab3acb9db0f5c50058d79919d29fc2db8898"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d6417c0194bd83acfc2e1304c12dce3c48c4377f27f5b85061b4f2b3be82382"></a>

## vmware.not_managed.node_list.interface_list.bond_interface.active_backup — vmware.not_managed.node_list.interface_list.bond_interface.active_backup / d048e647c16e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-fa0212e02a884ac9787883b18ed631499dd1a3f80bb0f1ba692e2edeecdd9d30)
- vmware.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-cbaf1b0cce55c268a3ae86ff65af42801bdee5d7262c3d14f869f4f1c9729093"></a>

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

<a id="canonical-bd0c636f0cb84a4dbc1e6765cd17bfd9dad7758fae8784720b7983c2f9b21bd0"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.bond_interface.active_backup / d048e647c16e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dde87fb0680705d5b0637c0a912488449d13d6d8104bb0be6029eee9e83bfec5"></a>

## Next pages — vmware.not_managed.node_list.interface_list.bond_interface.active_backup / d048e647c16e / 4

- [vmware.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-fa0212e02a884ac9787883b18ed631499dd1a3f80bb0f1ba692e2edeecdd9d30)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5c7c981167b150f5998e233a568f263e0aec90277082296cebc39f8a7b5734e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1bd34753bb58fa2b319117745f8af1a3b881d011a025c74e91ec8226cc12b8a"></a>

## vmware.not_managed.node_list.interface_list.bond_interface.lacp — vmware.not_managed.node_list.interface_list.bond_interface.lacp / b2fdac8757e5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-fa0212e02a884ac9787883b18ed631499dd1a3f80bb0f1ba692e2edeecdd9d30)
- vmware.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-7182f67bfba017cdf5ba55a4f4cc7f6249f29245e4d5ff2bc29f7b090aba8d87"></a>

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

<a id="canonical-859480dcc1163ef22f7dca112760524a567b58c3994c1bbfb3b4012d83ebd009"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.bond_interface.lacp / b2fdac8757e5 / 3

<a id="canonical-560505d5cfb3eb11910d3d2e39bbb58a09d5327faf26d0fd063a5fae2c51b8ff"></a>

<a id="canonical-4c4defa71e53dae90e4ec8019284b168230edfba4c31894321749c7ed465ee73"></a>

## rate property — vmware.not_managed.node_list.interface_list.bond_interface.lacp / b2fdac8757e5 / 4

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

<a id="canonical-2c62f6e6750a91a5fc7ca2dd8da6f4aec32f8d49a479005243d7ef549a38fb46"></a>

## Next pages — vmware.not_managed.node_list.interface_list.bond_interface.lacp / b2fdac8757e5 / 5

- [vmware.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-017.md#canonical-fa0212e02a884ac9787883b18ed631499dd1a3f80bb0f1ba692e2edeecdd9d30)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c3990a634f865d7b66b6b9a0d44b5a364a2b6be61107aa3dacf618944f87462b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd4988fd40d3709f6a59781994a92b4a9a8ecbd13013085dd32be277b4034384"></a>

## vmware.not_managed.node_list.interface_list.dhcp_client — vmware.not_managed.node_list.interface_list.dhcp_client / 73576b7a69d6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-ab6e73534e3bf74a529c01d68014c0202e15bb7b977100569a5417ad85078302"></a>

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

<a id="canonical-5063dc1ab388f73b7e5e05e203c6e8c1b7edab24e0902ab0c59b18f82556b58e"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_client / 73576b7a69d6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ddcb8f123d4f778f7df15cbea256276362645b8f42f4eb4991ac796f3132551c"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_client / 73576b7a69d6 / 4

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23ff544f118fb15debde6f13c6053e576e9a3aac812348e76507133ede5a079f"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server — vmware.not_managed.node_list.interface_list.dhcp_server / 465d043f70a7 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-c46b417fe501fc45187db6828e0612c75fcab5dafe43575f65b028479c6da7a1"></a>

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

<a id="canonical-39239a259efd0404e0dd4248fc8a3520d8d5e0acbc2c566bf1a5ade2eded41c4"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server / 465d043f70a7 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-017.md#canonical-99d6ee464fccfe07d183951ab2684e78c93d061ebf734231b6eee776ff2935b9): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-017.md#canonical-4eb1226192e194a74198eee33d7465d654bd919bb8177c048d2da028e76a7e04): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-017.md#canonical-ffa33f77ca422d1e37317332c05cd31ba6d541a91502fb5ff91eba0335a54f27): complete subsection reference.

<a id="canonical-e98e504c626c8d9bcef8f34d9ac0795ac04df78b119e3589a8a798dd95ef9c60"></a>

<a id="canonical-b8040d74846d65ee8bdcf192921f5e7e529f086ccb54a9fad3ccff01e481657b"></a>

## dhcp_option82_tag property — vmware.not_managed.node_list.interface_list.dhcp_server / 465d043f70a7 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-e9a7d48b706323f950c12173c07e1492fbd1e6ab1af2efe779f0cac91f6bf6ce"></a>

<a id="canonical-6cfb13236a832d1b33990f847a974752b971eb91ff2561dc6c669d135f73c83a"></a>

## fixed_ip_map property — vmware.not_managed.node_list.interface_list.dhcp_server / 465d043f70a7 / 5

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-017.md#canonical-64e67fe4e8eda0abc4abbd4f7c5cc131a9a6cc368250fae11733b63daac05f7e): complete subsection reference.

<a id="canonical-18a5609e723570fa88c45c623a428471bd648c913b1d43baf1e3c189761f0309"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server / 465d043f70a7 / 6

- [vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-017.md#canonical-99d6ee464fccfe07d183951ab2684e78c93d061ebf734231b6eee776ff2935b9)
- [vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-017.md#canonical-4eb1226192e194a74198eee33d7465d654bd919bb8177c048d2da028e76a7e04)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-017.md#canonical-ffa33f77ca422d1e37317332c05cd31ba6d541a91502fb5ff91eba0335a54f27)
- [vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-017.md#canonical-64e67fe4e8eda0abc4abbd4f7c5cc131a9a6cc368250fae11733b63daac05f7e)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-99d6ee464fccfe07d183951ab2684e78c93d061ebf734231b6eee776ff2935b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd4301ab46cca553de641f997170f6f2178f30fbd22ae89dd591d5055423b5b6"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / cd0836543c81 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-017.md#canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b)
- vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-4c945df63148f353329586734aaf1f806de03b1ba26d2d8c9bda9768ff41be7e"></a>

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

<a id="canonical-a56b094921be0073cf78effbc0f740274e886be2b7aa4a7c83514e47c49ab254"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / cd0836543c81 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-18d6c5a9b16a8f178029a8c3286d8c4113cf2a3be5523452e2066116317928a5"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / cd0836543c81 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-017.md#canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-4eb1226192e194a74198eee33d7465d654bd919bb8177c048d2da028e76a7e04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15ff089095c6b0dd91741316e4b6cfb379cce42ee393b99c1a2a47af0382f42c"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 936cd26302e3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-017.md#canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b)
- vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-f8ee9d52407b009c4f5f7b9293136f96284701b2a9b4007a2fb30e7ab8635669"></a>

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

<a id="canonical-92f5a5549e8a1de17680d03f8b6eb4b3f5900c0502a4c1afef5272a654835ff0"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 936cd26302e3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1bd5e30c48fef1183f0cb61ff30383f9b0345b9f42a537b02c266f829f34dd39"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 936cd26302e3 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-017.md#canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ffa33f77ca422d1e37317332c05cd31ba6d541a91502fb5ff91eba0335a54f27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0569fcee95b6bfeb16b7618902f23acc854dd58a82bc1263d5a3464e4767734b"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 8a59f9dc941d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-017.md#canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-1665fb9ce361079c0dd1c1a03ed896ce8399c4bbee9c0521034f6a2c8add84ef"></a>

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

<a id="canonical-e6d7d4f5fe6909766c9468e2b36200ba5d075464d2652f64dc72d1e963942ad7"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 8a59f9dc941d / 3

<a id="canonical-0a5d212a0ca13e9406cd4c9faa90c9d24081056e5d928d7bbaa671c13c3a8746"></a>

<a id="canonical-2beed02f4f325b9d10b88e3fb5089c46b0e3f31f0df4ca185a1c97919777f948"></a>

## dgw_address property — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 8a59f9dc941d / 4

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

<a id="canonical-2bf1e3f8c9f6ff4646bc8b2fe4dd9fe4255be9207c36e586be4c0d785a4130a8"></a>

<a id="canonical-339e0c7baf28e823843d8fbed34ac3dc536539e3f2be2ec68762dcf75fa24508"></a>

## dns_address property — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 8a59f9dc941d / 5

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

- [first_address](resources--securemesh_site_v2--reference--group-017.md#canonical-60eae457c87c784716d55b869506832cf3f3e0a8e9387000d44a9b1e387e845b): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-017.md#canonical-72b3ba17a8a47f56ef481f17b1761fec28e056e6a6f841262a3c55364846ef29): complete subsection reference.

<a id="canonical-c696bcf35bfbd27ba53084fd2cc758034db90487f43eee2ee73084abee9f1b85"></a>

<a id="canonical-58754b4b79941b799ab7b7f67f9f6f9f8dc117132f5cdefe18c3e8068c49dfa4"></a>

## network_prefix property — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 8a59f9dc941d / 6

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

<a id="canonical-20486c59fb330e9ca2002a5b02b9d2e516c0d4bc508d0c9e80951c2ba159ed05"></a>

<a id="canonical-5515ea3fdbec0d17a01952519645d4cec42e95235e40098da34f94e6af184f50"></a>

## pool_settings property — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 8a59f9dc941d / 7

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

- [pools](resources--securemesh_site_v2--reference--group-017.md#canonical-ed6dbf592ce25f439847bb4210c56c5f80b84990ef847c5b764bb4f1cf427b22): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-017.md#canonical-30d7b7b856e4dce91017479fbf87f52ebfc36bc9dfb5818ea8529afd81f30819): complete subsection reference.

<a id="canonical-386a0dce4c24edf4bb9e1d848483e4d9101b8a6ae7b276bb157155297a5a8298"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 8a59f9dc941d / 8

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-017.md#canonical-60eae457c87c784716d55b869506832cf3f3e0a8e9387000d44a9b1e387e845b)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-017.md#canonical-72b3ba17a8a47f56ef481f17b1761fec28e056e6a6f841262a3c55364846ef29)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-017.md#canonical-ed6dbf592ce25f439847bb4210c56c5f80b84990ef847c5b764bb4f1cf427b22)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-017.md#canonical-30d7b7b856e4dce91017479fbf87f52ebfc36bc9dfb5818ea8529afd81f30819)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-017.md#canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-60eae457c87c784716d55b869506832cf3f3e0a8e9387000d44a9b1e387e845b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5312992ceaa7fe58899c7e6fc837bbfc179fe37f6f3693c6cf5d1c26cdb32077"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_addr / 75e57a01a580 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-017.md#canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-017.md#canonical-ffa33f77ca422d1e37317332c05cd31ba6d541a91502fb5ff91eba0335a54f27)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-cd2e7533dff0b837033367452d698c5370a387536e12e3bb341757de866e113d"></a>

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

<a id="canonical-a91b9b6b7318c1684e4d75f3223582c3159b387a010c171717733fe541c9e62c"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_addr / 75e57a01a580 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-29b27d65dadb70bcd026fb11af557d0da2b29deebc448aeb6a5fbb86cc5f3f4c"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_addr / 75e57a01a580 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-017.md#canonical-ffa33f77ca422d1e37317332c05cd31ba6d541a91502fb5ff91eba0335a54f27)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-72b3ba17a8a47f56ef481f17b1761fec28e056e6a6f841262a3c55364846ef29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42c7ffa8d3c163476c2cc3e88d5f9cf65766579f2deeb7d202a1a2c838623081"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addre / defe4e79eafc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-017.md#canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-017.md#canonical-ffa33f77ca422d1e37317332c05cd31ba6d541a91502fb5ff91eba0335a54f27)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-f5fc3ab0a5dfd9921ba6f73abc3e7e6ce3fa4a6e4a4c69abd9f9e4e23a11edc3"></a>

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

<a id="canonical-653c8fefa798e13536543929d8139a1825c6f1bf5abc789ff7ea0ba117aeb7ea"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addre / defe4e79eafc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fcd7fc7b07acfbde14437ada988638d26762d8b5325b28d5951acf21ebd9aed5"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addre / defe4e79eafc / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-017.md#canonical-ffa33f77ca422d1e37317332c05cd31ba6d541a91502fb5ff91eba0335a54f27)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ed6dbf592ce25f439847bb4210c56c5f80b84990ef847c5b764bb4f1cf427b22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a158d42ba93518bd7453cafa60a5ae3de4689afe66e60448c9ce55fda6cad7b"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / deb3d2fbc30f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-017.md#canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-017.md#canonical-ffa33f77ca422d1e37317332c05cd31ba6d541a91502fb5ff91eba0335a54f27)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-3a995efb555fef7eef22883bff864b4ba77a67ff8ec5096eabbded7edc185cd3"></a>

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

<a id="canonical-e5ee6456059286b9e9e20d2eadfeac264e14e6e8dd46aa4180bc21aa06be17db"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / deb3d2fbc30f / 3

<a id="canonical-d9773a35ec866614310d19a4020ee88a7b6e691d20b42088f8628fe18e65afc8"></a>

<a id="canonical-38b97ec9d4f63350839f9004bc654578585915f87d9fc634389ad56e87520cfc"></a>

## end_ip property — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / deb3d2fbc30f / 4

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

<a id="canonical-fd18278cdd8feddf7eb7cd22f00d5cea4d38fd138f8ff8e0b4a657b863211c5f"></a>

<a id="canonical-8abcabb3251bc5eadc85389f54e38cdf80e56e73764c38e92d2c05a6013e4b61"></a>

## exclude property — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / deb3d2fbc30f / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-3d69b39de724f97f0e88e200c51553c9b0030b0b1c2cfdc748b44f7e6c0b45ec"></a>

<a id="canonical-02c26d6f8a863ae7d34cb601a347012056dffd7ee8a5eefdca9a4f586b11e4e6"></a>

## start_ip property — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / deb3d2fbc30f / 6

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

<a id="canonical-713a73136dc817720b3a6ccdb83caa4426d4cce87272e42e9da5d5c5f292d08b"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / deb3d2fbc30f / 7

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-017.md#canonical-ffa33f77ca422d1e37317332c05cd31ba6d541a91502fb5ff91eba0335a54f27)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-30d7b7b856e4dce91017479fbf87f52ebfc36bc9dfb5818ea8529afd81f30819"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d03677c9935a9b615e88d7ea6c687cfe3de81f2893b5408b00a879fe20ee8dd"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dg / 0c1866c71d52 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-017.md#canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b)
- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-017.md#canonical-ffa33f77ca422d1e37317332c05cd31ba6d541a91502fb5ff91eba0335a54f27)
- vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-3fce4f33cddb9bba0dae9b86f454b4fa023ab627f3d763287436480425018777"></a>

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

<a id="canonical-805a37e43b7437b9f283fd68f86a0bce459937e63f0273f0934be3a19321b300"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dg / 0c1866c71d52 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe45b2eb3a92a3af435442c4c239cd69cc8cef6bced842ecc67fce5afba39c24"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dg / 0c1866c71d52 / 4

- [vmware.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-017.md#canonical-ffa33f77ca422d1e37317332c05cd31ba6d541a91502fb5ff91eba0335a54f27)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-64e67fe4e8eda0abc4abbd4f7c5cc131a9a6cc368250fae11733b63daac05f7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6cde13b1898a07f07b5f8f8d92e58bb2a3b5886cd45b456418c5c9988f2336b"></a>

## vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 8d08b93259eb / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-017.md#canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b)
- vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-0eb5fb05fbc169d0713ad6f0bc7a7196972c3ad9fc8aa01fb35152049a842508"></a>

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

<a id="canonical-a6d797680324502d290daa5048eeef87d91b37ffd3c8d2238328303aa51d0aa6"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 8d08b93259eb / 3

<a id="canonical-d58bd9a6a3dd5c602ee95f7be52735e682af801c53e5ca32a2e072c5a7c54c97"></a>

<a id="canonical-ba246a30b7d475f8d4af1deb08f27346f0179da56ec02055eb01bef04f87c91a"></a>

## interface_ip_map property — vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 8d08b93259eb / 4

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

<a id="canonical-9f8c96ae309e1173b93f90bb4fe81e970f699805707263b757dfb943fc06997e"></a>

## Next pages — vmware.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 8d08b93259eb / 5

- [vmware.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-017.md#canonical-7cae369f6453874ef55b3b47d8f2249ea1bb47adc247e1ce939a80112331926b)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-728b11da29730bc0b5e34c0970782488096a8d6f8760d7c9142ea82fcb480756"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc2f28475567ea4dfdff6595983aaccb84b9afec7d9961568fd7546018883784"></a>

## vmware.not_managed.node_list.interface_list.ethernet_interface — vmware.not_managed.node_list.interface_list.ethernet_interface / 7470c1fabd1b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-f8a2e0c0016b676bbb52ffa70694d5d02555e18bcdd110310d0d927ffb89a191"></a>

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

<a id="canonical-fa08c9bf8b1c9d01a509a8d7fb95450e19bdd3a27191f1fe8623ef0b3b16a6f7"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ethernet_interface / 7470c1fabd1b / 3

<a id="canonical-26374a59fe91325162e77e361d46c45d2a8b9ee255fb9faa573de5fe67aa05e8"></a>

<a id="canonical-63fefc32df236ccc47405acf993efa1e69c96de36a642e98ad1e11e86972e6e5"></a>

## device property — vmware.not_managed.node_list.interface_list.ethernet_interface / 7470c1fabd1b / 4

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

<a id="canonical-6f6e55f6bf0e876499110c950a42ff325fce98e9feefc4adcfca943bd4e5e1c5"></a>

<a id="canonical-16ba83512245fe3646f5a423114a19167fc3d1ba4ece73752dc636d0b9824426"></a>

## mac property — vmware.not_managed.node_list.interface_list.ethernet_interface / 7470c1fabd1b / 5

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

<a id="canonical-e869e5fb23da52144d0900d467df3820d57993468edac673ae02f16972461297"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ethernet_interface / 7470c1fabd1b / 6

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aee5476df9cf1cee0c61cc5be81fc7f2b48f49fdba7ce093820f322066d6cba5"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config — vmware.not_managed.node_list.interface_list.ipv6_auto_config / e1d1b965113e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-119cc6bdf5be47cf674695c641838c1eaa37d61db9e4b927b16a0705f3718c0a"></a>

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

<a id="canonical-6175bb831275f815bd5a369623daffad8d91d73991681018357b4095978d8953"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config / e1d1b965113e / 3

- [host](resources--securemesh_site_v2--reference--group-017.md#canonical-0b447f4e07cf6cea3c4ba246e69a601cbe58d786f749a4bc9b8d3b34e0c22919): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08): complete subsection reference.

<a id="canonical-29a59ed02589d9fafb081acffbc661376cbe7f3cabe1143ac5473127980dfa7d"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config / e1d1b965113e / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-017.md#canonical-0b447f4e07cf6cea3c4ba246e69a601cbe58d786f749a4bc9b8d3b34e0c22919)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-017.md#canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0b447f4e07cf6cea3c4ba246e69a601cbe58d786f749a4bc9b8d3b34e0c22919"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29e786ef72880fc9ae354309c95b2c2125699ee1f3cdc8b1786af24ab31b11e2"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.host — vmware.not_managed.node_list.interface_list.ipv6_auto_config.host / 5ebbb2ba8909 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-58081d8abaad62b95260751d3d89ed5ebe5ff39df6d9398ebce623835c055cf3"></a>

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

<a id="canonical-753abef3b849c02edaa651c804e1a2c2346fa9e2d139cccde45b95df121fa65c"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.host / 5ebbb2ba8909 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-86f36ba2b1d75ea556997c37357b9c382464cbae4fb267f11454b64b4c916381"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.host / 5ebbb2ba8909 / 4

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c5c7b3fe8db397038c37d95ab016e54688593303819651f1927e249b240ebe08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05aecba23514c9b57fb72b94e1fd51384c9d0a6b1da623162d479c73f80c6f54"></a>

## vmware.not_managed.node_list.interface_list.ipv6_auto_config.router — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router / 59d7a4b3916c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-c15f30455bb5de54af2535b836bd1740e2dd9208f108c2c19db6f3f57c04d83c)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-4cdaeabc5bd9766b7f4faad22d2b18078645bba71fad13c059b2550b8a63745f)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-2bfa0f50f635caa32211b0fb36d9c71961b10b8ea87d400854278f9b6f06f7a3)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-405c31a1575c9af8c5ddced60a1d594cde3e536c78ee0fc549f4a2a98c0576a5)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-832feb0c2a11bf10a11b26d48c2085b0049f89267e9126f662c47ea8f986d3dd"></a>

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

<a id="canonical-0f771fc793bc04f4e8261549a4f7dbcaefcfb4418508bbfcfedeed6e93b19c28"></a>

## Direct properties — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router / 59d7a4b3916c / 3

- [dns_config](resources--securemesh_site_v2--reference--group-017.md#canonical-3bc5c5da15314133ff03418fbec7fb825fb63d7b5ba1b72a8c28fceb10bb9739): complete subsection reference.

<a id="canonical-5979d8c3a8a4e96bd474f483c407442d0fc133db615e26051981e3a937618416"></a>

<a id="canonical-eedd5d3981ace18aa4b1a9c8b34cd75a2b9f563bcbfd9dc39176f515e70133e6"></a>

## network_prefix property — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router / 59d7a4b3916c / 4

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

- [stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-7c7d839a74b1dbce9e5eeb5e2072c6578f923d2e53c5ca1cbb463fce794b6d6e): complete subsection reference.

<a id="canonical-a63abeb1b22766f964edd87528106fb5e1791e4c430263d5c941a80c324ff80b"></a>

## Next pages — vmware.not_managed.node_list.interface_list.ipv6_auto_config.router / 59d7a4b3916c / 5

- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-017.md#canonical-3bc5c5da15314133ff03418fbec7fb825fb63d7b5ba1b72a8c28fceb10bb9739)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-018.md#canonical-7c7d839a74b1dbce9e5eeb5e2072c6578f923d2e53c5ca1cbb463fce794b6d6e)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-017.md#canonical-1ca26003927c6657202334c8a528e4b985397ff13caad46c5efcc976b7f7333f)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-3bc5c5da15314133ff03418fbec7fb825fb63d7b5ba1b72a8c28fceb10bb9739"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
