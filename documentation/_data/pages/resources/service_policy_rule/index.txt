---
page_title: "xcsh_service_policy_rule"
subcategory: ""
description: "Manages service_policy_rule creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "service policy rule", "upstream servers"], "body_bytes": 1722, "body_sha256": "sha256:57ef58e16b0d676d6afd9a33f68d03db9eae1e954d47f3f31b806b980549b527", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy_rule:reference", "xcsh-docs:resources:service_policy_rule:examples", "xcsh-docs:resources:service_policy_rule:import", "xcsh-docs:resources:service_policy_rule:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/service_policy_rule/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032", "registry_path": "docs/resources/service_policy_rule.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages service_policy_rule creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/lifecycle/timeouts/)
