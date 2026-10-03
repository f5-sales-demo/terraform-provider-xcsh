---
page_title: "rate_limit.custom_ip_allowed_list"
subcategory: "Load Balancing"
description: "IP Allowed list using existing ip_prefix_set objects."
xcsh_docs: {"aliases": ["rate limit custom ip allowed list"], "body_bytes": 1568, "body_sha256": "sha256:6429d3f1c1aeeb3a936221ddb40e278968cbb002b9f09f42983cbc0497fbe98b", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:custom_ip_allowed_list:rate_limiter_allowed_prefixes"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:custom_ip_allowed_list", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit", "path": "documentation/data-sources/cdn_loadbalancer/properties/rate_limit/custom_ip_allowed_list/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2212331123312213-2313113110021301-1103033012023210-2113110012322031-0321031332032312-1023121113211323-0211010202011201-0101011300302321", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rate_limit", "custom_ip_allowed_list"], "schema_version": 1, "sections": [{"aliases": ["rate limit custom ip allowed list rate limiter allowed prefixes"], "anchor": "section", "description": "References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:custom_ip_allowed_list:rate_limiter_allowed_prefixes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rate_limit", "custom_ip_allowed_list", "rate_limiter_allowed_prefixes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/rate_limit/custom_ip_allowed_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "IP Allowed list using existing ip_prefix_set objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.custom_ip_allowed_list

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/)
- rate_limit.custom_ip_allowed_list

<a id="section"></a>

Type: `"single"`. Computed.

IP Allowed list using existing ip\_prefix\_set objects.

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

- [rate_limiter_allowed_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/custom_ip_allowed_list/rate_limiter_allowed_prefixes/): complete subsection reference.

## Next pages

- [rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/custom_ip_allowed_list/rate_limiter_allowed_prefixes/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
