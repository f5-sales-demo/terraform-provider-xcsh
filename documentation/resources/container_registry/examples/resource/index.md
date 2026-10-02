---
page_title: "Resource"
subcategory: "Container"
description: "Resource for xcsh_container_registry."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1409, "body_sha256": "sha256:ef042c989c9b631759b4ce169bdda8071415063d55ac71012747575edbf3df6b", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:container_registry:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:333d8100537c966abe2c10d8b88b6ea4b034dc0115a736110297064d1894694d", "source_path": "examples/resources/xcsh_container_registry/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:container_registry:example:resource", "parent_id": "xcsh-docs:resources:container_registry:examples", "path": "documentation/resources/container_registry/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "container_registry", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2130301231230320-1301231032231233-0032311020313122-3013012322020212-0130212101221232-0012132331221100-0232022323323110-2033011011101232", "registry_path": "docs/guides/resources--container_registry--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/container_registry/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_container_registry.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["container_registryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_container_registry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/examples/)
- [xcsh_container_registry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/container_registry/)
