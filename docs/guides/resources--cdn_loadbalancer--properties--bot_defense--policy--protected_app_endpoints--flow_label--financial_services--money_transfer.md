---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1862, "body_sha256": "sha256:de8fd1e43075dd74119082ef68e478f252665c9a579219f664c4853107980641", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services:money_transfer", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services:money_transfer", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services", "path": "docs/guides/resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--financial_services--money_transfer.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "financial_services", "money_transfer"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/financial_services/money_transfer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [bot_defense](resources--cdn_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--cdn_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--financial_services.md)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for money transfer.

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
money_transfer = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--financial_services.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
