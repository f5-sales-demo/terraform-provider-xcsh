---
page_title: "enable_ai_enhancements"
subcategory: "Security"
description: "Actions complimented by the additional intelligence of the F5 AI Powered Risk-based analysis."
xcsh_docs: {"aliases": ["enable ai enhancements"], "body_bytes": 1279, "body_sha256": "sha256:c2274260aa9d27ee906d162e95b25645cec520959141c400a1cb836ed411c76d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:app_firewall:properties:enable_ai_enhancements:mitigate_high_medium_risk_action", "xcsh-docs:data-sources:app_firewall:properties:enable_ai_enhancements:mitigate_high_risk_action"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:enable_ai_enhancements", "parent_id": "xcsh-docs:data-sources:app_firewall:reference", "path": "documentation/data-sources/app_firewall/properties/enable_ai_enhancements/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3330323012211103-0031131130300122-2011021101121023-3112012123021123-3231300122313331-1101212033212032-3302322233212232-3200321220223211", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_ai_enhancements"], "schema_version": 1, "sections": [{"aliases": ["enable ai enhancements mitigate high medium risk action"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:enable_ai_enhancements:mitigate_high_medium_risk_action", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_ai_enhancements", "mitigate_high_medium_risk_action"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable ai enhancements mitigate high risk action"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:enable_ai_enhancements:mitigate_high_risk_action", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_ai_enhancements", "mitigate_high_risk_action"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/enable_ai_enhancements/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Actions complimented by the additional intelligence of the F5 AI Powered Risk-based analysis.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["app_firewallCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_ai_enhancements

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- enable_ai_enhancements

<a id="section"></a>

Type: `"single"`. Computed.

Actions complimented by the additional intelligence of the F5 AI Powered Risk-based analysis.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-risk_score_action_choice": "[\"mitigate_high_medium_risk_action\",\"mitigate_high_risk_action\"]"
}
```

## Direct properties

- [mitigate_high_medium_risk_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/enable_ai_enhancements/mitigate_high_medium_risk_action/): complete subsection reference.

- [mitigate_high_risk_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/enable_ai_enhancements/mitigate_high_risk_action/): complete subsection reference.
