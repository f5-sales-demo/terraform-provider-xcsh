---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_ike_phase1_profile."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1273, "body_sha256": "sha256:90ea274ada675d6f0600ef556ffdd228e2d5c8884937dac2384e475b7cf446f0", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike_phase1_profile:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8b0906176c52164cc70b7c3ef3103bf969192548192507a8a9df9fa2b687acf6", "source_path": "examples/resources/xcsh_ike_phase1_profile/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:ike_phase1_profile:example:resource", "parent_id": "xcsh-docs:resources:ike_phase1_profile:examples", "path": "documentation/resources/ike_phase1_profile/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2321102110313013-1133331010000002-2232020212000012-1003221021302333-3001330231230132-0112111103213021-0113020023311031-1003110212100232", "registry_path": "docs/guides/resources--ike_phase1_profile--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase1_profile/examples/resource/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Resource for xcsh_ike_phase1_profile.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_ike_phase1_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase1_profile/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ike_phase1_profile/resource.tf`; digest `sha256:8b0906176c52164cc70b7c3ef3103bf969192548192507a8a9df9fa2b687acf6`.

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
