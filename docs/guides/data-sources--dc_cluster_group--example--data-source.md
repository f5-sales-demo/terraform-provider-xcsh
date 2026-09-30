---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_dc_cluster_group."
xcsh_docs: {"aliases": [], "body_bytes": 1023, "body_sha256": "sha256:9fcff824cf07f3f23a82118be83b5228eb09a29d21eaf9c183e2a20c8cf89a8e", "canonical_id": "xcsh-docs:data-sources:dc_cluster_group:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:dc_cluster_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:82b040ec9e5742587a1d4d2feb953ac78e9dca3a538c42aff66a3020a25fe4c5", "source_path": "examples/data-sources/xcsh_dc_cluster_group/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dc_cluster_group:example:data-source", "parent_id": "xcsh-docs:data-sources:dc_cluster_group:examples", "path": "docs/guides/data-sources--dc_cluster_group--example--data-source.md", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dc_cluster_group/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_dc_cluster_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md)
- [Examples](data-sources--dc_cluster_group--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dc_cluster_group/data-source.tf`; digest `sha256:82b040ec9e5742587a1d4d2feb953ac78e9dca3a538c42aff66a3020a25fe4c5`.

```terraform
# DcClusterGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DcClusterGroup by name
data "xcsh_dc_cluster_group" "example" {
  name      = "example-dc-cluster-group"
  namespace = "system"
}

output "dc_cluster_group_id" {
  value = data.xcsh_dc_cluster_group.example.id
}
```

## Next pages

- [Examples](data-sources--dc_cluster_group--examples.md)
- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md)
