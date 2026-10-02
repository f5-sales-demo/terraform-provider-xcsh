---
page_title: "xcsh_enhanced_firewall_policy"
subcategory: ""
description: "Manages a Enhanced Firewall Policy resource in F5 Distributed Cloud for enhanced firewall policy specification. configuration."
xcsh_docs: {"aliases": ["enhanced firewall policy"], "body_bytes": 1740, "body_sha256": "sha256:bb9fbdf5f42abe4612077b3502e0059ef44d2f76b69b09bb706a2cf3dac0fcba", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:enhanced_firewall_policy:reference", "xcsh-docs:resources:enhanced_firewall_policy:examples", "xcsh-docs:resources:enhanced_firewall_policy:import", "xcsh-docs:resources:enhanced_firewall_policy:timeouts"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:enhanced_firewall_policy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/enhanced_firewall_policy/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120", "registry_path": "docs/resources/enhanced_firewall_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Enhanced Firewall Policy resource in F5 Distributed Cloud for enhanced firewall policy specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/lifecycle/timeouts/)
