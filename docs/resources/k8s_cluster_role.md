---
page_title: "xcsh_k8s_cluster_role"
subcategory: "Container"
description: "xcsh_k8s_cluster_role for xcsh_k8s_cluster_role."
xcsh_docs: {"aliases": [], "body_bytes": 1512, "body_sha256": "sha256:bd9f32179fa256fbb2ba983346ac6eff59ab81e923b7c3886a51e2f495e68f30", "canonical_id": "xcsh-docs:resources:k8s_cluster_role:fundamentals", "child_ids": ["xcsh-docs:resources:k8s_cluster_role:reference", "xcsh-docs:resources:k8s_cluster_role:examples", "xcsh-docs:resources:k8s_cluster_role:import", "xcsh-docs:resources:k8s_cluster_role:timeouts"], "collection_id": "xcsh-docs:resources:k8s_cluster_role:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role:fundamentals", "parent_id": null, "path": "docs/resources/k8s_cluster_role.md", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_k8s_cluster_role for xcsh_k8s_cluster_role.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](../guides/resources--k8s_cluster_role--reference.md)
- [Examples](../guides/resources--k8s_cluster_role--examples.md)
- [Import](../guides/resources--k8s_cluster_role--import.md)
- [Timeouts](../guides/resources--k8s_cluster_role--timeouts.md)
