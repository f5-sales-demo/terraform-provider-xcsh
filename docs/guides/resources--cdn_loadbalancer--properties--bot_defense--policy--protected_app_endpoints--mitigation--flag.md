---
page_title: "bot_defense.policy.protected_app_endpoints.mitigation.flag"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.mitigation.flag for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2521, "body_sha256": "sha256:2cf8e7c0e1b1b46b853b84bc8471f9e9d36b1d965b7416cd978f807a4dd2698b", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:append_headers", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:no_headers"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "path": "docs/guides/resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation", "flag"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.mitigation.flag for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.mitigation.flag

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [bot_defense](resources--cdn_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--cdn_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation.md)
- bot_defense.policy.protected_app_endpoints.mitigation.flag

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select Flag Bot Mitigation Action. Flag mitigation action.

Upstream description:

Flag mitigation action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_headers",
    "no_headers")}
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
  "x-ves-oneof-field-send_headers_choice": "[\"append_headers\",\"no_headers\"]"
}
```

Terraform syntax:

```terraform
flag {
  # Configure direct properties listed below.
}
```

## Direct properties

- [append_headers](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag--append_headers.md): complete subsection reference.

- [no_headers](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag--no_headers.md): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag--append_headers.md)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation--flag--no_headers.md)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--mitigation.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
