---
page_title: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info"
subcategory: ""
description: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info for xcsh_bigip_virtual_server."
xcsh_docs: {"aliases": [], "body_bytes": 3010, "body_sha256": "sha256:6e0e206a139a6e9a596badd36dc598121759ecbd0b9078179dd6fbb924374d23", "canonical_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "path": "docs/guides/data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info.md", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "blindfold_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/blindfold_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info for xcsh_bigip_virtual_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

Breadcrumbs:

- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md)
- [Property reference](data-sources--bigip_virtual_server--reference.md)
- [enable_api_discovery](data-sources--bigip_virtual_server--properties--enable_api_discovery.md)
- [enable_api_discovery.api_crawler](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler.md)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password.md)
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

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--bigip_virtual_server--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password.md)
- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md)
