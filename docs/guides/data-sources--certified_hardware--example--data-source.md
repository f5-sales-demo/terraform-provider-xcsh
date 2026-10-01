---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_certified_hardware."
xcsh_docs: {"aliases": [], "body_bytes": 1151, "body_sha256": "sha256:3944d14e4079cc4744f823b87a9b38750a666357bbbe00a6b7826b5c7c5e0d5b", "canonical_id": "xcsh-docs:data-sources:certified_hardware:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:44db2061b6da9d984930695f1d32baecee3d5ce0c74f32e0c377962e7ba246ba", "source_path": "examples/data-sources/xcsh_certified_hardware/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:certified_hardware:example:data-source", "parent_id": "xcsh-docs:data-sources:certified_hardware:examples", "path": "docs/guides/data-sources--certified_hardware--example--data-source.md", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_certified_hardware.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md)
- [Examples](data-sources--certified_hardware--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_certified_hardware/data-source.tf`; digest `sha256:44db2061b6da9d984930695f1d32baecee3d5ce0c74f32e0c377962e7ba246ba`.

```terraform
# CertifiedHardware Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CertifiedHardware by name
data "xcsh_certified_hardware" "example" {
  name      = "example-certified-hardware"
  namespace = "staging"
}

output "certified_hardware_id" {
  value = data.xcsh_certified_hardware.example.id
}
```

## Next pages

- [Examples](data-sources--certified_hardware--examples.md)
- [xcsh_certified_hardware](../data-sources/certified_hardware.md)
