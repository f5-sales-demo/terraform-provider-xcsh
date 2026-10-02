---
page_title: "voltstack_cluster"
subcategory: "Infrastructure"
description: "App Stack cluster of single interface GCP site."
xcsh_docs: {"aliases": ["voltstack cluster"], "body_bytes": 13803, "body_sha256": "sha256:570d99c59fe4b3a6ee6b1e386a96959a3444049f3c5b92947166fe724cea00d5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_enhanced_firewall_policies", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_network_policies", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:dc_cluster_group", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:default_storage", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:forward_proxy_allow_all", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:k8s_cluster", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_dc_cluster_group", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_forward_proxy", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_global_network", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_k8s_cluster", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_network_policy", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_outside_static_routes", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:sm_connection_public_ip", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:sm_connection_pvt_ip", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:storage_class_list"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster", "parent_id": "xcsh-docs:resources:gcp_vpc_site:reference", "path": "documentation/resources/gcp_vpc_site/properties/voltstack_cluster/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0030221310331323-0201022132301223-1111311321113231-1120211001321323-2310121021112311-0311123122202010-2003132323001230-3123103010121113", "registry_path": "docs/guides/resources--gcp_vpc_site--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:active_enhanced_firewall_policies,active_network_policies", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_enhanced_firewall_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:active_enhanced_firewall_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_enhanced_firewall_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:active_forward_proxy_policies,forward_proxy_allow_all", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:active_forward_proxy_policies,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:active_enhanced_firewall_policies,active_network_policies", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_network_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:active_network_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_network_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:dc_cluster_group,no_dc_cluster_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:dc_cluster_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:default_storage,storage_class_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:default_storage", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:active_forward_proxy_policies,forward_proxy_allow_all", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:forward_proxy_allow_all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:forward_proxy_allow_all,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:forward_proxy_allow_all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:global_network_list,no_global_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:k8s_cluster,no_k8s_cluster", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:k8s_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:dc_cluster_group,no_dc_cluster_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_dc_cluster_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:active_forward_proxy_policies,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_forward_proxy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:forward_proxy_allow_all,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_forward_proxy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:global_network_list,no_global_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_global_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:k8s_cluster,no_k8s_cluster", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_k8s_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:active_enhanced_firewall_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_network_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:active_network_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_network_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:no_outside_static_routes,outside_static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_outside_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:no_outside_static_routes,outside_static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:sm_connection_public_ip,sm_connection_pvt_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:sm_connection_public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:sm_connection_public_ip,sm_connection_pvt_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:sm_connection_pvt_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster:ConflictingObjectAttributes:default_storage,storage_class_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:storage_class_list", "type": "conflicts"}, {"anchor": "schema-voltstack_cluster--gcp_certified_hw", "enforcement": "provider-schema", "group": "voltstack_cluster:RequiredObjectAttributes:gcp_certified_hw,gcp_zone_names", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster", "type": "requires"}, {"anchor": "schema-voltstack_cluster--gcp_zone_names", "enforcement": "provider-schema", "group": "voltstack_cluster:RequiredObjectAttributes:gcp_certified_hw,gcp_zone_names", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster"], "schema_version": 1, "sections": [{"aliases": ["active enhanced firewall policies"], "anchor": "section", "description": "List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS available under firewall policies with an additional option for service insertion.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_enhanced_firewall_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.active_enhanced_firewall_policies:RequiredObjectAttributes:enhanced_firewall_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_enhanced_firewall_policies:enhanced_firewall_policies", "type": "requires"}], "schema_path": ["voltstack_cluster", "active_enhanced_firewall_policies"], "syntax": "block", "type": "object"}, {"aliases": ["active forward proxy policies"], "anchor": "section", "description": "Ordered List of Forward Proxy Policies active.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.active_forward_proxy_policies:RequiredObjectAttributes:forward_proxy_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies:forward_proxy_policies", "type": "requires"}], "schema_path": ["voltstack_cluster", "active_forward_proxy_policies"], "syntax": "block", "type": "object"}, {"aliases": ["active network policies"], "anchor": "section", "description": "List of firewall policy views.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_network_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.active_network_policies:RequiredObjectAttributes:network_policies", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_network_policies:network_policies", "type": "requires"}], "schema_path": ["voltstack_cluster", "active_network_policies"], "syntax": "block", "type": "object"}, {"aliases": ["dc cluster group"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:dc_cluster_group", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-voltstack_cluster--dc_cluster_group--name", "enforcement": "provider-schema", "group": "voltstack_cluster.dc_cluster_group:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:dc_cluster_group", "type": "requires"}], "schema_path": ["voltstack_cluster", "dc_cluster_group"], "syntax": "block", "type": "object"}, {"aliases": ["default storage"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:default_storage", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "default_storage"], "syntax": "attribute", "type": "object"}, {"aliases": ["forward proxy allow all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:forward_proxy_allow_all", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "forward_proxy_allow_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["gcp certified hw"], "anchor": "schema-voltstack_cluster--gcp_certified_hw", "description": "Name for GCP certified hardware.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "gcp_certified_hw"], "syntax": "attribute", "type": "string"}, {"aliases": ["gcp zone names"], "anchor": "schema-voltstack_cluster--gcp_zone_names", "description": "X-required List of zones when instances will be created, needs to match with region selected.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "gcp_zone_names"], "syntax": "attribute", "type": "list"}, {"aliases": ["global network list"], "anchor": "section", "description": "List of global network connections.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.global_network_list:RequiredObjectAttributes:global_network_connections", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list:global_network_connections", "type": "requires"}], "schema_path": ["voltstack_cluster", "global_network_list"], "syntax": "block", "type": "object"}, {"aliases": ["k8s cluster"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:k8s_cluster", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-voltstack_cluster--k8s_cluster--name", "enforcement": "provider-schema", "group": "voltstack_cluster.k8s_cluster:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:k8s_cluster", "type": "requires"}], "schema_path": ["voltstack_cluster", "k8s_cluster"], "syntax": "block", "type": "object"}, {"aliases": ["no dc cluster group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_dc_cluster_group", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "no_dc_cluster_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["no forward proxy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_forward_proxy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "no_forward_proxy"], "syntax": "attribute", "type": "object"}, {"aliases": ["no global network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_global_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "no_global_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["no k8s cluster"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_k8s_cluster", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "no_k8s_cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["no network policy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_network_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "no_network_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["no outside static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_outside_static_routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "no_outside_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["node number"], "anchor": "schema-voltstack_cluster--node_number", "description": "Number of main nodes to create, either 1 or 3.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "node_number"], "syntax": "attribute", "type": "number"}, {"aliases": ["outside static routes"], "anchor": "section", "description": "List of static routes.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.outside_static_routes:RequiredObjectAttributes:static_route_list", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes:static_route_list", "type": "requires"}], "schema_path": ["voltstack_cluster", "outside_static_routes"], "syntax": "block", "type": "object"}, {"aliases": ["site local network"], "anchor": "section", "description": "This defines choice about GCP VPC network for a view.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_network:ConflictingObjectAttributes:existing_network,new_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:existing_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_network:ConflictingObjectAttributes:existing_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:existing_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_network:ConflictingObjectAttributes:existing_network,new_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_network:ConflictingObjectAttributes:new_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_network:ConflictingObjectAttributes:existing_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network_autogenerate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_network:ConflictingObjectAttributes:new_network,new_network_autogenerate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network_autogenerate", "type": "conflicts"}], "schema_path": ["voltstack_cluster", "site_local_network"], "syntax": "block", "type": "object"}, {"aliases": ["site local subnet"], "anchor": "section", "description": "This defines choice about GCP VPC network for a view.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_subnet:ConflictingObjectAttributes:existing_subnet,new_subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet:existing_subnet", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "voltstack_cluster.site_local_subnet:ConflictingObjectAttributes:existing_subnet,new_subnet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet:new_subnet", "type": "conflicts"}], "schema_path": ["voltstack_cluster", "site_local_subnet"], "syntax": "block", "type": "object"}, {"aliases": ["sm connection public ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:sm_connection_public_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "sm_connection_public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["sm connection pvt ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:sm_connection_pvt_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "sm_connection_pvt_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage class list"], "anchor": "section", "description": "Add additional custom storage classes in Kubernetes for this site.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:storage_class_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "storage_class_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/voltstack_cluster/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "App Stack cluster of single interface GCP site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- voltstack_cluster

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

App Stack cluster of single interface GCP site.

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
  validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("default_storage",
    "storage_class_list"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("k8s_cluster",
    "no_k8s_cluster"),
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-k8s_cluster_choice": "[\"k8s_cluster\",\"no_k8s_cluster\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage\",\"storage_class_list\"]"
}
```

Terraform syntax:

```terraform
voltstack_cluster {
  # Configure direct properties listed below.
}
```

## Direct properties

- [active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/active_enhanced_firewall_policies/): complete subsection reference.

- [active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/): complete subsection reference.

- [active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/active_network_policies/): complete subsection reference.

- [dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/dc_cluster_group/): complete subsection reference.

- [default_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/default_storage/): complete subsection reference.

- [forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/forward_proxy_allow_all/): complete subsection reference.

<a id="schema-voltstack_cluster--gcp_certified_hw"></a>

### gcp_certified_hw property

Type: `"string"`. Optional.

\[Enum: gcp-byol-voltstack-combo\] GCP Certified Hardware. Name for GCP certified hardware. The only
possible value is \`gcp-byol-voltstack-combo\`.

Upstream description:

Name for GCP certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("gcp-byol-voltstack-combo"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "gcp-byol-voltstack-combo"
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
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="schema-voltstack_cluster--gcp_zone_names"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/global_network_list/): complete subsection reference.

- [k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/k8s_cluster/): complete subsection reference.

- [no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/no_dc_cluster_group/): complete subsection reference.

- [no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/no_forward_proxy/): complete subsection reference.

- [no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/no_global_network/): complete subsection reference.

- [no_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/no_k8s_cluster/): complete subsection reference.

- [no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/no_network_policy/): complete subsection reference.

- [no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/no_outside_static_routes/): complete subsection reference.

<a id="schema-voltstack_cluster--node_number"></a>

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

- [outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/outside_static_routes/): complete subsection reference.

- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/): complete subsection reference.

- [site_local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_subnet/): complete subsection reference.

- [sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/sm_connection_public_ip/): complete subsection reference.

- [sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/sm_connection_pvt_ip/): complete subsection reference.

- [storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/storage_class_list/): complete subsection reference.

## Next pages

- [voltstack_cluster.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/active_enhanced_firewall_policies/)
- [voltstack_cluster.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/)
- [voltstack_cluster.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/active_network_policies/)
- [voltstack_cluster.dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/dc_cluster_group/)
- [voltstack_cluster.default_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/default_storage/)
- [voltstack_cluster.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/forward_proxy_allow_all/)
- [voltstack_cluster.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/global_network_list/)
- [voltstack_cluster.k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/k8s_cluster/)
- [voltstack_cluster.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/no_dc_cluster_group/)
- [voltstack_cluster.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/no_forward_proxy/)
- [voltstack_cluster.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/no_global_network/)
- [voltstack_cluster.no_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/no_k8s_cluster/)
- [voltstack_cluster.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/no_network_policy/)
- [voltstack_cluster.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/no_outside_static_routes/)
- [voltstack_cluster.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/outside_static_routes/)
- [voltstack_cluster.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/)
- [voltstack_cluster.site_local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_subnet/)
- [voltstack_cluster.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/sm_connection_public_ip/)
- [voltstack_cluster.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/sm_connection_pvt_ip/)
- [voltstack_cluster.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/storage_class_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
