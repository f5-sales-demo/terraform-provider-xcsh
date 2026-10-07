---
page_title: "time_series_results"
subcategory: ""
description: "Collection of time series grouped by a series key."
xcsh_docs: {"aliases": ["time series results"], "body_bytes": 971, "body_sha256": "sha256:d17ea77446464ea58f719e580209acbc1f0fdbbb4a1ab0d3bc2e25929221ecd8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results:time_series"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results", "parent_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:reference", "path": "documentation/data-sources/device_intelligence_high_risk_transactions/properties/time_series_results/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_high_risk_transactions", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1302300303113012-0021110302030233-3333102210111132-1313202330311232-3220230301100221-3131210133121012-0223331023211203-0301303003330001", "registry_path": "docs/guides/data-sources--device_intelligence_high_risk_transactions--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["time_series_results"], "schema_version": 1, "sections": [{"aliases": ["time series results series key"], "anchor": "schema-time_series_results--series_key", "description": "Identifier for the time series (e.g., 'total', 'high_risk').", "document_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["time_series_results", "series_key"], "syntax": "attribute", "type": "string"}, {"aliases": ["time series results time series"], "anchor": "section", "description": "Sequence of timestamped values for this series.", "document_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results:time_series", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["time_series_results", "time_series"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_high_risk_transactions/properties/time_series_results/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Collection of time series grouped by a series key.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# time_series_results

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/properties/)
- time_series_results

<a id="section"></a>

Type: `"list"`. Computed.

Collection of time series grouped by a series key.

## Direct properties

<a id="schema-time_series_results--series_key"></a>

### series_key property

Type: `"string"`. Computed.

Identifier for the time series (e.g., 'total', 'high\_risk').

- [time_series](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/properties/time_series_results/time_series/): complete subsection reference.
