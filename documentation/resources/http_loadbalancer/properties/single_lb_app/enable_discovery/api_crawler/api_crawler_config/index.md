---
page_title: "single_lb_app.enable_discovery.api_crawler.api_crawler_config"
subcategory: "Load Balancing"
description: "single_lb_app.enable_discovery.api_crawler.api_crawler_config for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2165, "body_sha256": "sha256:9150989608af819da3e13c9a8d94e5b4bf1be6cafcb0459e7ca6f063809022cd", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler", "path": "documentation/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "single_lb_app.enable_discovery.api_crawler.api_crawler_config for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# single_lb_app.enable_discovery.api_crawler.api_crawler_config

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [single_lb_app](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/)
- [single_lb_app.enable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/)
- [single_lb_app.enable_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/)
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

- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/): complete subsection reference.

## Next pages

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/)
- [single_lb_app.enable_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
