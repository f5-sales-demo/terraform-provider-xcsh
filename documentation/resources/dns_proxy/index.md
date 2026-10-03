---
page_title: "xcsh_dns_proxy"
subcategory: ""
description: "Manages DNS Proxy in a given namespace. If one already exists it will give an error in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["dns proxy"], "body_bytes": 1488, "body_sha256": "sha256:81ef73751c2c64fe97a8c2ef856d2d9a67e78123f80ccb50029d78b41bfcc026", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_proxy:reference", "xcsh-docs:resources:dns_proxy:examples", "xcsh-docs:resources:dns_proxy:import", "xcsh-docs:resources:dns_proxy:timeouts"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/dns_proxy/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122", "registry_path": "docs/resources/dns_proxy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages DNS Proxy in a given namespace. If one already exists it will give an error in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_proxy

Breadcrumbs:

- xcsh_dns_proxy

Manages DNS Proxy in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSProxy Resource Example
# Manages DNS Proxy in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSProxy configuration
resource "xcsh_dns_proxy" "example" {
  name      = "example-dns-proxy"
  namespace = "system"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/lifecycle/timeouts/)
