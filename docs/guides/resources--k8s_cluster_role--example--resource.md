---
page_title: "Resource"
subcategory: "Container"
description: "Resource for xcsh_k8s_cluster_role."
xcsh_docs: {"aliases": [], "body_bytes": 1041, "body_sha256": "sha256:dce7619c9e808f27d4a7dc260342eda1a90c139156b5727f86575c3c6421e64a", "canonical_id": "xcsh-docs:resources:k8s_cluster_role:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:k8s_cluster_role:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f4470002db951c71b74b86c89c38ea07204065a7ca36bf4ba1c80737618e16ac", "source_path": "examples/resources/xcsh_k8s_cluster_role/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:k8s_cluster_role:example:resource", "parent_id": "xcsh-docs:resources:k8s_cluster_role:examples", "path": "docs/guides/resources--k8s_cluster_role--example--resource.md", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster_role/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_k8s_cluster_role.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md)
- [Examples](resources--k8s_cluster_role--examples.md)
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

## Next pages

- [Examples](resources--k8s_cluster_role--examples.md)
- [xcsh_k8s_cluster_role](../resources/k8s_cluster_role.md)
