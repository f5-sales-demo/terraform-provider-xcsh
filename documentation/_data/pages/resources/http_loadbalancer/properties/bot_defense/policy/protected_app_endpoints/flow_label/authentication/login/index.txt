---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.authentication.login"
subcategory: "Load Balancing"
description: "Bot Defense Transaction Result."
xcsh_docs: {"aliases": ["bot defense policy protected app endpoints flow label authentication login", "login", "login result", "sign in"], "body_bytes": 3657, "body_sha256": "sha256:05aebb84aafab08a7aacccd87933ee0163a356f75d9f9a7d864c523505e7497c", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:disable_transaction_result", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:transaction_result"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication", "path": "documentation/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-012.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.flow_label.authentication.login:ConflictingObjectAttributes:disable_transaction_result,transaction_result", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:disable_transaction_result", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.flow_label.authentication.login:ConflictingObjectAttributes:disable_transaction_result,transaction_result", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:transaction_result", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "authentication", "login"], "schema_version": 1, "sections": [{"aliases": ["bot defense policy protected app endpoints flow label authentication login disable transaction result", "login", "login result", "sign in"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:disable_transaction_result", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "authentication", "login", "disable_transaction_result"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy protected app endpoints flow label authentication login transaction result", "login", "login result", "sign in"], "anchor": "section", "description": "Bot Defense Transaction ResultType.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:transaction_result", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "authentication", "login", "transaction_result"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Bot Defense Transaction Result.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Upstream description:

Bot Defense Transaction Result.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_transaction_result",
    "transaction_result")}
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
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

Terraform syntax:

```terraform
login {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_transaction_result](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/disable_transaction_result/): complete subsection reference.

- [transaction_result](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/transaction_result/): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/disable_transaction_result/)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/transaction_result/)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
