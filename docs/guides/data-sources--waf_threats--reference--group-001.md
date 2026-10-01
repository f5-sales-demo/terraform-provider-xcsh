---
page_title: "xcsh_waf_threats reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_threats reference."
---

# xcsh_waf_threats reference

<a id="canonical-694a9838f896e006462dd3aa27eeba62fd0d019c7247bec80b961803c8526d8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-486c38bab49a2d04f96b0df8b4badcb99e67d84be555f20aa95a72b439a232eb"></a>

## Property reference — Property reference / 056d873416bc / 2

Breadcrumbs:

- [xcsh_waf_threats](../data-sources/waf_threats.md#canonical-ad4b2c34c0e4ec597b5ffdad9e167fc4c6ccccc97a9939911c04f34090ec91a0)
- Property reference

<a id="canonical-8070857cfeafdd05d7f6af7bd535f5d6d307514c81ca88f01da9e863883a88f0"></a>

## Direct properties — Property reference / 056d873416bc / 3

<a id="canonical-a765bf054877783c93ef18c6ea8b515aa48250e8480dc660126a841050cba196"></a>

<a id="canonical-1a7488683fb19f8bb8f4328e897029b3ba5ff7544ad4394344a162151c90d7e2"></a>

## cursor property — Property reference / 056d873416bc / 4

Type: `"string"`. Optional.

Opaque pagination cursor returned from previous response.

- [cve_ids](data-sources--waf_threats--reference--group-001.md#canonical-f495697c60862e8ccfe2732a1d7c2e9a01d895b7f1fe66fc61618f4968eb1015): complete subsection reference.

<a id="canonical-fafa7ded22a2863b269ce141e3baecec6f405be613947d8073c997ab2ce14c05"></a>

<a id="canonical-f1b2aa18242fe5023869987d8e530394446237f733f1dc9e841e3d0296fe51b8"></a>

## next_cursor property — Property reference / 056d873416bc / 5

Type: `"string"`. Computed.

Next Cursor. Opaque cursor for fetching next page.

<a id="canonical-6fbb746227e49ab2925e71525d2a955aba8aae8b38ecf0e3a68d6a31937f79df"></a>

<a id="canonical-ca9a737a9e1a883fbfb80f218346630bdce36e8ff86dfd4894715bc54b25da77"></a>

## primary_tag property — Property reference / 056d873416bc / 6

Type: `"string"`. Optional.

Exclusive with \[cve\_ids waf\_sec\_event\_id\] Primary tag to filter threats. A primary tag is a
high level categorization of a threat, such as an associated threat actor, malware family or CVE.

<a id="canonical-b3fdbae1cf063b54be2f148a69a42081b4782310db135002b32715f8bc46468b"></a>

<a id="canonical-6f22296dfa7732efb13e0bb9792e824c3d2027a97dfab1f5e6a7520085ebebd2"></a>

## report_fields property — Property reference / 056d873416bc / 7

Type: `["list", "string"]`. Optional.

Optional list of fields to include in threat representation.

<a id="canonical-5e7cc657c0f64284c90f7c6f78abb74d776ffab47186c6730ba8b73cdaedae21"></a>

<a id="canonical-a1d9a297995af424e4e9079a2172aac7ff8644d38a12d911b4ec04b0c0f224b4"></a>

## threats property — Property reference / 056d873416bc / 8

Type: `["list", "string"]`. Computed.

Threats. A list of threats that match the query.

<a id="canonical-38d4b53ab5d9af43160551949412c8f78e229d17e73a2be7e671d9f9d8f15eee"></a>

<a id="canonical-eb03d48145036b79f5627d67a7959b9c53db9d5cf173a81fcfc93e04d3dab618"></a>

## waf_sec_event_id property — Property reference / 056d873416bc / 9

Type: `"string"`. Optional.

Exclusive with \[cve\_ids primary\_tag\] WAF Security Event ID to GET associated threats
information.

<a id="canonical-30c5f602551c7c5ed0c7a8b7cd902a000bba2b73d2109a60c8f2deb6b8f60472"></a>

## All schema paths — Property reference / 056d873416bc / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `cursor` | [cursor](data-sources--waf_threats--reference--group-001.md#canonical-a765bf054877783c93ef18c6ea8b515aa48250e8480dc660126a841050cba196) |
| `cve_ids` | [cve_ids](data-sources--waf_threats--reference--group-001.md#canonical-93a66a02105712fd395e1c93594c78d380708690b2dd9d0179b75d8042afe42c) |
| `cve_ids.ids` | [cve_ids.ids](data-sources--waf_threats--reference--group-001.md#canonical-7d9c43fb31f75d1b29ce496aff066dd7bda5654efa6f27fb4569e5f971d06c6f) |
| `next_cursor` | [next_cursor](data-sources--waf_threats--reference--group-001.md#canonical-fafa7ded22a2863b269ce141e3baecec6f405be613947d8073c997ab2ce14c05) |
| `primary_tag` | [primary_tag](data-sources--waf_threats--reference--group-001.md#canonical-6fbb746227e49ab2925e71525d2a955aba8aae8b38ecf0e3a68d6a31937f79df) |
| `report_fields` | [report_fields](data-sources--waf_threats--reference--group-001.md#canonical-b3fdbae1cf063b54be2f148a69a42081b4782310db135002b32715f8bc46468b) |
| `threats` | [threats](data-sources--waf_threats--reference--group-001.md#canonical-5e7cc657c0f64284c90f7c6f78abb74d776ffab47186c6730ba8b73cdaedae21) |
| `waf_sec_event_id` | [waf_sec_event_id](data-sources--waf_threats--reference--group-001.md#canonical-38d4b53ab5d9af43160551949412c8f78e229d17e73a2be7e671d9f9d8f15eee) |

<a id="canonical-23a89917a93e77a457b92e367b137ccbaf2845459949c24efb8014913fbdadbf"></a>

## Next pages — Property reference / 056d873416bc / 11

- [cve_ids](data-sources--waf_threats--reference--group-001.md#canonical-f495697c60862e8ccfe2732a1d7c2e9a01d895b7f1fe66fc61618f4968eb1015)
- [xcsh_waf_threats](../data-sources/waf_threats.md#canonical-ad4b2c34c0e4ec597b5ffdad9e167fc4c6ccccc97a9939911c04f34090ec91a0)

<a id="canonical-f495697c60862e8ccfe2732a1d7c2e9a01d895b7f1fe66fc61618f4968eb1015"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e5bd00c5e0b0d8ca8a213920058aa5eabab0affefd5b066d067214e6e3e4b6d"></a>

## cve_ids — cve_ids / 28a5ea1c11f1 / 2

Breadcrumbs:

- [xcsh_waf_threats](../data-sources/waf_threats.md#canonical-ad4b2c34c0e4ec597b5ffdad9e167fc4c6ccccc97a9939911c04f34090ec91a0)
- [Property reference](data-sources--waf_threats--reference--group-001.md#canonical-694a9838f896e006462dd3aa27eeba62fd0d019c7247bec80b961803c8526d8e)
- cve_ids

<a id="canonical-93a66a02105712fd395e1c93594c78d380708690b2dd9d0179b75d8042afe42c"></a>

Type: `"single"`. Optional.

CVE ID List. A list of CVE IDs.

<a id="canonical-4ba8d41601045df8cc59255bd02e4930be11ad5e797d30a6dd5337efae5b410f"></a>

## Direct properties — cve_ids / 28a5ea1c11f1 / 3

<a id="canonical-7d9c43fb31f75d1b29ce496aff066dd7bda5654efa6f27fb4569e5f971d06c6f"></a>

<a id="canonical-efbb1081a5bf74cc61c9be2b18a8b0d0cd0d4c65f52b82077e7906dd97ea7fe8"></a>

## ids property — cve_ids / 28a5ea1c11f1 / 4

Type: `["list", "string"]`. Optional.

IDs. A list of CVE IDs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

<a id="canonical-dc41114cbe1ee1dc820fd51bded09fa7408a98b731522895f3d625270cb98db7"></a>

## Next pages — cve_ids / 28a5ea1c11f1 / 5

- [Property reference](data-sources--waf_threats--reference--group-001.md#canonical-694a9838f896e006462dd3aa27eeba62fd0d019c7247bec80b961803c8526d8e)
- [xcsh_waf_threats](../data-sources/waf_threats.md#canonical-ad4b2c34c0e4ec597b5ffdad9e167fc4c6ccccc97a9939911c04f34090ec91a0)
