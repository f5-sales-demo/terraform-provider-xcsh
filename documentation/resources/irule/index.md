---
page_title: "xcsh_irule"
subcategory: ""
description: "Manages iRule in a given namespace. If one already exists it will give an error in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["irule"], "body_bytes": 1573, "body_sha256": "sha256:c915a943c05e0f63f80e5b3a7fe4b9345f704d2ff24d344d3feb1c8c6f82f2be", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:irule:reference", "xcsh-docs:resources:irule:examples", "xcsh-docs:resources:irule:import", "xcsh-docs:resources:irule:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:irule:collection", "completeness": "complete", "id": "xcsh-docs:resources:irule:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/irule/index.md", "product": "distributed-cloud", "provider_name": "irule", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2011333112331031-0020123312211213-0023313313132301-1232121102103020-3220220031303222-0310302031122301-3220110033323311-1131002100310203", "registry_path": "docs/resources/irule.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/irule/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Manages iRule in a given namespace. If one already exists it will give an error in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["iruleCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_irule

Breadcrumbs:

- xcsh_irule

Manages iRule in a given namespace. If one already exists it will give an error in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Irule Resource Example
# Manages iRule in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Irule configuration
resource "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"

  description_spec = "example-value"
  irule            = "example-value"
}
```

## Root configuration

Required root properties: `description_spec`, `irule`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/lifecycle/timeouts/)
