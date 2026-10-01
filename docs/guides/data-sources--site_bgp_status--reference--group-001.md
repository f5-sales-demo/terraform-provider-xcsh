---
page_title: "xcsh_site_bgp_status reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_bgp_status reference."
---

# xcsh_site_bgp_status reference

<a id="canonical-b3d7a08e010462cf8b9e36f2952a57bc79b1f43a2d6740a60d4ab1649bb5cb27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-057d1fb318afbcf8a2bcf36c60c53389c2c8fca48bbf2ba530709b37e36955c7"></a>

## Property reference — Property reference / b6a3c3a053fc / 2

Breadcrumbs:

- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-e2748d411a0ce0890ae363c44505ececdca5700e631692a6a67cd6b16ab5132d)
- Property reference

<a id="canonical-db339cdcca7d4f8bee4301cc7bdf990d7328e7b398afbb2f10f009b9a1320eae"></a>

## Direct properties — Property reference / b6a3c3a053fc / 3

<a id="canonical-ccdc64cc03db68ed1948c59845f36555a29a5df36d98f3de850a1109f3acf90a"></a>

<a id="canonical-b102434b7d582d53af2ccc8400abcaf45c956c1cfe962539a3e721ff9b74befe"></a>

## bgp_routes_json property — Property reference / b6a3c3a053fc / 4

Type: `"string"`. Computed.

<a id="canonical-161b6da6de79aacd2de9ddff1756c08141b2c9005dbbecb754445ea37c0814df"></a>

<a id="canonical-2c103bc638b488588e654b3335d7091c0ade21c63cd637f0794ce64fb16e1e38"></a>

## converged property — Property reference / b6a3c3a053fc / 5

Type: `"bool"`. Computed.

<a id="canonical-be554a01a2ba988993ac275db49e6b858ab905d712cad923b0a71df917e71950"></a>

<a id="canonical-e20287cd24f0d3291266b9a0d447630d6b2eb4e3968d6f119df264094dca5540"></a>

## expected_exported_routes property — Property reference / b6a3c3a053fc / 6

Type: `["set", "string"]`. Required.

Exact prefixes that every expected node must export and carry in the selected SLO or SLI route view.

- [expected_peers](data-sources--site_bgp_status--reference--group-001.md#canonical-3ad3f30f9107724dcee3f0be9bfbfcb2da9755a223fab57030259ef8bdc63cce): complete subsection reference.

<a id="canonical-80da3077b51ce1a4159a456b7217bc5b186cb32dbec5f31fbb9fcd6a4bc691fc"></a>

<a id="canonical-7e0e9e0dd2cc4c95be644d984f9af5fa7ded6fb472fa7d90b88bbd8b680712e6"></a>

## id property — Property reference / b6a3c3a053fc / 7

Type: `"string"`. Computed.

<a id="canonical-adba78bad6c386f957d42aa87cff6db643eb601f5110292c00197b17997dfae1"></a>

<a id="canonical-39d6afb43629c64f4c7dc0d4aa7580bea74045ae32fdf23404d96a77a4b36013"></a>

## namespace property — Property reference / b6a3c3a053fc / 8

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.OneOf("system")}
```

- [peers](data-sources--site_bgp_status--reference--group-001.md#canonical-f6f9f0a95ad134f97f1ebf048904042d26508976abc8abd522b37ef4149ad59b): complete subsection reference.

<a id="canonical-5e29f5108a04042394122fcafe1655aa30c6a36e7c979c81c0da7106ce01cdd9"></a>

<a id="canonical-e1967b77f9a4a320b657a38e6f91ae140938ce72b19596098cca5f138165706e"></a>

## poll_interval_seconds property — Property reference / b6a3c3a053fc / 9

Type: `"number"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 60)}
```

<a id="canonical-f707f0587419c027c8a8f829114f44ca336ae3a376fba37affdce56cde37d3da"></a>

<a id="canonical-8ce99a761914ae1b8f4e5d54181e86b3ea881bc08122b57095fa702023020c70"></a>

## site property — Property reference / b6a3c3a053fc / 10

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="canonical-42dd903a9d278735f7da130918c7a7902dad44b3b920c04f82ef222593d9e2e9"></a>

<a id="canonical-4a318976677f3ebb30bc633eab9a47453f1f2be53eaab29b444b5a084010a4e7"></a>

## sli_routes_json property — Property reference / b6a3c3a053fc / 11

Type: `"string"`. Computed.

<a id="canonical-7b5f389564d5d27ec94edfa4d1300bbe0ea9c2e9f853c3532c717c878cd365a4"></a>

<a id="canonical-39cefb5fc07292dd9f000be76254602a92029d93b208c345aac15824f5f386dd"></a>

## slo_routes_json property — Property reference / b6a3c3a053fc / 12

Type: `"string"`. Computed.

<a id="canonical-ccbdb432b0868fbb52526986c189371a0d1a4b4e90276656b489ee898c664a58"></a>

<a id="canonical-b6515f64f6884c5383a2a4a80d6a0a766aa7e99c03053c670d49f6a8fd74b7e8"></a>

## timeout_seconds property — Property reference / b6a3c3a053fc / 13

Type: `"number"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 1800)}
```

<a id="canonical-b8998b0c26de06ff689bd51b67f00d99be0cf391172bf2b3eb31f16af0657067"></a>

## All schema paths — Property reference / b6a3c3a053fc / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `bgp_routes_json` | [bgp_routes_json](data-sources--site_bgp_status--reference--group-001.md#canonical-ccdc64cc03db68ed1948c59845f36555a29a5df36d98f3de850a1109f3acf90a) |
| `converged` | [converged](data-sources--site_bgp_status--reference--group-001.md#canonical-161b6da6de79aacd2de9ddff1756c08141b2c9005dbbecb754445ea37c0814df) |
| `expected_exported_routes` | [expected_exported_routes](data-sources--site_bgp_status--reference--group-001.md#canonical-be554a01a2ba988993ac275db49e6b858ab905d712cad923b0a71df917e71950) |
| `expected_peers` | [expected_peers](data-sources--site_bgp_status--reference--group-001.md#canonical-6a09c7bee93853d0470cefaf49ffb55539676e60d73c6cb23cb86c41ec4a9c23) |
| `expected_peers.expected_imported_routes` | [expected_peers.expected_imported_routes](data-sources--site_bgp_status--reference--group-001.md#canonical-be0f8d08009f2aa97dab7bea12ebb5a06befaf7abbfa7d889ff178f77b6360c9) |
| `expected_peers.mac` | [expected_peers.mac](data-sources--site_bgp_status--reference--group-001.md#canonical-351a5d3d8d115fc4cba3971e3cd9237c88ceab3352c3e9e8cf92df1f8010c5f6) |
| `expected_peers.node` | [expected_peers.node](data-sources--site_bgp_status--reference--group-001.md#canonical-6989325e5893d6b2d551c5c209f4782d60e88e87c74fd9880c41b94d26c35421) |
| `expected_peers.peer_address` | [expected_peers.peer_address](data-sources--site_bgp_status--reference--group-001.md#canonical-8e216b53f03bb33c7854961b714f4916f518e165540a9c5b493986ad081b0bce) |
| `expected_peers.role` | [expected_peers.role](data-sources--site_bgp_status--reference--group-001.md#canonical-ee4b85b466a9c812776d065714f7910825e40143bb12d3dda1ba3f31dacbff59) |
| `id` | [id](data-sources--site_bgp_status--reference--group-001.md#canonical-80da3077b51ce1a4159a456b7217bc5b186cb32dbec5f31fbb9fcd6a4bc691fc) |
| `namespace` | [namespace](data-sources--site_bgp_status--reference--group-001.md#canonical-adba78bad6c386f957d42aa87cff6db643eb601f5110292c00197b17997dfae1) |
| `peers` | [peers](data-sources--site_bgp_status--reference--group-001.md#canonical-b71f536392474871dda31f168f1c62c3a4d8bb0e23c34c8da8ac70ac797ff3ec) |
| `peers.advertised_prefix_count` | [peers.advertised_prefix_count](data-sources--site_bgp_status--reference--group-001.md#canonical-2fe61bd7679491b05866601b9c3f670afd1d34ae58ce8671cece48e5f9f61bad) |
| `peers.established` | [peers.established](data-sources--site_bgp_status--reference--group-001.md#canonical-418e8285a33748e64dfd2a5d6cd9133f1038f6c3f315a6bc905a3dda50c2cf39) |
| `peers.interface_name` | [peers.interface_name](data-sources--site_bgp_status--reference--group-001.md#canonical-1487ad63edec0f4f7e9ae809f1941266e128d736b11b246df259cb2382af23a3) |
| `peers.mac` | [peers.mac](data-sources--site_bgp_status--reference--group-001.md#canonical-f907061eee3a0a6215b6209bbb091d4efa7cb0d0fe6d53633b26512d7fb76c5a) |
| `peers.node` | [peers.node](data-sources--site_bgp_status--reference--group-001.md#canonical-04ecf2bff1a57ca22bea81059c98bc7c00d6ab549e2d2968d6a0c2500e964f6f) |
| `peers.peer_address` | [peers.peer_address](data-sources--site_bgp_status--reference--group-001.md#canonical-f14371ee900229a17918fad590e7ddeabee88cfa5015154b9908d197ae9121f8) |
| `peers.received_prefix_count` | [peers.received_prefix_count](data-sources--site_bgp_status--reference--group-001.md#canonical-2abd78c531044aca96101569ccf39cddeebe81987d02bb71fff10b9201c55e2a) |
| `peers.role` | [peers.role](data-sources--site_bgp_status--reference--group-001.md#canonical-efa116bdba4f1d5d41b3d380041b3a97a05c015acae393e32889de3a8165de39) |
| `peers.state` | [peers.state](data-sources--site_bgp_status--reference--group-001.md#canonical-28319e12db8c2b26b185d19a95c769dbb5b5bd738dfd0adce15d0f4cb3cbc85b) |
| `peers.state_changed_at` | [peers.state_changed_at](data-sources--site_bgp_status--reference--group-001.md#canonical-451dd444d501ee2f01a40bfbc5d2187aacb7adadb04f9e1433c2982a5e9c5ee0) |
| `poll_interval_seconds` | [poll_interval_seconds](data-sources--site_bgp_status--reference--group-001.md#canonical-5e29f5108a04042394122fcafe1655aa30c6a36e7c979c81c0da7106ce01cdd9) |
| `site` | [site](data-sources--site_bgp_status--reference--group-001.md#canonical-f707f0587419c027c8a8f829114f44ca336ae3a376fba37affdce56cde37d3da) |
| `sli_routes_json` | [sli_routes_json](data-sources--site_bgp_status--reference--group-001.md#canonical-42dd903a9d278735f7da130918c7a7902dad44b3b920c04f82ef222593d9e2e9) |
| `slo_routes_json` | [slo_routes_json](data-sources--site_bgp_status--reference--group-001.md#canonical-7b5f389564d5d27ec94edfa4d1300bbe0ea9c2e9f853c3532c717c878cd365a4) |
| `timeout_seconds` | [timeout_seconds](data-sources--site_bgp_status--reference--group-001.md#canonical-ccbdb432b0868fbb52526986c189371a0d1a4b4e90276656b489ee898c664a58) |

<a id="canonical-0cbc963a1b3d30723d7f5767d8097ca3762a0c9134fa3437f9c914d31e3b5b56"></a>

## Next pages — Property reference / b6a3c3a053fc / 15

- [expected_peers](data-sources--site_bgp_status--reference--group-001.md#canonical-3ad3f30f9107724dcee3f0be9bfbfcb2da9755a223fab57030259ef8bdc63cce)
- [peers](data-sources--site_bgp_status--reference--group-001.md#canonical-f6f9f0a95ad134f97f1ebf048904042d26508976abc8abd522b37ef4149ad59b)
- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-e2748d411a0ce0890ae363c44505ececdca5700e631692a6a67cd6b16ab5132d)

<a id="canonical-3ad3f30f9107724dcee3f0be9bfbfcb2da9755a223fab57030259ef8bdc63cce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f458e02672fb7e72a4a4d47d53270cc6395fe6d1ae54f62be8e7b6e4e50addec"></a>

## expected_peers — expected_peers / 2ccc0984144f / 2

Breadcrumbs:

- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-e2748d411a0ce0890ae363c44505ececdca5700e631692a6a67cd6b16ab5132d)
- [Property reference](data-sources--site_bgp_status--reference--group-001.md#canonical-b3d7a08e010462cf8b9e36f2952a57bc79b1f43a2d6740a60d4ab1649bb5cb27)
- expected_peers

<a id="canonical-6a09c7bee93853d0470cefaf49ffb55539676e60d73c6cb23cb86c41ec4a9c23"></a>

Type: `"map"`. Required.

<a id="canonical-4a71d52c8e0234bab9d7e190f95af7ce55c019114828cd9c5ac91f440584dd2d"></a>

## Direct properties — expected_peers / 2ccc0984144f / 3

<a id="canonical-be0f8d08009f2aa97dab7bea12ebb5a06befaf7abbfa7d889ff178f77b6360c9"></a>

<a id="canonical-a7fb0985beb156b1bffe6d363a7eadf68bfe4d774d7f343ff425bd5ef4a376cc"></a>

## expected_imported_routes property — expected_peers / 2ccc0984144f / 4

Type: `["set", "string"]`. Required.

Exact prefixes that must be imported from this remote peer on the expected node.

<a id="canonical-351a5d3d8d115fc4cba3971e3cd9237c88ceab3352c3e9e8cf92df1f8010c5f6"></a>

<a id="canonical-ca514b977059915e6e95af15c62b22048fee31ec505e3f90f0785c66d6d6f94c"></a>

## mac property — expected_peers / 2ccc0984144f / 5

Type: `"string"`. Required.

<a id="canonical-6989325e5893d6b2d551c5c209f4782d60e88e87c74fd9880c41b94d26c35421"></a>

<a id="canonical-ee65a36260263ff442ea56029076600920b8e3a739197fca95657a27a7b89345"></a>

## node property — expected_peers / 2ccc0984144f / 6

Type: `"string"`. Required.

<a id="canonical-8e216b53f03bb33c7854961b714f4916f518e165540a9c5b493986ad081b0bce"></a>

<a id="canonical-23663cbe7c3e031bac624dbd3bb24448f106ec74039a3cd48363c56a95999f9f"></a>

## peer_address property — expected_peers / 2ccc0984144f / 7

Type: `"string"`. Required.

<a id="canonical-ee4b85b466a9c812776d065714f7910825e40143bb12d3dda1ba3f31dacbff59"></a>

<a id="canonical-3477b3140caf55a550c91b3d5a44e02e1dea36822b9802f99507cdf110bcf46b"></a>

## role property — expected_peers / 2ccc0984144f / 8

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.OneOf("slo",
    "sli")}
```

<a id="canonical-9684571fb788bb49bf27306b9a5a0b6af04ab83f589acca63f056b1c7a2cb0fc"></a>

## Next pages — expected_peers / 2ccc0984144f / 9

- [Property reference](data-sources--site_bgp_status--reference--group-001.md#canonical-b3d7a08e010462cf8b9e36f2952a57bc79b1f43a2d6740a60d4ab1649bb5cb27)
- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-e2748d411a0ce0890ae363c44505ececdca5700e631692a6a67cd6b16ab5132d)

<a id="canonical-f6f9f0a95ad134f97f1ebf048904042d26508976abc8abd522b37ef4149ad59b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed9830e8ccdc8b85569320c613af105785b9a60915de759cb7b261099e705334"></a>

## peers — peers / 6a98b7bd4aeb / 2

Breadcrumbs:

- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-e2748d411a0ce0890ae363c44505ececdca5700e631692a6a67cd6b16ab5132d)
- [Property reference](data-sources--site_bgp_status--reference--group-001.md#canonical-b3d7a08e010462cf8b9e36f2952a57bc79b1f43a2d6740a60d4ab1649bb5cb27)
- peers

<a id="canonical-b71f536392474871dda31f168f1c62c3a4d8bb0e23c34c8da8ac70ac797ff3ec"></a>

Type: `"map"`. Computed.

<a id="canonical-c4ad1837640bc29abf6122fa841b6ee5e35e4c4c35a196f3c70995d8410e27da"></a>

## Direct properties — peers / 6a98b7bd4aeb / 3

<a id="canonical-2fe61bd7679491b05866601b9c3f670afd1d34ae58ce8671cece48e5f9f61bad"></a>

<a id="canonical-2971cfd25e476c5e6d08f7562105f8e43cf419c9f9452d24ba2128fabf069749"></a>

## advertised_prefix_count property — peers / 6a98b7bd4aeb / 4

Type: `"number"`. Computed.

<a id="canonical-418e8285a33748e64dfd2a5d6cd9133f1038f6c3f315a6bc905a3dda50c2cf39"></a>

<a id="canonical-398f6afbf00b558ad88e539d113ef32516b72f51622dacf4f6814352fc207acf"></a>

## established property — peers / 6a98b7bd4aeb / 5

Type: `"bool"`. Computed.

<a id="canonical-1487ad63edec0f4f7e9ae809f1941266e128d736b11b246df259cb2382af23a3"></a>

<a id="canonical-18fdfa90259bba32bcc9b03d0138e518484533e27a286b05d5952661fdf1d582"></a>

## interface_name property — peers / 6a98b7bd4aeb / 6

Type: `"string"`. Computed.

<a id="canonical-f907061eee3a0a6215b6209bbb091d4efa7cb0d0fe6d53633b26512d7fb76c5a"></a>

<a id="canonical-cc130e70d0a9e2bb96afcf8cb2fc5764185f349fa7534b9db6e52ca92188606a"></a>

## mac property — peers / 6a98b7bd4aeb / 7

Type: `"string"`. Computed.

<a id="canonical-04ecf2bff1a57ca22bea81059c98bc7c00d6ab549e2d2968d6a0c2500e964f6f"></a>

<a id="canonical-4b5a459f9dd9ef69983d5885ecace702db9ccf58d09b5ce1bed777b8785872d8"></a>

## node property — peers / 6a98b7bd4aeb / 8

Type: `"string"`. Computed.

<a id="canonical-f14371ee900229a17918fad590e7ddeabee88cfa5015154b9908d197ae9121f8"></a>

<a id="canonical-08f1b7559b733fa07deac475364e74e5ec53e72d62ee8f5e4bfa6ab78b7eaaaa"></a>

## peer_address property — peers / 6a98b7bd4aeb / 9

Type: `"string"`. Computed.

<a id="canonical-2abd78c531044aca96101569ccf39cddeebe81987d02bb71fff10b9201c55e2a"></a>

<a id="canonical-5623a91a840816439a014d7d7f6bea8bc4583784f5a49f0c401e2878d00a3431"></a>

## received_prefix_count property — peers / 6a98b7bd4aeb / 10

Type: `"number"`. Computed.

<a id="canonical-efa116bdba4f1d5d41b3d380041b3a97a05c015acae393e32889de3a8165de39"></a>

<a id="canonical-b5925ab58fc056a4cd5ff4c18efa9dc287247c5510277f4b2ba9b0dc91730a68"></a>

## role property — peers / 6a98b7bd4aeb / 11

Type: `"string"`. Computed.

<a id="canonical-28319e12db8c2b26b185d19a95c769dbb5b5bd738dfd0adce15d0f4cb3cbc85b"></a>

<a id="canonical-fa49ba62dc7212da521772a3aba0e17cbea1bde310a281f0641dabd1727b7d75"></a>

## state property — peers / 6a98b7bd4aeb / 12

Type: `"string"`. Computed.

<a id="canonical-451dd444d501ee2f01a40bfbc5d2187aacb7adadb04f9e1433c2982a5e9c5ee0"></a>

<a id="canonical-ff0b4ba894946708605d6fdf6b2e63f3e507d169ebec35d495ee572169d5e023"></a>

## state_changed_at property — peers / 6a98b7bd4aeb / 13

Type: `"string"`. Computed.

<a id="canonical-718cf564e91306d15703ea97a6fc859bef185a2d682da89bd3f32307257d8cb2"></a>

## Next pages — peers / 6a98b7bd4aeb / 14

- [Property reference](data-sources--site_bgp_status--reference--group-001.md#canonical-b3d7a08e010462cf8b9e36f2952a57bc79b1f43a2d6740a60d4ab1649bb5cb27)
- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md#canonical-e2748d411a0ce0890ae363c44505ececdca5700e631692a6a67cd6b16ab5132d)
