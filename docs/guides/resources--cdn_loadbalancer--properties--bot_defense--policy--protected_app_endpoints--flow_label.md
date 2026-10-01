---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.flow_label for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 5850, "body_sha256": "sha256:735b91a79fed6ced2126eb9a0546ac6d8d2068764dc8aec9525e6a730568ef84", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:flight", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:profile_management", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "path": "docs/guides/resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.flow_label for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [bot_defense](resources--cdn_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--cdn_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
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

- [account_management](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--account_management.md): complete subsection reference.

- [authentication](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication.md): complete subsection reference.

- [financial_services](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--financial_services.md): complete subsection reference.

- [flight](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--flight.md): complete subsection reference.

- [profile_management](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--profile_management.md): complete subsection reference.

- [search](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--search.md): complete subsection reference.

- [shopping_gift_cards](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards.md): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--account_management.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--financial_services.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--flight.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--profile_management.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--search.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards.md)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
