---
page_title: "app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["app type settings user behavior analysis setting enable detection include waf activity"], "body_bytes": 1565, "body_sha256": "sha256:7f69e985425189942931ac527df8e2685752d89f6c2c70e7b8eba4a21257f723", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_waf_activity", "parent_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection", "path": "documentation/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_waf_activity/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2321311012200003-3323332300100232-0110223212200132-0220232021222023-3133212033112223-3102230303211101-0031202030213333-1213200213303111", "registry_path": "docs/guides/resources--app_setting--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_waf_activity"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_waf_activity/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["app_settingCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/)
- [app_type_settings.user_behavior_analysis_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for include waf activity.

Additional upstream details:

This can be used for messages where no values are needed.

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
include_waf_activity = {}
```

This is an empty object or choice marker. It has no direct properties.
