---
page_title: "xcsh_irule"
subcategory: ""
description: "xcsh_irule for xcsh_irule."
xcsh_docs: {"aliases": [], "body_bytes": 1184, "body_sha256": "sha256:317cbc0dfe30b9d6819a0f1dc046a8a32a9d0aa844c16bd36446ef120e79b7dc", "child_ids": ["xcsh-docs:data-sources:irule:reference", "xcsh-docs:data-sources:irule:examples"], "collection_id": "xcsh-docs:data-sources:irule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:irule:fundamentals", "parent_id": null, "path": "documentation/data-sources/irule/index.md", "provider_name": "irule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/irule/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_irule for xcsh_irule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["iruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_irule

Breadcrumbs:

- xcsh_irule

Manages iRule in a given namespace. If one already exists it will give an error in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Irule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Irule by name
data "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"
}

output "irule_id" {
  value = data.xcsh_irule.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/irule/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/irule/examples/)
