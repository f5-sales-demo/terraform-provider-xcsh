---
page_title: "data.metric"
subcategory: ""
description: "Metric. List of metrics."
xcsh_docs: {"aliases": ["data metric"], "body_bytes": 922, "body_sha256": "sha256:012bbea055f683d5d04b7c8fe4ccc4ade2f8255cdd29ffae4590440abd7a624d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:key", "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data", "path": "documentation/data-sources/tmm_session_metrics/properties/data/metric/index.md", "product": "distributed-cloud", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2022123030212230-3223133203200121-3123211322331000-1023310331223000-2222101311323333-2330200031300120-3232223013121033-1301033330333100", "registry_path": "docs/guides/data-sources--tmm_session_metrics--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["data", "metric"], "schema_version": 1, "sections": [{"aliases": ["data metric key"], "anchor": "section", "description": "Key contains the name/value pair. 'name' is the label name defined in 'MetricLabel'.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:key", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["data", "metric", "key"], "syntax": "attribute", "type": "object"}, {"aliases": ["data metric value"], "anchor": "section", "description": "Value. List of metric values.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["data", "metric", "value"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/properties/data/metric/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Metric. List of metrics.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# data.metric

Breadcrumbs:

- [xcsh_tmm_session_metrics](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/)
- [data](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/)
- data.metric

<a id="section"></a>

Type: `"list"`. Computed.

Metric. List of metrics.

## Direct properties

- [key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/key/): complete subsection reference.

- [value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/value/): complete subsection reference.
