---
page_title: "app_type_settings.timeseries_analyses_setting.metric_selectors"
subcategory: ""
description: "Define the metric selection criteria, i.e. The metrics source and the actual metrics that should be included in the detection logic."
xcsh_docs: {"aliases": ["app type settings timeseries analyses setting metric selectors"], "body_bytes": 4411, "body_sha256": "sha256:5183ee872500a7c7529588f44da19e9c80a23121ff3eb745c44a457c40bcb2a9", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting:metric_selectors", "parent_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting", "path": "documentation/resources/app_setting/properties/app_type_settings/timeseries_analyses_setting/metric_selectors/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-1221212232310233-2012132010033212-2101230013021333-1133321010033232-1020012033022212-2103023201122113-0013323211302332-2202301113011123", "registry_path": "docs/guides/resources--app_setting--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "timeseries_analyses_setting", "metric_selectors"], "schema_version": 1, "sections": [{"aliases": ["metric"], "anchor": "schema-app_type_settings--timeseries_analyses_setting--metric_selectors--metric", "description": "Choose one or more metrics to be included in the detection logic.", "document_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting:metric_selectors", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "timeseries_analyses_setting", "metric_selectors", "metric"], "syntax": "attribute", "type": "list"}, {"aliases": ["metrics source"], "anchor": "schema-app_type_settings--timeseries_analyses_setting--metric_selectors--metrics_source", "description": "Supported sources from which Metrics can be analyzed All edges in the service mesh graph. Metrics are analyzed separately between all source and destination service combinations.", "document_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting:metric_selectors", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "timeseries_analyses_setting", "metric_selectors", "metrics_source"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/timeseries_analyses_setting/metric_selectors/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Define the metric selection criteria, i.e. The metrics source and the actual metrics that should be included in the detection logic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_settingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
