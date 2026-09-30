---
page_title: "Data source"
subcategory: "Infrastructure"
description: "Data source for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 1245, "body_sha256": "sha256:8d4c2ff6b9b582fc7b79ed808f13596936e6c4ea377844285fbfaf74af88d99e", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6a62066c747f97e4baff52829aa9bc36ad62c8b9304220fe45b6e8d9ef9839f1", "source_path": "examples/data-sources/xcsh_cloud_credentials/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cloud_credentials:example:data-source", "parent_id": "xcsh-docs:data-sources:cloud_credentials:examples", "path": "documentation/data-sources/cloud_credentials/examples/data-source/index.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_credentials/data-source.tf`; digest `sha256:6a62066c747f97e4baff52829aa9bc36ad62c8b9304220fe45b6e8d9ef9839f1`.

```terraform
# CloudCredentials Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudCredentials by name
data "xcsh_cloud_credentials" "example" {
  name      = "example-cloud-credentials"
  namespace = "staging"
}

output "cloud_credentials_id" {
  value = data.xcsh_cloud_credentials.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/examples/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_credentials/)
