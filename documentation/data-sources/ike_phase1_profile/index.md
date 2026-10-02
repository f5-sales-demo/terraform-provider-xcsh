---
page_title: "xcsh_ike_phase1_profile"
subcategory: ""
description: "Manages a IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification. configuration."
xcsh_docs: {"aliases": ["ike phase1 profile"], "body_bytes": 1419, "body_sha256": "sha256:548b29f0839e2f09757feff30a241f8ef84594bf4325f3e9e3ac94dc4ce8440e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:ike_phase1_profile:reference", "xcsh-docs:data-sources:ike_phase1_profile:examples"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike_phase1_profile:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike_phase1_profile:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/ike_phase1_profile/index.md", "product": "distributed-cloud", "provider_name": "ike_phase1_profile", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1332210000333121-0321333020120302-3121233211013331-1320231330211033-2111031001002303-1300211301302100-1031311212302312-1213122302330023", "registry_path": "docs/data-sources/ike_phase1_profile.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike_phase1_profile/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a IKE Phase1 Profile resource in F5 Distributed Cloud for ike phase1 profile specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike_phase1_profileCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase1_profile/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike_phase1_profile/examples/)
