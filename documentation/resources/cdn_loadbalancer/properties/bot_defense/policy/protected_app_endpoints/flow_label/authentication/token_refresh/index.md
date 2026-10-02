---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["authentication", "bot defense policy protected app endpoints flow label authentication token refresh", "credential setup", "credentials"], "body_bytes": 2281, "body_sha256": "sha256:57d47a3b22b9d992d8eba4feca6aa65dc969b22d60a1a4dba8527476bd67e408", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:token_refresh", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication", "path": "documentation/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/token_refresh/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0133212203133310-1201032011111112-3320330112303200-1103301310102232-0003031301102113-0233132230001030-0131230002001102-3002331120301112", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "authentication", "token_refresh"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/token_refresh/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for token refresh.

Upstream description:

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
token_refresh = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
