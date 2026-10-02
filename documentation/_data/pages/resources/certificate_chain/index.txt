---
page_title: "xcsh_certificate_chain"
subcategory: "Security"
description: "Manages a Certificate Chain resource in F5 Distributed Cloud for certificate chain configuration for TLS."
xcsh_docs: {"aliases": ["cert", "certificate", "certificate chain", "existing certificates", "tls certificates"], "body_bytes": 1735, "body_sha256": "sha256:988bd26b5a69dc241b4db1c5e13706027b65bb635606dd1a85485408315102b4", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:certificate_chain:reference", "xcsh-docs:resources:certificate_chain:examples", "xcsh-docs:resources:certificate_chain:import", "xcsh-docs:resources:certificate_chain:timeouts"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:certificate_chain:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate_chain:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/certificate_chain/index.md", "product": "distributed-cloud", "provider_name": "certificate_chain", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3300311221202132-1102121323122311-1122132320001300-2323321020311113-1011300012003132-2010130221021000-0203210321321112-0332331331210322", "registry_path": "docs/resources/certificate_chain.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate_chain/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Certificate Chain resource in F5 Distributed Cloud for certificate chain configuration for TLS.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["certificate_chainCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate_chain/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate_chain/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate_chain/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate_chain/lifecycle/timeouts/)
