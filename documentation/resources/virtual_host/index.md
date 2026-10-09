---
page_title: "xcsh_virtual_host"
subcategory: ""
description: "Manages virtual host in a given namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["virtual host"], "body_bytes": 1530, "body_sha256": "sha256:543840f1545f0212ae486ec3990e824088c8b1a827761f982ca9fb8cdcabd5e3", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:reference", "xcsh-docs:resources:virtual_host:examples", "xcsh-docs:resources:virtual_host:import", "xcsh-docs:resources:virtual_host:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/virtual_host/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222", "registry_path": "docs/resources/virtual_host.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Manages virtual host in a given namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/lifecycle/timeouts/)
