---
page_title: "xcsh_srv6_network_slice"
subcategory: ""
description: "xcsh_srv6_network_slice for xcsh_srv6_network_slice."
xcsh_docs: {"aliases": [], "body_bytes": 1312, "body_sha256": "sha256:4305a0b15a55e9f5ba03ea3c9000c40745f84b3fba2a6f87c452c6a2992f37c2", "child_ids": ["xcsh-docs:data-sources:srv6_network_slice:reference", "xcsh-docs:data-sources:srv6_network_slice:examples"], "collection_id": "xcsh-docs:data-sources:srv6_network_slice:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:srv6_network_slice:fundamentals", "parent_id": null, "path": "documentation/data-sources/srv6_network_slice/index.md", "provider_name": "srv6_network_slice", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/srv6_network_slice/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_srv6_network_slice for xcsh_srv6_network_slice.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["srv6_network_sliceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
