---
page_title: "response_cache.default_response_cache_parameters"
subcategory: "DNS"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["response cache default response cache parameters"], "body_bytes": 1327, "body_sha256": "sha256:8e29b2e94851355d57d81af4eb7e7c3907812cfc7b794d628740a6c772a5ec36", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache:default_response_cache_parameters", "parent_id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache", "path": "documentation/data-sources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0033011311220122-3232113303221303-1012003100001323-2223031022220221-3231103130102203-1333022100222212-3213121300313113-3331221303212220", "registry_path": "docs/guides/data-sources--dns_load_balancer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["response_cache", "default_response_cache_parameters"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# response_cache.default_response_cache_parameters

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/)
- [response_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/)
- response_cache.default_response_cache_parameters

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default response cache parameters.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [response_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/)
- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
