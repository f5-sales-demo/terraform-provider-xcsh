---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_registrations."
xcsh_docs: {"aliases": [], "body_bytes": 1174, "body_sha256": "sha256:1411ebd536a0d53da2ff54c90b0abf6d42bdac1fed4d6f0e79ed756ebe45ee65", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:382726257bf8ed644e4ba85a525a0636702e8ea87e1b309384adf95a431060c9", "source_path": "examples/data-sources/xcsh_site_registrations/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_registrations:example:data-source", "parent_id": "xcsh-docs:data-sources:site_registrations:examples", "path": "documentation/data-sources/site_registrations/examples/data-source/index.md", "provider_name": "site_registrations", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_site_registrations.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_registrations/data-source.tf`; digest `sha256:382726257bf8ed644e4ba85a525a0636702e8ea87e1b309384adf95a431060c9`.

```terraform
# SiteRegistrations DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations" "example" {
  namespace = "example-value"
}

output "site_registrations_result" {
  value = data.xcsh_site_registrations.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/examples/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
