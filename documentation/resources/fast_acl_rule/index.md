---
page_title: "xcsh_fast_acl_rule"
subcategory: ""
description: "xcsh_fast_acl_rule for xcsh_fast_acl_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1635, "body_sha256": "sha256:14918ccd5f4194c6e4b5114c2c5df55a8d95b5ed630f3404f3ad386f5e10b98d", "child_ids": ["xcsh-docs:resources:fast_acl_rule:reference", "xcsh-docs:resources:fast_acl_rule:examples", "xcsh-docs:resources:fast_acl_rule:import", "xcsh-docs:resources:fast_acl_rule:timeouts"], "collection_id": "xcsh-docs:resources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl_rule:fundamentals", "parent_id": null, "path": "documentation/resources/fast_acl_rule/index.md", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl_rule/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_fast_acl_rule for xcsh_fast_acl_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
