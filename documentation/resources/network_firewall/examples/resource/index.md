---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_network_firewall."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1107, "body_sha256": "sha256:cd4c8d860bbb4f52f125dbd165d416331e1f6ad8d1bb8dc198a7759b609338c8", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7eeae23d1d7b0c0adbae62b0e59aafa91676ffc023ed6aa05396e73c69627d84", "source_path": "examples/resources/xcsh_network_firewall/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:network_firewall:example:resource", "parent_id": "xcsh-docs:resources:network_firewall:examples", "path": "documentation/resources/network_firewall/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "network_firewall", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3131002303311013-0301222221021220-1112323130123222-3021021322310303-3202210313132233-3033332133000102-2203100030222302-0100331332220220", "registry_path": "docs/guides/resources--network_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_firewall/examples/resource/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Resource for xcsh_network_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["network_firewallCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
