---
page_title: "xcsh_cloud_credentials"
subcategory: "Infrastructure"
description: "xcsh_cloud_credentials for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 1267, "body_sha256": "sha256:5cca702c3305461804b258ec1ce0f63d82e3dc619124dfd7b84988aa58d1e1fb", "canonical_id": "xcsh-docs:data-sources:cloud_credentials:fundamentals", "child_ids": ["xcsh-docs:data-sources:cloud_credentials:reference", "xcsh-docs:data-sources:cloud_credentials:examples"], "collection_id": "xcsh-docs:data-sources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_credentials:fundamentals", "parent_id": null, "path": "docs/data-sources/cloud_credentials.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_credentials/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cloud_credentials for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_cloud_credentials

Breadcrumbs:

- xcsh_cloud_credentials

Manages a Cloud Credentials resource in F5 Distributed Cloud for api to create cloud\_credentials
object. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--cloud_credentials--reference.md)
- [Examples](../guides/data-sources--cloud_credentials--examples.md)
