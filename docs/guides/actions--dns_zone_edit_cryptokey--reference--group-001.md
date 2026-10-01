---
page_title: "xcsh_dns_zone_edit_cryptokey reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_edit_cryptokey reference."
---

# xcsh_dns_zone_edit_cryptokey reference

<a id="canonical-8c778cc8bc7098812305b349d8624d6e5e0355d7a00fe2afec983f56a25a2bb4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f047906cb6ad8fcb57b13019ae79f08aaa4ae70889e207e1c221a8877cee741"></a>

## Property reference — Property reference / 9a1ce105753d / 2

Breadcrumbs:

- [xcsh_dns_zone_edit_cryptokey](../actions/dns_zone_edit_cryptokey.md#canonical-5ea34fd0806e345b9da4127afca9297afdff2223ccf404858b3a86b9b6048b4c)
- Property reference

<a id="canonical-092f9427bf019ea05f2961349e524aeaaaf748168bd882982fbc08f13c564e7d"></a>

## Direct properties — Property reference / 9a1ce105753d / 3

<a id="canonical-34603e784b9653ef12a2458c77a90d39ca63d85893716cb8c42c597b16b4fb1a"></a>

<a id="canonical-b4c7175e0583af63e9c38eaf1e7e6cd5ccfe3ac24684a4131cb75f48cb4fee48"></a>

## active property — Property reference / 9a1ce105753d / 4

Type: `"bool"`. Optional.

Active. Indicates if the resource is active

<a id="canonical-c42bfe86e6fdeefd350e1af488af7688443e3192e86dbb5fde85605a7034fe5f"></a>

<a id="canonical-fa108b0f19d58c4f73af05a38ba37eba17a42567c853fb7157cd9a9071301753"></a>

## key_id property — Property reference / 9a1ce105753d / 5

Type: `"number"`. Optional.

Key ID. Unique identifier for this resource

<a id="canonical-31804d473a5f0cc8fe8d712e515bc4da097507a5d6ac9068e7728ff4db6b1d9e"></a>

<a id="canonical-43e63ae5ff16f952cf4e18c9b03638e29be962cbab397c772c27f3083d6a7e3a"></a>

## namespace property — Property reference / 9a1ce105753d / 6

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

<a id="canonical-db47e87ee1bb314656cf4912e3bfe23fa3f089025e0e652baf91687cc805680a"></a>

<a id="canonical-a2da4279cfb6a39cbeb3279d9c8c79f21c20d6e07bddc7dcab2e4a8f222d9dad"></a>

## zone_name property — Property reference / 9a1ce105753d / 7

Type: `"string"`. Optional.

Zone Name. Human-readable name for the resource

<a id="canonical-981a76ddf34dfcc15dbbee3f362046374125477435777c488e5d217b37580d15"></a>

## All schema paths — Property reference / 9a1ce105753d / 8

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active` | [active](actions--dns_zone_edit_cryptokey--reference--group-001.md#canonical-34603e784b9653ef12a2458c77a90d39ca63d85893716cb8c42c597b16b4fb1a) |
| `key_id` | [key_id](actions--dns_zone_edit_cryptokey--reference--group-001.md#canonical-c42bfe86e6fdeefd350e1af488af7688443e3192e86dbb5fde85605a7034fe5f) |
| `namespace` | [namespace](actions--dns_zone_edit_cryptokey--reference--group-001.md#canonical-31804d473a5f0cc8fe8d712e515bc4da097507a5d6ac9068e7728ff4db6b1d9e) |
| `zone_name` | [zone_name](actions--dns_zone_edit_cryptokey--reference--group-001.md#canonical-db47e87ee1bb314656cf4912e3bfe23fa3f089025e0e652baf91687cc805680a) |

<a id="canonical-35498dc095a87ab37bd60981f0ffedd11f968e3d95f356a2de79f0071028328a"></a>

## Next pages — Property reference / 9a1ce105753d / 9

- [xcsh_dns_zone_edit_cryptokey](../actions/dns_zone_edit_cryptokey.md#canonical-5ea34fd0806e345b9da4127afca9297afdff2223ccf404858b3a86b9b6048b4c)
