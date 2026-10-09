---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_ike_phase1_profile."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1109, "body_sha256": "sha256:e80fedb1f9a03e632f23e1c0997d10fe52967aa3d0f55da3d3ef83c169875345", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike_phase1_profile:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a22af56dca6a3b92413f9588bd02d5a8a0c4c469599c9023c9feb46580028ab5", "source_path": "examples/data-sources/xcsh_ike_phase1_profile/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:ike_phase1_profile:example:data-source", "parent_id": "xcsh-docs:data-sources:ike_phase1_profile:examples", "path": "documentation/data-sources/ike_phase1_profile/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3323210120231233-1120213301210313-1322301121031313-2031222010201011-0012332201233030-1121233113122111-3111220230221103-2211022201310130", "registry_path": "docs/guides/data-sources--ike_phase1_profile--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase1_profile/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_ike_phase1_profile.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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
