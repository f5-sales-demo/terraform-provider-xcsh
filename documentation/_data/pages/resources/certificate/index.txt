---
page_title: "xcsh_certificate"
subcategory: "Security"
description: "Manages a Certificate resource in F5 Distributed Cloud for certificate. configuration."
xcsh_docs: {"aliases": ["certificate"], "body_bytes": 1624, "body_sha256": "sha256:cf1505bd48a1077fe8242798982529aa34e08e16eea3f28908956302aac90b6b", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:certificate:reference", "xcsh-docs:resources:certificate:examples", "xcsh-docs:resources:certificate:import", "xcsh-docs:resources:certificate:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/certificate/index.md", "product": "distributed-cloud", "provider_name": "certificate", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130", "registry_path": "docs/resources/certificate.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages a Certificate resource in F5 Distributed Cloud for certificate. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["certificateCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
