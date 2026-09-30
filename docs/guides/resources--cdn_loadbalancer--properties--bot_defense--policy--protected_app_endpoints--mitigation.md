---
page_title: "bot_defense.policy.protected_app_endpoints.mitigation"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.mitigation for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2531, "body_sha256": "sha256:7d54e3c6df9ef7c48369a7663d75f09c2cb7abd4d8ce83149b878f55224d21a8", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:block", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:redirect"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "path": "docs/guides/resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.mitigation for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense.policy.protected_app_endpoints.mitigation

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [bot_defense](resources--cdn_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--cdn_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- bot_defense.policy.protected_app_endpoints.mitigation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Modify Bot Defense behavior for a matching request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "flag"),
  validators.ConflictingObjectAttributes("block",
    "redirect"),
  validators.ConflictingObjectAttributes("flag",
    "redirect")}
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
  "x-ves-oneof-field-action_type": "[\"block\",\"flag\",\"redirect\"]"
}
```

Terraform syntax:

```terraform
mitigation {
  # Configure direct properties listed below.
}
```

## Direct properties

- [block](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--block.md): complete subsection reference.

- [flag](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag.md): complete subsection reference.

- [redirect](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--redirect.md): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.mitigation.block](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--block.md)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag.md)
- [bot_defense.policy.protected_app_endpoints.mitigation.redirect](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--redirect.md)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
