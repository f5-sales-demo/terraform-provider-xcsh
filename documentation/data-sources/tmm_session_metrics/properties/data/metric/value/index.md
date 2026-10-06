---
page_title: "data.metric.value"
subcategory: ""
description: "Value. List of metric values."
xcsh_docs: {"aliases": ["data metric value"], "body_bytes": 1194, "body_sha256": "sha256:e52b3c110a3f3b961c935a7b1d2d4108bddf40bffbef8ca09da03d0572ce93f4", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value:trend_value"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric", "path": "documentation/data-sources/tmm_session_metrics/properties/data/metric/value/index.md", "product": "distributed-cloud", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1322312002310232-0123212212222132-1230223123330220-3100022130101101-0331233131210213-2012130033010301-0212223313330003-2132130301202033", "registry_path": "docs/guides/data-sources--tmm_session_metrics--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["data", "metric", "value"], "schema_version": 1, "sections": [{"aliases": ["data metric value timestamp"], "anchor": "schema-data--metric--value--timestamp", "description": "Timestamp. Timestamp", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data", "metric", "value", "timestamp"], "syntax": "attribute", "type": "number"}, {"aliases": ["data metric value trend value"], "anchor": "section", "description": "Trend value contains trend value, trend sentiment and trend calculation description and window size.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value:trend_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["data", "metric", "value", "trend_value"], "syntax": "attribute", "type": "object"}, {"aliases": ["data metric value value"], "anchor": "schema-data--metric--value--value", "description": "Value. Configuration parameter for value", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data", "metric", "value", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/properties/data/metric/value/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Value. List of metric values.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
