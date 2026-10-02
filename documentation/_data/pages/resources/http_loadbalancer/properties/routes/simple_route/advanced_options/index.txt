---
page_title: "routes.simple_route.advanced_options"
subcategory: "Load Balancing"
description: "Configure advanced OPTIONS for route like path rewrite, hash policy, etc."
xcsh_docs: {"aliases": ["routes simple route advanced options"], "body_bytes": 27456, "body_sha256": "sha256:a9e238bdd4056d46855ef1e595aac4e87f6bbce695c53135aba1a9b93a757e8c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:app_firewall", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:bot_defense_javascript_injection", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:buffer_policy", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:common_buffering", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:common_hash_policy", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:cors_policy", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:default_retry_policy", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_mirroring", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_prefix_rewrite", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_spdy", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_waf", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_web_socket_config", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:do_not_retract_cluster", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:enable_spdy", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:endpoint_subsets", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_bot_defense_javascript_injection", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_waf", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_waf_exclusion", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:no_retry_policy", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:regex_rewrite", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:request_cookies_to_add", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:request_headers_to_add", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_headers_to_add", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:retract_cluster", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:retry_policy", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:waf_exclusion_policy", "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:web_socket_config"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "path": "documentation/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-024.md", "relationships": [{"anchor": "schema-routes--simple_route--advanced_options--prefix_rewrite", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_prefix_rewrite,prefix_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "type": "conflicts"}, {"anchor": "schema-routes--simple_route--advanced_options--prefix_rewrite", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:prefix_rewrite,regex_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:app_firewall,disable_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:app_firewall", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:app_firewall,inherited_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:app_firewall", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:bot_defense_javascript_injection,inherited_bot_defense_javascript_injection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:bot_defense_javascript_injection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:buffer_policy,common_buffering", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:buffer_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:buffer_policy,common_buffering", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:common_buffering", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:common_hash_policy,specific_hash_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:common_hash_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:default_retry_policy,no_retry_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:default_retry_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:default_retry_policy,retry_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:default_retry_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_mirroring,mirror_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_mirroring", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_prefix_rewrite,prefix_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_prefix_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_prefix_rewrite,regex_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_prefix_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_spdy,enable_spdy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_spdy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:app_firewall,disable_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_waf,inherited_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_web_socket_config,web_socket_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_web_socket_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:do_not_retract_cluster,retract_cluster", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:do_not_retract_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_spdy,enable_spdy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:enable_spdy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:bot_defense_javascript_injection,inherited_bot_defense_javascript_injection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_bot_defense_javascript_injection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:app_firewall,inherited_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_waf,inherited_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:inherited_waf_exclusion,waf_exclusion_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_waf_exclusion", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_mirroring,mirror_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:default_retry_policy,no_retry_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:no_retry_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:no_retry_policy,retry_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:no_retry_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_prefix_rewrite,regex_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:regex_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:prefix_rewrite,regex_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:regex_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:do_not_retract_cluster,retract_cluster", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:retract_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:default_retry_policy,retry_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:retry_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:no_retry_policy,retry_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:retry_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:common_hash_policy,specific_hash_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:inherited_waf_exclusion,waf_exclusion_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:waf_exclusion_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options:ConflictingObjectAttributes:disable_web_socket_config,web_socket_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:web_socket_config", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options"], "schema_version": 1, "sections": [{"aliases": ["app firewall"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:app_firewall", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--simple_route--advanced_options--app_firewall--name", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.app_firewall:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:app_firewall", "type": "requires"}], "schema_path": ["routes", "simple_route", "advanced_options", "app_firewall"], "syntax": "block", "type": "object"}, {"aliases": ["bot defense javascript injection"], "anchor": "section", "description": "Bot Defense Javascript Injection Configuration for inline bot defense deployments.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:bot_defense_javascript_injection", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.bot_defense_javascript_injection:RequiredObjectAttributes:javascript_tags", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:bot_defense_javascript_injection:javascript_tags", "type": "requires"}], "schema_path": ["routes", "simple_route", "advanced_options", "bot_defense_javascript_injection"], "syntax": "block", "type": "object"}, {"aliases": ["buffer policy"], "anchor": "section", "description": "Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level bu", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:buffer_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "buffer_policy"], "syntax": "block", "type": "object"}, {"aliases": ["common buffering"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:common_buffering", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "common_buffering"], "syntax": "attribute", "type": "object"}, {"aliases": ["common hash policy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:common_hash_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "common_hash_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "cors policy", "duration", "operation timeout", "origin servers", "upstream servers"], "anchor": "section", "description": "Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route level configuration takes precedence. An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gec", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:cors_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "cors_policy"], "syntax": "block", "type": "object"}, {"aliases": ["backend servers", "csrf policy", "origin servers", "upstream servers"], "anchor": "section", "description": "To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its destination.the policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request (sou", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,custom_domain_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:all_load_balancer_domains", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:all_load_balancer_domains", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,custom_domain_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:custom_domain_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.csrf_policy:ConflictingObjectAttributes:custom_domain_list,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:custom_domain_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.csrf_policy:ConflictingObjectAttributes:custom_domain_list,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:disabled", "type": "conflicts"}], "schema_path": ["routes", "simple_route", "advanced_options", "csrf_policy"], "syntax": "block", "type": "object"}, {"aliases": ["default retry policy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:default_retry_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "default_retry_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable location add"], "anchor": "schema-routes--simple_route--advanced_options--disable_location_add", "description": "Disables append of x-F5 Distributed Cloud-location = <RE-site-name> at route level, if it is configured at virtual-host level. This configuration is ignored on CE sites.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "disable_location_add"], "syntax": "attribute", "type": "bool"}, {"aliases": ["disable mirroring"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_mirroring", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "disable_mirroring"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable prefix rewrite"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_prefix_rewrite", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "disable_prefix_rewrite"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable spdy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_spdy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "disable_spdy"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable waf"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_waf", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "disable_waf"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable web socket config"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:disable_web_socket_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "disable_web_socket_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["do not retract cluster"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:do_not_retract_cluster", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "do_not_retract_cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable spdy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:enable_spdy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "enable_spdy"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "endpoint subsets", "origin servers", "upstream servers"], "anchor": "section", "description": "Upstream origin pool may be configured to divide its origin servers into subsets based on metadata attached to the origin servers. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer For origin servers which are discovered in K8s or Consul cluster, the label of t", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:endpoint_subsets", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "endpoint_subsets"], "syntax": "block", "type": "object"}, {"aliases": ["inherited bot defense javascript injection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_bot_defense_javascript_injection", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "inherited_bot_defense_javascript_injection"], "syntax": "attribute", "type": "object"}, {"aliases": ["inherited waf"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_waf", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "inherited_waf"], "syntax": "attribute", "type": "object"}, {"aliases": ["inherited waf exclusion"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:inherited_waf_exclusion", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "inherited_waf_exclusion"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "mirror policy", "origin servers", "upstream servers"], "anchor": "section", "description": "MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is \"fire and forget\", meaning it will not wait for the shadow origin pool to respond before returning the response from the primary origin pool. All normal statistics are collected for the shadow origin pool making this featur", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "mirror_policy"], "syntax": "block", "type": "object"}, {"aliases": ["no retry policy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:no_retry_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "no_retry_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["prefix rewrite"], "anchor": "schema-routes--simple_route--advanced_options--prefix_rewrite", "description": "Exclusive with prefix_rewrite indicates that during forwarding, the matched prefix (or path) should be swapped with its value. When using regex path matching, the entire path (not including the query string) will be swapped with this value.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "prefix_rewrite"], "syntax": "attribute", "type": "string"}, {"aliases": ["priority"], "anchor": "schema-routes--simple_route--advanced_options--priority", "description": "Priority routing for each request. Different connection pools are used based on the priority selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based on selected priority. Default routing mechanism High-Priority routing mechanism.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "priority"], "syntax": "attribute", "type": "string"}, {"aliases": ["regex rewrite"], "anchor": "section", "description": "RegexMatchRewrite describes how to match a string and then produce a new string using a regular expression and a substitution string.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:regex_rewrite", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "regex_rewrite"], "syntax": "block", "type": "object"}, {"aliases": ["request cookies to add"], "anchor": "section", "description": "Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:request_cookies_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "request_cookies_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["request cookies to remove"], "anchor": "schema-routes--simple_route--advanced_options--request_cookies_to_remove", "description": "List of keys of Cookies to be removed from the HTTP request being sent towards upstream.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "request_cookies_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["request headers to add"], "anchor": "section", "description": "Headers are key-value pairs to be added to HTTP request being routed towards upstream.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:request_headers_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "request_headers_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["request headers to remove"], "anchor": "schema-routes--simple_route--advanced_options--request_headers_to_remove", "description": "List of keys of Headers to be removed from the HTTP request being sent towards upstream.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "request_headers_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["response cookies to add"], "anchor": "section", "description": "Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream. Cookies specified at this level are applied after cookies from matched Route are applied.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "response_cookies_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["response cookies to remove"], "anchor": "schema-routes--simple_route--advanced_options--response_cookies_to_remove", "description": "List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "response_cookies_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["response headers to add"], "anchor": "section", "description": "Headers are key-value pairs to be added to HTTP response being sent towards downstream.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_headers_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "response_headers_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["response headers to remove"], "anchor": "schema-routes--simple_route--advanced_options--response_headers_to_remove", "description": "List of keys of Headers to be removed from the HTTP response being sent towards downstream.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "response_headers_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["retract cluster"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:retract_cluster", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "retract_cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["retry policy"], "anchor": "section", "description": "Retry policy configuration for route destination.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:retry_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--simple_route--advanced_options--retry_policy--retry_condition", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.retry_policy:RequiredObjectAttributes:retry_condition", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:retry_policy", "type": "requires"}], "schema_path": ["routes", "simple_route", "advanced_options", "retry_policy"], "syntax": "block", "type": "object"}, {"aliases": ["specific hash policy"], "anchor": "section", "description": "List of hash policy rules.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.specific_hash_policy:RequiredObjectAttributes:hash_policy", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:specific_hash_policy:hash_policy", "type": "requires"}], "schema_path": ["routes", "simple_route", "advanced_options", "specific_hash_policy"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeout"], "anchor": "schema-routes--simple_route--advanced_options--timeout", "description": "The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0 (infinite timeout) for server-side streaming.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["waf exclusion policy"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:waf_exclusion_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--simple_route--advanced_options--waf_exclusion_policy--name", "enforcement": "provider-schema", "group": "routes.simple_route.advanced_options.waf_exclusion_policy:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:waf_exclusion_policy", "type": "requires"}], "schema_path": ["routes", "simple_route", "advanced_options", "waf_exclusion_policy"], "syntax": "block", "type": "object"}, {"aliases": ["web socket config"], "anchor": "section", "description": "Configuration to allow Websocket Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'websocket' With configuration to allow websocket upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:web_socket_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "simple_route", "advanced_options", "web_socket_config"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configure advanced OPTIONS for route like path rewrite, hash policy, etc.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/)
- routes.simple_route.advanced_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure advanced OPTIONS for route like path rewrite, hash policy, etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall",
    "disable_waf"),
  validators.ConflictingObjectAttributes("app_firewall",
    "inherited_waf"),
  validators.ConflictingObjectAttributes("bot_defense_javascript_injection",
    "inherited_bot_defense_javascript_injection"),
  validators.ConflictingObjectAttributes("buffer_policy",
    "common_buffering"),
  validators.ConflictingObjectAttributes("common_hash_policy",
    "specific_hash_policy"),
  validators.ConflictingObjectAttributes("default_retry_policy",
    "no_retry_policy"),
  validators.ConflictingObjectAttributes("default_retry_policy",
    "retry_policy"),
  validators.ConflictingObjectAttributes("disable_mirroring",
    "mirror_policy"),
  validators.ConflictingObjectAttributes("disable_prefix_rewrite",
    "prefix_rewrite"),
  validators.ConflictingObjectAttributes("disable_prefix_rewrite",
    "regex_rewrite"),
  validators.ConflictingObjectAttributes("disable_spdy",
    "enable_spdy"),
  validators.ConflictingObjectAttributes("disable_waf",
    "inherited_waf"),
  validators.ConflictingObjectAttributes("disable_web_socket_config",
    "web_socket_config"),
  validators.ConflictingObjectAttributes("do_not_retract_cluster",
    "retract_cluster"),
  validators.ConflictingObjectAttributes("inherited_waf_exclusion",
    "waf_exclusion_policy"),
  validators.ConflictingObjectAttributes("no_retry_policy",
    "retry_policy"),
  validators.ConflictingObjectAttributes("prefix_rewrite",
    "regex_rewrite")}
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
  "x-ves-oneof-field-bot_defense_javascript_injection_choice": "[\"bot_defense_javascript_injection\",\"inherited_bot_defense_javascript_injection\"]",
  "x-ves-oneof-field-buffer_choice": "[\"buffer_policy\",\"common_buffering\"]",
  "x-ves-oneof-field-cluster_retract_choice": "[\"do_not_retract_cluster\",\"retract_cluster\"]",
  "x-ves-oneof-field-hash_policy_choice": "[\"common_hash_policy\",\"specific_hash_policy\"]",
  "x-ves-oneof-field-mirroring_choice": "[\"disable_mirroring\",\"mirror_policy\"]",
  "x-ves-oneof-field-retry_policy_choice": "[\"default_retry_policy\",\"no_retry_policy\",\"retry_policy\"]",
  "x-ves-oneof-field-rewrite_choice": "[\"disable_prefix_rewrite\",\"prefix_rewrite\",\"regex_rewrite\"]",
  "x-ves-oneof-field-spdy_choice": "[\"disable_spdy\",\"enable_spdy\"]",
  "x-ves-oneof-field-waf_choice": "[\"app_firewall\",\"disable_waf\",\"inherited_waf\"]",
  "x-ves-oneof-field-waf_exclusion_choice": "[\"inherited_waf_exclusion\",\"waf_exclusion_policy\"]",
  "x-ves-oneof-field-websocket_choice": "[\"disable_web_socket_config\",\"web_socket_config\"]"
}
```

Terraform syntax:

```terraform
advanced_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/app_firewall/): complete subsection reference.

- [bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/bot_defense_javascript_injection/): complete subsection reference.

- [buffer_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/buffer_policy/): complete subsection reference.

- [common_buffering](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/common_buffering/): complete subsection reference.

- [common_hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/common_hash_policy/): complete subsection reference.

- [cors_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/cors_policy/): complete subsection reference.

- [csrf_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/csrf_policy/): complete subsection reference.

- [default_retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/default_retry_policy/): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--disable_location_add"></a>

### disable_location_add property

Type: `"bool"`. Optional.

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

Upstream description:

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

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

- [disable_mirroring](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/disable_mirroring/): complete subsection reference.

- [disable_prefix_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/disable_prefix_rewrite/): complete subsection reference.

- [disable_spdy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/disable_spdy/): complete subsection reference.

- [disable_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/disable_waf/): complete subsection reference.

- [disable_web_socket_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/disable_web_socket_config/): complete subsection reference.

- [do_not_retract_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/do_not_retract_cluster/): complete subsection reference.

- [enable_spdy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/enable_spdy/): complete subsection reference.

- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/endpoint_subsets/): complete subsection reference.

- [inherited_bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/inherited_bot_defense_javascript_injection/): complete subsection reference.

- [inherited_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/inherited_waf/): complete subsection reference.

- [inherited_waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/inherited_waf_exclusion/): complete subsection reference.

- [mirror_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/): complete subsection reference.

- [no_retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/no_retry_policy/): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--prefix_rewrite"></a>

### prefix_rewrite property

Type: `"string"`. Optional.

Exclusive with \[disable\_prefix\_rewrite regex\_rewrite\] prefix\_rewrite indicates that during
forwarding, the matched prefix (or path) should be swapped with its value. When using regex path
matching, the entire path (not including the query string) will be swapped with this value.

Upstream description:

Exclusive with \[disable\_prefix\_rewrite regex\_rewrite\] prefix\_rewrite indicates that during
forwarding, the matched prefix (or path) should be swapped with its value. When using regex path
matching, the entire path (not including the query string) will be swapped with this value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-routes--simple_route--advanced_options--priority"></a>

### priority property

Type: `"string"`. Optional.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Upstream description:

Priority routing for each request. Different connection pools are used based on the priority
selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based
on selected priority.

Default routing mechanism High-Priority routing mechanism.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DEFAULT",
    "HIGH"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DEFAULT",
  "enum": [
    "DEFAULT",
    "HIGH"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [regex_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/regex_rewrite/): complete subsection reference.

- [request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/request_cookies_to_add/): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--request_cookies_to_remove"></a>

### request_cookies_to_remove property

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/request_headers_to_add/): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--request_headers_to_remove"></a>

### request_headers_to_remove property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--response_cookies_to_remove"></a>

### response_cookies_to_remove property

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_headers_to_add/): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--response_headers_to_remove"></a>

### response_headers_to_remove property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [retract_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/retract_cluster/): complete subsection reference.

- [retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/retry_policy/): complete subsection reference.

- [specific_hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/): complete subsection reference.

<a id="schema-routes--simple_route--advanced_options--timeout"></a>

### timeout property

Type: `"number"`. Optional.

The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0
(infinite timeout) for server-side streaming.

Upstream description:

The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0
(infinite timeout) for server-side streaming.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/waf_exclusion_policy/): complete subsection reference.

- [web_socket_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/web_socket_config/): complete subsection reference.

## Next pages

- [routes.simple_route.advanced_options.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/app_firewall/)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/bot_defense_javascript_injection/)
- [routes.simple_route.advanced_options.buffer_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/buffer_policy/)
- [routes.simple_route.advanced_options.common_buffering](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/common_buffering/)
- [routes.simple_route.advanced_options.common_hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/common_hash_policy/)
- [routes.simple_route.advanced_options.cors_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/cors_policy/)
- [routes.simple_route.advanced_options.csrf_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/csrf_policy/)
- [routes.simple_route.advanced_options.default_retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/default_retry_policy/)
- [routes.simple_route.advanced_options.disable_mirroring](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/disable_mirroring/)
- [routes.simple_route.advanced_options.disable_prefix_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/disable_prefix_rewrite/)
- [routes.simple_route.advanced_options.disable_spdy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/disable_spdy/)
- [routes.simple_route.advanced_options.disable_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/disable_waf/)
- [routes.simple_route.advanced_options.disable_web_socket_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/disable_web_socket_config/)
- [routes.simple_route.advanced_options.do_not_retract_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/do_not_retract_cluster/)
- [routes.simple_route.advanced_options.enable_spdy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/enable_spdy/)
- [routes.simple_route.advanced_options.endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/endpoint_subsets/)
- [routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/inherited_bot_defense_javascript_injection/)
- [routes.simple_route.advanced_options.inherited_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/inherited_waf/)
- [routes.simple_route.advanced_options.inherited_waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/inherited_waf_exclusion/)
- [routes.simple_route.advanced_options.mirror_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/)
- [routes.simple_route.advanced_options.no_retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/no_retry_policy/)
- [routes.simple_route.advanced_options.regex_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/regex_rewrite/)
- [routes.simple_route.advanced_options.request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/request_cookies_to_add/)
- [routes.simple_route.advanced_options.request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/request_headers_to_add/)
- [routes.simple_route.advanced_options.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/)
- [routes.simple_route.advanced_options.response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_headers_to_add/)
- [routes.simple_route.advanced_options.retract_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/retract_cluster/)
- [routes.simple_route.advanced_options.retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/retry_policy/)
- [routes.simple_route.advanced_options.specific_hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/specific_hash_policy/)
- [routes.simple_route.advanced_options.waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/waf_exclusion_policy/)
- [routes.simple_route.advanced_options.web_socket_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/web_socket_config/)
- [routes.simple_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/simple_route/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
