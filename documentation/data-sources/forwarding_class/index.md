---
page_title: "xcsh_forwarding_class"
subcategory: ""
description: "Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users in system namespace. configuration."
xcsh_docs: {"aliases": ["forwarding class"], "body_bytes": 1423, "body_sha256": "sha256:40470a43698f6fb4687cf5cedac7135ddbf3d93c72a4926df58dfd440710e45e", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:forwarding_class:reference", "xcsh-docs:data-sources:forwarding_class:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:forwarding_class:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forwarding_class:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/forwarding_class/index.md", "product": "distributed-cloud", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203", "registry_path": "docs/data-sources/forwarding_class.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forwarding_class/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users in system namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_forwarding_class

Breadcrumbs:

- xcsh_forwarding_class

Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users
in system namespace. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ForwardingClass Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ForwardingClass by name
data "xcsh_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "staging"
}

output "forwarding_class_id" {
  value = data.xcsh_forwarding_class.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/examples/)
