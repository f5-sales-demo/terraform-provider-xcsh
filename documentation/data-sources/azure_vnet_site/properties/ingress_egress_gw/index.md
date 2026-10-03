---
page_title: "ingress_egress_gw"
subcategory: "Infrastructure"
description: "Two interface Azure ingress/egress site."
xcsh_docs: {"aliases": ["ingress egress gw"], "body_bytes": 12543, "body_sha256": "sha256:c43abccf1493f3914deae12a570ff7bc02924552b3b0d1ef3f8935cf7f6fc046", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_enhanced_firewall_policies", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_network_policies", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:dc_cluster_group_inside_vn", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:dc_cluster_group_outside_vn", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:forward_proxy_allow_all", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:global_network_list", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_dc_cluster_group", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_forward_proxy", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_global_network", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_inside_static_routes", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_network_policy", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_outside_static_routes", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:not_hub", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:outside_static_routes", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:performance_enhancement_mode", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:sm_connection_public_ip", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:sm_connection_pvt_ip"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1202233120032131-2011321120221120-3331021002331300-1032320232233133-3331232223020020-0112003120021030-1322232221203323-2122003022330112", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw"], "schema_version": 1, "sections": [{"aliases": ["ingress egress gw accelerated networking"], "anchor": "section", "description": "Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "accelerated_networking"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw active enhanced firewall policies"], "anchor": "section", "description": "List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS available under firewall policies with an additional option for service insertion.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_enhanced_firewall_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "active_enhanced_firewall_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw active forward proxy policies"], "anchor": "section", "description": "Ordered List of Forward Proxy Policies active.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "active_forward_proxy_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw active network policies"], "anchor": "section", "description": "List of firewall policy views.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:active_network_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "active_network_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw az nodes"], "anchor": "section", "description": "Only Single AZ or Three AZ(s) nodes are supported currently.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ingress_egress_gw", "az_nodes"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw azure certified hw"], "anchor": "schema-ingress_egress_gw--azure_certified_hw", "description": "Name for Azure certified hardware.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "azure_certified_hw"], "syntax": "attribute", "type": "string"}, {"aliases": ["ingress egress gw dc cluster group inside vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:dc_cluster_group_inside_vn", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "dc_cluster_group_inside_vn"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw dc cluster group outside vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:dc_cluster_group_outside_vn", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "dc_cluster_group_outside_vn"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw forward proxy allow all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:forward_proxy_allow_all", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "forward_proxy_allow_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw global network list"], "anchor": "section", "description": "List of global network connections.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:global_network_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "global_network_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw hub"], "anchor": "section", "description": "Hub VNet type.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "hub"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw inside static routes"], "anchor": "section", "description": "List of static routes.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "inside_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw no dc cluster group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_dc_cluster_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "no_dc_cluster_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw no forward proxy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_forward_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "no_forward_proxy"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw no global network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_global_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "no_global_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw no inside static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_inside_static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "no_inside_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw no network policy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_network_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "no_network_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw no outside static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:no_outside_static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "no_outside_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw not hub"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:not_hub", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "not_hub"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw outside static routes"], "anchor": "section", "description": "List of static routes.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:outside_static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "outside_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw performance enhancement mode"], "anchor": "section", "description": "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:performance_enhancement_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "performance_enhancement_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw sm connection public ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:sm_connection_public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "sm_connection_public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw sm connection pvt ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:sm_connection_pvt_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "sm_connection_pvt_ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Two interface Azure ingress/egress site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
