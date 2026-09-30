---
page_title: "server_spec.locations"
subcategory: ""
description: "server_spec.locations for xcsh_nginx_server."
xcsh_docs: {"aliases": [], "body_bytes": 1552, "body_sha256": "sha256:8a4b7807b73b9c3210a20c979acce6ed5f23cf66beb3d2c5568ce821e726ed51", "canonical_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations", "child_ids": ["xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:api_discovery_spec", "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec"], "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations", "parent_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "path": "docs/guides/data-sources--nginx_server--properties--server_spec--locations.md", "provider_name": "nginx_server", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["server_spec", "locations"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/server_spec/locations/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "server_spec.locations for xcsh_nginx_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# server_spec.locations

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md)
- [Property reference](data-sources--nginx_server--reference.md)
- [server_spec](data-sources--nginx_server--properties--server_spec.md)
- server_spec.locations

<a id="section"></a>

Type: `"list"`. Computed.

Configuration of the set of locations corresponding to this server.

## Direct properties

- [api_discovery_spec](data-sources--nginx_server--properties--server_spec--locations--api_discovery_spec.md): complete subsection reference.

<a id="schema-server_spec--locations--definition"></a>

### definition property

Type: `"string"`. Computed.

Location definition specified as the attributes of $\{location\} block in NGINX config. This
includes both the optional\_modifier and the location\_match combined. A location can either be
defined by a prefix string, or by a regular expression.

<a id="schema-server_spec--locations--name"></a>

### name property

Type: `"string"`. Computed.

Uniqueness identifier for a location definition.

- [waf_spec](data-sources--nginx_server--properties--server_spec--locations--waf_spec.md): complete subsection reference.

## Next pages

- [server_spec.locations.api_discovery_spec](data-sources--nginx_server--properties--server_spec--locations--api_discovery_spec.md)
- [server_spec.locations.waf_spec](data-sources--nginx_server--properties--server_spec--locations--waf_spec.md)
- [server_spec](data-sources--nginx_server--properties--server_spec.md)
- [xcsh_nginx_server](../data-sources/nginx_server.md)
