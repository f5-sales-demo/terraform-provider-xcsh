---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_ike_phase2_profile."
xcsh_docs: {"aliases": [], "body_bytes": 1221, "body_sha256": "sha256:6762f4cc85d5d103fbec646d5aff5c0a774c69f626129bbb5db9de148009e43e", "canonical_id": "xcsh-docs:resources:ike_phase2_profile:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:ike_phase2_profile:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:80e0e1a4beab03b8fb465c2b172d0e27a9d0b25515c600ec2364548afb3d3ffd", "source_path": "examples/resources/xcsh_ike_phase2_profile/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:ike_phase2_profile:example:resource", "parent_id": "xcsh-docs:resources:ike_phase2_profile:examples", "path": "docs/guides/resources--ike_phase2_profile--example--resource.md", "provider_name": "ike_phase2_profile", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase2_profile/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_ike_phase2_profile.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase2_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md)
- [Examples](resources--ike_phase2_profile--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ike_phase2_profile/resource.tf`; digest `sha256:80e0e1a4beab03b8fb465c2b172d0e27a9d0b25515c600ec2364548afb3d3ffd`.

```terraform
# IKEPhase2Profile Resource Example
# Manages a IKE Phase2 Profile resource in F5 Distributed Cloud for ike phase2 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IKEPhase2Profile configuration
resource "xcsh_ike_phase2_profile" "example" {
  name      = "example-ike-phase2-profile"
  namespace = "staging"

  authentication_algos = ["example-value"]
  encryption_algos     = ["example-value"]
}
```

## Next pages

- [Examples](resources--ike_phase2_profile--examples.md)
- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md)
