---
page_title: "xcsh_ike_phase2_profile"
subcategory: ""
description: "Manages a IKE Phase2 Profile resource in F5 Distributed Cloud for ike phase2 profile specification. configuration."
xcsh_docs: {"aliases": ["ike phase2 profile"], "body_bytes": 1419, "body_sha256": "sha256:99956ef809e03b1bd48552584b4e91501c21de1db3d1c588247d458579d267fd", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:ike_phase2_profile:reference", "xcsh-docs:data-sources:ike_phase2_profile:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike_phase2_profile:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike_phase2_profile:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/ike_phase2_profile/index.md", "product": "distributed-cloud", "provider_name": "ike_phase2_profile", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3110330013310011-0121322330132310-2300311200301300-0223012032021201-1220120002002021-3100202000033331-0211031202223101-3022233301301130", "registry_path": "docs/data-sources/ike_phase2_profile.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase2_profile/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a IKE Phase2 Profile resource in F5 Distributed Cloud for ike phase2 profile specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike_phase2_profileCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_ike_phase2_profile

Breadcrumbs:

- xcsh_ike_phase2_profile

Manages a IKE Phase2 Profile resource in F5 Distributed Cloud for ike phase2 profile specification.
configuration.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase2_profile/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase2_profile/examples/)
