---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_fast_acl."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1226, "body_sha256": "sha256:60a4ca912a20bf4cdc931e2e9a6b8f5b84ebb287829018d8c09e4bc65a895cd7", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8fad5fdbb88c4dc30f2314eb3a4717e4d1278f83525eb02233ddf9ff899b963a", "source_path": "examples/data-sources/xcsh_fast_acl/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:fast_acl:example:data-source", "parent_id": "xcsh-docs:data-sources:fast_acl:examples", "path": "documentation/data-sources/fast_acl/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3033302120112100-1131213013222223-1301013212132100-1101102233301022-3120003130000031-1201303311111122-1002332203203321-2323321120031222", "registry_path": "docs/guides/data-sources--fast_acl--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_fast_acl.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_fast_acl/data-source.tf`; digest `sha256:8fad5fdbb88c4dc30f2314eb3a4717e4d1278f83525eb02233ddf9ff899b963a`.

```terraform
# FastACL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FastACL by name
data "xcsh_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"
}

output "fast_acl_id" {
  value = data.xcsh_fast_acl.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/examples/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
