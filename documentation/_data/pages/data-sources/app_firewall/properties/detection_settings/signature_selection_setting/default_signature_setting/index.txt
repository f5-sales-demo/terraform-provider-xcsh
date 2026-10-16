---
page_title: "detection_settings.signature_selection_setting.default_signature_setting"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["detection settings signature selection setting default signature setting"], "body_bytes": 1267, "body_sha256": "sha256:cb397ed7fbebe643139d8394b7cd6b367ed957c8c2447ea44911fa2c09021a22", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:default_signature_setting", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting", "path": "documentation/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/default_signature_setting/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-0230113311030022-1101331313222002-1013302130232132-3032103232012100-1213100102002001-2111200121002020-3201301133203202-1103130032102301", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["detection_settings", "signature_selection_setting", "default_signature_setting"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/default_signature_setting/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["app_firewallCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.signature_selection_setting.default_signature_setting

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/)
- [detection_settings.signature_selection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/)
- detection_settings.signature_selection_setting.default_signature_setting

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default signature setting.

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

This is an empty object or choice marker. It has no direct properties.
