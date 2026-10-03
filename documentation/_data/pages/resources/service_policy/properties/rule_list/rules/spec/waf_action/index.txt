---
page_title: "rule_list.rules.spec.waf_action"
subcategory: "Security"
description: "Modify App Firewall behavior for a matching request. The modification could either be to entirely skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall Rule Control settings."
xcsh_docs: {"aliases": ["rule list rules spec waf action"], "body_bytes": 3562, "body_sha256": "sha256:dd749c488ccd60fb1ee7dd09363b42890f275adb8e84be21ce1ffd90e683b801", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:none", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:waf_skip_processing"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "path": "documentation/resources/service_policy/properties/rule_list/rules/spec/waf_action/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1330120330031000-3202113320223100-2303020333222232-1113322320113320-2131320031302300-3021303213200102-0321303213312133-3121220001012010", "registry_path": "docs/guides/resources--service_policy--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.waf_action:ConflictingObjectAttributes:app_firewall_detection_control,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.waf_action:ConflictingObjectAttributes:app_firewall_detection_control,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.waf_action:ConflictingObjectAttributes:app_firewall_detection_control,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.waf_action:ConflictingObjectAttributes:none,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.waf_action:ConflictingObjectAttributes:app_firewall_detection_control,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:waf_skip_processing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.waf_action:ConflictingObjectAttributes:none,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:waf_skip_processing", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec", "waf_action"], "schema_version": 1, "sections": [{"aliases": ["rule list rules spec waf action app firewall detection control"], "anchor": "section", "description": "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:app_firewall_detection_control", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "spec", "waf_action", "app_firewall_detection_control"], "syntax": "block", "type": "object"}, {"aliases": ["rule list rules spec waf action none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:none", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "waf_action", "none"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list rules spec waf action waf skip processing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:waf_action:waf_skip_processing", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "waf_action", "waf_skip_processing"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/waf_action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Modify App Firewall behavior for a matching request. The modification could either be to entirely skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall Rule Control settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.waf_action

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/)
- rule_list.rules.spec.waf_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Upstream description:

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall_detection_control",
    "none"),
  validators.ConflictingObjectAttributes("app_firewall_detection_control",
    "waf_skip_processing"),
  validators.ConflictingObjectAttributes("none",
    "waf_skip_processing")}
```

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

Terraform syntax:

```terraform
waf_action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/): complete subsection reference.

- [none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/none/): complete subsection reference.

- [waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/waf_skip_processing/): complete subsection reference.

## Next pages

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/app_firewall_detection_control/)
- [rule_list.rules.spec.waf_action.none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/none/)
- [rule_list.rules.spec.waf_action.waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/waf_action/waf_skip_processing/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
