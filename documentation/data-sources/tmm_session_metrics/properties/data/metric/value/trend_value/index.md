---
page_title: "data.metric.value.trend_value"
subcategory: ""
description: "Trend value contains trend value, trend sentiment and trend calculation description and window size."
xcsh_docs: {"aliases": ["data metric value trend value"], "body_bytes": 2438, "body_sha256": "sha256:2c7a997099bc45b19c036056e28dac9fc26118d4d34e42a18982272fa8e773a9", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value:trend_value", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value", "path": "documentation/data-sources/tmm_session_metrics/properties/data/metric/value/trend_value/index.md", "product": "distributed-cloud", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2211330301100331-3121020210000012-0011133012222123-0320022102331112-0231231002003110-2133132122212032-1330033333111022-3332212122000021", "registry_path": "docs/guides/data-sources--tmm_session_metrics--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["data", "metric", "value", "trend_value"], "schema_version": 1, "sections": [{"aliases": ["description spec"], "anchor": "schema-data--metric--value--trend_value--description_spec", "description": "Description of the method used to calculate trend.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value:trend_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data", "metric", "value", "trend_value", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["previous value"], "anchor": "schema-data--metric--value--trend_value--previous_value", "description": "Configuration parameter for previous value.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value:trend_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data", "metric", "value", "trend_value", "previous_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["sentiment"], "anchor": "schema-data--metric--value--trend_value--sentiment", "description": "Trend sentiment Indicates trend sentiment is positive Indicates trend sentiment is negative. Possible values are `TREND_SENTIMENT_NONE`, `TREND_SENTIMENT_POSITIVE`, `TREND_SENTIMENT_NEGATIVE`. Defaults to `TREND_SENTIMENT_NONE`.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value:trend_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data", "metric", "value", "trend_value", "sentiment"], "syntax": "attribute", "type": "string"}, {"aliases": ["value"], "anchor": "schema-data--metric--value--trend_value--value", "description": "Value. Configuration parameter for value", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value:trend_value", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data", "metric", "value", "trend_value", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/properties/data/metric/value/trend_value/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Trend value contains trend value, trend sentiment and trend calculation description and window size.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# data.metric.value.trend_value

Breadcrumbs:

- [xcsh_tmm_session_metrics](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/)
- [data](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/)
- [data.metric](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/)
- [data.metric.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/value/)
- data.metric.value.trend_value

<a id="section"></a>

Type: `"single"`. Computed.

Trend value contains trend value, trend sentiment and trend calculation description and window size.

## Direct properties

<a id="schema-data--metric--value--trend_value--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description of the method used to calculate trend.

<a id="schema-data--metric--value--trend_value--previous_value"></a>

### previous_value property

Type: `"string"`. Computed.

Configuration parameter for previous value.

<a id="schema-data--metric--value--trend_value--sentiment"></a>

### sentiment property

Type: `"string"`. Computed.

\[Enum: TREND\_SENTIMENT\_NONE|TREND\_SENTIMENT\_POSITIVE|TREND\_SENTIMENT\_NEGATIVE\] Trend
sentiment Indicates trend sentiment is positive Indicates trend sentiment is negative. Possible
values are \`TREND\_SENTIMENT\_NONE\`, \`TREND\_SENTIMENT\_POSITIVE\`,
\`TREND\_SENTIMENT\_NEGATIVE\`. Defaults to \`TREND\_SENTIMENT\_NONE\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TREND_SENTIMENT_NONE",
    "TREND_SENTIMENT_POSITIVE",
    "TREND_SENTIMENT_NEGATIVE"),
}
```

<a id="schema-data--metric--value--trend_value--value"></a>

### value property

Type: `"string"`. Computed.

Value. Configuration parameter for value

## Next pages

- [data.metric.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/value/)
- [xcsh_tmm_session_metrics](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/)
