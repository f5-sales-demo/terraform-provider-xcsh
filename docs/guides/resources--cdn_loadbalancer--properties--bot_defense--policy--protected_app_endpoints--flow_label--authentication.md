---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.authentication"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.flow_label.authentication for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4372, "body_sha256": "sha256:4c1c9cf47854aba22666a2a663dfef5a1e42eafe16ce7658e625dd679c4592ca", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login_mfa", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:login_partner", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:logout", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication:token_refresh"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:authentication", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "path": "docs/guides/resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "authentication"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/authentication/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.flow_label.authentication for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.authentication

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [bot_defense](resources--cdn_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--cdn_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Authentication Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("login",
    "login_mfa"),
  validators.ConflictingObjectAttributes("login",
    "login_partner"),
  validators.ConflictingObjectAttributes("login",
    "logout"),
  validators.ConflictingObjectAttributes("login",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_mfa",
    "login_partner"),
  validators.ConflictingObjectAttributes("login_mfa",
    "logout"),
  validators.ConflictingObjectAttributes("login_mfa",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_partner",
    "logout"),
  validators.ConflictingObjectAttributes("login_partner",
    "token_refresh"),
  validators.ConflictingObjectAttributes("logout",
    "token_refresh")}
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
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

Terraform syntax:

```terraform
authentication {
  # Configure direct properties listed below.
}
```

## Direct properties

- [login](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--login.md): complete subsection reference.

- [login_mfa](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--login_mfa.md): complete subsection reference.

- [login_partner](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--login_partner.md): complete subsection reference.

- [logout](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--logout.md): complete subsection reference.

- [token_refresh](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--token_refresh.md): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--login.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--login_mfa.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--login_partner.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--logout.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--authentication--token_refresh.md)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
