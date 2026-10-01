---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_addon_service_activation_status."
xcsh_docs: {"aliases": [], "body_bytes": 1285, "body_sha256": "sha256:7d6bbc8190aaf69e615b14d3a19d8cab9bf1710e46be35341929f063a17709a0", "canonical_id": "xcsh-docs:data-sources:addon_service_activation_status:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:addon_service_activation_status:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:73b6eb2a08b57df6e1279cca2e1bd3688fc9db605ed1feb8ab5a2418a8f37cc0", "source_path": "examples/data-sources/xcsh_addon_service_activation_status/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:addon_service_activation_status:example:data-source", "parent_id": "xcsh-docs:data-sources:addon_service_activation_status:examples", "path": "docs/guides/data-sources--addon_service_activation_status--example--data-source.md", "provider_name": "addon_service_activation_status", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/addon_service_activation_status/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_addon_service_activation_status.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md)
- [Examples](data-sources--addon_service_activation_status--examples.md)
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

## Next pages

- [Examples](data-sources--addon_service_activation_status--examples.md)
- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md)
