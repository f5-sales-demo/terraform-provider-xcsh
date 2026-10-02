---
page_title: "Data source"
subcategory: "Infrastructure"
description: "Data source for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1315, "body_sha256": "sha256:e5410abf6113e26e670339fba1925b0c50ecd4c57b1ead5bd1e909344dba63eb", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c010494f98bc90e6b2715fc8293de32922c6d80e43d822b5b3e3b2a72ebc3f82", "source_path": "examples/data-sources/xcsh_azure_vnet_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:azure_vnet_site:example:data-source", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:examples", "path": "documentation/data-sources/azure_vnet_site/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1313003303031321-2133111010020232-1103320102023312-0200303221310322-2233231131001332-2320021202102211-3311332330131003-2301132303000332", "registry_path": "docs/guides/data-sources--azure_vnet_site--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_azure_vnet_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/examples/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
