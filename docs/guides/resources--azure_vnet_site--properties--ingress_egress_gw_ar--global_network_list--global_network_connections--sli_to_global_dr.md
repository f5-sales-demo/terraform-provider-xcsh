---
page_title: "ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1833, "body_sha256": "sha256:8d15dbc618389d6933dfdda4f06de6f8e36619d8dbef4061adb04050d86a313a", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:global_network_list:global_network_connections:sli_to_global_dr", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:global_network_list:global_network_connections:sli_to_global_dr:global_vn"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:global_network_list:global_network_connections:sli_to_global_dr", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:global_network_list:global_network_connections", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections--sli_to_global_dr.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "global_network_list", "global_network_connections", "sli_to_global_dr"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/global_network_connections/sli_to_global_dr/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](resources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_egress_gw_ar.global_network_list](resources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list.md)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](resources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections.md)
- ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr

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

- [global_vn](resources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn.md)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](resources--azure_vnet_site--properties--ingress_egress_gw_ar--global_network_list--global_network_connections.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
