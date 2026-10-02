---
page_title: "xcsh_srv6_network_slice"
subcategory: ""
description: "Manages srv6_network_slice creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "srv6 network slice", "upstream servers"], "body_bytes": 1411, "body_sha256": "sha256:918a46c0cf6107068479a1f54233f45165cbf53c3991d7339f90c8ed4ec0e952", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:srv6_network_slice:reference", "xcsh-docs:data-sources:srv6_network_slice:examples"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:srv6_network_slice:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:srv6_network_slice:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/srv6_network_slice/index.md", "product": "distributed-cloud", "provider_name": "srv6_network_slice", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0001322321013132-3303233021210322-1322230110201110-3330310302210111-3123030332122130-1230012021313000-1121212210223211-3123021213301223", "registry_path": "docs/data-sources/srv6_network_slice.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/srv6_network_slice/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages srv6_network_slice creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["srv6_network_sliceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_srv6_network_slice

Breadcrumbs:

- xcsh_srv6_network_slice

Manages srv6\_network\_slice creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Srv6NetworkSlice Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Srv6NetworkSlice by name
data "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"
}

output "srv6_network_slice_id" {
  value = data.xcsh_srv6_network_slice.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/examples/)
