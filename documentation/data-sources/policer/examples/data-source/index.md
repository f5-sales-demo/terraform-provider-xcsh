---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_policer."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1003, "body_sha256": "sha256:b56198cab6b5ddc4d912d4775a9a140c384b16ba306cba07efddf4deee65c7f5", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:policer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:bf10d06a3b33dc40f3d87d76f2625b6d1d04a36df75db4e7eceed8ef9f19362a", "source_path": "examples/data-sources/xcsh_policer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:policer:example:data-source", "parent_id": "xcsh-docs:data-sources:policer:examples", "path": "documentation/data-sources/policer/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "policer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1002021032231023-1113203102320202-1101332212133120-2000300203301113-0313311230123111-0213020212200231-1113221002333333-0103232201133122", "registry_path": "docs/guides/data-sources--policer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policer/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_policer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["policerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policer/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_policer/data-source.tf`; digest `sha256:bf10d06a3b33dc40f3d87d76f2625b6d1d04a36df75db4e7eceed8ef9f19362a`.

```terraform
# Policer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Policer by name
data "xcsh_policer" "example" {
  name      = "example-policer"
  namespace = "staging"
}

output "policer_id" {
  value = data.xcsh_policer.example.id
}
```
