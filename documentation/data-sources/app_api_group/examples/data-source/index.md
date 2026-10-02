---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_app_api_group."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1290, "body_sha256": "sha256:72e6b77d379053308d6de4644adf21a42f7dead371afeeb8f978e9fdd6ef5e25", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_api_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ea62726df10ead3d5992d45e3c012bb70708bc94e4271b50bd5cb8013ce78312", "source_path": "examples/data-sources/xcsh_app_api_group/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_api_group:example:data-source", "parent_id": "xcsh-docs:data-sources:app_api_group:examples", "path": "documentation/data-sources/app_api_group/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "app_api_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3313023333231021-0222020033001332-3103101120320231-3021031120132230-1111122322111110-3301023103232020-0230000020312220-0031200222103313", "registry_path": "docs/guides/data-sources--app_api_group--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_api_group/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_app_api_group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_api_group/data-source.tf`; digest `sha256:ea62726df10ead3d5992d45e3c012bb70708bc94e4271b50bd5cb8013ce78312`.

```terraform
# AppAPIGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppAPIGroup by name
data "xcsh_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"
}

output "app_api_group_id" {
  value = data.xcsh_app_api_group.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/examples/)
- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/)
