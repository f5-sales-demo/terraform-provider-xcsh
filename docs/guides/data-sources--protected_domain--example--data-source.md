---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_protected_domain."
xcsh_docs: {"aliases": [], "body_bytes": 1125, "body_sha256": "sha256:5626a20d5c02c7bbc2b9da095aaf5ac8d5703c110192beb8d33337f06446e2f1", "canonical_id": "xcsh-docs:data-sources:protected_domain:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:protected_domain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:54de5a860ec2aa40dfd3370e44fe450b1b713886394cacf655ec2b67b6b0edd5", "source_path": "examples/data-sources/xcsh_protected_domain/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:protected_domain:example:data-source", "parent_id": "xcsh-docs:data-sources:protected_domain:examples", "path": "docs/guides/data-sources--protected_domain--example--data-source.md", "provider_name": "protected_domain", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_domain/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_protected_domain.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_domainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_protected_domain](../data-sources/protected_domain.md)
- [Examples](data-sources--protected_domain--examples.md)
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

- [Examples](data-sources--protected_domain--examples.md)
- [xcsh_protected_domain](../data-sources/protected_domain.md)
