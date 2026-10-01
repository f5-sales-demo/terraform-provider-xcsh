---
page_title: "single_lb_app.enable_discovery.api_crawler"
subcategory: "Load Balancing"
description: "single_lb_app.enable_discovery.api_crawler for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2130, "body_sha256": "sha256:4bccfd12cca22b7d067439555afab3009d33230f103ccf68068b043a5d5124ce", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config", "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:disable_api_crawler"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery", "path": "docs/guides/resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_crawler"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "single_lb_app.enable_discovery.api_crawler for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery.api_crawler

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [single_lb_app](resources--http_loadbalancer--properties--single_lb_app.md)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--properties--single_lb_app--enable_discovery.md)
- single_lb_app.enable_discovery.api_crawler

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

API Crawling. API Crawler message.

Upstream description:

API Crawler message.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("api_crawler_config",
    "disable_api_crawler")}
```

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

Terraform syntax:

```terraform
api_crawler {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_crawler_config](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config.md): complete subsection reference.

- [disable_api_crawler](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--disable_api_crawler.md): complete subsection reference.

## Next pages

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config.md)
- [single_lb_app.enable_discovery.api_crawler.disable_api_crawler](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--disable_api_crawler.md)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--properties--single_lb_app--enable_discovery.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
