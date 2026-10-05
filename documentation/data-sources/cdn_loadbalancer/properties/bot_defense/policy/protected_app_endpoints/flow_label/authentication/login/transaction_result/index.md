---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result"
subcategory: "Load Balancing"
description: "Bot Defense Transaction ResultType."
xcsh_docs: {"aliases": ["bot defense policy protected app endpoints flow label authentication login transaction result", "login", "login result", "sign in"], "body_bytes": 3663, "body_sha256": "sha256:aa2f7cea357c259b203efe1d703db9d3b22aa549e213efb035ed2d5d25a6c12d", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:transaction_result:failure_conditions", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:transaction_result:success_conditions"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:transaction_result", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login", "path": "documentation/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/transaction_result/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0000221300122333-3101003110122012-2331311331003032-3333033101012222-3202132202131211-1231202120000322-2302211323310213-1221210312023313", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "authentication", "login", "transaction_result"], "schema_version": 1, "sections": [{"aliases": ["bot defense policy protected app endpoints flow label authentication login transaction result failure conditions", "login", "login result", "sign in"], "anchor": "section", "description": "Failure Conditions.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:transaction_result:failure_conditions", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "authentication", "login", "transaction_result", "failure_conditions"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy protected app endpoints flow label authentication login transaction result success conditions", "login", "login result", "login success", "sign in", "succeeded", "success", "successful"], "anchor": "section", "description": "Success Conditions.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:transaction_result:success_conditions", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "authentication", "login", "transaction_result", "success_conditions"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/transaction_result/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Bot Defense Transaction ResultType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

Upstream description:

Bot Defense Transaction ResultType.

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

## Direct properties

- [failure_conditions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/transaction_result/failure_conditions/): complete subsection reference.

- [success_conditions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/transaction_result/success_conditions/): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/transaction_result/failure_conditions/)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/transaction_result/success_conditions/)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
