---
page_title: "xcsh_service_policy_rule"
subcategory: ""
description: "xcsh_service_policy_rule for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1351, "body_sha256": "sha256:5b7427b01a609c1159150ff06dcadc2f48b2ef5572769e3ae6b9d2df8f9dd2ea", "canonical_id": "xcsh-docs:data-sources:service_policy_rule:fundamentals", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:reference", "xcsh-docs:data-sources:service_policy_rule:examples"], "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:fundamentals", "parent_id": null, "path": "docs/data-sources/service_policy_rule.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_service_policy_rule for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Property reference](../guides/data-sources--service_policy_rule--reference.md)
- [Examples](../guides/data-sources--service_policy_rule--examples.md)
