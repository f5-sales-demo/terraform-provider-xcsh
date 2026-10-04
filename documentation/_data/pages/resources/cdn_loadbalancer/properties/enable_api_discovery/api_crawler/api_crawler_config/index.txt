---
page_title: "enable_api_discovery.api_crawler.api_crawler_config"
subcategory: "Load Balancing"
description: "Crawler Configure."
xcsh_docs: {"aliases": ["enable api discovery api crawler api crawler config"], "body_bytes": 2015, "body_sha256": "sha256:5f145c716f99e4bfa8eb85890a4e5b463480d366486e9441b2065bed80c1509e", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler", "path": "documentation/resources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/api_crawler_config/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery.api_crawler.api_crawler_config:RequiredObjectAttributes:domains", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config"], "schema_version": 1, "sections": [{"aliases": ["enable api discovery api crawler api crawler config domains"], "anchor": "section", "description": "Enter domains and their credentials to allow authenticated API crawling. You can only include domains you own that are associated with this Load Balancer.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-enable_api_discovery--api_crawler--api_crawler_config--domains--domain", "enforcement": "provider-schema", "group": "enable_api_discovery.api_crawler.api_crawler_config.domains:RequiredListObjectAttributes:domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains", "type": "requires"}], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config", "domains"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/api_crawler_config/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Crawler Configure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler.api_crawler_config

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/)
- [enable_api_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/)
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

- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/): complete subsection reference.

## Next pages

- [enable_api_discovery.api_crawler.api_crawler_config.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/api_crawler_config/domains/)
- [enable_api_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
