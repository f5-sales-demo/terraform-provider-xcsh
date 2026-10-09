---
page_title: "waf_action"
subcategory: ""
description: "Modify App Firewall behavior for a matching request. The modification could either be to entirely skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall Rule Control settings."
xcsh_docs: {"aliases": ["waf action"], "body_bytes": 1524, "body_sha256": "sha256:2d9447fba9711e12fe03758b21cf3c1ddfdd369719e0a7d08b7ddacc8623016c", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:none", "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:waf_skip_processing"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action", "parent_id": "xcsh-docs:data-sources:service_policy_rule:reference", "path": "documentation/data-sources/service_policy_rule/properties/waf_action/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322", "registry_path": "docs/guides/data-sources--service_policy_rule--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_action"], "schema_version": 1, "sections": [{"aliases": ["waf action app firewall detection control"], "anchor": "section", "description": "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.", "document_id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["waf_action", "app_firewall_detection_control"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf action none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:none", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_action", "none"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf action waf skip processing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:service_policy_rule:properties:waf_action:waf_skip_processing", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_action", "waf_skip_processing"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/waf_action/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Modify App Firewall behavior for a matching request. The modification could either be to entirely skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall Rule Control settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_action

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- waf_action

<a id="section"></a>

Type: `"single"`. Computed.

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"app_firewall_detection_control\",\"none\",\"waf_skip_processing\"]"
}
```

## Direct properties

- [app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/waf_action/app_firewall_detection_control/): complete subsection reference.

- [none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/waf_action/none/): complete subsection reference.

- [waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/waf_action/waf_skip_processing/): complete subsection reference.
