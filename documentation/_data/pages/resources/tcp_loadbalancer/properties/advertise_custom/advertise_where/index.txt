---
page_title: "advertise_custom.advertise_where"
subcategory: "Load Balancing"
description: "Where should this load balancer be available."
xcsh_docs: {"aliases": ["advertise custom advertise where"], "body_bytes": 11321, "body_sha256": "sha256:e1b69abcbc928cfc51f44682d39d0bbf7ccf1de53d14e98cdbc166eb33de9566", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_dualstack_on_public", "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_on_public", "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:site", "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:use_default_port", "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site", "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip", "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom", "path": "documentation/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2220031130203102-2210011003013002-3311212001302200-0330220102220213-0312033302222203-3301311320102322-2002200311110023-0013020221203332", "registry_path": "docs/guides/resources--tcp_loadbalancer--reference--group-001.md", "relationships": [{"anchor": "schema-advertise_custom--advertise_where--port", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where", "type": "conflicts"}, {"anchor": "schema-advertise_custom--advertise_where--port", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:port,use_default_port", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where", "type": "conflicts"}, {"anchor": "schema-advertise_custom--advertise_where--port_ranges", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where", "type": "conflicts"}, {"anchor": "schema-advertise_custom--advertise_where--port_ranges", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:port_ranges,use_default_port", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_dualstack_on_public,advertise_on_public", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_dualstack_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_dualstack_on_public,advertise_v6_on_public", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_dualstack_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_dualstack_on_public,site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_dualstack_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_dualstack_on_public,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_dualstack_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_dualstack_on_public,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_dualstack_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_dualstack_on_public,virtual_site_with_vip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_dualstack_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_dualstack_on_public,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_dualstack_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_dualstack_on_public,advertise_on_public", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_on_public,advertise_v6_on_public", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_on_public,site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_on_public,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_on_public,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_on_public,virtual_site_with_vip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_on_public,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_dualstack_on_public,advertise_v6_on_public", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_on_public,advertise_v6_on_public", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_v6_on_public,site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_v6_on_public,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_v6_on_public,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_v6_on_public,virtual_site_with_vip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_v6_on_public,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_dualstack_on_public,site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_on_public,site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_v6_on_public,site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:site,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:site,virtual_site_with_vip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:port,use_default_port", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:use_default_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:port_ranges,use_default_port", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:use_default_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_dualstack_on_public,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_on_public,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_v6_on_public,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:site,virtual_network", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_network,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_network,virtual_site_with_vip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_network,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_dualstack_on_public,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_on_public,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_v6_on_public,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_network,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_site,virtual_site_with_vip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_dualstack_on_public,virtual_site_with_vip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_on_public,virtual_site_with_vip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_v6_on_public,virtual_site_with_vip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:site,virtual_site_with_vip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_network,virtual_site_with_vip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_site,virtual_site_with_vip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_site_with_vip,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_dualstack_on_public,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_on_public,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:advertise_v6_on_public,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_network,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_site,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where:ConflictingListObjectAttributes:virtual_site_with_vip,vk8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["advertise_custom", "advertise_where"], "schema_version": 1, "sections": [{"aliases": ["advertise dualstack on public"], "anchor": "section", "description": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_dualstack_on_public", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "advertise_dualstack_on_public"], "syntax": "block", "type": "object"}, {"aliases": ["advertise on public"], "anchor": "section", "description": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_on_public", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "advertise_on_public"], "syntax": "block", "type": "object"}, {"aliases": ["advertise v6 on public"], "anchor": "section", "description": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:advertise_v6_on_public", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "advertise_v6_on_public"], "syntax": "block", "type": "object"}, {"aliases": ["port"], "anchor": "schema-advertise_custom--advertise_where--port", "description": "Exclusive with Port to Listen.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["port ranges"], "anchor": "schema-advertise_custom--advertise_where--port_ranges", "description": "Exclusive with A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by \"-\".", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "port_ranges"], "syntax": "attribute", "type": "string"}, {"aliases": ["site"], "anchor": "section", "description": "This defines a reference to a CE site along with network type and an optional IP address where a load balancer could be advertised.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "site"], "syntax": "block", "type": "object"}, {"aliases": ["use default port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:use_default_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "use_default_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual network"], "anchor": "section", "description": "Parameters to advertise on a given virtual network.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-advertise_custom--advertise_where--virtual_network--specific_v6_vip", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where.virtual_network:ConflictingObjectAttributes:default_v6_vip,specific_v6_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "type": "conflicts"}, {"anchor": "schema-advertise_custom--advertise_where--virtual_network--specific_vip", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where.virtual_network:ConflictingObjectAttributes:default_vip,specific_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where.virtual_network:ConflictingObjectAttributes:default_v6_vip,specific_v6_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:default_v6_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where.virtual_network:ConflictingObjectAttributes:default_vip,specific_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_network:default_vip", "type": "conflicts"}], "schema_path": ["advertise_custom", "advertise_where", "virtual_network"], "syntax": "block", "type": "object"}, {"aliases": ["virtual site"], "anchor": "section", "description": "This defines a reference to a customer site virtual site along with network type where a load balancer could be advertised.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "virtual_site"], "syntax": "block", "type": "object"}, {"aliases": ["virtual site with vip"], "anchor": "section", "description": "This defines a reference to a customer site virtual site along with network type and IP where a load balancer could be advertised.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:virtual_site_with_vip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advertise_custom", "advertise_where", "virtual_site_with_vip"], "syntax": "block", "type": "object"}, {"aliases": ["vk8s service"], "anchor": "section", "description": "This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where.vk8s_service:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advertise_custom.advertise_where.vk8s_service:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom:advertise_where:vk8s_service:virtual_site", "type": "conflicts"}], "schema_path": ["advertise_custom", "advertise_where", "vk8s_service"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Where should this load balancer be available.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advertise_custom.advertise_where

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/)
- advertise_custom.advertise_where

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingListObjectAttributes("port",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("port_ranges",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site_with_vip",
    "vk8s_service")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_dualstack_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/advertise_dualstack_on_public/): complete subsection reference.

- [advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/advertise_on_public/): complete subsection reference.

- [advertise_v6_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/advertise_v6_on_public/): complete subsection reference.

<a id="schema-advertise_custom--advertise_where--port"></a>

### port property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Upstream description:

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-advertise_custom--advertise_where--port_ranges"></a>

### port_ranges property

Type: `"string"`. Optional.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/site/): complete subsection reference.

- [use_default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/use_default_port/): complete subsection reference.

- [virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/virtual_network/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/virtual_site/): complete subsection reference.

- [virtual_site_with_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/virtual_site_with_vip/): complete subsection reference.

- [vk8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/vk8s_service/): complete subsection reference.

## Next pages

- [advertise_custom.advertise_where.advertise_dualstack_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/advertise_dualstack_on_public/)
- [advertise_custom.advertise_where.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/advertise_on_public/)
- [advertise_custom.advertise_where.advertise_v6_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/advertise_v6_on_public/)
- [advertise_custom.advertise_where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/site/)
- [advertise_custom.advertise_where.use_default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/use_default_port/)
- [advertise_custom.advertise_where.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/virtual_network/)
- [advertise_custom.advertise_where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/virtual_site/)
- [advertise_custom.advertise_where.virtual_site_with_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/virtual_site_with_vip/)
- [advertise_custom.advertise_where.vk8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/advertise_where/vk8s_service/)
- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/advertise_custom/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
