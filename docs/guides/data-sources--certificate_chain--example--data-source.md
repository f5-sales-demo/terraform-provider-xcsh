---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_certificate_chain."
xcsh_docs: {"aliases": [], "body_bytes": 1039, "body_sha256": "sha256:9dea8b68bd62c8ec9d8ab044aa1d0f63c971965169bbf13e7c1d65eaac1a7df3", "canonical_id": "xcsh-docs:data-sources:certificate_chain:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:certificate_chain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:af7fc7589403e5f8724fbb3e9b221e9a801d9fa6878ef97a8d7e1b6b86b1c396", "source_path": "examples/data-sources/xcsh_certificate_chain/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:certificate_chain:example:data-source", "parent_id": "xcsh-docs:data-sources:certificate_chain:examples", "path": "docs/guides/data-sources--certificate_chain--example--data-source.md", "provider_name": "certificate_chain", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certificate_chain/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_certificate_chain.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["certificate_chainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_certificate_chain](../data-sources/certificate_chain.md)
- [Examples](data-sources--certificate_chain--examples.md)
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

- [Examples](data-sources--certificate_chain--examples.md)
- [xcsh_certificate_chain](../data-sources/certificate_chain.md)
