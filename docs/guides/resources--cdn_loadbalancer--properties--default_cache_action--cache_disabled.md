---
page_title: "default_cache_action.cache_disabled"
subcategory: "Load Balancing"
description: "default_cache_action.cache_disabled for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 960, "body_sha256": "sha256:abad93249d8db385397961d74c04f4cb857c10a6b4463a2ab4f3d6c62504004e", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_cache_action:cache_disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_cache_action:cache_disabled", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:default_cache_action", "path": "docs/guides/resources--cdn_loadbalancer--properties--default_cache_action--cache_disabled.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_cache_action", "cache_disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/default_cache_action/cache_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_cache_action.cache_disabled for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# default_cache_action.cache_disabled

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [default_cache_action](resources--cdn_loadbalancer--properties--default_cache_action.md)
- default_cache_action.cache_disabled

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
cache_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [default_cache_action](resources--cdn_loadbalancer--properties--default_cache_action.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
