---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1418, "body_sha256": "sha256:bd4c87d3a2189e06294231f3cce7ecf314784536d495bc43b937735c59277e65", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:de754ad7eba55318affca16381d010b7282e358ea583d9fad1b36b6208539006", "source_path": "examples/resources/xcsh_enhanced_firewall_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:enhanced_firewall_policy:example:resource", "parent_id": "xcsh-docs:resources:enhanced_firewall_policy:examples", "path": "documentation/resources/enhanced_firewall_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1110202103311302-3120003312311211-0300112112111000-2313110111313033-2330220202320211-2232211200331203-2121322111231302-3032332113231121", "registry_path": "docs/guides/resources--enhanced_firewall_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_enhanced_firewall_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/examples/)
- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
