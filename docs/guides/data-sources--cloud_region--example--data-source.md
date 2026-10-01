---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cloud_region."
xcsh_docs: {"aliases": [], "body_bytes": 1073, "body_sha256": "sha256:7dd5e418396ed4079e051254944952a21752d19c45207ad6e678dd0148f8e330", "canonical_id": "xcsh-docs:data-sources:cloud_region:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cloud_region:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:78c1e9c6e8676d021e9b37f2f6a06e7aa0e27e7d2ed71f73c7951ef98b928563", "source_path": "examples/data-sources/xcsh_cloud_region/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cloud_region:example:data-source", "parent_id": "xcsh-docs:data-sources:cloud_region:examples", "path": "docs/guides/data-sources--cloud_region--example--data-source.md", "provider_name": "cloud_region", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_region/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_cloud_region.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cloud_region](../data-sources/cloud_region.md)
- [Examples](data-sources--cloud_region--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_region/data-source.tf`; digest `sha256:78c1e9c6e8676d021e9b37f2f6a06e7aa0e27e7d2ed71f73c7951ef98b928563`.

```terraform
# CloudRegion Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudRegion by name
data "xcsh_cloud_region" "example" {
  name      = "example-cloud-region"
  namespace = "staging"
}

output "cloud_region_id" {
  value = data.xcsh_cloud_region.example.id
}
```

## Next pages

- [Examples](data-sources--cloud_region--examples.md)
- [xcsh_cloud_region](../data-sources/cloud_region.md)
