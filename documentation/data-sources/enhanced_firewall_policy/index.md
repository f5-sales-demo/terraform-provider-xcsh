---
page_title: "xcsh_enhanced_firewall_policy"
subcategory: ""
description: "Reads Enhanced Firewall Policy information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["enhanced firewall policy"], "body_bytes": 1447, "body_sha256": "sha256:4289a772f666b0e2c2156683f227d4542b7cae94b47575c65c2f59e15da31485", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:enhanced_firewall_policy:reference", "xcsh-docs:data-sources:enhanced_firewall_policy:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:enhanced_firewall_policy:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/enhanced_firewall_policy/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200", "registry_path": "docs/data-sources/enhanced_firewall_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/enhanced_firewall_policy/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reads Enhanced Firewall Policy information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_enhanced_firewall_policy

Breadcrumbs:

- xcsh_enhanced_firewall_policy

Reads Enhanced Firewall Policy information from F5 Distributed Cloud.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/examples/)
