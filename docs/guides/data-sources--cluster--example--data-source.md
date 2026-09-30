---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 911, "body_sha256": "sha256:da97eee81c3a565aad6462736d72ad726fba92e876b1d655c3246e5e877f7b84", "canonical_id": "xcsh-docs:data-sources:cluster:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:657c300c6f1b904145d0c5ed8c458dbdc20cbe652caf3ef6b0a40bc6ebb5e4ee", "source_path": "examples/data-sources/xcsh_cluster/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cluster:example:data-source", "parent_id": "xcsh-docs:data-sources:cluster:examples", "path": "docs/guides/data-sources--cluster--example--data-source.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md)
- [Examples](data-sources--cluster--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cluster/data-source.tf`; digest `sha256:657c300c6f1b904145d0c5ed8c458dbdc20cbe652caf3ef6b0a40bc6ebb5e4ee`.

```terraform
# Cluster Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cluster by name
data "xcsh_cluster" "example" {
  name      = "example-cluster"
  namespace = "staging"
}

output "cluster_id" {
  value = data.xcsh_cluster.example.id
}
```

## Next pages

- [Examples](data-sources--cluster--examples.md)
- [xcsh_cluster](../data-sources/cluster.md)
