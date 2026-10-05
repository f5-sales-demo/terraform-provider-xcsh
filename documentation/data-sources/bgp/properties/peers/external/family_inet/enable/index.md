---
page_title: "peers.external.family_inet.enable"
subcategory: ""
description: "IPv4 Unicast."
xcsh_docs: {"aliases": ["peers external family inet enable"], "body_bytes": 1675, "body_sha256": "sha256:20036c1de1073d62c5e3c808e3e3226e0fd1f7cfbca9e032660adc0dd1418398", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet", "path": "documentation/data-sources/bgp/properties/peers/external/family_inet/enable/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1202221221321332-1012011021203233-0322011222210001-1211000120000333-0221213030222220-0313313120000111-2231112031231303-3222212323102132", "registry_path": "docs/guides/data-sources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "external", "family_inet", "enable"], "schema_version": 1, "sections": [{"aliases": ["peers external family inet enable aggregation"], "anchor": "section", "description": "BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing table and applies to outbound advertisements.", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["peers", "external", "family_inet", "enable", "aggregation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/external/family_inet/enable/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "IPv4 Unicast.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.family_inet.enable

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/)
- [peers.external.family_inet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/)
- peers.external.family_inet.enable

<a id="section"></a>

Type: `"single"`. Computed.

Unicast IPv4. IPv4 Unicast.

Upstream description:

IPv4 Unicast.

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

- [aggregation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/enable/aggregation/): complete subsection reference.

## Next pages

- [peers.external.family_inet.enable.aggregation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/enable/aggregation/)
- [peers.external.family_inet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
