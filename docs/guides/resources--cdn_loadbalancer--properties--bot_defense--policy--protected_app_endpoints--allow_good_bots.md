---
page_title: "bot_defense.policy.protected_app_endpoints.allow_good_bots"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.allow_good_bots for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1385, "body_sha256": "sha256:e4ce25ff6850fbdf850d195059a368c1acdba33e1ed770a8b339ac102ddd6f5a", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:allow_good_bots", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:allow_good_bots", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "path": "docs/guides/resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--allow_good_bots.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "allow_good_bots"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/allow_good_bots/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.allow_good_bots for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.allow_good_bots

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [bot_defense](resources--cdn_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--cdn_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- bot_defense.policy.protected_app_endpoints.allow_good_bots

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for allow good bots.

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
allow_good_bots = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
