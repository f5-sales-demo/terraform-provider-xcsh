---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_addon_service."
xcsh_docs: {"aliases": [], "body_bytes": 1193, "body_sha256": "sha256:5968805cd5e5c788a6f1cc0c04f57647c27015b971707e8b09f7cb9b4f220d11", "child_ids": [], "collection_id": "xcsh-docs:data-sources:addon_service:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:809191858b0cd3ac6042ff132999dd4d2b0f0773a32efcde70b3a2eff7c8fa98", "source_path": "examples/data-sources/xcsh_addon_service/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:addon_service:example:data-source", "parent_id": "xcsh-docs:data-sources:addon_service:examples", "path": "documentation/data-sources/addon_service/examples/data-source/index.md", "provider_name": "addon_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/addon_service/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_addon_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_addon_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_addon_service/data-source.tf`; digest `sha256:809191858b0cd3ac6042ff132999dd4d2b0f0773a32efcde70b3a2eff7c8fa98`.

```terraform
# AddonService Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AddonService by name
data "xcsh_addon_service" "example" {
  name      = "example-addon-service"
  namespace = "staging"
}

output "addon_service_id" {
  value = data.xcsh_addon_service.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/examples/)
- [xcsh_addon_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/)
