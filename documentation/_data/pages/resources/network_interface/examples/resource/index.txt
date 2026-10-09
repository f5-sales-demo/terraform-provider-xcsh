---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_network_interface."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1123, "body_sha256": "sha256:3ff3fea514459c1670c0a84e2eef81ef15fe657f0a51c76090ae9f849ba88c82", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7844384dc87eb9b41829fe1bcb33636a61a1469e9b6c183c89c2bb605fd224c5", "source_path": "examples/resources/xcsh_network_interface/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:network_interface:example:resource", "parent_id": "xcsh-docs:resources:network_interface:examples", "path": "documentation/resources/network_interface/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2003300332031011-0310201123023101-2001130020212130-0133100130023231-0122000312012030-0213003300322000-3313221322331232-0303200100032021", "registry_path": "docs/guides/resources--network_interface--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/examples/resource/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Resource for xcsh_network_interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_interface/resource.tf`; digest `sha256:7844384dc87eb9b41829fe1bcb33636a61a1469e9b6c183c89c2bb605fd224c5`.

```terraform
# NetworkInterface Resource Example
# Manages a Network Interface resource in F5 Distributed Cloud for network interface represents configuration of a network device.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkInterface configuration
resource "xcsh_network_interface" "example" {
  name      = "example-network-interface"
  namespace = "staging"
}
```
