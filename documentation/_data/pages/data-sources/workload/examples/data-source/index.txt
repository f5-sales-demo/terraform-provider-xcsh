---
page_title: "Data source"
subcategory: "Container"
description: "Data source for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 1229, "body_sha256": "sha256:2ced002b26d9b8d61a8a1b5e1b0ddb87897df03774237beead663a94c120d838", "child_ids": [], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:cc2cda57d20aa47819844a194fd93abc59f8f0880458fe8ebeae6e323411954a", "source_path": "examples/data-sources/xcsh_workload/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:workload:example:data-source", "parent_id": "xcsh-docs:data-sources:workload:examples", "path": "documentation/data-sources/workload/examples/data-source/index.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
