---
page_title: "xcsh_authentication"
subcategory: ""
description: "Manages an Authentication resource in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials"], "body_bytes": 1537, "body_sha256": "sha256:273905230a99458a862af8bbb4167feaf6a45a8b4129ab71c629c4a18db1bb71", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:authentication:reference", "xcsh-docs:resources:authentication:examples", "xcsh-docs:resources:authentication:import", "xcsh-docs:resources:authentication:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/authentication/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210", "registry_path": "docs/resources/authentication.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Manages an Authentication resource in F5 Distributed Cloud.", "tasks": ["authentication", "configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["authenticationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_authentication

Breadcrumbs:

- xcsh_authentication

Manages an Authentication resource in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Authentication Resource Example
# Manages a Authentication resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Authentication configuration
resource "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/authentication/lifecycle/timeouts/)
