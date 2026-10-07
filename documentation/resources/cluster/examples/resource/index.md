---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cluster."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1036, "body_sha256": "sha256:0a149a88b604c85d68382b7851a9d7d0854431b33b74cb478cf1f37f634d91b0", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5826c5405f2445e4c35c5efebde0e90198d315b39c9333da9092af648210a59e", "source_path": "examples/resources/xcsh_cluster/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cluster:example:resource", "parent_id": "xcsh-docs:resources:cluster:examples", "path": "documentation/resources/cluster/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0331332022232123-1322210312322323-3011020222012033-3110310002112202-1232010201201122-0302122102013332-0301020330012110-0312002302322022", "registry_path": "docs/guides/resources--cluster--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/examples/resource/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Resource for xcsh_cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["clusterCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cluster/resource.tf`; digest `sha256:5826c5405f2445e4c35c5efebde0e90198d315b39c9333da9092af648210a59e`.

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
