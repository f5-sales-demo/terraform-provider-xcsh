---
page_title: "xcsh_ike_phase2_profile"
subcategory: ""
description: "xcsh_ike_phase2_profile for xcsh_ike_phase2_profile."
xcsh_docs: {"aliases": [], "body_bytes": 1235, "body_sha256": "sha256:f75d4f101f21d3c3c5bea130f23ebc2fa80227952b1f1574f6b3b13004d21cef", "canonical_id": "xcsh-docs:data-sources:ike_phase2_profile:fundamentals", "child_ids": ["xcsh-docs:data-sources:ike_phase2_profile:reference", "xcsh-docs:data-sources:ike_phase2_profile:examples"], "collection_id": "xcsh-docs:data-sources:ike_phase2_profile:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike_phase2_profile:fundamentals", "parent_id": null, "path": "docs/data-sources/ike_phase2_profile.md", "provider_name": "ike_phase2_profile", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase2_profile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_ike_phase2_profile for xcsh_ike_phase2_profile.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase2_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_ike_phase2_profile

Breadcrumbs:

- xcsh_ike_phase2_profile

Manages a IKE Phase2 Profile resource in F5 Distributed Cloud for ike phase2 profile specification.
configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--ike_phase2_profile--reference.md)
- [Examples](../guides/data-sources--ike_phase2_profile--examples.md)
