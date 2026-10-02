---
page_title: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface"
subcategory: ""
description: "Ethernet Interface Configuration."
xcsh_docs: {"aliases": ["custom storage config storage interface list storage interfaces storage interface"], "body_bytes": 17931, "body_sha256": "sha256:4a635f56a43711599350aaa9049e7441f0a09d4a71a6f5c39fe9b384ddd06b45", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:cluster", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_client", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:is_primary", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:monitor", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:monitor_disabled", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:no_ipv6_address", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:not_primary", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:site_local_inside_network", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:site_local_network", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ipv6_address", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:storage_network", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:untagged"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001", "registry_path": "docs/guides/resources--voltstack_site--reference--group-008.md", "relationships": [{"anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--node", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:cluster,node", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--vlan_id", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:untagged,vlan_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:cluster,node", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:dhcp_client,dhcp_server", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:dhcp_client,static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:dhcp_client,dhcp_server", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:dhcp_server,static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:ipv6_auto_config,no_ipv6_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:ipv6_auto_config,static_ipv6_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:is_primary,not_primary", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:is_primary", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:monitor,monitor_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:monitor", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:monitor,monitor_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:monitor_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:ipv6_auto_config,no_ipv6_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:no_ipv6_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:no_ipv6_address,static_ipv6_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:no_ipv6_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:is_primary,not_primary", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:not_primary", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:site_local_inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:site_local_inside_network,storage_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:site_local_inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:site_local_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:site_local_network,storage_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:site_local_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:dhcp_client,static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:dhcp_server,static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:ipv6_auto_config,static_ipv6_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ipv6_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:no_ipv6_address,static_ipv6_address", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ipv6_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:site_local_inside_network,storage_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:storage_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:site_local_network,storage_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:storage_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:ConflictingObjectAttributes:untagged,vlan_id", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:untagged", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--device", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface:RequiredObjectAttributes:device", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface"], "schema_version": 1, "sections": [{"aliases": ["cluster"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:cluster", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["device"], "anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--device", "description": "Interface configuration for the ethernet device.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "device"], "syntax": "attribute", "type": "string"}, {"aliases": ["dhcp client"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_client", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "dhcp_client"], "syntax": "attribute", "type": "object"}, {"aliases": ["dhcp server"], "anchor": "section", "description": "Configuration parameter for dhcp server.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server:ConflictingObjectAttributes:automatic_from_end,automatic_from_start", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:automatic_from_end", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server:ConflictingObjectAttributes:automatic_from_end,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:automatic_from_end", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server:ConflictingObjectAttributes:automatic_from_end,automatic_from_start", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:automatic_from_start", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server:ConflictingObjectAttributes:automatic_from_start,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:automatic_from_start", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server:ConflictingObjectAttributes:automatic_from_end,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:interface_ip_map", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server:ConflictingObjectAttributes:automatic_from_start,interface_ip_map", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:interface_ip_map", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server:RequiredObjectAttributes:dhcp_networks", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks", "type": "requires"}], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "dhcp_server"], "syntax": "block", "type": "object"}, {"aliases": ["ipv6 auto config"], "anchor": "section", "description": "IPV6AutoConfigType.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config:ConflictingObjectAttributes:host,router", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:host", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config:ConflictingObjectAttributes:host,router", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config:router", "type": "conflicts"}], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "ipv6_auto_config"], "syntax": "block", "type": "object"}, {"aliases": ["is primary"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:is_primary", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "is_primary"], "syntax": "attribute", "type": "object"}, {"aliases": ["monitor"], "anchor": "section", "description": "Link Quality Monitoring configuration for a network interface.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:monitor", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "monitor"], "syntax": "attribute", "type": "object"}, {"aliases": ["monitor disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:monitor_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "monitor_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["mtu"], "anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--mtu", "description": "Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between 512 and 9000.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "mtu"], "syntax": "attribute", "type": "number"}, {"aliases": ["no ipv6 address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:no_ipv6_address", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "no_ipv6_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["node"], "anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--node", "description": "Exclusive with Configuration will apply to a device on the given node.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "node"], "syntax": "attribute", "type": "string"}, {"aliases": ["not primary"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:not_primary", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "not_primary"], "syntax": "attribute", "type": "object"}, {"aliases": ["priority"], "anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--priority", "description": "Priority of the network interface when multiple network interfaces are present in outside network Greater the value, higher the priority.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "priority"], "syntax": "attribute", "type": "number"}, {"aliases": ["site local inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:site_local_inside_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "site_local_inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["site local network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:site_local_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "site_local_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["static ip"], "anchor": "section", "description": "Configure Static IP parameters.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip:cluster_static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip:node_static_ip", "type": "conflicts"}], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "static_ip"], "syntax": "block", "type": "object"}, {"aliases": ["static ipv6 address"], "anchor": "section", "description": "Configure Static IP parameters.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ipv6_address", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ipv6_address:cluster_static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ipv6_address:node_static_ip", "type": "conflicts"}], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "static_ipv6_address"], "syntax": "block", "type": "object"}, {"aliases": ["storage network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:storage_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "storage_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["untagged"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:untagged", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "untagged"], "syntax": "attribute", "type": "object"}, {"aliases": ["vlan id"], "anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--vlan_id", "description": "Exclusive with Configure a VLAN tagged ethernet interface.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "vlan_id"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Ethernet Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_interface_list.storage_interfaces.storage_interface

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/)
- [custom_storage_config.storage_interface_list.storage_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for storage interface.

Upstream description:

Ethernet Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node"),
  validators.ConflictingObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingObjectAttributes("is_primary",
    "not_primary"),
  validators.ConflictingObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "storage_network"),
  validators.ConflictingObjectAttributes("site_local_network",
    "storage_network"),
  validators.ConflictingObjectAttributes("untagged",
    "vlan_id")}
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
  "x-ves-oneof-field-address_choice": "[\"dhcp_client\",\"dhcp_server\",\"static_ip\"]",
  "x-ves-oneof-field-ipv6_address_choice": "[\"ipv6_auto_config\",\"no_ipv6_address\",\"static_ipv6_address\"]",
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\",\"storage_network\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]",
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

Terraform syntax:

```terraform
storage_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/cluster/): complete subsection reference.

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--device"></a>

### device property

Type: `"string"`. Optional.

Interface configuration for the ethernet device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [dhcp_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_client/): complete subsection reference.

- [dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/): complete subsection reference.

- [ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/ipv6_auto_config/): complete subsection reference.

- [is_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/is_primary/): complete subsection reference.

- [monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/monitor/): complete subsection reference.

- [monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/monitor_disabled/): complete subsection reference.

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--mtu"></a>

### mtu property

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

- [no_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/no_ipv6_address/): complete subsection reference.

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--node"></a>

### node property

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [not_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/not_primary/): complete subsection reference.

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--priority"></a>

### priority property

Type: `"number"`. Optional.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/site_local_inside_network/): complete subsection reference.

- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/site_local_network/): complete subsection reference.

- [static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/static_ip/): complete subsection reference.

- [static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/static_ipv6_address/): complete subsection reference.

- [storage_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/storage_network/): complete subsection reference.

- [untagged](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/untagged/): complete subsection reference.

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--vlan_id"></a>

### vlan_id property

Type: `"number"`. Optional.

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

## Next pages

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/cluster/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_client/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/ipv6_auto_config/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.is_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/is_primary/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/monitor/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/monitor_disabled/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.no_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/no_ipv6_address/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.not_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/not_primary/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/site_local_inside_network/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/site_local_network/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/static_ip/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/static_ipv6_address/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.storage_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/storage_network/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.untagged](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/untagged/)
- [custom_storage_config.storage_interface_list.storage_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
