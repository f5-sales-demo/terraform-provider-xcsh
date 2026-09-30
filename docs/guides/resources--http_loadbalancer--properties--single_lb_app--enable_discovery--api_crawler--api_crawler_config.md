---
page_title: "single_lb_app.enable_discovery.api_crawler.api_crawler_config"
subcategory: "Load Balancing"
description: "single_lb_app.enable_discovery.api_crawler.api_crawler_config for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1717, "body_sha256": "sha256:54a5e887fd1ba9c7c300f37005a7c02166b92cf07150ca3cff34fffd17ca24d6", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler", "path": "docs/guides/resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "single_lb_app.enable_discovery.api_crawler.api_crawler_config for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# single_lb_app.enable_discovery.api_crawler.api_crawler_config

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [single_lb_app](resources--http_loadbalancer--properties--single_lb_app.md)
- [single_lb_app.enable_discovery](resources--http_loadbalancer--properties--single_lb_app--enable_discovery.md)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler.md)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config

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

- [domains](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains.md): complete subsection reference.

## Next pages

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler--api_crawler_config--domains.md)
- [single_lb_app.enable_discovery.api_crawler](resources--http_loadbalancer--properties--single_lb_app--enable_discovery--api_crawler.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
