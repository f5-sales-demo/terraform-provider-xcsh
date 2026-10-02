---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_dns_lb_pool."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1215, "body_sha256": "sha256:18068d9caa20bf099bed1c3427bbd88a1ca94d37f8db9c1c0c8172888d46326f", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a0f8bbd985a30a9aa4d49bcfaf3107e9c25c24df2c7857f8b4c1d06b0812d763", "source_path": "examples/resources/xcsh_dns_lb_pool/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dns_lb_pool:example:resource", "parent_id": "xcsh-docs:resources:dns_lb_pool:examples", "path": "documentation/resources/dns_lb_pool/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0131031232300121-2212232023211202-2122022301200203-0330133212020000-3310112012123003-2232233201322301-1220111133231000-0121033120100332", "registry_path": "docs/guides/resources--dns_lb_pool--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_dns_lb_pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_lb_pool/resource.tf`; digest `sha256:a0f8bbd985a30a9aa4d49bcfaf3107e9c25c24df2c7857f8b4c1d06b0812d763`.

```terraform
# DNSLBPool Resource Example
# Manages DNS Load Balancer Pool in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSLBPool configuration
resource "xcsh_dns_lb_pool" "example" {
  name      = "example-dns-lb-pool"
  namespace = "system"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/examples/)
- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
