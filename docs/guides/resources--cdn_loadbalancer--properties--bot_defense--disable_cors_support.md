---
page_title: "bot_defense.disable_cors_support"
subcategory: "Load Balancing"
description: "bot_defense.disable_cors_support for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1023, "body_sha256": "sha256:8c0b244878709fe65b116ddab4b5b0e9b10216e98e57d2f721514e35347ae0fc", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:disable_cors_support", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:disable_cors_support", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense", "path": "docs/guides/resources--cdn_loadbalancer--properties--bot_defense--disable_cors_support.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "disable_cors_support"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/disable_cors_support/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.disable_cors_support for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.disable_cors_support

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [bot_defense](resources--cdn_loadbalancer--properties--bot_defense.md)
- bot_defense.disable_cors_support

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
disable_cors_support = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [bot_defense](resources--cdn_loadbalancer--properties--bot_defense.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
