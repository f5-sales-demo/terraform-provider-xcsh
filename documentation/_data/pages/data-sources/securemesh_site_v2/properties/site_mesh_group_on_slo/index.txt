---
page_title: "site_mesh_group_on_slo"
subcategory: ""
description: "site_mesh_group_on_slo for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2802, "body_sha256": "sha256:9dd6171aa52dbc0f6e13b58398042cb53c76f5e27ed321d5a5594cb8a08bd8da", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:no_site_mesh_group", "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:site_mesh_group", "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_public_ip", "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_pvt_ip"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:site_mesh_group_on_slo", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "documentation/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/index.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["site_mesh_group_on_slo"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_mesh_group_on_slo for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_mesh_group_on_slo

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- site_mesh_group_on_slo

<a id="section"></a>

Type: `"single"`. Computed.

Select how the site mesh group will be connected. By default, public IPs of the control nodes of the
site will be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-site_mesh_group_choice": "[\"no_site_mesh_group\",\"site_mesh_group\"]",
  "x-ves-oneof-field-site_mesh_group_ip_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

## Direct properties

- [no_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/no_site_mesh_group/): complete subsection reference.

- [site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/site_mesh_group/): complete subsection reference.

- [sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/sm_connection_public_ip/): complete subsection reference.

- [sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/sm_connection_pvt_ip/): complete subsection reference.

## Next pages

- [site_mesh_group_on_slo.no_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/no_site_mesh_group/)
- [site_mesh_group_on_slo.site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/site_mesh_group/)
- [site_mesh_group_on_slo.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/sm_connection_public_ip/)
- [site_mesh_group_on_slo.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/site_mesh_group_on_slo/sm_connection_pvt_ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
