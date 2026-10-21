---
page_title: "xcsh_cluster"
subcategory: ""
description: "Reads Cluster information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["cluster"], "body_bytes": 1264, "body_sha256": "sha256:d8aec630bf07febaeff832241ca8ed980e69f48a99bab95910dcf1bd7ebbab54", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:cluster:reference", "xcsh-docs:data-sources:cluster:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/cluster/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0232210231030021-1231211333330303-2032130210102220-2102012032113202-2202232210201313-2210112130121033-3133211302133121-1202210031003013", "registry_path": "docs/data-sources/cluster.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Reads Cluster information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["clusterCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_cluster

Breadcrumbs:

- xcsh_cluster

Reads Cluster information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Cluster Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cluster by name
data "xcsh_cluster" "example" {
  name      = "example-cluster"
  namespace = "staging"
}

output "cluster_id" {
  value = data.xcsh_cluster.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/examples/)
