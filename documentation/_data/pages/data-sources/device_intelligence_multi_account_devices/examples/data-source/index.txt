---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_device_intelligence_multi_account_devices."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1523, "body_sha256": "sha256:296de72305fd354e96b941eecd275e98c79306d0f60b82cd3bb5b5127a50da79", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:83db70480833f20bc19383610323bcf199f0707dd687a03af490d0098255cfcb", "source_path": "examples/data-sources/xcsh_device_intelligence_multi_account_devices/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:example:data-source", "parent_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:examples", "path": "documentation/data-sources/device_intelligence_multi_account_devices/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_multi_account_devices", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1300111002212133-0100000022222012-3220100313023223-2121110011022211-1202021322301122-2203213212102321-1120311111330033-0122320021123312", "registry_path": "docs/guides/data-sources--device_intelligence_multi_account_devices--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_multi_account_devices/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_device_intelligence_multi_account_devices.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/examples/)
- [xcsh_device_intelligence_multi_account_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/)
