---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1160, "body_sha256": "sha256:a13bac82f1dc7d3464e79042c68fb6dce7898719e93b3b1e2f3e3121328e989e", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:de754ad7eba55318affca16381d010b7282e358ea583d9fad1b36b6208539006", "source_path": "examples/resources/xcsh_enhanced_firewall_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:enhanced_firewall_policy:example:resource", "parent_id": "xcsh-docs:resources:enhanced_firewall_policy:examples", "path": "documentation/resources/enhanced_firewall_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1110202103311302-3120003312311211-0300112112111000-2313110111313033-2330220202320211-2232211200331203-2121322111231302-3032332113231121", "registry_path": "docs/guides/resources--enhanced_firewall_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Resource for xcsh_enhanced_firewall_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_enhanced_firewall_policy/resource.tf`; digest `sha256:de754ad7eba55318affca16381d010b7282e358ea583d9fad1b36b6208539006`.

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
