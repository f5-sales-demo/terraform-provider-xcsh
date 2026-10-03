---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_dns_proxy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1182, "body_sha256": "sha256:91f6e804feb7cb194edd2ff66bdf9f9c129471df2ff02cf516b0196f5decb1c4", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:943728fe8d11ffeb7bc560a6a0e3b2ff777ab7fb03003f4fd0ef80073c303500", "source_path": "examples/resources/xcsh_dns_proxy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dns_proxy:example:resource", "parent_id": "xcsh-docs:resources:dns_proxy:examples", "path": "documentation/resources/dns_proxy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1023023213302301-1013332103111020-1223023020331033-1222112203110332-1121320322130100-1232001012320001-2002332201010123-1323321100230330", "registry_path": "docs/guides/resources--dns_proxy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_dns_proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/examples/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_proxy/)
