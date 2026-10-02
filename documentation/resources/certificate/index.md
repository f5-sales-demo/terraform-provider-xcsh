---
page_title: "xcsh_certificate"
subcategory: "Security"
description: "Manages a Certificate resource in F5 Distributed Cloud for certificate. configuration."
xcsh_docs: {"aliases": ["certificate"], "body_bytes": 1624, "body_sha256": "sha256:cf1505bd48a1077fe8242798982529aa34e08e16eea3f28908956302aac90b6b", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:certificate:reference", "xcsh-docs:resources:certificate:examples", "xcsh-docs:resources:certificate:import", "xcsh-docs:resources:certificate:timeouts"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/certificate/index.md", "product": "distributed-cloud", "provider_name": "certificate", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130", "registry_path": "docs/resources/certificate.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Certificate resource in F5 Distributed Cloud for certificate. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["certificateCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_certificate

Breadcrumbs:

- xcsh_certificate

Manages a Certificate resource in F5 Distributed Cloud for certificate. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Certificate Resource Example
# Manages a Certificate resource in F5 Distributed Cloud for certificate.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Certificate configuration
resource "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"

  certificate_url = "example-value"
}
```

## Root configuration

Required root properties: `certificate_url`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/lifecycle/timeouts/)
