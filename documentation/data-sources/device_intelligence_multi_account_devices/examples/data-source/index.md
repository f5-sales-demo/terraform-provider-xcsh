---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_device_intelligence_multi_account_devices."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1208, "body_sha256": "sha256:4043e171adef45ee00dec914cd66559a02681d904f0e02e0aca281518c323792", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:83db70480833f20bc19383610323bcf199f0707dd687a03af490d0098255cfcb", "source_path": "examples/data-sources/xcsh_device_intelligence_multi_account_devices/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:example:data-source", "parent_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:examples", "path": "documentation/data-sources/device_intelligence_multi_account_devices/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_multi_account_devices", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1300111002212133-0100000022222012-3220100313023223-2121110011022211-1202021322301122-2203213212102321-1120311111330033-0122320021123312", "registry_path": "docs/guides/data-sources--device_intelligence_multi_account_devices--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_multi_account_devices/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_device_intelligence_multi_account_devices.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_multi_account_devices/data-source.tf`; digest `sha256:83db70480833f20bc19383610323bcf199f0707dd687a03af490d0098255cfcb`.

```terraform
# DeviceIntelligenceMultiAccountDevices DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_multi_account_devices" "example" {
  namespace = "example-value"
}

output "device_intelligence_multi_account_devices_result" {
  value = data.xcsh_device_intelligence_multi_account_devices.example
}
```
