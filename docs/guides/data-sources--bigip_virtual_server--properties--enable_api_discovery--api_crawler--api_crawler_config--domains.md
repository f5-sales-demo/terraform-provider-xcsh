---
page_title: "enable_api_discovery.api_crawler.api_crawler_config.domains"
subcategory: ""
description: "enable_api_discovery.api_crawler.api_crawler_config.domains for xcsh_bigip_virtual_server."
xcsh_docs: {"aliases": [], "body_bytes": 1835, "body_sha256": "sha256:afa1872b17d63f651c4d284968a35d512b2535f894f99b7d4a6c566746f28e1c", "canonical_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains", "child_ids": ["xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login"], "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config", "path": "docs/guides/data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains.md", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery.api_crawler.api_crawler_config.domains for xcsh_bigip_virtual_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config.domains

Breadcrumbs:

- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md)
- [Property reference](data-sources--bigip_virtual_server--reference.md)
- [enable_api_discovery](data-sources--bigip_virtual_server--properties--enable_api_discovery.md)
- [enable_api_discovery.api_crawler](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler.md)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config.md)
- enable_api_discovery.api_crawler.api_crawler_config.domains

<a id="section"></a>

Type: `"list"`. Computed.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

## Direct properties

<a id="schema-enable_api_discovery--api_crawler--api_crawler_config--domains--domain"></a>

### domain property

Type: `"string"`. Computed.

Select the domain to execute API Crawling with given credentials.

- [simple_login](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login.md): complete subsection reference.

## Next pages

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login.md)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config.md)
- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md)
