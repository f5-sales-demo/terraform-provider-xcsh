---
page_title: "data.metric.value"
subcategory: ""
description: "Value. List of metric values."
xcsh_docs: {"aliases": ["data metric value"], "body_bytes": 1626, "body_sha256": "sha256:050d3bb48c140c612e33dac7389e7bdfa37b6bad75b9859240845fbdcff14690", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value:trend_value"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric", "path": "documentation/data-sources/tmm_session_metrics/properties/data/metric/value/index.md", "product": "distributed-cloud", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1322312002310232-0123212212222132-1230223123330220-3100022130101101-0331233131210213-2012130033010301-0212223313330003-2132130301202033", "registry_path": "docs/guides/data-sources--tmm_session_metrics--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["data", "metric", "value"], "schema_version": 1, "sections": [{"aliases": ["timestamp"], "anchor": "schema-data--metric--value--timestamp", "description": "Timestamp. Timestamp", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data", "metric", "value", "timestamp"], "syntax": "attribute", "type": "number"}, {"aliases": ["trend value"], "anchor": "section", "description": "Trend value contains trend value, trend sentiment and trend calculation description and window size.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value:trend_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["data", "metric", "value", "trend_value"], "syntax": "attribute", "type": "object"}, {"aliases": ["value"], "anchor": "schema-data--metric--value--value", "description": "Value. Configuration parameter for value", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data", "metric", "value", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/properties/data/metric/value/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Value. List of metric values.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [data.metric.value.trend_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/value/trend_value/)
- [data.metric](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/)
- [xcsh_tmm_session_metrics](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/)
