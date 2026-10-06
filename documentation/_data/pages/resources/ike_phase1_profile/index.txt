---
page_title: "xcsh_ike_phase1_profile"
subcategory: ""
description: "Manages an IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification. configuration."
xcsh_docs: {"aliases": ["ike phase1 profile"], "body_bytes": 1906, "body_sha256": "sha256:a851e3914e0f2a46218dcae02cd0b62e5e7d6759ec227e56a3f5bb75b352119f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:ike_phase1_profile:reference", "xcsh-docs:resources:ike_phase1_profile:examples", "xcsh-docs:resources:ike_phase1_profile:import", "xcsh-docs:resources:ike_phase1_profile:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike_phase1_profile:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/ike_phase1_profile/index.md", "product": "distributed-cloud", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220", "registry_path": "docs/resources/ike_phase1_profile.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase1_profile/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Manages an IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_ike_phase1_profile

Breadcrumbs:

- xcsh_ike_phase1_profile

Manages an IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification.
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/lifecycle/timeouts/)
