---
page_title: "xcsh_healthcheck"
subcategory: "Monitoring"
description: "Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to determine if the given endpoint is healthy. single healthcheck object can be referred to by one or many cluster objects. configuration."
xcsh_docs: {"aliases": ["healthcheck"], "body_bytes": 1953, "body_sha256": "sha256:a3c1bc403d176c67859f552f10d83081258354a47a9be19cdb48989f602b9bad", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:healthcheck:reference", "xcsh-docs:resources:healthcheck:examples", "xcsh-docs:resources:healthcheck:import", "xcsh-docs:resources:healthcheck:timeouts"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:resources:healthcheck:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/healthcheck/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100", "registry_path": "docs/resources/healthcheck.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to determine if the given endpoint is healthy. single healthcheck object can be referred to by one or many cluster objects. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["healthcheckCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_healthcheck

Breadcrumbs:

- xcsh_healthcheck

Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to
determine if the given endpoint is healthy. single healthcheck object can be referred to by one or
many cluster objects. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Healthcheck Resource Example
# Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to determine if the given endpoint is healthy.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Healthcheck configuration
resource "xcsh_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "staging"

  healthy_threshold   = 1
  interval            = 1
  timeout             = 1
  unhealthy_threshold = 1
}
```

## Root configuration

Required root properties: `healthy_threshold`, `interval`, `name`, `namespace`, `timeout`, `unhealthy_threshold`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/lifecycle/timeouts/)
