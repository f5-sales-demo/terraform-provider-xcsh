---
page_title: "Data source"
subcategory: "Container"
description: "Data source for xcsh_container_registry."
xcsh_docs: {"aliases": [], "body_bytes": 1258, "body_sha256": "sha256:395bb6aff28134f7b834108b2ce186cfb005753c50a9d958ce58aa137abc70ae", "child_ids": [], "collection_id": "xcsh-docs:data-sources:container_registry:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:235790f0ed1e99e7cff01491265f3f6071b78251c3ba60e8925b5d01fe4a9772", "source_path": "examples/data-sources/xcsh_container_registry/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:container_registry:example:data-source", "parent_id": "xcsh-docs:data-sources:container_registry:examples", "path": "documentation/data-sources/container_registry/examples/data-source/index.md", "provider_name": "container_registry", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/container_registry/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_container_registry.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["container_registryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_container_registry](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/container_registry/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/container_registry/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_container_registry/data-source.tf`; digest `sha256:235790f0ed1e99e7cff01491265f3f6071b78251c3ba60e8925b5d01fe4a9772`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/container_registry/examples/)
- [xcsh_container_registry](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/container_registry/)
