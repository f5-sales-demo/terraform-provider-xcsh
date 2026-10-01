---
page_title: "bot_defense.policy.mobile_sdk_config"
subcategory: "Load Balancing"
description: "bot_defense.policy.mobile_sdk_config for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1450, "body_sha256": "sha256:a6f8c95ba637b8f894c1478cf74b9a3dc8b7d2b0b5661120d5f177f4d530fc09", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy", "path": "docs/guides/resources--http_loadbalancer--properties--bot_defense--policy--mobile_sdk_config.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "mobile_sdk_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.mobile_sdk_config for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.mobile_sdk_config

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [bot_defense](resources--http_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--http_loadbalancer--properties--bot_defense--policy.md)
- bot_defense.policy.mobile_sdk_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Mobile SDK Configuration. Mobile SDK configuration.

Upstream description:

Mobile SDK configuration.

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
mobile_sdk_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [mobile_identifier](resources--http_loadbalancer--properties--bot_defense--policy--mobile_sdk_config--mobile_identifier.md): complete subsection reference.

## Next pages

- [bot_defense.policy.mobile_sdk_config.mobile_identifier](resources--http_loadbalancer--properties--bot_defense--policy--mobile_sdk_config--mobile_identifier.md)
- [bot_defense.policy](resources--http_loadbalancer--properties--bot_defense--policy.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
