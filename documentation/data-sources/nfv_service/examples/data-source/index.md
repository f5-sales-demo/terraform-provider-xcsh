---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1167, "body_sha256": "sha256:e385d1e9f61878ca7311214c661e192e96b4f80b4c4d26e4d3270a77b1810d09", "child_ids": [], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e1bcf01b10f5271939ad8fda84b9d63610e3af5498aba69a8e2b34b2c60983a2", "source_path": "examples/data-sources/xcsh_nfv_service/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:nfv_service:example:data-source", "parent_id": "xcsh-docs:data-sources:nfv_service:examples", "path": "documentation/data-sources/nfv_service/examples/data-source/index.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nfv_service/data-source.tf`; digest `sha256:e1bcf01b10f5271939ad8fda84b9d63610e3af5498aba69a8e2b34b2c60983a2`.

```terraform
# NfvService Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NfvService by name
data "xcsh_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "staging"
}

output "nfv_service_id" {
  value = data.xcsh_nfv_service.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/examples/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
