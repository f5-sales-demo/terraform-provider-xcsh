---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_tmm_session_metrics."
xcsh_docs: {"aliases": ["tmm session metrics"], "body_bytes": 8635, "body_sha256": "sha256:4eaaa83a4a2f48d5e9720a40174cc075796dbdb984c3bd689d4c06d4f41dbc52", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:tmm_session_metrics:properties:data", "xcsh-docs:data-sources:tmm_session_metrics:properties:label_filter"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:reference", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:fundamentals", "path": "documentation/data-sources/tmm_session_metrics/properties/index.md", "product": "distributed-cloud", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1102133100011312-3223013330330031-0122231013233021-3302021011031130-2012012221022001-0001110032023221-2202110303032301-2312300333322010", "registry_path": "docs/guides/data-sources--tmm_session_metrics--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["data"], "anchor": "section", "description": "Data contains time-series TMM Session data.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:data", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["data"], "syntax": "attribute", "type": "object"}, {"aliases": ["end time"], "anchor": "schema-end_time", "description": "End time of metric collection from which data will be considered to build dashboard. Format: unix_timestamp|RFC 3339 Optional: If not specified, then the end_time will be evaluated to start_time+10m If start_time is not specified, then the end_time will be evaluated to <current time>.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["end_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["field selector"], "anchor": "schema-field_selector", "description": "Select fields to be returned in the response. Field_selector is used to specify the fields to be returned in the response, thereby limiting the amount of data returned in the response. Possible values are `METRIC_TYPE_NONE`, `METRIC_TYPE_ACTIVE`, `METRIC_TYPE_ALLOWED`, `METRIC_TYPE_DENIED`, `METRIC_TYPE_TOTAL`,", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["field_selector"], "syntax": "attribute", "type": "list"}, {"aliases": ["group by"], "anchor": "schema-group_by", "description": "Aggregate data by any or all of namespace, site and virtual_server. Optional: If not specified, then the data is aggregated/grouped by namespace and service. Possible values are `METRIC_LABEL_NONE`, `METRIC_LABEL_NAMESPACE`, `METRIC_LABEL_VIRTUAL_SERVER`. Defaults to `METRIC_LABEL_NONE`.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["group_by"], "syntax": "attribute", "type": "list"}, {"aliases": ["is trend request"], "anchor": "schema-is_trend_request", "description": "Trend value computation requested by the user Optional:. Defaults to `false`.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["is_trend_request"], "syntax": "attribute", "type": "bool"}, {"aliases": ["label filter"], "anchor": "section", "description": "List of label filter expressions of the form 'label key' QueryOp 'value'. Response will only contain data that matches all the conditions specified in the label_filter.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:properties:label_filter", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["label_filter"], "syntax": "attribute", "type": "object"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace namespace is used to scope session metrics. Only virtual server in given namespace will be considered.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["start time"], "anchor": "schema-start_time", "description": "Start time of metric collection from which data will be considered to build metrics dashboard. Format: unix_timestamp|RFC 3339 Optional: If not specified, then the start_time will be evaluated to end_time-10m If end_time is not specified, then the start_time will be evaluated to <current time>-10m.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["start_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["step"], "anchor": "schema-step", "description": "Step is the resolution width, which determines the number of the data points to be returned in the response. The timestamps in the response will be t1=start_time, t2=t1+step, .. Tn=tn-1+step, where tn <= end_time.", "document_id": "xcsh-docs:data-sources:tmm_session_metrics:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["step"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_tmm_session_metrics.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
