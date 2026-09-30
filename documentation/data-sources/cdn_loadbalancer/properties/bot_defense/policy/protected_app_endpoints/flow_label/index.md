---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.flow_label for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4883, "body_sha256": "sha256:a9bb72dbb9509d1e9f2ce13bf2617b7d1918ce20614d5d1fd7c59dc24d53e87a", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:flight", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:profile_management", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "path": "documentation/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.flow_label for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense.policy.protected_app_endpoints.flow_label

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- bot_defense.policy.protected_app_endpoints.flow_label

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

## Direct properties

- [account_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/): complete subsection reference.

- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/): complete subsection reference.

- [financial_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/financial_services/): complete subsection reference.

- [flight](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/flight/): complete subsection reference.

- [profile_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/profile_management/): complete subsection reference.

- [search](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/): complete subsection reference.

- [shopping_gift_cards](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/shopping_gift_cards/): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/financial_services/)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/flight/)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/profile_management/)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/shopping_gift_cards/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
