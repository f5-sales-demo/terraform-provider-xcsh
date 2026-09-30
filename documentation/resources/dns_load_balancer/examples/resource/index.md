---
page_title: "Resource"
subcategory: "DNS"
description: "Resource for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 1177, "body_sha256": "sha256:f4f8cd762066ac74f2ab4c44c922519db173f3f14317bfb55764b20e5390590d", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:103fabb6495b3c7befdb807ff9ebc6d11ed6c79f3f4a4083e4f2815709c3a01c", "source_path": "examples/resources/xcsh_dns_load_balancer/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dns_load_balancer:example:resource", "parent_id": "xcsh-docs:resources:dns_load_balancer:examples", "path": "documentation/resources/dns_load_balancer/examples/resource/index.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_load_balancer/resource.tf`; digest `sha256:103fabb6495b3c7befdb807ff9ebc6d11ed6c79f3f4a4083e4f2815709c3a01c`.

```terraform
# DNSLoadBalancer Resource Example
# Manages DNS Load Balancer in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSLoadBalancer configuration
resource "xcsh_dns_load_balancer" "example" {
  name      = "example-dns-load-balancer"
  namespace = "system"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/examples/)
- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
