---
page_title: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password"
subcategory: "Load Balancing"
description: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2812, "body_sha256": "sha256:9a537a5fa5409b83e2ca93b33322e03e9f8f3e53a8e505f2d90a2ba5418060f6", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:blindfold_secret_info", "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login", "password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [enable_api_discovery](data-sources--cdn_loadbalancer--properties--enable_api_discovery.md)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler.md)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login.md)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info.md): complete subsection reference.

## Next pages

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--blindfold_secret_info.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password--clear_secret_info.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
