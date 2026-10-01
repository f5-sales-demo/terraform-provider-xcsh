---
page_title: "bot_defense.policy.protected_app_endpoints.mitigation"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.mitigation for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2273, "body_sha256": "sha256:c06f3fcbff20f1714adcb0b530609ae8bae7630a98260ecb9d2494a0c30b9aa2", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:block", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:redirect"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "path": "docs/guides/data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.mitigation for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.mitigation

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [bot_defense](data-sources--http_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](data-sources--http_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
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

- [block](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--block.md): complete subsection reference.

- [flag](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag.md): complete subsection reference.

- [redirect](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--redirect.md): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.mitigation.block](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--block.md)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag.md)
- [bot_defense.policy.protected_app_endpoints.mitigation.redirect](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--redirect.md)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
