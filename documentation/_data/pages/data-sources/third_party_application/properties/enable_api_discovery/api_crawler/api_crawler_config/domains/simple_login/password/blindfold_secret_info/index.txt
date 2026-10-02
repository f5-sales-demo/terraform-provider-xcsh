---
page_title: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info"
subcategory: ""
description: "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management."
xcsh_docs: {"aliases": ["enable api discovery api crawler api crawler config domains simple login password blindfold secret info", "login", "login result", "sign in"], "body_bytes": 3538, "body_sha256": "sha256:850fc0a898c9af90ea8f85e2e8ac46bd78f701044a88394a62dbd3ed48b28e6e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "path": "documentation/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/blindfold_secret_info/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1211213132212113-1101110112212003-2130122322003011-0110130002333022-3310322101333012-0102020211000021-1012003222202101-2231033303013000", "registry_path": "docs/guides/data-sources--third_party_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "blindfold_secret_info"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "decryption provider", "origin servers", "upstream servers"], "anchor": "schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info--decryption_provider", "description": "Name of the Secret Management Access object that contains information about the backend Secret Management service.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "blindfold_secret_info", "decryption_provider"], "syntax": "attribute", "type": "string"}, {"aliases": ["location"], "anchor": "schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info--location", "description": "Location is the uri_ref. It could be in URL format for string:/// Or it could be a path if the store provider is an HTTP/HTTPS location.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "flags": ["computed", "sensitive"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "blindfold_secret_info", "location"], "syntax": "attribute", "type": "string"}, {"aliases": ["store provider"], "anchor": "schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info--store_provider", "description": "Name of the Secret Management Access object that contains information about the store to GET encrypted bytes This field needs to be provided only if the URL scheme is not string:///.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "blindfold_secret_info", "store_provider"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/blindfold_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/)
- [enable_api_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/)
- [enable_api_discovery.api_crawler.api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="section"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

## Direct properties

<a id="schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info--decryption_provider"></a>

### decryption_provider property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

<a id="schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info--location"></a>

### location property

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

<a id="schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info--store_provider"></a>

### store_provider property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

## Next pages

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/)
- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
