---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_device_intelligence_high_risk_transactions."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1534, "body_sha256": "sha256:e2f7bf2ca4be8cd78eadbc0e68624e0285bdfdbb779a231d4e9fb71929d552dc", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5f7ef7da7ff2c9994ec1cbab645c6c6f2e697ac5efc56e49ae4d09a5e717abc6", "source_path": "examples/data-sources/xcsh_device_intelligence_high_risk_transactions/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:example:data-source", "parent_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:examples", "path": "documentation/data-sources/device_intelligence_high_risk_transactions/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_high_risk_transactions", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2100102113233300-3021213113221123-0102013001111213-1133302320301201-0103131300021130-2310120022222123-2111031203210122-2223202131030332", "registry_path": "docs/guides/data-sources--device_intelligence_high_risk_transactions--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_high_risk_transactions/examples/data-source/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Data source for xcsh_device_intelligence_high_risk_transactions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_high_risk_transactions/data-source.tf`; digest `sha256:5f7ef7da7ff2c9994ec1cbab645c6c6f2e697ac5efc56e49ae4d09a5e717abc6`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/examples/)
- [xcsh_device_intelligence_high_risk_transactions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_high_risk_transactions/)
