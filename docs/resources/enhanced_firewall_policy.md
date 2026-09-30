---
page_title: "xcsh_enhanced_firewall_policy"
subcategory: ""
description: "xcsh_enhanced_firewall_policy for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1452, "body_sha256": "sha256:8d6780f2cebadaa350b74bd4075b000f2ef9971ad96524a96959005c8ee9a341", "canonical_id": "xcsh-docs:resources:enhanced_firewall_policy:fundamentals", "child_ids": ["xcsh-docs:resources:enhanced_firewall_policy:reference", "xcsh-docs:resources:enhanced_firewall_policy:examples", "xcsh-docs:resources:enhanced_firewall_policy:import", "xcsh-docs:resources:enhanced_firewall_policy:timeouts"], "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:enhanced_firewall_policy:fundamentals", "parent_id": null, "path": "docs/resources/enhanced_firewall_policy.md", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_enhanced_firewall_policy for xcsh_enhanced_firewall_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_enhanced_firewall_policy

Breadcrumbs:

- xcsh_enhanced_firewall_policy

Manages a Enhanced Firewall Policy resource in F5 Distributed Cloud for enhanced firewall policy
specification. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# EnhancedFirewallPolicy Resource Example
# Manages a Enhanced Firewall Policy resource in F5 Distributed Cloud for enhanced firewall policy specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic EnhancedFirewallPolicy configuration
resource "xcsh_enhanced_firewall_policy" "example" {
  name      = "example-enhanced-firewall-policy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--enhanced_firewall_policy--reference.md)
- [Examples](../guides/resources--enhanced_firewall_policy--examples.md)
- [Import](../guides/resources--enhanced_firewall_policy--import.md)
- [Timeouts](../guides/resources--enhanced_firewall_policy--timeouts.md)
