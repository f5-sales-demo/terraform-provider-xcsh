---
page_title: "xcsh_bigip_virtual_server"
subcategory: ""
description: "xcsh_bigip_virtual_server for xcsh_bigip_virtual_server."
xcsh_docs: {"aliases": [], "body_bytes": 1384, "body_sha256": "sha256:05c03fb9fc743f0f2a9f7d6cf5928a735deabb00af0456189d56e08bf396e4e6", "canonical_id": "xcsh-docs:data-sources:bigip_virtual_server:fundamentals", "child_ids": ["xcsh-docs:data-sources:bigip_virtual_server:reference", "xcsh-docs:data-sources:bigip_virtual_server:examples"], "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:fundamentals", "parent_id": null, "path": "docs/data-sources/bigip_virtual_server.md", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_bigip_virtual_server for xcsh_bigip_virtual_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bigip_virtual_server

Breadcrumbs:

- xcsh_bigip_virtual_server

Manages a BIG-IP Virtual Server resource in F5 Distributed Cloud for big-ip virtual server
specification. configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BigIPVirtualServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BigIPVirtualServer by name
data "xcsh_bigip_virtual_server" "example" {
  name      = "example-bigip-virtual-server"
  namespace = "staging"
}

output "bigip_virtual_server_id" {
  value = data.xcsh_bigip_virtual_server.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--bigip_virtual_server--reference.md)
- [Examples](../guides/data-sources--bigip_virtual_server--examples.md)
