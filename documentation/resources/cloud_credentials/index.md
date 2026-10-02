---
page_title: "xcsh_cloud_credentials"
subcategory: "Infrastructure"
description: "Manages a Cloud Credentials resource in F5 Distributed Cloud for api to create cloud_credentials object. configuration."
xcsh_docs: {"aliases": ["authentication", "cloud credentials", "credential setup", "credentials"], "body_bytes": 1693, "body_sha256": "sha256:78d0464781de149bbdb317d2e9adea191f72e514a2dfe7bcdf425237f1511cfa", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_credentials:reference", "xcsh-docs:resources:cloud_credentials:examples", "xcsh-docs:resources:cloud_credentials:import", "xcsh-docs:resources:cloud_credentials:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_credentials:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/cloud_credentials/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0130231023331222-3013013110330021-2003033121003000-3023101330100030-1121122101200120-3011301230122133-2001010122021211-1302313010320120", "registry_path": "docs/resources/cloud_credentials.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Cloud Credentials resource in F5 Distributed Cloud for api to create cloud_credentials object. configuration.", "tasks": ["authentication", "configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/lifecycle/timeouts/)
