---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_certificate_chain."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1101, "body_sha256": "sha256:71497dac2edbd0114b5458956220984d4d4e119fa269bff52cbb2a7fda72976c", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:certificate_chain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:af7fc7589403e5f8724fbb3e9b221e9a801d9fa6878ef97a8d7e1b6b86b1c396", "source_path": "examples/data-sources/xcsh_certificate_chain/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:certificate_chain:example:data-source", "parent_id": "xcsh-docs:data-sources:certificate_chain:examples", "path": "documentation/data-sources/certificate_chain/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "certificate_chain", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3323010121233202-0202231203101310-0112230320210311-3001220210100210-2233220030222223-0131320003303010-3132002232133331-2202101302212032", "registry_path": "docs/guides/data-sources--certificate_chain--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certificate_chain/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_certificate_chain.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["certificate_chainCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_certificate_chain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate_chain/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate_chain/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_certificate_chain/data-source.tf`; digest `sha256:af7fc7589403e5f8724fbb3e9b221e9a801d9fa6878ef97a8d7e1b6b86b1c396`.

```terraform
# CertificateChain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CertificateChain by name
data "xcsh_certificate_chain" "example" {
  name      = "example-certificate-chain"
  namespace = "staging"
}

output "certificate_chain_id" {
  value = data.xcsh_certificate_chain.example.id
}
```
