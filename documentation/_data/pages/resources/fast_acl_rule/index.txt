---
page_title: "xcsh_fast_acl_rule"
subcategory: ""
description: "Manages new Fast ACL rule, has specification to match source IP, source port and action to apply in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["fast acl rule"], "body_bytes": 1635, "body_sha256": "sha256:14918ccd5f4194c6e4b5114c2c5df55a8d95b5ed630f3404f3ad386f5e10b98d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl_rule:reference", "xcsh-docs:resources:fast_acl_rule:examples", "xcsh-docs:resources:fast_acl_rule:import", "xcsh-docs:resources:fast_acl_rule:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl_rule:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/fast_acl_rule/index.md", "product": "distributed-cloud", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302", "registry_path": "docs/resources/fast_acl_rule.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl_rule/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages new Fast ACL rule, has specification to match source IP, source port and action to apply in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_fast_acl_rule

Breadcrumbs:

- xcsh_fast_acl_rule

Manages new Fast ACL rule, has specification to match source IP, source port and action to apply in
F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FastACLRule Resource Example
# Manages new Fast ACL rule, has specification to match source IP, source port and action to apply in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FastACLRule configuration
resource "xcsh_fast_acl_rule" "example" {
  name      = "example-fast-acl-rule"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/lifecycle/timeouts/)
