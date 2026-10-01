---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.financial_services"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.flow_label.financial_services for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2295, "body_sha256": "sha256:5c9377fd642b78d13bfe746cd242ca8b2dc666808e397d4457a2641c439b01e2", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services:apply", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services:money_transfer"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:financial_services", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--financial_services.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "financial_services"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/financial_services/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.flow_label.financial_services for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.financial_services

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [bot_defense](data-sources--cdn_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](data-sources--cdn_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Financial Services Category.

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

## Direct properties

- [apply](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--financial_services--apply.md): complete subsection reference.

- [money_transfer](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--financial_services--money_transfer.md): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--financial_services--apply.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--financial_services--money_transfer.md)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
