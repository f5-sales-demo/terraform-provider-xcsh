---
page_title: "server_spec"
subcategory: ""
description: "server_spec for xcsh_nginx_server."
xcsh_docs: {"aliases": [], "body_bytes": 2390, "body_sha256": "sha256:b227dc531cfb4953bc152131aea46298018f175e7c8c18e913fba1f0660e112b", "canonical_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "child_ids": ["xcsh-docs:data-sources:nginx_server:properties:server_spec:api_discovery_spec", "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations", "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec"], "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "parent_id": "xcsh-docs:data-sources:nginx_server:reference", "path": "docs/guides/data-sources--nginx_server--properties--server_spec.md", "provider_name": "nginx_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["server_spec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/server_spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "server_spec for xcsh_nginx_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# server_spec

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md)
- [Property reference](data-sources--nginx_server--reference.md)
- server_spec

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for server\_spec.

## Direct properties

- [api_discovery_spec](data-sources--nginx_server--properties--server_spec--api_discovery_spec.md): complete subsection reference.

<a id="schema-server_spec--domains"></a>

### domains property

Type: `["list", "string"]`. Computed.

Server name list specified as $\{server\_name\} in NGINX config. If no value is specified
corresponding to this variable, 'default' is used Reference:
https&#58;//nginx.org/en/docs/HTTP/ngx\_http\_core\_module.html\#server.

- [locations](data-sources--nginx_server--properties--server_spec--locations.md): complete subsection reference.

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

- [waf_spec](data-sources--nginx_server--properties--server_spec--waf_spec.md): complete subsection reference.

## Next pages

- [server_spec.api_discovery_spec](data-sources--nginx_server--properties--server_spec--api_discovery_spec.md)
- [server_spec.locations](data-sources--nginx_server--properties--server_spec--locations.md)
- [server_spec.waf_spec](data-sources--nginx_server--properties--server_spec--waf_spec.md)
- [Property reference](data-sources--nginx_server--reference.md)
- [xcsh_nginx_server](../data-sources/nginx_server.md)
