---
page_title: "enable_api_discovery.api_crawler"
subcategory: "Load Balancing"
description: "API Crawler message."
xcsh_docs: {"aliases": ["enable api discovery api crawler"], "body_bytes": 2368, "body_sha256": "sha256:3d75b0810e31a39da8e61db66c56b9c54df813488908fb295f3e44d145cb55a3", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config", "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:disable_api_crawler"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery", "path": "documentation/resources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery.api_crawler:ConflictingObjectAttributes:api_crawler_config,disable_api_crawler", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery.api_crawler:ConflictingObjectAttributes:api_crawler_config,disable_api_crawler", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:disable_api_crawler", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "api_crawler"], "schema_version": 1, "sections": [{"aliases": ["api crawler config"], "anchor": "section", "description": "Crawler Configure.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery.api_crawler.api_crawler_config:RequiredObjectAttributes:domains", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:api_crawler_config:domains", "type": "requires"}], "schema_path": ["enable_api_discovery", "api_crawler", "api_crawler_config"], "syntax": "block", "type": "object"}, {"aliases": ["disable api crawler"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_crawler:disable_api_crawler", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler", "disable_api_crawler"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "API Crawler message.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_crawler

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/)
- enable_api_discovery.api_crawler

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

- [api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/api_crawler_config/): complete subsection reference.

- [disable_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/disable_api_crawler/): complete subsection reference.

## Next pages

- [enable_api_discovery.api_crawler.api_crawler_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/api_crawler_config/)
- [enable_api_discovery.api_crawler.disable_api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_crawler/disable_api_crawler/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
