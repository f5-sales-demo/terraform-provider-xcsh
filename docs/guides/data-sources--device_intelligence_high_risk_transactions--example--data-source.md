---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_device_intelligence_high_risk_transactions."
xcsh_docs: {"aliases": [], "body_bytes": 1328, "body_sha256": "sha256:335b48f6084caae3abc666e30ba7e8dcd4957c8251e87e4db004b7780ecf9050", "canonical_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5f7ef7da7ff2c9994ec1cbab645c6c6f2e697ac5efc56e49ae4d09a5e717abc6", "source_path": "examples/data-sources/xcsh_device_intelligence_high_risk_transactions/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:example:data-source", "parent_id": "xcsh-docs:data-sources:device_intelligence_high_risk_transactions:examples", "path": "docs/guides/data-sources--device_intelligence_high_risk_transactions--example--data-source.md", "provider_name": "device_intelligence_high_risk_transactions", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_high_risk_transactions/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_device_intelligence_high_risk_transactions.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md)
- [Examples](data-sources--device_intelligence_high_risk_transactions--examples.md)
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

- [Examples](data-sources--device_intelligence_high_risk_transactions--examples.md)
- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md)
