---
page_title: "xcsh_enhanced_firewall_policy"
subcategory: ""
description: "Manages a Enhanced Firewall Policy resource in F5 Distributed Cloud for enhanced firewall policy specification. configuration."
xcsh_docs: {"aliases": ["enhanced firewall policy"], "body_bytes": 1491, "body_sha256": "sha256:7915b6ee3bb8c006d062c84591d511bfa8a50c9045bc70228b94eac4c777c09a", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:enhanced_firewall_policy:reference", "xcsh-docs:data-sources:enhanced_firewall_policy:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:enhanced_firewall_policy:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/enhanced_firewall_policy/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200", "registry_path": "docs/data-sources/enhanced_firewall_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/enhanced_firewall_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Enhanced Firewall Policy resource in F5 Distributed Cloud for enhanced firewall policy specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
# EnhancedFirewallPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing EnhancedFirewallPolicy by name
data "xcsh_enhanced_firewall_policy" "example" {
  name      = "example-enhanced-firewall-policy"
  namespace = "staging"
}

output "enhanced_firewall_policy_id" {
  value = data.xcsh_enhanced_firewall_policy.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/examples/)
