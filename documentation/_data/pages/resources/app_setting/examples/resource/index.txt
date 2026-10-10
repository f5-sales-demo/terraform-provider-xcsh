---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_app_setting."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1037, "body_sha256": "sha256:39772bb1972d243fa6f445542e49cfad80872b8bc49ad3dc57c16df85c5cd753", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d31cbf807fc9dcf209556e05af346ee32ebdbf440fd52fffb2c11e50ae1d1f1a", "source_path": "examples/resources/xcsh_app_setting/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_setting:example:resource", "parent_id": "xcsh-docs:resources:app_setting:examples", "path": "documentation/resources/app_setting/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0300030103112322-2002201203210030-1013031211001003-0301213001221003-3202330000311300-2302322102003332-2322121113311303-3022201333120203", "registry_path": "docs/guides/resources--app_setting--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_app_setting.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["app_settingCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_setting/resource.tf`; digest `sha256:d31cbf807fc9dcf209556e05af346ee32ebdbf440fd52fffb2c11e50ae1d1f1a`.

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
