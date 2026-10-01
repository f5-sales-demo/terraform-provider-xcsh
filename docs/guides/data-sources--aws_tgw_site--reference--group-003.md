---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-7868772f6fddcea45c888fe39d488d4f601d651a32024223ce284eaf32aaad17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b243248a9c75e45a200979d28bd41b010a806fd0c7cb489c5614b2581358af7"></a>

## vn_config.allowed_vip_port_sli — vn_config.allowed_vip_port_sli / 2c4e89448f68 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- vn_config.allowed_vip_port_sli

<a id="canonical-86114df20d333a69267033a41329347c5917e8e4fab5f068240efd85a199e896"></a>

Type: `"single"`. Computed.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

<a id="canonical-ea147bf626649060af9f251b724cdaf0674a4f9bff101d520f22aa65c005a507"></a>

## Direct properties — vn_config.allowed_vip_port_sli / 2c4e89448f68 / 3

- [custom_ports](data-sources--aws_tgw_site--reference--group-003.md#canonical-42be77dc1ca772607baabcb098cb6724e9447e63a6a2b3a676180d5bdf94a4c4): complete subsection reference.

- [disable_allowed_vip_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-129dfe4c22e6a0f1956780d8e077e3b63b7bf6cd713731d2c3a534e64a9f24bb): complete subsection reference.

- [use_http_https_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-83cd82d833ac7e2d66811ab922e0e609567f8571f309389864e9ed16d8ec8cb2): complete subsection reference.

- [use_http_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-fc4baa26d1c6a22a1507f4e2af2d91ca41419ed0ab4d1ac9a9d784767c3da1d1): complete subsection reference.

- [use_https_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-753de8576754806eb7123493748b5d4851f108966e1c1343b1aebdd7ab000a7b): complete subsection reference.

<a id="canonical-78e04fe97d0763baaa3c8e46b74376aeca624ef94cfd29f657122ea0212f7477"></a>

## Next pages — vn_config.allowed_vip_port_sli / 2c4e89448f68 / 4

- [vn_config.allowed_vip_port_sli.custom_ports](data-sources--aws_tgw_site--reference--group-003.md#canonical-42be77dc1ca772607baabcb098cb6724e9447e63a6a2b3a676180d5bdf94a4c4)
- [vn_config.allowed_vip_port_sli.disable_allowed_vip_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-129dfe4c22e6a0f1956780d8e077e3b63b7bf6cd713731d2c3a534e64a9f24bb)
- [vn_config.allowed_vip_port_sli.use_http_https_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-83cd82d833ac7e2d66811ab922e0e609567f8571f309389864e9ed16d8ec8cb2)
- [vn_config.allowed_vip_port_sli.use_http_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-fc4baa26d1c6a22a1507f4e2af2d91ca41419ed0ab4d1ac9a9d784767c3da1d1)
- [vn_config.allowed_vip_port_sli.use_https_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-753de8576754806eb7123493748b5d4851f108966e1c1343b1aebdd7ab000a7b)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-42be77dc1ca772607baabcb098cb6724e9447e63a6a2b3a676180d5bdf94a4c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85e1dfd113b9b195ff6956ccf27026608b443725f77f1f89ec86b90cab2360e1"></a>

## vn_config.allowed_vip_port_sli.custom_ports — vn_config.allowed_vip_port_sli.custom_ports / 8ef24d139b37 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-7868772f6fddcea45c888fe39d488d4f601d651a32024223ce284eaf32aaad17)
- vn_config.allowed_vip_port_sli.custom_ports

<a id="canonical-9e8db0b2ca68c49489ed9ce4c305ff6b83b561045c87db275c4adab5a8b2796b"></a>

Type: `"single"`. Computed.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5c4f23c146df67be7647fc1972bcf52f67de4b5507b91dbd40306ef554863496"></a>

## Direct properties — vn_config.allowed_vip_port_sli.custom_ports / 8ef24d139b37 / 3

<a id="canonical-a0097a7edab31d767fcbbd9597517b6dd499163a5c7607ff194e4077802a3e5a"></a>

<a id="canonical-5a883f21770e2b2db9fa0a4887fb9788399df9a0375572cfaf2281c3a50ba3e5"></a>

## port_ranges property — vn_config.allowed_vip_port_sli.custom_ports / 8ef24d139b37 / 4

Type: `"string"`. Computed.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-b8d825c1c366cd8b5ec13de197bd58e446c7c6e1171008a96f6c25721a4940f7"></a>

## Next pages — vn_config.allowed_vip_port_sli.custom_ports / 8ef24d139b37 / 5

- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-7868772f6fddcea45c888fe39d488d4f601d651a32024223ce284eaf32aaad17)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-129dfe4c22e6a0f1956780d8e077e3b63b7bf6cd713731d2c3a534e64a9f24bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7cd9df14ffe97a199898a61ef61f5325d9dfc561ad9128314e8d15036dfb6c62"></a>

## vn_config.allowed_vip_port_sli.disable_allowed_vip_port — vn_config.allowed_vip_port_sli.disable_allowed_vip_port / b15a3d93fc9e / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-7868772f6fddcea45c888fe39d488d4f601d651a32024223ce284eaf32aaad17)
- vn_config.allowed_vip_port_sli.disable_allowed_vip_port

<a id="canonical-df2e9c3a0ab5d181c8f5cefda3c71c540726e9b585a967e9518415a1d644b743"></a>

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

<a id="canonical-31a793d1ad0fb32f069923f5133e91c2ca1d8cdcddbfacfb5f68a221d146bd25"></a>

## Direct properties — vn_config.allowed_vip_port_sli.disable_allowed_vip_port / b15a3d93fc9e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d94c385dab2773dd839fb4c7ce0eec3ede2e0481b8a834bbaba2ef94bb36d47d"></a>

## Next pages — vn_config.allowed_vip_port_sli.disable_allowed_vip_port / b15a3d93fc9e / 4

- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-7868772f6fddcea45c888fe39d488d4f601d651a32024223ce284eaf32aaad17)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-83cd82d833ac7e2d66811ab922e0e609567f8571f309389864e9ed16d8ec8cb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43151269f85e8b6dbb7583561d204786f57ec0bba07b11b5172669d8a1b44a75"></a>

## vn_config.allowed_vip_port_sli.use_http_https_port — vn_config.allowed_vip_port_sli.use_http_https_port / 8563da3376a6 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-7868772f6fddcea45c888fe39d488d4f601d651a32024223ce284eaf32aaad17)
- vn_config.allowed_vip_port_sli.use_http_https_port

<a id="canonical-18e65baba2dd97984b46eee7f81f110c93dbaa6669ca3c2bfe4aece4213a0419"></a>

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

<a id="canonical-233de5434796be24134f042ca2b4d2180ce26bae220b2dd04e8b5557178dd4fb"></a>

## Direct properties — vn_config.allowed_vip_port_sli.use_http_https_port / 8563da3376a6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dcda91e12249156607048538a154eaeb630642dd147c200c5890547911068fad"></a>

## Next pages — vn_config.allowed_vip_port_sli.use_http_https_port / 8563da3376a6 / 4

- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-7868772f6fddcea45c888fe39d488d4f601d651a32024223ce284eaf32aaad17)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-fc4baa26d1c6a22a1507f4e2af2d91ca41419ed0ab4d1ac9a9d784767c3da1d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6a1258345ef0b371e489e8d19235ff1101ae7b2729f276170b7ff85f5257a52"></a>

## vn_config.allowed_vip_port_sli.use_http_port — vn_config.allowed_vip_port_sli.use_http_port / 281cf3cc915c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-7868772f6fddcea45c888fe39d488d4f601d651a32024223ce284eaf32aaad17)
- vn_config.allowed_vip_port_sli.use_http_port

<a id="canonical-2582d626f45d902c0835897f699be2bf9ea5df58d3564530ab63d0c3904f8f80"></a>

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

<a id="canonical-6303458e4efaa95a136a0c177b3bd4cd2b0c8b79592783a7287bbc2a5663f8c4"></a>

## Direct properties — vn_config.allowed_vip_port_sli.use_http_port / 281cf3cc915c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0403987ddda600f0f20c81d5a4fd1af3d43046142b92f58c3bfc2e4a5eeb8be0"></a>

## Next pages — vn_config.allowed_vip_port_sli.use_http_port / 281cf3cc915c / 4

- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-7868772f6fddcea45c888fe39d488d4f601d651a32024223ce284eaf32aaad17)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-753de8576754806eb7123493748b5d4851f108966e1c1343b1aebdd7ab000a7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6ca82a9b75ad1c1d405c34c6f6a4e1c62bbf995a84395cc244ef126df4f702b"></a>

## vn_config.allowed_vip_port_sli.use_https_port — vn_config.allowed_vip_port_sli.use_https_port / bc16efc9c6c3 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-7868772f6fddcea45c888fe39d488d4f601d651a32024223ce284eaf32aaad17)
- vn_config.allowed_vip_port_sli.use_https_port

<a id="canonical-71f76f1657caa45ff53300830a5a0713f98932198d2c452dbb26ce6a62e9449d"></a>

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

<a id="canonical-e8e35acbb6d329e5318a164845745943b5a755119f39698d63469db2d955d30a"></a>

## Direct properties — vn_config.allowed_vip_port_sli.use_https_port / bc16efc9c6c3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-19ac98eabbd02cd590d497f1756ed65f94f766ce0f1bce13ba5c34e957acac9f"></a>

## Next pages — vn_config.allowed_vip_port_sli.use_https_port / bc16efc9c6c3 / 4

- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-7868772f6fddcea45c888fe39d488d4f601d651a32024223ce284eaf32aaad17)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-1b86e7c2f865a154fe26a67ddf1bbf01f7b8691b43ef36c63f9ab856d5dae408"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37d0f8a65e840225fbdce5ae4ebb8ee5b674fa88cd766ee77a7c9661b9c00e8a"></a>

## vn_config.dc_cluster_group_inside_vn — vn_config.dc_cluster_group_inside_vn / a465f4b3c67c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- vn_config.dc_cluster_group_inside_vn

<a id="canonical-07a495ad6894a04082761fe9a7e87646827c3aba90ff1e864cbd384b219d98c9"></a>

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

<a id="canonical-3c1c06e26207f8fb11a63af8a19683d290c3c723862ca8f7b09c620f723790b8"></a>

## Direct properties — vn_config.dc_cluster_group_inside_vn / a465f4b3c67c / 3

<a id="canonical-1c7061deb0bee5932cacd13b179aef2f77d52916d17e2e8ea8f33762d978cdb7"></a>

<a id="canonical-ccaa096bd91ed43a19acc8625947b182eca64ebaa91a56299e39445073f7e562"></a>

## name property — vn_config.dc_cluster_group_inside_vn / a465f4b3c67c / 4

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

<a id="canonical-b338afe8f8f729e23f5430480e947ea1052c3f556feef7c9826ed8ffe8f2fb71"></a>

<a id="canonical-f1d7de9677ee4c94b7eff9314f1c49312e047e80f2f174e6cb662273b9700115"></a>

## namespace property — vn_config.dc_cluster_group_inside_vn / a465f4b3c67c / 5

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

<a id="canonical-8f946d5c0338a5345e2bc37e35e44f8e73b3d9ac2014c8647d5c1c12cbc45fd0"></a>

<a id="canonical-cc1ec0e082306dc85552cfac8900064c53c93b5df9117a9696e51d7aae702200"></a>

## tenant property — vn_config.dc_cluster_group_inside_vn / a465f4b3c67c / 6

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

<a id="canonical-663580e0112bb26fa7f7ac10e0779397273cf03d676e0c6138611294dd852402"></a>

## Next pages — vn_config.dc_cluster_group_inside_vn / a465f4b3c67c / 7

- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-c61185272959ffb80f3d668e931bee0845ed0670570d5d5fdb0fd6cec8e70922"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20b565b5a3ee455e2bec1e03f220cf73167d6b58c6dbc03a728d276c6e1cecf2"></a>

## vn_config.dc_cluster_group_outside_vn — vn_config.dc_cluster_group_outside_vn / 57cd0ff9b99c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- vn_config.dc_cluster_group_outside_vn

<a id="canonical-2f7c595e9d8e123aec5deacc82185a031b73a22f98562ad758131e0c88110b46"></a>

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

<a id="canonical-e2c0bb2bcdab648f4a3e17cfe8bee243abe5dd066b5750ee4ea9573537dd1c54"></a>

## Direct properties — vn_config.dc_cluster_group_outside_vn / 57cd0ff9b99c / 3

<a id="canonical-9f4094dd7c3171c0f83ea52e2ed78c4d77a84e98dd9a2fe35931a04af21af16e"></a>

<a id="canonical-731fafc75754e11380c185c00d7d445fe74d77a01397ca14e4279c15c226895e"></a>

## name property — vn_config.dc_cluster_group_outside_vn / 57cd0ff9b99c / 4

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

<a id="canonical-8d8624961aac15ee3efa59df58d5576e4af525de215b7eab987ecad07363121b"></a>

<a id="canonical-99d2bac21108a5947b4a5a4973bc2490d8ce36dbf6b450b370e3881d20a6c08c"></a>

## namespace property — vn_config.dc_cluster_group_outside_vn / 57cd0ff9b99c / 5

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

<a id="canonical-8cbeef19c903b9afdc8cf2f1826fd7c8ed1bdeabaea2d6cc03194b89f7e15b26"></a>

<a id="canonical-fba1e050394e635e05c11945cbe9079e0b3c8086a06b11c44b7fb4dc8d006e87"></a>

## tenant property — vn_config.dc_cluster_group_outside_vn / 57cd0ff9b99c / 6

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

<a id="canonical-8b867d7f2fcf35e0c17bbe26f7c8ff0b0c2ff440b7f1ebbadbacf7cccd7317af"></a>

## Next pages — vn_config.dc_cluster_group_outside_vn / 57cd0ff9b99c / 7

- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-52586cb0d4e0ac88914ed7ecafc62403594f30a09f9862c374016a7f7e24f563"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-611598a6ebd797e3dfdd03d0870606f2b30acc70dcda756f9fb54aadf59e222c"></a>

## vn_config.global_network_list — vn_config.global_network_list / 1657d6efca8b / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- vn_config.global_network_list

<a id="canonical-19d95146fb7ff7e344dcf1febec3f9f7944a495af462f8785a9a0c5cf6472991"></a>

Type: `"single"`. Computed.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a3111b62b7f09932718f30c9705d7e8644f7a9d86492dde19b1df7db666bbcc4"></a>

## Direct properties — vn_config.global_network_list / 1657d6efca8b / 3

- [global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-ee6206d778b01df7d2a7798364d5bff4165511d30806de46944cb3579691ca12): complete subsection reference.

<a id="canonical-f4a4c0e33b022b453eb5d1237345de50dce129594e51938d762d44f00291228e"></a>

## Next pages — vn_config.global_network_list / 1657d6efca8b / 4

- [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-ee6206d778b01df7d2a7798364d5bff4165511d30806de46944cb3579691ca12)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-ee6206d778b01df7d2a7798364d5bff4165511d30806de46944cb3579691ca12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9926fa5b2e194f7f2b0ef2c73411caafa26f6a523839838647c991586bad9f70"></a>

## vn_config.global_network_list.global_network_connections — vn_config.global_network_list.global_network_connections / 5ea1c2b26869 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-52586cb0d4e0ac88914ed7ecafc62403594f30a09f9862c374016a7f7e24f563)
- vn_config.global_network_list.global_network_connections

<a id="canonical-ba255be613fdae79944d031ba220f074560c4e3b5e77ab84c8d0e271597a63a4"></a>

Type: `"list"`. Computed.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-d6f0c2d0c9643355ac2d1c1134ca2309c4378aa8521590156285e1f154757b21"></a>

## Direct properties — vn_config.global_network_list.global_network_connections / 5ea1c2b26869 / 3

- [sli_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-f19f48e60638282949ae91737b2f53d9decad95a3e5153c7f594ed7cbce5b4e8): complete subsection reference.

- [slo_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-b2714247159432d79fb3ae49cebf317577383cb2ffa4bf87affa300bfd8c451d): complete subsection reference.

<a id="canonical-d45b48bbe47c9be188a0190a5c4cd79fe03135a1782bc5b24b8f4b4602cf9727"></a>

## Next pages — vn_config.global_network_list.global_network_connections / 5ea1c2b26869 / 4

- [vn_config.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-f19f48e60638282949ae91737b2f53d9decad95a3e5153c7f594ed7cbce5b4e8)
- [vn_config.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-b2714247159432d79fb3ae49cebf317577383cb2ffa4bf87affa300bfd8c451d)
- [vn_config.global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-52586cb0d4e0ac88914ed7ecafc62403594f30a09f9862c374016a7f7e24f563)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-f19f48e60638282949ae91737b2f53d9decad95a3e5153c7f594ed7cbce5b4e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d91e2dc91eaf7a30ca4ecc76921635d79f6eec3e3aeaf4602ab4f8d2e10e6a8"></a>

## vn_config.global_network_list.global_network_connections.sli_to_global_dr — vn_config.global_network_list.global_network_connections.sli_to_global_dr / 77ddb15cba8f / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-52586cb0d4e0ac88914ed7ecafc62403594f30a09f9862c374016a7f7e24f563)
- [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-ee6206d778b01df7d2a7798364d5bff4165511d30806de46944cb3579691ca12)
- vn_config.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-510d66e0eb26ac744a950f15b1d48f8a63f5957f5068d37401d79f5660a9d232"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-abff902e0abfd84dfd0fc7e2ffac5c9fd6990f0173d99d35bdb0dc56745db7d6"></a>

## Direct properties — vn_config.global_network_list.global_network_connections.sli_to_global_dr / 77ddb15cba8f / 3

- [global_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-ddb1a926156c7dfa942850b991dd9947142765b155b6eb115317d5106cc4b55c): complete subsection reference.

<a id="canonical-8986f6cbbb4e370436ff3700519130d2b1c0947748b160d89bd213bdd54e1ee7"></a>

## Next pages — vn_config.global_network_list.global_network_connections.sli_to_global_dr / 77ddb15cba8f / 4

- [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-ddb1a926156c7dfa942850b991dd9947142765b155b6eb115317d5106cc4b55c)
- [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-ee6206d778b01df7d2a7798364d5bff4165511d30806de46944cb3579691ca12)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-ddb1a926156c7dfa942850b991dd9947142765b155b6eb115317d5106cc4b55c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-425a0ab5361dcc651e9eee8a761a74bca610ac1311859ee09afc0b7904321a94"></a>

## vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn — vn_config.global_network_list.global_network_connections.sli_to_global_dr.global / 7f2f82c2286c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-52586cb0d4e0ac88914ed7ecafc62403594f30a09f9862c374016a7f7e24f563)
- [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-ee6206d778b01df7d2a7798364d5bff4165511d30806de46944cb3579691ca12)
- [vn_config.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-f19f48e60638282949ae91737b2f53d9decad95a3e5153c7f594ed7cbce5b4e8)
- vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-2116b8fa7f9acf945028f6db8bd39335d521829e52ddb15ce328f244d486f69a"></a>

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

<a id="canonical-e14bbd486159ba9488430773f75dae88b05e0b8bb83306a376b0522e690c9fc1"></a>

## Direct properties — vn_config.global_network_list.global_network_connections.sli_to_global_dr.global / 7f2f82c2286c / 3

<a id="canonical-29e5a2cde78b1c2b0799dde611a9f9a393536a808db1e9a04772ae8d91b08a00"></a>

<a id="canonical-6c0f32f66a05b4c982d8873d5732f350897da8929357e61ee551662e397706a4"></a>

## name property — vn_config.global_network_list.global_network_connections.sli_to_global_dr.global / 7f2f82c2286c / 4

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

<a id="canonical-17212703e6995ee7f1bd649a885c858de288bb67d936b5b135cf57fe097743f4"></a>

<a id="canonical-5eab460d05fc7934aff7c07924a46588ec271b7aa32d4f28c5815eeb37689a07"></a>

## namespace property — vn_config.global_network_list.global_network_connections.sli_to_global_dr.global / 7f2f82c2286c / 5

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

<a id="canonical-eaa69a0b13d951c2f19e2daa3cea24c6f6d5ad0b1036652e93fe56cc2f650255"></a>

<a id="canonical-34e1adbd9d94dcc01645def757ba78f2a085a4a3dcf99ee45dc6938971d4b2f1"></a>

## tenant property — vn_config.global_network_list.global_network_connections.sli_to_global_dr.global / 7f2f82c2286c / 6

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

<a id="canonical-ece73998c8629d4847a13039085aa5d9d6e3f112940135da49241e0f174b5a0d"></a>

## Next pages — vn_config.global_network_list.global_network_connections.sli_to_global_dr.global / 7f2f82c2286c / 7

- [vn_config.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-f19f48e60638282949ae91737b2f53d9decad95a3e5153c7f594ed7cbce5b4e8)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-b2714247159432d79fb3ae49cebf317577383cb2ffa4bf87affa300bfd8c451d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99e98fd5d5a7853fcaf1d33b10997163c73055703c86e7608798332ca530a160"></a>

## vn_config.global_network_list.global_network_connections.slo_to_global_dr — vn_config.global_network_list.global_network_connections.slo_to_global_dr / 40c0f3a7e33a / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-52586cb0d4e0ac88914ed7ecafc62403594f30a09f9862c374016a7f7e24f563)
- [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-ee6206d778b01df7d2a7798364d5bff4165511d30806de46944cb3579691ca12)
- vn_config.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-ac51addef5cf435bb85972d4b82d0b5840ec976ccf9a73cf1324c4ced8f2cab9"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-011c2e7688b982de0352dd4d7a37dbff977b77bce260da83c5c208775100cfe9"></a>

## Direct properties — vn_config.global_network_list.global_network_connections.slo_to_global_dr / 40c0f3a7e33a / 3

- [global_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-f6d7f5a0e646fb26648d8f830fb187af7587426e8e08de6d705a502c25ef5529): complete subsection reference.

<a id="canonical-b1c9f3f5c9779289ace0839970befc11c937c42eba16dbf72c458ec1d35c06ee"></a>

## Next pages — vn_config.global_network_list.global_network_connections.slo_to_global_dr / 40c0f3a7e33a / 4

- [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-f6d7f5a0e646fb26648d8f830fb187af7587426e8e08de6d705a502c25ef5529)
- [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-ee6206d778b01df7d2a7798364d5bff4165511d30806de46944cb3579691ca12)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-f6d7f5a0e646fb26648d8f830fb187af7587426e8e08de6d705a502c25ef5529"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26dd8b5b518b0c59d146ccacb8b7f3cf122695b141419dd8974715dc7578e122"></a>

## vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn — vn_config.global_network_list.global_network_connections.slo_to_global_dr.global / 3984ad90ef4a / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-52586cb0d4e0ac88914ed7ecafc62403594f30a09f9862c374016a7f7e24f563)
- [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-ee6206d778b01df7d2a7798364d5bff4165511d30806de46944cb3579691ca12)
- [vn_config.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-b2714247159432d79fb3ae49cebf317577383cb2ffa4bf87affa300bfd8c451d)
- vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-178efda810336a5ac12f165c4f82ada8c554933cb1932ee53c1838458fb2a8bd"></a>

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

<a id="canonical-ae981efdb7a425fc28f94abdc3be83cb003c490cea38e3900547788ce11344f8"></a>

## Direct properties — vn_config.global_network_list.global_network_connections.slo_to_global_dr.global / 3984ad90ef4a / 3

<a id="canonical-1724bc704cae11616e438d898b66c68d30a025d7a5ac652e92407ee9a497a0d5"></a>

<a id="canonical-c9ddc37cd7c045028c5a012661bcb11258967ec5d6a3862282fee281ae7f5b44"></a>

## name property — vn_config.global_network_list.global_network_connections.slo_to_global_dr.global / 3984ad90ef4a / 4

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

<a id="canonical-963bea0b5f37cb209d77e36af97ef3f76fdb380188ca26f0bd308d8928d6335b"></a>

<a id="canonical-024ccb02470b8278f42f8b398a73da43d72ba2d8c5a33b1173b6d96950fe73b9"></a>

## namespace property — vn_config.global_network_list.global_network_connections.slo_to_global_dr.global / 3984ad90ef4a / 5

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

<a id="canonical-ef760c3856b8d7cb1b2ba01d517fd016a1610f2514c6b267463f23b69f34969f"></a>

<a id="canonical-3343916d7a9afa30ae67050adc4a7103cee8563668fc00affccda319ef7aa46e"></a>

## tenant property — vn_config.global_network_list.global_network_connections.slo_to_global_dr.global / 3984ad90ef4a / 6

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

<a id="canonical-8e603a82af96295472b74648b6d7c4f82f78db4ae1e3023e1708d48d2af96557"></a>

## Next pages — vn_config.global_network_list.global_network_connections.slo_to_global_dr.global / 3984ad90ef4a / 7

- [vn_config.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-b2714247159432d79fb3ae49cebf317577383cb2ffa4bf87affa300bfd8c451d)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5d30ac93dbe823120a40e54e0089c1a4e7f0523532b0a5dfe92949725315394"></a>

## vn_config.inside_static_routes — vn_config.inside_static_routes / 212960c6c228 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- vn_config.inside_static_routes

<a id="canonical-77395b6a6d1b2077e2a00a4c4fe80f620c5109175a808bb0feefcbd75ab195ea"></a>

Type: `"single"`. Computed.

Configuration parameter for inside static routes.

Upstream description:

List of static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-85ba6efdfde674f70b72db0cbb8ca33c6a4c92551dfacaf49fdca31c854742d7"></a>

## Direct properties — vn_config.inside_static_routes / 212960c6c228 / 3

- [static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7): complete subsection reference.

<a id="canonical-d05a316bf634c0425663dba44fee7ce10be695063016124b93c15f4f58a1a32f"></a>

## Next pages — vn_config.inside_static_routes / 212960c6c228 / 4

- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e95554dfeb54a1a7750a5240ef4598cbe4a046bf5392bbbb8b4f174892440bb"></a>

## vn_config.inside_static_routes.static_route_list — vn_config.inside_static_routes.static_route_list / 7e79d1dbfaf6 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- vn_config.inside_static_routes.static_route_list

<a id="canonical-38085c9a9a136e3727de8ea4db5a4bdecec927d3c207bba0232ddfffb2912686"></a>

Type: `"list"`. Computed.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-5996ba4ea9d4206418e07255c7f9d2f244c83016f8670b3001e79b8d67421241"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list / 7e79d1dbfaf6 / 3

- [custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb): complete subsection reference.

<a id="canonical-0028e82b9552f84599d5eb24689043c6224ab65cc2104dfcc96ff0918860ad36"></a>

<a id="canonical-d187a23a60a62a4fba27c50b86543f2d00986a848288ba576beace934e38ae85"></a>

## simple_static_route property — vn_config.inside_static_routes.static_route_list / 7e79d1dbfaf6 / 4

Type: `"string"`. Computed.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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

<a id="canonical-c7a8164364706606acb9bc3f750228cba5f156911e2e026019cb914c77c1d4ea"></a>

## Next pages — vn_config.inside_static_routes.static_route_list / 7e79d1dbfaf6 / 5

- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-465f2cf2ae6b9faa2d72b92950dcf66eb57aa12d8a9a5add1f7f1411616991a5"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route — vn_config.inside_static_routes.static_route_list.custom_static_route / 63101be701cd / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- vn_config.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-b4255be746d21e9ed6b08e0b35a042e1df40bca48660545fe7c12f98caa72555"></a>

Type: `"single"`. Computed.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-781081afc45b98155410d54a27c6dfebacd31118a8c006a3bd1df82ed15cbaf8"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route / 63101be701cd / 3

<a id="canonical-7e00f4cb2cf22dcb0ab3089a6fe8a8bcb096821f34af841c6d1f1351569e9b89"></a>

<a id="canonical-f229ba5f282fe149cb0d69183af2571b4540c1e813d59637fd33b4db9e1f3962"></a>

## attrs property — vn_config.inside_static_routes.static_route_list.custom_static_route / 63101be701cd / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](data-sources--aws_tgw_site--reference--group-003.md#canonical-e9347ff478f6c22289a7462a067208a5d362ed84dd6d4f35eb46ff4d7438b030): complete subsection reference.

- [nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-2288304af6369af4cf4a911cf4850704311ed601645b4336fbc80021633c1734): complete subsection reference.

- [subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-e7855e6bb4e1c05bcdbd841ff2120788338c00cd223f973950afb581c7dc02bd): complete subsection reference.

<a id="canonical-5cba4285d0a1341857cd890de924f2d32e043b859699740013f2de5be3c47523"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route / 63101be701cd / 5

- [vn_config.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_tgw_site--reference--group-003.md#canonical-e9347ff478f6c22289a7462a067208a5d362ed84dd6d4f35eb46ff4d7438b030)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-2288304af6369af4cf4a911cf4850704311ed601645b4336fbc80021633c1734)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-e7855e6bb4e1c05bcdbd841ff2120788338c00cd223f973950afb581c7dc02bd)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-e9347ff478f6c22289a7462a067208a5d362ed84dd6d4f35eb46ff4d7438b030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5d0e52bc7ae21143b6df6c9d93cfe6053dbbdc3266600a6c274e03fa96b0f01"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.labels — vn_config.inside_static_routes.static_route_list.custom_static_route.labels / 596dbce066cd / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- vn_config.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-9069cdda869b1fe17184519b84a4132f1420d6e6510a6b8c47b5b3ce0afb5e67"></a>

Type: `"single"`. Computed.

Add Labels for this Static Route, these labels can be used in network policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-180208e16ff32b08b140b5757850f8950edb44b33b6b11f8ee3af367ecdbeca6"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.labels / 596dbce066cd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa1ee7fbe3f68c4a0698063a37605866078db54c1a5be2e086e14a9c82ba9edb"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.labels / 596dbce066cd / 4

- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-2288304af6369af4cf4a911cf4850704311ed601645b4336fbc80021633c1734"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e299c2029dbff625183582ccb9429624b318356ba9d0abbb8ed56d0b667bdf3"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop / 4f5ee86f1fe0 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-d844ff7876b3a4e16e74b13b4f551e8d81572abc8ad5af2ab0618eee7e0f9ff9"></a>

Type: `"single"`. Computed.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ba8a679f35539309800bc426b46dcd04fd0430bdf9a5de68465642854744a099"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop / 4f5ee86f1fe0 / 3

- [interface](data-sources--aws_tgw_site--reference--group-003.md#canonical-fb130b7fcbb9071cdd4102daf78e407d6b1667f5f6f50c82587d6997593b5ea6): complete subsection reference.

- [nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-95cdbc1367fa73f7b7e2f37d317dc8808921f07fb2531b553da583cb70cf7073): complete subsection reference.

<a id="canonical-34ca86cd22203aec34245d4b7cc1259937b97a3decab967410b4b09e8f10f8f4"></a>

<a id="canonical-fea559ef4d6ef01428c19ef7625630406cc5b5b05f2ef9b2ba2602ab52bc8b7e"></a>

## type property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop / 4f5ee86f1fe0 / 4

Type: `"string"`. Computed.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ccbbfd9c74566f7a672c4e6ed7c9bbc8ce6ad2e7fee22017aa05cb94b82a60a7"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop / 4f5ee86f1fe0 / 5

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_tgw_site--reference--group-003.md#canonical-fb130b7fcbb9071cdd4102daf78e407d6b1667f5f6f50c82587d6997593b5ea6)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-95cdbc1367fa73f7b7e2f37d317dc8808921f07fb2531b553da583cb70cf7073)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-fb130b7fcbb9071cdd4102daf78e407d6b1667f5f6f50c82587d6997593b5ea6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38e68713e31127feb97a0f108d5ede73dc7c254b025d57a4d2dc316c22e721a2"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 87e9521807a4 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-2288304af6369af4cf4a911cf4850704311ed601645b4336fbc80021633c1734)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-bcaab18cbe117d2aeb6ae9fed6826f280a6cc1791c193b3ca965e08f3282543c"></a>

Type: `"list"`. Computed.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

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

<a id="canonical-9ee9b837c8bb10f8062f389154cc7649f5621b8fbf63afe59ceb28a0b656ea0c"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 87e9521807a4 / 3

<a id="canonical-6a4b44118e914655bdf6bfab8f74e5ad94c2d7858d3939eaec13a20337800f55"></a>

<a id="canonical-deac4fe6e4a4939189b2672c8ef9a9c43f79fdfaf7403d829a470975420a02d2"></a>

## kind property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 87e9521807a4 / 4

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

<a id="canonical-c65ea6904265eea7906d8664bd6593111f6732c56241ab0e5459e4295865ff97"></a>

<a id="canonical-f34cca6a4849888fb7a2ff1f2b3f311463321035d8e39bb4f41880601fe5efd1"></a>

## name property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 87e9521807a4 / 5

Type: `"string"`. Computed.

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

<a id="canonical-d1c533fa3087382233a6443a246466df15bdcb3ec44a3d9c964827139238b2a5"></a>

<a id="canonical-2481f7544b3684f2763a45c943ddca02d1f715d3f0f1932da42a1c21b3974096"></a>

## namespace property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 87e9521807a4 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-64cd0fec1eb02bd38deaee2a55b57f867a678ca2bdbf16c54d5f418e84c2cb94"></a>

<a id="canonical-1ae8770cf109adc41eaee98ff656e900857350abca93a170d298f770ca115379"></a>

## tenant property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 87e9521807a4 / 7

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

<a id="canonical-74f5a08924f780c8188c773a93af11ae384c9838156f9c56c5a03923d538fd7a"></a>

<a id="canonical-64734d3bb17d5e756e062ab242489e31f0db02e5a33d3a34439b861a8b5fdb73"></a>

## uid property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 87e9521807a4 / 8

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

<a id="canonical-c755d98309952491f5ea3b5c60a1b98abadbe3aa9708f2258af5e29de76fe101"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.int / 87e9521807a4 / 9

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-2288304af6369af4cf4a911cf4850704311ed601645b4336fbc80021633c1734)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-95cdbc1367fa73f7b7e2f37d317dc8808921f07fb2531b553da583cb70cf7073"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e85e6ef208a7bd4330c86478ee63d017a886b156ee9ad9dc17e1ba0fc4397caa"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / e30a789984cf / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-2288304af6369af4cf4a911cf4850704311ed601645b4336fbc80021633c1734)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-47ab70ce259ee88e9f92ad6411ab8af9204dcd63798e954b520a9370b3c99815"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

<a id="canonical-2047bb293335a33523fda1675dc7c29f200a58e20cee8b083f50dfeebd87e305"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / e30a789984cf / 3

- [dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-9b3fd15ef3c27f8c25d9f9c7fe60cf2ee65023c6a1ecee6608ee956cc8f99628): complete subsection reference.

- [ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-9df13ec6097e1ccbd9c44487d77c553c79c67ec3b63d59d9bd2ce7415771a7a9): complete subsection reference.

- [ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-00016b6afc3f42fdabfe403a37a0c4db7fa2956d03336b5e0c17eb041e6b5fd5): complete subsection reference.

<a id="canonical-d4de493ccd30f53e7707e10eabf94fa9ee391291e9fa552f76090807e62f63db"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / e30a789984cf / 4

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-9b3fd15ef3c27f8c25d9f9c7fe60cf2ee65023c6a1ecee6608ee956cc8f99628)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-9df13ec6097e1ccbd9c44487d77c553c79c67ec3b63d59d9bd2ce7415771a7a9)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-00016b6afc3f42fdabfe403a37a0c4db7fa2956d03336b5e0c17eb041e6b5fd5)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-2288304af6369af4cf4a911cf4850704311ed601645b4336fbc80021633c1734)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-9b3fd15ef3c27f8c25d9f9c7fe60cf2ee65023c6a1ecee6608ee956cc8f99628"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31d5ab8d5e76f58269dcc8b74c555844a01f4404b561b4a1742c287677440d25"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / bc5cd49aa397 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-2288304af6369af4cf4a911cf4850704311ed601645b4336fbc80021633c1734)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-95cdbc1367fa73f7b7e2f37d317dc8808921f07fb2531b553da583cb70cf7073)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-e6510a9e39310b7d8e64871fd3bfbd899048c86a8022bcce9ee94a6d47cd5fe9"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-88b308956bee5e829eea88f037d709dea71821aa589f0eb267d2179b726c8fb3"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / bc5cd49aa397 / 3

- [ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-d1cf1fb5bbece74259cb00984119921c0143b3919836e29c0fb57924a24f10bf): complete subsection reference.

- [ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-96fd640ee8e7ee8e7fb1db9eb83f35b9e038389fecc2b263c1c23c2fb6ff4ce4): complete subsection reference.

<a id="canonical-c90cd1470b0ad9a8d875bb2a8adbe8b1272887f8adf09cb987f170b6f597a4aa"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / bc5cd49aa397 / 4

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-d1cf1fb5bbece74259cb00984119921c0143b3919836e29c0fb57924a24f10bf)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-96fd640ee8e7ee8e7fb1db9eb83f35b9e038389fecc2b263c1c23c2fb6ff4ce4)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-95cdbc1367fa73f7b7e2f37d317dc8808921f07fb2531b553da583cb70cf7073)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-d1cf1fb5bbece74259cb00984119921c0143b3919836e29c0fb57924a24f10bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51c6618fecd3b770fc0096f290918f393e67ef982a15fc1c29e28798dba95eb9"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / ce093cd78508 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-2288304af6369af4cf4a911cf4850704311ed601645b4336fbc80021633c1734)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-95cdbc1367fa73f7b7e2f37d317dc8808921f07fb2531b553da583cb70cf7073)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-9b3fd15ef3c27f8c25d9f9c7fe60cf2ee65023c6a1ecee6608ee956cc8f99628)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-e1ed1d2721c33e535342aa72d9127c462565b0436b6534d6b7f6b154ca126ae3"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e1a9ad4ef25e177dbc1347be9dc41f315bf71818491ccce128c9f7946cadf1b5"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / ce093cd78508 / 3

<a id="canonical-58948e3494d73ed114b5bd32289ed06f42bfa4ea73437b49321da44c01546ab0"></a>

<a id="canonical-62184c76a4668e3830a6ec58189f10fda120c84f0e258a4e82ab7315087a5ad2"></a>

## addr property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / ce093cd78508 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-1a02374cd19a3719149b681da24e7d092a80638ae601ce07166f0e9001fdb1d4"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / ce093cd78508 / 5

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-9b3fd15ef3c27f8c25d9f9c7fe60cf2ee65023c6a1ecee6608ee956cc8f99628)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-96fd640ee8e7ee8e7fb1db9eb83f35b9e038389fecc2b263c1c23c2fb6ff4ce4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57bc54440b28a9e211cfffc97a8fb577487d9ce53028209da2d87336dbd40afb"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 0868cfe40ff0 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-2288304af6369af4cf4a911cf4850704311ed601645b4336fbc80021633c1734)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-95cdbc1367fa73f7b7e2f37d317dc8808921f07fb2531b553da583cb70cf7073)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-9b3fd15ef3c27f8c25d9f9c7fe60cf2ee65023c6a1ecee6608ee956cc8f99628)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-aa45d2dd132d13bc261310aebec11b9d2256461a6a4c86951490ba93adab0545"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-78c42a93da0a95d58ee9866b44314f0b4e90eb9d4a98ae4f10fcc1a2778725e6"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 0868cfe40ff0 / 3

<a id="canonical-85a74f1b64336852f64711f242b90875854dd0cbb237327134090769d866454e"></a>

<a id="canonical-79f16599811e6a5723ee80dff5b2a31aa5a35842aa89cf615d12d3589a59deba"></a>

## addr property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 0868cfe40ff0 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-ca355d4b850f601b56671304c70ae814aa4abf53e05141e01f8b021912ab100a"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 0868cfe40ff0 / 5

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-9b3fd15ef3c27f8c25d9f9c7fe60cf2ee65023c6a1ecee6608ee956cc8f99628)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-9df13ec6097e1ccbd9c44487d77c553c79c67ec3b63d59d9bd2ce7415771a7a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eeca3d172c6ac772cefbda7cfaa13fb4f83b62b2094877b3fd94b2b5b751c77c"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / f7a4237079b9 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-2288304af6369af4cf4a911cf4850704311ed601645b4336fbc80021633c1734)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-95cdbc1367fa73f7b7e2f37d317dc8808921f07fb2531b553da583cb70cf7073)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-dcfe85c3a29229032a6e3af58f765803738dda3367a0282ed947b50fcd0981bf"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9be51d3116336f9bbd37c32ad8457f850728b62503ee4a1e90f14c2f6c6e0084"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / f7a4237079b9 / 3

<a id="canonical-89632ca4de1d22027116222352b7d46014bf05c0db7b0c34518a944bc5353625"></a>

<a id="canonical-aa58007fa187df9057e3f4cb566af322d55b649687a2053849bf3d5707439cd0"></a>

## addr property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / f7a4237079b9 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-8108d13be27ebf898548b9dd558c70c28f8020528a016c0bfb25ed9620993b5d"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / f7a4237079b9 / 5

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-95cdbc1367fa73f7b7e2f37d317dc8808921f07fb2531b553da583cb70cf7073)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-00016b6afc3f42fdabfe403a37a0c4db7fa2956d03336b5e0c17eb041e6b5fd5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65fb02fb70ece7154e0a46ce6def9977a5f5b0654a5dbfae53e34169e4928a31"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 07cfcb532e62 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-2288304af6369af4cf4a911cf4850704311ed601645b4336fbc80021633c1734)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-95cdbc1367fa73f7b7e2f37d317dc8808921f07fb2531b553da583cb70cf7073)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-098f6bfd4ff7eaf789d3590e54b8852d0e214a89609b804ab123bb1fbced3da5"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3770d7405cbf857d15db3c187cc434c4f5a3888c900f13fdcde8d922b5651230"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 07cfcb532e62 / 3

<a id="canonical-deff29e8bceba55cde7ff455bce3c7098143da6d5693222ffda77b0ed322d1f0"></a>

<a id="canonical-9227194583215fa8bc74a80d41cb40abd85520486733d01f82a2e2a582994a1f"></a>

## addr property — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 07cfcb532e62 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-05eb6b73df9e273add6f136d0fd9ded743240785d0c4bdb25af25bb666bb5dbe"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nex / 07cfcb532e62 / 5

- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-95cdbc1367fa73f7b7e2f37d317dc8808921f07fb2531b553da583cb70cf7073)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-e7855e6bb4e1c05bcdbd841ff2120788338c00cd223f973950afb581c7dc02bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a48cc0520e1794ec8c8edd528fb7c2dfde14f6aa4e385770fcf74a00aed49480"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.subnets — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets / c14a4db62047 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- vn_config.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-e01dea2cc991d5254ae6e0cca8a5bbb84e5ecd6fe7a11eb59154755d3a88a3ea"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-fa46541ece251b73d89d956ada4827898e6b86715e1f3713d9591fd9ebee11c3"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets / c14a4db62047 / 3

- [ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-aa337f355b770caf3cd232151cf66e55a756ae3e44ec102925aa884264876cc8): complete subsection reference.

- [ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-8ba8e11c23c0cc73187ab9d2c8f0e0021185e31bbf39f2338540d4e5d07edb4e): complete subsection reference.

<a id="canonical-f7bf925f8d4f7176854e0ae2d5a42278179b08e01f84e9e9ed1b32a3f9ff875f"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets / c14a4db62047 / 4

- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-aa337f355b770caf3cd232151cf66e55a756ae3e44ec102925aa884264876cc8)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-8ba8e11c23c0cc73187ab9d2c8f0e0021185e31bbf39f2338540d4e5d07edb4e)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-aa337f355b770caf3cd232151cf66e55a756ae3e44ec102925aa884264876cc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3432bafd5d86b0810c8ed172b729c52373292b41e81af58f5b92168c5372d066"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 11b192a8e004 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-e7855e6bb4e1c05bcdbd841ff2120788338c00cd223f973950afb581c7dc02bd)
- vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-72a3f012479b3f64f6eaa51459dd0e569882e2870688b4d3b2fb254105999e54"></a>

Type: `"single"`. Computed.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6e978208fc9002b6e65e2fba47f06ee8a617e679a0222667bdb307d65568d592"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 11b192a8e004 / 3

<a id="canonical-4ba09d1fefd192e35bdde26e127b4f917ea00b7e69169593919181bd870c2971"></a>

<a id="canonical-0a36fe2dbfa38dcc9ef181a01a9d780df790fd3da1f711ff2395cf9785bd49a2"></a>

## plen property — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 11b192a8e004 / 4

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

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
    }
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

<a id="canonical-6cd5ea2b3284f4da4ac2db8229f9dca2e531c9429903aee90f20a3b784654b3a"></a>

<a id="canonical-ebaa934eeb8e725d0e6ac8d2d3df9a61d228950c7dd2b9a1b7c64f17e15cdba4"></a>

## prefix property — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 11b192a8e004 / 5

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-9bde19d15a5b9a061023a3b0d5423abb33e13fa4de839f2f9f71bfd110c9f89d"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 11b192a8e004 / 6

- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-e7855e6bb4e1c05bcdbd841ff2120788338c00cd223f973950afb581c7dc02bd)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-8ba8e11c23c0cc73187ab9d2c8f0e0021185e31bbf39f2338540d4e5d07edb4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ff0c7a1885204d9217133931911a20443da9669fca5956ff61ca71142259f8b"></a>

## vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 80784ee5b685 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-e1ec3ae5699fd3cc954fbcb4005b852065769fa90244dbdaeea40575c79e22b8)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-3a1f7b0c665044554f282b929d96be883a7972fa2b3896cb2d025c98e60674a7)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-cc31bda9e2712bc707f5f9ac4521633fd5009f5d93059a53abfbc5da200217fb)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-e7855e6bb4e1c05bcdbd841ff2120788338c00cd223f973950afb581c7dc02bd)
- vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-784c97cca1b5cd990e897e7c22f3a881e4d2737286ab6711b9478f2537a1010c"></a>

Type: `"single"`. Computed.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2df4cc9670cb3fb8e8cb04d7cf9166d343a4e296fd5fcb817e5d6ad357578b8d"></a>

## Direct properties — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 80784ee5b685 / 3

<a id="canonical-9f9c58234560fd4eda94dc0e660bdc0bd149fbfa3dc809b526b12692a8283c1d"></a>

<a id="canonical-7ae08c55dd7311b1362aaccd4a1d508d58a1b6c1e95493c65a1c660cf190e980"></a>

## plen property — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 80784ee5b685 / 4

Type: `"number"`. Computed.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-1b08f9c4ee3119852207fc5ab0c22e6c34a8c672b79ec71fd1ab6e0ed3028660"></a>

<a id="canonical-03475fc2e6371a2dbc8b488f5fb8b4c6e16b7e2339b0c3619e48d6337957a6f9"></a>

## prefix property — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 80784ee5b685 / 5

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

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

<a id="canonical-ac2c53b64a16e344ca4f0fd4ff63d38000fa268e5b781d509b800496eaadc82a"></a>

## Next pages — vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv / 80784ee5b685 / 6

- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-e7855e6bb4e1c05bcdbd841ff2120788338c00cd223f973950afb581c7dc02bd)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-fa3fb56fbfc055e5b3d06ffc4eefc3ef5fdd5fdc73cf09e46aacff2f3a27129f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8cf808a62d4a16bc5531e8d9fb6aaccc98276c46d42ee3c5bf50c586c8bc0e68"></a>

## vn_config.no_dc_cluster_group — vn_config.no_dc_cluster_group / 3c6aa7e68682 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- vn_config.no_dc_cluster_group

<a id="canonical-6978794e9f73e35b082bd014485a0732a9f0c6f4a003e7a1fb7e7e20823fa2d1"></a>

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

<a id="canonical-966adeb8d3940acf4d14c8a1c6ea46287d51769e8f0043ae9275d01c5e86bf94"></a>

## Direct properties — vn_config.no_dc_cluster_group / 3c6aa7e68682 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e7e9da64e8c856226f8c241369ea413252bb8aa9a77e3c31d027c275ba5dbe16"></a>

## Next pages — vn_config.no_dc_cluster_group / 3c6aa7e68682 / 4

- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-25b7302d7eebfcd4adf3f689ca1d7d323dae92f439034c5b5670ee1e0645763f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9f808f2af8f1e0a28e0b214391dd5bbde23404a6465977852e47105d90c33ec"></a>

## vn_config.no_global_network — vn_config.no_global_network / d190dfbec65c / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- vn_config.no_global_network

<a id="canonical-2e70c28e74d3d84bfc1a13aba5a46cce722710230d861c949626e86c16724d7f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no global network.

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

<a id="canonical-b0577667637ac0193408088654a5f554094858be8ef640f5646240931e4fee0f"></a>

## Direct properties — vn_config.no_global_network / d190dfbec65c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-66d924f85eb3c9182b79beeff7fc387b5038bc9a520753868806525a3e7ca729"></a>

## Next pages — vn_config.no_global_network / d190dfbec65c / 4

- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-9f4c0cd4ba79349015b403ab4688e69d9e338e96b5bc82fb4443ee6b2624ae9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb944397e1c70c8d8e67e595b06bd2501ec2252611baeec6978be650e10bd905"></a>

## vn_config.no_inside_static_routes — vn_config.no_inside_static_routes / a0aad7eca27b / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- vn_config.no_inside_static_routes

<a id="canonical-e59e2f2fc6dd976c8795649b8d25aa62451ca5093ada723e77bdbd8738525058"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no inside static routes.

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

<a id="canonical-1853aa12d2111e2aece305f56217b63d34dfd340e72045d628bc37e5134b3b3d"></a>

## Direct properties — vn_config.no_inside_static_routes / a0aad7eca27b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6c3a431ddfd7d98a053907d266b10ac0346212cf3861912d406ef577b1d6f7b0"></a>

## Next pages — vn_config.no_inside_static_routes / a0aad7eca27b / 4

- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-a8654d3af4e3abb9a00673afe8fbdd1d517fb0ca691537558ebc49634b43537f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-170877bb741a1bbb40edbcaec333ce7567a20cb3ce7e25ce82a3ed423e552a83"></a>

## vn_config.no_outside_static_routes — vn_config.no_outside_static_routes / 3ca240d94001 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- vn_config.no_outside_static_routes

<a id="canonical-13725b68d17fcc374360eff82993dc4a0f178f1f2129ff4bc62c27557f1f81da"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no outside static routes.

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

<a id="canonical-276502638289a762f201c42bdf53d54df4130ce2390cc36ca0fb0c107c2ea314"></a>

## Direct properties — vn_config.no_outside_static_routes / 3ca240d94001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b0a07522b1f325bd2767848c06ee34cbd903a2b03d0a5699071a53ec1ff0c852"></a>

## Next pages — vn_config.no_outside_static_routes / 3ca240d94001 / 4

- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3848398f9a16073e5678a5b9ca3aa97a70da1f29cff8ae5abb1da147fb669f10"></a>

## vn_config.outside_static_routes — vn_config.outside_static_routes / d548adceb0f8 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- vn_config.outside_static_routes

<a id="canonical-f6f7be2e5b45e94a34a64cfd268a67b4b45d00c3ab5f6e632f95cf0b2241b25d"></a>

Type: `"single"`. Computed.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3f8fc71f5b956c93cde8a47dabdb4219b1416c13a4f49e85d917757f3aecde0d"></a>

## Direct properties — vn_config.outside_static_routes / d548adceb0f8 / 3

- [static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca): complete subsection reference.

<a id="canonical-6b97f44a77274b8e0fd7ab5b984237230ee8491460f8df3ef13e89de4a430ecd"></a>

## Next pages — vn_config.outside_static_routes / d548adceb0f8 / 4

- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b1ea20fd3ad942d1c3da62d5e79b8abec505682975d2bb94bfb0bbc3717d94e"></a>

## vn_config.outside_static_routes.static_route_list — vn_config.outside_static_routes.static_route_list / ff1504966afc / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- vn_config.outside_static_routes.static_route_list

<a id="canonical-17bf8c248d588df892025a7d223fa406821a7fada821333d5c0a40d7dea51e77"></a>

Type: `"list"`. Computed.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-51a980a9ab2572312edba4a17462576d4b263e8c38185a09ee157ba76da006a8"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list / ff1504966afc / 3

- [custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf): complete subsection reference.

<a id="canonical-32d3678a730ba3bf075b4615917a6a5a4b6e89274bb0709e71be35cd38becf58"></a>

<a id="canonical-80f877d5238dcf98ca7c5ac1bd8e8799e04647bdef2106ba8aee681758dba712"></a>

## simple_static_route property — vn_config.outside_static_routes.static_route_list / ff1504966afc / 4

Type: `"string"`. Computed.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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

<a id="canonical-77b1a8017b22bfe5e44c06d0d32e5c0711464efa69e99430e928e1c8611bbdd7"></a>

## Next pages — vn_config.outside_static_routes.static_route_list / ff1504966afc / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01d37d508a7863bd5803398896d1dca67211d5128fae72b7d4442c0f4698ca14"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route — vn_config.outside_static_routes.static_route_list.custom_static_route / 881c315811f1 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- vn_config.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-a799550dc83d4c1feadab4463d9ea02630adc6f321e81987f44e189f1d4d6c11"></a>

Type: `"single"`. Computed.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3f16a2e0a753071170fe89281b252eaac54bd4330b74ff8faf4c6c197b5d0236"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route / 881c315811f1 / 3

<a id="canonical-01cf3966c1567ce2ea27048453b53a656005ec6f97d7fa19d789fad261a392af"></a>

<a id="canonical-e625638d0fef1743339051e781d610ffd61502e300a4eecc382a405a37d8ea8b"></a>

## attrs property — vn_config.outside_static_routes.static_route_list.custom_static_route / 881c315811f1 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](data-sources--aws_tgw_site--reference--group-003.md#canonical-5056fe0c0d800449aee0881f01aff8efc60fd2d04af69a530ced444543176c5c): complete subsection reference.

- [nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-a7f871f23f326555861537ba44ac293fcb58ff8b5e0899cbd3af109faba07ed7): complete subsection reference.

- [subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-9cad8c1bb58d76db38273421360c7368fcfc653a2d0ae9052c3a387bf52cc0a1): complete subsection reference.

<a id="canonical-a4dd2d66c2b3a49a3fe078e882d7e273a65f0ea9f57f6fa8c0365392c9ca1f8b"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route / 881c315811f1 / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.labels](data-sources--aws_tgw_site--reference--group-003.md#canonical-5056fe0c0d800449aee0881f01aff8efc60fd2d04af69a530ced444543176c5c)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-a7f871f23f326555861537ba44ac293fcb58ff8b5e0899cbd3af109faba07ed7)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-9cad8c1bb58d76db38273421360c7368fcfc653a2d0ae9052c3a387bf52cc0a1)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-5056fe0c0d800449aee0881f01aff8efc60fd2d04af69a530ced444543176c5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-151c02e564c4ed3c4c67d29dfb9080f972652b69e7a844b44c9888f91d199501"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.labels — vn_config.outside_static_routes.static_route_list.custom_static_route.labels / bcfcec5a28ad / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- vn_config.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-608eb731ef84811cdcf6d89f7760720015133a9702dc278e18fb450f0d9abcae"></a>

Type: `"single"`. Computed.

Add Labels for this Static Route, these labels can be used in network policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-58b19d5ee2118e3464e2e8192b012cbb8a60d20da429f244d35ce9a6f89cf0e8"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.labels / bcfcec5a28ad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7ef55c77af8d4c9cdcf1a650dfa040a6c288b9ffe12a9b53b414d203bdceceb3"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.labels / bcfcec5a28ad / 4

- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-a7f871f23f326555861537ba44ac293fcb58ff8b5e0899cbd3af109faba07ed7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a13d1ec29fdc44eb85cf08e30a8d13dcbe4f4dbd2ff003b4ac33c8f9cc2e470"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop / da9ce62befa8 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-33cd500dd4a4cf86e5caca5d8ba721f8fca5f772df6a127ba99cf7df5ffbb8fe"></a>

Type: `"single"`. Computed.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4a3851f8265acf0b52cb315be5e513c562334aa8edb114d8243c08e0f2fdd937"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop / da9ce62befa8 / 3

- [interface](data-sources--aws_tgw_site--reference--group-003.md#canonical-3724217a856ed03b70f86da71662751dcb7c219c3e4ad62a9d7ad82f5eaced63): complete subsection reference.

- [nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-4d1526acce9e6cae571471ae42c117ab85b1bb44fe95073a1caf1e3de0cc8d50): complete subsection reference.

<a id="canonical-3e638279180a6560392271bff24b20a2c3a46ee4a1348ade28257d8a2640cde9"></a>

<a id="canonical-4031d0e679f07cd12ed957282ac65b71dcbe416840dee858c908e93199ef051a"></a>

## type property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop / da9ce62befa8 / 4

Type: `"string"`. Computed.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-df78824435df5cfd689fad9b01c424ff1b486b516a03e51374d020ed114b02da"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop / da9ce62befa8 / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_tgw_site--reference--group-003.md#canonical-3724217a856ed03b70f86da71662751dcb7c219c3e4ad62a9d7ad82f5eaced63)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-4d1526acce9e6cae571471ae42c117ab85b1bb44fe95073a1caf1e3de0cc8d50)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-3724217a856ed03b70f86da71662751dcb7c219c3e4ad62a9d7ad82f5eaced63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cbc89aee2b8d11f6580ca37699c2c40b8f3834cb747ad327a8ba2e83a2e9d429"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / 5f647f11ecf3 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-a7f871f23f326555861537ba44ac293fcb58ff8b5e0899cbd3af109faba07ed7)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-62f231cd56dfd0e515cdeaf0854cf85079b39495c125ee2534e43d5f591a6b33"></a>

Type: `"list"`. Computed.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

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

<a id="canonical-b95129e30ff6e23aa5a833509c607039157ea327de71fc80bdf7ba0254c31576"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / 5f647f11ecf3 / 3

<a id="canonical-5ac00e2e95eebf245814692708b32b74c433d187daf0865e87d80adb712ae01d"></a>

<a id="canonical-e4bd83b9aa7bad3f5581a03e6282042a5ece2147692cd061bc047ffd574fbc1d"></a>

## kind property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / 5f647f11ecf3 / 4

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

<a id="canonical-431065f1a56cd191ed8a6c94c966c6a59b24215ce7cb387abc36cdd081629475"></a>

<a id="canonical-cc4f4153fe4f57c33968ef3d7d76134b33c006e77182b8c3022377969936ad32"></a>

## name property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / 5f647f11ecf3 / 5

Type: `"string"`. Computed.

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

<a id="canonical-6357f8a6ae93b36378dfc39bfb130abb1bbf951e3534a8411d657c413e68ed28"></a>

<a id="canonical-f63b388cf405d476800c284948fcb11a87f399704ac1e8b02df85089fc6ed12c"></a>

## namespace property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / 5f647f11ecf3 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-75acb681d3d734b1f29a1736201faf04b277988fb51445195b73cd94e138a114"></a>

<a id="canonical-b0166bbb3f6e8345a5f18dd8c3c98167a76acd8f6c00d4d0fcfe6493cd8616b2"></a>

## tenant property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / 5f647f11ecf3 / 7

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

<a id="canonical-114486fd1c13e3f46824f44ffe9b42220d23b83ee03fd8cb677133319dc10ca9"></a>

<a id="canonical-1fa48bfed69e7115e574b0b374513740d648990186d804f05c15a4eaef24f410"></a>

## uid property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / 5f647f11ecf3 / 8

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

<a id="canonical-d794227f5f800d1671fe43cc8324f326792512ab37d399154e07297501407230"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.in / 5f647f11ecf3 / 9

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-a7f871f23f326555861537ba44ac293fcb58ff8b5e0899cbd3af109faba07ed7)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-4d1526acce9e6cae571471ae42c117ab85b1bb44fe95073a1caf1e3de0cc8d50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a700930de17a1e44f903a901bda2ea42711761a3a01a785cd8213f3e858786b"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 36febbae0987 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-a7f871f23f326555861537ba44ac293fcb58ff8b5e0899cbd3af109faba07ed7)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-afbee053028e331a9b9a34afe97677ca96c2d4b0c6553054ae6ecf482435c839"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

<a id="canonical-1065eaf60789a9b135edef0681b7796e739f44a7ee18769ee24c6fcbc9a5fb08"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 36febbae0987 / 3

- [dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-0d9e5be6600ba97126ae137b748befd063a85c0f348c1947a0ac3001823d9f0a): complete subsection reference.

- [ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-84f8daa4eb4b4b48a06ca1b658217437c459b4aeaed450735bda3328e2ca92b7): complete subsection reference.

- [ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-518aa9a897065155b92894bbd2c3e6e7314a083f6b8e4aa40c6694b626f1190c): complete subsection reference.

<a id="canonical-a8dca1e4c29dacffc32a9599b9318aa04670c4bf8664b34d621dc80c8eb4c48f"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 36febbae0987 / 4

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-0d9e5be6600ba97126ae137b748befd063a85c0f348c1947a0ac3001823d9f0a)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-84f8daa4eb4b4b48a06ca1b658217437c459b4aeaed450735bda3328e2ca92b7)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-518aa9a897065155b92894bbd2c3e6e7314a083f6b8e4aa40c6694b626f1190c)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-a7f871f23f326555861537ba44ac293fcb58ff8b5e0899cbd3af109faba07ed7)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-0d9e5be6600ba97126ae137b748befd063a85c0f348c1947a0ac3001823d9f0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1114b7522784f6515054616cf4e45bda82a88e8d12d7f02d6bdee01085aa9953"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / d83e93700325 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-a7f871f23f326555861537ba44ac293fcb58ff8b5e0899cbd3af109faba07ed7)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-4d1526acce9e6cae571471ae42c117ab85b1bb44fe95073a1caf1e3de0cc8d50)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-230b9464ad6f42ae941927e1fb7070d3c510e939892fcbd18e056c980006225c"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-cd7e045e05acae2822277b43a1e3456b6c0b171e80581fc4356ee367259d41fd"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / d83e93700325 / 3

- [ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-b5ff0dd924e522bab9b735b601bd45b4ed7a7ff13176d4372f7f82bf48f88818): complete subsection reference.

- [ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-39edb6c625008498993006e06d6aa5e87d14b9af34d9a66be14856b6b9dbc211): complete subsection reference.

<a id="canonical-6959fbee9ae71f3a3ce8521ef4e44b60aa7b0a04c2bc82244d121a6ffd84a0cd"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / d83e93700325 / 4

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-b5ff0dd924e522bab9b735b601bd45b4ed7a7ff13176d4372f7f82bf48f88818)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-39edb6c625008498993006e06d6aa5e87d14b9af34d9a66be14856b6b9dbc211)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-4d1526acce9e6cae571471ae42c117ab85b1bb44fe95073a1caf1e3de0cc8d50)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-b5ff0dd924e522bab9b735b601bd45b4ed7a7ff13176d4372f7f82bf48f88818"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7071b686b6fb7effea6c2940d4b8fab04808febc8c3c605f3ae9b1a9ded8e4d"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 0372e2135a0e / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-a7f871f23f326555861537ba44ac293fcb58ff8b5e0899cbd3af109faba07ed7)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-4d1526acce9e6cae571471ae42c117ab85b1bb44fe95073a1caf1e3de0cc8d50)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-0d9e5be6600ba97126ae137b748befd063a85c0f348c1947a0ac3001823d9f0a)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-34b26846d372d405a551bae592fa1971e6ba05a941ac1abcaf3a432297aac262"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6f5e9c715a6ec146f67a06a5c503e1e096328df36e1931503ecdde73c922c578"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 0372e2135a0e / 3

<a id="canonical-ab4ca23c72524d8bbbe482c9b5b800878716857f31ae6f3bd0861701579bc67f"></a>

<a id="canonical-ce0a01e33ec9c539b9cf195a8658a5f7acfbd6ad43c8bbe62af7a43cb9c9ea2b"></a>

## addr property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 0372e2135a0e / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-8115501ee869185549697a449beaba350df3e0d62334febc37f7020403dfb392"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 0372e2135a0e / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-0d9e5be6600ba97126ae137b748befd063a85c0f348c1947a0ac3001823d9f0a)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-39edb6c625008498993006e06d6aa5e87d14b9af34d9a66be14856b6b9dbc211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54bfd3d2e2ed9ef7a39f8d82145be7df7be1e9ca26fbee86581d964185f244e5"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 1cacc503735d / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-a7f871f23f326555861537ba44ac293fcb58ff8b5e0899cbd3af109faba07ed7)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-4d1526acce9e6cae571471ae42c117ab85b1bb44fe95073a1caf1e3de0cc8d50)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-0d9e5be6600ba97126ae137b748befd063a85c0f348c1947a0ac3001823d9f0a)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-694346bb645b208d5bffa02256a1e4239d88dc8fdad162ea58a8d25047fc2b5a"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-21b9b0a1c13f4103d366cad0164ba6d14530fde6a16cf63613368bd386cdd266"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 1cacc503735d / 3

<a id="canonical-b8e7b8cbb6865243bb1d468576e64e19023e6162a8ae171afa3ddf8b890407ea"></a>

<a id="canonical-d24e558a886d77d1db8a66124addea97370f885b80850a06f58f9c373b50886d"></a>

## addr property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 1cacc503735d / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-bcf9f98735313bb2a48a65d936bdc9e95fd42a2de35805018d8eb507b9ce0d6f"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 1cacc503735d / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-0d9e5be6600ba97126ae137b748befd063a85c0f348c1947a0ac3001823d9f0a)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-84f8daa4eb4b4b48a06ca1b658217437c459b4aeaed450735bda3328e2ca92b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc0a641fd5f91c0c26defb1e081954d0ff3056a188642a098334b7ef923fe4f0"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 0f4933b36553 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-a7f871f23f326555861537ba44ac293fcb58ff8b5e0899cbd3af109faba07ed7)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-4d1526acce9e6cae571471ae42c117ab85b1bb44fe95073a1caf1e3de0cc8d50)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-049644de7be974ec2f8a6f012b0ca586c50ce02991851156bb1e17cf31da0eb8"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-69a36de08d60765c6d6be27a7fb135ed60acdc3050d87647fa5fc1e944f5cdef"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 0f4933b36553 / 3

<a id="canonical-61ed817f48c82cb1d03628f758bc6cae9013d03b52877104a1ddc0a54404515b"></a>

<a id="canonical-be9c62e1bebe4c787218422619611f4eb2f6f5741f8162bf942a1f2bba5dc766"></a>

## addr property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 0f4933b36553 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-fc87984008ff070b183516f51b4c9a96e7230de7a87313981396104a663337e3"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 0f4933b36553 / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-4d1526acce9e6cae571471ae42c117ab85b1bb44fe95073a1caf1e3de0cc8d50)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-518aa9a897065155b92894bbd2c3e6e7314a083f6b8e4aa40c6694b626f1190c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97a2071cdbf8ca491464bda27e1f6401e029d32fbf6a1a96abdeb2c800abe241"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 70abba2ae751 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-a7f871f23f326555861537ba44ac293fcb58ff8b5e0899cbd3af109faba07ed7)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-4d1526acce9e6cae571471ae42c117ab85b1bb44fe95073a1caf1e3de0cc8d50)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-e8524d59ca1a2589da615a1d8b801491b528eea7c174b9d0cb9f1ac7d1ffc2d1"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-eadfb79c3facd037921597add7febfe29a101fb90162841b9180836b99b8a8b2"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 70abba2ae751 / 3

<a id="canonical-e159d9239d4a539a927fd0a8037ce7cf1bcd691c030aeffed118ffbac69b827b"></a>

<a id="canonical-63fbc1df35556cd12a134867d41a1542f35d7f1be502acc8b8d6dbda6742ee8f"></a>

## addr property — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 70abba2ae751 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-0bd32adc59726e17bde53780885c0865aedf53acc2e422f23ec7be5b66245787"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.ne / 70abba2ae751 / 5

- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-4d1526acce9e6cae571471ae42c117ab85b1bb44fe95073a1caf1e3de0cc8d50)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-9cad8c1bb58d76db38273421360c7368fcfc653a2d0ae9052c3a387bf52cc0a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11516f698e33a3ce7251ae27a0074013dd2c050f8ce587745f633f54cbe38c4a"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.subnets — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets / 5d67ecaf4216 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-257a7a9bf7672c034fcb584f91c93f7b6c15687e3a59fc313d97f79488b11eba"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-79905a9a7c6b68421f12f1c1ea59cd481e467759c06d1f1ea8eaf577ad57f6a7"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets / 5d67ecaf4216 / 3

- [ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-7f7c46a150bc3f6660bb72272a1296219773178a2fcded55da02b922b0152543): complete subsection reference.

- [ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-b675836210bbe0244f79a93ce7c05f241eca5469e738bd75ed04272f52e907ea): complete subsection reference.

<a id="canonical-94cccd32870f93f680990d0355d0446237a3d6d54339540971c2363c1ccef55b"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets / 5d67ecaf4216 / 4

- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-7f7c46a150bc3f6660bb72272a1296219773178a2fcded55da02b922b0152543)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-b675836210bbe0244f79a93ce7c05f241eca5469e738bd75ed04272f52e907ea)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-7f7c46a150bc3f6660bb72272a1296219773178a2fcded55da02b922b0152543"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac8b5555cf39c3927853a4a1322b7db04bb667bedbe3140cc55f2cae0bbdfa85"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 07fc3e64d6bc / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-9cad8c1bb58d76db38273421360c7368fcfc653a2d0ae9052c3a387bf52cc0a1)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-947838441b2543891d4cc5e76533764d64d49ef410211c0b8c831a86c99acae9"></a>

Type: `"single"`. Computed.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-42c3ee21e3cdee3c00e10961c7a1937728d090a4ae7e0aa2afa23d8faea6065f"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 07fc3e64d6bc / 3

<a id="canonical-d1322a713490a10ca60ef016badf9b194c7a63f2a0f7f5496cfeeb6c543894e7"></a>

<a id="canonical-4b25c028c1be97a4b12a8ab96a40816e695447ad14e2bea357fbb1015191795b"></a>

## plen property — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 07fc3e64d6bc / 4

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

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
    }
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

<a id="canonical-421022fea38be249e8397b96885c2ac47976114a9f26c6ae758ffc632ffb7ee7"></a>

<a id="canonical-e5a1c1884044a65997935e822fe583f552c78118a4abec6edb64b54281fb5165"></a>

## prefix property — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 07fc3e64d6bc / 5

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-fff07ad40cdb0acdcb93bd89d2c67471d027aa1e9c7ef2e82402d4ecb82363e6"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 07fc3e64d6bc / 6

- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-9cad8c1bb58d76db38273421360c7368fcfc653a2d0ae9052c3a387bf52cc0a1)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-b675836210bbe0244f79a93ce7c05f241eca5469e738bd75ed04272f52e907ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10e23d13695c5de24a7ce63ad114a227a83d099e881fd7d3de1bb3fbd308c317"></a>

## vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 39c0dc213044 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-591af35859f6df978c00ddc93dd2bc588c40111b981f2ef98348a829317182f1)
- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-e09ea0365ee48d211f64bd35b9cd94b4dee203e1afc23e437cceebb9b8cbe8ca)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-ff0d32b7bc07054ffdb565998e70c0035d29802adefe1b5c4075c92d687e3eaf)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-9cad8c1bb58d76db38273421360c7368fcfc653a2d0ae9052c3a387bf52cc0a1)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-4742df12e337e9d413117b4083bc236cff530ad764f079401da931c0bb5cd79c"></a>

Type: `"single"`. Computed.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-183385eebd2a9d61d189e1861e60c5fade7c6651abec7a756ba7947f86fd5073"></a>

## Direct properties — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 39c0dc213044 / 3

<a id="canonical-7e774c8b919b9a72b5e8b4624b2001c0b4cc4aaeb8f2008673fe19168d2b6d33"></a>

<a id="canonical-ceb7a07da309d91cb2f77186cf609bd6f10a10c4fc3400334285c4e50c68730b"></a>

## plen property — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 39c0dc213044 / 4

Type: `"number"`. Computed.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-40995b42f8baa88050c1865e3ec47b5f2fe82b8b900b7c1b220dc2eb7698cdcf"></a>

<a id="canonical-8fc297be28e4fe75f1e58250219f5e30b67fa2ef9bc00750fdd9a95c6a307d54"></a>

## prefix property — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 39c0dc213044 / 5

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

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

<a id="canonical-9147b39dd915aa257ec05af49f523da06ad8dbf618f32476f95acc1f3075421b"></a>

## Next pages — vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ip / 39c0dc213044 / 6

- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-9cad8c1bb58d76db38273421360c7368fcfc653a2d0ae9052c3a387bf52cc0a1)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-f4fdd17a3c4086ea8ce98a5495a7f9bb1c91e08a52dbec978d7d8c9924bd4b59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-948d01dd0eb996b099712ea9a9622bc08f693a6c36d3c70f4588a84332b74eda"></a>

## vn_config.sm_connection_public_ip — vn_config.sm_connection_public_ip / 67545abf3e80 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- vn_config.sm_connection_public_ip

<a id="canonical-50d3cf912cfaea739a7162aa74e95e3da5282f223ab577100a1506394880c9d2"></a>

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

<a id="canonical-6ec4a4677da955cf5d22ee4841df99e08647b001441314c854d9210e20468234"></a>

## Direct properties — vn_config.sm_connection_public_ip / 67545abf3e80 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aeeea1750600b3f8baf7dbad140664aa3948c846d05819c004045edfdb45b395"></a>

## Next pages — vn_config.sm_connection_public_ip / 67545abf3e80 / 4

- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-3c9cd6e4cf604d890b7093e8b72a560f553f8b1a2ea8ebdbb243c50bc000f7fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4137bd9c06c941411506b801743427a1b837cb4ea7b84b12a26361bb79539e09"></a>

## vn_config.sm_connection_pvt_ip — vn_config.sm_connection_pvt_ip / 82a6b05a601f / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- vn_config.sm_connection_pvt_ip

<a id="canonical-031c93baa4c02b1e6d0859ccda05cd8ddd86b0f57287727d6be962a578e3127b"></a>

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

<a id="canonical-714c84ace65ef71d6f172d7d0d19283efd408bffd21d9f791ea6450ad36dab0b"></a>

## Direct properties — vn_config.sm_connection_pvt_ip / 82a6b05a601f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bea124921c147b01a038d1dfc1cfe98a63d5b546c4d287c5549dcbd6b1a69639"></a>

## Next pages — vn_config.sm_connection_pvt_ip / 82a6b05a601f / 4

- [vn_config](data-sources--aws_tgw_site--reference--group-002.md#canonical-32b10d4c64984020cc625a0762991b797a846eac30bb0524537eaae245096e7c)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-411569971dc5e487a1b7eb5dd5026c6a951d476dc5d2fda3bedd09ca3b953633"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2285594db515fd4c67aabf046f288d170071b1344bf163408bcdd1faf44b754"></a>

## vpc_attachments — vpc_attachments / 0a560572eeee / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- vpc_attachments

<a id="canonical-8e6516a8c16e051650a02da73fb22e1d1c2ba7b9f3bbbaeb4b90e1519cfa4d43"></a>

Type: `"single"`. Computed.

Spoke VPCs to be attached to the AWS TGW Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b18047f734e97912cc02dfaaf6f9948e04aaa8c98d91d52f316b3f26cf3c944f"></a>

## Direct properties — vpc_attachments / 0a560572eeee / 3

- [vpc_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-104588b3ea300c40ea94103c1c27bc7107084b99d8659bb14ae72f72e3d6ad08): complete subsection reference.

<a id="canonical-a430c60413869d2aa9c7848b2a12d5b6657efe35157673ffc647ce7e02aa1b47"></a>

## Next pages — vpc_attachments / 0a560572eeee / 4

- [vpc_attachments.vpc_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-104588b3ea300c40ea94103c1c27bc7107084b99d8659bb14ae72f72e3d6ad08)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-66e7549b1e806db8b6b653eca0836b2fdee052e915e26788c200223cabca3805)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-8be3d7f650880a49403726f80b627b70cc0540cf66813de375cb0af9f65e5e06)

<a id="canonical-104588b3ea300c40ea94103c1c27bc7107084b99d8659bb14ae72f72e3d6ad08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
