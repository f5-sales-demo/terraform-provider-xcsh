---
page_title: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info"
subcategory: ""
description: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info for xcsh_third_party_application."
xcsh_docs: {"aliases": [], "body_bytes": 3046, "body_sha256": "sha256:ee0421cb422bb8ebb665af2bc8fc463d961727e4077d4ab4521ee5cc576cfe33", "canonical_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "child_ids": [], "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "path": "docs/guides/data-sources--third_party_application--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info.md", "provider_name": "third_party_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "blindfold_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/blindfold_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info for xcsh_third_party_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

Breadcrumbs:

- [xcsh_third_party_application](../data-sources/third_party_application.md)
- [Property reference](data-sources--third_party_application--reference.md)
- [enable_api_discovery](data-sources--third_party_application--properties--enable_api_discovery.md)
- [enable_api_discovery.api_crawler](data-sources--third_party_application--properties--enable_api_discovery--api_crawler.md)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--third_party_application--properties--enable_api_discovery--api_crawler--api_crawler_config.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--third_party_application--properties--enable_api_discovery--api_crawler--api_crawler_config--domains.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--third_party_application--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--third_party_application--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password.md)
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

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--third_party_application--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password.md)
- [xcsh_third_party_application](../data-sources/third_party_application.md)
