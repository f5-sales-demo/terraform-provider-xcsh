---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.account_management"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.flow_label.account_management for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2561, "body_sha256": "sha256:84f8d4fbaadfb16e88f85c71ab74c6314820d8d79b76d0875272550792d25d5c", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management:create", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management:password_reset"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:account_management", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "path": "docs/guides/resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--account_management.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "account_management"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/account_management/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.flow_label.account_management for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.account_management

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [bot_defense](resources--cdn_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--cdn_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Account Management Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "password_reset")}
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
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

Terraform syntax:

```terraform
account_management {
  # Configure direct properties listed below.
}
```

## Direct properties

- [create](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--account_management--create.md): complete subsection reference.

- [password_reset](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--account_management--password_reset.md): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.create](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--account_management--create.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--account_management--password_reset.md)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
