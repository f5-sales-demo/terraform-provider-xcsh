---
page_title: "single_lb_app.enable_discovery.api_crawler.api_crawler_config"
subcategory: "Load Balancing"
description: "Crawler Configure."
xcsh_docs: {"aliases": ["single lb app enable discovery api crawler api crawler config"], "body_bytes": 2018, "body_sha256": "sha256:a48663db8a90403e25159abb1f7518d47d01cec40c3e80b4777de0710bea0bb0", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler", "path": "documentation/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0133120322312222-2312223333031212-1032032132030213-3301300112323232-2302012312301022-3110302310020121-2131323233322010-2212313212231211", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-026.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config"], "schema_version": 1, "sections": [{"aliases": ["authentication", "credential setup", "credentials", "domains"], "anchor": "section", "description": "Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config:domains", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config", "domains"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Crawler Configure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery.api_crawler.api_crawler_config

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [single_lb_app](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/)
- [single_lb_app.enable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/)
- [single_lb_app.enable_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config

<a id="section"></a>

Type: `"single"`. Computed.

Crawler Configure.

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

- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/): complete subsection reference.

## Next pages

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/domains/)
- [single_lb_app.enable_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
