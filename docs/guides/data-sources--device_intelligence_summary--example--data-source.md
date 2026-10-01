---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_device_intelligence_summary."
xcsh_docs: {"aliases": [], "body_bytes": 1165, "body_sha256": "sha256:780deb1801f205dc209d820d3688bf3d9300b648a88d3d981392afe20a67bec6", "canonical_id": "xcsh-docs:data-sources:device_intelligence_summary:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:device_intelligence_summary:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a3b5af58871245c4f3794bc2ed886a4421185806d4cb6b7c6eb5328aa94877da", "source_path": "examples/data-sources/xcsh_device_intelligence_summary/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:device_intelligence_summary:example:data-source", "parent_id": "xcsh-docs:data-sources:device_intelligence_summary:examples", "path": "docs/guides/data-sources--device_intelligence_summary--example--data-source.md", "provider_name": "device_intelligence_summary", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_summary/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_device_intelligence_summary.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_device_intelligence_summary](../data-sources/device_intelligence_summary.md)
- [Examples](data-sources--device_intelligence_summary--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_summary/data-source.tf`; digest `sha256:a3b5af58871245c4f3794bc2ed886a4421185806d4cb6b7c6eb5328aa94877da`.

```terraform
# DeviceIntelligenceSummary DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_summary" "example" {
  namespace = "example-value"
}

output "device_intelligence_summary_result" {
  value = data.xcsh_device_intelligence_summary.example
}
```

## Next pages

- [Examples](data-sources--device_intelligence_summary--examples.md)
- [xcsh_device_intelligence_summary](../data-sources/device_intelligence_summary.md)
