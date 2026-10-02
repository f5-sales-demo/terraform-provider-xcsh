---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_protected_domain."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1331, "body_sha256": "sha256:dd260ba829be13ff22d31ceb29c08e6cd9a0ebedc820e7ec62f2e36b5f0785e8", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_domain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:54de5a860ec2aa40dfd3370e44fe450b1b713886394cacf655ec2b67b6b0edd5", "source_path": "examples/data-sources/xcsh_protected_domain/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:protected_domain:example:data-source", "parent_id": "xcsh-docs:data-sources:protected_domain:examples", "path": "documentation/data-sources/protected_domain/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "protected_domain", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0331333123202210-0111123103133310-3212211032231312-0332100021230212-0312010101102121-1210000212010123-0332123323203200-2022111221211102", "registry_path": "docs/guides/data-sources--protected_domain--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_domain/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_protected_domain.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protected_domainCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_domain/examples/)
- [xcsh_protected_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_domain/)
