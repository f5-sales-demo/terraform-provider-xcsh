---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cloud_user_account."
xcsh_docs: {"aliases": [], "body_bytes": 1050, "body_sha256": "sha256:34c490b2a74b193b5ce9078ae5124bdf0ad37c890d7b8f51b3f8bdb246e88333", "canonical_id": "xcsh-docs:resources:cloud_user_account:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:10ccdaeaddaf605a3535a274963418653193279aea9eade41b140aeacaaa4d59", "source_path": "examples/resources/xcsh_cloud_user_account/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cloud_user_account:example:resource", "parent_id": "xcsh-docs:resources:cloud_user_account:examples", "path": "docs/guides/resources--cloud_user_account--example--resource.md", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_cloud_user_account.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md)
- [Examples](resources--cloud_user_account--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_user_account/resource.tf`; digest `sha256:10ccdaeaddaf605a3535a274963418653193279aea9eade41b140aeacaaa4d59`.

```terraform
# CloudUserAccount Resource Example
# Manages a Cloud User Account resource in F5 Distributed Cloud for cloud user account object create specifications.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudUserAccount configuration
resource "xcsh_cloud_user_account" "example" {
  name      = "example-cloud-user-account"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--cloud_user_account--examples.md)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md)
