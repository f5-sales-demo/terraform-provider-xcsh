---
page_title: "xcsh_certificate_chain"
subcategory: "Security"
description: "xcsh_certificate_chain for xcsh_certificate_chain."
xcsh_docs: {"aliases": [], "body_bytes": 1546, "body_sha256": "sha256:3dc4c8c46de9aff78836d9784672dcbf67b410621cef4cb02649ef1f82a96589", "canonical_id": "xcsh-docs:resources:certificate_chain:fundamentals", "child_ids": ["xcsh-docs:resources:certificate_chain:reference", "xcsh-docs:resources:certificate_chain:examples", "xcsh-docs:resources:certificate_chain:import", "xcsh-docs:resources:certificate_chain:timeouts"], "collection_id": "xcsh-docs:resources:certificate_chain:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate_chain:fundamentals", "parent_id": null, "path": "docs/resources/certificate_chain.md", "provider_name": "certificate_chain", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate_chain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_certificate_chain for xcsh_certificate_chain.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["certificate_chainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_certificate_chain

Breadcrumbs:

- xcsh_certificate_chain

Manages a Certificate Chain resource in F5 Distributed Cloud for certificate chain configuration for
TLS.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `certificate_url`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--certificate_chain--reference.md)
- [Examples](../guides/resources--certificate_chain--examples.md)
- [Import](../guides/resources--certificate_chain--import.md)
- [Timeouts](../guides/resources--certificate_chain--timeouts.md)
