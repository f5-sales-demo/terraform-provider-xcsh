---
page_title: "xcsh_bot_peer_threat_types reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_threat_types reference."
---

# xcsh_bot_peer_threat_types reference

<a id="canonical-fe41abbeffdeb22d2d0f00d523f6ba4a5c0487f7ff396036ada01cc5a9aac87e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edd89eb1552efa70307a405631d6f2518242765423bfd3f255fb849e821032ff"></a>

## Property reference — Property reference / 02eb73af5b74 / 2

Breadcrumbs:

- [xcsh_bot_peer_threat_types](../data-sources/bot_peer_threat_types.md#canonical-87108d47396ce895c0d9b659f1d45c2967514dee0cd6d925cfc6ed3f040bf641)
- Property reference

<a id="canonical-5ed77476200ba59fa0b19af50a167f4c5b4f9d2d3ce46bd012df9f356ab3c7b9"></a>

## Direct properties — Property reference / 02eb73af5b74 / 3

- [details](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-a5556dad9d01aba6c200161da921238fa71ced8bcbaacb369b7e710eb29feb54): complete subsection reference.

<a id="canonical-4ba3e6a74407d04c96ebaf3fb144bbae720e3c92687b4b29d3a71179f5a707b2"></a>

<a id="canonical-b91223f14a62864d55f6e529290862adeb158540008eb4d2d224b880621e8f12"></a>

## end_time property — Property reference / 02eb73af5b74 / 4

Type: `"string"`. Optional.

End Time. End time of the query period.

<a id="canonical-c4b7bd554a82fed313ee3882d2c0a380a291c8e6d7c4d09eadc4ab31fffdd268"></a>

<a id="canonical-0c8d3a4bf7831ed320158697f5c55f6c5d32a3a27bef8bd66b1ad0a7be0534d8"></a>

## limit property — Property reference / 02eb73af5b74 / 5

Type: `"number"`. Optional.

Limits the number of transactions returned in the response Optional: If not specified (with default
value 0), all transactions that match the query will be returned in the response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 500),
}
```

<a id="canonical-76301aa3d1c19a2a3a150dfad8cbc2e9b5e8919c1094a21736cf19c5081e32cc"></a>

<a id="canonical-866794e160aec935c97fde65a939ee3822c46bb45caff5d953c991332fea945c"></a>

## namespace property — Property reference / 02eb73af5b74 / 6

Type: `"string"`. Required.

Namespace. namespace is used to scope the query. Only virtual\_host in given namespace will be
considered.

<a id="canonical-4447bbd3d2b25f1c9bd2c6c7ba67ef7e2a773c97d4751d3209238a8628929ed7"></a>

<a id="canonical-f405f2932007a6118820511edd6b2f82b87daddd4fb5fece1e6ff6ae416a6024"></a>

## rank_by property — Property reference / 02eb73af5b74 / 7

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

<a id="canonical-11f0794097af59cee59157f3dfcd2062556295352e32ce00906bbb9a633a0db0"></a>

<a id="canonical-c1ec1f415a1b93562f2150002a70ed752bef140340c5e0b0649d0bbcc7bb7361"></a>

## start_time property — Property reference / 02eb73af5b74 / 8

Type: `"string"`. Optional.

Start Time. Start time of the query period.

<a id="canonical-42772cfd0515ae731f941054f9ee471181e601254d5adeb1144f98c81a90599a"></a>

## All schema paths — Property reference / 02eb73af5b74 / 9

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `details` | [details](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-74fa768d4bff725f2ec2e30a0a4ba69e5d24d82bfaa9c3d48b5c7d61307c06b2) |
| `details.name` | [details.name](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-758d857cd12ec50bd743e49506b148199e22d94752096c8b7287b7553614fc41) |
| `details.peer_count` | [details.peer_count](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-54b47cdce91648ebc93bf59d74ef39c9e6b9ff04f858de9015212ef1ee6a5827) |
| `details.peer_percentage` | [details.peer_percentage](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-e55e5fd2028134d4e9e22926ec32e621afa7cd3b35e9dc11d65a509b16ac2427) |
| `details.self_count` | [details.self_count](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-7bfd35114b04a8cd8b4c9e775346c47f5ec011a4c4976d6cfe3ef361be0ef00c) |
| `details.self_percentage` | [details.self_percentage](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-ad607d89e624d28f46c51f1c3168ceca29e0dce2511ddfe3d6645c446ff00f07) |
| `end_time` | [end_time](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-4ba3e6a74407d04c96ebaf3fb144bbae720e3c92687b4b29d3a71179f5a707b2) |
| `limit` | [limit](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-c4b7bd554a82fed313ee3882d2c0a380a291c8e6d7c4d09eadc4ab31fffdd268) |
| `namespace` | [namespace](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-76301aa3d1c19a2a3a150dfad8cbc2e9b5e8919c1094a21736cf19c5081e32cc) |
| `rank_by` | [rank_by](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-4447bbd3d2b25f1c9bd2c6c7ba67ef7e2a773c97d4751d3209238a8628929ed7) |
| `start_time` | [start_time](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-11f0794097af59cee59157f3dfcd2062556295352e32ce00906bbb9a633a0db0) |

<a id="canonical-5a551ed5e067a7dfdbd03660e409304925a61c45b82e5d097edc3f13dca96584"></a>

## Next pages — Property reference / 02eb73af5b74 / 10

- [details](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-a5556dad9d01aba6c200161da921238fa71ced8bcbaacb369b7e710eb29feb54)
- [xcsh_bot_peer_threat_types](../data-sources/bot_peer_threat_types.md#canonical-87108d47396ce895c0d9b659f1d45c2967514dee0cd6d925cfc6ed3f040bf641)

<a id="canonical-a5556dad9d01aba6c200161da921238fa71ced8bcbaacb369b7e710eb29feb54"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2edde7a35bc0259dc3b73dc43e2bfb8aa3e3efa2928b56b7adaa8f4f7a4d3e60"></a>

## details — details / 9ff64c422482 / 2

Breadcrumbs:

- [xcsh_bot_peer_threat_types](../data-sources/bot_peer_threat_types.md#canonical-87108d47396ce895c0d9b659f1d45c2967514dee0cd6d925cfc6ed3f040bf641)
- [Property reference](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-fe41abbeffdeb22d2d0f00d523f6ba4a5c0487f7ff396036ada01cc5a9aac87e)
- details

<a id="canonical-74fa768d4bff725f2ec2e30a0a4ba69e5d24d82bfaa9c3d48b5c7d61307c06b2"></a>

Type: `"list"`. Computed.

Peer Group Top Good Bots Details. Configuration parameter for details

<a id="canonical-a163dd2caf9081b5a80629004c6f04e9a258fddad43b1bba8a48e0c71d770d0d"></a>

## Direct properties — details / 9ff64c422482 / 3

<a id="canonical-758d857cd12ec50bd743e49506b148199e22d94752096c8b7287b7553614fc41"></a>

<a id="canonical-2a9e555730a583243b8b3a685c492f4b40a5e474054f4f8f5657caf613034d7e"></a>

## name property — details / 9ff64c422482 / 4

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

<a id="canonical-54b47cdce91648ebc93bf59d74ef39c9e6b9ff04f858de9015212ef1ee6a5827"></a>

<a id="canonical-481f87b4329496d702cf0cbc4868f6a5ae9156ee3633d08dad584365b2e264e8"></a>

## peer_count property — details / 9ff64c422482 / 5

Type: `"string"`. Computed.

Peer Count. The count of Peer of the Item.

<a id="canonical-e55e5fd2028134d4e9e22926ec32e621afa7cd3b35e9dc11d65a509b16ac2427"></a>

<a id="canonical-3c0d039aadb90906c80af0f65aa434bc3e5c7ac248a71e3bb37089707fe20761"></a>

## peer_percentage property — details / 9ff64c422482 / 6

Type: `"number"`. Computed.

Peer Percentage. The Peer Percentage of the Item.

<a id="canonical-7bfd35114b04a8cd8b4c9e775346c47f5ec011a4c4976d6cfe3ef361be0ef00c"></a>

<a id="canonical-34830910f47ae3d7dcb56f5be99f26423c9e824ea96fb63095f780e32185b6ac"></a>

## self_count property — details / 9ff64c422482 / 7

Type: `"string"`. Computed.

Self Count. The count of Self of the item.

<a id="canonical-ad607d89e624d28f46c51f1c3168ceca29e0dce2511ddfe3d6645c446ff00f07"></a>

<a id="canonical-2073dcbd45858f09295ec730bd9e7d602b90a8115fee4a5171e0056eeb94118c"></a>

## self_percentage property — details / 9ff64c422482 / 8

Type: `"number"`. Computed.

Self Percentage. The Self Percentage of the Item.

<a id="canonical-886487806189e94b370d21dc2104bcd7307cffbfbe6be4da16dce544443c21e2"></a>

## Next pages — details / 9ff64c422482 / 9

- [Property reference](data-sources--bot_peer_threat_types--reference--group-001.md#canonical-fe41abbeffdeb22d2d0f00d523f6ba4a5c0487f7ff396036ada01cc5a9aac87e)
- [xcsh_bot_peer_threat_types](../data-sources/bot_peer_threat_types.md#canonical-87108d47396ce895c0d9b659f1d45c2967514dee0cd6d925cfc6ed3f040bf641)
