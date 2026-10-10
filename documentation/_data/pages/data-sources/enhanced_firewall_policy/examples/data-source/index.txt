---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1169, "body_sha256": "sha256:749884f8e485bd938c41dea23a3dbe077f826fb42b95bda7d5db902364255106", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:enhanced_firewall_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4fad9180d1e11ef8fde0a09a6c59611ed259b2e0dc9e6684bd5a7a00db211172", "source_path": "examples/data-sources/xcsh_enhanced_firewall_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:enhanced_firewall_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:enhanced_firewall_policy:examples", "path": "documentation/data-sources/enhanced_firewall_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3212301120123002-2120002222312122-0030221131010231-2202311220223233-2003202123012223-3303322102000200-3203311323100031-2131223212103132", "registry_path": "docs/guides/data-sources--enhanced_firewall_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/enhanced_firewall_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_enhanced_firewall_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_enhanced_firewall_policy/data-source.tf`; digest `sha256:4fad9180d1e11ef8fde0a09a6c59611ed259b2e0dc9e6684bd5a7a00db211172`.

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
