---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.financial_services"
subcategory: "Load Balancing"
description: "Bot Defense Flow Label Financial Services Category."
xcsh_docs: {"aliases": ["bot defense policy protected app endpoints flow label financial services"], "body_bytes": 2310, "body_sha256": "sha256:792f988b8485d93660c35b8200a444f30f8af4a55fbda13d1518460999b53552", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services:apply", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services:money_transfer"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "path": "documentation/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/financial_services/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2313330222120110-2222031320011333-2210303331220110-1210112000113221-0120310301320110-3320203021001300-3323133332212333-0032021010221112", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.flow_label.financial_services:ConflictingObjectAttributes:apply,money_transfer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services:apply", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.flow_label.financial_services:ConflictingObjectAttributes:apply,money_transfer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services:money_transfer", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "financial_services"], "schema_version": 1, "sections": [{"aliases": ["bot defense policy protected app endpoints flow label financial services apply"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services:apply", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "financial_services", "apply"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy protected app endpoints flow label financial services money transfer"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services:money_transfer", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "financial_services", "money_transfer"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/financial_services/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Bot Defense Flow Label Financial Services Category.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.financial_services

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Financial Services Category.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("apply",
    "money_transfer")}
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
  "x-ves-oneof-field-label_choice": "[\"apply\",\"money_transfer\"]"
}
```

Terraform syntax:

```terraform
financial_services {
  # Configure direct properties listed below.
}
```

## Direct properties

- [apply](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/financial_services/apply/): complete subsection reference.

- [money_transfer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/financial_services/money_transfer/): complete subsection reference.
