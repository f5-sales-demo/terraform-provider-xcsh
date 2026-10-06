---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_dns_zone_cryptokeys."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1004, "body_sha256": "sha256:5b90e62b89f87d3086d749d9b84a08dc03376bda02e1b824e16bccdfb4ac504a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:b7a0c08b1e74769c7e06fd735c8df6c0ed430cc8d5516a712859e61adfe28bc8", "source_path": "examples/data-sources/xcsh_dns_zone_cryptokeys/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dns_zone_cryptokeys:example:data-source", "parent_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:examples", "path": "documentation/data-sources/dns_zone_cryptokeys/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_cryptokeys", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0221323200003000-0201121202321133-1030213300123123-1121031321031132-0022112333000122-2121303111222113-3101312100231313-1123200101000122", "registry_path": "docs/guides/data-sources--dns_zone_cryptokeys--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone_cryptokeys/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_dns_zone_cryptokeys.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
