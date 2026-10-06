---
page_title: "xcsh_device_intelligence_high_risk_transactions"
subcategory: ""
description: "Reads Device Intelligence High Risk Transactions information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["device intelligence high risk transactions"], "body_bytes": 1504, "body_sha256": "sha256:e1e40b3bc48a7117429d3adc100569a66694162b87d161e1f1d26c7402a2d158", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_high_risk_transactions:reference", "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/device_intelligence_high_risk_transactions/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_high_risk_transactions", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2323323100323133-2031230011120020-3320233202222001-1321322211033211-2221211100303331-0220103011323320-1220023022213033-0133011203321103", "registry_path": "docs/data-sources/device_intelligence_high_risk_transactions.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_high_risk_transactions/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reads Device Intelligence High Risk Transactions information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_device_intelligence_high_risk_transactions

Breadcrumbs:

- xcsh_device_intelligence_high_risk_transactions

Reads Device Intelligence High Risk Transactions information from F5 Distributed Cloud.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/examples/)
