---
page_title: "rule_list.rules.spec.waf_action"
subcategory: "Security"
description: "Modify App Firewall behavior for a matching request. The modification could either be to entirely skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall Rule Control settings."
xcsh_docs: {"aliases": ["rule list rules spec waf action"], "body_bytes": 3118, "body_sha256": "sha256:1c3d2a3b86e78b72d8760c29505aa6e64c97f138e03c0b8b7375da84a5152981", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:none", "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:waf_skip_processing"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action", "parent_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec", "path": "documentation/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100", "registry_path": "docs/guides/data-sources--service_policy--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec", "waf_action"], "schema_version": 1, "sections": [{"aliases": ["app firewall detection control"], "anchor": "section", "description": "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "waf_action", "app_firewall_detection_control"], "syntax": "attribute", "type": "object"}, {"aliases": ["none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:none", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "waf_action", "none"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf skip processing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:service_policy:properties:rule_list:rules:spec:waf_action:waf_skip_processing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "waf_action", "waf_skip_processing"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Modify App Firewall behavior for a matching request. The modification could either be to entirely skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall Rule Control settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.waf_action

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/)
- rule_list.rules.spec.waf_action

<a id="section"></a>

Type: `"single"`. Computed.

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Upstream description:

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

- [app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/): complete subsection reference.

- [none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/none/): complete subsection reference.

- [waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/waf_skip_processing/): complete subsection reference.

## Next pages

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/)
- [rule_list.rules.spec.waf_action.none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/none/)
- [rule_list.rules.spec.waf_action.waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/waf_action/waf_skip_processing/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/properties/rule_list/rules/spec/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy/)
