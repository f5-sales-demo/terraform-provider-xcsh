---
page_title: "Data source"
subcategory: "Container"
description: "Data source for xcsh_container_registry."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1111, "body_sha256": "sha256:5b060fb0cb3373be7ec7b58fc3b52c0befde87d173968a9a418f3ba3b24b25ac", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:container_registry:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:235790f0ed1e99e7cff01491265f3f6071b78251c3ba60e8925b5d01fe4a9772", "source_path": "examples/data-sources/xcsh_container_registry/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:container_registry:example:data-source", "parent_id": "xcsh-docs:data-sources:container_registry:examples", "path": "documentation/data-sources/container_registry/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "container_registry", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1113010313311302-1131103033021122-2130031232233213-2020311233321332-0000303202001013-3203113012013120-1101210133102301-1303022020330121", "registry_path": "docs/guides/data-sources--container_registry--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/container_registry/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_container_registry.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["container_registryCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
