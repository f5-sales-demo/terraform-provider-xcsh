---
page_title: "xcsh_api_testing"
subcategory: ""
description: "xcsh_api_testing for xcsh_api_testing."
xcsh_docs: {"aliases": [], "body_bytes": 1197, "body_sha256": "sha256:2e2f8b3bbda0faf6eac7a27bbcf72e7e76900073c51bd5f7e04cc7354bbdcb2d", "canonical_id": "xcsh-docs:resources:api_testing:fundamentals", "child_ids": ["xcsh-docs:resources:api_testing:reference", "xcsh-docs:resources:api_testing:examples", "xcsh-docs:resources:api_testing:import", "xcsh-docs:resources:api_testing:timeouts"], "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_testing:fundamentals", "parent_id": null, "path": "docs/resources/api_testing.md", "provider_name": "api_testing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_api_testing for xcsh_api_testing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_api_testing

Breadcrumbs:

- xcsh_api_testing

Manages a API Testing resource in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APITesting Resource Example
# Manages a API Testing resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APITesting configuration
resource "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--api_testing--reference.md)
- [Examples](../guides/resources--api_testing--examples.md)
- [Import](../guides/resources--api_testing--import.md)
- [Timeouts](../guides/resources--api_testing--timeouts.md)
