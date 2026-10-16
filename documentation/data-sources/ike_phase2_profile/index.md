---
page_title: "xcsh_ike_phase2_profile"
subcategory: ""
description: "Reads IKE Phase2 Profile information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["ike phase2 profile"], "body_bytes": 1381, "body_sha256": "sha256:4889d639e0b283d5e24cb21618cc6ca85f5ffc5d93ff6c25a65a34dae9a3b0a2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:ike_phase2_profile:reference", "xcsh-docs:data-sources:ike_phase2_profile:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike_phase2_profile:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike_phase2_profile:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/ike_phase2_profile/index.md", "product": "distributed-cloud", "provider_name": "ike_phase2_profile", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-3110330013310011-0121322330132310-2300311200301300-0223012032021201-1220120002002021-3100202000033331-0211031202223101-3022233301301130", "registry_path": "docs/data-sources/ike_phase2_profile.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase2_profile/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Reads IKE Phase2 Profile information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["ike_phase2_profileCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_ike_phase2_profile

Breadcrumbs:

- xcsh_ike_phase2_profile

Reads IKE Phase2 Profile information from F5 Distributed Cloud.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase2_profile/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase2_profile/examples/)
