---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 1037, "body_sha256": "sha256:658255c9eb7fca68c1b8c48377fcfaac20c3c7668214ce4f53de45fa725ac478", "canonical_id": "xcsh-docs:resources:cluster:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5826c5405f2445e4c35c5efebde0e90198d315b39c9333da9092af648210a59e", "source_path": "examples/resources/xcsh_cluster/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cluster:example:resource", "parent_id": "xcsh-docs:resources:cluster:examples", "path": "docs/guides/resources--cluster--example--resource.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md)
- [Examples](resources--cluster--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cluster/resource.tf`; digest `sha256:5826c5405f2445e4c35c5efebde0e90198d315b39c9333da9092af648210a59e`.

```terraform
# Cluster Resource Example
# Manages cluster will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Cluster configuration
resource "xcsh_cluster" "example" {
  name      = "example-cluster"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--cluster--examples.md)
- [xcsh_cluster](../resources/cluster.md)
