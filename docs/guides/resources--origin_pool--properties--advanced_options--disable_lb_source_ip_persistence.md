---
page_title: "advanced_options.disable_lb_source_ip_persistence"
subcategory: "Load Balancing"
description: "advanced_options.disable_lb_source_ip_persistence for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 922, "body_sha256": "sha256:0d44adbba3844e0a26a7d114f5f91391b26ab5f83cc3b16f8b0b5d48fd0c7cc8", "canonical_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_lb_source_ip_persistence", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_lb_source_ip_persistence", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "path": "docs/guides/resources--origin_pool--properties--advanced_options--disable_lb_source_ip_persistence.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options", "disable_lb_source_ip_persistence"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/disable_lb_source_ip_persistence/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.disable_lb_source_ip_persistence for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# advanced_options.disable_lb_source_ip_persistence

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- advanced_options.disable_lb_source_ip_persistence

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

IP address configuration

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
disable_lb_source_ip_persistence = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
