---
page_title: "bot_defense.policy.mobile_sdk_config.mobile_identifier"
subcategory: "Load Balancing"
description: "bot_defense.policy.mobile_sdk_config.mobile_identifier for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1678, "body_sha256": "sha256:1aee60eca4e522e0e7e324ea0c2a64c9db2c1ac35e9ce68d96f59e661e5696b9", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config", "path": "docs/guides/resources--http_loadbalancer--properties--bot_defense--policy--mobile_sdk_config--mobile_identifier.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "mobile_sdk_config", "mobile_identifier"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.mobile_sdk_config.mobile_identifier for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.mobile_sdk_config.mobile_identifier

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [bot_defense](resources--http_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--http_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--properties--bot_defense--policy--mobile_sdk_config.md)
- bot_defense.policy.mobile_sdk_config.mobile_identifier

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Mobile Traffic Identifier. Mobile traffic identifier type.

Upstream description:

Mobile traffic identifier type.

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
mobile_identifier {
  # Configure direct properties listed below.
}
```

## Direct properties

- [headers](resources--http_loadbalancer--properties--bot_defense--policy--mobile_sdk_config--mobile_identifier--headers.md): complete subsection reference.

## Next pages

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](resources--http_loadbalancer--properties--bot_defense--policy--mobile_sdk_config--mobile_identifier--headers.md)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--properties--bot_defense--policy--mobile_sdk_config.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
