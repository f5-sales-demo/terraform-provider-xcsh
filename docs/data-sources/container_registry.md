---
page_title: "xcsh_container_registry"
subcategory: "Container"
description: "xcsh_container_registry for xcsh_container_registry."
xcsh_docs: {"aliases": [], "body_bytes": 1262, "body_sha256": "sha256:11e8ac9cffe5a5768743ecd6d4da2108d98d2263d62b0a5356a094faa9a083b4", "canonical_id": "xcsh-docs:data-sources:container_registry:fundamentals", "child_ids": ["xcsh-docs:data-sources:container_registry:reference", "xcsh-docs:data-sources:container_registry:examples"], "collection_id": "xcsh-docs:data-sources:container_registry:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:container_registry:fundamentals", "parent_id": null, "path": "docs/data-sources/container_registry.md", "provider_name": "container_registry", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/container_registry/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_container_registry for xcsh_container_registry.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["container_registryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
# ContainerRegistry Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ContainerRegistry by name
data "xcsh_container_registry" "example" {
  name      = "example-container-registry"
  namespace = "staging"
}

output "container_registry_id" {
  value = data.xcsh_container_registry.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--container_registry--reference.md)
- [Examples](../guides/data-sources--container_registry--examples.md)
