---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-e081d19285c81fe6af7e0a57d9f61dd8ea9d7b954efb9c2b37e9bc02d90e296a"></a>

## f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps — f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps / f2699b73c02a / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.market_place_image](data-sources--nfv_service--reference--group-001.md#canonical-a436a856fd52a1a49e3f7d06f8669b081d25cc482b3c200ec9cda77881219996)
- f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps

<a id="canonical-c89651796244f6551e8b7fa2f0f954726acbbe25ef777ed511ee8b9813d7033b"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for AWAFPayG3Gbps.

<a id="canonical-337639ab3058c9b51041a32cdfb8a5502109cba4b29742b791c78d82bb8bc0fd"></a>

## Direct properties — f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps / f2699b73c02a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0ff3e0f0822b13f225ec81bd8b1cfc7b7f08f5af91b3c743efcf6483c6b950ce"></a>

## Next pages — f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps / f2699b73c02a / 4

- [f5_big_ip_aws_service.market_place_image](data-sources--nfv_service--reference--group-001.md#canonical-a436a856fd52a1a49e3f7d06f8669b081d25cc482b3c200ec9cda77881219996)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-dde503fe987104949100eb9e9885b6d2faf030c7d243f017d7a1775751685bfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1067ae7c6aec3ec8e0ee4b787e16de037c5b340b934318be1a86c68fd399e886"></a>

## f5_big_ip_aws_service.nodes — f5_big_ip_aws_service.nodes / 753a6cdaab10 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- f5_big_ip_aws_service.nodes

<a id="canonical-cfa47f171548c18963f853af07ad82af5443d6f9c53baa9b0270f9fde9cafe9e"></a>

Type: `"list"`. Computed.

Specify how and where the service nodes are spawned.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 2,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 2,
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
    "ves.io.schema.rules.repeated.max_items": "2",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "2",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-a111c4879d5fe7d09dfe13d6015fd00e9e7bb46c7a47a6ad0bb0d121d06027a4"></a>

## Direct properties — f5_big_ip_aws_service.nodes / 753a6cdaab10 / 3

- [automatic_prefix](data-sources--nfv_service--reference--group-002.md#canonical-f1183240d442f07a997a663b288d0fe7d9cdeaa8abb7b3e5824839b57c1a8da3): complete subsection reference.

<a id="canonical-baaeae1a41544b6e265ed4916e90eed4f8156c3c81c56cf4520e6e6ff16ae66d"></a>

<a id="canonical-2d3596bcebb9e622ae61285d771e786e5dccca91501cdd6d4862c792b923348c"></a>

## aws_az_name property — f5_big_ip_aws_service.nodes / 753a6cdaab10 / 4

Type: `"string"`. Computed.

The AWS Availability Zone must be consistent with the AWS Region chosen. Please select an AZ in the
same Region as your TGW Site.

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
    "pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  }
}
```

- [mgmt_subnet](data-sources--nfv_service--reference--group-002.md#canonical-b375fc2e691fe0edb2cf3f6f4bcde66b5ca0df98c30367ca3cc9332cb69df8b4): complete subsection reference.

<a id="canonical-65de6a9a58c40b708e1fa30c47031c84ca734586256aa7a25b284c92d0d18ba9"></a>

<a id="canonical-8565c7ac096d44b8c02360340d49f066b0a931c1427760c5de22e0d0b32327ac"></a>

## node_name property — f5_big_ip_aws_service.nodes / 753a6cdaab10 / 5

Type: `"string"`. Computed.

Node Name will be used to assign as hostname to the service.

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

- [reserved_mgmt_subnet](data-sources--nfv_service--reference--group-002.md#canonical-7dd0808d20cc4c25c8023899e6c77ba1ffa594817b96b0cd24f1013b91a240d6): complete subsection reference.

<a id="canonical-a85099e65fc028729646e684b340c3861670bff168832f2d8c6955ee85cd4ff1"></a>

<a id="canonical-f7c9bd60c4234fa64b3ee82504b1bd0b095866d5cc94f7e18405b785aa886591"></a>

## tunnel_prefix property — f5_big_ip_aws_service.nodes / 753a6cdaab10 / 6

Type: `"string"`. Computed.

Exclusive with \[automatic\_prefix\] Enter IP prefix for the tunnel, it has to be /30.

Upstream description:

Exclusive with \[automatic\_prefix\] Enter IP prefix for the tunnel, it has to be /30.

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

<a id="canonical-f66a4067cb27114cd586a646ac2e7c2ee543e7fac95a842c209079d949f32217"></a>

## Next pages — f5_big_ip_aws_service.nodes / 753a6cdaab10 / 7

- [f5_big_ip_aws_service.nodes.automatic_prefix](data-sources--nfv_service--reference--group-002.md#canonical-f1183240d442f07a997a663b288d0fe7d9cdeaa8abb7b3e5824839b57c1a8da3)
- [f5_big_ip_aws_service.nodes.mgmt_subnet](data-sources--nfv_service--reference--group-002.md#canonical-b375fc2e691fe0edb2cf3f6f4bcde66b5ca0df98c30367ca3cc9332cb69df8b4)
- [f5_big_ip_aws_service.nodes.reserved_mgmt_subnet](data-sources--nfv_service--reference--group-002.md#canonical-7dd0808d20cc4c25c8023899e6c77ba1ffa594817b96b0cd24f1013b91a240d6)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-f1183240d442f07a997a663b288d0fe7d9cdeaa8abb7b3e5824839b57c1a8da3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eaffc562c395cc6fca0275f7f37804e7dfe550d0d4cfedd337f0fec6acf57fe8"></a>

## f5_big_ip_aws_service.nodes.automatic_prefix — f5_big_ip_aws_service.nodes.automatic_prefix / ead1ed0d23b0 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.nodes](data-sources--nfv_service--reference--group-002.md#canonical-dde503fe987104949100eb9e9885b6d2faf030c7d243f017d7a1775751685bfb)
- f5_big_ip_aws_service.nodes.automatic_prefix

<a id="canonical-f1f10f5d63d8b0a87c2bcfaeb7b5ff662fa103ef1cbc919e3bede6f0b0c37d6e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic prefix.

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

<a id="canonical-35c6d5ae88a939131d225c5c62ab560b4be7ea65f5a3a6bc42ff31371240e8d0"></a>

## Direct properties — f5_big_ip_aws_service.nodes.automatic_prefix / ead1ed0d23b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-60491aad21e615184f9565f98c42bd7515234dd0f20d30607e01ebee95edae47"></a>

## Next pages — f5_big_ip_aws_service.nodes.automatic_prefix / ead1ed0d23b0 / 4

- [f5_big_ip_aws_service.nodes](data-sources--nfv_service--reference--group-002.md#canonical-dde503fe987104949100eb9e9885b6d2faf030c7d243f017d7a1775751685bfb)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-b375fc2e691fe0edb2cf3f6f4bcde66b5ca0df98c30367ca3cc9332cb69df8b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b23d9133e5d7d2bc193c32d8070d9bee27831b4dd991c20a841e77cac3251db"></a>

## f5_big_ip_aws_service.nodes.mgmt_subnet — f5_big_ip_aws_service.nodes.mgmt_subnet / 90167e36e723 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.nodes](data-sources--nfv_service--reference--group-002.md#canonical-dde503fe987104949100eb9e9885b6d2faf030c7d243f017d7a1775751685bfb)
- f5_big_ip_aws_service.nodes.mgmt_subnet

<a id="canonical-ee148c512e7de3d604065153a8eeea3f1638d7acba7b4c580ef0c61e5ff0a2f4"></a>

Type: `"single"`. Computed.

Configuration parameter for mgmt subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-f0944f699e95077d89c02c5338286f2701ffb868209fe2f81e0304f25956ec5d"></a>

## Direct properties — f5_big_ip_aws_service.nodes.mgmt_subnet / 90167e36e723 / 3

<a id="canonical-86cf4b294be1a38f18b424bb7cc581b6ab37db5d98bc5a1d0950eca1936be808"></a>

<a id="canonical-04669f066400a946eb72783b0761b55906487f25d8310e4e6b88793c680ecdf7"></a>

## existing_subnet_id property — f5_big_ip_aws_service.nodes.mgmt_subnet / 90167e36e723 / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--nfv_service--reference--group-002.md#canonical-ab613e173ec44f1a136a038259a51ce10ba3b44669a28ab9df932007e6a83b59): complete subsection reference.

<a id="canonical-c3c0c170c7a23674a2878b59f05228a8e73d2ec1c9e88d43b367730964f9a981"></a>

## Next pages — f5_big_ip_aws_service.nodes.mgmt_subnet / 90167e36e723 / 5

- [f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param](data-sources--nfv_service--reference--group-002.md#canonical-ab613e173ec44f1a136a038259a51ce10ba3b44669a28ab9df932007e6a83b59)
- [f5_big_ip_aws_service.nodes](data-sources--nfv_service--reference--group-002.md#canonical-dde503fe987104949100eb9e9885b6d2faf030c7d243f017d7a1775751685bfb)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-ab613e173ec44f1a136a038259a51ce10ba3b44669a28ab9df932007e6a83b59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-166e9d09b5ec3d5f86a882e06c1c5643bd38a1c70a9e3b1d91cc9cedc608626e"></a>

## f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param — f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param / 794ce15eef4a / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.nodes](data-sources--nfv_service--reference--group-002.md#canonical-dde503fe987104949100eb9e9885b6d2faf030c7d243f017d7a1775751685bfb)
- [f5_big_ip_aws_service.nodes.mgmt_subnet](data-sources--nfv_service--reference--group-002.md#canonical-b375fc2e691fe0edb2cf3f6f4bcde66b5ca0df98c30367ca3cc9332cb69df8b4)
- f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param

<a id="canonical-1e0251c39ea60d7d7d3bade7ddd1477a842a09f676cd5a8344fbc60beeb91c33"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a41ca5ed6ba6dfd8591bdf78088960228d12f2412dd2f48e736b1181b8266da1"></a>

## Direct properties — f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param / 794ce15eef4a / 3

<a id="canonical-c6b31d332e8c66c1368d779a010375bd8c16903056c405e27cd4d1b52a3f96c4"></a>

<a id="canonical-b592253f314f022c3b9cd522800d36dc7ccb235b654cbfeb3c429c5291166e06"></a>

## ipv4 property — f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param / 794ce15eef4a / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-40178f1b70a7ebc24924a662f56cd86bdece84404bf154b5347a37d183cd4104"></a>

## Next pages — f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param / 794ce15eef4a / 5

- [f5_big_ip_aws_service.nodes.mgmt_subnet](data-sources--nfv_service--reference--group-002.md#canonical-b375fc2e691fe0edb2cf3f6f4bcde66b5ca0df98c30367ca3cc9332cb69df8b4)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-7dd0808d20cc4c25c8023899e6c77ba1ffa594817b96b0cd24f1013b91a240d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7cf6ee5b366f77b681829062e211c6ec25a70bd44b0a00561ed894f97edbba87"></a>

## f5_big_ip_aws_service.nodes.reserved_mgmt_subnet — f5_big_ip_aws_service.nodes.reserved_mgmt_subnet / cec9aae1393a / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [f5_big_ip_aws_service](data-sources--nfv_service--reference--group-001.md#canonical-d98e2b92f1592fdf93cd0b4279ad5a21d3d5720033f82ea1c9f57d1fa1a74aa9)
- [f5_big_ip_aws_service.nodes](data-sources--nfv_service--reference--group-002.md#canonical-dde503fe987104949100eb9e9885b6d2faf030c7d243f017d7a1775751685bfb)
- f5_big_ip_aws_service.nodes.reserved_mgmt_subnet

<a id="canonical-d8ff73141281b0ba0356a72c537de012c1aeb441a312dbc0a1d8bf0fb28038af"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for reserved mgmt subnet.

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

<a id="canonical-6185b2c648e5488bd2a5ad44809cd83768f2de4f899a298b3ee185cd5bd6dfa8"></a>

## Direct properties — f5_big_ip_aws_service.nodes.reserved_mgmt_subnet / cec9aae1393a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-986bb90ec37c5785e2411e9590987ff7614a84530a43c4c9f151b1ffe4c20839"></a>

## Next pages — f5_big_ip_aws_service.nodes.reserved_mgmt_subnet / cec9aae1393a / 4

- [f5_big_ip_aws_service.nodes](data-sources--nfv_service--reference--group-002.md#canonical-dde503fe987104949100eb9e9885b6d2faf030c7d243f017d7a1775751685bfb)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8dac804e4a81083a173008210e1d1d174d4c632a19f73583099fd0ddaaaf536d"></a>

## https_management — https_management / 98c2fbf8cff9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- https_management

<a id="canonical-fd4a121fbf60553361382d690ccb12ad338f58b82514814fbcb963bede6c6944"></a>

Type: `"single"`. Computed.

Configuration parameter for https management.

Upstream description:

HTTPS based configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_on_internet\",\"advertise_on_internet_default_vip\",\"advertise_on_sli_vip\",\"advertise_on_slo_internet_vip\",\"advertise_on_slo_sli\",\"advertise_on_slo_vip\"]",
  "x-ves-oneof-field-internet_choice": "[]",
  "x-ves-oneof-field-port_choice": "[\"default_https_port\",\"https_port\"]"
}
```

<a id="canonical-02df41b54c0aeca66ce65d7e7fd2aeef9f1f17ab179a237553e74a4da63552a9"></a>

## Direct properties — https_management / 98c2fbf8cff9 / 3

- [advertise_on_internet](data-sources--nfv_service--reference--group-002.md#canonical-45534132ba081c7187375c22b68a63278a8ab7e16f265db31224bbe140b6974f): complete subsection reference.

- [advertise_on_internet_default_vip](data-sources--nfv_service--reference--group-002.md#canonical-63ad95008e007a85138a397c330cbdb8423c4d7ea9c330043c0f192497de2f95): complete subsection reference.

- [advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6): complete subsection reference.

- [advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f): complete subsection reference.

- [advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f): complete subsection reference.

- [advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7): complete subsection reference.

- [default_https_port](data-sources--nfv_service--reference--group-003.md#canonical-1ea6170c2e0093e64ef8c11cfeadbca923ba5e0edcba2b15fbd606e1aa0c37e7): complete subsection reference.

<a id="canonical-1d1281b9d20046e1a8e87ce545f615f3d362388d1b5fd03291ccc08e23496f7c"></a>

<a id="canonical-c5f962bb7c315d381497886d0ef154c57e498ce9bc3d034cd4af4d535f144fa4"></a>

## domain_suffix property — https_management / 98c2fbf8cff9 / 4

Type: `"string"`. Computed.

Domain suffix will be used along with node name to form URL to access node management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-776b0bae7a9d1863d090b2e595ad414c658f2412a5198a6953acceb8671c8e6b"></a>

<a id="canonical-278847076c2b0960841cd6ebb05270755e06cf485ba517268e47de3323a8a0ee"></a>

## https_port property — https_management / 98c2fbf8cff9 / 5

Type: `"number"`. Computed.

Exclusive with \[default\_https\_port\] Enter TCP port number.

Upstream description:

Exclusive with \[default\_https\_port\] Enter TCP port number.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-38ab7c4297f1f785433369be49ab7d300940a222265b286cf957e7d5dabc5e82"></a>

## Next pages — https_management / 98c2fbf8cff9 / 6

- [https_management.advertise_on_internet](data-sources--nfv_service--reference--group-002.md#canonical-45534132ba081c7187375c22b68a63278a8ab7e16f265db31224bbe140b6974f)
- [https_management.advertise_on_internet_default_vip](data-sources--nfv_service--reference--group-002.md#canonical-63ad95008e007a85138a397c330cbdb8423c4d7ea9c330043c0f192497de2f95)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--reference--group-003.md#canonical-9cb7d734443534e78f77f5035f885e992d26930de647b960b4a5a8542d04f2a7)
- [https_management.default_https_port](data-sources--nfv_service--reference--group-003.md#canonical-1ea6170c2e0093e64ef8c11cfeadbca923ba5e0edcba2b15fbd606e1aa0c37e7)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-45534132ba081c7187375c22b68a63278a8ab7e16f265db31224bbe140b6974f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c18493401b50b47ddad84594c3c42c0683aa0f2ce8ca53c39e5ade4060ae2c9"></a>

## https_management.advertise_on_internet — https_management.advertise_on_internet / a7576c6e7f03 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- https_management.advertise_on_internet

<a id="canonical-3a1ad6a2dd30b631ce8c5e2324e81ec5e61920f3c9042b41aebbc5c663dccdd7"></a>

Type: `"single"`. Computed.

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

<a id="canonical-516b582da0be8fc7173a330db39f568b12f92e5b57b7b5e468f84b592f7c9582"></a>

## Direct properties — https_management.advertise_on_internet / a7576c6e7f03 / 3

- [public_ip](data-sources--nfv_service--reference--group-002.md#canonical-c27700bce841b2e3f1f388856d9841e4a87960ce2970994d72c88dcfef737afa): complete subsection reference.

<a id="canonical-ebea1d0759efc1ac39868a6feb319187d3f4cc21cd4496af773a0d8176c4054d"></a>

## Next pages — https_management.advertise_on_internet / a7576c6e7f03 / 4

- [https_management.advertise_on_internet.public_ip](data-sources--nfv_service--reference--group-002.md#canonical-c27700bce841b2e3f1f388856d9841e4a87960ce2970994d72c88dcfef737afa)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-c27700bce841b2e3f1f388856d9841e4a87960ce2970994d72c88dcfef737afa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4bcf76cbfe7e79cc6620c79f15f41090f399709b3899b475306920d6e1443d84"></a>

## https_management.advertise_on_internet.public_ip — https_management.advertise_on_internet.public_ip / 8cfd0dec67d2 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_internet](data-sources--nfv_service--reference--group-002.md#canonical-45534132ba081c7187375c22b68a63278a8ab7e16f265db31224bbe140b6974f)
- https_management.advertise_on_internet.public_ip

<a id="canonical-27f5fab27dbaa703d83d1dd4c29801eb5270f826508d99f24d2a97a3518f46ff"></a>

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

<a id="canonical-7abec123f19c157280f742fd28dc8c0af4ddee97d48bab910282c462fa414706"></a>

## Direct properties — https_management.advertise_on_internet.public_ip / 8cfd0dec67d2 / 3

<a id="canonical-da02994d0e3a88e121b5c4cbcd660c8bba3e5ac2f38e8d6146625c27c26d0d55"></a>

<a id="canonical-44d604f616b9a72239843a0e70c22da00f9ed836120f5dd8ce2f32d1f742e3b9"></a>

## name property — https_management.advertise_on_internet.public_ip / 8cfd0dec67d2 / 4

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

<a id="canonical-21112d4949a2381134be679f140bc0fa4eaf63668402eb418b314d13e0ffd36e"></a>

<a id="canonical-4d608509f614c7da6dce1bb1a1f5bb8b7fb68a77baaea0f886c7520e241bde35"></a>

## namespace property — https_management.advertise_on_internet.public_ip / 8cfd0dec67d2 / 5

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

<a id="canonical-456fa9da8325d2b45fefb1239fe3a7dee9226c68a883d7f96e0d0d504e60ba62"></a>

<a id="canonical-76e36f7971ff780135fe28b480a0664d08e1fbb85b96bdbfaf4b98ba047a337f"></a>

## tenant property — https_management.advertise_on_internet.public_ip / 8cfd0dec67d2 / 6

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

<a id="canonical-e2bcba1a49aaca72cb17dac347c8c44ac3fe07124d1cbcaf92fc3826c1bd359d"></a>

## Next pages — https_management.advertise_on_internet.public_ip / 8cfd0dec67d2 / 7

- [https_management.advertise_on_internet](data-sources--nfv_service--reference--group-002.md#canonical-45534132ba081c7187375c22b68a63278a8ab7e16f265db31224bbe140b6974f)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-63ad95008e007a85138a397c330cbdb8423c4d7ea9c330043c0f192497de2f95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4745b18bc0f18895d2f40721582d71037227ca58cfb92fd1a58ae0ff3f77b72d"></a>

## https_management.advertise_on_internet_default_vip — https_management.advertise_on_internet_default_vip / f400b6ebe444 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- https_management.advertise_on_internet_default_vip

<a id="canonical-623487111cbb318eeb6c751ff782ec75b9dbdcbe8e2b002319b42b6ba8bf7f9c"></a>

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

<a id="canonical-80d720209ff1dd6de52a16b44e1db9f3eef88cd899f725ceca6a1d9bebb0011b"></a>

## Direct properties — https_management.advertise_on_internet_default_vip / f400b6ebe444 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9c07ae5401b60b9ac85209225c96fdc1b64149149f8e20286a461f05b3d42d8c"></a>

## Next pages — https_management.advertise_on_internet_default_vip / f400b6ebe444 / 4

- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8def00ad8dcac3d029a7f67e3654f0411fe79b2bc89be129e8387fdc1113bd1"></a>

## https_management.advertise_on_sli_vip — https_management.advertise_on_sli_vip / 7c4bcaeb0da1 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- https_management.advertise_on_sli_vip

<a id="canonical-617d3e85a39c22b32f2f0ec63beccf7bd21be0e448654b9e36c63fc829299ded"></a>

Type: `"single"`. Computed.

Inline TLS Parameters. Inline TLS parameters.

Upstream description:

Inline TLS parameters.

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

<a id="canonical-de6ab9a792edb1b2f5e85d163918a6d956235e77f74db52da073455613e116d8"></a>

## Direct properties — https_management.advertise_on_sli_vip / 7c4bcaeb0da1 / 3

- [no_mtls](data-sources--nfv_service--reference--group-002.md#canonical-4c078eb10bee4746e319220ccb09f5a19c190cb49d47f8b43502167f1516abb6): complete subsection reference.

- [tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-78b0acdd236b4d4baf24f2dd5c2d86533834db3912a5be431e4a5007b981548d): complete subsection reference.

- [tls_config](data-sources--nfv_service--reference--group-002.md#canonical-14ea4ec4883e65a8a90708fa693c6313cb4f4e6ef48470d9ce830ca32f742ce1): complete subsection reference.

- [use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-b09cbd9adc3a73d35b0e8158063c9ef822f97cfa300783551725395ffed3ae32): complete subsection reference.

<a id="canonical-c0009981ba57686c37327cf18e99403f41ae7a9b1542d23a9d35d9166660d1b7"></a>

## Next pages — https_management.advertise_on_sli_vip / 7c4bcaeb0da1 / 4

- [https_management.advertise_on_sli_vip.no_mtls](data-sources--nfv_service--reference--group-002.md#canonical-4c078eb10bee4746e319220ccb09f5a19c190cb49d47f8b43502167f1516abb6)
- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-78b0acdd236b4d4baf24f2dd5c2d86533834db3912a5be431e4a5007b981548d)
- [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-14ea4ec4883e65a8a90708fa693c6313cb4f4e6ef48470d9ce830ca32f742ce1)
- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-b09cbd9adc3a73d35b0e8158063c9ef822f97cfa300783551725395ffed3ae32)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-4c078eb10bee4746e319220ccb09f5a19c190cb49d47f8b43502167f1516abb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03611fa52dbee38e2dc9125fbde55ef3e94868497b7eb17e197fb23c6f74c2d2"></a>

## https_management.advertise_on_sli_vip.no_mtls — https_management.advertise_on_sli_vip.no_mtls / 099b9864fa0e / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- https_management.advertise_on_sli_vip.no_mtls

<a id="canonical-ce56e73411d50115233f57f7fb3145a9161b1cf3523babb2aa329de8f5163436"></a>

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

<a id="canonical-d632a8321a7d7b7e59769df22404f80db0169e7a7f780d6e0e7d7771974a5469"></a>

## Direct properties — https_management.advertise_on_sli_vip.no_mtls / 099b9864fa0e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-02de99a50b7d0e2811eb15e2c516f10e9df55b9f039a0f3b4f1950c3412a6333"></a>

## Next pages — https_management.advertise_on_sli_vip.no_mtls / 099b9864fa0e / 4

- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-78b0acdd236b4d4baf24f2dd5c2d86533834db3912a5be431e4a5007b981548d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef1272a2c76e8e3677fc01c30645a14920d5ea3e735e040e99800e8a07278de1"></a>

## https_management.advertise_on_sli_vip.tls_certificates — https_management.advertise_on_sli_vip.tls_certificates / f7d78a6c8f17 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- https_management.advertise_on_sli_vip.tls_certificates

<a id="canonical-8217ec8f93128b419d7abc05501050cb0618bb35c0ef3ca940cd0c791fbe9c5f"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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

<a id="canonical-ba95d1999898295e79dd3b1efbb4b85799aeb59986f64aed155921ebea2a2340"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_certificates / f7d78a6c8f17 / 3

<a id="canonical-b6d9a237a07f0d282d9f38df94ceaa834ec28a7f3f69c04dc2499af41d2c4422"></a>

<a id="canonical-4146427dcd7396604dbf5c87b43f2811bb1b7420afa8ac70d394c25d86c3c722"></a>

## certificate_url property — https_management.advertise_on_sli_vip.tls_certificates / f7d78a6c8f17 / 4

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

- [custom_hash_algorithms](data-sources--nfv_service--reference--group-002.md#canonical-2fac76379466fe62309acdf4a24b8e445f33aadd1f7f948b2d38bcad9cf122fd): complete subsection reference.

<a id="canonical-537fb4f258ae572f9260eda1209b0bd0bccbf44564d9f2874f36ab9ea80f72b5"></a>

<a id="canonical-f557aaca70f080ec1ca3d2dd8ec7dcbcccc7bccf778de14db20bf585e7cc3d08"></a>

## description_spec property — https_management.advertise_on_sli_vip.tls_certificates / f7d78a6c8f17 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--nfv_service--reference--group-002.md#canonical-6d20e56734b109d921adbbe603e8547a5a5da5b5ee7b0ac7d06558f885ea828a): complete subsection reference.

- [private_key](data-sources--nfv_service--reference--group-002.md#canonical-e21efdce5ae452e3d44836c89ba96c6ed7bde7271b979dbe67f0c78881b6899c): complete subsection reference.

- [use_system_defaults](data-sources--nfv_service--reference--group-002.md#canonical-8abc4cd42131893e1ee9cd2aa9f3338766fd0a20a4ad79d51345c766230349bb): complete subsection reference.

<a id="canonical-f6c75f4901adcb1b79708401d5616b22831ef52053732b0908f77e0c1b0d1999"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_certificates / f7d78a6c8f17 / 6

- [https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--reference--group-002.md#canonical-2fac76379466fe62309acdf4a24b8e445f33aadd1f7f948b2d38bcad9cf122fd)
- [https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--reference--group-002.md#canonical-6d20e56734b109d921adbbe603e8547a5a5da5b5ee7b0ac7d06558f885ea828a)
- [https_management.advertise_on_sli_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-e21efdce5ae452e3d44836c89ba96c6ed7bde7271b979dbe67f0c78881b6899c)
- [https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults](data-sources--nfv_service--reference--group-002.md#canonical-8abc4cd42131893e1ee9cd2aa9f3338766fd0a20a4ad79d51345c766230349bb)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-2fac76379466fe62309acdf4a24b8e445f33aadd1f7f948b2d38bcad9cf122fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee615ce8e5fb710bca3528c19fea796ce9fd0e749625950074891dfa6ea0a138"></a>

## https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms — https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms / 83b7022b4cb9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-78b0acdd236b4d4baf24f2dd5c2d86533834db3912a5be431e4a5007b981548d)
- https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms

<a id="canonical-9ffce2de81dc9ac671ba3c6bcccfdb7ffff80ad6f5b99b60d6fc7977577e4f3c"></a>

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

<a id="canonical-f4e51d06ce1e0d900a3076a175cc5bf21c0fa82b50550bb2facaf143e88313c6"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms / 83b7022b4cb9 / 3

<a id="canonical-94a35b1d4b889aa307ffa2a84e7c279696e8b81d8d5b7fef21b61b573ea5faed"></a>

<a id="canonical-8944a613d28c7b4068cb2372d1d6d38d1bb62a695960a8bb64e5dca567ce681d"></a>

## hash_algorithms property — https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms / 83b7022b4cb9 / 4

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

<a id="canonical-71eddd9c31658bde02b0769d3eee1b6d075ea2fa9773ff6ec1eb479fe2cd5853"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms / 83b7022b4cb9 / 5

- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-78b0acdd236b4d4baf24f2dd5c2d86533834db3912a5be431e4a5007b981548d)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-6d20e56734b109d921adbbe603e8547a5a5da5b5ee7b0ac7d06558f885ea828a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62503a97364088ae080231a7d538ef7b675caa5b2b6a2cbfd56946b25d74af7d"></a>

## https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling — https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling / 079ff63efea7 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-78b0acdd236b4d4baf24f2dd5c2d86533834db3912a5be431e4a5007b981548d)
- https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling

<a id="canonical-3b8984ed91c1a3fab4a4a48347ca8af7449c32f8dbdc1ade9f88f3981784459a"></a>

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

<a id="canonical-7ab935174fbeee47492815be988e846cedebc03237ee59bdaf1112af12cb5c2e"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling / 079ff63efea7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f2eabf10408c64e19e88c355bb5aa339939f8626a383e11b2e2f733441fd2a3b"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling / 079ff63efea7 / 4

- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-78b0acdd236b4d4baf24f2dd5c2d86533834db3912a5be431e4a5007b981548d)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-e21efdce5ae452e3d44836c89ba96c6ed7bde7271b979dbe67f0c78881b6899c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-443a63b5382a6194a8fdebc86ed113b3c62c1c3bcbdf724c7cba403c7db5e9d7"></a>

## https_management.advertise_on_sli_vip.tls_certificates.private_key — https_management.advertise_on_sli_vip.tls_certificates.private_key / 5006f0758f84 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-78b0acdd236b4d4baf24f2dd5c2d86533834db3912a5be431e4a5007b981548d)
- https_management.advertise_on_sli_vip.tls_certificates.private_key

<a id="canonical-832c64e9b321e74824904304cf37c3fa52be58d7c88969cddb583ef7d120ebf7"></a>

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

<a id="canonical-c8686c0d885a8cb92c3e567a6191db65ec1b19817995b23f16c15272ebcce5f0"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_certificates.private_key / 5006f0758f84 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-34145e93b74ddddfa34544cdb9c2cf86a0ba87f3c1d3054839bbf53b95d7fbb6): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-e4821399ae580efd7153834e45089a706fac9220ad4380b2f0038adabdd718db): complete subsection reference.

<a id="canonical-6abb446f27d5404600e6ff304c4a736cf03b113592487332273fa13d72ee48c1"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_certificates.private_key / 5006f0758f84 / 4

- [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-34145e93b74ddddfa34544cdb9c2cf86a0ba87f3c1d3054839bbf53b95d7fbb6)
- [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-e4821399ae580efd7153834e45089a706fac9220ad4380b2f0038adabdd718db)
- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-78b0acdd236b4d4baf24f2dd5c2d86533834db3912a5be431e4a5007b981548d)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-34145e93b74ddddfa34544cdb9c2cf86a0ba87f3c1d3054839bbf53b95d7fbb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00b9a57be14dec5c20f729298e1f4ed69e8daf6996e558fb18f4002fc9c7012b"></a>

## https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info — https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_sec / 647c2178f340 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-78b0acdd236b4d4baf24f2dd5c2d86533834db3912a5be431e4a5007b981548d)
- [https_management.advertise_on_sli_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-e21efdce5ae452e3d44836c89ba96c6ed7bde7271b979dbe67f0c78881b6899c)
- https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-db66e505f7d7798670f3cfd191ec33b67906b91796385e6c8645b3fd6a972ebe"></a>

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

<a id="canonical-0ec567d0554967cd31fdfd20e1dae70a58013808130b90dc7596bb50ed6da614"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_sec / 647c2178f340 / 3

<a id="canonical-e3c86ce8cf335439ace3c9fb20367fa6a570f916480def3a50bfbb03241f3a45"></a>

<a id="canonical-6161a87312d5a7ecd4a996225f5e0b445a6d6e3240a29da271fb075992b6683c"></a>

## decryption_provider property — https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_sec / 647c2178f340 / 4

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

<a id="canonical-92b08cf668151060f13fe1b7c1620b37ef6855535569e9cf8dfad733feccf921"></a>

<a id="canonical-99bc67bc69f89097440a9d62f17cd4f930fecf75aa2baec8edd1b8eb744c740a"></a>

## location property — https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_sec / 647c2178f340 / 5

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

<a id="canonical-7b66bf798882c956c3c5b42f1ccc3b4894447fd8d5dad9e80b06c04491abdd06"></a>

<a id="canonical-22c6fb4afdf6e00f0d25f6efaedd63e0c9a93d41a3e8346627aaefcfe2e0ad22"></a>

## store_provider property — https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_sec / 647c2178f340 / 6

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

<a id="canonical-df8ea8e2ecaa2f5654767945336cf64f829d52c15c21eb453c70cd6bf0dd556f"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_sec / 647c2178f340 / 7

- [https_management.advertise_on_sli_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-e21efdce5ae452e3d44836c89ba96c6ed7bde7271b979dbe67f0c78881b6899c)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-e4821399ae580efd7153834e45089a706fac9220ad4380b2f0038adabdd718db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26371fa34821f30364f78d21b9a546d3e89cb9267760d1e69d168b4e83f3ad3c"></a>

## https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info — https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_ / fc3a3fc04ef7 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-78b0acdd236b4d4baf24f2dd5c2d86533834db3912a5be431e4a5007b981548d)
- [https_management.advertise_on_sli_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-e21efdce5ae452e3d44836c89ba96c6ed7bde7271b979dbe67f0c78881b6899c)
- https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info

<a id="canonical-c37fa882674973a7b39a44e2043aa15a70a75a8d5ff1556491b3439960c18953"></a>

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

<a id="canonical-6b7f3e2bb0b7e709872a23fe414dcf47b812d3b859fbeb8720016178179a6ade"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_ / fc3a3fc04ef7 / 3

<a id="canonical-a8c07eb14dcff3cce31842faf8ff01b44b9627b84dd22ca67530724017c0f4b1"></a>

<a id="canonical-f464cef63b1b8432cb5ec267933f3743c7a5bda9a4ed012b80251cad05a4bed8"></a>

## provider_ref property — https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_ / fc3a3fc04ef7 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-ca40d21865bf8b6c90ee0537cf937047739d2eb6d20880f1234b90491d8db8c1"></a>

<a id="canonical-5613201dcdb2d8d863a6eac6faf7811af2cf5652285d2f06aea804a13433fa42"></a>

## url property — https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_ / fc3a3fc04ef7 / 5

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

<a id="canonical-aa5ca1495a3b17714f63d185d6f539248272dffe3f73ac3662abb21303d1ddcb"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_ / fc3a3fc04ef7 / 6

- [https_management.advertise_on_sli_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-e21efdce5ae452e3d44836c89ba96c6ed7bde7271b979dbe67f0c78881b6899c)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-8abc4cd42131893e1ee9cd2aa9f3338766fd0a20a4ad79d51345c766230349bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81cfd999c122a8f05945442b3b1f46c8530cdf0fa43b77ced1180353ada5a314"></a>

## https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults — https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults / 73449cff47b0 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-78b0acdd236b4d4baf24f2dd5c2d86533834db3912a5be431e4a5007b981548d)
- https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults

<a id="canonical-089f402daf7850da817d582ccf57b5910be4e28628cb320908124d01ef429a9b"></a>

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

<a id="canonical-665b48834ab59e3a8f8ce7e0eaf8cfec6cacb0043ca1ad6221f7d402049a7ff4"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults / 73449cff47b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fd67f064c0dbf19be2b8ee66c48e0a14da88a18d7457dfb214975dc24e893c58"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults / 73449cff47b0 / 4

- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-78b0acdd236b4d4baf24f2dd5c2d86533834db3912a5be431e4a5007b981548d)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-14ea4ec4883e65a8a90708fa693c6313cb4f4e6ef48470d9ce830ca32f742ce1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a668a7dc3658157446f2823c9053b074dcfe6ed5b915094a45c22d83fe6ae9d"></a>

## https_management.advertise_on_sli_vip.tls_config — https_management.advertise_on_sli_vip.tls_config / 2fbad5b2f318 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- https_management.advertise_on_sli_vip.tls_config

<a id="canonical-15606791ee664c18e6e0f3bba67607a29467a2fd8c67475b77d0c908b8cf86b4"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

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

<a id="canonical-808c14c7ee7272ca7be54a9c922b5b87f259dee7bec50868dc699a3f66337d98"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_config / 2fbad5b2f318 / 3

- [custom_security](data-sources--nfv_service--reference--group-002.md#canonical-c1ccb3372442ab3c00fb107fc9ba23f26ec4f19c5123d360c93e49a84c7e5814): complete subsection reference.

- [default_security](data-sources--nfv_service--reference--group-002.md#canonical-c3c1272866a4f9eee932235298f58cdf955c9a2357a2aa9b5bac44f6c570eed5): complete subsection reference.

- [low_security](data-sources--nfv_service--reference--group-002.md#canonical-085fadfd9a7f32952c9109b7b9312b67bb5d4b156548d8f20c7816b9ac3d9fbc): complete subsection reference.

- [medium_security](data-sources--nfv_service--reference--group-002.md#canonical-faaa9b558bd35ba05e765fc098f0f4b8821f3911ced24894c11b8ddaecc440d7): complete subsection reference.

<a id="canonical-66065886b576fe263159fd0de947554765e96d7b54ba0bc0057279821ac801ba"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_config / 2fbad5b2f318 / 4

- [https_management.advertise_on_sli_vip.tls_config.custom_security](data-sources--nfv_service--reference--group-002.md#canonical-c1ccb3372442ab3c00fb107fc9ba23f26ec4f19c5123d360c93e49a84c7e5814)
- [https_management.advertise_on_sli_vip.tls_config.default_security](data-sources--nfv_service--reference--group-002.md#canonical-c3c1272866a4f9eee932235298f58cdf955c9a2357a2aa9b5bac44f6c570eed5)
- [https_management.advertise_on_sli_vip.tls_config.low_security](data-sources--nfv_service--reference--group-002.md#canonical-085fadfd9a7f32952c9109b7b9312b67bb5d4b156548d8f20c7816b9ac3d9fbc)
- [https_management.advertise_on_sli_vip.tls_config.medium_security](data-sources--nfv_service--reference--group-002.md#canonical-faaa9b558bd35ba05e765fc098f0f4b8821f3911ced24894c11b8ddaecc440d7)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-c1ccb3372442ab3c00fb107fc9ba23f26ec4f19c5123d360c93e49a84c7e5814"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f86c192cd96a51b55dec27390ac0fc96ca0b106ba68acfc7536a419cf268dbb9"></a>

## https_management.advertise_on_sli_vip.tls_config.custom_security — https_management.advertise_on_sli_vip.tls_config.custom_security / d24ecca16839 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-14ea4ec4883e65a8a90708fa693c6313cb4f4e6ef48470d9ce830ca32f742ce1)
- https_management.advertise_on_sli_vip.tls_config.custom_security

<a id="canonical-7b8d00c798347ae3999b7956117a75b220cafd519c56d3c64fd042d44dcd71a4"></a>

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

<a id="canonical-31c78c671a06ce277b8369e251f39e827cf2a80f456b45ba0991ba49052fe85a"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_config.custom_security / d24ecca16839 / 3

<a id="canonical-c74b5331c8c59d2ae05b6c1109a7580ae75a337ac5068cae501ed6343efb277d"></a>

<a id="canonical-1cbb7244a0e28325bf17ff244736e4cf7d3ada84b2361f5389cd93cab078b84d"></a>

## cipher_suites property — https_management.advertise_on_sli_vip.tls_config.custom_security / d24ecca16839 / 4

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

<a id="canonical-614150add9d866c692b2dfc792e240bc71921fa50dd0747555571056a532ebec"></a>

<a id="canonical-5415dcc14fe1107d4a0fe517c4acf8908f5c6c8076021998530c460a62e1d599"></a>

## max_version property — https_management.advertise_on_sli_vip.tls_config.custom_security / d24ecca16839 / 5

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

<a id="canonical-bb6d9beb59422c65ab4ebc5515112101899a523261b82ece803d4a505304c436"></a>

<a id="canonical-e0a847eac94229413f083a4b7dbce32322375fa986303a27f0e2eb7761104540"></a>

## min_version property — https_management.advertise_on_sli_vip.tls_config.custom_security / d24ecca16839 / 6

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

<a id="canonical-37ef9021c0f7bb92b480ab5f715658daf873320e577c9902492048c45e611a24"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_config.custom_security / d24ecca16839 / 7

- [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-14ea4ec4883e65a8a90708fa693c6313cb4f4e6ef48470d9ce830ca32f742ce1)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-c3c1272866a4f9eee932235298f58cdf955c9a2357a2aa9b5bac44f6c570eed5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cdfd9ee427119bbffd06b3c5304da1ecfa58773ef8d1abae94b19569f7a7c8e7"></a>

## https_management.advertise_on_sli_vip.tls_config.default_security — https_management.advertise_on_sli_vip.tls_config.default_security / 6ce35636bbdb / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-14ea4ec4883e65a8a90708fa693c6313cb4f4e6ef48470d9ce830ca32f742ce1)
- https_management.advertise_on_sli_vip.tls_config.default_security

<a id="canonical-055469d52dddda8c20be3ae21d9b5bcf4e649dae634d28541c75b4e55b7fbf29"></a>

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

<a id="canonical-161ecbbeec7c183dde9e9ecfe1c94474d4ccc3deac1bc231d29990773752c427"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_config.default_security / 6ce35636bbdb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-46ea45e25387c8500f81d532d544fae599631551171c26ca7e21dbd517f9ad8c"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_config.default_security / 6ce35636bbdb / 4

- [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-14ea4ec4883e65a8a90708fa693c6313cb4f4e6ef48470d9ce830ca32f742ce1)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-085fadfd9a7f32952c9109b7b9312b67bb5d4b156548d8f20c7816b9ac3d9fbc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ea1f58f61e38070d2e2b8119486ba5ff6b14fdfb24cf1372087f41b64d82cef"></a>

## https_management.advertise_on_sli_vip.tls_config.low_security — https_management.advertise_on_sli_vip.tls_config.low_security / a45cfa6ba922 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-14ea4ec4883e65a8a90708fa693c6313cb4f4e6ef48470d9ce830ca32f742ce1)
- https_management.advertise_on_sli_vip.tls_config.low_security

<a id="canonical-098503c98aef1caf8114de8ac0c07516f3d3adec3edca6df231ec5530dafdbd7"></a>

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

<a id="canonical-d7c2d9fdd31cb2e9df66d5dc358dd1c1404693e13d80c21450e4cabbaa8cb901"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_config.low_security / a45cfa6ba922 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-54e35a6e79cab1296aeedde6249f9076cf44f05bb23df122c5205c88b14c5b9c"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_config.low_security / a45cfa6ba922 / 4

- [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-14ea4ec4883e65a8a90708fa693c6313cb4f4e6ef48470d9ce830ca32f742ce1)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-faaa9b558bd35ba05e765fc098f0f4b8821f3911ced24894c11b8ddaecc440d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a99ae1cab8b62f02f0fb894970a555dbd5007428b14ea7bd500c741067e910d3"></a>

## https_management.advertise_on_sli_vip.tls_config.medium_security — https_management.advertise_on_sli_vip.tls_config.medium_security / f8b77f2b68ca / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-14ea4ec4883e65a8a90708fa693c6313cb4f4e6ef48470d9ce830ca32f742ce1)
- https_management.advertise_on_sli_vip.tls_config.medium_security

<a id="canonical-42e5f4707adbafb3cebbfaeb84a1782da818aa1ed745918b820db86d5b40e984"></a>

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

<a id="canonical-a2de2ec10211beb098635aed44a084fedd6cc3c02cce0953882278d3238bb5e0"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_config.medium_security / f8b77f2b68ca / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f1ba50a1c90a8f077c7c1069db29c78a7f168a618f402229662510d156936ab1"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_config.medium_security / f8b77f2b68ca / 4

- [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-14ea4ec4883e65a8a90708fa693c6313cb4f4e6ef48470d9ce830ca32f742ce1)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-b09cbd9adc3a73d35b0e8158063c9ef822f97cfa300783551725395ffed3ae32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbfb3c0840b70bdca67a5d5c0c5bc6d54f67c8c46a4029fa32afc3e0f88dcd40"></a>

## https_management.advertise_on_sli_vip.use_mtls — https_management.advertise_on_sli_vip.use_mtls / 2a35b8a44944 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- https_management.advertise_on_sli_vip.use_mtls

<a id="canonical-41282ee243a138e9d2e5aaab686fd7cf00023932b5bab22dd5f1484ec925417e"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

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

<a id="canonical-85d91744ab40a03826bfa91e34edef55de933ca2bb231703186a5aed35295463"></a>

## Direct properties — https_management.advertise_on_sli_vip.use_mtls / 2a35b8a44944 / 3

<a id="canonical-7902133232464c99e26c2e080d30eb5cfa172a6cc4e7a9c103f74d0b33333396"></a>

<a id="canonical-59958018a1f0b455f236021858582c28801625a6fa60c25a7bf5625d1dea7d8f"></a>

## client_certificate_optional property — https_management.advertise_on_sli_vip.use_mtls / 2a35b8a44944 / 4

Type: `"bool"`. Computed.

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

- [crl](data-sources--nfv_service--reference--group-002.md#canonical-5a5cb1191862667ada0aa09070992dd76e27a80eed62129404978138cc60d9c0): complete subsection reference.

- [no_crl](data-sources--nfv_service--reference--group-002.md#canonical-1c1a5e719a24e9f39148c9514a9a749635a25e961d8f5fd07c00c13f340f1ae5): complete subsection reference.

- [trusted_ca](data-sources--nfv_service--reference--group-002.md#canonical-7207bef3af1a22bde0f5acc50fdd59f90b047a05869152c68f0328f558177680): complete subsection reference.

<a id="canonical-57d89f8e3a777dd84a31a5e0b571922206c7e96c588e91342a8dc75d9a5ddb8d"></a>

<a id="canonical-d917a531feda441dff797de5bbe5710eb0e291ecba3d0f361eb786a95b930885"></a>

## trusted_ca_url property — https_management.advertise_on_sli_vip.use_mtls / 2a35b8a44944 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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

- [xfcc_disabled](data-sources--nfv_service--reference--group-002.md#canonical-2fe7d8299de08f26094b48b7f37b104bec6bcf04eae559f698ded95c5454b83a): complete subsection reference.

- [xfcc_options](data-sources--nfv_service--reference--group-002.md#canonical-b0e130df6937ff590fad347f56ecb6ad659e1d57a9769e0dddddfc1e01822c94): complete subsection reference.

<a id="canonical-4b0a36a56c61f5e6f1039c97a2fc1f20dc47692a773748317f10916b1782a388"></a>

## Next pages — https_management.advertise_on_sli_vip.use_mtls / 2a35b8a44944 / 6

- [https_management.advertise_on_sli_vip.use_mtls.crl](data-sources--nfv_service--reference--group-002.md#canonical-5a5cb1191862667ada0aa09070992dd76e27a80eed62129404978138cc60d9c0)
- [https_management.advertise_on_sli_vip.use_mtls.no_crl](data-sources--nfv_service--reference--group-002.md#canonical-1c1a5e719a24e9f39148c9514a9a749635a25e961d8f5fd07c00c13f340f1ae5)
- [https_management.advertise_on_sli_vip.use_mtls.trusted_ca](data-sources--nfv_service--reference--group-002.md#canonical-7207bef3af1a22bde0f5acc50fdd59f90b047a05869152c68f0328f558177680)
- [https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled](data-sources--nfv_service--reference--group-002.md#canonical-2fe7d8299de08f26094b48b7f37b104bec6bcf04eae559f698ded95c5454b83a)
- [https_management.advertise_on_sli_vip.use_mtls.xfcc_options](data-sources--nfv_service--reference--group-002.md#canonical-b0e130df6937ff590fad347f56ecb6ad659e1d57a9769e0dddddfc1e01822c94)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-5a5cb1191862667ada0aa09070992dd76e27a80eed62129404978138cc60d9c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6a030eca24a3720ba71a28ebb8ba08319cfa11d1984d13e9564efe4dc8d9e00"></a>

## https_management.advertise_on_sli_vip.use_mtls.crl — https_management.advertise_on_sli_vip.use_mtls.crl / ba182e19dff3 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-b09cbd9adc3a73d35b0e8158063c9ef822f97cfa300783551725395ffed3ae32)
- https_management.advertise_on_sli_vip.use_mtls.crl

<a id="canonical-95d43dd626bd334bcd0f533291014e3fe83fc979731fe9d86526a07d243a3ba4"></a>

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

<a id="canonical-7e97010489ef1b48348c4cd1cb87b1c6f4cb8e1f903d782660440b464eda4dbc"></a>

## Direct properties — https_management.advertise_on_sli_vip.use_mtls.crl / ba182e19dff3 / 3

<a id="canonical-35d34760c3ef015fe50f5ca9c04e8a1249718185f0d5287d336b99d1db6ca4d9"></a>

<a id="canonical-ac59f7fa73d4b0ca37f3515e5e947efef10ad463d730bb15908baa5885b89ab2"></a>

## name property — https_management.advertise_on_sli_vip.use_mtls.crl / ba182e19dff3 / 4

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

<a id="canonical-ebd9bcd892ac7543c2d644c240f700a53029952eca39d6586de0703c30ccce91"></a>

<a id="canonical-6c0b6da601f97f0757c4a6eb2baa0683238d08a584bf730a5d07c7b6f3ec7add"></a>

## namespace property — https_management.advertise_on_sli_vip.use_mtls.crl / ba182e19dff3 / 5

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

<a id="canonical-59761ff4c94e8778a89c0577beb1acc078d1bc8393241611546f016cc327cd1e"></a>

<a id="canonical-e6ce6c57b69a417772956abf5a840719a267ed9e105ee3e6d60f2cb97660bb66"></a>

## tenant property — https_management.advertise_on_sli_vip.use_mtls.crl / ba182e19dff3 / 6

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

<a id="canonical-b3d24544115bb0b8183154592ba85105e0f1ee0ed7efbc54fbd8e0119696bfb9"></a>

## Next pages — https_management.advertise_on_sli_vip.use_mtls.crl / ba182e19dff3 / 7

- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-b09cbd9adc3a73d35b0e8158063c9ef822f97cfa300783551725395ffed3ae32)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-1c1a5e719a24e9f39148c9514a9a749635a25e961d8f5fd07c00c13f340f1ae5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96c1fd4b9bc70ee7affbca9fdc01a208f93ca340ba2124bd18a86184e5159954"></a>

## https_management.advertise_on_sli_vip.use_mtls.no_crl — https_management.advertise_on_sli_vip.use_mtls.no_crl / 9fdc44e25ed8 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-b09cbd9adc3a73d35b0e8158063c9ef822f97cfa300783551725395ffed3ae32)
- https_management.advertise_on_sli_vip.use_mtls.no_crl

<a id="canonical-634960c3aae56271ad89b5fca146eed226a07bf85925b8f5e8f953eaf251eda0"></a>

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

<a id="canonical-253b385dcfc6301b913878950163067efeadc2c75609bc2d343ae8dc1bb218e6"></a>

## Direct properties — https_management.advertise_on_sli_vip.use_mtls.no_crl / 9fdc44e25ed8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a3182bb70247f782fdb10767b432b092a103f9664d5916bbb08e6f5ad929a773"></a>

## Next pages — https_management.advertise_on_sli_vip.use_mtls.no_crl / 9fdc44e25ed8 / 4

- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-b09cbd9adc3a73d35b0e8158063c9ef822f97cfa300783551725395ffed3ae32)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-7207bef3af1a22bde0f5acc50fdd59f90b047a05869152c68f0328f558177680"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc9938f074d44a79d27981224c370dbb37d42ff98b0ce81784b7dd5b3d021f0b"></a>

## https_management.advertise_on_sli_vip.use_mtls.trusted_ca — https_management.advertise_on_sli_vip.use_mtls.trusted_ca / cb98c775eee7 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-b09cbd9adc3a73d35b0e8158063c9ef822f97cfa300783551725395ffed3ae32)
- https_management.advertise_on_sli_vip.use_mtls.trusted_ca

<a id="canonical-5748f0476cc3b5a2b34bc1459c614163303ff2ede137c43fcb88db94fa6483d3"></a>

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

<a id="canonical-d72aecbedad6fa0a01d2ff4e9811de2ccef2d6c62615862e4f19ca0a4f30262a"></a>

## Direct properties — https_management.advertise_on_sli_vip.use_mtls.trusted_ca / cb98c775eee7 / 3

<a id="canonical-5a38884dcd8d872af4866460320fe8ed090544496bbe42bd7309e17f1a62b7e7"></a>

<a id="canonical-88f11739259ae3e9bdf76f62ab58d6378f243e359b0039b4cd2994771c2c4a19"></a>

## name property — https_management.advertise_on_sli_vip.use_mtls.trusted_ca / cb98c775eee7 / 4

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

<a id="canonical-18af490b8b7f2ec672a5b498dd60dab83ae9471c1fac00b385812c4a8f2241e5"></a>

<a id="canonical-6bb338b539ed2f761e8735a6b9702137bbfac087da7dd08c92f3945b3d95bc7f"></a>

## namespace property — https_management.advertise_on_sli_vip.use_mtls.trusted_ca / cb98c775eee7 / 5

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

<a id="canonical-a769332707a430594b38fb62ef215e70d1f96efa3d6c58a4bb1a4e8a2ce089bc"></a>

<a id="canonical-cfdfa81e5aa5ed48999258ba2682a0501513a76cd820fc0e010162e1751902e5"></a>

## tenant property — https_management.advertise_on_sli_vip.use_mtls.trusted_ca / cb98c775eee7 / 6

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

<a id="canonical-a8b5d9bd2b29a3f16a539d4875b574bd00130b2075935fa9ff0fe3a166041045"></a>

## Next pages — https_management.advertise_on_sli_vip.use_mtls.trusted_ca / cb98c775eee7 / 7

- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-b09cbd9adc3a73d35b0e8158063c9ef822f97cfa300783551725395ffed3ae32)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-2fe7d8299de08f26094b48b7f37b104bec6bcf04eae559f698ded95c5454b83a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13b5abb6549f4deae66b42c3671576fb1fe196a403e094bd7521a358eb3dd1ed"></a>

## https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled — https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled / ae34b59ae367 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-b09cbd9adc3a73d35b0e8158063c9ef822f97cfa300783551725395ffed3ae32)
- https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled

<a id="canonical-2d81f0f2a05035a62e4488c2ff331d3f4e7272f774790ba084750369c884f47f"></a>

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

<a id="canonical-a8747b50bc0f5a7f4a2a791fb6231a1e6f8bec7101dd3a5e4cbe5807e8019814"></a>

## Direct properties — https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled / ae34b59ae367 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-489b884e364acaeb7be3cde2c8062d9aa074a1a6eadacbe89c833471a69054ba"></a>

## Next pages — https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled / ae34b59ae367 / 4

- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-b09cbd9adc3a73d35b0e8158063c9ef822f97cfa300783551725395ffed3ae32)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-b0e130df6937ff590fad347f56ecb6ad659e1d57a9769e0dddddfc1e01822c94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f486d1199a1ee7c42b5cc9814fe7dc6832e4e4ab1089273dc213e0f0efd4a45"></a>

## https_management.advertise_on_sli_vip.use_mtls.xfcc_options — https_management.advertise_on_sli_vip.use_mtls.xfcc_options / d308423578f9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-002.md#canonical-93bd04204c977b2fb2cb5c1f39eeeb1ea5a2c5d7381f34a5cc397d8585671ed6)
- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-b09cbd9adc3a73d35b0e8158063c9ef822f97cfa300783551725395ffed3ae32)
- https_management.advertise_on_sli_vip.use_mtls.xfcc_options

<a id="canonical-63158af95a0886672e9573b2c2e09caf843ce19c02469968030cefa69b5af26c"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-deb2cf7a795e071369277d5251f1c50e819377f1fe76e29d3c152df1423eac84"></a>

## Direct properties — https_management.advertise_on_sli_vip.use_mtls.xfcc_options / d308423578f9 / 3

<a id="canonical-608ba2bcdddc191c1d537768d99de5fa61799385303542b3df6244ed45116489"></a>

<a id="canonical-f06f254163fdbfd1824c6428eadf2b3f39271ccc2923bc5b5380c135b159f556"></a>

## xfcc_header_elements property — https_management.advertise_on_sli_vip.use_mtls.xfcc_options / d308423578f9 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-5bc7456fecb0563c279814da2d8bea5f7e5adb013e47b2723f713e21bea84d2e"></a>

## Next pages — https_management.advertise_on_sli_vip.use_mtls.xfcc_options / d308423578f9 / 5

- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-b09cbd9adc3a73d35b0e8158063c9ef822f97cfa300783551725395ffed3ae32)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-388ed78de0053b47119f14a1aba32a425360024117aca4064a0d2046b2f374b7"></a>

## https_management.advertise_on_slo_internet_vip — https_management.advertise_on_slo_internet_vip / dd20baefb812 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- https_management.advertise_on_slo_internet_vip

<a id="canonical-63d481ba4f4be297c47c9c63cebc2ce640cbf2a29c56ffd45509e4d72f593461"></a>

Type: `"single"`. Computed.

Inline TLS Parameters. Inline TLS parameters.

Upstream description:

Inline TLS parameters.

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

<a id="canonical-8f8e5b917ec780955f3c806deb4fa79dd0ee8f4f53181a63c391aaf71c07a2f3"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip / dd20baefb812 / 3

- [no_mtls](data-sources--nfv_service--reference--group-002.md#canonical-aad68a184f7f3c1232d7eea90f4684c4f6c3d78471be88c49dadd410e5a788bd): complete subsection reference.

- [tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-31f7d3af5f0bd351fc9e921d2e2fbbe51ee4c42e5d62447f010d20e61ddab436): complete subsection reference.

- [tls_config](data-sources--nfv_service--reference--group-002.md#canonical-78f0ca926c5842a0af13f5d6e2f298a65322948981f7d069e7ade2213d9b77c7): complete subsection reference.

- [use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-9fd77b3bb6352a172d395c7191f2b909686330ea314aec3eef82dc410240657a): complete subsection reference.

<a id="canonical-40f82edd66c267a27d2e0385232411a7dbf40830f778e712f6ac3f06d5290a6a"></a>

## Next pages — https_management.advertise_on_slo_internet_vip / dd20baefb812 / 4

- [https_management.advertise_on_slo_internet_vip.no_mtls](data-sources--nfv_service--reference--group-002.md#canonical-aad68a184f7f3c1232d7eea90f4684c4f6c3d78471be88c49dadd410e5a788bd)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-31f7d3af5f0bd351fc9e921d2e2fbbe51ee4c42e5d62447f010d20e61ddab436)
- [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-78f0ca926c5842a0af13f5d6e2f298a65322948981f7d069e7ade2213d9b77c7)
- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-9fd77b3bb6352a172d395c7191f2b909686330ea314aec3eef82dc410240657a)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-aad68a184f7f3c1232d7eea90f4684c4f6c3d78471be88c49dadd410e5a788bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81910c5bb35ed8e35a9026b9562c7d530a4a5d6d269ef0d282f7b8adebf6900f"></a>

## https_management.advertise_on_slo_internet_vip.no_mtls — https_management.advertise_on_slo_internet_vip.no_mtls / 671aed3b0e6f / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- https_management.advertise_on_slo_internet_vip.no_mtls

<a id="canonical-ee5627be6eab856b4a638ec02eccd2f65a2aae15468260ead52d1632e4dfb473"></a>

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

<a id="canonical-6e6a975ea42ed6347f14da2115ac84a112e76a2d776dd6316c791876f47566e9"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.no_mtls / 671aed3b0e6f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-73249e374523e250d14953520b8deffea29084952673f6156a5ad7cb3ab59760"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.no_mtls / 671aed3b0e6f / 4

- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-31f7d3af5f0bd351fc9e921d2e2fbbe51ee4c42e5d62447f010d20e61ddab436"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ae32d68f2b37543d834a117357d4e721796ef4edde4107abe8be81b6ca0da75"></a>

## https_management.advertise_on_slo_internet_vip.tls_certificates — https_management.advertise_on_slo_internet_vip.tls_certificates / c13f0f8465bb / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- https_management.advertise_on_slo_internet_vip.tls_certificates

<a id="canonical-5c3eec0736bdb6b1f090153cfa8f6e0007d16d339e37ae0bbd49c337df86b731"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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

<a id="canonical-53aae58354b51a4d69efe44997531232e2e0075260902f81797f80de5bcdbc43"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_certificates / c13f0f8465bb / 3

<a id="canonical-448b1f48dae15d619ab980b8a715b546d486e26d4ce4954ab1c814499a66ed9a"></a>

<a id="canonical-e578dd317f41739c3c7c7d43e3bae390b16d81e6b99f3985a7addccdfa4ec9e6"></a>

## certificate_url property — https_management.advertise_on_slo_internet_vip.tls_certificates / c13f0f8465bb / 4

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

- [custom_hash_algorithms](data-sources--nfv_service--reference--group-002.md#canonical-731da62a332679ce5426f60612eb1c7c44c16727e3cdf9f958b99acf459dc39c): complete subsection reference.

<a id="canonical-c3bff427c7b4f59d4e66a4bc6769abb1e1f1a3171bd89ff2f82040ae906aa7ed"></a>

<a id="canonical-32d6d184f0db21619a794d0c70020340ad17fe90aead40b671b9370df2e37bca"></a>

## description_spec property — https_management.advertise_on_slo_internet_vip.tls_certificates / c13f0f8465bb / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--nfv_service--reference--group-002.md#canonical-e3f4c5397dd4491caed7ae272240aad739c9f428bd0962aa70a2ea7b3751caaf): complete subsection reference.

- [private_key](data-sources--nfv_service--reference--group-002.md#canonical-73c6183db20ea1f374bea730f2f73a2777759eaf80680d288ddc63255a96c68c): complete subsection reference.

- [use_system_defaults](data-sources--nfv_service--reference--group-002.md#canonical-8974fa8413b00affe03802a5560527bc4ab79ef6335815afe10d423c68176727): complete subsection reference.

<a id="canonical-a0895eb3b535d889a4c4f37acd7f36508ed0efc4aac1995f45dc798a9038d99c"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_certificates / c13f0f8465bb / 6

- [https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--reference--group-002.md#canonical-731da62a332679ce5426f60612eb1c7c44c16727e3cdf9f958b99acf459dc39c)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--reference--group-002.md#canonical-e3f4c5397dd4491caed7ae272240aad739c9f428bd0962aa70a2ea7b3751caaf)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-73c6183db20ea1f374bea730f2f73a2777759eaf80680d288ddc63255a96c68c)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults](data-sources--nfv_service--reference--group-002.md#canonical-8974fa8413b00affe03802a5560527bc4ab79ef6335815afe10d423c68176727)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-731da62a332679ce5426f60612eb1c7c44c16727e3cdf9f958b99acf459dc39c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-526e44d84f168729c2593609bf62649f50f76a41afd9538709625f871e6c2e6e"></a>

## https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms — https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algo / 58aba57a370b / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-31f7d3af5f0bd351fc9e921d2e2fbbe51ee4c42e5d62447f010d20e61ddab436)
- https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms

<a id="canonical-d5dab0f6be34a79019fc4979707e30e7f9a5b31b43b36c6f706fd2d6ad5ae170"></a>

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

<a id="canonical-63abe8e5462eecbef3731e17acffc6ce5b04ea61f16046159972a3deb6bd15a2"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algo / 58aba57a370b / 3

<a id="canonical-cd895daeefeec2a58de6a8bfed8348f9d4e42b3e9e027d88214f11876cb01fe2"></a>

<a id="canonical-5024374fa92b9e29fc2b1e213f9b5de67005f71c8655eeeafb991b98e6cc16c4"></a>

## hash_algorithms property — https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algo / 58aba57a370b / 4

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

<a id="canonical-a303c78b662d897a3d44ba27f83d41a8e968db3047f1b47540fe22053d7804b0"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algo / 58aba57a370b / 5

- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-31f7d3af5f0bd351fc9e921d2e2fbbe51ee4c42e5d62447f010d20e61ddab436)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-e3f4c5397dd4491caed7ae272240aad739c9f428bd0962aa70a2ea7b3751caaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c2ab12b338492ec3bfad0c8a0f381746ee9a9f19c883f8c01d2fe3c0f0784ab"></a>

## https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling — https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_sta / 4b8fde073fb1 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-31f7d3af5f0bd351fc9e921d2e2fbbe51ee4c42e5d62447f010d20e61ddab436)
- https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling

<a id="canonical-97d29eb5c07de3e88070a8d2d9c68afd00c513da19ed985daaabe2b5b8d548be"></a>

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

<a id="canonical-ecd8b0f8915b782ba3b8bf4b704b6d42ee60ce064c2642c3b3e64c4c4d2e6340"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_sta / 4b8fde073fb1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-711d509e5028aca6246ae1a937085e856045ef92c758041c3f364fd15e181e84"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_sta / 4b8fde073fb1 / 4

- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-31f7d3af5f0bd351fc9e921d2e2fbbe51ee4c42e5d62447f010d20e61ddab436)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-73c6183db20ea1f374bea730f2f73a2777759eaf80680d288ddc63255a96c68c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c7ae9a875d8dbab3daecb9708e4895a548bef5420632b5509a8e6efe09824a9"></a>

## https_management.advertise_on_slo_internet_vip.tls_certificates.private_key — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key / 86a543a4a956 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-31f7d3af5f0bd351fc9e921d2e2fbbe51ee4c42e5d62447f010d20e61ddab436)
- https_management.advertise_on_slo_internet_vip.tls_certificates.private_key

<a id="canonical-c508d096f3cd25a2fe82533bd0b673caa5e79c9b08b85dc8755d4f5427891f0d"></a>

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

<a id="canonical-80c284feece041ee93cea3ad3b7321a9688798b022b43494c14c965d110c1984"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key / 86a543a4a956 / 3

- [blindfold_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-ea61ff2df9b7212c288a7533b523cb35a52e89eb73d66883f77b52856e7c03f2): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-764f015ffd5484c8a11b51957f226ee8a811fb3d787006d946591d4575aa3fbb): complete subsection reference.

<a id="canonical-22e4c0b81ff399d36060d70598330d797350bdb3c95859476811dfa891980ec0"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key / 86a543a4a956 / 4

- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-ea61ff2df9b7212c288a7533b523cb35a52e89eb73d66883f77b52856e7c03f2)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-764f015ffd5484c8a11b51957f226ee8a811fb3d787006d946591d4575aa3fbb)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-31f7d3af5f0bd351fc9e921d2e2fbbe51ee4c42e5d62447f010d20e61ddab436)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-ea61ff2df9b7212c288a7533b523cb35a52e89eb73d66883f77b52856e7c03f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48a53f2a7aa99e1cf5d56409a2c855e594c06a8d649d422c2fe6d55e040a95a3"></a>

## https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blin / 5278845a75b7 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-31f7d3af5f0bd351fc9e921d2e2fbbe51ee4c42e5d62447f010d20e61ddab436)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-73c6183db20ea1f374bea730f2f73a2777759eaf80680d288ddc63255a96c68c)
- https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-6798fab4bbc53734bd1fde479032848f3ee2d8602861d34021f98538fcab3fd5"></a>

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

<a id="canonical-16d8de22f2560ff70ede0e56bccf9bc89f35aa41da4c2d1ee83d05468ea5a7a5"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blin / 5278845a75b7 / 3

<a id="canonical-0b2cf4fb0593f5c2f7502296379048e92a87f40cf62dc743d09a730c9d82bd95"></a>

<a id="canonical-35f997f60eb590982107cd75c157ecc8087e063da45e0e611678753b0f36f20d"></a>

## decryption_provider property — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blin / 5278845a75b7 / 4

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

<a id="canonical-c5652f051ebaf3d443f9a02d4bc462088fc67878ad9441e041e74d19b3feedc8"></a>

<a id="canonical-f298d67733a24ba89c2210ea9ea79723fc37dda8d6c366664c915ddfd751db38"></a>

## location property — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blin / 5278845a75b7 / 5

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

<a id="canonical-87a768f1958a8afddf8539377f4124fd1a01810cd85d95bfb56d4a81d2e35c33"></a>

<a id="canonical-c0800e20558620988b0ecd13f658c09867b4e8f4a843a5dbeb74485fe6e96851"></a>

## store_provider property — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blin / 5278845a75b7 / 6

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

<a id="canonical-665b6a296e136140598ec25b909e33f04ed112a5e3cfe85a2c5abdc6080dd42d"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blin / 5278845a75b7 / 7

- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-73c6183db20ea1f374bea730f2f73a2777759eaf80680d288ddc63255a96c68c)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-764f015ffd5484c8a11b51957f226ee8a811fb3d787006d946591d4575aa3fbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-077b71ee21fc9311808652062e8b985da268c28222de254d9332202e9b78a35f"></a>

## https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clea / e5503c1bde05 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-31f7d3af5f0bd351fc9e921d2e2fbbe51ee4c42e5d62447f010d20e61ddab436)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-73c6183db20ea1f374bea730f2f73a2777759eaf80680d288ddc63255a96c68c)
- https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info

<a id="canonical-7390502a546007c1fa440d8c503a9090df2aa93e8bb61d9da1f461d2fe5f9674"></a>

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

<a id="canonical-e8dd3c667f578ba1169215614244415a757e55f1df6f24715223e1cee72bbb7f"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clea / e5503c1bde05 / 3

<a id="canonical-4008344de00fc61c30d75c7dc892e562d8a43477c8ec770d218fef429d1c55f1"></a>

<a id="canonical-ee4184a4fd2bcd28f54bb50f6efff3ec3e4d3c98eddf373c811c7b685ddc5a4b"></a>

## provider_ref property — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clea / e5503c1bde05 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-09e5533779e2f6dce31acc31735f1691fdc3631c3bb4d74a2f0a78560de63cac"></a>

<a id="canonical-3e6ec31b5549e9711ff7dd5629074867f99fcd7bb52687e538970f655f09e279"></a>

## url property — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clea / e5503c1bde05 / 5

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

<a id="canonical-eca62cd24f8d932bf8e73295f5a958d1d289724e9163f599ab7b207449627927"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clea / e5503c1bde05 / 6

- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-73c6183db20ea1f374bea730f2f73a2777759eaf80680d288ddc63255a96c68c)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-8974fa8413b00affe03802a5560527bc4ab79ef6335815afe10d423c68176727"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a4706b6a4769cede20c3ba5c5c5538eb012449ec6a30594503b5cbf701dd7d0"></a>

## https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults — https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defau / 6a563ad32087 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-31f7d3af5f0bd351fc9e921d2e2fbbe51ee4c42e5d62447f010d20e61ddab436)
- https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults

<a id="canonical-87b6db5582baf32fc5940587dca4c60e6f1e872cb5797bab7c79555c6f324d71"></a>

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

<a id="canonical-38c58d97731cd97c829fdcefba4dae8f59d9ff910343ab0fb62bb8df11c97ed5"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defau / 6a563ad32087 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-74f03c4d02909010bd476728569c84a3ecfb0845365f78ea8cc9e62f2de3430b"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defau / 6a563ad32087 / 4

- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-31f7d3af5f0bd351fc9e921d2e2fbbe51ee4c42e5d62447f010d20e61ddab436)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-78f0ca926c5842a0af13f5d6e2f298a65322948981f7d069e7ade2213d9b77c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f04a0a3f9b894b91be7629585f90b1feb5bda56bb07af417dffbc359e202a70a"></a>

## https_management.advertise_on_slo_internet_vip.tls_config — https_management.advertise_on_slo_internet_vip.tls_config / 5e71873c63bd / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- https_management.advertise_on_slo_internet_vip.tls_config

<a id="canonical-9f4794a7f17fad83392fceb46da5239015a93f311924f1cff4209344351af1af"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

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

<a id="canonical-dc33e14ac3d0714c5a3a593c7c0f50dc370c90f7534a08c410be400da9e754de"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_config / 5e71873c63bd / 3

- [custom_security](data-sources--nfv_service--reference--group-002.md#canonical-bd43e3aa4659d130acc5185619122d9bdd7b5f7c38b05ffb3de88f0717dda7a1): complete subsection reference.

- [default_security](data-sources--nfv_service--reference--group-002.md#canonical-092e2d7f8dfc739d12fd29d398ae12daa4d832b550097bdfec6c6b51bfe6dc15): complete subsection reference.

- [low_security](data-sources--nfv_service--reference--group-002.md#canonical-4948f67d8b0d5444b8ed3373a96f91ea6b493284a1f7386cd15e506cbc253837): complete subsection reference.

- [medium_security](data-sources--nfv_service--reference--group-002.md#canonical-5a758c7b15322e9b67e0355ff0205986c4da0a58e5d56669fe1ed6653d6eb0a4): complete subsection reference.

<a id="canonical-3b50a561f659df520865da4e531b0713e4ae28871598284730a59f8ff4e2f988"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_config / 5e71873c63bd / 4

- [https_management.advertise_on_slo_internet_vip.tls_config.custom_security](data-sources--nfv_service--reference--group-002.md#canonical-bd43e3aa4659d130acc5185619122d9bdd7b5f7c38b05ffb3de88f0717dda7a1)
- [https_management.advertise_on_slo_internet_vip.tls_config.default_security](data-sources--nfv_service--reference--group-002.md#canonical-092e2d7f8dfc739d12fd29d398ae12daa4d832b550097bdfec6c6b51bfe6dc15)
- [https_management.advertise_on_slo_internet_vip.tls_config.low_security](data-sources--nfv_service--reference--group-002.md#canonical-4948f67d8b0d5444b8ed3373a96f91ea6b493284a1f7386cd15e506cbc253837)
- [https_management.advertise_on_slo_internet_vip.tls_config.medium_security](data-sources--nfv_service--reference--group-002.md#canonical-5a758c7b15322e9b67e0355ff0205986c4da0a58e5d56669fe1ed6653d6eb0a4)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-bd43e3aa4659d130acc5185619122d9bdd7b5f7c38b05ffb3de88f0717dda7a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0600c0ccf8957d71d8c680c528a78b433606332319f61d0bcbdb9c3ee925330a"></a>

## https_management.advertise_on_slo_internet_vip.tls_config.custom_security — https_management.advertise_on_slo_internet_vip.tls_config.custom_security / c10962180388 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-78f0ca926c5842a0af13f5d6e2f298a65322948981f7d069e7ade2213d9b77c7)
- https_management.advertise_on_slo_internet_vip.tls_config.custom_security

<a id="canonical-5092810a7ff67352951137c0fdcaf00fe53864aa4a57a3b23ecf4f7aceee2b6b"></a>

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

<a id="canonical-95ff59f74fd9581d8ae5866b0b8f78fb4e0931776b8e85092756257612367b65"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_config.custom_security / c10962180388 / 3

<a id="canonical-86dffbd531e3d3411162cd922b28aa95ab044c705fbb787711140993f14bc994"></a>

<a id="canonical-87308410d7445cd510d6bcbfad3d5ad35e741f86038b25bfe513a4a670cd59d6"></a>

## cipher_suites property — https_management.advertise_on_slo_internet_vip.tls_config.custom_security / c10962180388 / 4

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

<a id="canonical-23df65288c70a8b0075425da31f115352859b190f7969bd8f6641e0628caa062"></a>

<a id="canonical-c51243d1572b84c31ed1fb81b1b40d89e290933fefe62663b6e1496b7629c863"></a>

## max_version property — https_management.advertise_on_slo_internet_vip.tls_config.custom_security / c10962180388 / 5

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

<a id="canonical-2c68acb00eed048ae4b679e196ede55d4015fc298043a82c29deb2862729acc7"></a>

<a id="canonical-9d15587f9d5a1c101e48b4ffd88ff61249d03a07ac054201cecaf2a4a614c106"></a>

## min_version property — https_management.advertise_on_slo_internet_vip.tls_config.custom_security / c10962180388 / 6

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

<a id="canonical-e7cadf524403a62d4dc6591ece705a584a8aacaae818db1fde4ca77aea427adc"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_config.custom_security / c10962180388 / 7

- [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-78f0ca926c5842a0af13f5d6e2f298a65322948981f7d069e7ade2213d9b77c7)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-092e2d7f8dfc739d12fd29d398ae12daa4d832b550097bdfec6c6b51bfe6dc15"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ef4dfb84a095fc262589ca8b75b23f129fcfef642cdaea451dc67bc1084b6f0"></a>

## https_management.advertise_on_slo_internet_vip.tls_config.default_security — https_management.advertise_on_slo_internet_vip.tls_config.default_security / 1a3c738f667b / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-78f0ca926c5842a0af13f5d6e2f298a65322948981f7d069e7ade2213d9b77c7)
- https_management.advertise_on_slo_internet_vip.tls_config.default_security

<a id="canonical-c49cbddb6ee07ecf2f3c0951f28727334a4cee27db21f534935886237dfb6920"></a>

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

<a id="canonical-34034aca922682eb72819e7518e1668919c4439cae697c1f5bdf53b0a9ab09ce"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_config.default_security / 1a3c738f667b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8ba63bb1fd40d8fadb99b16342b91b9b8771cb9391e1a17d044279eb090c98da"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_config.default_security / 1a3c738f667b / 4

- [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-78f0ca926c5842a0af13f5d6e2f298a65322948981f7d069e7ade2213d9b77c7)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-4948f67d8b0d5444b8ed3373a96f91ea6b493284a1f7386cd15e506cbc253837"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-032887470d2bee13b4a4a0fe71ccd7648fdcb4a2a9263512fb2acf572d3cc2bf"></a>

## https_management.advertise_on_slo_internet_vip.tls_config.low_security — https_management.advertise_on_slo_internet_vip.tls_config.low_security / 2d1b041ad7c1 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-78f0ca926c5842a0af13f5d6e2f298a65322948981f7d069e7ade2213d9b77c7)
- https_management.advertise_on_slo_internet_vip.tls_config.low_security

<a id="canonical-e656f410b6dd4d0919cb9ea79b14deba73ab36e54194e73cd1459b1dcaedf5fb"></a>

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

<a id="canonical-ca9b9498c94d4c91490085235370be45a3bcd51513aa2e939372c3cde6dbb8ea"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_config.low_security / 2d1b041ad7c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cdf73bdf202cb10c57f7a21029f3b059eb6b7306b48074963544bb9ab1741182"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_config.low_security / 2d1b041ad7c1 / 4

- [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-78f0ca926c5842a0af13f5d6e2f298a65322948981f7d069e7ade2213d9b77c7)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-5a758c7b15322e9b67e0355ff0205986c4da0a58e5d56669fe1ed6653d6eb0a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3f778e544a3a94e9cc9affd22b6155086c4f189f7d05363d2fbcd38d4bc5ede"></a>

## https_management.advertise_on_slo_internet_vip.tls_config.medium_security — https_management.advertise_on_slo_internet_vip.tls_config.medium_security / ebfe2a53b22a / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-78f0ca926c5842a0af13f5d6e2f298a65322948981f7d069e7ade2213d9b77c7)
- https_management.advertise_on_slo_internet_vip.tls_config.medium_security

<a id="canonical-f0334b77f00644aad38200fc7ef1bd243d259ce4ed7c397688c3b5e40151fefd"></a>

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

<a id="canonical-f0932b9da6de7324fa676bb5715fe3e26c6b11c588f59079f9d426ce5b8077a3"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_config.medium_security / ebfe2a53b22a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4f549c853db5f402c6e19efb1f1bbd3a7d8af70217eaec2ba4b985e65d99085"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_config.medium_security / ebfe2a53b22a / 4

- [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-78f0ca926c5842a0af13f5d6e2f298a65322948981f7d069e7ade2213d9b77c7)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-9fd77b3bb6352a172d395c7191f2b909686330ea314aec3eef82dc410240657a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fa664d4639912473ee2e00e7bff969fe190301aedd559b04faf3a9516d814c3"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls — https_management.advertise_on_slo_internet_vip.use_mtls / 9d5c722b0d1f / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- https_management.advertise_on_slo_internet_vip.use_mtls

<a id="canonical-f74056e015954befeae3133f9c2fff765468ee360b882e0e623580a210df359d"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

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

<a id="canonical-e0325488bc7648e095c2df815ccff638f030e86789671a9f6060dbc440a1ebcf"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.use_mtls / 9d5c722b0d1f / 3

<a id="canonical-6bc84d02ceb3904b319d06b34f2f6f133da3c71b33703c0ab9607f4e2d722086"></a>

<a id="canonical-384c567138e414e1d8f9f4afc4dda69d032a3e1657c20055ae00d6fa5cae52e9"></a>

## client_certificate_optional property — https_management.advertise_on_slo_internet_vip.use_mtls / 9d5c722b0d1f / 4

Type: `"bool"`. Computed.

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

- [crl](data-sources--nfv_service--reference--group-002.md#canonical-66dc51a9f24f0f88e5d6aca234cec88690018782294555cdaee98582b441d077): complete subsection reference.

- [no_crl](data-sources--nfv_service--reference--group-002.md#canonical-8b0f49a43f85f354428be82e2aa84ec60ce15a9bfec227571214d2cf23965ec9): complete subsection reference.

- [trusted_ca](data-sources--nfv_service--reference--group-002.md#canonical-8da4c19b79394c1de58b6be272a8cc7e61ab4f6e69da23690608af37b8958482): complete subsection reference.

<a id="canonical-23c23911c6c389cb744fef52ca147b2556e2566833649e9228196a7a09025762"></a>

<a id="canonical-26a8f166dbdc583e53c1d6a9dbdf190d29bc3b02c259910ec6b6391866aed97d"></a>

## trusted_ca_url property — https_management.advertise_on_slo_internet_vip.use_mtls / 9d5c722b0d1f / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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

- [xfcc_disabled](data-sources--nfv_service--reference--group-002.md#canonical-5a2a019898c56885f0a91852a7b7bc987b640bd93db5177a32dbb5cdad326c87): complete subsection reference.

- [xfcc_options](data-sources--nfv_service--reference--group-002.md#canonical-dfa1bc020e18dfe472985361b11b48c21892a338877a6f10cefedf3e7704be23): complete subsection reference.

<a id="canonical-876e53980f9a7998b3c775e0001fb83b599118c87966906915c3b0588638b5c7"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.use_mtls / 9d5c722b0d1f / 6

- [https_management.advertise_on_slo_internet_vip.use_mtls.crl](data-sources--nfv_service--reference--group-002.md#canonical-66dc51a9f24f0f88e5d6aca234cec88690018782294555cdaee98582b441d077)
- [https_management.advertise_on_slo_internet_vip.use_mtls.no_crl](data-sources--nfv_service--reference--group-002.md#canonical-8b0f49a43f85f354428be82e2aa84ec60ce15a9bfec227571214d2cf23965ec9)
- [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca](data-sources--nfv_service--reference--group-002.md#canonical-8da4c19b79394c1de58b6be272a8cc7e61ab4f6e69da23690608af37b8958482)
- [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled](data-sources--nfv_service--reference--group-002.md#canonical-5a2a019898c56885f0a91852a7b7bc987b640bd93db5177a32dbb5cdad326c87)
- [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options](data-sources--nfv_service--reference--group-002.md#canonical-dfa1bc020e18dfe472985361b11b48c21892a338877a6f10cefedf3e7704be23)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-66dc51a9f24f0f88e5d6aca234cec88690018782294555cdaee98582b441d077"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f87922ba66678fba2b841d67e68cb48755f7c5bfd752e26238fbed2912164f62"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.crl — https_management.advertise_on_slo_internet_vip.use_mtls.crl / e34e763ae2fe / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-9fd77b3bb6352a172d395c7191f2b909686330ea314aec3eef82dc410240657a)
- https_management.advertise_on_slo_internet_vip.use_mtls.crl

<a id="canonical-5049403c77ecc9172cbbd8522860ca499467bcf067085314ba31aa8f7d1db448"></a>

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

<a id="canonical-89732a1c807077293fe5d21034b41c8b8987259b4ddd0bc06fa54e8116080bd1"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.use_mtls.crl / e34e763ae2fe / 3

<a id="canonical-66791cf4e7253d028830b2352a5a7d33663d33c762d4aa197803c4ec80f45d26"></a>

<a id="canonical-ea4219669e9d16e351c8c49208fb57af5d4f6b265e563ac052e2448f71546841"></a>

## name property — https_management.advertise_on_slo_internet_vip.use_mtls.crl / e34e763ae2fe / 4

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

<a id="canonical-946413cd17ad9650bd70489eb7be231003aa9f7861f65a910ef0366cb53fc7a3"></a>

<a id="canonical-979ba8e4b9cf0f06fed7d2795c660dd88a577c3cc0f535d9d98ee1e87e1346cb"></a>

## namespace property — https_management.advertise_on_slo_internet_vip.use_mtls.crl / e34e763ae2fe / 5

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

<a id="canonical-b2062e6e7e3b0136eaa73d7bef9074a753f8b7f623b111bd0de1df625588c4d7"></a>

<a id="canonical-b1d3d43831a106f2b64290878653d3d91cf86f2c9bf6b14e958552fa1861184d"></a>

## tenant property — https_management.advertise_on_slo_internet_vip.use_mtls.crl / e34e763ae2fe / 6

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

<a id="canonical-78ae0796823d13967b4bceddc719048025ca319a6141bc04e6f8243d1eb668ac"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.use_mtls.crl / e34e763ae2fe / 7

- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-9fd77b3bb6352a172d395c7191f2b909686330ea314aec3eef82dc410240657a)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-8b0f49a43f85f354428be82e2aa84ec60ce15a9bfec227571214d2cf23965ec9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-059869737fdcf5f4ca3814856dbc1f85453a0a38cd2edb4af365e7256ec7cab0"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.no_crl — https_management.advertise_on_slo_internet_vip.use_mtls.no_crl / b1a9b57b29d3 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-9fd77b3bb6352a172d395c7191f2b909686330ea314aec3eef82dc410240657a)
- https_management.advertise_on_slo_internet_vip.use_mtls.no_crl

<a id="canonical-ff1d63e4c23a1e9aba6434d44e99f8801f5e937c2337767981a30fffdba41ebb"></a>

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

<a id="canonical-caf82472dc3167c888d9b2fae11a364874cf3404ca41b81d54d09a438c2a9377"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.use_mtls.no_crl / b1a9b57b29d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e7165c1a8c3a5f38ebb39681159069dadc95c4dedd19c6fa8b52e639b092d81c"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.use_mtls.no_crl / b1a9b57b29d3 / 4

- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-9fd77b3bb6352a172d395c7191f2b909686330ea314aec3eef82dc410240657a)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-8da4c19b79394c1de58b6be272a8cc7e61ab4f6e69da23690608af37b8958482"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1af34b0c229b20b9a539ac9e4ff2cb96f476f23f4b78a3e90edcd9834937603"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca — https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca / 477094841dfc / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-9fd77b3bb6352a172d395c7191f2b909686330ea314aec3eef82dc410240657a)
- https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca

<a id="canonical-f48000760af114e7332429e399979158e311e69b1305778ac49a4997d15d5b84"></a>

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

<a id="canonical-3ba8ad64815ffb7f01e20514a291e62f681bea11e7a6ba85b3dfe42c545c25e9"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca / 477094841dfc / 3

<a id="canonical-c88c0b322bc1bb82a37f003b60dce495fd9ac2be51a4f5e844d263a0bdd72cba"></a>

<a id="canonical-06c28896eb584619fabaca047c0cfc300768403e923f6d025a7473209af060c7"></a>

## name property — https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca / 477094841dfc / 4

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

<a id="canonical-5fc2ac55442c4fd23f87fd7109a528b9c70521763df8f4cc97e504d6b36385ab"></a>

<a id="canonical-2372dd6116cd4d3b57927298b9cef24effae87883e50ac719a9d010831ca007d"></a>

## namespace property — https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca / 477094841dfc / 5

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

<a id="canonical-0cca91838fbb30d4a4da8677d778af4abf6606b95d849b864593b4d395ce92d5"></a>

<a id="canonical-d094b94c39febb0cf777ac0fe01811c2beddbdc48589cf564928c0d6331ca9d5"></a>

## tenant property — https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca / 477094841dfc / 6

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

<a id="canonical-30dff36a91cb3886de89a6097015f10d1427a77a7552bd67ea93ae68826835c2"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca / 477094841dfc / 7

- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-9fd77b3bb6352a172d395c7191f2b909686330ea314aec3eef82dc410240657a)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-5a2a019898c56885f0a91852a7b7bc987b640bd93db5177a32dbb5cdad326c87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bfb022c0a0be08c5ffc02fe88e9ea63956baa2f99dc47980e66fb41e08e53af9"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled — https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled / e96e4acf1cd5 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-9fd77b3bb6352a172d395c7191f2b909686330ea314aec3eef82dc410240657a)
- https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled

<a id="canonical-eb6cc07a4dc44dd20153e369d7f25bf58e171e1c2282760700e8aa6f6af62c7a"></a>

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

<a id="canonical-3a9f06d276d0473727603cb8cdae8284347ef8c4e14d18ac91a249fc1086ea57"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled / e96e4acf1cd5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9eb9d783feed7d00d6f59711528f4a39557dd26342dab90963ab3452dc939e5d"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled / e96e4acf1cd5 / 4

- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-9fd77b3bb6352a172d395c7191f2b909686330ea314aec3eef82dc410240657a)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-dfa1bc020e18dfe472985361b11b48c21892a338877a6f10cefedf3e7704be23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f686ea8d375038bc98cf32146c4629542da814ce4c344c5e6f5859494eb2ef66"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options — https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options / 8a1d7904cdc9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-14b050436def827a452ae02bb8fb652aad8cdbadd4acffd6b4d0d0cdf5b7d41f)
- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-9fd77b3bb6352a172d395c7191f2b909686330ea314aec3eef82dc410240657a)
- https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options

<a id="canonical-26fa4dcd0e07437257e9e5f91981cbc73fa24ebd88fd73eb389ef40d7ebc7633"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d3903d11c2216fd3dd1484cb096fd0d9c707d45237e013b73927b0f700c6e14d"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options / 8a1d7904cdc9 / 3

<a id="canonical-8510ce34115352996363da133753aafdc95377533fb71651a3f13cc0c15f8df6"></a>

<a id="canonical-2b0e0c0f874fc76f31cb3afa0c095d704338c574cd521594b05b9699eadc98b7"></a>

## xfcc_header_elements property — https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options / 8a1d7904cdc9 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-b12bdf654c8e02dfded30d6eed2958dad0a20b9d74a204047384cb86f25e3811"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options / 8a1d7904cdc9 / 5

- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-9fd77b3bb6352a172d395c7191f2b909686330ea314aec3eef82dc410240657a)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f094e3a12112fd1833c7dd483cf08da930e911a5d0596dc9b798e447d202409"></a>

## https_management.advertise_on_slo_sli — https_management.advertise_on_slo_sli / 9d3eb1d97099 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- https_management.advertise_on_slo_sli

<a id="canonical-dd95e1072b6d17e62d99dbf6406caa7f414e7b01b19a15e08fb71b759342f545"></a>

Type: `"single"`. Computed.

Configuration parameter for advertise on slo sli.

Upstream description:

Inline TLS parameters.

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

<a id="canonical-baf7914084b000135135fb1999d84d8bc55c2cba868761a8bafe356cbeb2013e"></a>

## Direct properties — https_management.advertise_on_slo_sli / 9d3eb1d97099 / 3

- [no_mtls](data-sources--nfv_service--reference--group-002.md#canonical-862bec11fbce1c877bc6a82ed0cac2dc7c9a6f0d87193e08cb2b875ce9e713d4): complete subsection reference.

- [tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-5964b4cebf3482d0880b7c54d204993a7c0648f12baf3a8071e4184a0f93bbea): complete subsection reference.

- [tls_config](data-sources--nfv_service--reference--group-003.md#canonical-079f523bfa7f99351fd61f57cc767cb2eb1e33ea71ab5e84e2155d35d598ba92): complete subsection reference.

- [use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-a81b93841729eecca0b56bf195a182d25d620ef5fb5b7bd2eebf84ece5ac3601): complete subsection reference.

<a id="canonical-7b3932acbab45445f556a0a006a965e637b52bac68f2a21bfad5f2281e2077e2"></a>

## Next pages — https_management.advertise_on_slo_sli / 9d3eb1d97099 / 4

- [https_management.advertise_on_slo_sli.no_mtls](data-sources--nfv_service--reference--group-002.md#canonical-862bec11fbce1c877bc6a82ed0cac2dc7c9a6f0d87193e08cb2b875ce9e713d4)
- [https_management.advertise_on_slo_sli.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-5964b4cebf3482d0880b7c54d204993a7c0648f12baf3a8071e4184a0f93bbea)
- [https_management.advertise_on_slo_sli.tls_config](data-sources--nfv_service--reference--group-003.md#canonical-079f523bfa7f99351fd61f57cc767cb2eb1e33ea71ab5e84e2155d35d598ba92)
- [https_management.advertise_on_slo_sli.use_mtls](data-sources--nfv_service--reference--group-003.md#canonical-a81b93841729eecca0b56bf195a182d25d620ef5fb5b7bd2eebf84ece5ac3601)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-862bec11fbce1c877bc6a82ed0cac2dc7c9a6f0d87193e08cb2b875ce9e713d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d70489a766108791dbfec1d7865e9499493aa5a7190963f246ab276d26fd6901"></a>

## https_management.advertise_on_slo_sli.no_mtls — https_management.advertise_on_slo_sli.no_mtls / cbd3b0003f25 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- https_management.advertise_on_slo_sli.no_mtls

<a id="canonical-796ad0dc70e9b53048a21e59bf7804f28777b2feada1e7abb8420eb03da0decf"></a>

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

<a id="canonical-5decfcb2d1d8041dafb68a5c1ec2b9700dfb1ba152a6458a571233ca49217fcd"></a>

## Direct properties — https_management.advertise_on_slo_sli.no_mtls / cbd3b0003f25 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-97a80ab7d555b92341843161f06448873bde14ff5a9833d0ad84a07265157c25"></a>

## Next pages — https_management.advertise_on_slo_sli.no_mtls / cbd3b0003f25 / 4

- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-5964b4cebf3482d0880b7c54d204993a7c0648f12baf3a8071e4184a0f93bbea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edabe288741c2dde77bf4f2e8a6e4a4b2214822f7895366b5fb93db03f13ae68"></a>

## https_management.advertise_on_slo_sli.tls_certificates — https_management.advertise_on_slo_sli.tls_certificates / 4eb18c169d46 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-b713feddc432dd59b1ed8c061ae4887dc8ae5a0a106502c5bbe75c3b1a41e0c8)
- [https_management](data-sources--nfv_service--reference--group-002.md#canonical-3aad6fe5fb30d6951d0aa769826fe1b8075876088c48548c08b041462dc4fb56)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--reference--group-002.md#canonical-086d88128f5e5786a03853766b2c2649c0af61792f065c9bcfa6cdbb298fd85f)
- https_management.advertise_on_slo_sli.tls_certificates

<a id="canonical-ea2fdbf237cbafca299b4bd2e0aacd949b2bf3bcd7c6df7a7e3e6537ee2891ad"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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

<a id="canonical-cfee35ef412e84f7db237b582dd321501f35907f486d18905107679e673d4cde"></a>

## Direct properties — https_management.advertise_on_slo_sli.tls_certificates / 4eb18c169d46 / 3

<a id="canonical-04f831a43d83c9399f1616c67cd23f2873ae036d511ca061986ae98b93e03daf"></a>
