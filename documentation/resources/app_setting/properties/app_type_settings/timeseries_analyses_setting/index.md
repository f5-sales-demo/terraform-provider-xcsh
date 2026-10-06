---
page_title: "app_type_settings.timeseries_analyses_setting"
subcategory: ""
description: "Configuration for DDoS Detection."
xcsh_docs: {"aliases": ["app type settings timeseries analyses setting"], "body_bytes": 1271, "body_sha256": "sha256:b8a2ddb874fab1a5c51aab66fa8e0e948f2218c2a7f8bcf2a89420b60c43277f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting:metric_selectors"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting", "parent_id": "xcsh-docs:resources:app_setting:properties:app_type_settings", "path": "documentation/resources/app_setting/properties/app_type_settings/timeseries_analyses_setting/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2032330103110022-0012233022213111-3321320300011031-3121331201230233-1321332320000100-1210033032002113-3101023132030320-2312022210011020", "registry_path": "docs/guides/resources--app_setting--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "timeseries_analyses_setting"], "schema_version": 1, "sections": [{"aliases": ["app type settings timeseries analyses setting metric selectors"], "anchor": "section", "description": "Define the metric selection criteria, i.e. The metrics source and the actual metrics that should be included in the detection logic.", "document_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting:metric_selectors", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["app_type_settings", "timeseries_analyses_setting", "metric_selectors"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/timeseries_analyses_setting/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration for DDoS Detection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.timeseries_analyses_setting

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/)
- app_type_settings.timeseries_analyses_setting

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
timeseries_analyses_setting {
  # Configure direct properties listed below.
}
```

## Direct properties

- [metric_selectors](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/timeseries_analyses_setting/metric_selectors/): complete subsection reference.
