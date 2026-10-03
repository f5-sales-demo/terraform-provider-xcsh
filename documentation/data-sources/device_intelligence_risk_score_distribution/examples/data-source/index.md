---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_device_intelligence_risk_score_distribution."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1545, "body_sha256": "sha256:a7d4606a71e1a8d5d5cd57b3de29d1b18316dd8c5bc58155f32bc4ef55b8d43f", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c10972aeff2a32c44eb299757977d3dc58259ac93f292c8fdb4d0bb8272ceb96", "source_path": "examples/data-sources/xcsh_device_intelligence_risk_score_distribution/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:example:data-source", "parent_id": "xcsh-docs:data-sources:device_intelligence_risk_score_distribution:examples", "path": "documentation/data-sources/device_intelligence_risk_score_distribution/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_risk_score_distribution", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2201101332003111-2103021123300311-3303003233302120-3023332312221122-0232231133000323-2312023311212200-1310301330220333-0131023003231110", "registry_path": "docs/guides/data-sources--device_intelligence_risk_score_distribution--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_risk_score_distribution/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_device_intelligence_risk_score_distribution.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_device_intelligence_risk_score_distribution](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_risk_score_distribution/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_risk_score_distribution/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_risk_score_distribution/data-source.tf`; digest `sha256:c10972aeff2a32c44eb299757977d3dc58259ac93f292c8fdb4d0bb8272ceb96`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_risk_score_distribution/examples/)
- [xcsh_device_intelligence_risk_score_distribution](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_risk_score_distribution/)
