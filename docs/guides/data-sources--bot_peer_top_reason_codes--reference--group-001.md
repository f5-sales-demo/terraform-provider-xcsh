---
page_title: "xcsh_bot_peer_top_reason_codes reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_top_reason_codes reference."
---

# xcsh_bot_peer_top_reason_codes reference

<a id="canonical-1458b8c00d28b21f8f369757866bee112100ec9dea2320a099dc3c9ccabea02c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ece44c3ce95cdd12840cdec0329c08de2972fc54a74ffeb34a56e85b058e8fe8"></a>

## Property reference — Property reference / c5990c14a1c1 / 2

Breadcrumbs:

- [xcsh_bot_peer_top_reason_codes](../data-sources/bot_peer_top_reason_codes.md#canonical-07a07206937eae30b859c6d2292d9376c48db6e56deabe8cee1ccc1272a3593f)
- Property reference

<a id="canonical-9b188251bb646ac3be8a57cbcd676807f463be2a137fa6786e6b1a23e40ef298"></a>

## Direct properties — Property reference / c5990c14a1c1 / 3

- [details](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-23a4e39c8bda29c9ea1a62f16cfe16940e97aeadedbc0c5788e12a61e4e7c742): complete subsection reference.

<a id="canonical-cba898eec44bf7e29d547b01463e60b6cd943a898b0ce408e38789c0ce0cbe11"></a>

<a id="canonical-9c8864ba1790214f4f41aa12e96fd8dcf5ca9270f9f2371fb39be9e25c7cca93"></a>

## end_time property — Property reference / c5990c14a1c1 / 4

Type: `"string"`. Optional.

End Time. End time of the query period.

<a id="canonical-02beefb73409e9be3d09c1be47fb13cfd52c4857c0f841b5d535ec15ca404865"></a>

<a id="canonical-ad01fc5301dc1fc741a8dbe2c9d72cb20a8ef21429e1c324655126d2395b2000"></a>

## limit property — Property reference / c5990c14a1c1 / 5

Type: `"number"`. Optional.

Limits the number of transactions returned in the response Optional: If not specified (with default
value 0), all transactions that match the query will be returned in the response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 500),
}
```

<a id="canonical-90c47505a0d9fd7a275b734809d6601e3bf871f1230116fc3f51cf874a6c1417"></a>

<a id="canonical-5acf9f072eaa33bb2e73f8f4c1ac36433d2ad3042a14b1f225088ec37ffd9d4a"></a>

## namespace property — Property reference / c5990c14a1c1 / 6

Type: `"string"`. Required.

Namespace. namespace is used to scope the query. Only virtual\_host in given namespace will be
considered.

<a id="canonical-d3811fe02156394ec0588ea8fd4d803d7d58dfab46c1982f7154f3a5f804d1ee"></a>

<a id="canonical-9d9073e9e0fcb1a835cf24c636e8fd588d1c60dfeb8126fd8d5020b4c0ecb3e6"></a>

## rank_by property — Property reference / c5990c14a1c1 / 7

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

<a id="canonical-95f3943c8f2e9e24b6530b14e60549638b10f454b242a21b998bef67b5e33e85"></a>

<a id="canonical-922e67a6f6ed4bc25f6a244a398d0b166aad89c7ffa3bb4d4a8f5560d476defb"></a>

## start_time property — Property reference / c5990c14a1c1 / 8

Type: `"string"`. Optional.

Start Time. Start time of the query period.

<a id="canonical-65c760ba40eda3e6e1861bafdfaaac6c59049a5ad0e04ce70cfe05fc5692fec6"></a>

## All schema paths — Property reference / c5990c14a1c1 / 9

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `details` | [details](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-ffb37437a96fdcd1f045b57754b46d1f8d1a15e321d7fade93aeeee9c2da9080) |
| `details.name` | [details.name](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-5f98995be5e55ca192ad4b3b63dd406e95276ea1ec97f3aa1151a0da0df33f57) |
| `details.peer_count` | [details.peer_count](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-a387ca0d11e43e66f868961ac9b7bade916ca700875c2173904fa253a9676eee) |
| `details.peer_percentage` | [details.peer_percentage](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-d0ccfa2e2aac0a459fc58b1618c1a9adff9824b4536ab9852c1d6f3f295208e6) |
| `details.self_count` | [details.self_count](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-8d8a16f0ced05f537a9efb16d314d5784c795328014d4344507d1509c6fc09b9) |
| `details.self_percentage` | [details.self_percentage](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-1e2081e7d10d596e7a6b89f6145655387fee699e1e2f569134a5d10065e94528) |
| `end_time` | [end_time](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-cba898eec44bf7e29d547b01463e60b6cd943a898b0ce408e38789c0ce0cbe11) |
| `limit` | [limit](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-02beefb73409e9be3d09c1be47fb13cfd52c4857c0f841b5d535ec15ca404865) |
| `namespace` | [namespace](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-90c47505a0d9fd7a275b734809d6601e3bf871f1230116fc3f51cf874a6c1417) |
| `rank_by` | [rank_by](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-d3811fe02156394ec0588ea8fd4d803d7d58dfab46c1982f7154f3a5f804d1ee) |
| `start_time` | [start_time](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-95f3943c8f2e9e24b6530b14e60549638b10f454b242a21b998bef67b5e33e85) |

<a id="canonical-358162c5d2dc1b295fd9a2ca03cf98bca3140b259bb4429955c565bc0237556e"></a>

## Next pages — Property reference / c5990c14a1c1 / 10

- [details](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-23a4e39c8bda29c9ea1a62f16cfe16940e97aeadedbc0c5788e12a61e4e7c742)
- [xcsh_bot_peer_top_reason_codes](../data-sources/bot_peer_top_reason_codes.md#canonical-07a07206937eae30b859c6d2292d9376c48db6e56deabe8cee1ccc1272a3593f)

<a id="canonical-23a4e39c8bda29c9ea1a62f16cfe16940e97aeadedbc0c5788e12a61e4e7c742"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ff4658987a0747979b4bc78bfa5df10a8defa56551ed90e4462f99bb4140dfe"></a>

## details — details / eb88a982bb87 / 2

Breadcrumbs:

- [xcsh_bot_peer_top_reason_codes](../data-sources/bot_peer_top_reason_codes.md#canonical-07a07206937eae30b859c6d2292d9376c48db6e56deabe8cee1ccc1272a3593f)
- [Property reference](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-1458b8c00d28b21f8f369757866bee112100ec9dea2320a099dc3c9ccabea02c)
- details

<a id="canonical-ffb37437a96fdcd1f045b57754b46d1f8d1a15e321d7fade93aeeee9c2da9080"></a>

Type: `"list"`. Computed.

Peer Group Top Good Bots Details. Configuration parameter for details

<a id="canonical-0bdb643dca667d9712a40f2819c4dcf5d75bfcaf971f01c24cc8f1a9d2226796"></a>

## Direct properties — details / eb88a982bb87 / 3

<a id="canonical-5f98995be5e55ca192ad4b3b63dd406e95276ea1ec97f3aa1151a0da0df33f57"></a>

<a id="canonical-917bf90d125f1a6200ac826b0f420d907f52bcf76144e1c580bef04cbe402657"></a>

## name property — details / eb88a982bb87 / 4

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

<a id="canonical-a387ca0d11e43e66f868961ac9b7bade916ca700875c2173904fa253a9676eee"></a>

<a id="canonical-c204f9ef9756b43f0e980725ef2dc4762a05f20b086c4f97f109229c5b75f650"></a>

## peer_count property — details / eb88a982bb87 / 5

Type: `"string"`. Computed.

Peer Count. The count of Peer of the Item.

<a id="canonical-d0ccfa2e2aac0a459fc58b1618c1a9adff9824b4536ab9852c1d6f3f295208e6"></a>

<a id="canonical-5dc22f4b2c209c5ce6b7641c357349a02733669218453ed77a42317ac3aa05f7"></a>

## peer_percentage property — details / eb88a982bb87 / 6

Type: `"number"`. Computed.

Peer Percentage. The Peer Percentage of the Item.

<a id="canonical-8d8a16f0ced05f537a9efb16d314d5784c795328014d4344507d1509c6fc09b9"></a>

<a id="canonical-dfffeb817055bb8abde62fe664d722bfc9e9e9892176ee8f7c615d7a0b5d8ef1"></a>

## self_count property — details / eb88a982bb87 / 7

Type: `"string"`. Computed.

Self Count. The count of Self of the item.

<a id="canonical-1e2081e7d10d596e7a6b89f6145655387fee699e1e2f569134a5d10065e94528"></a>

<a id="canonical-d733a37efe730e78cf488f2a00941cc7a14c6438818b1bccca803fca8f330140"></a>

## self_percentage property — details / eb88a982bb87 / 8

Type: `"number"`. Computed.

Self Percentage. The Self Percentage of the Item.

<a id="canonical-475efd5ea6020d1bb9702a767029d099386ab87a55deaa15d04259c1ca897a06"></a>

## Next pages — details / eb88a982bb87 / 9

- [Property reference](data-sources--bot_peer_top_reason_codes--reference--group-001.md#canonical-1458b8c00d28b21f8f369757866bee112100ec9dea2320a099dc3c9ccabea02c)
- [xcsh_bot_peer_top_reason_codes](../data-sources/bot_peer_top_reason_codes.md#canonical-07a07206937eae30b859c6d2292d9376c48db6e56deabe8cee1ccc1272a3593f)
