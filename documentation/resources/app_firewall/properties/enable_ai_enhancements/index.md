---
page_title: "enable_ai_enhancements"
subcategory: "Security"
description: "Actions complimented by the additional intelligence of the F5 AI Powered Risk-based analysis."
xcsh_docs: {"aliases": ["enable ai enhancements"], "body_bytes": 2241, "body_sha256": "sha256:91cbb3879be94cf8cd2686ff404a28cd7bc304076465e20faae59393bc4b9732", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:app_firewall:properties:enable_ai_enhancements:mitigate_high_medium_risk_action", "xcsh-docs:resources:app_firewall:properties:enable_ai_enhancements:mitigate_high_risk_action"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:enable_ai_enhancements", "parent_id": "xcsh-docs:resources:app_firewall:reference", "path": "documentation/resources/app_firewall/properties/enable_ai_enhancements/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-1212321310122210-3103000112213223-1213303113222233-3133233012310201-3202312230230022-1003203010121000-2313120111303003-3030300312002202", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_ai_enhancements:ConflictingObjectAttributes:mitigate_high_medium_risk_action,mitigate_high_risk_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:enable_ai_enhancements:mitigate_high_medium_risk_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_ai_enhancements:ConflictingObjectAttributes:mitigate_high_medium_risk_action,mitigate_high_risk_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:enable_ai_enhancements:mitigate_high_risk_action", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_ai_enhancements"], "schema_version": 1, "sections": [{"aliases": ["mitigate high medium risk action"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:app_firewall:properties:enable_ai_enhancements:mitigate_high_medium_risk_action", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_ai_enhancements", "mitigate_high_medium_risk_action"], "syntax": "attribute", "type": "object"}, {"aliases": ["mitigate high risk action"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:app_firewall:properties:enable_ai_enhancements:mitigate_high_risk_action", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_ai_enhancements", "mitigate_high_risk_action"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/enable_ai_enhancements/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Actions complimented by the additional intelligence of the F5 AI Powered Risk-based analysis.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_ai_enhancements

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- enable_ai_enhancements

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Actions complimented by the additional intelligence of the F5 AI Powered Risk-based analysis.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("mitigate_high_medium_risk_action",
    "mitigate_high_risk_action")}
```

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

Terraform syntax:

```terraform
enable_ai_enhancements {
  # Configure direct properties listed below.
}
```

## Direct properties

- [mitigate_high_medium_risk_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/enable_ai_enhancements/mitigate_high_medium_risk_action/): complete subsection reference.

- [mitigate_high_risk_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/enable_ai_enhancements/mitigate_high_risk_action/): complete subsection reference.

## Next pages

- [enable_ai_enhancements.mitigate_high_medium_risk_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/enable_ai_enhancements/mitigate_high_medium_risk_action/)
- [enable_ai_enhancements.mitigate_high_risk_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/enable_ai_enhancements/mitigate_high_risk_action/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
