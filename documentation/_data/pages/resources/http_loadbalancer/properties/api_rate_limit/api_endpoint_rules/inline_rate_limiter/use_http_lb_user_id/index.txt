---
page_title: "api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["api rate limit api endpoint rules inline rate limiter use http lb user id"], "body_bytes": 1836, "body_sha256": "sha256:591135aaa5396aeea0bc9eb16a74cf9c7e64b5ab01cd59799ac8ee6912e498ac", "capabilities": ["load-balancing", "security.rate-limiting"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter:use_http_lb_user_id", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:inline_rate_limiter", "path": "documentation/resources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/inline_rate_limiter/use_http_lb_user_id/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3330031100020111-3030203233132221-1000321133200203-1321313131010003-3122213211133010-3300211303020132-1200031300300131-2210032020311321", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "api_endpoint_rules", "inline_rate_limiter", "use_http_lb_user_id"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/inline_rate_limiter/use_http_lb_user_id/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/inline_rate_limiter/)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
use_http_lb_user_id = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/inline_rate_limiter/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
