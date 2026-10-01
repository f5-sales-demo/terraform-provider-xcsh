---
page_title: "xcsh_network_regional_edges reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_regional_edges reference."
---

# xcsh_network_regional_edges reference

<a id="canonical-e079a6e575809cb77caea1bfb1f018d98eced62e8e5590c688ca25690a4e079a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31a020b79843150f79d7635a6bd7497edce4634cae9ac8e1168706a1d6855f85"></a>

## Property reference — Property reference / 28a2240a0595 / 2

Breadcrumbs:

- [xcsh_network_regional_edges](../data-sources/network_regional_edges.md#canonical-4b07492253f8f10e61287f15adccadbfc866bfc8be3823bb27bb4e02b5d71089)
- Property reference

<a id="canonical-71c5e4f53e5c34715c0c359d4c5593749d278afedd407842f27fd996620613c3"></a>

## Direct properties — Property reference / 28a2240a0595 / 3

<a id="canonical-03eeb31109638d1e04c2afc7ebffec1f662bc88499441469b65b239f38debb08"></a>

<a id="canonical-f5db7ed7bb1353a9cec02b4065635c08a527e4176648ac2b04c7a298e45c126f"></a>

## api_release_tag property — Property reference / 28a2240a0595 / 4

Type: `"string"`. Computed.

Pinned api-specs-enriched release tag compiled into this provider.

<a id="canonical-21575448a6f4b5f028a3025ae29d0515bc3f1ae157401a0a4d824f1f38991c80"></a>

<a id="canonical-e60b1d6cb13d8d0d8ed1c183995508667363a4c12e7ac603b0d2bd8e1689ddb0"></a>

## cidr_blocks property — Property reference / 28a2240a0595 / 5

Type: `["list", "string"]`. Computed.

Sorted unique IPv4 CIDRs across selected regions. Individual IPv4 addresses are normalized to /32.

<a id="canonical-eb1d173e52ff97ddc0c04ee9c03349b42b6f829372c1c581a91ec8cb5dc42119"></a>

<a id="canonical-eb6a6f388e3641f3a9f2b5c956446ac45bd411d4de727c9db11d9cf717ba3506"></a>

## cidr_blocks_by_region property — Property reference / 28a2240a0595 / 6

Type: `["map", ["list", "string"]]`. Computed.

Sorted unique IPv4 CIDRs keyed by selected published region.

<a id="canonical-e320b2a22b0d16b616d0570c1e4de3dafea2c25c381bb2af7c01633ce1abb022"></a>

<a id="canonical-f779aa137d181c0fc5eabf144dc63c1bc5801bfd0cc55bdfbe986f867ab79886"></a>

## id property — Property reference / 28a2240a0595 / 7

Type: `"string"`. Computed.

Stable identifier derived from the pinned source digest, data-source group, and selected regions.

<a id="canonical-aab5966293dfa9543b6adc7b2112ac268872ac1adf012b932c8c86b6f703d79a"></a>

<a id="canonical-b0a29936ea50b32320cc5d21b1feeac51239af6c24ebbc3df9bb00dc956afeb6"></a>

## manifest_generated_at property — Property reference / 28a2240a0595 / 8

Type: `"string"`. Computed.

Generation timestamp reported by the published network allowlist manifest.

<a id="canonical-127d4ce3ad9cfceb9ea1a082d67601d770cce38e468a54464fb9d47272c5ab34"></a>

<a id="canonical-2a6e811dd7e04830a1d0fec98957b04cea77a61e9cce2c9ac55a0ff8de5df573"></a>

## regions property — Property reference / 28a2240a0595 / 9

Type: `["set", "string"]`. Optional, Computed.

Published regions to include. Omit to select every region.

<a id="canonical-be98a5d8aa25674f566106ed2a34a5741196fcde3477972f49e70e0863442f33"></a>

<a id="canonical-566b5fde765ec036a5065adbfc824cccdcc6a8faeb484c58469f7447a0b23b32"></a>

## source_entries property — Property reference / 28a2240a0595 / 10

Type: `["list", "string"]`. Computed.

Selected IPv4 entries exactly as represented by the published regions, ordered by region name and
source order.

<a id="canonical-c4d83c3b7c5830cdb0856ba6249468cac8ce0fbb15da070f49a396e34de55831"></a>

<a id="canonical-0db447ad6301f6fb91d39ec7bc66014881ecaeeec3ad2e14933326ac86e4353a"></a>

## source_sha256 property — Property reference / 28a2240a0595 / 11

Type: `"string"`. Computed.

SHA-256 digest of the canonical published network allowlist manifest.

<a id="canonical-e9677830b089cb20a94ccc844a5a9b59f5a10baa309127302be0276a0d0541a9"></a>

<a id="canonical-d30ae320cc5c87640a482edee7ed49c244e121da885acdd22405d5d2a409a15b"></a>

## source_url property — Property reference / 28a2240a0595 / 12

Type: `"string"`. Computed.

Published F5 network allowlist source URL.

<a id="canonical-3f86673ec82a0602a58d2b668cf64339113eb20d920cd6cfac18845eaa355d57"></a>

## All schema paths — Property reference / 28a2240a0595 / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `api_release_tag` | [api_release_tag](data-sources--network_regional_edges--reference--group-001.md#canonical-03eeb31109638d1e04c2afc7ebffec1f662bc88499441469b65b239f38debb08) |
| `cidr_blocks` | [cidr_blocks](data-sources--network_regional_edges--reference--group-001.md#canonical-21575448a6f4b5f028a3025ae29d0515bc3f1ae157401a0a4d824f1f38991c80) |
| `cidr_blocks_by_region` | [cidr_blocks_by_region](data-sources--network_regional_edges--reference--group-001.md#canonical-eb1d173e52ff97ddc0c04ee9c03349b42b6f829372c1c581a91ec8cb5dc42119) |
| `id` | [id](data-sources--network_regional_edges--reference--group-001.md#canonical-e320b2a22b0d16b616d0570c1e4de3dafea2c25c381bb2af7c01633ce1abb022) |
| `manifest_generated_at` | [manifest_generated_at](data-sources--network_regional_edges--reference--group-001.md#canonical-aab5966293dfa9543b6adc7b2112ac268872ac1adf012b932c8c86b6f703d79a) |
| `regions` | [regions](data-sources--network_regional_edges--reference--group-001.md#canonical-127d4ce3ad9cfceb9ea1a082d67601d770cce38e468a54464fb9d47272c5ab34) |
| `source_entries` | [source_entries](data-sources--network_regional_edges--reference--group-001.md#canonical-be98a5d8aa25674f566106ed2a34a5741196fcde3477972f49e70e0863442f33) |
| `source_sha256` | [source_sha256](data-sources--network_regional_edges--reference--group-001.md#canonical-c4d83c3b7c5830cdb0856ba6249468cac8ce0fbb15da070f49a396e34de55831) |
| `source_url` | [source_url](data-sources--network_regional_edges--reference--group-001.md#canonical-e9677830b089cb20a94ccc844a5a9b59f5a10baa309127302be0276a0d0541a9) |

<a id="canonical-58e8f7c8e97d25252b5e3ab56723125b15bc76d0e4a3588f913055d4ab6dd5f4"></a>

## Next pages — Property reference / 28a2240a0595 / 14

- [xcsh_network_regional_edges](../data-sources/network_regional_edges.md#canonical-4b07492253f8f10e61287f15adccadbfc866bfc8be3823bb27bb4e02b5d71089)
