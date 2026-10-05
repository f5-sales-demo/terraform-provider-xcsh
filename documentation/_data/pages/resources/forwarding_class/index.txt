---
page_title: "xcsh_forwarding_class"
subcategory: ""
description: "Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users in system namespace. configuration."
xcsh_docs: {"aliases": ["forwarding class"], "body_bytes": 1682, "body_sha256": "sha256:aeaeb49a9828af9b36b4922537ab28051c1f450bd732e1bd268701ed00dbb477", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:forwarding_class:reference", "xcsh-docs:resources:forwarding_class:examples", "xcsh-docs:resources:forwarding_class:import", "xcsh-docs:resources:forwarding_class:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:forwarding_class:collection", "completeness": "complete", "id": "xcsh-docs:resources:forwarding_class:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/forwarding_class/index.md", "product": "distributed-cloud", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212", "registry_path": "docs/resources/forwarding_class.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forwarding_class/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users in system namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
# ForwardingClass Resource Example
# Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ForwardingClass configuration
resource "xcsh_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forwarding_class/lifecycle/timeouts/)
