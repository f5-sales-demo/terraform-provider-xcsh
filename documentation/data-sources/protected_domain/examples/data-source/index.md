---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_protected_domain."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1091, "body_sha256": "sha256:343e578236f15d5cd266f5cbcba175d5e02135713b703aedfd191806af7aca41", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_domain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:54de5a860ec2aa40dfd3370e44fe450b1b713886394cacf655ec2b67b6b0edd5", "source_path": "examples/data-sources/xcsh_protected_domain/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:protected_domain:example:data-source", "parent_id": "xcsh-docs:data-sources:protected_domain:examples", "path": "documentation/data-sources/protected_domain/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "protected_domain", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0331333123202210-0111123103133310-3212211032231312-0332100021230212-0312010101102121-1210000212010123-0332123323203200-2022111221211102", "registry_path": "docs/guides/data-sources--protected_domain--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_domain/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_protected_domain.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["protected_domainCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_protected_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_domain/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_domain/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_protected_domain/data-source.tf`; digest `sha256:54de5a860ec2aa40dfd3370e44fe450b1b713886394cacf655ec2b67b6b0edd5`.

```terraform
# ProtectedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtectedDomain by name
data "xcsh_protected_domain" "example" {
  name      = "example-protected-domain"
  namespace = "staging"
}

output "protected_domain_id" {
  value = data.xcsh_protected_domain.example.id
}
```
