---
page_title: "xcsh_protected_domain"
subcategory: ""
description: "Manages Domain to protect in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["protected domain"], "body_bytes": 1581, "body_sha256": "sha256:21a8b814fea895a48ae9d7f75773774f205fb20b07c2848f9340ce871e40ba73", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_domain:reference", "xcsh-docs:resources:protected_domain:examples", "xcsh-docs:resources:protected_domain:import", "xcsh-docs:resources:protected_domain:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_domain:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_domain:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/protected_domain/index.md", "product": "distributed-cloud", "provider_name": "protected_domain", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3133012330313310-3233013213001311-1212311200333023-2121112332020022-2232211202113213-1220123021122333-3111220011333221-3312301231323110", "registry_path": "docs/resources/protected_domain.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_domain/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages Domain to protect in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["protected_domainCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_protected_domain

Breadcrumbs:

- xcsh_protected_domain

Manages Domain to protect in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtectedDomain Resource Example
# Manages Domain to protect in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtectedDomain configuration
resource "xcsh_protected_domain" "example" {
  name      = "example-protected-domain"
  namespace = "staging"

  protected_domain = "example.com"
}
```

## Root configuration

Required root properties: `name`, `namespace`, `protected_domain`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_domain/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_domain/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_domain/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_domain/lifecycle/timeouts/)
