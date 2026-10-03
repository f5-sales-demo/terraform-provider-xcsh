---
page_title: "waf_action"
subcategory: ""
description: "Modify App Firewall behavior for a matching request. The modification could either be to entirely skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall Rule Control settings."
xcsh_docs: {"aliases": ["waf action"], "body_bytes": 2980, "body_sha256": "sha256:4fdad58e0816417d4eb7a707128f93a17157f4ed643886d216bfba41e825a020", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "xcsh-docs:resources:service_policy_rule:properties:waf_action:none", "xcsh-docs:resources:service_policy_rule:properties:waf_action:waf_skip_processing"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:waf_action", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "documentation/resources/service_policy_rule/properties/waf_action/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "waf_action:ConflictingObjectAttributes:app_firewall_detection_control,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_action:ConflictingObjectAttributes:app_firewall_detection_control,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_action:ConflictingObjectAttributes:app_firewall_detection_control,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_action:ConflictingObjectAttributes:none,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_action:ConflictingObjectAttributes:app_firewall_detection_control,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:waf_skip_processing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_action:ConflictingObjectAttributes:none,waf_skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:waf_skip_processing", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_action"], "schema_version": 1, "sections": [{"aliases": ["waf action app firewall detection control"], "anchor": "section", "description": "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:app_firewall_detection_control", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["waf_action", "app_firewall_detection_control"], "syntax": "block", "type": "object"}, {"aliases": ["waf action none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:none", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_action", "none"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf action waf skip processing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:waf_action:waf_skip_processing", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_action", "waf_skip_processing"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/waf_action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Modify App Firewall behavior for a matching request. The modification could either be to entirely skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall Rule Control settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_action

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- waf_action

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

- [app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/): complete subsection reference.

- [none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/none/): complete subsection reference.

- [waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/waf_skip_processing/): complete subsection reference.

## Next pages

- [waf_action.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/app_firewall_detection_control/)
- [waf_action.none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/none/)
- [waf_action.waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/waf_action/waf_skip_processing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
