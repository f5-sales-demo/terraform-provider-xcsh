---
page_title: "xcsh_namespace"
subcategory: ""
description: "Reads Namespace information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["namespace"], "body_bytes": 1303, "body_sha256": "sha256:91a30341b2487f3a57cc785515ede7b66407680f617ebd5bcd07fd3c30aa2c58", "capabilities": ["administration"], "category": "administration", "child_ids": ["xcsh-docs:data-sources:namespace:reference", "xcsh-docs:data-sources:namespace:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:namespace:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:namespace:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/namespace/index.md", "product": "distributed-cloud", "provider_name": "namespace", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-2013030003332230-3323112220123010-1133000333310232-3133330332200222-0123031332332100-3300310323320232-0203323000033132-1310310302002112", "registry_path": "docs/data-sources/namespace.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/namespace/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Reads Namespace information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["namespaceCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_namespace

Breadcrumbs:

- xcsh_namespace

Reads Namespace information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Namespace Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Credentials are supplied externally.
provider "xcsh" {}

# Look up an existing Namespace by name
data "xcsh_namespace" "example" {
  name = "example-namespace"
}

output "namespace_id" {
  value = data.xcsh_namespace.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/namespace/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/namespace/examples/)
