---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1082, "body_sha256": "sha256:86fd0a21f8bb132ad3eb29fea7f277229415c0036d3793341e574457843cd1a7", "canonical_id": "xcsh-docs:resources:k8s_cluster:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4b234e2e10f8643667634d9286b365cede58a97ccb384ebff96249fbe01b28a1", "source_path": "examples/resources/xcsh_k8s_cluster/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:k8s_cluster:example:resource", "parent_id": "xcsh-docs:resources:k8s_cluster:examples", "path": "docs/guides/resources--k8s_cluster--example--resource.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
- [Examples](resources--k8s_cluster--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_k8s_cluster/resource.tf`; digest `sha256:4b234e2e10f8643667634d9286b365cede58a97ccb384ebff96249fbe01b28a1`.

```terraform
# K8SCluster Resource Example
# Manages k8s_cluster will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SCluster configuration
resource "xcsh_k8s_cluster" "example" {
  name      = "example-k8s-cluster"
  namespace = "system"
}
```

## Next pages

- [Examples](resources--k8s_cluster--examples.md)
- [xcsh_k8s_cluster](../resources/k8s_cluster.md)
