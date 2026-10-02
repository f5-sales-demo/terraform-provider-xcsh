---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.authentication.login"
subcategory: "Load Balancing"
description: "Bot Defense Transaction Result."
xcsh_docs: {"aliases": ["authentication", "bot defense policy protected app endpoints flow label authentication login", "credential setup", "credentials", "login", "login result", "sign in"], "body_bytes": 3373, "body_sha256": "sha256:ac81819ae40956af49cf4adb9be8728cafeaf51dc6e26a3d9378f12ae12e76bb", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:disable_transaction_result", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:transaction_result"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication", "path": "documentation/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3131113030220122-0202331010323001-2122322201031222-2313130113302010-2320111320101310-1113132210221330-1210023001313110-1031131310202220", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "authentication", "login"], "schema_version": 1, "sections": [{"aliases": ["disable transaction result"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:disable_transaction_result", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "authentication", "login", "disable_transaction_result"], "syntax": "attribute", "type": "object"}, {"aliases": ["transaction result"], "anchor": "section", "description": "Bot Defense Transaction ResultType.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:transaction_result", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "authentication", "login", "transaction_result"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Bot Defense Transaction Result.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Upstream description:

Bot Defense Transaction Result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

## Direct properties

- [disable_transaction_result](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/disable_transaction_result/): complete subsection reference.

- [transaction_result](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/transaction_result/): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/disable_transaction_result/)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/transaction_result/)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
