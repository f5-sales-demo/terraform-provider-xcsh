---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_app_setting."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1256, "body_sha256": "sha256:9dcbc868991e81c5d5d9eafb50d2b643943326d024dd408205d1b7b49106c60b", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d31cbf807fc9dcf209556e05af346ee32ebdbf440fd52fffb2c11e50ae1d1f1a", "source_path": "examples/resources/xcsh_app_setting/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_setting:example:resource", "parent_id": "xcsh-docs:resources:app_setting:examples", "path": "documentation/resources/app_setting/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0300030103112322-2002201203210030-1013031211001003-0301213001221003-3202330000311300-2302322102003332-2322121113311303-3022201333120203", "registry_path": "docs/guides/resources--app_setting--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_app_setting.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/examples/)
- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
