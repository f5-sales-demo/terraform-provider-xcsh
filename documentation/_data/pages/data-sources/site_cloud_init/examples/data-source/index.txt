---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_cloud_init."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1298, "body_sha256": "sha256:b118161683794faef5b05a65f90de2fa7dae88411a583344b7d8489c75b5dc9e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_cloud_init:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:edce0cae6fa8845014088889818debc08fd6df550bae0c288c669538441379ae", "source_path": "examples/data-sources/xcsh_site_cloud_init/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_cloud_init:example:data-source", "parent_id": "xcsh-docs:data-sources:site_cloud_init:examples", "path": "documentation/data-sources/site_cloud_init/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "site_cloud_init", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3023031122223133-3022120121113110-2321333301133030-2300003101220121-2011303302310123-1320021012330331-3211001111311321-1210212212231330", "registry_path": "docs/guides/data-sources--site_cloud_init--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_cloud_init/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_site_cloud_init.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_cloud_init](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/examples/)
- [xcsh_site_cloud_init](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/)
