---
page_title: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login"
subcategory: ""
description: "Configuration parameter for simple login."
xcsh_docs: {"aliases": ["enable api discovery api crawler api crawler config domains simple login", "login", "login result", "sign in"], "body_bytes": 1847, "body_sha256": "sha256:18ac0de011105c796e2675854ceb56547ba236634499ed3270df628aacaf0f4e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains", "path": "documentation/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0032223030331112-2210312110231233-3131010113321113-1322322131303303-1223133113030022-0211123333100211-0133201220032320-2102312101113132", "registry_path": "docs/guides/data-sources--bigip_virtual_server--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login"], "schema_version": 1, "sections": [{"aliases": ["enable api discovery api crawler api crawler config domains simple login password", "login", "login result", "sign in"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable api discovery api crawler api crawler config domains simple login user", "login", "login result", "sign in"], "anchor": "schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--user", "description": "Enter the username to assign credentials for the selected domain to crawl.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "user"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Configuration parameter for simple login.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

Breadcrumbs:

- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/)
- [enable_api_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/)
- [enable_api_discovery.api_crawler.api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for simple login.

## Direct properties

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/): complete subsection reference.

<a id="schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--user"></a>

### user property

Type: `"string"`. Computed.

Enter the username to assign credentials for the selected domain to crawl.
