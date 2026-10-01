---
page_title: "xcsh_bot_peer_traffic_overview reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_traffic_overview reference."
---

# xcsh_bot_peer_traffic_overview reference

<a id="canonical-53367a814368233d59c0d60053b9993843031752de8473de41d0df93b435773a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fd02542e6266c840199e2f3b98ecec21e97708cbadf640e3c61b914a6cbd751"></a>

## Property reference — Property reference / 819c50e1e319 / 2

Breadcrumbs:

- [xcsh_bot_peer_traffic_overview](../data-sources/bot_peer_traffic_overview.md#canonical-5b05326e4d591c752fdaa94722e4badaea615d250171e751c47a2266294991d7)
- Property reference

<a id="canonical-0d10112057fc18425e97870fdc932ffd1ac23d59b2a07a2718a3099fb90f65a7"></a>

## Direct properties — Property reference / 819c50e1e319 / 3

- [details](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-3ff6c2003052e5c0e764bfcdc89c57174450b2e337057daa66e9edc9142fd041): complete subsection reference.

<a id="canonical-ba35f57439787b58889b01a54dc67dc6ea2338c09ac90b1108199bec07717740"></a>

<a id="canonical-7cb2aeeb82aeaa65c51294c59016a25ff53126973b23a0ff450229de320998c0"></a>

## end_time property — Property reference / 819c50e1e319 / 4

Type: `"string"`. Optional.

End Time. End time of the query period.

<a id="canonical-aec2916674dc0f03316926742c44c2c3a10c651216aee9306811a3c5ae9d6b1a"></a>

<a id="canonical-0e009f24807a7275410a1071a25a2e427bdb17db7337032caeba65277bc9341c"></a>

## limit property — Property reference / 819c50e1e319 / 5

Type: `"number"`. Optional.

Limits the number of transactions returned in the response Optional: If not specified (with default
value 0), all transactions that match the query will be returned in the response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 500),
}
```

<a id="canonical-e5cd4e5991280ebc078dad4b7c932a2ea89ed6fcef9384b73b0efbaa512f868c"></a>

<a id="canonical-8453054afbb6644a4d76345557024b6c718693e4171e33fbdf6a95c005fca2e1"></a>

## namespace property — Property reference / 819c50e1e319 / 6

Type: `"string"`. Required.

Namespace. namespace is used to scope the query. Only virtual\_host in given namespace will be
considered.

<a id="canonical-42fac99e6a78b0b6dd781ded4cc0168f04dada00284288eb499b6887ddfb9076"></a>

<a id="canonical-58b692c538af9dab33e9b7b35d4b52319b05152d0e09485b81f9f61e0d12547e"></a>

## peer_percentage property — Property reference / 819c50e1e319 / 7

Type: `"number"`. Computed.

Peer Percentage. The Peer Percentage of the Item.

<a id="canonical-b69a4e80905d778c00fad32363aaf8cddafccbc6020ce55487ce5bebb11ac01d"></a>

<a id="canonical-035fec62a3c84c5d61e501366f5128213d86060d59ea43a775d9c2f39fd637dd"></a>

## peer_total property — Property reference / 819c50e1e319 / 8

Type: `"string"`. Computed.

Peer Total. The total number of Peer of the Item.

<a id="canonical-43476f22d0e32bdf342353d0b5fda35deb8287c8f0cb8ba4874cdff490c03b72"></a>

<a id="canonical-7337bf70f5fdd3b98ab3e8f81ceb71bf757de6b00e459b8dc950f136716da343"></a>

## rank_by property — Property reference / 819c50e1e319 / 9

Type: `"string"`. Optional.

\[Enum: SELF|PEER\] Or 'PEER' Key for ranking - SELF: Rank by self Use Self as key to query Use Peer
as key to query. Possible values are \`SELF\`, \`PEER\`. Defaults to \`SELF\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SELF",
    "PEER"),
}
```

<a id="canonical-fe2dac3b9196ab8f34293441d19bd3c5944a538ed060a178dd230dbc13a6ba31"></a>

<a id="canonical-4d9f6da20ce0e9f61e4a6976725eaa107857c177962a2842b66486087d6a3e83"></a>

## self_percentage property — Property reference / 819c50e1e319 / 10

Type: `"number"`. Computed.

Self Percentage. The Self Percentage of the Item.

<a id="canonical-fc79550a63f1d37206531799fcc703a731d34b4495831367ffc545fe711915c0"></a>

<a id="canonical-47b2c9561dc4d2d96a31a1ca0bc65b881674351d1e1f15f80eaff3a1af369c88"></a>

## self_total property — Property reference / 819c50e1e319 / 11

Type: `"string"`. Computed.

Self Total. The total number of Self of the item.

<a id="canonical-1f8f3db1266d674e70ec30401ebb18251b998d784279d07070407d4ee42a2f91"></a>

<a id="canonical-a77a71344786e08384723e57b05cfabfdc7424ae613bf6580f952c7dbd0ca882"></a>

## start_time property — Property reference / 819c50e1e319 / 12

Type: `"string"`. Optional.

Start Time. Start time of the query period.

<a id="canonical-f982f25bd6f97714d6e2e5552c93bade28d9113d2f60530c16495a52d91412aa"></a>

## All schema paths — Property reference / 819c50e1e319 / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `details` | [details](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-d5d8b22929f87befb3da95fb63f149e466809b41cef2cd076d93b22d64e98d8b) |
| `details.name` | [details.name](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-d89e028e82c7885b55a4fe97deb1a1223ebe9fcbada0b79cfea151040c0c952d) |
| `details.peer_count` | [details.peer_count](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-94d7182761e3644e6611ecffe085996a40b62ab317c1d8c96f01bd870585df8e) |
| `details.peer_percentage` | [details.peer_percentage](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-2142e1c0b71a59f8aaffea4a347e0afba3d74c2e36ca4e9156a0bded77ce17d1) |
| `details.self_count` | [details.self_count](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-9e5be1f493b8c85f046135d93d6289a1adffde83b9c565701f7f12257d4e6dac) |
| `details.self_percentage` | [details.self_percentage](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-0a9e26d509c429a0e80e9148738a8a415f835b21e2d3fd631e3bf85340997f07) |
| `end_time` | [end_time](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-ba35f57439787b58889b01a54dc67dc6ea2338c09ac90b1108199bec07717740) |
| `limit` | [limit](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-aec2916674dc0f03316926742c44c2c3a10c651216aee9306811a3c5ae9d6b1a) |
| `namespace` | [namespace](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-e5cd4e5991280ebc078dad4b7c932a2ea89ed6fcef9384b73b0efbaa512f868c) |
| `peer_percentage` | [peer_percentage](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-42fac99e6a78b0b6dd781ded4cc0168f04dada00284288eb499b6887ddfb9076) |
| `peer_total` | [peer_total](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-b69a4e80905d778c00fad32363aaf8cddafccbc6020ce55487ce5bebb11ac01d) |
| `rank_by` | [rank_by](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-43476f22d0e32bdf342353d0b5fda35deb8287c8f0cb8ba4874cdff490c03b72) |
| `self_percentage` | [self_percentage](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-fe2dac3b9196ab8f34293441d19bd3c5944a538ed060a178dd230dbc13a6ba31) |
| `self_total` | [self_total](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-fc79550a63f1d37206531799fcc703a731d34b4495831367ffc545fe711915c0) |
| `start_time` | [start_time](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-1f8f3db1266d674e70ec30401ebb18251b998d784279d07070407d4ee42a2f91) |

<a id="canonical-17646cc4d2cf28e0df4bf4f2fd7a2e951b977eee0d5f86c0163a164087e7498e"></a>

## Next pages — Property reference / 819c50e1e319 / 14

- [details](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-3ff6c2003052e5c0e764bfcdc89c57174450b2e337057daa66e9edc9142fd041)
- [xcsh_bot_peer_traffic_overview](../data-sources/bot_peer_traffic_overview.md#canonical-5b05326e4d591c752fdaa94722e4badaea615d250171e751c47a2266294991d7)

<a id="canonical-3ff6c2003052e5c0e764bfcdc89c57174450b2e337057daa66e9edc9142fd041"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0432f394d3a766795107013ac3e982a1e265559c59e3f13d8fd558b273c52513"></a>

## details — details / aa1dc2901926 / 2

Breadcrumbs:

- [xcsh_bot_peer_traffic_overview](../data-sources/bot_peer_traffic_overview.md#canonical-5b05326e4d591c752fdaa94722e4badaea615d250171e751c47a2266294991d7)
- [Property reference](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-53367a814368233d59c0d60053b9993843031752de8473de41d0df93b435773a)
- details

<a id="canonical-d5d8b22929f87befb3da95fb63f149e466809b41cef2cd076d93b22d64e98d8b"></a>

Type: `"list"`. Computed.

Peer Group Traffic Overview Details. Configuration parameter for details

<a id="canonical-c165c13e479d530b695f0d044ae7d994f745805ced40298252500fbe230c7460"></a>

## Direct properties — details / aa1dc2901926 / 3

<a id="canonical-d89e028e82c7885b55a4fe97deb1a1223ebe9fcbada0b79cfea151040c0c952d"></a>

<a id="canonical-bb5807af2a43cb53ed7c87f04d3b21ff83d25253633a3cfd539810df5d54b16d"></a>

## name property — details / aa1dc2901926 / 4

Type: `"string"`. Computed.

Name. The Name of Item.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-94d7182761e3644e6611ecffe085996a40b62ab317c1d8c96f01bd870585df8e"></a>

<a id="canonical-ca642fae35442177806f2af32e52376baed26839d470cbbf5acee7883fc45517"></a>

## peer_count property — details / aa1dc2901926 / 5

Type: `"string"`. Computed.

Peer Count. The count of Peer of the Item.

<a id="canonical-2142e1c0b71a59f8aaffea4a347e0afba3d74c2e36ca4e9156a0bded77ce17d1"></a>

<a id="canonical-92f35e4e7f99517aa9ab03be52dd1ad9a0170c0fcea85f8f933b97b0a6043603"></a>

## peer_percentage property — details / aa1dc2901926 / 6

Type: `"number"`. Computed.

Peer Percentage. The Peer Percentage of the Item.

<a id="canonical-9e5be1f493b8c85f046135d93d6289a1adffde83b9c565701f7f12257d4e6dac"></a>

<a id="canonical-ab38b545701d5b3fe9fed1a9923d6b896544a05f0df36b4defb7b732ce99d4b9"></a>

## self_count property — details / aa1dc2901926 / 7

Type: `"string"`. Computed.

Self Count. The count of Self of the item.

<a id="canonical-0a9e26d509c429a0e80e9148738a8a415f835b21e2d3fd631e3bf85340997f07"></a>

<a id="canonical-c7d1fd029bdd134cfd47cbbfbcb4fddd0303a0a8b5b5db271d742f89da65c55f"></a>

## self_percentage property — details / aa1dc2901926 / 8

Type: `"number"`. Computed.

Self Percentage. The Self Percentage of the Item.

<a id="canonical-f100362b625c3f6f7192f5fa5908ba21fd4a981083cd5b16c5e8fa42fff64af2"></a>

## Next pages — details / aa1dc2901926 / 9

- [Property reference](data-sources--bot_peer_traffic_overview--reference--group-001.md#canonical-53367a814368233d59c0d60053b9993843031752de8473de41d0df93b435773a)
- [xcsh_bot_peer_traffic_overview](../data-sources/bot_peer_traffic_overview.md#canonical-5b05326e4d591c752fdaa94722e4badaea615d250171e751c47a2266294991d7)
