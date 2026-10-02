---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.account_management"
subcategory: "Load Balancing"
description: "Bot Defense Flow Label Account Management Category."
xcsh_docs: {"aliases": ["bot defense policy protected app endpoints flow label account management"], "body_bytes": 2899, "body_sha256": "sha256:6e9f7d3bf07fe5872d10ee4daa2a21b7898569d5f0f01026c4b31581819f7070", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management:create", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management:password_reset"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "path": "documentation/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0222003330120203-0023033212212033-0001102031020013-2202103300233003-1320322202210012-0223300222003003-1221013010300201-0013211001201210", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "account_management"], "schema_version": 1, "sections": [{"aliases": ["create"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management:create", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "account_management", "create"], "syntax": "attribute", "type": "object"}, {"aliases": ["password reset"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management:password_reset", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "account_management", "password_reset"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Bot Defense Flow Label Account Management Category.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.account_management

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
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

- [create](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/create/): complete subsection reference.

- [password_reset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/password_reset/): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/create/)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/password_reset/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
