---
page_title: "data.metric.value.trend_value"
subcategory: ""
description: "data.metric.value.trend_value for xcsh_tmm_session_metrics."
xcsh_docs: {"aliases": [], "body_bytes": 2084, "body_sha256": "sha256:1d1d1078c95362bb8e0f8c871a31816b7690998236e2e6e593efcbf49791c340", "canonical_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value:trend_value", "child_ids": [], "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value:trend_value", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data:metric:value", "path": "docs/guides/data-sources--tmm_session_metrics--properties--data--metric--value--trend_value.md", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["data", "metric", "value", "trend_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/properties/data/metric/value/trend_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "data.metric.value.trend_value for xcsh_tmm_session_metrics.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# data.metric.value.trend_value

Breadcrumbs:

- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md)
- [Property reference](data-sources--tmm_session_metrics--reference.md)
- [data](data-sources--tmm_session_metrics--properties--data.md)
- [data.metric](data-sources--tmm_session_metrics--properties--data--metric.md)
- [data.metric.value](data-sources--tmm_session_metrics--properties--data--metric--value.md)
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

- [data.metric.value](data-sources--tmm_session_metrics--properties--data--metric--value.md)
- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md)
