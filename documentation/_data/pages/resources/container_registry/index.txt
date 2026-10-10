---
page_title: "xcsh_container_registry"
subcategory: "Container"
description: "Manages a Container Registry resource in F5 Distributed Cloud for container image registry configuration."
xcsh_docs: {"aliases": ["container registry"], "body_bytes": 1788, "body_sha256": "sha256:70f219898ea203491d0d313a40ee58d23e1bdf87291f90d3dcfe6a47e1118278", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:container_registry:reference", "xcsh-docs:resources:container_registry:examples", "xcsh-docs:resources:container_registry:import", "xcsh-docs:resources:container_registry:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:container_registry:collection", "completeness": "complete", "id": "xcsh-docs:resources:container_registry:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/container_registry/index.md", "product": "distributed-cloud", "provider_name": "container_registry", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3121031110330001-1021233032301331-1200313112012333-2030023322330002-2021231201222110-1023032102203110-1103221200322013-1102022332013332", "registry_path": "docs/resources/container_registry.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/container_registry/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Manages a Container Registry resource in F5 Distributed Cloud for container image registry configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["container_registryCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_container_registry

Breadcrumbs:

- xcsh_container_registry

Manages a Container Registry resource in F5 Distributed Cloud for container image registry
configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ContainerRegistry Resource Example
# Manages a Container Registry resource in F5 Distributed Cloud for container image registry configuration.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ContainerRegistry configuration
resource "xcsh_container_registry" "example" {
  name      = "example-container-registry"
  namespace = "staging"

  registry  = "example-value"
  user_name = "example-value"
}
```

## Root configuration

Required root properties: `name`, `namespace`, `registry`, `user_name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/lifecycle/timeouts/)
