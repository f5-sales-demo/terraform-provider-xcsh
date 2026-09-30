---
page_title: "advanced_options.http1_config"
subcategory: "Load Balancing"
description: "advanced_options.http1_config for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1137, "body_sha256": "sha256:95635b64337d37570f4977ba7233ed994b69151150df5df5ca5decb3e8c31961", "canonical_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config", "child_ids": ["xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config:header_transformation"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "path": "docs/guides/resources--origin_pool--properties--advanced_options--http1_config.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options", "http1_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/http1_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.http1_config for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# advanced_options.http1_config

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- advanced_options.http1_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for upstream connections.

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
http1_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [header_transformation](resources--origin_pool--properties--advanced_options--http1_config--header_transformation.md): complete subsection reference.

## Next pages

- [advanced_options.http1_config.header_transformation](resources--origin_pool--properties--advanced_options--http1_config--header_transformation.md)
- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
