---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_network_data_intelligence."
xcsh_docs: {"aliases": ["network data intelligence"], "body_bytes": 3596, "body_sha256": "sha256:dc01844427a8a22c70a2f8d1f3c5d5ecf0302c10bf8f614509ca1dc36272c495", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_data_intelligence:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_data_intelligence:reference", "parent_id": "xcsh-docs:data-sources:network_data_intelligence:fundamentals", "path": "documentation/data-sources/network_data_intelligence/properties/index.md", "product": "distributed-cloud", "provider_name": "network_data_intelligence", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3012112022123323-3332203123233211-3331023122230032-2300033323333020-2001232011111103-0002122020120221-0122010012220211-0130201010131103", "registry_path": "docs/guides/data-sources--network_data_intelligence--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["api release tag"], "anchor": "schema-api_release_tag", "description": "Pinned api-specs-enriched release tag compiled into this provider.", "document_id": "xcsh-docs:data-sources:network_data_intelligence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_release_tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["cidr blocks"], "anchor": "schema-cidr_blocks", "description": "Sorted unique IPv4 CIDRs across selected regions. Individual IPv4 addresses are normalized to /32.", "document_id": "xcsh-docs:data-sources:network_data_intelligence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cidr_blocks"], "syntax": "attribute", "type": "list"}, {"aliases": ["cidr blocks by region"], "anchor": "schema-cidr_blocks_by_region", "description": "Sorted unique IPv4 CIDRs keyed by selected published region.", "document_id": "xcsh-docs:data-sources:network_data_intelligence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cidr_blocks_by_region"], "syntax": "attribute", "type": "map"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Stable identifier derived from the pinned source digest, data-source group, and selected regions.", "document_id": "xcsh-docs:data-sources:network_data_intelligence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["manifest generated at"], "anchor": "schema-manifest_generated_at", "description": "Generation timestamp reported by the published network allowlist manifest.", "document_id": "xcsh-docs:data-sources:network_data_intelligence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["manifest_generated_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["regions"], "anchor": "schema-regions", "description": "Published regions to include. Omit to select every region.", "document_id": "xcsh-docs:data-sources:network_data_intelligence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["regions"], "syntax": "attribute", "type": "set"}, {"aliases": ["source entries"], "anchor": "schema-source_entries", "description": "Selected IPv4 entries exactly as represented by the published regions, ordered by region name and source order.", "document_id": "xcsh-docs:data-sources:network_data_intelligence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["source_entries"], "syntax": "attribute", "type": "list"}, {"aliases": ["source sha256"], "anchor": "schema-source_sha256", "description": "SHA-256 digest of the canonical published network allowlist manifest.", "document_id": "xcsh-docs:data-sources:network_data_intelligence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["source_sha256"], "syntax": "attribute", "type": "string"}, {"aliases": ["source url"], "anchor": "schema-source_url", "description": "Published F5 network allowlist source URL.", "document_id": "xcsh-docs:data-sources:network_data_intelligence:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["source_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_data_intelligence/properties/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Property reference for xcsh_network_data_intelligence.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_network_data_intelligence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/)
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
| `api_release_tag` | [api_release_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/properties/#schema-api_release_tag) |
| `cidr_blocks` | [cidr_blocks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/properties/#schema-cidr_blocks) |
| `cidr_blocks_by_region` | [cidr_blocks_by_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/properties/#schema-cidr_blocks_by_region) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/properties/#schema-id) |
| `manifest_generated_at` | [manifest_generated_at](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/properties/#schema-manifest_generated_at) |
| `regions` | [regions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/properties/#schema-regions) |
| `source_entries` | [source_entries](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/properties/#schema-source_entries) |
| `source_sha256` | [source_sha256](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/properties/#schema-source_sha256) |
| `source_url` | [source_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/properties/#schema-source_url) |
