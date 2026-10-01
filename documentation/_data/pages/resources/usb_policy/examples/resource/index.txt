---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_usb_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1209, "body_sha256": "sha256:35cda0b5a6a7fd79c315147accd167168b803929761f20c86476488fb1f0ff70", "child_ids": [], "collection_id": "xcsh-docs:resources:usb_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d81199b0fb106ffd6cc8d506207e603f1169d2cbf87d6b31f0ac720b8b44794e", "source_path": "examples/resources/xcsh_usb_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:usb_policy:example:resource", "parent_id": "xcsh-docs:resources:usb_policy:examples", "path": "documentation/resources/usb_policy/examples/resource/index.md", "provider_name": "usb_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/usb_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_usb_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["usb_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_usb_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_usb_policy/resource.tf`; digest `sha256:d81199b0fb106ffd6cc8d506207e603f1169d2cbf87d6b31f0ac720b8b44794e`.

```terraform
# UsbPolicy Resource Example
# Manages new USB policy object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UsbPolicy configuration
resource "xcsh_usb_policy" "example" {
  name      = "example-usb-policy"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/examples/)
- [xcsh_usb_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/)
