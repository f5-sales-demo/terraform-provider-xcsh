---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_smsv2_kvm_runtime."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1458, "body_sha256": "sha256:1288d3b561f78fd812306f0d9005e09af78022a9db2158b1a4288421e76a49c4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:59afa846179726aea892d691cd1d6b9fa1a3b1c0e460e92807612ddcb56e2ba6", "source_path": "examples/data-sources/xcsh_smsv2_kvm_runtime/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:smsv2_kvm_runtime:example:data-source", "parent_id": "xcsh-docs:data-sources:smsv2_kvm_runtime:examples", "path": "documentation/data-sources/smsv2_kvm_runtime/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "smsv2_kvm_runtime", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2120112323021230-0321032311122313-2320331030113323-0303130013123001-1303231022102231-0203013112213213-1200131131101112-2332121310100202", "registry_path": "docs/guides/data-sources--smsv2_kvm_runtime--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_kvm_runtime/examples/data-source/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Data source for xcsh_smsv2_kvm_runtime.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/examples/)
- [xcsh_smsv2_kvm_runtime](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/smsv2_kvm_runtime/)
