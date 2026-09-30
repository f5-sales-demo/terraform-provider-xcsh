---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.flow_label for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3894, "body_sha256": "sha256:70b261942ef88ac4030cb4187573c4e1d5e95198576a04d20c96548325c07584", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:flight", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:profile_management", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "path": "docs/guides/data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.flow_label for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense.policy.protected_app_endpoints.flow_label

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [bot_defense](data-sources--http_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](data-sources--http_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
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

- [account_management](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--account_management.md): complete subsection reference.

- [authentication](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication.md): complete subsection reference.

- [financial_services](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--financial_services.md): complete subsection reference.

- [flight](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--flight.md): complete subsection reference.

- [profile_management](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--profile_management.md): complete subsection reference.

- [search](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--search.md): complete subsection reference.

- [shopping_gift_cards](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards.md): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--account_management.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--financial_services.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--flight.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--profile_management.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--search.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards.md)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
