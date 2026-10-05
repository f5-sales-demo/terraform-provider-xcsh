---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_dns_lb_pool."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1263, "body_sha256": "sha256:1d1ccb52b05bd2392339df20e00e5a07b947e0fcf192db5c7088af5b6597ebf4", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_lb_pool:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9dddbd7b299906535a7d4a7f3de8f5eb6fae61134e8af1bbe57c031c456e2613", "source_path": "examples/data-sources/xcsh_dns_lb_pool/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dns_lb_pool:example:data-source", "parent_id": "xcsh-docs:data-sources:dns_lb_pool:examples", "path": "documentation/data-sources/dns_lb_pool/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3023200330133301-3130231230213011-0313212203321200-0110222110132101-1212200320230102-3323131102332331-1103313303213021-1033012000322312", "registry_path": "docs/guides/data-sources--dns_lb_pool--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_pool/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_dns_lb_pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_lb_pool/data-source.tf`; digest `sha256:9dddbd7b299906535a7d4a7f3de8f5eb6fae61134e8af1bbe57c031c456e2613`.

```terraform
# DNSLBPool Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSLBPool by name
data "xcsh_dns_lb_pool" "example" {
  name      = "example-dns-lb-pool"
  namespace = "system"
}

output "dns_lb_pool_id" {
  value = data.xcsh_dns_lb_pool.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/examples/)
- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/)
