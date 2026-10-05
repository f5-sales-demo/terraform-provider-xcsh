---
page_title: "xcsh_virtual_host"
subcategory: ""
description: "Manages virtual host in a given namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["virtual host"], "body_bytes": 1517, "body_sha256": "sha256:bb83266a12fa37c2858c21933052e7daa004f937046a4f76c3a7e9a055fff5c0", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:reference", "xcsh-docs:resources:virtual_host:examples", "xcsh-docs:resources:virtual_host:import", "xcsh-docs:resources:virtual_host:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/virtual_host/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222", "registry_path": "docs/resources/virtual_host.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages virtual host in a given namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_virtual_host

Breadcrumbs:

- xcsh_virtual_host

Manages virtual host in a given namespace in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualHost Resource Example
# Manages virtual host in a given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualHost configuration
resource "xcsh_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/lifecycle/timeouts/)
