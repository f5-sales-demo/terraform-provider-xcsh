---
page_title: "xcsh_secret_management_access"
subcategory: ""
description: "xcsh_secret_management_access for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 1698, "body_sha256": "sha256:fa2b238f74f4b70c36b83de13422114c0896c2d6e70872d3a365d6c520eb56c1", "child_ids": ["xcsh-docs:resources:secret_management_access:reference", "xcsh-docs:resources:secret_management_access:examples", "xcsh-docs:resources:secret_management_access:import", "xcsh-docs:resources:secret_management_access:timeouts"], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:fundamentals", "parent_id": null, "path": "documentation/resources/secret_management_access/index.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_secret_management_access for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_secret_management_access

Breadcrumbs:

- xcsh_secret_management_access

Manages secret\_management\_access creates a new object in storage backend for metadata.namespace in
F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecretManagementAccess Resource Example
# Manages secret_management_access creates a new object in storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecretManagementAccess configuration
resource "xcsh_secret_management_access" "example" {
  name      = "example-secret-management-access"
  namespace = "staging"

  provider_name = "example-value"
}
```

## Root configuration

Required root properties: `name`, `namespace`, `provider_name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/lifecycle/timeouts/)
