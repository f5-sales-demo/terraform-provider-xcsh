---
page_title: "xcsh_token"
subcategory: "Identity"
description: "xcsh_token for xcsh_token."
xcsh_docs: {"aliases": [], "body_bytes": 1599, "body_sha256": "sha256:3d357efa612560828fbceaa4566b65a5c17391cd92e32eb6e184154be5d59e9f", "child_ids": ["xcsh-docs:resources:token:reference", "xcsh-docs:resources:token:examples", "xcsh-docs:resources:token:import", "xcsh-docs:resources:token:timeouts"], "collection_id": "xcsh-docs:resources:token:collection", "completeness": "complete", "id": "xcsh-docs:resources:token:fundamentals", "parent_id": null, "path": "documentation/resources/token/index.md", "provider_name": "token", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_token for xcsh_token.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tokenCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_token

Breadcrumbs:

- xcsh_token

Manages new token. Token object is used to manage site admission. User must generate token before
provisioning and pass this token to site during it's registration in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Token Resource Example
# Manages new token.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Token configuration
resource "xcsh_token" "example" {
  name      = "example-token"
  namespace = "system"
  type      = 1
  site_name = "example-securemesh-site"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/lifecycle/timeouts/)
