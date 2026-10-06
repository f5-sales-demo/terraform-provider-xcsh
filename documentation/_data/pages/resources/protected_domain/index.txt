---
page_title: "xcsh_protected_domain"
subcategory: ""
description: "Manages Domain to protect in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["protected domain"], "body_bytes": 1594, "body_sha256": "sha256:bad89081a80d65ca92751a586a05648ec8768881487a2654f5ba1ffc0d85e93d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_domain:reference", "xcsh-docs:resources:protected_domain:examples", "xcsh-docs:resources:protected_domain:import", "xcsh-docs:resources:protected_domain:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_domain:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_domain:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/protected_domain/index.md", "product": "distributed-cloud", "provider_name": "protected_domain", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3133012330313310-3233013213001311-1212311200333023-2121112332020022-2232211202113213-1220123021122333-3111220011333221-3312301231323110", "registry_path": "docs/resources/protected_domain.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_domain/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Manages Domain to protect in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["protected_domainCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_domain/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_domain/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_domain/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_domain/lifecycle/timeouts/)
