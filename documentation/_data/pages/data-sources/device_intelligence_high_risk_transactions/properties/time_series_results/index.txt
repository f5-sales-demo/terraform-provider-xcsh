---
page_title: "time_series_results"
subcategory: ""
description: "Collection of time series grouped by a series key."
xcsh_docs: {"aliases": ["time series results"], "body_bytes": 1494, "body_sha256": "sha256:49f345a53d025b2dc7ef0df8fb8861d5ba543433f8c719e6d1323e6454c0e09c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results:time_series"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results", "parent_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:reference", "path": "documentation/data-sources/device_intelligence_high_risk_transactions/properties/time_series_results/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_high_risk_transactions", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1302300303113012-0021110302030233-3333102210111132-1313202330311232-3220230301100221-3131210133121012-0223331023211203-0301303003330001", "registry_path": "docs/guides/data-sources--device_intelligence_high_risk_transactions--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["time_series_results"], "schema_version": 1, "sections": [{"aliases": ["series key"], "anchor": "schema-time_series_results--series_key", "description": "Identifier for the time series (e.g., 'total', 'high_risk').", "document_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["time_series_results", "series_key"], "syntax": "attribute", "type": "string"}, {"aliases": ["time series"], "anchor": "section", "description": "Sequence of timestamped values for this series.", "document_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:properties:time_series_results:time_series", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["time_series_results", "time_series"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_high_risk_transactions/properties/time_series_results/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Collection of time series grouped by a series key.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [time_series_results.time_series](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/properties/time_series_results/time_series/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/properties/)
- [xcsh_device_intelligence_high_risk_transactions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/)
