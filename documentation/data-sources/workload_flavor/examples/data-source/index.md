---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_workload_flavor."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1317, "body_sha256": "sha256:877eea152a98fe12abd1ba56833d68884d4ddb00e31154db3bd24b916261977d", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:workload_flavor:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6a1b8a9b6022de8c3cc75f425f01b978df5bf741969fc4bc0998b22eb6f2b173", "source_path": "examples/data-sources/xcsh_workload_flavor/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:workload_flavor:example:data-source", "parent_id": "xcsh-docs:data-sources:workload_flavor:examples", "path": "documentation/data-sources/workload_flavor/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "workload_flavor", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3203031020011131-3202033303100303-1322213220233302-3302132331023112-3113123022323233-1123022212331032-0111222222310331-0012310120313013", "registry_path": "docs/guides/data-sources--workload_flavor--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload_flavor/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_workload_flavor.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workload_flavorCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_workload_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload_flavor/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload_flavor/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_workload_flavor/data-source.tf`; digest `sha256:6a1b8a9b6022de8c3cc75f425f01b978df5bf741969fc4bc0998b22eb6f2b173`.

```terraform
# WorkloadFlavor Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing WorkloadFlavor by name
data "xcsh_workload_flavor" "example" {
  name      = "example-workload-flavor"
  namespace = "shared"
}

output "workload_flavor_id" {
  value = data.xcsh_workload_flavor.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload_flavor/examples/)
- [xcsh_workload_flavor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload_flavor/)
