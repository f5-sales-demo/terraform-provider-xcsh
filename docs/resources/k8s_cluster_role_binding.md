---
page_title: "xcsh_k8s_cluster_role_binding"
subcategory: ""
description: "xcsh_k8s_cluster_role_binding for xcsh_k8s_cluster_role_binding."
xcsh_docs: {"aliases": [], "body_bytes": 1587, "body_sha256": "sha256:4e5baa1a55f229e9ef154ca4ccd057e8842475ddbf70a2f1eeecff098d3e84d5", "canonical_id": "xcsh-docs:resources:k8s_cluster_role_binding:fundamentals", "child_ids": ["xcsh-docs:resources:k8s_cluster_role_binding:reference", "xcsh-docs:resources:k8s_cluster_role_binding:examples", "xcsh-docs:resources:k8s_cluster_role_binding:import", "xcsh-docs:resources:k8s_cluster_role_binding:timeouts"], "collection_id": "xcsh-docs:resources:k8s_cluster_role_binding:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster_role_binding:fundamentals", "parent_id": null, "path": "docs/resources/k8s_cluster_role_binding.md", "provider_name": "k8s_cluster_role_binding", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role_binding/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_k8s_cluster_role_binding for xcsh_k8s_cluster_role_binding.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_role_bindingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](../guides/resources--k8s_cluster_role_binding--reference.md)
- [Examples](../guides/resources--k8s_cluster_role_binding--examples.md)
- [Import](../guides/resources--k8s_cluster_role_binding--import.md)
- [Timeouts](../guides/resources--k8s_cluster_role_binding--timeouts.md)
