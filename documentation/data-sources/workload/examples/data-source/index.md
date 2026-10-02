---
page_title: "Data source"
subcategory: "Container"
description: "Data source for xcsh_workload."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1229, "body_sha256": "sha256:2ced002b26d9b8d61a8a1b5e1b0ddb87897df03774237beead663a94c120d838", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:cc2cda57d20aa47819844a194fd93abc59f8f0880458fe8ebeae6e323411954a", "source_path": "examples/data-sources/xcsh_workload/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:workload:example:data-source", "parent_id": "xcsh-docs:data-sources:workload:examples", "path": "documentation/data-sources/workload/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0021101122100012-0321032002123223-3311132110313023-3022101022212002-1133333033021220-3331101233200213-3111312331210312-2120022012122033", "registry_path": "docs/guides/data-sources--workload--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_workload/data-source.tf`; digest `sha256:cc2cda57d20aa47819844a194fd93abc59f8f0880458fe8ebeae6e323411954a`.

```terraform
# Workload Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Workload by name
data "xcsh_workload" "example" {
  name      = "example-workload"
  namespace = "staging"
}

output "workload_id" {
  value = data.xcsh_workload.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/examples/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
