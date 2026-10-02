---
page_title: "xcsh_smsv2_kvm_runtime"
subcategory: ""
description: "Resolves one realized KVM Secure Mesh Site v2 SLO network interface through site UID ownership, the live registration hostname and device, and an expected MAC address. The name is observed, never guessed."
xcsh_docs: {"aliases": ["smsv2 kvm runtime"], "body_bytes": 1618, "body_sha256": "sha256:6df1d95e9b13bad3e2ed7035bf382e1243cdeb820eaead47855f07123b1a35fe", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:smsv2_kvm_runtime:reference", "xcsh-docs:data-sources:smsv2_kvm_runtime:examples"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_kvm_runtime:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/smsv2_kvm_runtime/index.md", "product": "distributed-cloud", "provider_name": "smsv2_kvm_runtime", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3112321232333312-1310102333312112-1232330010223001-2220301030033011-2013200023031111-0001200213232313-3011013301302013-2000100013012202", "registry_path": "docs/data-sources/smsv2_kvm_runtime.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_kvm_runtime/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resolves one realized KVM Secure Mesh Site v2 SLO network interface through site UID ownership, the live registration hostname and device, and an expected MAC address. The name is observed, never guessed.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/examples/)
