---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_lma_region."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1253, "body_sha256": "sha256:169adde0909745c4f389dcaccad01a9c3e5cf507a434549e8444defb19068e6c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6e44cbfbf0065cf8526c29152a263e7eb100be46cd3be2092eca1d815333bf89", "source_path": "examples/data-sources/xcsh_lma_region/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:lma_region:example:data-source", "parent_id": "xcsh-docs:data-sources:lma_region:examples", "path": "documentation/data-sources/lma_region/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-1231011002111003-3313012131132321-3230212301000031-2201000110013133-2223103203300302-2002132103232003-1020001030201221-0113310210300231", "registry_path": "docs/guides/data-sources--lma_region--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_lma_region.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_lma_region/data-source.tf`; digest `sha256:6e44cbfbf0065cf8526c29152a263e7eb100be46cd3be2092eca1d815333bf89`.

```terraform
# LmaRegion Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing LmaRegion by name
data "xcsh_lma_region" "example" {
  name      = "example-lma-region"
  namespace = "staging"
}

output "lma_region_id" {
  value = data.xcsh_lma_region.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/examples/)
- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
