---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_usb_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1253, "body_sha256": "sha256:e1021f32143078567a2bdb9822b38d685cdebcda05d496e67c737a1fe87937fb", "child_ids": [], "collection_id": "xcsh-docs:data-sources:usb_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:40aac30ea686e00cce92cecb2aee01cb69e14aef947ae725b686bf0b6e9af7b0", "source_path": "examples/data-sources/xcsh_usb_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:usb_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:usb_policy:examples", "path": "documentation/data-sources/usb_policy/examples/data-source/index.md", "provider_name": "usb_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/usb_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_usb_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["usb_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_usb_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/usb_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/usb_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_usb_policy/data-source.tf`; digest `sha256:40aac30ea686e00cce92cecb2aee01cb69e14aef947ae725b686bf0b6e9af7b0`.

```terraform
# UsbPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UsbPolicy by name
data "xcsh_usb_policy" "example" {
  name      = "example-usb-policy"
  namespace = "staging"
}

output "usb_policy_id" {
  value = data.xcsh_usb_policy.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/usb_policy/examples/)
- [xcsh_usb_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/usb_policy/)
