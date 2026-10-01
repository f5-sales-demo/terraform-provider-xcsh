---
page_title: "xcsh_udp_loadbalancer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer reference."
---

# xcsh_udp_loadbalancer reference

<a id="canonical-cc8bd872d4b0a168dd5cf6a214a1cdd7218374adfe2f7ca0d9456fce112d1d6f"></a>

## create property — timeouts / 6a0f045370d9 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-aad5010b11f3777368a49591eae0ffe43b7cae41a9afe936f245bb40f8e53127"></a>

<a id="canonical-5167f01bbc9e1104ded1a16e2df15a3d6618e3d481288c2fcd723f4fd382dccd"></a>

## delete property — timeouts / 6a0f045370d9 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-a008e7da5d3bc52306de89977122267bae3338412452280fae0f0b63e9eec139"></a>

<a id="canonical-1e6ee1264fac7a77fbb0f814d8e3c3d41ce3892aea05d39e81da008433658f02"></a>

## read property — timeouts / 6a0f045370d9 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-717d775e3bf523c3de93b50b3301a6afe937251ea478f596a8c52b54f4a674f4"></a>

<a id="canonical-7df0efce4bb28ecb7f97df09d16d1193864289a8ec24caf2ba1fc1996aa70207"></a>

## update property — timeouts / 6a0f045370d9 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f5a1bcb0258f4eea08eaba5714432be0e4bc8a106d0f2ca7907c20c110325605"></a>

## Next pages — timeouts / 6a0f045370d9 / 8

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-d95efc1645c481a1dff9b350771ac01bbb10588a5a3ef681c1dce0e98f8af44c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50c757992c88b7094c4aa2e9de6c16f9d89274b876f06dfb5310e9c739b97dab"></a>

## udp — udp / 2de1fd40d785 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- udp

<a id="canonical-679f3148a643ae01c68df2c6228142bc4fbde22a45250bdfa3beb537fbb80270"></a>

Type: `"object"`. single nested block, Optional.

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
udp {}
```

<a id="canonical-a557202080f226f53ec20b9d5db051f0cc468b0456117c46de70ef5f804ba9c3"></a>

## Direct properties — udp / 2de1fd40d785 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3f0bcbbd586211c08dc06492c65bbdf641df3831f444de3a27aaa1c328bf5703"></a>

## Next pages — udp / 2de1fd40d785 / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
