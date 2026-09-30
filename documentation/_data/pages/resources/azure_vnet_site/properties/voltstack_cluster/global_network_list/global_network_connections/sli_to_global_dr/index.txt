---
page_title: "voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr"
subcategory: "Infrastructure"
description: "voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2242, "body_sha256": "sha256:3f85cda98bcb0876b2454b40331730e594ad24ba90429be05df9807bf58391be", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:global_network_list:global_network_connections:sli_to_global_dr:global_vn"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:global_network_list:global_network_connections:sli_to_global_dr", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:voltstack_cluster:global_network_list:global_network_connections", "path": "documentation/resources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["voltstack_cluster", "global_network_list", "global_network_connections", "sli_to_global_dr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/)
- [voltstack_cluster.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/global_network_list/)
- [voltstack_cluster.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr

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
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

## Direct properties

- [global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/global_vn/): complete subsection reference.

## Next pages

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/global_vn/)
- [voltstack_cluster.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
