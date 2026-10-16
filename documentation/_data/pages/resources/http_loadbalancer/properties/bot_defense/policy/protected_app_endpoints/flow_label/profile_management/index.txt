---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.profile_management"
subcategory: "Load Balancing"
description: "Bot Defense Flow Label Profile Management Category."
xcsh_docs: {"aliases": ["bot defense policy protected app endpoints flow label profile management"], "body_bytes": 2316, "body_sha256": "sha256:3f838ba0e771f0d3c3572e839986d4bfc91884cde2531b3631f29b046d8b17e5", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:profile_management:create", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:profile_management:update", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:profile_management:view"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:profile_management", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "path": "documentation/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/profile_management/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "profile_management"], "schema_version": 1, "sections": [{"aliases": ["bot defense policy protected app endpoints flow label profile management create"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:profile_management:create", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "profile_management", "create"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy protected app endpoints flow label profile management update"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:profile_management:update", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "profile_management", "update"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy protected app endpoints flow label profile management view"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:profile_management:view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "profile_management", "view"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/profile_management/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Bot Defense Flow Label Profile Management Category.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.profile_management

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Profile Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"update\",\"view\"]"
}
```

Terraform syntax:

```terraform
profile_management {
  # Configure direct properties listed below.
}
```

## Direct properties

- [create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/profile_management/create/): complete subsection reference.

- [update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/profile_management/update/): complete subsection reference.

- [view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/profile_management/view/): complete subsection reference.
