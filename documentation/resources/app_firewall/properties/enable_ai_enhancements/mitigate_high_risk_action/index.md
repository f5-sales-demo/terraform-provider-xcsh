---
page_title: "enable_ai_enhancements.mitigate_high_risk_action"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["enable ai enhancements mitigate high risk action"], "body_bytes": 1057, "body_sha256": "sha256:a1e27e64ab04cdefe2c00c8a350ad9415efb1721d19bee6f735d400a41061dae", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:enable_ai_enhancements:mitigate_high_risk_action", "parent_id": "xcsh-docs:resources:app_firewall:properties:enable_ai_enhancements", "path": "documentation/resources/app_firewall/properties/enable_ai_enhancements/mitigate_high_risk_action/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3223233223310233-2330130333121002-1323103311003101-0023011211000332-3300332301001200-2112200302113111-1132330020312310-0113123311023221", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_ai_enhancements", "mitigate_high_risk_action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/enable_ai_enhancements/mitigate_high_risk_action/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["app_firewallCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_ai_enhancements.mitigate_high_risk_action

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [enable_ai_enhancements](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/enable_ai_enhancements/)
- enable_ai_enhancements.mitigate_high_risk_action

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
mitigate_high_risk_action = {}
```

This is an empty object or choice marker. It has no direct properties.
