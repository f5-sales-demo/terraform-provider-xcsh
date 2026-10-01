---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_ike_phase2_profile."
xcsh_docs: {"aliases": [], "body_bytes": 1149, "body_sha256": "sha256:d16eedbbc9213bf415f377e22085233266aad82a06e0a0584b0c2d0bbafe4fe7", "canonical_id": "xcsh-docs:data-sources:ike_phase2_profile:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:ike_phase2_profile:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a5a7c67365793986b4f75d4d921baeb441a3835387aa16c35286802039150551", "source_path": "examples/data-sources/xcsh_ike_phase2_profile/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:ike_phase2_profile:example:data-source", "parent_id": "xcsh-docs:data-sources:ike_phase2_profile:examples", "path": "docs/guides/data-sources--ike_phase2_profile--example--data-source.md", "provider_name": "ike_phase2_profile", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase2_profile/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_ike_phase2_profile.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase2_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md)
- [Examples](data-sources--ike_phase2_profile--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ike_phase2_profile/data-source.tf`; digest `sha256:a5a7c67365793986b4f75d4d921baeb441a3835387aa16c35286802039150551`.

```terraform
# IKEPhase2Profile Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IKEPhase2Profile by name
data "xcsh_ike_phase2_profile" "example" {
  name      = "example-ike-phase2-profile"
  namespace = "staging"
}

output "ike_phase2_profile_id" {
  value = data.xcsh_ike_phase2_profile.example.id
}
```

## Next pages

- [Examples](data-sources--ike_phase2_profile--examples.md)
- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md)
