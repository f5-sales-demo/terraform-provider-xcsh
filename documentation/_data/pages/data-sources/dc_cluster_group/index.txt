---
page_title: "xcsh_dc_cluster_group"
subcategory: ""
description: "Manages DC Cluster group in given namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["dc cluster group"], "body_bytes": 1339, "body_sha256": "sha256:b46882aae390de93417326ae80dd9b2113d2620ea0c40cd8999febbe5232a1c4", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:dc_cluster_group:reference", "xcsh-docs:data-sources:dc_cluster_group:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dc_cluster_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dc_cluster_group:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/dc_cluster_group/index.md", "product": "distributed-cloud", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2200021000201211-3130032310312033-3013132131200022-0333300021021222-1220133033212100-0102010030313222-0332101231122310-0203313010012313", "registry_path": "docs/data-sources/dc_cluster_group.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dc_cluster_group/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages DC Cluster group in given namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
# DcClusterGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DcClusterGroup by name
data "xcsh_dc_cluster_group" "example" {
  name      = "example-dc-cluster-group"
  namespace = "system"
}

output "dc_cluster_group_id" {
  value = data.xcsh_dc_cluster_group.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dc_cluster_group/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dc_cluster_group/examples/)
