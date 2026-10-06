---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_dns_proxy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 969, "body_sha256": "sha256:d403ae5f74a6c1ad1caaa96e4a4678c2340e43b95e807fc1bf647f5fd27ba4d6", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:943728fe8d11ffeb7bc560a6a0e3b2ff777ab7fb03003f4fd0ef80073c303500", "source_path": "examples/resources/xcsh_dns_proxy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dns_proxy:example:resource", "parent_id": "xcsh-docs:resources:dns_proxy:examples", "path": "documentation/resources/dns_proxy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1023023213302301-1013332103111020-1223023020331033-1222112203110332-1121320322130100-1232001012320001-2002332201010123-1323321100230330", "registry_path": "docs/guides/resources--dns_proxy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/examples/resource/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Resource for xcsh_dns_proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_proxy/resource.tf`; digest `sha256:943728fe8d11ffeb7bc560a6a0e3b2ff777ab7fb03003f4fd0ef80073c303500`.

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
