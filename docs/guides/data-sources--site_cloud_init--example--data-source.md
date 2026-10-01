---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_cloud_init."
xcsh_docs: {"aliases": [], "body_bytes": 1092, "body_sha256": "sha256:f623ea58c144856186e8dd6f303a3cb85232a31845b25e9da0f7a959e90cc573", "canonical_id": "xcsh-docs:data-sources:site_cloud_init:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_cloud_init:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:edce0cae6fa8845014088889818debc08fd6df550bae0c288c669538441379ae", "source_path": "examples/data-sources/xcsh_site_cloud_init/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_cloud_init:example:data-source", "parent_id": "xcsh-docs:data-sources:site_cloud_init:examples", "path": "docs/guides/data-sources--site_cloud_init--example--data-source.md", "provider_name": "site_cloud_init", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_cloud_init/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_site_cloud_init.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_cloud_init](../data-sources/site_cloud_init.md)
- [Examples](data-sources--site_cloud_init--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_cloud_init/data-source.tf`; digest `sha256:edce0cae6fa8845014088889818debc08fd6df550bae0c288c669538441379ae`.

```terraform
# SiteCloudInit DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_cloud_init" "example" {
  provider_ref = "example-value"
  site_name    = "example-value"
}

output "site_cloud_init_result" {
  value     = data.xcsh_site_cloud_init.example
  sensitive = true
}
```

## Next pages

- [Examples](data-sources--site_cloud_init--examples.md)
- [xcsh_site_cloud_init](../data-sources/site_cloud_init.md)
