---
page_title: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login"
subcategory: "Load Balancing"
description: "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2927, "body_sha256": "sha256:abcc590f70bc6f0ea801fe60c37090914a8e2c7d8a3ee42fae491fcfd6c5fa72", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login:password"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains:simple_login", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains", "simple_login"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/simple_login/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [enable_api_discovery](data-sources--cdn_loadbalancer--properties--enable_api_discovery.md)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler.md)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains.md)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for simple login.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [password](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password.md): complete subsection reference.

<a id="schema-enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--user"></a>

### user property

Type: `"string"`. Computed.

Enter the username to assign credentials for the selected domain to crawl.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

## Next pages

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains--simple_login--password.md)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
