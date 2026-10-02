---
page_title: "caching_policy.custom_cache_rule"
subcategory: "Load Balancing"
description: "Caching policies for CDN."
xcsh_docs: {"aliases": ["caching policy custom cache rule"], "body_bytes": 1572, "body_sha256": "sha256:e1fe997cf7e2cf4034675ee3ed21217347b20a284bbfa7b65b816388667fd61c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy:custom_cache_rule:cdn_cache_rules"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy:custom_cache_rule", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy", "path": "documentation/data-sources/http_loadbalancer/properties/caching_policy/custom_cache_rule/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2103101010232231-0110012110211132-0322322122312321-3300211033231023-1310100312303203-2311312023231103-2232010032031201-0312132102010220", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["caching_policy", "custom_cache_rule"], "schema_version": 1, "sections": [{"aliases": ["cdn cache rules"], "anchor": "section", "description": "Reference to CDN Cache Rule configuration object.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy:custom_cache_rule:cdn_cache_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["caching_policy", "custom_cache_rule", "cdn_cache_rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/caching_policy/custom_cache_rule/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Caching policies for CDN.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# caching_policy.custom_cache_rule

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [caching_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/caching_policy/)
- caching_policy.custom_cache_rule

<a id="section"></a>

Type: `"single"`. Computed.

Custom Cache Rules. Caching policies for CDN.

Upstream description:

Caching policies for CDN.

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

- [cdn_cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/caching_policy/custom_cache_rule/cdn_cache_rules/): complete subsection reference.

## Next pages

- [caching_policy.custom_cache_rule.cdn_cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/caching_policy/custom_cache_rule/cdn_cache_rules/)
- [caching_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/caching_policy/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
