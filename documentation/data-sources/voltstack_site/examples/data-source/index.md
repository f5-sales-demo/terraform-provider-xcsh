---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1206, "body_sha256": "sha256:5e2bef98c0ff81a496a56319f4518e4eaf8a15489d58c9d8947240e59ceaa75c", "child_ids": [], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9d6c1f9af1750a76c0f8df78e6df342564c3af32fa86a487ff12441eece04681", "source_path": "examples/data-sources/xcsh_voltstack_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:voltstack_site:example:data-source", "parent_id": "xcsh-docs:data-sources:voltstack_site:examples", "path": "documentation/data-sources/voltstack_site/examples/data-source/index.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_voltstack_site/data-source.tf`; digest `sha256:9d6c1f9af1750a76c0f8df78e6df342564c3af32fa86a487ff12441eece04681`.

```terraform
# VoltstackSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VoltstackSite by name
data "xcsh_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "staging"
}

output "voltstack_site_id" {
  value = data.xcsh_voltstack_site.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/examples/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
