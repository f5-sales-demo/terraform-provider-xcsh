---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_registration."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1281, "body_sha256": "sha256:507e62c23aba3e45bc4a9fd609ba217db967836beb164ecd77b7e2088136fbf5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:887504fe79e23b9f0b1ae62f2ad6a04c459d805004c72b40be81620b6a16261f", "source_path": "examples/data-sources/xcsh_registration/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:registration:example:data-source", "parent_id": "xcsh-docs:data-sources:registration:examples", "path": "documentation/data-sources/registration/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3002131203233232-1020331232213323-2011131113233312-3303311323331302-3133132013002211-2233013223212112-0310212021011031-1101133030233022", "registry_path": "docs/guides/data-sources--registration--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_registration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["registrationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_registration/data-source.tf`; digest `sha256:887504fe79e23b9f0b1ae62f2ad6a04c459d805004c72b40be81620b6a16261f`.

```terraform
# Registration Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Registration by name
data "xcsh_registration" "example" {
  name      = "example-registration"
  namespace = "staging"
}

output "registration_id" {
  value = data.xcsh_registration.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/examples/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
