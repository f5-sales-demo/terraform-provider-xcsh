---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_allowed_domain."
xcsh_docs: {"aliases": [], "body_bytes": 1206, "body_sha256": "sha256:a4d0fb36636b15c054446958e1665ab5c6bb9a3ffd01baf49d2844a5d94bb6ce", "child_ids": [], "collection_id": "xcsh-docs:data-sources:allowed_domain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e22d7f6871a55bc1c9dd4658347afbf10c80e91173e9c9082eacc38fb5713006", "source_path": "examples/data-sources/xcsh_allowed_domain/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:allowed_domain:example:data-source", "parent_id": "xcsh-docs:data-sources:allowed_domain:examples", "path": "documentation/data-sources/allowed_domain/examples/data-source/index.md", "provider_name": "allowed_domain", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/allowed_domain/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_allowed_domain.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["allowed_domainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_allowed_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/allowed_domain/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/allowed_domain/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_allowed_domain/data-source.tf`; digest `sha256:e22d7f6871a55bc1c9dd4658347afbf10c80e91173e9c9082eacc38fb5713006`.

```terraform
# AllowedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AllowedDomain by name
data "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"
}

output "allowed_domain_id" {
  value = data.xcsh_allowed_domain.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/allowed_domain/examples/)
- [xcsh_allowed_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/allowed_domain/)
