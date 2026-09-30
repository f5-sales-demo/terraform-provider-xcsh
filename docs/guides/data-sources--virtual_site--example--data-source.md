---
page_title: "Data source"
subcategory: "Infrastructure"
description: "Data source for xcsh_virtual_site."
xcsh_docs: {"aliases": [], "body_bytes": 974, "body_sha256": "sha256:bab68c775f381fc9b0559ff2fdfbdfbbae0d68dc8658bd6d7bb8a18cad4d77cd", "canonical_id": "xcsh-docs:data-sources:virtual_site:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:virtual_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:319f1c9eed60a6b55a583796c0636dc04effd666a48047cf3b9134b1863d1477", "source_path": "examples/data-sources/xcsh_virtual_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:virtual_site:example:data-source", "parent_id": "xcsh-docs:data-sources:virtual_site:examples", "path": "docs/guides/data-sources--virtual_site--example--data-source.md", "provider_name": "virtual_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_virtual_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_virtual_site](../data-sources/virtual_site.md)
- [Examples](data-sources--virtual_site--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_site/data-source.tf`; digest `sha256:319f1c9eed60a6b55a583796c0636dc04effd666a48047cf3b9134b1863d1477`.

```terraform
# VirtualSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualSite by name
data "xcsh_virtual_site" "example" {
  name      = "example-virtual-site"
  namespace = "staging"
}

output "virtual_site_id" {
  value = data.xcsh_virtual_site.example.id
}
```

## Next pages

- [Examples](data-sources--virtual_site--examples.md)
- [xcsh_virtual_site](../data-sources/virtual_site.md)
