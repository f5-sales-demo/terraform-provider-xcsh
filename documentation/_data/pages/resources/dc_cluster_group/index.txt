---
page_title: "xcsh_dc_cluster_group"
subcategory: ""
description: "Manages DC Cluster group in given namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["dc cluster group"], "body_bytes": 1558, "body_sha256": "sha256:36a153869560a887c89acb5335535daa894b2e62e9366fc7a87e91e6c0fba37d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:dc_cluster_group:reference", "xcsh-docs:resources:dc_cluster_group:examples", "xcsh-docs:resources:dc_cluster_group:import", "xcsh-docs:resources:dc_cluster_group:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dc_cluster_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:dc_cluster_group:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/dc_cluster_group/index.md", "product": "distributed-cloud", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0023121103022030-2302020033112202-0212333320231232-3001211200201011-0122330223232010-0110110123001022-0123201000313123-3333311302021321", "registry_path": "docs/resources/dc_cluster_group.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dc_cluster_group/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Manages DC Cluster group in given namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/lifecycle/timeouts/)
