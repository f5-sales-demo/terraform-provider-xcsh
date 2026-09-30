---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_smsv2_kvm_runtime_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1276, "body_sha256": "sha256:087027e756a75de09ae55abf58cf49843daeda6d89474fe8aed4a2bc96e86e8f", "child_ids": [], "collection_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3203711bd7a88fd8f1838fcf238c8f13492e522a7f5e46972dc68cf409895eb4", "source_path": "examples/resources/xcsh_smsv2_kvm_runtime_interface/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:example:resource", "parent_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:examples", "path": "documentation/resources/smsv2_kvm_runtime_interface/examples/resource/index.md", "provider_name": "smsv2_kvm_runtime_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/smsv2_kvm_runtime_interface/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_smsv2_kvm_runtime_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_smsv2_kvm_runtime_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_smsv2_kvm_runtime_interface/resource.tf`; digest `sha256:3203711bd7a88fd8f1838fcf238c8f13492e522a7f5e46972dc68cf409895eb4`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/examples/)
- [xcsh_smsv2_kvm_runtime_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/smsv2_kvm_runtime_interface/)
