---
page_title: "Data source"
subcategory: "Infrastructure"
description: "Data source for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1010, "body_sha256": "sha256:eaddc529f9840f70df51fda3a6267022111ae0773e718ce4bd1959b274bddada", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c010494f98bc90e6b2715fc8293de32922c6d80e43d822b5b3e3b2a72ebc3f82", "source_path": "examples/data-sources/xcsh_azure_vnet_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:azure_vnet_site:example:data-source", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:examples", "path": "docs/guides/data-sources--azure_vnet_site--example--data-source.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Examples](data-sources--azure_vnet_site--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_azure_vnet_site/data-source.tf`; digest `sha256:c010494f98bc90e6b2715fc8293de32922c6d80e43d822b5b3e3b2a72ebc3f82`.

```terraform
# AzureVNETSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AzureVNETSite by name
data "xcsh_azure_vnet_site" "example" {
  name      = "example-azure-vnet-site"
  namespace = "system"
}

output "azure_vnet_site_id" {
  value = data.xcsh_azure_vnet_site.example.id
}
```

## Next pages

- [Examples](data-sources--azure_vnet_site--examples.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
