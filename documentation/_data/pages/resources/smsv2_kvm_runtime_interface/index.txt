---
page_title: "xcsh_smsv2_kvm_runtime_interface"
subcategory: ""
description: "Adopts one existing XC-owned KVM Secure Mesh Site v2 SLI child and manages only its DHCP/static IPv4 mode. It never creates or deletes the runtime child."
xcsh_docs: {"aliases": ["smsv2 kvm runtime interface"], "body_bytes": 1485, "body_sha256": "sha256:55d640155e7fce0bada53e3aabee92cc4c4439506784b788cef556459962624f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:smsv2_kvm_runtime_interface:reference", "xcsh-docs:resources:smsv2_kvm_runtime_interface:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/smsv2_kvm_runtime_interface/index.md", "product": "distributed-cloud", "provider_name": "smsv2_kvm_runtime_interface", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0233003001122211-1020213201010101-2002322101121221-3331312002100322-3220301113102331-2301333011333001-3210222301200020-1102130003212303", "registry_path": "docs/resources/smsv2_kvm_runtime_interface.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/smsv2_kvm_runtime_interface/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Adopts one existing XC-owned KVM Secure Mesh Site v2 SLI child and manages only its DHCP/static IPv4 mode. It never creates or deletes the runtime child.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": [], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
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
