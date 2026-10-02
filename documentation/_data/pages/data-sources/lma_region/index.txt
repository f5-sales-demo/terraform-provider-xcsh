---
page_title: "xcsh_lma_region"
subcategory: ""
description: "Manages a Lma Region resource in F5 Distributed Cloud for lma region specification. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["lma region"], "body_bytes": 1349, "body_sha256": "sha256:989618b4ed5bd8e0facd56ea4be72638dfeb75f01d8d52f9703dd8f42c723d7b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:lma_region:reference", "xcsh-docs:data-sources:lma_region:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/lma_region/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2212011320333232-1123311200100001-1321311330220000-3033022231230311-3012011130132021-2120003020013220-1120200123320031-0302331121003101", "registry_path": "docs/data-sources/lma_region.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Lma Region resource in F5 Distributed Cloud for lma region specification. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_lma_region

Breadcrumbs:

- xcsh_lma_region

Manages a Lma Region resource in F5 Distributed Cloud for lma region specification. configuration.
(read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/examples/)
