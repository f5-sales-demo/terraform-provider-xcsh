---
page_title: "xcsh_smsv2_kvm_runtime"
subcategory: ""
description: "xcsh_smsv2_kvm_runtime for xcsh_smsv2_kvm_runtime."
xcsh_docs: {"aliases": [], "body_bytes": 1434, "body_sha256": "sha256:87b50594961f2800b9c9cf711b50bf7b50ccce33a76867c3c0f293b5c7571b7c", "canonical_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:fundamentals", "child_ids": ["xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "xcsh-docs:data-sources:smsv2_kvm_runtime:examples"], "collection_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_kvm_runtime:fundamentals", "parent_id": null, "path": "docs/data-sources/smsv2_kvm_runtime.md", "provider_name": "smsv2_kvm_runtime", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_kvm_runtime/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_smsv2_kvm_runtime for xcsh_smsv2_kvm_runtime.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_smsv2_kvm_runtime

Breadcrumbs:

- xcsh_smsv2_kvm_runtime

Resolves one realized KVM Secure Mesh Site v2 SLO network interface through site UID ownership, the
live registration hostname and device, and an expected MAC address. The name is observed, never
guessed.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Resolve the exact XC interface object created for one registered KVM CE.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 9.5.2"
    }
  }
}

data "xcsh_smsv2_kvm_runtime" "ce" {
  namespace    = "system"
  site         = "example-kvm-smsv2-site"
  expected_mac = "52:54:00:10:00:11"
}

output "kvm_interface_name" {
  value = data.xcsh_smsv2_kvm_runtime.ce.interface_name
}

output "kvm_registration_device" {
  value = data.xcsh_smsv2_kvm_runtime.ce.device
}
```

## Root configuration

Required root properties: `expected_mac`, `site`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--smsv2_kvm_runtime--reference.md)
- [Examples](../guides/data-sources--smsv2_kvm_runtime--examples.md)
