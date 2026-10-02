---
page_title: "xcsh_k8s_cluster_role_binding"
subcategory: ""
description: "Manages k8s_cluster_role_binding will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["backend servers", "k8s cluster role binding", "origin servers", "upstream servers"], "body_bytes": 1776, "body_sha256": "sha256:754a0f8b4ad1bae02ed50f2ba03e79b99feb4f63b79984c7b717abe7bdc95e2d", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_cluster_role_binding:reference", "xcsh-docs:resources:k8s_cluster_role_binding:examples", "xcsh-docs:resources:k8s_cluster_role_binding:import", "xcsh-docs:resources:k8s_cluster_role_binding:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster_role_binding:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role_binding:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/k8s_cluster_role_binding/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster_role_binding", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1203120300333222-3203223200312120-0002210222112313-1113120331003323-3102000331030303-1133200312133122-1031101033120033-3233030130121012", "registry_path": "docs/resources/k8s_cluster_role_binding.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role_binding/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages k8s_cluster_role_binding will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_role_bindingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_k8s_cluster_role_binding

Breadcrumbs:

- xcsh_k8s_cluster_role_binding

Manages k8s\_cluster\_role\_binding will create the object in the storage backend for namespace
metadata.namespace in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SClusterRoleBinding Resource Example
# Manages k8s_cluster_role_binding will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SClusterRoleBinding configuration
resource "xcsh_k8s_cluster_role_binding" "example" {
  name      = "example-k8s-cluster-role-binding"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster_role_binding/lifecycle/timeouts/)
