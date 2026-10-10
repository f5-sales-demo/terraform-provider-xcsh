---
page_title: "rate_limit.policies"
subcategory: "Load Balancing"
description: "List of rate limiter policies to be applied."
xcsh_docs: {"aliases": ["rate limit policies"], "body_bytes": 987, "body_sha256": "sha256:2667fd5872302ae4b69e1a8b45c6bad3c84512d238edf1dafd87cf8cec7ef6cc", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:policies:policies"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:policies", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit", "path": "documentation/data-sources/cdn_loadbalancer/properties/rate_limit/policies/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2313310002111202-1332131303132010-2111300111031203-3121312022221032-3203333230000233-0300233203011031-2310330201113301-1120323203132331", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-015.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rate_limit", "policies"], "schema_version": 1, "sections": [{"aliases": ["rate limit policies policies"], "anchor": "section", "description": "Ordered list of rate limiter policies.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:rate_limit:policies:policies", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rate_limit", "policies", "policies"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/rate_limit/policies/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of rate limiter policies to be applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.policies

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/)
- rate_limit.policies

<a id="section"></a>

Type: `"single"`. Computed.

List of rate limiter policies to be applied.

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

- [policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/rate_limit/policies/policies/): complete subsection reference.
