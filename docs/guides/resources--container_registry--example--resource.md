---
page_title: "Resource"
subcategory: "Container"
description: "Resource for xcsh_container_registry."
xcsh_docs: {"aliases": [], "body_bytes": 1203, "body_sha256": "sha256:83286bcd5fc9c157dcef876a7daa92e65e7d33898474698da9899a514680b245", "canonical_id": "xcsh-docs:resources:container_registry:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:container_registry:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:333d8100537c966abe2c10d8b88b6ea4b034dc0115a736110297064d1894694d", "source_path": "examples/resources/xcsh_container_registry/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:container_registry:example:resource", "parent_id": "xcsh-docs:resources:container_registry:examples", "path": "docs/guides/resources--container_registry--example--resource.md", "provider_name": "container_registry", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/container_registry/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_container_registry.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["container_registryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md)
- [Examples](resources--container_registry--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_container_registry/resource.tf`; digest `sha256:333d8100537c966abe2c10d8b88b6ea4b034dc0115a736110297064d1894694d`.

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

## Next pages

- [Examples](resources--container_registry--examples.md)
- [xcsh_container_registry](../resources/container_registry.md)
