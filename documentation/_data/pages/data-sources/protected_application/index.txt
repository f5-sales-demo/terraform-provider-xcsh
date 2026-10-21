---
page_title: "xcsh_protected_application"
subcategory: ""
description: "Reads Protected Application information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["protected application"], "body_bytes": 1416, "body_sha256": "sha256:45ee92c1ce123af398cb8b8ed944a26bb02b8d539e6f0b1da103ff03ec88094f", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:reference", "xcsh-docs:data-sources:protected_application:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/protected_application/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022", "registry_path": "docs/data-sources/protected_application.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Reads Protected Application information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_protected_application

Breadcrumbs:

- xcsh_protected_application

Reads Protected Application information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtectedApplication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtectedApplication by name
data "xcsh_protected_application" "example" {
  name      = "example-protected-application"
  namespace = "staging"
}

output "protected_application_id" {
  value = data.xcsh_protected_application.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/examples/)
