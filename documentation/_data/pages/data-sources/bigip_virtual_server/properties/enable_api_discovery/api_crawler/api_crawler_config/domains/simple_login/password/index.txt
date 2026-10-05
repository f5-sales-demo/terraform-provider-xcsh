---
page_title: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["enable api discovery api crawler api crawler config domains simple login password", "login", "login result", "sign in"], "body_bytes": 3230, "body_sha256": "sha256:218a9c5b1678c274ba20ffa54f20509174a382863d2325af504f361e1ebdbed1", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login", "path": "documentation/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3232311233331103-3221332100300133-0222023110023330-3122312213121223-2210022221203310-0232322121012331-3230120001032002-1113311220013203", "registry_path": "docs/guides/data-sources--bigip_virtual_server--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password"], "schema_version": 1, "sections": [{"aliases": ["enable api discovery api crawler api crawler config domains simple login password blindfold secret info", "login", "login result", "sign in"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable api discovery api crawler api crawler config domains simple login password clear secret info", "login", "login result", "sign in"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

Breadcrumbs:

- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/)
- [enable_api_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/)
- [enable_api_discovery.api_crawler.api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/clear_secret_info/): complete subsection reference.

## Next pages

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/blindfold_secret_info/)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/clear_secret_info/)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/)
- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
