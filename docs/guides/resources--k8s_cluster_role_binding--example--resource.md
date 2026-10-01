---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_k8s_cluster_role_binding."
xcsh_docs: {"aliases": [], "body_bytes": 1235, "body_sha256": "sha256:f8416832eb77e1eb621eccfd14582adabf95434f51191094c377604e7725e125", "canonical_id": "xcsh-docs:resources:k8s_cluster_role_binding:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:k8s_cluster_role_binding:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5eb9f86d55919610546b8f56916f90f46bab49e1d41c08733c6886ac7ecd18db", "source_path": "examples/resources/xcsh_k8s_cluster_role_binding/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:k8s_cluster_role_binding:example:resource", "parent_id": "xcsh-docs:resources:k8s_cluster_role_binding:examples", "path": "docs/guides/resources--k8s_cluster_role_binding--example--resource.md", "provider_name": "k8s_cluster_role_binding", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role_binding/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_k8s_cluster_role_binding.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_role_bindingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md)
- [Examples](resources--k8s_cluster_role_binding--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_k8s_cluster_role_binding/resource.tf`; digest `sha256:5eb9f86d55919610546b8f56916f90f46bab49e1d41c08733c6886ac7ecd18db`.

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

## Next pages

- [Examples](resources--k8s_cluster_role_binding--examples.md)
- [xcsh_k8s_cluster_role_binding](../resources/k8s_cluster_role_binding.md)
