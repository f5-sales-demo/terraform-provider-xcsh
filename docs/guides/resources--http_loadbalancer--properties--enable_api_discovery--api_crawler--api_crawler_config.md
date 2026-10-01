---
page_title: "enable_api_discovery.api_crawler.api_crawler_config"
subcategory: "Load Balancing"
description: "enable_api_discovery.api_crawler.api_crawler_config for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1623, "body_sha256": "sha256:410531062cfeb7b0fd76a6f2a94131d82ecaed2071367af623eb09343ac693e4", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:enable_api_discovery:api_crawler", "path": "docs/guides/resources--http_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/enable_api_discovery/api_crawler/api_crawler_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_api_discovery.api_crawler.api_crawler_config for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [enable_api_discovery](resources--http_loadbalancer--properties--enable_api_discovery.md)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--properties--enable_api_discovery--api_crawler.md)
- enable_api_discovery.api_crawler.api_crawler_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Crawler Configure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
```

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

Terraform syntax:

```terraform
api_crawler_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [domains](resources--http_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains.md): complete subsection reference.

## Next pages

- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--properties--enable_api_discovery--api_crawler--api_crawler_config--domains.md)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--properties--enable_api_discovery--api_crawler.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
