---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_k8s_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1059, "body_sha256": "sha256:1ed980b196b5f3c7e14452bbaf144130272c8e2b6a4f1dabf5fbc78677d3fcef", "canonical_id": "xcsh-docs:data-sources:k8s_cluster:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1ffe54fa3443473ef381429123673c75783a908cc4dae691369d96028c415887", "source_path": "examples/data-sources/xcsh_k8s_cluster/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:k8s_cluster:example:data-source", "parent_id": "xcsh-docs:data-sources:k8s_cluster:examples", "path": "docs/guides/data-sources--k8s_cluster--example--data-source.md", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_k8s_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md)
- [Examples](data-sources--k8s_cluster--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_k8s_cluster/data-source.tf`; digest `sha256:1ffe54fa3443473ef381429123673c75783a908cc4dae691369d96028c415887`.

```terraform
# K8SCluster Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SCluster by name
data "xcsh_k8s_cluster" "example" {
  name      = "example-k8s-cluster"
  namespace = "system"
}

output "k8s_cluster_id" {
  value = data.xcsh_k8s_cluster.example.id
}
```

## Next pages

- [Examples](data-sources--k8s_cluster--examples.md)
- [xcsh_k8s_cluster](../data-sources/k8s_cluster.md)
