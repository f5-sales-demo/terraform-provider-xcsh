---
page_title: "xcsh_cluster"
subcategory: ""
description: "Manages cluster will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["cluster"], "body_bytes": 1575, "body_sha256": "sha256:e468791ad036faebbc9114a56b57e31a0dfa256a0f499cba8b0fe0a7faf801ac", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:cluster:reference", "xcsh-docs:resources:cluster:examples", "xcsh-docs:resources:cluster:import", "xcsh-docs:resources:cluster:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/cluster/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003", "registry_path": "docs/resources/cluster.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages cluster will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_cluster

Breadcrumbs:

- xcsh_cluster

Manages cluster will create the object in the storage backend for namespace metadata.namespace in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Cluster Resource Example
# Manages cluster will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Cluster configuration
resource "xcsh_cluster" "example" {
  name      = "example-cluster"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/lifecycle/timeouts/)
