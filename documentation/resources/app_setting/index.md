---
page_title: "xcsh_app_setting"
subcategory: ""
description: "Manages App setting configuration in namespace metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["app setting"], "body_bytes": 1555, "body_sha256": "sha256:ff89e767f448fac682521c54371aff238f1f992f22baa7c89f49e382e38089bc", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:app_setting:reference", "xcsh-docs:resources:app_setting:examples", "xcsh-docs:resources:app_setting:import", "xcsh-docs:resources:app_setting:timeouts"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/app_setting/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110", "registry_path": "docs/resources/app_setting.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Manages App setting configuration in namespace metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["app_settingCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_app_setting

Breadcrumbs:

- xcsh_app_setting

Manages App setting configuration in namespace metadata.namespace in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppSetting Resource Example
# Manages App setting configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppSetting configuration
resource "xcsh_app_setting" "example" {
  name      = "example-app-setting"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/lifecycle/timeouts/)
