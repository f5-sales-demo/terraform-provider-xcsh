---
page_title: "xcsh_k8s_cluster_role"
subcategory: "Container"
description: "Manages k8s_cluster_role will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["backend servers", "k8s cluster role", "origin servers", "upstream servers"], "body_bytes": 1701, "body_sha256": "sha256:eb02bb0e52d09a05678a12258ee1df0da6083c737df61487118f93f7a089f1b4", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_cluster_role:reference", "xcsh-docs:resources:k8s_cluster_role:examples", "xcsh-docs:resources:k8s_cluster_role:import", "xcsh-docs:resources:k8s_cluster_role:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/k8s_cluster_role/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1032203102313310-1310133031312032-2032211220123023-3022101331120133-3300202302000301-2221012031203122-0233002122331003-2331111132121223", "registry_path": "docs/resources/k8s_cluster_role.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages k8s_cluster_role will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_k8s_cluster_role

Breadcrumbs:

- xcsh_k8s_cluster_role

Manages k8s\_cluster\_role will create the object in the storage backend for namespace
metadata.namespace in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role/lifecycle/timeouts/)
