---
page_title: "xcsh_dns_zone_delete_cryptokey reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_delete_cryptokey reference."
---

# xcsh_dns_zone_delete_cryptokey reference

<a id="canonical-53390e8981cf8014bda7b0838c903adf614b4fc57c56d4927129396bb9cc67aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3cc9aadc725d6a78798e7ad9900f532f36235a4fcea5f3c9bb394342deedddc7"></a>

## Property reference — Property reference / 0f23171a0858 / 2

Breadcrumbs:

- [xcsh_dns_zone_delete_cryptokey](../actions/dns_zone_delete_cryptokey.md#canonical-d0f11b9ff55eda3e0a21dbaf9c1ae6f999a193083bac154e03bd47c8807d0123)
- Property reference

<a id="canonical-0a6730b568f5dad141512992d19617e862f2bf3d17374515e8a471b9e96ba945"></a>

## Direct properties — Property reference / 0f23171a0858 / 3

<a id="canonical-f1943d3315f7d35928044cdc85b3fed62dcf4c4eba8db4c4ae851fdc6ed0a736"></a>

<a id="canonical-6ec1ffc83b611b28aae74a7d95ed68d8ab1a2661b5df47745a23e937af724720"></a>

## key_id property — Property reference / 0f23171a0858 / 4

Type: `"number"`. Optional.

Key ID. Unique identifier for this resource

<a id="canonical-ebe4d31e1f2495a61608623dea603435862082914b630f92aacc365c67fcde03"></a>

<a id="canonical-b93d82e1ecdeb96953dfd7d9445d14dbc762f7dd77471f6d6612eee619532008"></a>

## namespace property — Property reference / 0f23171a0858 / 5

Type: `"string"`. Optional.

Namespace is always system for dns\_zone.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-a2087800757573f212fdb2c5670af215c3173b7a21228c5ec2bfc7144552645f"></a>

<a id="canonical-3fc5e900597005b996e7784b6e3fcc5e881213ca9063b43d534a5a2cfe79c3f1"></a>

## zone_name property — Property reference / 0f23171a0858 / 6

Type: `"string"`. Optional.

Zone Name. Human-readable name for the resource

<a id="canonical-305e07b10650e612391831bc96eb6af89a12e726a7c09aa1cb2fd3575bfbbaeb"></a>

## All schema paths — Property reference / 0f23171a0858 / 7

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `key_id` | [key_id](actions--dns_zone_delete_cryptokey--reference--group-001.md#canonical-f1943d3315f7d35928044cdc85b3fed62dcf4c4eba8db4c4ae851fdc6ed0a736) |
| `namespace` | [namespace](actions--dns_zone_delete_cryptokey--reference--group-001.md#canonical-ebe4d31e1f2495a61608623dea603435862082914b630f92aacc365c67fcde03) |
| `zone_name` | [zone_name](actions--dns_zone_delete_cryptokey--reference--group-001.md#canonical-a2087800757573f212fdb2c5670af215c3173b7a21228c5ec2bfc7144552645f) |

<a id="canonical-d353b3a04dc5c8e71cb2bd1c94dd202c273b43c0a586629094ffff7f134fe77a"></a>

## Next pages — Property reference / 0f23171a0858 / 8

- [xcsh_dns_zone_delete_cryptokey](../actions/dns_zone_delete_cryptokey.md#canonical-d0f11b9ff55eda3e0a21dbaf9c1ae6f999a193083bac154e03bd47c8807d0123)
