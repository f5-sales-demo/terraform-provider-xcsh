---
page_title: "Resource"
subcategory: "Container"
description: "Resource for xcsh_virtual_k8s."
xcsh_docs: {"aliases": [], "body_bytes": 1190, "body_sha256": "sha256:8ed57a4f8140dcc8cf154858da73968ddf7531f83e167bb8d172781e99f6ae1d", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_k8s:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a5065caf12c5cd73a6083207375f30b6129bce86fd82163c8845c9e0ab44c400", "source_path": "examples/resources/xcsh_virtual_k8s/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:virtual_k8s:example:resource", "parent_id": "xcsh-docs:resources:virtual_k8s:examples", "path": "documentation/resources/virtual_k8s/examples/resource/index.md", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_k8s/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_virtual_k8s.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_virtual_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_k8s/resource.tf`; digest `sha256:a5065caf12c5cd73a6083207375f30b6129bce86fd82163c8845c9e0ab44c400`.

```terraform
# VirtualK8S Resource Example
# Manages virtual_k8s will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualK8S configuration
resource "xcsh_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/examples/)
- [xcsh_virtual_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_k8s/)
