---
page_title: "rate_limit.rate_limiter.leaky_bucket"
subcategory: "Load Balancing"
description: "Leaky-Bucket is the default rate limiter algorithm for F5."
xcsh_docs: {"aliases": ["rate limit rate limiter leaky bucket"], "body_bytes": 1416, "body_sha256": "sha256:a03ecbdfc185761a9ca0fb1c40a234275454b29c44045ebecf9d11a6755476e5", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter:leaky_bucket", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:rate_limit:rate_limiter", "path": "documentation/resources/http_loadbalancer/properties/rate_limit/rate_limiter/leaky_bucket/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0123123131000310-3001121302331023-2333103110313033-2231001002030010-1032020221100330-0020022233032020-3333212003233230-1203231010321320", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-024.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rate_limit", "rate_limiter", "leaky_bucket"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/rate_limit/rate_limiter/leaky_bucket/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Leaky-Bucket is the default rate limiter algorithm for F5.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.rate_limiter.leaky_bucket

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/)
- [rate_limit.rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/rate_limiter/)
- rate_limit.rate_limiter.leaky_bucket

<a id="section"></a>

Type: `["object", {}]`. Optional.

Leaky-Bucket is the default rate limiter algorithm for F5.

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
leaky_bucket = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rate_limit.rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/rate_limit/rate_limiter/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
