---
page_title: "server_spec"
subcategory: ""
description: "server_spec for xcsh_nginx_server."
xcsh_docs: {"aliases": [], "body_bytes": 2799, "body_sha256": "sha256:b4bfc8bd8fe03e3c07b8a5bda68dfd2c9d01507d02ea139135482690266248e3", "child_ids": ["xcsh-docs:data-sources:nginx_server:properties:server_spec:api_discovery_spec", "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations", "xcsh-docs:data-sources:nginx_server:properties:server_spec:waf_spec"], "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "parent_id": "xcsh-docs:data-sources:nginx_server:reference", "path": "documentation/data-sources/nginx_server/properties/server_spec/index.md", "provider_name": "nginx_server", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["server_spec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/server_spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "server_spec for xcsh_nginx_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

## Next pages

- [server_spec.api_discovery_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/api_discovery_spec/)
- [server_spec.locations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/)
- [server_spec.waf_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/waf_spec/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/)
- [xcsh_nginx_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/)
