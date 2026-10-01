---
page_title: "advanced_options.disable_subsets"
subcategory: "Load Balancing"
description: "advanced_options.disable_subsets for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1101, "body_sha256": "sha256:651263326c174af6694c13b1f41b341ee57ced6454dc5866eca22c2371e8e4f7", "canonical_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_subsets", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_subsets", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "path": "docs/guides/resources--origin_pool--properties--advanced_options--disable_subsets.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options", "disable_subsets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/disable_subsets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.disable_subsets for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.disable_subsets

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- advanced_options.disable_subsets

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable subsets. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
disable_subsets = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
