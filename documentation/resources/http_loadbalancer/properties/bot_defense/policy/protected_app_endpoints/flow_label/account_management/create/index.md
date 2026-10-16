---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.account_management.create"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["bot defense policy protected app endpoints flow label account management create"], "body_bytes": 1880, "body_sha256": "sha256:9478e77388751e1422bc74cda404a8722e6f786aa0a9d545948366728e95ac3e", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management:create", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management", "path": "documentation/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/create/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-2301020011313203-3212210013121110-1231032200222233-0103330212020133-1001030121111321-0223030213011232-0102201021123103-3330230323022223", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "account_management", "create"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/create/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.account_management.create

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.create

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
create = {}
```

This is an empty object or choice marker. It has no direct properties.
