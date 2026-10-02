---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_filter_set."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1253, "body_sha256": "sha256:7913cda5c71a9f65ddbb251217efc5d8067ee8f834e48d89ad8cb125f01bc0ca", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:filter_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:03a6f567d57f7b6dda2ca579596922f0b9b282f98f909656ba616ef93c767b0d", "source_path": "examples/data-sources/xcsh_filter_set/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:filter_set:example:data-source", "parent_id": "xcsh-docs:data-sources:filter_set:examples", "path": "documentation/data-sources/filter_set/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-2120310123232131-1010111201313002-2222112022000303-1312202013021330-2121202223020110-3313022212330131-1112230210131311-2011010312313322", "registry_path": "docs/guides/data-sources--filter_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/filter_set/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_filter_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["filter_setCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_filter_set/data-source.tf`; digest `sha256:03a6f567d57f7b6dda2ca579596922f0b9b282f98f909656ba616ef93c767b0d`.

```terraform
# FilterSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FilterSet by name
data "xcsh_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "staging"
}

output "filter_set_id" {
  value = data.xcsh_filter_set.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/examples/)
- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/)
