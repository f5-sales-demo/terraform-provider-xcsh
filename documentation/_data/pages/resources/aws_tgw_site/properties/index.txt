---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": ["aws tgw site"], "body_bytes": 91654, "body_sha256": "sha256:4f2ccdc4c2841712200018d9c9a3da3b510273c245d3a4fd63c50ea2f7da3d9f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "xcsh-docs:resources:aws_tgw_site:properties:block_all_services", "xcsh-docs:resources:aws_tgw_site:properties:blocked_services", "xcsh-docs:resources:aws_tgw_site:properties:coordinates", "xcsh-docs:resources:aws_tgw_site:properties:custom_dns", "xcsh-docs:resources:aws_tgw_site:properties:default_blocked_services", "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_disabled", "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled", "xcsh-docs:resources:aws_tgw_site:properties:kubernetes_upgrade_drain", "xcsh-docs:resources:aws_tgw_site:properties:log_receiver", "xcsh-docs:resources:aws_tgw_site:properties:logs_streaming_disabled", "xcsh-docs:resources:aws_tgw_site:properties:offline_survivability_mode", "xcsh-docs:resources:aws_tgw_site:properties:os", "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode", "xcsh-docs:resources:aws_tgw_site:properties:private_connectivity", "xcsh-docs:resources:aws_tgw_site:properties:sw", "xcsh-docs:resources:aws_tgw_site:properties:tgw_security", "xcsh-docs:resources:aws_tgw_site:properties:timeouts", "xcsh-docs:resources:aws_tgw_site:properties:vn_config", "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments", "xcsh-docs:resources:aws_tgw_site:properties:waf_signatures"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:reference", "parent_id": "xcsh-docs:resources:aws_tgw_site:fundamentals", "path": "documentation/resources/aws_tgw_site/properties/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:aws_tgw_site:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["aws parameters"], "anchor": "section", "description": "Setup AWS services VPC, transit gateway and site.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_parameters--nodes_per_az", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:no_worker_nodes,nodes_per_az", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "conflicts"}, {"anchor": "schema-aws_parameters--nodes_per_az", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:nodes_per_az,total_nodes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "conflicts"}, {"anchor": "schema-aws_parameters--total_nodes", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:no_worker_nodes,total_nodes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "conflicts"}, {"anchor": "schema-aws_parameters--total_nodes", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:nodes_per_az,total_nodes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "conflicts"}, {"anchor": "schema-aws_parameters--vpc_id", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:new_vpc,vpc_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:custom_security_group,f5xc_security_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:custom_security_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:disable_encryption,enable_encryption", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:disable_encryption", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:disable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:disable_encryption,enable_encryption", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:enable_encryption", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:enable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:existing_tgw,new_tgw", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:existing_tgw", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:custom_security_group,f5xc_security_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:f5xc_security_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:existing_tgw,new_tgw", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_tgw", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:new_vpc,vpc_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:new_vpc", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:no_worker_nodes,nodes_per_az", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:no_worker_nodes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:no_worker_nodes,total_nodes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:no_worker_nodes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:reserved_tgw_cidr,tgw_cidr", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:reserved_tgw_cidr", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:ConflictingObjectAttributes:reserved_tgw_cidr,tgw_cidr", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:tgw_cidr", "type": "conflicts"}, {"anchor": "schema-aws_parameters--aws_region", "enforcement": "provider-schema", "group": "aws_parameters:RequiredObjectAttributes:aws_region,az_nodes,instance_type,ssh_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "requires"}, {"anchor": "schema-aws_parameters--instance_type", "enforcement": "provider-schema", "group": "aws_parameters:RequiredObjectAttributes:aws_region,az_nodes,instance_type,ssh_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "requires"}, {"anchor": "schema-aws_parameters--ssh_key", "enforcement": "provider-schema", "group": "aws_parameters:RequiredObjectAttributes:aws_region,az_nodes,instance_type,ssh_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_parameters:RequiredObjectAttributes:aws_region,az_nodes,instance_type,ssh_key", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:aws_parameters:az_nodes", "type": "requires"}], "schema_path": ["aws_parameters"], "syntax": "block", "type": "object"}, {"aliases": ["block all services"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:block_all_services", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["block_all_services"], "syntax": "attribute", "type": "object"}, {"aliases": ["blocked services"], "anchor": "section", "description": "Disable node local services on this site. Note: The chosen services will GET disabled on all nodes in the site.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:blocked_services", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["blocked_services"], "syntax": "block", "type": "object"}, {"aliases": ["coordinates"], "anchor": "section", "description": "Coordinates of the site which provides the site physical location.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:coordinates", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["coordinates"], "syntax": "block", "type": "object"}, {"aliases": ["custom dns"], "anchor": "section", "description": "Custom DNS is the configured for specify CE site.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:custom_dns", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_dns"], "syntax": "block", "type": "object"}, {"aliases": ["default blocked services"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:default_blocked_services", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_blocked_services"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:aws_tgw_site:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["direct connect disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["direct_connect_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["direct connect enabled"], "anchor": "section", "description": "Direct Connect Configuration.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-direct_connect_enabled--custom_asn", "enforcement": "provider-schema", "group": "direct_connect_enabled:ConflictingObjectAttributes:auto_asn,custom_asn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled:ConflictingObjectAttributes:auto_asn,custom_asn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:auto_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled:ConflictingObjectAttributes:hosted_vifs,standard_vifs", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled:ConflictingObjectAttributes:hosted_vifs,standard_vifs", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:standard_vifs", "type": "conflicts"}], "schema_path": ["direct_connect_enabled"], "syntax": "block", "type": "object"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:aws_tgw_site:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:aws_tgw_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["kubernetes upgrade drain"], "anchor": "section", "description": "Specify how worker nodes within a site will be upgraded.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:kubernetes_upgrade_drain", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "kubernetes_upgrade_drain:ConflictingObjectAttributes:disable_upgrade_drain,enable_upgrade_drain", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kubernetes_upgrade_drain:ConflictingObjectAttributes:disable_upgrade_drain,enable_upgrade_drain", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "type": "conflicts"}], "schema_path": ["kubernetes_upgrade_drain"], "syntax": "block", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:aws_tgw_site:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["log receiver"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:log_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-log_receiver--name", "enforcement": "provider-schema", "group": "log_receiver:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:log_receiver", "type": "requires"}], "schema_path": ["log_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["logs streaming disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:logs_streaming_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["logs_streaming_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:aws_tgw_site:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:aws_tgw_site:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["offline survivability mode"], "anchor": "section", "description": "Offline Survivability allows the Site to continue functioning normally without traffic loss during periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this feature is enabled, a site can continue to function as is with existing configuration for upto 7 days, even when the site is", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:offline_survivability_mode", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "offline_survivability_mode:ConflictingObjectAttributes:enable_offline_survivability_mode,no_offline_survivability_mode", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:offline_survivability_mode:enable_offline_survivability_mode", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "offline_survivability_mode:ConflictingObjectAttributes:enable_offline_survivability_mode,no_offline_survivability_mode", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:offline_survivability_mode:no_offline_survivability_mode", "type": "conflicts"}], "schema_path": ["offline_survivability_mode"], "syntax": "block", "type": "object"}, {"aliases": ["os"], "anchor": "section", "description": "Select the F5XC Operating System Version for the site. By default, latest available OS Version will be used. Refer to release notes to find required released OS versions.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:os", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-os--operating_system_version", "enforcement": "provider-schema", "group": "os:ConflictingObjectAttributes:default_os_version,operating_system_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:os", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "os:ConflictingObjectAttributes:default_os_version,operating_system_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:os:default_os_version", "type": "conflicts"}], "schema_path": ["os"], "syntax": "block", "type": "object"}, {"aliases": ["performance enhancement mode"], "anchor": "section", "description": "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "performance_enhancement_mode:ConflictingObjectAttributes:perf_mode_l3_enhanced,perf_mode_l7_enhanced", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "performance_enhancement_mode:ConflictingObjectAttributes:perf_mode_l3_enhanced,perf_mode_l7_enhanced", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "type": "conflicts"}], "schema_path": ["performance_enhancement_mode"], "syntax": "block", "type": "object"}, {"aliases": ["private connectivity"], "anchor": "section", "description": "Private Connect Configuration.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:private_connectivity", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "private_connectivity:ConflictingObjectAttributes:inside,outside", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:private_connectivity:inside", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "private_connectivity:ConflictingObjectAttributes:inside,outside", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:private_connectivity:outside", "type": "conflicts"}], "schema_path": ["private_connectivity"], "syntax": "block", "type": "object"}, {"aliases": ["sw"], "anchor": "section", "description": "Select the F5XC Software Version for the site. By default, latest available F5XC Software Version will be used. Refer to release notes to find required released SW versions.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:sw", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-sw--volterra_software_version", "enforcement": "provider-schema", "group": "sw:ConflictingObjectAttributes:default_sw_version,volterra_software_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:sw", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "sw:ConflictingObjectAttributes:default_sw_version,volterra_software_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:sw:default_sw_version", "type": "conflicts"}], "schema_path": ["sw"], "syntax": "block", "type": "object"}, {"aliases": ["tags"], "anchor": "schema-tags", "description": "AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify, organize, search for, and filter resources in AWS console.", "document_id": "xcsh-docs:resources:aws_tgw_site:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tags"], "syntax": "attribute", "type": "map"}, {"aliases": ["tgw security"], "anchor": "section", "description": "Security Configuration for transit gateway.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:active_east_west_service_policies,east_west_service_policy_allow_all", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:active_east_west_service_policies,no_east_west_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:active_enhanced_firewall_policies,active_network_policies", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_enhanced_firewall_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:active_enhanced_firewall_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_enhanced_firewall_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:active_forward_proxy_policies,forward_proxy_allow_all", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_forward_proxy_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:active_forward_proxy_policies,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_forward_proxy_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:active_enhanced_firewall_policies,active_network_policies", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_network_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:active_network_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_network_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:active_east_west_service_policies,east_west_service_policy_allow_all", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:east_west_service_policy_allow_all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:east_west_service_policy_allow_all,no_east_west_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:east_west_service_policy_allow_all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:active_forward_proxy_policies,forward_proxy_allow_all", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:forward_proxy_allow_all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:forward_proxy_allow_all,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:forward_proxy_allow_all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:active_east_west_service_policies,no_east_west_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_east_west_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:east_west_service_policy_allow_all,no_east_west_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_east_west_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:active_forward_proxy_policies,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_forward_proxy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:forward_proxy_allow_all,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_forward_proxy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:active_enhanced_firewall_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_network_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tgw_security:ConflictingObjectAttributes:active_network_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:no_network_policy", "type": "conflicts"}], "schema_path": ["tgw_security"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["vn config"], "anchor": "section", "description": "Virtual Network Configuration.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "vn_config:ConflictingObjectAttributes:dc_cluster_group_inside_vn,dc_cluster_group_outside_vn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:dc_cluster_group_inside_vn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config:ConflictingObjectAttributes:dc_cluster_group_inside_vn,no_dc_cluster_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:dc_cluster_group_inside_vn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config:ConflictingObjectAttributes:dc_cluster_group_inside_vn,dc_cluster_group_outside_vn", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:dc_cluster_group_outside_vn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config:ConflictingObjectAttributes:dc_cluster_group_outside_vn,no_dc_cluster_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:dc_cluster_group_outside_vn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config:ConflictingObjectAttributes:global_network_list,no_global_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config:ConflictingObjectAttributes:inside_static_routes,no_inside_static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:inside_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config:ConflictingObjectAttributes:dc_cluster_group_inside_vn,no_dc_cluster_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:no_dc_cluster_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config:ConflictingObjectAttributes:dc_cluster_group_outside_vn,no_dc_cluster_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:no_dc_cluster_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config:ConflictingObjectAttributes:global_network_list,no_global_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:no_global_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config:ConflictingObjectAttributes:inside_static_routes,no_inside_static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:no_inside_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config:ConflictingObjectAttributes:no_outside_static_routes,outside_static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:no_outside_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config:ConflictingObjectAttributes:no_outside_static_routes,outside_static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config:ConflictingObjectAttributes:sm_connection_public_ip,sm_connection_pvt_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:sm_connection_public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config:ConflictingObjectAttributes:sm_connection_public_ip,sm_connection_pvt_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:sm_connection_pvt_ip", "type": "conflicts"}], "schema_path": ["vn_config"], "syntax": "block", "type": "object"}, {"aliases": ["vpc attachments"], "anchor": "section", "description": "Spoke VPCs to be attached to the AWS TGW Site.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vpc_attachments"], "syntax": "block", "type": "object"}, {"aliases": ["waf signatures"], "anchor": "section", "description": "Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied manually. Refer to release notes for details about available Signatures update modes.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:waf_signatures", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "waf_signatures:ConflictingObjectAttributes:automatic,manual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:waf_signatures:automatic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_signatures:ConflictingObjectAttributes:automatic,manual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:waf_signatures:manual", "type": "conflicts"}], "schema_path": ["waf_signatures"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_aws_tgw_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/): complete subsection reference.

- [block_all_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/block_all_services/): complete subsection reference.

- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/blocked_services/): complete subsection reference.

- [coordinates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/coordinates/): complete subsection reference.

- [custom_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/custom_dns/): complete subsection reference.

- [default_blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/default_blocked_services/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [direct_connect_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_disabled/): complete subsection reference.

- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/): complete subsection reference.

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/kubernetes_upgrade_drain/): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

- [log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/log_receiver/): complete subsection reference.

- [logs_streaming_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/logs_streaming_disabled/): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the AWS TGW Site. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the AWS TGW Site is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/offline_survivability_mode/): complete subsection reference.

- [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/os/): complete subsection reference.

- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/): complete subsection reference.

- [private_connectivity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/): complete subsection reference.

- [sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/sw/): complete subsection reference.

<a id="schema-tags"></a>

### tags property

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{
  validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":40},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":127,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"127\",\"ves.io.schema.rules.map.max_pairs\":\"40\",\"ves.io.schema.rules.map.values.string.max_len\":\"255\"},\"values\":{\"maxLength\":255,\"type\":\"string\"}}"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 40
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 127,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "127",
      "ves.io.schema.rules.map.max_pairs": "40",
      "ves.io.schema.rules.map.values.string.max_len": "255"
    },
    "values": {
      "maxLength": 255,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

- [tgw_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/timeouts/): complete subsection reference.

- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/): complete subsection reference.

- [vpc_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vpc_attachments/): complete subsection reference.

- [waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/waf_signatures/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/#schema-annotations) |
| `aws_parameters` | [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/#section) |
| `aws_parameters.admin_password` | [aws_parameters.admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/#section) |
| `aws_parameters.admin_password.blindfold_secret_info` | [aws_parameters.admin_password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/blindfold_secret_info/#section) |
| `aws_parameters.admin_password.blindfold_secret_info.decryption_provider` | [aws_parameters.admin_password.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/blindfold_secret_info/#schema-aws_parameters--admin_password--blindfold_secret_info--decryption_provider) |
| `aws_parameters.admin_password.blindfold_secret_info.location` | [aws_parameters.admin_password.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/blindfold_secret_info/#schema-aws_parameters--admin_password--blindfold_secret_info--location) |
| `aws_parameters.admin_password.blindfold_secret_info.store_provider` | [aws_parameters.admin_password.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/blindfold_secret_info/#schema-aws_parameters--admin_password--blindfold_secret_info--store_provider) |
| `aws_parameters.admin_password.clear_secret_info` | [aws_parameters.admin_password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/clear_secret_info/#section) |
| `aws_parameters.admin_password.clear_secret_info.provider_ref` | [aws_parameters.admin_password.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/clear_secret_info/#schema-aws_parameters--admin_password--clear_secret_info--provider_ref) |
| `aws_parameters.admin_password.clear_secret_info.url` | [aws_parameters.admin_password.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/admin_password/clear_secret_info/#schema-aws_parameters--admin_password--clear_secret_info--url) |
| `aws_parameters.aws_cred` | [aws_parameters.aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/aws_cred/#section) |
| `aws_parameters.aws_cred.name` | [aws_parameters.aws_cred.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/aws_cred/#schema-aws_parameters--aws_cred--name) |
| `aws_parameters.aws_cred.namespace` | [aws_parameters.aws_cred.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/aws_cred/#schema-aws_parameters--aws_cred--namespace) |
| `aws_parameters.aws_cred.tenant` | [aws_parameters.aws_cred.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/aws_cred/#schema-aws_parameters--aws_cred--tenant) |
| `aws_parameters.aws_region` | [aws_parameters.aws_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/#schema-aws_parameters--aws_region) |
| `aws_parameters.az_nodes` | [aws_parameters.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/#section) |
| `aws_parameters.az_nodes.aws_az_name` | [aws_parameters.az_nodes.aws_az_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/#schema-aws_parameters--az_nodes--aws_az_name) |
| `aws_parameters.az_nodes.inside_subnet` | [aws_parameters.az_nodes.inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/inside_subnet/#section) |
| `aws_parameters.az_nodes.inside_subnet.existing_subnet_id` | [aws_parameters.az_nodes.inside_subnet.existing_subnet_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/inside_subnet/#schema-aws_parameters--az_nodes--inside_subnet--existing_subnet_id) |
| `aws_parameters.az_nodes.inside_subnet.subnet_param` | [aws_parameters.az_nodes.inside_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/inside_subnet/subnet_param/#section) |
| `aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.inside_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/inside_subnet/subnet_param/#schema-aws_parameters--az_nodes--inside_subnet--subnet_param--ipv4) |
| `aws_parameters.az_nodes.outside_subnet` | [aws_parameters.az_nodes.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/outside_subnet/#section) |
| `aws_parameters.az_nodes.outside_subnet.existing_subnet_id` | [aws_parameters.az_nodes.outside_subnet.existing_subnet_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/outside_subnet/#schema-aws_parameters--az_nodes--outside_subnet--existing_subnet_id) |
| `aws_parameters.az_nodes.outside_subnet.subnet_param` | [aws_parameters.az_nodes.outside_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/outside_subnet/subnet_param/#section) |
| `aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.outside_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/outside_subnet/subnet_param/#schema-aws_parameters--az_nodes--outside_subnet--subnet_param--ipv4) |
| `aws_parameters.az_nodes.reserved_inside_subnet` | [aws_parameters.az_nodes.reserved_inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/reserved_inside_subnet/#section) |
| `aws_parameters.az_nodes.workload_subnet` | [aws_parameters.az_nodes.workload_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/workload_subnet/#section) |
| `aws_parameters.az_nodes.workload_subnet.existing_subnet_id` | [aws_parameters.az_nodes.workload_subnet.existing_subnet_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/workload_subnet/#schema-aws_parameters--az_nodes--workload_subnet--existing_subnet_id) |
| `aws_parameters.az_nodes.workload_subnet.subnet_param` | [aws_parameters.az_nodes.workload_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/workload_subnet/subnet_param/#section) |
| `aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4` | [aws_parameters.az_nodes.workload_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/az_nodes/workload_subnet/subnet_param/#schema-aws_parameters--az_nodes--workload_subnet--subnet_param--ipv4) |
| `aws_parameters.custom_security_group` | [aws_parameters.custom_security_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/custom_security_group/#section) |
| `aws_parameters.custom_security_group.inside_security_group_id` | [aws_parameters.custom_security_group.inside_security_group_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/custom_security_group/#schema-aws_parameters--custom_security_group--inside_security_group_id) |
| `aws_parameters.custom_security_group.outside_security_group_id` | [aws_parameters.custom_security_group.outside_security_group_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/custom_security_group/#schema-aws_parameters--custom_security_group--outside_security_group_id) |
| `aws_parameters.disable_encryption` | [aws_parameters.disable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/disable_encryption/#section) |
| `aws_parameters.disable_internet_vip` | [aws_parameters.disable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/disable_internet_vip/#section) |
| `aws_parameters.disk_size` | [aws_parameters.disk_size](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/#schema-aws_parameters--disk_size) |
| `aws_parameters.enable_encryption` | [aws_parameters.enable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/enable_encryption/#section) |
| `aws_parameters.enable_encryption.kms_key_id` | [aws_parameters.enable_encryption.kms_key_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/enable_encryption/#schema-aws_parameters--enable_encryption--kms_key_id) |
| `aws_parameters.enable_internet_vip` | [aws_parameters.enable_internet_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/enable_internet_vip/#section) |
| `aws_parameters.existing_tgw` | [aws_parameters.existing_tgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/existing_tgw/#section) |
| `aws_parameters.existing_tgw.tgw_asn` | [aws_parameters.existing_tgw.tgw_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/existing_tgw/#schema-aws_parameters--existing_tgw--tgw_asn) |
| `aws_parameters.existing_tgw.tgw_id` | [aws_parameters.existing_tgw.tgw_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/existing_tgw/#schema-aws_parameters--existing_tgw--tgw_id) |
| `aws_parameters.existing_tgw.volterra_site_asn` | [aws_parameters.existing_tgw.volterra_site_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/existing_tgw/#schema-aws_parameters--existing_tgw--volterra_site_asn) |
| `aws_parameters.f5xc_security_group` | [aws_parameters.f5xc_security_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/f5xc_security_group/#section) |
| `aws_parameters.instance_type` | [aws_parameters.instance_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/#schema-aws_parameters--instance_type) |
| `aws_parameters.new_tgw` | [aws_parameters.new_tgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/#section) |
| `aws_parameters.new_tgw.system_generated` | [aws_parameters.new_tgw.system_generated](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/system_generated/#section) |
| `aws_parameters.new_tgw.user_assigned` | [aws_parameters.new_tgw.user_assigned](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/user_assigned/#section) |
| `aws_parameters.new_tgw.user_assigned.tgw_asn` | [aws_parameters.new_tgw.user_assigned.tgw_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/user_assigned/#schema-aws_parameters--new_tgw--user_assigned--tgw_asn) |
| `aws_parameters.new_tgw.user_assigned.volterra_site_asn` | [aws_parameters.new_tgw.user_assigned.volterra_site_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_tgw/user_assigned/#schema-aws_parameters--new_tgw--user_assigned--volterra_site_asn) |
| `aws_parameters.new_vpc` | [aws_parameters.new_vpc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_vpc/#section) |
| `aws_parameters.new_vpc.autogenerate` | [aws_parameters.new_vpc.autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_vpc/autogenerate/#section) |
| `aws_parameters.new_vpc.name_tag` | [aws_parameters.new_vpc.name_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_vpc/#schema-aws_parameters--new_vpc--name_tag) |
| `aws_parameters.new_vpc.primary_ipv4` | [aws_parameters.new_vpc.primary_ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/new_vpc/#schema-aws_parameters--new_vpc--primary_ipv4) |
| `aws_parameters.no_worker_nodes` | [aws_parameters.no_worker_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/no_worker_nodes/#section) |
| `aws_parameters.nodes_per_az` | [aws_parameters.nodes_per_az](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/#schema-aws_parameters--nodes_per_az) |
| `aws_parameters.reserved_tgw_cidr` | [aws_parameters.reserved_tgw_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/reserved_tgw_cidr/#section) |
| `aws_parameters.ssh_key` | [aws_parameters.ssh_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/#schema-aws_parameters--ssh_key) |
| `aws_parameters.tgw_cidr` | [aws_parameters.tgw_cidr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/tgw_cidr/#section) |
| `aws_parameters.tgw_cidr.ipv4` | [aws_parameters.tgw_cidr.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/tgw_cidr/#schema-aws_parameters--tgw_cidr--ipv4) |
| `aws_parameters.total_nodes` | [aws_parameters.total_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/#schema-aws_parameters--total_nodes) |
| `aws_parameters.vpc_id` | [aws_parameters.vpc_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/aws_parameters/#schema-aws_parameters--vpc_id) |
| `block_all_services` | [block_all_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/block_all_services/#section) |
| `blocked_services` | [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/blocked_services/#section) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/blocked_services/blocked_service/#section) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/blocked_services/blocked_service/dns/#section) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/blocked_services/blocked_service/#schema-blocked_services--blocked_service--network_type) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/blocked_services/blocked_service/ssh/#section) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/blocked_services/blocked_service/web_user_interface/#section) |
| `coordinates` | [coordinates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/coordinates/#section) |
| `coordinates.latitude` | [coordinates.latitude](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/coordinates/#schema-coordinates--latitude) |
| `coordinates.longitude` | [coordinates.longitude](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/coordinates/#schema-coordinates--longitude) |
| `custom_dns` | [custom_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/custom_dns/#section) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/custom_dns/#schema-custom_dns--inside_nameserver) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/custom_dns/#schema-custom_dns--outside_nameserver) |
| `default_blocked_services` | [default_blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/default_blocked_services/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/#schema-description) |
| `direct_connect_disabled` | [direct_connect_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_disabled/#section) |
| `direct_connect_enabled` | [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/#section) |
| `direct_connect_enabled.auto_asn` | [direct_connect_enabled.auto_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/auto_asn/#section) |
| `direct_connect_enabled.custom_asn` | [direct_connect_enabled.custom_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/#schema-direct_connect_enabled--custom_asn) |
| `direct_connect_enabled.hosted_vifs` | [direct_connect_enabled.hosted_vifs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/#section) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_direct_connect/#section) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_direct_connect/#schema-direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect--cloudlink_network_name) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_internet` | [direct_connect_enabled.hosted_vifs.site_registration_over_internet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_internet/#section) |
| `direct_connect_enabled.hosted_vifs.vif_list` | [direct_connect_enabled.hosted_vifs.vif_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/#section) |
| `direct_connect_enabled.hosted_vifs.vif_list.other_region` | [direct_connect_enabled.hosted_vifs.vif_list.other_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/#schema-direct_connect_enabled--hosted_vifs--vif_list--other_region) |
| `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` | [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/same_as_site_region/#section) |
| `direct_connect_enabled.hosted_vifs.vif_list.vif_id` | [direct_connect_enabled.hosted_vifs.vif_list.vif_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/#schema-direct_connect_enabled--hosted_vifs--vif_list--vif_id) |
| `direct_connect_enabled.standard_vifs` | [direct_connect_enabled.standard_vifs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/standard_vifs/#section) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/#schema-id) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/kubernetes_upgrade_drain/#section) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/kubernetes_upgrade_drain/disable_upgrade_drain/#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/disable_vega_upgrade_mode/#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/enable_vega_upgrade_mode/#section) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/#schema-labels) |
| `log_receiver` | [log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/log_receiver/#section) |
| `log_receiver.name` | [log_receiver.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/log_receiver/#schema-log_receiver--name) |
| `log_receiver.namespace` | [log_receiver.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/log_receiver/#schema-log_receiver--namespace) |
| `log_receiver.tenant` | [log_receiver.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/log_receiver/#schema-log_receiver--tenant) |
| `logs_streaming_disabled` | [logs_streaming_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/logs_streaming_disabled/#section) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/#schema-namespace) |
| `offline_survivability_mode` | [offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/offline_survivability_mode/#section) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/offline_survivability_mode/enable_offline_survivability_mode/#section) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/offline_survivability_mode/no_offline_survivability_mode/#section) |
| `os` | [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/os/#section) |
| `os.default_os_version` | [os.default_os_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/os/default_os_version/#section) |
| `os.operating_system_version` | [os.operating_system_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/os/#schema-os--operating_system_version) |
| `performance_enhancement_mode` | [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l3_enhanced/#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l3_enhanced/jumbo/#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l3_enhanced/no_jumbo/#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/#section) |
| `private_connectivity` | [private_connectivity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/#section) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/cloud_link/#section) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/cloud_link/#schema-private_connectivity--cloud_link--name) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/cloud_link/#schema-private_connectivity--cloud_link--namespace) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/cloud_link/#schema-private_connectivity--cloud_link--tenant) |
| `private_connectivity.inside` | [private_connectivity.inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/inside/#section) |
| `private_connectivity.outside` | [private_connectivity.outside](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/outside/#section) |
| `sw` | [sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/sw/#section) |
| `sw.default_sw_version` | [sw.default_sw_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/sw/default_sw_version/#section) |
| `sw.volterra_software_version` | [sw.volterra_software_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/sw/#schema-sw--volterra_software_version) |
| `tags` | [tags](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/#schema-tags) |
| `tgw_security` | [tgw_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/#section) |
| `tgw_security.active_east_west_service_policies` | [tgw_security.active_east_west_service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/#section) |
| `tgw_security.active_east_west_service_policies.service_policies` | [tgw_security.active_east_west_service_policies.service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/service_policies/#section) |
| `tgw_security.active_east_west_service_policies.service_policies.name` | [tgw_security.active_east_west_service_policies.service_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/service_policies/#schema-tgw_security--active_east_west_service_policies--service_policies--name) |
| `tgw_security.active_east_west_service_policies.service_policies.namespace` | [tgw_security.active_east_west_service_policies.service_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/service_policies/#schema-tgw_security--active_east_west_service_policies--service_policies--namespace) |
| `tgw_security.active_east_west_service_policies.service_policies.tenant` | [tgw_security.active_east_west_service_policies.service_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/service_policies/#schema-tgw_security--active_east_west_service_policies--service_policies--tenant) |
| `tgw_security.active_enhanced_firewall_policies` | [tgw_security.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_enhanced_firewall_policies/#section) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_enhanced_firewall_policies/enhanced_firewall_policies/#section) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-tgw_security--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `tgw_security.active_forward_proxy_policies` | [tgw_security.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_forward_proxy_policies/#section) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_forward_proxy_policies/forward_proxy_policies/#section) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.name` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_forward_proxy_policies/forward_proxy_policies/#schema-tgw_security--active_forward_proxy_policies--forward_proxy_policies--name) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_forward_proxy_policies/forward_proxy_policies/#schema-tgw_security--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant` | [tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_forward_proxy_policies/forward_proxy_policies/#schema-tgw_security--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `tgw_security.active_network_policies` | [tgw_security.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_network_policies/#section) |
| `tgw_security.active_network_policies.network_policies` | [tgw_security.active_network_policies.network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_network_policies/network_policies/#section) |
| `tgw_security.active_network_policies.network_policies.name` | [tgw_security.active_network_policies.network_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_network_policies/network_policies/#schema-tgw_security--active_network_policies--network_policies--name) |
| `tgw_security.active_network_policies.network_policies.namespace` | [tgw_security.active_network_policies.network_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_network_policies/network_policies/#schema-tgw_security--active_network_policies--network_policies--namespace) |
| `tgw_security.active_network_policies.network_policies.tenant` | [tgw_security.active_network_policies.network_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_network_policies/network_policies/#schema-tgw_security--active_network_policies--network_policies--tenant) |
| `tgw_security.east_west_service_policy_allow_all` | [tgw_security.east_west_service_policy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/east_west_service_policy_allow_all/#section) |
| `tgw_security.forward_proxy_allow_all` | [tgw_security.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/forward_proxy_allow_all/#section) |
| `tgw_security.no_east_west_policy` | [tgw_security.no_east_west_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/no_east_west_policy/#section) |
| `tgw_security.no_forward_proxy` | [tgw_security.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/no_forward_proxy/#section) |
| `tgw_security.no_network_policy` | [tgw_security.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/no_network_policy/#section) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/timeouts/#schema-timeouts--update) |
| `vn_config` | [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/#section) |
| `vn_config.allowed_vip_port` | [vn_config.allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/#section) |
| `vn_config.allowed_vip_port.custom_ports` | [vn_config.allowed_vip_port.custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/custom_ports/#section) |
| `vn_config.allowed_vip_port.custom_ports.port_ranges` | [vn_config.allowed_vip_port.custom_ports.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/custom_ports/#schema-vn_config--allowed_vip_port--custom_ports--port_ranges) |
| `vn_config.allowed_vip_port.disable_allowed_vip_port` | [vn_config.allowed_vip_port.disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/disable_allowed_vip_port/#section) |
| `vn_config.allowed_vip_port.use_http_https_port` | [vn_config.allowed_vip_port.use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/use_http_https_port/#section) |
| `vn_config.allowed_vip_port.use_http_port` | [vn_config.allowed_vip_port.use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/use_http_port/#section) |
| `vn_config.allowed_vip_port.use_https_port` | [vn_config.allowed_vip_port.use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port/use_https_port/#section) |
| `vn_config.allowed_vip_port_sli` | [vn_config.allowed_vip_port_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/#section) |
| `vn_config.allowed_vip_port_sli.custom_ports` | [vn_config.allowed_vip_port_sli.custom_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/custom_ports/#section) |
| `vn_config.allowed_vip_port_sli.custom_ports.port_ranges` | [vn_config.allowed_vip_port_sli.custom_ports.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/custom_ports/#schema-vn_config--allowed_vip_port_sli--custom_ports--port_ranges) |
| `vn_config.allowed_vip_port_sli.disable_allowed_vip_port` | [vn_config.allowed_vip_port_sli.disable_allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/disable_allowed_vip_port/#section) |
| `vn_config.allowed_vip_port_sli.use_http_https_port` | [vn_config.allowed_vip_port_sli.use_http_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_http_https_port/#section) |
| `vn_config.allowed_vip_port_sli.use_http_port` | [vn_config.allowed_vip_port_sli.use_http_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_http_port/#section) |
| `vn_config.allowed_vip_port_sli.use_https_port` | [vn_config.allowed_vip_port_sli.use_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/allowed_vip_port_sli/use_https_port/#section) |
| `vn_config.dc_cluster_group_inside_vn` | [vn_config.dc_cluster_group_inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/dc_cluster_group_inside_vn/#section) |
| `vn_config.dc_cluster_group_inside_vn.name` | [vn_config.dc_cluster_group_inside_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/dc_cluster_group_inside_vn/#schema-vn_config--dc_cluster_group_inside_vn--name) |
| `vn_config.dc_cluster_group_inside_vn.namespace` | [vn_config.dc_cluster_group_inside_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/dc_cluster_group_inside_vn/#schema-vn_config--dc_cluster_group_inside_vn--namespace) |
| `vn_config.dc_cluster_group_inside_vn.tenant` | [vn_config.dc_cluster_group_inside_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/dc_cluster_group_inside_vn/#schema-vn_config--dc_cluster_group_inside_vn--tenant) |
| `vn_config.dc_cluster_group_outside_vn` | [vn_config.dc_cluster_group_outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/dc_cluster_group_outside_vn/#section) |
| `vn_config.dc_cluster_group_outside_vn.name` | [vn_config.dc_cluster_group_outside_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/dc_cluster_group_outside_vn/#schema-vn_config--dc_cluster_group_outside_vn--name) |
| `vn_config.dc_cluster_group_outside_vn.namespace` | [vn_config.dc_cluster_group_outside_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/dc_cluster_group_outside_vn/#schema-vn_config--dc_cluster_group_outside_vn--namespace) |
| `vn_config.dc_cluster_group_outside_vn.tenant` | [vn_config.dc_cluster_group_outside_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/dc_cluster_group_outside_vn/#schema-vn_config--dc_cluster_group_outside_vn--tenant) |
| `vn_config.global_network_list` | [vn_config.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/#section) |
| `vn_config.global_network_list.global_network_connections` | [vn_config.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/#section) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/sli_to_global_dr/#section) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#section) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-vn_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/slo_to_global_dr/#section) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#section) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-vn_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `vn_config.inside_static_routes` | [vn_config.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/#section) |
| `vn_config.inside_static_routes.static_route_list` | [vn_config.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route` | [vn_config.inside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.attrs` | [vn_config.inside_static_routes.static_route_list.custom_static_route.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--attrs) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.labels` | [vn_config.inside_static_routes.static_route_list.custom_static_route.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/labels/#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/nexthop/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/subnets/#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#section) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-vn_config--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `vn_config.inside_static_routes.static_route_list.simple_static_route` | [vn_config.inside_static_routes.static_route_list.simple_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/inside_static_routes/static_route_list/#schema-vn_config--inside_static_routes--static_route_list--simple_static_route) |
| `vn_config.no_dc_cluster_group` | [vn_config.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/no_dc_cluster_group/#section) |
| `vn_config.no_global_network` | [vn_config.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/no_global_network/#section) |
| `vn_config.no_inside_static_routes` | [vn_config.no_inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/no_inside_static_routes/#section) |
| `vn_config.no_outside_static_routes` | [vn_config.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/no_outside_static_routes/#section) |
| `vn_config.outside_static_routes` | [vn_config.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/#section) |
| `vn_config.outside_static_routes.static_route_list` | [vn_config.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route` | [vn_config.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.attrs` | [vn_config.outside_static_routes.static_route_list.custom_static_route.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.labels` | [vn_config.outside_static_routes.static_route_list.custom_static_route.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/labels/#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/nexthop/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/subnets/#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#section) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-vn_config--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `vn_config.outside_static_routes.static_route_list.simple_static_route` | [vn_config.outside_static_routes.static_route_list.simple_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/#schema-vn_config--outside_static_routes--static_route_list--simple_static_route) |
| `vn_config.sm_connection_public_ip` | [vn_config.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/sm_connection_public_ip/#section) |
| `vn_config.sm_connection_pvt_ip` | [vn_config.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/sm_connection_pvt_ip/#section) |
| `vpc_attachments` | [vpc_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vpc_attachments/#section) |
| `vpc_attachments.vpc_list` | [vpc_attachments.vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vpc_attachments/vpc_list/#section) |
| `vpc_attachments.vpc_list.labels` | [vpc_attachments.vpc_list.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vpc_attachments/vpc_list/labels/#section) |
| `vpc_attachments.vpc_list.vpc_id` | [vpc_attachments.vpc_list.vpc_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vpc_attachments/vpc_list/#schema-vpc_attachments--vpc_list--vpc_id) |
| `waf_signatures` | [waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/waf_signatures/#section) |
| `waf_signatures.automatic` | [waf_signatures.automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/waf_signatures/automatic/#section) |
| `waf_signatures.manual` | [waf_signatures.manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/waf_signatures/manual/#section) |
