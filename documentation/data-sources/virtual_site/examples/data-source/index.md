---
page_title: "Data source"
subcategory: "Infrastructure"
description: "Data source for xcsh_virtual_site."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1279, "body_sha256": "sha256:80c554858a749ee6ecc3625b0ff402ec99220a601b2da5fce5893fae079a206e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:virtual_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:319f1c9eed60a6b55a583796c0636dc04effd666a48047cf3b9134b1863d1477", "source_path": "examples/data-sources/xcsh_virtual_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:virtual_site:example:data-source", "parent_id": "xcsh-docs:data-sources:virtual_site:examples", "path": "documentation/data-sources/virtual_site/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "virtual_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1000133010023010-0301230230021210-3312232022020211-2010102212332200-0100020110033011-2023211012013003-1223311121030321-2110231103331112", "registry_path": "docs/guides/data-sources--virtual_site--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_virtual_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_site/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_site/examples/)
- [xcsh_virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_site/)
