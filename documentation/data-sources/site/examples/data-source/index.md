---
page_title: "Data source"
subcategory: "Infrastructure"
description: "Data source for xcsh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1177, "body_sha256": "sha256:7d359a6b26f4f875c68a2a9ab3646953bcb6d28d0d77a528de4f64513283f0b6", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3e384d82af36bf37fe3ff1c335c8824183b28374dc3c5fb9a56016a0094aa348", "source_path": "examples/data-sources/xcsh_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site:example:data-source", "parent_id": "xcsh-docs:data-sources:site:examples", "path": "documentation/data-sources/site/examples/data-source/index.md", "provider_name": "site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site/data-source.tf`; digest `sha256:3e384d82af36bf37fe3ff1c335c8824183b28374dc3c5fb9a56016a0094aa348`.

```terraform
# Site Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Site by name
data "xcsh_site" "example" {
  name      = "example-site"
  namespace = "staging"
}

output "site_id" {
  value = data.xcsh_site.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/examples/)
- [xcsh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site/)
