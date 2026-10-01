---
page_title: "advanced_options.enable_subsets.default_subset"
subcategory: "Load Balancing"
description: "advanced_options.enable_subsets.default_subset for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1480, "body_sha256": "sha256:eec1249d53851b6782ed4721ba0cda9221e472a567dcc764f33702ceeb0310fb", "canonical_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:default_subset", "child_ids": ["xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:default_subset:default_subset"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:default_subset", "parent_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets", "path": "docs/guides/resources--origin_pool--properties--advanced_options--enable_subsets--default_subset.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options", "enable_subsets", "default_subset"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/enable_subsets/default_subset/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options.enable_subsets.default_subset for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options.enable_subsets.default_subset

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- [advanced_options.enable_subsets](resources--origin_pool--properties--advanced_options--enable_subsets.md)
- advanced_options.enable_subsets.default_subset

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default subset.

Upstream description:

Default Subset definition.

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
default_subset {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_subset](resources--origin_pool--properties--advanced_options--enable_subsets--default_subset--default_subset.md): complete subsection reference.

## Next pages

- [advanced_options.enable_subsets.default_subset.default_subset](resources--origin_pool--properties--advanced_options--enable_subsets--default_subset--default_subset.md)
- [advanced_options.enable_subsets](resources--origin_pool--properties--advanced_options--enable_subsets.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
