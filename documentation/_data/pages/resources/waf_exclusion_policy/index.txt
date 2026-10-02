---
page_title: "xcsh_waf_exclusion_policy"
subcategory: ""
description: "Manages WAF exclusion policy in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["waf exclusion policy"], "body_bytes": 1569, "body_sha256": "sha256:55a1b3d4f77224e9c39758d910482ec467da5dfe1beb772673c890e7777cb3a3", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:waf_exclusion_policy:reference", "xcsh-docs:resources:waf_exclusion_policy:examples", "xcsh-docs:resources:waf_exclusion_policy:import", "xcsh-docs:resources:waf_exclusion_policy:timeouts"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:waf_exclusion_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:waf_exclusion_policy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/waf_exclusion_policy/index.md", "product": "distributed-cloud", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3330030113020300-2333011200112232-2103211201021223-3323101232201322-0333023102000133-0023211202130100-1022130022032200-1213203320332021", "registry_path": "docs/resources/waf_exclusion_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/waf_exclusion_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages WAF exclusion policy in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_waf_exclusion_policy

Breadcrumbs:

- xcsh_waf_exclusion_policy

Manages WAF exclusion policy in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFExclusionPolicy Resource Example
# Manages WAF exclusion policy in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic WAFExclusionPolicy configuration
resource "xcsh_waf_exclusion_policy" "example" {
  name      = "example-waf-exclusion-policy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/waf_exclusion_policy/lifecycle/timeouts/)
