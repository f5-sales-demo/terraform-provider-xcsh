---
page_title: "xcsh_ike_phase1_profile"
subcategory: ""
description: "Reads IKE Phase1 Profile information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["ike phase1 profile"], "body_bytes": 1381, "body_sha256": "sha256:98a5e88d5dd04c88831c3c9da079cbbb49b3c7f4bfe5b37ab3b7cd4293e9a575", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:ike_phase1_profile:reference", "xcsh-docs:data-sources:ike_phase1_profile:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike_phase1_profile:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/ike_phase1_profile/index.md", "product": "distributed-cloud", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-1332210000333121-0321333020120302-3121233211013331-1320231330211033-2111031001002303-1300211301302100-1031311212302312-1213122302330023", "registry_path": "docs/data-sources/ike_phase1_profile.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase1_profile/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Reads IKE Phase1 Profile information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_ike_phase1_profile

Breadcrumbs:

- xcsh_ike_phase1_profile

Reads IKE Phase1 Profile information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase1_profile/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase1_profile/examples/)
