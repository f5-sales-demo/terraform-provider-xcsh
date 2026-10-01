---
page_title: "xcsh_ike_phase1_profile"
subcategory: ""
description: "xcsh_ike_phase1_profile for xcsh_ike_phase1_profile."
xcsh_docs: {"aliases": [], "body_bytes": 1703, "body_sha256": "sha256:f7416f795ab4f63a4f148d9e2aa23e9fbd84f3584b2cd4d101fa592e964d6c06", "canonical_id": "xcsh-docs:resources:ike_phase1_profile:fundamentals", "child_ids": ["xcsh-docs:resources:ike_phase1_profile:reference", "xcsh-docs:resources:ike_phase1_profile:examples", "xcsh-docs:resources:ike_phase1_profile:import", "xcsh-docs:resources:ike_phase1_profile:timeouts"], "collection_id": "xcsh-docs:resources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike_phase1_profile:fundamentals", "parent_id": null, "path": "docs/resources/ike_phase1_profile.md", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase1_profile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_ike_phase1_profile for xcsh_ike_phase1_profile.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_ike_phase1_profile

Breadcrumbs:

- xcsh_ike_phase1_profile

Manages a IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification.
configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IKEPhase1Profile Resource Example
# Manages a IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IKEPhase1Profile configuration
resource "xcsh_ike_phase1_profile" "example" {
  name      = "example-ike-phase1-profile"
  namespace = "staging"

  authentication_algos = ["example-value"]
  dh_group             = ["example-value"]
  encryption_algos     = ["example-value"]
  prf                  = ["example-value"]
}
```

## Root configuration

Required root properties: `authentication_algos`, `dh_group`, `encryption_algos`, `name`, `namespace`, `prf`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--ike_phase1_profile--reference.md)
- [Examples](../guides/resources--ike_phase1_profile--examples.md)
- [Import](../guides/resources--ike_phase1_profile--import.md)
- [Timeouts](../guides/resources--ike_phase1_profile--timeouts.md)
