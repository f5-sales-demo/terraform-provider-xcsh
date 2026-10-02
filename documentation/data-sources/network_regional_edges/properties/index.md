---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_regional_edges."
xcsh_docs: {"aliases": ["network regional edges"], "body_bytes": 3705, "body_sha256": "sha256:08efc0de2c093b568dce10ae7ecddeac80ed9f910f7a28d67627bc8f9e8babf7", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_regional_edges:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_regional_edges:reference", "parent_id": "xcsh-docs:data-sources:network_regional_edges:fundamentals", "path": "documentation/data-sources/network_regional_edges/properties/index.md", "product": "distributed-cloud", "provider_name": "network_regional_edges", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3200132122123211-1311200021302313-1330223222012333-2301330001203121-2032303231120232-2032111121003012-2020302202111221-0022103200132122", "registry_path": "docs/guides/data-sources--network_regional_edges--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["api release tag"], "anchor": "schema-api_release_tag", "description": "Pinned api-specs-enriched release tag compiled into this provider.", "document_id": "xcsh-docs:data-sources:network_regional_edges:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_release_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["cidr blocks"], "anchor": "schema-cidr_blocks", "description": "Sorted unique IPv4 CIDRs across selected regions. Individual IPv4 addresses are normalized to /32.", "document_id": "xcsh-docs:data-sources:network_regional_edges:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cidr_blocks"], "syntax": "attribute", "type": "list"}, {"aliases": ["cidr blocks by region"], "anchor": "schema-cidr_blocks_by_region", "description": "Sorted unique IPv4 CIDRs keyed by selected published region.", "document_id": "xcsh-docs:data-sources:network_regional_edges:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cidr_blocks_by_region"], "syntax": "attribute", "type": "map"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Stable identifier derived from the pinned source digest, data-source group, and selected regions.", "document_id": "xcsh-docs:data-sources:network_regional_edges:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["manifest generated at"], "anchor": "schema-manifest_generated_at", "description": "Generation timestamp reported by the published network allowlist manifest.", "document_id": "xcsh-docs:data-sources:network_regional_edges:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["manifest_generated_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["regions"], "anchor": "schema-regions", "description": "Published regions to include. Omit to select every region.", "document_id": "xcsh-docs:data-sources:network_regional_edges:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["regions"], "syntax": "attribute", "type": "set"}, {"aliases": ["source entries"], "anchor": "schema-source_entries", "description": "Selected IPv4 entries exactly as represented by the published regions, ordered by region name and source order.", "document_id": "xcsh-docs:data-sources:network_regional_edges:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["source_entries"], "syntax": "attribute", "type": "list"}, {"aliases": ["source sha256"], "anchor": "schema-source_sha256", "description": "SHA-256 digest of the canonical published network allowlist manifest.", "document_id": "xcsh-docs:data-sources:network_regional_edges:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["source_sha256"], "syntax": "attribute", "type": "string"}, {"aliases": ["source url"], "anchor": "schema-source_url", "description": "Published F5 network allowlist source URL.", "document_id": "xcsh-docs:data-sources:network_regional_edges:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["source_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_regional_edges/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_network_regional_edges.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
