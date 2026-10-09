---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_smsv2_kvm_runtime."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1215, "body_sha256": "sha256:a156b055b361ff1ea114ad4bd1ad9ad5c749a31867f58e449b0b79bcc1e9b97c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:59afa846179726aea892d691cd1d6b9fa1a3b1c0e460e92807612ddcb56e2ba6", "source_path": "examples/data-sources/xcsh_smsv2_kvm_runtime/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:smsv2_kvm_runtime:example:data-source", "parent_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:examples", "path": "documentation/data-sources/smsv2_kvm_runtime/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "smsv2_kvm_runtime", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2120112323021230-0321032311122313-2320331030113323-0303130013123001-1303231022102231-0203013112213213-1200131131101112-2332121310100202", "registry_path": "docs/guides/data-sources--smsv2_kvm_runtime--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_kvm_runtime/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_smsv2_kvm_runtime.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_smsv2_kvm_runtime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_smsv2_kvm_runtime/data-source.tf`; digest `sha256:59afa846179726aea892d691cd1d6b9fa1a3b1c0e460e92807612ddcb56e2ba6`.

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
