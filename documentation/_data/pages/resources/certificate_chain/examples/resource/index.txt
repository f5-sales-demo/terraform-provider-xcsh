---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_certificate_chain."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1374, "body_sha256": "sha256:98d696b6c2248ffbf2abb2b602f70eb75a6ef54d63f5b733dc55f42fcae4f484", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:certificate_chain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:b97518ab6ca31864cd9125be1aa4c371e7fb2f6d12ef9376188442d135d4d743", "source_path": "examples/resources/xcsh_certificate_chain/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:certificate_chain:example:resource", "parent_id": "xcsh-docs:resources:certificate_chain:examples", "path": "documentation/resources/certificate_chain/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "certificate_chain", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2310023133111213-3302123200210231-1100230003202222-2213202211121212-0221200000102112-2111311200202031-0013202200033210-1221312113320221", "registry_path": "docs/guides/resources--certificate_chain--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate_chain/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_certificate_chain.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["certificate_chainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_certificate_chain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate_chain/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate_chain/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_certificate_chain/resource.tf`; digest `sha256:b97518ab6ca31864cd9125be1aa4c371e7fb2f6d12ef9376188442d135d4d743`.

```terraform
# CertificateChain Resource Example
# Manages a Certificate Chain resource in F5 Distributed Cloud for certificate chain configuration for TLS.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CertificateChain configuration
resource "xcsh_certificate_chain" "example" {
  name      = "example-certificate-chain"
  namespace = "staging"

  certificate_url = "example-value"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate_chain/examples/)
- [xcsh_certificate_chain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate_chain/)
