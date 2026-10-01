---
page_title: "xcsh_dns_lb_pool"
subcategory: ""
description: "xcsh_dns_lb_pool for xcsh_dns_lb_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1255, "body_sha256": "sha256:d7f518ce479b15a5219941ff8ff962f820810266e1fa71e5d76eb946f0dad0b8", "canonical_id": "xcsh-docs:data-sources:dns_lb_pool:fundamentals", "child_ids": ["xcsh-docs:data-sources:dns_lb_pool:reference", "xcsh-docs:data-sources:dns_lb_pool:examples"], "collection_id": "xcsh-docs:data-sources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_pool:fundamentals", "parent_id": null, "path": "docs/data-sources/dns_lb_pool.md", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_dns_lb_pool for xcsh_dns_lb_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_lb_pool

Breadcrumbs:

- xcsh_dns_lb_pool

Manages DNS Load Balancer Pool in a given namespace. If one already exist it will give a error in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--dns_lb_pool--reference.md)
- [Examples](../guides/data-sources--dns_lb_pool--examples.md)
