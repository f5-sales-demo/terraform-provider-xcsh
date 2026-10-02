---
page_title: "xcsh_cloud_user_account"
subcategory: ""
description: "Manages a Cloud User Account resource in F5 Distributed Cloud for cloud user account object create specifications. configuration."
xcsh_docs: {"aliases": ["cloud user account"], "body_bytes": 1686, "body_sha256": "sha256:584f2ebf1b8fa473a137297e0787d9aecff70137b3e6258dfac9714bac8c4880", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:cloud_user_account:reference", "xcsh-docs:resources:cloud_user_account:examples", "xcsh-docs:resources:cloud_user_account:import", "xcsh-docs:resources:cloud_user_account:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_user_account:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/cloud_user_account/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321", "registry_path": "docs/resources/cloud_user_account.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Cloud User Account resource in F5 Distributed Cloud for cloud user account object create specifications. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_cloud_user_account

Breadcrumbs:

- xcsh_cloud_user_account

Manages a Cloud User Account resource in F5 Distributed Cloud for cloud user account object create
specifications. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/lifecycle/timeouts/)
