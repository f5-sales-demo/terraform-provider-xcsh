---
page_title: "xcsh_dns_zone_add_cryptokey reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_add_cryptokey reference."
---

# xcsh_dns_zone_add_cryptokey reference

<a id="canonical-a2054223eb59702f779639edabed8972fc3dc1170f0410220f43b85cfc6e4f06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad8d3cd152854f3f92d9e5467c496b3df73ebe218efdb9ec490b8ba90a0cc04e"></a>

## Property reference — Property reference / 2fff315da8ab / 2

Breadcrumbs:

- [xcsh_dns_zone_add_cryptokey](../actions/dns_zone_add_cryptokey.md#canonical-85ff7b5ce0689a3d0497afdd728eedfead55ce6ef5288db5ba75c83a0e169984)
- Property reference

<a id="canonical-4ea27ba0da73bde51fa7b4d20f8ff9e624426298a031036ce1405313de87fbe6"></a>

## Direct properties — Property reference / 2fff315da8ab / 3

<a id="canonical-c3935d8683bd7ce877bbc375beda55988494294927b40b3d03448fc73ec780ce"></a>

<a id="canonical-6cbdabf19ae4566bd1c9b0075b0337bc66b2dad85691bbc1cc41329d974b71ba"></a>

## key_type property — Property reference / 2fff315da8ab / 4

Type: `"string"`. Optional.

Key Type. Type or category classification

<a id="canonical-c7b9ea0317cef58a3c33e2adc9bceb9ed11077f69d51133387e00f6029e540b2"></a>

<a id="canonical-de74bed08495c13d1ce57a9313d5ac7a7a6f284ed71c9ee8260b80be35c1b024"></a>

## namespace property — Property reference / 2fff315da8ab / 5

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

<a id="canonical-2e4506e4e3ecfe08db8a1d61bb93e7048a5ee702f5017e835de4f3134f8f6488"></a>

<a id="canonical-e31c72a149b6d3f8d660ea632c3c210196ca6c7d2b58f1cd37ebe45ff5723b00"></a>

## zone_name property — Property reference / 2fff315da8ab / 6

Type: `"string"`. Optional.

Zone Name. Human-readable name for the resource

<a id="canonical-e599fc075ae690d9551912917a7445b2fccc53dd144ed331a054e53b512c0373"></a>

## All schema paths — Property reference / 2fff315da8ab / 7

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `key_type` | [key_type](actions--dns_zone_add_cryptokey--reference--group-001.md#canonical-c3935d8683bd7ce877bbc375beda55988494294927b40b3d03448fc73ec780ce) |
| `namespace` | [namespace](actions--dns_zone_add_cryptokey--reference--group-001.md#canonical-c7b9ea0317cef58a3c33e2adc9bceb9ed11077f69d51133387e00f6029e540b2) |
| `zone_name` | [zone_name](actions--dns_zone_add_cryptokey--reference--group-001.md#canonical-2e4506e4e3ecfe08db8a1d61bb93e7048a5ee702f5017e835de4f3134f8f6488) |

<a id="canonical-4b529c8e7361638b5b99c974f4215a842ae7d1f47150714de2cae0df0a280f15"></a>

## Next pages — Property reference / 2fff315da8ab / 8

- [xcsh_dns_zone_add_cryptokey](../actions/dns_zone_add_cryptokey.md#canonical-85ff7b5ce0689a3d0497afdd728eedfead55ce6ef5288db5ba75c83a0e169984)
