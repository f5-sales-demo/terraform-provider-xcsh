---
page_title: "xcsh_dns_zone_cryptokeys reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_cryptokeys reference."
---

# xcsh_dns_zone_cryptokeys reference

<a id="canonical-9b19874018ee99a46c60d650f6c9b63d534c81dd0ca88d63f5aa54f3b74462eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f65bc3a5da226b35e7862b072a40aa2e11eedbed1d0e6fceb2245be03737af27"></a>

## Property reference — Property reference / 25a93ab24a54 / 2

Breadcrumbs:

- [xcsh_dns_zone_cryptokeys](../data-sources/dns_zone_cryptokeys.md#canonical-f52d736244c299ba5f4852780f6f61f432e002050062705e557a852e68c45332)
- Property reference

<a id="canonical-458affdaae068a18f872547ae34da804257e06c15f43fe0908caf5acf65a17f3"></a>

## Direct properties — Property reference / 25a93ab24a54 / 3

- [keys](data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-6a6e0edd017f7645baeebbedf5a11b2843eb8312c2fd20ab52b8b40857673b2f): complete subsection reference.

<a id="canonical-ff40b03f5cacdfbdd0b6774258f92ad9f24a82400780b059c0f64c89e3471f82"></a>

<a id="canonical-ced8b9f3bcf0c74f293e32f98978234399d90970855d5bce4982ff77b0ebf337"></a>

## namespace property — Property reference / 25a93ab24a54 / 4

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

<a id="canonical-62847fca115ada614a4e92a3d39f63ee43ed94c83c69ccc63f07bd8bfa83f124"></a>

<a id="canonical-38499b43c865fbdbe30087f0475e07cdca2df02720350b4b90fdb2675ae962b9"></a>

## zone_name property — Property reference / 25a93ab24a54 / 5

Type: `"string"`. Optional.

Zone Name. Human-readable name for the resource

<a id="canonical-4fb1dab0d900e8d12d0bfa96fa6e867ac7f4c3f3e36d748a6cc2563d62e1272e"></a>

## All schema paths — Property reference / 25a93ab24a54 / 6

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `keys` | [keys](data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-b81b6527098f1465f1eaefcd3b90c7d11109ece2e3cfaa5adbaab2f64adac13f) |
| `keys.active` | [keys.active](data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-d7b9b16d0da9d7eb1e74d41700d1f1af090a4552bae67ee290f5429f2eb21e33) |
| `keys.algorithm` | [keys.algorithm](data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-a5683fca9f1406eedb391d7275acf3bc36634f3ef9e4b9094e65c656b0047235) |
| `keys.dnskey` | [keys.dnskey](data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-3648bc6ec690f82ec7551d20f7ed4c2ff2dc69f2b0d43b46f93315e0ed010246) |
| `keys.key_id` | [keys.key_id](data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-634742aae3ba620eefed8cdbc6b902c094b6ccb5ab7959497b40d48b57087003) |
| `keys.key_type` | [keys.key_type](data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-6260ea807148790e332aff1616ae216ab16c94c84630def640bab2fc77a941e5) |
| `keys.published` | [keys.published](data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-2d983011dfa77c360710e502ca44db665ec309f4d61dccf8f178d858fe22abf1) |
| `keys.type` | [keys.type](data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-b62901bcc7755c29c3880f4c491d6251029270fba1732ec317f7e313666f371c) |
| `namespace` | [namespace](data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-ff40b03f5cacdfbdd0b6774258f92ad9f24a82400780b059c0f64c89e3471f82) |
| `zone_name` | [zone_name](data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-62847fca115ada614a4e92a3d39f63ee43ed94c83c69ccc63f07bd8bfa83f124) |

<a id="canonical-88ca4bc449b0511bbae2fe43ce2c8d9de197cae7008fc7c45ab7412bab0c84af"></a>

## Next pages — Property reference / 25a93ab24a54 / 7

- [keys](data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-6a6e0edd017f7645baeebbedf5a11b2843eb8312c2fd20ab52b8b40857673b2f)
- [xcsh_dns_zone_cryptokeys](../data-sources/dns_zone_cryptokeys.md#canonical-f52d736244c299ba5f4852780f6f61f432e002050062705e557a852e68c45332)

<a id="canonical-6a6e0edd017f7645baeebbedf5a11b2843eb8312c2fd20ab52b8b40857673b2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40e45148cd24fa84affc5f4152acffad6f537b2a73fd11cc38d23f07d18b39db"></a>

## keys — keys / d108170e7da7 / 2

Breadcrumbs:

- [xcsh_dns_zone_cryptokeys](../data-sources/dns_zone_cryptokeys.md#canonical-f52d736244c299ba5f4852780f6f61f432e002050062705e557a852e68c45332)
- [Property reference](data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-9b19874018ee99a46c60d650f6c9b63d534c81dd0ca88d63f5aa54f3b74462eb)
- keys

<a id="canonical-b81b6527098f1465f1eaefcd3b90c7d11109ece2e3cfaa5adbaab2f64adac13f"></a>

Type: `"list"`. Computed.

Configuration parameter for keys

<a id="canonical-81b2a8776be4a215a6251d077111f4c643fab75ceddb3b3a9a584d5e90d2ce59"></a>

## Direct properties — keys / d108170e7da7 / 3

<a id="canonical-d7b9b16d0da9d7eb1e74d41700d1f1af090a4552bae67ee290f5429f2eb21e33"></a>

<a id="canonical-9cb793bec5f102db82f8189b1f590cc77483388136c5c0dc55b7d23d3cf903a8"></a>

## active property — keys / d108170e7da7 / 4

Type: `"bool"`. Computed.

Whether the key is currently active for signing.

<a id="canonical-a5683fca9f1406eedb391d7275acf3bc36634f3ef9e4b9094e65c656b0047235"></a>

<a id="canonical-0b51e8f84825ea4b0fa7222f8e2ec4366e5028f92aaba7b7ff4f46f3bcf104f5"></a>

## algorithm property — keys / d108170e7da7 / 5

Type: `"string"`. Computed.

The DNSSEC signing algorithm used by this key.

<a id="canonical-3648bc6ec690f82ec7551d20f7ed4c2ff2dc69f2b0d43b46f93315e0ed010246"></a>

<a id="canonical-98362b03f6de5e3f9d62195c111452919f3423fc664bb21a795638d22116dac4"></a>

## dnskey property — keys / d108170e7da7 / 6

Type: `"string"`. Computed.

DNSKEY. The DNSKEY record data for this key.

<a id="canonical-634742aae3ba620eefed8cdbc6b902c094b6ccb5ab7959497b40d48b57087003"></a>

<a id="canonical-74a80d36ed2c96952836ffc3d824fe15c6c968b39b7bfaa85091d148d8db9b39"></a>

## key_id property — keys / d108170e7da7 / 7

Type: `"number"`. Computed.

Unique identifier for the cryptographic key.

<a id="canonical-6260ea807148790e332aff1616ae216ab16c94c84630def640bab2fc77a941e5"></a>

<a id="canonical-9754a5d86898ecde8de980630af4bdcd5ac7fcb2b4d39e151a52b6dd7e657774"></a>

## key_type property — keys / d108170e7da7 / 8

Type: `"string"`. Computed.

The cryptographic key type (e.g., CSK, KSK, ZSK).

<a id="canonical-2d983011dfa77c360710e502ca44db665ec309f4d61dccf8f178d858fe22abf1"></a>

<a id="canonical-5f2e637202eda639e48a23cb2a2bc2aefad797fb9c094f84a9ea1535366f324b"></a>

## published property — keys / d108170e7da7 / 9

Type: `"bool"`. Computed.

Whether the key is published in the DNSKEY RRset.

<a id="canonical-b62901bcc7755c29c3880f4c491d6251029270fba1732ec317f7e313666f371c"></a>

<a id="canonical-3e4557f354fced17b7aa3f027419952f8a242b674a18aba43acf68e4fb1f051b"></a>

## type property — keys / d108170e7da7 / 10

Type: `"string"`. Computed.

Type. Should always be 'CryptoKey'

<a id="canonical-f874ba01270adefcebf44cf735e16d8572253cd90a4a4c789a4cb90771beb09f"></a>

## Next pages — keys / d108170e7da7 / 11

- [Property reference](data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-9b19874018ee99a46c60d650f6c9b63d534c81dd0ca88d63f5aa54f3b74462eb)
- [xcsh_dns_zone_cryptokeys](../data-sources/dns_zone_cryptokeys.md#canonical-f52d736244c299ba5f4852780f6f61f432e002050062705e557a852e68c45332)
