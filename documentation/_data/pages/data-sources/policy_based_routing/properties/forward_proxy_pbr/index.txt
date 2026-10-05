---
page_title: "forward_proxy_pbr"
subcategory: ""
description: "Network(L3/L4) routing policy rule."
xcsh_docs: {"aliases": ["forward proxy pbr"], "body_bytes": 1860, "body_sha256": "sha256:53a2dec24234c685c71094eca24261b4404fba23baa0900a6e8464fde0e11a26", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr", "parent_id": "xcsh-docs:data-sources:policy_based_routing:reference", "path": "documentation/data-sources/policy_based_routing/properties/forward_proxy_pbr/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3223233112321212-1021022011013010-3323201332221223-2030103110313232-0332221032132101-0032001220122113-2313122122020010-3321223211310302", "registry_path": "docs/guides/data-sources--policy_based_routing--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["forward_proxy_pbr"], "schema_version": 1, "sections": [{"aliases": ["forward proxy pbr forward proxy pbr rules"], "anchor": "section", "description": "Network(L3/L4) routing policy rules.", "document_id": "xcsh-docs:data-sources:policy_based_routing:properties:forward_proxy_pbr:forward_proxy_pbr_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["forward_proxy_pbr", "forward_proxy_pbr_rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/forward_proxy_pbr/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Network(L3/L4) routing policy rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# forward_proxy_pbr

Breadcrumbs:

- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/)
- forward_proxy_pbr

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: forward\_proxy\_pbr, network\_pbr\] Configuration parameter for forward proxy pbr.

Upstream description:

Network(L3/L4) routing policy rule.

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

OneOf alternatives in this subsection:

- [forward_proxy_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/#section)
- [network_pbr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/network_pbr/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [forward_proxy_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/): complete subsection reference.

## Next pages

- [forward_proxy_pbr.forward_proxy_pbr_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/forward_proxy_pbr/forward_proxy_pbr_rules/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/)
- [xcsh_policy_based_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/)
