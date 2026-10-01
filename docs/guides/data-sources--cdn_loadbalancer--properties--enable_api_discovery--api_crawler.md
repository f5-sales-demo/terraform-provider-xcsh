---
page_title: "enable_api_discovery.api_crawler"
subcategory: "Load Balancing"
description: "enable_api_discovery.api_crawler for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1635, "body_sha256": "sha256:77e913002855375dff288e04f3fee0c332e088f3f0eacbb76570902c6dda17f1", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config", "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:disable_api_crawler"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:enable_api_discovery", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery.api_crawler for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [enable_api_discovery](data-sources--cdn_loadbalancer--properties--enable_api_discovery.md)
- enable_api_discovery.api_crawler

<a id="section"></a>

Type: `"single"`. Computed.

API Crawling. API Crawler message.

Upstream description:

API Crawler message.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

## Direct properties

- [api_crawler_config](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config.md): complete subsection reference.

- [disable_api_crawler](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--disable_api_crawler.md): complete subsection reference.

## Next pages

- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config.md)
- [enable_api_discovery.api_crawler.disable_api_crawler](data-sources--cdn_loadbalancer--properties--enable_api_discovery--api_crawler--disable_api_crawler.md)
- [enable_api_discovery](data-sources--cdn_loadbalancer--properties--enable_api_discovery.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
