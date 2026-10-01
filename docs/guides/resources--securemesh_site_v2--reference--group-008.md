---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-59c27ed663a3ad49c30632cffac7f25c9a0e0e0681450601cc3d2fb28097ac4b"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.monitor_disabled / ad016b8fb3a9 / 4

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-4635bf9274a9d8f090b4d1f7167c939581652a0ca999e4cdc3e1d263f84b939f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce46294b3c578f6e0e387cc60b647f45739b36b7bda3e84e7e526526af5e81d9"></a>

## eks_k8s.not_managed.node_list.interface_list.network_option — eks_k8s.not_managed.node_list.interface_list.network_option / 2466cc144696 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0a71cc062cfc0d2f7a3e7af6f9442dd36174c372898c92591f150e614cc351d0)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-5459ad007ffa47652f345f67082963a5d24f40d6322417dda277e25a2ad47204)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-993b7eabff037dffd429554f118618d451ca2820fc05ccb505c992bd9a66c46d)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- eks_k8s.not_managed.node_list.interface_list.network_option

<a id="canonical-a4c03736fb1aca065a686b3ae39c90cd2667ef51fa0adf915ef0005a8fbb96e4"></a>

Type: `"object"`. single nested block, Optional.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional.

Upstream description:

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
network_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-dde35336ea61fb141c356acd6c5e319e58c1f103c1cca794198c78c1e0bab208"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.network_option / 2466cc144696 / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-008.md#canonical-d510e2164966dd4ab859dc1c7091ac44e91dcc2cbef868120f2df6bada8bfd8b): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-008.md#canonical-872a1f536ebfe2c8f656f922405bc74baef37f79737fa59cb741bc4ed60eb55d): complete subsection reference.

<a id="canonical-92d9f05a663751b1dc2f38d47d210729e24de2bef40e1f4413b8a2ef4b34708e"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.network_option / 2466cc144696 / 4

- [eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-008.md#canonical-d510e2164966dd4ab859dc1c7091ac44e91dcc2cbef868120f2df6bada8bfd8b)
- [eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-008.md#canonical-872a1f536ebfe2c8f656f922405bc74baef37f79737fa59cb741bc4ed60eb55d)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d510e2164966dd4ab859dc1c7091ac44e91dcc2cbef868120f2df6bada8bfd8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-030992405324fe88a0bded23d3740911f27545fd39e65f4605f221efb523095c"></a>

## eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_network — eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_ne / 5c0b5ac690ec / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0a71cc062cfc0d2f7a3e7af6f9442dd36174c372898c92591f150e614cc351d0)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-5459ad007ffa47652f345f67082963a5d24f40d6322417dda277e25a2ad47204)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-993b7eabff037dffd429554f118618d451ca2820fc05ccb505c992bd9a66c46d)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- [eks_k8s.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-008.md#canonical-4635bf9274a9d8f090b4d1f7167c939581652a0ca999e4cdc3e1d263f84b939f)
- eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-9a662c27fdce85fb977ac6bbc7dafe53c1c1e70c998525e794c78e6e6763047e"></a>

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
site_local_inside_network = {}
```

<a id="canonical-736e4d147c3929670a66b043709a1759b70fd8d949b246a86ae63e17f8c78561"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_ne / 5c0b5ac690ec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-553115f60c32b39c6a5b3123bea14225185ef18a61e29e86a88b92be4b73ef07"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_ne / 5c0b5ac690ec / 4

- [eks_k8s.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-008.md#canonical-4635bf9274a9d8f090b4d1f7167c939581652a0ca999e4cdc3e1d263f84b939f)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-872a1f536ebfe2c8f656f922405bc74baef37f79737fa59cb741bc4ed60eb55d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f1e7dd54e965bf2c537a9a9618ce1ec9744b3ae205164661d2c4ddb96fb81ad"></a>

## eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network — eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network / 0e33ca431b08 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0a71cc062cfc0d2f7a3e7af6f9442dd36174c372898c92591f150e614cc351d0)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-5459ad007ffa47652f345f67082963a5d24f40d6322417dda277e25a2ad47204)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-993b7eabff037dffd429554f118618d451ca2820fc05ccb505c992bd9a66c46d)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- [eks_k8s.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-008.md#canonical-4635bf9274a9d8f090b4d1f7167c939581652a0ca999e4cdc3e1d263f84b939f)
- eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-c8c5b17d72ec13997dbbc8e269590ec157776e56bdb778e1dda244601dfbca4f"></a>

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
site_local_network = {}
```

<a id="canonical-d9eb7eb9ddaed6648941e1ac93e42e692089220d6fdae6002e889aa2543c7fe4"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network / 0e33ca431b08 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c4725a32a1b8c9e390b7c2be4a802551c4ff2acd78ca088c571c4a0cdd2bb302"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network / 0e33ca431b08 / 4

- [eks_k8s.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-008.md#canonical-4635bf9274a9d8f090b4d1f7167c939581652a0ca999e4cdc3e1d263f84b939f)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-b135a2fbb632f7ee031cf33bebb52a2ad90dd48a449c1f9d1e2399a3d911560b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78c4a30cf1ff3a4a4446d3ce41f16b7cadea3c83a11cbfbad9647f5e7f9cea40"></a>

## eks_k8s.not_managed.node_list.interface_list.no_ipv4_address — eks_k8s.not_managed.node_list.interface_list.no_ipv4_address / 1b7cbc3fdcc8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0a71cc062cfc0d2f7a3e7af6f9442dd36174c372898c92591f150e614cc351d0)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-5459ad007ffa47652f345f67082963a5d24f40d6322417dda277e25a2ad47204)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-993b7eabff037dffd429554f118618d451ca2820fc05ccb505c992bd9a66c46d)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- eks_k8s.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-aca0548a0e95720c14087ea0c5a1a1b0c401cfde73b5dca72310735421fc5c78"></a>

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
no_ipv4_address = {}
```

<a id="canonical-b0710df1ff920969dfbe58461078e9b18f88021d619ddbac411cc146825ce343"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.no_ipv4_address / 1b7cbc3fdcc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b67649002d23a4ed8a03f6a1965171431fcbf3f5807b69dfc56e5167820fef40"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.no_ipv4_address / 1b7cbc3fdcc8 / 4

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5c91714421ee3f5d95899baf02b8308257a79aa8cea581d871db6268e5fd8186"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02017b1b5e36a720a853ade064c21458332ab55cfdf209e9653f1e7edbd2ec6a"></a>

## eks_k8s.not_managed.node_list.interface_list.no_ipv6_address — eks_k8s.not_managed.node_list.interface_list.no_ipv6_address / 1de92d932f00 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0a71cc062cfc0d2f7a3e7af6f9442dd36174c372898c92591f150e614cc351d0)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-5459ad007ffa47652f345f67082963a5d24f40d6322417dda277e25a2ad47204)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-993b7eabff037dffd429554f118618d451ca2820fc05ccb505c992bd9a66c46d)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- eks_k8s.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-06a96b51997f6fa7fba6fdbf4ef7f24092e3d09c1e523d0e4779c6083b6e530d"></a>

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
no_ipv6_address = {}
```

<a id="canonical-84ec3a713d380402ff13285bc85d9ae7f53a7f7632e81416bd6aa4c907048e9a"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.no_ipv6_address / 1de92d932f00 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ca5ec571390a635d6d7975160fb77bda79c7df3ae00402d2a1e53f6b4e2daef6"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.no_ipv6_address / 1de92d932f00 / 4

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-13fa44aa11dc66d7d4ac4f2d6f90df9a688587324ca353c07dd0f09fe7b6def4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d5e4630933ad9d1f98085679f8a8f8d0cbdc9ec8e332b4e2083edb42ca7c04a"></a>

## eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface / 9b9ebaaf63f1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0a71cc062cfc0d2f7a3e7af6f9442dd36174c372898c92591f150e614cc351d0)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-5459ad007ffa47652f345f67082963a5d24f40d6322417dda277e25a2ad47204)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-993b7eabff037dffd429554f118618d451ca2820fc05ccb505c992bd9a66c46d)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-ad0b84f860431a89a0b1104d9e20e9f8afb88fb680b410e5079fbf058c1dc732"></a>

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
site_to_site_connectivity_interface_disabled = {}
```

<a id="canonical-ed6ed63a3bc984d66b0bc13afde473f2be4fe3050d5ec676470d7e9a0d227fea"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface / 9b9ebaaf63f1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9609ef2537d6992e5b677f42d59451a9947c109e345cf0ba4d9f4762d4dfdfdc"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface / 9b9ebaaf63f1 / 4

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-8d88fc0653c09b1d9de9b265d50d9bfd19f66426aa2684bdb5e795af37ee82d6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d95a6041eeafad14147b359b2950854b9e2033af6afa731dec85e7ded5f526a2"></a>

## eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface / 3db1d143daeb / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0a71cc062cfc0d2f7a3e7af6f9442dd36174c372898c92591f150e614cc351d0)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-5459ad007ffa47652f345f67082963a5d24f40d6322417dda277e25a2ad47204)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-993b7eabff037dffd429554f118618d451ca2820fc05ccb505c992bd9a66c46d)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-7f52a999d52c6eb9b8949860a38c98bb23fda727f0fa4c7789b4739a53f781bb"></a>

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
site_to_site_connectivity_interface_enabled = {}
```

<a id="canonical-589f673bb2d30bfbb681bff16f3bf9b402a34e7557fe9898568ffa7b2a9c58c3"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface / 3db1d143daeb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-26c5f7fb65ae46badbf76c1db35d558a83dc954b06c45f2a1ac5081e6746b275"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface / 3db1d143daeb / 4

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-cc94e4b1073a7863bdab9ff80567050333eec8b96c08ffa9ff763385f11c7856"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69b8302e9aafe8fb88934eadb5410fe6f3a6ea74340974d39c09f8242455b536"></a>

## eks_k8s.not_managed.node_list.interface_list.static_ip — eks_k8s.not_managed.node_list.interface_list.static_ip / 35d0e029f6c1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0a71cc062cfc0d2f7a3e7af6f9442dd36174c372898c92591f150e614cc351d0)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-5459ad007ffa47652f345f67082963a5d24f40d6322417dda277e25a2ad47204)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-993b7eabff037dffd429554f118618d451ca2820fc05ccb505c992bd9a66c46d)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- eks_k8s.not_managed.node_list.interface_list.static_ip

<a id="canonical-204b31b755edb01240bd8ba4f89d442beb9d23e2cb2a8c224663eef48b71bdf1"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0ed7ba0880ebab487a05670d90a703f2f36de6c7f3e06d077d7808a38f6dfa69"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.static_ip / 35d0e029f6c1 / 3

<a id="canonical-d0635997024babe9fb98cad9683b779ee527462c62e86c02c120bec7f5b0139d"></a>

<a id="canonical-5d12dc2f2c0b440af94c62a358aa8ec7fd7dabff0a425cbc5977d5310130b2a7"></a>

## default_gw property — eks_k8s.not_managed.node_list.interface_list.static_ip / 35d0e029f6c1 / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-faeebc68fafb6947f5b0f98274d8d7a11017c8d5c9f83046fd90afc459ee2e84"></a>

<a id="canonical-382911d4a7d4a3dd6763450858ea44bdf6d1b0f18bb1f5acdd6719ed59e5a2b8"></a>

## dns_server property — eks_k8s.not_managed.node_list.interface_list.static_ip / 35d0e029f6c1 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-f86fbbb6b0bbd4b38f7a6944e877dc7a90d440bb3d6d8824629108e3e6609c3c"></a>

<a id="canonical-63f5e86d0837a147853d2132a8a00083c889b7540cd5ba130802271be06e2a3b"></a>

## ip_address property — eks_k8s.not_managed.node_list.interface_list.static_ip / 35d0e029f6c1 / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-8ea37f94906ffc18ceb1ebd814cdc3711478bf5474a6781f3f779a17ab5760c9"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.static_ip / 35d0e029f6c1 / 7

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9a2601a73eb1e7941d80e8a90f878ddf0a26d8bca1402fafb2ab071ddf8e4d29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-439b7cf816ebfef84acd508f36ad6371cc3e8ef068477713a1ed08a5e2423ec6"></a>

## eks_k8s.not_managed.node_list.interface_list.static_ipv6_address — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address / 13b4374abbdf / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0a71cc062cfc0d2f7a3e7af6f9442dd36174c372898c92591f150e614cc351d0)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-5459ad007ffa47652f345f67082963a5d24f40d6322417dda277e25a2ad47204)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-993b7eabff037dffd429554f118618d451ca2820fc05ccb505c992bd9a66c46d)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- eks_k8s.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-2a76bf8190d2a17c770b8f8a8601edc94b3f9cfef1c27ef2f2b717b20bcdf639"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
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
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ipv6_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-8017366b957549037f1b6b4173e1658457201c9b541681e761766cd5c600884f"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address / 13b4374abbdf / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-008.md#canonical-3ac18035400000e756e45e6f168e0c0c31cf590f96177994435ad4db2f8c023b): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-008.md#canonical-6ee2bdf70b58623f5bb65630699c9b5f60bcd558b2216c6b43c1f7c5c8719835): complete subsection reference.

<a id="canonical-58cf929aa391dce7ea3a0ca18b527bdfcda9a65da611e3a62c8b28e125375d0f"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address / 13b4374abbdf / 4

- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-008.md#canonical-3ac18035400000e756e45e6f168e0c0c31cf590f96177994435ad4db2f8c023b)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-008.md#canonical-6ee2bdf70b58623f5bb65630699c9b5f60bcd558b2216c6b43c1f7c5c8719835)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-3ac18035400000e756e45e6f168e0c0c31cf590f96177994435ad4db2f8c023b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-236487535607af99974b9924d07ea74532e4e179e168297a9d4fda423a3c3537"></a>

## eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / 0a682873e4ad / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0a71cc062cfc0d2f7a3e7af6f9442dd36174c372898c92591f150e614cc351d0)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-5459ad007ffa47652f345f67082963a5d24f40d6322417dda277e25a2ad47204)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-993b7eabff037dffd429554f118618d451ca2820fc05ccb505c992bd9a66c46d)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-008.md#canonical-9a2601a73eb1e7941d80e8a90f878ddf0a26d8bca1402fafb2ab071ddf8e4d29)
- eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-9f94c48e18eaea39d0286752b97b8463c37c13c7e8f19c9185f95725e2b6e7c0"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for cluster.

Receipt-pinned upstream constraints:

```json
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
cluster_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-60b9dab2d0d2b7b31135d63a3d259dfef4ab50faa0781d6a068a2a34aeb9a294"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / 0a682873e4ad / 3

<a id="canonical-db2c76598203f9746d7c2297adeeba6073ef4549b7759f1f696f4962fafdd2fc"></a>

<a id="canonical-c9b721ec42f7bb67ccf99f5a810158c2f69bc2923f63699c6feae1bd2eddb0c6"></a>

## interface_ip_map property — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / 0a682873e4ad / 4

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

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
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-5e8e210a062d7a9e94bcbad134c334fbfa0f9d703317bfac699de8941f314d19"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / 0a682873e4ad / 5

- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-008.md#canonical-9a2601a73eb1e7941d80e8a90f878ddf0a26d8bca1402fafb2ab071ddf8e4d29)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6ee2bdf70b58623f5bb65630699c9b5f60bcd558b2216c6b43c1f7c5c8719835"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d192a49842fbd31035c81b14006b9486e2bee63cce119447ee9e292708c575ed"></a>

## eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 1390b6f8a85e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0a71cc062cfc0d2f7a3e7af6f9442dd36174c372898c92591f150e614cc351d0)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-5459ad007ffa47652f345f67082963a5d24f40d6322417dda277e25a2ad47204)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-993b7eabff037dffd429554f118618d451ca2820fc05ccb505c992bd9a66c46d)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-008.md#canonical-9a2601a73eb1e7941d80e8a90f878ddf0a26d8bca1402fafb2ab071ddf8e4d29)
- eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-89afa80074b5c8460417df82ae37806c6c03229a01f1ee1d83009da0d734029e"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
node_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-ebffb57451c89ab4a9428dc36363a42e7161f875f4075cd33389e2a30486d8bd"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 1390b6f8a85e / 3

<a id="canonical-426e2023375cf7f63ebe456042323459d9e092ce6edfbb231a7790512ccfbaa1"></a>

<a id="canonical-1dcf07e8f6743cb4cf9fa2a3416f72ffe0ca05a6869ae278851d84890fb78484"></a>

## default_gw property — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 1390b6f8a85e / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-dfd52058f9cb5a58709d1c5b228e9f4639ad929d7d724120e22ee48d357c05dd"></a>

<a id="canonical-b2a2191f28f0eda3f87957ad6f7e7957486cdbe3cd8f1ea56a8050b7921afb72"></a>

## dns_server property — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 1390b6f8a85e / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-8a6a824aee5ce65f2f05ce4d6abc950de910cfeb20956295121b181ca76045cd"></a>

<a id="canonical-03e910d5292dcc258f40f398b7666cc0b8b1a73829ba70fd5ca39cc507ac458e"></a>

## ip_address property — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 1390b6f8a85e / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-71e403c877625042ee129164f65dae9c9b4c91d83dbc2abb210d6eda27cadea7"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 1390b6f8a85e / 7

- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-008.md#canonical-9a2601a73eb1e7941d80e8a90f878ddf0a26d8bca1402fafb2ab071ddf8e4d29)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-3139bbab50353e7b1fa33bc12bf2d582ecc95f128f00fdbfb2bde58dda90d891"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bff54897f572a45e2be346ba269095ffaef47710d0d0ab607bed57943f949a9"></a>

## eks_k8s.not_managed.node_list.interface_list.vlan_interface — eks_k8s.not_managed.node_list.interface_list.vlan_interface / ebf5a14c023b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0a71cc062cfc0d2f7a3e7af6f9442dd36174c372898c92591f150e614cc351d0)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-5459ad007ffa47652f345f67082963a5d24f40d6322417dda277e25a2ad47204)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-993b7eabff037dffd429554f118618d451ca2820fc05ccb505c992bd9a66c46d)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- eks_k8s.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-22e95dc142b017db99ffd803b92365865d308a000f22ce36a4390b2db975a30d"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vlan interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device",
    "vlan_id")}
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
vlan_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-e72cff6d7726474c150354339c5d381bcb674f7e13eab603c94bb1b2ce42ccbd"></a>

## Direct properties — eks_k8s.not_managed.node_list.interface_list.vlan_interface / ebf5a14c023b / 3

<a id="canonical-6e4cf59e82c52c7cee7e299f21f14257ea882bd35ed7b5a7037b66ed0fea7511"></a>

<a id="canonical-16a50175a10b98dcf78161fbea18a9ae369253bcef51230f64da24d565284918"></a>

## device property — eks_k8s.not_managed.node_list.interface_list.vlan_interface / ebf5a14c023b / 4

Type: `"string"`. Optional.

Select a parent interface from the dropdown.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-d59fa61dbe3397a6c74c4f8a22d9560c25e9389495180d15773ce5ba78798f82"></a>

<a id="canonical-643f06996aa467312987176c941d896707c055e46efd38cee13d0381bb2b9029"></a>

## vlan_id property — eks_k8s.not_managed.node_list.interface_list.vlan_interface / ebf5a14c023b / 5

Type: `"number"`. Optional.

Configure the VLAN tag for this interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-cb381efe2ee07e0f4382d740a6bd852d374ce2f6914403b3e19831702fa32086"></a>

## Next pages — eks_k8s.not_managed.node_list.interface_list.vlan_interface / ebf5a14c023b / 6

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-a88e13e2592b4034c25f3738d86891f703715cea3c46c1ce7265131387b87ee4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-58d87ed581de59f3c7eb9469b207629142c00856112d3a9917b12e7d7ebd180b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f33af56731bdbf7472edb9a4992d16c06cdc380d6828a4acb9c5b5f5a409cb89"></a>

## enable_advanced_delivery — enable_advanced_delivery / fe93f1dae8e1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- enable_advanced_delivery

<a id="canonical-687720315a3d4b373a5d5e7fe87148ba8b20ac5ca18af7bce589d215a0751e38"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable advanced delivery.

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
enable_advanced_delivery = {}
```

<a id="canonical-8be016988d6878e0d93c854ca65a78f81e0aace7313e9aced5fa92023db3cb50"></a>

## Direct properties — enable_advanced_delivery / fe93f1dae8e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8bf6ed21860a8df3b7a6e7436306d539f14d0c6048beed896cc84ef99664b93f"></a>

## Next pages — enable_advanced_delivery / fe93f1dae8e1 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-687f3b62d8ed4274294ed0a7726f16941786952f8fe35efd00640ef5dbe8a2de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f290b7d855dcdd7aa8d9e6f0a9e3563c048d38890dff3359dbb9427251f4e62"></a>

## enable_ha — enable_ha / d3cc3c4ecf15 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- enable_ha

<a id="canonical-bc044734ce2286958aaf0f933402b07d621ede0f900604fb48a7f7e1adcb708b"></a>

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
enable_ha = {}
```

<a id="canonical-baf9fa51f58cc1ac31c9c75d07301a37fade87e1129ff8e83a8830f74a13f85c"></a>

## Direct properties — enable_ha / d3cc3c4ecf15 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cabf80d979c13ce0a66c5d1bd7358e3f725e947976a0c4813bafccfd495ed56d"></a>

## Next pages — enable_ha / d3cc3c4ecf15 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f7fd420a7bf79edf5db7e0403c835c6b9609c913a68a4ed301ebdcc856ff1750"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3775d38ea0214530b88175b7ff6f7bc09c99b7523af47acd6114d156642df29"></a>

## enable_log_anonymization — enable_log_anonymization / 97eb8edc8e2b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- enable_log_anonymization

<a id="canonical-5e2c8841e47f64b2837038dd457fd7bc02c0dec0f5149cbbdaf9245d451e577b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable log anonymization.

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
enable_log_anonymization = {}
```

<a id="canonical-39b495a5da403165beac9386c69dc13258f8a8a6dbe4ecbc46bbdeb16e5fb7db"></a>

## Direct properties — enable_log_anonymization / 97eb8edc8e2b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6de0fe985a58c59172412646dea31300c15f3f07f024dfdb8e1787d2f26680ef"></a>

## Next pages — enable_log_anonymization / 97eb8edc8e2b / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d165033fedf36409580e6d7d7e956850c8ba9cc31e82f5d4457d87fd173f4398"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afb5910f42b63cc363eb00279be04d9a493ab268a1785174a5fc514aa355b186"></a>

## enable_management_network — enable_management_network / 238b835aa8e1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- enable_management_network

<a id="canonical-369a2b433de9c462e81b1d3f7c0e6882d561728dffb6f65800d6ac61dcb390ca"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable management network.

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
enable_management_network = {}
```

<a id="canonical-1cd35c27e6487becc6bd837351130150d35cb7501ca57156241aded05fe2be3d"></a>

## Direct properties — enable_management_network / 238b835aa8e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-831c43f737bde4f8650d86cd45124a19c4dddb07d03f4c3b08093d2237eb98c5"></a>

## Next pages — enable_management_network / 238b835aa8e1 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-942a730157531e299183f328ff838a8dc1db2259df9c81d356098acc27c24b9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-527531aa716874a5ee6413ca58d7afe3150c3ae18fc00f7736fc45a00505466a"></a>

## enable_url_categorization — enable_url_categorization / d7c9837d3e24 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- enable_url_categorization

<a id="canonical-3ac0afc9cde68a03d905b0fde4c35a2d97a5f00455cba0f8f55fdfb017ac072b"></a>

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
enable_url_categorization = {}
```

<a id="canonical-43061e8399808d9c43c2f994e0af3efb24f467f3c9e26c4261cc118e7a2d10d2"></a>

## Direct properties — enable_url_categorization / d7c9837d3e24 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0b7b194048eb66d2b8f01a41d66ec05a89eb5b3684fd8274f0163404550a375f"></a>

## Next pages — enable_url_categorization / d7c9837d3e24 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-345994128f5ad690d9f5f0d39d0e8e2bfac7c15828fbff60915be8de51e480f4"></a>

## equinix — equinix / e1e2fb1798bd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- equinix

<a id="canonical-5c00d11005f9e53ead02087bd775458f54e69d08f48c42b97aec719563a70dad"></a>

Type: `"object"`. single nested block, Optional.

Equinix Provider Type. Equinix Provider Type.

Upstream description:

Equinix Provider Type.

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
equinix {
  # Configure direct properties listed below.
}
```

<a id="canonical-19313281c28a2ac859b05fa31fea70bdd76ef88eaa09a9283864724a43d1f0a0"></a>

## Direct properties — equinix / e1e2fb1798bd / 3

- [not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90): complete subsection reference.

<a id="canonical-56190ae2640166c2f8a4c87ddd2ff815b439d449459792711ab77a14d2d28832"></a>

## Next pages — equinix / e1e2fb1798bd / 4

- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70c99205eb8010c128b284b2c938f41a937ec60934b63fa9a16700cc09088f1d"></a>

## equinix.not_managed — equinix.not_managed / d1f6a073dc3b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- equinix.not_managed

<a id="canonical-3dfdd123a633c08d294b8ba29ffa211a7348e912635e21d626345b3e40259a09"></a>

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

<a id="canonical-00b31bf81e087f390acd390ab8e5330f8643682557aa2eb1a7c12b3de4fe71e5"></a>

## Direct properties — equinix.not_managed / d1f6a073dc3b / 3

- [node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f): complete subsection reference.

<a id="canonical-12da953ab9992fc3cd5889bc43ea051b8a36a39079953f26238337cb995af9a1"></a>

## Next pages — equinix.not_managed / d1f6a073dc3b / 4

- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6596adb0095cb2d77ffbb01cb214262c78d356f6f6fd55c2f36bdddf370ded56"></a>

## equinix.not_managed.node_list — equinix.not_managed.node_list / 0979db505f3b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- equinix.not_managed.node_list

<a id="canonical-e6006d43f915b84be0b523481de8fe73622928bdc0a8d6347e8f8962a8be7379"></a>

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

<a id="canonical-190e0a677299e4774d6f96ede52f00566d7bf0be58e58e04dcc2455efe9c1084"></a>

## Direct properties — equinix.not_managed.node_list / 0979db505f3b / 3

<a id="canonical-ebf0dfe39de4cbca57efde52f5b40730925a1b88c0dba3406ae042c3e75fad6e"></a>

<a id="canonical-a0acafffd996d4ec4781e4368406fddaaef27c587c218e236ea8f7ea1e9ba7e8"></a>

## hostname property — equinix.not_managed.node_list / 0979db505f3b / 4

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

- [interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c): complete subsection reference.

<a id="canonical-b83e3d40c44d949a9583b94b2dd136359b27e60f0d564c38aa15e2a5dbf7b11c"></a>

<a id="canonical-df0177f758339ede443128de66772b35bbd861662ced06c50e7a15f2f24b09f6"></a>

## public_ip property — equinix.not_managed.node_list / 0979db505f3b / 5

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

<a id="canonical-e3f1845b6fa8e01b96a614831da7d35c0b70d7926eadbaf49de199e9c5c9d4b9"></a>

<a id="canonical-840c2cf1850b071168a34882a6cc8637216bda2cae840a99ebb493d427b81869"></a>

## type property — equinix.not_managed.node_list / 0979db505f3b / 6

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

<a id="canonical-ad3c736ecd637bc5033d079fa5653e6ddd24ab93de83238d9085d5b951ed5c97"></a>

## Next pages — equinix.not_managed.node_list / 0979db505f3b / 7

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a4fa38a4db47533aa1470f0a5871575a05dc08214a235bc0beff2eccafca11c"></a>

## equinix.not_managed.node_list.interface_list — equinix.not_managed.node_list.interface_list / f4d03dc4b9d2 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- equinix.not_managed.node_list.interface_list

<a id="canonical-d198f43e8153242dbe238cb0058262fad08fbdb8ea6de89b262dc78d94e165b6"></a>

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

<a id="canonical-98dca55a14e558ae366530f60866d5754503f560f2a9f24a81404e9c633253d6"></a>

## Direct properties — equinix.not_managed.node_list.interface_list / f4d03dc4b9d2 / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-8b24d95a7f8c48ac43b74945ef9f3c06cf98468fe528b53acede29da71f95968): complete subsection reference.

<a id="canonical-7f115492ab2dbd90934fb8e792fe80f5836c0241771bb8561cf752e1d6d0db24"></a>

<a id="canonical-82cd1bb00ac4efe22bc9041107573cd035d23f3eafbc1050207063fc626f545e"></a>

## description_spec property — equinix.not_managed.node_list.interface_list / f4d03dc4b9d2 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-008.md#canonical-012b7d2ce30c67cf98551d53b35c545be0f844539f99a53e9a2e06bc7d55d96a): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-fa344d5d9a8d955c01d1e0e0c94ca03461b0ae99cf28f49e181a3978f96ee1c2): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6): complete subsection reference.

<a id="canonical-952eec6829aabd36f4a9c2020f091d4171091f7ce1bc0fc4e529552206fffac6"></a>

<a id="canonical-beea0fba1daab9acd81b5eeb40015681faca633cfeeb9636b2429589bdfe16e8"></a>

## is_management property — equinix.not_managed.node_list.interface_list / f4d03dc4b9d2 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-0e889722fa1213eda2a2ed479d5decebd1cff61715a56cb0fce93b9e340167e7"></a>

<a id="canonical-a3ed97603d9c0bf7b99e8952ffc51f54278d88331472f3e626cee7e5eac8ef63"></a>

## is_primary property — equinix.not_managed.node_list.interface_list / f4d03dc4b9d2 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-c4b23540c13fcc0f3b0bfa812107a7902403f46f5f0ce5fc166f944fc3b0d358"></a>

<a id="canonical-367c959756e95e35e9e43ab9ac34b536516d431e42b242cbbf170b8eb75c6651"></a>

## labels property — equinix.not_managed.node_list.interface_list / f4d03dc4b9d2 / 7

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

- [monitor](resources--securemesh_site_v2--reference--group-009.md#canonical-fbf5804b938e70b250d4662bdd3a52ebfa361ea69199b210816c52e72aba3464): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-009.md#canonical-58f9af6210d386d904b25e71cf30a85287899ea13c6464aee92f1a9a8d7d7f12): complete subsection reference.

<a id="canonical-289fd2d637ccbb002c40a67a424e41f6410b2401550422887549c0f7e10d0a7f"></a>

<a id="canonical-c7d1dfb92ea27c87402d385d0a022b415721420565afbb90d0497fc7b689b6eb"></a>

## mtu property — equinix.not_managed.node_list.interface_list / f4d03dc4b9d2 / 8

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

<a id="canonical-79ba736ce1385b9cf54110ec4bbea93cbb423df6f299d5bd89b58e9dea43c85c"></a>

<a id="canonical-1b7a48eaa381ca1df1e3df8a2ba462b5528be28f36e7313a83c045d411528462"></a>

## name property — equinix.not_managed.node_list.interface_list / f4d03dc4b9d2 / 9

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

- [network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-6aec47591d0b4b6ab9b29ac34dd0a58e15c9ace78e3f58baaeb0aa3dee04db18): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-009.md#canonical-576e9a39c18b846647c47d4f4518fbe43ab40bfac003f13930a0aebfe06a661f): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-65a381955fa5cd94cc5fa2fbc222152f00bc574e39921048c235ddd028d24854): complete subsection reference.

<a id="canonical-b2d5d4d6cce20923998c1c79ce767d03c453b78a3363640c88dda2e1e1b49129"></a>

<a id="canonical-c43b7b9cdfe728f70ad2dad8c97c893f03ba5986cd3f470013955eed82f51a6d"></a>

## priority property — equinix.not_managed.node_list.interface_list / f4d03dc4b9d2 / 10

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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-009.md#canonical-40cba71be545e59450d86e3f8ff40174e36f79e098e9ded35f2e282a5c77d951): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-009.md#canonical-b7349d6306fa8949ba028853754c52c8cfd146d76baed5a41c8267bb705a072a): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-89207d7f8b2c06f093ea659df00de8be4131ba92fa4513d8bf017732eebb960b): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-90a8505774e15c93615dbbcd1864cd8bb1dcd473856f3b8c346fad00ae52edbc): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-7884ed7284e85608b088cc98b43a4f6bd2fdb09a059944df8d2c819dc7536ed9): complete subsection reference.

<a id="canonical-311598a3ca785d3bf0319cff69c4c93466c7b7f3f564fe4bef15957790ba4dd0"></a>

## Next pages — equinix.not_managed.node_list.interface_list / f4d03dc4b9d2 / 11

- [equinix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-8b24d95a7f8c48ac43b74945ef9f3c06cf98468fe528b53acede29da71f95968)
- [equinix.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-008.md#canonical-012b7d2ce30c67cf98551d53b35c545be0f844539f99a53e9a2e06bc7d55d96a)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7)
- [equinix.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-fa344d5d9a8d955c01d1e0e0c94ca03461b0ae99cf28f49e181a3978f96ee1c2)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- [equinix.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-009.md#canonical-fbf5804b938e70b250d4662bdd3a52ebfa361ea69199b210816c52e72aba3464)
- [equinix.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-009.md#canonical-58f9af6210d386d904b25e71cf30a85287899ea13c6464aee92f1a9a8d7d7f12)
- [equinix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-6aec47591d0b4b6ab9b29ac34dd0a58e15c9ace78e3f58baaeb0aa3dee04db18)
- [equinix.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-009.md#canonical-576e9a39c18b846647c47d4f4518fbe43ab40bfac003f13930a0aebfe06a661f)
- [equinix.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-65a381955fa5cd94cc5fa2fbc222152f00bc574e39921048c235ddd028d24854)
- [equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-009.md#canonical-40cba71be545e59450d86e3f8ff40174e36f79e098e9ded35f2e282a5c77d951)
- [equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-009.md#canonical-b7349d6306fa8949ba028853754c52c8cfd146d76baed5a41c8267bb705a072a)
- [equinix.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-89207d7f8b2c06f093ea659df00de8be4131ba92fa4513d8bf017732eebb960b)
- [equinix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-90a8505774e15c93615dbbcd1864cd8bb1dcd473856f3b8c346fad00ae52edbc)
- [equinix.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-7884ed7284e85608b088cc98b43a4f6bd2fdb09a059944df8d2c819dc7536ed9)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-8b24d95a7f8c48ac43b74945ef9f3c06cf98468fe528b53acede29da71f95968"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d1fe8a6a16123d105b82a7389af30db3b3f005e3a912f4aefbbb18bfa87e7b0"></a>

## equinix.not_managed.node_list.interface_list.bond_interface — equinix.not_managed.node_list.interface_list.bond_interface / c5d4b270742f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.bond_interface

<a id="canonical-86019d5edc9464176609c971344ccd046645c766c32034dcd93de3383179156c"></a>

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

<a id="canonical-6396393f468867a98218a0c0cfe61024babdff53c81d3b397d407cf014aeba82"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.bond_interface / c5d4b270742f / 3

- [active_backup](resources--securemesh_site_v2--reference--group-008.md#canonical-1869535154ad156e368c9987804009cfc1bef8309803c7affb18fb408d1377c0): complete subsection reference.

<a id="canonical-7221d49fa7777c14c5a6eed0c2cd6608d66fed9fcb3b307a3a1d4f0937d6a9c6"></a>

<a id="canonical-6e04785627bf27c073498a2db57125388210804701c140bc36f265f5d45282cb"></a>

## devices property — equinix.not_managed.node_list.interface_list.bond_interface / c5d4b270742f / 4

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

- [lacp](resources--securemesh_site_v2--reference--group-008.md#canonical-61d47cf9226a4aa6afeb0546bea22f576f129728877bcf25310512c0de12512f): complete subsection reference.

<a id="canonical-fafec8f914089fbcaa6b38aa54a9dc8736f6be575d3f66cdf57504436e0bbfe3"></a>

<a id="canonical-b027a87c57a2b1f58e60d177dd78b6c26d0939586992eb4585c2792190b07e88"></a>

## link_polling_interval property — equinix.not_managed.node_list.interface_list.bond_interface / c5d4b270742f / 5

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

<a id="canonical-531641fb9ece26fa54a7493e715a4d5b66f4a8f17f294401c3328c58e5aadf88"></a>

<a id="canonical-14942a60be5644cdbeed68606355ef6d96827fc676ce4afd4159db0f316544f3"></a>

## link_up_delay property — equinix.not_managed.node_list.interface_list.bond_interface / c5d4b270742f / 6

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

<a id="canonical-ed7660d9670db7dfed9a71210e2eae6201c6b9948f47c8258431b68ede82fd64"></a>

<a id="canonical-504cb09a19722d3c4d105a1cecf49954073a314bb7cce91b0055a4def05a1e85"></a>

## name property — equinix.not_managed.node_list.interface_list.bond_interface / c5d4b270742f / 7

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

<a id="canonical-e3cca4e61c05484c4f0d4e04490265e11b47d218740cf64574dbfcd9735f25f4"></a>

## Next pages — equinix.not_managed.node_list.interface_list.bond_interface / c5d4b270742f / 8

- [equinix.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-008.md#canonical-1869535154ad156e368c9987804009cfc1bef8309803c7affb18fb408d1377c0)
- [equinix.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-008.md#canonical-61d47cf9226a4aa6afeb0546bea22f576f129728877bcf25310512c0de12512f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1869535154ad156e368c9987804009cfc1bef8309803c7affb18fb408d1377c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2afd2714ce9aba13ad5f0d4373bb761576de75e13a37529bbc58da8a8ba73e48"></a>

## equinix.not_managed.node_list.interface_list.bond_interface.active_backup — equinix.not_managed.node_list.interface_list.bond_interface.active_backup / 4efabd1055df / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-8b24d95a7f8c48ac43b74945ef9f3c06cf98468fe528b53acede29da71f95968)
- equinix.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-6e0d056852430d1f13bda2d03f60766f52078fe1923236b56d65c4f58ef6fa78"></a>

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

<a id="canonical-ac2b0315e7d8d2ff66979e1077fa1aee018078f8af2f2f8d48cf319c94ac48c2"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.bond_interface.active_backup / 4efabd1055df / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-11c48cb6457fb0c854ec2af894a6a968dfef74af68c47a7799a278d5d489ca33"></a>

## Next pages — equinix.not_managed.node_list.interface_list.bond_interface.active_backup / 4efabd1055df / 4

- [equinix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-8b24d95a7f8c48ac43b74945ef9f3c06cf98468fe528b53acede29da71f95968)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-61d47cf9226a4aa6afeb0546bea22f576f129728877bcf25310512c0de12512f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b940ba41a11d0df9cbc3c6b18f8dbb430df17fc2f5a5cca37601d70243a20949"></a>

## equinix.not_managed.node_list.interface_list.bond_interface.lacp — equinix.not_managed.node_list.interface_list.bond_interface.lacp / 36ca329fe6d6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-8b24d95a7f8c48ac43b74945ef9f3c06cf98468fe528b53acede29da71f95968)
- equinix.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-8a148855c9b1123f5c8ea8e5ab4ac7e66a5883e724b88b02a1d8f75e9f194dc2"></a>

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

<a id="canonical-f08eb836a1cbc41a8282eef0b83a25dba271aebed6b0926a1a3d432c47ea1c97"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.bond_interface.lacp / 36ca329fe6d6 / 3

<a id="canonical-59decee006dbe1b3cee1ee58a6a808b06d6ab55cfdc2231b1eb69b63d0eb9bbd"></a>

<a id="canonical-0b8f88969b32ddbc7b7043bd3af4c51c5b9d2401e81375d626abbbf1b6243d60"></a>

## rate property — equinix.not_managed.node_list.interface_list.bond_interface.lacp / 36ca329fe6d6 / 4

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

<a id="canonical-5cc83da6c012e7fa21ed20fa5ae1d0225438688d7ec4cc1f4190341a82190894"></a>

## Next pages — equinix.not_managed.node_list.interface_list.bond_interface.lacp / 36ca329fe6d6 / 5

- [equinix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-8b24d95a7f8c48ac43b74945ef9f3c06cf98468fe528b53acede29da71f95968)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-012b7d2ce30c67cf98551d53b35c545be0f844539f99a53e9a2e06bc7d55d96a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2bc579d53afca7e5c52d6db00a8f2ab1b66ca28238cde41fda846cda1c59f13"></a>

## equinix.not_managed.node_list.interface_list.dhcp_client — equinix.not_managed.node_list.interface_list.dhcp_client / 491329a30317 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-a3a7665e82783cca38e4413ef95e670fd369c0faaa655c36eaab5082e8dd19c9"></a>

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

<a id="canonical-966ff5c179f896deecf6a131def43ad4899723c145e67e7c9e250ab50add383e"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_client / 491329a30317 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-97e3c383f6dcfd9d9848645d97e4279f1a752b52040f752ed2fe89781f63cd11"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_client / 491329a30317 / 4

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60699148637560ef4d5d6c4f0767b71cd3d5fcefcaddb6eefc45f245b458518e"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server — equinix.not_managed.node_list.interface_list.dhcp_server / 87c00774410c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-8d308a791c047813bfbea46f2b537ce0f4a3a28b62b9d6ed6e6b3d99f0be0b93"></a>

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

<a id="canonical-ad500c76fa1053f8029e6962f74086c0085f526b17d83af79ee79a9cc51d94ce"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server / 87c00774410c / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-008.md#canonical-4d6f35e4e496aa4bff82a448b0016b6b717c939110dff109d2bf863809390467): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-008.md#canonical-28000afe26522ced81d309e638297f1d5913848cb0f06978f7f916af9982431d): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-2721d20ca3e3d0acd13cfa4dc587d59da8a946eadf4611cc6a3ed457117f25a4): complete subsection reference.

<a id="canonical-02d068546e3462d05b119cde80bb93d9ad5eccf33c3b66fae3b933530469a76c"></a>

<a id="canonical-440f1fcfe77b58e101b2d22082f0a5c6e71173ae0561ffdc9963b34a3a00cb5a"></a>

## dhcp_option82_tag property — equinix.not_managed.node_list.interface_list.dhcp_server / 87c00774410c / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-e654629de4b1b5e133a237092f6960108c7f601bd68e362000ec6ed9c8bd8d0d"></a>

<a id="canonical-0e9d5d597941bc20fd30d2cf0dc26c3a555d9f73a0630e89dea77954bdde2927"></a>

## fixed_ip_map property — equinix.not_managed.node_list.interface_list.dhcp_server / 87c00774410c / 5

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-008.md#canonical-f6eb574936577fdce71b4012797b90ebde301adf2ed6e476448eccf4a4e198db): complete subsection reference.

<a id="canonical-eb489b168c0d45a27c6dab0c2b78661447abf827392b376ded25f15b11ba1383"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server / 87c00774410c / 6

- [equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-008.md#canonical-4d6f35e4e496aa4bff82a448b0016b6b717c939110dff109d2bf863809390467)
- [equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-008.md#canonical-28000afe26522ced81d309e638297f1d5913848cb0f06978f7f916af9982431d)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-2721d20ca3e3d0acd13cfa4dc587d59da8a946eadf4611cc6a3ed457117f25a4)
- [equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-008.md#canonical-f6eb574936577fdce71b4012797b90ebde301adf2ed6e476448eccf4a4e198db)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-4d6f35e4e496aa4bff82a448b0016b6b717c939110dff109d2bf863809390467"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a9acb1ab370943e47b0e83748564c9fd673fa3a9a5d6f37ee3ffbf294759c31"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / e7c3849579e1 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7)
- equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-6e5b742afa8322f0e45bf375c0902dd4c2fb104f3496f76804f7fa66ebaa8219"></a>

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

<a id="canonical-011ac5231a1f180bb064c558f8c169b561e09a331b71da30f97f2d23400f7d55"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / e7c3849579e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d74b1372953bed5880f44c67ccf2f2c5a8c4cc580caae6a6c9cc76fc49cd198b"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / e7c3849579e1 / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-28000afe26522ced81d309e638297f1d5913848cb0f06978f7f916af9982431d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eeb267189e4c75e56993596e9c9c4472109067642165a7dc1408d3d04ecc8f51"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 5a356dd01756 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7)
- equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-c9a12c1bae303ed9dbd06b6b6636b132b9097ceb0ff9f2127d6e5ac27d98ef89"></a>

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

<a id="canonical-44118461c05b7e278e20bfee992b88c194ebab378e37393badcb3d7eb630033a"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 5a356dd01756 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6124df067602bce4bb161296195075dcb601b69874f602189cce18fb038dc50f"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / 5a356dd01756 / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-2721d20ca3e3d0acd13cfa4dc587d59da8a946eadf4611cc6a3ed457117f25a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e63b8134c9520f3763537de0ee89f0e98d27f46408669d500ec04b63b4930738"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 575854b19ab3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-962b2a8ff60dafb8ff79ceda26f6141c64064b99aa7bb58e292b552105abf9ec"></a>

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

<a id="canonical-f4cb524f1184f256b8cedf067cd8f46731bba46e6599f9abb120fd802e724657"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 575854b19ab3 / 3

<a id="canonical-c86187120ac8258bb8f70661476f144b9531a1e812771e2bd54579910b63ecd1"></a>

<a id="canonical-3b253fb0358822c607d1fb0852e559c9d4132913d95f24bb5440a4116a75468e"></a>

## dgw_address property — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 575854b19ab3 / 4

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

<a id="canonical-2e898ad561f6e77cb6f09ca8ced76d49ae78ca130e10d6ec078ebd52928f2d1f"></a>

<a id="canonical-2543974fc8357272fa8d0c457ed270d89cb4c5cbf86eb0cf5e075ceb4489d16f"></a>

## dns_address property — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 575854b19ab3 / 5

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

- [first_address](resources--securemesh_site_v2--reference--group-008.md#canonical-b6ebeabc5b7a4be416b3872936c184a31f316a9d1a362a331043bb2b4cadfa67): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-008.md#canonical-9ea3d2a0d81224e7c68202fe5b4985e61cb5df9169965aa0c11477a2012d5bb1): complete subsection reference.

<a id="canonical-8d14213e8982cbc65dd29875eb4bc39909ebea25ec4fa0f2c4e1cd5588ecbeca"></a>

<a id="canonical-b9c5129381371203376c25ad0d415fc9f9be2bfb6a462ebe02e2ab6ab4dbbe0e"></a>

## network_prefix property — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 575854b19ab3 / 6

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

<a id="canonical-79460864a16513e5cfa650f2eb68384c8bd606ce8388d0cc73db4cc5b7b3a49f"></a>

<a id="canonical-6dd9537b277b59331d4874292bd352fa6eb0ab0c0dc79154f1a6f432852d00d5"></a>

## pool_settings property — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 575854b19ab3 / 7

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

- [pools](resources--securemesh_site_v2--reference--group-008.md#canonical-ded6a9bc628e6215cc6aa4e0367bbd115a0d67b93de5f76269b9d2020b94c33e): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-008.md#canonical-2a5da2cb579d48a02f0cacafd5719193d64226bd982bdecab634269b9b94b5c7): complete subsection reference.

<a id="canonical-52b2d877ef1789f5213934f4bdebfd661d62ab6e5910082e94afd1e1aa60df5c"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 575854b19ab3 / 8

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-008.md#canonical-b6ebeabc5b7a4be416b3872936c184a31f316a9d1a362a331043bb2b4cadfa67)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-008.md#canonical-9ea3d2a0d81224e7c68202fe5b4985e61cb5df9169965aa0c11477a2012d5bb1)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-008.md#canonical-ded6a9bc628e6215cc6aa4e0367bbd115a0d67b93de5f76269b9d2020b94c33e)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-008.md#canonical-2a5da2cb579d48a02f0cacafd5719193d64226bd982bdecab634269b9b94b5c7)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-b6ebeabc5b7a4be416b3872936c184a31f316a9d1a362a331043bb2b4cadfa67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62a6db2014013d1fc1c8110ca3786185d52b717c702d5fb166886111747d8671"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_add / f4175fbc7a73 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-2721d20ca3e3d0acd13cfa4dc587d59da8a946eadf4611cc6a3ed457117f25a4)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-c26a15e1b29c27d847b49427cf71dfd635f9552341300e99f6007ff767c784d4"></a>

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

<a id="canonical-689da909dc93646c383f41e0339472be6f741eaab9995680f9ab51927a0d437e"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_add / f4175fbc7a73 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-559298cf1117409c4d79a64d51119a85ac570e70b21de8fbfd862a7a28414956"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_add / f4175fbc7a73 / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-2721d20ca3e3d0acd13cfa4dc587d59da8a946eadf4611cc6a3ed457117f25a4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9ea3d2a0d81224e7c68202fe5b4985e61cb5df9169965aa0c11477a2012d5bb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0544584ed11b50dfb221b3b103e813588526c1846c00bd323644d6d8b6ab073"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addr / 9ced37de8cac / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-2721d20ca3e3d0acd13cfa4dc587d59da8a946eadf4611cc6a3ed457117f25a4)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-2eb5f06d14243cc2f07e117bbc9fbbfe34f5346954f0e85d8e3a7486bf7b443d"></a>

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

<a id="canonical-8f2e4bdd06d798c7e1d3f4f123bdd80939a888182429cd12d30e20d159140996"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addr / 9ced37de8cac / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-845fa872e8d0bcae2a722632ca35aac4b17d424e39956531dd7b58b87d39b18d"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_addr / 9ced37de8cac / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-2721d20ca3e3d0acd13cfa4dc587d59da8a946eadf4611cc6a3ed457117f25a4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ded6a9bc628e6215cc6aa4e0367bbd115a0d67b93de5f76269b9d2020b94c33e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d3538ba70a6172176a2979a7d73e3790f6b7af1c5845bb9cd3762f0fbe1a5ec"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 9f94c86f8963 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-2721d20ca3e3d0acd13cfa4dc587d59da8a946eadf4611cc6a3ed457117f25a4)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-46966890c3c07fee9fc33149e44e415ea3c5ed79aa4f001551594193daedc7f2"></a>

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

<a id="canonical-09aa740f8859f5fe541b41a851bbfcff5d781c26ca11330fbfba223dcabaa729"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 9f94c86f8963 / 3

<a id="canonical-65841a96c4c393cb0e55d5d2dc682f4af14e08d76108f23d36a6ea42f727c44e"></a>

<a id="canonical-69719a954ca23c56e163eac991b6ed2438759950147e5c931b14df4c88a4a4c0"></a>

## end_ip property — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 9f94c86f8963 / 4

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

<a id="canonical-a75d63994f13effaf9ac335689470000331c7d54e5afc3b7852b470af1621e1a"></a>

<a id="canonical-e7ffb536b8be00059f1c3ce5dfe7ed9695480fe3230dc4f6ab8adc5348b45ad2"></a>

## exclude property — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 9f94c86f8963 / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-3f10766bdbc1662a98501e3f870fe744ff257d09830ec96a0561496f41005526"></a>

<a id="canonical-75ed853a68825295f3bc78943939238ca9459f49c37ab00c138da046c93e8206"></a>

## start_ip property — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 9f94c86f8963 / 6

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

<a id="canonical-efcc8771aa45907ed918c915147a3319fc68433bd198d883826f07d3826ac6a9"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 9f94c86f8963 / 7

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-2721d20ca3e3d0acd13cfa4dc587d59da8a946eadf4611cc6a3ed457117f25a4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-2a5da2cb579d48a02f0cacafd5719193d64226bd982bdecab634269b9b94b5c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d92ee48538254776be16928e5cbcc263d269e8fc3d1711d2751a4368a2eb1760"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_d / 36e0496d2c10 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7)
- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-2721d20ca3e3d0acd13cfa4dc587d59da8a946eadf4611cc6a3ed457117f25a4)
- equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-5c46818519732bcb104b9143921b33bf03276b6a17ab217a2fc1314a3cd2756b"></a>

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

<a id="canonical-44235445c387a576b3a0b657fb3fe4085b9b8a6c733b51184bab2255767de41f"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_d / 36e0496d2c10 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bd5535046114431da930f4b2ce22a0902b6d7e29181219689a84bc6f476e29d1"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_d / 36e0496d2c10 / 4

- [equinix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-2721d20ca3e3d0acd13cfa4dc587d59da8a946eadf4611cc6a3ed457117f25a4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f6eb574936577fdce71b4012797b90ebde301adf2ed6e476448eccf4a4e198db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f1270939b05f64964e81aac83fac6063193b5cd15c8e67ab26f9bf14db1bd25"></a>

## equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 834ad65ba9ee / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7)
- equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-9e55fe2aa214bb2856276ca617208dee4bc083d410c423673ab5a14fe2d6d835"></a>

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

<a id="canonical-2e15e94bd223a991b47bfedc6cfd1e5ec463f4e41314dbeb37bcb4cc909293b0"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 834ad65ba9ee / 3

<a id="canonical-519839c154f5fd81f654d923a9d3a1e782a07b0d9dcf5fbd8654100da47e1417"></a>

<a id="canonical-3d4f6ef8b2844184918a8d314e4047830a316e36fed01b5020dc4781be05c343"></a>

## interface_ip_map property — equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 834ad65ba9ee / 4

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

<a id="canonical-236ec6285ff9c4dd361540e499ce54fd56d34eccaffb0d1729c829ee95921843"></a>

## Next pages — equinix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / 834ad65ba9ee / 5

- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-008.md#canonical-dc2aa4a98dbf413b0d34b7915a3b2e6f206fdad70c73e3c0b06a079a9d4743d7)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-fa344d5d9a8d955c01d1e0e0c94ca03461b0ae99cf28f49e181a3978f96ee1c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-813d59fb92de6e08dd5551405ba9ac8cf42e1c0c250106438c832e08c7c47d1b"></a>

## equinix.not_managed.node_list.interface_list.ethernet_interface — equinix.not_managed.node_list.interface_list.ethernet_interface / a02292289cee / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-112f947989db730e2c1de8b1458b056012fee3d61cfd819cac7a1c3730f4444d"></a>

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

<a id="canonical-6e749e5ed83a230c818d5451b6ccf0e3abbf8010d30a70724e4de337bedaa925"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ethernet_interface / a02292289cee / 3

<a id="canonical-6fcf7bdc956a3c4263972fc81b634020a95455db91ec1143982ba7d6ea1b353c"></a>

<a id="canonical-9db575d25dc816bea6c045553559b9905dc75b083b47caf6da673f0e2426b509"></a>

## device property — equinix.not_managed.node_list.interface_list.ethernet_interface / a02292289cee / 4

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

<a id="canonical-cf22e414c63c294f19121ffe3cd827b3392908185443f2277b525eab0b0ae885"></a>

<a id="canonical-91ac27f204b849b51e47c5da82b1a40daa88670a25c54068ecc5329e21e61da2"></a>

## mac property — equinix.not_managed.node_list.interface_list.ethernet_interface / a02292289cee / 5

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

<a id="canonical-6b849c4cc6181b8fdaaf9d93109c0708a399c4a6afa5fdcc177152816f3f5daf"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ethernet_interface / a02292289cee / 6

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-062a732a1f5ddf6e39b1297ac6d630c92fd769073c2963ecafca12c7a123c6b3"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config — equinix.not_managed.node_list.interface_list.ipv6_auto_config / 7606f126a7be / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-f1f033c71c8c29ac3bf96554b6b2a5db81d2085389ab965fc6fa892c8d0a4208"></a>

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

<a id="canonical-3068d1ee5cc0d14ce8935ecfb528f90d9843c6bca03b10cb6bb32e1b7d7dfe75"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config / 7606f126a7be / 3

- [host](resources--securemesh_site_v2--reference--group-008.md#canonical-390f96f388045f23726bd8d9c0f303d01653dbd2b91a6d5311b59f731d1869ab): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5): complete subsection reference.

<a id="canonical-668b50607976d6166291c976f756cb16dd5b04107fa3624b09a91c718622c86f"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config / 7606f126a7be / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-008.md#canonical-390f96f388045f23726bd8d9c0f303d01653dbd2b91a6d5311b59f731d1869ab)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-390f96f388045f23726bd8d9c0f303d01653dbd2b91a6d5311b59f731d1869ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96663849230335c431eda7aaf7c36f7948de1766eacd33b37ed1776d8a667b4a"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.host — equinix.not_managed.node_list.interface_list.ipv6_auto_config.host / 58053cb7e440 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-a13e5e6dd7d4e3e8c95592255d3817265461a2a6485d6abfa36385d298343967"></a>

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

<a id="canonical-a0ba3439255973e0507fafae1579e66a69f44c9f3535e4587157918d882f570e"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.host / 58053cb7e440 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-74db1df9357adb1003c48407ec8da51658d5e8b3db9b9fd50007826dcc9fa239"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.host / 58053cb7e440 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-734a8559e174ba1dc7112308ccae155e92e6bd96d3a424e64a1d8c5d132bbddd"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router / 88e9492b1f3d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-21b860b59cb6901798127d8413a3328ac2b9e8971e69568bc769e67702755668"></a>

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

<a id="canonical-6c67c27eba8ea72c7830468aa3c6368d412e6777aaa3889eb7ac4699f5c804bc"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router / 88e9492b1f3d / 3

- [dns_config](resources--securemesh_site_v2--reference--group-008.md#canonical-64243a799a165f40cd98225a0246fae2eec6fc62b322bfd49e556cd266ff48aa): complete subsection reference.

<a id="canonical-d6d32806761cc4df0e4a0e19afa36fefdfa3d3bf6dfa0e50c95e6578aad55bc2"></a>

<a id="canonical-7081d505d5132b86169d43654e0a04c4e475fb47bbed3376202fc612f24796a5"></a>

## network_prefix property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router / 88e9492b1f3d / 4

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

- [stateful](resources--securemesh_site_v2--reference--group-008.md#canonical-cbea60ed74d72383917cc47841d496e8d6c9430cd7bb42490e28bcc441ac0081): complete subsection reference.

<a id="canonical-d6e8458e3457759ac0da1b02618ed8cc6549777cf999655c02e21cebd52875cc"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router / 88e9492b1f3d / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-008.md#canonical-64243a799a165f40cd98225a0246fae2eec6fc62b322bfd49e556cd266ff48aa)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-008.md#canonical-cbea60ed74d72383917cc47841d496e8d6c9430cd7bb42490e28bcc441ac0081)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-64243a799a165f40cd98225a0246fae2eec6fc62b322bfd49e556cd266ff48aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5c438ff6b31b0d16e1704b6cc94d3e4812b3baaf2c264dbfac8d13c9ab0f07c"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 30291e7352e2 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-d2ff8e0fef2e6a80609a686f0a6d3f7e2a1e433878effb77857ee3f660ede89e"></a>

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

<a id="canonical-6689809cf2018487c584f5b2ad3611efa4020a5d54fabfdb01c083237ce458eb"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 30291e7352e2 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-008.md#canonical-20836d44cf51ac8f280abb254d5e388aa9a78495bf55ba871a22d12f85fef02d): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-008.md#canonical-1e6680ca4d96fa38e874fdd3d5cb259b2b831c302c9cd2c0d95d76c883812f17): complete subsection reference.

<a id="canonical-4aa15ac5d61dd7c11943b1504828a1a8a7fc24ef9bcc2634d949639f6afd1635"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 30291e7352e2 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-008.md#canonical-20836d44cf51ac8f280abb254d5e388aa9a78495bf55ba871a22d12f85fef02d)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-008.md#canonical-1e6680ca4d96fa38e874fdd3d5cb259b2b831c302c9cd2c0d95d76c883812f17)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-20836d44cf51ac8f280abb254d5e388aa9a78495bf55ba871a22d12f85fef02d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf7cdc365ad03a67ed18fd1adbb43abcf05f61f969ec090a08fb98f31d6e29ae"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 9be705a547c4 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-008.md#canonical-64243a799a165f40cd98225a0246fae2eec6fc62b322bfd49e556cd266ff48aa)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-640c1169251e8e779f05c1aeb8eb9b78fa3e05ed7f81969bd9b051b0ee4f0ba4"></a>

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

<a id="canonical-a351e7a1bd73ca24a1bbda7eeb17b2d1b2ae4957cf0627af4198badf2dc18512"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 9be705a547c4 / 3

<a id="canonical-5d4d0cc49b624d4af6d7c9ef12e3a8c6657ec49a00230092e8c92951c987f71b"></a>

<a id="canonical-1d03e2267a10bfacd0d82bc1cef859fc039a7a1dee4c8a8a03011de9d7801ded"></a>

## dns_list property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 9be705a547c4 / 4

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

<a id="canonical-08034359c415a72985eaa7aa820a4374099239697e8581f90464238bb1737743"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 9be705a547c4 / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-008.md#canonical-64243a799a165f40cd98225a0246fae2eec6fc62b322bfd49e556cd266ff48aa)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1e6680ca4d96fa38e874fdd3d5cb259b2b831c302c9cd2c0d95d76c883812f17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e203edc6308375171b3cc74a453c8e975bc5cf84bba2bc9566927a571a087d9"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 373f4d7838ef / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-008.md#canonical-64243a799a165f40cd98225a0246fae2eec6fc62b322bfd49e556cd266ff48aa)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-76008f09baf6d6eba22d6466b6d7ba88a3bf9530d25204a05401d82221582927"></a>

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

<a id="canonical-f20d73900e7d1b671d08912ead0eff3072b5f90437bfad9b8ea15ba4c38669c4"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 373f4d7838ef / 3

<a id="canonical-25b4f885c46559049b79340002d063f081dc4995153d555df00b8316788e6a5f"></a>

<a id="canonical-d3000730625b1e58d786f8a8895096e7de9ec35d6bc74e781c9714e1c4746031"></a>

## configured_address property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 373f4d7838ef / 4

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

- [first_address](resources--securemesh_site_v2--reference--group-008.md#canonical-a40edad7b770eadb2a6f667b610dbaedb9da797c5b3efeb859884dd9fd51ecf3): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-008.md#canonical-060e80a8c62c133f55d2f2d25b86e229a97b22b1ae5a51a48e2bc3dd5c9abe07): complete subsection reference.

<a id="canonical-7eb830813d9d0207672add3fda87b4a49233a4e97490ce849d6d3dcbd75a446f"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 373f4d7838ef / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-008.md#canonical-a40edad7b770eadb2a6f667b610dbaedb9da797c5b3efeb859884dd9fd51ecf3)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-008.md#canonical-060e80a8c62c133f55d2f2d25b86e229a97b22b1ae5a51a48e2bc3dd5c9abe07)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-008.md#canonical-64243a799a165f40cd98225a0246fae2eec6fc62b322bfd49e556cd266ff48aa)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-a40edad7b770eadb2a6f667b610dbaedb9da797c5b3efeb859884dd9fd51ecf3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff5f450fe2cdea28d3e03c6f04d1324e26d2112d4c00274ca32e7df84864b144"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / dbce542c2896 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-008.md#canonical-64243a799a165f40cd98225a0246fae2eec6fc62b322bfd49e556cd266ff48aa)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-008.md#canonical-1e6680ca4d96fa38e874fdd3d5cb259b2b831c302c9cd2c0d95d76c883812f17)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-2d62d009928eba156adb5d52d70ffe5b4aec8d3aa37555177da60a55097f9c87"></a>

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

<a id="canonical-a38c82751562a7e84b32ca77bcab52f2ef79862f49a03a136636cdfee14b4865"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / dbce542c2896 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6792c3ac97cc398a1c15e8229cb0e01f46977c6ee9b5c1f5a4f3d69de2b865bd"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / dbce542c2896 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-008.md#canonical-1e6680ca4d96fa38e874fdd3d5cb259b2b831c302c9cd2c0d95d76c883812f17)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-060e80a8c62c133f55d2f2d25b86e229a97b22b1ae5a51a48e2bc3dd5c9abe07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-492c1df35592262c4647c39a7947a92eb4a5a0cff23e72715b9ecda2828917ba"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 79ff3c8e1158 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-008.md#canonical-64243a799a165f40cd98225a0246fae2eec6fc62b322bfd49e556cd266ff48aa)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-008.md#canonical-1e6680ca4d96fa38e874fdd3d5cb259b2b831c302c9cd2c0d95d76c883812f17)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-12aea33a3004233ed1f225fa37faee15d2d00a5e559f43797009870dfd2632d9"></a>

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

<a id="canonical-aa3a32e1587a227c2975a41e6a28df60b8f7daf7bb95b6bdd950930068f70432"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 79ff3c8e1158 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a89289deaef86f0f364affdec58784efe9766ac9a8ab3e62aef2da8e303d241a"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config. / 79ff3c8e1158 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-008.md#canonical-1e6680ca4d96fa38e874fdd3d5cb259b2b831c302c9cd2c0d95d76c883812f17)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-cbea60ed74d72383917cc47841d496e8d6c9430cd7bb42490e28bcc441ac0081"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6ebd9a10d672b32ef394f72525cd3433f4cf5ab76335b3318c6f29653e828f0"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 564f137d3196 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-2054f6caf3398815c844166592786c39fda6fa0ebd99f60c4ae4a1089fc90e35"></a>

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

<a id="canonical-3e181ca8203b1cb759eed9e9a87722131a7bbe01f37a77f3a76bf0064ec3c6fd"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 564f137d3196 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-008.md#canonical-7d70bf0a37236aeae815b16c5754f6f83c5b64a92dad8c4110fd2cca57ba30b1): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-008.md#canonical-f3a05989fba3b4d2f1af92e0ab65ea94ec5e954190b9f42cdddfc4c25f4c5562): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-5948b425d72cd4c5cdc03867645116417c4971439c13f1ef39d4f94149e3940e): complete subsection reference.

<a id="canonical-88ec2f07c05de3fdece97b544f4977a83f9c6e4ccf38e7794ea6cb77c3b6d1df"></a>

<a id="canonical-eccf05edd389ac821acb3c07ac51ad2033c9b316fe668cc89f89d42be72e6c1f"></a>

## fixed_ip_map property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 564f137d3196 / 4

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-009.md#canonical-b00bb0bf541e65537eba77487f37944348350895c53af02d6b163aecd82724e7): complete subsection reference.

<a id="canonical-160eef3ee3316139d53cb49993b810533b1cce077af42ec612702f395856781a"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 564f137d3196 / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-008.md#canonical-7d70bf0a37236aeae815b16c5754f6f83c5b64a92dad8c4110fd2cca57ba30b1)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-008.md#canonical-f3a05989fba3b4d2f1af92e0ab65ea94ec5e954190b9f42cdddfc4c25f4c5562)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-5948b425d72cd4c5cdc03867645116417c4971439c13f1ef39d4f94149e3940e)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-009.md#canonical-b00bb0bf541e65537eba77487f37944348350895c53af02d6b163aecd82724e7)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-7d70bf0a37236aeae815b16c5754f6f83c5b64a92dad8c4110fd2cca57ba30b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-131fd3d1d5454d915d2d4d77222bb939a0366e848759f425a4eddfe37821bbb4"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 906f7cb1ae51 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-008.md#canonical-cbea60ed74d72383917cc47841d496e8d6c9430cd7bb42490e28bcc441ac0081)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-1828022422beb2530eb2e883880ac6e8fa021af1d9ccd46ed18c97c678f3b476"></a>

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

<a id="canonical-41153c1a8570b495943edeb3c5c47958520831eb4847746b7ed2b421925d8f82"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 906f7cb1ae51 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-908b449e2bd98c6470b68a85ffb5e7c60336f6a031786fb6716328ff69cb135c"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 906f7cb1ae51 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-008.md#canonical-cbea60ed74d72383917cc47841d496e8d6c9430cd7bb42490e28bcc441ac0081)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f3a05989fba3b4d2f1af92e0ab65ea94ec5e954190b9f42cdddfc4c25f4c5562"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39699900264503baf1def2e7412cc2908225392c615c9da6c368aacb369e2176"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 58d623f01b7f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-008.md#canonical-cbea60ed74d72383917cc47841d496e8d6c9430cd7bb42490e28bcc441ac0081)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-721780f8613b0259eacc4e531c58c8231434cbb43abf1440f9b60783505f95bf"></a>

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

<a id="canonical-287cd422d3c4da1539297c967f94caaf1262c23aae533425fa5fc849686b953a"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 58d623f01b7f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-af05f16830bf9f37481ca07fe04327c2cda581f813c338b5eecefcf60d0a7742"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.au / 58d623f01b7f / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-008.md#canonical-cbea60ed74d72383917cc47841d496e8d6c9430cd7bb42490e28bcc441ac0081)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5948b425d72cd4c5cdc03867645116417c4971439c13f1ef39d4f94149e3940e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
