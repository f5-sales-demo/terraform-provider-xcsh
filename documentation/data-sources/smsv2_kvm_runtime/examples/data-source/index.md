---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_smsv2_kvm_runtime."
xcsh_docs: {"aliases": [], "body_bytes": 1359, "body_sha256": "sha256:fe2cb158fec4432d93bea89c1a3a284ce068a53cab2868f57aa8cc326e8c8c14", "child_ids": [], "collection_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:59afa846179726aea892d691cd1d6b9fa1a3b1c0e460e92807612ddcb56e2ba6", "source_path": "examples/data-sources/xcsh_smsv2_kvm_runtime/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:smsv2_kvm_runtime:example:data-source", "parent_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:examples", "path": "documentation/data-sources/smsv2_kvm_runtime/examples/data-source/index.md", "provider_name": "smsv2_kvm_runtime", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_kvm_runtime/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_smsv2_kvm_runtime.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/examples/)
- [xcsh_smsv2_kvm_runtime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/)
