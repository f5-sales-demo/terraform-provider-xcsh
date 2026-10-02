---
page_title: "enable_api_discovery.api_crawler.api_crawler_config.domains"
subcategory: ""
description: "Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials", "enable api discovery api crawler api crawler config domains"], "body_bytes": 2316, "body_sha256": "sha256:20204ce7d8233332d3037f43c1fd4ab4ad55c2d63e138c5091d7194d11f742c3", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config", "path": "documentation/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2020021102120312-2200212233023101-3132313123002203-2033022002233221-2203111011013123-2301231332213210-0300212031232102-2233321232213203", "registry_path": "docs/guides/data-sources--third_party_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains"], "schema_version": 1, "sections": [{"aliases": ["authentication", "credential setup", "credentials", "domain"], "anchor": "schema-enable_api_discovery--api_crawler--api_crawler_config--domains--domain", "description": "Select the domain to execute API Crawling with given credentials.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["login", "login result", "sign in", "simple login"], "anchor": "section", "description": "Configuration parameter for simple login.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config.domains

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/)
- [enable_api_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/)
- [enable_api_discovery.api_crawler.api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/)
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

- [simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/): complete subsection reference.

## Next pages

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/)
- [enable_api_discovery.api_crawler.api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/)
- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
