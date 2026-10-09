---
page_title: "type.data_plane_mesh"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["type data plane mesh"], "body_bytes": 971, "body_sha256": "sha256:1c50d0e69af0301a8de3999ad2b7bf1b5a9084cc646a56272abc7dc67aff95ba", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dc_cluster_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:dc_cluster_group:properties:type:data_plane_mesh", "parent_id": "xcsh-docs:resources:dc_cluster_group:properties:type", "path": "documentation/resources/dc_cluster_group/properties/type/data_plane_mesh/index.md", "product": "distributed-cloud", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2320232132000312-1010113003001200-1311022122203122-2211313203022221-3112122033021320-3313203000022331-1320121233021221-3321202230320213", "registry_path": "docs/guides/resources--dc_cluster_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["type", "data_plane_mesh"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dc_cluster_group/properties/type/data_plane_mesh/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# type.data_plane_mesh

Breadcrumbs:

- [xcsh_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/properties/)
- [type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/properties/type/)
- type.data_plane_mesh

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.
