---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cloud_region."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1279, "body_sha256": "sha256:f19a520b7cd634bf48ad9770d07d60962a60b76ffe74cb8c09ab08ce372c94f7", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_region:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:78c1e9c6e8676d021e9b37f2f6a06e7aa0e27e7d2ed71f73c7951ef98b928563", "source_path": "examples/data-sources/xcsh_cloud_region/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cloud_region:example:data-source", "parent_id": "xcsh-docs:data-sources:cloud_region:examples", "path": "documentation/data-sources/cloud_region/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "cloud_region", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3101303132200022-0123003311322231-3330303011110000-1220020103221231-3030222230132110-3013130112303132-2010020233230323-2003131010100310", "registry_path": "docs/guides/data-sources--cloud_region--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_region/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_cloud_region.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cloud_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_region/data-source.tf`; digest `sha256:78c1e9c6e8676d021e9b37f2f6a06e7aa0e27e7d2ed71f73c7951ef98b928563`.

```terraform
# CloudRegion Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudRegion by name
data "xcsh_cloud_region" "example" {
  name      = "example-cloud-region"
  namespace = "staging"
}

output "cloud_region_id" {
  value = data.xcsh_cloud_region.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/examples/)
- [xcsh_cloud_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_region/)
