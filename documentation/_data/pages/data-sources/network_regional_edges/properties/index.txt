---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_regional_edges."
xcsh_docs: {"aliases": [], "body_bytes": 3705, "body_sha256": "sha256:08efc0de2c093b568dce10ae7ecddeac80ed9f910f7a28d67627bc8f9e8babf7", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_regional_edges:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_regional_edges:reference", "parent_id": "xcsh-docs:data-sources:network_regional_edges:fundamentals", "path": "documentation/data-sources/network_regional_edges/properties/index.md", "provider_name": "network_regional_edges", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_regional_edges/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_network_regional_edges.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_regional_edges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/)
- Property reference

## Direct properties

<a id="schema-api_release_tag"></a>

### api_release_tag property

Type: `"string"`. Computed.

Pinned api-specs-enriched release tag compiled into this provider.

<a id="schema-cidr_blocks"></a>

### cidr_blocks property

Type: `["list", "string"]`. Computed.

Sorted unique IPv4 CIDRs across selected regions. Individual IPv4 addresses are normalized to /32.

<a id="schema-cidr_blocks_by_region"></a>

### cidr_blocks_by_region property

Type: `["map", ["list", "string"]]`. Computed.

Sorted unique IPv4 CIDRs keyed by selected published region.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Stable identifier derived from the pinned source digest, data-source group, and selected regions.

<a id="schema-manifest_generated_at"></a>

### manifest_generated_at property

Type: `"string"`. Computed.

Generation timestamp reported by the published network allowlist manifest.

<a id="schema-regions"></a>

### regions property

Type: `["set", "string"]`. Optional, Computed.

Published regions to include. Omit to select every region.

<a id="schema-source_entries"></a>

### source_entries property

Type: `["list", "string"]`. Computed.

Selected IPv4 entries exactly as represented by the published regions, ordered by region name and
source order.

<a id="schema-source_sha256"></a>

### source_sha256 property

Type: `"string"`. Computed.

SHA-256 digest of the canonical published network allowlist manifest.

<a id="schema-source_url"></a>

### source_url property

Type: `"string"`. Computed.

Published F5 network allowlist source URL.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `api_release_tag` | [api_release_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/properties/#schema-api_release_tag) |
| `cidr_blocks` | [cidr_blocks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/properties/#schema-cidr_blocks) |
| `cidr_blocks_by_region` | [cidr_blocks_by_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/properties/#schema-cidr_blocks_by_region) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/properties/#schema-id) |
| `manifest_generated_at` | [manifest_generated_at](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/properties/#schema-manifest_generated_at) |
| `regions` | [regions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/properties/#schema-regions) |
| `source_entries` | [source_entries](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/properties/#schema-source_entries) |
| `source_sha256` | [source_sha256](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/properties/#schema-source_sha256) |
| `source_url` | [source_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/properties/#schema-source_url) |

## Next pages

- [xcsh_network_regional_edges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_regional_edges/)
