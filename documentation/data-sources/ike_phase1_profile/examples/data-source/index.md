---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_ike_phase1_profile."
xcsh_docs: {"aliases": [], "body_bytes": 1355, "body_sha256": "sha256:c1365bd872d5086aeb0cf9da3cf7ab39e5656ef9190bf680db73234cfd7b50f4", "child_ids": [], "collection_id": "xcsh-docs:data-sources:ike_phase1_profile:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a22af56dca6a3b92413f9588bd02d5a8a0c4c469599c9023c9feb46580028ab5", "source_path": "examples/data-sources/xcsh_ike_phase1_profile/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:ike_phase1_profile:example:data-source", "parent_id": "xcsh-docs:data-sources:ike_phase1_profile:examples", "path": "documentation/data-sources/ike_phase1_profile/examples/data-source/index.md", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase1_profile/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_ike_phase1_profile.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase1_profile/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase1_profile/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ike_phase1_profile/data-source.tf`; digest `sha256:a22af56dca6a3b92413f9588bd02d5a8a0c4c469599c9023c9feb46580028ab5`.

```terraform
# IKEPhase1Profile Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IKEPhase1Profile by name
data "xcsh_ike_phase1_profile" "example" {
  name      = "example-ike-phase1-profile"
  namespace = "staging"
}

output "ike_phase1_profile_id" {
  value = data.xcsh_ike_phase1_profile.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase1_profile/examples/)
- [xcsh_ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase1_profile/)
