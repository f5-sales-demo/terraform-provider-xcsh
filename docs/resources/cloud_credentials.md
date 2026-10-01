---
page_title: "xcsh_cloud_credentials"
subcategory: "Infrastructure"
description: "xcsh_cloud_credentials for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 1504, "body_sha256": "sha256:01264cc665d0cbb81e3a73e809c040077c49850ff892f578bf1aee4bc4f1c6ed", "canonical_id": "xcsh-docs:resources:cloud_credentials:fundamentals", "child_ids": ["xcsh-docs:resources:cloud_credentials:reference", "xcsh-docs:resources:cloud_credentials:examples", "xcsh-docs:resources:cloud_credentials:import", "xcsh-docs:resources:cloud_credentials:timeouts"], "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:fundamentals", "parent_id": null, "path": "docs/resources/cloud_credentials.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cloud_credentials for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
# CloudCredentials Resource Example
# Manages a Cloud Credentials resource in F5 Distributed Cloud for api to create cloud_credentials object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudCredentials configuration
resource "xcsh_cloud_credentials" "example" {
  name      = "example-cloud-credentials"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--cloud_credentials--reference.md)
- [Examples](../guides/resources--cloud_credentials--examples.md)
- [Import](../guides/resources--cloud_credentials--import.md)
- [Timeouts](../guides/resources--cloud_credentials--timeouts.md)
