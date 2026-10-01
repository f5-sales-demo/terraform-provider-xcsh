---
page_title: "xcsh_device_intelligence_risk_score_distribution"
subcategory: ""
description: "xcsh_device_intelligence_risk_score_distribution for xcsh_device_intelligence_risk_score_distribution."
xcsh_docs: {"aliases": [], "body_bytes": 1355, "body_sha256": "sha256:6209a6ac4afe15b9dc1f147f93a7e2584e699f03d0d45955008832e863b9df75", "canonical_id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:fundamentals", "child_ids": ["xcsh-docs:data-sources:device_intelligence_risk_score_distribution:reference", "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:examples"], "collection_id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:fundamentals", "parent_id": null, "path": "docs/data-sources/device_intelligence_risk_score_distribution.md", "provider_name": "device_intelligence_risk_score_distribution", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_risk_score_distribution/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_device_intelligence_risk_score_distribution for xcsh_device_intelligence_risk_score_distribution.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_device_intelligence_risk_score_distribution

Breadcrumbs:

- xcsh_device_intelligence_risk_score_distribution

Resource creation operation.

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

## Next pages

- [Property reference](../guides/data-sources--device_intelligence_risk_score_distribution--reference.md)
- [Examples](../guides/data-sources--device_intelligence_risk_score_distribution--examples.md)
