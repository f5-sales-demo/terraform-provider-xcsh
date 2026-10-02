---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_policer."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1216, "body_sha256": "sha256:68b12a5a0314ecc429530bd124d5282d6f43db6d3dd20470711ad5360cec2256", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:policer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:bf10d06a3b33dc40f3d87d76f2625b6d1d04a36df75db4e7eceed8ef9f19362a", "source_path": "examples/data-sources/xcsh_policer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:policer:example:data-source", "parent_id": "xcsh-docs:data-sources:policer:examples", "path": "documentation/data-sources/policer/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "policer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1002021032231023-1113203102320202-1101332212133120-2000300203301113-0313311230123111-0213020212200231-1113221002333333-0103232201133122", "registry_path": "docs/guides/data-sources--policer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policer/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_policer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policer/examples/)
- [xcsh_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policer/)
