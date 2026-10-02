---
page_title: "api_rate_limit.custom_ip_allowed_list"
subcategory: "Load Balancing"
description: "IP Allowed list using existing ip_prefix_set objects."
xcsh_docs: {"aliases": ["api rate limit custom ip allowed list"], "body_bytes": 1604, "body_sha256": "sha256:77798a4fd7195b82b51b438b048ea9c546e6ff003e96f8d2d1df4a94791cd5e7", "capabilities": ["cdn", "security.rate-limiting"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list:rate_limiter_allowed_prefixes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit", "path": "documentation/data-sources/cdn_loadbalancer/properties/api_rate_limit/custom_ip_allowed_list/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-1010130033000023-0003001221210201-0012121132223103-0120030303211233-1033132223023030-1321021102203331-2100133012021110-2303033132131131", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "custom_ip_allowed_list"], "schema_version": 1, "sections": [{"aliases": ["rate limiter allowed prefixes"], "anchor": "section", "description": "References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list:rate_limiter_allowed_prefixes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_rate_limit", "custom_ip_allowed_list", "rate_limiter_allowed_prefixes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_rate_limit/custom_ip_allowed_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "IP Allowed list using existing ip_prefix_set objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.custom_ip_allowed_list

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/)
- api_rate_limit.custom_ip_allowed_list

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

- [rate_limiter_allowed_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/custom_ip_allowed_list/rate_limiter_allowed_prefixes/): complete subsection reference.

## Next pages

- [api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/custom_ip_allowed_list/rate_limiter_allowed_prefixes/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
