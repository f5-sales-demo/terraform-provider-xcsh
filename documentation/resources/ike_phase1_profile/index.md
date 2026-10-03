---
page_title: "xcsh_ike_phase1_profile"
subcategory: ""
description: "Manages a IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification. configuration."
xcsh_docs: {"aliases": ["ike phase1 profile"], "body_bytes": 1892, "body_sha256": "sha256:2c2b074a34cd1e27de4e86ae3a9815da2e6be629b4f560fbcea65674c1b2b29b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:ike_phase1_profile:reference", "xcsh-docs:resources:ike_phase1_profile:examples", "xcsh-docs:resources:ike_phase1_profile:import", "xcsh-docs:resources:ike_phase1_profile:timeouts"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike_phase1_profile:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/ike_phase1_profile/index.md", "product": "distributed-cloud", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220", "registry_path": "docs/resources/ike_phase1_profile.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase1_profile/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/lifecycle/timeouts/)
