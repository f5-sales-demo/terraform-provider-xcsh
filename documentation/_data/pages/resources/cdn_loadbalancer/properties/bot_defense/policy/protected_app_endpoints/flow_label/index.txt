---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.flow_label for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 6862, "body_sha256": "sha256:6b427ebe6222d931b0c3df52ea0a70d659df1cad8961bd16bac1e33d9edb01d3", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:flight", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:profile_management", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "path": "documentation/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.flow_label for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- bot_defense.policy.protected_app_endpoints.flow_label

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("account_management",
    "authentication"),
  validators.ConflictingObjectAttributes("account_management",
    "financial_services"),
  validators.ConflictingObjectAttributes("account_management",
    "flight"),
  validators.ConflictingObjectAttributes("account_management",
    "profile_management"),
  validators.ConflictingObjectAttributes("account_management",
    "search"),
  validators.ConflictingObjectAttributes("account_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("authentication",
    "financial_services"),
  validators.ConflictingObjectAttributes("authentication",
    "flight"),
  validators.ConflictingObjectAttributes("authentication",
    "profile_management"),
  validators.ConflictingObjectAttributes("authentication",
    "search"),
  validators.ConflictingObjectAttributes("authentication",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("financial_services",
    "flight"),
  validators.ConflictingObjectAttributes("financial_services",
    "profile_management"),
  validators.ConflictingObjectAttributes("financial_services",
    "search"),
  validators.ConflictingObjectAttributes("financial_services",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("flight",
    "profile_management"),
  validators.ConflictingObjectAttributes("flight",
    "search"),
  validators.ConflictingObjectAttributes("flight",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("profile_management",
    "search"),
  validators.ConflictingObjectAttributes("profile_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("search",
    "shopping_gift_cards")}
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
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

Terraform syntax:

```terraform
flow_label {
  # Configure direct properties listed below.
}
```

## Direct properties

- [account_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/): complete subsection reference.

- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/): complete subsection reference.

- [financial_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/financial_services/): complete subsection reference.

- [flight](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/flight/): complete subsection reference.

- [profile_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/profile_management/): complete subsection reference.

- [search](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/): complete subsection reference.

- [shopping_gift_cards](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/shopping_gift_cards/): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/financial_services/)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/flight/)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/profile_management/)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/shopping_gift_cards/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
