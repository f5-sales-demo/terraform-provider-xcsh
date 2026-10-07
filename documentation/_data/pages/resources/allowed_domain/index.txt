---
page_title: "xcsh_allowed_domain"
subcategory: ""
description: "Manages allowed domain in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["allowed domain"], "body_bytes": 1566, "body_sha256": "sha256:1daa7943fa2fe0f3e2dfebf4c62806ce2a632fd21890970d98f0e54adf30c3c7", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:allowed_domain:reference", "xcsh-docs:resources:allowed_domain:examples", "xcsh-docs:resources:allowed_domain:import", "xcsh-docs:resources:allowed_domain:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:allowed_domain:collection", "completeness": "complete", "id": "xcsh-docs:resources:allowed_domain:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/allowed_domain/index.md", "product": "distributed-cloud", "provider_name": "allowed_domain", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3123322023133102-1010022310110102-2030101233201213-3122230300233213-1133020330121130-0211003110101101-3301233223103200-0030203112113300", "registry_path": "docs/resources/allowed_domain.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/allowed_domain/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Manages allowed domain in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["allowed_domainCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/lifecycle/timeouts/)
