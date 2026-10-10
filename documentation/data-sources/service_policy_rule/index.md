---
page_title: "xcsh_service_policy_rule"
subcategory: ""
description: "Reads Service Policy Rule information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["service policy rule"], "body_bytes": 1392, "body_sha256": "sha256:6bc941ba07232518bc64499d74b51f1f2311c3aa0eaef2c322f88f26ae88c50a", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:reference", "xcsh-docs:data-sources:service_policy_rule:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/service_policy_rule/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133", "registry_path": "docs/data-sources/service_policy_rule.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reads Service Policy Rule information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_service_policy_rule

Breadcrumbs:

- xcsh_service_policy_rule

Reads Service Policy Rule information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ServicePolicyRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicyRule by name
data "xcsh_service_policy_rule" "example" {
  name      = "example-service-policy-rule"
  namespace = "staging"
}

output "service_policy_rule_id" {
  value = data.xcsh_service_policy_rule.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/examples/)
