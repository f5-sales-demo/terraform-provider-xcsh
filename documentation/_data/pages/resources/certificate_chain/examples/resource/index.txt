---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_certificate_chain."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1137, "body_sha256": "sha256:3d83a6f7e02e7ca193125a4ce32b1befcb61b9884ff189aa3aa42e67b339dfa6", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:certificate_chain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:b97518ab6ca31864cd9125be1aa4c371e7fb2f6d12ef9376188442d135d4d743", "source_path": "examples/resources/xcsh_certificate_chain/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:certificate_chain:example:resource", "parent_id": "xcsh-docs:resources:certificate_chain:examples", "path": "documentation/resources/certificate_chain/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "certificate_chain", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-2310023133111213-3302123200210231-1100230003202222-2213202211121212-0221200000102112-2111311200202031-0013202200033210-1221312113320221", "registry_path": "docs/guides/resources--certificate_chain--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate_chain/examples/resource/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Resource for xcsh_certificate_chain.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["certificate_chainCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
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
