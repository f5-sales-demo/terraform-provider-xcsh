---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-91f56c07b8be289ceb8111c7385f7df7da4bc6858727ebfb9d26a0687da5a999"></a>

## port_ranges property — https / 230ec52ee4e6 / 9

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

<a id="canonical-e3c4432876d605aa4932f51500838838cdbe7a91a7bfaf955bc4cee7cbcb433f"></a>

<a id="canonical-122bbe54b5cc58ca84672ed93bd63b26c54963294ccb75495d39cf8abf30ca76"></a>

## server_name property — https / 230ec52ee4e6 / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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

- [tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2): complete subsection reference.

- [tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d): complete subsection reference.

<a id="canonical-17e709854920025a439fd43087ab6f6ea7c6f295f4aa2984ffc02df5e3b9df90"></a>

## Next pages — https / 230ec52ee4e6 / 11

- [https.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-41a23dbccd1c97f2dc1e26c46e237d1a27fb06d4debe54e928158e55e838b0be)
- [https.default_header](data-sources--http_loadbalancer--reference--group-018.md#canonical-45a289eb0b2b5755f9439993b2190f90f49954fd8c1bc921503aa970bf4dddaf)
- [https.default_loadbalancer](data-sources--http_loadbalancer--reference--group-018.md#canonical-5b0efb11f9fd0b7666a36992cab5d99e5ead20fbcc2aa2434c8092674ca51ad1)
- [https.disable_path_normalize](data-sources--http_loadbalancer--reference--group-018.md#canonical-77b368d92971088c9c30dd54008c34933cacc279e688507abbf7369af2972041)
- [https.enable_path_normalize](data-sources--http_loadbalancer--reference--group-018.md#canonical-e5127c793408a0cfb866a5a6fa3cd627e1b66df8afb6371ae5bb2c998660acaf)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-032234c0bf40bef5559c9957b38561f39bb4cd94e21c9e1efcf689c954b3a809)
- [https.non_default_loadbalancer](data-sources--http_loadbalancer--reference--group-018.md#canonical-51f361932fee4d694f1924fb28f3c46f54163b6aad91f8294a4062ebd6122245)
- [https.pass_through](data-sources--http_loadbalancer--reference--group-018.md#canonical-e4ba0983025b8422431c4ada41aeace4ba6fe934fa5855e9b51e2c4a21edc80d)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-41a23dbccd1c97f2dc1e26c46e237d1a27fb06d4debe54e928158e55e838b0be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f406478cadba73b770db022317f13e8a08c62e9d0e07433fc113db59d734b9e3"></a>

## https.coalescing_options — https.coalescing_options / a0beee63b19e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- https.coalescing_options

<a id="canonical-22ad8666b57076dd282cb912d0b673d7672828646a56241fa7d7702556879861"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

<a id="canonical-45eb190316dd70f3d3708cdee1eaaf61eb20738db0010566546c52d2e5c312ce"></a>

## Direct properties — https.coalescing_options / a0beee63b19e / 3

- [default_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-6c35deb023f3d766590452021c910c721850463fe8fd392e9538bfa8b09dbd68): complete subsection reference.

- [strict_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-b441f2eda76dee4598ec13ba7887ba3bb2acd39b8aff8f3ccc23ace5845b8b26): complete subsection reference.

<a id="canonical-9b89525d06d8daec7a7f97085d2e16e2f3e9fac1e1249594b5ba5994a00e1c56"></a>

## Next pages — https.coalescing_options / a0beee63b19e / 4

- [https.coalescing_options.default_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-6c35deb023f3d766590452021c910c721850463fe8fd392e9538bfa8b09dbd68)
- [https.coalescing_options.strict_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-b441f2eda76dee4598ec13ba7887ba3bb2acd39b8aff8f3ccc23ace5845b8b26)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6c35deb023f3d766590452021c910c721850463fe8fd392e9538bfa8b09dbd68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c15039253d7715a62cdb733bb82aaf446bc889b00e8cb64b8f07d2255532e362"></a>

## https.coalescing_options.default_coalescing — https.coalescing_options.default_coalescing / ec4394703b3b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-41a23dbccd1c97f2dc1e26c46e237d1a27fb06d4debe54e928158e55e838b0be)
- https.coalescing_options.default_coalescing

<a id="canonical-f89dc5b4af5d4f4c9fd2421d472565e35fec0b08d486a84b845a787ef7bff833"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-449837ebff4649509472a5c83e18ccf70d64a0bff9761333344a24992599a0a5"></a>

## Direct properties — https.coalescing_options.default_coalescing / ec4394703b3b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1a0cd5f9eb81a6c07b9072c5d119ed9f615f54a2d0727229facba48d0c5d4618"></a>

## Next pages — https.coalescing_options.default_coalescing / ec4394703b3b / 4

- [https.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-41a23dbccd1c97f2dc1e26c46e237d1a27fb06d4debe54e928158e55e838b0be)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b441f2eda76dee4598ec13ba7887ba3bb2acd39b8aff8f3ccc23ace5845b8b26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f93eb162efbca3d857c02b6b07138c3ad5a02cfb661fa381b8827edde82153a"></a>

## https.coalescing_options.strict_coalescing — https.coalescing_options.strict_coalescing / cc99ddd44dc7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-41a23dbccd1c97f2dc1e26c46e237d1a27fb06d4debe54e928158e55e838b0be)
- https.coalescing_options.strict_coalescing

<a id="canonical-41a3f8649f9f088a11d461250eba177cf824ffd54254717d9ac8b7ecb7220c92"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-f7991d57c40e6dd189882d3a3ac7e78a0a41cd3c3b073c6157761462f71eafb5"></a>

## Direct properties — https.coalescing_options.strict_coalescing / cc99ddd44dc7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-172544d2dfba6b6981093cf3e9650ecd887c593d57ce40a610da08e71b0a9d8a"></a>

## Next pages — https.coalescing_options.strict_coalescing / cc99ddd44dc7 / 4

- [https.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-41a23dbccd1c97f2dc1e26c46e237d1a27fb06d4debe54e928158e55e838b0be)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-45a289eb0b2b5755f9439993b2190f90f49954fd8c1bc921503aa970bf4dddaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-872ae6220df15cc8952e5bcc53d6299531223d62810e1021543b5e84b7e86d9f"></a>

## https.default_header — https.default_header / 52298fa05763 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- https.default_header

<a id="canonical-2a4deee4e416644a824af6429adb989fdb99892e2999324460d710059a9c2ef7"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-8152109aed1a4186a94dd351a00aa383a5e7a7fa41cdbfb2fc365fb2f284a8a0"></a>

## Direct properties — https.default_header / 52298fa05763 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-656061e8d9d155b6c361ec7d3375ecc52b08d15a7b43aa9317f6b1273683e0b0"></a>

## Next pages — https.default_header / 52298fa05763 / 4

- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5b0efb11f9fd0b7666a36992cab5d99e5ead20fbcc2aa2434c8092674ca51ad1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-166a73d43e8299caf5b3e46440480dabadd0c12bbed939013158cc444a63779f"></a>

## https.default_loadbalancer — https.default_loadbalancer / 2ecfa820102c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- https.default_loadbalancer

<a id="canonical-bd37c97b62cd47d0abde51a577da0f975d919633fe440a8f80765df8e1467ddf"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

<a id="canonical-82a68df1c7a09b8e1302265c90e3544c89ff64b83054f1513e70ba37ca667572"></a>

## Direct properties — https.default_loadbalancer / 2ecfa820102c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9bc39856702cd43c31d91b9f5ac11c6cd047ba0f58471af714270b04acd3db94"></a>

## Next pages — https.default_loadbalancer / 2ecfa820102c / 4

- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-77b368d92971088c9c30dd54008c34933cacc279e688507abbf7369af2972041"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d295c09a1a1bcaef637645ea0edc8360d6755ef3d1cb88a28a13a3629107f73d"></a>

## https.disable_path_normalize — https.disable_path_normalize / e1d5f3cba5f5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- https.disable_path_normalize

<a id="canonical-066f49719a9c3390e642b4f54ead21e3f56f8a95cdc49713b38c806fe9fd3c66"></a>

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

<a id="canonical-0f6e8841fe7d0b4b83c2ae4824da7f373361915460cb0c24a624507ad39931e7"></a>

## Direct properties — https.disable_path_normalize / e1d5f3cba5f5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-219ccc715f0d894707095aee06c62e6d9a326c139538afb937b5aaa55d6acd96"></a>

## Next pages — https.disable_path_normalize / e1d5f3cba5f5 / 4

- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e5127c793408a0cfb866a5a6fa3cd627e1b66df8afb6371ae5bb2c998660acaf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79575a8be5668cb450943f0ded961f307787c45ddbfd5d73e5224e8b88d7e743"></a>

## https.enable_path_normalize — https.enable_path_normalize / 52ea090cd4c5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- https.enable_path_normalize

<a id="canonical-432f2a2ab2f6f0f1972ac711e9e2ec1cd8e314101ff04e793f31a81c8e227e3d"></a>

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

<a id="canonical-b92f8bc6149fcc56cf2ba079688fa4b1692c06c33585326ba6a492d35f6bc6ab"></a>

## Direct properties — https.enable_path_normalize / 52ea090cd4c5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7bda1084268b0738c4623af9bd4ec8c4b6312385917735a5bea4817b91d4a16f"></a>

## Next pages — https.enable_path_normalize / 52ea090cd4c5 / 4

- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-032234c0bf40bef5559c9957b38561f39bb4cd94e21c9e1efcf689c954b3a809"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d8fdd8266631ab10cd973d95acd79979774fd8ae682dfc0b1aa3138df2d67d9"></a>

## https.http_protocol_options — https.http_protocol_options / 4f8925f2dc47 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- https.http_protocol_options

<a id="canonical-3dfe49d44012c268a276c40bb5d5d55f3e610aeeaebe28710ece76bfb7ff4307"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

<a id="canonical-7f37f84aea40c9e53d7c93f8c6c693e00b68f93d8db645dfff6e44c301abf24e"></a>

## Direct properties — https.http_protocol_options / 4f8925f2dc47 / 3

- [http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-9db72b05f4eaef658c07c46523e7e11764263766af2ec88e20bf821228d878fa): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--http_loadbalancer--reference--group-018.md#canonical-dbfe546f8597c5835a20d9623ea57e42200b029dc33d476e2c4669ed69bb6f8a): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-5c0dc5fbfbc307f742db02474b6807a12d46181fbd443841e81687bea5ebe59f): complete subsection reference.

<a id="canonical-117476ba317631e86067ad8bb27b88fb4453c4ab07edd027ab6618241f21a10b"></a>

## Next pages — https.http_protocol_options / 4f8925f2dc47 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-9db72b05f4eaef658c07c46523e7e11764263766af2ec88e20bf821228d878fa)
- [https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--http_loadbalancer--reference--group-018.md#canonical-dbfe546f8597c5835a20d9623ea57e42200b029dc33d476e2c4669ed69bb6f8a)
- [https.http_protocol_options.http_protocol_enable_v2_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-5c0dc5fbfbc307f742db02474b6807a12d46181fbd443841e81687bea5ebe59f)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9db72b05f4eaef658c07c46523e7e11764263766af2ec88e20bf821228d878fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c96f18e44f41d819b48dc099468e5751ed9bac55d46152f18cc612493f585b88"></a>

## https.http_protocol_options.http_protocol_enable_v1_only — https.http_protocol_options.http_protocol_enable_v1_only / 4f980bd4cddd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-032234c0bf40bef5559c9957b38561f39bb4cd94e21c9e1efcf689c954b3a809)
- https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-097772634808951116dc007500bd87e0fd9131a7551fb735b1d2cb1a5e9760d3"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-28841dad2c2bda03dad361059b31d48f0c4feeb229e8741421d18819ea87ac52"></a>

## Direct properties — https.http_protocol_options.http_protocol_enable_v1_only / 4f980bd4cddd / 3

- [header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-3426dd5e34f8427dc86fc06e8a07782aa8be5c94ec40a46f268a67e10ba5054c): complete subsection reference.

<a id="canonical-4109fed29b338120f6c359b45de8057a7786adcc82599212af9557fcda7e9432"></a>

## Next pages — https.http_protocol_options.http_protocol_enable_v1_only / 4f980bd4cddd / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-3426dd5e34f8427dc86fc06e8a07782aa8be5c94ec40a46f268a67e10ba5054c)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-032234c0bf40bef5559c9957b38561f39bb4cd94e21c9e1efcf689c954b3a809)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3426dd5e34f8427dc86fc06e8a07782aa8be5c94ec40a46f268a67e10ba5054c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9765740c8daa1d7ed302df5b22a4775f0d7cf31417327377736e93ffad4d2c03"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation / f69120cf8773 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-032234c0bf40bef5559c9957b38561f39bb4cd94e21c9e1efcf689c954b3a809)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-9db72b05f4eaef658c07c46523e7e11764263766af2ec88e20bf821228d878fa)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-054e60d0da07ba142bcf526fdd8c7d643136dd080ea5cd5dd1c1ee6910fac22e"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

<a id="canonical-02ddbe0d9fa2860f04e0c48aea481666b2b3453ec0adb07a98df05cd2908928e"></a>

## Direct properties — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation / f69120cf8773 / 3

- [default_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-26b56b73d4fa27aa960672360a507ee4d5ac0467ff911635dd4ab3c51473ccf5): complete subsection reference.

- [preserve_case_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-11367fda5aec3684661496a2de501b764a91fcd6a69df05c567e73091841f98b): complete subsection reference.

- [proper_case_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-5d185b8dc5a1d34dabdef067081064f689cfdaf1d57ea2e6e3b7bf8426b38ddf): complete subsection reference.

<a id="canonical-e56fcd1b4e276f7e23af659e7f5272e48fcf3841fc6f7e7d3f27eb65ba933d3e"></a>

## Next pages — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation / f69120cf8773 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-26b56b73d4fa27aa960672360a507ee4d5ac0467ff911635dd4ab3c51473ccf5)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-11367fda5aec3684661496a2de501b764a91fcd6a69df05c567e73091841f98b)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-5d185b8dc5a1d34dabdef067081064f689cfdaf1d57ea2e6e3b7bf8426b38ddf)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-9db72b05f4eaef658c07c46523e7e11764263766af2ec88e20bf821228d878fa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-26b56b73d4fa27aa960672360a507ee4d5ac0467ff911635dd4ab3c51473ccf5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e890d398c7b060885c8c349a29cfd303505f0611647c129d4ee7e28513d49cc1"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.d / 02a285a74a56 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-032234c0bf40bef5559c9957b38561f39bb4cd94e21c9e1efcf689c954b3a809)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-9db72b05f4eaef658c07c46523e7e11764263766af2ec88e20bf821228d878fa)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-3426dd5e34f8427dc86fc06e8a07782aa8be5c94ec40a46f268a67e10ba5054c)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-2b9da6cd63f4c7d48f01e58669fdc86f1b46900054dbe6ab684592758af5a68a"></a>

Type: `["object", {}]`. Computed.

Use the platform's current default HTTP header transformation behavior.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1f1352d460af2e2bfb18bed1060215b48da2cbfe12eac8bf7fea8980d65a242a"></a>

## Direct properties — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.d / 02a285a74a56 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7fee6223a7f09ccaac258bff60ea0fb08eb85a7e89c290f939225f7f53bf8d37"></a>

## Next pages — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.d / 02a285a74a56 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-3426dd5e34f8427dc86fc06e8a07782aa8be5c94ec40a46f268a67e10ba5054c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-11367fda5aec3684661496a2de501b764a91fcd6a69df05c567e73091841f98b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23655cb42347f8676950e671fd2cae8e7595af7163f21d77c7b817821b3b3c03"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.p / d2d143132117 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-032234c0bf40bef5559c9957b38561f39bb4cd94e21c9e1efcf689c954b3a809)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-9db72b05f4eaef658c07c46523e7e11764263766af2ec88e20bf821228d878fa)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-3426dd5e34f8427dc86fc06e8a07782aa8be5c94ec40a46f268a67e10ba5054c)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-1a0907a5403e9edd9c84ff35d9c9431d0f0f61f44ffb100d1653687b676b9eb5"></a>

Type: `["object", {}]`. Computed.

Preserve HTTP header-name case when upstream case must remain unchanged.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1c2eaea42b686d9ba5c89b7f18c6b57b750ae70a275a1fd4579228e6686d70a6"></a>

## Direct properties — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.p / d2d143132117 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-948dcc7916ea31e9b30cbe3c68261be4e64b2ed17ce99b45e16779eaecc615e4"></a>

## Next pages — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.p / d2d143132117 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-3426dd5e34f8427dc86fc06e8a07782aa8be5c94ec40a46f268a67e10ba5054c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5d185b8dc5a1d34dabdef067081064f689cfdaf1d57ea2e6e3b7bf8426b38ddf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c70a8c855729acc1f1a47de206f210e3e93654687577bfc026ed69f3568a98f6"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.p / 845e6b2d41d9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-032234c0bf40bef5559c9957b38561f39bb4cd94e21c9e1efcf689c954b3a809)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-9db72b05f4eaef658c07c46523e7e11764263766af2ec88e20bf821228d878fa)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-3426dd5e34f8427dc86fc06e8a07782aa8be5c94ec40a46f268a67e10ba5054c)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-5e633f2ba7c83b7312ff6719e854bb2edbabafbf543ca3842a1027635d2ae183"></a>

Type: `["object", {}]`. Computed.

Transform HTTP header names to proper case when explicit transformation is required.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6d06c52362f42c93b333e5d8c765ffad23f47102d4ca8db43ca2d288f0eb0b94"></a>

## Direct properties — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.p / 845e6b2d41d9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1854a07d46d06fc155fcc30d7b8bc75445097ef202bc2c1da96eb49d275363d0"></a>

## Next pages — https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.p / 845e6b2d41d9 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-3426dd5e34f8427dc86fc06e8a07782aa8be5c94ec40a46f268a67e10ba5054c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-dbfe546f8597c5835a20d9623ea57e42200b029dc33d476e2c4669ed69bb6f8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf79bd27c87498f7cc9b1e2e119025914866fffe38b3e7a1bc81075714c123a5"></a>

## https.http_protocol_options.http_protocol_enable_v1_v2 — https.http_protocol_options.http_protocol_enable_v1_v2 / 6b5a939096a2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-032234c0bf40bef5559c9957b38561f39bb4cd94e21c9e1efcf689c954b3a809)
- https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-8ad6fdc53b0226ea9926206080acc19fab345ecccbcaf5ad44fe9c3710ace713"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

<a id="canonical-497ad4738120c80b9dfc96b95f824ac3685b14903ff27106200ef6ba6d365d52"></a>

## Direct properties — https.http_protocol_options.http_protocol_enable_v1_v2 / 6b5a939096a2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9cf44d96c25eb82312fcebfdaec4087978280e21d6ddf4f8ad20b6b8f583a244"></a>

## Next pages — https.http_protocol_options.http_protocol_enable_v1_v2 / 6b5a939096a2 / 4

- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-032234c0bf40bef5559c9957b38561f39bb4cd94e21c9e1efcf689c954b3a809)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5c0dc5fbfbc307f742db02474b6807a12d46181fbd443841e81687bea5ebe59f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73de2b2de6cd67e12f2fa3a570b2bf105ec20a69619785b33744d8769eb3d92f"></a>

## https.http_protocol_options.http_protocol_enable_v2_only — https.http_protocol_options.http_protocol_enable_v2_only / b8b2c8c044e9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-032234c0bf40bef5559c9957b38561f39bb4cd94e21c9e1efcf689c954b3a809)
- https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-a24675da51180b46803e53c5e1d53caec6a390bf57a42f77eab54f87ed0165fe"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-cfc006f39fbb81b36f4d9de87f6575498dc049143a95d5c8d1b45ef2d5ba23bb"></a>

## Direct properties — https.http_protocol_options.http_protocol_enable_v2_only / b8b2c8c044e9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e6fb28edaedcc069f3c145289eb3fd3400ba8026daa3a07d67630d882c8e148c"></a>

## Next pages — https.http_protocol_options.http_protocol_enable_v2_only / b8b2c8c044e9 / 4

- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-032234c0bf40bef5559c9957b38561f39bb4cd94e21c9e1efcf689c954b3a809)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-51f361932fee4d694f1924fb28f3c46f54163b6aad91f8294a4062ebd6122245"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-035fc1fd35d2817e00deeac4ca4f5df1a1f89d7ff51535e2716d3b23a950235e"></a>

## https.non_default_loadbalancer — https.non_default_loadbalancer / f155daaf954b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- https.non_default_loadbalancer

<a id="canonical-cafba7e1da57ea74b3d56360c6f9a4d9fcc838c709d1341bbc5d3159ca5251b8"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

<a id="canonical-1f54b9672cb3d5644f192861a67bf858798bc4517dbed32646c877015dd92601"></a>

## Direct properties — https.non_default_loadbalancer / f155daaf954b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b34abecc1c6eea01bb6532a786d7c31399be437b7584c3728f841e040e4bfe4d"></a>

## Next pages — https.non_default_loadbalancer / f155daaf954b / 4

- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e4ba0983025b8422431c4ada41aeace4ba6fe934fa5855e9b51e2c4a21edc80d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4dcd818f95e2e8921f61fc9c39f0285436ad326d3208b7a12450d8f16a6c96a"></a>

## https.pass_through — https.pass_through / e6efdea36eb5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- https.pass_through

<a id="canonical-69caaca5058403c938113a2394376ea0d7a9a4926840cb90279a392b2cec181f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

<a id="canonical-c9c7b625dc5aa0d1e6ab28e45d2b1468166ffcc864caed1aded2c2253ee47d97"></a>

## Direct properties — https.pass_through / e6efdea36eb5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-83d2e4d42a39d97acb8e2de170d577bcd847ba88fea9acc1229b0931991c400d"></a>

## Next pages — https.pass_through / e6efdea36eb5 / 4

- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb1affafb0351fc1ab92d1a7d1dcb7bb4e00e60936989cc9f7873f22352417c6"></a>

## https.tls_cert_params — https.tls_cert_params / 2d3ece441508 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- https.tls_cert_params

<a id="canonical-5d33a6acb9c89b2062c9bdaa5bc9aa2808a3bbf9878f1864789a0b568b712ea2"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

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

<a id="canonical-b72701ac0f7e6e3db2f7b8cc05195aa6f49b86c886e0449ef4b38fa7a85a8faf"></a>

## Direct properties — https.tls_cert_params / 2d3ece441508 / 3

- [certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-b68e119f5e7fe21416f9f43aa6820b09233d44327149147ce0e13d3f5bab743f): complete subsection reference.

- [no_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-2b7d2d88ef8c809fb310a9e1e49506a8e9733496b3a03c988beb177950698248): complete subsection reference.

- [tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-3c9c2404e46f71e4154a80936a806e9a7b8dca7b711ecbcdd445aebf0f3466c1): complete subsection reference.

- [use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-109437c68a412d0b677ccd3f56d7ce7b57300844c679372dd4cffee7325e0f90): complete subsection reference.

<a id="canonical-74458db2e0be16dfc1dd474cc3ce01544966cff368f66797172174a0b171f43b"></a>

## Next pages — https.tls_cert_params / 2d3ece441508 / 4

- [https.tls_cert_params.certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-b68e119f5e7fe21416f9f43aa6820b09233d44327149147ce0e13d3f5bab743f)
- [https.tls_cert_params.no_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-2b7d2d88ef8c809fb310a9e1e49506a8e9733496b3a03c988beb177950698248)
- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-3c9c2404e46f71e4154a80936a806e9a7b8dca7b711ecbcdd445aebf0f3466c1)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-109437c68a412d0b677ccd3f56d7ce7b57300844c679372dd4cffee7325e0f90)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b68e119f5e7fe21416f9f43aa6820b09233d44327149147ce0e13d3f5bab743f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f32ea38a2cb574a46a7fc581babad67d2fccf48fe84b4f0313cf47c28671cc9"></a>

## https.tls_cert_params.certificates — https.tls_cert_params.certificates / cfbbb831a7a7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- https.tls_cert_params.certificates

<a id="canonical-5849161bb86cf2944b34813ecbf83371736f4427bb7788580b3f6ce9aebce8e5"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

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

<a id="canonical-caf20bb37d6056c0ccfec16094297b7aa48262f9e391131058ddf72ae14ec39b"></a>

## Direct properties — https.tls_cert_params.certificates / cfbbb831a7a7 / 3

<a id="canonical-edb07e142a22a4dbb4b88290ab2b4f2be83532061ddefaf4200cf2a8c7ce47a7"></a>

<a id="canonical-63bd6601509efb583e253cc2143448121f05ba772295bbfdb03237b41e40fed5"></a>

## name property — https.tls_cert_params.certificates / cfbbb831a7a7 / 4

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

<a id="canonical-cee7274622589729015879bd35a669f768c1910f0d6f8939100ecae77afcdee3"></a>

<a id="canonical-7d21966859b968cba081bc482901e08d33e30132d4e320e0620f693f268a6bc5"></a>

## namespace property — https.tls_cert_params.certificates / cfbbb831a7a7 / 5

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

<a id="canonical-4301a2d77a76a46549e47bb074f919c392e06ba4d96cd99642142ba14952a1d2"></a>

<a id="canonical-74abbd301d9a5cfb0fce3267826c8fffc4897a9341c0e432ab6a3a63d8898d88"></a>

## tenant property — https.tls_cert_params.certificates / cfbbb831a7a7 / 6

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

<a id="canonical-88e6349ba9a06795c626e8825da3647854bb5f6c1dd6abed28d2c1063a9ff2fc"></a>

## Next pages — https.tls_cert_params.certificates / cfbbb831a7a7 / 7

- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2b7d2d88ef8c809fb310a9e1e49506a8e9733496b3a03c988beb177950698248"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c04f6c3b4e7b028c636608d3cbcc398795e9a5f0bc46f813a6863f665dadf6dc"></a>

## https.tls_cert_params.no_mtls — https.tls_cert_params.no_mtls / 8785603c760e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- https.tls_cert_params.no_mtls

<a id="canonical-66711c7060d5a82b5206b34065433ed3f0937092f165d27caa41f2fd2d93e157"></a>

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

<a id="canonical-6460d46f5e268dbcb159076ae3ecdd283c684035dece704ecc03992f64b2ea71"></a>

## Direct properties — https.tls_cert_params.no_mtls / 8785603c760e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2c898af38109bc5a1720ad595aee1fcef771b319750453c575968fdd30ef4eb0"></a>

## Next pages — https.tls_cert_params.no_mtls / 8785603c760e / 4

- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3c9c2404e46f71e4154a80936a806e9a7b8dca7b711ecbcdd445aebf0f3466c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30aa5746b0cc67c57dbd31b770d71e889fdc5ffc1c8dfeb9c8e55ca608cc9c93"></a>

## https.tls_cert_params.tls_config — https.tls_cert_params.tls_config / 84a8427430ae / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- https.tls_cert_params.tls_config

<a id="canonical-defbbfbb7c8554acb3ab11031488a04ae5d063a24a078bcd6c269dff4a599454"></a>

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

<a id="canonical-b554a1b9a80dc382f94e3fce7bf581fbc04e1de3b59bbfb30f6bf3f7922c7fd2"></a>

## Direct properties — https.tls_cert_params.tls_config / 84a8427430ae / 3

- [custom_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-b86db2c7a4a1ac8034231874f66abe4cb91bd6a87eb6860e9f3662f82cd824c3): complete subsection reference.

- [default_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-766d90139d20b9f920e1d32671ec3d98e1307e3d395fc1e868ad211a0fc29bd3): complete subsection reference.

- [low_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-02a1437a802bc08b643ceea1aa4b0373d22314dd991ded09269c64dea1f22e5a): complete subsection reference.

- [medium_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-e3397b4536758220a0d0a85a1c26926af6e811726d97ba68f81fb3ea693c773e): complete subsection reference.

<a id="canonical-bcb260f96137851ed93b44b7ce791e7f84afdc3fbc44032ed4499d8f75fbbd56"></a>

## Next pages — https.tls_cert_params.tls_config / 84a8427430ae / 4

- [https.tls_cert_params.tls_config.custom_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-b86db2c7a4a1ac8034231874f66abe4cb91bd6a87eb6860e9f3662f82cd824c3)
- [https.tls_cert_params.tls_config.default_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-766d90139d20b9f920e1d32671ec3d98e1307e3d395fc1e868ad211a0fc29bd3)
- [https.tls_cert_params.tls_config.low_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-02a1437a802bc08b643ceea1aa4b0373d22314dd991ded09269c64dea1f22e5a)
- [https.tls_cert_params.tls_config.medium_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-e3397b4536758220a0d0a85a1c26926af6e811726d97ba68f81fb3ea693c773e)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b86db2c7a4a1ac8034231874f66abe4cb91bd6a87eb6860e9f3662f82cd824c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bc64992f9d1c554b8c515cc14a5f3dd266307e01524c4dbccc32a6044e90fe9"></a>

## https.tls_cert_params.tls_config.custom_security — https.tls_cert_params.tls_config.custom_security / 335cec033c66 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-3c9c2404e46f71e4154a80936a806e9a7b8dca7b711ecbcdd445aebf0f3466c1)
- https.tls_cert_params.tls_config.custom_security

<a id="canonical-9e7a4f9a7335dde3d8a91468d1067f8357c44d2a8b97122451cddbdfa24687d4"></a>

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

<a id="canonical-bea47329937aaa3c02be1564fb248b472d20b4239ac3c58d4df76b0048dd5d5c"></a>

## Direct properties — https.tls_cert_params.tls_config.custom_security / 335cec033c66 / 3

<a id="canonical-d503ca91402fe1f52ef7e76c9af6e5604846cb426eb1a331b8dd0678425260eb"></a>

<a id="canonical-1cdd4ed3e7a3b7acaca45b768341cfd87c68f732b432d51ca8c8a81f7eee4557"></a>

## cipher_suites property — https.tls_cert_params.tls_config.custom_security / 335cec033c66 / 4

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

<a id="canonical-6a77d36ec4a60fec01cb0bf7c8b230d048c366189b6b752be6b35818c2934e69"></a>

<a id="canonical-e6b1dc5825822576ce83b311cbc6b18a9730ba45629793414a22a2b47182f6e7"></a>

## max_version property — https.tls_cert_params.tls_config.custom_security / 335cec033c66 / 5

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

<a id="canonical-a95df658483804d18fb83f687884d9415903867087b986fa5a9bad2f615302ce"></a>

<a id="canonical-f75a4d02f68d30abe2ed06f08b067fad229164c5234212e0bf4cfbd602ac6113"></a>

## min_version property — https.tls_cert_params.tls_config.custom_security / 335cec033c66 / 6

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

<a id="canonical-8274b9593aa333f7578988e461342be31f3acb4d0e1277337c8348567033272a"></a>

## Next pages — https.tls_cert_params.tls_config.custom_security / 335cec033c66 / 7

- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-3c9c2404e46f71e4154a80936a806e9a7b8dca7b711ecbcdd445aebf0f3466c1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-766d90139d20b9f920e1d32671ec3d98e1307e3d395fc1e868ad211a0fc29bd3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29fd985fd585913f5c48b8fd63f65544fc36d02cad5b94e96c994a2d70a6a101"></a>

## https.tls_cert_params.tls_config.default_security — https.tls_cert_params.tls_config.default_security / 663eadd912f9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-3c9c2404e46f71e4154a80936a806e9a7b8dca7b711ecbcdd445aebf0f3466c1)
- https.tls_cert_params.tls_config.default_security

<a id="canonical-fa83a622a001902ef5ea978299dd358f31f8082a931f0840d003394548f0eb5d"></a>

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

<a id="canonical-2d9e918fca622cc908a3d2215f4616b6831f10f86d7a9b8cb4c0999714654d0f"></a>

## Direct properties — https.tls_cert_params.tls_config.default_security / 663eadd912f9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0342ac838bc592ed126a263914747df44b34c35e14c4d7fefa703a7cd31ae9cc"></a>

## Next pages — https.tls_cert_params.tls_config.default_security / 663eadd912f9 / 4

- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-3c9c2404e46f71e4154a80936a806e9a7b8dca7b711ecbcdd445aebf0f3466c1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-02a1437a802bc08b643ceea1aa4b0373d22314dd991ded09269c64dea1f22e5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41445696295a1d3a7a8d00359a7d555e7aa932509bd565792ec971405f68b424"></a>

## https.tls_cert_params.tls_config.low_security — https.tls_cert_params.tls_config.low_security / 694ef3d5d772 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-3c9c2404e46f71e4154a80936a806e9a7b8dca7b711ecbcdd445aebf0f3466c1)
- https.tls_cert_params.tls_config.low_security

<a id="canonical-15926dd886c765f7f6d0f837e0102ca8f5cfd77907868e63634c8a4d2b38a33b"></a>

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

<a id="canonical-6652a62ce10db723da7a983173cc5ec95592093bd1acdf1f825bef95c2eeb2f8"></a>

## Direct properties — https.tls_cert_params.tls_config.low_security / 694ef3d5d772 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5686fc5e8476a8e86ac8f193b8879fa4e7abc7b840db6a5961bdfa3d69773d37"></a>

## Next pages — https.tls_cert_params.tls_config.low_security / 694ef3d5d772 / 4

- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-3c9c2404e46f71e4154a80936a806e9a7b8dca7b711ecbcdd445aebf0f3466c1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e3397b4536758220a0d0a85a1c26926af6e811726d97ba68f81fb3ea693c773e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03623bcea2f29a64be3ab0c1f3e761ca38de6c0b5dddc07b4a7b34ac7e0ec813"></a>

## https.tls_cert_params.tls_config.medium_security — https.tls_cert_params.tls_config.medium_security / 4035d5b5f944 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-3c9c2404e46f71e4154a80936a806e9a7b8dca7b711ecbcdd445aebf0f3466c1)
- https.tls_cert_params.tls_config.medium_security

<a id="canonical-3f3d16e108e99339fe7f981ae0e9644e64e8da382ba618b9670622d4e6522c2d"></a>

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

<a id="canonical-d860fa2618e45004aea8a61a63b4ad5189baf93a0330324314c65d5c7e6a1492"></a>

## Direct properties — https.tls_cert_params.tls_config.medium_security / 4035d5b5f944 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-674e1e35f99005c55baaf919c382bfed0e7f127d599c9c277fa7261ac583e990"></a>

## Next pages — https.tls_cert_params.tls_config.medium_security / 4035d5b5f944 / 4

- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-3c9c2404e46f71e4154a80936a806e9a7b8dca7b711ecbcdd445aebf0f3466c1)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-109437c68a412d0b677ccd3f56d7ce7b57300844c679372dd4cffee7325e0f90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d94264dd0abd1d8fd4d9c4ce9cf885995e8dc176574e9f1a57f89e93d4b0dc5b"></a>

## https.tls_cert_params.use_mtls — https.tls_cert_params.use_mtls / 48e7c93923b8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- https.tls_cert_params.use_mtls

<a id="canonical-413f6b26653b97e8d5c2bf4c26e34b861ca0fa63a568ef08a1874d1ea4ab90fb"></a>

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

<a id="canonical-a88b0a95701e5ae6dada992f64f7e2a1ef82d7ff4ac2e36bc5618f22baf45b56"></a>

## Direct properties — https.tls_cert_params.use_mtls / 48e7c93923b8 / 3

<a id="canonical-7096ef9f7487d38819d7014ca8d79c0c07622500b1d126710f3d9fb705ad6695"></a>

<a id="canonical-f8da3d77407e89a223fea3cedb41f0fe22c364cbc06f62daa073c920ffca5144"></a>

## client_certificate_optional property — https.tls_cert_params.use_mtls / 48e7c93923b8 / 4

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

- [crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-9e4f14895e4da46cdb2a416838327bbdb5e6302329c997a85bb6523b4ef44eba): complete subsection reference.

- [no_crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-e77b079a70be8330a0060021eb26d2a14dc71bbc055d92e0e2b36ec230477e4e): complete subsection reference.

- [trusted_ca](data-sources--http_loadbalancer--reference--group-018.md#canonical-69e4bceb3760b19dd0ee81e996049025b2d4a7f3a0cd0f57e0d41ae07b5e7f9c): complete subsection reference.

<a id="canonical-f895b4a415fa405b2dc3350599348864875249c2aba9728a6d30fffc5db10731"></a>

<a id="canonical-085080450739a8560063499ac8b71e56484880edb1c34ebe718f6d07097c3652"></a>

## trusted_ca_url property — https.tls_cert_params.use_mtls / 48e7c93923b8 / 5

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

- [xfcc_disabled](data-sources--http_loadbalancer--reference--group-018.md#canonical-115fe7e085800b0d0e878431bf01379919a570c4e89663b84b3551a6576fb6c3): complete subsection reference.

- [xfcc_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-c7630ff79e850d0c74a1ab8cf27647d01f33d29db0913277bad5bf6b1bb058ab): complete subsection reference.

<a id="canonical-cd2254fb3569bfe588839d6f5f55a36f6936c389f9123459ba48db2888e762ba"></a>

## Next pages — https.tls_cert_params.use_mtls / 48e7c93923b8 / 6

- [https.tls_cert_params.use_mtls.crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-9e4f14895e4da46cdb2a416838327bbdb5e6302329c997a85bb6523b4ef44eba)
- [https.tls_cert_params.use_mtls.no_crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-e77b079a70be8330a0060021eb26d2a14dc71bbc055d92e0e2b36ec230477e4e)
- [https.tls_cert_params.use_mtls.trusted_ca](data-sources--http_loadbalancer--reference--group-018.md#canonical-69e4bceb3760b19dd0ee81e996049025b2d4a7f3a0cd0f57e0d41ae07b5e7f9c)
- [https.tls_cert_params.use_mtls.xfcc_disabled](data-sources--http_loadbalancer--reference--group-018.md#canonical-115fe7e085800b0d0e878431bf01379919a570c4e89663b84b3551a6576fb6c3)
- [https.tls_cert_params.use_mtls.xfcc_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-c7630ff79e850d0c74a1ab8cf27647d01f33d29db0913277bad5bf6b1bb058ab)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9e4f14895e4da46cdb2a416838327bbdb5e6302329c997a85bb6523b4ef44eba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6dd32743ebe82446cd703e0065382aca117d99a29743273560a3f958eab99f7"></a>

## https.tls_cert_params.use_mtls.crl — https.tls_cert_params.use_mtls.crl / c67fdf16de50 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-109437c68a412d0b677ccd3f56d7ce7b57300844c679372dd4cffee7325e0f90)
- https.tls_cert_params.use_mtls.crl

<a id="canonical-64aa35e23f198a0fcced76055f567e5f81f8091ce8d468d1d08448190845f070"></a>

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

<a id="canonical-2468da465f0449d1c45310aeaeae896c896bf5c06bc027f7e0619c77ca408d35"></a>

## Direct properties — https.tls_cert_params.use_mtls.crl / c67fdf16de50 / 3

<a id="canonical-9a8e1e1365b64ccbf7f04e673d08ec3c51e6af78809094a0257c375d0183205f"></a>

<a id="canonical-513c8c4a7b309786fc2c9f32ab3a71118ee122852f20673817899d2971e6e496"></a>

## name property — https.tls_cert_params.use_mtls.crl / c67fdf16de50 / 4

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

<a id="canonical-b371593557df1c6b3472d85958d527586ddee2d2319923bf0de1c50c72eb7f77"></a>

<a id="canonical-d8b660f9942624828ebb9507778ead013667b0f7c0335a10dabcfe8b5271bcd5"></a>

## namespace property — https.tls_cert_params.use_mtls.crl / c67fdf16de50 / 5

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

<a id="canonical-fefc5f71a32b8328ad7b41d3cb58b684fb56f9cc19ce0d28fb68f00b5db633c6"></a>

<a id="canonical-74b4ad73ebaa12922ce43607fc3154670471848d38156819388a6be1b72fe080"></a>

## tenant property — https.tls_cert_params.use_mtls.crl / c67fdf16de50 / 6

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

<a id="canonical-a1123bbd9f486571f316e39947ad7cca0958c800a521a97c019ef930b9334a2a"></a>

## Next pages — https.tls_cert_params.use_mtls.crl / c67fdf16de50 / 7

- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-109437c68a412d0b677ccd3f56d7ce7b57300844c679372dd4cffee7325e0f90)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e77b079a70be8330a0060021eb26d2a14dc71bbc055d92e0e2b36ec230477e4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f05f4fe3e17fbf72232312d4e87546abe5d181bdd79c65de11fd252f9cfe8c82"></a>

## https.tls_cert_params.use_mtls.no_crl — https.tls_cert_params.use_mtls.no_crl / a2ea40999f55 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-109437c68a412d0b677ccd3f56d7ce7b57300844c679372dd4cffee7325e0f90)
- https.tls_cert_params.use_mtls.no_crl

<a id="canonical-6ced433599c51731317291c026c8a6ca708a3042c20253e1a7779674584e530d"></a>

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

<a id="canonical-f895a1fc09d37e1dc61866e3c1c4ffdc439c1b059618f2f3058a29dd11f09b4e"></a>

## Direct properties — https.tls_cert_params.use_mtls.no_crl / a2ea40999f55 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-998e7329c835bffa8b073ac6f30c217b295fe722cfe3dfac8eb38b4b73e6da61"></a>

## Next pages — https.tls_cert_params.use_mtls.no_crl / a2ea40999f55 / 4

- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-109437c68a412d0b677ccd3f56d7ce7b57300844c679372dd4cffee7325e0f90)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-69e4bceb3760b19dd0ee81e996049025b2d4a7f3a0cd0f57e0d41ae07b5e7f9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b9ca8cfdbcc344007246c1cce75b86f01e8c8798fb962c346ab90151e669879"></a>

## https.tls_cert_params.use_mtls.trusted_ca — https.tls_cert_params.use_mtls.trusted_ca / f76c14b667f3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-109437c68a412d0b677ccd3f56d7ce7b57300844c679372dd4cffee7325e0f90)
- https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-1a25681655b1b7aa172873f86ee57f27b69242d9a61d725c06bb424bd5a56be1"></a>

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

<a id="canonical-c18ddc6d7e52adcdfb778893603bd50480642a889e27335648e759096d3a9248"></a>

## Direct properties — https.tls_cert_params.use_mtls.trusted_ca / f76c14b667f3 / 3

<a id="canonical-2ec6aca19d3f2888971510be1118a6f9d9a0f9c07381cc14d30a5ad6f82e40b6"></a>

<a id="canonical-1eb7c8eb6a5ce887e06ae975c44c0c3a960281e435547ffa86ec76429841c6ac"></a>

## name property — https.tls_cert_params.use_mtls.trusted_ca / f76c14b667f3 / 4

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

<a id="canonical-11f0accaa7d41f46ba03f46c4460302a8db95460b5fa44a9db2da55f31bc3229"></a>

<a id="canonical-b45048231ed2aaa16399e504d71da30337c2ac448a12400c9583af61fa26bfdd"></a>

## namespace property — https.tls_cert_params.use_mtls.trusted_ca / f76c14b667f3 / 5

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

<a id="canonical-eb7a8fdcdfa00d42bea0720879c8b23deb82d96bb3ab62f3e193d9249ae416bb"></a>

<a id="canonical-6eec5da1181956fbfb6bef4a14a9b0f927a35ed5057006a4e98a912c81502e91"></a>

## tenant property — https.tls_cert_params.use_mtls.trusted_ca / f76c14b667f3 / 6

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

<a id="canonical-ff8e3795a243daf2056d53411ddbc0dad090859aaa98cec89c5de1d5585d0f31"></a>

## Next pages — https.tls_cert_params.use_mtls.trusted_ca / f76c14b667f3 / 7

- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-109437c68a412d0b677ccd3f56d7ce7b57300844c679372dd4cffee7325e0f90)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-115fe7e085800b0d0e878431bf01379919a570c4e89663b84b3551a6576fb6c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5159289799b2cc6f255273a354ae587fc80c5ccca4f8ff1c94a5b48e16899301"></a>

## https.tls_cert_params.use_mtls.xfcc_disabled — https.tls_cert_params.use_mtls.xfcc_disabled / 4b475b71cf2b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-109437c68a412d0b677ccd3f56d7ce7b57300844c679372dd4cffee7325e0f90)
- https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-ba10d39df61dcf6612d6e614bcc0390c48bdf9a6104859c00ead7a572c1a85f2"></a>

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

<a id="canonical-3e808f367cff3ab09ff7ffc7c0d6a9eabc9be8969778d4632a3dc64653eb59fb"></a>

## Direct properties — https.tls_cert_params.use_mtls.xfcc_disabled / 4b475b71cf2b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-38519b898799ca93cacc1e1708eee110816c23f90ca0be55cbf34c291fb97564"></a>

## Next pages — https.tls_cert_params.use_mtls.xfcc_disabled / 4b475b71cf2b / 4

- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-109437c68a412d0b677ccd3f56d7ce7b57300844c679372dd4cffee7325e0f90)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c7630ff79e850d0c74a1ab8cf27647d01f33d29db0913277bad5bf6b1bb058ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eafe6546dc5a8131b2b41d5ec02c09b7c982b13aa164d197d3120810e57febe6"></a>

## https.tls_cert_params.use_mtls.xfcc_options — https.tls_cert_params.use_mtls.xfcc_options / e73f46e61938 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0564fbe8f589658128fd5e9aa78badc2e74eed6fa048cda70e62d557e1b23eb2)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-109437c68a412d0b677ccd3f56d7ce7b57300844c679372dd4cffee7325e0f90)
- https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-9c0671d6136e47ffd725d59826da29c3ec7e63b5fdca2d6705a844637254c879"></a>

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

<a id="canonical-6be9eb5871af02ffa3e539914f7723585ddbb5dfeeea108a373250222aaecac6"></a>

## Direct properties — https.tls_cert_params.use_mtls.xfcc_options / e73f46e61938 / 3

<a id="canonical-ffd2ddc9351d53838f850b0444292b118e3966b2a77a106550cdcdd1e6547d96"></a>

<a id="canonical-de2dc5fec5a946543c89c193089a3d22654f64b9353fd9097f3e966aa8794cc7"></a>

## xfcc_header_elements property — https.tls_cert_params.use_mtls.xfcc_options / e73f46e61938 / 4

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

<a id="canonical-91c15fedd837b69dac0ff393ed042549292a4e90a37a5d4fbe172643a0b27b97"></a>

## Next pages — https.tls_cert_params.use_mtls.xfcc_options / e73f46e61938 / 5

- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-109437c68a412d0b677ccd3f56d7ce7b57300844c679372dd4cffee7325e0f90)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e984e44a650f74877d240b26a61261f1dd62cdbb88b3d18215b2c964aff0dd2"></a>

## https.tls_parameters — https.tls_parameters / 1deb97bcf48a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- https.tls_parameters

<a id="canonical-08c5c48b32d7da409d638151d1ceb84f52bb0972c845a1186319dd895f97e273"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

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

<a id="canonical-e3ef325840075062ac58d4b1fcbfd921799661cdbe8c358e822ec13c222e3b7d"></a>

## Direct properties — https.tls_parameters / 1deb97bcf48a / 3

- [no_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-e1666d80784efb3f8bd6b0d712b8c1a8d1532b792a1c5a29dcc7efa7d12d993f): complete subsection reference.

- [tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-6f6067b6bb0b79822cbe13ac0c3570cc2b2829500c276631cba2407ec97752bb): complete subsection reference.

- [tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-46568524921ac4b0eb6f50b6f3b5fb9524433874a79df32f6881461c4431b66c): complete subsection reference.

- [use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-c5172c223c579dc3bc4e8925d406495ac392e82c5be6a24d35df6c8999afb1ae): complete subsection reference.

<a id="canonical-223649658ca75f348c6fe90d9dd1abb11749ec05522e30d13cd12b6b506942f7"></a>

## Next pages — https.tls_parameters / 1deb97bcf48a / 4

- [https.tls_parameters.no_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-e1666d80784efb3f8bd6b0d712b8c1a8d1532b792a1c5a29dcc7efa7d12d993f)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-6f6067b6bb0b79822cbe13ac0c3570cc2b2829500c276631cba2407ec97752bb)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-46568524921ac4b0eb6f50b6f3b5fb9524433874a79df32f6881461c4431b66c)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-c5172c223c579dc3bc4e8925d406495ac392e82c5be6a24d35df6c8999afb1ae)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e1666d80784efb3f8bd6b0d712b8c1a8d1532b792a1c5a29dcc7efa7d12d993f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad2608e3f8aac7f2b746d942a068ef469ecb9bc84a437e0fd7935afff45a0ab6"></a>

## https.tls_parameters.no_mtls — https.tls_parameters.no_mtls / 1d42a1a01113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- https.tls_parameters.no_mtls

<a id="canonical-5f1dec4f06a58e301d626b4a33405c858f8cad89c3148f8f0096954442fa003a"></a>

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

<a id="canonical-e7cf47670fd2b5eaf55382f8c350dbe74cef16948062d362b8908d6b1a27de7e"></a>

## Direct properties — https.tls_parameters.no_mtls / 1d42a1a01113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cf13fead9d7de611b152581d2557b347a3d1bc92eb644dd4e5cea7ae38fa9fd3"></a>

## Next pages — https.tls_parameters.no_mtls / 1d42a1a01113 / 4

- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6f6067b6bb0b79822cbe13ac0c3570cc2b2829500c276631cba2407ec97752bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b34bb295d01c0d3882c72b94b829cdc8965eed54a72006ed509512dc0ea2a048"></a>

## https.tls_parameters.tls_certificates — https.tls_parameters.tls_certificates / 68b939bae166 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- https.tls_parameters.tls_certificates

<a id="canonical-bdc2f6479cc18336bf55824ddcf188012f35e568da8e26135bd815348b5726f2"></a>

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

<a id="canonical-5bf174eb38c2c42d64e9c0b2a4566c617c0990fe439f06b86886d0e695e586bd"></a>

## Direct properties — https.tls_parameters.tls_certificates / 68b939bae166 / 3

<a id="canonical-39b9ed311e692cd885caf70c1fd106e1a55480cc2c3299211ef4d121a48cdc6c"></a>

<a id="canonical-96beb2b15abadad79af2de01689238960970afde0a6821dbe3f110aad15f359f"></a>

## certificate_url property — https.tls_parameters.tls_certificates / 68b939bae166 / 4

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

- [custom_hash_algorithms](data-sources--http_loadbalancer--reference--group-018.md#canonical-9cd0141f683a2eafcb470f048a9c9faf3d33a362334d6e0015b9631d60a15d6e): complete subsection reference.

<a id="canonical-ff8e1558936a0b260e71ac7d6635af4a1296f35861d3e611fa8a89472dcd94f4"></a>

<a id="canonical-611bdb2eb440c4e6a08eb73aa85b6b402b2e288c1029d1d9fcfeae3040c2379b"></a>

## description_spec property — https.tls_parameters.tls_certificates / 68b939bae166 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--http_loadbalancer--reference--group-018.md#canonical-d9c4a4e20892c9216f872ee8cd7c89fb264b69b544d6099117b66b48c1bfedea): complete subsection reference.

- [private_key](data-sources--http_loadbalancer--reference--group-018.md#canonical-0df7a56a4df90971bff110a6bcf80c70f8a00364e466a52457e63d028a0a2dc3): complete subsection reference.

- [use_system_defaults](data-sources--http_loadbalancer--reference--group-018.md#canonical-41882326889fa677ea61d68c91c0360c4b4104ae8b6e6d489e164058b68d7c04): complete subsection reference.

<a id="canonical-847a52100f453056d0d75d21b361e5a7ab466bcb8a7ddb6de88246cdf52d69b0"></a>

## Next pages — https.tls_parameters.tls_certificates / 68b939bae166 / 6

- [https.tls_parameters.tls_certificates.custom_hash_algorithms](data-sources--http_loadbalancer--reference--group-018.md#canonical-9cd0141f683a2eafcb470f048a9c9faf3d33a362334d6e0015b9631d60a15d6e)
- [https.tls_parameters.tls_certificates.disable_ocsp_stapling](data-sources--http_loadbalancer--reference--group-018.md#canonical-d9c4a4e20892c9216f872ee8cd7c89fb264b69b544d6099117b66b48c1bfedea)
- [https.tls_parameters.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-018.md#canonical-0df7a56a4df90971bff110a6bcf80c70f8a00364e466a52457e63d028a0a2dc3)
- [https.tls_parameters.tls_certificates.use_system_defaults](data-sources--http_loadbalancer--reference--group-018.md#canonical-41882326889fa677ea61d68c91c0360c4b4104ae8b6e6d489e164058b68d7c04)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9cd0141f683a2eafcb470f048a9c9faf3d33a362334d6e0015b9631d60a15d6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a468d41c6dc25a0d7efd5f611daa17a6cee6e06f5641112c84955cbb87cad73b"></a>

## https.tls_parameters.tls_certificates.custom_hash_algorithms — https.tls_parameters.tls_certificates.custom_hash_algorithms / 930456009b05 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-6f6067b6bb0b79822cbe13ac0c3570cc2b2829500c276631cba2407ec97752bb)
- https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-b11b9edb59ff6a5ba10e1a23eb66f346126172cbf646febe1069d45e8815de26"></a>

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

<a id="canonical-3170d15cbd9d8e1b86d92169a289118f22f605aa603480cc55350b4fe89eb679"></a>

## Direct properties — https.tls_parameters.tls_certificates.custom_hash_algorithms / 930456009b05 / 3

<a id="canonical-c95d06b20072421481b35bef8dce46bb3c5de7d99fc831bae06de42342416634"></a>

<a id="canonical-78a1c14181184f5b3731e50882c32390ba1b52fbf1433c17bc4bc3f9994e0ff4"></a>

## hash_algorithms property — https.tls_parameters.tls_certificates.custom_hash_algorithms / 930456009b05 / 4

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

<a id="canonical-2c48ebf36c17afcc9c85f7d6afb39454639824222318e0161b526d754d2196a2"></a>

## Next pages — https.tls_parameters.tls_certificates.custom_hash_algorithms / 930456009b05 / 5

- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-6f6067b6bb0b79822cbe13ac0c3570cc2b2829500c276631cba2407ec97752bb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d9c4a4e20892c9216f872ee8cd7c89fb264b69b544d6099117b66b48c1bfedea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25fa560fef482f61b03b1a1fa24e175a2db014245cda36d8a0eea8d3f06060fb"></a>

## https.tls_parameters.tls_certificates.disable_ocsp_stapling — https.tls_parameters.tls_certificates.disable_ocsp_stapling / ce6e803293c6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-6f6067b6bb0b79822cbe13ac0c3570cc2b2829500c276631cba2407ec97752bb)
- https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-b3e608816cce2ff997425ae3d644aa0f0af6149c4d198f2f45e56de76bc1784f"></a>

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

<a id="canonical-e142a685134d5dfd0f648f72af95f3703d1716e57cd3f4b1b0d2cfb7aac76a8a"></a>

## Direct properties — https.tls_parameters.tls_certificates.disable_ocsp_stapling / ce6e803293c6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-15fc271a0031e7b15fc07c7a0d5c6e3a195cdd9db81e03461df7001fc4f00ad3"></a>

## Next pages — https.tls_parameters.tls_certificates.disable_ocsp_stapling / ce6e803293c6 / 4

- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-6f6067b6bb0b79822cbe13ac0c3570cc2b2829500c276631cba2407ec97752bb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0df7a56a4df90971bff110a6bcf80c70f8a00364e466a52457e63d028a0a2dc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd7b7b3d52f75cbc13b8c3c7819e1c1cedca4a35d325861a1881f389ab217578"></a>

## https.tls_parameters.tls_certificates.private_key — https.tls_parameters.tls_certificates.private_key / 6897afd379e6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-6f6067b6bb0b79822cbe13ac0c3570cc2b2829500c276631cba2407ec97752bb)
- https.tls_parameters.tls_certificates.private_key

<a id="canonical-b705e26c988618b823be51d2ef4b1274a5ce2afcb841a1e951b2aa56b36d2688"></a>

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

<a id="canonical-fcb126cbc9b72997f9a424c7cefde34f3f0cd87a832cf7bd4a02c2788e292e61"></a>

## Direct properties — https.tls_parameters.tls_certificates.private_key / 6897afd379e6 / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-018.md#canonical-11c65e5d4745d2faa1e0f9a1c77c693a13d8655ebb99e77085c69e285cc3eaec): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-018.md#canonical-e0a34b721b3088439258594b4f7b4238b093085130db4bfca8cabf99adb02bb0): complete subsection reference.

<a id="canonical-601fb3952c8a8248f86d299a168d19008223871e89c8d813168083d84ff0b849"></a>

## Next pages — https.tls_parameters.tls_certificates.private_key / 6897afd379e6 / 4

- [https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-018.md#canonical-11c65e5d4745d2faa1e0f9a1c77c693a13d8655ebb99e77085c69e285cc3eaec)
- [https.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--http_loadbalancer--reference--group-018.md#canonical-e0a34b721b3088439258594b4f7b4238b093085130db4bfca8cabf99adb02bb0)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-6f6067b6bb0b79822cbe13ac0c3570cc2b2829500c276631cba2407ec97752bb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-11c65e5d4745d2faa1e0f9a1c77c693a13d8655ebb99e77085c69e285cc3eaec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2017a2190d6aa3ef7a8f4884a79a987d04ade34932e8d8c5becc64df8b18c82"></a>

## https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — https.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 7f3e6c671cef / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-6f6067b6bb0b79822cbe13ac0c3570cc2b2829500c276631cba2407ec97752bb)
- [https.tls_parameters.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-018.md#canonical-0df7a56a4df90971bff110a6bcf80c70f8a00364e466a52457e63d028a0a2dc3)
- https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-22ac4de36325180489d8906d07e19c910605690c9a05eb739b013df2ffe19064"></a>

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

<a id="canonical-bb04179bcb95a8b3c7feeb5ca1739303f7ee346f629398da268993209d618973"></a>

## Direct properties — https.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 7f3e6c671cef / 3

<a id="canonical-74c67c99c114983325366abd41fdf0413cb6ce98094802c1aee593c78436ac3f"></a>

<a id="canonical-b713ca8e0c94d9d5d5629c62e56c57b6e5b282553af25d5b98f686be7fcdd732"></a>

## decryption_provider property — https.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 7f3e6c671cef / 4

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

<a id="canonical-4472c0d467cb7230dcc6f8c26b519ff7481fd985050b8ac9bd8fac8f7bc4a0c1"></a>

<a id="canonical-e8643b56365ce856c3b733d6aef5c27bbd30de10489755dc34197001ce6671ea"></a>

## location property — https.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 7f3e6c671cef / 5

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

<a id="canonical-da04ea3bafd8f9da5e32ebbe8ce6b3165f6509a92ea422735d13427607280db7"></a>

<a id="canonical-8611cf6583342a6489e9a52c4117d1bddd48dd85094a952c352a7f96f01bc0f9"></a>

## store_provider property — https.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 7f3e6c671cef / 6

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

<a id="canonical-71ce6e520535ccc5e3836d50721ec8897e47b6946555a2c436e7745cc2e61676"></a>

## Next pages — https.tls_parameters.tls_certificates.private_key.blindfold_secret_info / 7f3e6c671cef / 7

- [https.tls_parameters.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-018.md#canonical-0df7a56a4df90971bff110a6bcf80c70f8a00364e466a52457e63d028a0a2dc3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e0a34b721b3088439258594b4f7b4238b093085130db4bfca8cabf99adb02bb0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd8a19eb3665a488285ad2aca776e9ae9c15212ab80e21c5fca747863dda5408"></a>

## https.tls_parameters.tls_certificates.private_key.clear_secret_info — https.tls_parameters.tls_certificates.private_key.clear_secret_info / 1647d609deba / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-6f6067b6bb0b79822cbe13ac0c3570cc2b2829500c276631cba2407ec97752bb)
- [https.tls_parameters.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-018.md#canonical-0df7a56a4df90971bff110a6bcf80c70f8a00364e466a52457e63d028a0a2dc3)
- https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-c3001126020aac71a5d41b8b50dee02050851919847e2345a55223dd331670a3"></a>

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

<a id="canonical-202122e669a8fa29b66d0ecb6942032a27147c3879651c83742502e79bc13428"></a>

## Direct properties — https.tls_parameters.tls_certificates.private_key.clear_secret_info / 1647d609deba / 3

<a id="canonical-17ddb1e0c4313c93fcee72b97b74e3a1b8261a4c79e55635687c44730d7d3ed6"></a>

<a id="canonical-eb755c55e53d39adef15816938aeedf9bb6089038550cd9f749d85fa9670eadb"></a>

## provider_ref property — https.tls_parameters.tls_certificates.private_key.clear_secret_info / 1647d609deba / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-95ee8df5bb6455d88061157ebb675b3cbd147f8e8315390dac0abe9c170010a9"></a>

<a id="canonical-c1828459d0e81e1bdf419f5fa2364b6d4a785a7ca7b3c676ae8fdc0d7ba6bd49"></a>

## url property — https.tls_parameters.tls_certificates.private_key.clear_secret_info / 1647d609deba / 5

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

<a id="canonical-ce253f268c4a5e08e24a4e99b6f862ca1ff38b24c66ac01611970841a8854d49"></a>

## Next pages — https.tls_parameters.tls_certificates.private_key.clear_secret_info / 1647d609deba / 6

- [https.tls_parameters.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-018.md#canonical-0df7a56a4df90971bff110a6bcf80c70f8a00364e466a52457e63d028a0a2dc3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-41882326889fa677ea61d68c91c0360c4b4104ae8b6e6d489e164058b68d7c04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36686cd9229827add0b0f49ac599d6d6c920c94b703ca8efd5263bbde528fc75"></a>

## https.tls_parameters.tls_certificates.use_system_defaults — https.tls_parameters.tls_certificates.use_system_defaults / 856a14590fd7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-6f6067b6bb0b79822cbe13ac0c3570cc2b2829500c276631cba2407ec97752bb)
- https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-98aa436bc9b5bdbf21f3de53a70ebcc3d063ad29ee771dc65312ab60ef1de7d1"></a>

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

<a id="canonical-5b34d88caef32141987deb6ce39b1449ab2a4c3a260a538acfae6bc19c441568"></a>

## Direct properties — https.tls_parameters.tls_certificates.use_system_defaults / 856a14590fd7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3fce64ccf351469ec62fdcc5cf3a632ff556fd40abee97d5f7ab470865303532"></a>

## Next pages — https.tls_parameters.tls_certificates.use_system_defaults / 856a14590fd7 / 4

- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-6f6067b6bb0b79822cbe13ac0c3570cc2b2829500c276631cba2407ec97752bb)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-46568524921ac4b0eb6f50b6f3b5fb9524433874a79df32f6881461c4431b66c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e01712798c6e37055ec8c9e81a62d6f6f38c48c0be12089a085a322cbac3259a"></a>

## https.tls_parameters.tls_config — https.tls_parameters.tls_config / caaa2b994e12 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- https.tls_parameters.tls_config

<a id="canonical-a81775a096bbd6975da4e0a997bc53b9e87cb7e9e7b3522e0c472e337fbb2f57"></a>

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

<a id="canonical-b0a1ef756646390f86ac59dc084d8ac91dcf7efe10e65df48652df7fca686ea9"></a>

## Direct properties — https.tls_parameters.tls_config / caaa2b994e12 / 3

- [custom_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-185c37656ae2ea3953a8f36b6f1d39d019d5518667f510419150d07e3a1908db): complete subsection reference.

- [default_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-71a2a72838f960e3c1e6f9b0228de19fbf7a5b1cb781b376a59b2ae9544e428b): complete subsection reference.

- [low_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-90b6b2e7a705e4497ec536aed1ba2257abce08f9bb4b180460d43cd05e9de168): complete subsection reference.

- [medium_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-6d75871836e0cca6b304b387fbf0a924c667f8fa208026d8e807eb83a3c669b5): complete subsection reference.

<a id="canonical-5afcf67fc8abe65a0dc7a250b6fe7b75c1eede3a54d5bd45e8f06292d3dc09be"></a>

## Next pages — https.tls_parameters.tls_config / caaa2b994e12 / 4

- [https.tls_parameters.tls_config.custom_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-185c37656ae2ea3953a8f36b6f1d39d019d5518667f510419150d07e3a1908db)
- [https.tls_parameters.tls_config.default_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-71a2a72838f960e3c1e6f9b0228de19fbf7a5b1cb781b376a59b2ae9544e428b)
- [https.tls_parameters.tls_config.low_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-90b6b2e7a705e4497ec536aed1ba2257abce08f9bb4b180460d43cd05e9de168)
- [https.tls_parameters.tls_config.medium_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-6d75871836e0cca6b304b387fbf0a924c667f8fa208026d8e807eb83a3c669b5)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-185c37656ae2ea3953a8f36b6f1d39d019d5518667f510419150d07e3a1908db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d77d485350b7846cc41d21f390eaa9105069d3b405dbf6f7d147441b5a4abcd"></a>

## https.tls_parameters.tls_config.custom_security — https.tls_parameters.tls_config.custom_security / 7af8e67e16b8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-46568524921ac4b0eb6f50b6f3b5fb9524433874a79df32f6881461c4431b66c)
- https.tls_parameters.tls_config.custom_security

<a id="canonical-db5f669bbbc955ddcbe356638d0314aa99a696bd0b12f8461c1ad4579ae72ce2"></a>

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

<a id="canonical-4d1c40c2d82150937f83947958d4ba657d204956c9a8380919ce6c9c693b610d"></a>

## Direct properties — https.tls_parameters.tls_config.custom_security / 7af8e67e16b8 / 3

<a id="canonical-8ab371ea3c3ed66befc01a66d7d85c563a3678d72146539dfa1db5f16d88d908"></a>

<a id="canonical-c82d3e058ffc3bdec80b4506d12b847caabea445d88e9211c41af149841e8db4"></a>

## cipher_suites property — https.tls_parameters.tls_config.custom_security / 7af8e67e16b8 / 4

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

<a id="canonical-04e71858b757bc6657ea61de8357f6c1321f31256d42fa1176d5724550fbb98e"></a>

<a id="canonical-5baac22866ceb34c6308b1aceb858785f52daae134599628c6c6dc4ab253d1c2"></a>

## max_version property — https.tls_parameters.tls_config.custom_security / 7af8e67e16b8 / 5

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

<a id="canonical-739fccb6ad218b102790318263aca3c0adcfc1b05cab57fc723e169a4a99ce29"></a>

<a id="canonical-5f98e414c59ef08c3b3736db172293efbe04765c4d28a8caf2c959d3ddd9cf2b"></a>

## min_version property — https.tls_parameters.tls_config.custom_security / 7af8e67e16b8 / 6

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

<a id="canonical-66c0c7531d0221fc0103fd48d8d543c018fecb74d64ee9a74c7b4176d0d502c1"></a>

## Next pages — https.tls_parameters.tls_config.custom_security / 7af8e67e16b8 / 7

- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-46568524921ac4b0eb6f50b6f3b5fb9524433874a79df32f6881461c4431b66c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-71a2a72838f960e3c1e6f9b0228de19fbf7a5b1cb781b376a59b2ae9544e428b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d52201783e15117fddf055ad53c398f18c02695a3112918762c8559e6ec08bb"></a>

## https.tls_parameters.tls_config.default_security — https.tls_parameters.tls_config.default_security / d000d9d3f7b9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-46568524921ac4b0eb6f50b6f3b5fb9524433874a79df32f6881461c4431b66c)
- https.tls_parameters.tls_config.default_security

<a id="canonical-4fe2041193d5000eda33a126cc87ea72dbc7b25213e251b769c5db25d4578a4d"></a>

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

<a id="canonical-f1b59c0a02c61138e9629eef826219febf76866a04b961993502f939fef262e8"></a>

## Direct properties — https.tls_parameters.tls_config.default_security / d000d9d3f7b9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5aff2659a30b614b7e595ad08dd72016d01e8339f32693f409c33a439cc5a1aa"></a>

## Next pages — https.tls_parameters.tls_config.default_security / d000d9d3f7b9 / 4

- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-46568524921ac4b0eb6f50b6f3b5fb9524433874a79df32f6881461c4431b66c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-90b6b2e7a705e4497ec536aed1ba2257abce08f9bb4b180460d43cd05e9de168"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93a6d8aa68dd50bd7899b579bec1d1204ea2a7c2a7a29b40b4625623d2d5cef1"></a>

## https.tls_parameters.tls_config.low_security — https.tls_parameters.tls_config.low_security / b716ddb212ef / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-46568524921ac4b0eb6f50b6f3b5fb9524433874a79df32f6881461c4431b66c)
- https.tls_parameters.tls_config.low_security

<a id="canonical-94d7e6d3aae895c20659524e9f49d6f2a40926cef976e7dd0171d663c1f2c00e"></a>

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

<a id="canonical-dcb315156b22af33bdde65ee8f0c72051d4702c511bf17332dfd34f0a95f4280"></a>

## Direct properties — https.tls_parameters.tls_config.low_security / b716ddb212ef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-de916b6432e99a22a92e29e3be966d5d62b697af0166da8923ee67ed017bd396"></a>

## Next pages — https.tls_parameters.tls_config.low_security / b716ddb212ef / 4

- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-46568524921ac4b0eb6f50b6f3b5fb9524433874a79df32f6881461c4431b66c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6d75871836e0cca6b304b387fbf0a924c667f8fa208026d8e807eb83a3c669b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7015a9c4245c111ed98c111560488d552fa1d535b158e27ae9ae4d8321e45520"></a>

## https.tls_parameters.tls_config.medium_security — https.tls_parameters.tls_config.medium_security / 96ad1ec59aad / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-46568524921ac4b0eb6f50b6f3b5fb9524433874a79df32f6881461c4431b66c)
- https.tls_parameters.tls_config.medium_security

<a id="canonical-a6db1d4d13545288c0810b0ea818298663750a4ed910c68a35140ab4a38e521f"></a>

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

<a id="canonical-20044c4063edfbd0bfea239e72eea8eea7e4a40308bf9e5e51fee2b590c8af27"></a>

## Direct properties — https.tls_parameters.tls_config.medium_security / 96ad1ec59aad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-17945043e2bc481e8cf40ea31908bb94b6b8c25b05126833e97301b6bf97f21c"></a>

## Next pages — https.tls_parameters.tls_config.medium_security / 96ad1ec59aad / 4

- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-46568524921ac4b0eb6f50b6f3b5fb9524433874a79df32f6881461c4431b66c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c5172c223c579dc3bc4e8925d406495ac392e82c5be6a24d35df6c8999afb1ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e0a6ef9417f91b562658c3c0cd757496da540f3945d028345fff2d8ae9db3b7"></a>

## https.tls_parameters.use_mtls — https.tls_parameters.use_mtls / 63dbb853ca1f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- https.tls_parameters.use_mtls

<a id="canonical-fa53e2a0ec5559f8a814d2f148d85e6bf9b1a91b30ff3bd43c78f4cac979dd9d"></a>

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

<a id="canonical-8e849eff607a85c0d68108861d2181908dc0f9389fc8e8a61f639fcab1cbca19"></a>

## Direct properties — https.tls_parameters.use_mtls / 63dbb853ca1f / 3

<a id="canonical-9f4990c1053a86a8597a816b3ce753c2582c6823c1c99e531e0fb1375ae280be"></a>

<a id="canonical-759f023b34c7bce7e4ae1c92d575642d34d5a0264239c407677baf77b7195cb8"></a>

## client_certificate_optional property — https.tls_parameters.use_mtls / 63dbb853ca1f / 4

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

- [crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-988345739b0be3ab7829335d9dded24febcc3727be80ac00f0c43588d903b87a): complete subsection reference.

- [no_crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-f85f5feb3dd61ed5ecd1268ed4a4a96ee97e1378e9dc59a0172f1b9727b370c3): complete subsection reference.

- [trusted_ca](data-sources--http_loadbalancer--reference--group-018.md#canonical-dc32a19ae0061d47b8be0ae78b208cfed03cc4f6e97ca783f533c625fca3bb0d): complete subsection reference.

<a id="canonical-110c22c2ef629f18a00821f253d2423630e7459f257327fb45bcecad9d8dae43"></a>

<a id="canonical-b3012d1ec465064332d30a1280c1b4fb79f0758ccd297a21636218a90dbb35df"></a>

## trusted_ca_url property — https.tls_parameters.use_mtls / 63dbb853ca1f / 5

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

- [xfcc_disabled](data-sources--http_loadbalancer--reference--group-018.md#canonical-258614468e3065bf4bda7a0ed2fb08f1253c8061cd7c1d59cf0244dca4d12241): complete subsection reference.

- [xfcc_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-7a2fba13185563eef141a4404a2d2333ffc26f23a9a89b9b8e620e97891993d2): complete subsection reference.

<a id="canonical-cbf1b77723496cbfce2f6732c6fbfc72956f27628c10d1e6b7b70cd9e2cb4e5e"></a>

## Next pages — https.tls_parameters.use_mtls / 63dbb853ca1f / 6

- [https.tls_parameters.use_mtls.crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-988345739b0be3ab7829335d9dded24febcc3727be80ac00f0c43588d903b87a)
- [https.tls_parameters.use_mtls.no_crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-f85f5feb3dd61ed5ecd1268ed4a4a96ee97e1378e9dc59a0172f1b9727b370c3)
- [https.tls_parameters.use_mtls.trusted_ca](data-sources--http_loadbalancer--reference--group-018.md#canonical-dc32a19ae0061d47b8be0ae78b208cfed03cc4f6e97ca783f533c625fca3bb0d)
- [https.tls_parameters.use_mtls.xfcc_disabled](data-sources--http_loadbalancer--reference--group-018.md#canonical-258614468e3065bf4bda7a0ed2fb08f1253c8061cd7c1d59cf0244dca4d12241)
- [https.tls_parameters.use_mtls.xfcc_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-7a2fba13185563eef141a4404a2d2333ffc26f23a9a89b9b8e620e97891993d2)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-988345739b0be3ab7829335d9dded24febcc3727be80ac00f0c43588d903b87a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2861d44c25ab45b1cbbc628c6c165cbebf97d3b2f072d2e4766fa0b31dbb8efc"></a>

## https.tls_parameters.use_mtls.crl — https.tls_parameters.use_mtls.crl / 915a1384f9f8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-c5172c223c579dc3bc4e8925d406495ac392e82c5be6a24d35df6c8999afb1ae)
- https.tls_parameters.use_mtls.crl

<a id="canonical-7e203dad69afc1c38d8379aa57aa38056b68ef4d7969206f8663a63945bbbf1e"></a>

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

<a id="canonical-dc093dbadd79c933c89bd840f1e9a324652e3fd6c3b26f1a754887f6d57a9f10"></a>

## Direct properties — https.tls_parameters.use_mtls.crl / 915a1384f9f8 / 3

<a id="canonical-fd089cb13ceb4865758b3d871fa30897dca3074dbda4e3b52a78e239042f125c"></a>

<a id="canonical-114dfe072f338b5ab416890df00be7b27de68adb8cadd47eb0c510f1c44ba7f4"></a>

## name property — https.tls_parameters.use_mtls.crl / 915a1384f9f8 / 4

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

<a id="canonical-33dad67014a2f538bcb2732d1a44af7cd9045c3050f400a666548205cb2f28e0"></a>

<a id="canonical-b15dd9139454ecbf7ebd508264fea5d470e2528b6e1c445a377d92f00a20dfcd"></a>

## namespace property — https.tls_parameters.use_mtls.crl / 915a1384f9f8 / 5

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

<a id="canonical-87f6b6631e8c8901246a16b7bab2bb0c41c41d1ed070cc8764d8bf5992b1d29f"></a>

<a id="canonical-3375adb57b2d00eac6774faaa406e2b37d1e1179cff423c231285a59e7816632"></a>

## tenant property — https.tls_parameters.use_mtls.crl / 915a1384f9f8 / 6

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

<a id="canonical-191ab0b91593bf794b188e3cf3a1103513735f3f122072806179a706b9c7633c"></a>

## Next pages — https.tls_parameters.use_mtls.crl / 915a1384f9f8 / 7

- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-c5172c223c579dc3bc4e8925d406495ac392e82c5be6a24d35df6c8999afb1ae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f85f5feb3dd61ed5ecd1268ed4a4a96ee97e1378e9dc59a0172f1b9727b370c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1c0e8dfc029a785b7fc4416619c8a6e2ed308f65eb956b9c044a890eb5ce506"></a>

## https.tls_parameters.use_mtls.no_crl — https.tls_parameters.use_mtls.no_crl / e19f75502796 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-c5172c223c579dc3bc4e8925d406495ac392e82c5be6a24d35df6c8999afb1ae)
- https.tls_parameters.use_mtls.no_crl

<a id="canonical-d4ad07348e20357cc8725de0369670f56ed79c966a68e08472c307e12f56c2e7"></a>

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

<a id="canonical-b9c66dada5eae0d5f2efca1c705f3bbc754e06e3d8c1dcd5c677648de72ccd3a"></a>

## Direct properties — https.tls_parameters.use_mtls.no_crl / e19f75502796 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64f56f555e23c0253c4c433e9fab30c7ae838f0228d321ff1eb755b158c786eb"></a>

## Next pages — https.tls_parameters.use_mtls.no_crl / e19f75502796 / 4

- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-c5172c223c579dc3bc4e8925d406495ac392e82c5be6a24d35df6c8999afb1ae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-dc32a19ae0061d47b8be0ae78b208cfed03cc4f6e97ca783f533c625fca3bb0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4effbcb597eb3efc1c57ffe0b478009bfc5dd38c4623e836061c70b8e31e893"></a>

## https.tls_parameters.use_mtls.trusted_ca — https.tls_parameters.use_mtls.trusted_ca / 3477a57f41db / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-c5172c223c579dc3bc4e8925d406495ac392e82c5be6a24d35df6c8999afb1ae)
- https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-16814366a4011b0195ee8aeaf193eaa7615fc8708daad22d07a383240ea3372f"></a>

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

<a id="canonical-18977d8e99982f4c078949e6dc2ee14d676367c8a9d2ff9979962efd9b5175d1"></a>

## Direct properties — https.tls_parameters.use_mtls.trusted_ca / 3477a57f41db / 3

<a id="canonical-129c992367cee059f423ddb84534b0c6174ba7517885b6034d505928b7972f96"></a>

<a id="canonical-c274a0fffd99ae2fefafe3f5a26a31d392ec98d1b4d3d387cdc4f95bf7472395"></a>

## name property — https.tls_parameters.use_mtls.trusted_ca / 3477a57f41db / 4

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

<a id="canonical-5d45d632899644e0c7d0f92e607eca8e03aa8f5b6e6824f7cb8fde9e1fcb6d67"></a>

<a id="canonical-af79d9eb0e63472efb08465ad0622f51a510a65600aef98c2ae6fd6127b8bf19"></a>

## namespace property — https.tls_parameters.use_mtls.trusted_ca / 3477a57f41db / 5

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

<a id="canonical-ef148c397ecff902b72e36af9b498850a616bd274565950d8aeadc76c0c03f55"></a>

<a id="canonical-cda1699feb14eee57aa9adc80d9c0f5642d07368a263c22f2aec2121ad531fc9"></a>

## tenant property — https.tls_parameters.use_mtls.trusted_ca / 3477a57f41db / 6

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

<a id="canonical-3460c1239776ab161f4853b12897ed9589e6963950bf9a6657b7b4002ddc42e7"></a>

## Next pages — https.tls_parameters.use_mtls.trusted_ca / 3477a57f41db / 7

- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-c5172c223c579dc3bc4e8925d406495ac392e82c5be6a24d35df6c8999afb1ae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-258614468e3065bf4bda7a0ed2fb08f1253c8061cd7c1d59cf0244dca4d12241"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c9a0032428e5c428f33faf9062077b849aa49202ad4e214209f1b93395089f2"></a>

## https.tls_parameters.use_mtls.xfcc_disabled — https.tls_parameters.use_mtls.xfcc_disabled / cbe2f2236e58 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-c5172c223c579dc3bc4e8925d406495ac392e82c5be6a24d35df6c8999afb1ae)
- https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-7d3b12c1efa5d8a7c78c3bd0d0d2f0025f954decb37a1cae3906f46fe3cc193e"></a>

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

<a id="canonical-b9ee1a927e9c8476a3df8d8f44be4605703e623893c95d6b0eab6a980d030477"></a>

## Direct properties — https.tls_parameters.use_mtls.xfcc_disabled / cbe2f2236e58 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-417cb3339419ede86ac73f5c393bfc76ccafb3146d47d07d673a061de56b5741"></a>

## Next pages — https.tls_parameters.use_mtls.xfcc_disabled / cbe2f2236e58 / 4

- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-c5172c223c579dc3bc4e8925d406495ac392e82c5be6a24d35df6c8999afb1ae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7a2fba13185563eef141a4404a2d2333ffc26f23a9a89b9b8e620e97891993d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2c278b025bdbc9bb50c864c5045641c1d75a7c5cddbea35f1eff36f172dab91"></a>

## https.tls_parameters.use_mtls.xfcc_options — https.tls_parameters.use_mtls.xfcc_options / fbd50a680af2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https](data-sources--http_loadbalancer--reference--group-017.md#canonical-16acff651d12ae362819957fbb0eb6204f1a67c04a94bd3664e8cfd0778b6171)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-ab56f2cd52f56425a70b8af47932489211a13b484a76e043dc9d820d5717240d)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-c5172c223c579dc3bc4e8925d406495ac392e82c5be6a24d35df6c8999afb1ae)
- https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-db62e917588413b8958b0b756f1735d70f59571e934bbc5099f859c3294fc68c"></a>

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

<a id="canonical-68801de960b81bee7857b1e294e07095908c7fc67f239483b41ee9d0f88ad3fd"></a>

## Direct properties — https.tls_parameters.use_mtls.xfcc_options / fbd50a680af2 / 3

<a id="canonical-615773f3b146a7cd5d6de69f4ef84b3afc0a72963b67ee16758b70c87bbbb14b"></a>

<a id="canonical-d7e06b754197d8423378ec81d075b57cc58f9f311108082b1d40a1af17154c3a"></a>

## xfcc_header_elements property — https.tls_parameters.use_mtls.xfcc_options / fbd50a680af2 / 4

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

<a id="canonical-903c9bed0af44cbf027670eecbbd83ab450958583cf76ce229f62b2d79674cb6"></a>

## Next pages — https.tls_parameters.use_mtls.xfcc_options / fbd50a680af2 / 5

- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-c5172c223c579dc3bc4e8925d406495ac392e82c5be6a24d35df6c8999afb1ae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e2ff5feee3a1ae293158939a97729795c3d8a7d9b4042a95c2f12867960d7cdd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c2d35168570086c2f7a08101d866090427edd2f3cca0fecd1f886c6843f54a3"></a>

## https_auto_cert — https_auto_cert / 9a86dc860a37 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- https_auto_cert

<a id="canonical-94bcef80ac42c4b17d5d9e65955f67b3339f56bd0be651f841e5fe2c67636962"></a>

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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]"
}
```

<a id="canonical-bb0835dc675b9c0f5a32c82368df869bfc2b49801919104aa267964d3c3b7430"></a>

## Direct properties — https_auto_cert / 9a86dc860a37 / 3

<a id="canonical-35ccff33f4b32e55f4475d712fc298d7235f25c33c539f122aaf3cb1935f2b75"></a>

<a id="canonical-d1faa041cb869cd6f1e3dfcc78d32a31b7c13117af817ad2dce185201fa81215"></a>

## add_hsts property — https_auto_cert / 9a86dc860a37 / 4

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header. Defaults to \`false\`. Server applies default
when omitted.

Upstream description:

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

<a id="canonical-3deb42273b68729a9e175117fffb540338a7c47f54f5bfb6c9be4872ac96fe1a"></a>

<a id="canonical-e49c6035bf5301329406cb55daf059034f8f1a0cafde4e5b4864daf5f4f697db"></a>

## append_server_name property — https_auto_cert / 9a86dc860a37 / 5

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

- [coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-fc24498d21cafcec765385dad56e1ac1ccb981e91d3873fee067e48bd77bc3dc): complete subsection reference.

<a id="canonical-4a6401c6a46a1b113bb4e840d6a2a87849da7f017669e5d958261994cb1ec09b"></a>

<a id="canonical-573cb565a2e463935b0b1f3a02e1897cea5c04653f727225eea136dfd70bc654"></a>

## connection_idle_timeout property — https_auto_cert / 9a86dc860a37 / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Server
applies default when omitted.

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

- [default_header](data-sources--http_loadbalancer--reference--group-018.md#canonical-f784dc759277c3a4f7b03f5524656d83772059bf09a0bae3adb8e5df6eab1aa0): complete subsection reference.

- [default_loadbalancer](data-sources--http_loadbalancer--reference--group-019.md#canonical-5370dd835d8b8e15da3b3fb81f1b50bf18ecec3524f51e241588f1beb9fb3969): complete subsection reference.

- [disable_path_normalize](data-sources--http_loadbalancer--reference--group-019.md#canonical-ac3b635fc7541a5cfe0dcc5e0fd11c3abc98224fbd690a0d7792c714cff1ce21): complete subsection reference.

- [enable_path_normalize](data-sources--http_loadbalancer--reference--group-019.md#canonical-289da877ce1d9348c55378310a94862dde2014ed2bdc312ec5de318b410b6ff7): complete subsection reference.

- [http_protocol_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-2ba73566b5d2f8460ac6670d0936192f46e26ad25347fecf4549bd80759abb11): complete subsection reference.

<a id="canonical-b01f4f11ad5b6f14a86c93d3594c67d00a12ab1b44b9f73b668def1e934d8fb0"></a>

<a id="canonical-6c479b516ce1240b82db9191ae864db6d9003e77dba2953f5495ebcafbb5f7cb"></a>

## http_redirect property — https_auto_cert / 9a86dc860a37 / 7

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS. Defaults to \`false\`. Server applies
default when omitted.

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

- [no_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-92c674b0dc9736d465350a71117143b3efa33055ffd5823ac08b0bfda705e14f): complete subsection reference.

- [non_default_loadbalancer](data-sources--http_loadbalancer--reference--group-019.md#canonical-f5e31eb273a07c8f3e56020b510a838abdc255f11cade33b6c97b437e448857d): complete subsection reference.

- [pass_through](data-sources--http_loadbalancer--reference--group-019.md#canonical-a98e6a4ab6e5f5e26acc643f9b4ccded462734d7eec0f5e940b573536276a97e): complete subsection reference.

<a id="canonical-6b781dcb5e71c3c7ab5d2e076faac303d43d23f7b7006df4f6c38507f1317ba3"></a>

<a id="canonical-f678ccdbde900c57d3f78df03c30af73e70d8dc0c3382f4b4dcfe2f206e8bdf0"></a>

## port property — https_auto_cert / 9a86dc860a37 / 8

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

<a id="canonical-a93f96bdb879ca5d93f49619aecbc927d46013e2f2a112df8502597fed30d283"></a>

<a id="canonical-47d29ca3635acd273c4b3f37b351218d71e33094b7f770d982fabc690a86e8a8"></a>

## port_ranges property — https_auto_cert / 9a86dc860a37 / 9

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

<a id="canonical-efd63135c8745d8d79ec1e920c2d69fbcfac984110cd4d3899d01537c7fd9851"></a>

<a id="canonical-f86b2af835993ec093cf22bbebee1f02f7b1a7e3ae69b264d7f7ecc80a6afdcc"></a>

## server_name property — https_auto_cert / 9a86dc860a37 / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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

- [tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-d063c10b8a5e2a7c99b5c5167ef68b0739078a5c4ea93e8ac7e0fef7efb4edde): complete subsection reference.

- [use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-d6cc7070a4da0f1e6653eedda19de11f6b8ce9f24ac85b7160b2085bf9ca4f7c): complete subsection reference.

<a id="canonical-16f46a2ac2cb8aaa88c8621b4c2fe05fe22679369f55a7bf2056e48bbc980b42"></a>

## Next pages — https_auto_cert / 9a86dc860a37 / 11

- [https_auto_cert.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-fc24498d21cafcec765385dad56e1ac1ccb981e91d3873fee067e48bd77bc3dc)
- [https_auto_cert.default_header](data-sources--http_loadbalancer--reference--group-018.md#canonical-f784dc759277c3a4f7b03f5524656d83772059bf09a0bae3adb8e5df6eab1aa0)
- [https_auto_cert.default_loadbalancer](data-sources--http_loadbalancer--reference--group-019.md#canonical-5370dd835d8b8e15da3b3fb81f1b50bf18ecec3524f51e241588f1beb9fb3969)
- [https_auto_cert.disable_path_normalize](data-sources--http_loadbalancer--reference--group-019.md#canonical-ac3b635fc7541a5cfe0dcc5e0fd11c3abc98224fbd690a0d7792c714cff1ce21)
- [https_auto_cert.enable_path_normalize](data-sources--http_loadbalancer--reference--group-019.md#canonical-289da877ce1d9348c55378310a94862dde2014ed2bdc312ec5de318b410b6ff7)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-2ba73566b5d2f8460ac6670d0936192f46e26ad25347fecf4549bd80759abb11)
- [https_auto_cert.no_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-92c674b0dc9736d465350a71117143b3efa33055ffd5823ac08b0bfda705e14f)
- [https_auto_cert.non_default_loadbalancer](data-sources--http_loadbalancer--reference--group-019.md#canonical-f5e31eb273a07c8f3e56020b510a838abdc255f11cade33b6c97b437e448857d)
- [https_auto_cert.pass_through](data-sources--http_loadbalancer--reference--group-019.md#canonical-a98e6a4ab6e5f5e26acc643f9b4ccded462734d7eec0f5e940b573536276a97e)
- [https_auto_cert.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-d063c10b8a5e2a7c99b5c5167ef68b0739078a5c4ea93e8ac7e0fef7efb4edde)
- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-d6cc7070a4da0f1e6653eedda19de11f6b8ce9f24ac85b7160b2085bf9ca4f7c)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fc24498d21cafcec765385dad56e1ac1ccb981e91d3873fee067e48bd77bc3dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-addce61a0f99a3d3c788279c0a8abbdc8a2995601a8d620f13262e9effcc6a89"></a>

## https_auto_cert.coalescing_options — https_auto_cert.coalescing_options / 6ff1bda7da4b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-e2ff5feee3a1ae293158939a97729795c3d8a7d9b4042a95c2f12867960d7cdd)
- https_auto_cert.coalescing_options

<a id="canonical-c9d56e449b9e93904aa722736f67979cfdd562252c974e9ff2e55713a2bf050a"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

<a id="canonical-0f0d8c5d9fce0b0a5a5663697b25584ae94b391296383c7db3289ed8705dc4f8"></a>

## Direct properties — https_auto_cert.coalescing_options / 6ff1bda7da4b / 3

- [default_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-20be2efe87f697b4057a52e61ce59a5859a238bb27ff0a165c80d39e6fb23403): complete subsection reference.

- [strict_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-d50ff4af37a3c6e4be391145e7659e576779e861d27ba6d6d3a83aae8de3a14b): complete subsection reference.

<a id="canonical-b6e205b8abb5a5c1309d957e70e0ea72108e072fa1684545a660b10e89fcb7d3"></a>

## Next pages — https_auto_cert.coalescing_options / 6ff1bda7da4b / 4

- [https_auto_cert.coalescing_options.default_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-20be2efe87f697b4057a52e61ce59a5859a238bb27ff0a165c80d39e6fb23403)
- [https_auto_cert.coalescing_options.strict_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-d50ff4af37a3c6e4be391145e7659e576779e861d27ba6d6d3a83aae8de3a14b)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-e2ff5feee3a1ae293158939a97729795c3d8a7d9b4042a95c2f12867960d7cdd)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-20be2efe87f697b4057a52e61ce59a5859a238bb27ff0a165c80d39e6fb23403"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb8d4914fc23716e98d4827e5d07c5a8f4a6d919db34e629102ef1f49723e810"></a>

## https_auto_cert.coalescing_options.default_coalescing — https_auto_cert.coalescing_options.default_coalescing / 657ed2a48d85 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-e2ff5feee3a1ae293158939a97729795c3d8a7d9b4042a95c2f12867960d7cdd)
- [https_auto_cert.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-fc24498d21cafcec765385dad56e1ac1ccb981e91d3873fee067e48bd77bc3dc)
- https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-c096e96da5dc6da7f069356346cb1fb74a1ba12391a7f7478842e1eb1f232323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-f011de3544d21f1ba6a5b59c8bdec17638c6f5af9ae6d7c45bdcc5789d7cf72b"></a>

## Direct properties — https_auto_cert.coalescing_options.default_coalescing / 657ed2a48d85 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2d323777802a22a59346b48b7fa6201d92ff08c06b87cef52a8b45ce9bef4c41"></a>

## Next pages — https_auto_cert.coalescing_options.default_coalescing / 657ed2a48d85 / 4

- [https_auto_cert.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-fc24498d21cafcec765385dad56e1ac1ccb981e91d3873fee067e48bd77bc3dc)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d50ff4af37a3c6e4be391145e7659e576779e861d27ba6d6d3a83aae8de3a14b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-437846be7ec60977325d85f32c8c013ee232ed574ffb93eed658f290500f4ab0"></a>

## https_auto_cert.coalescing_options.strict_coalescing — https_auto_cert.coalescing_options.strict_coalescing / 0c52cb0b4a06 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-e2ff5feee3a1ae293158939a97729795c3d8a7d9b4042a95c2f12867960d7cdd)
- [https_auto_cert.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-fc24498d21cafcec765385dad56e1ac1ccb981e91d3873fee067e48bd77bc3dc)
- https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-6e953ab96721f9d26b3963e33abc6c1d6f8cf90e92c0fdeefc5f99eae59ede24"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-04613a08addb8874354f25bd3703e319a48af2fa44901d97b3c4542dd0ba9dae"></a>

## Direct properties — https_auto_cert.coalescing_options.strict_coalescing / 0c52cb0b4a06 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aabac8e921de2fc0213247bf00a8216cbc0d0890e816d27de1d15bcc05803d09"></a>

## Next pages — https_auto_cert.coalescing_options.strict_coalescing / 0c52cb0b4a06 / 4

- [https_auto_cert.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-fc24498d21cafcec765385dad56e1ac1ccb981e91d3873fee067e48bd77bc3dc)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f784dc759277c3a4f7b03f5524656d83772059bf09a0bae3adb8e5df6eab1aa0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
