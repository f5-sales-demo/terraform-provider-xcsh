---
page_title: "Data source"
subcategory: "DNS"
description: "Data source for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 1341, "body_sha256": "sha256:c5449b514e0bd6f8d938206f70a9148b911c59a71b9fad84b8ec9981958dd460", "child_ids": [], "collection_id": "xcsh-docs:data-sources:dns_load_balancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:94e6c40d8f676ae588608f07a90134d4cad49dd9a1bfd873d996f0534dfb0e30", "source_path": "examples/data-sources/xcsh_dns_load_balancer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dns_load_balancer:example:data-source", "parent_id": "xcsh-docs:data-sources:dns_load_balancer:examples", "path": "documentation/data-sources/dns_load_balancer/examples/data-source/index.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_load_balancer/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_load_balancer/data-source.tf`; digest `sha256:94e6c40d8f676ae588608f07a90134d4cad49dd9a1bfd873d996f0534dfb0e30`.

```terraform
# DNSLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSLoadBalancer by name
data "xcsh_dns_load_balancer" "example" {
  name      = "example-dns-load-balancer"
  namespace = "system"
}

output "dns_load_balancer_id" {
  value = data.xcsh_dns_load_balancer.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/examples/)
- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
