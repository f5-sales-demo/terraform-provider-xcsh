---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_user_identification."
xcsh_docs: {"aliases": [], "body_bytes": 1065, "body_sha256": "sha256:09e8a82e3118a3a8a7b461ff56d16cdd6fda51676c07638051fce46bae975242", "canonical_id": "xcsh-docs:data-sources:user_identification:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:user_identification:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c30020dfd6aecd49ffb0273704b9b9e1c85f494c0e9e28cd883c479b6b6234a3", "source_path": "examples/data-sources/xcsh_user_identification/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:user_identification:example:data-source", "parent_id": "xcsh-docs:data-sources:user_identification:examples", "path": "docs/guides/data-sources--user_identification--example--data-source.md", "provider_name": "user_identification", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/user_identification/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_user_identification.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md)
- [Examples](data-sources--user_identification--examples.md)
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

- [Examples](data-sources--user_identification--examples.md)
- [xcsh_user_identification](../data-sources/user_identification.md)
