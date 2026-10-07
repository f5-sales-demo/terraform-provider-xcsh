---
page_title: "Resource"
subcategory: "Networking"
description: "Resource for xcsh_virtual_network."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1045, "body_sha256": "sha256:3fe30db77947ff4fc4ebe4caa3ac0955196031cc945e298f981e6e27665a4e19", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:virtual_network:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:24ff8ad3e56ba4ece18f3beaaabdcc55805f2e25c74acd17556111d54902f473", "source_path": "examples/resources/xcsh_virtual_network/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:virtual_network:example:resource", "parent_id": "xcsh-docs:resources:virtual_network:examples", "path": "documentation/resources/virtual_network/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "virtual_network", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1322123120312230-3132110232130120-0302230100113022-0331311030201303-3002210133220221-3332123213020313-3120112033301002-2020030121000232", "registry_path": "docs/guides/resources--virtual_network--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_network/examples/resource/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Resource for xcsh_virtual_network.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_network/resource.tf`; digest `sha256:24ff8ad3e56ba4ece18f3beaaabdcc55805f2e25c74acd17556111d54902f473`.

```terraform
# VirtualNetwork Resource Example
# Manages virtual network in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualNetwork configuration
resource "xcsh_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"
}
```
