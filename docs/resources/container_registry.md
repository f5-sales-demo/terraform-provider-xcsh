---
page_title: "xcsh_container_registry"
subcategory: "Container"
description: "xcsh_container_registry for xcsh_container_registry."
xcsh_docs: {"aliases": [], "body_bytes": 1586, "body_sha256": "sha256:9003234e195daae699ec4ce5ab63f4b26db30f190ebe2e1501fe4eda573c4c09", "canonical_id": "xcsh-docs:resources:container_registry:fundamentals", "child_ids": ["xcsh-docs:resources:container_registry:reference", "xcsh-docs:resources:container_registry:examples", "xcsh-docs:resources:container_registry:import", "xcsh-docs:resources:container_registry:timeouts"], "collection_id": "xcsh-docs:resources:container_registry:collection", "completeness": "complete", "id": "xcsh-docs:resources:container_registry:fundamentals", "parent_id": null, "path": "docs/resources/container_registry.md", "provider_name": "container_registry", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/container_registry/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_container_registry for xcsh_container_registry.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["container_registryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Property reference](../guides/resources--container_registry--reference.md)
- [Examples](../guides/resources--container_registry--examples.md)
- [Import](../guides/resources--container_registry--import.md)
- [Timeouts](../guides/resources--container_registry--timeouts.md)
