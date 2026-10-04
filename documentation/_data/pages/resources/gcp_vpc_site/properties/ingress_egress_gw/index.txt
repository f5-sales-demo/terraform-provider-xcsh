---
page_title: "ingress_egress_gw"
subcategory: "Infrastructure"
description: "Two interface GCP ingress/egress site."
xcsh_docs: {"aliases": ["ingress egress gw"], "body_bytes": 15433, "body_sha256": "sha256:bc2139c2c98d23a59f035c58cdd7b61a331ec50fad91d3402141d3c7085be14b", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_enhanced_firewall_policies", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_forward_proxy_policies", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_network_policies", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:dc_cluster_group_inside_vn", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:dc_cluster_group_outside_vn", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:forward_proxy_allow_all", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:global_network_list", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_network", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_dc_cluster_group", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_forward_proxy", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_global_network", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_inside_static_routes", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_network_policy", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_outside_static_routes", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_network", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:sm_connection_public_ip", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:sm_connection_pvt_ip"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw", "parent_id": "xcsh-docs:resources:gcp_vpc_site:reference", "path": "documentation/resources/gcp_vpc_site/properties/ingress_egress_gw/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3123100130121022-2313112013232222-0213021230302120-3230300230213031-3101322300112331-0011013220032322-1311222312201210-2213020302010100", "registry_path": "docs/guides/resources--gcp_vpc_site--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:active_enhanced_firewall_policies,active_network_policies", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_enhanced_firewall_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:active_enhanced_firewall_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_enhanced_firewall_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:active_forward_proxy_policies,forward_proxy_allow_all", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_forward_proxy_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:active_forward_proxy_policies,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_forward_proxy_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:active_enhanced_firewall_policies,active_network_policies", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_network_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:active_network_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_network_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:dc_cluster_group_inside_vn,dc_cluster_group_outside_vn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:dc_cluster_group_inside_vn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:dc_cluster_group_inside_vn,no_dc_cluster_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:dc_cluster_group_inside_vn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:dc_cluster_group_inside_vn,dc_cluster_group_outside_vn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:dc_cluster_group_outside_vn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:dc_cluster_group_outside_vn,no_dc_cluster_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:dc_cluster_group_outside_vn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:active_forward_proxy_policies,forward_proxy_allow_all", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:forward_proxy_allow_all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:forward_proxy_allow_all,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:forward_proxy_allow_all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:global_network_list,no_global_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:global_network_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:inside_static_routes,no_inside_static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:dc_cluster_group_inside_vn,no_dc_cluster_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_dc_cluster_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:dc_cluster_group_outside_vn,no_dc_cluster_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_dc_cluster_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:active_forward_proxy_policies,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_forward_proxy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:forward_proxy_allow_all,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_forward_proxy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:global_network_list,no_global_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_global_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:inside_static_routes,no_inside_static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_inside_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:active_enhanced_firewall_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_network_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:active_network_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_network_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:no_outside_static_routes,outside_static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_outside_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:no_outside_static_routes,outside_static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:sm_connection_public_ip,sm_connection_pvt_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:sm_connection_public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw:ConflictingObjectAttributes:sm_connection_public_ip,sm_connection_pvt_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:sm_connection_pvt_ip", "type": "conflicts"}, {"anchor": "schema-ingress_egress_gw--gcp_certified_hw", "enforcement": "provider-schema", "group": "ingress_egress_gw:RequiredObjectAttributes:gcp_certified_hw,gcp_zone_names", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw", "type": "requires"}, {"anchor": "schema-ingress_egress_gw--gcp_zone_names", "enforcement": "provider-schema", "group": "ingress_egress_gw:RequiredObjectAttributes:gcp_certified_hw,gcp_zone_names", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw"], "schema_version": 1, "sections": [{"aliases": ["ingress egress gw active enhanced firewall policies"], "anchor": "section", "description": "List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS available under firewall policies with an additional option for service insertion.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_enhanced_firewall_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.active_enhanced_firewall_policies:RequiredObjectAttributes:enhanced_firewall_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_enhanced_firewall_policies:enhanced_firewall_policies", "type": "requires"}], "schema_path": ["ingress_egress_gw", "active_enhanced_firewall_policies"], "syntax": "block", "type": "object"}, {"aliases": ["ingress egress gw active forward proxy policies"], "anchor": "section", "description": "Ordered List of Forward Proxy Policies active.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_forward_proxy_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.active_forward_proxy_policies:RequiredObjectAttributes:forward_proxy_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_forward_proxy_policies:forward_proxy_policies", "type": "requires"}], "schema_path": ["ingress_egress_gw", "active_forward_proxy_policies"], "syntax": "block", "type": "object"}, {"aliases": ["ingress egress gw active network policies"], "anchor": "section", "description": "List of firewall policy views.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_network_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.active_network_policies:RequiredObjectAttributes:network_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:active_network_policies:network_policies", "type": "requires"}], "schema_path": ["ingress_egress_gw", "active_network_policies"], "syntax": "block", "type": "object"}, {"aliases": ["ingress egress gw dc cluster group inside vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:dc_cluster_group_inside_vn", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw--dc_cluster_group_inside_vn--name", "enforcement": "provider-schema", "group": "ingress_egress_gw.dc_cluster_group_inside_vn:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:dc_cluster_group_inside_vn", "type": "requires"}], "schema_path": ["ingress_egress_gw", "dc_cluster_group_inside_vn"], "syntax": "block", "type": "object"}, {"aliases": ["ingress egress gw dc cluster group outside vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:dc_cluster_group_outside_vn", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-ingress_egress_gw--dc_cluster_group_outside_vn--name", "enforcement": "provider-schema", "group": "ingress_egress_gw.dc_cluster_group_outside_vn:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:dc_cluster_group_outside_vn", "type": "requires"}], "schema_path": ["ingress_egress_gw", "dc_cluster_group_outside_vn"], "syntax": "block", "type": "object"}, {"aliases": ["ingress egress gw forward proxy allow all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:forward_proxy_allow_all", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "forward_proxy_allow_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw gcp certified hw"], "anchor": "schema-ingress_egress_gw--gcp_certified_hw", "description": "Name for GCP certified hardware.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "gcp_certified_hw"], "syntax": "attribute", "type": "string"}, {"aliases": ["ingress egress gw gcp zone names"], "anchor": "schema-ingress_egress_gw--gcp_zone_names", "description": "X-required List of zones when instances will be created, needs to match with region selected.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "gcp_zone_names"], "syntax": "attribute", "type": "list"}, {"aliases": ["ingress egress gw global network list"], "anchor": "section", "description": "List of global network connections.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:global_network_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.global_network_list:RequiredObjectAttributes:global_network_connections", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:global_network_list:global_network_connections", "type": "requires"}], "schema_path": ["ingress_egress_gw", "global_network_list"], "syntax": "block", "type": "object"}, {"aliases": ["ingress egress gw inside network"], "anchor": "section", "description": "This defines choice about GCP VPC network for a view.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_network", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_network:ConflictingObjectAttributes:existing_network,new_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_network:existing_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_network:ConflictingObjectAttributes:existing_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_network:existing_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_network:ConflictingObjectAttributes:existing_network,new_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_network:new_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_network:ConflictingObjectAttributes:new_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_network:new_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_network:ConflictingObjectAttributes:existing_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_network:new_network_autogenerate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_network:ConflictingObjectAttributes:new_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_network:new_network_autogenerate", "type": "conflicts"}], "schema_path": ["ingress_egress_gw", "inside_network"], "syntax": "block", "type": "object"}, {"aliases": ["ingress egress gw inside static routes"], "anchor": "section", "description": "List of static routes.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_static_routes:RequiredObjectAttributes:static_route_list", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list", "type": "requires"}], "schema_path": ["ingress_egress_gw", "inside_static_routes"], "syntax": "block", "type": "object"}, {"aliases": ["ingress egress gw inside subnet"], "anchor": "section", "description": "This defines choice about GCP VPC network for a view.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_subnet:ConflictingObjectAttributes:existing_subnet,new_subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:existing_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.inside_subnet:ConflictingObjectAttributes:existing_subnet,new_subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:new_subnet", "type": "conflicts"}], "schema_path": ["ingress_egress_gw", "inside_subnet"], "syntax": "block", "type": "object"}, {"aliases": ["ingress egress gw no dc cluster group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_dc_cluster_group", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "no_dc_cluster_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw no forward proxy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_forward_proxy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "no_forward_proxy"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw no global network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_global_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "no_global_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw no inside static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_inside_static_routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "no_inside_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw no network policy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_network_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "no_network_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw no outside static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:no_outside_static_routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "no_outside_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw node number"], "anchor": "schema-ingress_egress_gw--node_number", "description": "Number of main nodes to create, either 1 or 3.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "node_number"], "syntax": "attribute", "type": "number"}, {"aliases": ["ingress egress gw outside network"], "anchor": "section", "description": "This defines choice about GCP VPC network for a view.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_network", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_network:ConflictingObjectAttributes:existing_network,new_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_network:existing_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_network:ConflictingObjectAttributes:existing_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_network:existing_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_network:ConflictingObjectAttributes:existing_network,new_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_network:new_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_network:ConflictingObjectAttributes:new_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_network:new_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_network:ConflictingObjectAttributes:existing_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_network:new_network_autogenerate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_network:ConflictingObjectAttributes:new_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_network:new_network_autogenerate", "type": "conflicts"}], "schema_path": ["ingress_egress_gw", "outside_network"], "syntax": "block", "type": "object"}, {"aliases": ["ingress egress gw outside static routes"], "anchor": "section", "description": "List of static routes.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_static_routes:RequiredObjectAttributes:static_route_list", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes:static_route_list", "type": "requires"}], "schema_path": ["ingress_egress_gw", "outside_static_routes"], "syntax": "block", "type": "object"}, {"aliases": ["ingress egress gw outside subnet"], "anchor": "section", "description": "This defines choice about GCP VPC network for a view.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_subnet:ConflictingObjectAttributes:existing_subnet,new_subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet:existing_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.outside_subnet:ConflictingObjectAttributes:existing_subnet,new_subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet:new_subnet", "type": "conflicts"}], "schema_path": ["ingress_egress_gw", "outside_subnet"], "syntax": "block", "type": "object"}, {"aliases": ["ingress egress gw performance enhancement mode"], "anchor": "section", "description": "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.performance_enhancement_mode:ConflictingObjectAttributes:perf_mode_l3_enhanced,perf_mode_l7_enhanced", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode:perf_mode_l3_enhanced", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_egress_gw.performance_enhancement_mode:ConflictingObjectAttributes:perf_mode_l3_enhanced,perf_mode_l7_enhanced", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode:perf_mode_l7_enhanced", "type": "conflicts"}], "schema_path": ["ingress_egress_gw", "performance_enhancement_mode"], "syntax": "block", "type": "object"}, {"aliases": ["ingress egress gw sm connection public ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:sm_connection_public_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "sm_connection_public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw sm connection pvt ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:sm_connection_pvt_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_egress_gw", "sm_connection_pvt_ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/ingress_egress_gw/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Two interface GCP ingress/egress site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- ingress_egress_gw

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ingress\_egress\_gw, ingress\_gw, voltstack\_cluster\] Configuration parameter for ingress
egress gw.

Upstream description:

Two interface GCP ingress/egress site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("gcp_certified_hw",
    "gcp_zone_names"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "dc_cluster_group_outside_vn"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("dc_cluster_group_outside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
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
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

OneOf alternatives in this subsection:

- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/#section)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_gw/#section)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ingress_egress_gw {
  # Configure direct properties listed below.
}
```

## Direct properties

- [active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/): complete subsection reference.

- [active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/active_forward_proxy_policies/): complete subsection reference.

- [active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/active_network_policies/): complete subsection reference.

- [dc_cluster_group_inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/dc_cluster_group_inside_vn/): complete subsection reference.

- [dc_cluster_group_outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/dc_cluster_group_outside_vn/): complete subsection reference.

- [forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/forward_proxy_allow_all/): complete subsection reference.

<a id="schema-ingress_egress_gw--gcp_certified_hw"></a>

### gcp_certified_hw property

Type: `"string"`. Optional.

\[Enum: gcp-byol-multi-nic-voltmesh\] GCP Certified Hardware. Name for GCP certified hardware. The
only possible value is \`gcp-byol-multi-nic-voltmesh\`.

Upstream description:

Name for GCP certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("gcp-byol-multi-nic-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "gcp-byol-multi-nic-voltmesh"
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="schema-ingress_egress_gw--gcp_zone_names"></a>

### gcp_zone_names property

Type: `["list", "string"]`. Optional.

X-required List of zones when instances will be created, needs to match with region selected.

Upstream description:

X-required List of zones when instances will be created, needs to match with region selected.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(3),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/global_network_list/): complete subsection reference.

- [inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_network/): complete subsection reference.

- [inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/): complete subsection reference.

- [inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_subnet/): complete subsection reference.

- [no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/no_dc_cluster_group/): complete subsection reference.

- [no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/no_forward_proxy/): complete subsection reference.

- [no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/no_global_network/): complete subsection reference.

- [no_inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/no_inside_static_routes/): complete subsection reference.

- [no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/no_network_policy/): complete subsection reference.

- [no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/no_outside_static_routes/): complete subsection reference.

<a id="schema-ingress_egress_gw--node_number"></a>

### node_number property

Type: `"number"`. Optional.

Number of main nodes to create, either 1 or 3.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

- [outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_network/): complete subsection reference.

- [outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/): complete subsection reference.

- [outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_subnet/): complete subsection reference.

- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/): complete subsection reference.

- [sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/sm_connection_public_ip/): complete subsection reference.

- [sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/sm_connection_pvt_ip/): complete subsection reference.

## Next pages

- [ingress_egress_gw.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/)
- [ingress_egress_gw.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/active_forward_proxy_policies/)
- [ingress_egress_gw.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/active_network_policies/)
- [ingress_egress_gw.dc_cluster_group_inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/dc_cluster_group_inside_vn/)
- [ingress_egress_gw.dc_cluster_group_outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/dc_cluster_group_outside_vn/)
- [ingress_egress_gw.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/forward_proxy_allow_all/)
- [ingress_egress_gw.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/global_network_list/)
- [ingress_egress_gw.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_network/)
- [ingress_egress_gw.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_static_routes/)
- [ingress_egress_gw.inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_subnet/)
- [ingress_egress_gw.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/no_dc_cluster_group/)
- [ingress_egress_gw.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/no_forward_proxy/)
- [ingress_egress_gw.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/no_global_network/)
- [ingress_egress_gw.no_inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/no_inside_static_routes/)
- [ingress_egress_gw.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/no_network_policy/)
- [ingress_egress_gw.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/no_outside_static_routes/)
- [ingress_egress_gw.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_network/)
- [ingress_egress_gw.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_static_routes/)
- [ingress_egress_gw.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_subnet/)
- [ingress_egress_gw.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/)
- [ingress_egress_gw.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/sm_connection_public_ip/)
- [ingress_egress_gw.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/ingress_egress_gw/sm_connection_pvt_ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
