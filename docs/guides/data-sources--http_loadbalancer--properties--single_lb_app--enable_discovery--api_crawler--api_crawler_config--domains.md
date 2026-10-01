---
page_title: "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains"
subcategory: "Load Balancing"
description: "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3657, "body_sha256": "sha256:28b6eb5e73caf77eb8ff7b001805d2cfa0cd5db2ddd357906259af1a620d36b3", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains:simple_login"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config", "path": "docs/guides/data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config", "domains"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [single_lb_app](data-sources--http_loadbalancer--properties--single_lb_app.md)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery.md)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler.md)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config.md)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains

<a id="section"></a>

Type: `"list"`. Computed.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

## Direct properties

<a id="schema-single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--domain"></a>

### domain property

Type: `"string"`. Computed.

Select the domain to execute API Crawling with given credentials.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login.md): complete subsection reference.

## Next pages

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains--simple_login.md)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
