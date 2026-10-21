---
page_title: "type.control_and_data_plane_mesh"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["type control and data plane mesh"], "body_bytes": 1007, "body_sha256": "sha256:bc6829cfa8534e57e163b6d3e17b226e151346125e9c1f44772f98e796efc77f", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dc_cluster_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:dc_cluster_group:properties:type:control_and_data_plane_mesh", "parent_id": "xcsh-docs:resources:dc_cluster_group:properties:type", "path": "documentation/resources/dc_cluster_group/properties/type/control_and_data_plane_mesh/index.md", "product": "distributed-cloud", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0200311211231101-3231031202003133-3121110031113210-2023321233130123-3002332022110202-2110230332101211-0310100113102021-0122022132222311", "registry_path": "docs/guides/resources--dc_cluster_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["type", "control_and_data_plane_mesh"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dc_cluster_group/properties/type/control_and_data_plane_mesh/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# type.control_and_data_plane_mesh

Breadcrumbs:

- [xcsh_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/properties/)
- [type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/properties/type/)
- type.control_and_data_plane_mesh

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
control_and_data_plane_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.
