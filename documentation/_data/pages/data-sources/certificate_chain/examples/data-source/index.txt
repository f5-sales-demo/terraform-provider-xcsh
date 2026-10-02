---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_certificate_chain."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1344, "body_sha256": "sha256:1e8f375dd3f7e2ec991d0d9ba91e9a88d8d7cb2b8aa172c674660fbe818da014", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:certificate_chain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:af7fc7589403e5f8724fbb3e9b221e9a801d9fa6878ef97a8d7e1b6b86b1c396", "source_path": "examples/data-sources/xcsh_certificate_chain/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:certificate_chain:example:data-source", "parent_id": "xcsh-docs:data-sources:certificate_chain:examples", "path": "documentation/data-sources/certificate_chain/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "certificate_chain", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3323010121233202-0202231203101310-0112230320210311-3001220210100210-2233220030222223-0131320003303010-3132002232133331-2202101302212032", "registry_path": "docs/guides/data-sources--certificate_chain--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certificate_chain/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_certificate_chain.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["certificate_chainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate_chain/examples/)
- [xcsh_certificate_chain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate_chain/)
