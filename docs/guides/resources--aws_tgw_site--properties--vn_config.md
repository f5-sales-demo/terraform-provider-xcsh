---
page_title: "vn_config"
subcategory: ""
description: "vn_config for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 5376, "body_sha256": "sha256:65e281d98380b4ec6962b89203c8b3297e4528eabd8e9a8979f09fa54bc2532d", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:allowed_vip_port_sli", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:dc_cluster_group_inside_vn", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:dc_cluster_group_outside_vn", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:inside_static_routes", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:no_dc_cluster_group", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:no_global_network", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:no_inside_static_routes", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:no_outside_static_routes", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:sm_connection_public_ip", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:sm_connection_pvt_ip"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config", "parent_id": "xcsh-docs:resources:aws_tgw_site:reference", "path": "docs/guides/resources--aws_tgw_site--properties--vn_config.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- vn_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Virtual Network Configuration. Virtual Network Configuration.

Upstream description:

Virtual Network Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "dc_cluster_group_outside_vn"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("dc_cluster_group_outside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("inside_static_routes",
    "no_inside_static_routes"),
  validators.ConflictingObjectAttributes("no_outside_static_routes",
    "outside_static_routes"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

Terraform syntax:

```terraform
vn_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [allowed_vip_port](resources--aws_tgw_site--properties--vn_config--allowed_vip_port.md): complete subsection reference.

- [allowed_vip_port_sli](resources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli.md): complete subsection reference.

- [dc_cluster_group_inside_vn](resources--aws_tgw_site--properties--vn_config--dc_cluster_group_inside_vn.md): complete subsection reference.

- [dc_cluster_group_outside_vn](resources--aws_tgw_site--properties--vn_config--dc_cluster_group_outside_vn.md): complete subsection reference.

- [global_network_list](resources--aws_tgw_site--properties--vn_config--global_network_list.md): complete subsection reference.

- [inside_static_routes](resources--aws_tgw_site--properties--vn_config--inside_static_routes.md): complete subsection reference.

- [no_dc_cluster_group](resources--aws_tgw_site--properties--vn_config--no_dc_cluster_group.md): complete subsection reference.

- [no_global_network](resources--aws_tgw_site--properties--vn_config--no_global_network.md): complete subsection reference.

- [no_inside_static_routes](resources--aws_tgw_site--properties--vn_config--no_inside_static_routes.md): complete subsection reference.

- [no_outside_static_routes](resources--aws_tgw_site--properties--vn_config--no_outside_static_routes.md): complete subsection reference.

- [outside_static_routes](resources--aws_tgw_site--properties--vn_config--outside_static_routes.md): complete subsection reference.

- [sm_connection_public_ip](resources--aws_tgw_site--properties--vn_config--sm_connection_public_ip.md): complete subsection reference.

- [sm_connection_pvt_ip](resources--aws_tgw_site--properties--vn_config--sm_connection_pvt_ip.md): complete subsection reference.

## Next pages

- [vn_config.allowed_vip_port](resources--aws_tgw_site--properties--vn_config--allowed_vip_port.md)
- [vn_config.allowed_vip_port_sli](resources--aws_tgw_site--properties--vn_config--allowed_vip_port_sli.md)
- [vn_config.dc_cluster_group_inside_vn](resources--aws_tgw_site--properties--vn_config--dc_cluster_group_inside_vn.md)
- [vn_config.dc_cluster_group_outside_vn](resources--aws_tgw_site--properties--vn_config--dc_cluster_group_outside_vn.md)
- [vn_config.global_network_list](resources--aws_tgw_site--properties--vn_config--global_network_list.md)
- [vn_config.inside_static_routes](resources--aws_tgw_site--properties--vn_config--inside_static_routes.md)
- [vn_config.no_dc_cluster_group](resources--aws_tgw_site--properties--vn_config--no_dc_cluster_group.md)
- [vn_config.no_global_network](resources--aws_tgw_site--properties--vn_config--no_global_network.md)
- [vn_config.no_inside_static_routes](resources--aws_tgw_site--properties--vn_config--no_inside_static_routes.md)
- [vn_config.no_outside_static_routes](resources--aws_tgw_site--properties--vn_config--no_outside_static_routes.md)
- [vn_config.outside_static_routes](resources--aws_tgw_site--properties--vn_config--outside_static_routes.md)
- [vn_config.sm_connection_public_ip](resources--aws_tgw_site--properties--vn_config--sm_connection_public_ip.md)
- [vn_config.sm_connection_pvt_ip](resources--aws_tgw_site--properties--vn_config--sm_connection_pvt_ip.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
