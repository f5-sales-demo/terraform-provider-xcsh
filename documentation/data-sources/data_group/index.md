---
page_title: "xcsh_data_group"
subcategory: ""
description: "Manages data group in a given namespace. If one already exists it will give an error in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["data group"], "body_bytes": 1336, "body_sha256": "sha256:fedb7f16357d5db37b70606b6cbdc393bc7ea0624a7d72219b954d68342a56d4", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:data_group:reference", "xcsh-docs:data-sources:data_group:examples"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_group:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/data_group/index.md", "product": "distributed-cloud", "provider_name": "data_group", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0222002101202123-1221202102003112-2210120001000231-0023132202013201-3203321232210332-0212210130133110-2010020210021103-2203321203313002", "registry_path": "docs/data-sources/data_group.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_group/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages data group in a given namespace. If one already exists it will give an error in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["data_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_data_group

Breadcrumbs:

- xcsh_data_group

Manages data group in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DataGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DataGroup by name
data "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}

output "data_group_id" {
  value = data.xcsh_data_group.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/examples/)
