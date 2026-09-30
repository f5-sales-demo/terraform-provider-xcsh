---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_regional_edges."
xcsh_docs: {"aliases": [], "body_bytes": 1229, "body_sha256": "sha256:5b4c24bba84290590dcb666a78ecc066389639b2f78f176c333e80dc47111c28", "canonical_id": "xcsh-docs:data-sources:network_regional_edges:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_regional_edges:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1744edb321716a35476a9a25cade01cfd9671a46b90a9dfb2f90ccff555099d6", "source_path": "examples/data-sources/xcsh_network_regional_edges/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_regional_edges:example:data-source", "parent_id": "xcsh-docs:data-sources:network_regional_edges:examples", "path": "docs/guides/data-sources--network_regional_edges--example--data-source.md", "provider_name": "network_regional_edges", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_regional_edges/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_regional_edges.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_network_regional_edges](../data-sources/network_regional_edges.md)
- [Examples](data-sources--network_regional_edges--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_regional_edges/data-source.tf`; digest `sha256:1744edb321716a35476a9a25cade01cfd9671a46b90a9dfb2f90ccff555099d6`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

# Select the Regional Edge source networks that may initiate HTTPS connections
# to an origin. Omit regions to return all published regions.
data "xcsh_network_regional_edges" "origin_ingress" {
  regions = ["americas", "europe"]
}

output "https_origin_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_regional_edges.origin_ingress.cidr_blocks
  }
}
```

## Next pages

- [Examples](data-sources--network_regional_edges--examples.md)
- [xcsh_network_regional_edges](../data-sources/network_regional_edges.md)
