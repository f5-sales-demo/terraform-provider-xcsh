---
page_title: "xcsh_network_data_intelligence reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_data_intelligence reference."
---

# xcsh_network_data_intelligence reference

<a id="canonical-c658a6fbfe8dbbe5fd2dab0eb03fbfc881b85553026886291a106a251c844753"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f9504c31f49972c544ad7f3e9c519c2c968852532076989fbd8885f98ca3bcf"></a>

## Property reference — Property reference / 45f1c0b4109e / 2

Breadcrumbs:

- [xcsh_network_data_intelligence](../data-sources/network_data_intelligence.md#canonical-db1c6e7c4254880692807159269c6d76d21575ca6575b838319cf6c66703a647)
- Property reference

<a id="canonical-024a85e288cc29786449628610684d6b6c3089dfdb10e948715cff3161e56344"></a>

## Direct properties — Property reference / 45f1c0b4109e / 3

<a id="canonical-da8fc611191654c28af39accc28e4f76ed22db31a31aa18f6bb34c1948c7a72f"></a>

<a id="canonical-e02d3c1c125f552677bca34e3d037ed99a52e99172a332461f44e5d550d31a34"></a>

## api_release_tag property — Property reference / 45f1c0b4109e / 4

Type: `"string"`. Computed.

Pinned api-specs-enriched release tag compiled into this provider.

<a id="canonical-2e10f9aeeebe426b3395d3c82276af38b804b83a37032c07fa8f0993f0bbdfec"></a>

<a id="canonical-dccf6833845408f81f87ef79be34da644096462a0585adc38a819bcad574248b"></a>

## cidr_blocks property — Property reference / 45f1c0b4109e / 5

Type: `["list", "string"]`. Computed.

Sorted unique IPv4 CIDRs across selected regions. Individual IPv4 addresses are normalized to /32.

<a id="canonical-f9906ef2d8fda372c799c40fbdb8189743fe93b3854039732ac40f16759bb6f9"></a>

<a id="canonical-ebaa923995485aaf4659f4f4241c239a5386d0e4ffb0c24e3356092cf7bc8be9"></a>

## cidr_blocks_by_region property — Property reference / 45f1c0b4109e / 6

Type: `["map", ["list", "string"]]`. Computed.

Sorted unique IPv4 CIDRs keyed by selected published region.

<a id="canonical-2bf02ce8bc04e13d4f9795350de31134407e15e6896dfb2d3e85426730ca854d"></a>

<a id="canonical-ab7318c77c33ffe485855235de2af79e7b68e7580b69c86e9b42d00f7a856f11"></a>

## id property — Property reference / 45f1c0b4109e / 7

Type: `"string"`. Computed.

Stable identifier derived from the pinned source digest, data-source group, and selected regions.

<a id="canonical-4cc4a9d8208ac0db3573f8210d9e4e6ff3b93741d6ce1585a09d326a9d578aa9"></a>

<a id="canonical-98eb1b2b35e58dfda73f88f194db59d6761365fd931f894d7f10ea3a1608a512"></a>

## manifest_generated_at property — Property reference / 45f1c0b4109e / 8

Type: `"string"`. Computed.

Generation timestamp reported by the published network allowlist manifest.

<a id="canonical-ec391f06f7569cfec1ef0868c7924cf9af8a2ac076d5a86ee6a89c73d7c1b0f9"></a>

<a id="canonical-b4e2ad1f4abd7382cefff03476550689e390d38cca8409135f5080704415fe92"></a>

## regions property — Property reference / 45f1c0b4109e / 9

Type: `["set", "string"]`. Optional, Computed.

Published regions to include. Omit to select every region.

<a id="canonical-e97ea91ab19077745fa23c762173fd576b3c03179980df66c52bd66c11a6fa92"></a>

<a id="canonical-38facd1424bf12835b71776191d5afbfa78bbbd1090c1c38194f69ddacd5fa65"></a>

## source_entries property — Property reference / 45f1c0b4109e / 10

Type: `["list", "string"]`. Computed.

Selected IPv4 entries exactly as represented by the published regions, ordered by region name and
source order.

<a id="canonical-c75f1ed790480ad915c0e74672b9000fab5aad4fe361da23dbcf1211c5a1aa5f"></a>

<a id="canonical-adc6f09cf41561da2b10ff46d5c15f939584e2b57d7ce6a8ba5e6fbcb079aaa5"></a>

## source_sha256 property — Property reference / 45f1c0b4109e / 11

Type: `"string"`. Computed.

SHA-256 digest of the canonical published network allowlist manifest.

<a id="canonical-8f456ec966e25995b2f783a4724a2d2068540517c9e581fb9bb149848bed9b51"></a>

<a id="canonical-2a5bd5c01ab0d6d285eedc3785ccbd62d36efbbda69dfca366bcfa6521371dfe"></a>

## source_url property — Property reference / 45f1c0b4109e / 12

Type: `"string"`. Computed.

Published F5 network allowlist source URL.

<a id="canonical-5a826c9d7ba44275a9a7da488dc19fd4b437750364158616af68b95821f0b043"></a>

## All schema paths — Property reference / 45f1c0b4109e / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `api_release_tag` | [api_release_tag](data-sources--network_data_intelligence--reference--group-001.md#canonical-da8fc611191654c28af39accc28e4f76ed22db31a31aa18f6bb34c1948c7a72f) |
| `cidr_blocks` | [cidr_blocks](data-sources--network_data_intelligence--reference--group-001.md#canonical-2e10f9aeeebe426b3395d3c82276af38b804b83a37032c07fa8f0993f0bbdfec) |
| `cidr_blocks_by_region` | [cidr_blocks_by_region](data-sources--network_data_intelligence--reference--group-001.md#canonical-f9906ef2d8fda372c799c40fbdb8189743fe93b3854039732ac40f16759bb6f9) |
| `id` | [id](data-sources--network_data_intelligence--reference--group-001.md#canonical-2bf02ce8bc04e13d4f9795350de31134407e15e6896dfb2d3e85426730ca854d) |
| `manifest_generated_at` | [manifest_generated_at](data-sources--network_data_intelligence--reference--group-001.md#canonical-4cc4a9d8208ac0db3573f8210d9e4e6ff3b93741d6ce1585a09d326a9d578aa9) |
| `regions` | [regions](data-sources--network_data_intelligence--reference--group-001.md#canonical-ec391f06f7569cfec1ef0868c7924cf9af8a2ac076d5a86ee6a89c73d7c1b0f9) |
| `source_entries` | [source_entries](data-sources--network_data_intelligence--reference--group-001.md#canonical-e97ea91ab19077745fa23c762173fd576b3c03179980df66c52bd66c11a6fa92) |
| `source_sha256` | [source_sha256](data-sources--network_data_intelligence--reference--group-001.md#canonical-c75f1ed790480ad915c0e74672b9000fab5aad4fe361da23dbcf1211c5a1aa5f) |
| `source_url` | [source_url](data-sources--network_data_intelligence--reference--group-001.md#canonical-8f456ec966e25995b2f783a4724a2d2068540517c9e581fb9bb149848bed9b51) |

<a id="canonical-89929c840beac384445a6adf97bc8c6d92df70d1a7735969b98deea5a69e761c"></a>

## Next pages — Property reference / 45f1c0b4109e / 14

- [xcsh_network_data_intelligence](../data-sources/network_data_intelligence.md#canonical-db1c6e7c4254880692807159269c6d76d21575ca6575b838319cf6c66703a647)
