---
page_title: "Data source"
subcategory: "Container"
description: "Data source for xcsh_k8s_cluster_role."
xcsh_docs: {"aliases": [], "body_bytes": 1122, "body_sha256": "sha256:a3147e5249c909975b8a2fda2270b40f7ff4f6544232eb789dab638fc7bf0fe2", "canonical_id": "xcsh-docs:data-sources:k8s_cluster_role:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:k8s_cluster_role:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3a2a8196bb01ed845c9acfb23f02e2761ef76a0ab6bfc6ad04f8767ce9170112", "source_path": "examples/data-sources/xcsh_k8s_cluster_role/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:k8s_cluster_role:example:data-source", "parent_id": "xcsh-docs:data-sources:k8s_cluster_role:examples", "path": "docs/guides/data-sources--k8s_cluster_role--example--data-source.md", "provider_name": "k8s_cluster_role", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster_role/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_k8s_cluster_role.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["k8s_cluster_roleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md)
- [Examples](data-sources--k8s_cluster_role--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_k8s_cluster_role/data-source.tf`; digest `sha256:3a2a8196bb01ed845c9acfb23f02e2761ef76a0ab6bfc6ad04f8767ce9170112`.

```terraform
# K8SClusterRole Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SClusterRole by name
data "xcsh_k8s_cluster_role" "example" {
  name      = "example-k8s-cluster-role"
  namespace = "system"
}

output "k8s_cluster_role_id" {
  value = data.xcsh_k8s_cluster_role.example.id
}
```

## Next pages

- [Examples](data-sources--k8s_cluster_role--examples.md)
- [xcsh_k8s_cluster_role](../data-sources/k8s_cluster_role.md)
