---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.account_management"
subcategory: "Load Balancing"
description: "Bot Defense Flow Label Account Management Category."
xcsh_docs: {"aliases": ["bot defense policy protected app endpoints flow label account management"], "body_bytes": 2885, "body_sha256": "sha256:d40d004fc490fe2d9ae16245f2bdbbb123cd36cf192174c1b37ca5ac5fa0e686", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management:create", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management:password_reset"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "path": "documentation/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3012232111033320-2231113022320000-3013203130211113-3211022000030112-2123320300103221-2123022311030010-2012200000232313-2033302023012203", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "account_management"], "schema_version": 1, "sections": [{"aliases": ["create"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management:create", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "account_management", "create"], "syntax": "attribute", "type": "object"}, {"aliases": ["password reset"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management:password_reset", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "account_management", "password_reset"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Bot Defense Flow Label Account Management Category.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.account_management

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Account Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

## Direct properties

- [create](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/create/): complete subsection reference.

- [password_reset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/password_reset/): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/create/)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/password_reset/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
