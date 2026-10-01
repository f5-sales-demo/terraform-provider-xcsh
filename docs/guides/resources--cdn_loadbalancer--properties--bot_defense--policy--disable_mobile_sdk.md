---
page_title: "bot_defense.policy.disable_mobile_sdk"
subcategory: "Load Balancing"
description: "bot_defense.policy.disable_mobile_sdk for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1134, "body_sha256": "sha256:7b35a981715f16e7014d010caaa8af25d6c5ae7c2e168c6715bc9e7ba7766b47", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:disable_mobile_sdk", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:disable_mobile_sdk", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy", "path": "docs/guides/resources--cdn_loadbalancer--properties--bot_defense--policy--disable_mobile_sdk.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "disable_mobile_sdk"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/disable_mobile_sdk/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.disable_mobile_sdk for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.disable_mobile_sdk

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [bot_defense](resources--cdn_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--cdn_loadbalancer--properties--bot_defense--policy.md)
- bot_defense.policy.disable_mobile_sdk

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
disable_mobile_sdk = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [bot_defense.policy](resources--cdn_loadbalancer--properties--bot_defense--policy.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
