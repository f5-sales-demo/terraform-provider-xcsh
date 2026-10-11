---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_ike_phase2_profile."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1187, "body_sha256": "sha256:f2ce22a0968d993f10bb5088e9c050707d16b3adb540598b479ba1c4911ccc8d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike_phase2_profile:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:80e0e1a4beab03b8fb465c2b172d0e27a9d0b25515c600ec2364548afb3d3ffd", "source_path": "examples/resources/xcsh_ike_phase2_profile/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:ike_phase2_profile:example:resource", "parent_id": "xcsh-docs:resources:ike_phase2_profile:examples", "path": "documentation/resources/ike_phase2_profile/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "ike_phase2_profile", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1220231231012030-2030203220222323-3133212110332020-0000003000321201-3132330323300003-2332102320201201-2210233330213220-2101030123033212", "registry_path": "docs/guides/resources--ike_phase2_profile--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase2_profile/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_ike_phase2_profile.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["ike_phase2_profileCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/examples/)
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
