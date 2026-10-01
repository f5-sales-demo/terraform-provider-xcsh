---
page_title: "server_spec.locations"
subcategory: ""
description: "server_spec.locations for xcsh_nginx_server."
xcsh_docs: {"aliases": [], "body_bytes": 1651, "body_sha256": "sha256:a03b21f833d7420b034076a2e95e63da1d87192118aab0fdc9fbbd3700ef2018", "canonical_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations", "child_ids": ["xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:api_discovery_spec", "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec"], "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations", "parent_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "path": "docs/guides/data-sources--nginx_server--properties--server_spec--locations.md", "provider_name": "nginx_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["server_spec", "locations"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/server_spec/locations/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "server_spec.locations for xcsh_nginx_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
