---
page_title: "xcsh_usb_policy"
subcategory: ""
description: "xcsh_usb_policy for xcsh_usb_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1284, "body_sha256": "sha256:041916fd9de919a0491e54e56bfbed0a28ce5c289a8f8ab56b7c8f7a144b0b27", "canonical_id": "xcsh-docs:resources:usb_policy:fundamentals", "child_ids": ["xcsh-docs:resources:usb_policy:reference", "xcsh-docs:resources:usb_policy:examples", "xcsh-docs:resources:usb_policy:import", "xcsh-docs:resources:usb_policy:timeouts"], "collection_id": "xcsh-docs:resources:usb_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:usb_policy:fundamentals", "parent_id": null, "path": "docs/resources/usb_policy.md", "provider_name": "usb_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/usb_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_usb_policy for xcsh_usb_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["usb_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_usb_policy

Breadcrumbs:

- xcsh_usb_policy

Manages new USB policy object in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--usb_policy--reference.md)
- [Examples](../guides/resources--usb_policy--examples.md)
- [Import](../guides/resources--usb_policy--import.md)
- [Timeouts](../guides/resources--usb_policy--timeouts.md)
