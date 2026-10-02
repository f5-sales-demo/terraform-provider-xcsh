---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_user_identification."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1370, "body_sha256": "sha256:45a9b3d6de88a37e9caf21343637bb1dec935b661de0e96aeb98f99d562504e7", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:user_identification:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c30020dfd6aecd49ffb0273704b9b9e1c85f494c0e9e28cd883c479b6b6234a3", "source_path": "examples/data-sources/xcsh_user_identification/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:user_identification:example:data-source", "parent_id": "xcsh-docs:data-sources:user_identification:examples", "path": "documentation/data-sources/user_identification/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1113012020110212-3203110103303312-2010033223110313-3013231211121132-3210300300233322-3300021101211023-1120313212023220-0333121030121331", "registry_path": "docs/guides/data-sources--user_identification--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/user_identification/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_user_identification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["user_identificationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_user_identification/data-source.tf`; digest `sha256:c30020dfd6aecd49ffb0273704b9b9e1c85f494c0e9e28cd883c479b6b6234a3`.

```terraform
# UserIdentification Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UserIdentification by name
data "xcsh_user_identification" "example" {
  name      = "example-user-identification"
  namespace = "staging"
}

output "user_identification_id" {
  value = data.xcsh_user_identification.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/examples/)
- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/user_identification/)
