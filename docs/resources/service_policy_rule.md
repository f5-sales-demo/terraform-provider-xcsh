---
page_title: "xcsh_service_policy_rule"
subcategory: ""
description: "xcsh_service_policy_rule for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1434, "body_sha256": "sha256:14852497e8bfd7e5bdab22f6f03cfbb74309817bada0af085274f69669836d16", "canonical_id": "xcsh-docs:resources:service_policy_rule:fundamentals", "child_ids": ["xcsh-docs:resources:service_policy_rule:reference", "xcsh-docs:resources:service_policy_rule:examples", "xcsh-docs:resources:service_policy_rule:import", "xcsh-docs:resources:service_policy_rule:timeouts"], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:fundamentals", "parent_id": null, "path": "docs/resources/service_policy_rule.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_service_policy_rule for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_service_policy_rule

Breadcrumbs:

- xcsh_service_policy_rule

Manages service\_policy\_rule creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ServicePolicyRule Resource Example
# Manages service_policy_rule creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ServicePolicyRule configuration
resource "xcsh_service_policy_rule" "example" {
  name      = "example-service-policy-rule"
  namespace = "staging"

  action = "DENY"
}
```

## Root configuration

Required root properties: `action`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--service_policy_rule--reference.md)
- [Examples](../guides/resources--service_policy_rule--examples.md)
- [Import](../guides/resources--service_policy_rule--import.md)
- [Timeouts](../guides/resources--service_policy_rule--timeouts.md)
