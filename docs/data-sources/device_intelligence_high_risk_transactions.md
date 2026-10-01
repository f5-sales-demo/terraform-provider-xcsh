---
page_title: "xcsh_device_intelligence_high_risk_transactions"
subcategory: ""
description: "xcsh_device_intelligence_high_risk_transactions for xcsh_device_intelligence_high_risk_transactions."
xcsh_docs: {"aliases": [], "body_bytes": 1347, "body_sha256": "sha256:663bd11eab3506c5dbaee3f3474219604dbe5045b4c74e3aadbb0025bd3f1b39", "canonical_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:fundamentals", "child_ids": ["xcsh-docs:data-sources:device_intelligence_high_risk_transactions:reference", "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:examples"], "collection_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:fundamentals", "parent_id": null, "path": "docs/data-sources/device_intelligence_high_risk_transactions.md", "provider_name": "device_intelligence_high_risk_transactions", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_high_risk_transactions/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_device_intelligence_high_risk_transactions for xcsh_device_intelligence_high_risk_transactions.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_device_intelligence_high_risk_transactions

Breadcrumbs:

- xcsh_device_intelligence_high_risk_transactions

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceHighRiskTransactions DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_high_risk_transactions" "example" {
  namespace = "example-value"
}

output "device_intelligence_high_risk_transactions_result" {
  value = data.xcsh_device_intelligence_high_risk_transactions.example
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--device_intelligence_high_risk_transactions--reference.md)
- [Examples](../guides/data-sources--device_intelligence_high_risk_transactions--examples.md)
