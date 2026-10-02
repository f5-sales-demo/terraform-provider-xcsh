---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_device_intelligence_device_history."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1477, "body_sha256": "sha256:abe23d958059d40a1e925b9e90bbf1c7f2c501b6a87a1b521ac3bc1542f5d29d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_device_history:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f95f8328fecc18ee316fb7732f60f6fec73cc07388f5c6067f1bad91b8ac0d75", "source_path": "examples/data-sources/xcsh_device_intelligence_device_history/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:device_intelligence_device_history:example:data-source", "parent_id": "xcsh-docs:data-sources:device_intelligence_device_history:examples", "path": "documentation/data-sources/device_intelligence_device_history/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_device_history", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1130232223300130-2002111313021222-2201321310031221-0112313123310100-3221002130323312-3203003300110200-0003033003323222-3313202022112122", "registry_path": "docs/guides/data-sources--device_intelligence_device_history--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_history/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_device_intelligence_device_history.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_device_intelligence_device_history](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_device_history/data-source.tf`; digest `sha256:f95f8328fecc18ee316fb7732f60f6fec73cc07388f5c6067f1bad91b8ac0d75`.

```terraform
# DeviceIntelligenceDeviceHistory DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_device_history" "example" {
  device_id = "example-value"
  namespace = "example-value"
}

output "device_intelligence_device_history_result" {
  value = data.xcsh_device_intelligence_device_history.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/examples/)
- [xcsh_device_intelligence_device_history](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/)
