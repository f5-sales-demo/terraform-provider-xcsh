---
page_title: "spoke_mesh.data_plane_mesh"
subcategory: "Infrastructure"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["spoke mesh data plane mesh"], "body_bytes": 1252, "body_sha256": "sha256:16400f6eaf70dfa9d22cd18d33fed435428bb33c2e9e69c889f26905e456f04c", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:data_plane_mesh", "parent_id": "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh", "path": "documentation/resources/site_mesh_group/properties/spoke_mesh/data_plane_mesh/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3313130132312313-3323213312213100-2112302012030121-2212303020012102-3300032113133100-2311213210213213-3033001231302012-0323132112303313", "registry_path": "docs/guides/resources--site_mesh_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["spoke_mesh", "data_plane_mesh"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/properties/spoke_mesh/data_plane_mesh/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# spoke_mesh.data_plane_mesh

Breadcrumbs:

- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/)
- [spoke_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/)
- spoke_mesh.data_plane_mesh

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
data_plane_mesh = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [spoke_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/)
- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/)
