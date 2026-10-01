---
page_title: "bot_defense.policy.protected_app_endpoints.mitigation"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.mitigation for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2258, "body_sha256": "sha256:1afa0336f1a15abc4fc59d22693e1b266710af768a6f99dfa5916824b1a74758", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:block", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:redirect"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.mitigation for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.mitigation

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [bot_defense](data-sources--cdn_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](data-sources--cdn_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- bot_defense.policy.protected_app_endpoints.mitigation

<a id="section"></a>

Type: `"single"`. Computed.

Modify Bot Defense behavior for a matching request.

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

## Direct properties

- [block](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--block.md): complete subsection reference.

- [flag](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag.md): complete subsection reference.

- [redirect](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--redirect.md): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.mitigation.block](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--block.md)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag.md)
- [bot_defense.policy.protected_app_endpoints.mitigation.redirect](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--redirect.md)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
