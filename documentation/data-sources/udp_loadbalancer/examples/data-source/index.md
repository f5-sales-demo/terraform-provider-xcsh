---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1331, "body_sha256": "sha256:f5826d8a4041734daedf4076bc065b97e31b0590f2a71a246e6132bafe599f9d", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:udp_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:90436c5e0e42b785b65d7370ba89219e7c36c5236adccb1abeed194590bc9460", "source_path": "examples/data-sources/xcsh_udp_loadbalancer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:udp_loadbalancer:example:data-source", "parent_id": "xcsh-docs:data-sources:udp_loadbalancer:examples", "path": "documentation/data-sources/udp_loadbalancer/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0111330112313303-3132033201120320-3023232301303230-0011332301323002-0010010222331002-0222112320131300-2100121111310203-2021302010230113", "registry_path": "docs/guides/data-sources--udp_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/udp_loadbalancer/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_udp_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_udp_loadbalancer/data-source.tf`; digest `sha256:90436c5e0e42b785b65d7370ba89219e7c36c5236adccb1abeed194590bc9460`.

```terraform
# UDPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UDPLoadBalancer by name
data "xcsh_udp_loadbalancer" "example" {
  name      = "example-udp-loadbalancer"
  namespace = "staging"
}

output "udp_loadbalancer_id" {
  value = data.xcsh_udp_loadbalancer.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/examples/)
- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/)
