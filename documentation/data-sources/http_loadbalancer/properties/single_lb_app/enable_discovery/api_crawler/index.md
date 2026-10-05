---
page_title: "single_lb_app.enable_discovery.api_crawler"
subcategory: "Load Balancing"
description: "API Crawler message."
xcsh_docs: {"aliases": ["single lb app enable discovery api crawler"], "body_bytes": 2351, "body_sha256": "sha256:b674e0d06b8697d885bcce72fafad550a753ffe9425562022695234e3ddea3a2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config", "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:disable_api_crawler"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery", "path": "documentation/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1122300132122203-1113231131101200-0021200301202132-2030211130001200-3100101010012101-1120303001023123-3223212121001201-1011322201031133", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-026.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["single_lb_app", "enable_discovery", "api_crawler"], "schema_version": 1, "sections": [{"aliases": ["single lb app enable discovery api crawler api crawler config"], "anchor": "section", "description": "Crawler Configure.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:api_crawler_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "api_crawler_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["single lb app enable discovery api crawler disable api crawler"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:single_lb_app:enable_discovery:api_crawler:disable_api_crawler", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["single_lb_app", "enable_discovery", "api_crawler", "disable_api_crawler"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "API Crawler message.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# single_lb_app.enable_discovery.api_crawler

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [single_lb_app](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/)
- [single_lb_app.enable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/)
- single_lb_app.enable_discovery.api_crawler

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

- [api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/): complete subsection reference.

- [disable_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/disable_api_crawler/): complete subsection reference.

## Next pages

- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/api_crawler_config/)
- [single_lb_app.enable_discovery.api_crawler.disable_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/api_crawler/disable_api_crawler/)
- [single_lb_app.enable_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/single_lb_app/enable_discovery/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
