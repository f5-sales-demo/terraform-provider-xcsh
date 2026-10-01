---
page_title: "bot_defense.policy.protected_app_endpoints.mitigation.flag"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.mitigation.flag for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2270, "body_sha256": "sha256:1fb2612e82f8f1537ade33873d52910d71fe0720150499f731aafc81aadbfa0c", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:append_headers", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:no_headers"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation", "flag"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.mitigation.flag for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.mitigation.flag

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [bot_defense](data-sources--cdn_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](data-sources--cdn_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [bot_defense.policy.protected_app_endpoints.mitigation](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation.md)
- bot_defense.policy.protected_app_endpoints.mitigation.flag

<a id="section"></a>

Type: `"single"`. Computed.

Select Flag Bot Mitigation Action. Flag mitigation action.

Upstream description:

Flag mitigation action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-send_headers_choice": "[\"append_headers\",\"no_headers\"]"
}
```

## Direct properties

- [append_headers](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag--append_headers.md): complete subsection reference.

- [no_headers](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag--no_headers.md): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag--append_headers.md)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag--no_headers.md)
- [bot_defense.policy.protected_app_endpoints.mitigation](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
