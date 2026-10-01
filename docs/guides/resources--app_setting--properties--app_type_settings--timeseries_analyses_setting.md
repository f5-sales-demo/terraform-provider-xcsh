---
page_title: "app_type_settings.timeseries_analyses_setting"
subcategory: ""
description: "app_type_settings.timeseries_analyses_setting for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 1378, "body_sha256": "sha256:7c3fe999248e13af8d970de0fe6feeceba27e2e437d8e3a4a1928bbb84ca0558", "canonical_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting", "child_ids": ["xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting:metric_selectors"], "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting", "parent_id": "xcsh-docs:resources:app_setting:properties:app_type_settings", "path": "docs/guides/resources--app_setting--properties--app_type_settings--timeseries_analyses_setting.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["app_type_settings", "timeseries_analyses_setting"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/timeseries_analyses_setting/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "app_type_settings.timeseries_analyses_setting for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.timeseries_analyses_setting

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md)
- [Property reference](resources--app_setting--reference.md)
- [app_type_settings](resources--app_setting--properties--app_type_settings.md)
- app_type_settings.timeseries_analyses_setting

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
timeseries_analyses_setting {
  # Configure direct properties listed below.
}
```

## Direct properties

- [metric_selectors](resources--app_setting--properties--app_type_settings--timeseries_analyses_setting--metric_selectors.md): complete subsection reference.

## Next pages

- [app_type_settings.timeseries_analyses_setting.metric_selectors](resources--app_setting--properties--app_type_settings--timeseries_analyses_setting--metric_selectors.md)
- [app_type_settings](resources--app_setting--properties--app_type_settings.md)
- [xcsh_app_setting](../resources/app_setting.md)
