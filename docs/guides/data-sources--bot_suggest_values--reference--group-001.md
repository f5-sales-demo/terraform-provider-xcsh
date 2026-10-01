---
page_title: "xcsh_bot_suggest_values reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_suggest_values reference."
---

# xcsh_bot_suggest_values reference

<a id="canonical-1fd46f22c3a31585d917103e1b4692467b3651cd75d9fe25d8cf70497503b572"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-205a7b6cabbe09e1cd487cb3635587cfb9ebda616696510edfebbf12751a5e62"></a>

## Property reference — Property reference / d32a6ab13e68 / 2

Breadcrumbs:

- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md#canonical-eb5fdcd8e576fd0893458000f2f7792dfce1074e73bded33582900b2e388a473)
- Property reference

<a id="canonical-b7d60d0ebeff4f0b3a2f0a1d66d77c3692c2be514e4fb14f8f739c4166c40ee6"></a>

## Direct properties — Property reference / d32a6ab13e68 / 3

<a id="canonical-ef55b3649a45e1d9772c3fa2cdfa01f2cc259d01385bd5676b47eb28ec65ccd4"></a>

<a id="canonical-37d30ac22c256bb665b161032a21f67e0baa65cabc045e160d285dca1edd01e8"></a>

## field_path property — Property reference / d32a6ab13e68 / 4

Type: `"string"`. Optional.

JSON path of the field for which the suggested values are being requested.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

- [items](data-sources--bot_suggest_values--reference--group-001.md#canonical-5525495cda2ea7ed853ce566ef5caf44c1a3ba52a9e64a3d68e0a6e67f002a69): complete subsection reference.

<a id="canonical-4c1d4cd8b491fc4d280d75a7ca2bb4f3300abafe1fcea83e94c6faae609de449"></a>

<a id="canonical-0002c792ceb306ffe29f1f875801ba161c0eae96213f5d07987c9ec2693dc2be"></a>

## match_value property — Property reference / d32a6ab13e68 / 5

Type: `"string"`. Optional.

Substring that must be present in either the value or description of each SuggestedItem in the
response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-c9c99046ec4e3ba1daee92152c4c0fac1051246c97a4eb26c2a6b44f7ec44cdd"></a>

<a id="canonical-df9b1f73beb72ebabafa6495bb47a52ae93070b86ba642cc881359014b4c776d"></a>

## namespace property — Property reference / d32a6ab13e68 / 6

Type: `"string"`. Required.

Namespace Namespace in which the suggestions are scoped.

- [request_body](data-sources--bot_suggest_values--reference--group-001.md#canonical-b20993a3bce85dda40d81bbcfa72db5f0eee3674bf37e1dd9ad913470ef04f1b): complete subsection reference.

<a id="canonical-8e8c1e7822b792c14dcf3785aac915b4a9e50f7d1f232ac2ad088b9816c77903"></a>

## All schema paths — Property reference / d32a6ab13e68 / 7

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `field_path` | [field_path](data-sources--bot_suggest_values--reference--group-001.md#canonical-ef55b3649a45e1d9772c3fa2cdfa01f2cc259d01385bd5676b47eb28ec65ccd4) |
| `items` | [items](data-sources--bot_suggest_values--reference--group-001.md#canonical-f471b1d93eb5a784647a2c4392000da5886eaf970c93d11f3893cf99c1b336a1) |
| `items.description_spec` | [items.description_spec](data-sources--bot_suggest_values--reference--group-001.md#canonical-a2dd8519ba19aa521048919a2a637761cd6de4c84308a55d8ec069bc169d9990) |
| `items.ref_value` | [items.ref_value](data-sources--bot_suggest_values--reference--group-001.md#canonical-6cf951e374e90f1c3a33ba2f5b18652d67d433ff42659dc11f40aff1b448a091) |
| `items.ref_value.name` | [items.ref_value.name](data-sources--bot_suggest_values--reference--group-001.md#canonical-e7d380398126d837f5046eb411c683f53e01999d44830508f3a9539bce79e4cb) |
| `items.ref_value.namespace` | [items.ref_value.namespace](data-sources--bot_suggest_values--reference--group-001.md#canonical-cc932cd01c8edcc5aeafe5e321cd6c6ff1833431a0547fbae93744304365f9a8) |
| `items.ref_value.tenant` | [items.ref_value.tenant](data-sources--bot_suggest_values--reference--group-001.md#canonical-1a4e299ede393d0b79474275b8f50fcbe9c63f2b79c3635c9dad8fb145f012ab) |
| `items.str_value` | [items.str_value](data-sources--bot_suggest_values--reference--group-001.md#canonical-c732f7b503155812ae7726489ecaec64a6637446c7c493425e75a973447be490) |
| `items.title` | [items.title](data-sources--bot_suggest_values--reference--group-001.md#canonical-995477810e4cea6ee3d9f434b676ee601428b02493010c35ccf558f5fd2615d9) |
| `items.value` | [items.value](data-sources--bot_suggest_values--reference--group-001.md#canonical-a0c4bdd5e8a0ea8d5eff98dae77af0625f16346a478ca01041c12eca764dc2f6) |
| `match_value` | [match_value](data-sources--bot_suggest_values--reference--group-001.md#canonical-4c1d4cd8b491fc4d280d75a7ca2bb4f3300abafe1fcea83e94c6faae609de449) |
| `namespace` | [namespace](data-sources--bot_suggest_values--reference--group-001.md#canonical-c9c99046ec4e3ba1daee92152c4c0fac1051246c97a4eb26c2a6b44f7ec44cdd) |
| `request_body` | [request_body](data-sources--bot_suggest_values--reference--group-001.md#canonical-8f3d259dc114c9afc7c4ce9bbcc80d802ac7ed1f15bfd65903c6eaad38709825) |
| `request_body.type_url` | [request_body.type_url](data-sources--bot_suggest_values--reference--group-001.md#canonical-fadc7c15cab22ec80cc121ec90cf2c9677a3895cde021458d9641e60104cc3bd) |
| `request_body.value` | [request_body.value](data-sources--bot_suggest_values--reference--group-001.md#canonical-6fdc993dcba6c9831c90da8de16924851e3608022ed5aaffa7f7e99f7d034ed4) |

<a id="canonical-dc34711dacacc35f9c7feac0d723aeb739c5b5911a0076fb4b5634291bf7aa03"></a>

## Next pages — Property reference / d32a6ab13e68 / 8

- [items](data-sources--bot_suggest_values--reference--group-001.md#canonical-5525495cda2ea7ed853ce566ef5caf44c1a3ba52a9e64a3d68e0a6e67f002a69)
- [request_body](data-sources--bot_suggest_values--reference--group-001.md#canonical-b20993a3bce85dda40d81bbcfa72db5f0eee3674bf37e1dd9ad913470ef04f1b)
- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md#canonical-eb5fdcd8e576fd0893458000f2f7792dfce1074e73bded33582900b2e388a473)

<a id="canonical-5525495cda2ea7ed853ce566ef5caf44c1a3ba52a9e64a3d68e0a6e67f002a69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e509e00a514659a8dcc627059c64a9e3eb2885b3a81c623279a6a959c0ebe48"></a>

## items — items / 776b65fcf749 / 2

Breadcrumbs:

- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md#canonical-eb5fdcd8e576fd0893458000f2f7792dfce1074e73bded33582900b2e388a473)
- [Property reference](data-sources--bot_suggest_values--reference--group-001.md#canonical-1fd46f22c3a31585d917103e1b4692467b3651cd75d9fe25d8cf70497503b572)
- items

<a id="canonical-f471b1d93eb5a784647a2c4392000da5886eaf970c93d11f3893cf99c1b336a1"></a>

Type: `"list"`. Computed.

Suggested Items. List of suggested items.

<a id="canonical-8ca0331330ee2c5c4c5e292cfb00f3a3bc09f20c7201bc18b7d5dfc72b887b6f"></a>

## Direct properties — items / 776b65fcf749 / 3

<a id="canonical-a2dd8519ba19aa521048919a2a637761cd6de4c84308a55d8ec069bc169d9990"></a>

<a id="canonical-f97bd034a21916f22be621978b7f9315053805b93f282d0e72457ad0b3c54acf"></a>

## description_spec property — items / 776b65fcf749 / 4

Type: `"string"`. Computed.

Optional description for the suggested value.

- [ref_value](data-sources--bot_suggest_values--reference--group-001.md#canonical-eec009c32f5d57d623ea8dc22541fa11fe4cffaf361a32de8943aef801cd7879): complete subsection reference.

<a id="canonical-c732f7b503155812ae7726489ecaec64a6637446c7c493425e75a973447be490"></a>

<a id="canonical-94c427c8307c57229d343f35f1a72728c9e7b1b93ec72d316e5f4d4efb8574f1"></a>

## str_value property — items / 776b65fcf749 / 5

Type: `"string"`. Computed.

String. Exclusive with \[ref\_value\]

<a id="canonical-995477810e4cea6ee3d9f434b676ee601428b02493010c35ccf558f5fd2615d9"></a>

<a id="canonical-92a07434d3bd796028f22b28e890f8bc2847525132cb76346396a83dab6d4c4e"></a>

## title property — items / 776b65fcf749 / 6

Type: `"string"`. Computed.

Optional title to be displayed instead of the actual value. Used when pure value doesn't contain all
the needed information to be displayed, or when display titles should be customized.

<a id="canonical-a0c4bdd5e8a0ea8d5eff98dae77af0625f16346a478ca01041c12eca764dc2f6"></a>

<a id="canonical-6d3c9031e46ee9db148e39de426e7d31a1b9f07e7b9be0327891009274c6e2bc"></a>

## value property — items / 776b65fcf749 / 7

Type: `"string"`. Computed.

Suggested value for the field. Should use value\_choice.str\_value instead.

<a id="canonical-7d97be5024d1852bdfd5cf28b5a0238a0b570b9104be7f5bfb2c22a081a69091"></a>

## Next pages — items / 776b65fcf749 / 8

- [items.ref_value](data-sources--bot_suggest_values--reference--group-001.md#canonical-eec009c32f5d57d623ea8dc22541fa11fe4cffaf361a32de8943aef801cd7879)
- [Property reference](data-sources--bot_suggest_values--reference--group-001.md#canonical-1fd46f22c3a31585d917103e1b4692467b3651cd75d9fe25d8cf70497503b572)
- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md#canonical-eb5fdcd8e576fd0893458000f2f7792dfce1074e73bded33582900b2e388a473)

<a id="canonical-eec009c32f5d57d623ea8dc22541fa11fe4cffaf361a32de8943aef801cd7879"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-694db62f41155a1b7cbd0af219c98f75aa95b272c6d7deaba0d89197fc2404e8"></a>

## items.ref_value — items.ref_value / 87bf68b38450 / 2

Breadcrumbs:

- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md#canonical-eb5fdcd8e576fd0893458000f2f7792dfce1074e73bded33582900b2e388a473)
- [Property reference](data-sources--bot_suggest_values--reference--group-001.md#canonical-1fd46f22c3a31585d917103e1b4692467b3651cd75d9fe25d8cf70497503b572)
- [items](data-sources--bot_suggest_values--reference--group-001.md#canonical-5525495cda2ea7ed853ce566ef5caf44c1a3ba52a9e64a3d68e0a6e67f002a69)
- items.ref_value

<a id="canonical-6cf951e374e90f1c3a33ba2f5b18652d67d433ff42659dc11f40aff1b448a091"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

<a id="canonical-b17b549626c81db119a91e8aa1f1cbf8a453115f9b26af4007ab5e5f6596dc62"></a>

## Direct properties — items.ref_value / 87bf68b38450 / 3

<a id="canonical-e7d380398126d837f5046eb411c683f53e01999d44830508f3a9539bce79e4cb"></a>

<a id="canonical-593752a484274e35801b4296637626d580de847beef4396342511720ad89c745"></a>

## name property — items.ref_value / 87bf68b38450 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

<a id="canonical-cc932cd01c8edcc5aeafe5e321cd6c6ff1833431a0547fbae93744304365f9a8"></a>

<a id="canonical-22461867ae631d632b1cbdb4e87214e30dbb7ebb50fbf5ad3f165f4283efd98f"></a>

## namespace property — items.ref_value / 87bf68b38450 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

<a id="canonical-1a4e299ede393d0b79474275b8f50fcbe9c63f2b79c3635c9dad8fb145f012ab"></a>

<a id="canonical-fedbff78cdab224c11bb5b1db6cfb21865a7c55e4a62c9c833af8ea618b8e0b1"></a>

## tenant property — items.ref_value / 87bf68b38450 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

<a id="canonical-7e0a7d8645bfd5891cff9f6fecaab66182a884ef413c76c6f4302076012dafba"></a>

## Next pages — items.ref_value / 87bf68b38450 / 7

- [items](data-sources--bot_suggest_values--reference--group-001.md#canonical-5525495cda2ea7ed853ce566ef5caf44c1a3ba52a9e64a3d68e0a6e67f002a69)
- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md#canonical-eb5fdcd8e576fd0893458000f2f7792dfce1074e73bded33582900b2e388a473)

<a id="canonical-b20993a3bce85dda40d81bbcfa72db5f0eee3674bf37e1dd9ad913470ef04f1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8542b50b11b35edb78d3d16c02299224bcb42b7763b8666c675f9d7322eec1c0"></a>

## request_body — request_body / 92d94034913b / 2

Breadcrumbs:

- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md#canonical-eb5fdcd8e576fd0893458000f2f7792dfce1074e73bded33582900b2e388a473)
- [Property reference](data-sources--bot_suggest_values--reference--group-001.md#canonical-1fd46f22c3a31585d917103e1b4692467b3651cd75d9fe25d8cf70497503b572)
- request_body

<a id="canonical-8f3d259dc114c9afc7c4ce9bbcc80d802ac7ed1f15bfd65903c6eaad38709825"></a>

Type: `"single"`. Optional.

Contains an arbitrary serialized protocol buffer message along with a URL that describes the type of
the serialized message. Protobuf library provides support to pack/unpack Any values in the form of
utility functions or additional generated methods of the Any type. Example 1: Pack and unpack a..

<a id="canonical-a1ca41d376dafa8623af91f630ea997b03ebf44eecf0fbc9c751b600070e38d5"></a>

## Direct properties — request_body / 92d94034913b / 3

<a id="canonical-fadc7c15cab22ec80cc121ec90cf2c9677a3895cde021458d9641e60104cc3bd"></a>

<a id="canonical-e06ecb9d82f5de794d83165b471dad56e241ecf574d155bdfd20055ea1448ab4"></a>

## type_url property — request_body / 92d94034913b / 4

Type: `"string"`. Optional.

URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This
string must contain at least one '/' character. The last segment of the URL path must represent the
fully qualified name of the type (as in ).

<a id="canonical-6fdc993dcba6c9831c90da8de16924851e3608022ed5aaffa7f7e99f7d034ed4"></a>

<a id="canonical-00415e2e56d66fe070e7b4ac5809ade81e8d1c59835366530d5a250efbf5e443"></a>

## value property — request_body / 92d94034913b / 5

Type: `"string"`. Optional.

Must be a valid serialized protocol buffer of the above specified type.

<a id="canonical-9cda1fcf2fd4e93e718156a05d3293fb4c50905f1f33302481396b34f22d48cb"></a>

## Next pages — request_body / 92d94034913b / 6

- [Property reference](data-sources--bot_suggest_values--reference--group-001.md#canonical-1fd46f22c3a31585d917103e1b4692467b3651cd75d9fe25d8cf70497503b572)
- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md#canonical-eb5fdcd8e576fd0893458000f2f7792dfce1074e73bded33582900b2e388a473)
