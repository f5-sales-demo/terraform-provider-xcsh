---
page_title: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info"
subcategory: ""
description: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info for xcsh_third_party_application."
xcsh_docs: {"aliases": [], "body_bytes": 2742, "body_sha256": "sha256:5cfc0002ec056911d33e20135f9c4e0af660dbf7084ffcca326bc20c0481e269", "canonical_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info", "child_ids": [], "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "path": "docs/guides/data-sources--third_party_application--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info.md", "provider_name": "third_party_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password", "clear_secret_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/clear_secret_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info for xcsh_third_party_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

Breadcrumbs:

- [xcsh_third_party_application](../data-sources/third_party_application.md)
- [Property reference](data-sources--third_party_application--reference.md)
- [enable_api_discovery](data-sources--third_party_application--properties--enable_api_discovery.md)
- [enable_api_discovery.api_crawler](data-sources--third_party_application--properties--enable_api_discovery--api_crawler.md)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--third_party_application--properties--enable_api_discovery--api_crawler--api_crawler_config.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--third_party_application--properties--enable_api_discovery--api_crawler--api_crawler_config--domains.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--third_party_application--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--third_party_application--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password.md)
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

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--third_party_application--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password.md)
- [xcsh_third_party_application](../data-sources/third_party_application.md)
