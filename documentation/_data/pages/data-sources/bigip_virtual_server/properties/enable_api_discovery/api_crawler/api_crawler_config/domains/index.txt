---
page_title: "enable_api_discovery.api_crawler.api_crawler_config.domains"
subcategory: ""
description: "Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials", "enable api discovery api crawler api crawler config domains"], "body_bytes": 2283, "body_sha256": "sha256:a6613bd65dd08d3a1b9a2a553589936ac4fc4a93ff65bfb0bc4159f39395ffbc", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config", "path": "documentation/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0110113103100310-0002102020033030-1303030113331333-0232203332320011-0203131322022230-1110223103333301-1323300323220300-2232233203131010", "registry_path": "docs/guides/data-sources--bigip_virtual_server--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains"], "schema_version": 1, "sections": [{"aliases": ["authentication", "credential setup", "credentials", "domain"], "anchor": "schema-enable_api_discovery--api_crawler--api_crawler_config--domains--domain", "description": "Select the domain to execute API Crawling with given credentials.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["login", "login result", "sign in", "simple login"], "anchor": "section", "description": "Configuration parameter for simple login.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config.domains

Breadcrumbs:

- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/)
- [enable_api_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/)
- [enable_api_discovery.api_crawler.api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/)
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

- [simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/): complete subsection reference.

## Next pages

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/)
- [enable_api_discovery.api_crawler.api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/)
- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
