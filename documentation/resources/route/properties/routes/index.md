---
page_title: "routes"
subcategory: ""
description: "List of routes to match for incoming request."
xcsh_docs: {"aliases": ["routes"], "body_bytes": 12499, "body_sha256": "sha256:75edea0bf7b7fc3537674ecdd85d7679239e009fb665ddbcd8e7debab9352457", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:bot_defense_javascript_injection", "xcsh-docs:resources:route:properties:routes:inherited_bot_defense_javascript_injection", "xcsh-docs:resources:route:properties:routes:inherited_waf_exclusion", "xcsh-docs:resources:route:properties:routes:match", "xcsh-docs:resources:route:properties:routes:request_cookies_to_add", "xcsh-docs:resources:route:properties:routes:request_headers_to_add", "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "xcsh-docs:resources:route:properties:routes:response_headers_to_add", "xcsh-docs:resources:route:properties:routes:route_destination", "xcsh-docs:resources:route:properties:routes:route_direct_response", "xcsh-docs:resources:route:properties:routes:route_redirect", "xcsh-docs:resources:route:properties:routes:service_policy", "xcsh-docs:resources:route:properties:routes:waf_exclusion_policy", "xcsh-docs:resources:route:properties:routes:waf_type"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes", "parent_id": "xcsh-docs:resources:route:reference", "path": "documentation/resources/route/properties/routes/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130", "registry_path": "docs/guides/resources--route--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:bot_defense_javascript_injection,inherited_bot_defense_javascript_injection", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:bot_defense_javascript_injection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:bot_defense_javascript_injection,inherited_bot_defense_javascript_injection", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:inherited_bot_defense_javascript_injection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:inherited_waf_exclusion,waf_exclusion_policy", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:inherited_waf_exclusion", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:route_destination,route_direct_response", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:route_destination,route_redirect", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:route_destination,route_direct_response", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_direct_response", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:route_direct_response,route_redirect", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_direct_response", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:route_destination,route_redirect", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_redirect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:route_direct_response,route_redirect", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_redirect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:inherited_waf_exclusion,waf_exclusion_policy", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_exclusion_policy", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes"], "schema_version": 1, "sections": [{"aliases": ["routes bot defense javascript injection"], "anchor": "section", "description": "Bot Defense Javascript Injection Configuration for inline bot defense deployments.", "document_id": "xcsh-docs:resources:route:properties:routes:bot_defense_javascript_injection", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.bot_defense_javascript_injection:RequiredObjectAttributes:javascript_tags", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:bot_defense_javascript_injection:javascript_tags", "type": "requires"}], "schema_path": ["routes", "bot_defense_javascript_injection"], "syntax": "block", "type": "object"}, {"aliases": ["routes disable location add"], "anchor": "schema-routes--disable_location_add", "description": "Disables append of x-F5 Distributed Cloud-location = <RE-site-name> at route level, if it is configured at virtual-host level. This configuration is ignored on CE sites.", "document_id": "xcsh-docs:resources:route:properties:routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "disable_location_add"], "syntax": "attribute", "type": "bool"}, {"aliases": ["routes inherited bot defense javascript injection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:route:properties:routes:inherited_bot_defense_javascript_injection", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "inherited_bot_defense_javascript_injection"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes inherited waf exclusion"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:route:properties:routes:inherited_waf_exclusion", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "inherited_waf_exclusion"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes match"], "anchor": "section", "description": "Route match condition.", "document_id": "xcsh-docs:resources:route:properties:routes:match", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "match"], "syntax": "block", "type": "object"}, {"aliases": ["routes request cookies to add"], "anchor": "section", "description": "Cookies are key-value pairs to be added to HTTP request being routed towards upstream.", "document_id": "xcsh-docs:resources:route:properties:routes:request_cookies_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-routes--request_cookies_to_add--value", "enforcement": "provider-schema", "group": "routes.request_cookies_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:request_cookies_to_add", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.request_cookies_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:request_cookies_to_add:secret_value", "type": "conflicts"}, {"anchor": "schema-routes--request_cookies_to_add--name", "enforcement": "provider-schema", "group": "routes.request_cookies_to_add:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:request_cookies_to_add", "type": "requires"}], "schema_path": ["routes", "request_cookies_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["routes request cookies to remove"], "anchor": "schema-routes--request_cookies_to_remove", "description": "List of keys of Cookies to be removed from the HTTP request being sent towards upstream.", "document_id": "xcsh-docs:resources:route:properties:routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "request_cookies_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["routes request headers to add"], "anchor": "section", "description": "Headers are key-value pairs to be added to HTTP requests being sent towards upstream. Headers specified at this level are applied before headers from the enclosing VirtualHost object level.", "document_id": "xcsh-docs:resources:route:properties:routes:request_headers_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-routes--request_headers_to_add--value", "enforcement": "provider-schema", "group": "routes.request_headers_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:request_headers_to_add", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.request_headers_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:request_headers_to_add:secret_value", "type": "conflicts"}, {"anchor": "schema-routes--request_headers_to_add--name", "enforcement": "provider-schema", "group": "routes.request_headers_to_add:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:request_headers_to_add", "type": "requires"}], "schema_path": ["routes", "request_headers_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["routes request headers to remove"], "anchor": "schema-routes--request_headers_to_remove", "description": "List of keys of Headers to be removed from the HTTP request being sent towards upstream.", "document_id": "xcsh-docs:resources:route:properties:routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "request_headers_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["routes response cookies to add"], "anchor": "section", "description": "Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream.", "document_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-routes--response_cookies_to_add--add_domain", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:add_domain,ignore_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-routes--response_cookies_to_add--add_expiry", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:add_expiry,ignore_expiry", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-routes--response_cookies_to_add--add_path", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:add_path,ignore_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-routes--response_cookies_to_add--max_age_value", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:ignore_max_age,max_age_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-routes--response_cookies_to_add--value", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:ignore_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-routes--response_cookies_to_add--value", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:add_httponly,ignore_httponly", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:add_httponly", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:add_partitioned,ignore_partitioned", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:add_partitioned", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:add_secure,ignore_secure", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:add_secure", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:add_domain,ignore_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:add_expiry,ignore_expiry", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_expiry", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:add_httponly,ignore_httponly", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_httponly", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:ignore_max_age,max_age_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_max_age", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:add_partitioned,ignore_partitioned", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_partitioned", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:add_path,ignore_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_lax", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:add_secure,ignore_secure", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_secure", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:ignore_value,secret_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_value", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:ignore_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:ignore_value", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_lax", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:samesite_lax,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:samesite_lax,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:samesite_lax,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:samesite_none,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:samesite_lax,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:samesite_none,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:ignore_value,secret_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:secret_value", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add:secret_value", "type": "conflicts"}, {"anchor": "schema-routes--response_cookies_to_add--name", "enforcement": "provider-schema", "group": "routes.response_cookies_to_add:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_cookies_to_add", "type": "requires"}], "schema_path": ["routes", "response_cookies_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["routes response cookies to remove"], "anchor": "schema-routes--response_cookies_to_remove", "description": "List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed.", "document_id": "xcsh-docs:resources:route:properties:routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "response_cookies_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["routes response headers to add"], "anchor": "section", "description": "Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers specified at this level are applied before headers from the enclosing VirtualHost object level.", "document_id": "xcsh-docs:resources:route:properties:routes:response_headers_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-routes--response_headers_to_add--value", "enforcement": "provider-schema", "group": "routes.response_headers_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_headers_to_add", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.response_headers_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_headers_to_add:secret_value", "type": "conflicts"}, {"anchor": "schema-routes--response_headers_to_add--name", "enforcement": "provider-schema", "group": "routes.response_headers_to_add:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:response_headers_to_add", "type": "requires"}], "schema_path": ["routes", "response_headers_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["routes response headers to remove"], "anchor": "schema-routes--response_headers_to_remove", "description": "List of keys of Headers to be removed from the HTTP response being sent towards downstream.", "document_id": "xcsh-docs:resources:route:properties:routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "response_headers_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["routes route destination"], "anchor": "section", "description": "List of destination to choose if the route is match.", "document_id": "xcsh-docs:resources:route:properties:routes:route_destination", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--route_destination--auto_host_rewrite", "enforcement": "provider-schema", "group": "routes.route_destination:ConflictingObjectAttributes:auto_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination", "type": "conflicts"}, {"anchor": "schema-routes--route_destination--host_rewrite", "enforcement": "provider-schema", "group": "routes.route_destination:ConflictingObjectAttributes:auto_host_rewrite,host_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination", "type": "conflicts"}, {"anchor": "schema-routes--route_destination--prefix_rewrite", "enforcement": "provider-schema", "group": "routes.route_destination:ConflictingObjectAttributes:prefix_rewrite,regex_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination:ConflictingObjectAttributes:do_not_retract_cluster,retract_cluster", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:do_not_retract_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination:ConflictingObjectAttributes:prefix_rewrite,regex_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:regex_rewrite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination:ConflictingObjectAttributes:do_not_retract_cluster,retract_cluster", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:retract_cluster", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_destination:RequiredObjectAttributes:destinations", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination:destinations", "type": "requires"}], "schema_path": ["routes", "route_destination"], "syntax": "block", "type": "object"}, {"aliases": ["routes route direct response"], "anchor": "section", "description": "Send this direct response in case of route match action is direct response.", "document_id": "xcsh-docs:resources:route:properties:routes:route_direct_response", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--route_direct_response--response_code", "enforcement": "provider-schema", "group": "routes.route_direct_response:RequiredObjectAttributes:response_code", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_direct_response", "type": "requires"}], "schema_path": ["routes", "route_direct_response"], "syntax": "block", "type": "object"}, {"aliases": ["routes route redirect"], "anchor": "section", "description": "Route redirect parameters when match action is redirect.", "document_id": "xcsh-docs:resources:route:properties:routes:route_redirect", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--route_redirect--path_redirect", "enforcement": "provider-schema", "group": "routes.route_redirect:ConflictingObjectAttributes:path_redirect,prefix_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_redirect", "type": "conflicts"}, {"anchor": "schema-routes--route_redirect--prefix_rewrite", "enforcement": "provider-schema", "group": "routes.route_redirect:ConflictingObjectAttributes:path_redirect,prefix_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_redirect", "type": "conflicts"}, {"anchor": "schema-routes--route_redirect--replace_params", "enforcement": "provider-schema", "group": "routes.route_redirect:ConflictingObjectAttributes:remove_all_params,replace_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_redirect", "type": "conflicts"}, {"anchor": "schema-routes--route_redirect--replace_params", "enforcement": "provider-schema", "group": "routes.route_redirect:ConflictingObjectAttributes:replace_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_redirect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_redirect:ConflictingObjectAttributes:remove_all_params,replace_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_redirect:remove_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_redirect:ConflictingObjectAttributes:remove_all_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_redirect:remove_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_redirect:ConflictingObjectAttributes:remove_all_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_redirect:retain_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.route_redirect:ConflictingObjectAttributes:replace_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_redirect:retain_all_params", "type": "conflicts"}], "schema_path": ["routes", "route_redirect"], "syntax": "block", "type": "object"}, {"aliases": ["routes service policy"], "anchor": "section", "description": "ServicePolicy configuration details at route level.", "document_id": "xcsh-docs:resources:route:properties:routes:service_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "service_policy"], "syntax": "block", "type": "object"}, {"aliases": ["routes waf exclusion policy"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:route:properties:routes:waf_exclusion_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--waf_exclusion_policy--name", "enforcement": "provider-schema", "group": "routes.waf_exclusion_policy:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_exclusion_policy", "type": "requires"}], "schema_path": ["routes", "waf_exclusion_policy"], "syntax": "block", "type": "object"}, {"aliases": ["routes waf type"], "anchor": "section", "description": "WAF instance will be pointing to an app_firewall object.", "document_id": "xcsh-docs:resources:route:properties:routes:waf_type", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:app_firewall,disable_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:app_firewall,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:app_firewall", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:app_firewall,disable_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:disable_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:disable_waf,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:disable_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:app_firewall,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:inherit_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.waf_type:ConflictingObjectAttributes:disable_waf,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_type:inherit_waf", "type": "conflicts"}], "schema_path": ["routes", "waf_type"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of routes to match for incoming request.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- routes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of routes to match for incoming request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bot_defense_javascript_injection",
    "inherited_bot_defense_javascript_injection"),
  validators.ConflictingListObjectAttributes("inherited_waf_exclusion",
    "waf_exclusion_policy"),
  validators.ConflictingListObjectAttributes("route_destination",
    "route_direct_response"),
  validators.ConflictingListObjectAttributes("route_destination",
    "route_redirect"),
  validators.ConflictingListObjectAttributes("route_direct_response",
    "route_redirect")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 257,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 257,
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
    "ves.io.schema.rules.repeated.max_items": "257"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "257"
  }
}
```

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/bot_defense_javascript_injection/): complete subsection reference.

<a id="schema-routes--disable_location_add"></a>

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

- [inherited_bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/inherited_bot_defense_javascript_injection/): complete subsection reference.

- [inherited_waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/inherited_waf_exclusion/): complete subsection reference.

- [match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/): complete subsection reference.

- [request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/): complete subsection reference.

<a id="schema-routes--request_cookies_to_remove"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/): complete subsection reference.

<a id="schema-routes--request_headers_to_remove"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

- [response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/): complete subsection reference.

<a id="schema-routes--response_cookies_to_remove"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/): complete subsection reference.

<a id="schema-routes--response_headers_to_remove"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

- [route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/): complete subsection reference.

- [route_direct_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_direct_response/): complete subsection reference.

- [route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_redirect/): complete subsection reference.

- [service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/service_policy/): complete subsection reference.

- [waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_exclusion_policy/): complete subsection reference.

- [waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/): complete subsection reference.

## Next pages

- [routes.bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/bot_defense_javascript_injection/)
- [routes.inherited_bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/inherited_bot_defense_javascript_injection/)
- [routes.inherited_waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/inherited_waf_exclusion/)
- [routes.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/)
- [routes.request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/)
- [routes.request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/)
- [routes.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/)
- [routes.response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/)
- [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/)
- [routes.route_direct_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_direct_response/)
- [routes.route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_redirect/)
- [routes.service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/service_policy/)
- [routes.waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_exclusion_policy/)
- [routes.waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
