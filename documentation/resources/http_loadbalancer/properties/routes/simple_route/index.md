---
page_title: "routes.simple_route"
subcategory: "Load Balancing"
description: "A simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the matching traffic to the associated pools."
xcsh_docs: {"aliases": ["routes simple route"], "body_bytes": 7865, "body_sha256": "sha256:e601add5e94c6add8a3a1cb40fc26221226240c36c577a78143336dce715babf", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:auto_host_rewrite", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:caching_disable", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:caching_inherit", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:disable_host_rewrite", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:headers", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:origin_pools", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:path", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:query_params"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes", "path": "documentation/resources/http_loadbalancer/properties/routes/simple_route/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-024.md", "relationships": [{"anchor": "schema-routes--simple_route--host_rewrite", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:auto_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--host_rewrite", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:disable_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:auto_host_rewrite,disable_host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:auto_host_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:auto_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:auto_host_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:caching_disable,caching_inherit", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:caching_disable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:caching_disable,caching_inherit", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:caching_inherit", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:auto_host_rewrite,disable_host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:disable_host_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route:ConflictingObjectAttributes:disable_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:disable_host_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route:RequiredObjectAttributes:origin_pools", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:origin_pools", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "simple_route"], "schema_version": 1, "sections": [{"aliases": ["advanced options"], "anchor": "section", "description": "Configure advanced OPTIONS for route like path rewrite, hash policy, etc.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--simple_route--advanced_options--prefix_rewrite", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_prefix_rewrite,prefix_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--advanced_options--prefix_rewrite", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:prefix_rewrite,regex_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:app_firewall,disable_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:app_firewall", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:app_firewall,inherited_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:app_firewall", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:bot_defense_javascript_injection,inherited_bot_defense_javascript_injection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:bot_defense_javascript_injection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:buffer_policy,common_buffering", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:buffer_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:buffer_policy,common_buffering", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:common_buffering", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:common_hash_policy,specific_hash_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:common_hash_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:default_retry_policy,no_retry_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:default_retry_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:default_retry_policy,retry_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:default_retry_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_mirroring,mirror_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_mirroring", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_prefix_rewrite,prefix_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_prefix_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_prefix_rewrite,regex_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_prefix_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_spdy,enable_spdy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_spdy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:app_firewall,disable_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_waf,inherited_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_web_socket_config,web_socket_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_web_socket_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:do_not_retract_cluster,retract_cluster", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:do_not_retract_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_spdy,enable_spdy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:enable_spdy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:bot_defense_javascript_injection,inherited_bot_defense_javascript_injection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_bot_defense_javascript_injection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:app_firewall,inherited_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_waf,inherited_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:inherited_waf_exclusion,waf_exclusion_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_waf_exclusion", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_mirroring,mirror_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:default_retry_policy,no_retry_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:no_retry_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:no_retry_policy,retry_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:no_retry_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_prefix_rewrite,regex_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:regex_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:prefix_rewrite,regex_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:regex_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:do_not_retract_cluster,retract_cluster", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:retract_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:default_retry_policy,retry_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:retry_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:no_retry_policy,retry_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:retry_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:common_hash_policy,specific_hash_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:inherited_waf_exclusion,waf_exclusion_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:waf_exclusion_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_web_socket_config,web_socket_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:web_socket_config", "type": "conflicts"}], "schema_path": ["routes", "simple_route", "advanced_options"], "syntax": "block", "type": "object"}, {"aliases": ["auto host rewrite"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:auto_host_rewrite", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "auto_host_rewrite"], "syntax": "attribute", "type": "object"}, {"aliases": ["caching disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:caching_disable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "caching_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["caching inherit"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:caching_inherit", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "caching_inherit"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable host rewrite"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:disable_host_rewrite", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "disable_host_rewrite"], "syntax": "attribute", "type": "object"}, {"aliases": ["headers"], "anchor": "section", "description": "List of (key, value) headers.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:headers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "simple_route", "headers"], "syntax": "block", "type": "object"}, {"aliases": ["host rewrite"], "anchor": "schema-routes--simple_route--host_rewrite", "description": "Exclusive with Host header will be swapped with this value.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "host_rewrite"], "syntax": "attribute", "type": "string"}, {"aliases": ["http method"], "anchor": "schema-routes--simple_route--http_method", "description": "Specifies the HTTP method used to access a resource. Any HTTP Method.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "http_method"], "syntax": "attribute", "type": "string"}, {"aliases": ["incoming port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--simple_route--incoming_port--port", "enforcement": "provider-schema", "group": "routes.simple_route.incoming_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--incoming_port--port", "enforcement": "provider-schema", "group": "routes.simple_route.incoming_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--incoming_port--port_ranges", "enforcement": "provider-schema", "group": "routes.simple_route.incoming_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--incoming_port--port_ranges", "enforcement": "provider-schema", "group": "routes.simple_route.incoming_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.incoming_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port:no_port_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.incoming_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:incoming_port:no_port_match", "type": "conflicts"}], "schema_path": ["routes", "simple_route", "incoming_port"], "syntax": "block", "type": "object"}, {"aliases": ["backend servers", "origin pools", "origin servers", "upstream servers"], "anchor": "section", "description": "Origin Pools for this route.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:origin_pools", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "simple_route", "origin_pools"], "syntax": "block", "type": "object"}, {"aliases": ["path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--simple_route--path--path", "enforcement": "provider-schema", "group": "routes.simple_route.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:path", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--path--path", "enforcement": "provider-schema", "group": "routes.simple_route.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:path", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--path--prefix", "enforcement": "provider-schema", "group": "routes.simple_route.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:path", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--path--prefix", "enforcement": "provider-schema", "group": "routes.simple_route.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:path", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--path--regex", "enforcement": "provider-schema", "group": "routes.simple_route.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:path", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--path--regex", "enforcement": "provider-schema", "group": "routes.simple_route.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:path", "type": "conflicts"}], "schema_path": ["routes", "simple_route", "path"], "syntax": "block", "type": "object"}, {"aliases": ["query params"], "anchor": "section", "description": "Handling of incoming query parameters in simple route.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:query_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--simple_route--query_params--replace_params", "enforcement": "provider-schema", "group": "routes.simple_route.query_params:ConflictingObjectAttributes:remove_all_params,replace_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:query_params", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--query_params--replace_params", "enforcement": "provider-schema", "group": "routes.simple_route.query_params:ConflictingObjectAttributes:replace_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:query_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.query_params:ConflictingObjectAttributes:remove_all_params,replace_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:query_params:remove_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.query_params:ConflictingObjectAttributes:remove_all_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:query_params:remove_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.query_params:ConflictingObjectAttributes:remove_all_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:query_params:retain_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.query_params:ConflictingObjectAttributes:replace_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:query_params:retain_all_params", "type": "conflicts"}], "schema_path": ["routes", "simple_route", "query_params"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "A simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the matching traffic to the associated pools.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- routes.simple_route

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the
matching traffic to the associated pools.

Upstream description:

A simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the
matching traffic to the associated pools.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_pools"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("caching_disable",
    "caching_inherit"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
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
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]",
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
simple_route {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/): complete subsection reference.

- [auto_host_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/auto_host_rewrite/): complete subsection reference.

- [caching_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/caching_disable/): complete subsection reference.

- [caching_inherit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/caching_inherit/): complete subsection reference.

- [disable_host_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/disable_host_rewrite/): complete subsection reference.

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/headers/): complete subsection reference.

<a id="schema-routes--simple_route--host_rewrite"></a>

### host_rewrite property

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="schema-routes--simple_route--http_method"></a>

### http_method property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/incoming_port/): complete subsection reference.

- [origin_pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/origin_pools/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/path/): complete subsection reference.

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/query_params/): complete subsection reference.

## Next pages

- [routes.simple_route.advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/)
- [routes.simple_route.auto_host_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/auto_host_rewrite/)
- [routes.simple_route.caching_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/caching_disable/)
- [routes.simple_route.caching_inherit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/caching_inherit/)
- [routes.simple_route.disable_host_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/disable_host_rewrite/)
- [routes.simple_route.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/headers/)
- [routes.simple_route.incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/incoming_port/)
- [routes.simple_route.origin_pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/origin_pools/)
- [routes.simple_route.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/path/)
- [routes.simple_route.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/query_params/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
