---
page_title: "xcsh_healthcheck"
subcategory: "Monitoring"
description: "Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to determine if the given endpoint is healthy. single healthcheck object can be referred to by one or many cluster objects. configuration."
xcsh_docs: {"aliases": ["healthcheck"], "body_bytes": 1966, "body_sha256": "sha256:810c1d97f28758d59f25fb8030a24690d194f8f3ad015b39f685623d25383ca8", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:healthcheck:reference", "xcsh-docs:resources:healthcheck:examples", "xcsh-docs:resources:healthcheck:import", "xcsh-docs:resources:healthcheck:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:resources:healthcheck:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/healthcheck/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3300320023110233-3202122221011131-2102110331222313-3011102021101332-1311231011303011-1030300103121111-0010331313321201-0112100103120100", "registry_path": "docs/resources/healthcheck.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to determine if the given endpoint is healthy. single healthcheck object can be referred to by one or many cluster objects. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/lifecycle/timeouts/)
