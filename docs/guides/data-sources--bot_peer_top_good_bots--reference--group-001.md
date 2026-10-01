---
page_title: "xcsh_bot_peer_top_good_bots reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_top_good_bots reference."
---

# xcsh_bot_peer_top_good_bots reference

<a id="canonical-6ac86678f40f40275bf9c57605bc9b4e11cf7b3003d83b65b2d5f4fdfbda158e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f727275d57f7aa63a58cd6cbac4dd5e80d1f8060d0fe62e567201822608b036c"></a>

## Property reference — Property reference / 5e846e06b634 / 2

Breadcrumbs:

- [xcsh_bot_peer_top_good_bots](../data-sources/bot_peer_top_good_bots.md#canonical-cc077f9c205b2a2bad9388a3f1f0bfb9445ec52b5c9b1a27c681c552089ee8f2)
- Property reference

<a id="canonical-00f1cee96e02e5314070ca2eb0988881286f1f9f962cac503488909695a53592"></a>

## Direct properties — Property reference / 5e846e06b634 / 3

- [details](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-1896fcf1d9fc69f6db0ad60a80699d6d2590856ca4902aa0c8f24c92f4d5c419): complete subsection reference.

<a id="canonical-97649e57726816568c9df3ef0e8d3bfbc201df0f6c3fa31563ef661e2dbd456c"></a>

<a id="canonical-ca59e951adf439b0610edda7e62dedda43ea0225aef58219961f9abc763a4595"></a>

## end_time property — Property reference / 5e846e06b634 / 4

Type: `"string"`. Optional.

End Time. End time of the query period.

<a id="canonical-ebabe1f950cd796bfd76954ca74d13483dae489da2d88f951529b2a291c709c1"></a>

<a id="canonical-eea72fa97a3d5fbc8c208f21d0f430ffd96004d4d77faa556f33637b4fa0692c"></a>

## limit property — Property reference / 5e846e06b634 / 5

Type: `"number"`. Optional.

Limits the number of transactions returned in the response Optional: If not specified (with default
value 0), all transactions that match the query will be returned in the response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 500),
}
```

<a id="canonical-15356e977c060cf015e066f94f903ec99efcff1168c1acfedfbcb90b54faa1be"></a>

<a id="canonical-05f474babaa12cf832e6a67315400ede35622a67e7d55314dc8d558c91c2777d"></a>

## namespace property — Property reference / 5e846e06b634 / 6

Type: `"string"`. Required.

Namespace. namespace is used to scope the query. Only virtual\_host in given namespace will be
considered.

<a id="canonical-188e6722d26cf83db1ce18bc3cf287077e17c1139acd3a4f18d57e7048e49456"></a>

<a id="canonical-8922d7b708c13a4c44c4184860195603c6ea275d3716135e50b032465052b36f"></a>

## rank_by property — Property reference / 5e846e06b634 / 7

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

<a id="canonical-a583392025ea79bea88d120f8314405be700b9ff472bfef95874fb79527a7ec1"></a>

<a id="canonical-954f8f8a1a12be89ed04f1f9edcc2d9d625f75a46d232f6681e96013f13299ec"></a>

## start_time property — Property reference / 5e846e06b634 / 8

Type: `"string"`. Optional.

Start Time. Start time of the query period.

<a id="canonical-f97a4c7b5a313b1c7d23ab465ca195ed34d9c65f8922702bc0dfe11ea560d7f3"></a>

## All schema paths — Property reference / 5e846e06b634 / 9

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `details` | [details](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-67bf51d735ef1eb00c794cfb3cebead5ea090f948b6c140c105bf78e02a9a93d) |
| `details.name` | [details.name](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-4423acd7960b7391addd2c0460e147b67e02c978fd1104668f4d242ee5ede4c1) |
| `details.peer_count` | [details.peer_count](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-552063750503c3e2f15e8b7a0e96174951ec7ccc61d531657a24cd75faa03936) |
| `details.peer_percentage` | [details.peer_percentage](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-b6389abf39c3a25dd6278a80bfb9a55eebdce1e14e7c0469b77aade2cd6b1372) |
| `details.self_count` | [details.self_count](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-3c3c7cc8002b330bdb2b2e4d0609a135b18b9cc93fe6a17fea5a312016a3b95b) |
| `details.self_percentage` | [details.self_percentage](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-fe8e3cd89641ccbe34a3276051d05bd9f9562180ea280617b556caad22f0707b) |
| `end_time` | [end_time](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-97649e57726816568c9df3ef0e8d3bfbc201df0f6c3fa31563ef661e2dbd456c) |
| `limit` | [limit](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-ebabe1f950cd796bfd76954ca74d13483dae489da2d88f951529b2a291c709c1) |
| `namespace` | [namespace](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-15356e977c060cf015e066f94f903ec99efcff1168c1acfedfbcb90b54faa1be) |
| `rank_by` | [rank_by](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-188e6722d26cf83db1ce18bc3cf287077e17c1139acd3a4f18d57e7048e49456) |
| `start_time` | [start_time](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-a583392025ea79bea88d120f8314405be700b9ff472bfef95874fb79527a7ec1) |

<a id="canonical-1e364bc587b50a7c15471ebee1981cc2cff7fa6e76262e6691a182ecaa9ebf19"></a>

## Next pages — Property reference / 5e846e06b634 / 10

- [details](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-1896fcf1d9fc69f6db0ad60a80699d6d2590856ca4902aa0c8f24c92f4d5c419)
- [xcsh_bot_peer_top_good_bots](../data-sources/bot_peer_top_good_bots.md#canonical-cc077f9c205b2a2bad9388a3f1f0bfb9445ec52b5c9b1a27c681c552089ee8f2)

<a id="canonical-1896fcf1d9fc69f6db0ad60a80699d6d2590856ca4902aa0c8f24c92f4d5c419"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b22d5b48a37a3da83ed012b763c7c81e33db0a32f739579094b5820d15a04067"></a>

## details — details / 4f16162b5429 / 2

Breadcrumbs:

- [xcsh_bot_peer_top_good_bots](../data-sources/bot_peer_top_good_bots.md#canonical-cc077f9c205b2a2bad9388a3f1f0bfb9445ec52b5c9b1a27c681c552089ee8f2)
- [Property reference](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-6ac86678f40f40275bf9c57605bc9b4e11cf7b3003d83b65b2d5f4fdfbda158e)
- details

<a id="canonical-67bf51d735ef1eb00c794cfb3cebead5ea090f948b6c140c105bf78e02a9a93d"></a>

Type: `"list"`. Computed.

Peer Group Top Good Bots Details. Configuration parameter for details

<a id="canonical-a57f242ff1d299ad17a3842fb6b0edd306dd481db7801c02d032d62c7b867b75"></a>

## Direct properties — details / 4f16162b5429 / 3

<a id="canonical-4423acd7960b7391addd2c0460e147b67e02c978fd1104668f4d242ee5ede4c1"></a>

<a id="canonical-db16006f8a1b3a5ac47d07208b7bb3c28ffa60c65f763b9cf668dd2f5b69385c"></a>

## name property — details / 4f16162b5429 / 4

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

<a id="canonical-552063750503c3e2f15e8b7a0e96174951ec7ccc61d531657a24cd75faa03936"></a>

<a id="canonical-6944a5ec3bb84924184c5252a682a7bcf0ff189f137bbc316b9951a02ea4a9aa"></a>

## peer_count property — details / 4f16162b5429 / 5

Type: `"string"`. Computed.

Peer Count. The count of Peer of the Item.

<a id="canonical-b6389abf39c3a25dd6278a80bfb9a55eebdce1e14e7c0469b77aade2cd6b1372"></a>

<a id="canonical-b55acc7a2a25be9b6b44c1027941e20a8d731ab2123b60a52bc13181b9238621"></a>

## peer_percentage property — details / 4f16162b5429 / 6

Type: `"number"`. Computed.

Peer Percentage. The Peer Percentage of the Item.

<a id="canonical-3c3c7cc8002b330bdb2b2e4d0609a135b18b9cc93fe6a17fea5a312016a3b95b"></a>

<a id="canonical-2cc0029bd9837a58dfac6f1f5765ade687626fd2f5bd32690607fdaa0fcbbfb5"></a>

## self_count property — details / 4f16162b5429 / 7

Type: `"string"`. Computed.

Self Count. The count of Self of the item.

<a id="canonical-fe8e3cd89641ccbe34a3276051d05bd9f9562180ea280617b556caad22f0707b"></a>

<a id="canonical-0eb59c655d51771435202eb6a44ed07968c4e04fda256e4a94a31232830c0435"></a>

## self_percentage property — details / 4f16162b5429 / 8

Type: `"number"`. Computed.

Self Percentage. The Self Percentage of the Item.

<a id="canonical-ba07ac912275d81f2e5d1a3fcb3a11b8ef74e9ac7ed3d3cf7c667b4965564067"></a>

## Next pages — details / 4f16162b5429 / 9

- [Property reference](data-sources--bot_peer_top_good_bots--reference--group-001.md#canonical-6ac86678f40f40275bf9c57605bc9b4e11cf7b3003d83b65b2d5f4fdfbda158e)
- [xcsh_bot_peer_top_good_bots](../data-sources/bot_peer_top_good_bots.md#canonical-cc077f9c205b2a2bad9388a3f1f0bfb9445ec52b5c9b1a27c681c552089ee8f2)
