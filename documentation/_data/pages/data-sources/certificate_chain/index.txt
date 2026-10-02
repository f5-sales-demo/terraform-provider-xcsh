---
page_title: "xcsh_certificate_chain"
subcategory: "Security"
description: "Manages a Certificate Chain resource in F5 Distributed Cloud for certificate chain configuration for TLS."
xcsh_docs: {"aliases": ["cert", "certificate", "certificate chain", "existing certificates", "tls certificates"], "body_bytes": 1436, "body_sha256": "sha256:ef41534ccb7635de979f6150c0ca9b5e6af1368257d127a0c7a80e8d5ee49aae", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:certificate_chain:reference", "xcsh-docs:data-sources:certificate_chain:examples"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:certificate_chain:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certificate_chain:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/certificate_chain/index.md", "product": "distributed-cloud", "provider_name": "certificate_chain", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3032033102223330-1001120313322113-3332030202222110-0322003202122313-3333310012310100-1102220222021332-3222113133323101-1133210103103321", "registry_path": "docs/data-sources/certificate_chain.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certificate_chain/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Certificate Chain resource in F5 Distributed Cloud for certificate chain configuration for TLS.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["certificate_chainCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate_chain/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certificate_chain/examples/)
