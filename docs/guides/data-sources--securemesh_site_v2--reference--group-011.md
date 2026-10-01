---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-229a3eca82bdcfbec9d34f31387f33e72b815c9978933e644a33b44e325d2ab5"></a>

## kvm.not_managed.node_list.interface_list.monitor_disabled — kvm.not_managed.node_list.interface_list.monitor_disabled / c57259af7629 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-671e616ce176e8a5200edfede43edb42df057387de83e93fc1735eb82d31614e)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-04f6a372f08b7175e8445bffc2db3eb26e20e4f140f83e65c802fefc3e800891)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-fd6f11f31db03235a76c4418e02fcf7235c56c56f4704f4e67db9f264fa465c0)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- kvm.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-6ca40e3b553618dd0bc524994b12ea02433fcfdf7adff08fd47b8ec495b8996a"></a>

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

<a id="canonical-a8898288056530b0fa2dea6deb1262b1c4623234577ea47b80584d73c49fb4ca"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.monitor_disabled / c57259af7629 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0fff2c8445f7e4887df2ad8edaf86b4d2cf20f4997057063b6e7fdd932cecb44"></a>

## Next pages — kvm.not_managed.node_list.interface_list.monitor_disabled / c57259af7629 / 4

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a2445faf25d2361234fb597ab1d9200ee16fa1fd2bf1899c228f173cf25b2022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-241bccd9735f42451cd5a87ebeb23c9ec552561f8f47a41b8623370dfb79cfe4"></a>

## kvm.not_managed.node_list.interface_list.network_option — kvm.not_managed.node_list.interface_list.network_option / 989a5404669e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-671e616ce176e8a5200edfede43edb42df057387de83e93fc1735eb82d31614e)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-04f6a372f08b7175e8445bffc2db3eb26e20e4f140f83e65c802fefc3e800891)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-fd6f11f31db03235a76c4418e02fcf7235c56c56f4704f4e67db9f264fa465c0)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- kvm.not_managed.node_list.interface_list.network_option

<a id="canonical-416d508ae77e6e0f590b234bbdc62ff4d49992640d3e0e088347dc5897df974c"></a>

Type: `"single"`. Computed.

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

<a id="canonical-960723ef0b10cdefbb2406e3a3ee8756dc2ce3b9d15e56d709e935f05e1f2b6d"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.network_option / 989a5404669e / 3

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c52494eb2a12a4f0caca87c920d2c52edd4979b2b7763b103d9bdf840efc9f27): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-011.md#canonical-a913742bfcea082d6e586da017c1b4ab6d153f53a3be0d5d01dfcae36ca1f465): complete subsection reference.

<a id="canonical-49a5613896e97602d12578be4133988a73bd732ec78fb890c18125bfbeb79d32"></a>

## Next pages — kvm.not_managed.node_list.interface_list.network_option / 989a5404669e / 4

- [kvm.not_managed.node_list.interface_list.network_option.site_local_inside_network](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c52494eb2a12a4f0caca87c920d2c52edd4979b2b7763b103d9bdf840efc9f27)
- [kvm.not_managed.node_list.interface_list.network_option.site_local_network](data-sources--securemesh_site_v2--reference--group-011.md#canonical-a913742bfcea082d6e586da017c1b4ab6d153f53a3be0d5d01dfcae36ca1f465)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c52494eb2a12a4f0caca87c920d2c52edd4979b2b7763b103d9bdf840efc9f27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-169c4316e79325a2d407039cf6272244e0a3e04c37ed5d60093ae739a51fb889"></a>

## kvm.not_managed.node_list.interface_list.network_option.site_local_inside_network — kvm.not_managed.node_list.interface_list.network_option.site_local_inside_networ / eb5ed22f86f7 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-671e616ce176e8a5200edfede43edb42df057387de83e93fc1735eb82d31614e)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-04f6a372f08b7175e8445bffc2db3eb26e20e4f140f83e65c802fefc3e800891)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-fd6f11f31db03235a76c4418e02fcf7235c56c56f4704f4e67db9f264fa465c0)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- [kvm.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-011.md#canonical-a2445faf25d2361234fb597ab1d9200ee16fa1fd2bf1899c228f173cf25b2022)
- kvm.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-9aa3cd7bfa45ebb882d5ef8e079bde5cf2b3837bc8abf8329fc372a9cd907a9b"></a>

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

<a id="canonical-0264bbfbfaefe7a01ba7532f8c45567a5ef3a86dada3610d83fe6d73ce02e03c"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.network_option.site_local_inside_networ / eb5ed22f86f7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-71e04a05214caf2f7bab6c47aa24b76293aa14a5bb713bfc341f1887b4847d5c"></a>

## Next pages — kvm.not_managed.node_list.interface_list.network_option.site_local_inside_networ / eb5ed22f86f7 / 4

- [kvm.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-011.md#canonical-a2445faf25d2361234fb597ab1d9200ee16fa1fd2bf1899c228f173cf25b2022)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a913742bfcea082d6e586da017c1b4ab6d153f53a3be0d5d01dfcae36ca1f465"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae0eeac7456a7f9387a70a3009647e2e31542fdc9664eb2164d5cb40a827accc"></a>

## kvm.not_managed.node_list.interface_list.network_option.site_local_network — kvm.not_managed.node_list.interface_list.network_option.site_local_network / eb2298b7aba0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-671e616ce176e8a5200edfede43edb42df057387de83e93fc1735eb82d31614e)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-04f6a372f08b7175e8445bffc2db3eb26e20e4f140f83e65c802fefc3e800891)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-fd6f11f31db03235a76c4418e02fcf7235c56c56f4704f4e67db9f264fa465c0)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- [kvm.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-011.md#canonical-a2445faf25d2361234fb597ab1d9200ee16fa1fd2bf1899c228f173cf25b2022)
- kvm.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-36b92964dd89da8ba9051b8ec509dd12cd08938c1595e3fe46025db7c81ad1e4"></a>

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

<a id="canonical-2b6d68e6c2f43f4491b54a41f175f47aabdbbf7464f9abaf2d42c9630455977d"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.network_option.site_local_network / eb2298b7aba0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-137430d816046451585c24ddc61393166bb188aa7479ee35b233ded16bf8a94d"></a>

## Next pages — kvm.not_managed.node_list.interface_list.network_option.site_local_network / eb2298b7aba0 / 4

- [kvm.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-011.md#canonical-a2445faf25d2361234fb597ab1d9200ee16fa1fd2bf1899c228f173cf25b2022)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-88c0b8792683cc7601f078f36f586d2cdd73b96cae7a00d35001193a2c536342"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2dea437b30e41f59a462f6028e779b05cf9d6693bdc842b59aadf99517639f6a"></a>

## kvm.not_managed.node_list.interface_list.no_ipv4_address — kvm.not_managed.node_list.interface_list.no_ipv4_address / 48756121ce40 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-671e616ce176e8a5200edfede43edb42df057387de83e93fc1735eb82d31614e)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-04f6a372f08b7175e8445bffc2db3eb26e20e4f140f83e65c802fefc3e800891)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-fd6f11f31db03235a76c4418e02fcf7235c56c56f4704f4e67db9f264fa465c0)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- kvm.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-77fc186502475fbcbb9296ae953f3b622d77cde9b8234be474c69d11172a48fc"></a>

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

<a id="canonical-0b60c19e86f53c89da411e10015dcd00c7aa969034f17739c17034ab2f988d7a"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.no_ipv4_address / 48756121ce40 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e99436054126baa04285b24b8519a41ce5a07f9ac8f1e01e52c32641b3dd62b3"></a>

## Next pages — kvm.not_managed.node_list.interface_list.no_ipv4_address / 48756121ce40 / 4

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-4f2dcd9b9aad63572f8ea2e5c6f5de13f7cf3c812da09373b32d20897770c5fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-967f87be97e594d1df6129ddf7ac89eaf325a9a1eae25e8a92f16dbf0c1bfa56"></a>

## kvm.not_managed.node_list.interface_list.no_ipv6_address — kvm.not_managed.node_list.interface_list.no_ipv6_address / 5ef9cc42122c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-671e616ce176e8a5200edfede43edb42df057387de83e93fc1735eb82d31614e)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-04f6a372f08b7175e8445bffc2db3eb26e20e4f140f83e65c802fefc3e800891)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-fd6f11f31db03235a76c4418e02fcf7235c56c56f4704f4e67db9f264fa465c0)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- kvm.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-e5db9ff94d394a37a8497aeb6d380072ae55a81cf85d48a7985a080b84a85cc4"></a>

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

<a id="canonical-74fd24745bf23deb859290f9173f6002d7d311b637bbe89d97883e71e2f56caa"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.no_ipv6_address / 5ef9cc42122c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4089efe0e7d986768bb4c71538be15115ddbf7c332bd5c40c06906a8c132e582"></a>

## Next pages — kvm.not_managed.node_list.interface_list.no_ipv6_address / 5ef9cc42122c / 4

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-423169b592e389a47bd79ccbab6833d780fbd50f10d4baca8e9a566682feeae5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9458f8abb940d067ee01135e6ee90f763cc5f7424999c8aa6c88a19ae9a1c6ad"></a>

## kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / b593713cf056 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-671e616ce176e8a5200edfede43edb42df057387de83e93fc1735eb82d31614e)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-04f6a372f08b7175e8445bffc2db3eb26e20e4f140f83e65c802fefc3e800891)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-fd6f11f31db03235a76c4418e02fcf7235c56c56f4704f4e67db9f264fa465c0)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-e6950c1879b73f336e64e36cde0937446b3a6429e512011f0e25579c487721c7"></a>

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

<a id="canonical-07f65c6522e5d2afc4f3fbd623f54aae4873febdd95425e3ef4e4b06707bd57f"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / b593713cf056 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2037b0de7d4ce9f78922f1682a782b14b339fd2a7e8992bb7e66bb3780077690"></a>

## Next pages — kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / b593713cf056 / 4

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-4241e9a794f52d7cb0e067a6a98a069df2ac154e12d80e6b2c53981e09b84689"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9264bd842242a1f70a4327b4aa5d3b35136c9884d03ee63a2ca16c64c497c2a7"></a>

## kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ena / cd80a4aeae9e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-671e616ce176e8a5200edfede43edb42df057387de83e93fc1735eb82d31614e)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-04f6a372f08b7175e8445bffc2db3eb26e20e4f140f83e65c802fefc3e800891)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-fd6f11f31db03235a76c4418e02fcf7235c56c56f4704f4e67db9f264fa465c0)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-eaa1fad5cdef38ed69ddfb5bcbe8306c0f81b7c4a84018a7d6d2547330a8fde8"></a>

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

<a id="canonical-073b09ddd811ce6fae0c2c14bf622b42689e96db00bb654707f7fe4cea7ca2b2"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ena / cd80a4aeae9e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f6c735330754a607f29c6ad163afb593fb8151c01a2a042acf467bb4d32e5257"></a>

## Next pages — kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ena / cd80a4aeae9e / 4

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-aa9b8a78722225cae0d3c7ee406483d9d74292b0674e097ad5ec2caabda110bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d163e2da6ac9fa576102372aac291a12428b3a1bed3d40de464ea6916fc2838"></a>

## kvm.not_managed.node_list.interface_list.static_ip — kvm.not_managed.node_list.interface_list.static_ip / 92ac4be73372 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-671e616ce176e8a5200edfede43edb42df057387de83e93fc1735eb82d31614e)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-04f6a372f08b7175e8445bffc2db3eb26e20e4f140f83e65c802fefc3e800891)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-fd6f11f31db03235a76c4418e02fcf7235c56c56f4704f4e67db9f264fa465c0)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- kvm.not_managed.node_list.interface_list.static_ip

<a id="canonical-05408dd3e1f64c05985a1a5760d75c1a66062fea518e74c13d1298514819ce19"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f80cd78b0998e6101bf7f8345b7cb7b892f348d8ac8feb90bf91efed245768ec"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.static_ip / 92ac4be73372 / 3

<a id="canonical-a2f1758c05c65bd6960dd605e2c699b5199e0b52cc9dfabbfd54763da0943986"></a>

<a id="canonical-ebb87fca5edca461b1cef83852a7fc93eed0a247d14cb84ae4d5315bb843bd40"></a>

## default_gw property — kvm.not_managed.node_list.interface_list.static_ip / 92ac4be73372 / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

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

<a id="canonical-95cc05737a80b1520445dacf41d234586392664eced2db8554f5b6e0c073fa7d"></a>

<a id="canonical-61378de606189f67ccb5c6d0e40b57fa78f07e90d60da5b18fec447fb05066b9"></a>

## dns_server property — kvm.not_managed.node_list.interface_list.static_ip / 92ac4be73372 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-49af8239f4ea7f25c5e210ebdffb21734c18e33b5989a31deb9fa7a3a96f3eff"></a>

<a id="canonical-aa0e3a770f905ed5656211e557f0e388d783e05f0647a6eef212b12e3c81efa7"></a>

## ip_address property — kvm.not_managed.node_list.interface_list.static_ip / 92ac4be73372 / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

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

<a id="canonical-91c738b2c25cb9ffa13c8c02555e2fe65ca18424926da09439b79f319828f3c3"></a>

## Next pages — kvm.not_managed.node_list.interface_list.static_ip / 92ac4be73372 / 7

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-2763117f5cfc5c5f6ab3085a0a2823cbe18d309bb07e0a03ac85d0d11a4936ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afba71ada2e5aef6914ba64f098c02d0c96b4d7715b49a2b898930480aa60e97"></a>

## kvm.not_managed.node_list.interface_list.static_ipv6_address — kvm.not_managed.node_list.interface_list.static_ipv6_address / 29e06029c7a0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-671e616ce176e8a5200edfede43edb42df057387de83e93fc1735eb82d31614e)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-04f6a372f08b7175e8445bffc2db3eb26e20e4f140f83e65c802fefc3e800891)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-fd6f11f31db03235a76c4418e02fcf7235c56c56f4704f4e67db9f264fa465c0)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- kvm.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-2978928565fe2810659d95708f15396b94cd202005c75add990de777941c7822"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

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

<a id="canonical-4c7a812c68bfe91cd06aaa58a9ab0c3bc0b13546340fbb414d2b0333d520ebc5"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.static_ipv6_address / 29e06029c7a0 / 3

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-011.md#canonical-32b8f5753e1975b2fb02e46c6baef726da19aa5ffb5a59242f1c7998464dd9f8): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-011.md#canonical-8dfd47fbc44226fd766005caaca1a399718f96b141d4c1489b6ab07b21bd13ae): complete subsection reference.

<a id="canonical-40e3e7204bd16fc30bb29199656f74877c610a27ce902c0a3f3336b0cb41c929"></a>

## Next pages — kvm.not_managed.node_list.interface_list.static_ipv6_address / 29e06029c7a0 / 4

- [kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site_v2--reference--group-011.md#canonical-32b8f5753e1975b2fb02e46c6baef726da19aa5ffb5a59242f1c7998464dd9f8)
- [kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](data-sources--securemesh_site_v2--reference--group-011.md#canonical-8dfd47fbc44226fd766005caaca1a399718f96b141d4c1489b6ab07b21bd13ae)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-32b8f5753e1975b2fb02e46c6baef726da19aa5ffb5a59242f1c7998464dd9f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fff832a101c6d8aa9403d65cc79468c8dc78889e702b332bec0da86e55f65159"></a>

## kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / 86f37e810244 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-671e616ce176e8a5200edfede43edb42df057387de83e93fc1735eb82d31614e)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-04f6a372f08b7175e8445bffc2db3eb26e20e4f140f83e65c802fefc3e800891)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-fd6f11f31db03235a76c4418e02fcf7235c56c56f4704f4e67db9f264fa465c0)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- [kvm.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2763117f5cfc5c5f6ab3085a0a2823cbe18d309bb07e0a03ac85d0d11a4936ca)
- kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-6d35ac58664d13b9f82240884ef5e71f1b3e103b3260214dd963a5293397fbb8"></a>

Type: `"single"`. Computed.

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

<a id="canonical-054a37fd8dd494be96e356f824f73331a1bab96b96c726a463c2c59f5057ff52"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / 86f37e810244 / 3

<a id="canonical-f4bff698ab50f439c49edf3f5dc7c5f436d9f25a9213701e8ed1ee2c2aaa0482"></a>

<a id="canonical-55f573ea62f8719ea8f0a114440e0575400a11ce26d5445d3e764d12b3adc24b"></a>

## interface_ip_map property — kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / 86f37e810244 / 4

Type: `["map", "string"]`. Computed.

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

<a id="canonical-63075e227c0b3d0ed641dad5719dd60285359c97b0ab49361cb1ba1c7d9955ab"></a>

## Next pages — kvm.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / 86f37e810244 / 5

- [kvm.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2763117f5cfc5c5f6ab3085a0a2823cbe18d309bb07e0a03ac85d0d11a4936ca)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8dfd47fbc44226fd766005caaca1a399718f96b141d4c1489b6ab07b21bd13ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0e8c7f522c4970b383cf74c92d60d30794894e0aaa219e53d93cd25dae238c9"></a>

## kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 0c0d44eb4c8b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-671e616ce176e8a5200edfede43edb42df057387de83e93fc1735eb82d31614e)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-04f6a372f08b7175e8445bffc2db3eb26e20e4f140f83e65c802fefc3e800891)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-fd6f11f31db03235a76c4418e02fcf7235c56c56f4704f4e67db9f264fa465c0)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- [kvm.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2763117f5cfc5c5f6ab3085a0a2823cbe18d309bb07e0a03ac85d0d11a4936ca)
- kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-1beae9c858a3f70bde0b80b07a8852cb8eb864d24859f6112a557e1ae4ad3058"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4c6b92cd95a000a781d4fdf4fafeee8009db5ddc3ee2ce4cc031f6375fc5a072"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 0c0d44eb4c8b / 3

<a id="canonical-b38e29be57a9e7b2924952d6a7415070c8e05a1302861c8891c802c7f5a4f6cb"></a>

<a id="canonical-643a94cf25c9f9c47b1de6d43e07bb7af8e4c827e9a692788b1bcdfb27a157bb"></a>

## default_gw property — kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 0c0d44eb4c8b / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

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

<a id="canonical-2b133d4dc20909e7bac61e79a4bfd92e76ba67da3e20672789055509aa28b102"></a>

<a id="canonical-1dcdbaa47c6e275bcc3dd4bc40d3401259f5009184d314a9c6cffc5297f985b5"></a>

## dns_server property — kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 0c0d44eb4c8b / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-f87d2af609a5a083dc59bace4a472eb0c6f2bf6e70b2755e59cc703e8ee5d479"></a>

<a id="canonical-37f33a13607abcbcb9d4ffc44df9bc048d1373ab81d3c5c695cbdd6522e19647"></a>

## ip_address property — kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 0c0d44eb4c8b / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

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

<a id="canonical-1aa6b00235ca9ed9f1b8485c35a981fa386aa077f140c83e3960ae6fa1d5fcdb"></a>

## Next pages — kvm.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 0c0d44eb4c8b / 7

- [kvm.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2763117f5cfc5c5f6ab3085a0a2823cbe18d309bb07e0a03ac85d0d11a4936ca)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-7e19993526595a754c1d6e4c64912106f1b715eaab0bca8ca4bf63cb8b2668af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-605aa3965f2b19f0a8d108363fcc6851bedb7ad82b6e4a9fdb53a848cca3884d"></a>

## kvm.not_managed.node_list.interface_list.vlan_interface — kvm.not_managed.node_list.interface_list.vlan_interface / fbbfef8d8980 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-671e616ce176e8a5200edfede43edb42df057387de83e93fc1735eb82d31614e)
- [kvm.not_managed](data-sources--securemesh_site_v2--reference--group-010.md#canonical-04f6a372f08b7175e8445bffc2db3eb26e20e4f140f83e65c802fefc3e800891)
- [kvm.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-fd6f11f31db03235a76c4418e02fcf7235c56c56f4704f4e67db9f264fa465c0)
- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- kvm.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-96684613f37bdcb15f5de1da942774b553879b5312641ab6ad39cbc23402d91f"></a>

Type: `"single"`. Computed.

Configuration parameter for vlan interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-59dd28e2f547c9bd4490f1b24153158df716c3f4aeb93dd103884a891adb4b50"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.vlan_interface / fbbfef8d8980 / 3

<a id="canonical-c999ff8595595ab9f925e36e62bd7ce58757d80876e76e092e0fb6d545b25168"></a>

<a id="canonical-2d34ec4cee31dc7829a4dbaa65b2d0fe76390cd7dad09eb6c5894f0904a4db16"></a>

## device property — kvm.not_managed.node_list.interface_list.vlan_interface / fbbfef8d8980 / 4

Type: `"string"`. Computed.

Select a parent interface from the dropdown.

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

<a id="canonical-901333460cedc8af6d8f4fd1815d6a8f707d842703947e73dd88783dc7ae0a72"></a>

<a id="canonical-4a1917f182db2b93c69d7988d59136d4f60f4f20bad6772e6d4c93958adffa3c"></a>

## vlan_id property — kvm.not_managed.node_list.interface_list.vlan_interface / fbbfef8d8980 / 5

Type: `"number"`. Computed.

Configure the VLAN tag for this interface.

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

<a id="canonical-da6bf07d104c73c512cd21ad25ad0844fdb5423e6cb08d9f788dbcfb1c4561bd"></a>

## Next pages — kvm.not_managed.node_list.interface_list.vlan_interface / fbbfef8d8980 / 6

- [kvm.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-010.md#canonical-14d07af97ecce8f925110825a578529ef32a83109a8ecdcda224053801ddd4f9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-19ad5681f96657661c8c290f68e21059592b727992303e34d6da3f181f75a994"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f6a27e709eaa8b0b4fabbf86d8ff4a7c0193f53e68d4dc34e5bdb9f7bf014f4"></a>

## load_balancing — load_balancing / c3299b3520dc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- load_balancing

<a id="canonical-4f1723583a02a56db729c517d0b790fffc1e56ec5f76826ffb13aeb4da31ea35"></a>

Type: `"single"`. Computed.

Section contains settings on the site that relate to Load Balancing functionality.

Upstream description:

This section contains settings on the site that relate to Load Balancing functionality.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-09c0fe5adfba6ee45808e5f4dec8e917864737074bc2003442968d95e424ce2a"></a>

## Direct properties — load_balancing / c3299b3520dc / 3

<a id="canonical-3cfc061beafb1983ef8d9f21783eab68d47dc30999831477a18dfc7d03cd3230"></a>

<a id="canonical-89e990caaa6d2e18793cb5be974d35a10b93d42644bb4a33119ec35d51db3bf9"></a>

## vip_vrrp_mode property — load_balancing / c3299b3520dc / 4

Type: `"string"`. Computed.

\[Enum: VIP\_VRRP\_INVALID|VIP\_VRRP\_ENABLE|VIP\_VRRP\_DISABLE\] VRRP advertisement mode for VIP
Invalid VRRP mode. Possible values are \`VIP\_VRRP\_INVALID\`, \`VIP\_VRRP\_ENABLE\`,
\`VIP\_VRRP\_DISABLE\`. Defaults to \`VIP\_VRRP\_INVALID\`.

Upstream description:

VRRP advertisement mode for VIP

Invalid VRRP mode.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIP_VRRP_INVALID",
  "enum": [
    "VIP_VRRP_INVALID",
    "VIP_VRRP_ENABLE",
    "VIP_VRRP_DISABLE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6080782df9b62e1fbf9956ba40907dc9436254bd8ee5d8a953ccaed6400d0177"></a>

## Next pages — load_balancing / c3299b3520dc / 5

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f46300387495682b29e0ca5fc13e02e5603b6a3614aaf9a88f731fd74ad3b1d"></a>

## local_vrf — local_vrf / 5b195ecc1301 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- local_vrf

<a id="canonical-843f0546d8f3cbdf3df9b6e4d2f876bba7ce650ac73cd0630e177c369a03029b"></a>

Type: `"single"`. Computed.

There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect
WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All
sites are required to have an SLO local VRF.

Upstream description:

There can be two local VRFs on each site. The Site Local Outside (SLO) local VRF is used to connect
WAN side workloads to this site and to connect the site to F5 Distributed Cloud for management. All
sites are required to have an SLO local VRF. The Site Local Inside (SLI) local VRF is used to
connect LAN side workloads to this site. SLI local VRF is optional.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-sli_choice": "[\"default_sli_config\",\"sli_config\"]",
  "x-ves-oneof-field-slo_choice": "[\"default_config\",\"slo_config\"]"
}
```

<a id="canonical-10023957eda93a0289a4a38c012ca02f889483578f162fb475d9385a41af7796"></a>

## Direct properties — local_vrf / 5b195ecc1301 / 3

- [default_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-769a7421e634893db92e018858d6018c095a0dd98c250ed317554acca92102a4): complete subsection reference.

- [default_sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0c315e97e27f77c66a39fd740313f5b0e1da4be36c6703fd88d50babbd62d5ef): complete subsection reference.

- [sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a): complete subsection reference.

- [slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870): complete subsection reference.

<a id="canonical-fdc98dee51c905a56a28f3a92dd83d1c2f58e53f150d12286080f63b4fca75ea"></a>

## Next pages — local_vrf / 5b195ecc1301 / 4

- [local_vrf.default_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-769a7421e634893db92e018858d6018c095a0dd98c250ed317554acca92102a4)
- [local_vrf.default_sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0c315e97e27f77c66a39fd740313f5b0e1da4be36c6703fd88d50babbd62d5ef)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-769a7421e634893db92e018858d6018c095a0dd98c250ed317554acca92102a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29125c2f0c076c9f31840405f964d8945b08bc15a2aeb6a38b650a680a6b45d0"></a>

## local_vrf.default_config — local_vrf.default_config / 7fed9974275c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- local_vrf.default_config

<a id="canonical-bc1af7e2c136a20076a80f4ab47e33194bf6f245d869e150811595f5ecee30b3"></a>

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

<a id="canonical-f9e7d57692bab9def6ef3adbd35058ab49bf13cff7d70daf513b0371cb7bd00c"></a>

## Direct properties — local_vrf.default_config / 7fed9974275c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9f2d4d0ce020092dd6f588118c738cdb1a861a01ff1b026eabe37b9cdf0d335a"></a>

## Next pages — local_vrf.default_config / 7fed9974275c / 4

- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0c315e97e27f77c66a39fd740313f5b0e1da4be36c6703fd88d50babbd62d5ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5c1972ecc06e34db2353dccc012c58d896c968df47a564e1712c6bd6e15b816"></a>

## local_vrf.default_sli_config — local_vrf.default_sli_config / fb8c6a3aa21f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- local_vrf.default_sli_config

<a id="canonical-8d1f5b093b696e46c3ef791294b91ef2aeee14a84bbd766e4564eea4fbf81de7"></a>

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

<a id="canonical-80866f12d67796e288efc0820ab971ca1dcdd2975acbed733957c6d520b86779"></a>

## Direct properties — local_vrf.default_sli_config / fb8c6a3aa21f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eaa907a922d70d6a71daf05b535fae5e486c85c97e97ebf3bb6143636d2f1d06"></a>

## Next pages — local_vrf.default_sli_config / fb8c6a3aa21f / 4

- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c2de8b268a6d75a1d332aea67ca63c9fb7601cdb64a1ead8595790d12adc4ab"></a>

## local_vrf.sli_config — local_vrf.sli_config / e2fb6a830188 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- local_vrf.sli_config

<a id="canonical-e85c7bf2564d3f149abbd7670b6caa9b248c2d13d99d56fa0c92228d920723b9"></a>

Type: `"single"`. Computed.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

<a id="canonical-a7c445087b730682b4613c438d2a15e0fdafc94aa0138115b17e6034e32640de"></a>

## Direct properties — local_vrf.sli_config / e2fb6a830188 / 3

<a id="canonical-287f00f275ffcc669db5a85b319440f8d1c7864d9e78a0ebf49b04e9243c1873"></a>

<a id="canonical-60d51bcde0af51847279d59bc5f1eebb0bec2a3b736ec45589e6844aae146dd8"></a>

## labels property — local_vrf.sli_config / e2fb6a830188 / 4

Type: `["map", "string"]`. Computed.

Add Labels for this network, these labels can be used in firewall policy.

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

<a id="canonical-eec501e760edef68458a401451efb1c2bb298c8cb9842926cdce61976b73bd53"></a>

<a id="canonical-78121b75bbbfc992a3420e2c421e63b09e2103424e21ae19b85ba836b882dee9"></a>

## nameserver property — local_vrf.sli_config / e2fb6a830188 / 5

Type: `"string"`. Computed.

Optional IPv4 DNS server to be used for name resolution.

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

- [no_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-47344ee0f760271973fd9051749a2a041faa494653e9ffa152ff325975e02e88): complete subsection reference.

- [no_v6_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-8237b382c33ffeeb8bf28b3da0408daa323052008d4689166bddd66ac6cd1837): complete subsection reference.

<a id="canonical-aed184821395036d055bc83b2d53c2c47ab585b907d77207f361593bde231595"></a>

<a id="canonical-79c544de8212e8b80346ef912b24d8fe71751207790be70c4f7221b6d198112d"></a>

## secondary_nameserver property — local_vrf.sli_config / e2fb6a830188 / 6

Type: `"string"`. Computed.

Optional Secondary IPv4 DNS server to be used for name resolution.

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

- [static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-93ea8120d47ebe43795f1bc7631e5287e64f1d51aedcc41ca54172ebd8a1278f): complete subsection reference.

- [static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3704e98825faca68e00e9b9fc724f1062824ae979806ea5ae7f6664216dfab06): complete subsection reference.

<a id="canonical-68fce12ab98c7bd9092ee1b5499d6b44c6603a4214afe6f48d5d8e53498c7d46"></a>

<a id="canonical-77a5429693d2cc4ad86b8a982c6a66a74a4350a4e6e9917da240ece45e72eac6"></a>

## vip property — local_vrf.sli_config / e2fb6a830188 / 7

Type: `"string"`. Computed.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

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

<a id="canonical-840e908aeac258bc800aec8d2d88680ab36396f2c7962bcd6183547a0dda8100"></a>

## Next pages — local_vrf.sli_config / e2fb6a830188 / 8

- [local_vrf.sli_config.no_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-47344ee0f760271973fd9051749a2a041faa494653e9ffa152ff325975e02e88)
- [local_vrf.sli_config.no_v6_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-8237b382c33ffeeb8bf28b3da0408daa323052008d4689166bddd66ac6cd1837)
- [local_vrf.sli_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-93ea8120d47ebe43795f1bc7631e5287e64f1d51aedcc41ca54172ebd8a1278f)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3704e98825faca68e00e9b9fc724f1062824ae979806ea5ae7f6664216dfab06)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-47344ee0f760271973fd9051749a2a041faa494653e9ffa152ff325975e02e88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbf55151cdfe1185e5ae818f2f97aa7b1f7de145f6fee631ec15d62b99ac93e3"></a>

## local_vrf.sli_config.no_static_routes — local_vrf.sli_config.no_static_routes / 2d7cf07643ba / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- local_vrf.sli_config.no_static_routes

<a id="canonical-0f27f3a8816f86a1e4aa285d62ec3cf7186fab41379f6a5d04087066ec950216"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no static routes.

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

<a id="canonical-7d6b5a42035fdc13b5ee3c319a5345f4774c8d22a412ec930705d7ab78cb6519"></a>

## Direct properties — local_vrf.sli_config.no_static_routes / 2d7cf07643ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d025acc3a0f41f7a49a32774ece104972fe4bd772ee60ae3ce18edcff6886f58"></a>

## Next pages — local_vrf.sli_config.no_static_routes / 2d7cf07643ba / 4

- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-8237b382c33ffeeb8bf28b3da0408daa323052008d4689166bddd66ac6cd1837"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59ba592ae30a7dfde0630c196621c0ace3cba3105492ab64985cd618d4b0f876"></a>

## local_vrf.sli_config.no_v6_static_routes — local_vrf.sli_config.no_v6_static_routes / 0e7f84fb2d14 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- local_vrf.sli_config.no_v6_static_routes

<a id="canonical-70cf803ddc2be61e08d96904f8fdedc1be898d1f0e014d8625ab5c9e98a1d726"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no v6 static routes.

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

<a id="canonical-74d0c4061237e0df37af3dd1883a8c12e1f29c7f9606adcf75ce0cabe1eebc57"></a>

## Direct properties — local_vrf.sli_config.no_v6_static_routes / 0e7f84fb2d14 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bf3930c595833c76dc602127b686aa91fcc0e60d2508099ae39b8f60833ad73e"></a>

## Next pages — local_vrf.sli_config.no_v6_static_routes / 0e7f84fb2d14 / 4

- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-93ea8120d47ebe43795f1bc7631e5287e64f1d51aedcc41ca54172ebd8a1278f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d00ec5ace9654254e34493b09fe981692fb8a2de31ce47eb213c6725f6bb31cc"></a>

## local_vrf.sli_config.static_routes — local_vrf.sli_config.static_routes / db46592b384c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- local_vrf.sli_config.static_routes

<a id="canonical-2373f547e965c5e5ec91701add53d19598521eb837a60a13a109b049f4a64dda"></a>

Type: `"single"`. Computed.

Configuration parameter for static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ba9caf8aed2fe259bc9b62b7d817822591c6890e2c3f0f206d518d4f522db056"></a>

## Direct properties — local_vrf.sli_config.static_routes / db46592b384c / 3

- [static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-5f46f43931cb082577bb8ecbaf221c81660981d5af400e0124a883dd8ce099c4): complete subsection reference.

<a id="canonical-5fb82706382c653256038283f1f11e81acaad3c2bd6efbbc648f350e977316e8"></a>

## Next pages — local_vrf.sli_config.static_routes / db46592b384c / 4

- [local_vrf.sli_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-5f46f43931cb082577bb8ecbaf221c81660981d5af400e0124a883dd8ce099c4)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-5f46f43931cb082577bb8ecbaf221c81660981d5af400e0124a883dd8ce099c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-562fb794efe731b6367e8c82b2595611bf1fbdd306d3d8b43bd40ec03a183247"></a>

## local_vrf.sli_config.static_routes.static_routes — local_vrf.sli_config.static_routes.static_routes / 27f0aaf00676 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [local_vrf.sli_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-93ea8120d47ebe43795f1bc7631e5287e64f1d51aedcc41ca54172ebd8a1278f)
- local_vrf.sli_config.static_routes.static_routes

<a id="canonical-a5b7ea828258a5eb62fa0985869ef374788ad8577213ada2be2caa8977fd7ba3"></a>

Type: `"list"`. Computed.

Configuration parameter for static routes.

Upstream description:

Configuration parameter for static routes

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

<a id="canonical-4064c207400928189f9a17aeef80155376e74623acea203691a8b48a8beac367"></a>

## Direct properties — local_vrf.sli_config.static_routes.static_routes / 27f0aaf00676 / 3

<a id="canonical-e1b8d4aa4334114c2e2f32acccb3367138e2c4076dc322a53644487cb659ff16"></a>

<a id="canonical-40d07030522bc0e5452acee4cb4e16cc9c1435c6b7f0543a2ddc5040207888e8"></a>

## attrs property — local_vrf.sli_config.static_routes.static_routes / 27f0aaf00676 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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

- [default_gateway](data-sources--securemesh_site_v2--reference--group-011.md#canonical-30e64e2ea2476db06154878c49554852e71247afcfac7e283bdc740f6f81d8d9): complete subsection reference.

<a id="canonical-180e889d59f916ef2ec9effcbcb813758bb033336f545b76980ddbe634938a7e"></a>

<a id="canonical-1812e679c544fe4ab884de09d694251a76e42e6e7a522a5c7b682e573145275f"></a>

## ip_address property — local_vrf.sli_config.static_routes.static_routes / 27f0aaf00676 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-4c99c4691eb7797663d6bc3c85ed47eac2e4c437f949bb3ea3daa04b7b63c279"></a>

<a id="canonical-e64cc04740dd5469a0826f3de62ab6f3f44f39f75904d3f19f396e7d39ef7d80"></a>

## ip_prefixes property — local_vrf.sli_config.static_routes.static_routes / 27f0aaf00676 / 6

Type: `["list", "string"]`. Computed.

List of route prefixes that have common next hop and attributes.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-4f4eb478f369bbfaa941201fabd8a675a9076d0e980d7a882b8c4cab19552591): complete subsection reference.

<a id="canonical-9bb900adb41026e59fdec9627b3fbfe1175d51abb27358ee35c07bbe2fc5b1e9"></a>

## Next pages — local_vrf.sli_config.static_routes.static_routes / 27f0aaf00676 / 7

- [local_vrf.sli_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site_v2--reference--group-011.md#canonical-30e64e2ea2476db06154878c49554852e71247afcfac7e283bdc740f6f81d8d9)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-4f4eb478f369bbfaa941201fabd8a675a9076d0e980d7a882b8c4cab19552591)
- [local_vrf.sli_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-93ea8120d47ebe43795f1bc7631e5287e64f1d51aedcc41ca54172ebd8a1278f)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-30e64e2ea2476db06154878c49554852e71247afcfac7e283bdc740f6f81d8d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01c3e69fe6f565ad3edf09f57947a84e9780d8915f47a6993dac99c3d0f35cb2"></a>

## local_vrf.sli_config.static_routes.static_routes.default_gateway — local_vrf.sli_config.static_routes.static_routes.default_gateway / ee06d485e7be / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [local_vrf.sli_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-93ea8120d47ebe43795f1bc7631e5287e64f1d51aedcc41ca54172ebd8a1278f)
- [local_vrf.sli_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-5f46f43931cb082577bb8ecbaf221c81660981d5af400e0124a883dd8ce099c4)
- local_vrf.sli_config.static_routes.static_routes.default_gateway

<a id="canonical-463591c6023464fad0b61516bf7fe4f179519991d8d9a7b7c5a3050a899b10a8"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-57bc8c85559b094bc8aed21ce55609b63cdacaa49289ad8f4f55e88ff7107530"></a>

## Direct properties — local_vrf.sli_config.static_routes.static_routes.default_gateway / ee06d485e7be / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-92a844c55c60d8cd81297ec3f7f1833eb97bf7b447a4fc8f42af78c9a3f807cc"></a>

## Next pages — local_vrf.sli_config.static_routes.static_routes.default_gateway / ee06d485e7be / 4

- [local_vrf.sli_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-5f46f43931cb082577bb8ecbaf221c81660981d5af400e0124a883dd8ce099c4)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-4f4eb478f369bbfaa941201fabd8a675a9076d0e980d7a882b8c4cab19552591"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60ff102cc87e771eb9571cf6c18b19787a30d3904e0c384823a4ddd621b4d586"></a>

## local_vrf.sli_config.static_routes.static_routes.node_interface — local_vrf.sli_config.static_routes.static_routes.node_interface / 3eb7d7608de2 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [local_vrf.sli_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-93ea8120d47ebe43795f1bc7631e5287e64f1d51aedcc41ca54172ebd8a1278f)
- [local_vrf.sli_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-5f46f43931cb082577bb8ecbaf221c81660981d5af400e0124a883dd8ce099c4)
- local_vrf.sli_config.static_routes.static_routes.node_interface

<a id="canonical-effd5e213af8ecb1d4c411f2a0371b057607238452d99be73639e14feca0f75b"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1a1b58c4c3da8a18e4332ae1fbced5d9bb4fc91856b54e99abbb18c7e0d06200"></a>

## Direct properties — local_vrf.sli_config.static_routes.static_routes.node_interface / 3eb7d7608de2 / 3

- [list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-29105fb9fc9ae663bf7dc953b5663a3928629557859422e1c3e83dee930fa0ee): complete subsection reference.

<a id="canonical-a0ccf48ba04dd0efce609ebe1a9b74f7808a7d199f3d0a70f63a4bea112c6a55"></a>

## Next pages — local_vrf.sli_config.static_routes.static_routes.node_interface / 3eb7d7608de2 / 4

- [local_vrf.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-29105fb9fc9ae663bf7dc953b5663a3928629557859422e1c3e83dee930fa0ee)
- [local_vrf.sli_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-5f46f43931cb082577bb8ecbaf221c81660981d5af400e0124a883dd8ce099c4)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-29105fb9fc9ae663bf7dc953b5663a3928629557859422e1c3e83dee930fa0ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e924240511328e399f0ac0062c140c4496bb56704e1a71cbc35d56f0e2956163"></a>

## local_vrf.sli_config.static_routes.static_routes.node_interface.list — local_vrf.sli_config.static_routes.static_routes.node_interface.list / c0cfd43748fb / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [local_vrf.sli_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-93ea8120d47ebe43795f1bc7631e5287e64f1d51aedcc41ca54172ebd8a1278f)
- [local_vrf.sli_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-5f46f43931cb082577bb8ecbaf221c81660981d5af400e0124a883dd8ce099c4)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-4f4eb478f369bbfaa941201fabd8a675a9076d0e980d7a882b8c4cab19552591)
- local_vrf.sli_config.static_routes.static_routes.node_interface.list

<a id="canonical-1bd039096f5727ab75908d75759b145b1a437d0b2a68534daa049a3aa2dd24fc"></a>

Type: `"list"`. Computed.

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

<a id="canonical-4f6b0c604dcd25590fe5ad8977e5c325049e1fdbae6b49cdf0a64325038630e1"></a>

## Direct properties — local_vrf.sli_config.static_routes.static_routes.node_interface.list / c0cfd43748fb / 3

- [interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cf48e802dc9a666e161a72f83cf1ecdc1894b4fddfeeed131b94ed0d74318613): complete subsection reference.

<a id="canonical-cd116e24c7e97736287921182a6bac93c5ba01a382b079db14a9a16a8115ab8f"></a>

<a id="canonical-ff5af3cf06f93955367c4d798e6322d521be8f8e03694dbb46f63d7d56107c19"></a>

## node property — local_vrf.sli_config.static_routes.static_routes.node_interface.list / c0cfd43748fb / 4

Type: `"string"`. Computed.

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

<a id="canonical-3ef0bfcd9448ac40ba754aeba2b666738e3021e58ff2a75a414c41640c321a3b"></a>

## Next pages — local_vrf.sli_config.static_routes.static_routes.node_interface.list / c0cfd43748fb / 5

- [local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cf48e802dc9a666e161a72f83cf1ecdc1894b4fddfeeed131b94ed0d74318613)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-4f4eb478f369bbfaa941201fabd8a675a9076d0e980d7a882b8c4cab19552591)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-cf48e802dc9a666e161a72f83cf1ecdc1894b4fddfeeed131b94ed0d74318613"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-837f159cbe27648c0151226cbdea48c67c7b392114f76a57f47892819d676e50"></a>

## local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a5f5806ed07b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [local_vrf.sli_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-93ea8120d47ebe43795f1bc7631e5287e64f1d51aedcc41ca54172ebd8a1278f)
- [local_vrf.sli_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-5f46f43931cb082577bb8ecbaf221c81660981d5af400e0124a883dd8ce099c4)
- [local_vrf.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-4f4eb478f369bbfaa941201fabd8a675a9076d0e980d7a882b8c4cab19552591)
- [local_vrf.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-29105fb9fc9ae663bf7dc953b5663a3928629557859422e1c3e83dee930fa0ee)
- local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-fb0e61ab3670a8d2e4dfe87c4fe5426a4fcd53d3dde9b4fa2df39f08a5197e17"></a>

Type: `"list"`. Computed.

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

<a id="canonical-1e16628e1eafc6b908a857bdf7e6402bdba9ef4093fa310530832d09a177f2b8"></a>

## Direct properties — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a5f5806ed07b / 3

<a id="canonical-eee9d9c174d13e4b58ad0ffb6a4c1e001f867ffd394189f2e1eea25891832f15"></a>

<a id="canonical-11449e38f99665ef83e9dd6207826ecef1d53e99bf584b72de18492e8b85ce4c"></a>

## kind property — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a5f5806ed07b / 4

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

<a id="canonical-b3bc3dcef10d1d4dd36f709c3a49159e0da7cc08fcd9841aa3d38b71ef22d415"></a>

<a id="canonical-b9506b906fb54eb3cfdfabd20c3595c7b01e8b9529fdd0c7703e9f7063093c78"></a>

## name property — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a5f5806ed07b / 5

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

<a id="canonical-6e894e373ff2136685d47434959a24b6e0a8af54a9094c26cde3c390cbc103fa"></a>

<a id="canonical-74860bb6ad7d71b961f6f13915989065f638d4c90e678fd728ff4b74c78f7555"></a>

## namespace property — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a5f5806ed07b / 6

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

<a id="canonical-929d2490bb5ba126539a3364b5e3a9e22e0821506664540eb7b80ebb1389757a"></a>

<a id="canonical-ac7cd5ae12752c409af3b40e45fc6c058f8b73ceab4b2f429c74b4f5cc3f9b3e"></a>

## tenant property — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a5f5806ed07b / 7

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

<a id="canonical-f4b5d11eac497c22122c4b586cf9ee629b2fe9aab5870fdb0d5379bd4f4c1e89"></a>

<a id="canonical-eadb23a21a89b37f8a078142faeccf792edf44f5c1816547cae33ac2f784c104"></a>

## uid property — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a5f5806ed07b / 8

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

<a id="canonical-6ee586671b46be047bd95a02d1c89205aea2dc5ae8660ab903bafc59f3ae590b"></a>

## Next pages — local_vrf.sli_config.static_routes.static_routes.node_interface.list.interface / a5f5806ed07b / 9

- [local_vrf.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-29105fb9fc9ae663bf7dc953b5663a3928629557859422e1c3e83dee930fa0ee)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-3704e98825faca68e00e9b9fc724f1062824ae979806ea5ae7f6664216dfab06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad225926d28383205ca90c65c7c9ac9f6316483c7370ffd99313933721b4a70c"></a>

## local_vrf.sli_config.static_v6_routes — local_vrf.sli_config.static_v6_routes / 7ae49fb40e17 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- local_vrf.sli_config.static_v6_routes

<a id="canonical-fb44c533a7f296e87dcfedda1db6302864e8db002ebe99672b0804d0da046c6b"></a>

Type: `"single"`. Computed.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2cde05ed395e94913c070c97a4722602dbf63d8589aa68ea33d28965b2a17aae"></a>

## Direct properties — local_vrf.sli_config.static_v6_routes / 7ae49fb40e17 / 3

- [static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-239831c8a1f0ea3cf6e5126ce51332d8b790ddbc2d187fdff54d0e79cbc809ec): complete subsection reference.

<a id="canonical-167d3fb19e05e8c75947aa320bb523ecb6164fd0de722c1b7521d11c5505051f"></a>

## Next pages — local_vrf.sli_config.static_v6_routes / 7ae49fb40e17 / 4

- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-239831c8a1f0ea3cf6e5126ce51332d8b790ddbc2d187fdff54d0e79cbc809ec)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-239831c8a1f0ea3cf6e5126ce51332d8b790ddbc2d187fdff54d0e79cbc809ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-480aaf5df6a6925a676490bf7041649b7cca1b98d275b5d9594f0e0417450acb"></a>

## local_vrf.sli_config.static_v6_routes.static_routes — local_vrf.sli_config.static_v6_routes.static_routes / 4885a8cbf736 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3704e98825faca68e00e9b9fc724f1062824ae979806ea5ae7f6664216dfab06)
- local_vrf.sli_config.static_v6_routes.static_routes

<a id="canonical-2a4ebe424e2e08728d5fa8160d574043603e44eaf0a2016c84bb29bb552e87de"></a>

Type: `"list"`. Computed.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-92026b1c23a34c291dce41f30507181dcad0e6ba0ea4e113bcaa95a97176f532"></a>

## Direct properties — local_vrf.sli_config.static_v6_routes.static_routes / 4885a8cbf736 / 3

<a id="canonical-b3f4c26305fe91e221380ae35da3564b9a8aedc718a8fe72295aca195bcc4baa"></a>

<a id="canonical-5553d6ea99cfbceb298199a7df82cae0b2b7ad64e1d973cb92db8da9e3d3e573"></a>

## attrs property — local_vrf.sli_config.static_v6_routes.static_routes / 4885a8cbf736 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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

- [default_gateway](data-sources--securemesh_site_v2--reference--group-011.md#canonical-9f505b2af915a40c2805649c6167bdb8f9f82e761dd398c6d2b4b2965da88338): complete subsection reference.

<a id="canonical-f7b2d042f53a95914f9f64b44497737b600b65558b3a0ec850d112643c5af49c"></a>

<a id="canonical-983b09c7e86f7ec05d40f2670d5469bcebca1d1d617d1eabc262474ef1786c3e"></a>

## ip_address property — local_vrf.sli_config.static_v6_routes.static_routes / 4885a8cbf736 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

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

<a id="canonical-387e573de2847544deb96d80ed6bf6f10cc43983d754d47dec970a525d06df5c"></a>

<a id="canonical-ddd5a258555b87b6c746eed2dad6c4976b1c7c045fc6b460547070ed1a366132"></a>

## ip_prefixes property — local_vrf.sli_config.static_v6_routes.static_routes / 4885a8cbf736 / 6

Type: `["list", "string"]`. Computed.

List of IPv6 route prefixes that have common next hop and attributes.

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

- [node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-a7632944e7a590751f7ec6d3662a307e906389045e4bb940c1a4803c7978c8de): complete subsection reference.

<a id="canonical-7e69bc3781bbbeb3db75968c84ae4cb4dfadf48743e25fd043abeebf711e34f8"></a>

## Next pages — local_vrf.sli_config.static_v6_routes.static_routes / 4885a8cbf736 / 7

- [local_vrf.sli_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site_v2--reference--group-011.md#canonical-9f505b2af915a40c2805649c6167bdb8f9f82e761dd398c6d2b4b2965da88338)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-a7632944e7a590751f7ec6d3662a307e906389045e4bb940c1a4803c7978c8de)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3704e98825faca68e00e9b9fc724f1062824ae979806ea5ae7f6664216dfab06)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9f505b2af915a40c2805649c6167bdb8f9f82e761dd398c6d2b4b2965da88338"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67fda428220bf13fb5d63fb7d6426d1b80cf073a52492f44dc00a43622b47388"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.default_gateway — local_vrf.sli_config.static_v6_routes.static_routes.default_gateway / b74f9e1044c6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3704e98825faca68e00e9b9fc724f1062824ae979806ea5ae7f6664216dfab06)
- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-239831c8a1f0ea3cf6e5126ce51332d8b790ddbc2d187fdff54d0e79cbc809ec)
- local_vrf.sli_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-a4b4137a49aec7783e244acb75b042b33fbf507e2fdfab73a8bfccf6456e242b"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-31725b5ae70f83013e707d078bd9a7e95d4acb11945d8dcc7f703ed59ed6a524"></a>

## Direct properties — local_vrf.sli_config.static_v6_routes.static_routes.default_gateway / b74f9e1044c6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dc93c863702a35371dab0e77b77649082bc2008a7ac0df9ca9ed97c0ead5f30e"></a>

## Next pages — local_vrf.sli_config.static_v6_routes.static_routes.default_gateway / b74f9e1044c6 / 4

- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-239831c8a1f0ea3cf6e5126ce51332d8b790ddbc2d187fdff54d0e79cbc809ec)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-a7632944e7a590751f7ec6d3662a307e906389045e4bb940c1a4803c7978c8de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64568e0e9c544e6e63cd2d6446fe8472267a33153c3fede35d913746399a8c17"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.node_interface — local_vrf.sli_config.static_v6_routes.static_routes.node_interface / f22011fcda46 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3704e98825faca68e00e9b9fc724f1062824ae979806ea5ae7f6664216dfab06)
- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-239831c8a1f0ea3cf6e5126ce51332d8b790ddbc2d187fdff54d0e79cbc809ec)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface

<a id="canonical-abad184fe20cda89d22a8f8b514da0d99db38b67076bab444fdef80936f12995"></a>

Type: `"single"`. Computed.

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

<a id="canonical-4df14e16cd93f1ad6f31a6f58ffd78985fecbcf251b84fc4eec02789779548f4"></a>

## Direct properties — local_vrf.sli_config.static_v6_routes.static_routes.node_interface / f22011fcda46 / 3

- [list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-396d211ef79082c727aec14addaf1413ad32daa95e0de50e83c045f0ffcca3d5): complete subsection reference.

<a id="canonical-9a8585229b1b1ffcb8aa1ecb8ab3a25d6e26bdb807d1bdb945b0d3f36d52d475"></a>

## Next pages — local_vrf.sli_config.static_v6_routes.static_routes.node_interface / f22011fcda46 / 4

- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-396d211ef79082c727aec14addaf1413ad32daa95e0de50e83c045f0ffcca3d5)
- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-239831c8a1f0ea3cf6e5126ce51332d8b790ddbc2d187fdff54d0e79cbc809ec)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-396d211ef79082c727aec14addaf1413ad32daa95e0de50e83c045f0ffcca3d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6172f628e0df2deb6eebe8d907a79880100e86eaad4fba534d875a666b8a66bb"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list / 7752c1ad37a3 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3704e98825faca68e00e9b9fc724f1062824ae979806ea5ae7f6664216dfab06)
- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-239831c8a1f0ea3cf6e5126ce51332d8b790ddbc2d187fdff54d0e79cbc809ec)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-a7632944e7a590751f7ec6d3662a307e906389045e4bb940c1a4803c7978c8de)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-a37c3918ad5bc5548cab597252879c4f3d1db09cfe60965ce6e3df48910485b8"></a>

Type: `"list"`. Computed.

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

<a id="canonical-bfe2c09aa7538335bcbf02bccab976fbe24a66fdc54fcac8c594606f782b3b04"></a>

## Direct properties — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list / 7752c1ad37a3 / 3

- [interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-f88e437ce121c569640c0e2fe8b197806c5b72d2ab3e9093c5f579d5171305c7): complete subsection reference.

<a id="canonical-6a105c068abb428b19c141743400edbfd62f12f5f4948d8949f37ba0bcbd6bd3"></a>

<a id="canonical-4ad1f3d9732b3d60ef0ef3a85d5b20c62c3b856ef03c869e2574e997ebd93d87"></a>

## node property — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list / 7752c1ad37a3 / 4

Type: `"string"`. Computed.

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

<a id="canonical-923d44f59114be4f9498717441b64cd9878cfdc4b02528f917c6ac3728d52a05"></a>

## Next pages — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list / 7752c1ad37a3 / 5

- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-f88e437ce121c569640c0e2fe8b197806c5b72d2ab3e9093c5f579d5171305c7)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-a7632944e7a590751f7ec6d3662a307e906389045e4bb940c1a4803c7978c8de)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f88e437ce121c569640c0e2fe8b197806c5b72d2ab3e9093c5f579d5171305c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-62a498008736c319c8d39fabc359e40158a128ca36e3f4c193f35307283a7c1c"></a>

## local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 693db71339fa / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.sli_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-42e832b854df6d0d16adc26a7a9e4da57a90d8ff0ef3941950937211935eb01a)
- [local_vrf.sli_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3704e98825faca68e00e9b9fc724f1062824ae979806ea5ae7f6664216dfab06)
- [local_vrf.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-239831c8a1f0ea3cf6e5126ce51332d8b790ddbc2d187fdff54d0e79cbc809ec)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-a7632944e7a590751f7ec6d3662a307e906389045e4bb940c1a4803c7978c8de)
- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-396d211ef79082c727aec14addaf1413ad32daa95e0de50e83c045f0ffcca3d5)
- local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-83e193d28d695e2b0ab96483c77cdd51e8a0a8c24b9d815354e5ac211eeb9ca0"></a>

Type: `"list"`. Computed.

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

<a id="canonical-fdb574f8680da539fdf4bab64b590d7f5344afdab1a20dea902bc6e155be0c2d"></a>

## Direct properties — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 693db71339fa / 3

<a id="canonical-56d09b42fc17c8bbf4ec6273dade760fb8ce7e895fa56f773e4073f10bf8b75a"></a>

<a id="canonical-0333724403d0ce77c0b5c3698fa910788c5e883a575f82ad5ca08bed9c630dd1"></a>

## kind property — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 693db71339fa / 4

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

<a id="canonical-2a2dc3968d3f86db3b9dc7d479ffe3c69f18f8f78077df222349c6da8e05e599"></a>

<a id="canonical-eb8b2f67d7e96fd2c94982b1a0117ba4aaed869b28963d728dd42c2e5e1bf460"></a>

## name property — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 693db71339fa / 5

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

<a id="canonical-2dac834feb3a834db660c3994a7643186802212271a2c236ce7ab5c39fe4fd4b"></a>

<a id="canonical-9b4b408c97ff08e2c860772f6e69975c0aa3a714d02a543c3aeead5781a80606"></a>

## namespace property — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 693db71339fa / 6

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

<a id="canonical-761d05d80e5bd4fce7a908b7a12825aa125b77d48c25ea9e3ce76f59d8fe9709"></a>

<a id="canonical-a44d86c70024a16661fc9d06c21d1c8990312bef69fb0e2dea1d8dd008a7cf02"></a>

## tenant property — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 693db71339fa / 7

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

<a id="canonical-52d929ff42d1a7362b97952fa4cb78aaa779131a7c65b180cee436494ed9eb28"></a>

<a id="canonical-e5e6b06316f35f19a3feac846322f4f702a54df33d39dbdd56027d9c5f86205d"></a>

## uid property — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 693db71339fa / 8

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

<a id="canonical-0ce328f6851dd7b8a45b3faacc126e4e556d57ef10d1316240d4066cbc9d340a"></a>

## Next pages — local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list.interfac / 693db71339fa / 9

- [local_vrf.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-396d211ef79082c727aec14addaf1413ad32daa95e0de50e83c045f0ffcca3d5)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84b356407a1521279443396b354da99e1263a3fd134cb73792a5b4df70a5ab24"></a>

## local_vrf.slo_config — local_vrf.slo_config / 625a854f6a19 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- local_vrf.slo_config

<a id="canonical-47c8393844e76886570ac34637e9353c1cf1427e3b3af70ad34aff18922baa3d"></a>

Type: `"single"`. Computed.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

<a id="canonical-2283bcccb9bca532fc1b012c72507a9405cd6ee5c6642b656d6cfc5a506fb10b"></a>

## Direct properties — local_vrf.slo_config / 625a854f6a19 / 3

<a id="canonical-036d928947709f91cde1e751f41f2398c01cd98a908cf9c06aeb04050aea9a44"></a>

<a id="canonical-bd56bd7b8237dcb2528375b0b107aa49206b8b6db9a1c9dc957e3b01e97db1d4"></a>

## labels property — local_vrf.slo_config / 625a854f6a19 / 4

Type: `["map", "string"]`. Computed.

Add Labels for this network, these labels can be used in firewall policy.

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

<a id="canonical-9000d931ec8cb0a6ad5294262677bd39c32cdd6e767c04be5457fce181653e55"></a>

<a id="canonical-9504b00eddf12bab3bfa21708c50236aa36de411a42ec38bfb928546debe36d5"></a>

## nameserver property — local_vrf.slo_config / 625a854f6a19 / 5

Type: `"string"`. Computed.

Optional IPv4 DNS server to be used for name resolution.

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

- [no_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-9e5f0823006ee3d2ee05ac8d83a8019ec34421863299d437f9fc87c6707ed962): complete subsection reference.

- [no_v6_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-59e97910b947bd19b3943117887ab667d3462178463ce281eff23204280bf298): complete subsection reference.

<a id="canonical-d9d9f2be11a22ba6321d1d1a21a9d1079f02df99ef6eb56b369eab7f7b6f2575"></a>

<a id="canonical-978764cac17350ff11b1b96a079cd68ab0bbeca24832a1d7a5538f86a9622d83"></a>

## secondary_nameserver property — local_vrf.slo_config / 625a854f6a19 / 6

Type: `"string"`. Computed.

Optional Secondary IPv4 DNS server to be used for name resolution.

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

- [static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-578c9a0a507f4f8ff8a98d4e74207e3561a39b1a2b43e75128c954325ec66634): complete subsection reference.

- [static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cee7717acf2f93282deb1242c0e8f0c270191bfff1980a588b87c3198f229a2a): complete subsection reference.

<a id="canonical-c5623c6dfb6939843c77f9872e1fc65f088201c967262b37ea440118cfe59634"></a>

<a id="canonical-1dbcf6ef67b9081eebb30264e47650e35bcde8015c9214ee155cea793208599e"></a>

## vip property — local_vrf.slo_config / 625a854f6a19 / 7

Type: `"string"`. Computed.

Optional common virtual V4 IP across all nodes to be used as automatic VIP.

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

<a id="canonical-b5af4e074330cd3858c66b2322a5d51ea596e6a4fa73d9750ae1fff0d3e0427e"></a>

## Next pages — local_vrf.slo_config / 625a854f6a19 / 8

- [local_vrf.slo_config.no_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-9e5f0823006ee3d2ee05ac8d83a8019ec34421863299d437f9fc87c6707ed962)
- [local_vrf.slo_config.no_v6_static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-59e97910b947bd19b3943117887ab667d3462178463ce281eff23204280bf298)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-578c9a0a507f4f8ff8a98d4e74207e3561a39b1a2b43e75128c954325ec66634)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cee7717acf2f93282deb1242c0e8f0c270191bfff1980a588b87c3198f229a2a)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9e5f0823006ee3d2ee05ac8d83a8019ec34421863299d437f9fc87c6707ed962"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fec418e4f53a104a05b36fcc7a79dfc59da07ce9fa99b13299764226c14d5438"></a>

## local_vrf.slo_config.no_static_routes — local_vrf.slo_config.no_static_routes / a2eb403b4129 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- local_vrf.slo_config.no_static_routes

<a id="canonical-029af0c23a5d01c0f27e2741b254f3dc801fbcb7cd714ac172fcb84fc31051e0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no static routes.

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

<a id="canonical-3dd69a47e34902b91eef60d1b5c907af2e7eb297028a6674d9ee0466c084443e"></a>

## Direct properties — local_vrf.slo_config.no_static_routes / a2eb403b4129 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d580e5fc6b24dd98ce80b7bbf5af3e2145aa80a2d09f45af852127dd2c3ce631"></a>

## Next pages — local_vrf.slo_config.no_static_routes / a2eb403b4129 / 4

- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-59e97910b947bd19b3943117887ab667d3462178463ce281eff23204280bf298"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79506ed28ab267944320be5f23ad3a54321cdc0b39a06b2d8766d6be406304ed"></a>

## local_vrf.slo_config.no_v6_static_routes — local_vrf.slo_config.no_v6_static_routes / 7569d4d9463a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- local_vrf.slo_config.no_v6_static_routes

<a id="canonical-d77e99614cbbd3f33b556cfc12f58b96a60ec757528306c243d9876a41d70793"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no v6 static routes.

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

<a id="canonical-ad2d3a1d62fb5120e7fa32ea364cb531bcd5997ed92a09f3285295c9f85fbc36"></a>

## Direct properties — local_vrf.slo_config.no_v6_static_routes / 7569d4d9463a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-691fb7b4504ffb6ce1b746463f8ea9449301d5510035124ee2b98ee0f2c192e0"></a>

## Next pages — local_vrf.slo_config.no_v6_static_routes / 7569d4d9463a / 4

- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-578c9a0a507f4f8ff8a98d4e74207e3561a39b1a2b43e75128c954325ec66634"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5df5e1b84c3dc5c4b1075173970feadfffd628a69d5314a202e895368adca40b"></a>

## local_vrf.slo_config.static_routes — local_vrf.slo_config.static_routes / 2a0f26fc8567 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- local_vrf.slo_config.static_routes

<a id="canonical-6536978cdf9462776783351a1236e546e42290a1ad14ba720cf48b0e7c5dde96"></a>

Type: `"single"`. Computed.

Configuration parameter for static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3681c238b2331f0e7231d21e6d53c2997bac6fd292ebce70176a02eeb1197b80"></a>

## Direct properties — local_vrf.slo_config.static_routes / 2a0f26fc8567 / 3

- [static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-84fc052204d42a2877b6f41c53143fb56e8e8b739afcb75f44e564fb23f5cb87): complete subsection reference.

<a id="canonical-fb46e5fd255215451508bb75759dd4d975c2d371cc62a3da57234a417ae677d3"></a>

## Next pages — local_vrf.slo_config.static_routes / 2a0f26fc8567 / 4

- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-84fc052204d42a2877b6f41c53143fb56e8e8b739afcb75f44e564fb23f5cb87)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-84fc052204d42a2877b6f41c53143fb56e8e8b739afcb75f44e564fb23f5cb87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d60302a7ae754415038c88fd9ed96d78d16bb3da52665a5e280b6aad8436bf3"></a>

## local_vrf.slo_config.static_routes.static_routes — local_vrf.slo_config.static_routes.static_routes / 013778d61884 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-578c9a0a507f4f8ff8a98d4e74207e3561a39b1a2b43e75128c954325ec66634)
- local_vrf.slo_config.static_routes.static_routes

<a id="canonical-ba56ff85a95af3ed0011e6a4aa12117f15f683a5c4308a5867179972ab822396"></a>

Type: `"list"`. Computed.

Configuration parameter for static routes.

Upstream description:

Configuration parameter for static routes

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

<a id="canonical-0b07fdbb1e311e71761a00cd3e4adc9ed784f3d94c13c3c49335081fec956e21"></a>

## Direct properties — local_vrf.slo_config.static_routes.static_routes / 013778d61884 / 3

<a id="canonical-8f237a0618138fa4f986e76c645f1c2efae40b7612dbaccf960f745070f36827"></a>

<a id="canonical-9ac32f7ede67a4deffe1097de688f7498065b2584ead8cc789d712733bf00329"></a>

## attrs property — local_vrf.slo_config.static_routes.static_routes / 013778d61884 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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

- [default_gateway](data-sources--securemesh_site_v2--reference--group-011.md#canonical-640bdad8d46b6e3915e29e45049aed852ae612730c2653c1486b65b5a7789819): complete subsection reference.

<a id="canonical-669d76561eedc3abb9b97f0ce6c309a89301e8352fe6df5bd2f2556e36eede9a"></a>

<a id="canonical-f4cea8a28f91a47a46e5ab6886ce79dc372bfde8ad073a0b28207787d6b2048d"></a>

## ip_address property — local_vrf.slo_config.static_routes.static_routes / 013778d61884 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-5d75d651b44e01b36679f375ab957409a3161c8034d1f39fc42cbbf50370bd8c"></a>

<a id="canonical-7a1e06f25357da0fe9b7cc6a6b1054b6b21753d13e064ed66e4f6a51d7f452de"></a>

## ip_prefixes property — local_vrf.slo_config.static_routes.static_routes / 013778d61884 / 6

Type: `["list", "string"]`. Computed.

List of route prefixes that have common next hop and attributes.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-9facb647b641a54c10201a51367bf93112163c191cb7861da1c22195bdf59341): complete subsection reference.

<a id="canonical-082070b6533a1d3792d90edc60b598b3bd601537bdf372725cd670999cde41ce"></a>

## Next pages — local_vrf.slo_config.static_routes.static_routes / 013778d61884 / 7

- [local_vrf.slo_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site_v2--reference--group-011.md#canonical-640bdad8d46b6e3915e29e45049aed852ae612730c2653c1486b65b5a7789819)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-9facb647b641a54c10201a51367bf93112163c191cb7861da1c22195bdf59341)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-578c9a0a507f4f8ff8a98d4e74207e3561a39b1a2b43e75128c954325ec66634)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-640bdad8d46b6e3915e29e45049aed852ae612730c2653c1486b65b5a7789819"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da77771d3003e0c27beff38e13194c3521459fa635cda88fe393b672d6a72287"></a>

## local_vrf.slo_config.static_routes.static_routes.default_gateway — local_vrf.slo_config.static_routes.static_routes.default_gateway / 1344df2d4747 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-578c9a0a507f4f8ff8a98d4e74207e3561a39b1a2b43e75128c954325ec66634)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-84fc052204d42a2877b6f41c53143fb56e8e8b739afcb75f44e564fb23f5cb87)
- local_vrf.slo_config.static_routes.static_routes.default_gateway

<a id="canonical-e8f01c522526e4df68dba63d5e62009cf4934a496975d606ce4d415ccb1ace20"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-5ddb6786d8fae8b7f1a683210fcf93e2fe61871ce1d5dfbfe51b8faefd2a00c4"></a>

## Direct properties — local_vrf.slo_config.static_routes.static_routes.default_gateway / 1344df2d4747 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5983258f2c286721d9df8fc35b8f159a26c8a58d5c28a11bf4842b25e3d80ec7"></a>

## Next pages — local_vrf.slo_config.static_routes.static_routes.default_gateway / 1344df2d4747 / 4

- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-84fc052204d42a2877b6f41c53143fb56e8e8b739afcb75f44e564fb23f5cb87)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9facb647b641a54c10201a51367bf93112163c191cb7861da1c22195bdf59341"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0a9c5372c917bdb5d23d0bdc31930f4b2a55e17c240a20c216a66889b3e9b82"></a>

## local_vrf.slo_config.static_routes.static_routes.node_interface — local_vrf.slo_config.static_routes.static_routes.node_interface / d7285de00d09 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-578c9a0a507f4f8ff8a98d4e74207e3561a39b1a2b43e75128c954325ec66634)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-84fc052204d42a2877b6f41c53143fb56e8e8b739afcb75f44e564fb23f5cb87)
- local_vrf.slo_config.static_routes.static_routes.node_interface

<a id="canonical-9447ae7b982a4a5a4bd8f80e4a5bd82a5e4d9eac043b29963260523058df3c0a"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3013b782f5a4a59c4e3b54adbeaec1d7072c1ed96dfa30dd4e9289f02420d5d6"></a>

## Direct properties — local_vrf.slo_config.static_routes.static_routes.node_interface / d7285de00d09 / 3

- [list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-468f8edea5791d537f47ab43561621e1676eb401ecf2b46b4958f09f208dedb9): complete subsection reference.

<a id="canonical-8013954ae11a877b8aa5459cefb20cdc12d6c58e588cca8cfefce794fb5943c9"></a>

## Next pages — local_vrf.slo_config.static_routes.static_routes.node_interface / d7285de00d09 / 4

- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-468f8edea5791d537f47ab43561621e1676eb401ecf2b46b4958f09f208dedb9)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-84fc052204d42a2877b6f41c53143fb56e8e8b739afcb75f44e564fb23f5cb87)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-468f8edea5791d537f47ab43561621e1676eb401ecf2b46b4958f09f208dedb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-826c12a1eaa67d62c982721d676034695be3af1731ef5d5b2ff095897148b7fd"></a>

## local_vrf.slo_config.static_routes.static_routes.node_interface.list — local_vrf.slo_config.static_routes.static_routes.node_interface.list / 667340117453 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-578c9a0a507f4f8ff8a98d4e74207e3561a39b1a2b43e75128c954325ec66634)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-84fc052204d42a2877b6f41c53143fb56e8e8b739afcb75f44e564fb23f5cb87)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-9facb647b641a54c10201a51367bf93112163c191cb7861da1c22195bdf59341)
- local_vrf.slo_config.static_routes.static_routes.node_interface.list

<a id="canonical-1d15fd5af33cb7a4f64ef8606562417bb06dc9cd9b07fc06f24c665a616cd6cd"></a>

Type: `"list"`. Computed.

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

<a id="canonical-88624838c9fe4962e49915a10f2b68cc38187438bf42ff84096f47747c64a6f6"></a>

## Direct properties — local_vrf.slo_config.static_routes.static_routes.node_interface.list / 667340117453 / 3

- [interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-ea52f888724d0b047fb3dd52b224e59c115e655c6e115511783ca138f028d572): complete subsection reference.

<a id="canonical-5e32bc17878d7c423d3f059ec93a18e007cad2fb6dbe889cdf663b81f11eea7a"></a>

<a id="canonical-5ca538eb47947be96c52023f8060fcd790f88d94de93abc41027ddae5cb639bb"></a>

## node property — local_vrf.slo_config.static_routes.static_routes.node_interface.list / 667340117453 / 4

Type: `"string"`. Computed.

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

<a id="canonical-ecd0ba3dfb465410600c0d16c6b39219769d75275ef59418984a95070efd07d8"></a>

## Next pages — local_vrf.slo_config.static_routes.static_routes.node_interface.list / 667340117453 / 5

- [local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-ea52f888724d0b047fb3dd52b224e59c115e655c6e115511783ca138f028d572)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-9facb647b641a54c10201a51367bf93112163c191cb7861da1c22195bdf59341)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-ea52f888724d0b047fb3dd52b224e59c115e655c6e115511783ca138f028d572"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb1ac013d9445878f03f77142175001151789fafe36dacae0f9ce06dedb52e03"></a>

## local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / 7107264a4418 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-578c9a0a507f4f8ff8a98d4e74207e3561a39b1a2b43e75128c954325ec66634)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-84fc052204d42a2877b6f41c53143fb56e8e8b739afcb75f44e564fb23f5cb87)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-9facb647b641a54c10201a51367bf93112163c191cb7861da1c22195bdf59341)
- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-468f8edea5791d537f47ab43561621e1676eb401ecf2b46b4958f09f208dedb9)
- local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-d28050a20bc0714b3805dec4aa2b9f67e727a1cefe5a9a9a053b0e41656be3ba"></a>

Type: `"list"`. Computed.

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

<a id="canonical-1d183d7177f707fec496a7ede21a3b2febb12c48275b4d5694365e086605f86d"></a>

## Direct properties — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / 7107264a4418 / 3

<a id="canonical-c19fed29c372fbe4a53f9ca716e9c4c3598879b47fad8c49619031c613e0180a"></a>

<a id="canonical-1bee8a769e2da843be9d7e160a7268c7d1cd8f1579d0bd3b10b614c119e35fae"></a>

## kind property — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / 7107264a4418 / 4

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

<a id="canonical-8acae55e2362cce22c31bfbaf6da578d89ed1d085aee3197531135390bf23720"></a>

<a id="canonical-9ef151fecc5f5fc32e94109ab7aa3ae9168d25c7aea14d87066fd4015e9814bc"></a>

## name property — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / 7107264a4418 / 5

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

<a id="canonical-f2ef5c976c9dc59e92dccdd5366f8fcb211fcbbf320eb9294681ea1e7e46b926"></a>

<a id="canonical-0f0bb295dbf9bfa8ba9267375233d92970af37081222c36b9fbd96b8e1d05073"></a>

## namespace property — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / 7107264a4418 / 6

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

<a id="canonical-89bf3b66d0fb99b5603b9b30d310bd4195687d945570364de2de2fc9ce8fcc10"></a>

<a id="canonical-3f1ec0fd08131d3cf2ffb90df26c0c46ba8c1194a103a4d27f7b279b906b815a"></a>

## tenant property — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / 7107264a4418 / 7

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

<a id="canonical-6cef4233f706461d338fca1b2c08000946e6330a514f360a0a86aad1e103c998"></a>

<a id="canonical-200c3cc2584c66416947f6b50368dbadbccd22ee1a1972babf16f9fb2b9bcb33"></a>

## uid property — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / 7107264a4418 / 8

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

<a id="canonical-923efa973688584a8b27a4c8235fccf18a8aee184316df9c7321e0dafb70174e"></a>

## Next pages — local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface / 7107264a4418 / 9

- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-468f8edea5791d537f47ab43561621e1676eb401ecf2b46b4958f09f208dedb9)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-cee7717acf2f93282deb1242c0e8f0c270191bfff1980a588b87c3198f229a2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56ed690eaaf7ef85e357310c164aa2c510b9f70fe1003079aeb2cdaeacfcd0fb"></a>

## local_vrf.slo_config.static_v6_routes — local_vrf.slo_config.static_v6_routes / 06efa8478726 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- local_vrf.slo_config.static_v6_routes

<a id="canonical-c9017f01f95fa06545b7a98ad7641e431b9af30fda17125548064e9fe446da50"></a>

Type: `"single"`. Computed.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-84dd00a88d913a2df2e2bcb8829e3779c62bf7648eebc0855dd86447e4c551f3"></a>

## Direct properties — local_vrf.slo_config.static_v6_routes / 06efa8478726 / 3

- [static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-af4854f9eec5985ac007d5b4c4a616a37dfe431fda08d1f8eaf3746fd296f326): complete subsection reference.

<a id="canonical-6d04f26d8a4252c69b872edd588ec7c42f1c92ddbbf59d0b52be90753a84ad22"></a>

## Next pages — local_vrf.slo_config.static_v6_routes / 06efa8478726 / 4

- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-af4854f9eec5985ac007d5b4c4a616a37dfe431fda08d1f8eaf3746fd296f326)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-af4854f9eec5985ac007d5b4c4a616a37dfe431fda08d1f8eaf3746fd296f326"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-108e45f8241f0391052834d60e2563e0bf680fdf54fd1d0ead185a9f3496f268"></a>

## local_vrf.slo_config.static_v6_routes.static_routes — local_vrf.slo_config.static_v6_routes.static_routes / 246818f60332 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cee7717acf2f93282deb1242c0e8f0c270191bfff1980a588b87c3198f229a2a)
- local_vrf.slo_config.static_v6_routes.static_routes

<a id="canonical-d2873c26d84fa5a16858249f3c63f20ae7c14344a07b11815628738ed9d3a2bd"></a>

Type: `"list"`. Computed.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-34b4065d431e94730b8a42430207d9d3dd1d7f70493f5c97c756dd25d1cd0217"></a>

## Direct properties — local_vrf.slo_config.static_v6_routes.static_routes / 246818f60332 / 3

<a id="canonical-3b371a81ff1319ba7a1a4493ddc5e9bdd320ce6bb99a790c2c014f63b80bc847"></a>

<a id="canonical-e5a96a895d8c37ba4ab20c07032d6f87e12840581834f7363b8eb3f63488dae6"></a>

## attrs property — local_vrf.slo_config.static_v6_routes.static_routes / 246818f60332 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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

- [default_gateway](data-sources--securemesh_site_v2--reference--group-011.md#canonical-f1bf2e34d939096ba13e0614dfccc6e672e3d29d0abf8b235606296a55da4ecb): complete subsection reference.

<a id="canonical-d7ffebc0b389241c7824387928496cd448e4e2c51be3c4e39bab4aa4025a36bb"></a>

<a id="canonical-eaf5a40da178b579adc0f2b0f8514c77941407915f341338cc701cc9260a8e24"></a>

## ip_address property — local_vrf.slo_config.static_v6_routes.static_routes / 246818f60332 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

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

<a id="canonical-9ad9f79fd21e42a1e06244c7183744a2fb3c50c8c19533ae415dd3fd129b1173"></a>

<a id="canonical-26c5b7c1c8b66f4911e1b2a231e2c45f1f1fa30ad90163f1825174177d31910b"></a>

## ip_prefixes property — local_vrf.slo_config.static_v6_routes.static_routes / 246818f60332 / 6

Type: `["list", "string"]`. Computed.

List of IPv6 route prefixes that have common next hop and attributes.

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

- [node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-bd09ea380f6603f6e1fbbe3371f913337387062749568cfd6ed63f02eef6fc12): complete subsection reference.

<a id="canonical-23fc7835d46771be819fe3ed629f51359a8d5ba9128bc8356cab5057ccd07644"></a>

## Next pages — local_vrf.slo_config.static_v6_routes.static_routes / 246818f60332 / 7

- [local_vrf.slo_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site_v2--reference--group-011.md#canonical-f1bf2e34d939096ba13e0614dfccc6e672e3d29d0abf8b235606296a55da4ecb)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-bd09ea380f6603f6e1fbbe3371f913337387062749568cfd6ed63f02eef6fc12)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cee7717acf2f93282deb1242c0e8f0c270191bfff1980a588b87c3198f229a2a)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f1bf2e34d939096ba13e0614dfccc6e672e3d29d0abf8b235606296a55da4ecb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fb490c3ef5d60412bb38907ea3e25cf2a578bf9e404c981d31796c927e0ae9d"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.default_gateway — local_vrf.slo_config.static_v6_routes.static_routes.default_gateway / cbe4ef09c785 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cee7717acf2f93282deb1242c0e8f0c270191bfff1980a588b87c3198f229a2a)
- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-af4854f9eec5985ac007d5b4c4a616a37dfe431fda08d1f8eaf3746fd296f326)
- local_vrf.slo_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-81ab4f0c42234c6cd9d447c1a497acf03115741e399bcb6aec506ec8a7ee5b53"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-163e5d10702550adecfcd217b7af0a41565cb3ab082a4912857634803ce49d96"></a>

## Direct properties — local_vrf.slo_config.static_v6_routes.static_routes.default_gateway / cbe4ef09c785 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b2cbf958f70cfa2b0bbcd66f38167644a33b11e50dfff6ed698e5c08f7beb4e1"></a>

## Next pages — local_vrf.slo_config.static_v6_routes.static_routes.default_gateway / cbe4ef09c785 / 4

- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-af4854f9eec5985ac007d5b4c4a616a37dfe431fda08d1f8eaf3746fd296f326)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-bd09ea380f6603f6e1fbbe3371f913337387062749568cfd6ed63f02eef6fc12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef62300f3a2e88c7c37bee6fe3d972d63af6ef36dd19abae136ce1c748ac2479"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.node_interface — local_vrf.slo_config.static_v6_routes.static_routes.node_interface / 2923921fa046 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cee7717acf2f93282deb1242c0e8f0c270191bfff1980a588b87c3198f229a2a)
- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-af4854f9eec5985ac007d5b4c4a616a37dfe431fda08d1f8eaf3746fd296f326)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface

<a id="canonical-482b1bb67d7dcf65115103b4109b391a845c629fea72312da8f525dd46f2281d"></a>

Type: `"single"`. Computed.

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

<a id="canonical-d8ed5ddb16077e5962423f6f494b12d1b58bfbecdc9cd9a78b679051c4a7bafb"></a>

## Direct properties — local_vrf.slo_config.static_v6_routes.static_routes.node_interface / 2923921fa046 / 3

- [list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0bf9a4c890336eaf45630ecbbcea03f89962e4c7ca3d2faf0c1c631ceae913ad): complete subsection reference.

<a id="canonical-3f168593f35d2e95e189a7876aff769e02f78364baa4eee470296d463fc0f491"></a>

## Next pages — local_vrf.slo_config.static_v6_routes.static_routes.node_interface / 2923921fa046 / 4

- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0bf9a4c890336eaf45630ecbbcea03f89962e4c7ca3d2faf0c1c631ceae913ad)
- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-af4854f9eec5985ac007d5b4c4a616a37dfe431fda08d1f8eaf3746fd296f326)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-0bf9a4c890336eaf45630ecbbcea03f89962e4c7ca3d2faf0c1c631ceae913ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2e0204da7ca1852de415549b0c396e5074dba529c71c3c0825f7d0e14332567"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list / ac11559aee48 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cee7717acf2f93282deb1242c0e8f0c270191bfff1980a588b87c3198f229a2a)
- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-af4854f9eec5985ac007d5b4c4a616a37dfe431fda08d1f8eaf3746fd296f326)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-bd09ea380f6603f6e1fbbe3371f913337387062749568cfd6ed63f02eef6fc12)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-ad09a6b822fb2669716c745c554dc5092171e33d2c4d8998695062488aaa5379"></a>

Type: `"list"`. Computed.

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

<a id="canonical-11ce88a68c75815d13d9e90a3908dfdc7a258e1989a58d840984a33b86cfd38a"></a>

## Direct properties — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list / ac11559aee48 / 3

- [interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-9c90b341f8a0826d29b985d83a7f4d147e2d4d41e63ef48fb48bb86440903bb5): complete subsection reference.

<a id="canonical-dc0bee2aa40d39832792dbf9a9ec1eef3f463a6fbc5f1e5ff8810ce471518f65"></a>

<a id="canonical-834e53706458e19154ce893065969df783f760f7214358be9545b9d1872c6c04"></a>

## node property — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list / ac11559aee48 / 4

Type: `"string"`. Computed.

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

<a id="canonical-eb9524546ec6f22acc41227c9b5023edd6aca08ba0ab1edc1e9ca72f4d8a1245"></a>

## Next pages — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list / ac11559aee48 / 5

- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-9c90b341f8a0826d29b985d83a7f4d147e2d4d41e63ef48fb48bb86440903bb5)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-bd09ea380f6603f6e1fbbe3371f913337387062749568cfd6ed63f02eef6fc12)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-9c90b341f8a0826d29b985d83a7f4d147e2d4d41e63ef48fb48bb86440903bb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-984706e57e0d56c358d4437d9eca9373ff28fadfd81c949588dbb831b84bf096"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / d6c042072704 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c7c9219c60f0f08d3543c1fe1c22406fec50388bc3e517b5265f68bd87167b25)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-b754167bb4da3d7c1e3473a1d6df496313c7235de86016b384b3ef5aaacd5870)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cee7717acf2f93282deb1242c0e8f0c270191bfff1980a588b87c3198f229a2a)
- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-af4854f9eec5985ac007d5b4c4a616a37dfe431fda08d1f8eaf3746fd296f326)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-011.md#canonical-bd09ea380f6603f6e1fbbe3371f913337387062749568cfd6ed63f02eef6fc12)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0bf9a4c890336eaf45630ecbbcea03f89962e4c7ca3d2faf0c1c631ceae913ad)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-3f0064b246ca3ac7a2b3e7ad29c0a58cfad62a64460244e496d39dedd3012347"></a>

Type: `"list"`. Computed.

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

<a id="canonical-9c94e56968ab5fa59e8f7adf9c05fd67025d172c73b7de0a3b1568696d5f48d0"></a>

## Direct properties — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / d6c042072704 / 3

<a id="canonical-29f0db020ba339e3a2c0bfae6f59fbab7678ffaac9c7daf939751dfd633d0026"></a>

<a id="canonical-1264fc1a459e02395123d099f13ce44d2d174692ad03bcf76a8b1318031e7bf1"></a>

## kind property — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / d6c042072704 / 4

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

<a id="canonical-4e600f6b4623aa15044723460052c852c38e04a88a4b1137246c642bda573b6e"></a>

<a id="canonical-315e68b4dcbc248ef67b3dcb83a399bce0cea2cfb4374f958b4cdb02e34433f7"></a>

## name property — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / d6c042072704 / 5

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

<a id="canonical-a82cca0728060a6eb9ee975373e40b053686182867fd3385cb340bfc9fce9500"></a>

<a id="canonical-b1ff4cf6868c941824719902c61c0e4577db7df5413b311a64434954e0850651"></a>

## namespace property — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / d6c042072704 / 6

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

<a id="canonical-a8949f1b865c227a7dce37a45f1615ea4baf8d4f30f40cd0ce5f3094aaae0357"></a>

<a id="canonical-ddb7c7ecfedcb0fecd8e5a27a4e435d9573baf56c0c739db367637ef1fc8e883"></a>

## tenant property — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / d6c042072704 / 7

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

<a id="canonical-dad6ac85793053c76f0d2d7390a5786e8698f41d8018996b605c08c4d290cf78"></a>

<a id="canonical-6654a471c73b484d2f6b9f832d68a42ad47d7d48f3cb4d1b4c30dbd824afa24a"></a>

## uid property — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / d6c042072704 / 8

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

<a id="canonical-ed0cbb07c05b742cb630e92df46077130caf187a7e9e6b985d9040db7b658929"></a>

## Next pages — local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interfac / d6c042072704 / 9

- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0bf9a4c890336eaf45630ecbbcea03f89962e4c7ca3d2faf0c1c631ceae913ad)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-cb521801b7b2b15af35a5400de6226f064cc01542e4f27303d5c4a4313dd55be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e87e7f64a27c3aa1ec051c44914749c93eb234d204d14b85ad2d9f24f14b45e"></a>

## log_receiver_with_net — log_receiver_with_net / 917f38cf15bc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- log_receiver_with_net

<a id="canonical-319530c937c39d8f98ec1be68c38e2e933cd0c2faa1f2f66757b3b4b5b8544ac"></a>

Type: `"single"`. Computed.

\[OneOf: log\_receiver\_with\_net, logs\_streaming\_disabled\] Select log receiver for logs
streaming with network option.

Upstream description:

Select log receiver for logs streaming with network option.

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

- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-011.md#canonical-319530c937c39d8f98ec1be68c38e2e933cd0c2faa1f2f66757b3b4b5b8544ac)
- [logs_streaming_disabled](data-sources--securemesh_site_v2--reference--group-011.md#canonical-17066e1e628011241463dde6dfc2f4a4e7bad6188a4e6b13a29ac1c31e72f370)

Select alternatives according to the provider validators above.

<a id="canonical-b5739b511fa52c5940cd55ca81e391f537dd033698b0d9ca8cdaa83bc35e74be"></a>

## Direct properties — log_receiver_with_net / 917f38cf15bc / 3

- [log_receiver](data-sources--securemesh_site_v2--reference--group-011.md#canonical-253e089a5b459219fc506a9c07d6220f149dafb1b115fbdc288888019854e01b): complete subsection reference.

- [use_management_network](data-sources--securemesh_site_v2--reference--group-011.md#canonical-91d570d8864297394f21acfaf3dbbd924ada9b6cbdc72b3b90956e31ba58fc33): complete subsection reference.

- [use_slo_sli](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c8f01ce4773434ed85961e6bf8f322545058a023b8b483138b5b92454a7662ae): complete subsection reference.

<a id="canonical-9e6863e4ca067d8ac143e580c8659d1344225314a151ec2d3f53dbf669a1d018"></a>

## Next pages — log_receiver_with_net / 917f38cf15bc / 4

- [log_receiver_with_net.log_receiver](data-sources--securemesh_site_v2--reference--group-011.md#canonical-253e089a5b459219fc506a9c07d6220f149dafb1b115fbdc288888019854e01b)
- [log_receiver_with_net.use_management_network](data-sources--securemesh_site_v2--reference--group-011.md#canonical-91d570d8864297394f21acfaf3dbbd924ada9b6cbdc72b3b90956e31ba58fc33)
- [log_receiver_with_net.use_slo_sli](data-sources--securemesh_site_v2--reference--group-011.md#canonical-c8f01ce4773434ed85961e6bf8f322545058a023b8b483138b5b92454a7662ae)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-253e089a5b459219fc506a9c07d6220f149dafb1b115fbdc288888019854e01b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c34f216f6ba5979b8e731bd5cc608e1060f639ba88a2a59fe42547a3456345b0"></a>

## log_receiver_with_net.log_receiver — log_receiver_with_net.log_receiver / 9fe675038c46 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cb521801b7b2b15af35a5400de6226f064cc01542e4f27303d5c4a4313dd55be)
- log_receiver_with_net.log_receiver

<a id="canonical-65aeecfc46572b3c3b3348a6c072321df6cebe8e5fc058f72eb8716532b37e3e"></a>

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

<a id="canonical-91aef2e0e026f89ac0aab92d4e5656e5fad90586fed6a6acb74504968d470a7d"></a>

## Direct properties — log_receiver_with_net.log_receiver / 9fe675038c46 / 3

<a id="canonical-8d8a31b16b915cedac74ec655c98e93785c2f070799d316eea1b02b6374098a2"></a>

<a id="canonical-22a5d08ab0125d5405bf0ee48415e9d49aec292fb33083dfb1e2b8556c74921e"></a>

## name property — log_receiver_with_net.log_receiver / 9fe675038c46 / 4

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

<a id="canonical-7e924399c796b5dcded06697f55076f61f6e47a2379a36ba49489fdefb898508"></a>

<a id="canonical-494d143d6b97352e5d162f49a25736b33eae6d2384036b76f2d6451a6def7269"></a>

## namespace property — log_receiver_with_net.log_receiver / 9fe675038c46 / 5

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

<a id="canonical-bbd3e3774d7046b15f5d6ea13abf5229263cc0da5795808e5859259c3fe924c2"></a>

<a id="canonical-2e3c27e0a22699d3adce0f0e37ffd51e4a2d2547aeaa2927fba06f5102adf494"></a>

## tenant property — log_receiver_with_net.log_receiver / 9fe675038c46 / 6

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

<a id="canonical-d59078c06e6d2c4a364c41b111df425af309179ca31b2e8c7675ca78f950b663"></a>

## Next pages — log_receiver_with_net.log_receiver / 9fe675038c46 / 7

- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cb521801b7b2b15af35a5400de6226f064cc01542e4f27303d5c4a4313dd55be)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-91d570d8864297394f21acfaf3dbbd924ada9b6cbdc72b3b90956e31ba58fc33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-527a4f730bde536859c1e8c0e6aebaab79591e74437f3dd949ef9a34cb973d1b"></a>

## log_receiver_with_net.use_management_network — log_receiver_with_net.use_management_network / 856385ca5321 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cb521801b7b2b15af35a5400de6226f064cc01542e4f27303d5c4a4313dd55be)
- log_receiver_with_net.use_management_network

<a id="canonical-aebc6e6aa782a184c4b14ab32736c70fb199a5d43acbf1af7d4debe93e6505c3"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1f55cf7ca1f2adc2c73d8d9d25c8013b7cf946251eb504401ac9cc93692bb484"></a>

## Direct properties — log_receiver_with_net.use_management_network / 856385ca5321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7505137305ad53873ae87c7977853f9d424626774f1f1d974e4e437f85629de8"></a>

## Next pages — log_receiver_with_net.use_management_network / 856385ca5321 / 4

- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cb521801b7b2b15af35a5400de6226f064cc01542e4f27303d5c4a4313dd55be)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-c8f01ce4773434ed85961e6bf8f322545058a023b8b483138b5b92454a7662ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a77a63f79141138c289f94e7ac32be8ebcea11207e75f2b351ebf325831677f"></a>

## log_receiver_with_net.use_slo_sli — log_receiver_with_net.use_slo_sli / 8448e6bc15c8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cb521801b7b2b15af35a5400de6226f064cc01542e4f27303d5c4a4313dd55be)
- log_receiver_with_net.use_slo_sli

<a id="canonical-a23969d4cee16defb06176e72013f14713402189923ecb2890ff5f7d5b5f1110"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-56958245c1d57d0be97dfe022feaf5a5ecf48b2de01f2f09183ed38da9438991"></a>

## Direct properties — log_receiver_with_net.use_slo_sli / 8448e6bc15c8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9f1024c418e8bb83ffa054b108440d397fc8dad63eb3bb866b3a24065d91bf43"></a>

## Next pages — log_receiver_with_net.use_slo_sli / 8448e6bc15c8 / 4

- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-011.md#canonical-cb521801b7b2b15af35a5400de6226f064cc01542e4f27303d5c4a4313dd55be)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-f6fc226cb76785de226517bf48defedf6623821c38740083e7a7933885135110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14313d612096154dcf97d695259d3af0a8159d6cf121f98a2d80756210501ce4"></a>

## logs_streaming_disabled — logs_streaming_disabled / 6de19fb36e29 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- logs_streaming_disabled

<a id="canonical-17066e1e628011241463dde6dfc2f4a4e7bad6188a4e6b13a29ac1c31e72f370"></a>

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

<a id="canonical-666ca486f69e5851c4260040309da7222d7f7988ae30c7b4ba4e1ea9c36f4c04"></a>

## Direct properties — logs_streaming_disabled / 6de19fb36e29 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64f34b3a7ce96d6454323a3bde2ce746fde4a61b080378b2a667b5d98c2d6e78"></a>

## Next pages — logs_streaming_disabled / 6de19fb36e29 / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-46ac4384a3ca0b62d8075b9c924d0fcfd2957353fe73c01a0f81c5acc436f882)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-cf134a3bd4d44ad284286867a84f381180caf1b4d78022ad5b1b080803ffc441)

<a id="canonical-2ffe5fefb62142719b26f3c6c44bede4eddcfc1859f2c8d964484226f2f01664"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
