---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_network_firewall."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1107, "body_sha256": "sha256:cd4c8d860bbb4f52f125dbd165d416331e1f6ad8d1bb8dc198a7759b609338c8", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7eeae23d1d7b0c0adbae62b0e59aafa91676ffc023ed6aa05396e73c69627d84", "source_path": "examples/resources/xcsh_network_firewall/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:network_firewall:example:resource", "parent_id": "xcsh-docs:resources:network_firewall:examples", "path": "documentation/resources/network_firewall/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3131002303311013-0301222221021220-1112323130123222-3021021322310303-3202210313132233-3033332133000102-2203100030222302-0100331332220220", "registry_path": "docs/guides/resources--network_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/examples/resource/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Resource for xcsh_network_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["network_firewallCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_network_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_firewall/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_firewall/resource.tf`; digest `sha256:7eeae23d1d7b0c0adbae62b0e59aafa91676ffc023ed6aa05396e73c69627d84`.

```terraform
# NetworkFirewall Resource Example
# Manages a Network Firewall resource in F5 Distributed Cloud for network firewall is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkFirewall configuration
resource "xcsh_network_firewall" "example" {
  name      = "example-network-firewall"
  namespace = "system"
}
```
