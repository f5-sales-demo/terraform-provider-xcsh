---
page_title: "xcsh_allowed_domain"
subcategory: ""
description: "Manages allowed domain in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["allowed domain"], "body_bytes": 1553, "body_sha256": "sha256:7cc57e0d99e02f77b2ff2e36a79d3af49d381176e4e6e686ebbf2374d0427abb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:allowed_domain:reference", "xcsh-docs:resources:allowed_domain:examples", "xcsh-docs:resources:allowed_domain:import", "xcsh-docs:resources:allowed_domain:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:allowed_domain:collection", "completeness": "complete", "id": "xcsh-docs:resources:allowed_domain:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/allowed_domain/index.md", "product": "distributed-cloud", "provider_name": "allowed_domain", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3123322023133102-1010022310110102-2030101233201213-3122230300233213-1133020330121130-0211003110101101-3301233223103200-0030203112113300", "registry_path": "docs/resources/allowed_domain.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/allowed_domain/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages allowed domain in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["allowed_domainCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_allowed_domain

Breadcrumbs:

- xcsh_allowed_domain

Manages allowed domain in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AllowedDomain Resource Example
# Manages allowed domain in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AllowedDomain configuration
resource "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"

  allowed_domain = "example-value"
}
```

## Root configuration

Required root properties: `allowed_domain`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/lifecycle/timeouts/)
