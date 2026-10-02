---
page_title: "xcsh_dc_cluster_group"
subcategory: ""
description: "Manages DC Cluster group in given namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["dc cluster group"], "body_bytes": 1339, "body_sha256": "sha256:b46882aae390de93417326ae80dd9b2113d2620ea0c40cd8999febbe5232a1c4", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:dc_cluster_group:reference", "xcsh-docs:data-sources:dc_cluster_group:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dc_cluster_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dc_cluster_group:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/dc_cluster_group/index.md", "product": "distributed-cloud", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2200021000201211-3130032310312033-3013132131200022-0333300021021222-1220133033212100-0102010030313222-0332101231122310-0203313010012313", "registry_path": "docs/data-sources/dc_cluster_group.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dc_cluster_group/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages DC Cluster group in given namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
