---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_ike_phase1_profile."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1355, "body_sha256": "sha256:c1365bd872d5086aeb0cf9da3cf7ab39e5656ef9190bf680db73234cfd7b50f4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike_phase1_profile:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a22af56dca6a3b92413f9588bd02d5a8a0c4c469599c9023c9feb46580028ab5", "source_path": "examples/data-sources/xcsh_ike_phase1_profile/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:ike_phase1_profile:example:data-source", "parent_id": "xcsh-docs:data-sources:ike_phase1_profile:examples", "path": "documentation/data-sources/ike_phase1_profile/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3323210120231233-1120213301210313-1322301121031313-2031222010201011-0012332201233030-1121233113122111-3111220230221103-2211022201310130", "registry_path": "docs/guides/data-sources--ike_phase1_profile--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase1_profile/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_ike_phase1_profile.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
