---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_workload_flavor."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1317, "body_sha256": "sha256:877eea152a98fe12abd1ba56833d68884d4ddb00e31154db3bd24b916261977d", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:workload_flavor:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6a1b8a9b6022de8c3cc75f425f01b978df5bf741969fc4bc0998b22eb6f2b173", "source_path": "examples/data-sources/xcsh_workload_flavor/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:workload_flavor:example:data-source", "parent_id": "xcsh-docs:data-sources:workload_flavor:examples", "path": "documentation/data-sources/workload_flavor/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "workload_flavor", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3203031020011131-3202033303100303-1322213220233302-3302132331023112-3113123022323233-1123022212331032-0111222222310331-0012310120313013", "registry_path": "docs/guides/data-sources--workload_flavor--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload_flavor/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_workload_flavor.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workload_flavorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
