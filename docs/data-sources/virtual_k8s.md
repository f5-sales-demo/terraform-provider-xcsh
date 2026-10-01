---
page_title: "xcsh_virtual_k8s"
subcategory: "Container"
description: "xcsh_virtual_k8s for xcsh_virtual_k8s."
xcsh_docs: {"aliases": [], "body_bytes": 1397, "body_sha256": "sha256:906da7aa620430b42bd4aa59c57b1e1264f8a39bf5a6e56a4927c53d44c032f3", "canonical_id": "xcsh-docs:data-sources:virtual_k8s:fundamentals", "child_ids": ["xcsh-docs:data-sources:virtual_k8s:reference", "xcsh-docs:data-sources:virtual_k8s:examples"], "collection_id": "xcsh-docs:data-sources:virtual_k8s:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_k8s:fundamentals", "parent_id": null, "path": "docs/data-sources/virtual_k8s.md", "provider_name": "virtual_k8s", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_k8s/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_virtual_k8s for xcsh_virtual_k8s.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_k8sCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_virtual_k8s

Breadcrumbs:

- xcsh_virtual_k8s

Manages virtual\_k8s will create the object in the storage backend for namespace metadata.namespace
in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `workload`.

- workload: Container workloads in this namespace

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualK8S Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualK8S by name
data "xcsh_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "staging"
}

output "virtual_k8s_id" {
  value = data.xcsh_virtual_k8s.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--virtual_k8s--reference.md)
- [Examples](../guides/data-sources--virtual_k8s--examples.md)
