---
page_title: "voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr"
subcategory: "Infrastructure"
description: "voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1833, "body_sha256": "sha256:57a543dc545ff97e2a0628dafc9e32c6f488ec06dbe4f677a5bac669e6b9a20a", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:global_network_list:global_network_connections:slo_to_global_dr", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:global_network_list:global_network_connections:slo_to_global_dr:global_vn"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:global_network_list:global_network_connections:slo_to_global_dr", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster_ar:global_network_list:global_network_connections", "path": "docs/guides/resources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster_ar", "global_network_list", "global_network_connections", "slo_to_global_dr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/slo_to_global_dr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [voltstack_cluster_ar](resources--azure_vnet_site--properties--voltstack_cluster_ar.md)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list.md)
- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections.md)
- voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

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
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

## Direct properties

- [global_vn](resources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn.md)
- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--properties--voltstack_cluster_ar--global_network_list--global_network_connections.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
