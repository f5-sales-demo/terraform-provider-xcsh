---
page_title: "Resource"
subcategory: "Container"
description: "Resource for xcsh_k8s_cluster_role."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1112, "body_sha256": "sha256:758660292c8e57b083e2316dbecc8c7946f323a79836655685ab4da4df5f73db", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:k8s_cluster_role:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f4470002db951c71b74b86c89c38ea07204065a7ca36bf4ba1c80737618e16ac", "source_path": "examples/resources/xcsh_k8s_cluster_role/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:k8s_cluster_role:example:resource", "parent_id": "xcsh-docs:resources:k8s_cluster_role:examples", "path": "documentation/resources/k8s_cluster_role/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0210001011031233-2213030330023033-0230032213111100-0020111321332013-1121320113302011-1231132013303212-1110223132313312-1310110320133101", "registry_path": "docs/guides/resources--k8s_cluster_role--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role/examples/resource/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Resource for xcsh_k8s_cluster_role.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_k8s_cluster_role](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_k8s_cluster_role/resource.tf`; digest `sha256:f4470002db951c71b74b86c89c38ea07204065a7ca36bf4ba1c80737618e16ac`.

```terraform
# K8SClusterRole Resource Example
# Manages k8s_cluster_role will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SClusterRole configuration
resource "xcsh_k8s_cluster_role" "example" {
  name      = "example-k8s-cluster-role"
  namespace = "system"
}
```
