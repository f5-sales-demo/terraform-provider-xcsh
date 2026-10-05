---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_dns_zone_cryptokeys."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1253, "body_sha256": "sha256:8061bfd903f1a6e7b459770893d156f65908539cb9cf7276726f3dc584362f3c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:b7a0c08b1e74769c7e06fd735c8df6c0ed430cc8d5516a712859e61adfe28bc8", "source_path": "examples/data-sources/xcsh_dns_zone_cryptokeys/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dns_zone_cryptokeys:example:data-source", "parent_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:examples", "path": "documentation/data-sources/dns_zone_cryptokeys/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_cryptokeys", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0221323200003000-0201121202321133-1030213300123123-1121031321031132-0022112333000122-2121303111222113-3101312100231313-1123200101000122", "registry_path": "docs/guides/data-sources--dns_zone_cryptokeys--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone_cryptokeys/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_dns_zone_cryptokeys.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_dns_zone_cryptokeys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_zone_cryptokeys/data-source.tf`; digest `sha256:b7a0c08b1e74769c7e06fd735c8df6c0ed430cc8d5516a712859e61adfe28bc8`.

```terraform
# DNSZoneCryptokeys DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_dns_zone_cryptokeys" "example" {
}

output "dns_zone_cryptokeys_result" {
  value = data.xcsh_dns_zone_cryptokeys.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/examples/)
- [xcsh_dns_zone_cryptokeys](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/)
