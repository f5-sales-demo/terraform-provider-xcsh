---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_device_intelligence_device_summary."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1183, "body_sha256": "sha256:38ce8a638a231f3d642715c36c48fc5b6bfde22a1d6dcd6ff5b238e86bad3cca", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_device_summary:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:04010440b0365c5df9d1e3244c5a3f2f792550e1e47b47ef7b1ff17bae9b7077", "source_path": "examples/data-sources/xcsh_device_intelligence_device_summary/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:device_intelligence_device_summary:example:data-source", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_summary:examples", "path": "documentation/data-sources/device_intelligence_device_summary/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_device_summary", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-1233033221000300-2123033110213020-3220330330333213-3322202120233011-2233320031300311-2110222320011022-1223313122230133-0220223033032323", "registry_path": "docs/guides/data-sources--device_intelligence_device_summary--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_summary/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_device_intelligence_device_summary.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_device_intelligence_device_summary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_device_summary/data-source.tf`; digest `sha256:04010440b0365c5df9d1e3244c5a3f2f792550e1e47b47ef7b1ff17bae9b7077`.

```terraform
# DeviceIntelligenceDeviceSummary DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_device_summary" "example" {
  device_id = "example-value"
  namespace = "example-value"
}

output "device_intelligence_device_summary_result" {
  value = data.xcsh_device_intelligence_device_summary.example
}
```
