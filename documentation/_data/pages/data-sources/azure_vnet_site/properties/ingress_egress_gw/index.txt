---
page_title: "ingress_egress_gw"
subcategory: "Infrastructure"
description: "ingress_egress_gw for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 12444, "body_sha256": "sha256:c766956db67d349770942d043d0a9f3dc3b7bbabf49768f5315d62011deb24a5", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_enhanced_firewall_policies", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_network_policies", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:dc_cluster_group_inside_vn", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:dc_cluster_group_outside_vn", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:forward_proxy_allow_all", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:global_network_list", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_dc_cluster_group", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_forward_proxy", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_global_network", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_inside_static_routes", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_network_policy", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_outside_static_routes", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:not_hub", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:outside_static_routes", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:performance_enhancement_mode", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:sm_connection_public_ip", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:sm_connection_pvt_ip"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["ingress_egress_gw"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- ingress_egress_gw

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: ingress\_egress\_gw, ingress\_egress\_gw\_ar, ingress\_gw, ingress\_gw\_ar,
voltstack\_cluster, voltstack\_cluster\_ar\] Two interface Azure ingress/egress site.

Upstream description:

Two interface Azure ingress/egress site.

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
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-hub_choice": "[\"hub\",\"not_hub\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

OneOf alternatives in this subsection:

- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/#section)
- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/#section)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/#section)
- [ingress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/#section)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/#section)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/): complete subsection reference.

- [active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/): complete subsection reference.

- [active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_forward_proxy_policies/): complete subsection reference.

- [active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_network_policies/): complete subsection reference.

- [az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/): complete subsection reference.

<a id="schema-ingress_egress_gw--azure_certified_hw"></a>

### azure_certified_hw property

Type: `"string"`. Computed.

\[Enum: azure-byol-multi-nic-voltmesh\] Azure Certified Hardware. Name for Azure certified hardware.
The only possible value is \`azure-byol-multi-nic-voltmesh\`.

Upstream description:

Name for Azure certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-multi-nic-voltmesh"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [dc_cluster_group_inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/dc_cluster_group_inside_vn/): complete subsection reference.

- [dc_cluster_group_outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/dc_cluster_group_outside_vn/): complete subsection reference.

- [forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/forward_proxy_allow_all/): complete subsection reference.

- [global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/): complete subsection reference.

- [hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/): complete subsection reference.

- [inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/): complete subsection reference.

- [no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_dc_cluster_group/): complete subsection reference.

- [no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_forward_proxy/): complete subsection reference.

- [no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_global_network/): complete subsection reference.

- [no_inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_inside_static_routes/): complete subsection reference.

- [no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_network_policy/): complete subsection reference.

- [no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_outside_static_routes/): complete subsection reference.

- [not_hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/not_hub/): complete subsection reference.

- [outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/): complete subsection reference.

- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/performance_enhancement_mode/): complete subsection reference.

- [sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/sm_connection_public_ip/): complete subsection reference.

- [sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/sm_connection_pvt_ip/): complete subsection reference.

## Next pages

- [ingress_egress_gw.accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/)
- [ingress_egress_gw.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/)
- [ingress_egress_gw.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_forward_proxy_policies/)
- [ingress_egress_gw.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_network_policies/)
- [ingress_egress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/)
- [ingress_egress_gw.dc_cluster_group_inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/dc_cluster_group_inside_vn/)
- [ingress_egress_gw.dc_cluster_group_outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/dc_cluster_group_outside_vn/)
- [ingress_egress_gw.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/forward_proxy_allow_all/)
- [ingress_egress_gw.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/)
- [ingress_egress_gw.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/)
- [ingress_egress_gw.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/)
- [ingress_egress_gw.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_dc_cluster_group/)
- [ingress_egress_gw.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_forward_proxy/)
- [ingress_egress_gw.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_global_network/)
- [ingress_egress_gw.no_inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_inside_static_routes/)
- [ingress_egress_gw.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_network_policy/)
- [ingress_egress_gw.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_outside_static_routes/)
- [ingress_egress_gw.not_hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/not_hub/)
- [ingress_egress_gw.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/)
- [ingress_egress_gw.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/performance_enhancement_mode/)
- [ingress_egress_gw.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/sm_connection_public_ip/)
- [ingress_egress_gw.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/sm_connection_pvt_ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
