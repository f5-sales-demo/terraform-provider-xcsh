---
page_title: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info"
subcategory: ""
description: "ClearSecretInfoType specifies information about the Secret that is not encrypted."
xcsh_docs: {"aliases": ["enable api discovery api crawler api crawler config domains simple login password clear secret info", "login", "login result", "sign in"], "body_bytes": 3234, "body_sha256": "sha256:d1e7fd2fe18d6e977484673d0330df6515bba23f25fab49b271dec2b593bb290", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "path": "documentation/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/clear_secret_info/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2033312330322321-2303002322322133-2003212200001131-2010011000013231-1322020221023102-3210032111030312-2331311200331110-3320333223001003", "registry_path": "docs/guides/data-sources--third_party_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "clear_secret_info"], "schema_version": 1, "sections": [{"aliases": ["enable api discovery api crawler api crawler config domains simple login password clear secret info provider ref", "login", "login result", "sign in"], "anchor": "schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info--provider_ref", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "clear_secret_info", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["enable api discovery api crawler api crawler config domains simple login password clear secret info url", "login", "login result", "sign in"], "anchor": "schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info--url", "description": "URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after Base64 decoding.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "clear_secret_info", "url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/clear_secret_info/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/)
- [enable_api_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/)
- [enable_api_discovery.api_crawler.api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="section"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

## Direct properties

<a id="schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info--url"></a>

### url property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

## Next pages

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/)
- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
