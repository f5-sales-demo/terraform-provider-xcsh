---
page_title: "xcsh_cluster reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster reference."
---

# xcsh_cluster reference

<a id="canonical-3a4ecaa57759e02be634551ff05d0c39c0d1053c7dbead0a9a3dfebfa246b814"></a>

## tls_parameters.use_host_header_as_sni — tls_parameters.use_host_header_as_sni / e526dff04deb / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- tls_parameters.use_host_header_as_sni

<a id="canonical-8753f5686a4b82b8e9d0088e463c4bfc5f50f7d4daeb556935641dbc909f8315"></a>

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

<a id="canonical-8ca7fe1a34bda38f5315480d3fcc654367d52541fa37d621512d230d58d9857d"></a>

## Direct properties — tls_parameters.use_host_header_as_sni / e526dff04deb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9fc14c0eec5512a6270822b8b77173349d217eeda4881aa4177cf59dd16f39dc"></a>

## Next pages — tls_parameters.use_host_header_as_sni / e526dff04deb / 4

- [tls_parameters](data-sources--cluster--reference--group-001.md#canonical-c0f5461277487bb91d92c5f204918cbc1896f3326805166d27a9234851250ccd)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-3706587128d7e7984d788fb07d46cccc9496e78ece7b56b721fccb29942c6506"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32bf8b66a87e92fab6ead72391e6514935fd0576ee0af1807e1c3734c963e9e2"></a>

## upstream_conn_pool_reuse_type — upstream_conn_pool_reuse_type / cde2ad87a714 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- upstream_conn_pool_reuse_type

<a id="canonical-0aa2f695104716438b0339fd7260249546d56960117d588d5fa935ff6994a028"></a>

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

<a id="canonical-17d81470010b00bf36597ce358f09b5a387c8c25881fa43d6e80816bf093fd71"></a>

## Direct properties — upstream_conn_pool_reuse_type / cde2ad87a714 / 3

- [disable_conn_pool_reuse](data-sources--cluster--reference--group-002.md#canonical-d0f0a7c575ad60b100d0817532871b0165495e338948ba9689725236efcd46c4): complete subsection reference.

- [enable_conn_pool_reuse](data-sources--cluster--reference--group-002.md#canonical-612d82a25f7d5d3c65aff0d7a0c4944f5d3a866a7c2278c6b9e8fa03497a8af3): complete subsection reference.

<a id="canonical-c66b1171836a5c2de03839ff78c1d0da7ab648463594aac415afa869d899308f"></a>

## Next pages — upstream_conn_pool_reuse_type / cde2ad87a714 / 4

- [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](data-sources--cluster--reference--group-002.md#canonical-d0f0a7c575ad60b100d0817532871b0165495e338948ba9689725236efcd46c4)
- [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](data-sources--cluster--reference--group-002.md#canonical-612d82a25f7d5d3c65aff0d7a0c4944f5d3a866a7c2278c6b9e8fa03497a8af3)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-d0f0a7c575ad60b100d0817532871b0165495e338948ba9689725236efcd46c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e77444e677425621218674170e0482d18df584c17a1a9bb36601dead6ca014f"></a>

## upstream_conn_pool_reuse_type.disable_conn_pool_reuse — upstream_conn_pool_reuse_type.disable_conn_pool_reuse / c6f3974feae6 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [upstream_conn_pool_reuse_type](data-sources--cluster--reference--group-002.md#canonical-3706587128d7e7984d788fb07d46cccc9496e78ece7b56b721fccb29942c6506)
- upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-25bc61fac320825cf64fc9d6ef46709c9d0e0486e5f462d396c1f552258cee24"></a>

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

<a id="canonical-7f6301075722ab3babd2a5cc3acc9daf7eab3dccdb5ec0752625e4d138554023"></a>

## Direct properties — upstream_conn_pool_reuse_type.disable_conn_pool_reuse / c6f3974feae6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-73a0f9271d872407980b0c8d3bf43e9e1019e0cf2b0f6970441d3e3f7b907529"></a>

## Next pages — upstream_conn_pool_reuse_type.disable_conn_pool_reuse / c6f3974feae6 / 4

- [upstream_conn_pool_reuse_type](data-sources--cluster--reference--group-002.md#canonical-3706587128d7e7984d788fb07d46cccc9496e78ece7b56b721fccb29942c6506)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-612d82a25f7d5d3c65aff0d7a0c4944f5d3a866a7c2278c6b9e8fa03497a8af3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a792c9112a591be33a504317da4a068cf10d346a5489ddc2652641e2079ef63"></a>

## upstream_conn_pool_reuse_type.enable_conn_pool_reuse — upstream_conn_pool_reuse_type.enable_conn_pool_reuse / edd040bd10db / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Property reference](data-sources--cluster--reference--group-001.md#canonical-88d6cbc7b1b7f80d007d47f4a04ed948a4b4f71b1f18ce47a8a781250b97b660)
- [upstream_conn_pool_reuse_type](data-sources--cluster--reference--group-002.md#canonical-3706587128d7e7984d788fb07d46cccc9496e78ece7b56b721fccb29942c6506)
- upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-557ed43439c68e3f9658e13df87c0faac86a857df1eab5065ca634aeaba0e771"></a>

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

<a id="canonical-66b7492e62681705899c5058d056357c39b4f73d988cfe546545be6885c10e04"></a>

## Direct properties — upstream_conn_pool_reuse_type.enable_conn_pool_reuse / edd040bd10db / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2346969f5ee64e8a091b16b2a51d3522a2581bbbcc61ab37bfca1fc3cf6562d5"></a>

## Next pages — upstream_conn_pool_reuse_type.enable_conn_pool_reuse / edd040bd10db / 4

- [upstream_conn_pool_reuse_type](data-sources--cluster--reference--group-002.md#canonical-3706587128d7e7984d788fb07d46cccc9496e78ece7b56b721fccb29942c6506)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
