---
page_title: "server_spec"
subcategory: ""
description: "Configuration for server_spec."
xcsh_docs: {"aliases": ["server spec"], "body_bytes": 2213, "body_sha256": "sha256:be80fbd70a03f657f1ec2d69364136c8538d46996336c29e003246633887bc43", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_server:properties:server_spec:api_discovery_spec", "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations", "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "parent_id": "xcsh-docs:data-sources:nginx_server:reference", "path": "documentation/data-sources/nginx_server/properties/server_spec/index.md", "product": "distributed-cloud", "provider_name": "nginx_server", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0122302310010021-1321221030311133-1100310212321130-2322212212032232-3203223221022103-2302223302323122-2021001303002122-1323202112301331", "registry_path": "docs/guides/data-sources--nginx_server--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["server_spec"], "schema_version": 1, "sections": [{"aliases": ["server spec api discovery spec"], "anchor": "section", "description": "Configuration for api_discovery_spec.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:api_discovery_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["server_spec", "api_discovery_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["server spec domains"], "anchor": "schema-server_spec--domains", "description": "Server name list specified as ${server_name} in NGINX config. If no value is specified corresponding to this variable, 'default' is used Reference: https://nginx.org/en/docs/HTTP/ngx_http_core_module.html#server.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "domains"], "syntax": "attribute", "type": "list"}, {"aliases": ["server spec locations"], "anchor": "section", "description": "Configuration of the set of locations corresponding to this server.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["server_spec", "locations"], "syntax": "attribute", "type": "object"}, {"aliases": ["server spec nginx one object id"], "anchor": "schema-server_spec--nginx_one_object_id", "description": "Signifies the uniqueness identifier for NGINX One representation of this NGINX server.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "nginx_one_object_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["server spec nginx one object name"], "anchor": "schema-server_spec--nginx_one_object_name", "description": "Hostname value set for Instance or Name for a Config Sync Group in NGINX One.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "nginx_one_object_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["server spec port"], "anchor": "schema-server_spec--port", "description": "Signifies the port configured for the NGINX server.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["server spec server name"], "anchor": "schema-server_spec--server_name", "description": "Signifies the combination of first element in domains array and the port configured for the NGINX server.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "server_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["server spec total routes"], "anchor": "schema-server_spec--total_routes", "description": "Total locations configured in the NGINX Server.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "total_routes"], "syntax": "attribute", "type": "number"}, {"aliases": ["server spec waf spec"], "anchor": "section", "description": "Configuration for waf_spec.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["server_spec", "waf_spec"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/server_spec/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Configuration for server_spec.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# server_spec

Breadcrumbs:

- [xcsh_nginx_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/)
- server_spec

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for server\_spec.

## Direct properties

- [api_discovery_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/api_discovery_spec/): complete subsection reference.

<a id="schema-server_spec--domains"></a>

### domains property

Type: `["list", "string"]`. Computed.

Server name list specified as $\{server\_name\} in NGINX config. If no value is specified
corresponding to this variable, 'default' is used Reference:
https&#58;//nginx.org/en/docs/HTTP/ngx\_http\_core\_module.html\#server.

- [locations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/): complete subsection reference.

<a id="schema-server_spec--nginx_one_object_id"></a>

### nginx_one_object_id property

Type: `"string"`. Computed.

Signifies the uniqueness identifier for NGINX One representation of this NGINX server.

<a id="schema-server_spec--nginx_one_object_name"></a>

### nginx_one_object_name property

Type: `"string"`. Computed.

Hostname value set for Instance or Name for a Config Sync Group in NGINX One.

<a id="schema-server_spec--port"></a>

### port property

Type: `"number"`. Computed.

Signifies the port configured for the NGINX server.

<a id="schema-server_spec--server_name"></a>

### server_name property

Type: `"string"`. Computed.

Signifies the combination of first element in domains array and the port configured for the NGINX
server.

<a id="schema-server_spec--total_routes"></a>

### total_routes property

Type: `"number"`. Computed.

Total locations configured in the NGINX Server.

- [waf_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/): complete subsection reference.
