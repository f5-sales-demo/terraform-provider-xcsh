---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_addon_service_activation_status."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1206, "body_sha256": "sha256:8ed8e35d8c07d1f9e010043d4e8f4d2318bed6d8d7b409a12a0492dc4b7bdaed", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:addon_service_activation_status:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:73b6eb2a08b57df6e1279cca2e1bd3688fc9db605ed1feb8ab5a2418a8f37cc0", "source_path": "examples/data-sources/xcsh_addon_service_activation_status/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:addon_service_activation_status:example:data-source", "parent_id": "xcsh-docs:data-sources:addon_service_activation_status:examples", "path": "documentation/data-sources/addon_service_activation_status/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "addon_service_activation_status", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3213023131322221-0101000021122220-0110100212033121-1332202133222033-3123311202220332-0112123331320111-1233223210022300-2211121222130212", "registry_path": "docs/guides/data-sources--addon_service_activation_status--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/addon_service_activation_status/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_addon_service_activation_status.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_addon_service_activation_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service_activation_status/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service_activation_status/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_addon_service_activation_status/data-source.tf`; digest `sha256:73b6eb2a08b57df6e1279cca2e1bd3688fc9db605ed1feb8ab5a2418a8f37cc0`.

```terraform
# AddonServiceActivationStatus Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Check the tenant's Client-Side Defense subscription.
data "xcsh_addon_service_activation_status" "example" {
  addon_service = "f5xc-client-side-defense-standard"
}

output "addon_service_activation_state" {
  value = data.xcsh_addon_service_activation_status.example.state
}
```
