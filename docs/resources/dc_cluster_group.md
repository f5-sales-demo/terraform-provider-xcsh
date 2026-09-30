---
page_title: "xcsh_dc_cluster_group"
subcategory: ""
description: "xcsh_dc_cluster_group for xcsh_dc_cluster_group."
xcsh_docs: {"aliases": [], "body_bytes": 1257, "body_sha256": "sha256:96585db7bef7a9ad7976c0ac8499fc7a7bae0a9a7332740b2d7b1b48ef1fec88", "canonical_id": "xcsh-docs:resources:dc_cluster_group:fundamentals", "child_ids": ["xcsh-docs:resources:dc_cluster_group:reference", "xcsh-docs:resources:dc_cluster_group:examples", "xcsh-docs:resources:dc_cluster_group:import", "xcsh-docs:resources:dc_cluster_group:timeouts"], "collection_id": "xcsh-docs:resources:dc_cluster_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:dc_cluster_group:fundamentals", "parent_id": null, "path": "docs/resources/dc_cluster_group.md", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dc_cluster_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_dc_cluster_group for xcsh_dc_cluster_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_dc_cluster_group

Breadcrumbs:

- xcsh_dc_cluster_group

Manages DC Cluster group in given namespace in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DcClusterGroup Resource Example
# Manages DC Cluster group in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DcClusterGroup configuration
resource "xcsh_dc_cluster_group" "example" {
  name      = "example-dc-cluster-group"
  namespace = "system"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--dc_cluster_group--reference.md)
- [Examples](../guides/resources--dc_cluster_group--examples.md)
- [Import](../guides/resources--dc_cluster_group--import.md)
- [Timeouts](../guides/resources--dc_cluster_group--timeouts.md)
