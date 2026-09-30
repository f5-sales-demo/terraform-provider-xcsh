---
page_title: "xcsh_tunnel"
subcategory: ""
description: "xcsh_tunnel for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1193, "body_sha256": "sha256:5171e13fa6899ea0b93faabfb14077e36a0ef4772b656799e12fe28662ad952b", "child_ids": ["xcsh-docs:data-sources:tunnel:reference", "xcsh-docs:data-sources:tunnel:examples"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:fundamentals", "parent_id": null, "path": "documentation/data-sources/tunnel/index.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_tunnel for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_tunnel

Breadcrumbs:

- xcsh_tunnel

Manages tunnel in a given namespace. If one already exist it will give a error in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Tunnel Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Tunnel by name
data "xcsh_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "staging"
}

output "tunnel_id" {
  value = data.xcsh_tunnel.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/examples/)
