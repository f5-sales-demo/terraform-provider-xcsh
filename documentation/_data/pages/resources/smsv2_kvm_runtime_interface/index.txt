---
page_title: "xcsh_smsv2_kvm_runtime_interface"
subcategory: ""
description: "xcsh_smsv2_kvm_runtime_interface for xcsh_smsv2_kvm_runtime_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1485, "body_sha256": "sha256:55d640155e7fce0bada53e3aabee92cc4c4439506784b788cef556459962624f", "child_ids": ["xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "xcsh-docs:resources:smsv2_kvm_runtime_interface:examples"], "collection_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:fundamentals", "parent_id": null, "path": "documentation/resources/smsv2_kvm_runtime_interface/index.md", "provider_name": "smsv2_kvm_runtime_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/smsv2_kvm_runtime_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_smsv2_kvm_runtime_interface for xcsh_smsv2_kvm_runtime_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_smsv2_kvm_runtime_interface

Breadcrumbs:

- xcsh_smsv2_kvm_runtime_interface

Adopts one existing XC-owned KVM Secure Mesh Site v2 SLI child and manages only its DHCP/static IPv4
mode. It never creates or deletes the runtime child.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Adopt the exact XC-owned SLI child discovered after a KVM CE registers.

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 10.0.0"
    }
  }
}

resource "xcsh_smsv2_kvm_runtime_interface" "sli" {
  namespace    = "system"
  site         = "onprem-example-kvm"
  expected_mac = "52:54:00:20:00:11"
  ipv4_cidr    = "10.201.0.11/24"
}
```

## Root configuration

Required root properties: `expected_mac`, `ipv4_cidr`, `site`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/examples/)
