---
page_title: "ring_hash"
subcategory: "Load Balancing"
description: "List of hash policy rules."
xcsh_docs: {"aliases": ["ring hash"], "body_bytes": 1296, "body_sha256": "sha256:ee30488db0f40b8c49ce772c87b66b4096ffaa1905256a472f59c303bc218b62", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:ring_hash:hash_policy"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:ring_hash", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/ring_hash/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3223003312110002-0221122213310322-3103023023321131-1323121100223120-2100103010223110-1231210111013100-0101202220221310-2300203220330323", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-023.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ring_hash"], "schema_version": 1, "sections": [{"aliases": ["ring hash hash policy"], "anchor": "section", "description": "Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated individually and the combined result is used to route the request.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:ring_hash:hash_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ring_hash", "hash_policy"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/ring_hash/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "List of hash policy rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ring_hash

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- ring_hash

<a id="section"></a>

Type: `"single"`. Computed.

Hash Policy List. List of hash policy rules.

Upstream description:

List of hash policy rules.

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

- [hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/ring_hash/hash_policy/): complete subsection reference.

## Next pages

- [ring_hash.hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/ring_hash/hash_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
