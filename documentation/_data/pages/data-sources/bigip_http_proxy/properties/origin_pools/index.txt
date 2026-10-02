---
page_title: "origin_pools"
subcategory: ""
description: "List of Origin Pools."
xcsh_docs: {"aliases": ["backend servers", "origin pools", "origin servers", "upstream servers"], "body_bytes": 1271, "body_sha256": "sha256:746eab2323367f61e67faf6534b5a22268090f3dace1c25fc1a6413fb846065b", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:reference", "path": "documentation/data-sources/bigip_http_proxy/properties/origin_pools/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pools"], "schema_version": 1, "sections": [{"aliases": ["backend servers", "origin servers", "pools", "upstream servers"], "anchor": "section", "description": "List of Origin Pools.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools:pools", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["origin_pools", "pools"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/origin_pools/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of Origin Pools.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- origin_pools

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for origin pools.

Upstream description:

List of Origin Pools.

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

- [pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/): complete subsection reference.

## Next pages

- [origin_pools.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
