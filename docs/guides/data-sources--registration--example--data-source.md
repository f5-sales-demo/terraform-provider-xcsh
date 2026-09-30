---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 976, "body_sha256": "sha256:7681ffa409f964c83b25ba0a6d6faa3217d8a37bee9ec035167c3556bba8f33e", "canonical_id": "xcsh-docs:data-sources:registration:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:887504fe79e23b9f0b1ae62f2ad6a04c459d805004c72b40be81620b6a16261f", "source_path": "examples/data-sources/xcsh_registration/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:registration:example:data-source", "parent_id": "xcsh-docs:data-sources:registration:examples", "path": "docs/guides/data-sources--registration--example--data-source.md", "provider_name": "registration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_registration](../data-sources/registration.md)
- [Examples](data-sources--registration--examples.md)
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

- [Examples](data-sources--registration--examples.md)
- [xcsh_registration](../data-sources/registration.md)
