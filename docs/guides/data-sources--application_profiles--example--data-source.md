---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_application_profiles."
xcsh_docs: {"aliases": [], "body_bytes": 1177, "body_sha256": "sha256:e016819c0548c1c674d537abe9dd73d4073c7c6d75bf4160f7d796feb5aa4cab", "canonical_id": "xcsh-docs:data-sources:application_profiles:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:653702219239c8ed2bc8016fb3a0ffd69a3b06f91231711272fb74352db423cc", "source_path": "examples/data-sources/xcsh_application_profiles/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:application_profiles:example:data-source", "parent_id": "xcsh-docs:data-sources:application_profiles:examples", "path": "docs/guides/data-sources--application_profiles--example--data-source.md", "provider_name": "application_profiles", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_application_profiles.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["application_profilesCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md)
- [Examples](data-sources--application_profiles--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_application_profiles/data-source.tf`; digest `sha256:653702219239c8ed2bc8016fb3a0ffd69a3b06f91231711272fb74352db423cc`.

```terraform
# ApplicationProfiles Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ApplicationProfiles by name
data "xcsh_application_profiles" "example" {
  name      = "example-application-profiles"
  namespace = "staging"
}

output "application_profiles_id" {
  value = data.xcsh_application_profiles.example.id
}
```

## Next pages

- [Examples](data-sources--application_profiles--examples.md)
- [xcsh_application_profiles](../data-sources/application_profiles.md)
