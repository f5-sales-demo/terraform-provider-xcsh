---
page_title: "server_spec.locations"
subcategory: ""
description: "Configuration of the set of locations corresponding to this server."
xcsh_docs: {"aliases": ["server spec locations"], "body_bytes": 2104, "body_sha256": "sha256:2f6c1c5767486d36e06b3f9b03fa20933b01555ee35b860fffabd2aaf5d09408", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:api_discovery_spec", "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations", "parent_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec", "path": "documentation/data-sources/nginx_server/properties/server_spec/locations/index.md", "product": "distributed-cloud", "provider_name": "nginx_server", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2203001311300220-2301213332031013-3223121201121313-1232332231300212-3302201103310311-2332230322131121-2210000030003210-2220203223201203", "registry_path": "docs/guides/data-sources--nginx_server--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["server_spec", "locations"], "schema_version": 1, "sections": [{"aliases": ["server spec locations api discovery spec"], "anchor": "section", "description": "Configuration for api_discovery_spec.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:api_discovery_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["server_spec", "locations", "api_discovery_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["server spec locations definition"], "anchor": "schema-server_spec--locations--definition", "description": "Location definition specified as the attributes of ${location} block in NGINX config. This includes both the optional_modifier and the location_match combined. A location can either be defined by a prefix string, or by a regular expression.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "locations", "definition"], "syntax": "attribute", "type": "string"}, {"aliases": ["server spec locations name"], "anchor": "schema-server_spec--locations--name", "description": "Uniqueness identifier for a location definition.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_spec", "locations", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["server spec locations waf spec"], "anchor": "section", "description": "Configuration for waf_spec.", "document_id": "xcsh-docs:data-sources:nginx_server:properties:server_spec:locations:waf_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["server_spec", "locations", "waf_spec"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/properties/server_spec/locations/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Configuration of the set of locations corresponding to this server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": [], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# server_spec.locations

Breadcrumbs:

- [xcsh_nginx_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/)
- [server_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/)
- server_spec.locations

<a id="section"></a>

Type: `"list"`. Computed.

Configuration of the set of locations corresponding to this server.

## Direct properties

- [api_discovery_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/api_discovery_spec/): complete subsection reference.

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

- [waf_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/): complete subsection reference.

## Next pages

- [server_spec.locations.api_discovery_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/api_discovery_spec/)
- [server_spec.locations.waf_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/locations/waf_spec/)
- [server_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/properties/server_spec/)
- [xcsh_nginx_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/)
