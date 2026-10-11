---
page_title: "xcsh_device_intelligence_risk_score_distribution"
subcategory: ""
description: "Reads Device Intelligence Risk Score Distribution information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["device intelligence risk score distribution"], "body_bytes": 1513, "body_sha256": "sha256:b880b45670542a93a1cae0088dd99ee133299dab724ee16b7889efd939bb8cca", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_risk_score_distribution:reference", "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/device_intelligence_risk_score_distribution/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_risk_score_distribution", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-0220221212202100-3300002033312230-0123021230311320-1120110032303212-1000013303201203-1310101222000102-0330103220003033-0002223100110032", "registry_path": "docs/data-sources/device_intelligence_risk_score_distribution.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_risk_score_distribution/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reads Device Intelligence Risk Score Distribution information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_device_intelligence_risk_score_distribution

Breadcrumbs:

- xcsh_device_intelligence_risk_score_distribution

Reads Device Intelligence Risk Score Distribution information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceRiskScoreDistribution DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_risk_score_distribution" "example" {
  namespace = "example-value"
}

output "device_intelligence_risk_score_distribution_result" {
  value = data.xcsh_device_intelligence_risk_score_distribution.example
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_risk_score_distribution/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_risk_score_distribution/examples/)
