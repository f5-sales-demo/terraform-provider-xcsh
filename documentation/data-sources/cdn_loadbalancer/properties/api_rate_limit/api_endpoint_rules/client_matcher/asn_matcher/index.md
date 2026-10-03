---
page_title: "api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher"
subcategory: "Load Balancing"
description: "Match any AS number contained in the list of bgp_asn_sets."
xcsh_docs: {"aliases": ["api rate limit api endpoint rules client matcher asn matcher"], "body_bytes": 2080, "body_sha256": "sha256:45fbf3eea734fa119a283f59d42c1493489752b1c0c5fd48aceab50a6dab0377", "capabilities": ["cdn", "security.rate-limiting"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:asn_matcher:asn_sets"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:asn_matcher", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher", "path": "documentation/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/asn_matcher/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0020001330021010-2211130231300323-0122200312322012-2110223200311233-3032113033302312-3121033320321300-0002013031200320-1030300231233020", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher", "asn_matcher"], "schema_version": 1, "sections": [{"aliases": ["api rate limit api endpoint rules client matcher asn matcher asn sets"], "anchor": "section", "description": "A list of references to bgp_asn_set objects.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:asn_matcher:asn_sets", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher", "asn_matcher", "asn_sets"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/asn_matcher/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Match any AS number contained in the list of bgp_asn_sets.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/)
- [api_rate_limit.api_endpoint_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

- [asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/asn_matcher/asn_sets/): complete subsection reference.

## Next pages

- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/asn_matcher/asn_sets/)
- [api_rate_limit.api_endpoint_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
