---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1355, "body_sha256": "sha256:37f21768288e516d2b92a6ef06b3774270bb76c306adc564b49a94e6624800e8", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp_routing_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:b3d6f35be703edce6abe0703f58285587174e0deae9414b75725a102a2aa63cb", "source_path": "examples/data-sources/xcsh_bgp_routing_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bgp_routing_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:bgp_routing_policy:examples", "path": "documentation/data-sources/bgp_routing_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1132013131121233-1203223013000330-1030133200231322-2301021203322013-0001012113023232-1332301210223102-3231101103003203-2222101010212112", "registry_path": "docs/guides/data-sources--bgp_routing_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp_routing_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_bgp_routing_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bgp_routing_policy/data-source.tf`; digest `sha256:b3d6f35be703edce6abe0703f58285587174e0deae9414b75725a102a2aa63cb`.

```terraform
# BGPRoutingPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGPRoutingPolicy by name
data "xcsh_bgp_routing_policy" "example" {
  name      = "example-bgp-routing-policy"
  namespace = "staging"
}

output "bgp_routing_policy_id" {
  value = data.xcsh_bgp_routing_policy.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/examples/)
- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/)
