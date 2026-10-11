---
page_title: "data.metric.value"
subcategory: ""
description: "Value. List of metric values."
xcsh_docs: {"aliases": ["data metric value"], "body_bytes": 1194, "body_sha256": "sha256:e52b3c110a3f3b961c935a7b1d2d4108bddf40bffbef8ca09da03d0572ce93f4", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value:trend_value"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric", "path": "documentation/data-sources/tmm_session_metrics/properties/data/metric/value/index.md", "product": "distributed-cloud", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-1322312002310232-0123212212222132-1230223123330220-3100022130101101-0331233131210213-2012130033010301-0212223313330003-2132130301202033", "registry_path": "docs/guides/data-sources--tmm_session_metrics--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["data", "metric", "value"], "schema_version": 1, "sections": [{"aliases": ["data metric value timestamp"], "anchor": "schema-data--metric--value--timestamp", "description": "Timestamp. Timestamp", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data", "metric", "value", "timestamp"], "syntax": "attribute", "type": "number"}, {"aliases": ["data metric value trend value"], "anchor": "section", "description": "Trend value contains trend value, trend sentiment and trend calculation description and window size.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value:trend_value", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["data", "metric", "value", "trend_value"], "syntax": "attribute", "type": "object"}, {"aliases": ["data metric value value"], "anchor": "schema-data--metric--value--value", "description": "Value. Configuration parameter for value", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data", "metric", "value", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/properties/data/metric/value/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Value. List of metric values.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# data.metric.value

Breadcrumbs:

- [xcsh_tmm_session_metrics](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/)
- [data](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/)
- [data.metric](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/)
- data.metric.value

<a id="section"></a>

Type: `"list"`. Computed.

Value. List of metric values.

## Direct properties

<a id="schema-data--metric--value--timestamp"></a>

### timestamp property

Type: `"number"`. Computed.

Timestamp. Timestamp

- [trend_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/value/trend_value/): complete subsection reference.

<a id="schema-data--metric--value--value"></a>

### value property

Type: `"string"`. Computed.

Value. Configuration parameter for value
