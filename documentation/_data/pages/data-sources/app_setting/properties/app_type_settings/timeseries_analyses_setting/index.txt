---
page_title: "app_type_settings.timeseries_analyses_setting"
subcategory: ""
description: "Configuration for DDoS Detection."
xcsh_docs: {"aliases": ["app type settings timeseries analyses setting"], "body_bytes": 1147, "body_sha256": "sha256:9e3d6d61aaafc414d98d1a4bb147758da0c6cabd638762f4e53e62b4eb90ccac", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:app_setting:properties:app_type_settings:timeseries_analyses_setting:metric_selectors"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:timeseries_analyses_setting", "parent_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings", "path": "documentation/data-sources/app_setting/properties/app_type_settings/timeseries_analyses_setting/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1322231230201010-1123332213003323-1200010112002212-1012313313333213-2203333030120312-3330311313101321-3011123003302232-1021032230212323", "registry_path": "docs/guides/data-sources--app_setting--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "timeseries_analyses_setting"], "schema_version": 1, "sections": [{"aliases": ["app type settings timeseries analyses setting metric selectors"], "anchor": "section", "description": "Define the metric selection criteria, i.e. The metrics source and the actual metrics that should be included in the detection logic.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:timeseries_analyses_setting:metric_selectors", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["app_type_settings", "timeseries_analyses_setting", "metric_selectors"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/properties/app_type_settings/timeseries_analyses_setting/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Configuration for DDoS Detection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["app_settingCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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

Additional upstream details:

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
