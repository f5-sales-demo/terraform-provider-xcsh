---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.authentication.login"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.flow_label.authentication.login for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2931, "body_sha256": "sha256:58e59424ae31ae0cf348c31668babe6bfd2e5cb6261ede2d8d18ab1bcc7dc7f3", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:disable_transaction_result", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login:transaction_result"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication", "path": "docs/guides/resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--login.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "authentication", "login"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/login/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.flow_label.authentication.login for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [bot_defense](resources--http_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--http_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication.md)
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

- [disable_transaction_result](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--login--disable_transaction_result.md): complete subsection reference.

- [transaction_result](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--login--transaction_result.md): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--login--disable_transaction_result.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--login--transaction_result.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
