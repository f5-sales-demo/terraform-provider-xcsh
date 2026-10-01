---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_dc_cluster_group."
xcsh_docs: {"aliases": [], "body_bytes": 1286, "body_sha256": "sha256:0e4c3037d1dd7b21815745a234af5ccb6ba7d84e631938b0ba01e410cd20c4cd", "child_ids": [], "collection_id": "xcsh-docs:resources:dc_cluster_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8a0876385980d3e39be1eb4ce88a5f57c0c3917ba68858a952fc5d4f99aaef53", "source_path": "examples/resources/xcsh_dc_cluster_group/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dc_cluster_group:example:resource", "parent_id": "xcsh-docs:resources:dc_cluster_group:examples", "path": "documentation/resources/dc_cluster_group/examples/resource/index.md", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dc_cluster_group/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_dc_cluster_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dc_cluster_group/resource.tf`; digest `sha256:8a0876385980d3e39be1eb4ce88a5f57c0c3917ba68858a952fc5d4f99aaef53`.

```terraform
# DcClusterGroup Resource Example
# Manages DC Cluster group in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DcClusterGroup configuration
resource "xcsh_dc_cluster_group" "example" {
  name      = "example-dc-cluster-group"
  namespace = "system"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/examples/)
- [xcsh_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dc_cluster_group/)
