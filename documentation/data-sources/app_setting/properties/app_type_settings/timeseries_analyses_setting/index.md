---
page_title: "app_type_settings.timeseries_analyses_setting"
subcategory: ""
description: "Configuration for DDoS Detection."
xcsh_docs: {"aliases": ["app type settings timeseries analyses setting"], "body_bytes": 1618, "body_sha256": "sha256:e9b8f085fa05f709a94f3574beacdeeb08d28f82f08ad96622062b925456af04", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:app_setting:properties:app_type_settings:timeseries_analyses_setting:metric_selectors"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:timeseries_analyses_setting", "parent_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings", "path": "documentation/data-sources/app_setting/properties/app_type_settings/timeseries_analyses_setting/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1322231230201010-1123332213003323-1200010112002212-1012313313333213-2203333030120312-3330311313101321-3011123003302232-1021032230212323", "registry_path": "docs/guides/data-sources--app_setting--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "timeseries_analyses_setting"], "schema_version": 1, "sections": [{"aliases": ["app type settings timeseries analyses setting metric selectors"], "anchor": "section", "description": "Define the metric selection criteria, i.e. The metrics source and the actual metrics that should be included in the detection logic.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:timeseries_analyses_setting:metric_selectors", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["app_type_settings", "timeseries_analyses_setting", "metric_selectors"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/properties/app_type_settings/timeseries_analyses_setting/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration for DDoS Detection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.timeseries_analyses_setting

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/)
- app_type_settings.timeseries_analyses_setting

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for timeseries analyses setting.

Upstream description:

Configuration for DDoS Detection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [metric_selectors](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/timeseries_analyses_setting/metric_selectors/): complete subsection reference.

## Next pages

- [app_type_settings.timeseries_analyses_setting.metric_selectors](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/timeseries_analyses_setting/metric_selectors/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/)
- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/)
