---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_smsv2_kvm_runtime_interface."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1108, "body_sha256": "sha256:c09c9e928ea33c45dd13b945fcac8ee2866d182d2678725d07b899d7cea4bf7e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3203711bd7a88fd8f1838fcf238c8f13492e522a7f5e46972dc68cf409895eb4", "source_path": "examples/resources/xcsh_smsv2_kvm_runtime_interface/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:example:resource", "parent_id": "xcsh-docs:resources:smsv2_kvm_runtime_interface:examples", "path": "documentation/resources/smsv2_kvm_runtime_interface/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "smsv2_kvm_runtime_interface", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2101230030301211-1101032323033100-1030311030320223-1330001120102223-0113331110210010-0323230232320002-0022332030232231-3220223233211313", "registry_path": "docs/guides/resources--smsv2_kvm_runtime_interface--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/smsv2_kvm_runtime_interface/examples/resource/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Resource for xcsh_smsv2_kvm_runtime_interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
