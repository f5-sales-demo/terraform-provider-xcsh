---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_tmm_session_metrics."
xcsh_docs: {"aliases": [], "body_bytes": 8635, "body_sha256": "sha256:4eaaa83a4a2f48d5e9720a40174cc075796dbdb984c3bd689d4c06d4f41dbc52", "child_ids": ["xcsh-docs:data-sources:tmm_session_metrics:properties:data", "xcsh-docs:data-sources:tmm_session_metrics:properties:label_filter"], "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:reference", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:fundamentals", "path": "documentation/data-sources/tmm_session_metrics/properties/index.md", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_tmm_session_metrics.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_tmm_session_metrics](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/)
- Property reference

## Direct properties

- [data](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/): complete subsection reference.

<a id="schema-end_time"></a>

### end_time property

Type: `"string"`. Optional.

End time of metric collection from which data will be considered to build dashboard. Format:
unix\_timestamp|RFC 3339 Optional: If not specified, then the end\_time will be evaluated to
start\_time+10m If start\_time is not specified, then the end\_time will be evaluated to &lt;current
time&gt;.

<a id="schema-field_selector"></a>

### field_selector property

Type: `["list", "string"]`. Optional.

\[Enum:
METRIC\_TYPE\_NONE|METRIC\_TYPE\_ACTIVE|METRIC\_TYPE\_ALLOWED|METRIC\_TYPE\_DENIED|METRIC\_TYPE\_TOTAL|METRIC\_TYPE\_LOGOUT|METRIC\_TYPE\_ESTABLISHED\_TIMEOUT|METRIC\_TYPE\_EVALUATION\_TIMEOUT|METRIC\_TYPE\_ADMIN\_TERMINATED\]
Select fields to be returned in the response. Field\_selector is used to specify the fields to be
returned in the response, thereby limiting the amount of data returned in the response. Possible
values are \`METRIC\_TYPE\_NONE\`, \`METRIC\_TYPE\_ACTIVE\`, \`METRIC\_TYPE\_ALLOWED\`,
\`METRIC\_TYPE\_DENIED\`, \`METRIC\_TYPE\_TOTAL\`, \`METRIC\_TYPE\_LOGOUT\`,
\`METRIC\_TYPE\_ESTABLISHED\_TIMEOUT\`, \`METRIC\_TYPE\_EVALUATION\_TIMEOUT\`,
\`METRIC\_TYPE\_ADMIN\_TERMINATED\`. Defaults to \`METRIC\_TYPE\_NONE\`.

<a id="schema-group_by"></a>

### group_by property

Type: `["list", "string"]`. Optional.

\[Enum: METRIC\_LABEL\_NONE|METRIC\_LABEL\_NAMESPACE|METRIC\_LABEL\_VIRTUAL\_SERVER\] Aggregate data
by any or all of namespace, site and virtual\_server. Optional: If not specified, then the data is
aggregated/grouped by namespace and service. Possible values are \`METRIC\_LABEL\_NONE\`,
\`METRIC\_LABEL\_NAMESPACE\`, \`METRIC\_LABEL\_VIRTUAL\_SERVER\`. Defaults to
\`METRIC\_LABEL\_NONE\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

<a id="schema-is_trend_request"></a>

### is_trend_request property

Type: `"bool"`. Optional.

Trend value computation requested by the user Optional:. Defaults to \`false\`.

- [label_filter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/label_filter/): complete subsection reference.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace namespace is used to scope session metrics. Only virtual server in given namespace will be
considered.

<a id="schema-start_time"></a>

### start_time property

Type: `"string"`. Optional.

Start time of metric collection from which data will be considered to build metrics dashboard.
Format: unix\_timestamp|RFC 3339 Optional: If not specified, then the start\_time will be evaluated
to end\_time-10m If end\_time is not specified, then the start\_time will be evaluated to
&lt;current time&gt;-10m.

<a id="schema-step"></a>

### step property

Type: `"string"`. Optional.

Step is the resolution width, which determines the number of the data points \[x-axis (time)\] to be
returned in the response. The timestamps in the response will be t1=start\_time, t2=t1+step, ..
Tn=tn-1+step, where tn &lt;= end\_time.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `data` | [data](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/#section) |
| `data.metric` | [data.metric](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/#section) |
| `data.metric.key` | [data.metric.key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/key/#section) |
| `data.metric.value` | [data.metric.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/value/#section) |
| `data.metric.value.timestamp` | [data.metric.value.timestamp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/value/#schema-data--metric--value--timestamp) |
| `data.metric.value.trend_value` | [data.metric.value.trend_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/value/trend_value/#section) |
| `data.metric.value.trend_value.description_spec` | [data.metric.value.trend_value.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/value/trend_value/#schema-data--metric--value--trend_value--description_spec) |
| `data.metric.value.trend_value.previous_value` | [data.metric.value.trend_value.previous_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/value/trend_value/#schema-data--metric--value--trend_value--previous_value) |
| `data.metric.value.trend_value.sentiment` | [data.metric.value.trend_value.sentiment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/value/trend_value/#schema-data--metric--value--trend_value--sentiment) |
| `data.metric.value.trend_value.value` | [data.metric.value.trend_value.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/value/trend_value/#schema-data--metric--value--trend_value--value) |
| `data.metric.value.value` | [data.metric.value.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/metric/value/#schema-data--metric--value--value) |
| `data.type` | [data.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/#schema-data--type) |
| `data.unit` | [data.unit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/#schema-data--unit) |
| `end_time` | [end_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/#schema-end_time) |
| `field_selector` | [field_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/#schema-field_selector) |
| `group_by` | [group_by](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/#schema-group_by) |
| `is_trend_request` | [is_trend_request](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/#schema-is_trend_request) |
| `label_filter` | [label_filter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/label_filter/#section) |
| `label_filter.label` | [label_filter.label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/label_filter/#schema-label_filter--label) |
| `label_filter.op` | [label_filter.op](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/label_filter/#schema-label_filter--op) |
| `label_filter.value` | [label_filter.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/label_filter/#schema-label_filter--value) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/#schema-namespace) |
| `start_time` | [start_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/#schema-start_time) |
| `step` | [step](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/#schema-step) |

## Next pages

- [data](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/data/)
- [label_filter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/label_filter/)
- [xcsh_tmm_session_metrics](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/)
