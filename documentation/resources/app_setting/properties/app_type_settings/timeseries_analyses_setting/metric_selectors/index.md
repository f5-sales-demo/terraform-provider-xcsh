---
page_title: "app_type_settings.timeseries_analyses_setting.metric_selectors"
subcategory: ""
description: "Define the metric selection criteria, i.e. The metrics source and the actual metrics that should be included in the detection logic."
xcsh_docs: {"aliases": ["app type settings timeseries analyses setting metric selectors"], "body_bytes": 4411, "body_sha256": "sha256:29308a4b01ac47df4c158e92267c9808dd41270229bb8303c836f71d6b26f707", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting:metric_selectors", "parent_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting", "path": "documentation/resources/app_setting/properties/app_type_settings/timeseries_analyses_setting/metric_selectors/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1221212232310233-2012132010033212-2101230013021333-1133321010033232-1020012033022212-2103023201122113-0013323211302332-2202301113011123", "registry_path": "docs/guides/resources--app_setting--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "timeseries_analyses_setting", "metric_selectors"], "schema_version": 1, "sections": [{"aliases": ["app type settings timeseries analyses setting metric selectors metric"], "anchor": "schema-app_type_settings--timeseries_analyses_setting--metric_selectors--metric", "description": "Choose one or more metrics to be included in the detection logic.", "document_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting:metric_selectors", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "timeseries_analyses_setting", "metric_selectors", "metric"], "syntax": "attribute", "type": "list"}, {"aliases": ["app type settings timeseries analyses setting metric selectors metrics source"], "anchor": "schema-app_type_settings--timeseries_analyses_setting--metric_selectors--metrics_source", "description": "Supported sources from which Metrics can be analyzed All edges in the service mesh graph. Metrics are analyzed separately between all source and destination service combinations.", "document_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting:metric_selectors", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "timeseries_analyses_setting", "metric_selectors", "metrics_source"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/timeseries_analyses_setting/metric_selectors/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Define the metric selection criteria, i.e. The metrics source and the actual metrics that should be included in the detection logic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.timeseries_analyses_setting.metric_selectors

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/)
- [app_type_settings.timeseries_analyses_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/timeseries_analyses_setting/)
- app_type_settings.timeseries_analyses_setting.metric_selectors

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Define the metric selection criteria, i.e. The metrics source and the actual metrics that should be
included in the detection logic.

Upstream description:

Define the metric selection criteria, i.e. The metrics source and the actual metrics that should be
included in the detection logic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
metric_selectors {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-app_type_settings--timeseries_analyses_setting--metric_selectors--metric"></a>

### metric property

Type: `["list", "string"]`. Optional.

\[Enum: NO\_METRICS|REQUEST\_RATE|ERROR\_RATE|LATENCY|THROUGHPUT\] Choose one or more metrics to be
included in the detection logic. Possible values are \`NO\_METRICS\`, \`REQUEST\_RATE\`,
\`ERROR\_RATE\`, \`LATENCY\`, \`THROUGHPUT\`. Defaults to \`NO\_METRICS\`.

Upstream description:

Choose one or more metrics to be included in the detection logic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-app_type_settings--timeseries_analyses_setting--metric_selectors--metrics_source"></a>

### metrics_source property

Type: `"string"`. Optional.

\[Enum: NONE|NODES|EDGES|VIRTUAL\_HOSTS\] Supported sources from which Metrics can be analyzed All
edges in the service mesh graph. Metrics are analyzed separately between all source and destination
service combinations. Possible values are \`NONE\`, \`NODES\`, \`EDGES\`, \`VIRTUAL\_HOSTS\`.

Upstream description:

Supported sources from which Metrics can be analyzed

All edges in the service mesh graph. Metrics are analyzed separately between all source and
destination service combinations.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NONE",
    "NODES",
    "EDGES",
    "VIRTUAL_HOSTS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NONE",
  "enum": [
    "NONE",
    "NODES",
    "EDGES",
    "VIRTUAL_HOSTS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [app_type_settings.timeseries_analyses_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/timeseries_analyses_setting/)
- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
