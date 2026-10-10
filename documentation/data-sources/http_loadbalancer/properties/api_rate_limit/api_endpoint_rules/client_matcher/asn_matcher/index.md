---
page_title: "api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher"
subcategory: "Load Balancing"
description: "Match any AS number contained in the list of bgp_asn_sets."
xcsh_docs: {"aliases": ["api rate limit api endpoint rules client matcher asn matcher"], "body_bytes": 1513, "body_sha256": "sha256:cbc112d12eece76eca1566398a16cb158bca52d5c9353929e4882bf7a5cf3672", "capabilities": ["load-balancing", "security.rate-limiting"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:asn_matcher:asn_sets"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:asn_matcher", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher", "path": "documentation/data-sources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/asn_matcher/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0330233120213000-3202021203001020-2331313312313330-0111303031020112-0200121231020322-2312023231332312-0320031122231211-2313101031013232", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher", "asn_matcher"], "schema_version": 1, "sections": [{"aliases": ["api rate limit api endpoint rules client matcher asn matcher asn sets"], "anchor": "section", "description": "A list of references to bgp_asn_set objects.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:api_endpoint_rules:client_matcher:asn_matcher:asn_sets", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_rate_limit", "api_endpoint_rules", "client_matcher", "asn_matcher", "asn_sets"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/asn_matcher/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Match any AS number contained in the list of bgp_asn_sets.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.api_endpoint_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/)
- [api_rate_limit.api_endpoint_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/)
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

- [asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/api_endpoint_rules/client_matcher/asn_matcher/asn_sets/): complete subsection reference.
